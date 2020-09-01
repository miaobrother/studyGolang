package main

import (
        "github.com/aws/aws-sdk-go/aws"
        "github.com/aws/aws-sdk-go/aws/session"
        "fmt"
        "os"
        "github.com/aws/aws-sdk-go/service/s3"
        "github.com/aws/aws-sdk-go/service/s3/s3manager"
        "github.com/aws/aws-sdk-go/aws/awserr"
        //"path/filepath"
        "time"
        //"strings"
        //"io/ioutil"
        "path/filepath"
)

func exitErrorf(msg string, args ...interface{}) {
        fmt.Fprintf(os.Stderr, msg+"\n", args...)
        os.Exit(1)
}

func getListBucket() {
        sess := session.Must(session.NewSession(&aws.Config{}))

        svc := s3.New(sess, &aws.Config{Region: aws.String("cn-north-1")})
        input := &s3.ListBucketsInput{}

        result, err := svc.ListBuckets(input)
        if err != nil {
                if aerr, ok := err.(awserr.Error); ok {
                        switch aerr.Code() {
                        default:
                                fmt.Println(aerr.Error())
                        }
                } else {
                        // Print the error, cast err to awserr.Error to get the Code and
                        // Message from an error.
                        fmt.Println(err.Error())
                }
                return
        }

        fmt.Println(result)

}

func createBucket() {
        bucket := "test-backup-test"
        sess := session.Must(session.NewSession(&aws.Config{Region: aws.String("cn-north-1")}))
        svc := s3.New(sess)

        _, err := svc.CreateBucket(&s3.CreateBucketInput{
                Bucket: &bucket,
        })
        if err != nil {
                fmt.Println("Failed to create bucket", err)
                return
        }

}

func getNeedUploadData() (pathDate string) {

        //Get 当天前一天日期

        currentDate := time.Now()
        beforeDate := currentDate.AddDate(0, 0, -1).Format("20060102")
        fmt.Printf("The beforeDate is %v\n", beforeDate)
        //var LocalBackDir = "/Users/zhangyalei/" + beforeDate + "/"
        //fmt.Printf("The localbackdir is:%v\n", LocalBackDir)
        return beforeDate
        //var TestBackDir = "/data/backupdb/" + "currentDate"


        //dirs, _ := ioutil.ReadDir(LocalBackDir)
        //for _, dir := range dirs {
        //
        //      if strings.HasPrefix(dir.Name(), ".") {
        //              continue
        //      } else {
        //              fmt.Printf("The dirs name is: %v , size is: %v\n", dir.Name(), dir.Size())
        //      }
        //
        //}

}

type DirectoryIterator struct {
        filePaths []string
        bucket    string
        next struct {
                path string
                f    *os.File
        }
        err error
}

func NewDirectoryIterator(bucket, dir string) s3manager.BatchUploadIterator {
        var paths []string
        filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
                if !info.IsDir() {
                        paths = append(paths, path)
                }
                return nil
        })

        return &DirectoryIterator{
                filePaths: paths,
                bucket:    bucket,
        }
}

func (di *DirectoryIterator) Next() bool {
        if len(di.filePaths) == 0 {
                di.next.f = nil
                return false
        }

        f, err := os.Open(di.filePaths[0])
        di.err = err
        di.next.f = f
        di.next.path = di.filePaths[0]
        di.filePaths = di.filePaths[1:]

        return true && di.Err() == nil
}

// Err returns error of DirectoryIterator
func (di *DirectoryIterator) Err() error {
        return di.err
}

// UploadObject uploads a file
func (di *DirectoryIterator) UploadObject() s3manager.BatchUploadObject {
        f := di.next.f
        return s3manager.BatchUploadObject{
                Object: &s3manager.UploadInput{
                        Bucket: &di.bucket,
                        Key:    &di.next.path,
                        Body:   f,
                },
                After: func() error {
                        return f.Close()
                },
        }
}

func main() {
        //createBucket()
        //getListBucket()
        // 获取前一天日志字符串
        pathTest := getNeedUploadData()
        // bucket := "data-backup-data"
        bucket := "data-test-backup"
        sess := session.Must(session.NewSession(&aws.Config{Region: aws.String("cn-north-1")}))
        uploader := s3manager.NewUploader(sess)
        // 循环上传目录
        dirSlice := []string{"/data/backupapp/business","/data/backupdb/test"}
        // dirSlice := []string{"/data/backupdb/test"}
        for _,dir := range dirSlice{
                fmt.Println("传输当前目录:",dir+"/"+pathTest)
                di := NewDirectoryIterator(bucket, dir+ "/"+pathTest)
                // di := NewDirectoryIterator(bucket,dir)
                if err := uploader.UploadWithIterator(aws.BackgroundContext(), di); err != nil {
                        exitErrorf("failed to upload %q, %v", err)
                }

        }
        fmt.Printf("successfully uploaded %q to %q", pathTest, bucket)


}

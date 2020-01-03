package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"

	"github.com/IBM/ibm-cos-sdk-go/aws"
	"github.com/IBM/ibm-cos-sdk-go/aws/awsutil"
	"github.com/IBM/ibm-cos-sdk-go/aws/credentials/ibmiam"
	"github.com/IBM/ibm-cos-sdk-go/aws/session"
	"github.com/IBM/ibm-cos-sdk-go/service/s3"
)

const (
	apiKey            = "<API-KEY>"
	serviceInstanceID = "<SERVICE-ID>"
	authEndpoint      = "<AUTH-ENDPOINT>"
	serviceEndpoint   = "<ICOS-SERVICES-ENDPOINT>"
	icosRegion        = "<REGION>" // EJEMPLO: "us-south"
	bucketName        = "<BUCKET-NAME>"
)

func main() {

	conf := aws.NewConfig().
		WithRegion(icosRegion).
		WithEndpoint(serviceEndpoint).
		WithCredentials(ibmiam.NewStaticCredentials(aws.NewConfig(), authEndpoint, apiKey, serviceInstanceID)).
		WithS3ForcePathStyle(true)

	sess := session.Must(session.NewSession())
	client := s3.New(sess, conf)

	// Retrieve the list of available buckets
	bklist, err := client.ListBuckets(nil)
	if err != nil {
		exitErrorf("Unable to list buckets, %v", err)
	}

	fmt.Println("Buckets:")

	for _, b := range bklist.Buckets {
		fmt.Printf("* %s created on %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}

	//vamos a subir un archivo

	fmt.Printf("svc: %+v \n", client)

	file, err := os.Open("./lista_obj.go")
	if err != nil {
		fmt.Printf("err opening file: %s", err)
	}
	defer file.Close()
	fileInfo, _ := file.Stat()
	size := fileInfo.Size()
	buffer := make([]byte, size) // read file content to buffer

	file.Read(buffer)
	fileBytes := bytes.NewReader(buffer)
	fileType := http.DetectContentType(buffer)
	path := file.Name()
	params := &s3.PutObjectInput{
		Bucket:        aws.String(bucketName),
		Key:           aws.String(path),
		Body:          fileBytes,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(fileType),
	}
	resp, err := client.PutObject(params)
	if err != nil {
		fmt.Printf("bad response: %s \n", err)
	}
	fmt.Printf("response %s", awsutil.StringValue(resp))

}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/debug"

	apis3ls "github.com/neatflowcv/s3/pkg/s3ls"
)

func version() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "unknown"
	}

	return info.Main.Version
}

func main() {
	log.Println("version", version())

	endpoint := flag.String("endpoint", "http://192.168.56.176:80", "Ceph RGW endpoint (http/https)")
	bucket := flag.String("bucket", "", "S3 bucket name")
	prefix := flag.String("prefix", "", "Object key prefix filter")
	accessKey := flag.String("access-key", os.Getenv("AWS_ACCESS_KEY_ID"), "Access key (default from AWS_ACCESS_KEY_ID)")
	secretKey := flag.String(
		"secret-key",
		os.Getenv("AWS_SECRET_ACCESS_KEY"),
		"Secret key (default from AWS_SECRET_ACCESS_KEY)",
	)
	flag.Parse()

	if *bucket == "" {
		log.Fatalf("-bucket 은 필수입니다")
	}

	if *accessKey == "" || *secretKey == "" {
		log.Fatalf("access key/secret key 가 필요합니다 (플래그 또는 환경변수 설정)")
	}

	creds := apis3ls.Credentials{AccessKey: *accessKey, SecretKey: *secretKey}

	ctx := context.Background()

	objects, err := apis3ls.ListAllObjects(ctx, *endpoint, creds, *bucket, *prefix)
	if err != nil {
		log.Fatalf("list objects 실패: %v", err)
	}

	for _, obj := range objects {
		size := fmt.Sprintf("%d", obj.Size)

		key := ""
		if obj.Key != nil {
			key = *obj.Key
		}
		// 단순 출력: Size Key
		fmt.Printf("%s\t%s\n", size, key) //nolint:forbidigo
	}
}

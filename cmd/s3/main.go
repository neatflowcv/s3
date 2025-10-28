package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"runtime/debug"

	"github.com/joho/godotenv"
	"github.com/neatflowcv/s3/internal/pkg/domain"
	apis3ls "github.com/neatflowcv/s3/pkg/s3ls"
	"github.com/urfave/cli/v3"
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

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("load .env file: %v", err)
	}

	var (
		endpoint string
		access   string
		secret   string
		bucket   string
	)

	app := &cli.Command{ //nolint:exhaustruct
		Name: "s3",
		Flags: []cli.Flag{
			&cli.StringFlag{ //nolint:exhaustruct
				Name:        "endpoint",
				Value:       "",
				Usage:       "Ceph RGW endpoint (http/https)",
				Sources:     cli.EnvVars("ENDPOINT"),
				Destination: &endpoint,
			},
			&cli.StringFlag{ //nolint:exhaustruct
				Name:        "access",
				Value:       "",
				Usage:       "Access key",
				Sources:     cli.EnvVars("ACCESS"),
				Destination: &access,
			},
			&cli.StringFlag{ //nolint:exhaustruct
				Name:        "secret",
				Value:       "",
				Usage:       "Secret key",
				Sources:     cli.EnvVars("SECRET"),
				Destination: &secret,
			},
		},
		Commands: []*cli.Command{
			{
				Name: "ls",
				Arguments: []cli.Argument{
					&cli.StringArg{ //nolint:exhaustruct
						Name:        "bucket",
						UsageText:   "S3 bucket name",
						Destination: &bucket,
					},
				},
				Action: func(ctx context.Context, c *cli.Command) error {
					return ls(ctx, endpoint, access, secret, bucket)
				},
			},
		},
	}

	err = app.Run(context.Background(), os.Args)
	if err != nil {
		log.Fatalf("run app: %v", err)
	}
}

func ls(ctx context.Context, endpoint, access, secret, bucket string) error {
	creds := domain.NewCredentials(access, secret)

	objects, err := apis3ls.ListAllObjects(ctx, endpoint, creds, bucket, "")
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

	return nil
}

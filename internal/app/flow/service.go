package flow

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/neatflowcv/s3/internal/pkg/domain"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) ListObjects(
	ctx context.Context,
	endpoint string,
	creds *domain.Credentials,
	bucket string,
	prefix string,
) ([]s3types.Object, error) {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("us-east-1"), // region은 ceph rgw 호환을 위해 기본값으로 us-east-1 사용
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			creds.AccessKey(), creds.SecretKey(), "",
		)),
	)
	if err != nil {
		return nil, fmt.Errorf("load default config: %w", err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		// Path-style 강제 및 커스텀 엔드포인트 지정
		o.UsePathStyle = true
		o.BaseEndpoint = aws.String(endpoint)
	})

	pager := s3.NewListObjectsV2Paginator(client, &s3.ListObjectsV2Input{ //nolint:exhaustruct
		Bucket: &bucket,
		Prefix: &prefix,
	})

	var all []s3types.Object

	for pager.HasMorePages() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("list objects: %w", err)
		}

		all = append(all, page.Contents...)
	}

	return all, nil
}

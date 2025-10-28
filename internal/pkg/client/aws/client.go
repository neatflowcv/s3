package aws

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/neatflowcv/s3/internal/pkg/client"
	"github.com/neatflowcv/s3/internal/pkg/domain"
)

var _ client.Client = (*Client)(nil)

type Client struct {
	client *s3.Client
}

func NewClient(ctx context.Context, endpoint string, accessKey, secretKey string) (*Client, error) {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("us-east-1"), // region은 ceph rgw 호환을 위해 기본값으로 us-east-1 사용
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			accessKey, secretKey, "",
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

	return &Client{
		client: client,
	}, nil
}

func (c *Client) ListObjects(ctx context.Context, bucket string) ([]*domain.Object, error) {
	pager := s3.NewListObjectsV2Paginator(c.client, &s3.ListObjectsV2Input{ //nolint:exhaustruct
		Bucket: &bucket,
	})

	var all []s3types.Object

	for pager.HasMorePages() {
		page, err := pager.NextPage(ctx)
		if err != nil {
			return nil, fmt.Errorf("list objects: %w", err)
		}

		all = append(all, page.Contents...)
	}

	var ret []*domain.Object

	for _, obj := range all {
		key := ""
		if obj.Key != nil {
			key = *obj.Key
		}

		size := uint64(0)

		if obj.Size != nil {
			if *obj.Size < 0 {
				return nil, ErrObjectSizeNegative
			}

			size = uint64(*obj.Size)
		}

		ret = append(ret, domain.NewObject(key, size))
	}

	return ret, nil
}

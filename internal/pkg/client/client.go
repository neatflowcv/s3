package client

import (
	"context"

	"github.com/neatflowcv/s3/internal/pkg/domain"
)

type Client interface {
	ListObjects(ctx context.Context, bucket string) ([]*domain.Object, error)
	HeadObject(ctx context.Context, bucket, key string) (*domain.Head, error)
}

package flow

import (
	"context"
	"fmt"

	"github.com/neatflowcv/s3/internal/pkg/client"
)

type Service struct {
	client client.Client
}

func NewService(client client.Client) *Service {
	return &Service{
		client: client,
	}
}

func (s *Service) ListObjects(
	ctx context.Context,
	bucket string,
) ([]*Object, error) {
	objects, err := s.client.ListObjects(ctx, bucket)
	if err != nil {
		return nil, fmt.Errorf("list objects: %w", err)
	}

	return fromObjects(objects), nil
}

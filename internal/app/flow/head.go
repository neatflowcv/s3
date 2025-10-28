package flow

import "github.com/neatflowcv/s3/internal/pkg/domain"

type Head struct {
	Key         string
	ContentType string
}

func fromHead(head *domain.Head) *Head {
	return &Head{
		Key:         head.Key(),
		ContentType: head.ContentType(),
	}
}

func fromHeads(heads []*domain.Head) []*Head {
	var ret []*Head
	for _, head := range heads {
		ret = append(ret, fromHead(head))
	}

	return ret
}

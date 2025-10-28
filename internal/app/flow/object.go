package flow

import "github.com/neatflowcv/s3/internal/pkg/domain"

type Object struct {
	Key  string
	Size uint64
}

func fromObject(obj *domain.Object) *Object {
	return &Object{
		Key:  obj.Key(),
		Size: obj.Size(),
	}
}

func fromObjects(objs []*domain.Object) []*Object {
	var ret []*Object
	for _, obj := range objs {
		ret = append(ret, fromObject(obj))
	}

	return ret
}

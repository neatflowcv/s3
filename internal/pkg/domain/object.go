package domain

type Object struct {
	key  string
	size uint64
}

func NewObject(key string, size uint64) *Object {
	return &Object{
		key:  key,
		size: size,
	}
}

func (o *Object) Key() string {
	return o.key
}

func (o *Object) Size() uint64 {
	return o.size
}

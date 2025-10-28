package domain

type Head struct {
	key         string
	contentType string
}

func NewHead(key, contentType string) *Head {
	if key == "" {
		panic("key is empty")
	}

	if contentType == "" {
		panic("contentType is empty")
	}

	return &Head{
		key:         key,
		contentType: contentType,
	}
}

func (h *Head) Key() string {
	return h.key
}

func (h *Head) ContentType() string {
	return h.contentType
}

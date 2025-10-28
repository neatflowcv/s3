package domain

type Head struct {
	key         string
	contentType string
}

func NewHead(key, contentType string) *Head {
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

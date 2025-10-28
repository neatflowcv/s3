package domain

type Credentials struct {
	accessKey string
	secretKey string
}

func NewCredentials(accessKey, secretKey string) *Credentials {
	return &Credentials{
		accessKey: accessKey,
		secretKey: secretKey,
	}
}

func (c *Credentials) AccessKey() string {
	return c.accessKey
}

func (c *Credentials) SecretKey() string {
	return c.secretKey
}

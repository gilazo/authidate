package auth

type Encryptable interface {
	Encrypt() string
}

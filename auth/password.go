package auth

type Password struct {
	Value string
}

func (p Password) Encrypt() string {
	return "*encrypted*"
}

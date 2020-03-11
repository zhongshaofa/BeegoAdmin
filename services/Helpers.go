package services

import (
	"crypto/md5"
	"fmt"
)

type Helpers struct {
}

func (s *Helpers) Password(password string) string {
	data := []byte(password)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	return md5str1
}

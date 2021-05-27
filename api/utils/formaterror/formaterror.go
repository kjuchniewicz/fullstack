package formaterror

import (
	"errors"
	"strings"
)

func FormatError(err string) error {
	if strings.Contains(err, "nickname") {
		return errors.New("nick już zabrany")
	}
	if strings.Contains(err, "email") {
		return errors.New("email już zabrany")
	}
	if strings.Contains(err, "title") {
		return errors.New("tytuł już zabrany")
	}
	if strings.Contains(err, "hashedPassword") {
		return errors.New("nieprawidłowe hasło")
	}
	return errors.New("nieprawidłowe szczególy")
}

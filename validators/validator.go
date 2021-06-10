package validators

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

func ValidateComment(comment struct {
	Body string `json:"body"`
}) error {
	return validation.ValidateStruct(&comment,
		validation.Field(&comment.Body, validation.Required, validation.Length(2, 50)),
	)
}

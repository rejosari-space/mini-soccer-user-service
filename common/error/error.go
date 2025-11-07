package error

import (
	"errors"
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

type ValidationResponse struct {
	Field   string `json:"field,omitempty"`
	Message string `json:"message,omitempty"`
}

var ErrValidator = map[string]string{}

func ErrValidationResponse(err error) (validationResponse []ValidationResponse) {
	var fieldErrors validator.ValidationErrors

	if errors.As(err, &fieldErrors) {
		for _, err := range fieldErrors {
			switch err.Tag() {
			case "required":
				validationResponse = append(validationResponse, ValidationResponse{
					Field:   err.Field(),
					Message: fmt.Sprintf("%s is required", err.Field()),
				})

			case "email":
				validationResponse = append(validationResponse, ValidationResponse{
					Field:   err.Field(),
					Message: fmt.Sprintf("%s is not valid email", err.Field()),
				})

			default:
				errValidator, ok := ErrValidator[err.Tag()]

				if ok {

					count := strings.Count(errValidator, "%s")

					if count == 1 {

						validationResponse = append(validationResponse, ValidationResponse{
							Field:   err.Field(),
							Message: fmt.Sprintf(errValidator, err.Field()),
						})
					} else {

						validationResponse = append(validationResponse, ValidationResponse{
							Field:   err.Field(),
							Message: fmt.Sprintf(errValidator, err.Field(), err.Param()),
						})

					}
				} else {
					validationResponse = append(validationResponse, ValidationResponse{
						Field:   err.Field(),
						Message: fmt.Sprintf("something wrong on %s; %s,", err.Field(), err.Tag()),
					})
				}

			}
		}
	}

	return validationResponse
}

func WrapError(err error) error {
	logrus.Printf("error %v", err)
	return err
}

// READABLE CODE

// // ValidationResponse adalah bentuk respon untuk setiap field yang gagal validasi.
// type ValidationResponse struct {
// 	Field   string `json:"field,omitempty"`
// 	Message string `json:"message,omitempty"`
// }

// // validationMessages berisi template pesan untuk tag validator kustom.
// // Template boleh menggunakan %s untuk field dan %s untuk param (param optional).
// var validationMessages = map[string]string{
// 	// contoh:
// 	// "min": "%s must be at least %s characters",
// 	// "max": "%s cannot be longer than %s characters",
// }

// // ErrValidationResponse mengubah error dari validator menjadi slice ValidationResponse.
// func ErrValidationResponse(err error) []ValidationResponse {
// 	var out []ValidationResponse

// 	// Periksa apakah err adalah validator.ValidationErrors
// 	var fieldErrors validator.ValidationErrors
// 	if !errors.As(err, &fieldErrors) {
// 		return out
// 	}

// 	for _, fe := range fieldErrors {
// 		// fe adalah validator.FieldError
// 		switch fe.Tag() {
// 		case "required":
// 			out = append(out, ValidationResponse{
// 				Field:   fe.Field(),
// 				Message: fmt.Sprintf("%s is required", fe.Field()),
// 			})
// 		case "email":
// 			out = append(out, ValidationResponse{
// 				Field:   fe.Field(),
// 				Message: fmt.Sprintf("%s is not a valid email", fe.Field()),
// 			})
// 		default:
// 			// Coba ambil pesan dari map; jika tidak ada, pakai fallback
// 			if tmpl, ok := validationMessages[fe.Tag()]; ok {
// 				// aman: selalu pass field dan param; fmt.Sprintf mengabaikan arg ekstra
// 				out = append(out, ValidationResponse{
// 					Field:   fe.Field(),
// 					Message: fmt.Sprintf(tmpl, fe.Field(), fe.Param()),
// 				})
// 			} else {
// 				out = append(out, ValidationResponse{
// 					Field:   fe.Field(),
// 					Message: fmt.Sprintf("validation failed for %s: %s", fe.Field(), fe.Tag()),
// 				})
// 			}
// 		}
// 	}

// 	return out
// }

// // WrapError mencatat error dan mengembalikannya (boleh dihapus atau diperluas).
// func WrapError(err error) error {
// 	if err == nil {
// 		return nil
// 	}
// 	logrus.WithError(err).Error("wrapped error")
// 	return err
// }

package validation

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/CHenrique-Oliv/estudo-go/src/config/rest_err"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	ptBr_transl "github.com/go-playground/validator/v10/translations/pt_BR"
)

var (
	Validate = validator.New()

	transl ut.Translator
)

func init() {
	if val, ok := binding.Validator.Engine().(*validator.Validate); ok {
		ptBr := pt_BR.New()
		en := en.New()
		unt := ut.New(ptBr, en)

		var found bool
		transl, found = unt.GetTranslator("pt_BR")

		if !found {
			fmt.Println("Alerta de Validação: Tradutor 'pt_BR' não encontrado. As mensagens de erro não serão traduzidas.")
		}

		if transl != nil {
			err := ptBr_transl.RegisterDefaultTranslations(val, transl)
			if err != nil {
				fmt.Printf("Erro ao registrar traduções padrão: %v\n", err)
			}
		}
	}
}

func ValidateUserErro(validation_err error) *rest_err.RestErr {
	var jsonErr *json.UnmarshalTypeError
	var jsonValidationError validator.ValidationErrors

	if errors.As(validation_err, &jsonErr) {
		return rest_err.NewBadRequestError(fmt.Sprintf(
			"Invalid field type: Field %s must be of type %s",
			jsonErr.Field,
			jsonErr.Type.String(),
		))

	} else if errors.As(validation_err, &jsonValidationError) {
		errorsCauses := []rest_err.Causes{}

		for _, e := range jsonValidationError {
			var translatedMessage string

			if transl != nil {
				translatedMessage = e.Translate(transl)
			} else {
				translatedMessage = fmt.Sprintf("Validation failed on the '%s' tag for field '%s'", e.Tag(), e.Field())
			}

			cause := rest_err.Causes{
				Message: translatedMessage,
				Field:   e.Field(),
			}
			errorsCauses = append(errorsCauses, cause)
		}

		return rest_err.NewBadRequestValidationError("Some fields are invalid", errorsCauses)

	} else {
		return rest_err.NewBadRequestError("Error trying to process the request body")
	}
}

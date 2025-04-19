package validator

import (
	"reflect"
	"strings"

	validate "github.com/gobuffalo/validate/v3"
	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"
)

// объединение *validate.Errors в одну ошибку *echo.HTTPError
func collectHttpError(validateErrors *validate.Errors) error {
	if len(validateErrors.Errors) > 0 {
		// словарь для ошибок
		errMap := make(map[string]string, len(validateErrors.Errors))

		for key, value := range validateErrors.Errors {
			errMap[key] = value[0]
		}
		// возвращаем *echo.HTTPError
		httpError := echo.NewHTTPError(400, errMap)
		return httpError
	}
	return nil
}

// возврат ошибок валидации
// принимает интерфейс валидируемыемых структур с входными данными (у которых есть метод IsValid(errors *validate.Errors))
func Validate(dataToValidate validate.Validator) error {
	validateErrors := validate.NewErrors()

	// базовая валидация входных данных по тегам структуры
	baseValidator(dataToValidate, validateErrors)
	// если базовая валидация не прошла, то возвращаем ошибку
	if err := collectHttpError(validateErrors); err != nil {
		return err
	}

	// дополнительная валидация входных данных (в методе IsValid у структуры)
	validateErrors = validate.Validate(dataToValidate)
	// если дополнительная валидация не прошла, то возвращаем ошибку
	if err := collectHttpError(validateErrors); err != nil {
		return err
	}
	return nil
}

// базовая валидация структуры по тегам
func baseValidator(givenStruct validate.Validator, errors *validate.Errors) {
	// Получаем значение структуры (с разыменовыванием поинтера через Elem())
	var structValue reflect.Value = reflect.ValueOf(givenStruct).Elem()

	// полная информация по полю самого типа структуры (название, тип, значение и т.д.)
	var fieldInfo reflect.StructField
	// значение поля данного объекта структуры
	var fieldValue reflect.Value
	// тип поля в виде строки
	var fieldType string
	// значение тега myvalid
	var myvalidTag string
	// найден ли тег у поля структуры
	var isFound bool

	// перебираем все поля структуры и проверяем теги каждого поля
	for i := 0; i < structValue.NumField(); i++ {
		fieldInfo = structValue.Type().Field(i)
		fieldValue = structValue.Field(i)
		fieldType = fieldInfo.Type.String()

		myvalidTag, isFound = fieldInfo.Tag.Lookup("myvalid")
		if isFound {
			// перебор доступных для валидации типов полей структуры
			switch {
			// строка
			case fieldType == "string":
				// валидация для поля структуры строкового типа
				myStringValidator(fieldInfo, fieldValue.String(), myvalidTag, errors)

			// целое число
			case strings.HasPrefix(fieldType, "int"):
				// валидация для поля структуры целочислленного типа
				myIntValidator(fieldInfo, fieldValue.Int(), myvalidTag, errors)

			// вещественное число
			case strings.HasPrefix(fieldType, "float"):
				// валидация для поля структуры вещественного типа
				myFloatValidator(fieldInfo, fieldValue.Float(), myvalidTag, errors)

			// uuid
			case fieldType == "uuid.UUID":
				// валидация для поля структуры типа uuid
				myUUIDValidator(fieldInfo, fieldValue.Interface().(uuid.UUID), myvalidTag, errors)
			}
		}
	}
}

package validator

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	validate "github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
)

func myStringValidator(fieldInfo reflect.StructField, fieldValue string, validateTagValues string, errors *validate.Errors) {
	// имя поля для составления ошибки (выбирает значение из тега json; если такого нет - берёт собственно имя поля)
	fieldNameForError, isFound := fieldInfo.Tag.Lookup("json")
	if !isFound {
		fieldNameForError = fieldInfo.Name
	}

	// перебираем значения тега validateTagValues
	for _, tagValue := range strings.Split(validateTagValues, "|") {
		switch {
		// обязательное поле
		case tagValue == "required":
			// валидация средствами библиотеки
			errors.Append(validate.Validate(
				&validators.StringIsPresent{
					Name:    fieldInfo.Name, // название поля
					Field:   fieldValue,     // значение поля
					Message: fmt.Sprintf("%s field must not be blank", fieldNameForError),
				},
			))

		// валидация email
		case tagValue == "email":
			// валидация средствами библиотеки
			errors.Append(validate.Validate(
				&validators.EmailIsPresent{
					Name:    fieldInfo.Name, // название поля
					Field:   fieldValue,     // значение поля
					Message: "Email is not in the right format",
				},
			))

		// длина больше чем ... (пример, "min:8")
		case strings.HasPrefix(tagValue, "min"):
			// парсинг минимальной длины из тега
			minLenInt, err := strconv.Atoi(strings.TrimPrefix(tagValue, "min:"))
			if err != nil {
				continue
			}
			// проверка значения поля на соответствие минимальной длине
			if len(fieldValue) < minLenInt {
				errors.Add(fieldNameForError, fmt.Sprintf("%s field must contain at least %d symbols", fieldNameForError, minLenInt))
			}

		// длина меньше чем ... (пример, "max:100")
		case strings.HasPrefix(tagValue, "max"):
			// парсинг максимальной длины из тега
			maxLenInt, err := strconv.Atoi(strings.TrimPrefix(tagValue, "max:"))
			if err != nil {
				continue
			}
			// проверка значения поля на соответствие минимальной длине
			if len(fieldValue) > maxLenInt {
				errors.Add(fieldNameForError, fmt.Sprintf("%s field must contain less than %d symbols", fieldNameForError, maxLenInt))
			}
		}
	}
}

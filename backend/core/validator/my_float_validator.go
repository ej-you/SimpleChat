package validator

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	validate "github.com/gobuffalo/validate/v3"
)

func myFloatValidator(fieldInfo reflect.StructField, fieldValue float64, validateTagValues string, errors *validate.Errors) {
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
			// если значение поля не указано (0.0 по умолчанию)
			if fieldValue == 0.0 {
				errors.Add(fieldNameForError, fmt.Sprintf("%s field must not be blank", fieldNameForError))
			}

		// число (float) больше чем ... (пример, "min:8")
		case strings.HasPrefix(tagValue, "min"):
			// парсинг минимального значения из тега (в float64)
			minFloat, err := strconv.ParseFloat(strings.TrimPrefix(tagValue, "min:"), 64)
			if err != nil {
				continue
			}
			// проверка значения поля на соответствие минимальному значению
			if fieldValue < minFloat {
				errors.Add(fieldNameForError, fmt.Sprintf("%s field must be greater than or equal to %f", fieldNameForError, minFloat))
			}

		// число (float) меньше чем ... (пример, "max:100")
		case strings.HasPrefix(tagValue, "max"):
			// парсинг максимального значения из тега (в float64)
			maxFloat, err := strconv.ParseFloat(strings.TrimPrefix(tagValue, "max:"), 64)
			if err != nil {
				continue
			}
			// проверка значения поля на соответствие максимальному значению
			if fieldValue > maxFloat {
				errors.Add(fieldNameForError, fmt.Sprintf("%s field must be less than or equal to %f", fieldNameForError, maxFloat))
			}
		}
	}
}

package validator

import (
	"fmt"
	"reflect"
	"strings"

	validate "github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
    googleUUID "github.com/google/uuid"
    gofrsUUID "github.com/gofrs/uuid"
)


func myUUIDValidator(fieldInfo reflect.StructField, fieldValue googleUUID.UUID, validateTagValues string, errors *validate.Errors) {
	// имя поля для составления ошибки (выбирает значение из тега json; если такого нет - берёт собственно имя поля)
	fieldNameForError, isFound := fieldInfo.Tag.Lookup("json")
	if !isFound {
		fieldNameForError = fieldInfo.Name
	}

	// получаем значение uuid по структуре из другого пакета (github.com/gofrs/uuid)
	fieldValueBytes, _ := fieldValue.MarshalBinary()
	fieldValueGofrsUUID, _ := gofrsUUID.FromBytes(fieldValueBytes)

	// перебираем значения тега validateTagValues
	for _, tagValue := range strings.Split(validateTagValues, "|") {
		switch {
			// обязательное поле
			case tagValue == "required":
				// валидация средствами библиотеки
				errors.Append(validate.Validate(
					&validators.UUIDIsPresent{
						Name: fieldInfo.Name, // название поля
						Field: fieldValueGofrsUUID , // значение поля
						Message: fmt.Sprintf("%s field must not be blank", fieldNameForError),
					},
				))
		}
	}
}

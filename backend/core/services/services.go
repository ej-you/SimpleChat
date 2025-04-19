package services

import (
	"SimpleChat/backend/core/db/models"
)

// нахождение пересечения двух срезов с чатами юзеров
func IntersectUserChats(firstSlice, secondSlice []models.Chat) []models.Chat {
	intersect := make([]models.Chat, 0)

	// перебор элементов первого среза на совпадение с элементом из второго среза
	for _, first := range firstSlice {
		for _, second := range secondSlice {
			if first.ID == second.ID {
				intersect = append(intersect, first)
				break
			}
		}
	}
	return intersect
}

package services

import (
	"hex_to_color_name/models"
)

func GetColorName(colorNameApiResponse *models.ColorNameApiResponse) string {
	return colorNameApiResponse.GetColor().GetName()
}

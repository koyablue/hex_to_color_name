package models

type Rgb struct {
	R int `json:"r"`
	G int `json:"g"`
	B int `json:"b"`
}

type Color struct {
	Name         string  `json:"name"`
	Hex          string  `json:"hex"`
	Rgb          Rgb     `json:"rgb"`
	Distance     float64 `json:"distance"`
	Luminance    float64 `json:"luminance"`
	RequestedHex string  `json:"requestedHex"`
}

//API戻り値の型
type ColorNameApiResponse struct {
	Colors []Color `json:"colors"`
}

func (apiResponse *ColorNameApiResponse) GetColor() *Color {
	return &apiResponse.Colors[0]
}

func (color *Color) GetName() string {
	return color.Name
}

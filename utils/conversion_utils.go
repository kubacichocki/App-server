package utils

func KgToPounds(kg float64) float64 {
	return kg * 2.20462
}

func PoundsToKg(pounds float64) float64 {
	return pounds / 2.20462
}

func KmToMiles(km float64) float64 {
	return km * 0.62137119
}

func MilesToKm(miles float64) float64 {
	return miles * 1.609344
}

func CmToFeetInches(cm float64) (int, float64) {
	totalInches := cm / 2.54
	feet := int(totalInches) / 12
	inches := totalInches - float64(feet*12)
	return feet, inches
}

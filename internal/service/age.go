package service

import "time"

func CalculateAge(dob time.Time) int {
	now := time.Now()
	age := now.Year() - dob.Year()

	// Adjust if birthday hasn't occurred yet this year
	if now.YearDay() < dob.YearDay() {
		age--
	}

	return age
}

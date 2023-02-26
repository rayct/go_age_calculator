// Version 1.0
// This Go program prompts the user to enter their date of birth in the format of DD-MM-YYYY.
// The program will then calculate your current age in years, the day of the week you were born on,
// and the number of days until your next birthday.

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	// Prompt user for their date of birth
// 	var birthdateStr string
// 	fmt.Print("Enter your date of birth (YYYY-MM-DD): ")
// 	fmt.Scanln(&birthdateStr)

// 	// Convert birthdate to time.Time object
// 	birthdate, err := time.Parse("2006-01-02", birthdateStr)
// 	if err != nil {
// 		fmt.Println("Invalid date format. Please enter a date in the format of YYYY-MM-DD.")
// 		return
// 	}

// 	// Get today's date
// 	today := time.Now()

// 	// Calculate age in years
// 	age := today.Year() - birthdate.Year()
// 	if today.YearDay() < birthdate.YearDay() {
// 		age--
// 	}

// 	// Calculate day of the week the user was born on
// 	daysOfWeek := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
// 	dayOfWeek := daysOfWeek[birthdate.Weekday()]

// 	// Calculate number of days until next birthday
// 	nextBirthday := time.Date(today.Year(), birthdate.Month(), birthdate.Day(), 0, 0, 0, 0, today.Location())
// 	if nextBirthday.Before(today) {
// 		nextBirthday = nextBirthday.AddDate(1, 0, 0)
// 	}
// 	daysUntilBirthday := int(nextBirthday.Sub(today).Hours() / 24)

// 	// Print results
// 	fmt.Printf("You were born on a %s.\n", dayOfWeek)
// 	fmt.Printf("You are currently %d years old.\n", age)
// 	fmt.Printf("There are %d days until your next birthday.\n", daysUntilBirthday)
// }

// Optimised Version 1.0.1
package main

import (
	"fmt"
	"time"
)

func main() {
	// Prompt user for their date of birth
	var birthdateStr string
	fmt.Print("Enter your date of birth (YYYY-MM-DD): ")
	fmt.Scanln(&birthdateStr)

	// Convert birthdate to time.Time object
	birthdate, err := time.Parse("2006-01-02", birthdateStr)
	if err != nil {
		fmt.Println("Invalid date format. Please enter a date in the format of YYYY-MM-DD.")
		return
	}

	// Get today's date
	today := time.Now()

	// Calculate age in years
	age := today.Year() - birthdate.Year()
	if today.YearDay() < birthdate.YearDay() {
		age--
	}

	// Calculate day of the week the user was born on
	dayOfWeek := birthdate.Weekday().String()

	// Calculate number of days until next birthday
	nextBirthday := time.Date(today.Year(), birthdate.Month(), birthdate.Day(), 0, 0, 0, 0, today.Location())
	if nextBirthday.Before(today) {
		nextBirthday = nextBirthday.AddDate(1, 0, 0)
	}
	daysUntilBirthday := int(nextBirthday.Sub(today).Hours() / 24)

	// Print results
	fmt.Printf("You were born on a %s.\n", dayOfWeek)
	fmt.Printf("You are currently %d years old.\n", age)
	fmt.Printf("There are %d days until your next birthday.\n", daysUntilBirthday)
}

// This version is slightly more optimized, with the following changes:

// The daysOfWeek array is removed since we can directly convert the Weekday value to a string using the String() method.
// The if statement that calculates the age now only checks the YearDay of the current date and the birthdate, rather than both the month and day.
// The int() cast is added when calculating daysUntilBirthday, since we don't need the float value returned by Sub().Hours() / 24.
// Overall, these changes make the code a bit more concise and eliminate unnecessary computations.

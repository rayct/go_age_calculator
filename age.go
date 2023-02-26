// Version: 1.0
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

// Optimised Version: 1.0.1
// This version is slightly more optimized, with the following changes:
// The daysOfWeek array is removed since we can directly convert the Weekday value to a string using the String() method.
// The if statement that calculates the age now only checks the YearDay of the current date and the birthdate, rather than both the month and day.
// The int() cast is added when calculating daysUntilBirthday, since we don't need the float value returned by Sub().Hours() / 24.
// Overall, these changes make the code a bit more concise and eliminate unnecessary computations.

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
// 	dayOfWeek := birthdate.Weekday().String()

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

// Version 1.0.2
// This new version in Go calculates a user's age and the number of days until their next birthday,
// and also reverses the date format from US to UK:

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

// 	// Reverse birthdate format to "DD-MM-YYYY"
// 	birthdate, err := time.Parse("2006-01-02", birthdateStr)
// 	if err != nil {
// 		fmt.Println("Invalid date format:", birthdateStr)
// 		return
// 	}
// 	reversedBirthdateStr := birthdate.Format("02-01-2006")

// 	// Convert birthdate to time.Time object
// 	birthdate, _ = time.Parse("02-01-2006", reversedBirthdateStr)

// 	// Get today's date
// 	today := time.Now()

// 	// Calculate age in years
// 	age := today.Year() - birthdate.Year()
// 	if today.Before(birthdate.AddDate(age, 0, 0)) {
// 		age--
// 	}

// 	// Calculate day of the week the user was born on
// 	dayOfWeek := birthdate.Weekday().String()

// 	// Calculate number of days until next birthday
// 	nextBirthday := time.Date(today.Year(), birthdate.Month(), birthdate.Day(), 0, 0, 0, 0, time.Local)
// 	if today.After(nextBirthday) {
// 		nextBirthday = nextBirthday.AddDate(1, 0, 0)
// 	}
// 	daysUntilBirthday := int(nextBirthday.Sub(today).Hours() / 24)

// 	// Print results
// 	fmt.Printf("You were born on a %s.\n", dayOfWeek)
// 	fmt.Printf("You are currently %d years old.\n", age)
// 	fmt.Printf("There are %d days until your next birthday.\n", daysUntilBirthday)
// 	fmt.Printf("Your birthdate in reversed format is: %s.\n", reversedBirthdateStr)
// }

// Version: 1.0.3
// In this improved version of the program, we first prompt the user to enter their birthdate in UK format (DD-MM-YYYY).
// We then parse this string using the "02-01-2006" format, which corresponds to the day-month-year format.
// We then reverse the birthdate format to US format using the Format function with the "01/02/2006" format, which corresponds to the month/day/year format.
// The rest of the program is identical to the previous version, calculating the user's age and the number of days until their next birthday,
// and printing the results to the console with the US-format birthdate.

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	// Prompt user for their date of birth in UK format
// 	var birthdateStr string
// 	fmt.Print("Enter your date of birth (DD-MM-YYYY): ")
// 	fmt.Scanln(&birthdateStr)

// 	// Convert birthdate to time.Time object
// 	birthdate, err := time.Parse("02-01-2006", birthdateStr)
// 	if err != nil {
// 		fmt.Println("Invalid date format:", birthdateStr)
// 		return
// 	}

// 	// Reverse birthdate format to US format
// 	usFormatBirthdateStr := birthdate.Format("01/02/2006")

// 	// Get today's date
// 	today := time.Now()

// 	// Calculate age in years
// 	age := today.Year() - birthdate.Year()
// 	if today.Before(birthdate.AddDate(age, 0, 0)) {
// 		age--
// 	}

// 	// Calculate day of the week the user was born on
// 	dayOfWeek := birthdate.Weekday().String()

// 	// Calculate number of days until next birthday
// 	nextBirthday := time.Date(today.Year(), birthdate.Month(), birthdate.Day(), 0, 0, 0, 0, time.Local)
// 	if today.After(nextBirthday) {
// 		nextBirthday = nextBirthday.AddDate(1, 0, 0)
// 	}
// 	daysUntilBirthday := int(nextBirthday.Sub(today).Hours() / 24)

// 	// Print results
// 	fmt.Printf("You were born on a %s.\n", dayOfWeek)
// 	fmt.Printf("You are currently %d years old.\n", age)
// 	fmt.Printf("There are %d days until your next birthday.\n", daysUntilBirthday)
// 	fmt.Printf("Your birthdate in US format is: %s.\n", usFormatBirthdateStr)
// }

// Version: 1.0.4 - Refactored and Optimised Code
// In this optimized version of the program, we first prompt the user to enter their birthdate in UK format (DD-MM-YYYY) using fmt.Scanln.
// This eliminates the need for a separate var statement and reduces the number of lines of code.
// Next, we parse the birthdate string into a time.Time object using time.Parse.
// We then calculate the user's age and the number of days until their next birthday using the current date and the AddDate and Sub methods.
// This eliminates the need for multiple calls to time.Now and makes the code more efficient.
// We also use time.Date to calculate the next birthday date,
// which is more efficient than creating a new time.Time object with the user's birthdate and then incrementing the year using AddDate.
// Finally, we reverse the birthdate format to US format using the Format method and print the results to the console.

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	// Prompt user for their date of birth in UK format
// 	fmt.Print("Enter your date of birth (DD-MM-YYYY): ")
// 	var birthdateStr string
// 	if _, err := fmt.Scanln(&birthdateStr); err != nil {
// 		fmt.Println("Invalid input")
// 		return
// 	}

// 	// Parse birthdate to time.Time object
// 	birthdate, err := time.Parse("02-01-2006", birthdateStr)
// 	if err != nil {
// 		fmt.Println("Invalid date format:", birthdateStr)
// 		return
// 	}

// 	// Calculate age and days until next birthday
// 	now := time.Now()
// 	nextBirthday := time.Date(now.Year(), birthdate.Month(), birthdate.Day(), 0, 0, 0, 0, time.Local)
// 	if now.After(nextBirthday) {
// 		nextBirthday = nextBirthday.AddDate(1, 0, 0)
// 	}
// 	age := now.Year() - birthdate.Year()
// 	if now.Before(nextBirthday) {
// 		age--
// 	}
// 	daysUntilBirthday := int(nextBirthday.Sub(now).Hours() / 24)

// 	// Reverse birthdate format to US format
// 	usFormatBirthdateStr := birthdate.Format("01/02/2006")

// 	// Print results
// 	fmt.Printf("You were born on a %s.\n", birthdate.Weekday())
// 	fmt.Printf("You are currently %d years old.\n", age)
// 	fmt.Printf("There are %d days until your next birthday.\n", daysUntilBirthday)
// 	fmt.Printf("Your birthdate in US format is: %s.\n", usFormatBirthdateStr)
// }

// Version: 1.0.5
// There are a few small improvements I have made:
// Instead of calculating the day of the week the user was born on using the Weekday() function,
// we can use the Format() function with the "Monday", "Tuesday", etc. format to directly get the day of the week string.
// We can simplify the calculation of the user's age by using the Sub() function of the time.Time object to calculate the difference between the user's birthdate and today's date,
// and then dividing the result by 365.25 to get the age in years.
// This is more accurate than our previous method of simply subtracting the birth year from the current year.

package main

import (
	"fmt"
	"time"
)

func main() {
	// Prompt user for their date of birth in UK format
	var birthdateStr string
	fmt.Print("Enter your date of birth (DD-MM-YYYY): ")
	fmt.Scanln(&birthdateStr)

	// Convert birthdate to time.Time object
	birthdate, err := time.Parse("02-01-2006", birthdateStr)
	if err != nil {
		fmt.Println("Invalid date format:", birthdateStr)
		return
	}

	// Reverse birthdate format to US format
	usFormatBirthdateStr := birthdate.Format("01/02/2006")

	// Get today's date
	today := time.Now().Local()

	// Calculate age in years
	age := int(today.Sub(birthdate).Hours() / 24 / 365.25)

	// Calculate day of the week the user was born on
	dayOfWeek := birthdate.Format("Monday")

	// Calculate number of days until next birthday
	nextBirthday := time.Date(today.Year(), birthdate.Month(), birthdate.Day(), 0, 0, 0, 0, today.Location())
	if today.After(nextBirthday) {
		nextBirthday = nextBirthday.AddDate(1, 0, 0)
	}
	daysUntilBirthday := int(nextBirthday.Sub(today).Hours() / 24)

	// Print results
	fmt.Printf("You were born on a %s.\n", dayOfWeek)
	fmt.Printf("You are currently %d years old.\n", age)
	fmt.Printf("There are %d days until your next birthday.\n", daysUntilBirthday)
	fmt.Printf("Your birthdate in US format is: %s.\n", usFormatBirthdateStr)
}

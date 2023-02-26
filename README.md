# This program will prompt the user for their date of birth, and then it will calculate current age, the day of the week you were born on, and the number of days until your next birthday.

## To use the program, simply run it in a Go environment and follow the prompts.
# Version: 1.0.4
## In this optimized version of the program, we first prompt the user to enter their birthdate in UK format (DD-MM-YYYY) using fmt.Scanln.
## This eliminates the need for a separate var statement and reduces the number of lines of code.
## Next, we parse the birthdate string into a time.Time object using time.Parse.
## We then calculate the user's age and the number of days until their next birthday using the current date and the AddDate and Sub methods.
## This eliminates the need for multiple calls to time.Now and makes the code more efficient.
## We also use time.Date to calculate the next birthday date,
## which is more efficient than creating a new time.Time object with the user's birthdate and then incrementing the year using AddDate.
## Finally, we reverse the birthdate format to US format using the Format method and print the results to the console.

## Ray C. Turner
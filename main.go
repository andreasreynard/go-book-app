package main

import (
	"fmt"
	"strings"
)

func main() {
	confName := "Go Conference"
	const confTickets uint = 54
	var remTickets uint = confTickets
	remTickets -= 4
	bookings := []string{"aReynard Samsico", "mIqbal Ramadhan", "Delai R.S.", "mochAvin noLastName"}

	// fmt.Printf("Note 1: confName is %T, confTickets is %T, remTickets is %T\n", confName, confTickets, remTickets)
	// fmt.Printf("Note 2: confName's ponter name is %v, remTickets' is %v\n", &confName, &remTickets)

	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confTickets, remTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint
		// ask user for their personal data
		fmt.Print("Enter your first name: ")
		fmt.Scan(&firstName)
		fmt.Print("Enter your last name: ")
		fmt.Scan(&lastName)
		fmt.Print("Enter your email address: ")
		fmt.Scan(&email)
		fmt.Print("Enter your number of tickets: ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNum := userTickets > 0 && userTickets < remTickets

		if isValidName && isValidEmail && isValidTicketNum {
			remTickets -= userTickets
			bookings = append(bookings, firstName+" "+lastName)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n%v tickets remaining for %v.\n", firstName, lastName, userTickets, email, remTickets, confName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}
			fmt.Printf("The 1st names of bookings are: %v\n", firstNames)
			// fmt.Printf("The first value: %v\n", bookings[0])
			// fmt.Printf("Slice type: %T\n", bookings)
			// fmt.Printf("Slice length: %v\n", len(bookings))

			if remTickets == 0 {
				// end program
				fmt.Println("Our conference's booked out. Come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("1st/last name you entered is too short.")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered doesn't contain @ sign.")
			}
			if !isValidTicketNum {
				fmt.Println("Number of tickets you entered is invalid.")
			}
		}
	}

	// city := "London"

	// switch city {
	// case "New York":
	// 	//...
	// case "Singapore", "Hong Kong":
	// 	//...
	// case "London", "Berlin":
	// 	//...
	// case "Mexico City":
	// 	//...
	// default:
	// 	fmt.Print("No valid city Selected")
	// }
}

package main

import (
	"fmt"
	"sync"
	"time"
)

const confTickets uint = 50

var confName = "Go Conference"
var remTickets uint = confTickets
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	// fmt.Printf("Note 1: confName is %T, confTickets is %T, remTickets is %T\n", confName, confTickets, remTickets)
	// fmt.Printf("Note 2: confName's ponter name is %v, remTickets' is %v\n", &confName, &remTickets)
	greetUsers()

	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNum := ValidateUserInput(firstName, lastName, email, userTickets, remTickets)

	if isValidName && isValidEmail && isValidTicketNum {
		bookTicket(firstName, lastName, email, userTickets)
		wg.Add(1)
		go sendTicket(firstName, lastName, email, userTickets)

		fmt.Printf("The 1st names of bookings are: %v\n", get1stNames())

		if remTickets == 0 {
			fmt.Println("Our conference's booked out. Come back next year.")
			// break
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
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", confName)
	fmt.Printf("We have total of %v tickets and %v are still available.\n", confTickets, remTickets)
	fmt.Println("Get your tickets here to attend")
}

func get1stNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var f string
	fmt.Print("Enter your first name: ")
	fmt.Scan(&f)

	var l string
	fmt.Print("Enter your last name: ")
	fmt.Scan(&l)

	var e string
	fmt.Print("Enter your email address: ")
	fmt.Scan(&e)

	var u uint
	fmt.Print("Enter your number of tickets: ")
	fmt.Scan(&u)

	return f, l, e, u
}

func bookTicket(f string, l string, e string, u uint) {
	remTickets -= u

	var userData = UserData{
		firstName:       f,
		lastName:        l,
		email:           e,
		numberOfTickets: u,
	}

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v.\n%v tickets remaining for %v.\n", f, l, u, e, remTickets, confName)
}

func sendTicket(f string, l string, e string, u uint) {
	time.Sleep(50 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", u, f, l)
	fmt.Println("################")
	fmt.Printf("Sending ticket:\n %v \nto email adress %v\n", ticket, e)
	fmt.Println("################")
	wg.Done()
}

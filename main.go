package main

import "fmt"

func main() {

	var conferenceName string = "Go Conference"
	const conferenceTicket uint = 50
	var remainingTickets uint = 50

	var booking [50]string

	println("Welcome to", conferenceName, "booking application")
	println("we have total of", conferenceTicket, "tickets and", remainingTickets, " Tickets are still remaining  ")
	println("Get your Ticket here to attend")
	fmt.Printf("welcome to %s booking", conferenceName)

	var username string
	var lastname string
	var email string
	var userTicket uint

	for {
		//user input
		fmt.Println("\nEnter your first name: ")
		fmt.Scan(&username)

		//username = "Tom"
		//userTicket = 2
		fmt.Println("Enter your last name")
		fmt.Scan(&lastname)

		fmt.Println("enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("enter number of tickets: ")
		fmt.Scan(&userTicket)

		//if user type more than 50 ticket at ones then this condition
		if userTicket > remainingTickets {
			fmt.Printf("we have only %v tickets remaining, so you cannt book %v tickets \n", remainingTickets, userTicket)
			continue
		}

		//remainingTickets left to book
		remainingTickets = remainingTickets - userTicket

		//using Arrays and slicing to store user who is booking tickets
		booking[0] = username + " " + lastname
		//pringing whole array
		fmt.Printf("the whleo array: %v\n", booking)
		//printing first value of array
		fmt.Printf("the first value : %v\n", booking[0])
		//finding total array
		fmt.Printf("Array Type : %T\n", booking)
		//finding tootal len of array
		fmt.Printf("Array length %v\n", len(booking))

		//fmt.Printf("user %v booked %v tickets", username, userTicket)
		fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a connfirmation email at %v \n", username, lastname, userTicket, email)
		fmt.Printf("%v tickets remaining for %v \n", remainingTickets, conferenceName)

		//slicing print

		//var booking []string  //for slicing line no  11 should be comment out
		//booking =append(booking , username +" " + lastname)
		//fmt.Println("Booking :", booking)

		//checking ticket left is zero or not
		//if ticket is booked the first asked name and give message no ticket
		fmt.Println("the first name of booking are : %v \n", username)
		if remainingTickets == 0 {
			//end pogram if ticket is zero
			fmt.Println("our conference is booked out. come back next year.")
			break

		}
	}
}

package main

import (
	"fmt"
	"strings"
	"ticketbookingapp/UserInputValidator"
)

var firstname string
var lastname string
var mob_no string
var emailid string
var total_avail_tickets uint = 50
var ticket_order uint
var current_avail_ticket uint = 50
var bookings [50]string
var test_bookings []string

func main() {

	for current_avail_ticket > 0 || len(bookings) < 50 {
		userinput()

		isValidName, isValidmail, isValidMob, isValidOrder := UserInputValidator.Validateuserinput(firstname, lastname, emailid, ticket_order, current_avail_ticket, mob_no)

		if isValidName == true && isValidmail == true && isValidOrder == true && isValidMob == true {
			current_avail_ticket = current_avail_ticket - ticket_order
			display_msg()

			test_bookings = append(test_bookings, firstname+" "+lastname)

			firstnames := printfirstname(test_bookings)
			fmt.Println(" ")
			fmt.Println("Booking is done by -  %v ", firstnames)

			if current_avail_ticket == 0 {
				fmt.Println("All the ticket Sold Out, Sorry for the inconveinece")
				break
			}
		} else {
			user_input_reply(isValidMob, isValidmail, isValidOrder, isValidName)
			fmt.Printf("Sorry, We only have %v tickets, you have ordered for %v \n", current_avail_ticket, ticket_order)
			fmt.Println("Do you want still continue buying ticket - Yes / No")
			var retry string
			fmt.Scan(&retry)
			if retry == "Yes" || retry == "Y" {
				continue
			} else {
				break
			}
		}

	}
}

func userinput() {
	fmt.Println("*** Please Enter Your Details ****")
	fmt.Println(" ")
	fmt.Printf("Enter Your First Name: ")
	fmt.Scan(&firstname)

	fmt.Printf("Enter Your Last Name: ")
	fmt.Scan(&lastname)

	fmt.Printf("Enter Your Email Id: ")
	fmt.Scan(&emailid)

	fmt.Printf("Enter Your Contact No: ")
	fmt.Scan(&mob_no)

	fmt.Printf("Enter Your No Of Orders: ")
	fmt.Scan(&ticket_order)
	fmt.Println(" ")
}

func display_msg() {
	fmt.Println(" Please Check The Details Below ")
	fmt.Println("First Name :  %v", firstname)
	fmt.Println("Last Name : %v", lastname)
	fmt.Println("Mobile No : %v", mob_no)
	fmt.Println("Mobile No : %v", emailid)
	fmt.Println("Ticket Booked :  %v", ticket_order)
	fmt.Println("Current Tickets In Stock: %v", current_avail_ticket)
	fmt.Println("Total Ticket Count :  %v", total_avail_tickets)
	fmt.Println(" ")
}

func user_input_reply(isValidName bool, isValidMob bool, isValidmail bool, isValidOrder bool) {
	if !isValidName {
		fmt.Println("You have written a short name")
	}
	if !isValidMob {
		fmt.Println("You have written a short mobile, please enter 10 digit")
	}
	if !isValidmail {
		fmt.Println("You have written incorrect emailid")
	}
	if !isValidOrder {
		fmt.Println("You have written incorrect order number")
	}
}

func printfirstname(test_bookings []string) []string {
	var firstnames []string
	for _, booking := range test_bookings {
		var names = strings.Fields(booking)
		firstnames = append(firstnames, names[0])
	}
	return firstnames

}

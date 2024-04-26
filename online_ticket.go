package main

import (
	"booking_app/helper"
	"fmt"
	"strings"
)

var firstname string
var lastname string
var mob_no string
var emailid string
var city string
var total_avail_tickets uint = 50
var ticket_order uint
var current_avail_ticket uint = 50
var bookings [50]string
var test_bookings []string

func main() {

	for current_avail_ticket > 0 || len(bookings) < 50 {
		userinput()

		selectcity(city)
		isValidName, isValidmail, isValidMob, isValidOrder := helper.Validateuserinput(firstname, lastname, emailid, ticket_order, current_avail_ticket, mob_no)

		fmt.Println(isValidName, isValidmail, isValidMob, isValidOrder)
		if isValidName == true && isValidmail == true && isValidOrder == true && isValidMob == true {
			current_avail_ticket = current_avail_ticket - ticket_order
			display_msg()
			/*
				bookings[0] = firstname + " " + lastname
				fmt.Println("The Whole Array %v", bookings)
				fmt.Println("The First Array %v", bookings[0])
				fmt.Printf("Array Type %T \n", bookings)
				fmt.Println("Length Of Array %v", len(bookings))
			*/

			test_bookings = append(test_bookings, firstname+" "+lastname)
			slice_test()
			firstnames := printfirstname(test_bookings)
			fmt.Println("This are all our bookings %v ", test_bookings)
			fmt.Println("This are all our bookings with only firstnames %v ", firstnames)

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
	fmt.Printf("Enter Your First Name: ")
	fmt.Scan(&firstname)

	fmt.Printf("Enter Your Last Name: ")
	fmt.Scan(&lastname)

	fmt.Printf("Enter Your Email Id: ")
	fmt.Scan(&emailid)

	fmt.Printf("Enter Your Contact No: ")
	fmt.Scan(&mob_no)

	fmt.Printf("Enter Your City: ")
	fmt.Scan(&city)

	fmt.Printf("Enter Your No Of Orders: ")
	fmt.Scan(&ticket_order)

}

func display_msg() {
	fmt.Println("First Name :  %v", firstname)
	fmt.Println("Last Name : %v", lastname)
	fmt.Println("Mobile No : %v", mob_no)
	fmt.Println("Mobile No : %v", emailid)
	fmt.Println("City : %v", city)
	fmt.Println("Ticket Booked :  %v", ticket_order)
	fmt.Println("Current Tickets In Stock: %v", current_avail_ticket)
	fmt.Println("Total Ticket Count :  %v", total_avail_tickets)

}

func slice_test() {
	fmt.Println("The Whole Slice %v", test_bookings)
	fmt.Println("The First Slice %v", test_bookings[0])
	fmt.Printf("Slice Type %T \n", test_bookings)
	fmt.Println("Length Of Slice %v", len(test_bookings))
}

func selectcity(city string) {

	switch city {
	case "Mumbai", "Goa":
		fmt.Println("You have choosen Mumbai / Goa City")
	case "Delhi":
		fmt.Println("You have choosen Delhi City")
	case "Punjab":
		fmt.Println("You have choosen Goa City")
	case "Chennai":
		fmt.Println("You have choosen Chennai City")
	default:
		fmt.Println("Wrong Input")

	}
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

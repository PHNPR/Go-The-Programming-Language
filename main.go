package main

import (
	"fmt"
	"strconv"
	"strings"
)

const con_tickets int = 50

var con_name string = "Go Conference"
var rem_tickets uint = 50
var bookings = make([]map[string]string , 0) 

func main() {
	greetUsers()

	for {
		firstname, lastname, email, usertickets := getuserinput()
		
		isvalidname, isvalidmail, isvalidtickets := validation(firstname, lastname, email, usertickets)

		if isvalidname && isvalidmail && isvalidtickets {

			bookTicket(usertickets, firstname, lastname, email)

			first_names := getfirstnames()
			fmt.Printf("The first names of bookings are : %v\n", first_names)

			if rem_tickets == 0 {
				fmt.Println("Our Conference is booked out .Comeback later.")
				break
			}
		} else {
			if !isvalidname {
				fmt.Println("Username is too short")
			}
			if !isvalidmail {
				fmt.Println("Email address you entered doesn't contain @")
			}
			if !isvalidtickets {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", con_name)
	fmt.Printf("We have total of %v tickets and %v are still available\n", con_tickets, rem_tickets)
	fmt.Println("Get your tickets here to attend")
}

func getuserinput() (string, string, string, uint) {
	var firstname string
	var lastname string
	var email string
	var usertickets uint

	fmt.Print("Enter your first name : ")
	fmt.Scan(&firstname)

	fmt.Print("Enter your last name : ")
	fmt.Scan(&lastname)

	fmt.Print("Enter your email : ")
	fmt.Scan(&email)

	fmt.Print("Enter number of tickets : ")
	fmt.Scan(&usertickets)

	return firstname, lastname, email, usertickets
}

func validation(firstname string, lastname string, email string, usertickets uint) (bool, bool, bool) {
	isvalidname := len(firstname) >= 2 && len(lastname) >= 2
	isvalidmail := strings.Contains(email, "@")
	isvalidtickets := usertickets > 0 && rem_tickets >= usertickets
	return isvalidname, isvalidmail, isvalidtickets
}

func bookTicket(usertickets uint , firstname string , lastname string , email string) {
	rem_tickets = rem_tickets - usertickets
	var userdata = make(map[string]string)

	userdata["first_name"] = firstname
	userdata["last_name"] = lastname
	userdata["email_id"] = email
	userdata["ticket_count"] = strconv.FormatUint(uint64(usertickets) , 10)

	bookings = append(bookings , userdata)
	fmt.Printf("List of bookings is %v\n" , bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation mail at %v.\n" , firstname , lastname , usertickets , email)
	fmt.Printf("%v tickets are remaining for %v.\n" , rem_tickets , con_name)
}

func getfirstnames() []string {   
	firstnames := []string{}
	for _ , booking := range bookings {
		firstnames = append(firstnames, booking["first_name"])
	}
	return firstnames
}
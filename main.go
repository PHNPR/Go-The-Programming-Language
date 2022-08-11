package main

import (
	"fmt"
	"strings"
)

func main() {
	con_name := "Go Conference"   		 
	const con_tickets int = 50			 
	var rem_tickets uint = 50			
	bookings := []string{}				 

	fmt.Printf("Welcome to %v booking application\n" , con_name)
	fmt.Printf("We have total of %v tickets and %v are still available\n" , con_tickets , rem_tickets)
	fmt.Println("Get your tickets here to attend")

	for{
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

		isvalidname := len(firstname) >= 2 && len(lastname) >= 2
		isvalidmail := strings.Contains(email,"@")
		isvalidtickets := usertickets > 0 && rem_tickets >= usertickets

		if isvalidname && isvalidmail && isvalidtickets {
			rem_tickets = rem_tickets - usertickets

			bookings = append(bookings , firstname + " " + lastname)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation mail at %v.\n" , firstname , lastname , usertickets , email)
			fmt.Printf("%v tickets are remaining for %v.\n" , rem_tickets , con_name)

			firstnames := []string{}

			for _ , booking := range bookings {
				var names = strings.Fields(booking)
				firstnames = append(firstnames, names[0])
			}
			fmt.Printf("The first names of bookings are : %v\n" , firstnames)

			if rem_tickets == 0 {
				fmt.Println("Our Conference is booked out .Comeback later.")
				break
			}
		} else {
			if !isvalidname{ 
				fmt.Println("Username is too short")
			}
			if !isvalidmail{ 
				fmt.Println("Email address you entered doesn't contain @")
			}
			if !isvalidtickets{ 
				fmt.Println("Number of tickets you entered is invalid")
			}
		}
	} 
}
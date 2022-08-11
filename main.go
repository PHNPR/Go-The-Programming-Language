package main

import (
	"fmt"
	"strings"
)

func main() {
	con_name := "Go Conference"   		 // sweeter syntax but cant specify type and cant declare constants
	const con_tickets int = 50			 // constants cannot be changed
	var rem_tickets uint = 50			 // variables can be changed
	bookings := []string{}				 // this is a SLICE - slice is dynamic in size :)) = python list

	// Different types of numerics in GO - screenshot : int , uint ala 

	//fmt.Printf("con_name is %T , con_tickets is %T , rem_tickets is %T \n" , con_name , con_tickets , rem_tickets) // Type :)

	fmt.Printf("Welcome to %v booking application\n" , con_name)
	fmt.Printf("We have total of %v tickets and %v are still available\n" , con_tickets , rem_tickets)
	fmt.Println("Get your tickets here to attend")

	//var bookings = [50]string{}  // Array of size 50-only fixed size arrays . only same data type can be stored.
	
	// for len(di) > 7 && boks < 45 { // conditions }

	for{
		var firstname string    		// When talking user input
		var lastname string    		
		var email string    			
		var usertickets uint     

		fmt.Print("Enter your first name : ") // pointer is back to fuck  - pass the reference to the scan function
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
		// isvalidcity := city == "mumbai" || city == "chennai"

		if isvalidname && isvalidmail && isvalidtickets {
			rem_tickets = rem_tickets - usertickets

			//bookings[0] = firstname + " " + lastname
			bookings = append(bookings , firstname + " " + lastname)

			// fmt.Printf("the whole slice is %v\n" , bookings)
			// fmt.Printf("the first value of slice is %v\n" , bookings[0])
			// fmt.Printf("the slice type is %T\n" , bookings)
			// fmt.Printf("the slice length is %v\n" , len(bookings))

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation mail at %v.\n" , firstname , lastname , usertickets , email)
			fmt.Printf("%v tickets are remaining for %v.\n" , rem_tickets , con_name)

			// Arrays and Slices - 2 important data types :))
			firstnames := []string{}

			for _ , booking := range bookings {
				var names = strings.Fields(booking)
				firstnames = append(firstnames, names[0])
			}
			fmt.Printf("The first names of bookings are : %v\n" , firstnames)

			//var noTicketsRemaining bool = rem_tickets == 0
			// noTicketsRemaining := rem_tickets == 0
			// if noTicketsRemaining {
			// 	fmt.Println("Our Conference is booked out .Comeback later.")
			// 	break
			// }
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
	// city := "London"
	
	// switch city{
	// 	case "New york" :
	// 		// code
	// 	case "Singapore" ,"California" :
	// 		//code
	// 	case "London" :
	// 		//code
	// 	default :
	// 		fmt.Println("NO valid city is selected")
	// }
}
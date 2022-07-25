package main

import (
	"fmt"
	"booking-app/helper"
	"time"
	"sync"
)


//const keyword cannot be changed
const conferenceTickets = 50
var ConferenceName = "Go Conference"
//uint are variable types that cannot have negative values
var remainingTickets uint = 50

//creating an array to get a list of bookings
//var bookings [50]string
//creating a slice to get a list
var bookings = make([]userData, 0)


type userData struct {
	firstName string
	lastName string
	email string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	

	greetUsers()

		//printing the variable types
		//fmt.Printf("conferenceTicket is %T, remaiingTickets is %T, conferenceName is %T\n",conferenceTickets, remainingTickets, ConferenceName)



	
		
	firstName, lastName, email, userTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)


	if isValidName && isValidEmail && isValidTicketNumber {

			/*fmt.Printf("The whole slice: %v\n",bookings)
			fmt.Printf("The first value: %v\n",bookings[0])
			fmt.Printf("Slice type: %T\n",bookings)
			fmt.Printf("Slice length: %v\n",len(bookings))
			*/

			//looping through usernames to get their firstnames
			
		bookTicket(userTickets, firstName, lastName, email)
			
		wg.Add(1)
		//adding concurrncy to our app(being able to multitask)
		go sendTicket(userTickets, firstName, lastName, email)          

			//call function print firstnames
		firstNames := getFirstNames()
		fmt.Printf("The first names of bookings: %v\n", firstNames)


		if remainingTickets == 0{
				// end program
			fmt.Println("Our conference is booked out. Come back next year")
				//----break
		}
	}else{
		if !isValidName {
			fmt.Println("First name or last name you entered is to short")
		}
		if !isValidEmail {
			fmt.Println("Email address you entered doesnt contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("Number of tickets you entered is not correct")
		}
	}
	wg.Wait()
	

}
	


func greetUsers(){
	fmt.Printf("Welcome to %v booking application", ConferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets are left\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}


func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}


//

func getUserInput() (string, string, string, uint) {
	var firstName string
		var lastName string
		var email string
		var userTickets uint

		// ask user for their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName)

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email address: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string){
			remainingTickets = remainingTickets - userTickets


			// create a map for a user
			var userData = userData {
				firstName: firstName,
				lastName: lastName,
				email: email,
				numberOfTickets: userTickets,
			}

			//adding element to array
			//bookings[0] = firstName + " " + lastName

			//adding elements to our slice
			bookings = append(bookings, userData)
			fmt.Printf("List of bookings is %v\n", bookings)

			fmt.Printf("Thank you %v %v for booking %v tickets. You will recieve a confirmation email at %v\n",firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, ConferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(20 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("####################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("####################")
	wg.Done()
}
package main

import (
	"fmt"
	"strings"

	"github.com/manveru/faker"
)

func main() {

	var recordsRequired int

	fmt.Print("Please enter the number of user data records to generate: ")
	fmt.Scan(&recordsRequired)
	fmt.Print("\nStarting data generation...\n\n")

	printCounter := 0
	countryCounter := 0

	for i := 0; i < recordsRequired; i++ {

		fake, _ := faker.New("en")
		fmt.Printf("Record %v\n==========================================\n", i+1)

		fmt.Printf("Name:         %s\n", strings.ToUpper(fake.NamePrefix()+" "+fake.FirstName()+" "+fake.LastName()))
		fmt.Println("Address:      " + fake.StreetAddress())
		fmt.Println("City          " + fake.City())
		fmt.Println("State/County: " + fake.State())
		fmt.Println("Postcode:     " + fake.PostCode())

		country := fake.Country()
		if country == "United States of America" || country == "United Kingdom" {
			countryCounter++
		}

		fmt.Println("Country:      " + fake.Country())
		fmt.Println("Email:        " + fake.Email())
		fmt.Println("Home Phone:   " + fake.PhoneNumber())
		fmt.Println("Mobile Phone: " + fake.CellPhoneNumber())

		fmt.Printf("==========================================\n\n")
		printCounter++
	}

	fmt.Printf("%v records successfully printed, of which %v were based in the US or UK\n", printCounter, countryCounter)

}

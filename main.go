package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/manveru/faker"
	"github.com/willf/pad"
)

var (
	minRecords     = 1
	maxRecords     = 1000
	printCounter   = 0
	countryCounter = 0
)

func main() {

	recordsRequired := getInputFromUser()

	fmt.Print("\nStarting data generation...\n\n")

	for i := 0; i < recordsRequired; i++ {

		fake, _ := faker.New("en")
		fmt.Printf("Record %v\n==========================================\n", i+1)

		printItem("Name", strings.ToUpper(fake.NamePrefix()+" "+fake.FirstName()+" "+fake.LastName()))
		printItem("Address", fake.StreetAddress())
		printItem("City", fake.City())
		printItem("State/County", fake.State())
		printItem("Postcode", fake.PostCode())
		printItem("Country", fake.Country())
		printItem("Email", fake.Email())
		printItem("Home Phone", fake.PhoneNumber())
		printItem("Mobile Phone", fake.CellPhoneNumber())

		fmt.Printf("==========================================\n\n")

		evaluateCountry(fake.Country())
		printCounter++
	}

	fmt.Printf("%v records successfully printed, of which %v were based in the US or UK\n", printCounter, countryCounter)

}

func getInputFromUser() int {
	var input int

	fmt.Print("Please enter the number of user data records to generate: ")
	fmt.Scan(&input)

	if input < minRecords || input > maxRecords {
		fmt.Printf("Please enter a number between %d and %d\n", minRecords, maxRecords)
		os.Exit(1)
	}
	return input
}

func evaluateCountry(country string) {
	if country == "United States of America" || country == "United Kingdom" {
		countryCounter++
	}
}

func printItem(label string, data string) {
	fmt.Println(pad.Right(label+":", 15, " ") + data)
}

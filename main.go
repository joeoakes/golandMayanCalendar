package main

import (
	"fmt"
	"time"
)

// Tzolk'in and Haab' day names and months
var tzolkinNames = []string{"Imix'", "Ik'", "Ak'b'al", "K'an", "Chikchan", "Kimi", "Manik'", "Lamat", "Muluk", "Ok", "Chuwen", "Eb'", "B'en", "Ix", "Men", "K'ib'", "Kab'an", "Etz'nab'", "Kawak", "Ajaw"}
var haabMonths = []string{"Pop", "Wo'", "Sip", "Sotz'", "Sek", "Xul", "Yaxk'in'", "Mol", "Ch'en", "Yax", "Sak'", "Keh", "Mak", "K'ank'in", "Muwan", "Pax", "K'ayab", "Kumk'u", "Wayeb"}

// Mayan calendar epoch reference date (Gregorian date: 11 August 3114 BCE)
var mayanEpoch = time.Date(-3113, 8, 11, 0, 0, 0, 0, time.UTC)

// Days elapsed since Mayan epoch
func daysSinceMayanEpoch(date time.Time) int {
	return int(date.Sub(mayanEpoch).Hours() / 24)
}

// Calculate Tzolk'in date
func tzolkinDate(days int) (int, string) {
	number := (days % 13) + 1
	dayName := tzolkinNames[days%20]
	return number, dayName
}

// Calculate Haab' date
func haabDate(days int) (int, string) {
	dayOfYear := days % 365
	month := dayOfYear / 20
	day := dayOfYear % 20
	return day, haabMonths[month]
}

// Main function
func main() {
	// Get today's date and calculate the number of days since the Mayan epoch
	//daysSinceMayanEpoch: Calculates the number of days from the Mayan epoch
	//(August 11, 3114 BCE) to the present date.
	today := time.Now()
	days := daysSinceMayanEpoch(today)

	// Calculate Tzolk'in and Haab' dates
	//tzolkinDate: Computes the Tzolk'in date, using modular arithmetic to determine the day number (1-13) and day name.
	tzNumber, tzDayName := tzolkinDate(days)
	//haabDate: Computes the Haab' date by breaking down the 365-day cycle into months and extra days.
	haabDay, haabMonth := haabDate(days)

	// Display the results
	fmt.Printf("Today's Mayan Date:\n")
	fmt.Printf("Tzolk'in: %d %s\n", tzNumber, tzDayName)
	fmt.Printf("Haab': %d %s\n", haabDay, haabMonth)
}

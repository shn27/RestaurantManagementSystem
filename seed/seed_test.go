package seed

import (
	"fmt"
	"github.com/shn27/RestaurantManagementSystem/seed/utils"
	"strings"
	"testing"
)

func Test_splitString(t *testing.T) {
	str := "Mon, Weds 10:30 am - 6 pm / Tues 9:30 am - 12:45 am / Thurs 12:30 pm - 6:45 pm / Fri 9:45 am - 2:30 am / Sat 3 pm - 12:30 am / Sun 7 am - 12 am"
	str = strings.ToLower(str)
	words := strings.Fields(str)
	length := len(words)
	for i := 3; i < length-2; i++ {
		word := words[i]
		if word == "-" {
			closingTime := words[i+1]
			if closingTime[0] >= '0' && closingTime[0] <= '9' {
				closingTimeAmPM := words[i+2]
				openingTimeAmPM := words[i-1]
				openingTime := words[i-2]
				day := words[i-3]
				flag, day, openingHour, openingMinite, closingHour, closingMinite := utils.CheckValidity(day, closingTime, openingTime, openingTimeAmPM, closingTimeAmPM)
				if flag {
					fmt.Println("Day:", day)
					fmt.Println("Opening time: ", openingHour, ":", openingMinite)
					fmt.Println("Closing time: ", closingHour, ":", closingMinite)
					fmt.Println()
				}
			}
		}
	}
}

func Test_seed(t *testing.T) {
	processMenu()
	//processUsersWithPurChaseHistory()
}

package groupietrackers

import (
	"fmt"
	"strconv"
)

//*This struct is used to store the data from the Bing API (Inspired by Killian) :) Thanks Killian for the API
type ForBingAPI struct {
	ResourceSets []struct {
		Resources []struct {
			Point struct {
				Coordinates []float64
			}
		}
	}
}

func ToHex(entry string) string {
	/*
	* A function we turn a string into a hex string for the map use
	 */
	return fmt.Sprintf("%x", entry)
}

func DateCompare(date1 string, date2 string) bool {
	/*
	* This function compare two dates and return true if the first date is before the second date
	 */
	year1, _ := strconv.Atoi(date1[6:10])
	year2, _ := strconv.Atoi(date2[6:10])
	if year1 < year2 {
		return false
	} else if year1 == year2 {
		month1, _ := strconv.Atoi(date1[3:5])
		month2, _ := strconv.Atoi(date2[3:5])
		if month1 < month2 {
			return false
		} else if month1 == month2 {
			day1, _ := strconv.Atoi(date1[0:2])
			day2, _ := strconv.Atoi(date2[0:2])
			if day1 < day2 {
				return false
			}
		}
	}

	return true
}

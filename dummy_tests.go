package main

func largestOddNumber(num string) string {
	for i := len(num); i >= 0; i-- {
		if (num[i]-'0')%2 != 0 {
			return num[:i+1]
		}
	}

	return ""
}

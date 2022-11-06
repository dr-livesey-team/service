package date

import (
	"fmt"
)

func NextDate(date string) string {
	year := NextYear(date)
	month, day := NextMonthDay(date)

	return fmt.Sprintf("%c%c%c%c-%c%c-%c%c",
		year[0],
		year[1],
		year[2],
		year[3],
		month[0],
		month[1],
		day[0],
		day[1],
	)
}

func NextYear(date string) string {
	year := date[0:4]
	month := date[5:7]
	day := date[8:10]

	if month == "12" && day == "31" {
		next := map[byte]byte {
			'1': '2',
			'2': '3',
			'3': '4',
			'4': '5',
			'5': '6',
			'6': '7',
			'7': '8',
			'8': '9',
			'9': '0',
		}
		
		if year[0] == '9' && year[1] == '9' && year[2] == '9' {
			return fmt.Sprintf("%c000", next[year[0]])
		}

		if year[0] == '9' && year[1] == '9' {
			return fmt.Sprintf("%c%c00", year[0], next[year[1]])
		}

		if year[0] == '9' {
			return fmt.Sprintf("%c%c%c0", year[0], year[1], next[year[2]])
		}

		return fmt.Sprintf("%c%c%c%c", year[0], year[1], year[2], next[year[3]])
	}
	
	return year
}

func NextMonthDay(date string) (string, string) {
	month := date[5:7]
	day := date[8:10]

	if month == "01" && day == "31"{
		return "02", "01"
	}

	if month == "02" && day == "28"{
		return "03", "01"
	}

	if month == "03" && day == "31"{
		return "04", "01"
	}

	if month == "04" && day == "30" {
		return "05", "01"
	}

	if month == "05" && day == "31" {
		return "06", "01"
	}

	if month == "06" && day == "30" {
		return "07", "01"
	}

	if month == "07" && day == "31" {
		return "08", "01"
	}

	if month == "08" && day == "31" {
		return "09", "01"
	}

	if month == "09" && day == "30" {
		return "10", "01"
	}

	if month == "10" && day == "31" {
		return "11", "01"
	}

	if month == "11" && day == "30" {
		return "12", "01"
	}

	if month == "12" && day == "31" {
		return "01", "01"
	}

	next := map[string]string{
		"01": "02",
		"02": "03",
		"03": "04",
		"04": "05",
		"05": "06",
		"06": "07",
		"07": "08",
		"08": "09",
		"09": "10",
		"10": "11",
		"11": "12",
		"12": "13",
		"13": "14",
		"14": "15",
		"15": "16",
		"16": "17",
		"17": "18",
		"18": "19",
		"19": "20",
		"20": "21",
		"21": "22",
		"22": "23",
		"23": "24",
		"24": "25",
		"25": "26",
		"26": "27",
		"27": "28",
		"28": "29",
		"29": "30",
		"30": "31",
	}

	return month, next[day]
}

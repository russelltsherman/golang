package main

import (
    "encoding/csv"
    "fmt"
    "log"
    "os"
		"strconv"
		"strings"
)

type Class struct {
	Id int32
	GymId int32
	ClassName string
	ClassCategories []string
	FirstTimeSpecial bool
	RemainingSlots int32
	ClassFreqency string
	InstructorID int32
	InstructorName string
}

func main() {
    // open file
    f, err := os.Open("./GymClasses.csv")
    if err != nil {
        log.Fatal(err)
    }

    // remember to close the file at the end of the program
    defer f.Close()

    // read csv values using csv.Reader
    csvReader := csv.NewReader(f)
    data, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

		dataOut := []Class{}
		monthCount := make(map[string]int32)

		for i := 0; i < len(data); i++{
			row := data[i]

			if i == 0 {
				continue
			}

			v, err := strconv.ParseBool(row[4])
			if err != nil {
				panic(err)
			}
			// strings.Split
			rowOut := Class{
				Id:  mustParseInt(row[0]),
				GymId: mustParseInt(row[1]),
				ClassName: row[2],
				ClassCategories: strings.Split(row[3], ";"),
				FirstTimeSpecial: v,
				RemainingSlots: mustParseInt(row[5]),
				ClassFreqency: row[6],
				InstructorID: mustParseInt(row[7]),
				InstructorName: row[8],
			}

			if row[6] == "monthly" {
				monthCount[row[8]] = monthCount[row[8]] + 1
			}

			fmt.Println(monthCount)
			dataOut = append(dataOut, rowOut)
		}

    // print the array
    // fmt.Printf("%+v\n", data)
		// fmt.Println(dataOut)

		// How many monthly classes does Kate teach?
}

func mustParseInt(a string) (int32) {
	b, err := strconv.ParseInt(a, 10, 32)
	if err != nil {
		panic(err)
	}
	return int32(b)
}

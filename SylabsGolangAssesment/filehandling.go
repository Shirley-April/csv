package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type CsvFields struct{
	ID            string
	UserID        string
	TimeReceived  time.Time
	TimeBegan     time.Time
	TimeFinished  time.Time
	DeletedBuilds bool
	ExitBuildProcess  int
	ImageSize     int

}

var fields []*CsvFields

//BuildRemoteService() and SuccessRate() to do the analyzing and main() to print the output
func main(){
	if err := readStats("stats.csv"); err != nil {
		log.Fatal(err)
	}

	users := BuildRemoteService()
	fmt.Println("\nThese are the top 5 users:")
	fmt.Println(users)
	//for i, userPair := range users {
	//	if i == 5 {
	//		break
	//	}
	//
	//	fmt.Printf("%d. User %s: %d builds.\n", i + 1, userPair.Key, userPair.Value)

	rate := SuccessRate()
	fmt.Printf("\nSuccess Rate: %.2f%%\n", rate)//round off the result to 2 decimal places
	}

//func BuildTimeWindow()
//Go through the builds
//Input:
//Time builds began
//Time builds finished
//Interval of when the builds happen
//Output: No of builds within an interval


//Top 5 users and builds executed in a time window
//Go through the builds
//If user exists add 1 too their value else add a new value which is 1
//return the
func BuildRemoteService ()  int {

	ranking := 0 //store userID also how many times they've used a service
	for _, data := range fields{  //checks if the user is in the list
		if data.UserID == data.UserID{
			fmt.Println(data.UserID)
		}else {
			ranking = 1 //If not give them a new value
		}
	}
	return ranking
}

//Go over the data
//Count the successful and total builds
//Get the success rate
func SuccessRate() float64{
	builds, successful := 0, 0

	for _, data := range fields {

		if data.ExitBuildProcess == 0 { //successful build output 0 count them
			successful++
		}
			//fmt.Printf("%d top exit codes\n", data.ExitBuildProcess)// print exit codes of successful builds

		}
	builds++// count total builds


	rate := (float64(successful)/float64(builds)) * 100 //find the success rates
	return rate
}

func readStats(filename string) error {
//opens the stats.csv file
	file, err := os.Open(filename)
	if err != nil {
		return err
	} //if any errors break and return those errors

//We run the defer function before we exit
	defer func() {
		err = file.Close()
	}()

//contents of the csv go to the lines variable
	lines, err := csv.NewReader(file).ReadAll()
	if err != nil {
		return err
	}

	for _, line := range lines {
        // convert to proper types so that they can be stored in the struct(it has different formats)
		received, _ := time.Parse(time.RFC3339, line[2])
		began, _ := time.Parse(time.RFC3339, line[3])
		finished, _ := time.Parse(time.RFC3339, line[4])
		deleted, _ := strconv.ParseBool(line[5])//converting to a boolean
		process, _ := strconv.Atoi(line[6])//converting alphanumeric to an int
		size, _ := strconv.Atoi(line[7])
//creating a new objects from whatever we are reading from each line and each slice is a particular column
		data := &CsvFields{
			ID:            line[0],
			UserID:        line[1],
			TimeReceived:  received,
			TimeBegan:     began,
			TimeFinished:  finished,
			DeletedBuilds: deleted,
			ExitBuildProcess:  process,
			ImageSize:     size,
		}
//appending fields to the slice
		fields = append(fields, data)
	}

	return nil
}
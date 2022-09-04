package lib

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
	"github.com/gin-gonic/gin"
)

func Booking(c *gin.Context){
	var request RequestBooking
	var response ResponseBooking

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(200, gin.H{"errcode": http.StatusBadRequest, "description": "Post Data Err"})
		return
	}

	err = saveData(request)
	fmt.Println(err)
	if err != nil {
		response.Status = false
		response.Message = err.Error()
		response.Data = request
		c.JSON(200, response)
		return
	}

	response.Status = true
	response.Message = "Success"
	response.Data = request

	c.JSON(200, response)
}

func saveData(in RequestBooking) (err error) {
	var timezone string
	var item CSVRecord
	var value []CSVRecord

	data := readCSV("lib/data.csv")
	id := data[len(data)-1]
	newID := id.ID + 1
	for _, v := range data{
		if v.Name == in.Name && v.Day == in.Day {
			availableAt := convertTimezone(in.AvailableAt, localTime)
			availableUntil := convertTimezone(in.AvailableUntil, localTime)
			start := convertTimezone(strings.TrimSpace(v.AvailableAt), v.Timezone)
			end := convertTimezone(strings.TrimSpace(v.AvailableUntil), v.Timezone)
	
			//check range time
			valid := inTimeSpan(start, end, availableAt, availableUntil)
	
			if !valid{
				err = errors.New("failed when appointment")
				return err
			}
	
			timezone = localTime
		}
	}

	//write to csv
	item.ID = newID
	item.Name = in.Name
	item.Timezone = timezone
	item.Day = in.Day
	item.AvailableAt = in.AvailableAt
	item.AvailableUntil = in.AvailableUntil
	item.Status = 0
	
	value = append(value, item)
	writeData("data.csv", value)

	return nil
}

func convertTimezone(h, timezone string) string {
	local := 7
	zone := timezone[4:7]
	gmt, _ := strconv.Atoi(zone)

	if gmt == local {
		return h
	}else{

		diff := 0
		if gmt > 0 {
			diff = local-gmt
		}else{
			diff = local+gmt
		}
	
		fmt.Println("h", h)
		fmt.Println("diff", diff)
	
		//get different time
		t := time.Date(0, 0, 0, diff, 0, 0, 0, time.UTC)
	  
		// Calling Hour method
		hour, _ := time.Parse(time.Kitchen, h)
		hour = hour.Add(time.Duration(t.Hour()) * time.Hour)
	
		return hour.Format(time.Kitchen)
	}

}

func inTimeSpan(start, end, at, until string) bool {
	startTime, _ := time.Parse(time.Kitchen, start)		
	endTime, _ := time.Parse(time.Kitchen, end)		
	atTime, _ := time.Parse(time.Kitchen, at)		
	untilTime, _ := time.Parse(time.Kitchen, until)

	if startTime.Equal(atTime) {
        return false
    }
	if endTime.Equal(untilTime){
		return false
	}

	return true	
}

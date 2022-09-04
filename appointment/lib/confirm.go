package lib

import (
	"fmt"
	"net/http"
	"sort"
	"strings"
	"github.com/gin-gonic/gin"
)

func Conform(c *gin.Context){
	var request RequestApprove
	var response ResponseApprove
	var newData []CSVRecord

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(200, gin.H{"errcode": http.StatusBadRequest, "description": "Post Data Err"})
		return
	}

	data := readCSV("lib/data.csv")
	arr := filterByID(request.ID, data)
	if arr.ID == 0 {
		response.Status = false
		response.Message = "Data not found"
		response.Data = request
		response.Data.NewSchedule.Name = arr.Name
		c.JSON(200, response)
		return
	}

	if arr.Status > 0 {
		response.Status = false
		response.Message = "Data have been approved"
		response.Data = request
		response.Data.NewSchedule.Name = arr.Name
		c.JSON(200, response)
		return
	}

	for _, v := range data {
		if arr.ID != v.ID {
			newData = append(newData, v)
		}
	}
	//update status data if rejected
	if !request.Approve {
		arr.Status = 3
		newData = append(newData, arr)
		//sorting
		sort.SliceStable(newData, func(i, j int) bool {
			return newData[i].ID < newData[j].ID
		})

		//store to csv
		updateData("data.csv", newData)

		response.Status = false
		response.Message = "Appointment rejected!"
		response.Data = request
		response.Data.NewSchedule.Name = arr.Name
		c.JSON(200, response)
		return
	}
	
	if request.Reschedule {
		//check range time
		for _, v := range data {
			if v.Name == request.NewSchedule.Name && v.Day == request.NewSchedule.Day {
				availableAt := convertTimezone(request.NewSchedule.AvailableAt, localTime)
				availableUntil := convertTimezone(request.NewSchedule.AvailableUntil, localTime)
				start := convertTimezone(strings.TrimSpace(v.AvailableAt), v.Timezone)
				end := convertTimezone(strings.TrimSpace(v.AvailableUntil), v.Timezone)
		
				//check range time
				valid := inTimeSpan(start, end, availableAt, availableUntil)
		
				if !valid{
					response.Status = false
					response.Message = "failed when reschedule, please choose another time"
					c.JSON(200, response)
					return
				}		
			}
		}

		arr.Status = 2
		arr.Day = request.NewSchedule.Day
		arr.AvailableAt = request.NewSchedule.AvailableAt
		arr.AvailableUntil = request.NewSchedule.AvailableUntil

		newData = append(newData, arr)
		//sorting
		sort.SliceStable(newData, func(i, j int) bool {
			return newData[i].ID < newData[j].ID
		})
		//store to csv
		updateData("data.csv", newData)			

		response.Status = true
		response.Message = "success confirmation appointment"
		response.Data = request
		response.Data.NewSchedule.Name = arr.Name
		c.JSON(200, response)
		return			
	}

	arr.Status = 1
	newData = append(newData, arr)
	//sorting
	sort.SliceStable(newData, func(i, j int) bool {
		return newData[i].ID < newData[j].ID
	})

	//store to csv
	updateData("data.csv", newData)			

	response.Status = true
	response.Message = "success confirmation appointment"
	response.Data = request
	response.Data.NewSchedule.Name = arr.Name
	c.JSON(200, response)

}

func filterByID(id int, list []CSVRecord) CSVRecord{
	for _, v := range list{
		if v.ID == id {
			return v
		}
	}
	return CSVRecord{}
}

func UserConform(c *gin.Context){
	var request RequestUserApproved
	var response ResponseApprove
	var newData []CSVRecord

	err := c.BindJSON(&request)
	if err != nil {
		c.JSON(200, gin.H{"errcode": http.StatusBadRequest, "description": "Post Data Err"})
		return
	}

	data := readCSV("lib/data.csv")
	
	arr := filterByID(request.ID, data)

	//response
	response.Status = true
	response.Message = "Success confirmation reschedule"
	response.Data.Approve = true
	response.Data.Reschedule = true
	response.Data.ID = arr.ID
	response.Data.NewSchedule.AvailableAt = arr.AvailableAt
	response.Data.NewSchedule.AvailableUntil = arr.AvailableUntil
	response.Data.NewSchedule.Day = arr.Day
	response.Data.NewSchedule.Name = arr.Name

	for _, v := range data {
		if arr.ID != v.ID {
			newData = append(newData, v)
		}
	}
	fmt.Println("arr", arr)
	if arr.ID == 0 || arr.Status != 2{
		response.Status = false
		response.Message = "Data not found"

		c.JSON(200, response)
		return
	}

	if !request.Approved{
		arr.Status = 3
		newData = append(newData, arr)
		sort.SliceStable(newData, func(i, j int) bool {
			return newData[i].ID < newData[j].ID
		})

		//store to csv
		updateData("data.csv", newData)			


		c.JSON(200, response)
		return
	}
	
	arr.Status = 1
	newData = append(newData, arr)
	sort.SliceStable(newData, func(i, j int) bool {
		return newData[i].ID < newData[j].ID
	})

	//store to csv
	updateData("data.csv", newData)			


	c.JSON(200, response)
}


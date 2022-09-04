package lib

import (
	"encoding/csv"
	"log"
	"os"
	"github.com/gin-gonic/gin"
    "strconv"
)

func ReadData(c *gin.Context){
    list := readCSV("lib/data.csv")

	c.JSON(200, list)
}

func readCSV(in string) []CSVRecord {
    var list []CSVRecord

	f, err := os.Open(in)
    if err != nil {
        log.Fatal(err)
    }

    // remember to close the file at the end of the program
    defer f.Close()

    // read csv values using csv.Reader
    csvReader := csv.NewReader(f)
    datas, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

    for i, line := range datas {
        if i > 0 { // omit header line
            var rec CSVRecord
            for j, field := range line {
                if j == 0 {
                    rec.ID, _ = strconv.Atoi(field)
                } else if j == 1 {
                    rec.Name = field
                } else if j == 2 {
                    rec.Timezone = field
                }else if j == 3 {
					rec.Day = field
				}else if j == 4 {
					rec.AvailableAt = field
				}else if j == 5 {
					rec.AvailableUntil = field
				}else if j == 6 {
                    rec.Status, _ = strconv.Atoi(field)
				}
            }
            list = append(list, rec)
        }
    }
    return list
}

func writeData(file string, in []CSVRecord) {
    path := "lib/" + file
    f, err := os.OpenFile(path, os.O_CREATE|os.O_APPEND, os.ModePerm)
    defer f.Close()
    if err != nil {
        log.Fatalln("failed to open file", err)
    }
    w := csv.NewWriter(f)
    defer w.Flush()
    
    var data [][]string
    for _, v := range in {
        row := []string{strconv.Itoa(v.ID), v.Name, v.Timezone, v.Day, v.AvailableAt, v.AvailableUntil, strconv.Itoa(v.Status),}
        data = append(data, row)
    }
    w.WriteAll(data)
}

func updateData(file string, in []CSVRecord){
    path := "lib/" + file
    // delete file
    err := os.Remove(path)
    if err != nil {
        return
    }

    // check if file exists
    f, err := os.OpenFile(path,os.O_CREATE|os.O_APPEND, os.ModePerm)
    defer f.Close()
    if err != nil {
        log.Fatalln("failed to open file", err)
    }
    w := csv.NewWriter(f)
    defer w.Flush()
    
    headerRow := []string{
        "ID","Name","Timezone","Day","AvailableAt","AvailableUntil","Status",
    }
    var data [][]string
    data = append(data, headerRow)
    for _, v := range in {
        row := []string{strconv.Itoa(v.ID), v.Name, v.Timezone, v.Day, v.AvailableAt, v.AvailableUntil, strconv.Itoa(v.Status),}
        data = append(data, row)
    }
    w.WriteAll(data)
}


package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// myJobHandler fetch json resources from remote server and write it into csv file.
func myJobHandler(job Job) {
	res, err := http.Get(job.URL)
	if err != nil {
		log.Println(err)
		return
	}

	defer res.Body.Close()

	// decoding json response into defined models.
	var models []map[string]interface{}
	if err := json.NewDecoder(res.Body).Decode(&models); err != nil {
		log.Println(err)
		return
	}

	// create/open csv file.
	file, err := os.OpenFile(job.Dir+job.Filename, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Println(err)
		return
	}

	var contents [][]string
	columns := getColumns(models)

	// add csv header
	contents = append(contents, columns)

	// append csv body.
	for _, row := range models {
		values := make([]string, 0)

		for _, col := range columns {
			val := row[col]
			switch v := val.(type) {
			case int:
				values = append(values, strconv.Itoa(v))
			case float64:
				values = append(values, strconv.Itoa(int(v)))
			case string:
				// comma in csv is reserved.
				values = append(values, strings.ReplaceAll(v, ",", ""))
			default:
				log.Printf("unsupported type %T for value %v\n", v, v)
			}
		}

		contents = append(contents, values)
	}

	writer := csv.NewWriter(file)
	writer.WriteAll(contents)

	if err := writer.Error(); err != nil {
		log.Println("error writing csv:", err)
	}
}

func getColumns(models []map[string]interface{}) []string {
	var columns []string
	for column := range models[0] {
		columns = append(columns, column)
	}

	return columns
}

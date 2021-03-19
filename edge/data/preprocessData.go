package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	pwd, _ := os.Getwd()
	path := pwd + "/edge/data/dataset"
	writeF, err := os.OpenFile(path+"/data.csv", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	writer := csv.NewWriter(writeF)

	files, _ := ioutil.ReadDir(path)

	for _, f := range files {
		fmt.Println(f.Name())
		if strings.Contains(f.Name(), ".csv") && strings.Contains(f.Name(), "_") {
			readF, err := os.Open(path + "/" + f.Name())
			if err != nil {
				fmt.Println("Error: ", err)
			}
			reader := csv.NewReader(readF)
			reader.Comma = '\t'
			reader.FieldsPerRecord = -1
			bearType := extractBearType(f.Name())
			for index := 0; ; index++ {
				line, err := reader.Read()
				if err == io.EOF {
					break
				} else if err != nil {
					fmt.Println("Error: ", err)
					return
				}
				if index < 17 || len(line) < 8 {
					continue
				}
				line[len(line)-1] = bearType
				writer.Write(line)
			}
		}
	}
	writer.Flush()
	if err = writer.Error(); err != nil {
		fmt.Println(err)
	}
}

func extractBearType(name string) string {
	var bearType string
	bearType = strings.Split(name, "_")[0]
	return bearType
}

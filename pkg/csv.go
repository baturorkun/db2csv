package pkg

import (
	"encoding/csv"
	"log"
	"os"
	"time"
)

func PrepareData(cols Cols, data Data) (res []Cols) {

	//return append([]cols, data...)
	var prepend []Cols

	prepend = append(prepend, cols)

	return append(prepend, data...)

}

func CreateCSVFile(data []Cols) {

	var filename string
	dt := time.Now().Format("01-02-2006 15:04:05")

	if AppSetting.OverwriteFile {
		filename = AppSetting.FilesPath + "/" + AppSetting.Filename + ".csv"
	} else {
		filename = AppSetting.FilesPath + "/" +  AppSetting.Filename + "_" + dt + ".csv"
	}

	file, err := os.Create(filename)
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

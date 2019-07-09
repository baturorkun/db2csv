package main

import (
	pkg "db2csv/pkg"
)

//import "log"

func main() {

	pkg.SetupINI()
	pkg.SetupDB()

	cols, data := pkg.Populate()

	csvDataArray := pkg.PrepareData(cols, data)

	pkg.CreateCSVFile(csvDataArray)

}
package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
)


func test1() {
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var  err error

	file = xlsx.NewFile()
	sheet,_= file.AddSheet("sheet1")
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "000101"
	cell = row.AddCell()
	cell.Value = "中文"

	err = file.Save("MyXLSFile.xlsx")
	if err !=nil {
		fmt.Println(err.Error())
	}
}

func test2(){
	var file *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell
	var err error

	file,_ = xlsx.OpenFile("MyXLSFile.xlsx")
	sheet = file.Sheet["sheet1"]
	row = sheet.AddRow()
	cell = row.AddCell()
	cell.Value = "000101"
	cell = row.AddCell()
	cell.Value = "中文歌"
	err = file.Save("MyXLSFile.xlsx")
	if err != nil{
		fmt.Printf(err.Error())
	}
}

func main(){
	test1()
	test2()
	fmt.Println("kkkkkkk")
}
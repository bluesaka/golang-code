/**
go get -u github.com/360EntSecGroup-Skylar/excelize/v2
@link https://www.bookstack.cn/read/excelize-2.2-zh/b2e09ee3ee36ed31.md
*/

package myexcelize

import (
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"log"
	"math/rand"

	// 引入image下的包，否则excel无法添加png等图片
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

// WriteToExcel write data to excel
func WriteToExcel() {
	file := excelize.NewFile()

	// 创建一个 sheet
	index := file.NewSheet("Sheet2")

	// 向单元格中写入值
	_ = file.SetCellValue("Sheet2", "A2", "porsche")
	_ = file.SetCellValue("Sheet1", "B2", 100)

	// 设置文件打开后显示哪个 sheet， 0 表示 sheet1
	file.SetActiveSheet(index)

	// 保存到文件
	if err := file.SaveAs("../file/car.xlsx"); err != nil {
		log.Println("excel save error", err)
	}
}

// ReadFromExcel read data from excel
func ReadFromExcel() {
	//打开文件
	file, err := excelize.OpenFile("../file/car.xlsx")
	if err != nil {
		panic(err)
	}

	// 获取单元格内容
	cell, err := file.GetCellValue("Sheet2", "A2")
	if err != nil {
		log.Println("get cell value error:", err)
	} else {
		log.Println("Sheet2 A2 cell value:", cell)
	}

	cell, err = file.GetCellValue("Sheet1", "A2")
	if err != nil {
		log.Println("get cell value error:", err)
	} else {
		log.Println("Sheet1 A2 cell value:", cell)
	}

	// 获取sheet1中所有的行
	rows, err := file.GetRows("Sheet1")
	if err != nil {
		log.Println("get rows error:", err)
	} else {
		for _, row := range rows {
			for _, cell := range row {
				log.Println("cell value is", cell)
			}
		}
	}
}

// WriteChartToExcel write chart to excel
func WriteChartToExcel() {
	categories := map[string]string{
		"A2": "Small", "A3": "Normal", "A4": "Large",
		"B1": "Apple", "C1": "Orange", "D1": "Pear",
	}
	values := map[string]int{
		"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8,
	}

	file := excelize.NewFile()
	for k, v := range categories {
		file.SetCellValue("Sheet1", k, v)
	}
	for k, v := range values {
		file.SetCellValue("Sheet1", k, v)
	}

	// 添加图表
	if err := file.AddChart("Sheet1", "E1", `{
		"type": "col3DClustered",
		"series": [{
			"name": "Sheet1!$A$2",
			"categories": "Sheet1!$B$1:$D$1",
			"values":"Sheet1!$B$2:$D$2"
		},
        {
            "name": "Sheet1!$A$3",
            "categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$B$3:$D$3"
        },
        {
            "name": "Sheet1!$A$4",
            "categories": "Sheet1!$B$1:$D$1",
            "values": "Sheet1!$B$4:$D$4"
        }],
		"title": {
			"name": "Fruit 3D Chart"
		}
	}`); err != nil {
		log.Println("save chart to excel error:", err)
		return
	}

	if err := file.SaveAs("../file/fruitChart.xlsx"); err != nil {
		log.Println("save excel error:", err)
	}
}

// WriteImageToExcel write image to excel
func WriteImageToExcel() {
	file := excelize.NewFile()

	if err := file.AddPicture("Sheet1", "A2", "../file/sea.png", ""); err != nil {
		log.Println("save image to excel error:", err)
	}

	if err := file.AddPicture("Sheet1", "A2", "../file/sea.png", `{"x_scale": 0.5, "y_scale": 0.5}`); err != nil {
		log.Println("save image to excel error:", err)
	}

	if err := file.SaveAs("../file/pic.xlsx"); err != nil {
		log.Println("save excel error:", err)
	}
}

// StreamWrite 流式写入exel
// 适合大数据写入
func StreamWrite() {
	file := excelize.NewFile()
	streamWriter, err := file.NewStreamWriter("Sheet1")
	if err != nil {
		panic(err)
	}

	styleID, err := file.NewStyle(`{"font":{"color":"#777777"}}`)
	if err != nil {
		log.Println("new style error:", err)
	}

	if err := streamWriter.SetRow("A1", []interface{}{excelize.Cell{
		StyleID: styleID,
		Value:   "data a1",
	}}); err != nil {
		log.Println("set row data error:", err)
	}

	for rowID := 2; rowID <= 10; rowID++ {
		row := make([]interface{}, 5)
		for colID := 0; colID < 5; colID++ {
			row[colID] = rand.Intn(100000)
		}
		cell, _ := excelize.CoordinatesToCellName(1, rowID)
		if err := streamWriter.SetRow(cell, row); err != nil {
			log.Println("set row error", err)
		}
	}

	if err := streamWriter.Flush(); err != nil {
		log.Println("stream flush error", err)
	}

	if err := file.SaveAs("../file/stream.xlsx"); err != nil {
		log.Println("save excel error:", err)
	}
}

package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"net/http"
	"strconv"
)

func Formato1011(w http.ResponseWriter, r *http.Request) {
	listaformato := []miFormatoExcel{}
	listaformato = generarformato("1011", false)
	f := excelize.NewFile()

	if err := f.SetColWidth("Sheet1", "A", "A", 6); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "B", "B", 10); err != nil {
		fmt.Println(err)
		return
	}
	// titulos
	f.SetCellValue("Sheet1", "A1", "CONCEPTO")
	f.SetCellValue("Sheet1", "B1", "COLUMNA1")

	var filaExcel = 2
	//var a string
	//	var a = ""
	//var van int
	estiloTexto, err := f.NewStyle(`{"font":{"bold":false,"italic":false,"family":"Arial","size":8,"color":"##000000"}}`)

	estiloNumeroDetalle, err := f.NewStyle(`{"number_format": 1,"font":{"bold":false,"italic":false,"family":"Arial","size":8,"color":"##000000"}}`)

	if err != nil {
		fmt.Println(err)
	}

	var saldo float64

	listaconcepto := []FormatoConcepto{}
	var consulta = "select distinct concepto,formato from exogena where formato=$1"
	err2 := db.Select(&listaconcepto, consulta, "1011")
	if err2 != nil {
		fmt.Println(err2)
	}

	for i, miConcepto := range listaconcepto {
		fmt.Println("Concepto")
		saldo = 0
		for _, miFila := range listaformato {
			if miFila.Concepto == miConcepto.Concepto {
				saldo += Flotante(miFila.Columna1)
			}
		}
		var a = strconv.Itoa(filaExcel + i)
		f.SetCellValue("Sheet1", "A"+a, miConcepto.Concepto)
		f.SetCellValue("Sheet1", "B"+a, saldo)

		f.SetCellStyle("Sheet1", "A"+a, "A"+a, estiloTexto)
		f.SetCellStyle("Sheet1", "B"+a, "B"+a, estiloNumeroDetalle)
	}

	// Set the headers necessary to get browsers to interpret the downloadable file
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment;filename=userInputData.xlsx")
	w.Header().Set("File-Name", "userInputData.xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	err = f.Write(w)
	if err != nil {
		panic(err.Error())
	}

}

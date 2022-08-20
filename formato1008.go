package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"net/http"
	"strconv"
)

func Formato1008(w http.ResponseWriter, r *http.Request) {
	listaformato := []miFormatoExcel{}
	listaformato = generarformato("1008", false)
	f := excelize.NewFile()

	if err := f.SetColWidth("Sheet1", "A", "A", 4); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.SetColWidth("Sheet1", "B", "B", 10); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "C", "C", 4); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.SetColWidth("Sheet1", "D", "D", 10); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.SetColWidth("Sheet1", "E", "E", 10); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.SetColWidth("Sheet1", "F", "F", 10); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.SetColWidth("Sheet1", "G", "G", 10); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.SetColWidth("Sheet1", "H", "H", 10); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.SetColWidth("Sheet1", "I", "I", 10); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.SetColWidth("Sheet1", "J", "J", 10); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "K", "K", 10); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "L", "L", 10); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "M", "M", 10); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "N", "N", 10); err != nil {
		fmt.Println(err)
		return
	}
	// titulos
	f.SetCellValue("Sheet1", "A1", "CONCEPTO")
	f.SetCellValue("Sheet1", "B1", "TIPO")
	f.SetCellValue("Sheet1", "C1", "DOCUMENTO")
	f.SetCellValue("Sheet1", "D1", "DV")
	f.SetCellValue("Sheet1", "E1", "PRIMER APELLIDO")
	f.SetCellValue("Sheet1", "F1", "SEGUNDO APELLIDO")
	f.SetCellValue("Sheet1", "G1", "PRIMER NOMBRE")
	f.SetCellValue("Sheet1", "H1", "SEGUNDO NOMBRE")
	f.SetCellValue("Sheet1", "I1", "RAZON SOCIAL")
	f.SetCellValue("Sheet1", "J1", "DIRECCION")
	f.SetCellValue("Sheet1", "K1", "DEPARTAMENTO")
	f.SetCellValue("Sheet1", "L1", "CIUDAD")
	f.SetCellValue("Sheet1", "M1", "PAIS")
	f.SetCellValue("Sheet1", "N1", "COLUMNA1")

	var filaExcel = 2
	//var a string
	//	var a = ""
	//var van int
	estiloTexto, err := f.NewStyle(`{"font":{"bold":false,"italic":false,"family":"Arial","size":8,"color":"##000000"}}`)

	estiloNumeroDetalle, err := f.NewStyle(`{"number_format": 1,"font":{"bold":false,"italic":false,"family":"Arial","size":8,"color":"##000000"}}`)

	if err != nil {
		fmt.Println(err)
	}

	for i, miFila := range listaformato {
		var a = strconv.Itoa(filaExcel + i)
		f.SetCellValue("Sheet1", "A"+a, miFila.Concepto)
		f.SetCellValue("Sheet1", "B"+a, miFila.Documento)
		f.SetCellValue("Sheet1", "C"+a, miFila.Codigo)
		f.SetCellValue("Sheet1", "D"+a, miFila.Dv)
		f.SetCellValue("Sheet1", "E"+a, miFila.PrimerApllido)
		f.SetCellValue("Sheet1", "F"+a, miFila.SegundoApellido)
		f.SetCellValue("Sheet1", "G"+a, miFila.PrimerNombre)
		f.SetCellValue("Sheet1", "H"+a, miFila.SegundoNombre)
		f.SetCellValue("Sheet1", "I"+a, miFila.Juridica)
		f.SetCellValue("Sheet1", "J"+a, miFila.Direccion)
		f.SetCellValue("Sheet1", "K"+a, miFila.Departamento)
		f.SetCellValue("Sheet1", "L"+a, miFila.Ciudad)
		f.SetCellValue("Sheet1", "M"+a, miFila.Pais)
		f.SetCellValue("Sheet1", "N"+a, Flotante(miFila.Columna1))

		f.SetCellStyle("Sheet1", "A"+a, "A"+a, estiloTexto)
		f.SetCellStyle("Sheet1", "B"+a, "B"+a, estiloTexto)
		f.SetCellStyle("Sheet1", "C"+a, "C"+a, estiloTexto)
		f.SetCellStyle("Sheet1", "D"+a, "D"+a, estiloTexto)
		f.SetCellStyle("Sheet1", "E"+a, "E"+a, estiloTexto)
		f.SetCellStyle("Sheet1", "F"+a, "F"+a, estiloTexto)
		f.SetCellStyle("Sheet1", "G"+a, "G"+a, estiloTexto)
		f.SetCellStyle("Sheet1", "H"+a, "H"+a, estiloTexto)
		f.SetCellStyle("Sheet1", "I"+a, "I"+a, estiloTexto)
		f.SetCellStyle("Sheet1", "J"+a, "J"+a, estiloTexto)
		f.SetCellStyle("Sheet1", "K"+a, "K"+a, estiloTexto)
		f.SetCellStyle("Sheet1", "L"+a, "L"+a, estiloTexto)
		f.SetCellStyle("Sheet1", "M"+a, "M"+a, estiloTexto)
		f.SetCellStyle("Sheet1", "N"+a, "N"+a, estiloNumeroDetalle)

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

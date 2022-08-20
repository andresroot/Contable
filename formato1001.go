package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"net/http"
	"strconv"
)

func Formato1001(w http.ResponseWriter, r *http.Request) {
	listaformato := []miFormatoExcel{}
	listaformato = generarformato("1001", true)
	f := excelize.NewFile()
	if err := f.SetColWidth("Sheet1", "A", "A", 6); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "B", "B", 4); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.SetColWidth("Sheet1", "C", "C", 10); err != nil {
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

	// titulos

	f.SetCellValue("Sheet1", "A1", "CONCEPTO")
	f.SetCellValue("Sheet1", "B1", "TIPO")
	f.SetCellValue("Sheet1", "C1", "DOCUMENTO")
	f.SetCellValue("Sheet1", "D1", "PRIMER APELLIDO")
	f.SetCellValue("Sheet1", "E1", "SEGUNDO APELLIDO")
	f.SetCellValue("Sheet1", "F1", "PRIMER NOMBRE")
	f.SetCellValue("Sheet1", "G1", "SEGUNDO NOMBRE")
	f.SetCellValue("Sheet1", "H1", "RAZON SOCIAL")
	f.SetCellValue("Sheet1", "I1", "DIRECCION")
	f.SetCellValue("Sheet1", "J1", "DEPARTAMENTO")
	f.SetCellValue("Sheet1", "K1", "CIUDAD")
	f.SetCellValue("Sheet1", "L1", "PAIS")
	f.SetCellValue("Sheet1", "M1", "COLUMNA1")
	f.SetCellValue("Sheet1", "N1", "COLUMNA2")
	f.SetCellValue("Sheet1", "O1", "COLUMNA3")
	f.SetCellValue("Sheet1", "P1", "COLUMNA4")
	f.SetCellValue("Sheet1", "Q1", "COLUMNA5")
	f.SetCellValue("Sheet1", "R1", "COLUMNA6")
	f.SetCellValue("Sheet1", "S1", "COLUMNA7")
	f.SetCellValue("Sheet1", "T1", "COLUMNA8")

	// FILA QUE INICIA EL CONTENIDO
	var filaExcel = 2

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
		f.SetCellValue("Sheet1", "D"+a, miFila.PrimerApllido)
		f.SetCellValue("Sheet1", "E"+a, miFila.SegundoApellido)
		f.SetCellValue("Sheet1", "F"+a, miFila.PrimerNombre)
		f.SetCellValue("Sheet1", "G"+a, miFila.SegundoNombre)
		f.SetCellValue("Sheet1", "H"+a, miFila.Juridica)
		f.SetCellValue("Sheet1", "I"+a, miFila.Direccion)
		f.SetCellValue("Sheet1", "J"+a, miFila.Departamento)
		f.SetCellValue("Sheet1", "K"+a, miFila.Ciudad)
		f.SetCellValue("Sheet1", "L"+a, miFila.Pais)
		f.SetCellValue("Sheet1", "M"+a, Flotante(miFila.Columna1))
		f.SetCellValue("Sheet1", "N"+a, Flotante(miFila.Columna2))
		f.SetCellValue("Sheet1", "O"+a, Flotante(miFila.Columna3))
		f.SetCellValue("Sheet1", "P"+a, Flotante(miFila.Columna4))
		f.SetCellValue("Sheet1", "Q"+a, Flotante(miFila.Columna5))
		f.SetCellValue("Sheet1", "R"+a, Flotante(miFila.Columna6))
		f.SetCellValue("Sheet1", "S"+a, Flotante(miFila.Columna7))
		f.SetCellValue("Sheet1", "T"+a, Flotante(miFila.Columna8))

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
		f.SetCellStyle("Sheet1", "M"+a, "M"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "N"+a, "N"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "O"+a, "O"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "P"+a, "P"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "Q"+a, "Q"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "R"+a, "R"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "S"+a, "S"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "T"+a, "T"+a, estiloNumeroDetalle)
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

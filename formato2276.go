package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"net/http"
	"strconv"
)

func Formato2276(w http.ResponseWriter, r *http.Request) {
	listaformato := []miFormatoExcel{}
	listaformato = generarformato("2276", false)

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
	if err := f.SetColWidth("Sheet1", "M", "M", 10); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.SetColWidth("Sheet1", "N", "N", 10); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "O", "O", 10); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "P", "P", 10); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "Q", "Q", 10); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "R", "R", 10); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "S", "S", 10); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "T", "T", 10); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "U", "U", 10); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "V", "V", 10); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "W", "W", 10); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "X", "X", 10); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "Y", "Y", 10); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "Z", "Z", 10); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "AA", "AA", 10); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "AB", "AB", 10); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "AC", "AC", 10); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "AD", "AD", 10); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "AE", "AE", 10); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "AF", "AF", 10); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "AG", "AG", 10); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "AH", "AH", 10); err != nil {
		fmt.Println(err)
		return
	}

	// titulos

	f.SetCellValue("Sheet1", "A1", "ENTIDAD")
	f.SetCellValue("Sheet1", "B1", "TIPO")
	f.SetCellValue("Sheet1", "C1", "DOCUMENTO")
	f.SetCellValue("Sheet1", "D1", "PRIMER APELLIDO")
	f.SetCellValue("Sheet1", "E1", "SEGUNDO APELLIDO")
	f.SetCellValue("Sheet1", "F1", "PRIMER NOMBRE")
	f.SetCellValue("Sheet1", "G1", "SEGUNDO NOMBRE")
	f.SetCellValue("Sheet1", "H1", "DIRECCION")
	f.SetCellValue("Sheet1", "I1", "DEPARTAMENTO")
	f.SetCellValue("Sheet1", "J1", "CIUDAD")
	f.SetCellValue("Sheet1", "K1", "PAIS")
	f.SetCellValue("Sheet1", "L1", "COLUMNA1")
	f.SetCellValue("Sheet1", "M1", "COLUMNA2")
	f.SetCellValue("Sheet1", "N1", "COLUMNA3")
	f.SetCellValue("Sheet1", "O1", "COLUMNA4")
	f.SetCellValue("Sheet1", "P1", "COLUMNA5")
	f.SetCellValue("Sheet1", "Q1", "COLUMNA6")
	f.SetCellValue("Sheet1", "R1", "COLUMNA7")
	f.SetCellValue("Sheet1", "S1", "COLUMNA8")
	f.SetCellValue("Sheet1", "T1", "COLUMNA9")
	f.SetCellValue("Sheet1", "U1", "COLUMNA10")
	f.SetCellValue("Sheet1", "V1", "COLUMNA11")
	f.SetCellValue("Sheet1", "W1", "COLUMNA12")
	f.SetCellValue("Sheet1", "X1", "COLUMNA13")
	f.SetCellValue("Sheet1", "Y1", "COLUMNA14")
	f.SetCellValue("Sheet1", "Z1", "COLUMNA15")
	f.SetCellValue("Sheet1", "AA1", "COLUMNA16")
	f.SetCellValue("Sheet1", "AB1", "COLUMNA17")
	f.SetCellValue("Sheet1", "AC1", "COLUMNA18")
	f.SetCellValue("Sheet1", "AD1", "COLUMNA19")
	f.SetCellValue("Sheet1", "AE1", "COLUMNA20")
	f.SetCellValue("Sheet1", "AF1", "COLUMNA21")
	f.SetCellValue("Sheet1", "AG1", "COLUMNA22")
	f.SetCellValue("Sheet1", "AH1", "COLUMNA23")

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
		f.SetCellValue("Sheet1", "A"+a, 1)
		f.SetCellValue("Sheet1", "B"+a, miFila.Documento)
		f.SetCellValue("Sheet1", "C"+a, miFila.Codigo)
		f.SetCellValue("Sheet1", "D"+a, miFila.PrimerApllido)
		f.SetCellValue("Sheet1", "E"+a, miFila.SegundoApellido)
		f.SetCellValue("Sheet1", "F"+a, miFila.PrimerNombre)
		f.SetCellValue("Sheet1", "G"+a, miFila.SegundoNombre)
		f.SetCellValue("Sheet1", "H"+a, miFila.Direccion)
		f.SetCellValue("Sheet1", "I"+a, miFila.Departamento)
		f.SetCellValue("Sheet1", "J"+a, miFila.Ciudad)
		f.SetCellValue("Sheet1", "K"+a, miFila.Pais)
		f.SetCellValue("Sheet1", "L"+a, Flotante(miFila.Columna1))
		f.SetCellValue("Sheet1", "M"+a, Flotante(miFila.Columna2))
		f.SetCellValue("Sheet1", "N"+a, Flotante(miFila.Columna3))
		f.SetCellValue("Sheet1", "O"+a, Flotante(miFila.Columna4))
		f.SetCellValue("Sheet1", "P"+a, Flotante(miFila.Columna5))
		f.SetCellValue("Sheet1", "Q"+a, Flotante(miFila.Columna6))
		f.SetCellValue("Sheet1", "R"+a, Flotante(miFila.Columna7))
		f.SetCellValue("Sheet1", "S"+a, Flotante(miFila.Columna8))
		f.SetCellValue("Sheet1", "T"+a, Flotante(miFila.Columna9))
		f.SetCellValue("Sheet1", "U"+a, Flotante(miFila.Columna10))
		f.SetCellValue("Sheet1", "V"+a, Flotante(miFila.Columna11))
		f.SetCellValue("Sheet1", "W"+a, Flotante(miFila.Columna12))
		f.SetCellValue("Sheet1", "X"+a, Flotante(miFila.Columna13))
		f.SetCellValue("Sheet1", "Y"+a, Flotante(miFila.Columna14))
		f.SetCellValue("Sheet1", "Z"+a, Flotante(miFila.Columna15))
		f.SetCellValue("Sheet1", "AA"+a, Flotante(miFila.Columna16))
		f.SetCellValue("Sheet1", "AB"+a, Flotante(miFila.Columna17))
		f.SetCellValue("Sheet1", "AC"+a, Flotante(miFila.Columna18))
		f.SetCellValue("Sheet1", "AD"+a, Flotante(miFila.Columna19))
		f.SetCellValue("Sheet1", "AE"+a, Flotante(miFila.Columna20))
		f.SetCellValue("Sheet1", "AF"+a, Flotante(miFila.Columna21))
		f.SetCellValue("Sheet1", "AG"+a, Flotante(miFila.Columna22))
		f.SetCellValue("Sheet1", "AH"+a, Flotante(miFila.Columna23))

		f.SetCellStyle("Sheet1", "A"+a, "A"+a, estiloNumeroDetalle)
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
		f.SetCellStyle("Sheet1", "M"+a, "L"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "N"+a, "M"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "O"+a, "N"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "P"+a, "O"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "Q"+a, "P"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "R"+a, "Q"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "S"+a, "R"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "T"+a, "S"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "U"+a, "T"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "V"+a, "U"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "W"+a, "V"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "X"+a, "W"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "Y"+a, "X"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "Z"+a, "Y"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "AA"+a, "Z"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "AB"+a, "AA"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "AC"+a, "AB"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "AD"+a, "AC"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "AE"+a, "AD"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "AF"+a, "AE"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "AG"+a, "AF"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "AH"+a, "AG"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "AI"+a, "AH"+a, estiloNumeroDetalle)
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

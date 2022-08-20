package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
)

// CENTRO TABLA
type plandecuentaniif struct {
	Codigo      string
	Nombre      string
	Nivel       string
}

func PlandecuentaniifListaBuscar(w http.ResponseWriter, r *http.Request) {
	panel := mux.Vars(r)["panel"]
	codigo := mux.Vars(r)["codigo"]
	elemento := mux.Vars(r)["elemento"]

	db := dbConn()
	selDB, err := db.Query("SELECT * FROM plandecuentaniif ORDER BY codigo ASC")
	if err != nil {
		panic(err.Error())
	}
	res := []plandecuentaniif{}
	for selDB.Next() {
		var Codigo string
		var Nombre string
		var Nivel string
		err = selDB.Scan(&Codigo, &Nombre, &Nivel)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, plandecuentaniif{Codigo, Nombre, Nivel })
	}
	parametros := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
		"panel":panel,
		"codigo":codigo,
		"elemento":elemento,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/plandecuentaniif/plandecuentaniifListaBuscar.html")

	//fmt.Printf("%v, %v", t, err)
	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error compra nuevo 3")
	miTemplate.Execute(w, parametros)

	//tmp.Execute(w, varmap)
}
// CENTRO LISTA
func PlandecuentaniifLista(w http.ResponseWriter, r *http.Request) {
	panel := mux.Vars(r)["panel"]
	codigo := mux.Vars(r)["codigo"]
	elemento := mux.Vars(r)["elemento"]

	db := dbConn()
	selDB, err := db.Query("SELECT * FROM plandecuentaniif ORDER BY codigo ASC")
	if err != nil {
		panic(err.Error())
	}
	res := []plandecuentaniif{}
	for selDB.Next() {
		var Codigo string
		var Nombre string
		var Nivel string
		err = selDB.Scan(&Codigo, &Nombre, &Nivel)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, plandecuentaniif{Codigo, Nombre, Nivel })
	}
	parametros := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
		"panel":panel,
		"codigo":codigo,
		"elemento":elemento,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/plandecuentaniif/plandecuentaniifLista.html")

	//fmt.Printf("%v, %v", t, err)
	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error compra nuevo 3")
	miTemplate.Execute(w, parametros)

	//tmp.Execute(w, varmap)
}

// INICIA CENTRO PDF
func PlandecuentaniifPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	t := plandecuentaniif{}
	var e  empresa=ListaEmpresa()
	var c  ciudad=TraerCiudad(e.Ciudad)
	err := db.Get(&t, "SELECT * FROM plandecuentaniif where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	var buf bytes.Buffer
	var err1 error
	pdf := gofpdf.New("P", "mm", "Letter", cnFontDir)
	ene := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetHeaderFunc(func() {
		pdf.Image(imageFile("logo.png"), 20, 20, 40, 0, false,
			"", 0, "")
		pdf.SetY(15)
		//pdf.AddFont("Helvetica", "", "cp1251.map")
		pdf.SetFont("Helvetica", "", 10)
		pdf.CellFormat(190, 10, e.Nombre, "0", 0,
			"C", false, 0, "")
		pdf.Ln(4)

		pdf.CellFormat(190, 10, "Nit No. " +Coma(e.Codigo)+ " - "+e.Dv, "0", 0, "C",
			false, 0, "")
		pdf.Ln(4)
		pdf.CellFormat(190, 10, e.Iva+ " - "+e.ReteIva, "0", 0, "C", false, 0,
			"")
		pdf.Ln(4)
		pdf.CellFormat(190, 10, "Actividad Ica - "+e.ActividadIca, "0",
			0, "C", false, 0, "")
		pdf.Ln(4)
		pdf.CellFormat(190, 10, e.Direccion, "0", 0, "C", false, 0,
			"")
		pdf.Ln(4)
		log.Println("tercero 3")
		pdf.CellFormat(190, 10, ene(c.NombreCiudad+ " - "+c.NombreDepartamento), "0", 0, "C", false, 0,
			"")
		log.Println("tercero 4")
		pdf.Ln(10)
		pdf.CellFormat(190, 10, "Datos Centro de Costos", "0", 0,
			"C", false, 0, "")
		pdf.Ln(10)
	})
	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)
	pdf.SetX(30)
	pdf.CellFormat(40, 4, "Codigo", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Codigo, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(30)
	pdf.CellFormat(40, 4, "Nombre:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Nombre, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(30)
	pdf.CellFormat(40, 4, "Nivel:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Nivel, "", 0,
		"", false, 0, "")


	pdf.SetFooterFunc(func() {
		pdf.SetY(-20)
		pdf.SetFont("Arial", "", 9)
		pdf.SetX(30)
		pdf.CellFormat(90, 10, "Sadconf.com", "", 0,
			"L", false, 0, "")
		pdf.CellFormat(80, 10, fmt.Sprintf(" %d de {nb}", pdf.PageNo()), "",
			0, "R", false, 0, "")
	})
	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}
// TERMINA CENTRO PDF

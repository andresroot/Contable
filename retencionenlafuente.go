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

// RETENCION EN LA FUENTE
type retencionenlafuente struct {
	Codigo		string
	Nombre      string
	Valor       string
	Porcentaje  string
}


// RETENCION EN LA FUENTE LISTA
func RetencionenlafuenteLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/retencionenlafuente/retencionenlafuenteLista.html")
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM retencionenlafuente")
	if err != nil {
		panic(err.Error())
	}
	res := []retencionenlafuente{}
	for selDB.Next() {
		var Codigo     string
		var Nombre     string
		var Valor      string
		var Porcentaje string
		err = selDB.Scan(&Codigo, &Nombre, &Valor, &Porcentaje)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, retencionenlafuente{Codigo, Nombre, Valor, Porcentaje })
	}
	varmap := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
	}
	tmp.Execute(w, varmap)
}

// INICIA RETENCION EN LA FUENTE PDF
func RetencionenlafuentePdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	t := retencionenlafuente{}
	var e  empresa=ListaEmpresa()
	var c  ciudad=TraerCiudad(e.Ciudad)
	err := db.Get(&t, "SELECT * FROM retencionenlafuente where codigo=$1", Codigo)
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
	pdf.CellFormat(40, 4, t.Nombre, "", 0,
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
	pdf.CellFormat(40, 4, t.Porcentaje, "", 0,
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

package main

// INICIA COMPRA IMPORTAR PAQUETES
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
	"time"
)

// CENTRO TABLA
type inventario struct {
	Fecha       time.Time
	Tipo		string
	Codigo      string
	Bodega		string
	Producto	string
	Cantidad    string
	Precio      string
	Operacion   string
}


// CENTRO LISTA
func InventarioLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/inventario/inventarioLista.html")
	db := dbConn()
	selDB, err := db.Query("SELECT fecha, tipo, codigo, bodega, producto," +
		" cantidad, precio, operacion FROM inventario ORDER BY Fecha ")
	if err != nil {
		panic(err.Error())
	}
	res := []inventario{}
	for selDB.Next() {
		var Fecha       time.Time
		var Tipo		string
		var Codigo      string
		var Bodega		string
		var Producto	string
		var Cantidad    string
		var Precio      string
		var Operacion   string

		err = selDB.Scan(&Fecha, &Tipo, &Codigo, &Bodega, &Producto, &Cantidad,
			&Precio, &Operacion)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, inventario{Fecha, Tipo, Codigo,
			Bodega, Producto, Cantidad, Precio, Operacion})
	}
	varmap := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
	}
	tmp.Execute(w, varmap)
}

// INICIA CENTRO PDF
func inventarioPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Fecha := mux.Vars(r)["Fecha"]
	t := inventario{}
	var e  empresa=ListaEmpresa()
	var c  ciudad=TraerCiudad(e.Ciudad)
	err := db.Get(&t, "SELECT * FROM inventario where fecha=$1", Fecha)
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
	pdf.CellFormat(40, 4, "Fecha", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Fecha.Format("02/01/2006"), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(30)
	pdf.CellFormat(40, 4, "Tipo:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Tipo, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(30)
	pdf.CellFormat(40, 4, "Documento No.:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Codigo, "", 0,
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

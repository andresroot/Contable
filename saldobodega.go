package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
	"log"
	"math"
	"net/http"
	"strconv"
	"text/template"
	"time"
)
func SaldoBodegaLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/saldobodega/saldoBodegaLista.html")
	//	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	//res := []inventario{}
	//listadokardex := []kardex{}

	if Codigo == "False" {

	} else {

		//	FechaInicial := mux.Vars(r)["fechainicial"]

	}

	varmap := map[string]interface{}{
		//"res":     listadokardex,
		"hosting":  ruta,
		"bodega":   ListaBodega(),
		"producto": ListaProducto(),
	}
	tmp.Execute(w, varmap)
}
func SaldoBodegaDatosTodos(w http.ResponseWriter, r *http.Request) {
	var q string
	q="  select '' as filas ,producto,producto.nombre as NombreProducto,bodega.nombre as nombrebodega,bodega,saldo.cantidad from saldo "
	q+=" inner join bodega on bodega.codigo=saldo.bodega"
	q+=" inner join producto on producto.codigo=saldo.producto order by producto,bodega"
	saldofinal:= []saldo{}

	err2 := db.Select(&saldofinal, q)
	if err2 != nil {
		fmt.Println(err2)
	}


	//ProductoParametro := mux.Vars(r)["codigo"]
	//Discriminar := mux.Vars(r)["discriminar"]
	//FechaInicial := mux.Vars(r)["fechainicial"]
	//FechaFinal := mux.Vars(r)["fechafinal"]
	//
	//dateinicial, err := time.Parse("2006-01-02", FechaInicial)
	////datefinal, err := time.Parse("2006-01-02", fechaFinal)
	//
	//if err == nil {
	//	fmt.Println("Fecha Inicial suma"+dateinicial.String())
	//}
	//
	//log.Println("fecha Inicial : " + FechaInicial)
	//BodegaParametro := mux.Vars(r)["bodega"]
	//TipoParametro := mux.Vars(r)["tipo"]
	//listadokardexfinal := []kardex{}
	//listadokardexfinal=KardexDatosTodosGenerar(ProductoParametro,Discriminar,FechaInicial,FechaFinal,BodegaParametro,TipoParametro)

	if len(saldofinal)==0 {
		var slice []string
		slice = make([]string, 0)
		data, _ := json.Marshal(slice)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	} else {
		data, _ := json.Marshal(saldofinal)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}


// INICIA KARDEX TODOS PDF
func SaldoBodegaDatosCabecera(pdf *gofpdf.Fpdf ){
	pdf.SetFont("Arial", "", 10)
	// RELLENO TITULO
	pdf.SetY(66)
	pdf.SetFillColor(224,231,239)
	pdf.SetTextColor(0,0,0)
	pdf.Ln(7)
	pdf.SetX(23)
	pdf.CellFormat(247, 6, "No", "0", 0,
		"L", true, 0, "")
	pdf.SetX(35)
	pdf.CellFormat(40, 6, "Fecha", "0", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 6, "Tipo", "0", 0,
		"L", false, 0, "")
	pdf.SetX(71)
	pdf.CellFormat(40, 6, "Numero", "0", 0,
		"L", false, 0, "")
	pdf.SetX(86)

	pdf.CellFormat(40, 6, "B", "0", 0,
		"L", false, 0, "")
	pdf.SetX(90)
	pdf.CellFormat(40, 6, "Cantidad", "0", 0,
		"L", false, 0, "")
	pdf.SetX(115)

	pdf.CellFormat(40, 6, "Costo", "0", 0,
		"L", false, 0, "")
	pdf.SetX(140)
	pdf.CellFormat(40, 6, "Total", "0", 0,
		"L", false, 0, "")
	pdf.SetX(150)
	pdf.CellFormat(40, 6, "Cantidad", "0", 0,
		"L", false, 0, "")
	pdf.SetX(174)
	pdf.CellFormat(40, 6, "Costo", "0", 0,
		"L", false, 0, "")
	pdf.SetX(200)
	pdf.CellFormat(40, 6, "Total", "0", 0,
		"L", false, 0, "")
	pdf.SetX(210)
	pdf.CellFormat(40, 6, "Cantidad", "0", 0,
		"L", false, 0, "")
	pdf.SetX(235)
	pdf.CellFormat(40, 6, "Costo", "0", 0,
		"L", false, 0, "")
	pdf.SetX(259)
	pdf.CellFormat(40, 6, "Total", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}

func SaldoBoddegaDatosFilaDetalle(pdf *gofpdf.Fpdf,miFila kardex, a int ){
	pdf.SetFont("Arial", "", 8)
	pdf.SetTextColor(0,0,0)

	pdf.SetX(23)
	pdf.CellFormat(183, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(35)
	pdf.CellFormat(30, 4, miFila.Fecha, "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(30, 4, TraerTipoOperacionCorta(miFila.Operacion),  "", 0,
		"L", false, 0, "")
	pdf.SetX(56)
	pdf.CellFormat(30, 4, miFila.Codigo,  "", 0,
		"R", false, 0, "")
	pdf.SetX(60)
	pdf.CellFormat(30, 4, miFila.Bodega, "", 0,
		"R", false, 0, "")

	pdf.SetX(76)
	pdf.CellFormat(30, 4, FormatoFlotante( miFila.Cantidadentrada), "", 0,
		"R", false, 0, "")
	pdf.SetX(96)
	pdf.CellFormat(30, 4, FormatoFlotante( miFila.Precioentrada), "", 0,
		"R", false, 0, "")
	pdf.SetX(120)
	pdf.CellFormat(30, 4,  FormatoFlotante( miFila.Totalentrada), "", 0,
		"R", false, 0, "")

	pdf.SetX(136)
	pdf.CellFormat(30, 4, FormatoFlotante( miFila.Cantidadsalida), "", 0,
		"R", false, 0, "")
	pdf.SetX(156)
	pdf.CellFormat(30, 4, FormatoFlotante( miFila.Preciosalida), "", 0,
		"R", false, 0, "")
	pdf.SetX(180)
	pdf.CellFormat(30, 4,  FormatoFlotante( miFila.Totalsalida), "", 0,
		"R", false, 0, "")

	pdf.SetX(197)
	pdf.CellFormat(30, 4, FormatoFlotante( miFila.Cantidadsaldo), "", 0,
		"R", false, 0, "")
	pdf.SetX(217)
	pdf.CellFormat(30, 4, FormatoFlotante( miFila.Preciosaldo), "", 0,
		"R", false, 0, "")
	pdf.SetX(241)
	pdf.CellFormat(30, 4,  FormatoFlotante( miFila.Totalsaldo), "", 0,
		"R", false, 0, "")
	pdf.Ln(4)
}

func SaldoBodegaDatosPdf(w http.ResponseWriter, r *http.Request) {
	ProductoParametro := mux.Vars(r)["codigo"]
	Discriminar := mux.Vars(r)["discriminar"]
	FechaInicial := mux.Vars(r)["fechainicial"]
	FechaFinal := mux.Vars(r)["fechafinal"]
	dateinicial, err := time.Parse("2006-01-02", FechaInicial)
	datefinal, err := time.Parse("2006-01-02", FechaFinal)

	log.Println("fecha Inicial : " + FechaInicial)
	BodegaParametro := mux.Vars(r)["bodega"]
	TipoParametro := mux.Vars(r)["tipo"]
	listadokardexfinal := []kardex{}
	listadokardexfinal=KardexDatosTodosGenerar(ProductoParametro,Discriminar,FechaInicial,FechaFinal,BodegaParametro,TipoParametro)

	var productoNombre string
	var bodegaNombre string
	var tipoNombre string

	if ProductoParametro == "Todos" {
		productoNombre = "Todos"
	} else {
		db := dbConn()
		miProducto := producto{}
		err = db.Get(&miProducto, "SELECT * FROM producto where codigo=$1", ProductoParametro)
		if err != nil {
			log.Fatalln(err)
		}

		productoNombre = ProductoParametro+" - "+miProducto.Nombre
	}

	if BodegaParametro == "Todas" {
		bodegaNombre = "Todas"
	} else {
		bodegaNombre = BodegaParametro+ " - "+TraerBodega(BodegaParametro)
	}

	if TipoParametro == "Todos" {
		tipoNombre = "Todos"
	} else {
		tipoNombre =  TraerTipoOperacion(TipoParametro)
	}

	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	var buf bytes.Buffer
	var err1 error
	pdf := gofpdf.New("L", "mm", "Letter", cnFontDir)
	ene := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetHeaderFunc(func() {
		pdf.Image(imageFile("logo.png"), 20, 30, 40, 0, false,
			"", 0, "")
		pdf.SetY(17)
		pdf.SetX(55)
		pdf.SetFont("Arial", "", 10)
		pdf.CellFormat(190, 10, e.Nombre, "0", 0,
			"C", false, 0, "")
		pdf.Ln(4)
		pdf.SetX(55)
		pdf.CellFormat(190, 10, "Nit No. " +Coma(e.Codigo)+ " - "+e.Dv, "0", 0, "C",
			false, 0, "")
		pdf.Ln(4)
		pdf.SetX(55)
		pdf.CellFormat(190, 10, e.Iva+ " - "+e.ReteIva, "0", 0, "C", false, 0,
			"")
		pdf.Ln(4)
		pdf.SetX(55)
		pdf.CellFormat(190, 10, "Actividad Ica - "+e.ActividadIca, "0",
			0, "C", false, 0, "")
		pdf.Ln(4)
		pdf.SetX(55)
		pdf.CellFormat(190, 10, e.Direccion, "0", 0, "C", false, 0,
			"")
		pdf.Ln(4)
		pdf.SetX(55)
		pdf.CellFormat(190, 10, e.Telefono1+" - "+e.Telefono2, "0", 0, "C", false, 0,
			"")
		pdf.Ln(4)
		pdf.SetX(55)
		pdf.CellFormat(190, 10, ene(c.NombreCiudad+ " - "+c.NombreDepartamento), "0", 0, "C", false, 0,
			"")
		pdf.Ln(10)

		// RELLENO TITULO
		pdf.SetX(23)
		pdf.SetFillColor(224,231,239)
		pdf.SetTextColor(0,0,0)
		pdf.CellFormat(247, 6, "DATOS KARDEX PRODUCTO", "0", 0,
			"C", true, 0, "")
		pdf.Ln(6)
		pdf.SetX(23)
		pdf.CellFormat(20, 10, "Producto:", "0", 0, "L", false, 0,
			"")
		pdf.SetX(45)
		pdf.CellFormat(20, 10, productoNombre, "0", 0, "L", false, 0,
			"")
		pdf.SetX(170)
		pdf.CellFormat(20, 10, "Desde:", "0", 0, "L", false, 0,
			"")
		pdf.SetX(185)
		pdf.CellFormat(20, 10, dateinicial.Format("02/01/2006"), "0", 0, "L", false, 0,
			"")
		pdf.SetX(220)
		pdf.CellFormat(20, 10, "Hasta:", "0", 0, "L", false, 0,
			"")
		pdf.SetX(235)
		pdf.CellFormat(20, 10, datefinal.Format("02/01/2006"), "0", 0, "L", false, 0,
			"")
		pdf.Ln(6)
		pdf.SetX(23)
		pdf.CellFormat(20, 10, "Bodegas:", "0", 0, "L", false, 0,
			"")
		pdf.SetX(45)
		pdf.CellFormat(20, 10, bodegaNombre, "0", 0, "L", false, 0,
			"")
		pdf.SetX(170)
		pdf.CellFormat(20, 10, "Tipo:", "0", 0, "L", false, 0,
			"")
		pdf.SetX(180)
		pdf.CellFormat(20, 10, ene(tipoNombre), "0", 0, "L", false, 0,
			"")
		pdf.SetX(220)
		pdf.CellFormat(20, 10, "Discriminado:", "0", 0, "L", false, 0,
			"")
		pdf.SetX(245)
		pdf.CellFormat(20, 10, Discriminar, "0", 0, "L", false, 0,
			"")
		pdf.Ln(6)
	})

	pdf.SetFooterFunc(func() {
		pdf.SetFont("Arial", "", 9)
		pdf.SetTextColor(0,0,0)
		pdf.SetY(192)
		pdf.SetX(23)
		// LINEA
		pdf.Line(23,199,269,199)
		pdf.Ln(6)
		pdf.SetX(23)
		pdf.CellFormat(90, 10, "Sadconf.com", "", 0,
			"L", false, 0, "")
		pdf.CellFormat(161, 10, fmt.Sprintf(" %d de {nb}", pdf.PageNo()), "",
			0, "R", false, 0, "")
	})

	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 8)
	pdf.SetX(30)

	KardexDatosCabecera(pdf)
	// tercera pagina

	for i, miFila := range listadokardexfinal {
		KardexDatosFilaDetalle(pdf,miFila,i+1)

		if math.Mod(float64(i+1),28)==0 {
			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			KardexDatosCabecera(pdf)
		}

	}

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}
// TERMINA EMPRESA TODOS PDF


// KARDEX EXCEL
func SaldoBodegaDatosExcel(w http.ResponseWriter, r *http.Request) {
	ProductoParametro := mux.Vars(r)["codigo"]
	Discriminar := mux.Vars(r)["discriminar"]
	FechaInicial := mux.Vars(r)["fechainicial"]
	FechaFinal := mux.Vars(r)["fechafinal"]
	dateinicial, err := time.Parse("2006-01-02", FechaInicial)
	datefinal, err := time.Parse("2006-01-02", FechaFinal)

	if err == nil {
		fmt.Println("Fecha Inicial suma"+dateinicial.String())
	}

	log.Println("fecha Inicial : " + FechaInicial)
	BodegaParametro := mux.Vars(r)["bodega"]
	TipoParametro := mux.Vars(r)["tipo"]
	listadokardexfinal := []kardex{}
	listadokardexfinal=KardexDatosTodosGenerar(ProductoParametro,Discriminar,FechaInicial,FechaFinal,BodegaParametro,TipoParametro)

	var productoNombre string
	var bodegaNombre string
	var tipoNombre string

	if ProductoParametro == "Todos" {
		productoNombre = "Todos"
	} else {
		db := dbConn()
		miProducto := producto{}
		err = db.Get(&miProducto, "SELECT * FROM producto where codigo=$1", ProductoParametro)
		if err != nil {
			log.Fatalln(err)
		}

		productoNombre = ProductoParametro+" - "+miProducto.Nombre
	}

	if BodegaParametro == "Todas" {
		bodegaNombre = "Todas"
	} else {
		bodegaNombre = BodegaParametro+ " - "+TraerBodega(BodegaParametro)
	}

	if TipoParametro == "Todos" {
		tipoNombre = "Todos"
	} else {
		tipoNombre =  TraerTipoOperacion(TipoParametro)
	}

	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)

	f := excelize.NewFile()

	// FUNCION ANCHO DE LA COLUMNA
	if err =f.SetColWidth("Sheet1", "A", "A", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err =f.SetColWidth("Sheet1", "B", "B", 24); err != nil {
		fmt.Println(err)
		return
	}
	if err =f.SetColWidth("Sheet1", "C", "C", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err =f.SetColWidth("Sheet1", "D", "D", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err =f.SetColWidth("Sheet1", "E", "E", 15); err != nil {
		fmt.Println(err)
		return
	}

	if err =f.SetColWidth("Sheet1", "F", "F", 15); err != nil {
		fmt.Println(err)
		return
	}

	if err =f.SetColWidth("Sheet1", "G", "G", 15); err != nil {
		fmt.Println(err)
		return
	}

	if err =f.SetColWidth("Sheet1", "H", "H", 15); err != nil {
		fmt.Println(err)
		return
	}

	if err =f.SetColWidth("Sheet1", "I", "I", 15); err != nil {
		fmt.Println(err)
		return
	}

	if err =f.SetColWidth("Sheet1", "J", "J", 15); err != nil {
		fmt.Println(err)
		return
	}

	if err =f.SetColWidth("Sheet1", "K", "K", 15); err != nil {
		fmt.Println(err)
		return
	}

	if err =f.SetColWidth("Sheet1", "L", "L", 15); err != nil {
		fmt.Println(err)
		return
	}

	if err =f.SetColWidth("Sheet1", "M", "M", 15); err != nil {
		fmt.Println(err)
		return
	}
	// FUNCION PARA UNIR DOS CELDAS
	if err = f.MergeCell("Sheet1", "A1", "M1"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A2", "M2"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A3", "M3"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A4", "M4"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A5", "M5"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A6", "M6"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A7", "M7"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A8", "M8"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A9", "M9"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A10", "M10"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A13", "M13"); err != nil {
		fmt.Println(err)
		return
	}

	estiloTitulo, err := f.NewStyle(`{  "alignment":{"horizontal": "center"},"font":{"bold":false,"italic":false,"family":"Arial","size":10,"color":"##000000"}}`)

	// titulo
	f.SetCellValue("Sheet1", "A1", e.Nombre)
	f.SetCellValue("Sheet1", "A2","Nit No. "+Coma(e.Codigo)+" - "+e.Dv)
	f.SetCellValue("Sheet1", "A3",e.Iva+" - "+e.ReteIva)
	f.SetCellValue("Sheet1", "A4","Actividad Ica - "+e.ActividadIca)
	f.SetCellValue("Sheet1", "A5",e.Direccion)
	f.SetCellValue("Sheet1", "A6",e.Telefono1+" - "+e.Telefono2)
	f.SetCellValue("Sheet1", "A7",(c.NombreCiudad+" - "+c.NombreDepartamento))
	f.SetCellValue("Sheet1", "A9","DATOS KARDEX DEL "+dateinicial.Format("02/01/2006")+" AL "+datefinal.Format("02/01/2006"))

	f.SetCellStyle("Sheet1","A1","A1",estiloTitulo)
	f.SetCellStyle("Sheet1","A2","A2",estiloTitulo)
	f.SetCellStyle("Sheet1","A3","A3",estiloTitulo)
	f.SetCellStyle("Sheet1","A4","A4",estiloTitulo)
	f.SetCellStyle("Sheet1","A5","A5",estiloTitulo)
	f.SetCellStyle("Sheet1","A6","A6",estiloTitulo)
	f.SetCellStyle("Sheet1","A7","A7",estiloTitulo)
	f.SetCellStyle("Sheet1","A8","A8",estiloTitulo)
	f.SetCellStyle("Sheet1","A9","A9",estiloTitulo)
	f.SetCellStyle("Sheet1","A10","A10",estiloTitulo)

	var filaExcel=11

	estiloTexto, err := f.NewStyle(`{"font":{"bold":false,"italic":false,"family":"Arial","size":10,"color":"##000000"}}`)

	estiloNumeroDetalle, err := f.NewStyle(`{"number_format": 4,"font":{"bold":false,"italic":false,"family":"Arial","size":10,"color":"##000000"}}`)

	if err != nil {
		fmt.Println(err)
	}

	estiloCabecera, err := f.NewStyle(`{
"alignment":{"horizontal":"center"},
    "border": [
    {
        "type": "left",
        "color": "#000000",
        "style": 1
    },
    {
        "type": "top",
        "color": "#000000",
        "style": 1
    },
    {
        "type": "bottom",
        "color": "#000000",
        "style": 1
    },
    {
        "type": "right",
        "color": "#000000",
        "style": 1
    }]
}`)
	if err != nil {
		fmt.Println(err)
	}
	// datos
	f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel),"Producto:")
	f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel),productoNombre)

	f.SetCellValue("Sheet1", "G"+strconv.Itoa(filaExcel),"Desde:")
	f.SetCellValue("Sheet1", "H"+strconv.Itoa(filaExcel),dateinicial.Format("02/01/2006"))

	f.SetCellValue("Sheet1", "I"+strconv.Itoa(filaExcel),"Hasta:")
	f.SetCellValue("Sheet1", "J"+strconv.Itoa(filaExcel),datefinal.Format("02/01/2006"))

	filaExcel++

	f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel),"Bodegas:")
	f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel),bodegaNombre)

	f.SetCellValue("Sheet1", "G"+strconv.Itoa(filaExcel),"Tipo:")
	f.SetCellValue("Sheet1", "H"+strconv.Itoa(filaExcel),tipoNombre)

	f.SetCellValue("Sheet1", "J"+strconv.Itoa(filaExcel),"Discriminar:")
	f.SetCellValue("Sheet1", "K"+strconv.Itoa(filaExcel),Discriminar)

	filaExcel++
	filaExcel++

	//cabecera
	f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel),"Fecha")
	f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel), "Tipo")
	f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel), "Codigo")
	f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel), "Bodega")

	f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel), "Cantidad")
	f.SetCellValue("Sheet1", "F"+strconv.Itoa(filaExcel), "Precio")
	f.SetCellValue("Sheet1", "G"+strconv.Itoa(filaExcel), "Total")

	f.SetCellValue("Sheet1", "H"+strconv.Itoa(filaExcel), "Cantidad")
	f.SetCellValue("Sheet1", "I"+strconv.Itoa(filaExcel), "Precio")
	f.SetCellValue("Sheet1", "J"+strconv.Itoa(filaExcel), "Total")

	f.SetCellValue("Sheet1", "K"+strconv.Itoa(filaExcel), "Cantidad")
	f.SetCellValue("Sheet1", "L"+strconv.Itoa(filaExcel), "Precio")
	f.SetCellValue("Sheet1", "M"+strconv.Itoa(filaExcel), "Total")

	f.SetCellStyle("Sheet1","A"+strconv.Itoa(filaExcel),"A"+strconv.Itoa(filaExcel),estiloCabecera)
	f.SetCellStyle("Sheet1","B"+strconv.Itoa(filaExcel),"B"+strconv.Itoa(filaExcel),estiloCabecera)
	f.SetCellStyle("Sheet1","C"+strconv.Itoa(filaExcel),"C"+strconv.Itoa(filaExcel),estiloCabecera)
	f.SetCellStyle("Sheet1","D"+strconv.Itoa(filaExcel),"D"+strconv.Itoa(filaExcel),estiloCabecera)
	f.SetCellStyle("Sheet1","E"+strconv.Itoa(filaExcel),"E"+strconv.Itoa(filaExcel),estiloCabecera)
	f.SetCellStyle("Sheet1","F"+strconv.Itoa(filaExcel),"F"+strconv.Itoa(filaExcel),estiloCabecera)
	f.SetCellStyle("Sheet1","G"+strconv.Itoa(filaExcel),"G"+strconv.Itoa(filaExcel),estiloCabecera)
	f.SetCellStyle("Sheet1","H"+strconv.Itoa(filaExcel),"H"+strconv.Itoa(filaExcel),estiloCabecera)
	f.SetCellStyle("Sheet1","I"+strconv.Itoa(filaExcel),"I"+strconv.Itoa(filaExcel),estiloCabecera)
	f.SetCellStyle("Sheet1","J"+strconv.Itoa(filaExcel),"J"+strconv.Itoa(filaExcel),estiloCabecera)
	f.SetCellStyle("Sheet1","K"+strconv.Itoa(filaExcel),"K"+strconv.Itoa(filaExcel),estiloCabecera)
	f.SetCellStyle("Sheet1","L"+strconv.Itoa(filaExcel),"L"+strconv.Itoa(filaExcel),estiloCabecera)
	f.SetCellStyle("Sheet1","M"+strconv.Itoa(filaExcel),"M"+strconv.Itoa(filaExcel),estiloCabecera)

	filaExcel++

	for i, miFila := range listadokardexfinal{
		a:=strconv.Itoa(filaExcel+i)

		f.SetCellValue("Sheet1", "A"+a,miFila.Fecha)
		f.SetCellValue("Sheet1", "B"+a,miFila.Tipo)
		f.SetCellValue("Sheet1", "C"+a, miFila.Codigo)
		f.SetCellValue("Sheet1", "D"+a, miFila.Bodega)

		f.SetCellValue("Sheet1", "E"+a, miFila.Cantidadentrada)
		f.SetCellValue("Sheet1", "F"+a, miFila.Precioentrada)
		f.SetCellValue("Sheet1", "G"+a, miFila.Totalentrada)

		f.SetCellValue("Sheet1", "H"+a, miFila.Cantidadsalida)
		f.SetCellValue("Sheet1", "I"+a, miFila.Preciosalida)
		f.SetCellValue("Sheet1", "J"+a, miFila.Totalsalida)
		f.SetCellValue("Sheet1", "K"+a, miFila.Cantidadsaldo)
		f.SetCellValue("Sheet1", "L"+a, miFila.Preciosaldo)
		f.SetCellValue("Sheet1", "M"+a, miFila.Totalsaldo)

		f.SetCellStyle("Sheet1","A"+a,"A"+a,estiloTexto)
		f.SetCellStyle("Sheet1","B"+a,"B"+a,estiloTexto)
		f.SetCellStyle("Sheet1","C"+a,"C"+a,estiloTexto)
		f.SetCellStyle("Sheet1","D"+a,"D"+a,estiloTexto)
		f.SetCellStyle("Sheet1","E"+a,"E"+a,estiloTexto)
		f.SetCellStyle("Sheet1","F"+a,"F"+a,estiloNumeroDetalle)
		f.SetCellStyle("Sheet1","G"+a,"G"+a,estiloNumeroDetalle)
		f.SetCellStyle("Sheet1","H"+a,"H"+a,estiloNumeroDetalle)
		f.SetCellStyle("Sheet1","I"+a,"I"+a,estiloNumeroDetalle)
		f.SetCellStyle("Sheet1","J"+a,"J"+a,estiloNumeroDetalle)
		f.SetCellStyle("Sheet1","K"+a,"K"+a,estiloNumeroDetalle)
		f.SetCellStyle("Sheet1","L"+a,"L"+a,estiloNumeroDetalle)
		f.SetCellStyle("Sheet1","M"+a,"M"+a,estiloNumeroDetalle)

		//van=i
	}

	// Set the headers necessary to get browsers to interpret the downloadable file
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment;filename=Libro.xlsx")
	w.Header().Set("File-Name", "Libro.xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	err = f.Write(w)
	if err != nil {
		panic(err.Error())
	}
}

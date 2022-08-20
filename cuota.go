package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

// CUOTA TABLA
type Cuota struct {
	Filas           string  `json:"Filas"`
	Monto           string  `json:"Monto"`
	Plazo           string  `json:"Plazo"`
	Interes         string  `json:"Interes"`
	Pago 			float64 `json:"Cuota"`
	Fecha           string `json:"Fecha"`
	Inicial   		float64  `json:"Inicial"`
	Intereses       float64 `json:"Intereses"`
	Capital         float64  `json:"Capital"`
	Final         	float64  `json:"Final"`
}

type Prestamo struct {
	Monto           string  `json:"Monto"`
	Plazo           string  `json:"Plazo"`
	Interes         string  `json:"Interes"`
	Fecha           string `json:"Fecha"`
}

// CENTRO KARDEX
func CuotaLista(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	t := Prestamo{}
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM prestamo")
	err := row.Scan(&total)
	if err != nil {
		log.Fatal(err)
	}
	//var resultado bool
	if total > 0 {
		err := db.Get(&t, "SELECT * FROM prestamo")
		if err != nil {
			log.Fatal(err)
		}
		//resultado = true
	} else {
		//resultado = false
	}


	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/cuota/cuotaLista.html")
	varmap := map[string]interface{}{
		"hosting":  ruta,
		"prestamo":  t,

	}
	tmp.Execute(w, varmap)
}

func Calcularcuota(Monto string, Plazo string, parametroIntereses string, FechaInicial string)[]Cuota{
	log.Println("Monto : " + Monto)
	log.Println("Plazo : " + Plazo)
	log.Println("Intereses : " + parametroIntereses)
	log.Println("fecha Inicial : " + FechaInicial)

	var monto float64
	var plazo float64
	var interes float64
	var pago float64
	var fechainicial time.Time
	monto= Flotante(Monto)

	plazo,err1 := strconv.ParseFloat(Plazo,8)
	if err1 != nil {
		log.Fatalln(err1)
	}
	interes,err2 := strconv.ParseFloat(parametroIntereses,8)
	if err2 != nil {
		log.Fatalln(err2)
	}
	fechainicial,err := time.Parse("2006-01-02", FechaInicial)
	if err != nil {
		log.Fatalln(err)
	}

	// borra datos
	delForm, err := db.Prepare("DELETE from prestamo")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec()
	//inserat datos
	insForm, err := db.Prepare("INSERT INTO prestamo(monto,plazo,interes,fecha)VALUES($1, $2,$3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insForm.Exec(Monto,Plazo,parametroIntereses,fechainicial.Format("2006-01-02"))


	var arriba float64
	arriba=monto * ((interes/100)*((math.Pow(1+(interes/100),plazo))))
	fmt.Println("Arriba2")
	fmt.Println(arriba)
	var abajo float64

	abajo=(math.Pow((1+(interes/100)),plazo))-1

	pago = math.Round(arriba/abajo)
	fmt.Println("abajo")
	fmt.Println(abajo)

	fmt.Println("pago")
	fmt.Println(pago)

	var montolinea float64
	var intereslinea float64
	var capitallinea float64
	var finallinea float64
	montolinea=monto
	listadocuota := []Cuota{}

	for i:=1 ;i<int(plazo)+1;i++{
		fmt.Println("linea"+strconv.Itoa(i))
		intereslinea=math.Round(montolinea*(interes/100))
		capitallinea=pago-intereslinea
		finallinea=montolinea-capitallinea

		fmt.Println("interees")
		fmt.Println(intereslinea)
		fmt.Println("capital")
		fmt.Println(capitallinea)
		fmt.Println("final")
		fmt.Println(finallinea)

		listadocuota=append(listadocuota,Cuota{strconv.Itoa(i),Monto,Plazo,parametroIntereses,pago,fechainicial.AddDate(0,i-1,0).Format("02/01/2006"), capitallinea+finallinea, intereslinea,capitallinea, finallinea })
		montolinea=finallinea
	}
	return listadocuota
}


func CuotaDatos(w http.ResponseWriter, r *http.Request) {
	//db := dbConn()
	Monto := mux.Vars(r)["monto"]
	Plazo := mux.Vars(r)["plazo"]
	Intereses := mux.Vars(r)["intereses"]
	FechaInicial := mux.Vars(r)["fechainicial"]
	listadocuota := []Cuota{}
	listadocuota= Calcularcuota(Monto, Plazo, Intereses, FechaInicial)

		data, _ := json.Marshal(listadocuota)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
}


// CUOTA LISTA
func cuotaLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/kardex/kardexLista.html")
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


// INICIA COMPROBANTE TODOS PDF
func CuotaTodosCabecera(pdf *gofpdf.Fpdf){
	pdf.SetFont("Arial", "", 10)
	// RELLENO TITULO
	pdf.SetY(60)
	pdf.SetFillColor(224,231,239)
	pdf.SetTextColor(0,0,0)
	pdf.Ln(6)
	pdf.SetX(20)
	pdf.CellFormat(184, 6, "No", "0", 0,
		"L", true, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(190, 6, "Fecha", "0", 0,
		"L", false, 0, "")
	pdf.SetX(64)
	pdf.CellFormat(190, 6, "Saldo Inicial", "0", 0,
		"L", false, 0, "")
	pdf.SetX(98)
	pdf.CellFormat(190, 6, "Intereses", "0", 0,
		"L", false, 0, "")
	pdf.SetX(132)
	pdf.CellFormat(190, 6, "Capital", "0", 0,
		"L", false, 0, "")
	pdf.SetX(156)
	pdf.CellFormat(190, 6, "Saldo Final", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}

func CuotaTodosDetalle(pdf *gofpdf.Fpdf,miFila Cuota, a int ){
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(21)
	pdf.CellFormat(181, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(40, 4, miFila.Fecha, "", 0,
		"L", false, 0, "")
	pdf.SetX(45)
	pdf.CellFormat(40, 4, FormatoFlotante(miFila.Inicial), "", 0,"R", false, 0, "")
	pdf.SetX(75)
	pdf.CellFormat(40, 4, FormatoFlotante(miFila.Intereses), "", 0,
		"R", false, 0, "")
	pdf.SetX(105)
	pdf.CellFormat(40, 4, FormatoFlotante(miFila.Capital), "", 0,
		"R", false, 0, "")
	pdf.SetX(135)
	pdf.CellFormat(40, 4, FormatoFlotante(miFila.Final), "", 0,
		"R", false, 0, "")
	pdf.Ln(4)
}

func CuotaTodosPdf(w http.ResponseWriter, r *http.Request) {
	Monto := mux.Vars(r)["monto"]
	Plazo := mux.Vars(r)["plazo"]
	Intereses := mux.Vars(r)["intereses"]
	FechaInicial := mux.Vars(r)["fechainicial"]
	listadocuota := []Cuota{}
	listadocuota= Calcularcuota(Monto, Plazo, Intereses, FechaInicial)
	var Cuota float64
	for _, miFila := range listadocuota {
		Cuota= miFila.Pago
	}

	var e empresa=ListaEmpresa()
	var c ciudad=TraerCiudad(e.Ciudad)
	var buf bytes.Buffer
	var err1 error
	pdf := gofpdf.New("P", "mm", "Letter", cnFontDir)
	ene := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetHeaderFunc(func() {
		pdf.Image(imageFile("logo.png"), 20, 30, 40, 0, false,
			"", 0, "")
		pdf.SetY(17)
		pdf.SetFont("Arial", "", 10)
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
		pdf.CellFormat(190, 10, e.Telefono1+" - "+e.Telefono2, "0", 0, "C", false, 0,
			"")
		pdf.Ln(4)
		pdf.CellFormat(190, 10, ene(c.NombreCiudad+ " - "+c.NombreDepartamento), "0", 0, "C", false, 0,
			"")
		pdf.Ln(10)

		// RELLENO TITULO
		pdf.SetX(20)
		pdf.SetFillColor(224,231,239)
		pdf.SetTextColor(0,0,0)
		pdf.CellFormat(184, 6, "DATOS PRESTAMO", "0", 0,
			"C", true, 0, "")
		pdf.Ln(6)
		pdf.SetX(20)
		pdf.CellFormat(20, 10, "Monto:", "0", 0, "L", false, 0,
			"")
		pdf.SetX(35)
		pdf.CellFormat(20, 10, FormatoFlotante(Flotante(Monto)), "0", 0, "L", false, 0,
			"")
		pdf.SetX(65)
		pdf.CellFormat(20, 10, "Plazo:", "0", 0, "L", false, 0,
			"")
		pdf.SetX(78)
		pdf.CellFormat(20, 10, Plazo, "0", 0, "L", false, 0,
			"")
		pdf.SetX(92)
		pdf.CellFormat(20, 10, "Intereses:", "0", 0, "L", false, 0,
			"")
		pdf.SetX(111)
		pdf.CellFormat(20, 10, Intereses, "0", 0, "L", false, 0,
			"")
		pdf.SetX(125)
		pdf.CellFormat(20, 10, "Fecha:", "0", 0, "L", false, 0,
			"")
		pdf.SetX(141)
		pdf.CellFormat(20, 10, FechaInicial, "0", 0, "L", false, 0,
			"")
		pdf.SetX(165)
		pdf.CellFormat(20, 10, "Cuota:", "0", 0, "L", false, 0,
			"")
		pdf.SetX(180)
		pdf.CellFormat(20, 10, FormatoFlotante(Cuota), "0", 0, "L", false, 0,
			"")
		pdf.Ln(6)
	})

	pdf.SetFooterFunc(func() {
		pdf.SetTextColor(0,0,0)
		pdf.SetY(252)
		pdf.SetFont("Arial", "", 9)
		pdf.SetX(20)

		// LINEA
		pdf.Line(20,259,204,259)
		pdf.Ln(6)
		pdf.SetX(20)
		pdf.CellFormat(90, 10, "Sadconf.com", "", 0, "L", false, 0, "")
		pdf.SetX(129)
		pdf.CellFormat(80, 10, fmt.Sprintf(" %d de {nb}", pdf.PageNo()), "",
			0, "R", false, 0, "")
	})

	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)
	pdf.SetX(30)

	CuotaTodosCabecera(pdf)
	// tercera pagina

	for i, miFila := range listadocuota {
		CuotaTodosDetalle(pdf,miFila,i+1)
		if math.Mod(float64(i+1),46)==0 {
			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			CuotaTodosCabecera(pdf)
		}

	}

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}
// TERMINA COMPROBANTE TODOS PDF

// CUOTA EXCEL
func CuotaExcel(w http.ResponseWriter, r *http.Request) {
	Monto := mux.Vars(r)["monto"]
	Plazo := mux.Vars(r)["plazo"]
	Intereses := mux.Vars(r)["intereses"]
	FechaInicial := mux.Vars(r)["fechainicial"]
	listadocuota := []Cuota{}
	listadocuota= Calcularcuota(Monto, Plazo, Intereses, FechaInicial)
	var Cuota float64
	for _, miFila := range listadocuota {
		Cuota= miFila.Pago
	}

	var e  empresa=ListaEmpresa()
	var c  ciudad=TraerCiudad(e.Ciudad)

	f := excelize.NewFile()

	// FUNCION ANCHO DE LA COLUMNA
	if err :=f.SetColWidth("Sheet1", "A", "A", 13); err != nil {
		fmt.Println(err)
		return
	}
	if err :=f.SetColWidth("Sheet1", "B", "B", 13); err != nil {
		fmt.Println(err)
		return
	}
	if err :=f.SetColWidth("Sheet1", "C", "C", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err :=f.SetColWidth("Sheet1", "D", "D", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err :=f.SetColWidth("Sheet1", "E", "E", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err :=f.SetColWidth("Sheet1", "F", "F", 13); err != nil {
		fmt.Println(err)
		return
	}

	// FUNCION PARA UNIR DOS CELDAS
	if err := f.MergeCell("Sheet1", "A1", "F1"); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.MergeCell("Sheet1", "A2", "F2"); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.MergeCell("Sheet1", "A3", "F3"); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.MergeCell("Sheet1", "A4", "F4"); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.MergeCell("Sheet1", "A5", "F5"); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.MergeCell("Sheet1", "A6", "F6"); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.MergeCell("Sheet1", "A7", "F7"); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.MergeCell("Sheet1", "A8", "F8"); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.MergeCell("Sheet1", "A9", "F9"); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.MergeCell("Sheet1", "A14", "E14"); err != nil {
		fmt.Println(err)
		return
	}
	estiloTitulo, err := f.NewStyle(`{  "alignment":{"horizontal": "center"},"font":{"bold":false,"italic":false,"family":"Arial","size":10,"color":"##000000"}}`)
	//estiloTituloderecha, err := f.NewStyle(`{  "alignment":{"horizontal": "right"},"font":{"bold":false,"italic":false,"family":"Arial","size":10,"color":"##000000"}}`)
	estiloTituloizquierda, err := f.NewStyle(`{  "alignment":{"horizontal": "left"},"font":{"bold":false,"italic":false,"family":"Arial","size":10,"color":"##000000"}}`)
	estiloTexto, err := f.NewStyle(`{"font":{"bold":false,"italic":false,"family":"Arial","size":10,"color":"#000000"}}`)

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
	estiloNumero, err := f.NewStyle(`{"number_format": 3,"font":{"bold":false,"italic":false,"family":"Arial","size":10,"color":"##000000"}}`)

	if err != nil {
		fmt.Println(err)
	}
	// titulo
	f.SetCellValue("Sheet1", "A1", e.Nombre)
	f.SetCellValue("Sheet1", "A2","Nit No. "+Coma(e.Codigo)+" - "+e.Dv)
	f.SetCellValue("Sheet1", "A3",e.Iva+" - "+e.ReteIva)
	f.SetCellValue("Sheet1", "A4","Actividad Ica - "+e.ActividadIca)
	f.SetCellValue("Sheet1", "A5",e.Direccion)
	f.SetCellValue("Sheet1", "A6",(e.Telefono1+" - "+e.Telefono2))
	f.SetCellValue("Sheet1", "A7",(c.NombreCiudad+" - "+c.NombreDepartamento))
	f.SetCellValue("Sheet1", "A8","")
	f.SetCellValue("Sheet1", "A9","DATOS PRESTAMO")
	f.SetCellValue("Sheet1", "A10","Monto")
	f.SetCellValue("Sheet1", "A11","Plazo")
	f.SetCellValue("Sheet1", "A12","Intereses")
	f.SetCellValue("Sheet1", "A13","Cuota")
	f.SetCellValue("Sheet1", "A14","")

	f.SetCellStyle("Sheet1","A1","A1",estiloTitulo)
	f.SetCellStyle("Sheet1","A2","A2",estiloTitulo)
	f.SetCellStyle("Sheet1","A3","A3",estiloTitulo)
	f.SetCellStyle("Sheet1","A4","A4",estiloTitulo)
	f.SetCellStyle("Sheet1","A5","A5",estiloTitulo)
	f.SetCellStyle("Sheet1","A6","A6",estiloTitulo)
	f.SetCellStyle("Sheet1","A7","A7",estiloTitulo)
	f.SetCellStyle("Sheet1","A8","A8",estiloTitulo)
	f.SetCellStyle("Sheet1","A9","A9",estiloTitulo)
	f.SetCellStyle("Sheet1","A10","A10",estiloTituloizquierda)
	f.SetCellStyle("Sheet1","A11","A11",estiloTituloizquierda)
	f.SetCellStyle("Sheet1","A12","A12",estiloTituloizquierda)
	f.SetCellStyle("Sheet1","A13","A13",estiloTituloizquierda)
	f.SetCellStyle("Sheet1","A14","A14",estiloTitulo)


	f.SetCellValue("Sheet1", "B10",Flotante(Monto))
	f.SetCellValue("Sheet1", "B11",Flotante(Plazo))
	f.SetCellValue("Sheet1", "B12",Flotante(Intereses))
	f.SetCellValue("Sheet1", "B13",Cuota)
	f.SetCellValue("Sheet1", "B14","")

	f.SetCellStyle("Sheet1","B10","B10",estiloNumero)
	f.SetCellStyle("Sheet1","B11","B11",estiloNumero)
	f.SetCellStyle("Sheet1","B12","B12",estiloNumero)
	f.SetCellStyle("Sheet1","B13","B13",estiloNumero)
	f.SetCellStyle("Sheet1","B14","B14",estiloTexto)

	var filaExcel=15


	//cabecera
	f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel),"Fila")
	f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel),"Fecha")
	f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel), "Inicial")
	f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel), "Intereses")
	f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel), "Capital")
	f.SetCellValue("Sheet1", "F"+strconv.Itoa(filaExcel), "Saldo")

	f.SetCellStyle("Sheet1","A"+strconv.Itoa(filaExcel),"A"+strconv.Itoa(filaExcel),estiloCabecera)
	f.SetCellStyle("Sheet1","B"+strconv.Itoa(filaExcel),"B"+strconv.Itoa(filaExcel),estiloCabecera)
	f.SetCellStyle("Sheet1","C"+strconv.Itoa(filaExcel),"C"+strconv.Itoa(filaExcel),estiloCabecera)
	f.SetCellStyle("Sheet1","D"+strconv.Itoa(filaExcel),"D"+strconv.Itoa(filaExcel),estiloCabecera)
	f.SetCellStyle("Sheet1","E"+strconv.Itoa(filaExcel),"E"+strconv.Itoa(filaExcel),estiloCabecera)
	f.SetCellStyle("Sheet1","F"+strconv.Itoa(filaExcel),"F"+strconv.Itoa(filaExcel),estiloCabecera)
	filaExcel++

	for i, miFila := range listadocuota{
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel+i), miFila.Filas)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel+i), miFila.Fecha)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel+i), miFila.Inicial)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel+i), miFila.Intereses)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel+i), miFila.Capital)
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(filaExcel+i), miFila.Final)

		f.SetCellStyle("Sheet1","A"+strconv.Itoa(filaExcel+i),"A"+strconv.Itoa(filaExcel+i),estiloTexto)
		f.SetCellStyle("Sheet1","B"+strconv.Itoa(filaExcel+i),"B"+strconv.Itoa(filaExcel+i),estiloNumero)
		f.SetCellStyle("Sheet1","C"+strconv.Itoa(filaExcel+i),"C"+strconv.Itoa(filaExcel+i),estiloNumero)
		f.SetCellStyle("Sheet1","D"+strconv.Itoa(filaExcel+i),"D"+strconv.Itoa(filaExcel+i),estiloNumero)
		f.SetCellStyle("Sheet1","E"+strconv.Itoa(filaExcel+i),"E"+strconv.Itoa(filaExcel+i),estiloNumero)
		f.SetCellStyle("Sheet1","F"+strconv.Itoa(filaExcel+i),"F"+strconv.Itoa(filaExcel+i),estiloNumero)
		//van=i
	}

	// LINEA FINAL
	//a=strconv.Itoa(van+1+filaExcel)
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



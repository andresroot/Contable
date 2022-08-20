package main

import (
	"bytes"
	"database/sql"
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
)

type conceptoJson struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Nombre string `json:"nombre"`
}

// CONCEPTO TABLA
type concepto struct {
	Codigo      string
	Nombre      string
}

// CONCEPTO LISTA
func ConceptoLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/concepto/conceptoLista.html")
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM concepto ORDER BY cast(codigo as integer) ASC")
	if err != nil {
		panic(err.Error())
	}
	res := []concepto{}
	for selDB.Next() {
		var Codigo string
		var Nombre string
		err = selDB.Scan(&Codigo, &Nombre)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, concepto{Codigo, Nombre })
	}
	varmap := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
	}
	tmp.Execute(w, varmap)
}

// CONCEPTO NUEVO
func ConceptoNuevo(w http.ResponseWriter, r *http.Request) {
	// TRAER COPIA DE EDITAR
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	emp := concepto{}
	if Codigo == "False"{
	} else {
		err := db.Get(&emp, "SELECT * FROM concepto where codigo=$1", Codigo)
		if err != nil {
			log.Fatalln(err)
		}
	}

	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/concepto/conceptoNuevo.html")

	parametros := map[string]interface{}{
		"emp":     emp,
		"hosting": ruta,
		"codigo": Codigo,
	}
	tmp.Execute(w, parametros)
	// TERMINA TRAER COPIA DE EDITAR
}

// CONCEPTO INSERTAR
func ConceptoInsertar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		Codigo := r.FormValue("Codigo")
		Nombre := r.FormValue("Nombre")
		Nombre = Titulo(Nombre)
		insForm, err := db.Prepare("INSERT INTO concepto(codigo, nombre)VALUES($1, $2)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(Codigo, Nombre)
		log.Println("Nuevo Registro:" + Codigo + "," + Nombre)
	}
	http.Redirect(w, r, "/ConceptoLista", 301)
}

// CONCEPTO EXISTE
func ConceptoExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM concepto  WHERE codigo=$1", Codigo)
	err := row.Scan(&total)
	if err != nil {
		log.Fatal(err)
	}
	var resultado bool
	if total > 0 {
		resultado = true
	} else {
		resultado = false
	}
	js, err := json.Marshal(SomeStruct{resultado})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

// CONCEPTO EDITAR
func ConceptoEditar(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/concepto/conceptoEditar.html")
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	selDB, err := db.Query("SELECT * FROM concepto WHERE codigo=$1", Codigo)
	if err != nil {
		panic(err.Error())
	}
	emp := concepto{}
	for selDB.Next() {
		var codigo string
		var nombre string
		err = selDB.Scan(&codigo, &nombre)
		if err != nil {
			panic(err.Error())
		}
		emp.Codigo = codigo
		emp.Nombre = nombre
	}
	varmap := map[string]interface{}{
		"emp":     emp,
		"hosting": ruta,
	}
	tmp.Execute(w, varmap)
}

// CONCEPTO ACTUAL
func ConceptoActual(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	t := concepto{}
	var res []concepto
	err := db.Get(&t,"SELECT * FROM concepto where codigo=$1", Codigo)
	var siexiste bool
	siexiste=false
	switch err {
	case nil:
		log.Printf("centro found: %+v\n", t)
		siexiste=true
	case sql.ErrNoRows:
		log.Println("centro not found")
	default:
		log.Printf("centro error: %s\n", err)
	}

	if siexiste==true{
		log.Println("codigo nombre99" + t.Codigo + t.Nombre)
		res = append(res, t)
		data, err := json.Marshal(res)
		if err != nil {
			log.Fatal(err)
		}
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	} else {

		var slice []string
		slice = make([]string, 0)
		data, _ := json.Marshal(slice)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}


}

// CONCEPTO ACTUALIZAR
func ConceptoActualizar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		codigo := r.FormValue("Codigo")
		nombre := r.FormValue("Nombre")
		nombre = Titulo(nombre)
		insForm, err := db.Prepare("UPDATE concepto set	nombre=$2  " + " WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(codigo, nombre)
		log.Println("Registro Actualizado:" + codigo + "," +
			"" + nombre)
	}
	http.Redirect(w, r, "/ConceptoLista", 301)
}

// CONCEPTO BUSCAR
func ConceptoBuscar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	selDB, err := db.Query("SELECT codigo,"+
		"nombre FROM concepto where codigo LIKE '%' || $1 || '%' ORDER BY"+
		" codigo DESC", Codigo)
	if err != nil {
		panic(err.Error())
	}
	var resJson []conceptoJson
	var contar int
	contar = 0
	for selDB.Next() {
		contar++
		var id string
		var label string
		var value string
		var nombre string
		err = selDB.Scan(&id, &nombre)
		if err != nil {
			panic(err.Error())
		}
		value = id
		label = id + "  -  " + nombre
		resJson = append(resJson, conceptoJson{id, label, value, nombre})
	}
	if err := selDB.Err(); err != nil { // make sure that there was no issue during the process
		log.Println(err)
		return
	}
	if contar == 0 {
		var slice []string
		slice = make([]string, 0)
		data, _ := json.Marshal(slice)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	} else {
		data, _ := json.Marshal(resJson)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

// CONCEPTO BORRAR
func ConceptoBorrar(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/concepto/conceptoBorrar.html")
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	selDB, err := db.Query("SELECT * FROM concepto WHERE codigo=$1", Codigo)
	if err != nil {
		panic(err.Error())
	}
	emp := concepto{}
	for selDB.Next() {
		var codigo string
		var nombre string
		err = selDB.Scan(&codigo, &nombre)
		if err != nil {
			panic(err.Error())
		}
		emp.Codigo = codigo
		emp.Nombre = nombre
	}
	varmap := map[string]interface{}{
		"emp":     emp,
		"hosting": ruta,
	}
	tmp.Execute(w, varmap)
}

// CONCEPTO ELIMINAR
func ConceptoEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	emp := mux.Vars(r)["codigo"]
	//Codigo, _ := strconv.ParseInt(emp, 10, 0)
	delForm, err := db.Prepare("DELETE from concepto WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("Registro Eliminado" + emp)
	http.Redirect(w, r, "/ConceptoLista", 301)
}

// INICIA CONCEPTO PDF
func ConceptoPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	t := concepto{}
	var e  empresa=ListaEmpresa()
	var c  ciudad=TraerCiudad(e.Ciudad)
	err := db.Get(&t, "SELECT * FROM concepto where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	var buf bytes.Buffer
	var err1 error
	pdf := gofpdf.New("P", "mm", "Letter", cnFontDir)
	ene := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetHeaderFunc(func() {
		pdf.Image(imageFile("logo.png"), 20, 30, 40, 0, false,
			"", 0, "")
		pdf.SetY(17)

		pdf.SetFont("Arial", "", 10)
		pdf.CellFormat(190, 10, Mayuscula(e.Nombre), "0", 0,
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

		pdf.SetX(20)
		pdf.CellFormat(184, 6, "DATOS CONCEPTOS", "0", 0,
			"C", true, 0, "")
		pdf.Ln(8)
	})

	pdf.SetX(21)
	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.SetX(21)

	pdf.CellFormat(40, 4, "Codigo", "", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, t.Codigo, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Nombre:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Nombre, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)

	pdf.SetFooterFunc(func() {
		pdf.SetTextColor(0, 0, 0)
		pdf.SetY(252)
		pdf.SetFont("Arial", "", 9)
		pdf.SetX(20)

		// LINEA
		pdf.Line(20,259,204,259)
		pdf.Ln(6)
		pdf.SetX(20)
		pdf.CellFormat(90, 10, "Sadconf.com", "", 0,
			"L", false, 0, "")
		pdf.SetX(129)
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

// INICIA CONCEPTO TODOS PDF
func ConceptoTodosCabecera(pdf *gofpdf.Fpdf){
	pdf.SetFont("Arial", "", 10)
	// RELLENO TITULO
	pdf.SetY(50)
	pdf.SetFillColor(224,231,239)
	pdf.SetTextColor(0,0,0)
	pdf.Ln(7)
	pdf.SetX(20)
	pdf.CellFormat(184, 6, "No", "0", 0,
		"L", true, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(190, 6, "Codigo", "0", 0,
		"L", false, 0, "")
	pdf.SetX(55)
	pdf.CellFormat(190, 6, "Nombre", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}
func ConceptoTodosDetalle(pdf *gofpdf.Fpdf,miFila concepto, a int ){
	pdf.SetFont("Arial", "", 9)

	pdf.SetX(21)
	pdf.CellFormat(181, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(40, 4, Subcadena(miFila.Codigo,0,12), "", 0,
		"L", false, 0, "")
	pdf.SetX(55)
	pdf.CellFormat(40, 4, miFila.Nombre, "", 0,"L", false, 0, "")
	pdf.Ln(4)
}

func ConceptoTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	//	Codigo := mux.Vars(r)["codigo"]

	t := []concepto{}
	var e  empresa=ListaEmpresa()
	var c  ciudad=TraerCiudad(e.Ciudad)
	err := db.Select(&t, "SELECT * FROM concepto ORDER BY cast(codigo as integer) ")
	if err != nil {
		log.Fatalln(err)
	}
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
		pdf.CellFormat(190, 10, e.Telefono1+" "+e.Telefono2, "0", 0, "C", false, 0,
			"")
		pdf.Ln(4)
		pdf.CellFormat(190, 10, ene(c.NombreCiudad+ " - "+c.NombreDepartamento), "0", 0, "C", false, 0,
			"")
		pdf.Ln(6)
		pdf.CellFormat(190, 10, "DATOS CONCEPTO", "0", 0,
			"C", false, 0, "")
		pdf.Ln(10)
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
		pdf.CellFormat(90, 10, "Sadconf.com", "", 0,
			"L", false, 0, "")
		pdf.SetX(129)
		pdf.CellFormat(80, 10, fmt.Sprintf(" %d de {nb}", pdf.PageNo()), "",
			0, "R", false, 0, "")
	})

	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)
	pdf.SetX(30)

	ConceptoTodosCabecera(pdf)
	// tercera pagina

	for i, miFila := range t {
		var a = i + 1
		if math.Mod(float64(a),49)==0 {
			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			ConceptoTodosCabecera(pdf)
		}
		ConceptoTodosDetalle(pdf,miFila,a)
	}
	//BalancePieDePagina(pdf)

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}
// TERMINA DOCUMENTO TODOS PDF

// DOCUMENTO EXCEL
func ConceptoExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	t := []concepto{}
	var e  empresa=ListaEmpresa()
	var c  ciudad=TraerCiudad(e.Ciudad)
	err := db.Select(&t, "SELECT * FROM concepto ORDER BY cast(codigo as integer) ")
	if err != nil {
		log.Fatalln(err)
	}

	f := excelize.NewFile()

	// FUNCION ANCHO DE LA COLUMNA
	if err =f.SetColWidth("Sheet1", "A", "A", 13); err != nil {
		fmt.Println(err)
		return
	}
	if err =f.SetColWidth("Sheet1", "B", "B", 50); err != nil {
		fmt.Println(err)
		return
	}

	// FUNCION PARA UNIR DOS CELDAS TITULO
	if err = f.MergeCell("Sheet1", "A1", "B1"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A2", "B2"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A3", "B3"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A4", "B4"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A5", "B5"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A6", "B6"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A7", "B7"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A8", "B8"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A9", "B9"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A10", "B10"); err != nil {
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
	f.SetCellValue("Sheet1", "A6",(e.Telefono1+" - "+e.Telefono2))
	f.SetCellValue("Sheet1", "A7",(c.NombreCiudad+" - "+c.NombreDepartamento))
	f.SetCellValue("Sheet1", "A8","")
	f.SetCellValue("Sheet1", "A9","LISTADO DE CONCEPTOS")
	f.SetCellValue("Sheet1", "A10","")

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
	estiloNumeroDetalle, err := f.NewStyle(`{"number_format": 3,"font":{"bold":false,"italic":false,"family":"Arial","size":10,"color":"##000000"}}`)

	if err != nil {
		fmt.Println(err)
	}
	//cabecera
	f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel),"Codigo")
	f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel), "Nombre")


	f.SetCellStyle("Sheet1","A"+strconv.Itoa(filaExcel),"A"+strconv.Itoa(filaExcel),estiloCabecera)
	f.SetCellStyle("Sheet1","B"+strconv.Itoa(filaExcel),"B"+strconv.Itoa(filaExcel),estiloCabecera)
	filaExcel++

	for i, miFila := range t{
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel+i), Entero(miFila.Codigo))
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel+i), miFila.Nombre)

		f.SetCellStyle("Sheet1","A"+strconv.Itoa(filaExcel+i),"A"+strconv.Itoa(filaExcel+i),estiloNumeroDetalle)
		f.SetCellStyle("Sheet1","B"+strconv.Itoa(filaExcel+i),"B"+strconv.Itoa(filaExcel+i),estiloTexto)

		//van=i
	}

	// LIENA FINAL
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


// AVISO CRUZADO
//margin := 25.0
//markFontHt := 50.0
//markLineHt := pdf.PointToUnitConvert(markFontHt)
//markY := (297.0 - markLineHt) / 2.0
//ctrX := 210.0 / 2.0
//ctrY := 297.0 / 2.0
//pdf.SetFont("Arial", "B", markFontHt)
//pdf.SetTextColor(206, 216, 232)
//pdf.SetXY(margin, markY)
//pdf.TransformBegin()
//pdf.TransformRotate(45, ctrX, ctrY)
//pdf.CellFormat(0, markLineHt, "W A T E R M A R K   D E M O", "", 0, "C", false, 0, "")
//pdf.TransformEnd()
// FIN AVISO CRUZADO

// INICIA CUADRO EMPRESA
//pdf.SetY(20)
//pdf.SetX(20)
//pdf.SetDrawColor(119,134,153)
//pdf.CellFormat(184, 69, "", "1", 0, "C",
//false, 0, "")
//pdf.SetY(20)
// TERMINA CUADRO EMPRESA

// INICIA RELLENO TITULO COLOR AZUL CON BLANCO
//pdf.SetY(46)
//pdf.SetX(20)
//pdf.SetFillColor(59,99,146)
//pdf.SetDrawColor(119,134,153)
//pdf.SetTextColor(255,255,255)
//pdf.SetFont("Arial", "", 10)
//pdf.Ln(8)
//pdf.SetX(20)
//pdf.CellFormat(184, 5, "DATOS DEL PROVEEDOR", "0", 0,
//"L", true, 0, "")
//pdf.SetX(110)
//pdf.CellFormat(94, 5, "LUGAR DE ENTREGA O SERVICIO", "1", 0,
//"L", false, 0, "")
//pdf.Ln(8)
//pdf.SetX(20)
// TERMINA RELLENO TITULO COLOR AZUL CON BLANCO

//INICIA LINEA OSCURA Y BLANCA
//pdf.SetFont("Arial", "", 9)
//if math.Mod(float64(a),2)==0 {
//pdf.SetFillColor(224,231,239)
//pdf.SetTextColor(0,0,0)
//} else{
//pdf.SetFillColor(255,255,255)
//pdf.SetTextColor(0,0,0)
//}
//
//pdf.SetTextColor(0,0,0)
//pdf.SetX(21)
//var yinicial  float64
//yinicial=pdf.GetY()
//
//pdf.CellFormat(183, 4, strconv.Itoa(a), "", 0,
//"L", true, 0, "")
//pdf.SetX(30)
//pdf.CellFormat(40, 4, miFila.Codigoservicio, "", 0,
//"L", false, 0, "")
//var y float64
//y=pdf.GetY()
//pdf.SetX(42)
//pdf.MultiCell(80,4, Mayuscula(miFila.Nombreservicio), "","L", false)
//var yfinal float64
//yfinal=pdf.GetY()
//pdf.SetY(y)
//pdf.SetX(92)
//pdf.CellFormat(40, 4, Titulo(Subcadena(miFila.Unidadservicio, 0,6)), "", 0,
//"R", false, 0, "")
//pdf.SetX(121)
//pdf.CellFormat(40, 4, miFila.Cantidad, "", 0,
//"R", false, 0, "")
//pdf.SetX(143)
//pdf.CellFormat(40, 4, miFila.Precio, "", 0,
//"R", false, 0, "")
//pdf.SetX(163)
//pdf.CellFormat(40, 4, miFila.Subtotal, "", 0,
//"R", false, 0, "")
//pdf.Ln(yfinal-yinicial+3)
//INICIA LINEA OSCURA Y BLANCA

// INICIA LINEA
//pdf.Line(20,259,204,259)
//pdf.Ln(6)
//pdf.SetX(20)
// TERMINA LINEA

//INICIA CONTABILIZADO CRUZADO
//<a class="ribbon ribbon-top-right ribbon-amarillo ribbon-md"
//href="">CONTABILIZADO</a>
//TERMINA CONTABILIZADO CRUZADO
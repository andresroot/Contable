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
)

// ESTRUCTURA CARGO JSON
type cargoJson struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Nombre string `json:"nombre"`
}

// ESTRUCTURA CARGO
type cargo struct {
	Codigo      string
	Nombre      string
}

// CARGO LISTA
func CargoLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/cargo/cargoLista.html")
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM cargo ORDER BY cast(codigo as integer) ASC")
	if err != nil {
		panic(err.Error())
	}
	res := []cargo{}
	for selDB.Next() {
		var Codigo string
		var Nombre string
		err = selDB.Scan(&Codigo, &Nombre)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, cargo{Codigo, Nombre })
	}
	varmap := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
	}
	tmp.Execute(w, varmap)
}

// CARGO NUEVO
func CargoNuevo(w http.ResponseWriter, r *http.Request) {
	// TRAER COPIA DE EDITAR
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	emp := cargo{}
	if Codigo == "False"{
	} else {
		err := db.Get(&emp, "SELECT * FROM cargo where codigo=$1", Codigo)
		if err != nil {
			log.Fatalln(err)
		}
	}

	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/cargo/cargoNuevo.html")

	parametros := map[string]interface{}{
		"emp":     emp,
		"hosting": ruta,
		"codigo": Codigo,
	}
	tmp.Execute(w, parametros)
	// TERMINA TRAER COPIA DE EDITAR
}

// CARGO INSERTAR
func CargoInsertar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		Codigo := r.FormValue("Codigo")
		Nombre := r.FormValue("Nombre")
		Nombre = Titulo(Nombre)
		insForm, err := db.Prepare("INSERT INTO cargo(codigo, nombre)VALUES($1, $2)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(Codigo, Nombre)
		log.Println("Nuevo Registro:" + Codigo + "," + Nombre)
	}
	http.Redirect(w, r, "/CargoLista", 301)
}

// CARGO EXISTE
func CargoExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM cargo  WHERE codigo=$1", Codigo)
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

// CARGO EDITAR
func CargoEditar(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/cargo/cargoEditar.html")
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	selDB, err := db.Query("SELECT * FROM cargo WHERE codigo=$1", Codigo)
	if err != nil {
		panic(err.Error())
	}
	emp := cargo{}
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

// CARGO ACTUAL
func CargoActual(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	selDB, err := db.Query("SELECT * FROM cargo where codigo=$1", Codigo)
	if err != nil {
		panic(err.Error())
	}
	emp := cargo{}
	var res []cargo
	for selDB.Next() {
		var codigo string
		var nombre string
		err = selDB.Scan(&codigo, &nombre)
		if err != nil {
			panic(err.Error())
		}
		emp.Codigo = codigo
		emp.Nombre = nombre
		res = append(res, emp)
	}
	if err := selDB.Err(); err != nil { // make sure that there was no issue during the process
		log.Println(err)
		return
	}
	data, err := json.Marshal(res)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// CARGO ACTUALIZAR
func CargoActualizar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		codigo := r.FormValue("Codigo")
		nombre := r.FormValue("Nombre")
		nombre = Titulo(nombre)
		insForm, err := db.Prepare("UPDATE cargo set	nombre=$2  " + " WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(codigo, nombre)
		log.Println("Registro Actualizado:" + codigo + "," +
			"" + nombre)
	}
	http.Redirect(w, r, "/CargoLista", 301)
}

// CARGO BUSCAR
func CargoBuscar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	selDB, err := db.Query("SELECT codigo,"+
		"nombre FROM cargo where codigo LIKE '%' || $1 || '%' ORDER BY"+
		" codigo DESC", Codigo)
	if err != nil {
		panic(err.Error())
	}
	var resJson []cargoJson
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
		resJson = append(resJson, cargoJson{id, label, value, nombre})
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

// CARGO BORRAR
func CargoBorrar(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/cargo/cargoBorrar.html")
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	selDB, err := db.Query("SELECT * FROM cargo WHERE codigo=$1", Codigo)
	if err != nil {
		panic(err.Error())
	}
	emp := cargo{}
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

// CARGO ELIMINAR
func CargoEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	emp := mux.Vars(r)["codigo"]
	delForm, err := db.Prepare("DELETE from cargo WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("Registro Eliminado" + emp)
	http.Redirect(w, r, "/CargoLista", 301)
}

// INICIA CARGO PDF
func CargoPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	t := cargo{}
	var e  empresa=ListaEmpresa()
	var c  ciudad=TraerCiudad(e.Ciudad)
	err := db.Get(&t, "SELECT * FROM cargo where codigo=$1", Codigo)
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
		pdf.CellFormat(184, 6, "DATOS DEL CARGO", "0", 0,
			"C", true, 0, "")
		pdf.Ln(8)
	})

	pdf.SetTextColor(0,0,0)
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

// INICIA CARGO TODOS PDF
func CargoTodosCabecera(pdf *gofpdf.Fpdf){
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
	pdf.SetX(110)
	pdf.Ln(8)
}
func CargoTodosDetalle(pdf *gofpdf.Fpdf,miFila cargo, a int ){
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

func CargoTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	t := []cargo{}
	var e  empresa=ListaEmpresa()
	var c  ciudad=TraerCiudad(e.Ciudad)
	err := db.Select(&t, "SELECT * FROM cargo ORDER BY cast(codigo as integer) ")
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
		pdf.CellFormat(190, 10, "DATOS DE CARGOS", "0", 0,
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

	CargoTodosCabecera(pdf)
	// tercera pagina

	for i, miFila := range t {
		var a = i + 1
		if math.Mod(float64(a),49)==0 {
			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			DocumentoTodosCabecera(pdf)
		}
		CargoTodosDetalle(pdf,miFila,a)
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
func CargoExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	t := []cargo{}
	var e  empresa=ListaEmpresa()
	var c  ciudad=TraerCiudad(e.Ciudad)
	err := db.Select(&t, "SELECT * FROM cargo ORDER BY cast(codigo as integer) ")
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

	// FUNCION PARA UNIR DOS CELDAS
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
	f.SetCellValue("Sheet1", "A9","DATOS DE CARGOS")
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
	}

	// LINEA FINAL
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

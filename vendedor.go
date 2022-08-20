package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
	_ "github.com/lib/pq"
	"html/template"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
)

// INICIA VENDEDOR ESTRUCTURA JSON
type vendedorJson struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Nombre string `json:"nombre"`
}

// INICIA VENDEDOR ESTRUCTURA
type vendedor struct {
	Codigo          string
	Nombre          string
	Comision        string
}

// INICIA VENDEDOR LISTA
func VendedorLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/vendedor/vendedorLista.html")
	db := dbConn()
	res := []vendedor{}
	db.Select(&res, "SELECT * FROM vendedor ORDER BY cast(codigo as integer ) ASC")
	varmap := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
	}
	tmp.Execute(w, varmap)
}

// INICIA VENDEDOR NUEVO
func VendedorNuevo(w http.ResponseWriter, r *http.Request) {
	// TRAER COPIA DE EDITAR
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	Panel := mux.Vars(r)["panel"]
	Elemento := mux.Vars(r)["elemento"]
	emp := vendedor{}
	if Codigo == "False"{
	} else {
		if Elemento == "False" {
			err := db.Get(&emp, "SELECT * FROM vendedor where codigo=$1", Codigo)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}

	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/vendedor/vendedorNuevo.html",
		"vista/vendedor/autocompletaTerceroVendedor.html",)

	parametros := map[string]interface{}{
		"emp":                     emp,
		"hosting":                 ruta,
		"codigo":                  Codigo,
		"Panel":                   Panel,
		"Elemento":                Elemento,
	}
	tmp.Execute(w, parametros)
}

// INICIA VENDEDOR INSERTAR
func VendedorInsertar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	err := r.ParseForm()
	if err != nil {
		panic(err.Error())
	}
	var t vendedor
	err = decoder.Decode(&t, r.PostForm)
	if err != nil {
		// Handle error
		panic(err.Error())
	}
	var q string
	q = "insert into vendedor ("
	q += "Codigo,"
	q += "Nombre,"
	q += "Comision"
	q += " ) values("
	q += parametros(3)
	q += " ) "

	log.Println("Cadena SQL " + q)
	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	// INICIA GRABAR VENDEDOR INSERTAR
	t.Nombre = Titulo(t.Nombre)
	// TERMINA VENDEDOR GRABAR INSERTAR
	_, err = insForm.Exec(
		Quitacoma(t.Codigo),
		t.Nombre,
		t.Comision)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/VendedorLista", 301)
}

// INICIA VENDEDOR EXISTE
func VendedorExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM vendedor  WHERE codigo=$1", Codigo)
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

// INICIA VENDEDOR EDITAR
func VendedorEditar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/vendedor/vendedorEditar.html",
		"vista/autocompleta/autocompletaTercero.html")
	db := dbConn()
	t := vendedor{}
	err := db.Get(&t, "SELECT * FROM vendedor where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("codigo nombre99" + t.Codigo + t.Nombre)
	varmap := map[string]interface{}{
		// INICIA VENDEDOR EDITAR AUTOCOMPLETADO
		"emp":                     t,
		"hosting":                 ruta,
		"ciudad":                  ListaCiudad(),
	}
	tmp.Execute(w, varmap)
}

// VENDEDOR ACTUAL
func VendedorActual(w http.ResponseWriter, r *http.Request) {
	var tabla="vendedor"
	emp := buscarParametro{}
	var res []buscarParametro
	db := dbConn()
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var tempParametro buscarParametro
	err = json.Unmarshal(b, &tempParametro)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	err = db.Get(&emp, "SELECT codigo,nombre FROM "+tabla+" where codigo=$1 or UPPER(nombre)=$2", tempParametro.Codigo,Mayuscula(tempParametro.Nombre))
	switch err {
	case nil:
		log.Printf("tercero found: %+v\n", emp)
		res = append(res, emp)
		data, err := json.Marshal(res)
		if err != nil {
			log.Fatalln(err)
		}
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)

	case sql.ErrNoRows:
		log.Println("tercero NOT found, no error")
		var slice []string
		slice = make([]string, 0)
		data, _ := json.Marshal(slice)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	default:
		log.Printf("tercero error: %s\n", err)
	}
}

// INICIA VENDEDOR ACTUALIZAR
func VendedorActualizar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	err := r.ParseForm()
	if err != nil {
		panic(err.Error())
		// Handle error
	}
	var t vendedor
	// r.PostForm is a map of our POST form values
	err = decoder.Decode(&t, r.PostForm)
	if err != nil {
		// Handle error
		panic(err.Error())
	}
	var q string
	q = "UPDATE vendedor set "
	q += "Nombre=$2,"
	q += "Comision=$3"
	q += " where "
	q += "Codigo=$1"

	log.Println("cadena" + q)

	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	// INICIA GRABAR VENDEDOR ACTUALIZAR
	t.Nombre = Titulo(t.Nombre)
	// TERMINA GRABAR VENDEDOR ACTUALIZAR
	_, err = insForm.Exec(
		Quitacoma(t.Codigo),
		t.Nombre,
		t.Comision)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/VendedorLista", 301)

}

// VENDEDOR BUSCAR NOMBRE
func VendedorBuscarCrear(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := Mayuscula(mux.Vars(r)["codigo"])
	selDB, err := db.Query("SELECT codigo,"+
		"nombre FROM vendedor where codigo LIKE '%' || $1 || '%'  or  UPPER(nombre) LIKE '%' || $1 || '%' ORDER BY"+
		" codigo ASC", Codigo)
	if err != nil {
		panic(err.Error())
	}
	var resJson []documentoJson
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
		label = id + " - " + nombre
		resJson = append(resJson, documentoJson{id, label, value, nombre})
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

// INICIA VENDEDOR BUSCAR
func VendedorBuscar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := Mayuscula(mux.Vars(r)["codigo"])
	var tabla="vendedor"
	selDB, err := db.Query("SELECT codigo,"+
		"nombre FROM "+tabla+" where codigo LIKE '%' || $1 || '%'  or  Upper(nombre) LIKE '%' || $1 || '%' ORDER BY codigo ASC", Codigo)
	if err != nil {
		panic(err.Error())
	}
	var resJson []vendedorJson
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
		value = nombre
		label = id+"-"+nombre
		resJson = append(resJson, vendedorJson{id, label, value, nombre})
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

// INICIA VENDEDOR BORRAR
func VendedorBorrar(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/vendedor/vendedorBorrar.html")
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio vendedor borrar" + Codigo)
	db := dbConn()
	t := vendedor{}
	err := db.Get(&t, "SELECT * FROM vendedor where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("codigo nombre99 borrar" + t.Codigo + t.Nombre)
	varmap := map[string]interface{}{
		// INICIA VENDEDOR BORRAR AUTOCOMPLETADO
		"emp":                     t,
		"hosting":                 ruta,
		// TERMINA VENDEDOR BORRAR AUTOCOMPLETADO
	}
	tmp.Execute(w, varmap)
}

// INICIA VENDEDOR ELIMINAR
func VendedorEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	emp := mux.Vars(r)["codigo"]
	delForm, err := db.Prepare("DELETE from vendedor WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("Registro Eliminado" + emp)
	http.Redirect(w, r, "/VendedorLista", 301)
}
// TERMINA VENDEDOR ELIMINAR

// INICIA VENDEDOR PDF
func VendedorPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	t := vendedor{}
	var e  empresa=ListaEmpresa()
	var c  ciudad=TraerCiudad(e.Ciudad)
	err := db.Get(&t, "SELECT * FROM vendedor where codigo=$1", Codigo)
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
		pdf.CellFormat(184, 6, "DATOS DEL VENDEDOR", "0", 0,
			"C", true, 0, "")
		pdf.Ln(8)
	})

	pdf.SetTextColor(0,0,0)
	pdf.SetX(21)
	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Codigo:", "", 0,
		"", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, Coma(t.Codigo), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Nombre:", "", 0,
		"", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, t.Nombre, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Comision %:", "", 0,
		"", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, t.Comision, "", 0,
		"", false, 0, "")

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

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

// INICIA VENDEDOR TODOS PDF
func VendedorTodosCabecera(pdf *gofpdf.Fpdf){
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
	pdf.SetX(60)
	pdf.CellFormat(190, 6, "Nombre", "0", 0,
		"L", false, 0, "")
	pdf.SetX(130)
	pdf.CellFormat(190, 6, "Comision", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}

func VendedorTodosDetalle(pdf *gofpdf.Fpdf,miFila vendedor, a int ){
	pdf.SetFont("Arial", "", 9)

	pdf.SetX(21)
	pdf.CellFormat(181, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(40, 4, Coma(miFila.Codigo), "", 0,
		"L", false, 0, "")
	pdf.SetX(60)
	pdf.CellFormat(40, 4, miFila.Nombre, "", 0,
		"L", false, 0, "")
	pdf.SetX(130)
	pdf.CellFormat(40, 4, miFila.Comision, "", 0,
		"L", false, 0, "")
	pdf.Ln(4)
}

func VendedorTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	//	Codigo := mux.Vars(r)["codigo"]

	t := []vendedor{}
	var e  empresa=ListaEmpresa()
	var c  ciudad=TraerCiudad(e.Ciudad)
	err := db.Select(&t, "SELECT * FROM vendedor ORDER BY cast(codigo as integer) ")
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
		pdf.CellFormat(190, 10, "DATOS DEL VENDEDOR", "0", 0,
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

	VendedorTodosCabecera(pdf)
	// tercera pagina

	for i, miFila := range t {
		var a = i + 1
		if math.Mod(float64(a),49)==0 {
			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			VendedorTodosCabecera(pdf)
		}
		VendedorTodosDetalle(pdf,miFila,a)
	}
	//BalancePieDePagina(pdf)

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}
// TERMINA VENDEDOR TODOS PDF

// VENDEDOR EXCEL
func VendedorExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	t := []vendedor{}
	var e  empresa=ListaEmpresa()
	var c  ciudad=TraerCiudad(e.Ciudad)
	err := db.Select(&t, "SELECT * FROM vendedor ORDER BY cast(codigo as integer) ")
	if err != nil {
		log.Fatalln(err)
	}

	f := excelize.NewFile()

	// FUNCION ANCHO DE LA COLUMNA
	if err =f.SetColWidth("Sheet1", "A", "A", 20); err != nil {
		fmt.Println(err)
		return
	}
	if err =f.SetColWidth("Sheet1", "B", "B", 50); err != nil {
		fmt.Println(err)
		return
	}
	if err =f.SetColWidth("Sheet1", "C", "C", 10); err != nil {
		fmt.Println(err)
		return
	}


	// FUNCION PARA UNIR DOS CELDAS
	if err = f.MergeCell("Sheet1", "A1", "C1"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A2", "C2"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A3", "C3"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A4", "C4"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A5", "C5"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A6", "C6"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A7", "C7"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A8", "C8"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A9", "C9"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A10", "C10"); err != nil {
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
	f.SetCellValue("Sheet1", "A9","LISTADO DE VENDEDORES")

	f.SetCellStyle("Sheet1","A1","A1",estiloTitulo)
	f.SetCellStyle("Sheet1","A2","A2",estiloTitulo)
	f.SetCellStyle("Sheet1","A3","A3",estiloTitulo)
	f.SetCellStyle("Sheet1","A4","A4",estiloTitulo)
	f.SetCellStyle("Sheet1","A5","A5",estiloTitulo)
	f.SetCellStyle("Sheet1","A6","A6",estiloTitulo)
	f.SetCellStyle("Sheet1","A7","A7",estiloTitulo)
	f.SetCellStyle("Sheet1","A9","A8",estiloTitulo)

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
	f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel), "Comision")


	f.SetCellStyle("Sheet1","A"+strconv.Itoa(filaExcel),"A"+strconv.Itoa(filaExcel),estiloCabecera)
	f.SetCellStyle("Sheet1","B"+strconv.Itoa(filaExcel),"B"+strconv.Itoa(filaExcel),estiloCabecera)
	f.SetCellStyle("Sheet1","C"+strconv.Itoa(filaExcel),"C"+strconv.Itoa(filaExcel),estiloCabecera)
	filaExcel++


	for i, miFila := range t{
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel+i), Entero(miFila.Codigo))
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel+i), miFila.Nombre)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel+i), Entero(miFila.Comision))


		f.SetCellStyle("Sheet1","A"+strconv.Itoa(filaExcel+i),"A"+strconv.Itoa(filaExcel+i),estiloNumeroDetalle)
  	f.SetCellStyle("Sheet1","B"+strconv.Itoa(filaExcel+i),"B"+strconv.Itoa(filaExcel+i),estiloTexto)
		f.SetCellStyle("Sheet1","C"+strconv.Itoa(filaExcel+i),"C"+strconv.Itoa(filaExcel+i),estiloNumeroDetalle)

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

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
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
)

// SUBGRUPO TABLA
type Subgrupo struct {
	Codigo      string
	Nombre      string
	Grupo		string
}
// SUBGRUPO BUSCAR
func SubgrupoBuscar(w http.ResponseWriter, r *http.Request) {
	var tabla="subgrupo"
	db := dbConn()
	Codigo := Mayuscula( mux.Vars(r)["codigo"])
	selDB, err := db.Query("SELECT codigo,"+
		"nombre FROM "+tabla+" where codigo LIKE '%' || $1 || '%'  or  UPPER(nombre) LIKE '%' || $1 || '%'  ORDER BY"+
		" codigo ASC", Codigo)
	if err != nil {
		panic(err.Error())
	}
	var resJson []buscarJson
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
		label =id+" - "+nombre
		resJson = append(resJson, buscarJson{id, label, value, nombre})
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
func SubgrupoActual(w http.ResponseWriter, r *http.Request) {
	var tabla="subgrupo"
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
// SUBGRUPO LISTA
func SubgrupoLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/subgrupo/subgrupoLista.html")
	db := dbConn()

	res := []Subgrupo{}
	db.Select(&res, "SELECT * FROM subgrupo ORDER BY cast(codigo as integer) ASC")
	log.Println("Lista Grupo 1" )

	varmap := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
	}
	tmp.Execute(w, varmap)
}

// SUBGRUPO NUEVO
func SubgrupoNuevo(w http.ResponseWriter, r *http.Request) {
	// TRAER COPIA DE EDITAR
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	Panel := mux.Vars(r)["panel"]
	Elemento := mux.Vars(r)["elemento"]
	emp := Subgrupo{}
	if Codigo == "False"{
	} else {
		if Elemento == "False"{
			err := db.Get(&emp, "SELECT * FROM subgrupo where codigo=$1", Codigo)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}

	//Codigo := mux.Vars(r)["codigo"]
	//emp := Subgrupo{}
	//if Codigo == "False"{
	//} else {
	//	err := db.Get(&emp, "SELECT * FROM subgrupo where codigo=$1", Codigo)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//}

	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/autocompleta/autocompletaGrupo.html",
		"vista/subgrupo/subgrupoNuevo.html")

	parametros := map[string]interface{}{
		"emp":     emp,
		"hosting": ruta,
		"grupo": ListaGrupo(),
		"codigo": 	Codigo,
		"Panel":    Panel,
		"Elemento": Elemento,
	}
	tmp.Execute(w, parametros)
	// TERMINA TRAER COPIA DE EDITAR
}

// SUBGRUPO INSERTAR
func SubgrupoInsertar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		Codigo := r.FormValue("Codigo")
		Nombre := r.FormValue("Nombre")
		Grupo := r.FormValue("Grupo")

		Nombre = Titulo(Nombre)
		insForm, err := db.Prepare("INSERT INTO subgrupo(codigo, nombre, grupo)VALUES($1, $2, $3)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(Codigo, Nombre, Grupo)
		log.Println("Nuevo Registro:" + Codigo + "," + Nombre)
	}
	http.Redirect(w, r, "/SubgrupoLista", 301)
}

// SUBGRUPO EXISTE
func SubgrupoExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM subgrupo  WHERE codigo=$1", Codigo)
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

// SUBGRUPO EDITAR
func SubgrupoEditar(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/subgrupo/subgrupoEditar.html")
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	selDB, err := db.Query("SELECT * FROM subgrupo WHERE codigo=$1", Codigo)
	if err != nil {
		panic(err.Error())
	}
	emp := Subgrupo{}
	for selDB.Next() {
		var codigo string
		var nombre string
		var grupo string
		err = selDB.Scan(&codigo, &nombre, &grupo)
		if err != nil {
			panic(err.Error())
		}
		emp.Codigo = codigo
		emp.Nombre = nombre
		emp.Grupo = grupo
	}
	varmap := map[string]interface{}{
		"emp":     emp,
		"hosting": ruta,
		"grupo": ListaGrupo(),
	}
	//vistaSubgrupo.ExecuteTemplate(w, "SubgrupoEditar", varmap)
	tmp.Execute(w, varmap)
}

// SUBGRUPO ACTUALIZAR
func SubgrupoActualizar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		codigo := r.FormValue("Codigo")
		nombre := r.FormValue("Nombre")

		nombre = Titulo(nombre)
		insForm, err := db.Prepare("UPDATE subgrupo set	nombre=$2  " + " WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(codigo, nombre)
		log.Println("Registro Actualizado:" + codigo + "," +
			"" + nombre)
	}
	http.Redirect(w, r, "/SubgrupoLista", 301)
}

// SUBGRUPO BORRAR
func SubgrupoBorrar(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/subgrupo/subgrupoBorrar.html")
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	t := Subgrupo{}
	err := db.Get(&t, "SELECT * FROM subgrupo where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	varmap := map[string]interface{}{
		"emp":     t,
		"hosting": ruta,
		"grupo": ListaGrupo(),
	}
	tmp.Execute(w, varmap)
}

// SUBGRUPO ELIMINAR
func SubgrupoEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	emp := mux.Vars(r)["codigo"]
	//Codigo, _ := strconv.ParseInt(emp, 10, 0)
	delForm, err := db.Prepare("DELETE from subgrupo WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("Registro Eliminado" + emp)
	http.Redirect(w, r, "/SubgrupoLista", 301)
}

// SUBGRUPO PDF
func SubgrupoPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	t := Subgrupo{}
	var e  empresa=ListaEmpresa()
	var c  ciudad=TraerCiudad(e.Ciudad)
	err := db.Get(&t, "SELECT * FROM subgrupo where codigo=$1", Codigo)
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
		pdf.CellFormat(184, 6, "DATOS SUBGRUPO", "0", 0,
			"C", true, 0, "")
		pdf.Ln(8)
	})

	pdf.SetTextColor(0,0,0)
	pdf.SetX(21)
	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.SetX(21)

	pdf.CellFormat(40, 4, "Grupo", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, TraerGrupo(t.Grupo), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Codigo", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Codigo, "", 0,
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
// TERMINA SUBGRUPO PDF

// INICIA SUBGRUPO TODOS PDF
func SubgrupoTodosCabecera(pdf *gofpdf.Fpdf){
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
	pdf.SetX(50)
	pdf.CellFormat(190, 6, "Nombre", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}

func SubgrupoTodosDetalle(pdf *gofpdf.Fpdf,miFila Subgrupo, a int ){
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(21)
	pdf.CellFormat(181, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(40, 4, Subcadena(miFila.Codigo,0,12), "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, miFila.Nombre, "", 0,"L", false, 0, "")
	pdf.Ln(4)
}

func SubgrupoTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	t := []Subgrupo{}
	var e  empresa=ListaEmpresa()
	var c  ciudad=TraerCiudad(e.Ciudad)
	err := db.Select(&t, "SELECT * FROM subgrupo ORDER BY cast(codigo as integer) ")
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
		pdf.CellFormat(190, 10, "Nit No. "+Coma(e.Codigo)+" - "+e.Dv, "0", 0, "C",
			false, 0, "")
		pdf.Ln(4)
		pdf.CellFormat(190, 10, e.Iva+" - "+e.ReteIva, "0", 0, "C", false, 0,
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
		pdf.CellFormat(190, 10, "DATOS SUBGRUPO", "0", 0,
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

	SubgrupoTodosCabecera(pdf)
	// tercera pagina

	for i, miFila := range t {
		var a = i + 1
		if math.Mod(float64(a),49)==0 {
			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			SubgrupoTodosCabecera(pdf)
		}
		SubgrupoTodosDetalle(pdf,miFila,a)
	}

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}
// TERMINA DOCUMENTO TODOS PDF

// DOCUMENTO EXCEL
func SubgrupoExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	t := []Subgrupo{}
	var e  empresa=ListaEmpresa()
	var c  ciudad=TraerCiudad(e.Ciudad)
	err := db.Select(&t, "SELECT * FROM subgrupo ORDER BY cast(codigo as integer) ")
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
	f.SetCellValue("Sheet1", "A9","LISTADO DE SUBGRUPOS")
	f.SetCellValue("Sheet1", "A10","")

	f.SetCellStyle("Sheet1","A1","A1",estiloTitulo)
	f.SetCellStyle("Sheet1","A2","A2",estiloTitulo)
	f.SetCellStyle("Sheet1","A3","A3",estiloTitulo)
	f.SetCellStyle("Sheet1","A4","A4",estiloTitulo)
	f.SetCellStyle("Sheet1","A5","A5",estiloTitulo)
	f.SetCellStyle("Sheet1","A6","A6",estiloTitulo)
	f.SetCellStyle("Sheet1","A7","A7",estiloTitulo)
	f.SetCellStyle("Sheet1","A8","A8",estiloTitulo)
	f.SetCellStyle("Sheet1","A9","A7",estiloTitulo)
	f.SetCellStyle("Sheet1","A10","A8",estiloTitulo)

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
	//estiloNumeroDetalle, err := f.NewStyle(`{"number_format": 3,"font":{"bold":false,"italic":false,"family":"Arial","size":10,"color":"##000000"}}`)

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
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel+i), miFila.Codigo)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel+i), miFila.Nombre)

		f.SetCellStyle("Sheet1","A"+strconv.Itoa(filaExcel+i),"A"+strconv.Itoa(filaExcel+i),estiloTexto)
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

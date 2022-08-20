package main

import (
	"bytes"
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
	"log"
	"math"
	"net/http"
	"strconv"
)

type empresaJson struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Nombre string `json:"nombre"`
}

// INICIA EMPRESA ESTRUCTURA
type empresa struct {
	Codigo              string
	Dv                  string
	Nombre              string
	Iva                 string
	ReteIva             string
	Direccion           string
	ActividadIca        string
	Telefono1           string
	Telefono2           string
	Email1              string
	Email2              string
	Activa              string
	Licencia            string
	RepresentanteDv     string
	RepresentanteNombre string
	ContadorDv          string
	ContadorNombre      string
	RevisorDv           string
	RevisorNombre       string
	Ciudad              string
	ContadorNit         string
	Documento           string
	Fiscal              string
	Regimen             string
	RepresentanteNit    string
	RevisorNit          string
	Tipo                string
	Modulo              string
	Primernombre        string
	Segundonombre       string
	Primerapellido      string
	Segundoapellido     string
	Juridica            string
	Periodo             string
	PeriodoNombre       string
}

// INICIA EMPRESA LISTA
func EmpresaLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/empresa/empresaLista.html")
	db := dbConn()

	listado := []empresa{}
	err := db.Select(&listado, "SELECT * FROM empresa")
	if err != nil {
		panic(err.Error())
	}
	log.Println("inicio listado")
	for _, elem := range listado {
		log.Println("listado" + elem.Codigo)
	}

	varmap := map[string]interface{}{
		"listado": listado,
		"hosting": ruta,
	}
	log.Println("Error empresa888")
	tmp.Execute(w, varmap)
}

// INICIA EMPRESA NUEVO
func EmpresaNuevo(w http.ResponseWriter, r *http.Request) {
	// TRAER COPIA DE EDITAR
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	t := empresa{}
	if Codigo == "False" {
	} else {
		err := db.Get(&t, "SELECT * FROM empresa where codigo=$1", Quitacoma(Codigo))
		if err != nil {
			log.Fatalln(err)
		}
	}

	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/empresa/empresaNuevo.html",
		"vista/empresa/autocompletaTercero.html",
		"vista/empresa/autocompletaPeriodo.html")

	parametros := map[string]interface{}{
		"emp":                     t,
		"hosting":                 ruta,
		"ciudad":                  ListaCiudad(),
		"tipoorganizacion":        ListaTipoOrganizacion(),
		"regimenfiscal":           ListaRegimenFiscal(),
		"responsabilidadfiscal":   ListaResponsabilidadFiscal(),
		"documentoidentificacion": ListaDocumentoIdentificacion(),
		"codigo":                  Codigo,
	}
	tmp.Execute(w, parametros)
}

// INICIA EMPRESA INSERTAR
func EmpresaInsertar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	err := r.ParseForm()
	if err != nil {
		panic(err.Error())
	}
	var t empresa
	err = decoder.Decode(&t, r.PostForm)
	if err != nil {
		// Handle error
		panic(err.Error())
	}
	var q string
	q = "insert into empresa ("
	q += "Codigo,"
	q += "Dv,"
	q += "Nombre,"
	q += "Iva,"
	q += "ReteIva,"
	q += "Direccion,"
	q += "ActividadIca,"
	q += "Telefono1,"
	q += "Telefono2,"
	q += "Email1,"
	q += "Email2,"
	q += "Activa,"
	q += "Licencia,"
	q += "RepresentanteDv,"
	q += "RepresentanteNombre,"
	q += "ContadorDv,"
	q += "ContadorNombre,"
	q += "RevisorDv,"
	q += "RevisorNombre,"
	q += "Ciudad,"
	q += "ContadorNit,"
	q += "Documento,"
	q += "Fiscal,"
	q += "Regimen,"
	q += "RepresentanteNit,"
	q += "RevisorNit,"
	q += "Tipo,"
	q += "Modulo,"
	q += "Primernombre,"
	q += "Segundonombre,"
	q += "Primerapellido,"
	q += "Segundoapellido,"
	q += "Juridica,"
	q += "Periodo"
	q += " ) values("
	q += parametros(34)
	q += " ) "

	log.Println("Cadena SQL " + q)
	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	// INICIA GRABAR EMPRESA INSERTAR
	t.Codigo = Quitacoma(t.Codigo)
	t.Nombre = Mayuscula(t.Nombre)
	t.Direccion = Titulo(t.Direccion)
	t.ActividadIca = Titulo(t.ActividadIca)
	t.Licencia = Mayuscula(t.Licencia)
	t.Email1 = Minuscula(t.Email1)
	t.Email2 = Minuscula(t.Email2)
	t.RepresentanteNit = Quitacoma(t.RepresentanteNit)
	t.ContadorNit = Quitacoma(t.ContadorNit)
	t.RevisorNit = Quitacoma(t.RevisorNit)
	t.Primernombre = Titulo(t.Primernombre)
	t.Segundonombre = Titulo(t.Segundonombre)
	t.Primerapellido = Titulo(t.Primerapellido)
	t.Segundoapellido = Titulo(t.Segundoapellido)
	t.Juridica = Mayuscula(t.Juridica)

	if t.Tipo == "2" {
		t.Primernombre = Titulo(t.Primernombre)
		t.Segundonombre = Titulo(t.Segundonombre)
		t.Primerapellido = Titulo(t.Primerapellido)
		t.Segundoapellido = Titulo(t.Segundoapellido)
		t.Nombre = t.Primernombre + " " + t.Segundonombre + " " + t.Primerapellido + " " + t.Segundoapellido
	} else {
		t.Juridica = Mayuscula(t.Juridica)
		t.Nombre = Mayuscula(t.Juridica)
	}

	_, err = insForm.Exec(
		t.Codigo,
		t.Dv,
		t.Nombre,
		t.Iva,
		t.ReteIva,
		t.Direccion,
		t.ActividadIca,
		t.Telefono1,
		t.Telefono2,
		t.Email1,
		t.Email2,
		t.Activa,
		t.Licencia,
		t.RepresentanteDv,
		t.RepresentanteNombre,
		t.ContadorDv,
		t.ContadorNombre,
		t.RevisorDv,
		t.RevisorNombre,
		t.Ciudad,
		t.ContadorNit,
		t.Documento,
		t.Fiscal,
		t.Regimen,
		t.RepresentanteNit,
		t.RevisorNit,
		t.Tipo,
		t.Modulo,
		t.Primernombre,
		t.Segundonombre,
		t.Primerapellido,
		t.Segundoapellido,
		t.Juridica,
		t.Periodo)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/EmpresaLista", 301)
}

// TERMINA EMPRESA INSERTAR

// INICIA EMPRESA BUSCAR
func EmpresaBuscar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	selDB, err := db.Query("SELECT codigo,"+
		"nombre FROM empresa where codigo LIKE '%' || $1 || '%'  or  nombre LIKE '%' || $1 || '%' ORDER BY"+
		" codigo DESC", Codigo)
	if err != nil {
		panic(err.Error())
	}
	var resJson []empresaJson
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
		label = id + " " + nombre
		resJson = append(resJson, empresaJson{id, label, value, nombre})
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

// TERMINA EMPRESA BUSCAR

// INICIA EMPRESA EXISTE
func EmpresaExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	Codigo = Quitacoma(Codigo)
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM empresa  WHERE codigo=$1", Codigo)
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

// INICIA EMPRESA ACTUAL
func EmpresaActual(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	t := empresa{}
	var res []empresa
	err := db.Get(&t, "SELECT * FROM empresa where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("codigo nombre99" + t.Codigo + t.Nombre)
	res = append(res, t)
	data, err := json.Marshal(res)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// INICIA EMPRESA EDITAR
func EmpresaEditar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio empresa editar" + Codigo)
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/empresa/empresaEditar.html",
		"vista/empresa/autocompletaTercero.html",
		"vista/empresa/autocompletaPeriodo.html")
	db := dbConn()
	t := empresa{}
	err := db.Get(&t, "SELECT * FROM empresa where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("codigo nombre99" + t.Codigo + t.Nombre)
	varmap := map[string]interface{}{
		// INICIA EMPRESA EDITAR AUTOCOMPLETADO
		"emp":                     t,
		"hosting":                 ruta,
		"ciudad":                  ListaCiudad(),
		"tipoorganizacion":        ListaTipoOrganizacion(),
		"regimenfiscal":           ListaRegimenFiscal(),
		"responsabilidadfiscal":   ListaResponsabilidadFiscal(),
		"documentoidentificacion": ListaDocumentoIdentificacion(),
		"periodo":                 ListaPeriodo(),
	}
	tmp.Execute(w, varmap)
}

// INICIA EMPRESA ACTUALIZAR
func EmpresaActualizar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	err := r.ParseForm()
	if err != nil {
		panic(err.Error())
		// Handle error
	}
	var t empresa
	// r.PostForm is a map of our POST form values
	err = decoder.Decode(&t, r.PostForm)
	if err != nil {
		// Handle error
		panic(err.Error())
	}
	var q string
	q = "UPDATE empresa set "
	q += "Dv=$2,"
	q += "Nombre=$3,"
	q += "Iva=$4,"
	q += "ReteIva=$5,"
	q += "Direccion=$6,"
	q += "ActividadIca=$7,"
	q += "Telefono1=$8,"
	q += "Telefono2=$9,"
	q += "Email1=$10,"
	q += "Email2=$11,"
	q += "Activa=$12,"
	q += "Licencia=$13,"
	q += "RepresentanteDv=$14,"
	q += "RepresentanteNombre=$15,"
	q += "ContadorDv=$16,"
	q += "ContadorNombre=$17,"
	q += "RevisorDv=$18,"
	q += "RevisorNombre=$19,"
	q += "Ciudad=$20,"
	q += "ContadorNit=$21,"
	q += "Documento=$22,"
	q += "Fiscal=$23,"
	q += "Regimen=$24,"
	q += "RepresentanteNit=$25,"
	q += "RevisorNit=$26,"
	q += "Tipo=$27,"
	q += "Modulo=$28,"
	q += "Primernombre=$29,"
	q += "Segundonombre=$30,"
	q += "Primerapellido=$31,"
	q += "Segundoapellido=$32,"
	q += "Juridica=$33,"
	q += "Periodo=$34"
	q += " where "
	q += "Codigo=$1"
	log.Println("cadena" + q)

	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	// INICIA GRABAR EMPRESA ACTUALIZAR
	t.Codigo = Quitacoma(t.Codigo)
	t.Nombre = Mayuscula(t.Nombre)
	t.Direccion = Titulo(t.Direccion)
	t.Email1 = Minuscula(t.Email1)
	t.Email2 = Minuscula(t.Email2)
	t.ActividadIca = Titulo(t.ActividadIca)
	t.Licencia = Mayuscula(t.Licencia)
	t.RepresentanteNit = Quitacoma(t.RepresentanteNit)
	t.ContadorNit = Quitacoma(t.ContadorNit)
	t.RevisorNit = Quitacoma(t.RevisorNit)
	t.Primernombre = Titulo(t.Primernombre)
	t.Segundonombre = Titulo(t.Segundonombre)
	t.Primerapellido = Titulo(t.Primerapellido)
	t.Segundoapellido = Titulo(t.Segundoapellido)
	t.Juridica = Mayuscula(t.Juridica)

	if t.Tipo == "2" {
		t.Primernombre = Titulo(t.Primernombre)
		t.Segundonombre = Titulo(t.Segundonombre)
		t.Primerapellido = Titulo(t.Primerapellido)
		t.Segundoapellido = Titulo(t.Segundoapellido)
		t.Nombre = t.Primernombre + " " + t.Segundonombre + " " + t.Primerapellido + " " + t.Segundoapellido
	} else {
		t.Juridica = Mayuscula(t.Juridica)
		t.Nombre = Mayuscula(t.Juridica)
	}

	_, err = insForm.Exec(
		t.Codigo,
		t.Dv,
		t.Nombre,
		t.Iva,
		t.ReteIva,
		t.Direccion,
		t.ActividadIca,
		t.Telefono1,
		t.Telefono2,
		t.Email1,
		t.Email2,
		t.Activa,
		t.Licencia,
		t.RepresentanteDv,
		t.RepresentanteNombre,
		t.ContadorDv,
		t.ContadorNombre,
		t.RevisorDv,
		t.RevisorNombre,
		t.Ciudad,
		t.ContadorNit,
		t.Documento,
		t.Fiscal,
		t.Regimen,
		t.RepresentanteNit,
		t.RevisorNit,
		t.Tipo,
		t.Modulo,
		t.Primernombre,
		t.Segundonombre,
		t.Primerapellido,
		t.Segundoapellido,
		t.Juridica,
		t.Periodo)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/EmpresaLista", 301)
}

// INICIA EMPRESA BORRAR
func EmpresaBorrar(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/empresa/empresaBorrar.html")
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio empresa editar" + Codigo)
	db := dbConn()
	t := empresa{}
	err := db.Get(&t, "SELECT * FROM empresa where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("codigo nombre99" + t.Codigo + t.Nombre)
	varmap := map[string]interface{}{
		// INICIA EMPRESA BORRAR AUTOCOMPLETADO
		"emp":                     t,
		"hosting":                 ruta,
		"ciudad":                  ListaCiudad(),
		"tipoorganizacion":        ListaTipoOrganizacion(),
		"regimenfiscal":           ListaRegimenFiscal(),
		"responsabilidadfiscal":   ListaResponsabilidadFiscal(),
		"documentoidentificacion": ListaDocumentoIdentificacion(),
		// TERMINA EMPRESA BORRAR AUTOCOMPLETADO
	}
	tmp.Execute(w, varmap)
}

// INICIA EMPRESA ELIMINAR
func EmpresaEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	emp := mux.Vars(r)["codigo"]
	delForm, err := db.Prepare("DELETE from empresa WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("Registro Eliminado" + emp)
	http.Redirect(w, r, "/EmpresaLista", 301)
}

// INICIA EMPRESA PDF
func EmpresaPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	t := empresa{}
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	err := db.Get(&t, "SELECT * FROM empresa where codigo=$1", Codigo)
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
		pdf.SetY(15)

		pdf.SetFont("Arial", "", 10)
		pdf.CellFormat(190, 10, Mayuscula(e.Nombre), "0", 0,
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
		pdf.CellFormat(190, 10, e.Telefono1+" - "+e.Telefono2, "0", 0, "C", false, 0,
			"")
		pdf.Ln(4)
		pdf.CellFormat(190, 10, ene(c.NombreCiudad+" - "+c.NombreDepartamento), "0", 0, "C", false, 0,
			"")
		pdf.Ln(10)

		// RELLENO TITULO
		pdf.SetX(20)
		pdf.SetFillColor(224, 231, 239)
		pdf.SetTextColor(0, 0, 0)

		pdf.SetX(20)
		pdf.CellFormat(184, 6, "DATOS EMPRESA", "0", 0,
			"C", true, 0, "")
		pdf.Ln(8)
	})

	pdf.SetTextColor(0, 0, 0)
	pdf.SetX(21)
	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.SetX(21)

	pdf.CellFormat(40, 4, "Codigo", "", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, Coma(t.Codigo)+" - "+t.Dv, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Empresa:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Mayuscula(t.Juridica), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Primer Nombre:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Titulo(t.Primernombre), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Segundo Nombre:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Titulo(t.Segundonombre), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Primer Apellido:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Titulo(t.Primerapellido), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Segundo Apellido:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Titulo(t.Primerapellido), "", 0,
		"", false, 0, "")

	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Activa:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, t.Activa, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Impuesto de ventas:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, t.Iva, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Retencion de Iva:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, t.ReteIva, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Actividad Ica:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, t.ActividadIca, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Direccion:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, t.Direccion, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Ciudad:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, c.Nombre, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Licencia:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, t.Licencia, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Telefono 1:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, t.Telefono1, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Telefono 2:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, t.Telefono2, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "E-mail 1:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, t.Email1, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "E-mail 2:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, t.Email2, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Tipo:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, TraerTipo(t.Tipo), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Documento:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, TraerDocumentoIdentificacion(t.Documento), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Regimen:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, TraerRegimen(t.Regimen), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Fiscal:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, TraerFiscal(t.Fiscal), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Representante Legal:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, Coma(t.RepresentanteNit)+" - "+t.RepresentanteDv, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Nombre:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, t.RepresentanteNombre, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Revisor Fiscal:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, Coma(t.RevisorNit)+" - "+t.ContadorDv, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Nombre:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, t.RevisorNombre, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Contador:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, Coma(t.ContadorNit)+" - "+t.RevisorDv, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Nombre:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, t.ContadorNombre, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Modulo:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, t.Modulo, "", 0,
		"", false, 0, "")

	pdf.SetFooterFunc(func() {
		pdf.SetTextColor(0, 0, 0)
		pdf.SetY(252)
		pdf.SetFont("Arial", "", 9)
		pdf.SetX(20)

		// LINEA
		pdf.Line(20, 259, 204, 259)
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

// INICIA EMPRESA TODOS PDF
func EmpresaTodosCabecera(pdf *gofpdf.Fpdf) {
	pdf.SetFont("Arial", "", 10)
	// RELLENO TITULO
	pdf.SetY(50)
	pdf.SetFillColor(224, 231, 239)
	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(7)
	pdf.SetX(23)
	pdf.CellFormat(247, 6, "No", "0", 0,
		"L", true, 0, "")
	pdf.SetX(33)
	pdf.CellFormat(190, 6, "Codigo", "0", 0,
		"L", false, 0, "")
	pdf.SetX(63)
	pdf.CellFormat(190, 6, "Nombre", "0", 0,
		"L", false, 0, "")
	pdf.SetX(143)
	pdf.CellFormat(190, 6, "Direccion", "0", 0,
		"L", false, 0, "")
	pdf.SetX(193)
	pdf.CellFormat(190, 6, "Telefono", "0", 0,
		"L", false, 0, "")
	pdf.SetX(218)
	pdf.CellFormat(190, 6, "E-mail", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}
func EmpresaTodosDetalle(pdf *gofpdf.Fpdf, miFila empresa, a int) {
	pdf.SetFont("Arial", "", 9)

	pdf.SetX(23)
	pdf.CellFormat(180, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(33)
	pdf.CellFormat(40, 4, Coma(miFila.Codigo)+" - "+miFila.Dv, "", 0,
		"L", false, 0, "")
	pdf.SetX(63)
	pdf.CellFormat(40, 4, Subcadena(miFila.Nombre, 0, 40), "", 0, "L", false, 0, "")
	pdf.SetX(143)
	pdf.CellFormat(40, 4, miFila.Direccion, "", 0,
		"L", false, 0, "")
	pdf.SetX(193)
	pdf.CellFormat(40, 4, miFila.Telefono1, "", 0,
		"L", false, 0, "")
	pdf.SetX(218)
	pdf.CellFormat(40, 4, miFila.Email1, "", 0,
		"L", false, 0, "")
	pdf.Ln(4)
}

func EmpresaTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	//	Codigo := mux.Vars(r)["codigo"]

	t := []empresa{}
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	err := db.Select(&t, "SELECT * FROM empresa ORDER BY cast(codigo as integer) ")
	if err != nil {
		log.Fatalln(err)
	}
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
		pdf.CellFormat(190, 10, "Nit No. "+Coma(e.Codigo)+" - "+e.Dv, "0", 0, "C",
			false, 0, "")
		pdf.Ln(4)
		pdf.SetX(55)
		pdf.CellFormat(190, 10, e.Iva+" - "+e.ReteIva, "0", 0, "C", false, 0,
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
		pdf.CellFormat(190, 10, ene(c.NombreCiudad+" - "+c.NombreDepartamento), "0", 0, "C", false, 0,
			"")
		pdf.Ln(6)
		pdf.SetX(55)
		pdf.CellFormat(190, 10, "DATOS EMPRESA", "0", 0,
			"C", false, 0, "")
		pdf.Ln(10)
	})

	pdf.SetFooterFunc(func() {
		pdf.SetFont("Arial", "", 9)
		pdf.SetTextColor(0, 0, 0)
		pdf.SetY(192)
		pdf.SetX(23)
		// LINEA
		pdf.Line(23, 199, 269, 199)
		pdf.Ln(6)
		pdf.SetX(23)
		pdf.CellFormat(90, 10, "Sadconf.com", "", 0,
			"L", false, 0, "")
		pdf.CellFormat(161, 10, fmt.Sprintf(" %d de {nb}", pdf.PageNo()), "",
			0, "R", false, 0, "")
	})

	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)
	pdf.SetX(30)

	EmpresaTodosCabecera(pdf)
	// tercera pagina

	for i, miFila := range t {
		var a = i + 1
		if math.Mod(float64(a), 49) == 0 {
			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			EmpresaTodosCabecera(pdf)
		}
		EmpresaTodosDetalle(pdf, miFila, a)
	}

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

// TERMINA EMPRESA TODOS PDF

// EMPRESA EXCEL
func EmpresaExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	t := []empresa{}
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	err := db.Select(&t, "SELECT * FROM empresa ORDER BY cast(codigo as integer) ")
	if err != nil {
		log.Fatalln(err)
	}

	f := excelize.NewFile()

	// FUNCION ANCHO DE LA COLUMNA
	if err = f.SetColWidth("Sheet1", "A", "A", 15); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.SetColWidth("Sheet1", "B", "B", 30); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.SetColWidth("Sheet1", "C", "C", 30); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.SetColWidth("Sheet1", "D", "D", 20); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "E", "E", 30); err != nil {
		fmt.Println(err)
		return
	}

	// FUNCION PARA UNIR DOS CELDAS TITULO
	if err = f.MergeCell("Sheet1", "A1", "E1"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A2", "E2"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A3", "E3"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A4", "E4"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A5", "E5"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A6", "E6"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A7", "E7"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A8", "E8"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A9", "E9"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A10", "E10"); err != nil {
		fmt.Println(err)
		return
	}
	estiloTitulo, err := f.NewStyle(`{  "alignment":{"horizontal": "center"},"font":{"bold":false,"italic":false,"family":"Arial","size":10,"color":"##000000"}}`)

	// titulo
	f.SetCellValue("Sheet1", "A1", e.Nombre)
	f.SetCellValue("Sheet1", "A2", "Nit No. "+Coma(e.Codigo)+" - "+e.Dv)
	f.SetCellValue("Sheet1", "A3", e.Iva+" - "+e.ReteIva)
	f.SetCellValue("Sheet1", "A4", "Actividad Ica - "+e.ActividadIca)
	f.SetCellValue("Sheet1", "A5", e.Direccion)
	f.SetCellValue("Sheet1", "A6", (e.Telefono1 + " - " + e.Telefono2))
	f.SetCellValue("Sheet1", "A7", (c.NombreCiudad + " - " + c.NombreDepartamento))
	f.SetCellValue("Sheet1", "A8", "")
	f.SetCellValue("Sheet1", "A9", "LISTADO DE EMPRESAS")
	f.SetCellValue("Sheet1", "A10", "")

	f.SetCellStyle("Sheet1", "A1", "A1", estiloTitulo)
	f.SetCellStyle("Sheet1", "A2", "A2", estiloTitulo)
	f.SetCellStyle("Sheet1", "A3", "A3", estiloTitulo)
	f.SetCellStyle("Sheet1", "A4", "A4", estiloTitulo)
	f.SetCellStyle("Sheet1", "A5", "A5", estiloTitulo)
	f.SetCellStyle("Sheet1", "A6", "A6", estiloTitulo)
	f.SetCellStyle("Sheet1", "A7", "A7", estiloTitulo)
	f.SetCellStyle("Sheet1", "A8", "A8", estiloTitulo)
	f.SetCellStyle("Sheet1", "A9", "A9", estiloTitulo)
	f.SetCellStyle("Sheet1", "A10", "A10", estiloTitulo)

	var filaExcel = 11

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
	f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel), "Codigo")
	f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel), "Nombre")
	f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel), "Direccion")
	f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel), "Telefono")
	f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel), "E-mail")

	f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel), "A"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel), "B"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel), "C"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel), "D"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "E"+strconv.Itoa(filaExcel), "E"+strconv.Itoa(filaExcel), estiloCabecera)
	filaExcel++

	for i, miFila := range t {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel+i), Entero(miFila.Codigo))
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel+i), miFila.Nombre)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel+i), miFila.Direccion)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel+i), miFila.Telefono1)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel+i), miFila.Email1)

		f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel+i), "A"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel+i), "B"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel+i), "C"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel+i), "D"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "E"+strconv.Itoa(filaExcel+i), "E"+strconv.Itoa(filaExcel+i), estiloTexto)
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

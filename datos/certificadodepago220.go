package main

//
//// INICIA TERCERO IMPORTAR PAQUETES
//import (
//	"bytes"
//	"database/sql"
//	"encoding/json"
//	_ "encoding/json"
//	"fmt"
//	"github.com/360EntSecGroup-Skylar/excelize"
//	_ "github.com/bitly/go-simplejson"
//	"github.com/gorilla/mux"
//	_ "github.com/gorilla/mux"
//	"github.com/jung-kurt/gofpdf"
//	_ "github.com/lib/pq"
//	"html/template"
//	"log"
//	"math"
//	"net/http"
//	"strconv"
//)
//
//// TERMINA  a IMPORTAR PAQUETES
//
//// INICIA TERCERO ESTRUCTURA JSON
//type terceroJson struct {
//	Id     string `json:"id"`
//	Label  string `json:"label"`
//	Value  string `json:"value"`
//	Nombre string `json:"nombre"`
//}
//
//// TERMINA TERCERO ESTRUCTURA JSON
//type tercerolista struct {
//	Codigo string
//	Dv     string
//	Nombre string
//}
//
//// INICIA TERCERO ESTRUCTURA
//type tercero struct {
//	Codigo          string
//	Dv              string
//	Nombre          string
//	Juridica        string
//	PrimerNombre    string
//	SegundoNombre   string
//	PrimerApellido  string
//	SegundoApellido string
//	Direccion       string
//	Barrio          string
//	Telefono1       string
//	Telefono2       string
//	Email1          string
//	Email2          string
//	Contacto        string
//	Rut             string
//	Descuento1      string
//	Descuento2      string
//	Cuotap          string
//	Cuota1          string
//	Cuota2          string
//	Cuota3          string
//	Area            string
//	Factor          string
//	Matricula       string
//	Catastral       string
//	Banco           string
//	PhCodigo        string
//	PhDv            string
//	PhNombre        string
//	Ciudad          string
//	Documento       string
//	Fiscal          string
//	Regimen         string
//	Tipo            string
//	Ica             string
//	Ph              string
//}
//
//// TERMINA TERCERO ESTRUCTURA
//
//// INICIA TERCERO LISTA
//func TerceroLista(w http.ResponseWriter, r *http.Request) {
//	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
//		"vista/tercero/terceroLista.html")
//	log.Println("Error tercero 0")
//	db := dbConn()
//	res := []tercerolista{}
//	db.Select(&res, "SELECT codigo,dv,nombre FROM tercero ORDER BY cast(codigo as float) ASC")
//	varmap := map[string]interface{}{
//		"res":     res,
//		"hosting": ruta,
//	}
//	log.Println("Error tercero888")
//	tmp.Execute(w, varmap)
//}
//
//// TERMINA TERCERO LISTA
//
////INICIA TERCERO NUEVO
//func TerceroNuevo(w http.ResponseWriter, r *http.Request) {
//	Codigo := mux.Vars(r)["codigo"]
//	Panel := mux.Vars(r)["panel"]
//	Elemento := mux.Vars(r)["elemento"]
//	t := tercero{}
//
//	log.Println("Error tercero nuevo 2")
//	parametros := map[string]interface{}{
//		// INICIA TERCERO NUEVO AUTOCOMPLETADO
//		"Codigo":                  Codigo,
//		"Panel":                   Panel,
//		"Elemento":                Elemento,
//		"hosting":                 ruta,
//		"ciudad":                  ListaCiudad(),
//		"tipoorganizacion":        ListaTipoOrganizacion(),
//		"regimenfiscal":           ListaRegimenFiscal(),
//		"responsabilidadfiscal":   ListaResponsabilidadFiscal(),
//		"documentoidentificacion": ListaDocumentoIdentificacion(),
//		"copiar":                  "False",
//		"emp":                     t,
//		// TERMINA TERCERO NUEVO AUTOCOMPLETADO
//	}
//	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
//		"vista/tercero/terceroNuevo.html",
//		"vista/autocompleta/autocompletaTercero.html",
//		"vista/tercero/autocompletaTercerocrear.html")
//	log.Println("Error tercero nuevo 3")
//	tmp.Execute(w, parametros)
//}
//
//// INICIA TERCERO DUPLICAR
//func TerceroNuevoCopia(w http.ResponseWriter, r *http.Request) {
//	log.Println("Error tercero nuevo 1")
//	Codigo := "False"
//	Panel := "False"
//	Elemento := "False"
//
//	copiarCodigo := Quitacoma(mux.Vars(r)["copiacodigo"])
//	log.Println("inicio tercero editar" + Codigo)
//
//	db := dbConn()
//	t := tercero{}
//
//	if copiarCodigo == "False" {
//
//	} else {
//		// traer comprobante
//
//		err := db.Get(&t, "SELECT * FROM tercero WHERE codigo=$1", copiarCodigo)
//		if err != nil {
//			log.Fatalln(err)
//		}
//	}
//
//	log.Println("Error tercero nuevo 2")
//	parametros := map[string]interface{}{
//		// INICIA TERCERO NUEVO AUTOCOMPLETADO
//		"Codigo":                  Codigo,
//		"Panel":                   Panel,
//		"Elemento":                Elemento,
//		"hosting":                 ruta,
//		"ciudad":                  ListaCiudad(),
//		"tipoorganizacion":        ListaTipoOrganizacion(),
//		"regimenfiscal":           ListaRegimenFiscal(),
//		"responsabilidadfiscal":   ListaResponsabilidadFiscal(),
//		"documentoidentificacion": ListaDocumentoIdentificacion(),
//		"emp":                     t,
//		"copiar":                  "True",
//
//		// TERMINA TERCERO NUEVO AUTOCOMPLETADO
//	}
//	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
//		"vista/tercero/terceroNuevo.html",
//		"vista/autocompleta/autocompletaTercero.html",
//		"vista/tercero/autocompletaTercerocrear.html")
//	log.Println("Error tercero nuevo 3")
//	tmp.Execute(w, parametros)
//}
//
//// INICIA TERCERO INSERTAR
//func TerceroInsertar(w http.ResponseWriter, r *http.Request) {
//	db := dbConn()
//	err := r.ParseForm()
//	if err != nil {
//		panic(err.Error())
//	}
//	var t tercero
//	err = decoder.Decode(&t, r.PostForm)
//	if err != nil {
//		panic(err.Error())
//	}
//	var q string
//	q = "insert into tercero ("
//	q += "Codigo,"
//	q += "Dv,"
//	q += "Nombre,"
//	q += "Juridica,"
//	q += "PrimerNombre,"
//	q += "SegundoNombre,"
//	q += "PrimerApellido,"
//	q += "SegundoApellido,"
//	q += "Direccion,"
//	q += "Barrio,"
//	q += "Telefono1,"
//	q += "Telefono2,"
//	q += "Email1,"
//	q += "Email2,"
//	q += "Contacto,"
//	q += "Rut,"
//	q += "Ciudad,"
//	q += "Documento,"
//	q += "Fiscal,"
//	q += "Regimen,"
//	q += "Tipo,"
//	q += "Ica,"
//	q += "Descuento1,"
//	q += "Descuento2,"
//	q += "Cuotap,"
//	q += "Cuota1,"
//	q += "Cuota2,"
//	q += "Cuota3,"
//	q += "Area,"
//	q += "Factor,"
//	q += "Matricula,"
//	q += "Catastral,"
//	q += "Banco,"
//	q += "PhCodigo,"
//	q += "PhDv,"
//	q += "PhNombre,"
//	q += "Ph"
//	q += " ) values("
//	q += parametros(37)
//	q += " ) "
//
//	log.Println("Cadena SQL " + q)
//	insForm, err := db.Prepare(q)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	// INICIA GRABAR TERCERO INSERTAR
//	t.Codigo = Quitacoma(t.Codigo)
//	t.PhCodigo = Quitacoma(t.PhCodigo)
//	t.Nombre = t.Juridica
//	if t.Tipo == "2" {
//		t.PrimerNombre = Titulo(t.PrimerNombre)
//		t.SegundoNombre = Titulo(t.SegundoNombre)
//		t.PrimerApellido = Titulo(t.PrimerApellido)
//		t.SegundoApellido = Titulo(t.SegundoApellido)
//		t.Nombre = t.PrimerNombre + " " + t.SegundoNombre + " " + t.PrimerApellido + " " + t.SegundoApellido
//	} else {
//		t.Juridica = Titulo(t.Juridica)
//		t.Nombre = t.Juridica
//	}
//	t.Nombre = Titulo(t.Nombre)
//	t.Direccion = Titulo(t.Direccion)
//	t.Barrio = Titulo(t.Barrio)
//	t.Contacto = Titulo(t.Contacto)
//	t.Banco = Titulo(t.Banco)
//	t.PhNombre = Titulo(t.PhNombre)
//	t.Email1 = Minuscula(t.Email1)
//	t.Email2 = Minuscula(t.Email2)
//	// TERMINA TERCERO GRABAR INSERTAR
//	_, err = insForm.Exec(
//		t.Codigo,
//		t.Dv,
//		t.Nombre,
//		t.Juridica,
//		t.PrimerNombre,
//		t.SegundoNombre,
//		t.PrimerApellido,
//		t.SegundoApellido,
//		t.Direccion,
//		t.Barrio,
//		t.Telefono1,
//		t.Telefono2,
//		t.Email1,
//		t.Email2,
//		t.Contacto,
//		t.Rut,
//		t.Ciudad,
//		t.Documento,
//		t.Fiscal,
//		t.Regimen,
//		t.Tipo,
//		t.Ica,
//		"",
//		"",
//		"",
//		"",
//		"",
//		"",
//		"",
//		"",
//		"",
//		"",
//		"",
//		t.PhCodigo,
//		t.PhDv,
//		t.PhNombre,
//		t.Ph)
//
//	if err != nil {
//		panic(err)
//	}
//	http.Redirect(w, r, "/TerceroLista", 301)
//}
//
//// TERMINA TERCERO INSERTAR
//
//// INICIA TERCERO BUSCAR
//func TerceroBuscar(w http.ResponseWriter, r *http.Request) {
//	db := dbConn()
//	Codigo := mux.Vars(r)["codigo"]
//	Codigo = Quitacoma(Codigo)
//	selDB, err := db.Query("SELECT codigo,"+
//		"nombre FROM tercero where (codigo LIKE '%' || $1 || '%')  or  (upper(nombre) LIKE '%' || $1 || '%') ORDER BY"+
//		" codigo DESC", Mayuscula(Codigo))
//	if err != nil {
//		panic(err.Error())
//	}
//	var resJson []terceroJson
//	var contar int
//	contar = 0
//	for selDB.Next() {
//		contar++
//		var id string
//		var label string
//		var value string
//		var nombre string
//		err = selDB.Scan(&id, &nombre)
//		if err != nil {
//			panic(err.Error())
//		}
//		value = id
//		label = id + " - " + nombre
//		resJson = append(resJson, terceroJson{id, label, value, nombre})
//	}
//	if err := selDB.Err(); err != nil { // make sure that there was no issue during the process
//		log.Println(err)
//		return
//	}
//	if contar == 0 {
//		var slice []string
//		slice = make([]string, 0)
//		data, _ := json.Marshal(slice)
//		w.WriteHeader(200)
//		w.Header().Set("Content-Type", "application/json")
//		w.Write(data)
//	} else {
//		data, _ := json.Marshal(resJson)
//		w.WriteHeader(200)
//		w.Header().Set("Content-Type", "application/json")
//		w.Write(data)
//	}
//}
//
//// TERMINA TERCERO BUSCAR
////type ValorBoleano struct {
////	Result bool `json:"result,omitempty"`
////}
//// INICIA TERCERO EXISTE
//func TerceroExiste(w http.ResponseWriter, r *http.Request) {
//	db := dbConn()
//	Codigo := mux.Vars(r)["codigo"]
//	log.Println("codigo nombre 1" + Codigo)
//
//	var total int
//	row := db.QueryRow("SELECT COUNT(*) FROM tercero  WHERE codigo=$1", Codigo)
//	err := row.Scan(&total)
//	if err != nil {
//		log.Fatal(err)
//	}
//	var resultado bool
//	if total > 0 {
//		log.Println("si encontro")
//		resultado = true
//	} else {
//		resultado = false
//		log.Println("no encontro")
//	}
//	log.Println("codigo nombre 2")
//	js, err := json.Marshal(SomeStruct{resultado})
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	//log.Println(js )
//	w.Header().Set("Content-Type", "application/json")
//	//w.Write(js)
//	w.Write(js)
//
//}
//
//// TERMINA TERCERO EXISTE
//
//// INICIA TERCERO ACTUAL
//func TerceroActual(w http.ResponseWriter, r *http.Request) {
//	db := dbConn()
//	Codigo := mux.Vars(r)["codigo"]
//	Codigo = Quitacoma(Codigo)
//
//	t := tercero{}
//	var res []tercero
//	err := db.Get(&t, "SELECT * FROM tercero where codigo=$1", Codigo)
//
//	switch err {
//	case nil:
//		log.Printf("tercero found: %+v\n", t)
//		res = append(res, t)
//		data, err := json.Marshal(res)
//		if err != nil {
//			log.Fatalln(err)
//		}
//		w.WriteHeader(200)
//		w.Header().Set("Content-Type", "application/json")
//		w.Write(data)
//
//	case sql.ErrNoRows:
//		log.Println("tercero NOT found, no error")
//
//		var slice []string
//		slice = make([]string, 0)
//		data, _ := json.Marshal(slice)
//		w.WriteHeader(200)
//		w.Header().Set("Content-Type", "application/json")
//		w.Write(data)
//
//	default:
//		log.Printf("tercero error: %s\n", err)
//	}
//
//	log.Println("codigo nombre99" + t.Codigo + t.Nombre)
//
//}
//
//// TERMINA TERCERO ACTUAL
//
//// INICIA TERCERO EDITAR
//func TerceroEditar(w http.ResponseWriter, r *http.Request) {
//	Codigo := mux.Vars(r)["codigo"]
//	log.Println("inicio tercero editar" + Codigo)
//	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
//		"vista/tercero/terceroEditar.html",
//		"vista/autocompleta/autocompletaTercero.html")
//	db := dbConn()
//	t := tercero{}
//	err := db.Get(&t, "SELECT * FROM tercero where codigo=$1", Codigo)
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	log.Println("codigo nombre99" + t.Codigo + t.Nombre)
//	varmap := map[string]interface{}{
//		// INICIA TERCERO EDITAR AUTOCOMPLETADO
//		"emp":                     t,
//		"hosting":                 ruta,
//		"ciudad":                  ListaCiudad(),
//		"tipoorganizacion":        ListaTipoOrganizacion(),
//		"regimenfiscal":           ListaRegimenFiscal(),
//		"responsabilidadfiscal":   ListaResponsabilidadFiscal(),
//		"documentoidentificacion": ListaDocumentoIdentificacion(),
//		// TERMINA TERCERO EDITAR AUTOCOMPLETADO
//	}
//	tmp.Execute(w, varmap)
//}
//
//// INICIA TERCERO ACTUALIZAR
//func TerceroActualizar(w http.ResponseWriter, r *http.Request) {
//	db := dbConn()
//	err := r.ParseForm()
//	if err != nil {
//		panic(err.Error())
//		// Handle error
//	}
//	var t tercero
//	// r.PostForm is a map of our POST form values
//	err = decoder.Decode(&t, r.PostForm)
//	if err != nil {
//		// Handle error
//		panic(err.Error())
//	}
//	var q string
//	q = "UPDATE tercero set "
//	q += "Dv=$2,"
//	q += "Nombre=$3,"
//	q += "Juridica=$4,"
//	q += "PrimerNombre=$5,"
//	q += "SegundoNombre=$6,"
//	q += "PrimerApellido=$7,"
//	q += "SegundoApellido=$8,"
//	q += "Direccion=$9,"
//	q += "Barrio=$10,"
//	q += "Telefono1=$11,"
//	q += "Telefono2=$12,"
//	q += "Email1=$13,"
//	q += "Email2=$14,"
//	q += "Contacto=$15,"
//	q += "Rut=$16,"
//	q += "Ciudad=$17,"
//	q += "Documento=$18,"
//	q += "Fiscal=$19,"
//	q += "Regimen=$20,"
//	q += "Ica=$21,"
//	q += "Tipo=$22,"
//	q += "PhCodigo=$23,"
//	q += "PhDv=$24,"
//	q += "PhNombre=$25,"
//	q += "Ph=$26"
//	q += " where "
//	q += "Codigo=$1"
//
//	log.Println("cadena" + q)
//
//	insForm, err := db.Prepare(q)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	// INICIA GRABAR TERCERO ACTUALIZAR
//	t.Codigo = Quitacoma(t.Codigo)
//	t.PhCodigo = Quitacoma(t.PhCodigo)
//	t.Nombre = t.Juridica
//	if t.Tipo == "2" {
//		t.PrimerNombre = Titulo(t.PrimerNombre)
//		t.SegundoNombre = Titulo(t.SegundoNombre)
//		t.PrimerApellido = Titulo(t.PrimerApellido)
//		t.SegundoApellido = Titulo(t.SegundoApellido)
//		t.Nombre = t.PrimerNombre + " " + t.SegundoNombre + " " + t.PrimerApellido + " " + t.SegundoApellido
//	} else {
//		t.Juridica = Titulo(t.Juridica)
//		t.Nombre = t.Juridica
//	}
//	t.Nombre = Titulo(t.Nombre)
//	t.Direccion = Titulo(t.Direccion)
//	t.Barrio = Titulo(t.Barrio)
//	t.Contacto = Titulo(t.Contacto)
//	t.PhNombre = Titulo(t.PhNombre)
//	t.Email1 = Minuscula(t.Email1)
//	t.Email2 = Minuscula(t.Email2)
//	// TERMINA GRABAR TERCERO ACTUALIZAR
//
//	_, err = insForm.Exec(
//		t.Codigo,
//		t.Dv,
//		t.Nombre,
//		t.Juridica,
//		t.PrimerNombre,
//		t.SegundoNombre,
//		t.PrimerApellido,
//		t.SegundoApellido,
//		t.Direccion,
//		t.Barrio,
//		t.Telefono1,
//		t.Telefono2,
//		t.Email1,
//		t.Email2,
//		t.Contacto,
//		t.Rut,
//		t.Ciudad,
//		t.Documento,
//		t.Fiscal,
//		t.Regimen,
//		t.Ica,
//		t.Tipo,
//		t.PhCodigo,
//		t.PhDv,
//		t.PhNombre,
//		t.Ph)
//
//	if err != nil {
//		panic(err)
//	}
//	http.Redirect(w, r, "/TerceroLista", 301)
//
//}
//
//// INICIA TERCERO BORRAR
//func TerceroBorrar(w http.ResponseWriter, r *http.Request) {
//	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
//		"vista/tercero/terceroBorrar.html")
//	Codigo := mux.Vars(r)["codigo"]
//	log.Println("inicio tercero borrar" + Codigo)
//	db := dbConn()
//	t := tercero{}
//	err := db.Get(&t, "SELECT * FROM tercero where codigo=$1", Codigo)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	log.Println("codigo nombre99 borrar" + t.Codigo + t.Nombre)
//	varmap := map[string]interface{}{
//		// INICIA TERCERO BORRAR AUTOCOMPLETADO
//		"emp":                     t,
//		"hosting":                 ruta,
//		"ciudad":                  ListaCiudad(),
//		"tipoorganizacion":        ListaTipoOrganizacion(),
//		"regimenfiscal":           ListaRegimenFiscal(),
//		"responsabilidadfiscal":   ListaResponsabilidadFiscal(),
//		"documentoidentificacion": ListaDocumentoIdentificacion(),
//		// TERMINA TERCERO BORRAR AUTOCOMPLETADO
//	}
//	tmp.Execute(w, varmap)
//}
//
//// INICIA TERCERO ELIMINAR
//func TerceroEliminar(w http.ResponseWriter, r *http.Request) {
//	log.Println("Inicio Eliminar")
//	db := dbConn()
//	emp := mux.Vars(r)["codigo"]
//	delForm, err := db.Prepare("DELETE from tercero WHERE codigo=$1")
//	if err != nil {
//		panic(err.Error())
//	}
//	delForm.Exec(emp)
//	log.Println("Registro Eliminado" + emp)
//	http.Redirect(w, r, "/TerceroLista", 301)
//}
//
//// INICIA CERTIFICADO DE PAGO 220 PDF
//func Empleado220Pdf(w http.ResponseWriter, r *http.Request) {
//	db := dbConn()
//	Codigo := mux.Vars(r)["codigo"]
//	t := tercero{}
//	var e empresa = ListaEmpresa()
//	//var c  ciudad=TraerCiudad(e.Ciudad)
//	err := db.Get(&t, "SELECT * FROM tercero where codigo=$1", Codigo)
//	if err != nil {
//		log.Fatalln(err)
//	}
//	var buf bytes.Buffer
//	var err1 error
//	pdf := gofpdf.New("P", "mm", "Letter", cnFontDir)
//	ene := pdf.UnicodeTranslatorFromDescriptor("")
//
//	pdf.SetX(21)
//	pdf.AliasNbPages("")
//	pdf.AddPage()
//
//	pdf.Ln(2)
//	pdf.SetFont("Arial", "", 20)
//	//pdf.SetDrawColor(84,153,199)
//	pdf.SetDrawColor(95, 119, 146)
//	pdf.SetTextColor(0, 0, 0)
//
//	// LINEA HORIZONTAL
//	pdf.Line(21, 14, 201, 14)
//	pdf.Line(21, 26, 201, 26)
//	pdf.Line(21, 33, 201, 33)
//	pdf.Line(21, 55, 201, 55)
//	pdf.Line(21, 69, 201, 69)
//	pdf.Line(30, 45, 201, 45)
//	// LINEA VERTICAL
//	pdf.Line(47, 14, 47, 26)
//	pdf.Line(175, 14, 175, 26)
//	pdf.Line(21, 14, 21, 79)
//	pdf.Line(201, 14, 201, 79)
//	pdf.Line(111, 26, 111, 33)
//	pdf.Line(30, 33, 30, 55)
//	pdf.Line(30, 55, 30, 69)
//	pdf.Line(74, 69, 74, 79)
//	pdf.Line(105, 69, 105, 79)
//	pdf.Line(163, 69, 163, 79)
//	pdf.Line(178, 69, 178, 79)
//
//	// RETENEDOR
//	pdf.Line(85, 40, 85, 45)
//	pdf.Line(100, 40, 100, 45)
//	pdf.Line(125, 40, 125, 45)
//	pdf.Line(150, 40, 150, 45)
//	pdf.Line(175, 40, 175, 45)
//
//	// TRABAJADOR
//	pdf.Line(43, 64, 43, 69)
//	//pdf.Line(85,64,85,69)
//	pdf.Line(100, 64, 100, 69)
//	pdf.Line(125, 64, 125, 69)
//	pdf.Line(150, 64, 150, 69)
//	pdf.Line(175, 64, 175, 69)
//
//	pdf.SetY(15)
//	pdf.SetX(185)
//	pdf.Image(imageFile("Dian.png"), 25, 15, 20, 10, false,
//		"", 0, "")
//	pdf.SetFont("Arial", "B", 10)
//	pdf.SetY(15)
//	pdf.SetX(40)
//	pdf.CellFormat(140, 4, ene("Certificado de Ingresos y Retenciones Por Rentas de Trabajo y Pensiones"), "0", 0, "C", false, 0, "")
//	pdf.SetY(20)
//	pdf.SetX(40)
//	pdf.CellFormat(140, 4, ene("Año Gravable 2021"), "0", 0, "C", false, 0, "")
//	pdf.SetY(15)
//	pdf.SetX(177)
//	pdf.SetFont("Arial", "", 20)
//
//	pdf.SetTextColor(253, 254, 254)
//	pdf.SetFillColor(56, 100, 146)
//	pdf.CellFormat(22, 10, "220", "0", 0,
//		"C", true, 0, "")
//	pdf.SetTextColor(0, 0, 0)
//	pdf.Ln(-1)
//	pdf.SetX(23)
//	pdf.SetFont("Arial", "", 8)
//	pdf.CellFormat(95, 8, ene("Antes de diligenciar este formulario lea las instrucciones"), "0", 0,
//		"", false, 0, "")
//	pdf.CellFormat(89, 8, " 4. Formulario No."+" "+t.Codigo, "0", 0,
//		"", false, 0, "")
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.SetFont("Arial", "", 7)
//	pdf.TransformBegin()
//	pdf.TransformRotate(90, 28.5, 42.5)
//	pdf.CellFormat(15, 15, "Retenedor", "0", 0,
//		"", false, 0, "")
//	pdf.TransformEnd()
//	pdf.SetX(50)
//	pdf.SetY(35)
//	pdf.SetFont("Arial", "", 8)
//	pdf.CellFormat(100, 4, " 5. Numero de Identificacion Tributario", "0", 0, "C",
//		false, 0, "")
//	pdf.SetX(80)
//	pdf.CellFormat(20, 4, "6. Dv", "0", 0, "C",
//		false, 0, "")
//	pdf.SetY(40)
//	pdf.SetFont("Arial", "", 10)
//	pdf.CellFormat(70, 4, Coma(t.Codigo), "0", 0, "C",
//		false, 0, "")
//	pdf.SetX(35)
//	pdf.CellFormat(110, 4, t.Dv, "0", 0, "C",
//		false, 0, "")
//	pdf.SetFont("Arial", "", 8)
//	pdf.SetY(35)
//	pdf.CellFormat(280, 4, " 7. Primer Apellido  8. Segundo Apellido  9. Primer Nombre  10. Otros Nombres", "0", 0, "C",
//		false, 0, "")
//	pdf.SetY(40)
//	pdf.SetX(101)
//	pdf.SetFont("Arial", "", 10)
//	pdf.CellFormat(40, 4, t.PrimerNombre, "0", 0,
//		"", false, 0, "")
//	pdf.SetX(126)
//	pdf.CellFormat(40, 4, t.SegundoNombre, "0", 0,
//		"", false, 0, "")
//	pdf.SetX(151)
//	pdf.SetFont("Arial", "", 10)
//	pdf.CellFormat(40, 4, t.PrimerApellido, "0", 0,
//		"", false, 0, "")
//	pdf.SetX(176)
//	pdf.CellFormat(40, 4, t.SegundoApellido, "0", 0,
//		"", false, 0, "")
//	pdf.SetFont("Arial", "", 8)
//	pdf.SetY(45)
//	pdf.SetX(37)
//	pdf.CellFormat(80, 4, " 11. Razon Social", "0", 0, "",
//		false, 0, "")
//	pdf.SetFont("Arial", "", 10)
//	pdf.SetY(50)
//	pdf.SetX(40)
//	pdf.CellFormat(120, 4, e.Nombre, "0", 0,
//		"", false, 0, "")
//
//	pdf.SetX(21)
//	pdf.SetFont("Arial", "", 7)
//	pdf.TransformBegin()
//	pdf.TransformRotate(90, 29.5, 60.5)
//	pdf.CellFormat(15, 15, "Trabajador", "0", 0,
//		"", false, 0, "")
//	pdf.TransformEnd()
//	pdf.SetY(57)
//	pdf.SetX(32)
//	pdf.CellFormat(10, 4, " 24. Tipo", "0", 0, "R",
//		false, 0, "")
//	pdf.SetFont("Arial", "", 10)
//	pdf.SetY(62)
//	pdf.SetX(30)
//	pdf.CellFormat(10, 4, "13", "0", 0, "C",
//		false, 0, "")
//	pdf.SetX(150)
//	pdf.SetY(57)
//	pdf.SetFont("Arial", "", 8)
//	pdf.CellFormat(120, 4, " 25. Numero de Identificacion Tributario", "0", 0, "C",
//		false, 0, "")
//	pdf.SetY(62)
//	pdf.SetX(21)
//	pdf.SetFont("Arial", "", 10)
//	pdf.CellFormat(70, 4, Coma(t.Codigo), "0", 0, "C",
//		false, 0, "")
//	pdf.SetFont("Arial", "", 8)
//	pdf.SetY(57)
//	pdf.CellFormat(280, 4, "26. Primer Apellido  27. Segundo Apellido  28. Primer Nombre  29. Otros Nombres", "0", 0, "C",
//		false, 0, "")
//	pdf.SetY(62)
//	pdf.SetX(101)
//	pdf.SetFont("Arial", "", 10)
//	pdf.CellFormat(40, 4, t.PrimerNombre, "0", 0,
//		"", false, 0, "")
//	pdf.SetX(126)
//	pdf.CellFormat(40, 4, t.SegundoNombre, "0", 0,
//		"", false, 0, "")
//	pdf.SetX(151)
//	pdf.SetFont("Arial", "", 10)
//	pdf.CellFormat(40, 4, t.PrimerApellido, "0", 0,
//		"", false, 0, "")
//	pdf.SetX(176)
//	pdf.CellFormat(40, 4, t.SegundoApellido, "0", 0,
//		"", false, 0, "")
//
//	pdf.SetY(70)
//	pdf.SetX(23)
//	pdf.SetFont("Arial", "", 8)
//	pdf.CellFormat(40, 4, " Periodo de Certificacion", "0", 0,
//		"", false, 0, "")
//	pdf.SetX(75)
//	pdf.CellFormat(40, 4, " Fecha Expedicion", "0", 0,
//		"", false, 0, "")
//	pdf.SetX(107)
//	pdf.CellFormat(40, 4, "33. Lugar Donde Se Practico la Retencion", "0", 0,
//		"", false, 0, "")
//	pdf.SetX(165)
//	pdf.CellFormat(40, 4, "34. Cod.", "0", 0,
//		"", false, 0, "")
//	pdf.SetX(178)
//	pdf.CellFormat(40, 4, "35. Cod. Ciudad", "0", 0,
//		"", false, 0, "")
//
//	pdf.SetY(75)
//	pdf.SetX(21)
//	pdf.SetFont("Arial", "", 9)
//	pdf.CellFormat(40, 4, "30. de 01/01/2022 A 31. 31/12/2022", "0", 0,
//		"", false, 0, "")
//	pdf.SetX(70)
//	pdf.CellFormat(40, 4, "32. 30/03/2022", "0", 0,
//		"C", false, 0, "")
//	pdf.SetX(110)
//	pdf.CellFormat(40, 4, TraerCiudad(t.Ciudad).NombreCiudad+"-"+TraerCiudad(t.Ciudad).NombreDepartamento, "0", 0,
//		"", false, 0, "")
//	pdf.SetX(167)
//	pdf.CellFormat(40, 4, TraerCiudad(t.Ciudad).CodigoDepartamento, "0", 0,
//		"", false, 0, "")
//	pdf.SetX(185)
//	pdf.CellFormat(40, 4, TraerCiudad(t.Ciudad).CodigoCiudad, "0", 0,
//		"", false, 0, "")
//
//	pdf.SetFont("Arial", "B", 8)
//	pdf.SetFillColor(225, 232, 239)
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(142, 4, "Concepto de los Ingresos", "1", 0,
//		"C", true, 0, "")
//	pdf.CellFormat(38, 4, "Valor", "1", 0,
//		"C", true, 0, "")
//	pdf.SetFont("Arial", "", 8)
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(142, 4, "Pagos por salarios o emolumentos eclesiasticos", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(5, 4, "36", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "", 0,
//		"R", false, 0, "")
//	pdf.SetFillColor(225, 232, 239)
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(142, 4, "Pagos realizados con bonos electronicos o de papel de servicio, cheques, tarjetas, vales, etc", "", 0,
//		"", true, 0, "")
//	pdf.CellFormat(5, 4, "37", "", 0,
//		"", true, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "", 0,
//		"R", true, 0, "")
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(142, 4, "Pagos por Honorarios", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(5, 4, "38", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "", 0,
//		"R", false, 0, "")
//	pdf.SetFillColor(225, 232, 239)
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(142, 4, "Pagos por Servicios", "", 0,
//		"", true, 0, "")
//	pdf.CellFormat(5, 4, "39", "", 0,
//		"", true, 0, "")
//	pdf.CellFormat(33, 4, "333.000.000", "", 0,
//		"R", true, 0, "")
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(142, 4, "Pagos por Comisiones", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(5, 4, "40", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "", 0,
//		"R", false, 0, "")
//	pdf.SetFillColor(225, 232, 239)
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(142, 4, "Pagos por Prestaciones Sociales", "", 0,
//		"", true, 0, "")
//	pdf.CellFormat(5, 4, "41", "", 0,
//		"", true, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "", 0,
//		"R", true, 0, "")
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(142, 4, "Pagos por Viaticos", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(5, 4, "42", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "", 0,
//		"R", false, 0, "")
//	pdf.SetFillColor(225, 232, 239)
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(142, 4, "Pagos por Gastos de Representacion", "", 0,
//		"", true, 0, "")
//	pdf.CellFormat(5, 4, "43", "", 0,
//		"", true, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "", 0,
//		"R", true, 0, "")
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(142, 4, "Pagos por Compensaciones por el Trabajo Asociado Cooperativo", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(5, 4, "44", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "", 0,
//		"R", false, 0, "")
//	pdf.SetFillColor(225, 232, 239)
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(142, 4, "Otros Pagos", "", 0,
//		"", true, 0, "")
//	pdf.CellFormat(5, 4, "45", "", 0,
//		"", true, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "", 0,
//		"R", true, 0, "")
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(142, 4, "Cesantias e Intereses de Cesantias Efectivamente Pagados al Empleado", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(5, 4, "46", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "", 0,
//		"R", false, 0, "")
//	pdf.SetFillColor(225, 232, 239)
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(142, 4, "Cesantias Consignadas al Fondo de Cesantias", "", 0,
//		"", true, 0, "")
//	pdf.CellFormat(5, 4, "47", "", 0,
//		"", true, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "", 0,
//		"R", true, 0, "")
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(142, 4, "Pension de Jubilacion, Vejes o Invalidez", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(5, 4, "48", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "", 0,
//		"R", false, 0, "")
//	pdf.SetFillColor(225, 232, 239)
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(142, 4, "Total de Ingresos Brutos (Suma Casillas 36 a 48)", "", 0,
//		"", true, 0, "")
//	pdf.CellFormat(5, 4, "49", "", 0,
//		"", true, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "", 0,
//		"R", true, 0, "")
//	pdf.SetFillColor(225, 232, 239)
//	pdf.SetFont("Arial", "B", 8)
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(142, 4, "Concepto de los Aportes", "1", 0,
//		"C", true, 0, "")
//	pdf.CellFormat(38, 4, "Valor", "1", 0,
//		"C", true, 0, "")
//	pdf.SetFont("Arial", "", 8)
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(142, 4, "Aportes Obligatorios por Salud a Cargo del Trabajador", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(5, 4, "50", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "", 0,
//		"R", false, 0, "")
//	pdf.SetFillColor(225, 232, 239)
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(142, 4, "Aportes Obligatorios a Fondo de Pensiones y Solidaridad Pensional a Cargo del Trabajador", "", 0,
//		"", true, 0, "")
//	pdf.CellFormat(5, 4, "51", "", 0,
//		"", true, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "", 0,
//		"R", true, 0, "")
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(142, 4, "Aportes Obligatorios a Fondo de Pensiones y Solidaridad Pensional y Apaortes Voluntarios - RAIS", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(5, 4, "52", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "", 0,
//		"R", false, 0, "")
//	pdf.SetFillColor(225, 232, 239)
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(142, 4, "Aportes Voluntarios al Fondo de Pensiones", "", 0,
//		"", true, 0, "")
//	pdf.CellFormat(5, 4, "53", "", 0,
//		"", true, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "", 0,
//		"R", true, 0, "")
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(142, 4, "Aportes a Cuentas AFC", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(5, 4, "54", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "", 0,
//		"R", false, 0, "")
//	pdf.SetTextColor(253, 254, 254)
//	pdf.SetFillColor(56, 100, 146)
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(142, 4, "Valor de la Retencion en la Fuente Por Renta de Trabajo y Pensiones", "1", 0,
//		"", true, 0, "")
//	pdf.SetTextColor(0, 0, 0)
//	pdf.CellFormat(5, 4, "55", "1", 0,
//		"", false, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "1", 0,
//		"R", false, 0, "")
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(142, 4, "Pasivos Laborales Reales Consolidados en Cabeza del Trabajador", "1", 0,
//		"", false, 0, "")
//	pdf.CellFormat(5, 4, "56", "1", 0,
//		"", false, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "1", 0,
//		"R", false, 0, "")
//
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(180, 4, "Nombre del Pagador o Agente Retenedor:", "1", 0,
//		"", false, 0, "")
//	pdf.SetX(91)
//	pdf.CellFormat(180, 4, t.Nombre, "0", 0,
//		"", false, 0, "")
//	pdf.SetTextColor(253, 254, 254)
//	pdf.SetFillColor(56, 100, 146)
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(180, 4, "Datos a Cargo del Trabajador o Pensionado", "", 0,
//		"C", true, 0, "")
//	pdf.SetTextColor(0, 0, 0)
//	pdf.SetFillColor(225, 232, 239)
//	pdf.SetFont("Arial", "B", 8)
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(110, 4, "Concepto de Otros Ingresos", "1", 0,
//		"C", true, 0, "")
//	pdf.CellFormat(32, 4, "Valor Recibido", "1", 0,
//		"C", true, 0, "")
//	pdf.CellFormat(38, 4, "Valor Retenido", "1", 0,
//		"C", true, 0, "")
//	pdf.SetFont("Arial", "", 8)
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(110, 4, "Arrendamientos", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(5, 4, "57", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(27, 4, "333.900.000", "", 0,
//		"R", false, 0, "")
//	pdf.CellFormat(5, 4, "64", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "", 0,
//		"R", false, 0, "")
//	pdf.SetFillColor(225, 232, 239)
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(110, 4, "Honorarios, Comisiones y Servicios.", "", 0,
//		"", true, 0, "")
//	pdf.CellFormat(5, 4, "58", "", 0,
//		"", true, 0, "")
//	pdf.CellFormat(27, 4, "333.900.000", "", 0,
//		"R", true, 0, "")
//	pdf.CellFormat(5, 4, "65", "", 0,
//		"", true, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "", 0,
//		"R", true, 0, "")
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(110, 4, "Intereses y Rendimientos Financieros", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(5, 4, "59", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(27, 4, "333.900.000", "", 0,
//		"R", false, 0, "")
//	pdf.CellFormat(5, 4, "66", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "", 0,
//		"R", false, 0, "")
//	pdf.SetFillColor(225, 232, 239)
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(110, 4, "Enajenacion de Activos Fijos", "", 0,
//		"", true, 0, "")
//	pdf.CellFormat(5, 4, "60", "", 0,
//		"", true, 0, "")
//	pdf.CellFormat(27, 4, "333.900.000", "", 0,
//		"R", true, 0, "")
//	pdf.CellFormat(5, 4, "67", "", 0,
//		"", true, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "", 0,
//		"R", true, 0, "")
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(110, 4, "Loterias, Rifas, Apuestas y Similares", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(5, 4, "61", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(27, 4, "333.900.000", "", 0,
//		"R", false, 0, "")
//	pdf.CellFormat(5, 4, "68", "", 0,
//		"", false, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "", 0,
//		"R", false, 0, "")
//	pdf.SetFillColor(225, 232, 239)
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(110, 4, "Otros", "", 0,
//		"", true, 0, "")
//	pdf.CellFormat(5, 4, "62", "", 0,
//		"", true, 0, "")
//	pdf.CellFormat(27, 4, "333.900.000", "", 0,
//		"R", true, 0, "")
//	pdf.CellFormat(5, 4, "69", "", 0,
//		"", true, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "", 0,
//		"R", true, 0, "")
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(110, 4, "Totales (Valores Recibidos suma 57 a 62) (Valores Retenidos suma 64 a 69", "1", 0,
//		"", false, 0, "")
//	pdf.CellFormat(5, 4, "63", "1", 0,
//		"", false, 0, "")
//	pdf.CellFormat(27, 4, "333.900.000", "1", 0,
//		"R", false, 0, "")
//	pdf.CellFormat(5, 4, "70", "1", 0,
//		"", false, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "1", 0,
//		"R", false, 0, "")
//
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(142, 4, "Total Retenciones Año Gravable 2021 Suma casilla 55 + 70", "1", 0,
//		"", false, 0, "")
//	pdf.CellFormat(5, 4, "71", "1", 0,
//		"", false, 0, "")
//	pdf.CellFormat(33, 4, "333.900.000", "1", 0,
//		"R", false, 0, "")
//	pdf.SetTextColor(0, 0, 0)
//	pdf.SetFillColor(225, 232, 239)
//	pdf.SetFont("Arial", "B", 8)
//
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(7, 4, "Item", "1", 0,
//		"C", true, 0, "")
//	pdf.CellFormat(135, 4, "72. Identificacion de los Bienes Poseidos", "1", 0,
//		"C", true, 0, "")
//	pdf.CellFormat(38, 4, "73. Valor Patrimonial", "1", 0,
//		"C", true, 0, "")
//	pdf.SetFont("Arial", "", 8)
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(7, 4, "1", "1", 0,
//		"C", false, 0, "")
//	pdf.CellFormat(135, 4, "", "1", 0,
//		"C", false, 0, "")
//	pdf.CellFormat(38, 4, "", "1", 0,
//		"C", false, 0, "")
//	pdf.SetFillColor(225, 232, 239)
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(7, 4, "2", "1", 0,
//		"C", true, 0, "")
//	pdf.CellFormat(135, 4, "", "1", 0,
//		"C", true, 0, "")
//	pdf.CellFormat(38, 4, "", "1", 0,
//		"C", true, 0, "")
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(7, 4, "3", "1", 0,
//		"C", false, 0, "")
//	pdf.CellFormat(135, 4, "", "1", 0,
//		"C", false, 0, "")
//	pdf.CellFormat(38, 4, "", "1", 0,
//		"C", false, 0, "")
//	pdf.SetFillColor(225, 232, 239)
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(7, 4, "4", "1", 0,
//		"C", true, 0, "")
//	pdf.CellFormat(135, 4, "", "1", 0,
//		"C", true, 0, "")
//	pdf.CellFormat(38, 4, "", "1", 0,
//		"C", true, 0, "")
//
//	pdf.SetTextColor(253, 254, 254)
//	pdf.SetFillColor(56, 100, 146)
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(142, 4, "Deudas Vigentes a Diciembre 31 de 2021", "1", 0,
//		"L", true, 0, "")
//	pdf.SetTextColor(0, 0, 0)
//	pdf.CellFormat(5, 4, "74", "1", 0,
//		"C", false, 0, "")
//	pdf.CellFormat(33, 4, "", "1", 0,
//		"C", false, 0, "")
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(180, 4, "Identificacion del Dependiente Economico de acuerdo al paragrafo 2 del Articulo 387 del E. T.", "1", 0,
//		"C", false, 0, "")
//	pdf.SetFillColor(225, 232, 239)
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(45, 4, "75. Tipo Doc 76. Numero Doc", "1", 0,
//		"", true, 0, "")
//	pdf.CellFormat(97, 4, " 77. Apellidos y Nombres", "1", 0,
//		"", true, 0, "")
//	pdf.CellFormat(38, 4, " 78. Parentesco", "1", 0,
//		"", true, 0, "")
//	pdf.Ln(-1)
//	pdf.SetX(21)
//	pdf.CellFormat(45, 4, "13"+"- "+t.Codigo, "1", 0,
//		"", false, 0, "")
//	pdf.CellFormat(97, 4, t.Nombre, "1", 0,
//		"", false, 0, "")
//	pdf.CellFormat(38, 4, " Hijo", "1", 0,
//		"", false, 0, "")
//
//	pdf.SetFooterFunc(func() {
//
//		pdf.SetY(252)
//		pdf.SetX(25)
//		pdf.SetFont("Arial", "", 5)
//		pdf.CellFormat(142, 2, ene("Certifico que Durante el Año Gravable 2021:"), "0", 0,
//			"", false, 0, "")
//
//		pdf.SetY(252)
//		pdf.SetX(160)
//		pdf.SetFont("Arial", "", 6)
//		pdf.CellFormat(42, 2, ene("    Firma del Trabajador o Pensionado"), "0", 0,
//			"C", false, 0, "")
//
//		pdf.SetFont("Arial", "", 5)
//		pdf.Ln(-1)
//		pdf.SetX(25)
//		pdf.CellFormat(140, 2, ene("1. Mi Patrimonio no Excedio de 4.500 UVT"), "0", 0,
//			"", false, 0, "")
//
//		pdf.Ln(-1)
//		pdf.SetX(25)
//		pdf.SetFont("Arial", "", 5)
//		pdf.CellFormat(140, 2, ene("2. Mis Ingresos Fueron Inferiores a 1.400 UVT"), "0", 0,
//			"", false, 0, "")
//		pdf.Ln(-1)
//		pdf.SetX(25)
//		pdf.CellFormat(140, 2, ene("3. No Fui Responsable del Impuesto a Las Ventas"), "0", 0,
//			"", false, 0, "")
//
//		pdf.Ln(-1)
//		pdf.SetX(25)
//		pdf.CellFormat(140, 2, ene("4. Mis Consumos Mediante Tarjeta de Credito No Excedieron La Suma de 1.400 UVT"), "0", 0,
//			"", false, 0, "")
//
//		pdf.Ln(-1)
//		pdf.SetX(25)
//		pdf.CellFormat(140, 2, ene("5. Que el Total de Mis Compras y Consumos No Superaron la Suma de 1.400 UVT"), "0", 0,
//			"", false, 0, "")
//		pdf.Ln(-1)
//		pdf.SetX(25)
//		pdf.CellFormat(140, 2, ene("6. Que el Valor Total de Mis Consignaciones Bancarias, Depositos o Inversiones Financieras No Excedieron los 1.400 UVT"), "0", 0,
//			"", false, 0, "")
//		pdf.Ln(-1)
//		pdf.SetX(25)
//		pdf.CellFormat(140, 2, ene("Por lo Tanto, Manifiesto Que No Extoy Obligado a Presentar Declaracion de Renta y Complementarios por El Año Gravable 2021"), "0", 0,
//			"", false, 0, "")
//
//		// LINEA HORIZONTAL
//		pdf.Line(21, 269, 201, 269)
//		// LINEA VERTICAL
//		pdf.Line(21, 50, 21, 269)
//		pdf.Line(201, 50, 201, 269)
//		pdf.Line(163, 250, 163, 269)
//		pdf.Line(163, 70, 163, 170)
//
//		pdf.Line(168, 83, 168, 139)
//		pdf.Line(168, 143, 168, 169)
//
//		pdf.Line(168, 183, 168, 211)
//		pdf.Line(163, 183, 163, 211)
//
//		pdf.Line(131, 183, 131, 211)
//		pdf.Line(136, 183, 136, 211)
//
//		pdf.Ln(3)
//		pdf.SetFont("Arial", "", 8)
//		pdf.SetX(21)
//		pdf.SetY(269)
//		pdf.CellFormat(40, 4, "Sadconf.com", "", 0,
//			"C", false, 0, "")
//		pdf.SetX(165)
//		pdf.CellFormat(40, 4, fmt.Sprintf(" %d de {nb}", pdf.PageNo()), "",
//			0, "R", false, 0, "")
//	})
//
//	err1 = pdf.Output(&buf)
//	if err1 != nil {
//		panic(err1.Error())
//	}
//	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
//	w.Write(buf.Bytes())
//}
//
//// INICIA TERCEROS TODOS PDF
//func TerceroTodosCabecera(pdf *gofpdf.Fpdf) {
//	pdf.SetFont("Arial", "", 10)
//	// RELLENO TITULO
//	pdf.SetY(50)
//	pdf.SetFillColor(224, 231, 239)
//	pdf.SetTextColor(0, 0, 0)
//	pdf.Ln(7)
//	pdf.SetX(20)
//	pdf.CellFormat(181, 6, "No", "0", 0,
//		"L", true, 0, "")
//	pdf.SetX(30)
//	pdf.CellFormat(40, 6, "Codigo", "0", 0,
//		"L", false, 0, "")
//	pdf.SetX(60)
//	pdf.CellFormat(40, 6, "Nombre", "0", 0,
//		"L", false, 0, "")
//	pdf.SetX(120)
//	pdf.CellFormat(40, 6, "Direccion", "0", 0,
//		"L", false, 0, "")
//	pdf.SetX(171)
//	pdf.CellFormat(40, 6, "Telefono", "0", 0,
//		"L", false, 0, "")
//	pdf.Ln(8)
//}
//func TerceroTodosDetalle(pdf *gofpdf.Fpdf, t tercero, a int) {
//	pdf.SetFont("Arial", "", 9)
//
//	pdf.SetX(21)
//	pdf.CellFormat(181, 4, strconv.Itoa(a), "", 0,
//		"L", false, 0, "")
//	pdf.SetX(30)
//	pdf.CellFormat(40, 4, Coma(t.Codigo)+" - "+t.Dv, "", 0,
//		"L", false, 0, "")
//	pdf.SetX(60)
//	pdf.CellFormat(40, 4, t.Nombre, "", 0, "L", false, 0, "")
//	pdf.SetX(120)
//	pdf.CellFormat(40, 4, t.Direccion, "", 0,
//		"L", false, 0, "")
//	pdf.SetX(155)
//	pdf.CellFormat(40, 4, t.Telefono1, "", 0,
//		"R", false, 0, "")
//	pdf.Ln(4)
//}
//
//func TerceroTodosPdf(w http.ResponseWriter, r *http.Request) {
//	db := dbConn()
//	//	Codigo := mux.Vars(r)["codigo"]
//
//	t := []tercero{}
//	var e empresa = ListaEmpresa()
//	var c ciudad = TraerCiudad(e.Ciudad)
//	err := db.Select(&t, "SELECT * FROM tercero ORDER BY cast(codigo as integer) ")
//	if err != nil {
//		log.Fatalln(err)
//	}
//	var buf bytes.Buffer
//	var err1 error
//	pdf := gofpdf.New("P", "mm", "Letter", cnFontDir)
//	ene := pdf.UnicodeTranslatorFromDescriptor("")
//	pdf.SetHeaderFunc(func() {
//		pdf.Image(imageFile("logo.png"), 20, 30, 40, 0, false,
//			"", 0, "")
//		pdf.SetY(17)
//		pdf.SetFont("Arial", "", 10)
//		pdf.CellFormat(190, 10, e.Nombre, "0", 0,
//			"C", false, 0, "")
//		pdf.Ln(4)
//		pdf.CellFormat(190, 10, "Nit No. "+Coma(e.Codigo)+" - "+e.Dv, "0", 0, "C",
//			false, 0, "")
//		pdf.Ln(4)
//		pdf.CellFormat(190, 10, e.Iva+" - "+e.ReteIva, "0", 0, "C", false, 0,
//			"")
//		pdf.Ln(4)
//		pdf.CellFormat(190, 10, "Actividad Ica - "+e.ActividadIca, "0",
//			0, "C", false, 0, "")
//		pdf.Ln(4)
//		pdf.CellFormat(190, 10, e.Direccion, "0", 0, "C", false, 0,
//			"")
//		pdf.Ln(4)
//		pdf.CellFormat(190, 10, e.Telefono1+" "+e.Telefono2, "0", 0, "C", false, 0,
//			"")
//		pdf.Ln(4)
//		pdf.CellFormat(190, 10, ene(c.NombreCiudad+" - "+c.NombreDepartamento), "0", 0, "C", false, 0,
//			"")
//		pdf.Ln(6)
//		pdf.CellFormat(190, 10, "DATOS TERCERO", "0", 0,
//			"C", false, 0, "")
//		pdf.Ln(10)
//	})
//
//	pdf.SetFooterFunc(func() {
//		pdf.SetTextColor(0, 0, 0)
//		pdf.SetY(252)
//		pdf.SetFont("Arial", "", 9)
//		pdf.SetX(20)
//
//		// LINEA
//		//pdf.Line(20,239,204,259)
//		//pdf.Ln(6)
//		pdf.SetY(20)
//		pdf.SetX(20)
//		pdf.CellFormat(90, 10, "Sadconf.com", "", 0,
//			"L", false, 0, "")
//		pdf.SetX(129)
//		pdf.CellFormat(80, 10, fmt.Sprintf(" %d de {nb}", pdf.PageNo()), "",
//			0, "R", false, 0, "")
//	})
//
//	pdf.AliasNbPages("")
//	pdf.AddPage()
//	pdf.SetFont("Arial", "", 10)
//	pdf.SetX(30)
//
//	TerceroTodosCabecera(pdf)
//	// tercera pagina
//
//	for i, miFila := range t {
//		var a = i + 1
//		if math.Mod(float64(a), 49) == 0 {
//			pdf.AliasNbPages("")
//			pdf.AddPage()
//			pdf.SetFont("Arial", "", 10)
//			pdf.SetX(30)
//			TerceroTodosCabecera(pdf)
//		}
//		TerceroTodosDetalle(pdf, miFila, a)
//	}
//	//BalancePieDePagina(pdf)
//
//	err1 = pdf.Output(&buf)
//	if err1 != nil {
//		panic(err1.Error())
//	}
//	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
//	w.Write(buf.Bytes())
//}
//
//// TERMINA TERCERO TODOS PDF
//
//// TERCERO EXCEL
//func TerceroExcel(w http.ResponseWriter, r *http.Request) {
//	db := dbConn()
//	t := []tercero{}
//	var e empresa = ListaEmpresa()
//	var c ciudad = TraerCiudad(e.Ciudad)
//	err := db.Select(&t, "SELECT * FROM tercero ORDER BY cast(codigo as integer) ")
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	f := excelize.NewFile()
//
//	// FUNCION ANCHO DE LA COLUMNA
//	if err = f.SetColWidth("Sheet1", "A", "A", 15); err != nil {
//		fmt.Println(err)
//		return
//	}
//	if err = f.SetColWidth("Sheet1", "B", "B", 3); err != nil {
//		fmt.Println(err)
//		return
//	}
//	if err = f.SetColWidth("Sheet1", "C", "C", 30); err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	if err = f.SetColWidth("Sheet1", "D", "D", 20); err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	if err = f.SetColWidth("Sheet1", "E", "e", 20); err != nil {
//		fmt.Println(err)
//		return
//	}
//	// FUNCION PARA UNIR DOS CELDAS
//	if err = f.MergeCell("Sheet1", "A1", "E1"); err != nil {
//		fmt.Println(err)
//		return
//	}
//	if err = f.MergeCell("Sheet1", "A2", "E2"); err != nil {
//		fmt.Println(err)
//		return
//	}
//	if err = f.MergeCell("Sheet1", "A3", "E3"); err != nil {
//		fmt.Println(err)
//		return
//	}
//	if err = f.MergeCell("Sheet1", "A4", "E4"); err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	if err = f.MergeCell("Sheet1", "A5", "E5"); err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	if err = f.MergeCell("Sheet1", "A6", "E6"); err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	if err = f.MergeCell("Sheet1", "A7", "E7"); err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	if err = f.MergeCell("Sheet1", "A8", "E8"); err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	if err = f.MergeCell("Sheet1", "A9", "E9"); err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	if err = f.MergeCell("Sheet1", "A10", "E10"); err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	estiloTitulo, err := f.NewStyle(`{  "alignment":{"horizontal": "center"},"font":{"bold":false,"italic":false,"family":"Arial","size":10,"color":"##000000"}}`)
//
//	// titulo
//	f.SetCellValue("Sheet1", "A1", e.Nombre)
//	f.SetCellValue("Sheet1", "A2", "Nit No. "+Coma(e.Codigo)+" - "+e.Dv)
//	f.SetCellValue("Sheet1", "A3", e.Iva+" - "+e.ReteIva)
//	f.SetCellValue("Sheet1", "A4", "Actividad Ica - "+e.ActividadIca)
//	f.SetCellValue("Sheet1", "A5", e.Direccion)
//	f.SetCellValue("Sheet1", "A6", (e.Telefono1 + " - " + e.Telefono2))
//	f.SetCellValue("Sheet1", "A7", (c.NombreCiudad + " - " + c.NombreDepartamento))
//	f.SetCellValue("Sheet1", "A8", "")
//	f.SetCellValue("Sheet1", "A9", "LISTADO DE TERCEROS")
//	f.SetCellValue("Sheet1", "A10", "")
//
//	f.SetCellStyle("Sheet1", "A1", "A1", estiloTitulo)
//	f.SetCellStyle("Sheet1", "A2", "A2", estiloTitulo)
//	f.SetCellStyle("Sheet1", "A3", "A3", estiloTitulo)
//	f.SetCellStyle("Sheet1", "A4", "A4", estiloTitulo)
//	f.SetCellStyle("Sheet1", "A5", "A5", estiloTitulo)
//	f.SetCellStyle("Sheet1", "A6", "A6", estiloTitulo)
//	f.SetCellStyle("Sheet1", "A7", "A7", estiloTitulo)
//	f.SetCellStyle("Sheet1", "A8", "A8", estiloTitulo)
//	f.SetCellStyle("Sheet1", "A9", "A9", estiloTitulo)
//	f.SetCellStyle("Sheet1", "A10", "A10", estiloTitulo)
//
//	var filaExcel = 11
//
//	estiloTexto, err := f.NewStyle(`{"font":{"bold":false,"italic":false,"family":"Arial","size":10,"color":"#000000"}}`)
//
//	estiloCabecera, err := f.NewStyle(`{
//"alignment":{"horizontal":"center"},
//    "border": [
//    {
//        "type": "left",
//        "color": "#000000",
//        "style": 1
//    },
//    {
//        "type": "top",
//        "color": "#000000",
//        "style": 1
//    },
//    {
//        "type": "bottom",
//        "color": "#000000",
//        "style": 1
//    },
//    {
//        "type": "right",
//        "color": "#000000",
//        "style": 1
//    }]
//}`)
//	if err != nil {
//		fmt.Println(err)
//	}
//	estiloNumeroDetalle, err := f.NewStyle(`{"number_format": 3,"font":{"bold":false,"italic":false,"family":"Arial","size":10,"color":"##000000"}}`)
//
//	if err != nil {
//		fmt.Println(err)
//	}
//	//cabecera
//	f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel), "Codigo")
//	f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel), "Dv")
//	f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel), "Nombre")
//	f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel), "Direccion")
//	f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel), "Telefono")
//
//	f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel), "A"+strconv.Itoa(filaExcel), estiloCabecera)
//	f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel), "B"+strconv.Itoa(filaExcel), estiloCabecera)
//	f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel), "C"+strconv.Itoa(filaExcel), estiloCabecera)
//	f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel), "D"+strconv.Itoa(filaExcel), estiloCabecera)
//	f.SetCellStyle("Sheet1", "E"+strconv.Itoa(filaExcel), "E"+strconv.Itoa(filaExcel), estiloCabecera)
//	filaExcel++
//
//	for i, miFila := range t {
//		f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel+i), Entero(miFila.Codigo))
//		f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel+i), Entero(miFila.Dv))
//		f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel+i), miFila.Nombre)
//		f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel+i), miFila.Direccion)
//		f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel+i), miFila.Telefono1)
//
//		f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel+i), "A"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
//		f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel+i), "B"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
//		f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel+i), "C"+strconv.Itoa(filaExcel+i), estiloTexto)
//		f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel+i), "D"+strconv.Itoa(filaExcel+i), estiloTexto)
//		f.SetCellStyle("Sheet1", "E"+strconv.Itoa(filaExcel+i), "E"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
//
//		//van=i
//	}
//
//	// LIENA FINAL
//	//a=strconv.Itoa(van+1+filaExcel)
//	// Set the headers necessary to get browsers to interpret the downloadable file
//	w.Header().Set("Content-Type", "application/octet-stream")
//	w.Header().Set("Content-Disposition", "attachment;filename=userInputData.xlsx")
//	w.Header().Set("File-Name", "userInputData.xlsx")
//	w.Header().Set("Content-Transfer-Encoding", "binary")
//	w.Header().Set("Expires", "0")
//	err = f.Write(w)
//	if err != nil {
//		panic(err.Error())
//	}
//}
//
////
//// INICIA TERCERO PDF
////func TerceroPdf(w http.ResponseWriter, r *http.Request) {
////	db := dbConn()
////	Codigo := mux.Vars(r)["codigo"]
////	t := tercero{}
////	var e  empresa=ListaEmpresa()
////	var c  ciudad=TraerCiudad(e.Ciudad)
////	err := db.Get(&t, "SELECT * FROM tercero where codigo=$1", Codigo)
////	if err != nil {
////		log.Fatalln(err)
////	}
////	var buf bytes.Buffer
////	var err1 error
////	pdf := gofpdf.New("P", "mm", "Letter", cnFontDir)
////	ene := pdf.UnicodeTranslatorFromDescriptor("")
////	pdf.SetHeaderFunc(func() {
////		pdf.Image(imageFile("logo.png"), 20, 30, 40, 0, false,
////			"", 0, "")
////		pdf.SetY(17)
////		pdf.SetFont("Arial", "", 10)
////		pdf.CellFormat(190, 10, e.Nombre, "0", 0,
////			"C", false, 0, "")
////		pdf.Ln(4)
////		pdf.CellFormat(190, 10, "Nit No. " +Coma(e.Codigo)+ " - "+e.Dv, "0", 0, "C",
////			false, 0, "")
////		pdf.Ln(4)
////		pdf.CellFormat(190, 10, e.Iva+ " - "+e.ReteIva, "0", 0, "C", false, 0,
////			"")
////		pdf.Ln(4)
////		pdf.CellFormat(190, 10, "Actividad Ica - "+e.ActividadIca, "0",
////			0, "C", false, 0, "")
////		pdf.Ln(4)
////		pdf.CellFormat(190, 10, e.Direccion, "0", 0, "C", false, 0,
////			"")
////		pdf.Ln(4)
////		pdf.CellFormat(190, 10, e.Telefono1+" - "+e.Telefono2, "0", 0, "C", false, 0,
////			"")
////		pdf.Ln(4)
////		pdf.CellFormat(190, 10, ene(c.NombreCiudad+ " - "+c.NombreDepartamento), "0", 0, "C", false, 0,
////			"")
////		pdf.Ln(10)
////
////		// RELLENO TITULO
////		pdf.SetX(20)
////		pdf.SetFillColor(224,231,239)
////		pdf.SetTextColor(0,0,0)
////
////		pdf.SetX(20)
////		pdf.CellFormat(184, 6, "DATOS TERCERO", "0", 0,
////			"C", true, 0, "")
////		pdf.Ln(8)
////	})
////
////	pdf.SetTextColor(0,0,0)
////	pdf.SetX(21)
////	pdf.AliasNbPages("")
////	pdf.AddPage()
////	pdf.SetX(21)
////
////	pdf.CellFormat(40, 4, "Nit. No.", "", 0,
////		"", false, 0, "")
////	pdf.CellFormat(142, 4, Coma(t.Codigo)+ " - "+t.Dv, "", 0,
////		"", false, 0, "")
////	pdf.Ln(-1)
////	pdf.SetX(21)
////	pdf.CellFormat(40, 4, "Tipo:", "", 0,
////		"", false, 0, "")
////	pdf.CellFormat(40, 4, TraerTipo(t.Tipo), "", 0,
////		"", false, 0, "")
////	pdf.Ln(-1)
////	pdf.SetX(21)
////	pdf.CellFormat(40, 4, "Empresa:", "0", 0,
////		"", false, 0, "")
////	pdf.CellFormat(40, 4, t.Juridica, "", 0,
////		"", false, 0, "")
////	pdf.Ln(-1)
////	pdf.SetX(21)
////	pdf.CellFormat(40, 4, "Primer Nombre:", "", 0,
////		"", false, 0, "")
////	pdf.CellFormat(40, 4, t.PrimerNombre, "", 0,
////		"", false, 0, "")
////	pdf.Ln(-1)
////	pdf.SetX(21)
////	pdf.CellFormat(40, 4, "Segundo Nombre:", "0", 0,
////		"", false, 0, "")
////	pdf.CellFormat(40, 4, t.SegundoNombre, "", 0,
////		"", false, 0, "")
////	pdf.Ln(-1)
////	pdf.SetX(21)
////	pdf.CellFormat(40, 4, "Primer Apellido:", "", 0,
////		"", false, 0, "")
////	pdf.CellFormat(40, 4, t.PrimerApellido, "", 0,
////		"", false, 0, "")
////	pdf.Ln(-1)
////	pdf.SetX(21)
////	pdf.CellFormat(40, 4, "Segundo Apellido:", "0", 0,
////		"", false, 0, "")
////	pdf.CellFormat(40, 4, t.SegundoApellido, "", 0,
////		"", false, 0, "")
////	pdf.Ln(-1)
////	pdf.SetX(21)
////	pdf.CellFormat(40, 4, "Ret. Ica:", "", 0,
////		"", false, 0, "")
////	pdf.CellFormat(40, 4, (t.Ica), "", 0,
////		"", false, 0, "")
////	pdf.Ln(-1)
////	pdf.SetX(21)
////	pdf.CellFormat(40, 4, "Rut:", "", 0,
////		"", false, 0, "")
////	pdf.CellFormat(40, 4, t.Rut, "", 0,
////		"", false, 0, "")
////	pdf.Ln(-1)
////	pdf.SetX(21)
////	pdf.CellFormat(40, 4, "Direccion:", "0", 0,
////		"", false, 0, "")
////	pdf.CellFormat(40, 4, t.Direccion, "", 0,
////		"", false, 0, "")
////	pdf.Ln(-1)
////	pdf.SetX(21)
////	pdf.CellFormat(40, 4, "Barrio:", "", 0,
////		"", false, 0, "")
////	pdf.CellFormat(40, 4, t.Barrio, "", 0,
////		"", false, 0, "")
////	pdf.Ln(-1)
////	pdf.SetX(21)
////	pdf.CellFormat(40, 4, "Ciudad:", "0", 0,
////		"", false, 0, "")
////	pdf.CellFormat(40, 4, t.Ciudad, "", 0,
////		"", false, 0, "")
////	pdf.Ln(-1)
////	pdf.SetX(21)
////	pdf.CellFormat(40, 4, "E-mail 1:", "", 0,
////		"", false, 0, "")
////	pdf.CellFormat(40, 4, t.Email1, "", 0,
////		"", false, 0, "")
////	pdf.Ln(-1)
////	pdf.SetX(21)
////	pdf.CellFormat(40, 4, "E-mail 2:", "0", 0,
////		"", false, 0, "")
////	pdf.CellFormat(40, 4, t.Email2, "", 0,
////		"", false, 0, "")
////	pdf.Ln(-1)
////	pdf.SetX(21)
////	pdf.CellFormat(40, 4, "Contacto:", "", 0,
////		"", false, 0, "")
////	pdf.CellFormat(40, 4, t.Contacto, "", 0,
////		"", false, 0, "")
////	pdf.Ln(-1)
////	pdf.SetX(21)
////	pdf.CellFormat(40, 4, "Documento:", "0", 0,
////		"", false, 0, "")
////	pdf.CellFormat(40, 4, TraerDocumentoIdentificacion(t.Documento), "", 0,
////		"", false, 0, "")
////	pdf.Ln(-1)
////	pdf.SetX(21)
////	pdf.CellFormat(40, 4, "Regimen:", "", 0,
////		"", false, 0, "")
////	pdf.CellFormat(40, 4, TraerRegimen(t.Regimen), "", 0,
////		"", false, 0, "")
////	pdf.Ln(-1)
////	pdf.SetX(21)
////	pdf.CellFormat(40, 4, "Fiscal:", "0", 0,
////		"", false, 0, "")
////	pdf.CellFormat(40, 4, TraerFiscal(t.Fiscal), "", 0,
////		"", false, 0, "")
////	pdf.Ln(-1)
////	pdf.SetX(21)
////	pdf.CellFormat(40, 4, "Codigo No.", "0", 0,
////		"", false, 0, "")
////	pdf.CellFormat(40, 4, Coma(t.PhCodigo) + " - "+t.PhDv, "", 0,
////		"", false, 0, "")
////	pdf.Ln(-1)
////	pdf.SetX(21)
////	pdf.CellFormat(40, 4, "Nombre:", "0", 0,
////		"", false, 0, "")
////	pdf.CellFormat(40, 4, t.PhNombre, "", 0,
////		"", false, 0, "")
////	pdf.Ln(-1)
////	pdf.SetX(21)
////	pdf.CellFormat(40, 4, "Descuento 1:", "0", 0,
////		"", false, 0, "")
////	pdf.CellFormat(40, 4, Coma(t.Descuento1), "", 0,
////		"", false, 0, "")
////	pdf.Ln(-1)
////	pdf.SetX(21)
////	pdf.CellFormat(40, 4, "Descuento 2:", "", 0,
////		"", false, 0, "")
////	pdf.CellFormat(40, 4, Coma(t.Descuento2), "", 0,
////		"", false, 0, "")
////	pdf.Ln(-1)
////	pdf.SetX(21)
////	pdf.CellFormat(40, 4, "Cuota P:", "0", 0,
////		"", false, 0, "")
////	pdf.CellFormat(40, 4, Coma(t.Cuotap), "", 0,
////		"", false, 0, "")
////	pdf.Ln(-1)
////	pdf.SetX(21)
////	pdf.CellFormat(40, 4, "Cuota 1:", "", 0,
////		"", false, 0, "")
////	pdf.CellFormat(40, 4, Coma(t.Cuota1), "", 0,
////		"", false, 0, "")
////	pdf.Ln(-1)
////	pdf.SetX(21)
////	pdf.CellFormat(40, 4, "Cuota 2:", "0", 0,
////		"", false, 0, "")
////	pdf.CellFormat(40, 4, Coma(t.Cuota2), "", 0,
////		"", false, 0, "")
////	pdf.Ln(-1)
////	pdf.SetX(21)
////	pdf.CellFormat(40, 4, "Cuota 3:", "", 0,
////		"", false, 0, "")
////	pdf.CellFormat(40, 4, Coma(t.Cuota3), "", 0,
////		"", false, 0, "")
////	pdf.Ln(-1)
////	pdf.SetX(21)
////	pdf.CellFormat(40, 4, "Ph:", "", 0,
////		"", false, 0, "")
////	pdf.CellFormat(40, 4, t.Ph, "", 0,
////		"", false, 0, "")
////
////	pdf.SetFooterFunc(func() {
////		pdf.SetTextColor(0, 0, 0)
////		pdf.SetY(252)
////		pdf.SetFont("Arial", "", 9)
////		pdf.SetX(20)
////
////		// LINEA
////		pdf.Line(20,259,204,259)
////		pdf.Ln(6)
////		pdf.SetX(20)
////		pdf.CellFormat(90, 10, "Sadconf.com", "", 0,
////			"L", false, 0, "")
////		pdf.SetX(129)
////		pdf.CellFormat(80, 10, fmt.Sprintf(" %d de {nb}", pdf.PageNo()), "",
////			0, "R", false, 0, "")
////	})

package main

// INICIA TERCERO IMPORTAR PAQUETES
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
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

// TERMINA  a IMPORTAR PAQUETES

// INICIA RESIDENTE ESTRUCTURA JSON
type residenteJson struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Nombre string `json:"nombre"`
}

// TERMINA TERCERO ESTRUCTURA JSON
type residentelista struct {
	Fila      string
	Codigo    string
	Nit       string
	Direccion string
	Nombre    string
	Email1    string
}

// INICIA TERCERO ESTRUCTURA
type residente struct {
	Codigo           string
	Dv               string
	Nit              string
	Juridica         string
	PrimerNombre     string
	SegundoNombre    string
	PrimerApellido   string
	SegundoApellido  string
	Direccion        string
	Telefono1        string
	Telefono2        string
	Email1           string
	Email2           string
	Contacto         string
	Matricula        string
	Catastral        string
	Area             string
	Coeficiente      string
	Ciudad           string
	Descuento1       string
	Descuento2       string
	CuotaP           string
	Cuota1           string
	Cuota2           string
	Ciudadexpedicion string
	Fechaexpedicion  time.Time
	Fechanacimiento  time.Time
	Ciudadnacimiento string
	Nombre           string
	Extension        string
}

// INICIA TERCERO LISTA
func ResidenteLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/residente/residenteLista.html")
	log.Println("Error tercero 0")
	db := dbConn()
	res := []residentelista{}

	err := db.Select(&res, "SELECT ROW_NUMBER() OVER(ORDER BY codigo DESC)  as Fila ,codigo,nit, direccion, nombre, email1 FROM residente ORDER BY codigo DESC ")
	if err != nil {
		fmt.Println(err)
	}

	varmap := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
	}
	log.Println("Error residente888")
	tmp.Execute(w, varmap)
}

//INICIA TERCERO NUEVO
func ResidenteNuevo(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	Panel := mux.Vars(r)["panel"]
	Elemento := mux.Vars(r)["elemento"]
	t := residente{}

	log.Println("Error residente nuevo 2")
	parametros := map[string]interface{}{
		// INICIA TERCERO NUEVO AUTOCOMPLETADO
		"Codigo":   Codigo,
		"Panel":    Panel,
		"Elemento": Elemento,
		"hosting":  ruta,
		"ciudad":   ListaCiudad(),
		"copiar":   "False",
		"emp":      t,
		// TERMINA TERCERO NUEVO AUTOCOMPLETADO
	}
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/residente/residenteNuevo.html",
		"vista/residente/autocompletaTercerocrear.html",
		"vista/residente/autocompletaTercero.html")
	log.Println("Error tercero nuevo 3")
	tmp.Execute(w, parametros)
}

// INICIA TERCERO DUPLICAR
func ResidenteNuevoCopia(w http.ResponseWriter, r *http.Request) {
	log.Println("Error residente nuevo 1")
	Codigo := "False"
	Panel := "False"
	Elemento := "False"

	copiarCodigo := Quitacoma(mux.Vars(r)["copiacodigo"])
	log.Println("inicio tercero editar" + Codigo)

	db := dbConn()
	t := residente{}

	if copiarCodigo == "False" {

	} else {
		// traer comprobante

		err := db.Get(&t, "SELECT * FROM residente WHERE codigo=$1", copiarCodigo)
		if err != nil {
			log.Fatalln(err)
		}
	}

	log.Println("Error residente nuevo 2")
	parametros := map[string]interface{}{
		// INICIA TERCERO NUEVO AUTOCOMPLETADO
		"Codigo":   Codigo,
		"Panel":    Panel,
		"Elemento": Elemento,
		"hosting":  ruta,
		"ciudad":   ListaCiudad(),
		"emp":      t,
		"copiar":   "True",

		// TERMINA TERCERO NUEVO AUTOCOMPLETADO
	}
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/residente/residenteNuevo.html",
		"vista/residente/autocompletaTercero.html",
		"vista/residente/autocompletaTercerocrear.html")
	log.Println("Error residente nuevo 3")
	tmp.Execute(w, parametros)
}

// INICIA TERCERO INSERTAR
func ResidenteInsertar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	err := r.ParseForm()
	if err != nil {
		panic(err.Error())
	}
	var t residente

	decoder.RegisterConverter(time.Time{}, timeConverter)
	err = decoder.Decode(&t, r.PostForm)
	if err != nil {
		panic(err.Error())
	}
	var q string
	q = "insert into residente ("
	q += "Codigo,"
	q += "Dv,"
	q += "Nit,"
	q += "Juridica,"
	q += "PrimerNombre,"
	q += "SegundoNombre,"
	q += "PrimerApellido,"
	q += "SegundoApellido,"
	q += "Direccion,"
	q += "Telefono1,"
	q += "Telefono2,"
	q += "Email1,"
	q += "Email2,"
	q += "Contacto,"
	q += "Matricula,"
	q += "Catastral,"
	q += "Area,"
	q += "Coeficiente,"
	q += "Ciudad,"
	q += "Descuento1,"
	q += "Descuento2,"
	q += "CuotaP,"
	q += "Cuota1,"
	q += "Cuota2,"
	q += "Ciudadexpedicion,"
	q += "Fechaexpedicion,"
	q += "Fechanacimiento,"
	q += "Ciudadnacimiento,"
	q += "Nombre"
	q += " ) values("
	q += parametros(29)
	q += " ) "

	log.Println("Cadena SQL " + q)
	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	// INICIA GRABAR TERCERO INSERTAR
	t.Nit = Quitacoma(t.Nit)
	t.Descuento1 = Quitacoma(t.Descuento1)
	t.Descuento2 = Quitacoma(t.Descuento2)
	t.CuotaP = Quitacoma(t.CuotaP)
	t.Cuota1 = Quitacoma(t.Cuota1)
	t.Cuota2 = Quitacoma(t.Cuota2)
	t.PrimerNombre = Titulo(t.PrimerNombre)
	t.SegundoNombre = Titulo(t.SegundoNombre)
	t.PrimerApellido = Titulo(t.PrimerApellido)
	t.SegundoApellido = Titulo(t.SegundoApellido)
	t.Nombre = Titulo(t.PrimerNombre + " " + t.SegundoNombre + " " + t.PrimerApellido + " " + t.SegundoApellido)
	t.Direccion = Titulo(t.Direccion)
	t.Contacto = Titulo(t.Contacto)
	t.Email1 = Minuscula(t.Email1)
	t.Email2 = Minuscula(t.Email2)
	t.Ciudadexpedicion = Titulo(t.Ciudadexpedicion)
	t.Ciudadnacimiento = Titulo(t.Ciudadnacimiento)

	// TERMINA TERCERO GRABAR INSERTAR
	_, err = insForm.Exec(
		t.Codigo,
		t.Dv,
		t.Nit,
		t.Juridica,
		t.PrimerNombre,
		t.SegundoNombre,
		t.PrimerApellido,
		t.SegundoApellido,
		t.Direccion,
		t.Telefono1,
		t.Telefono2,
		t.Email1,
		t.Email2,
		t.Contacto,
		t.Matricula,
		t.Catastral,
		t.Area,
		t.Coeficiente,
		t.Ciudad,
		t.Descuento1,
		t.Descuento2,
		t.CuotaP,
		t.Cuota1,
		t.Cuota2,
		t.Ciudadexpedicion,
		t.Fechaexpedicion,
		t.Fechanacimiento,
		t.Ciudadnacimiento,
		t.Nombre)

	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/ResidenteLista", 301)
}

// INICIA TERCERO BUSCAR
func ResidenteBuscar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	Codigo = Quitacoma(Codigo)
	selDB, err := db.Query("SELECT codigo,"+
		"nombre FROM residente where (codigo LIKE '%' || $1 || '%')  or  (upper(nombre) LIKE '%' || $1 || '%') ORDER BY"+
		" codigo DESC", Mayuscula(Codigo))
	if err != nil {
		panic(err.Error())
	}
	var resJson []residenteJson
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
		resJson = append(resJson, residenteJson{id, label, value, nombre})
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

// TERMINA TERCERO BUSCAR

// INICIA TERCERO EXISTE
func ResidenteExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	Codigo = Quitacoma(Codigo)
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM residente  WHERE codigo=$1", Codigo)
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

// TERMINA TERCERO EXISTE

// INICIA TERCERO ACTUAL
func ResidenteActual(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	Codigo = Quitacoma(Codigo)

	t := residente{}
	var res []residente
	err := db.Get(&t, "SELECT * FROM residente where codigo=$1", Codigo)

	switch err {
	case nil:
		log.Printf("tercero found: %+v\n", t)
		res = append(res, t)
		data, err := json.Marshal(res)
		if err != nil {
			log.Fatalln(err)
		}
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)

	case sql.ErrNoRows:
		log.Println("residente NOT found, no error")

		var slice []string
		slice = make([]string, 0)
		data, _ := json.Marshal(slice)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)

	default:
		log.Printf("tercero error: %s\n", err)
	}

	log.Println("codigo nombre99" + t.Codigo)

}

// INICIA TERCERO EDITAR
func ResidenteEditar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio tercero editar" + Codigo)
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/residente/residenteEditar.html",
		"vista/residente/autocompletaTercerocrear.html",
		"vista/residente/autocompletaTercero.html")
	db := dbConn()
	t := residente{}
	err := db.Get(&t, "SELECT * FROM residente where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("codigo nombre99" + t.Codigo)
	varmap := map[string]interface{}{
		// INICIA TERCERO EDITAR AUTOCOMPLETADO
		"emp":     t,
		"hosting": ruta,
		"ciudad":  ListaCiudad(),

		// TERMINA TERCERO EDITAR AUTOCOMPLETADO
	}
	tmp.Execute(w, varmap)
}

// INICIA TERCERO ACTUALIZAR
func ResidenteActualizar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	err := r.ParseForm()
	if err != nil {
		panic(err.Error())
		// Handle error
	}
	var t residente
	// r.PostForm is a map of our POST form values
	// FORMATO DE FECHAS
	decoder.RegisterConverter(time.Time{}, timeConverter)
	err = decoder.Decode(&t, r.PostForm)
	if err != nil {
		// Handle error
		panic(err.Error())
	}
	var q string
	q = "UPDATE residente set "
	q += "Dv=$2,"
	q += "Nit=$3,"
	q += "Juridica=$4,"
	q += "PrimerNombre=$5,"
	q += "SegundoNombre=$6,"
	q += "PrimerApellido=$7,"
	q += "SegundoApellido=$8,"
	q += "Direccion=$9,"
	q += "Telefono2=$10,"
	q += "Telefono1=$11,"
	q += "Email1=$12,"
	q += "Email2=$13,"
	q += "Contacto=$14,"
	q += "Matricula=$15,"
	q += "Catastral=$16,"
	q += "Area=$17,"
	q += "Coeficiente=$18,"
	q += "Ciudad=$19,"
	q += "Descuento1=$20,"
	q += "Descuento2=$21,"
	q += "CuotaP=$22,"
	q += "Cuota1=$23,"
	q += "Cuota2=$24,"
	q += "Ciudadexpedicion=$25,"
	q += "Fechaexpedicion=$26,"
	q += "Fechanacimiento=$27,"
	q += "Ciudadnacimiento=$28,"
	q += "Nombre=$29"
	q += " where "
	q += "Codigo=$1"

	log.Println("cadena" + q)

	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	// INICIA GRABAR TERCERO ACTUALIZAR
	t.Nit = Quitacoma(t.Nit)
	t.Descuento1 = Quitacoma(t.Descuento2)
	t.Descuento2 = Quitacoma(t.Descuento2)
	t.CuotaP = Quitacoma(t.CuotaP)
	t.Cuota1 = Quitacoma(t.Cuota1)
	t.Cuota2 = Quitacoma(t.Cuota2)
	t.PrimerNombre = Titulo(t.PrimerNombre)
	t.SegundoNombre = Titulo(t.SegundoNombre)
	t.PrimerApellido = Titulo(t.PrimerApellido)
	t.SegundoApellido = Titulo(t.SegundoApellido)
	t.Nombre = Titulo(t.PrimerNombre + " " + t.SegundoNombre + " " + t.PrimerApellido + " " + t.SegundoApellido)
	t.Direccion = Titulo(t.Direccion)
	t.Contacto = Titulo(t.Contacto)
	t.Email1 = Minuscula(t.Email1)
	t.Email2 = Minuscula(t.Email2)
	t.Ciudadexpedicion = Titulo(t.Ciudadexpedicion)
	t.Ciudadnacimiento = Titulo(t.Ciudadnacimiento)

	// TERMINA GRABAR TERCERO ACTUALIZAR

	_, err = insForm.Exec(
		t.Codigo,
		t.Dv,
		t.Nit,
		t.Juridica,
		t.PrimerNombre,
		t.SegundoNombre,
		t.PrimerApellido,
		t.SegundoApellido,
		t.Direccion,
		t.Telefono1,
		t.Telefono2,
		t.Email1,
		t.Email2,
		t.Contacto,
		t.Matricula,
		t.Catastral,
		t.Area,
		t.Coeficiente,
		t.Ciudad,
		t.Descuento1,
		t.Descuento2,
		t.CuotaP,
		t.Cuota1,
		t.Cuota2,
		t.Ciudadexpedicion,
		t.Fechaexpedicion,
		t.Fechanacimiento,
		t.Ciudadnacimiento,
		t.Nombre)

	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/ResidenteLista", 301)

}

// INICIA TERCERO BORRAR
func ResidenteBorrar(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/residente/residenteBorrar.html")
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio tercero borrar" + Codigo)
	db := dbConn()
	t := residente{}
	err := db.Get(&t, "SELECT * FROM residente where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("codigo nombre99 borrar" + t.Codigo)
	varmap := map[string]interface{}{
		// INICIA TERCERO BORRAR AUTOCOMPLETADO
		"emp":     t,
		"hosting": ruta,
		"ciudad":  ListaCiudad(),
	}
	tmp.Execute(w, varmap)
}

// INICIA TERCERO ELIMINAR
func ResidenteEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	emp := mux.Vars(r)["codigo"]
	delForm, err := db.Prepare("DELETE from residente WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("Registro Eliminado" + emp)
	http.Redirect(w, r, "/ResidenteLista", 301)
}

// INICIA RESIDENTE PDF
func ResidentePdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	t := residente{}
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	log.Println("111999")
	err := db.Get(&t, "SELECT * FROM residente where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}

	var ciudadnacimiento ciudad = TraerCiudad(t.Ciudadnacimiento)
	var ciudadexpedicion ciudad = TraerCiudad(t.Ciudadexpedicion)
	var ciudad ciudad = TraerCiudad(t.Ciudad)

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
		pdf.CellFormat(184, 6, "DATOS RESIDENTES", "0", 0,
			"C", true, 0, "")
		pdf.Ln(8)
	})

	pdf.SetTextColor(0, 0, 0)
	pdf.SetX(21)
	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Codigo No.", "", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, Coma(t.Codigo), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Documento No.", "", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, Coma(t.Nit), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Dv:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Dv, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Fecha Expedicion:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Fechaexpedicion.Format("02-01-2006"), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Ciudad Expedicion:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, ene(ciudadexpedicion.CodigoDepartamento+ciudadexpedicion.CodigoCiudad+" - "+ciudadexpedicion.NombreCiudad+" - "+ciudadexpedicion.NombreDepartamento), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Fecha Nacimiento:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Fechanacimiento.Format("02-01-2006"), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Ciudad Nacimiento:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, ene(ciudadnacimiento.CodigoDepartamento+ciudadnacimiento.CodigoCiudad+" - "+ciudadnacimiento.NombreCiudad+" - "+ciudadnacimiento.NombreDepartamento), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Primer Nombre:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.PrimerNombre, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Segundo Nombre:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.SegundoNombre, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Primer Apellido:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.PrimerApellido, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Segundo Apellido:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.SegundoApellido, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Direccion:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, (t.Direccion), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Ciudad", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, ene(ciudad.CodigoDepartamento+ciudad.CodigoCiudad+" - "+ciudad.NombreCiudad+" - "+ciudad.NombreDepartamento), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Telefono:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Telefono1, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Telefono:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Telefono2, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "E-mail:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Email1, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "E-mail:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Email2, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Contacto:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Contacto, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Matricula:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Matricula, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Catastral:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Catastral, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Area Mts.:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Area, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Coeficiente:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Coeficiente, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Descuento 1:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Coma(t.Descuento1), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Descuento 2:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Coma(t.Descuento2), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Cuota:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Coma(t.CuotaP), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Cuota 1:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Coma(t.Cuota1), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Cuota 2:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Coma(t.Cuota2), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
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

// INICIA RESIDENTES TODOS PDF
func ResidenteTodosCabecera(pdf *gofpdf.Fpdf) {
	pdf.SetFont("Arial", "", 10)
	// RELLENO TITULO
	pdf.SetY(50)
	pdf.SetFillColor(224, 231, 239)
	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(7)
	pdf.SetX(20)
	pdf.CellFormat(181, 6, "No", "0", 0,
		"L", true, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(40, 6, "Codigo", "0", 0,
		"L", false, 0, "")
	pdf.SetX(60)
	pdf.CellFormat(40, 6, "Nombre", "0", 0,
		"L", false, 0, "")
	pdf.SetX(120)
	pdf.CellFormat(40, 6, "Direccion", "0", 0,
		"L", false, 0, "")
	pdf.SetX(171)
	pdf.CellFormat(40, 6, "Telefono", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}
func ResidenteTodosDetalle(pdf *gofpdf.Fpdf, t residente, a int) {
	pdf.SetFont("Arial", "", 9)

	pdf.SetX(21)
	pdf.CellFormat(181, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(40, 4, Coma(t.Codigo), "", 0,
		"L", false, 0, "")
	pdf.SetX(60)
	pdf.CellFormat(40, 4, t.Nombre, "", 0, "L", false, 0, "")
	pdf.SetX(120)
	pdf.CellFormat(40, 4, t.Direccion, "", 0,
		"L", false, 0, "")
	pdf.SetX(155)
	pdf.CellFormat(40, 4, t.Telefono1, "", 0,
		"R", false, 0, "")
	pdf.Ln(4)
}

func ResidenteTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	//	Codigo := mux.Vars(r)["codigo"]

	t := []residente{}
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	err := db.Select(&t, "SELECT * FROM residente ORDER BY cast(codigo as integer) ")
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
		pdf.CellFormat(190, 10, ene(c.NombreCiudad+" - "+c.NombreDepartamento), "0", 0, "C", false, 0,
			"")
		pdf.Ln(6)
		pdf.CellFormat(190, 10, "DATOS EMPLEDOS", "0", 0,
			"C", false, 0, "")
		pdf.Ln(10)
	})

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

	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)
	pdf.SetX(30)

	ResidenteTodosCabecera(pdf)
	// tercera pagina

	for i, miFila := range t {
		var a = i + 1
		if math.Mod(float64(a), 49) == 0 {
			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			ResidenteTodosCabecera(pdf)
		}
		ResidenteTodosDetalle(pdf, miFila, a)
	}
	//BalancePieDePagina(pdf)

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

// TERCERO EXCEL
func ResidenteExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	t := []residente{}
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	err := db.Select(&t, "SELECT * FROM residente ORDER BY cast(codigo as integer) ")
	if err != nil {
		log.Fatalln(err)
	}

	f := excelize.NewFile()

	// FUNCION ANCHO DE LA COLUMNA
	if err = f.SetColWidth("Sheet1", "A", "A", 15); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.SetColWidth("Sheet1", "B", "B", 50); err != nil {
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

	// FUNCION PARA UNIR DOS CELDAS
	if err = f.MergeCell("Sheet1", "A1", "D1"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A2", "D2"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A3", "D3"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A4", "D4"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A5", "D5"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A6", "D6"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A7", "D7"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A8", "D8"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A9", "D9"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A10", "D10"); err != nil {
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
	f.SetCellValue("Sheet1", "A9", "DATOS RESIDENTES")
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

	f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel), "A"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel), "B"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel), "C"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel), "D"+strconv.Itoa(filaExcel), estiloCabecera)
	filaExcel++

	for i, miFila := range t {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel+i), Entero(miFila.Codigo))
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel+i), miFila.Nombre)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel+i), miFila.Direccion)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel+i), miFila.Telefono1)

		f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel+i), "A"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel+i), "C"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel+i), "D"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel+i), "E"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)

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

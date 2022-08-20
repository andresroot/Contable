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

// INICIA EMPLEADO ESTRUCTURA JSON
type empleadoJson struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Nombre string `json:"nombre"`
}

// TERMINA TERCERO ESTRUCTURA JSON
type empleadolista struct {
	Fila         string
	Codigo       string
	Nombre       string
	Fechaingreso time.Time
	Sueldo       string
}

// INICIA TERCERO ESTRUCTURA
type empleado struct {
	Codigo                  string
	Tipodocumento           string
	PrimerNombre            string
	SegundoNombre           string
	PrimerApellido          string
	SegundoApellido         string
	Direccion               string
	Barrio                  string
	Telefono1               string
	Email1                  string
	Contacto                string
	Tipotrabajador          string
	Sueldo                  string
	Subtipotrabajador       string
	Fechaingreso            time.Time
	Fecharetiro             time.Time
	Tiempolaborado          string
	Salariointegral         string
	Altoriesgopension       string
	Tipocontrato            string
	Formadepago             string
	Metododepago            string
	Banco                   string
	Tipocuenta              string
	Fechaliquidacioninicio  time.Time
	Fechaliquidacionfinal   time.Time
	Ciudad                  string
	Ciudadexpedicion        string
	Fechanacimiento         time.Time
	Fechaexpedicion         time.Time
	Ciudadnacimiento        string
	Nombre                  string
	Transporte              string
	Numerocuenta            string
	Cargo                   string
	Tipopago                string
	Dotacion                string
	Salud                   string
	Saludnombre             string
	Pension                 string
	Pensionnombre           string
	Riesgos                 string
	Riesgosnombre           string
	Icbf                    string
	Icbfnombre              string
	Sena                    string
	Senanombre              string
	Caja                    string
	Cajanombre              string
	Telefonocontacto        string
	Direccionlaboral        string
	Ciudadlaboral           string
	Pensionvoluntaria       string
	Pensionvoluntarianombre string
}

// TERMINA TERCERO ESTRUCTURA

// INICIA TERCERO LISTA
func EmpleadoLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/empleado/empleadoLista.html")
	log.Println("Error tercero 0")
	db := dbConn()
	res := []empleadolista{}
	//resfila := []empleadolista{}

	db.Select(&res, "SELECT ROW_NUMBER() OVER(ORDER BY sueldo DESC)  as Fila ,codigo,nombre, fechaingreso, sueldo FROM empleado ORDER BY sueldo DESC ")

	varmap := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
	}
	log.Println("Error empleado888")
	tmp.Execute(w, varmap)
}

// TERMINA TERCERO LISTA

//INICIA TERCERO NUEVO
func EmpleadoNuevo(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	Panel := mux.Vars(r)["panel"]
	Elemento := mux.Vars(r)["elemento"]
	t := empleado{}

	log.Println("Error empleado nuevo 2")
	parametros := map[string]interface{}{
		// INICIA TERCERO NUEVO AUTOCOMPLETADO
		"Codigo":                  Codigo,
		"Panel":                   Panel,
		"Elemento":                Elemento,
		"hosting":                 ruta,
		"ciudad":                  ListaCiudad(),
		"tipoorganizacion":        ListaTipoOrganizacion(),
		"regimenfiscal":           ListaRegimenFiscal(),
		"responsabilidadfiscal":   ListaResponsabilidadFiscal(),
		"documentoidentificacion": ListaDocumentoIdentificacion(),
		"mediodepago":             ListaMedioDePago(),
		"formadepago":             ListaFormaDePago(),
		"copiar":                  "False",
		"emp":                     t,
		// TERMINA TERCERO NUEVO AUTOCOMPLETADO
	}
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/empleado/empleadoNuevo.html",
		"vista/empleado/autocompletaTercerocrear.html",
		"vista/empleado/autocompletaTercero.html")
	log.Println("Error tercero nuevo 3")
	tmp.Execute(w, parametros)
}

// INICIA TERCERO DUPLICAR
func EmpleadoNuevoCopia(w http.ResponseWriter, r *http.Request) {
	log.Println("Error empleado nuevo 1")
	Codigo := "False"
	Panel := "False"
	Elemento := "False"

	copiarCodigo := Quitacoma(mux.Vars(r)["copiacodigo"])
	log.Println("inicio tercero editar" + Codigo)

	db := dbConn()
	t := empleado{}

	if copiarCodigo == "False" {

	} else {
		// traer comprobante

		err := db.Get(&t, "SELECT * FROM empleado WHERE codigo=$1", copiarCodigo)
		if err != nil {
			log.Fatalln(err)
		}
	}

	log.Println("Error empleado nuevo 2")
	parametros := map[string]interface{}{
		// INICIA TERCERO NUEVO AUTOCOMPLETADO
		"Codigo":                  Codigo,
		"Panel":                   Panel,
		"Elemento":                Elemento,
		"hosting":                 ruta,
		"ciudad":                  ListaCiudad(),
		"tipoorganizacion":        ListaTipoOrganizacion(),
		"regimenfiscal":           ListaRegimenFiscal(),
		"responsabilidadfiscal":   ListaResponsabilidadFiscal(),
		"documentoidentificacion": ListaDocumentoIdentificacion(),
		"mediodepago":             ListaMedioDePago(),
		"formadepago":             ListaFormaDePago(),
		"emp":                     t,
		"copiar":                  "True",

		// TERMINA TERCERO NUEVO AUTOCOMPLETADO
	}
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/empleado/empleadoNuevo.html",
		"vista/empleado/autocompletaTercero.html",
		"vista/empleado/autocompletaTercerocrear.html")
	log.Println("Error empleado nuevo 3")
	tmp.Execute(w, parametros)
}

// INICIA TERCERO INSERTAR
func EmpleadoInsertar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	err := r.ParseForm()
	if err != nil {
		panic(err.Error())
	}
	var t empleado

	decoder.RegisterConverter(time.Time{}, timeConverter)
	err = decoder.Decode(&t, r.PostForm)
	if err != nil {
		panic(err.Error())
	}
	var q string
	q = "insert into empleado ("
	q += "Codigo,"
	q += "Tipodocumento,"
	q += "Fechaexpedicion,"
	q += "PrimerNombre,"
	q += "SegundoNombre,"
	q += "PrimerApellido,"
	q += "SegundoApellido,"
	q += "Direccion,"
	q += "Barrio,"
	q += "Telefono1,"
	q += "Email1,"
	q += "Contacto,"
	q += "Tipotrabajador,"
	q += "Sueldo,"
	q += "Subtipotrabajador,"
	q += "Fechaingreso,"
	q += "Fecharetiro,"
	q += "Tiempolaborado,"
	q += "Salariointegral,"
	q += "Altoriesgopension,"
	q += "Tipocontrato,"
	q += "Formadepago,"
	q += "Metododepago,"
	q += "Banco,"
	q += "Tipocuenta,"
	q += "Fechaliquidacioninicio,"
	q += "Fechaliquidacionfinal,"
	q += "Ciudad,"
	q += "Ciudadexpedicion,"
	q += "Fechanacimiento,"
	q += "Ciudadnacimiento,"
	q += "Nombre,"
	q += "Transporte,"
	q += "Numerocuenta,"
	q += "Cargo,"
	q += "Tipopago,"
	q += "Dotacion,"
	q += "Salud,"
	q += "Saludnombre,"
	q += "Pension,"
	q += "Pensionnombre,"
	q += "Riesgos,"
	q += "Riesgosnombre,"
	q += "Icbf,"
	q += "Icbfnombre,"
	q += "Sena,"
	q += "Senanombre,"
	q += "Caja,"
	q += "Cajanombre,"
	q += "Telefonocontacto,"
	q += "Direccionlaboral,"
	q += "Ciudadlaboral,"
	q += "Pensionvoluntaria,"
	q += "Pensionvoluntarianombre"
	q += " ) values("
	q += parametros(54)
	q += " ) "

	log.Println("Cadena SQL " + q)
	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	// INICIA GRABAR TERCERO INSERTAR
	t.Codigo = Quitacoma(t.Codigo)
	t.Tiempolaborado = Quitacoma(t.Tiempolaborado)
	t.Sueldo = Quitacoma(t.Sueldo)
	t.Transporte = Quitacoma(t.Transporte)
	t.Dotacion = Quitacoma(t.Dotacion)
	t.PrimerNombre = Titulo(t.PrimerNombre)
	t.SegundoNombre = Titulo(t.SegundoNombre)
	t.PrimerApellido = Titulo(t.PrimerApellido)
	t.SegundoApellido = Titulo(t.SegundoApellido)
	t.Nombre = Titulo(t.PrimerNombre + " " + t.SegundoNombre + " " + t.PrimerApellido + " " + t.SegundoApellido)
	t.Direccion = Titulo(t.Direccion)
	t.Barrio = Titulo(t.Barrio)
	t.Contacto = Titulo(t.Contacto)
	t.Email1 = Minuscula(t.Email1)
	t.Banco = Titulo(t.Banco)
	t.Cargo = Titulo(t.Cargo)
	t.Ciudadexpedicion = Titulo(t.Ciudadexpedicion)
	t.Ciudadnacimiento = Titulo(t.Ciudadnacimiento)
	t.Saludnombre = Titulo(t.Saludnombre)
	t.Pensionnombre = Titulo(t.Pensionnombre)
	t.Riesgosnombre = Titulo(t.Riesgosnombre)
	t.Icbfnombre = Titulo(t.Icbfnombre)
	t.Senanombre = Titulo(t.Senanombre)
	t.Cajanombre = Titulo(t.Cajanombre)
	t.Direccionlaboral = Titulo(t.Direccionlaboral)
	t.Ciudadlaboral = Titulo(t.Ciudadlaboral)
	t.Pensionvoluntarianombre = Titulo(t.Pensionvoluntarianombre)

	// TERMINA TERCERO GRABAR INSERTAR
	_, err = insForm.Exec(
		t.Codigo,
		t.Tipodocumento,
		t.Fechaexpedicion,
		t.PrimerNombre,
		t.SegundoNombre,
		t.PrimerApellido,
		t.SegundoApellido,
		t.Direccion,
		t.Barrio,
		t.Telefono1,
		t.Email1,
		t.Contacto,
		t.Tipotrabajador,
		t.Sueldo,
		t.Subtipotrabajador,
		t.Fechaingreso,
		t.Fecharetiro,
		t.Tiempolaborado,
		t.Salariointegral,
		t.Altoriesgopension,
		t.Tipocontrato,
		t.Formadepago,
		t.Metododepago,
		t.Banco,
		t.Tipocuenta,
		t.Fechaliquidacioninicio,
		t.Fechaliquidacionfinal,
		t.Ciudad,
		t.Ciudadexpedicion,
		t.Fechanacimiento,
		t.Ciudadnacimiento,
		t.Nombre,
		t.Transporte,
		t.Numerocuenta,
		t.Cargo,
		t.Tipopago,
		t.Dotacion,
		t.Salud,
		t.Saludnombre,
		t.Pension,
		t.Pensionnombre,
		t.Riesgos,
		t.Riesgosnombre,
		t.Icbf,
		t.Icbfnombre,
		t.Sena,
		t.Senanombre,
		t.Caja,
		t.Cajanombre,
		t.Telefonocontacto,
		t.Direccionlaboral,
		t.Ciudadlaboral,
		t.Pensionvoluntaria,
		t.Pensionvoluntarianombre)

	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/EmpleadoLista", 301)
}

// TERMINA TERCERO INSERTAR

// INICIA TERCERO BUSCAR
func EmpleadoBuscar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	Codigo = Quitacoma(Codigo)
	selDB, err := db.Query("SELECT codigo,"+
		"nombre FROM empleado where (codigo LIKE '%' || $1 || '%')  or  (upper(nombre) LIKE '%' || $1 || '%') ORDER BY"+
		" codigo DESC", Mayuscula(Codigo))
	if err != nil {
		panic(err.Error())
	}
	var resJson []empleadoJson
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
		resJson = append(resJson, empleadoJson{id, label, value, nombre})
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
func EmpleadoExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	Codigo = Quitacoma(Codigo)
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM empleado  WHERE codigo=$1", Codigo)
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
func EmpleadoActual(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	Codigo = Quitacoma(Codigo)

	t := empleado{}
	var res []empleado
	err := db.Get(&t, "SELECT * FROM empleado where codigo=$1", Codigo)

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
		log.Println("empleado NOT found, no error")

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
func EmpleadoEditar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio tercero editar" + Codigo)
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/empleado/empleadoEditar.html",
		"vista/empleado/autocompletaTercerocrear.html",
		"vista/empleado/autocompletaTercero.html")
	db := dbConn()
	t := empleado{}
	err := db.Get(&t, "SELECT * FROM empleado where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("codigo nombre99" + t.Codigo)
	varmap := map[string]interface{}{
		// INICIA TERCERO EDITAR AUTOCOMPLETADO
		"emp":                     t,
		"hosting":                 ruta,
		"ciudad":                  ListaCiudad(),
		"tipoorganizacion":        ListaTipoOrganizacion(),
		"regimenfiscal":           ListaRegimenFiscal(),
		"responsabilidadfiscal":   ListaResponsabilidadFiscal(),
		"documentoidentificacion": ListaDocumentoIdentificacion(),
		"mediodepago":             ListaMedioDePago(),
		"formadepago":             ListaFormaDePago(),

		// TERMINA TERCERO EDITAR AUTOCOMPLETADO
	}
	tmp.Execute(w, varmap)
}

// INICIA TERCERO ACTUALIZAR
func EmpleadoActualizar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	err := r.ParseForm()
	if err != nil {
		panic(err.Error())
		// Handle error
	}
	var t empleado
	// r.PostForm is a map of our POST form values
	// FORMATO DE FECHAS
	decoder.RegisterConverter(time.Time{}, timeConverter)
	err = decoder.Decode(&t, r.PostForm)
	if err != nil {
		// Handle error
		panic(err.Error())
	}
	var q string
	q = "UPDATE empleado set "
	q += "Tipodocumento=$2,"
	q += "Fechaexpedicion=$3,"
	q += "PrimerNombre=$4,"
	q += "SegundoNombre=$5,"
	q += "PrimerApellido=$6,"
	q += "SegundoApellido=$7,"
	q += "Direccion=$8,"
	q += "Barrio=$9,"
	q += "Telefono1=$10,"
	q += "Email1=$11,"
	q += "Contacto=$12,"
	q += "Tipotrabajador=$13,"
	q += "Sueldo=$14,"
	q += "Subtipotrabajador=$15,"
	q += "Fechaingreso=$16,"
	q += "Fecharetiro=$17,"
	q += "Tiempolaborado=$18,"
	q += "Salariointegral=$19,"
	q += "Altoriesgopension=$20,"
	q += "Tipocontrato=$21,"
	q += "Formadepago=$22,"
	q += "Metododepago=$23,"
	q += "Banco=$24,"
	q += "Tipocuenta=$25,"
	q += "Fechaliquidacioninicio=$26,"
	q += "Fechaliquidacionfinal=$27,"
	q += "Ciudad=$28,"
	q += "Ciudadexpedicion=$29,"
	q += "Fechanacimiento=$30,"
	q += "Ciudadnacimiento=$31,"
	q += "Nombre=$32,"
	q += "Transporte=$33,"
	q += "Numerocuenta=$34,"
	q += "Cargo=$35,"
	q += "Tipopago=$36,"
	q += "Dotacion=$37,"
	q += "Salud=$38,"
	q += "Saludnombre=$39,"
	q += "Pension=$40,"
	q += "Pensionnombre=$41,"
	q += "Riesgos=$42,"
	q += "Riesgosnombre=$43,"
	q += "Icbf=$44,"
	q += "Icbfnombre=$45,"
	q += "Sena=$46,"
	q += "Senanombre=$47,"
	q += "Caja=$48,"
	q += "Cajanombre=$49,"
	q += "Telefonocontacto=$50,"
	q += "Direccionlaboral=$51,"
	q += "Ciudadlaboral=$52,"
	q += "Pensionvoluntaria=$53,"
	q += "Pensionvoluntarianombre=$54"
	q += " where "
	q += "Codigo=$1"

	log.Println("cadena" + q)

	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	// INICIA GRABAR TERCERO ACTUALIZAR
	t.Codigo = Quitacoma(t.Codigo)
	t.Sueldo = Quitacoma(t.Sueldo)
	t.Dotacion = Quitacoma(t.Dotacion)
	t.Transporte = Quitacoma(t.Transporte)
	t.PrimerNombre = Titulo(t.PrimerNombre)
	t.SegundoNombre = Titulo(t.SegundoNombre)
	t.PrimerApellido = Titulo(t.PrimerApellido)
	t.SegundoApellido = Titulo(t.SegundoApellido)
	t.Nombre = Titulo(t.PrimerNombre + " " + t.SegundoNombre + " " + t.PrimerApellido + " " + t.SegundoApellido)
	t.Direccion = Titulo(t.Direccion)
	t.Barrio = Titulo(t.Barrio)
	t.Contacto = Titulo(t.Contacto)
	t.Email1 = Minuscula(t.Email1)
	t.Banco = Titulo(t.Banco)
	t.Cargo = Titulo(t.Cargo)
	t.Ciudadexpedicion = Titulo(t.Ciudadexpedicion)
	t.Ciudadnacimiento = Titulo(t.Ciudadnacimiento)
	t.Saludnombre = Titulo(t.Saludnombre)
	t.Pensionnombre = Titulo(t.Pensionnombre)
	t.Riesgosnombre = Titulo(t.Riesgosnombre)
	t.Icbfnombre = Titulo(t.Icbfnombre)
	t.Senanombre = Titulo(t.Senanombre)
	t.Cajanombre = Titulo(t.Cajanombre)
	t.Direccionlaboral = Titulo(t.Direccionlaboral)
	t.Ciudadlaboral = Titulo(t.Ciudadlaboral)
	t.Pensionvoluntarianombre = Titulo(t.Pensionvoluntarianombre)

	// TERMINA GRABAR TERCERO ACTUALIZAR

	_, err = insForm.Exec(
		t.Codigo,
		t.Tipodocumento,
		t.Fechaexpedicion,
		t.PrimerNombre,
		t.SegundoNombre,
		t.PrimerApellido,
		t.SegundoApellido,
		t.Direccion,
		t.Barrio,
		t.Telefono1,
		t.Email1,
		t.Contacto,
		t.Tipotrabajador,
		t.Sueldo,
		t.Subtipotrabajador,
		t.Fechaingreso,
		t.Fecharetiro,
		t.Tiempolaborado,
		t.Salariointegral,
		t.Altoriesgopension,
		t.Tipocontrato,
		t.Formadepago,
		t.Metododepago,
		t.Banco,
		t.Tipocuenta,
		t.Fechaliquidacioninicio,
		t.Fechaliquidacionfinal,
		t.Ciudad,
		t.Ciudadexpedicion,
		t.Fechanacimiento,
		t.Ciudadnacimiento,
		t.Nombre,
		t.Transporte,
		t.Numerocuenta,
		t.Cargo,
		t.Tipopago,
		t.Dotacion,
		t.Salud,
		t.Saludnombre,
		t.Pension,
		t.Pensionnombre,
		t.Riesgos,
		t.Riesgosnombre,
		t.Icbf,
		t.Icbfnombre,
		t.Sena,
		t.Senanombre,
		t.Caja,
		t.Cajanombre,
		t.Telefonocontacto,
		t.Direccionlaboral,
		t.Ciudadlaboral,
		t.Pensionvoluntaria,
		t.Pensionvoluntarianombre)

	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/EmpleadoLista", 301)

}

// INICIA TERCERO BORRAR
func EmpleadoBorrar(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/empleado/empleadoBorrar.html")
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio tercero borrar" + Codigo)
	db := dbConn()
	t := empleado{}
	err := db.Get(&t, "SELECT * FROM empleado where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("codigo nombre99 borrar" + t.Codigo)
	varmap := map[string]interface{}{
		// INICIA TERCERO BORRAR AUTOCOMPLETADO
		"emp":                     t,
		"hosting":                 ruta,
		"ciudad":                  ListaCiudad(),
		"tipoorganizacion":        ListaTipoOrganizacion(),
		"regimenfiscal":           ListaRegimenFiscal(),
		"responsabilidadfiscal":   ListaResponsabilidadFiscal(),
		"documentoidentificacion": ListaDocumentoIdentificacion(),
		"mediodepago":             ListaMedioDePago(),
		"formadepago":             ListaFormaDePago(),
		// TERMINA TERCERO BORRAR AUTOCOMPLETADO
	}
	tmp.Execute(w, varmap)
}

// INICIA TERCERO ELIMINAR
func EmpleadoEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	emp := mux.Vars(r)["codigo"]
	delForm, err := db.Prepare("DELETE from empleado WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("Registro Eliminado" + emp)
	http.Redirect(w, r, "/EmpleadoLista", 301)
}

// INICIA EMPLEADO PDF
func EmpleadoPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	t := empleado{}
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)

	err := db.Get(&t, "SELECT * FROM empleado where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}

	var ciudadnacimiento ciudad = TraerCiudad(t.Ciudadnacimiento)
	var ciudadexpedicion ciudad = TraerCiudad(t.Ciudadexpedicion)
	var ciudadlaboral ciudad = TraerCiudad(t.Ciudadlaboral)
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
		pdf.CellFormat(184, 6, "DATOS EMPLEADOS", "0", 0,
			"C", true, 0, "")
		pdf.Ln(8)
	})

	pdf.SetTextColor(0, 0, 0)
	pdf.SetX(21)
	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.SetX(21)

	pdf.CellFormat(50, 4, "Documento No.", "", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, Coma(t.Codigo), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Tipo de Documento:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, ene(TraerDocumentoIdentificacion(t.Tipodocumento)), "", 0,
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
	pdf.CellFormat(50, 4, "Barrio:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Barrio, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Ciudad Residencia:", "", 0,
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
	pdf.CellFormat(50, 4, "E-mail:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Email1, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Contacto:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Contacto, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Telefono Contacto:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Telefonocontacto, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Direccion Laboral:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Direccionlaboral, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Ciudad Laboral:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, ene(ciudadlaboral.CodigoDepartamento+ciudadlaboral.CodigoCiudad+" - "+ciudadlaboral.NombreCiudad+" - "+ciudadlaboral.NombreDepartamento), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Tipo de Trabajador:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Tipotrabajador, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Subtipo Trabajador:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Subtipotrabajador, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Cargo:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Cargo, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Sueldo:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Coma(t.Sueldo), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Transporte:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Coma(t.Transporte), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Fecha de Ingreso:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Fechaingreso.Format("02-01-2006"), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Fecha de Retiro:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Fecharetiro.Format("02-01-2006"), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Tiempo Laborado:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Tiempolaborado, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Salud:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Saludnombre, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Pension:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Pensionnombre, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Pension Voluntaria:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Pensionvoluntarianombre, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Riesgos laborales:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Riesgosnombre, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Icbf:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Icbfnombre, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Sena:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Senanombre, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Caja Compensacion:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Cajanombre, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Salario Integral:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Salariointegral, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Alto Riesgo de Pension:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Altoriesgopension, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Tipo Contrato:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Tipocontrato, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Forma de Pago:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Formadepago, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Metodo de Pago:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, TraerMediodepago(t.Metododepago), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Banco:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Banco, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Tipo de Cuenta:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Tipocuenta, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Numero de Cuenta:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Numerocuenta, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Tipo de Pago:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Tipopago, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Dotacion:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Coma(t.Dotacion), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Fecha Liquidacion Inicio:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Fechaliquidacioninicio.Format("02-01-2006"), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(50, 4, "Fecha de Liquidacion Final:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Fechaliquidacionfinal.Format("02-01-2006"), "", 0,
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

// INICIA EMPLEADOS TODOS PDF
func EmpleadoTodosCabecera(pdf *gofpdf.Fpdf) {
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
func EmpleadoTodosDetalle(pdf *gofpdf.Fpdf, t empleado, a int) {
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

func EmpleadoTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	//	Codigo := mux.Vars(r)["codigo"]

	t := []empleado{}
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	err := db.Select(&t, "SELECT * FROM empleado ORDER BY cast(codigo as integer) ")
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

	EmpleadoTodosCabecera(pdf)
	// tercera pagina

	for i, miFila := range t {
		var a = i + 1
		if math.Mod(float64(a), 49) == 0 {
			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			EmpleadoTodosCabecera(pdf)
		}
		EmpleadoTodosDetalle(pdf, miFila, a)
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
func EmpleadoExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	t := []empleado{}
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	err := db.Select(&t, "SELECT * FROM empleado ORDER BY cast(codigo as integer) ")
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
	f.SetCellValue("Sheet1", "A9", "DATOS EMPLEADOS")
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

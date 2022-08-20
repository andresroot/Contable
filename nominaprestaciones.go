package main

// INICIA NOMINA IMPORTAR PAQUETES
import (
	"bytes"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"math"

	//"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
	_ "github.com/lib/pq"
	"html/template"
	"io/ioutil"
	"log"
	//"math"
	"net/http"
	"strconv"
	"time"
)

// INICIA NOMINA ESTRUCTURA JSON
type nominaprestacionesJson struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Nombre string `json:"nombre"`
}

// INICIA NOMINA ESTRUCTURA
type nominaprestacionesLista struct {
	Numero       string
	Fechainicial time.Time
	Fechafinal   time.Time
	Cesantias    string
	Intereses    string
	Prima        string
	Vacaciones   string
	Dotaciones   string
	Total        string
}

// INICIA NOMINA ESTRUCTURA
type nominaprestacionesdetalle struct {
	Numero       string
	Fechainicial time.Time
	Fechafinal   time.Time
	Codigo       string
	Centro       string
	Sueldo       string
	Cesantias    string
	Intereses    string
	Prima        string
	Vacaciones   string
	Dotaciones   string
	Total        string
}

// INICIA PLANILLADETALLE ESTRUCTURA
type nominaprestaciones struct {
	Numero        string
	Fechainicial  time.Time
	Fechafinal    time.Time
	Cesantias     string
	Intereses     string
	Prima         string
	Vacaciones    string
	Dotaciones    string
	Total         string
	Detalle       []nominaprestacionesdetalle       `json:"Detalle"`
	DetalleEditar []nominaprestacionesdetalleeditar `json:"DetalleEditar"`
	Accion        string
}

// INICIA NOMINA DETALLE EDITARr
type nominaprestacionesdetalleeditar struct {
	Numero         string
	Fechainicial   time.Time
	Fechafinal     time.Time
	Codigo         string
	Empleadonombre string
	Centro         string
	Sueldo         string
	Cesantias      string
	Intereses      string
	Prima          string
	Vacaciones     string
	Dotaciones     string
	Total          string
}

// INICIA NOMINA CONSULTA DETALLE
func NominaprestacionesConsultaDetalle() string {
	var consulta = ""
	consulta = "select "
	consulta += "nominaprestacionesdetalle.Codigo as codigo,"
	consulta += "empleado.nombre as empleadonombre,"
	consulta += "nominaprestacionesdetalle.Centro as centro,"
	consulta += "nominaprestacionesdetalle.Sueldo as sueldo,"
	consulta += "nominaprestacionesdetalle.Cesantias as cesantias,"
	consulta += "nominaprestacionesdetalle.Intereses as intereses,"
	consulta += "nominaprestacionesdetalle.Prima as prima,"
	consulta += "nominaprestacionesdetalle.Vacaciones as vacaciones,"
	consulta += " nominaprestacionesdetalle.Dotaciones as dotaciones,"
	consulta += " nominaprestacionesdetalle.Total as total"
	consulta += " from nominaprestacionesdetalle "
	consulta += " inner join empleado on empleado.codigo=nominaprestacionesdetalle.codigo"
	consulta += " where nominaprestacionesdetalle.numero=$1 ORDER BY sueldo DESC"
	log.Println(consulta)
	return consulta
}

// INICIA NOMINA LISTA
func NominaprestacionesLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/nominaprestaciones/nominaprestacionesLista.html")
	log.Println("Error nominaprestaciones 0")
	var consulta string

	consulta = " SELECT nominaprestaciones.numero, nominaprestaciones.fechainicial, "
	consulta += " nominaprestaciones.fechafinal, nominaprestaciones.cesantias, nominaprestaciones.intereses,"
	consulta += " nominaprestaciones.prima, nominaprestaciones.vacaciones, nominaprestaciones.dotaciones, nominaprestaciones.total"
	consulta += " FROM nominaprestaciones"
	consulta += " ORDER BY cast(nominaprestaciones.numero as integer) ASC"

	db := dbConn()
	res := []nominaprestacionesLista{}

	err := db.Select(&res, consulta)

	if err != nil {
		fmt.Println(err)
	}

	varmap := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
	}
	log.Println("Error nominaprestaciones888")
	tmp.Execute(w, varmap)
}

// INICIA NOMINA NUEVO
func NominaprestacionesNuevo(w http.ResponseWriter, r *http.Request) {

	// TRAE COPIA DE EDITAR
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio nominaprestaciones editar" + Codigo)

	db := dbConn()
	empleados := []empleado{}
	nominaprestaciones := nominaprestaciones{}
	det := []nominaprestacionesdetalleeditar{}

	if Codigo == "False" {
		err := db.Select(&empleados, "SELECT * FROM empleado  ORDER BY sueldo DESC ")
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		// traer NOMINA

		err := db.Get(&nominaprestaciones, "SELECT * FROM nominaprestaciones where numero=$1", Codigo)
		if err != nil {
			log.Fatalln(err)
		}
		err2 := db.Select(&det, NominaprestacionesConsultaDetalle(), Codigo)
		if err2 != nil {
			fmt.Println(err2)
		}
	}

	parametros := map[string]interface{}{
		"codigo":             Codigo,
		"nominaprestaciones": nominaprestaciones,
		"empleado":           empleados,
		"detalle":            det,
		"hosting":            ruta,
	}
	//TERMINA TRAE COPIA DE EDITAR

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/nominaprestaciones/nominaprestacionesNuevo.html",
		"vista/nominaprestaciones/autocompletaCentro.html",
		"vista/nominaprestaciones/autocompletaempleado.html",
		"vista/nominaprestaciones/modalColumna.html",
		"vista/nominaprestaciones/nominaprestacionesScript.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error nominaprestaciones nuevo 3")
	miTemplate.Execute(w, parametros)
}

//INICIA NOMINA INSERTAR AJAX
func NominaprestacionesAgregar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var tempNominaprestaciones nominaprestaciones

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// carga informacion de la NOMINA
	err = json.Unmarshal(b, &tempNominaprestaciones)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	parametrosinventario := configuracioninventario{}
	err = db.Get(&parametrosinventario, "SELECT * FROM configuracioninventario")
	if err != nil {
		panic(err.Error())
	}

	var Codigoactual string

	Codigoactual = tempNominaprestaciones.Numero

	if tempNominaprestaciones.Accion == "Actualizar" {

		// borra detalle anterior
		delForm, err := db.Prepare("DELETE from nominaprestacionesdetalle WHERE numero=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(tempNominaprestaciones.Numero)

		// borra detalle inventario
		Borrarinventario(Codigoactual, "Nominaprestaciones")

		// borra cabecera anterior
		delForm1, err := db.Prepare("DELETE from nominaprestaciones WHERE numero=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm1.Exec(tempNominaprestaciones.Numero)
	}

	// INSERTA DETALLE
	for i, x := range tempNominaprestaciones.Detalle {
		var a = i
		var q string
		q = "insert into nominaprestacionesdetalle ("
		q += "Numero,"
		q += "Fechainicial,"
		q += "Fechafinal,"
		q += "Codigo,"
		q += "Centro,"
		q += "Sueldo,"
		q += "Cesantias,"
		q += "Intereses,"
		q += "Prima,"
		q += "Vacaciones,"
		q += "Dotaciones,"
		q += "Total"
		q += " ) values("
		q += parametros(12)
		q += ")"
		log.Println("Cadena SQL " + q)
		insForm, err := db.Prepare(q)
		if err != nil {
			panic(err.Error())
		}

		// TERMINA NOMINA GRABAR INSERTAR
		_, err = insForm.Exec(
			x.Numero,
			x.Fechainicial,
			x.Fechafinal,
			Quitacoma(x.Codigo),
			x.Centro,
			Quitacoma(x.Sueldo),
			Quitacoma(x.Cesantias),
			Quitacoma(x.Intereses),
			Quitacoma(x.Prima),
			Quitacoma(x.Vacaciones),
			Quitacoma(x.Dotaciones),
			Quitacoma(x.Total))

		if err != nil {
			panic(err)
		}
		log.Println("Insertar Codigo \n", x.Codigo, a)
	}

	// INICIA INSERTAR PLANILLAS
	var q string
	q = "insert into nominaprestaciones ("
	q += "Numero,"
	q += "Fechainicial,"
	q += "Fechafinal,"
	q += "Cesantias,"
	q += "Intereses,"
	q += "Prima,"
	q += "Vacaciones,"
	q += "Dotaciones,"
	q += "Total"
	q += " ) values("
	q += parametros(9)
	q += " ) "

	log.Println("Cadena SQL " + q)
	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	layout := "2006-01-02"
	//log.Println("Hora", tempNominaprestaciones.Fechainicial.Format("02/01/2006"))
	//currentTime := time.Now()

	_, err = insForm.Exec(
		tempNominaprestaciones.Numero,
		tempNominaprestaciones.Fechainicial.Format(layout),
		tempNominaprestaciones.Fechafinal.Format(layout),
		Quitacoma(tempNominaprestaciones.Cesantias),
		Quitacoma(tempNominaprestaciones.Intereses),
		Quitacoma(tempNominaprestaciones.Prima),
		Quitacoma(tempNominaprestaciones.Vacaciones),
		Quitacoma(tempNominaprestaciones.Dotaciones),
		Quitacoma(tempNominaprestaciones.Total))

	if err != nil {
		panic(err)
	}

	// INSERTAR COMPROBANTE CONTABILIDAD
	var tempComprobante comprobante
	var tempComprobanteDetalle comprobantedetalle
	tempComprobante.Documento = "18"
	tempComprobante.Numero = tempNominaprestaciones.Numero
	tempComprobante.Fecha = tempNominaprestaciones.Fechafinal
	tempComprobante.Fechaconsignacion = tempNominaprestaciones.Fechafinal
	tempComprobante.Debito = tempNominaprestaciones.Total
	tempComprobante.Credito = tempNominaprestaciones.Total
	tempComprobante.Periodo = ""
	tempComprobante.Licencia = ""
	tempComprobante.Usuario = ""
	tempComprobante.Estado = ""

	//borra detalle anterior
	delForm, err := db.Prepare("DELETE from comprobantedetalle WHERE documento=$1 and numero=$2")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(tempComprobante.Documento, tempComprobante.Numero)

	// borra cabecera anterior

	delForm1, err := db.Prepare("DELETE from comprobante WHERE documento=$1 and numero=$2 ")
	if err != nil {
		panic(err.Error())
	}
	delForm1.Exec(tempComprobante.Documento, tempComprobante.Numero)

	// INSERTAR CABECERA COMPROBANTE

	log.Println("Got %s age %s club %s\n", tempComprobante.Documento, tempComprobante.Numero)

	var totalDebito float64
	var totalCredito float64
	var fila int
	fila = 0
	totalDebito = 0
	totalCredito = 0

	for i, x := range tempNominaprestaciones.Detalle {

		// CESANTIAS
		if Flotante(x.Cesantias) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Cesantias)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Cesantias
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Cesantias))
			tempComprobanteDetalle.Credito = ""
			InsertaDetalleComprobanteNominaprestaciones(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// INTERESES
		if Flotante(x.Intereses) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Intereses)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Intereses
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Intereses))
			tempComprobanteDetalle.Credito = ""
			InsertaDetalleComprobanteNominaprestaciones(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// PRIMA
		if Flotante(x.Prima) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Prima)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Prima
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Prima))
			tempComprobanteDetalle.Credito = ""
			InsertaDetalleComprobanteNominaprestaciones(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// VACACIONES
		if Flotante(x.Vacaciones) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Vacaciones)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Vacaciones
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Vacaciones))
			tempComprobanteDetalle.Credito = ""
			InsertaDetalleComprobanteNominaprestaciones(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// DOTACIONES
		if Flotante(x.Dotaciones) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Dotaciones)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Dotaciones
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Dotaciones))
			tempComprobanteDetalle.Credito = ""
			InsertaDetalleComprobanteNominaprestaciones(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// CESANTIAS CREDITO
		if Flotante(x.Cesantias) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Cesantias)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Cesantiascxp
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Cesantias))
			InsertaDetalleComprobanteNominaprestaciones(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}

		// INTERESES CREDITO
		if Flotante(x.Intereses) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Intereses)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Interesescxp
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Intereses))
			InsertaDetalleComprobanteNominaprestaciones(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}

		// PRIMA CREDITO
		if Flotante(x.Prima) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Prima)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Primacxp
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Prima))
			InsertaDetalleComprobanteNominaprestaciones(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}

		// VACACIONES CREDITO
		if Flotante(x.Vacaciones) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Vacaciones)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Vacacionescxp
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Vacaciones))
			InsertaDetalleComprobanteNominaprestaciones(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}

		// DOTACIONES CREDITO
		if Flotante(x.Dotaciones) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Dotaciones)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Dotacionescxp
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Dotaciones))
			InsertaDetalleComprobanteNominaprestaciones(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}
	}

	// INSERTAR CUENTA DEBITO NOMINA

	var cadenaDebito = FormatoFlotante(totalDebito)
	var cadenaCredito = FormatoFlotante(totalCredito)

	q = "insert into comprobante ("
	q += "Documento,"
	q += "Numero,"
	q += "Fecha,"
	q += "Fechaconsignacion,"
	q += "Debito,"
	q += "Credito,"
	q += "Periodo,"
	q += "Licencia,"
	q += "Usuario,"
	q += "Estado"
	q += ") values("
	q += parametros(10)
	q += ")"

	log.Println("Cadena SQL " + q)
	insForm, err = db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	_, err = insForm.Exec(
		tempComprobante.Documento,
		tempComprobante.Numero,
		tempComprobante.Fecha.Format(layout),
		tempComprobante.Fechaconsignacion.Format(layout),
		Quitacoma(cadenaDebito),
		Quitacoma(cadenaCredito),
		tempComprobante.Periodo,
		tempComprobante.Licencia,
		tempComprobante.Usuario,
		tempComprobante.Estado)
	if err != nil {
		panic(err)
	}

	var resultado bool
	resultado = true

	js, err := json.Marshal(SomeStruct{resultado})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}
func InsertaDetalleComprobanteNominaprestaciones(miFilaComprobante comprobantedetalle, miComprobante comprobante, miCompra nominaprestacionesdetalle) {
	db := dbConn()

	// traer tercero
	miTercero := tercero{}
	err3 := db.Get(&miTercero, "SELECT * FROM tercero where codigo=$1", miCompra.Codigo)
	if err3 != nil {
		log.Fatalln(err3)
	}

	var q string
	q = "insert into comprobantedetalle ("
	q += "Fila,"
	q += "Cuenta,"
	q += "Tercero,"
	q += "Centro,"
	q += "Concepto,"
	q += "Factura,"
	q += "Debito,"
	q += "Credito,"
	q += "Documento,"
	q += "Numero,"
	q += "Fecha,"
	q += "Fechaconsignacion"
	q += " ) values("
	q += parametros(12)
	q += ")"
	log.Println("Cadena SQL " + q)
	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	if len(miFilaComprobante.Debito) > 0 {
		miFilaComprobante.Debito = miFilaComprobante.Debito
	}

	if len(miFilaComprobante.Credito) > 0 {
		miFilaComprobante.Credito = miFilaComprobante.Credito
	}

	// TERMINA COMPROBANTE GRABAR INSERTAR
	_, err = insForm.Exec(
		miFilaComprobante.Fila,
		miFilaComprobante.Cuenta,
		miCompra.Codigo,
		miCompra.Centro,
		miTercero.Nombre,
		"",
		Flotantedatabase(miFilaComprobante.Debito),
		Flotantedatabase(miFilaComprobante.Credito),
		miComprobante.Documento,
		miComprobante.Numero,
		miComprobante.Fecha,
		miComprobante.Fechaconsignacion)
	if err != nil {
		panic(err)
	}
}

// INICIA NOMINA EXISTE
func NominaprestacionesExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM nominaprestaciones  WHERE codigo=$1", Codigo)
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

//
//// INICIA NOMINA EDITAR
func NominaprestacionesEditar(w http.ResponseWriter, r *http.Request) {
	Numero := mux.Vars(r)["codigo"]
	log.Println("inicio nominaprestaciones editar" + Numero)
	db := dbConn()

	// traer NOMINA
	v := nominaprestaciones{}
	err := db.Get(&v, "SELECT * FROM nominaprestaciones where numero=$1", Numero)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []nominaprestacionesdetalleeditar{}

	err2 := db.Select(&det, NominaprestacionesConsultaDetalle(), Numero)
	if err2 != nil {
		fmt.Println(err2)
	}
	parametros := map[string]interface{}{
		"nominaprestaciones": v,
		"detalle":            det,
		"hosting":            ruta,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/nominaprestaciones/nominaprestacionesEditar.html",
		"vista/nominaprestaciones/autocompletaCentro.html",
		"vista/nominaprestaciones/autocompletaempleado.html",
		"vista/nominaprestaciones/modalColumna.html",
		"vista/nominaprestaciones/nominaprestacionesScript.html")

	fmt.Printf("%v, %v", miTemplate, err)

	miTemplate.Execute(w, parametros)
}

// INICIA NOMINA BORRAR
func NominaprestacionesBorrar(w http.ResponseWriter, r *http.Request) {
	Numero := mux.Vars(r)["codigo"]
	log.Println("inicio nominaprestaciones editar" + Numero)
	db := dbConn()

	// traer NOMINA
	v := nominaprestaciones{}
	err := db.Get(&v, "SELECT * FROM nominaprestaciones where numero=$1", Numero)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []nominaprestacionesdetalleeditar{}

	err2 := db.Select(&det, NominaprestacionesConsultaDetalle(), Numero)
	if err2 != nil {
		fmt.Println(err2)
	}
	parametros := map[string]interface{}{
		"codigo":             Numero,
		"nominaprestaciones": v,
		"detalle":            det,
		"hosting":            ruta,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/nominaprestaciones/nominaprestacionesBorrar.html",
		"vista/nominaprestaciones/autocompletaCentro.html",
		"vista/nominaprestaciones/autocompletaempleado.html",
		"vista/nominaprestaciones/modalColumna.html",
		"vista/nominaprestaciones/nominaprestacionesScript.html")

	fmt.Printf("%v, %v", miTemplate, err)

	miTemplate.Execute(w, parametros)
}

// INICIA NOMINA ELIMINAR
func NominaprestacionesEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	codigo := mux.Vars(r)["codigo"]

	// borrar NOMINA
	delForm, err := db.Prepare("DELETE from nominaprestaciones WHERE numero=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(codigo)

	// borar detalle
	delForm1, err := db.Prepare("DELETE from nominaprestacionesdetalle WHERE numero=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm1.Exec(codigo)

	// borar detalle invenario

	// borra detalle anterior
	delForm, err = db.Prepare("DELETE from comprobantedetalle WHERE documento=$1 and numero=$2")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec("6", codigo)

	// borra cabecera anterior

	delForm1, err = db.Prepare("DELETE from comprobante WHERE documento=$1 and numero=$2 ")
	if err != nil {
		panic(err.Error())
	}
	delForm1.Exec("6", codigo)

	log.Println("Registro Eliminado" + codigo)
	http.Redirect(w, r, "/NominaprestacionesLista", 301)
}

// INICIA NOMINA PDF
func NominaprestacionesPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Numero := mux.Vars(r)["numero"]

	// traer NOMINA
	miNominaprestaciones := nominaprestaciones{}
	err := db.Get(&miNominaprestaciones, "SELECT * FROM nominaprestaciones where numero=$1", Numero)
	if err != nil {
		log.Fatalln(err)
	}

	// traer detalle
	miDetalle := []nominaprestacionesdetalleeditar{}
	err2 := db.Select(&miDetalle, NominaprestacionesConsultaDetalle(), Numero)
	if err2 != nil {
		fmt.Println(err2)
	}
	//miDetalle.TerceroDetalle = TraerTerceroConsulta(miDetalle.Codigo)

	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
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
		pdf.CellFormat(190, 10, "Nit No. "+Coma(e.Codigo)+" - "+e.Dv, "0",
			0, "C", false, 0, "")
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
		pdf.CellFormat(190, 10, e.Telefono1+" - "+e.Telefono2, "0", 0, "C", false, 0, "")
		pdf.Ln(4)
		pdf.CellFormat(190, 10, ene(c.NombreCiudad+" - "+c.NombreDepartamento), "0", 0, "C", false, 0,
			"")
	})

	pdf.SetFooterFunc(func() {
		pdf.SetFont("Arial", "", 8)
		pdf.SetY(259)
		pdf.SetX(20)
		pdf.CellFormat(80, 10, "Andres Eduardo Ojeda Medina Nit."+
			" 80.853.536-7 SADCONF Derechos de Autor 13-16-230 de 30-06-2006  www.Sadconf.com.co", "",
			0, "L", false, 0, "")
		pdf.SetX(129)
		pdf.CellFormat(80, 10, fmt.Sprintf(" %d de {nb}", pdf.PageNo()), "",
			0, "R", false, 0, "")
	})

	// inicia suma de paginas
	pdf.AliasNbPages("")

	for a, miFila := range miDetalle {
		pdf.AddPage()
		pdf.Ln(1)
		NominaprestacionesCabecera(pdf, miFila)
		numerofila = 1
		NominaprestacionesFilaDetalle(pdf, miFila.Sueldo, a, "SUELDO")
		NominaprestacionesFilaDetalle(pdf, miFila.Vacaciones, a, "VACACIONES")
		NominaprestacionesFilaDetalle(pdf, miFila.Cesantias, a, "CESANTIAS")
		NominaprestacionesFilaDetalle(pdf, miFila.Intereses, a, "INTERESES")
		NominaprestacionesFilaDetalle(pdf, miFila.Prima, a, "PRIMA")
		NominaprestacionesFilaDetalle(pdf, miFila.Vacaciones, a, "VACACIONES")
		NominaprestacionesFilaDetalle(pdf, miFila.Dotaciones, a, "DOTACIONES")
		NominaprestacionesFilaDetalle(pdf, miFila.Total, a, "TOTAL")
	}

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

func NominaprestacionesCabecera(pdf *gofpdf.Fpdf, miNominadetalle nominaprestacionesdetalleeditar) {
	//ene := pdf.UnicodeTranslatorFromDescriptor("")
	// RELLENO TITULO

	pdf.SetFont("Arial", "", 10)
	pdf.SetY(44)
	pdf.SetX(20)
	pdf.SetFillColor(224, 231, 239)
	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(8)
	pdf.SetX(20)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(40, 4, "Nomina Numero:", "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, miNominadetalle.Numero, "", 0,
		"L", false, 0, "")
	pdf.SetX(60)
	pdf.CellFormat(40, 4, "Prefijo:", "", 0,
		"L", false, 0, "")
	pdf.SetX(75)
	//pdf.CellFormat(40, 4, miNominadetalle.Prefijo, "", 0,
	//	"L", false, 0, "")
	pdf.SetX(90)
	pdf.CellFormat(40, 4, "Fecha Inicial:", "", 0,
		"L", false, 0, "")
	pdf.SetX(115)
	pdf.CellFormat(40, 4, miNominadetalle.Fechainicial.Format("02/01/2006"), "", 0,
		"L", false, 0, "")
	pdf.SetX(138)
	pdf.CellFormat(40, 4, "Fecha Final:", "", 0,
		"L", false, 0, "")
	pdf.SetX(160)
	pdf.CellFormat(40, 4, miNominadetalle.Fechafinal.Format("02/01/2006"), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Cedula No.:", "", 0,
		"L", false, 0, "")
	pdf.SetX(35)
	pdf.CellFormat(40, 4, miNominadetalle.Codigo, "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, "Cedula No.:", "", 0,
		"L", false, 0, "")
	pdf.SetX(70)
	pdf.CellFormat(40, 4, miNominadetalle.Empleadonombre, "", 0,
		"L", false, 0, "")

	// RELLENO TITULO
	pdf.SetFont("Arial", "", 10)
	pdf.SetY(60)
	pdf.SetX(20)
	pdf.SetFillColor(224, 231, 239)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetX(20)

	pdf.CellFormat(184, 6, "ITEM", "0", 0,
		"L", true, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(190, 6, "", "0", 0,
		"L", false, 0, "")
	pdf.SetX(55)
	pdf.CellFormat(190, 6, "CONCEPTO", "0", 0,
		"L", false, 0, "")
	pdf.SetX(106)
	pdf.CellFormat(40, 6, "VALOR", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}
func NominaprestacionesFilaDetalle(pdf *gofpdf.Fpdf, miFila string, a int, miConcepto string) {

	if Flotante(miFila) != 0 {
		ene := pdf.UnicodeTranslatorFromDescriptor("")

		pdf.SetFont("Arial", "", 9)
		pdf.SetTextColor(0, 0, 0)
		pdf.SetX(21)
		pdf.CellFormat(183, 4, strconv.Itoa(numerofila), "", 0,
			"L", false, 0, "")
		pdf.SetX(30)
		pdf.CellFormat(40, 4, "", "", 0,
			"L", false, 0, "")
		pdf.SetX(55)
		pdf.CellFormat(40, 4, ene(miConcepto), "", 0,
			"L", false, 0, "")
		pdf.SetX(80)
		pdf.CellFormat(40, 4, Coma(miFila), "", 0,
			"R", false, 0, "")
		pdf.SetX(100)

		pdf.Ln(-1)
		numerofila++
	}

}
func NominaprestacionesPieDePagina(pdf *gofpdf.Fpdf, miNominaprestaciones nominaprestaciones) {

	Totalletras, err := IntLetra(Cadenaentero(miNominaprestaciones.Total))
	if err != nil {
		fmt.Println(err)
	}

	ene := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetFont("Arial", "", 8)
	pdf.SetY(222)
	pdf.SetX(20)
	pdf.MultiCell(184, 4, "SON: "+ene(Mayuscula(Totalletras))+" PESOS MDA. CTE.", "0", "L", false)
	pdf.SetFont("Arial", "", 9)
	pdf.SetY(230)

	// PRESENTA DATOS CON VALORES //
	//pdf.SetFont("Arial", "", 9)
	//var separador float64
	//var altoseparador float64
	//separador = 250
	//altoseparador = -4
}

func NominaprestacionesLinea(pdf *gofpdf.Fpdf) {
	pdf.SetY(259)
	pdf.Line(20, 259, 204, 259)
}

// INICIA NOMINA TODOS PDF
func NominaprestacionesTodosCabecera(pdf *gofpdf.Fpdf) {
	pdf.SetFont("Arial", "", 10)
	// RELLENO TITULO
	pdf.SetY(50)
	pdf.SetFillColor(224, 231, 239)
	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(7)
	pdf.SetX(20)
	pdf.CellFormat(184, 6, "No", "0", 0,
		"L", true, 0, "")
	pdf.SetX(32)
	pdf.CellFormat(190, 6, "Numero", "0", 0,
		"L", false, 0, "")
	pdf.SetX(55)
	pdf.CellFormat(190, 6, "Prefijo", "0", 0,
		"L", false, 0, "")
	pdf.SetX(75)
	pdf.CellFormat(190, 6, "Fecha Inicial", "0", 0,
		"L", false, 0, "")
	pdf.SetX(130)
	pdf.CellFormat(190, 6, "Fecha Final", "0", 0,
		"L", false, 0, "")
	pdf.SetX(193)
	pdf.CellFormat(190, 6, "Devengado", "0", 0,
		"L", false, 0, "")
	pdf.SetX(205)
	pdf.CellFormat(190, 6, "Neto", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}

func NominaprestacionesTodosDetalle(pdf *gofpdf.Fpdf, miFila nominaprestacionesLista, a int) {
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(21)
	pdf.CellFormat(181, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(32)
	pdf.CellFormat(40, 4, Subcadena(miFila.Numero, 0, 12), "", 0,
		"L", false, 0, "")
	pdf.SetX(55)
	pdf.CellFormat(40, 4, miFila.Fechainicial.Format("02/01/2006"), "", 0, "L", false, 0, "")
	pdf.SetX(75)
	pdf.CellFormat(40, 4, miFila.Fechafinal.Format("02/01/2006"), "", 0, "L", false, 0, "")
	pdf.SetX(164)
	pdf.CellFormat(40, 4, Coma(miFila.Total), "", 0,
		"R", false, 0, "")
	pdf.Ln(4)
}

func NominaprestacionesTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string

	consulta = " SELECT nominaprestaciones.numero, nominaprestaciones.prefijo, nominaprestaciones.fechainicial, "
	consulta += " nominaprestaciones.fechafinal, nominaprestaciones.devengado, nominaprestaciones.deducciones,"
	consulta += " nominaprestaciones.neto"
	consulta += " FROM nominaprestaciones"
	consulta += " ORDER BY cast(nominaprestaciones.numero as integer) ASC"

	t := []nominaprestacionesLista{}
	err := db.Select(&t, consulta)

	if err != nil {
		fmt.Println(err)
		return
	}

	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
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
		pdf.CellFormat(190, 10, "DATOS DOCUMENTO NOMINA", "0", 0,
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

	PedidoTodosCabecera(pdf)
	// tercera pagina

	for i, miFila := range t {
		var a = i + 1
		if math.Mod(float64(a), 49) == 0 {
			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			PedidoTodosCabecera(pdf)
		}
		NominaprestacionesTodosDetalle(pdf, miFila, a)
	}

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

// TERMINA PEDIDO TODOS PDF

// NOMINA EXCEL
func NominaprestacionesIndividualExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]

	// traer NOMINA
	miNominaprestaciones := nominaprestaciones{}
	err := db.Get(&miNominaprestaciones, "SELECT * FROM nominaprestaciones where numero=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}

	// traer detalle
	miDetalle := []nominaprestacionesdetalleeditar{}
	err2 := db.Select(&miDetalle, NominaprestacionesConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	f := excelize.NewFile()

	// FUNCION ANCHO DE LA COLUMNA
	if err = f.SetColWidth("Sheet1", "A", "A", 13); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.SetColWidth("Sheet1", "B", "B", 35); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.SetColWidth("Sheet1", "C", "C", 30); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "D", "D", 30); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "E", "E", 15); err != nil {
		fmt.Println(err)
		return
	}

	// FUNCION PARA UNIR DOS CELDAS
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
	f.SetCellValue("Sheet1", "A9", "LISTADO DE DOCUMENTOS NOMINA")
	f.SetCellValue("Sheet1", "A10", "")

	f.SetCellValue("Sheet1", "A12", "Numero")
	f.SetCellValue("Sheet1", "B12", miNominaprestaciones.Numero)
	f.SetCellValue("Sheet1", "C12", "Fecha Inicial")
	f.SetCellValue("Sheet1", "D12", miNominaprestaciones.Fechainicial)
	f.SetCellValue("Sheet1", "J12", "Fecha Final")
	f.SetCellValue("Sheet1", "K12", miNominaprestaciones.Fechafinal)

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
	var filaExcel = 15

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
	f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel), "Centro")
	f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel), "Sueldo")
	f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel), "Cesantias")
	f.SetCellValue("Sheet1", "F"+strconv.Itoa(filaExcel), "Intereses")
	f.SetCellValue("Sheet1", "G"+strconv.Itoa(filaExcel), "Prima")
	f.SetCellValue("Sheet1", "H"+strconv.Itoa(filaExcel), "Vacaciones")
	f.SetCellValue("Sheet1", "I"+strconv.Itoa(filaExcel), "Total")

	f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel), "A"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel), "B"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel), "C"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel), "D"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "E"+strconv.Itoa(filaExcel), "E"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "F"+strconv.Itoa(filaExcel), "F"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "G"+strconv.Itoa(filaExcel), "G"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "H"+strconv.Itoa(filaExcel), "H"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "I"+strconv.Itoa(filaExcel), "I"+strconv.Itoa(filaExcel), estiloCabecera)

	filaExcel++
	var totalSueldos float64
	var totalCesantias float64
	var totalIntereses float64
	var totalPrima float64
	var totalVacaciones float64
	var totalDotaciones float64
	var totalTotal float64

	var i int
	var ultima int

	for i, miFila := range miDetalle {
		totalSueldos += Flotante(miFila.Sueldo)
		totalCesantias += Flotante(miFila.Cesantias)
		totalIntereses += Flotante(miFila.Intereses)
		totalPrima += Flotante(miFila.Prima)
		totalVacaciones += Flotante(miFila.Vacaciones)
		totalDotaciones += Flotante(miFila.Dotaciones)
		totalTotal += Flotante(miFila.Total)

		f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel+i), miFila.Codigo)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel+i), miFila.Empleadonombre)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel+i), miFila.Centro)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel+i), Flotante(miFila.Sueldo))
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel+i), Flotante(miFila.Cesantias))
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(filaExcel+i), Flotante(miFila.Intereses))
		f.SetCellValue("Sheet1", "G"+strconv.Itoa(filaExcel+i), Flotante(miFila.Prima))
		f.SetCellValue("Sheet1", "H"+strconv.Itoa(filaExcel+i), Flotante(miFila.Vacaciones))
		f.SetCellValue("Sheet1", "I"+strconv.Itoa(filaExcel+i), Flotante(miFila.Dotaciones))
		f.SetCellValue("Sheet1", "J"+strconv.Itoa(filaExcel+i), Flotante(miFila.Total))

		f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel+i), "A"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel+i), "B"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel+i), "C"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel+i), "D"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "E"+strconv.Itoa(filaExcel+i), "E"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "F"+strconv.Itoa(filaExcel+i), "F"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "G"+strconv.Itoa(filaExcel+i), "G"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "H"+strconv.Itoa(filaExcel+i), "H"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "I"+strconv.Itoa(filaExcel+i), "I"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "J"+strconv.Itoa(filaExcel+i), "J"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)

		ultima = filaExcel + i

	}

	filaExcel = ultima + 1

	// totales
	f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel), "")
	f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel), "TOTALES")
	f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel), "")
	f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel), "")
	f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel), "")
	f.SetCellValue("Sheet1", "F"+strconv.Itoa(filaExcel), totalSueldos)
	f.SetCellValue("Sheet1", "G"+strconv.Itoa(filaExcel), totalCesantias)
	f.SetCellValue("Sheet1", "H"+strconv.Itoa(filaExcel), totalIntereses)
	f.SetCellValue("Sheet1", "I"+strconv.Itoa(filaExcel), totalPrima)
	f.SetCellValue("Sheet1", "J"+strconv.Itoa(filaExcel), totalCesantias)
	f.SetCellValue("Sheet1", "K"+strconv.Itoa(filaExcel), totalDotaciones)
	f.SetCellValue("Sheet1", "L"+strconv.Itoa(filaExcel), totalTotal)

	f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel), "A"+strconv.Itoa(filaExcel+i), estiloTexto)
	f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel), "B"+strconv.Itoa(filaExcel+i), estiloTexto)
	f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel), "C"+strconv.Itoa(filaExcel+i), estiloTexto)
	f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel), "D"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
	f.SetCellStyle("Sheet1", "E"+strconv.Itoa(filaExcel), "E"+strconv.Itoa(filaExcel+i), estiloTexto)
	f.SetCellStyle("Sheet1", "F"+strconv.Itoa(filaExcel), "F"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
	f.SetCellStyle("Sheet1", "G"+strconv.Itoa(filaExcel), "G"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
	f.SetCellStyle("Sheet1", "H"+strconv.Itoa(filaExcel), "H"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
	f.SetCellStyle("Sheet1", "I"+strconv.Itoa(filaExcel), "I"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
	f.SetCellStyle("Sheet1", "J"+strconv.Itoa(filaExcel), "J"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)

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

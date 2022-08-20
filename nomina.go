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
type nominaJson struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Nombre string `json:"nombre"`
}

// INICIA NOMINA ESTRUCTURA
type nominaLista struct {
	Numero       string
	Prefijo      string
	Fechainicial time.Time
	Fechafinal   time.Time
	Devengado    string
	Deducciones  string
	Neto         string
}

// INICIA NOMINA ESTRUCTURA
type nominadetalle struct {
	Numero            string
	Prefijo           string
	Fechainicial      time.Time
	Fechafinal        time.Time
	Codigo            string
	Centro            string
	Sueldo            string
	Dias              string
	Trabajado         string
	Transporte        string
	Cesantias         string
	Intereses         string
	Prima             string
	Vacaciones        string
	Viaticos          string
	Horasextras       string
	Incapacidades     string
	Licencias         string
	Bonificaciones    string
	Auxilios          string
	Huelgas           string
	Conceptos         string
	Compensaciones    string
	Bonos             string
	Comisiones        string
	Dotaciones        string
	Sostenimiento     string
	Teletrabajo       string
	Indemnizaciones   string
	Devengado         string
	Salud             string
	Pension           string
	Pensionrais       string
	Pensionvoluntaria string
	Solidaridad       string
	Subsistencia      string
	Sindicatos        string
	Sanciones         string
	Libranzas         string
	Terceros          string
	Anticipos         string
	Otras             string
	Retencion         string
	Afc               string
	Embargos          string
	Educacion         string
	Deuda             string
	Deducciones       string
	Neto              string
	TerceroDetalle    tercero
}

// INICIA PLANILLADETALLE ESTRUCTURA
type nomina struct {
	Numero        string
	Prefijo       string
	Fechainicial  time.Time
	Fechafinal    time.Time
	Detalle       []nominadetalle       `json:"Detalle"`
	DetalleEditar []nominadetalleeditar `json:"DetalleEditar"`
	Accion        string

	Trabajado         string
	Transporte        string
	Cesantias         string
	Intereses         string
	Prima             string
	Vacaciones        string
	Viaticos          string
	Horasextras       string
	Incapacidades     string
	Licencias         string
	Bonificaciones    string
	Auxilios          string
	Huelgas           string
	Conceptos         string
	Compensaciones    string
	Bonos             string
	Comisiones        string
	Dotaciones        string
	Sostenimiento     string
	Teletrabajo       string
	Indemnizaciones   string
	Devengado         string
	Salud             string
	Pension           string
	Pensionrais       string
	Pensionvoluntaria string
	Solidaridad       string
	Subsistencia      string
	Sindicatos        string
	Sanciones         string
	Libranzas         string
	Terceros          string
	Anticipos         string
	Otras             string
	Retencion         string
	Afc               string
	Embargos          string
	Educacion         string
	Deuda             string
	Deducciones       string
	Neto              string
}

// INICIA NOMINA DETALLE EDITARr
type nominadetalleeditar struct {
	Numero            string
	Prefijo           string
	Fechainicial      time.Time
	Fechafinal        time.Time
	Codigo            string
	Empleadonombre    string
	Centro            string
	Sueldo            string
	Dias              string
	Trabajado         string
	Transporte        string
	Cesantias         string
	Intereses         string
	Prima             string
	Vacaciones        string
	Viaticos          string
	Horasextras       string
	Incapacidades     string
	Licencias         string
	Bonificaciones    string
	Auxilios          string
	Huelgas           string
	Conceptos         string
	Compensaciones    string
	Bonos             string
	Comisiones        string
	Dotaciones        string
	Sostenimiento     string
	Teletrabajo       string
	Indemnizaciones   string
	Devengado         string
	Salud             string
	Pension           string
	Pensionrais       string
	Pensionvoluntaria string
	Solidaridad       string
	Subsistencia      string
	Sindicatos        string
	Sanciones         string
	Libranzas         string
	Terceros          string
	Anticipos         string
	Otras             string
	Retencion         string
	Afc               string
	Embargos          string
	Educacion         string
	Deuda             string
	Deducciones       string
	Neto              string
}

// INICIA NOMINA CONSULTA DETALLE
func NominaConsultaDetalle() string {
	var consulta = ""
	consulta = "select "
	consulta += "nominadetalle.Numero as numero,"
	consulta += "nominadetalle.Prefijo as prefijo,"
	consulta += "nominadetalle.Fechainicial as fechainicial,"
	consulta += "nominadetalle.Fechafinal as fechafinal,"
	consulta += "nominadetalle.Codigo as codigo,"
	consulta += "empleado.nombre as empleadonombre,"
	consulta += "nominadetalle.Centro as centro,"
	consulta += "nominadetalle.Sueldo as sueldo,"
	consulta += "nominadetalle.Dias as dias,"
	consulta += "nominadetalle.Trabajado as trabajado,"
	consulta += "nominadetalle.Transporte as transporte,"
	consulta += "nominadetalle.Cesantias as cesantias,"
	consulta += "nominadetalle.Intereses as intereses,"
	consulta += "nominadetalle.Prima as prima,"
	consulta += "nominadetalle.Vacaciones as vacaciones,"
	consulta += "nominadetalle.Viaticos as viaticos,"
	consulta += "nominadetalle.Horasextras as horasextras,"
	consulta += "nominadetalle.Incapacidades as incapacidades,"
	consulta += "nominadetalle.Licencias as Licencias,"
	consulta += "nominadetalle.Bonificaciones as bonificaciones,"
	consulta += "nominadetalle.Auxilios as auxilios,"
	consulta += "nominadetalle.Huelgas as huelgas,"
	consulta += "nominadetalle.Conceptos as conceptos,"
	consulta += "nominadetalle.Compensaciones as compensaciones,"
	consulta += "nominadetalle.Bonos as bonos,"
	consulta += "nominadetalle.Comisiones as comisiones,"
	consulta += " nominadetalle.Dotaciones as dotaciones,"
	consulta += " nominadetalle.Sostenimiento as sostenimiento,"
	consulta += " nominadetalle.Teletrabajo as teletrabajo,"
	consulta += " nominadetalle.indemnizaciones as indemnizaciones,"
	consulta += " nominadetalle.Devengado as devengado,"
	consulta += " nominadetalle.Salud as salud,"
	consulta += " nominadetalle.Pension as pension,"
	consulta += " nominadetalle.Pensionrais as pensionrais,"
	consulta += " nominadetalle.Pensionvoluntaria as pensionvoluntaria,"
	consulta += " nominadetalle.Solidaridad as solidaridad,"
	consulta += " nominadetalle.Subsistencia as subsistencia,"
	consulta += " nominadetalle.Sindicatos as Sindicatos,"
	consulta += " nominadetalle.Sanciones as sanciones,"
	consulta += " nominadetalle.Libranzas as libranzas,"
	consulta += " nominadetalle.Terceros as terceros,"
	consulta += " nominadetalle.Anticipos as anticipos,"
	consulta += " nominadetalle.Otras as otras,"
	consulta += " nominadetalle.Retencion as retencion,"
	consulta += " nominadetalle.Afc as afc,"
	consulta += " nominadetalle.Embargos as embargos,"
	consulta += " nominadetalle.Educacion as educacion,"
	consulta += " nominadetalle.Deuda as deuda,"
	consulta += " nominadetalle.Deducciones as deducciones,"
	consulta += " nominadetalle.Neto as neto"
	consulta += " from nominadetalle "
	consulta += " inner join empleado on empleado.codigo=nominadetalle.codigo"
	consulta += " where nominadetalle.numero=$1 ORDER BY sueldo DESC"
	log.Println(consulta)
	return consulta
}

// INICIA NOMINA LISTA
func NominaLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/nomina/nominaLista.html")
	log.Println("Error nomina 0")
	var consulta string

	consulta = " SELECT nomina.numero, nomina.prefijo, nomina.fechainicial, "
	consulta += " nomina.fechafinal, nomina.devengado, nomina.deducciones,"
	consulta += " nomina.neto"
	consulta += " FROM nomina"
	consulta += " ORDER BY cast(nomina.numero as integer) ASC"

	db := dbConn()
	res := []nominaLista{}

	err := db.Select(&res, consulta)

	if err != nil {
		fmt.Println(err)
	}

	varmap := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
	}
	log.Println("Error nomina888")
	tmp.Execute(w, varmap)
}

// INICIA NOMINA NUEVO
func NominaNuevo(w http.ResponseWriter, r *http.Request) {

	// TRAE COPIA DE EDITAR
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio nomina editar" + Codigo)

	db := dbConn()
	empleados := []empleado{}
	nomina := nomina{}
	det := []nominadetalleeditar{}

	if Codigo == "False" {
		err := db.Select(&empleados, "SELECT * FROM empleado  ORDER BY sueldo DESC ")
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		// traer NOMINA

		err := db.Get(&nomina, "SELECT * FROM nomina where numero=$1", Codigo)
		if err != nil {
			log.Fatalln(err)
		}
		err2 := db.Select(&det, NominaConsultaDetalle(), Codigo)
		if err2 != nil {
			fmt.Println(err2)
		}
	}

	parametros := map[string]interface{}{
		"codigo":   Codigo,
		"nomina":   nomina,
		"empleado": empleados,
		"detalle":  det,
		"hosting":  ruta,
	}
	//TERMINA TRAE COPIA DE EDITAR

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/nomina/nominaNuevo.html",
		"vista/nomina/autocompletaCentro.html",
		"vista/nomina/autocompletaempleado.html",
		"vista/nomina/modalColumna.html",
		"vista/nomina/modalHorasextras.html",
		"vista/nomina/nominaScript.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error nomina nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA INSERTAR COMPROBANTE DE NOMINA
//func InsertaDetalleComprobanteNomina(miFilaComprobante comprobantedetalle, miComprobante comprobante, miNomina nomina){
//	db := dbConn()
//// traer tercero
//
//
//	var q string
//	q = "insert into comprobantedetalle ("
//	q += "Fila,"
//	q += "Cuenta,"
//	q += "Tercero,"
//	q += "Centro,"
//	q += "Concepto,"
//	q += "Factura,"
//	q += "Debito,"
//	q += "Credito,"
//	q += "Documento,"
//	q += "Numero,"
//	q += "Fecha,"
//	q += "Fechaconsignacion"
//	q += " ) values("
//	q += parametros(12)
//	q += ")"
//	log.Println("Cadena SQL " + q)
//	insForm, err := db.Prepare(q)
//	if err != nil {
//		panic(err.Error())
//	}
//
//	if len(miFilaComprobante.Debito)>0 {
//		miFilaComprobante.Debito=	miFilaComprobante.Debito
//	}
//
//	if len(miFilaComprobante.Credito)>0 {
//		miFilaComprobante.Credito=	miFilaComprobante.Credito
//	}
//
//	// TERMINA COMPROBANTE GRABAR INSERTAR
//	_, err = insForm.Exec(
//		miFilaComprobante.Fila,
//	miFilaComprobante.Cuenta  ,
//	miNomina.Tercero,
//	miNomina.Centro,
//	miTercero.Nombre,
//	"",
//	Flotantedatabase(miFilaComprobante.Debito) ,
//	Flotantedatabase(miFilaComprobante.Credito) ,
//	miComprobante.Documento,
//	miComprobante.Numero,
//	miComprobante.Fecha,
//	miComprobante.Fechaconsignacion	)
//	if err != nil {
//	panic(err)
//	}
//}

//INICIA NOMINA INSERTAR AJAX
func NominaAgregar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var tempNomina nomina

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// carga informacion de la NOMINA
	err = json.Unmarshal(b, &tempNomina)
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

	Codigoactual = tempNomina.Numero

	if tempNomina.Accion == "Actualizar" {

		// borra detalle anterior
		delForm, err := db.Prepare("DELETE from nominadetalle WHERE numero=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(tempNomina.Numero)

		// borra detalle inventario
		Borrarinventario(Codigoactual, "Nomina")

		// borra cabecera anterior
		delForm1, err := db.Prepare("DELETE from nomina WHERE numero=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm1.Exec(tempNomina.Numero)
	}

	// INSERTA DETALLE
	for i, x := range tempNomina.Detalle {
		var a = i
		var q string
		q = "insert into nominadetalle ("
		q += "Numero,"
		q += "Prefijo,"
		q += "Fechainicial,"
		q += "Fechafinal,"
		q += "Codigo,"
		q += "Centro,"
		q += "Sueldo,"
		q += "Dias,"
		q += "Trabajado,"
		q += "Transporte,"

		q += "Cesantias,"
		q += "Intereses,"
		q += "Prima,"
		q += "Vacaciones,"
		q += "Viaticos,"
		q += "Horasextras,"
		q += "Incapacidades,"
		q += "Licencias,"
		q += "Bonificaciones,"
		q += "Auxilios,"

		q += "Huelgas,"
		q += "Conceptos,"
		q += "Compensaciones,"
		q += "Bonos,"
		q += "Comisiones,"
		q += "Dotaciones,"
		q += "Sostenimiento,"
		q += "Teletrabajo,"
		q += "Indemnizaciones,"
		q += "Devengado,"

		q += "Salud,"
		q += "Pension,"
		q += "Pensionrais,"
		q += "Pensionvoluntaria,"
		q += "Solidaridad,"
		q += "Subsistencia,"
		q += "Sindicatos,"
		q += "Sanciones,"
		q += "Libranzas,"
		q += "Terceros,"

		q += "Anticipos,"
		q += "Otras,"
		q += "Retencion,"
		q += "Afc,"
		q += "Embargos,"
		q += "Educacion,"
		q += "Deuda,"
		q += "Deducciones,"
		q += "Neto"
		q += " ) values("
		q += parametros(49)
		q += ")"
		log.Println("Cadena SQL " + q)
		insForm, err := db.Prepare(q)
		if err != nil {
			panic(err.Error())
		}

		// TERMINA NOMINA GRABAR INSERTAR
		_, err = insForm.Exec(
			x.Numero,
			Mayuscula(x.Prefijo),
			x.Fechainicial,
			x.Fechafinal,
			Quitacoma(x.Codigo),
			x.Centro,
			Quitacoma(x.Sueldo),
			x.Dias,
			Quitacoma(x.Trabajado),
			Quitacoma(x.Transporte),
			Quitacoma(x.Cesantias),
			Quitacoma(x.Intereses),
			Quitacoma(x.Prima),
			Quitacoma(x.Vacaciones),
			Quitacoma(x.Viaticos),
			Quitacoma(x.Horasextras),
			Quitacoma(x.Incapacidades),
			Quitacoma(x.Licencias),
			Quitacoma(x.Bonificaciones),
			Quitacoma(x.Auxilios),
			Quitacoma(x.Huelgas),
			Quitacoma(x.Conceptos),
			Quitacoma(x.Compensaciones),
			Quitacoma(x.Bonos),
			Quitacoma(x.Comisiones),
			Quitacoma(x.Dotaciones),
			Quitacoma(x.Sostenimiento),
			Quitacoma(x.Teletrabajo),
			Quitacoma(x.Indemnizaciones),
			Quitacoma(x.Devengado),
			Quitacoma(x.Salud),
			Quitacoma(x.Pension),
			Quitacoma(x.Pensionrais),
			Quitacoma(x.Pensionvoluntaria),
			Quitacoma(x.Solidaridad),
			Quitacoma(x.Subsistencia),
			Quitacoma(x.Sindicatos),
			Quitacoma(x.Sanciones),
			Quitacoma(x.Libranzas),
			Quitacoma(x.Terceros),
			Quitacoma(x.Anticipos),
			Quitacoma(x.Otras),
			Quitacoma(x.Retencion),
			Quitacoma(x.Afc),
			Quitacoma(x.Embargos),
			Quitacoma(x.Educacion),
			Quitacoma(x.Deuda),
			Quitacoma(x.Deducciones),
			Quitacoma(x.Neto))

		if err != nil {
			panic(err)
		}
		log.Println("Insertar Codigo \n", x.Codigo, a)
	}

	// INICIA INSERTAR PLANILLAS
	var q string
	q = "insert into nomina ("

	q += "Numero,"
	q += "Prefijo,"
	q += "Fechainicial,"
	q += "Fechafinal,"
	q += "Trabajado,"
	q += "Transporte,"
	q += "Cesantias,"
	q += "Intereses,"
	q += "Prima,"
	q += "Vacaciones,"

	q += "Viaticos,"
	q += "Horasextras,"
	q += "Incapacidades,"
	q += "Licencias,"
	q += "Bonificaciones,"
	q += "Auxilios,"
	q += "Huelgas,"
	q += "Conceptos,"
	q += "Compensaciones,"
	q += "Bonos,"

	q += "Comisiones,"
	q += "Dotaciones,"
	q += "Sostenimiento,"
	q += "Teletrabajo,"
	q += "Indemnizaciones,"
	q += "Devengado,"
	q += "Salud,"
	q += "Pension,"
	q += "Pensionrais,"
	q += "Pensionvoluntaria,"

	q += "Solidaridad,"
	q += "Subsistencia,"
	q += "Sindicatos,"
	q += "Sanciones,"
	q += "Libranzas,"
	q += "Terceros,"
	q += "Anticipos,"
	q += "Otras,"
	q += "Retencion,"
	q += "Afc,"

	q += "Embargos,"
	q += "Educacion,"
	q += "Deuda,"
	q += "Deducciones,"
	q += "Neto"

	q += " ) values("
	q += parametros(45)
	q += " ) "

	log.Println("Cadena SQL " + q)
	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	layout := "2006-01-02"

	_, err = insForm.Exec(
		tempNomina.Numero,
		Mayuscula(tempNomina.Prefijo),
		tempNomina.Fechainicial.Format(layout),
		tempNomina.Fechafinal.Format(layout),

		Quitacoma(tempNomina.Trabajado),
		Quitacoma(tempNomina.Transporte),
		Quitacoma(tempNomina.Cesantias),
		Quitacoma(tempNomina.Intereses),
		Quitacoma(tempNomina.Prima),
		Quitacoma(tempNomina.Vacaciones),
		Quitacoma(tempNomina.Viaticos),
		Quitacoma(tempNomina.Horasextras),
		Quitacoma(tempNomina.Incapacidades),
		Quitacoma(tempNomina.Licencias),
		Quitacoma(tempNomina.Bonificaciones),
		Quitacoma(tempNomina.Auxilios),
		Quitacoma(tempNomina.Huelgas),
		Quitacoma(tempNomina.Conceptos),
		Quitacoma(tempNomina.Compensaciones),
		Quitacoma(tempNomina.Bonos),
		Quitacoma(tempNomina.Comisiones),
		Quitacoma(tempNomina.Dotaciones),
		Quitacoma(tempNomina.Sostenimiento),
		Quitacoma(tempNomina.Teletrabajo),
		Quitacoma(tempNomina.Indemnizaciones),
		Quitacoma(tempNomina.Devengado),
		Quitacoma(tempNomina.Salud),
		Quitacoma(tempNomina.Pension),
		Quitacoma(tempNomina.Pensionrais),
		Quitacoma(tempNomina.Pensionvoluntaria),
		Quitacoma(tempNomina.Solidaridad),
		Quitacoma(tempNomina.Subsistencia),
		Quitacoma(tempNomina.Sindicatos),
		Quitacoma(tempNomina.Sanciones),
		Quitacoma(tempNomina.Libranzas),
		Quitacoma(tempNomina.Terceros),
		Quitacoma(tempNomina.Anticipos),
		Quitacoma(tempNomina.Otras),
		Quitacoma(tempNomina.Retencion),
		Quitacoma(tempNomina.Afc),
		Quitacoma(tempNomina.Embargos),
		Quitacoma(tempNomina.Educacion),
		Quitacoma(tempNomina.Deuda),
		Quitacoma(tempNomina.Deducciones),
		Quitacoma(tempNomina.Neto))

	if err != nil {
		panic(err)
	}

	// INSERTAR COMPROBANTE CONTABILIDAD
	var tempComprobante comprobante
	var tempComprobanteDetalle comprobantedetalle
	tempComprobante.Documento = "6"
	tempComprobante.Numero = tempNomina.Numero
	tempComprobante.Fecha = tempNomina.Fechafinal
	tempComprobante.Fechaconsignacion = tempNomina.Fechafinal
	tempComprobante.Debito = tempNomina.Neto
	tempComprobante.Credito = tempNomina.Neto
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

	for i, x := range tempNomina.Detalle {

		// TRABAJADO
		if Flotante(x.Trabajado) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Trabajado)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Trabajado
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Trabajado))
			tempComprobanteDetalle.Credito = ""
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// TRANSPORTE
		if Flotante(x.Transporte) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Transporte)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Transporte
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Transporte))
			tempComprobanteDetalle.Credito = ""
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// CESANTIAS
		if Flotante(x.Cesantias) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Cesantias)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Cesantias
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Cesantias))
			tempComprobanteDetalle.Credito = ""
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
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
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
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
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
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
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// VIATICOS
		if Flotante(x.Viaticos) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Viaticos)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Viaticos
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Viaticos))
			tempComprobanteDetalle.Credito = ""
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// HORAS EXTRAS
		if Flotante(x.Horasextras) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Horasextras)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Horasextras
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Horasextras))
			tempComprobanteDetalle.Credito = ""
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// INCAPACIDADES
		if Flotante(x.Incapacidades) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Incapacidades)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Incapacidades
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Incapacidades))
			tempComprobanteDetalle.Credito = ""
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// LICENCIAS
		if Flotante(x.Licencias) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Licencias)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Licencias
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Licencias))
			tempComprobanteDetalle.Credito = ""
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// BONIFICACIONES
		if Flotante(x.Bonificaciones) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Bonificaciones)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Bonificaciones
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Bonificaciones))
			tempComprobanteDetalle.Credito = ""
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// AUXILIOS
		if Flotante(x.Auxilios) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Auxilios)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Auxilios
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Auxilios))
			tempComprobanteDetalle.Credito = ""
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// HUELGAS
		if Flotante(x.Huelgas) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Huelgas)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Huelgas
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Huelgas))
			tempComprobanteDetalle.Credito = ""
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// CONCEPTOS
		if Flotante(x.Conceptos) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Conceptos)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Conceptos
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Conceptos))
			tempComprobanteDetalle.Credito = ""
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// COMPENSACIONES
		if Flotante(x.Compensaciones) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Compensaciones)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Compensaciones
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Compensaciones))
			tempComprobanteDetalle.Credito = ""
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// BONOS
		if Flotante(x.Bonos) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Bonos)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Bonos
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Bonos))
			tempComprobanteDetalle.Credito = ""
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// COMISIONES
		if Flotante(x.Comisiones) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Comisiones)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Comisiones
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Comisiones))
			tempComprobanteDetalle.Credito = ""
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
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
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// SOSTENIMIENTO
		if Flotante(x.Sostenimiento) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Sostenimiento)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Sostenimiento
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Sostenimiento))
			tempComprobanteDetalle.Credito = ""
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// TELETRABAJO
		if Flotante(x.Teletrabajo) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Teletrabajo)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Teletrabajo
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Teletrabajo))
			tempComprobanteDetalle.Credito = ""
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// INDEMNIZACIONES
		if Flotante(x.Indemnizaciones) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Indemnizaciones)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Indemnizaciones
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Indemnizaciones))
			tempComprobanteDetalle.Credito = ""
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// SALUD
		if Flotante(x.Salud) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Salud)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Salud
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Salud))
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}

		// PENSION
		if Flotante(x.Pension) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Pension)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Pension
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Pension))
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}

		// PENSION RAIS
		if Flotante(x.Pensionrais) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Pensionrais)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Pensionrais
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Pensionrais))
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}

		// PENSION VOLUNTARIA
		if Flotante(x.Pensionvoluntaria) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Pensionvoluntaria)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Pensionvoluntaria
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Pensionvoluntaria))
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}

		// SOLIDARIDAD
		if Flotante(x.Solidaridad) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Solidaridad)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Solidaridad
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Solidaridad))
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}

		// SUBSISTENCIA
		if Flotante(x.Subsistencia) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Subsistencia)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Subsistencia
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Subsistencia))
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}

		// SINDICATOS
		if Flotante(x.Sindicatos) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Sindicatos)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Sindicatos
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Sindicatos))
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}

		// SANCIONES
		if Flotante(x.Sanciones) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Sanciones)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Sanciones
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Sanciones))
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}

		// LIBRANZAS
		if Flotante(x.Libranzas) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Libranzas)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Libranzas
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Libranzas))
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}

		// TERCEROS
		if Flotante(x.Terceros) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Terceros)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Terceros
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Terceros))
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}

		// ANTICIPOS
		if Flotante(x.Anticipos) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Anticipos)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Anticipos
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Anticipos))
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}

		// OTRAS
		if Flotante(x.Otras) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Otras)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Otras
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Otras))
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}

		// RETENCION
		if Flotante(x.Retencion) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Retencion)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Retencion
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Retencion))
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}

		// AFC
		if Flotante(x.Afc) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Afc)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Afc
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Afc))
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}

		// EMBARGOS
		if Flotante(x.Embargos) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Embargos)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Embargos
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Embargos))
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}

		// EDUCACION
		if Flotante(x.Educacion) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Educacion)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Educacion
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Educacion))
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}

		// DEUDA
		if Flotante(x.Deuda) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Deuda)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Deuda
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Deuda))
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}

		// NETO
		if Flotante(x.Neto) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Neto)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Neto
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Neto))
			InsertaDetalleComprobanteNomina(tempComprobanteDetalle, tempComprobante, x)
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
func InsertaDetalleComprobanteNomina(miFilaComprobante comprobantedetalle, miComprobante comprobante, miCompra nominadetalle) {
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
func NominaExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM nomina  WHERE codigo=$1", Codigo)
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
func NominaEditar(w http.ResponseWriter, r *http.Request) {
	Numero := mux.Vars(r)["codigo"]
	log.Println("inicio nomina editar" + Numero)
	db := dbConn()

	// traer NOMINA
	v := nomina{}
	err := db.Get(&v, "SELECT * FROM nomina where numero=$1", Numero)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []nominadetalleeditar{}

	err2 := db.Select(&det, NominaConsultaDetalle(), Numero)
	if err2 != nil {
		fmt.Println(err2)
	}
	parametros := map[string]interface{}{
		"nomina":  v,
		"detalle": det,
		"hosting": ruta,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/nomina/nominaEditar.html",
		"vista/nomina/autocompletaCentro.html",
		"vista/nomina/autocompletaempleado.html",
		"vista/nomina/modalHorasextras.html",
		"vista/nomina/modalColumna.html",
		"vista/nomina/nominaScript.html")

	fmt.Printf("%v, %v", miTemplate, err)

	miTemplate.Execute(w, parametros)
}

// INICIA NOMINA BORRAR
func NominaBorrar(w http.ResponseWriter, r *http.Request) {
	Numero := mux.Vars(r)["codigo"]
	log.Println("inicio nomina editar" + Numero)
	db := dbConn()

	// traer NOMINA
	v := nomina{}
	err := db.Get(&v, "SELECT * FROM nomina where numero=$1", Numero)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []nominadetalleeditar{}

	err2 := db.Select(&det, NominaConsultaDetalle(), Numero)
	if err2 != nil {
		fmt.Println(err2)
	}
	parametros := map[string]interface{}{
		"codigo":  Numero,
		"nomina":  v,
		"detalle": det,
		"hosting": ruta,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/nomina/nominaBorrar.html",
		"vista/nomina/autocompletaCentro.html",
		"vista/nomina/autocompletaempleado.html",
		"vista/nomina/modalColumna.html",
		"vista/nomina/nominaScript.html")

	fmt.Printf("%v, %v", miTemplate, err)

	miTemplate.Execute(w, parametros)
}

// INICIA NOMINA ELIMINAR
func NominaEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	codigo := mux.Vars(r)["codigo"]

	// borrar NOMINA
	delForm, err := db.Prepare("DELETE from nomina WHERE numero=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(codigo)

	// borar detalle
	delForm1, err := db.Prepare("DELETE from nominadetalle WHERE numero=$1")
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
	http.Redirect(w, r, "/NominaLista", 301)
}

// INICIA NOMINA PDF
var numerofila = 0
var numeroitem int

var cordenadaitem float64

func NominaPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Numero := mux.Vars(r)["numero"]

	// traer NOMINA
	miNomina := nomina{}
	err := db.Get(&miNomina, "SELECT * FROM nomina where numero=$1", Numero)
	if err != nil {
		log.Fatalln(err)
	}

	// traer detalle
	miDetalle := []nominadetalleeditar{}
	err2 := db.Select(&miDetalle, NominaConsultaDetalle(), Numero)
	if err2 != nil {
		fmt.Println(err2)
	}

	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	var buf bytes.Buffer
	var err1 error
	pdf := gofpdf.New("P", "mm", "letter", cnFontDir)
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
		pdf.Ln(6)
	})

	pdf.SetFooterFunc(func() {
		pdf.SetFont("Arial", "", 8)
		pdf.SetY(128)
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
	numeroitem = 1

	var numerodevengado int
	var numerodeduccion int

	for a, miFila := range miDetalle {
		fmt.Println(a)
		var consecutivo string
		consecutivo = strconv.Itoa(a + 1)

		pdf.AddPage()
		pdf.SetY(30)
		pdf.SetX(150)
		pdf.SetFont("Arial", "", 11)
		pdf.CellFormat(40, 10, "SOPORTE DE PAGO No.", "0", 0, "C", false, 0, "")
		pdf.SetY(35)
		pdf.SetX(150)
		pdf.CellFormat(40, 10, miFila.Numero+" - "+miFila.Prefijo+" - "+consecutivo, "0", 0, "C",
			false, 0, "")
		pdf.Ln(1)
		NominaCabecera(pdf, miFila)
		numerodevengado = 0
		numerodeduccion = 0
		cordenadaitem = pdf.GetY()
		cordenada1 = pdf.GetY()
		cordenada2 = pdf.GetY()

		// lista devengado
		if Flotante(miFila.Trabajado) != 0 {
			pdf.SetY(cordenada1)
			pdf.SetX(28)
			pdf.CellFormat(40, 4, ene("Sueldo"), "", 0,
				"L", false, 0, "")
			pdf.SetX(50)
			pdf.CellFormat(40, 4, Coma(miFila.Trabajado), "", 0,
				"R", false, 0, "")
			cordenada1 = cordenada1 + 4
			numerodevengado = numerodevengado + 1
		}

		if Flotante(miFila.Transporte) != 0 {
			pdf.SetY(cordenada1)
			pdf.SetX(28)
			pdf.CellFormat(40, 4, ene("Auxilio de Transporte"), "", 0,
				"L", false, 0, "")
			pdf.SetX(50)
			pdf.CellFormat(40, 4, Coma(miFila.Transporte), "", 0,
				"R", false, 0, "")
			cordenada1 = cordenada1 + 4
			numerodevengado = numerodevengado + 1
		}

		if Flotante(miFila.Cesantias) != 0 {
			pdf.SetY(cordenada1)
			pdf.SetX(28)
			pdf.CellFormat(40, 4, ene("Cesantias"), "", 0,
				"L", false, 0, "")
			pdf.SetX(50)
			pdf.CellFormat(40, 4, Coma(miFila.Cesantias), "", 0,
				"R", false, 0, "")
			cordenada1 = cordenada1 + 4
			numerodevengado = numerodevengado + 1
		}

		if Flotante(miFila.Intereses) != 0 {
			pdf.SetY(cordenada1)
			pdf.SetX(28)
			pdf.CellFormat(40, 4, ene("Intereses Cesantias"), "", 0,
				"L", false, 0, "")
			pdf.SetX(50)
			pdf.CellFormat(40, 4, Coma(miFila.Intereses), "", 0,
				"R", false, 0, "")
			cordenada1 = cordenada1 + 4
			numerodevengado = numerodevengado + 1
		}

		if Flotante(miFila.Prima) != 0 {
			pdf.SetY(cordenada1)
			pdf.SetX(28)
			pdf.CellFormat(40, 4, ene("Prima de Servicios"), "", 0,
				"L", false, 0, "")
			pdf.SetX(50)
			pdf.CellFormat(40, 4, Coma(miFila.Prima), "", 0,
				"R", false, 0, "")
			cordenada1 = cordenada1 + 4
			numerodevengado = numerodevengado + 1
		}

		if Flotante(miFila.Vacaciones) != 0 {
			pdf.SetY(cordenada1)
			pdf.SetX(28)
			pdf.CellFormat(40, 4, ene("Vacaciones"), "", 0,
				"L", false, 0, "")
			pdf.SetX(50)
			pdf.CellFormat(40, 4, Coma(miFila.Vacaciones), "", 0,
				"R", false, 0, "")
			cordenada1 = cordenada1 + 4
			numerodevengado = numerodevengado + 1
		}

		if Flotante(miFila.Viaticos) != 0 {
			pdf.SetY(cordenada1)
			pdf.SetX(28)
			pdf.CellFormat(40, 4, ene("Viaticos"), "", 0,
				"L", false, 0, "")
			pdf.SetX(50)
			pdf.CellFormat(40, 4, Coma(miFila.Viaticos), "", 0,
				"R", false, 0, "")
			cordenada1 = cordenada1 + 4
			numerodevengado = numerodevengado + 1
		}

		if Flotante(miFila.Horasextras) != 0 {
			pdf.SetY(cordenada1)
			pdf.SetX(28)
			pdf.CellFormat(40, 4, ene("Horas Extras y Festivos"), "", 0,
				"L", false, 0, "")
			pdf.SetX(50)
			pdf.CellFormat(40, 4, Coma(miFila.Horasextras), "", 0,
				"R", false, 0, "")
			cordenada1 = cordenada1 + 4
			numerodevengado = numerodevengado + 1
		}

		if Flotante(miFila.Incapacidades) != 0 {
			pdf.SetY(cordenada1)
			pdf.SetX(28)
			pdf.CellFormat(40, 4, ene("Incapacidades"), "", 0,
				"L", false, 0, "")
			pdf.SetX(50)
			pdf.CellFormat(40, 4, Coma(miFila.Incapacidades), "", 0,
				"R", false, 0, "")
			cordenada1 = cordenada1 + 4
			numerodevengado = numerodevengado + 1
		}

		if Flotante(miFila.Licencias) != 0 {
			pdf.SetY(cordenada1)
			pdf.SetX(28)
			pdf.CellFormat(40, 4, ene("Licencias"), "", 0,
				"L", false, 0, "")
			pdf.SetX(50)
			pdf.CellFormat(40, 4, Coma(miFila.Licencias), "", 0,
				"R", false, 0, "")
			cordenada1 = cordenada1 + 4
			numerodevengado = numerodevengado + 1
		}

		if Flotante(miFila.Bonificaciones) != 0 {
			pdf.SetY(cordenada1)
			pdf.SetX(28)
			pdf.CellFormat(40, 4, ene("Bonificaciones"), "", 0,
				"L", false, 0, "")
			pdf.SetX(50)
			pdf.CellFormat(40, 4, Coma(miFila.Bonificaciones), "", 0,
				"R", false, 0, "")
			cordenada1 = cordenada1 + 4
			numerodevengado = numerodevengado + 1
		}

		if Flotante(miFila.Auxilios) != 0 {
			pdf.SetY(cordenada1)
			pdf.SetX(28)
			pdf.CellFormat(40, 4, ene("Auxilios"), "", 0,
				"L", false, 0, "")
			pdf.SetX(50)
			pdf.CellFormat(40, 4, Coma(miFila.Auxilios), "", 0,
				"R", false, 0, "")
			cordenada1 = cordenada1 + 4
			numerodevengado = numerodevengado + 1
		}

		if Flotante(miFila.Huelgas) != 0 {
			pdf.SetY(cordenada1)
			pdf.SetX(28)
			pdf.CellFormat(40, 4, ene("Huelgas"), "", 0,
				"L", false, 0, "")
			pdf.SetX(50)
			pdf.CellFormat(40, 4, Coma(miFila.Huelgas), "", 0,
				"R", false, 0, "")
			cordenada1 = cordenada1 + 4
			numerodevengado = numerodevengado + 1
		}

		if Flotante(miFila.Conceptos) != 0 {
			pdf.SetY(cordenada1)
			pdf.SetX(28)
			pdf.CellFormat(40, 4, ene("Conceptos"), "", 0,
				"L", false, 0, "")
			pdf.SetX(50)
			pdf.CellFormat(40, 4, Coma(miFila.Conceptos), "", 0,
				"R", false, 0, "")
			cordenada1 = cordenada1 + 4
			numerodevengado = numerodevengado + 1
		}

		if Flotante(miFila.Compensaciones) != 0 {
			pdf.SetY(cordenada1)
			pdf.SetX(28)
			pdf.CellFormat(40, 4, ene("Compensaciones"), "", 0,
				"L", false, 0, "")
			pdf.SetX(50)
			pdf.CellFormat(40, 4, Coma(miFila.Compensaciones), "", 0,
				"R", false, 0, "")
			cordenada1 = cordenada1 + 4
			numerodevengado = numerodevengado + 1
		}

		if Flotante(miFila.Bonos) != 0 {
			pdf.SetY(cordenada1)
			pdf.SetX(28)
			pdf.CellFormat(40, 4, ene("Bonos"), "", 0,
				"L", false, 0, "")
			pdf.SetX(50)
			pdf.CellFormat(40, 4, Coma(miFila.Bonos), "", 0,
				"R", false, 0, "")
			cordenada1 = cordenada1 + 4
			numerodevengado = numerodevengado + 1
		}

		if Flotante(miFila.Comisiones) != 0 {
			pdf.SetY(cordenada1)
			pdf.SetX(28)
			pdf.CellFormat(40, 4, ene("Comisiones"), "", 0,
				"L", false, 0, "")
			pdf.SetX(50)
			pdf.CellFormat(40, 4, Coma(miFila.Comisiones), "", 0,
				"R", false, 0, "")
			cordenada1 = cordenada1 + 4
			numerodevengado = numerodevengado + 1
		}

		if Flotante(miFila.Dotaciones) != 0 {
			pdf.SetY(cordenada1)
			pdf.SetX(28)
			pdf.CellFormat(40, 4, ene("Calzado y Vestido de Labor"), "", 0,
				"L", false, 0, "")
			pdf.SetX(50)
			pdf.CellFormat(40, 4, Coma(miFila.Dotaciones), "", 0,
				"R", false, 0, "")
			cordenada1 = cordenada1 + 4
			numerodevengado = numerodevengado + 1
		}

		if Flotante(miFila.Sostenimiento) != 0 {
			pdf.SetY(cordenada1)
			pdf.SetX(28)
			pdf.CellFormat(40, 4, ene("Sostenimiento"), "", 0,
				"L", false, 0, "")
			pdf.SetX(50)
			pdf.CellFormat(40, 4, Coma(miFila.Sostenimiento), "", 0,
				"R", false, 0, "")
			cordenada1 = cordenada1 + 4
			numerodevengado = numerodevengado + 1
		}

		if Flotante(miFila.Teletrabajo) != 0 {
			pdf.SetY(cordenada1)
			pdf.SetX(28)
			pdf.CellFormat(40, 4, ene("Teletrabajo"), "", 0,
				"L", false, 0, "")
			pdf.SetX(50)
			pdf.CellFormat(40, 4, Coma(miFila.Teletrabajo), "", 0,
				"R", false, 0, "")
			cordenada1 = cordenada1 + 4
			numerodevengado = numerodevengado + 1
		}

		if Flotante(miFila.Indemnizaciones) != 0 {
			pdf.SetY(cordenada1)
			pdf.SetX(28)
			pdf.CellFormat(40, 4, ene("Indemnizaciones"), "", 0,
				"L", false, 0, "")
			pdf.SetX(50)
			pdf.CellFormat(40, 4, Coma(miFila.Indemnizaciones), "", 0,
				"R", false, 0, "")
			cordenada1 = cordenada1 + 4
			numerodevengado = numerodevengado + 1
		}

		// lista deduccion
		if Flotante(miFila.Salud) != 0 {
			pdf.SetY(cordenada2)
			pdf.SetX(93)
			pdf.CellFormat(40, 4, ene("Aportes Salud"), "", 0,
				"L", false, 0, "")
			pdf.SetX(110)
			pdf.CellFormat(40, 4, Coma(miFila.Salud), "", 0,
				"R", false, 0, "")
			pdf.Ln(-1)
			cordenada2 = cordenada2 + 4
			numerodeduccion = numerodeduccion + 1
		}

		if Flotante(miFila.Pension) != 0 {
			pdf.SetY(cordenada2)
			pdf.SetX(93)
			pdf.CellFormat(40, 4, ene("Aportes Pension"), "", 0,
				"L", false, 0, "")
			pdf.SetX(110)
			pdf.CellFormat(40, 4, Coma(miFila.Pension), "", 0,
				"R", false, 0, "")
			pdf.Ln(-1)
			cordenada2 = cordenada2 + 4
			numerodeduccion = numerodeduccion + 1
		}

		if Flotante(miFila.Pensionrais) != 0 {
			pdf.SetY(cordenada2)
			pdf.SetX(93)
			pdf.CellFormat(40, 4, ene("Aportes Pension Rais"), "", 0,
				"L", false, 0, "")
			pdf.SetX(110)
			pdf.CellFormat(40, 4, Coma(miFila.Pensionrais), "", 0,
				"R", false, 0, "")
			pdf.Ln(-1)
			cordenada2 = cordenada2 + 4
			numerodeduccion = numerodeduccion + 1
		}

		if Flotante(miFila.Pensionvoluntaria) != 0 {
			pdf.SetY(cordenada2)
			pdf.SetX(93)
			pdf.CellFormat(40, 4, ene("Aportes Pension Voluntaria"), "", 0,
				"L", false, 0, "")
			pdf.SetX(110)
			pdf.CellFormat(40, 4, Coma(miFila.Pensionvoluntaria), "", 0,
				"R", false, 0, "")
			pdf.Ln(-1)
			cordenada2 = cordenada2 + 4
			numerodeduccion = numerodeduccion + 1
		}

		if Flotante(miFila.Solidaridad) != 0 {
			pdf.SetY(cordenada2)
			pdf.SetX(93)
			pdf.CellFormat(40, 4, ene("Fondo Solidaridad"), "", 0,
				"L", false, 0, "")
			pdf.SetX(110)
			pdf.CellFormat(40, 4, Coma(miFila.Solidaridad), "", 0,
				"R", false, 0, "")
			pdf.Ln(-1)
			cordenada2 = cordenada2 + 4
			numerodeduccion = numerodeduccion + 1
		}

		if Flotante(miFila.Subsistencia) != 0 {
			pdf.SetY(cordenada2)
			pdf.SetX(93)
			pdf.CellFormat(40, 4, ene("Subsistencia"), "", 0,
				"L", false, 0, "")
			pdf.SetX(110)
			pdf.CellFormat(40, 4, Coma(miFila.Subsistencia), "", 0,
				"R", false, 0, "")
			pdf.Ln(-1)
			cordenada2 = cordenada2 + 4
			numerodeduccion = numerodeduccion + 1
		}

		if Flotante(miFila.Sindicatos) != 0 {
			pdf.SetY(cordenada2)
			pdf.SetX(93)
			pdf.CellFormat(40, 4, ene("Cuota Sindicato"), "", 0,
				"L", false, 0, "")
			pdf.SetX(110)
			pdf.CellFormat(40, 4, Coma(miFila.Sindicatos), "", 0,
				"R", false, 0, "")
			pdf.Ln(-1)
			cordenada2 = cordenada2 + 4
			numerodeduccion = numerodeduccion + 1
		}

		if Flotante(miFila.Sanciones) != 0 {
			pdf.SetY(cordenada2)
			pdf.SetX(93)
			pdf.CellFormat(40, 4, ene("Sanciones"), "", 0,
				"L", false, 0, "")
			pdf.SetX(110)
			pdf.CellFormat(40, 4, Coma(miFila.Sanciones), "", 0,
				"R", false, 0, "")
			pdf.Ln(-1)
			cordenada2 = cordenada2 + 4
			numerodeduccion = numerodeduccion + 1
		}

		if Flotante(miFila.Libranzas) != 0 {
			pdf.SetY(cordenada2)
			pdf.SetX(93)
			pdf.CellFormat(40, 4, ene("Libranzas"), "", 0,
				"L", false, 0, "")
			pdf.SetX(110)
			pdf.CellFormat(40, 4, Coma(miFila.Libranzas), "", 0,
				"R", false, 0, "")
			pdf.Ln(-1)
			cordenada2 = cordenada2 + 4
			numerodeduccion = numerodeduccion + 1
		}

		if Flotante(miFila.Terceros) != 0 {
			pdf.SetY(cordenada2)
			pdf.SetX(93)
			pdf.CellFormat(40, 4, ene("Pagos a Terceros"), "", 0,
				"L", false, 0, "")
			pdf.SetX(110)
			pdf.CellFormat(40, 4, Coma(miFila.Terceros), "", 0,
				"R", false, 0, "")
			pdf.Ln(-1)
			cordenada2 = cordenada2 + 4
			numerodeduccion = numerodeduccion + 1
		}

		if Flotante(miFila.Anticipos) != 0 {
			pdf.SetY(cordenada2)
			pdf.SetX(93)
			pdf.CellFormat(40, 4, ene("Anticipos Sueldos"), "", 0,
				"L", false, 0, "")
			pdf.SetX(110)
			pdf.CellFormat(40, 4, Coma(miFila.Anticipos), "", 0,
				"R", false, 0, "")
			pdf.Ln(-1)
			cordenada2 = cordenada2 + 4
			numerodeduccion = numerodeduccion + 1
		}

		if Flotante(miFila.Otras) != 0 {
			pdf.SetY(cordenada2)
			pdf.SetX(93)
			pdf.CellFormat(40, 4, ene("Otras Deducciones"), "", 0,
				"L", false, 0, "")
			pdf.SetX(110)
			pdf.CellFormat(40, 4, Coma(miFila.Otras), "", 0,
				"R", false, 0, "")
			pdf.Ln(-1)
			cordenada2 = cordenada2 + 4
			numerodeduccion = numerodeduccion + 1
		}

		if Flotante(miFila.Retencion) != 0 {
			pdf.SetY(cordenada2)
			pdf.SetX(93)
			pdf.CellFormat(40, 4, ene("Retencion en la Fuente"), "", 0,
				"L", false, 0, "")
			pdf.SetX(110)
			pdf.CellFormat(40, 4, Coma(miFila.Retencion), "", 0,
				"R", false, 0, "")
			pdf.Ln(-1)
			cordenada2 = cordenada2 + 4
			numerodeduccion = numerodeduccion + 1
		}

		if Flotante(miFila.Afc) != 0 {
			pdf.SetY(cordenada2)
			pdf.SetX(93)
			pdf.CellFormat(40, 4, ene("Cuenta Afc"), "", 0,
				"L", false, 0, "")
			pdf.SetX(110)
			pdf.CellFormat(40, 4, Coma(miFila.Afc), "", 0,
				"R", false, 0, "")
			pdf.Ln(-1)
			cordenada2 = cordenada2 + 4
			numerodeduccion = numerodeduccion + 1
		}

		if Flotante(miFila.Embargos) != 0 {
			pdf.SetY(cordenada2)
			pdf.SetX(93)
			pdf.CellFormat(40, 4, ene("Embargos Judiciales"), "", 0,
				"L", false, 0, "")
			pdf.SetX(110)
			pdf.CellFormat(40, 4, Coma(miFila.Embargos), "", 0,
				"R", false, 0, "")
			pdf.Ln(-1)
			cordenada2 = cordenada2 + 4
			numerodeduccion = numerodeduccion + 1
		}

		if Flotante(miFila.Educacion) != 0 {
			pdf.SetY(cordenada2)
			pdf.SetX(93)
			pdf.CellFormat(40, 4, ene("Educacion"), "", 0,
				"L", false, 0, "")
			pdf.SetX(110)
			pdf.CellFormat(40, 4, Coma(miFila.Educacion), "", 0,
				"R", false, 0, "")
			pdf.Ln(-1)
			cordenada2 = cordenada2 + 4
			numerodeduccion = numerodeduccion + 1
		}

		if Flotante(miFila.Deuda) != 0 {
			pdf.SetY(cordenada2)
			pdf.SetX(93)
			pdf.CellFormat(40, 4, ene("Deuda"), "", 0,
				"L", false, 0, "")
			pdf.SetX(110)
			pdf.CellFormat(40, 4, Coma(miFila.Deuda), "", 0,

				"R", false, 0, "")
			pdf.Ln(-1)
			cordenada2 = cordenada2 + 4
			numerodeduccion = numerodeduccion + 1
		}

		// imprime item
		n := 1
		for n <= numerodevengado {
			pdf.SetY(cordenadaitem)

			pdf.SetFont("Arial", "", 9)
			pdf.SetTextColor(0, 0, 0)
			pdf.SetX(21)
			pdf.CellFormat(183, 4, strconv.Itoa(n), "", 0,
				"L", false, 0, "")
			cordenadaitem = cordenadaitem + 4
			n++
		}

		if numerodevengado >= numerodeduccion {
		} else {
			numerodevengado = numerodeduccion
		}

		if cordenada1 >= cordenada2 {
		} else {
			cordenada1 = cordenada2
		}

		//totales
		if Flotante(miFila.Devengado) != 0 {
			pdf.SetY(cordenada1)
			pdf.SetX(28)
			pdf.CellFormat(40, 4, ene("TOTAL DEVENGADO"), "", 0,
				"L", false, 0, "")
			pdf.SetX(50)
			pdf.CellFormat(40, 4, Coma(miFila.Devengado), "", 0,
				"R", false, 0, "")

		}

		if Flotante(miFila.Deducciones) != 0 {
			pdf.SetY(cordenada1)
			pdf.SetX(93)
			pdf.CellFormat(40, 4, ene("TOTAL DEDUCCIONES"), "", 0,
				"L", false, 0, "")
			pdf.SetX(110)
			pdf.CellFormat(40, 4, Coma(miFila.Deducciones), "", 0,
				"R", false, 0, "")
			pdf.Ln(-1)

		}

		cordenada1 = cordenada1 + 4

		if Flotante(miFila.Neto) != 0 {
			pdf.SetY(cordenada1)
			pdf.SetX(93)
			pdf.CellFormat(40, 4, ene("NETO A PAGAR"), "", 0,
				"L", false, 0, "")
			pdf.SetX(110)
			pdf.CellFormat(40, 4, Coma(miFila.Neto), "", 0,
				"R", false, 0, "")
			pdf.Ln(-1)
			cordenada2 = cordenada2 + 4
		}

		pdf.SetY(116)
		pdf.SetX(152)
		pdf.Line(152, 116, 204, 116)

		pdf.SetY(117)
		pdf.SetX(152)
		pdf.SetTextColor(0, 0, 0)
		pdf.CellFormat(40, 4, "C. C. No.:", "", 0,
			"L", false, 0, "")
		pdf.SetY(121)
		pdf.SetX(152)
		pdf.CellFormat(40, 4, "Recibo Conforme", "", 0,
			"L", false, 0, "")
	}

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

func NominaCabecera(pdf *gofpdf.Fpdf, miNominadetalle nominadetalleeditar) {
	ene := pdf.UnicodeTranslatorFromDescriptor("")
	// RELLENO TITULO
	pdf.SetFont("Arial", "", 10)
	pdf.SetY(50)
	pdf.SetX(20)
	pdf.SetFillColor(224, 231, 239)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(184, 6, "DATOS EMPLEADO NOMINA", "0", 0,
		"C", true, 0, "")
	pdf.Ln(8)
	pdf.SetX(20)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(40, 4, "Nomina Numero:", "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, miNominadetalle.Numero, "", 0,
		"L", false, 0, "")
	pdf.SetX(75)
	pdf.CellFormat(40, 4, "Prefijo:", "", 0,
		"L", false, 0, "")
	pdf.SetX(93)
	pdf.CellFormat(40, 4, miNominadetalle.Prefijo, "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Fecha Inicial:", "", 0,
		"L", false, 0, "")
	pdf.SetX(135)
	pdf.CellFormat(40, 4, miNominadetalle.Fechainicial.Format("02/01/2006"), "", 0,
		"L", false, 0, "")
	pdf.SetX(160)
	pdf.CellFormat(40, 4, "Fecha Final:", "", 0,
		"L", false, 0, "")
	pdf.SetX(185)
	pdf.CellFormat(40, 4, miNominadetalle.Fechafinal.Format("02/01/2006"), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Cedula No.:", "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, Coma(miNominadetalle.Codigo), "", 0,
		"L", false, 0, "")
	pdf.SetX(75)
	pdf.CellFormat(40, 4, "Nombre:", "", 0,
		"L", false, 0, "")
	pdf.SetX(93)
	pdf.CellFormat(40, 4, ene(miNominadetalle.Empleadonombre), "", 0,
		"L", false, 0, "")
	pdf.SetX(160)
	pdf.CellFormat(40, 4, "Centro:", "", 0,
		"L", false, 0, "")
	pdf.SetX(175)
	pdf.CellFormat(40, 4, miNominadetalle.Centro, "", 0,
		"L", false, 0, "")
	pdf.SetX(185)
	pdf.CellFormat(40, 4, "Dias:", "", 0,
		"L", false, 0, "")
	pdf.SetX(195)
	pdf.CellFormat(40, 4, miNominadetalle.Dias, "", 0,
		"L", false, 0, "")

	// RELLENO TITULO
	pdf.SetFont("Arial", "", 10)
	pdf.SetY(65)
	pdf.SetX(20)
	pdf.SetFillColor(224, 231, 239)
	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(4)
	pdf.SetX(20)

	pdf.CellFormat(184, 6, "No.", "0", 0,
		"L", true, 0, "")
	pdf.SetX(28)
	pdf.CellFormat(40, 6, "CONCEPTO", "0", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 6, "INGRESOS", "0", 0,
		"R", false, 0, "")
	pdf.SetX(93)
	pdf.CellFormat(40, 6, "CONCEPTO", "0", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 6, "EGRESOS", "0", 0,
		"R", false, 0, "")
	pdf.Ln(8)
}

// FILA1  Y FILA 2
var cordenada1 float64
var cordenada2 float64

func NominaFilaDetalleDevengado(pdf *gofpdf.Fpdf, miValor1, miConcepto1 string, a int) {
	ene := pdf.UnicodeTranslatorFromDescriptor("")

	if Flotante(miValor1) != 0 {

		if Flotante(miValor1) != 0 {
			pdf.SetY(cordenada1)
			pdf.SetX(30)
			pdf.CellFormat(40, 4, ene(miConcepto1), "", 0,
				"L", false, 0, "")
			pdf.SetX(45)
			pdf.CellFormat(40, 4, Coma(miValor1), "", 0,
				"R", false, 0, "")
			cordenada1 = cordenada1 + 4
		}
		numerofila++
	}
}

func NominaPieDePagina(pdf *gofpdf.Fpdf, miNomina nomina) {

	Totalletras, err := IntLetra(Cadenaentero(miNomina.Neto))
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
}

func NominaLinea(pdf *gofpdf.Fpdf) {
	pdf.SetY(150)
	pdf.SetX(160)
	pdf.Line(160, 150, 204, 150)
}

// INICIA NOMINA TODOS PDF
func NominaTodosCabecera(pdf *gofpdf.Fpdf) {
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
	pdf.SetX(105)
	pdf.CellFormat(190, 6, "Fecha Final", "0", 0,
		"L", false, 0, "")
	pdf.SetX(130)
	pdf.CellFormat(190, 6, "Devengado", "0", 0,
		"L", false, 0, "")
	pdf.SetX(155)
	pdf.CellFormat(190, 6, "Deducciones", "0", 0,
		"L", false, 0, "")
	pdf.SetX(193)
	pdf.CellFormat(190, 6, "Neto", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}

func NominaTodosDetalle(pdf *gofpdf.Fpdf, miFila nominaLista, a int) {
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(21)
	pdf.CellFormat(181, 4, strconv.Itoa(numerofila), "", 0,
		"L", false, 0, "")
	pdf.SetX(32)
	pdf.CellFormat(40, 4, miFila.Numero, "", 0,
		"L", false, 0, "")
	pdf.SetX(55)
	pdf.CellFormat(40, 4, miFila.Prefijo, "", 0,
		"L", false, 0, "")
	pdf.SetX(75)
	pdf.CellFormat(40, 4, miFila.Fechainicial.Format("02/01/2006"), "", 0, "L", false, 0, "")
	pdf.SetX(105)
	pdf.CellFormat(40, 4, miFila.Fechafinal.Format("02/01/2006"), "", 0, "L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, Coma(miFila.Devengado), "", 0,
		"R", false, 0, "")
	pdf.SetX(138)
	pdf.CellFormat(40, 4, Coma(miFila.Deducciones), "", 0,
		"R", false, 0, "")
	pdf.SetX(164)
	pdf.CellFormat(40, 4, Coma(miFila.Neto), "", 0,
		"R", false, 0, "")
	pdf.Ln(4)
	numerofila++
}
func NominaTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string

	consulta = " SELECT nomina.numero, nomina.prefijo, nomina.fechainicial, "
	consulta += " nomina.fechafinal, nomina.devengado, nomina.deducciones,"
	consulta += " nomina.neto"
	consulta += " FROM nomina"
	consulta += " ORDER BY cast(nomina.numero as integer) ASC"

	t := []nominaLista{}
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

	NominaTodosCabecera(pdf)
	// tercera pagina

	for i, miFila := range t {
		var a = i + 1
		if math.Mod(float64(a), 49) == 0 {
			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			NominaTodosCabecera(pdf)
		}
		NominaTodosDetalle(pdf, miFila, a)
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
func NominaExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	//Codigo := mux.Vars(r)["codigo"]

	// traer NOMINA
	miNomina := []nominaLista{}
	err := db.Select(&miNomina, "SELECT numero, prefijo, fechainicial, fechafinal, devengado, deducciones, neto FROM nomina ")
	if err != nil {
		log.Fatalln(err)
	}

	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	f := excelize.NewFile()

	// FUNCION ANCHO DE LA COLUMNA
	if err = f.SetColWidth("Sheet1", "A", "A", 10); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.SetColWidth("Sheet1", "B", "B", 10); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.SetColWidth("Sheet1", "C", "C", 15); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "D", "D", 15); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "E", "E", 15); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "F", "F", 15); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "G", "G", 15); err != nil {
		fmt.Println(err)
		return
	}

	// FUNCION PARA UNIR DOS CELDAS
	if err = f.MergeCell("Sheet1", "A1", "G1"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A2", "G2"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A3", "G3"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A4", "G4"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A5", "G5"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A6", "G6"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A7", "G7"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A8", "G8"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A9", "G9"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A10", "G10"); err != nil {
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
	f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel), "Numero")
	f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel), "Prefijo")
	f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel), "Fecha Inicial")
	f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel), "Fecha Final")
	f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel), "Devengado")
	f.SetCellValue("Sheet1", "F"+strconv.Itoa(filaExcel), "Deducciones")
	f.SetCellValue("Sheet1", "G"+strconv.Itoa(filaExcel), "Neto")

	f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel), "A"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel), "B"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel), "C"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel), "D"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "E"+strconv.Itoa(filaExcel), "E"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "F"+strconv.Itoa(filaExcel), "F"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "G"+strconv.Itoa(filaExcel), "G"+strconv.Itoa(filaExcel), estiloCabecera)

	filaExcel++

	for i, miFila := range miNomina {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel+i), Entero(miFila.Numero))
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel+i), miFila.Prefijo)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel+i), miFila.Fechainicial.Format("02/01/2006"))
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel+i), miFila.Fechafinal.Format("02/01/2006"))
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel+i), Entero(miFila.Devengado))
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(filaExcel+i), Entero(miFila.Deducciones))
		f.SetCellValue("Sheet1", "G"+strconv.Itoa(filaExcel+i), Entero(miFila.Neto))

		f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel+i), "A"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel+i), "B"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel+i), "C"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel+i), "D"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "E"+strconv.Itoa(filaExcel+i), "E"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "F"+strconv.Itoa(filaExcel+i), "F"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "G"+strconv.Itoa(filaExcel+i), "G"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
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

// NOMINA INDIVIDUAL EXCEL
func NominaIndividualExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	//var miColumna=1
	// traer NOMINA
	miNomina := nomina{}
	err := db.Get(&miNomina, "SELECT * FROM nomina where numero=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}

	// traer detalle
	miDetalle := []nominadetalleeditar{}
	err2 := db.Select(&miDetalle, NominaConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	f := excelize.NewFile()

	// FUNCION ANCHO DE LA COLUMNA
	if err = f.SetColWidth("Sheet1", "A", "A", 15); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.SetColWidth("Sheet1", "B", "B", 35); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.SetColWidth("Sheet1", "C", "C", 7); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "D", "D", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "E", "E", 7); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "F", "F", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "G", "G", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "H", "H", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "I", "I", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "J", "J", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "K", "K", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "L", "L", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "M", "M", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "N", "N", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "O", "O", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "P", "P", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "Q", "Q", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "R", "R", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "S", "S", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "T", "T", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "U", "U", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "V", "V", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "W", "W", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "X", "X", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "Y", "Y", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "Z", "Z", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "AA", "AA", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "AB", "AB", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "AC", "AC", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "AD", "AD", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "AE", "AE", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "AF", "AF", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "AG", "AG", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "AH", "AH", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "AI", "AI", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "AJ", "AJ", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "AK", "AK", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "AL", "AL", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "AM", "AM", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "AN", "AN", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "AO", "AO", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "AP", "AP", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "AQ", "AQ", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "AR", "AR", 13); err != nil {
		fmt.Println(err)
		return
	}

	// FUNCION PARA UNIR DOS CELDAS
	if err = f.MergeCell("Sheet1", "A1", "L1"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A2", "L2"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A3", "L3"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A4", "L4"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A5", "L5"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A6", "L6"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A7", "L7"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A8", "L8"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A9", "L9"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A10", "L10"); err != nil {
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
	f.SetCellValue("Sheet1", "B12", miNomina.Numero)

	f.SetCellValue("Sheet1", "D12", "Prefijo")
	f.SetCellValue("Sheet1", "E12", miNomina.Prefijo)
	f.SetCellValue("Sheet1", "G12", "Fecha Inicial")
	f.SetCellValue("Sheet1", "H12", miNomina.Fechainicial.Format("02/01/2006"))
	f.SetCellValue("Sheet1", "J12", "Fecha Final")
	f.SetCellValue("Sheet1", "K12", miNomina.Fechafinal.Format("02/01/2006"))

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

	//estiloTexto, err := f.NewStyle(`{"font":{"bold":false,"italic":false,"family":"Arial","size":10,"color":"#000000"}}`)

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
	f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel), "A"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel), "Nombre")
	f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel), "B"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel), "Centro")
	f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel), "C"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel), "Sueldo")
	f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel), "D"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel), "Dias")
	f.SetCellStyle("Sheet1", "E"+strconv.Itoa(filaExcel), "E"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellValue("Sheet1", "F"+strconv.Itoa(filaExcel), "Trabajado")
	f.SetCellStyle("Sheet1", "F"+strconv.Itoa(filaExcel), "F"+strconv.Itoa(filaExcel), estiloCabecera)
	var micolumna int
	micolumna = 7
	var micolumnanombre string

	if Flotante(miNomina.Transporte) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Transporte")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}

	if Flotante(miNomina.Cesantias) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Cesantias")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}

	if Flotante(miNomina.Intereses) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Intereses")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}
	if Flotante(miNomina.Prima) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Prima")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}
	if Flotante(miNomina.Vacaciones) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Vacaciones")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}

	if Flotante(miNomina.Viaticos) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Viaticos")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}
	if Flotante(miNomina.Horasextras) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Horasextras")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}
	if Flotante(miNomina.Incapacidades) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Incapacidades")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}
	if Flotante(miNomina.Licencias) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "LIcencias")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}
	if Flotante(miNomina.Bonificaciones) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Bonificaciones")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}
	if Flotante(miNomina.Auxilios) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Auxilios")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}
	if Flotante(miNomina.Huelgas) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Huelgas")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}
	if Flotante(miNomina.Conceptos) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Conceptos")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}
	if Flotante(miNomina.Compensaciones) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Compensaciones")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}
	if Flotante(miNomina.Bonos) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Bonos")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}
	if Flotante(miNomina.Comisiones) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Comisiones")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}
	if Flotante(miNomina.Dotaciones) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Dotaciones")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}
	if Flotante(miNomina.Sostenimiento) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Sostenimiento")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}
	if Flotante(miNomina.Teletrabajo) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Teletrabajo")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}

	if Flotante(miNomina.Indemnizaciones) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Indemnizaciones")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}
	// devengado
	micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
	f.SetCellValue("Sheet1", micolumnanombre, "Devengado")
	f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
	micolumna++

	if Flotante(miNomina.Salud) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Salud")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}

	if Flotante(miNomina.Pension) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Pension")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}

	if Flotante(miNomina.Pensionrais) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Pensionrais")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}

	if Flotante(miNomina.Pensionvoluntaria) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Pensionvoluntaria")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}

	if Flotante(miNomina.Solidaridad) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Solidaridad")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}

	if Flotante(miNomina.Subsistencia) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Subsistencia")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}

	if Flotante(miNomina.Sindicatos) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Sindicatos")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}

	if Flotante(miNomina.Sanciones) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Sanciones")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}

	if Flotante(miNomina.Libranzas) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Libranzas")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}

	if Flotante(miNomina.Terceros) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Terceros")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}

	if Flotante(miNomina.Anticipos) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Anticipos")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}

	if Flotante(miNomina.Otras) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Otras")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}

	if Flotante(miNomina.Retencion) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Retencion")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}

	if Flotante(miNomina.Afc) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Afc")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}

	if Flotante(miNomina.Embargos) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Embargos")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}

	if Flotante(miNomina.Educacion) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Educacion")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}

	if Flotante(miNomina.Deuda) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, "Deuda")
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
		micolumna++
	}

	// Deducciones
	micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
	f.SetCellValue("Sheet1", micolumnanombre, "Deducciones")
	f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
	micolumna++

	// Neto
	micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
	f.SetCellValue("Sheet1", micolumnanombre, "Neto")
	f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloCabecera)
	micolumna++

	filaExcel++
	var totalTrabajado float64
	var totalTransporte float64
	var totalDevengado float64
	var totalSalud float64
	var totalPension float64
	var totalDeducciones float64
	var totalNeto float64

	//var i int
	var ultima int

	for i, miFila := range miDetalle {
		totalTrabajado += Flotante(miFila.Trabajado)
		totalTransporte += Flotante(miFila.Transporte)
		totalDevengado += Flotante(miFila.Devengado)
		totalSalud += Flotante(miFila.Salud)
		totalPension += Flotante(miFila.Pension)
		totalDeducciones += Flotante(miFila.Deducciones)
		totalNeto += Flotante(miFila.Neto)

		f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel+i), Flotante(miFila.Codigo))
		f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel+i), "A"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel+i), miFila.Empleadonombre)
		f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel+i), "B"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel+i), Flotante(miFila.Centro))
		f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel+i), "C"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel+i), Flotante(miFila.Sueldo))
		f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel+i), "D"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel+i), Flotante(miFila.Dias))
		f.SetCellStyle("Sheet1", "E"+strconv.Itoa(filaExcel+i), "E"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(filaExcel+i), Flotante(miFila.Trabajado))
		f.SetCellStyle("Sheet1", "F"+strconv.Itoa(filaExcel+i), "F"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)

		micolumna = 7
		// columna transporte
		if Flotante(miNomina.Transporte) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Transporte))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Cesantias) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Cesantias))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Intereses) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Intereses))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Prima) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Prima))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Vacaciones) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Vacaciones))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Viaticos) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Viaticos))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Horasextras) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Horasextras))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Incapacidades) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Incapacidades))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Licencias) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Licencias))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Bonificaciones) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Bonificaciones))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Auxilios) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Auxilios))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Huelgas) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Huelgas))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Conceptos) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Conceptos))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Compensaciones) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Compensaciones))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Bonos) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Bonos))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Comisiones) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Comisiones))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Dotaciones) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Dotaciones))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Sostenimiento) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Sostenimiento))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Teletrabajo) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Teletrabajo))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Indemnizaciones) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Indemnizaciones))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		// devengados
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Devengado))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++

		if Flotante(miNomina.Salud) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Salud))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Pension) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Pension))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Pensionrais) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Pensionrais))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Pensionvoluntaria) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Pensionvoluntaria))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Solidaridad) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Solidaridad))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Subsistencia) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Subsistencia))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Sindicatos) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Sindicatos))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Sanciones) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Sanciones))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Libranzas) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Libranzas))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Terceros) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Terceros))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Anticipos) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Anticipos))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Otras) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Otras))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Retencion) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Retencion))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Afc) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Afc))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Embargos) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Embargos))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Educacion) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Educacion))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		if Flotante(miNomina.Deuda) != 0 {
			micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
			f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Deuda))
			f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
			micolumna++
		}

		// deducciones
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Deducciones))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++

		// Neto
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel+i)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miFila.Neto))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
		ultima = filaExcel + i

	}

	filaExcel = ultima + 1

	// totales
	f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel), "")
	f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel), "TOTALES")
	f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel), "")
	f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel), "")
	f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel), "")
	f.SetCellValue("Sheet1", "F"+strconv.Itoa(filaExcel), Flotante(miNomina.Trabajado))
	f.SetCellStyle("Sheet1", "F"+strconv.Itoa(filaExcel), "F"+strconv.Itoa(filaExcel), estiloNumeroDetalle)

	micolumna = 7
	// columna transporte
	if Flotante(miNomina.Transporte) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Transporte))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Cesantias) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Cesantias))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Intereses) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Intereses))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Prima) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Prima))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Vacaciones) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Vacaciones))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Viaticos) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Viaticos))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Horasextras) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Horasextras))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Incapacidades) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Incapacidades))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Licencias) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Licencias))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Bonificaciones) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Bonificaciones))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Auxilios) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Auxilios))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Huelgas) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Huelgas))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Conceptos) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Conceptos))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Compensaciones) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Compensaciones))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Bonos) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Bonos))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Comisiones) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Comisiones))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Dotaciones) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Dotaciones))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Sostenimiento) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Sostenimiento))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Teletrabajo) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Teletrabajo))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Indemnizaciones) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Indemnizaciones))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	// devengados
	micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
	f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Devengado))
	f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
	micolumna++

	// columna Salud
	if Flotante(miNomina.Salud) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Salud))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Pension) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Pension))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Pensionrais) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Pensionrais))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Pensionvoluntaria) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Pensionvoluntaria))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Solidaridad) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Solidaridad))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Subsistencia) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Subsistencia))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Sindicatos) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Sindicatos))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Sanciones) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Sanciones))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Libranzas) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Libranzas))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Terceros) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Terceros))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Anticipos) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Anticipos))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Otras) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Otras))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Retencion) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Retencion))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Afc) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Afc))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Embargos) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Embargos))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Educacion) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Educacion))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	if Flotante(miNomina.Deuda) != 0 {
		micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
		f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Deuda))
		f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
		micolumna++
	}

	// deducciones
	micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
	f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Deducciones))
	f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
	micolumna++

	// neto
	micolumnanombre, err = excelize.CoordinatesToCellName(micolumna, filaExcel)
	f.SetCellValue("Sheet1", micolumnanombre, Flotante(miNomina.Neto))
	f.SetCellStyle("Sheet1", micolumnanombre, micolumnanombre, estiloNumeroDetalle)
	micolumna++

	// LINEA FINAL
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment;filename=Informe.xlsx")
	w.Header().Set("File-Name", "Informe.xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	err = f.Write(w)
	if err != nil {
		panic(err.Error())
	}

}

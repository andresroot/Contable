package main

// INICIA NOMINA APORTES IMPORTAR PAQUETES
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

// INICIA NOMINA APORTES ESTRUCTURA JSON
type nominaaportesaportesJson struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Nombre string `json:"nombre"`
}

// INICIA NOMINA APORTES APORTES ESTRUCTURA
type nominaaportesLista struct {
	Numero       string
	Fechainicial time.Time
	Fechafinal   time.Time
	Saludgasto   string
	Pensiongasto string
	Riesgos      string
	Icbf         string
	Sena         string
	Caja         string
	Total        string
}

// INICIA  NOMINA APORTES APORTES ESTRUCTURA
type nominaaportesdetalle struct {
	Numero         string
	Fechainicial   time.Time
	Fechafinal     time.Time
	Codigo         string
	Centro         string
	Sueldo         string
	Saludgasto     string
	Pensiongasto   string
	Riesgos        string
	Icbf           string
	Sena           string
	Caja           string
	Total          string
	TerceroDetalle tercero
}

// INICIA PLANILLADETALLE ESTRUCTURA
type nominaaportes struct {
	Numero        string
	Fechainicial  time.Time
	Fechafinal    time.Time
	Codigo        string
	Centro        string
	Sueldo        string
	Saludgasto    string
	Pensiongasto  string
	Riesgos       string
	Icbf          string
	Sena          string
	Caja          string
	Total         string
	Detalle       []nominaaportesdetalle       `json:"Detalle"`
	DetalleEditar []nominaaportesdetalleeditar `json:"DetalleEditar"`
	Accion        string
}

// INICIA NOMINA APORTES DETALLE EDITARr
type nominaaportesdetalleeditar struct {
	Numero         string
	Fechainicial   time.Time
	Fechafinal     time.Time
	Codigo         string
	Empleadonombre string
	Centro         string
	Sueldo         string
	Saludgasto     string
	Pensiongasto   string
	Riesgos        string
	Icbf           string
	Sena           string
	Caja           string
	Total          string
}

// INICIA NOMINA APORTES CONSULTA DETALLE
func NominaaportesConsultaDetalle() string {
	var consulta = ""
	consulta = "select "
	consulta += "nominaaportesdetalle.Codigo as codigo,"
	consulta += "empleado.nombre as empleadonombre,"
	consulta += "nominaaportesdetalle.Centro as centro,"
	consulta += "nominaaportesdetalle.Sueldo as sueldo,"
	consulta += " nominaaportesdetalle.Saludgasto as saludgasto,"
	consulta += " nominaaportesdetalle.Pensiongasto as pensiongasto,"
	consulta += " nominaaportesdetalle.Riesgos as riesgos,"
	consulta += " nominaaportesdetalle.Icbf as icbf,"
	consulta += " nominaaportesdetalle.Sena as sena,"
	consulta += " nominaaportesdetalle.Caja as caja,"
	consulta += " nominaaportesdetalle.Total as total"
	consulta += " from nominaaportesdetalle "
	consulta += " inner join empleado on empleado.codigo=nominaaportesdetalle.codigo"
	consulta += " where nominaaportesdetalle.numero=$1 ORDER BY sueldo DESC"
	log.Println(consulta)
	return consulta
}

// INICIA NOMINA APORTES LISTA
func NominaaportesLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/nominaaportes/nominaaportesLista.html")
	log.Println("Error nominaaportes 0")
	var consulta string

	consulta = " SELECT nominaaportes.numero, nominaaportes.fechainicial, "
	consulta += " nominaaportes.fechafinal, nominaaportes.saludgasto, nominaaportes.pensiongasto,"
	consulta += " nominaaportes.riesgos, nominaaportes.icbf, nominaaportes.sena,"
	consulta += "nominaaportes.caja, nominaaportes.total"
	consulta += " FROM nominaaportes"
	consulta += " ORDER BY cast(nominaaportes.numero as integer) ASC"

	db := dbConn()
	res := []nominaaportesLista{}

	err := db.Select(&res, consulta)

	if err != nil {
		fmt.Println(err)
	}

	varmap := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
	}
	log.Println("Error nominaaportes888")
	tmp.Execute(w, varmap)
}

// INICIA NOMINA APORTES NUEVO
func NominaaportesNuevo(w http.ResponseWriter, r *http.Request) {

	// TRAE COPIA DE EDITAR
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio nominaaportes editar" + Codigo)

	db := dbConn()
	empleados := []empleado{}
	nominaaportes := nominaaportes{}
	det := []nominaaportesdetalleeditar{}

	if Codigo == "False" {
		err := db.Select(&empleados, "SELECT * FROM empleado  ORDER BY sueldo DESC ")
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		// traer NOMINA APORTES

		err := db.Get(&nominaaportes, "SELECT * FROM nominaaportes where numero=$1", Codigo)
		if err != nil {
			log.Fatalln(err)
		}
		err2 := db.Select(&det, NominaaportesConsultaDetalle(), Codigo)
		if err2 != nil {
			fmt.Println(err2)
		}
	}

	parametros := map[string]interface{}{
		"codigo":        Codigo,
		"nominaaportes": nominaaportes,
		"empleado":      empleados,
		"detalle":       det,
		"hosting":       ruta,
	}
	//TERMINA TRAE COPIA DE EDITAR

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/nominaaportes/nominaaportesNuevo.html",
		"vista/nominaaportes/autocompletaCentro.html",
		"vista/nominaaportes/autocompletaempleado.html",
		"vista/nominaaportes/modalColumna.html",
		"vista/nominaaportes/nominaaportesScript.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error nominaaportes nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA INSERTAR COMPROBANTE DE NOMINA APORTES
//func InsertaDetalleComprobanteNominaaportes(miFilaComprobante comprobantedetalle, miComprobante comprobante, miNominaaportes nominaaportes){
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
//	miNominaaportes.Tercero,
//	miNominaaportes.Centro,
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

//INICIA NOMINA APORTES INSERTAR AJAX
func NominaaportesAgregar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var tempNominaaportes nominaaportes

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// carga informacion de la NOMINA APORTES
	err = json.Unmarshal(b, &tempNominaaportes)
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

	Codigoactual = tempNominaaportes.Numero

	if tempNominaaportes.Accion == "Actualizar" {

		// borra detalle anterior
		delForm, err := db.Prepare("DELETE from nominaaportesdetalle WHERE numero=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(tempNominaaportes.Numero)

		// borra detalle inventario
		Borrarinventario(Codigoactual, "Nominaaportes")

		// borra cabecera anterior
		delForm1, err := db.Prepare("DELETE from nominaaportes WHERE numero=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm1.Exec(tempNominaaportes.Numero)
	}

	// INSERTA DETALLE
	for i, x := range tempNominaaportes.Detalle {
		var a = i
		var q string
		q = "insert into nominaaportesdetalle ("
		q += "Numero,"
		q += "Fechainicial,"
		q += "Fechafinal,"
		q += "Codigo,"
		q += "Centro,"
		q += "Sueldo,"
		q += "Saludgasto,"
		q += "Pensiongasto,"
		q += "Riesgos,"
		q += "Icbf,"
		q += "Sena,"
		q += "Caja,"
		q += "Total"
		q += " ) values("
		q += parametros(13)
		q += ")"
		log.Println("Cadena SQL " + q)
		insForm, err := db.Prepare(q)
		if err != nil {
			panic(err.Error())
		}

		// TERMINA NOMINA APORTES GRABAR INSERTAR
		_, err = insForm.Exec(
			x.Numero,
			x.Fechainicial,
			x.Fechafinal,
			Quitacoma(x.Codigo),
			x.Centro,
			Quitacoma(x.Sueldo),
			Quitacoma(x.Saludgasto),
			Quitacoma(x.Pensiongasto),
			Quitacoma(x.Riesgos),
			Quitacoma(x.Icbf),
			Quitacoma(x.Sena),
			Quitacoma(x.Caja),
			Quitacoma(x.Total))

		if err != nil {
			panic(err)
		}
		log.Println("Insertar Codigo \n", x.Codigo, a)
	}

	// INICIA INSERTAR PLANILLAS
	var q string
	q = "insert into nominaaportes ("
	q += "Numero,"
	q += "Fechainicial,"
	q += "Fechafinal,"
	q += "Saludgasto,"
	q += "Pensiongasto,"
	q += "Riesgos,"
	q += "Icbf,"
	q += "Sena,"
	q += "Caja,"
	q += "Total"
	q += " ) values("
	q += parametros(10)
	q += " ) "

	log.Println("Cadena SQL " + q)
	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	layout := "2006-01-02"
	//log.Println("Hora", tempNominaaportes.Fechainicial.Format("02/01/2006"))
	//currentTime := time.Now()

	_, err = insForm.Exec(
		tempNominaaportes.Numero,
		tempNominaaportes.Fechainicial.Format(layout),
		tempNominaaportes.Fechafinal.Format(layout),
		Quitacoma(tempNominaaportes.Saludgasto),
		Quitacoma(tempNominaaportes.Pensiongasto),
		Quitacoma(tempNominaaportes.Riesgos),
		Quitacoma(tempNominaaportes.Icbf),
		Quitacoma(tempNominaaportes.Sena),
		Quitacoma(tempNominaaportes.Caja),
		Quitacoma(tempNominaaportes.Total))

	if err != nil {
		panic(err)
	}

	// INSERTAR COMPROBANTE CONTABILIDAD
	var tempComprobante comprobante
	var tempComprobanteDetalle comprobantedetalle
	tempComprobante.Documento = "19"
	tempComprobante.Numero = tempNominaaportes.Numero
	tempComprobante.Fecha = tempNominaaportes.Fechafinal
	tempComprobante.Fechaconsignacion = tempNominaaportes.Fechafinal
	tempComprobante.Debito = tempNominaaportes.Total
	tempComprobante.Credito = tempNominaaportes.Total
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

	log.Println("INSERTA CONTABILIDAD %s\n", tempComprobante.Documento, tempComprobante.Numero)

	var totalDebito float64
	var totalCredito float64
	var fila int
	fila = 0
	totalDebito = 0
	totalCredito = 0

	for i, x := range tempNominaaportes.Detalle {

		// datos emppelado
		miEmpleado := empleado{}
		err = db.Get(&miEmpleado, "SELECT * FROM empleado where codigo=$1", x.Codigo)
		if err != nil {
			log.Fatalln(err)
		}

		// SALUD DEBITO
		if Flotante(x.Saludgasto) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Saludgasto)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Saludgasto
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Saludgasto))
			tempComprobanteDetalle.Credito = ""
			tempComprobanteDetalle.Tercero = miEmpleado.Salud
			tempComprobanteDetalle.Concepto = miEmpleado.Saludnombre
			InsertaDetalleComprobanteNominaaportes(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// PENSION DEBITO
		if Flotante(x.Pensiongasto) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Pensiongasto)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Pensiongasto
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Pensiongasto))
			tempComprobanteDetalle.Credito = ""
			tempComprobanteDetalle.Tercero = miEmpleado.Pension
			tempComprobanteDetalle.Concepto = miEmpleado.Pensionnombre
			InsertaDetalleComprobanteNominaaportes(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// RIESGOS DEBITO
		if Flotante(x.Riesgos) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Riesgos)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Riesgos
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Riesgos))
			tempComprobanteDetalle.Credito = ""
			tempComprobanteDetalle.Tercero = miEmpleado.Riesgos
			tempComprobanteDetalle.Concepto = miEmpleado.Riesgosnombre
			InsertaDetalleComprobanteNominaaportes(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// ICBF DEBITO
		if Flotante(x.Icbf) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Icbf)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Icbf
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Icbf))
			tempComprobanteDetalle.Credito = ""
			tempComprobanteDetalle.Tercero = miEmpleado.Icbf
			tempComprobanteDetalle.Concepto = miEmpleado.Icbfnombre
			InsertaDetalleComprobanteNominaaportes(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// SENA DEBITO
		if Flotante(x.Sena) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Sena)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Sena
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Sena))
			tempComprobanteDetalle.Credito = ""
			tempComprobanteDetalle.Tercero = miEmpleado.Sena
			tempComprobanteDetalle.Concepto = miEmpleado.Senanombre
			InsertaDetalleComprobanteNominaaportes(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// CAJA DEBITO
		if Flotante(x.Caja) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalDebito += Flotante(x.Caja)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Caja
			tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(x.Caja))
			tempComprobanteDetalle.Credito = ""
			tempComprobanteDetalle.Tercero = miEmpleado.Caja
			tempComprobanteDetalle.Concepto = miEmpleado.Cajanombre
			InsertaDetalleComprobanteNominaaportes(tempComprobanteDetalle, tempComprobante, x)
			log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito, i))
		}

		// SALUD CREDITO
		if Flotante(x.Saludgasto) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Saludgasto)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Salud
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Saludgasto))
			tempComprobanteDetalle.Tercero = miEmpleado.Salud
			tempComprobanteDetalle.Concepto = miEmpleado.Saludnombre
			InsertaDetalleComprobanteNominaaportes(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}

		// PENSION CREDITO
		if Flotante(x.Pensiongasto) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Pensiongasto)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Pension
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Pensiongasto))
			tempComprobanteDetalle.Tercero = miEmpleado.Pension
			tempComprobanteDetalle.Concepto = miEmpleado.Pensionnombre
			InsertaDetalleComprobanteNominaaportes(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}

		// RIESGOS CREDITO
		if Flotante(x.Riesgos) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Riesgos)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Riesgoscxp
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Riesgos))
			tempComprobanteDetalle.Tercero = miEmpleado.Riesgos
			tempComprobanteDetalle.Concepto = miEmpleado.Riesgosnombre
			InsertaDetalleComprobanteNominaaportes(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}

		// ICBF CREDITO
		if Flotante(x.Icbf) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Icbf)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Icbfcxp
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Icbf))
			tempComprobanteDetalle.Tercero = miEmpleado.Icbf
			tempComprobanteDetalle.Concepto = miEmpleado.Icbfnombre
			InsertaDetalleComprobanteNominaaportes(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}

		// SENA CREDITO
		if Flotante(x.Sena) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Sena)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Senacxp
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Sena))
			tempComprobanteDetalle.Tercero = miEmpleado.Sena
			tempComprobanteDetalle.Concepto = miEmpleado.Senanombre
			InsertaDetalleComprobanteNominaaportes(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}

		// CAJA CREDITO
		if Flotante(x.Caja) != 0 {
			fila = fila + 1
			tempComprobanteDetalle.Fila = strconv.Itoa(fila)
			totalCredito += Flotante(x.Caja)
			tempComprobanteDetalle.Cuenta = TraerParametrosContabilidad().Cajacxp
			tempComprobanteDetalle.Debito = ""
			tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(x.Caja))
			tempComprobanteDetalle.Tercero = miEmpleado.Caja
			tempComprobanteDetalle.Concepto = miEmpleado.Cajanombre
			InsertaDetalleComprobanteNominaaportes(tempComprobanteDetalle, tempComprobante, x)
			log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito, i))
		}

	}

	// INSERTAR CUENTA DEBITO NOMINA APORTES

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
func InsertaDetalleComprobanteNominaaportes(miFilaComprobante comprobantedetalle, miComprobante comprobante, miCompra nominaaportesdetalle) {
	db := dbConn()

	// traer tercero
	miTercero := tercero{}
	err3 := db.Get(&miTercero, "SELECT * FROM tercero where codigo=$1", miCompra.Codigo)
	if err3 != nil {
		log.Fatalln(err3)
	}

	if miFilaComprobante.Concepto == "" {
		miFilaComprobante.Concepto = miTercero.Nombre
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
		miFilaComprobante.Tercero,
		miCompra.Centro,
		miFilaComprobante.Concepto,
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

// INICIA NOMINA APORTES EXISTE
func NominaaportesExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM nominaaportes  WHERE codigo=$1", Codigo)
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
//// INICIA NOMINA APORTES EDITAR
func NominaaportesEditar(w http.ResponseWriter, r *http.Request) {
	Numero := mux.Vars(r)["codigo"]
	log.Println("inicio nominaaportes editar" + Numero)
	db := dbConn()

	// traer NOMINA APORTES
	v := nominaaportes{}
	err := db.Get(&v, "SELECT * FROM nominaaportes where numero=$1", Numero)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []nominaaportesdetalleeditar{}

	err2 := db.Select(&det, NominaaportesConsultaDetalle(), Numero)
	if err2 != nil {
		fmt.Println(err2)
	}
	parametros := map[string]interface{}{
		"nominaaportes": v,
		"detalle":       det,
		"hosting":       ruta,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/nominaaportes/nominaaportesEditar.html",
		"vista/nominaaportes/autocompletaCentro.html",
		"vista/nominaaportes/autocompletaempleado.html",
		"vista/nominaaportes/modalColumna.html",
		"vista/nominaaportes/nominaaportesScript.html")

	fmt.Printf("%v, %v", miTemplate, err)

	miTemplate.Execute(w, parametros)
}

// INICIA NOMINA APORTES BORRAR
func NominaaportesBorrar(w http.ResponseWriter, r *http.Request) {
	Numero := mux.Vars(r)["codigo"]
	log.Println("inicio nominaaportes editar" + Numero)
	db := dbConn()

	// traer NOMINA APORTES
	v := nominaaportes{}
	err := db.Get(&v, "SELECT * FROM nominaaportes where numero=$1", Numero)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []nominaaportesdetalleeditar{}

	err2 := db.Select(&det, NominaaportesConsultaDetalle(), Numero)
	if err2 != nil {
		fmt.Println(err2)
	}
	parametros := map[string]interface{}{
		"codigo":        Numero,
		"nominaaportes": v,
		"detalle":       det,
		"hosting":       ruta,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/nominaaportes/nominaaportesBorrar.html",
		"vista/nominaaportes/autocompletaCentro.html",
		"vista/nominaaportes/autocompletaempleado.html",
		"vista/nominaaportes/modalColumna.html",
		"vista/nominaaportes/nominaaportesScript.html")

	fmt.Printf("%v, %v", miTemplate, err)

	miTemplate.Execute(w, parametros)
}

// INICIA NOMINA APORTES ELIMINAR
func NominaaportesEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	codigo := mux.Vars(r)["codigo"]

	// borrar NOMINA APORTES
	delForm, err := db.Prepare("DELETE from nominaaportes WHERE numero=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(codigo)

	// borar detalle
	delForm1, err := db.Prepare("DELETE from nominaaportesdetalle WHERE numero=$1")
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
	http.Redirect(w, r, "/NominaaportesLista", 301)
}

// INICIA NOMINA APORTES PDF
//var numerofila = 0

func NominaaportesPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Numero := mux.Vars(r)["numero"]

	// traer NOMINA APORTES
	miNominaaportes := nominaaportes{}
	err := db.Get(&miNominaaportes, "SELECT * FROM nominaaportes where numero=$1", Numero)
	if err != nil {
		log.Fatalln(err)
	}

	// traer detalle
	miDetalle := []nominaaportesdetalleeditar{}
	err2 := db.Select(&miDetalle, NominaaportesConsultaDetalle(), Numero)
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
		NominaaportesCabecera(pdf, miFila)
		numerofila = 1
		NominaaportesFilaDetalle(pdf, miFila.Sueldo, a, "SUELDO")
		NominaaportesFilaDetalle(pdf, miFila.Saludgasto, a, "SALUD")
		NominaaportesFilaDetalle(pdf, miFila.Pensiongasto, a, "PENSION")
		NominaaportesFilaDetalle(pdf, miFila.Riesgos, a, "RIESGOS")
		NominaaportesFilaDetalle(pdf, miFila.Icbf, a, "ICBF")
		NominaaportesFilaDetalle(pdf, miFila.Sena, a, "SENA")
		NominaaportesFilaDetalle(pdf, miFila.Caja, a, "CAJA")
		NominaaportesFilaDetalle(pdf, miFila.Total, a, "TOTAL")
	}
	//var filas = len(miDetalle)
	// menos de 32
	//if filas <= 32 {
	//	for i, miFila := range miDetalle {
	//		var a = i + 1
	//		NominaaportesFilaDetalle(pdf, miFila, a)
	//	}
	//	NominaaportesPieDePagina(pdf, miNominaaportes)
	//} else {
	//	// mas de 32 y menos de 73
	//	if filas > 32 && filas <= 73 {
	//		// primera pagina
	//		for i, miFila := range miDetalle {
	//			var a = i + 1
	//			if a <= 41 {
	//				NominaaportesFilaDetalle(pdf, miFila, a)
	//			}
	//		}
	//		NominaaportesLinea(pdf)
	//		// segunda pagina
	//		pdf.AddPage()
	//		NominaaportesCabecera(pdf, miNominaaportes)
	//		for i, miFila := range miDetalle {
	//			var a = i + 1
	//			if a > 41 {
	//				NominaaportesFilaDetalle(pdf, miFila, a)
	//			}
	//		}
	//
	//		NominaaportesPieDePagina(pdf, miNominaaportes)
	//	} else {
	//		// mas de 80
	//
	//		// primera pagina
	//		for i, miFila := range miDetalle {
	//			var a = i + 1
	//			if a <= 41 {
	//				NominaaportesFilaDetalle(pdf, miFila, a)
	//			}
	//		}
	//		NominaaportesLinea(pdf)
	//		pdf.AddPage()
	//		NominaaportesCabecera(pdf, miNominaaportes)
	//		// segunda pagina
	//		for i, miFila := range miDetalle {
	//			var a = i + 1
	//			if a > 41 && a <= 82 {
	//				NominaaportesFilaDetalle(pdf, miFila, a)
	//			}
	//		}
	//		NominaaportesLinea(pdf)
	//		pdf.AddPage()
	//		NominaaportesCabecera(pdf, miNominaaportes)
	//		// tercera pagina
	//		for i, miFila := range miDetalle {
	//			var a = i + 1
	//			if a > 82 {
	//				NominaaportesFilaDetalle(pdf, miFila, a)
	//			}
	//		}
	//		NominaaportesPieDePagina(pdf, miNominaaportes)
	//	}
	//}
	//NominaaportesPieDePagina(pdf, miNominaaportes)
	// genera pdf
	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

func NominaaportesCabecera(pdf *gofpdf.Fpdf, miNominaaportesdetalle nominaaportesdetalleeditar) {
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
	pdf.CellFormat(40, 4, "Nominaaportes Numero:", "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, miNominaaportesdetalle.Numero, "", 0,
		"L", false, 0, "")
	pdf.SetX(90)
	pdf.CellFormat(40, 4, "Fecha Inicial:", "", 0,
		"L", false, 0, "")
	pdf.SetX(115)
	pdf.CellFormat(40, 4, miNominaaportesdetalle.Fechainicial.Format("02/01/2006"), "", 0,
		"L", false, 0, "")
	pdf.SetX(138)
	pdf.CellFormat(40, 4, "Fecha Final:", "", 0,
		"L", false, 0, "")
	pdf.SetX(160)
	pdf.CellFormat(40, 4, miNominaaportesdetalle.Fechafinal.Format("02/01/2006"), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Cedula No.:", "", 0,
		"L", false, 0, "")
	pdf.SetX(35)
	pdf.CellFormat(40, 4, miNominaaportesdetalle.Codigo, "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, "Cedula No.:", "", 0,
		"L", false, 0, "")
	pdf.SetX(70)
	pdf.CellFormat(40, 4, miNominaaportesdetalle.Empleadonombre, "", 0,
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
func NominaaportesFilaDetalle(pdf *gofpdf.Fpdf, miFila string, a int, miConcepto string) {

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

func NominaaportesPieDePagina(pdf *gofpdf.Fpdf, miNominaaportes nominaaportes) {

	Totalletras, err := IntLetra(Cadenaentero(miNominaaportes.Total))
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

func NominaaportesLinea(pdf *gofpdf.Fpdf) {
	pdf.SetY(259)
	pdf.Line(20, 259, 204, 259)
}

// INICIA NOMINA APORTES TODOS PDF
func NominaaportesTodosCabecera(pdf *gofpdf.Fpdf) {
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

func NominaaportesTodosDetalle(pdf *gofpdf.Fpdf, miFila nominaaportesLista, a int) {
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(21)
	pdf.CellFormat(181, 4, strconv.Itoa(numerofila), "", 0,
		"L", false, 0, "")
	pdf.SetX(32)
	pdf.CellFormat(40, 4, Subcadena(miFila.Numero, 0, 12), "", 0,
		"L", false, 0, "")
	pdf.SetX(55)
	pdf.CellFormat(40, 4, miFila.Fechainicial.Format("02/01/2006"), "", 0, "L", false, 0, "")
	pdf.SetX(75)
	pdf.CellFormat(40, 4, miFila.Fechafinal.Format("02/01/2006"), "", 0, "L", false, 0, "")
	pdf.SetX(130)
	pdf.CellFormat(40, 4, Coma(miFila.Saludgasto), "", 0,
		"L", false, 0, "")
	pdf.SetX(164)
	pdf.CellFormat(40, 4, Coma(miFila.Pensiongasto), "", 0,
		"R", false, 0, "")
	pdf.Ln(4)
	numerofila++
}
func NominaaportesTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string

	consulta = " SELECT nominaaportes.numero, nominaaportes.prefijo, nominaaportes.fechainicial, "
	consulta += " nominaaportes.fechafinal, nominaaportes.devengado, nominaaportes.deducciones,"
	consulta += " nominaaportes.neto"
	consulta += " FROM nominaaportes"
	consulta += " ORDER BY cast(nominaaportes.numero as integer) ASC"

	t := []nominaaportesLista{}
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
		pdf.CellFormat(190, 10, "DATOS DOCUMENTO NOMINA APORTES", "0", 0,
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
		NominaaportesTodosDetalle(pdf, miFila, a)
	}

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

// TERMINA PEDIDO TODOS PDF

// NOMINA APORTES EXCEL
func NominaaportesIndividualExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]

	// traer NOMINA APORTES
	miNominaaportes := nominaaportes{}
	err := db.Get(&miNominaaportes, "SELECT * FROM nominaaportes where numero=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}

	// traer detalle
	miDetalle := []nominaaportesdetalleeditar{}
	err2 := db.Select(&miDetalle, NominaaportesConsultaDetalle(), Codigo)
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
	f.SetCellValue("Sheet1", "A9", "LISTADO DE DOCUMENTOS NOMINA APORTES")
	f.SetCellValue("Sheet1", "A10", "")

	f.SetCellValue("Sheet1", "A12", "Numero")
	f.SetCellValue("Sheet1", "B12", miNominaaportes.Numero)
	f.SetCellValue("Sheet1", "C12", "Fecha Inicial")
	f.SetCellValue("Sheet1", "D12", miNominaaportes.Fechainicial)
	f.SetCellValue("Sheet1", "E12", "Fecha Final")
	f.SetCellValue("Sheet1", "F12", miNominaaportes.Fechafinal)

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
	f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel), "Dias")
	f.SetCellValue("Sheet1", "F"+strconv.Itoa(filaExcel), "Trabajado")
	f.SetCellValue("Sheet1", "G"+strconv.Itoa(filaExcel), "Transporte")
	f.SetCellValue("Sheet1", "H"+strconv.Itoa(filaExcel), "Devengado")
	f.SetCellValue("Sheet1", "I"+strconv.Itoa(filaExcel), "Salud")
	f.SetCellValue("Sheet1", "J"+strconv.Itoa(filaExcel), "Pension")
	f.SetCellValue("Sheet1", "K"+strconv.Itoa(filaExcel), "Deducciones")
	f.SetCellValue("Sheet1", "L"+strconv.Itoa(filaExcel), "Neto")

	f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel), "A"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel), "B"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel), "C"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel), "D"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "E"+strconv.Itoa(filaExcel), "E"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel), "F"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel), "G"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel), "H"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel), "I"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "E"+strconv.Itoa(filaExcel), "J"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel), "K"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel), "L"+strconv.Itoa(filaExcel), estiloCabecera)

	filaExcel++
	var totalSalud float64
	var totalPension float64
	var totalRiesgos float64
	var totalIcbf float64
	var totalSena float64
	var totalCaja float64
	var totalTotal float64

	var i int
	var ultima int

	for i, miFila := range miDetalle {
		totalSalud += Flotante(miFila.Saludgasto)
		totalPension += Flotante(miFila.Pensiongasto)
		totalRiesgos += Flotante(miFila.Riesgos)
		totalIcbf += Flotante(miFila.Icbf)
		totalSena += Flotante(miFila.Sena)
		totalCaja += Flotante(miFila.Caja)
		totalTotal += Flotante(miFila.Total)

		f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel+i), miFila.Codigo)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel+i), miFila.Empleadonombre)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel+i), miFila.Centro)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel+i), Flotante(miFila.Sueldo))
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel+i), Flotante(miFila.Saludgasto))
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(filaExcel+i), Flotante(miFila.Pensiongasto))
		f.SetCellValue("Sheet1", "G"+strconv.Itoa(filaExcel+i), Flotante(miFila.Riesgos))
		f.SetCellValue("Sheet1", "H"+strconv.Itoa(filaExcel+i), Flotante(miFila.Icbf))
		f.SetCellValue("Sheet1", "I"+strconv.Itoa(filaExcel+i), Flotante(miFila.Sena))
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
	f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel), totalSalud)
	f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel), totalPension)
	f.SetCellValue("Sheet1", "F"+strconv.Itoa(filaExcel), totalRiesgos)
	f.SetCellValue("Sheet1", "G"+strconv.Itoa(filaExcel), totalIcbf)
	f.SetCellValue("Sheet1", "H"+strconv.Itoa(filaExcel), totalSena)
	f.SetCellValue("Sheet1", "I"+strconv.Itoa(filaExcel), totalCaja)
	f.SetCellValue("Sheet1", "J"+strconv.Itoa(filaExcel), totalTotal)

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

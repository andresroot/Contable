package main

// INICIA SOPORTE SERVICIO IMPORTAR PAQUETES
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
	"time"
)

// INICIA SOPORTE SERVICIO ESTRUCTURA JSON
type soporteservicioJson struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Nombre string `json:"nombre"`
}

// INICIA SOPORTE SERVICIO ESTRUCTURA
type soporteservicioLista struct {
	Codigo            string
	Fecha             time.Time
	Neto              string
	Tercero           string
	TerceroNombre     string
	CentroNombre      string
	AlmacenistaNombre string
}

// INICIA SOPORTE SERVICIO ESTRUCTURA
type soporteservicio struct {
	Resolucionsoporte           string
	Codigo                      string
	Fecha                       time.Time
	Vence                       time.Time
	Hora                        string
	Descuento                   string
	Subtotaldescuento19         string
	Subtotaldescuento5          string
	Subtotaldescuento0          string
	Subtotal                    string
	Subtotal19                  string
	Subtotal5                   string
	Subtotal0                   string
	Subtotaliva19               string
	Subtotaliva5                string
	Subtotaliva0                string
	Subtotalbase19              string
	Subtotalbase5               string
	Subtotalbase0               string
	TotalIva                    string
	Total                       string
	PorcentajeRetencionFuente   string
	TotalRetencionFuente        string
	PorcentajeRetencionIca      string
	TotalRetencionIca           string
	Neto                        string
	Items                       string
	Formadepago                 string
	Mediodepago                 string
	Tercero                     string
	Almacenista                 string
	Accion                      string
	Detalle                     []soporteserviciodetalle       `json:"Detalle"`
	DetalleEditar               []soporteserviciodetalleeditar `json:"DetalleEditar"`
	TerceroDetalle              tercero
	Pedidosoporteservicio       string
	Tipo                        string
	Centro                      string
	Soporteserviciocuenta0      string
	Soporteservicionombre0      string
	Soporteserviciocuentaretfte string
	Soporteservicionombreretfte string
	FormadepagoDetalle          formadepago
	MediodepagoDetalle          mediodepago
	AlmacenistaDetalle          almacenista
	CentroDetalle               centro
}

// INICIA SOPORTE SERVICIODETALLE ESTRUCTURA
type soporteserviciodetalle struct {
	Id                string
	Codigo            string
	Fila              string
	Cantidad          string
	Precio            string
	Descuento         string
	Montodescuento    string
	Sigratis          string
	Subtotal          string
	Subtotaldescuento string
	Pagina            string
	Bodega            string
	Tipo              string
	Fecha             time.Time
	Nombreservicio    string
	Unidadservicio    string
	Codigoservicio    string
}

// INICIA SOPORTE SERVICIO DETALLE EDITARr
type soporteserviciodetalleeditar struct {
	Id                string
	Codigo            string
	Fila              string
	Cantidad          string
	Precio            string
	Descuento         string
	Montodescuento    string
	Sigratis          string
	Subtotal          string
	Subtotaldescuento string
	Pagina            string
	Bodega            string
	Tipo              string
	Fecha             time.Time
	Nombreservicio    string
	Unidadservicio    string
	Codigoservicio    string
}

// INICIA SOPORTE SERVICIO CONSULTA DETALLE
func SoporteservicioConsultaDetalle() string {
	var consulta = ""
	consulta = "select "
	consulta += "soporteserviciodetalle.Id as id ,"
	consulta += "soporteserviciodetalle.Codigo as codigo,"
	consulta += "soporteserviciodetalle.Fila as fila,"
	consulta += "soporteserviciodetalle.Cantidad as cantidad,"
	consulta += "soporteserviciodetalle.Precio as precio,"
	consulta += "soporteserviciodetalle.Descuento as descuento,"
	consulta += "soporteserviciodetalle.Montodescuento as montodescuento,"
	consulta += "soporteserviciodetalle.Sigratis as sigratis,"
	consulta += "soporteserviciodetalle.Subtotal as subtotal,"
	consulta += "soporteserviciodetalle.Subtotaldescuento  as subtotaldescuento,"
	consulta += "soporteserviciodetalle.Pagina as pagina ,"
	consulta += "soporteserviciodetalle.Bodega as bodega,"
	consulta += "soporteserviciodetalle.Fecha as fecha,"
	consulta += "soporteserviciodetalle.Nombreservicio,"
	consulta += "soporteserviciodetalle.Unidadservicio,"
	consulta += "soporteserviciodetalle.Codigoservicio"
	consulta += " from soporteserviciodetalle "
	consulta += " where soporteserviciodetalle.codigo=$1"
	log.Println(consulta)
	return consulta
}

// INICIA SOPORTE SERVICIO LISTA
func SoporteservicioLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/soporteservicio/soporteservicioLista.html")
	log.Println("Error soporteservicio 0")
	var consulta string
	var miperiodo = periodoSesion(r)
	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre,centro.nombre as CentroNombre,soporteservicio.neto,soporteservicio.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	consulta += " FROM soporteservicio "
	consulta += " inner join tercero on tercero.codigo=soporteservicio.tercero "
	consulta += " inner join centro on centro.codigo=soporteservicio.centro "
	consulta += " inner join almacenista on almacenista.codigo=soporteservicio.almacenista "
	consulta += " where  extract(year from fecha)=" + miperiodo + " "
	consulta += " ORDER BY cast(soporteservicio.codigo as integer) ASC"

	db := dbConn()
	res := []soporteservicioLista{}

	err := db.Select(&res, consulta)

	if err != nil {
		fmt.Println(err)
		return
	}

	varmap := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
	}
	log.Println("Error soporteservicio888")
	tmp.Execute(w, varmap)
}

// INICIA SOPORTE SERVICIO NUEVO
func SoporteservicioNuevo(w http.ResponseWriter, r *http.Request) {

	// TRAE COPIA DE EDITAR
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio soporteservicio editar" + Codigo)

	db := dbConn()
	v := soporteservicio{}
	det := []soporteserviciodetalleeditar{}
	if Codigo == "False" {

	} else {

		err := db.Get(&v, "SELECT * FROM soporteservicio where codigo=$1", Codigo)
		if err != nil {
			log.Fatalln(err)
		}

		v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
		v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
		v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
		v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
		v.CentroDetalle = TraerCentroConsulta(v.Centro)

		err2 := db.Select(&det, SoporteservicioConsultaDetalle(), Codigo)
		if err2 != nil {
			fmt.Println(err2)
		}
	}

	parametros := map[string]interface{}{
		"resolucionsoporte": ListaResolucionsoporte(),
		"codigo":            Codigo,
		"soporteservicio":   v,
		"almacenista":       PrimerAlmacenista(),
		"detalle":           det,
		"hosting":           ruta,
		"miperiodo":         periodoSesion(r),
		"retfte":            TraerParametrosInventario().Soportecuentaporcentajeretfte,
	}
	//TERMINA TRAE COPIA DE EDITAR

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/soporteservicio/soporteservicioNuevo.html",
		"vista/autocompleta/autocompletaPlandecuentaempresa.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/soporteservicio/soporteservicioScript.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error soporteservicio nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA INSERTAR COMPROBANTE DE SOPORTE SERVICIO
func InsertaDetalleComprobanteSoporteservicio(miFilaComprobante comprobantedetalle, miComprobante comprobante, miSoporteservicio soporteservicio) {
	db := dbConn()

	// traer tercero
	miTercero := tercero{}
	err3 := db.Get(&miTercero, "SELECT * FROM tercero where codigo=$1", miSoporteservicio.Tercero)
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
		miSoporteservicio.Tercero,
		miSoporteservicio.Centro,
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

// INICIA SOPORTE SERVICIO INSERTAR AJAX
func SoporteservicioAgregar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var tempSoporteservicio soporteservicio

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// carga informacion de la SOPORTE SERVICIO
	err = json.Unmarshal(b, &tempSoporteservicio)
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
	if tempSoporteservicio.Accion == "Nuevo" {
		log.Println("Resolucion " + tempSoporteservicio.Resolucionsoporte)
		Codigoactual = Numerosoporte(tempSoporteservicio.Resolucionsoporte)
		tempSoporteservicio.Codigo = Codigoactual
	} else {
		Codigoactual = tempSoporteservicio.Codigo
	}

	if tempSoporteservicio.Accion == "Actualizar" {

		// borra detalle anterior
		delForm, err := db.Prepare("DELETE from soporteserviciodetalle WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(tempSoporteservicio.Codigo)

		// borra cabecera anterior
		delForm1, err := db.Prepare("DELETE from soporteservicio WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm1.Exec(tempSoporteservicio.Codigo)
	}

	// INSERTA DETALLE
	for i, x := range tempSoporteservicio.Detalle {
		var a = i
		var q string
		q = "insert into soporteserviciodetalle ("
		q += "Id,"
		q += "Codigo,"
		q += "Fila,"
		q += "Cantidad,"
		q += "Precio,"
		q += "Subtotal,"
		q += "Pagina,"
		q += "Bodega,"
		q += "Descuento,"
		q += "Montodescuento,"
		q += "Sigratis,"
		q += "Subtotaldescuento,"
		q += "Tipo,"
		q += "Fecha,"
		q += "Nombreservicio,"
		q += "Unidadservicio,"
		q += "CodigoServicio"
		q += " ) values("
		q += parametros(17)
		q += ")"

		log.Println("Cadena SQL " + q)
		insForm, err := db.Prepare(q)
		if err != nil {
			panic(err.Error())
		}

		// TERMINA SOPORTE SERVICIO GRABAR INSERTAR
		_, err = insForm.Exec(
			x.Id,
			Codigoactual,
			x.Fila,
			Quitacoma(x.Cantidad),
			Quitacoma(x.Precio),
			Quitacoma(x.Subtotal),
			x.Pagina,
			x.Bodega,
			Quitacoma(x.Descuento),
			Quitacoma(x.Montodescuento),
			x.Sigratis,
			Quitacoma(x.Subtotaldescuento),
			x.Tipo,
			x.Fecha,
			Mayuscula(x.Nombreservicio),
			Titulo(x.Unidadservicio),
			x.Codigoservicio)
		if err != nil {
			panic(err)
		}

		log.Println("Insertar Producto \n", x.Nombreservicio, a)
	}

	// INICIA INSERTAR SOPORTE SERVICIOS
	log.Println("Got %s age %s club %s\n", tempSoporteservicio.Codigo, tempSoporteservicio.Tercero, tempSoporteservicio.Subtotal)
	var q string
	q = "insert into soporteservicio ("
	q += "Resolucionsoporte,"
	q += "Codigo,"
	q += "Fecha,"
	q += "Vence,"
	q += "Hora,"
	q += "Descuento,"
	q += "Subtotaldescuento19,"
	q += "Subtotaldescuento5,"
	q += "Subtotaldescuento0,"
	q += "Subtotal,"
	q += "Subtotal19,"
	q += "Subtotal5,"
	q += "Subtotal0,"
	q += "Subtotaliva19,"
	q += "Subtotaliva5,"
	q += "Subtotaliva0,"
	q += "Subtotalbase19,"
	q += "Subtotalbase5,"
	q += "Subtotalbase0,"
	q += "TotalIva,"
	q += "Total,"
	q += "Neto,"
	q += "Items,"
	q += "PorcentajeRetencionFuente,"
	q += "TotalRetencionFuente,"
	q += "PorcentajeRetencionIca,"
	q += "TotalRetencionIca,"
	q += "Formadepago,"
	q += "Mediodepago,"
	q += "Tercero,"
	q += "Almacenista,"
	q += "Pedidosoporteservicio,"
	q += "Centro,"
	q += "Tipo,"
	q += "Soporteserviciocuenta0,"
	q += "Soporteservicionombre0,"
	q += "Soporteserviciocuentaretfte,"
	q += "Soporteservicionombreretfte"
	q += ") values("
	q += parametros(38)
	q += ")"

	log.Println("Cadena SQL " + q)
	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	layout := "2006-01-02"
	log.Println("Hora", tempSoporteservicio.Fecha.Format("02/01/2006"))
	currentTime := time.Now()

	_, err = insForm.Exec(
		tempSoporteservicio.Resolucionsoporte,
		tempSoporteservicio.Codigo,
		tempSoporteservicio.Fecha.Format(layout),
		tempSoporteservicio.Vence.Format(layout),
		currentTime.Format("3:4:5 PM"),
		Quitacoma(tempSoporteservicio.Descuento),
		Quitacoma(tempSoporteservicio.Subtotaldescuento19),
		Quitacoma(tempSoporteservicio.Subtotaldescuento5),
		Quitacoma(tempSoporteservicio.Subtotaldescuento0),
		Quitacoma(tempSoporteservicio.Subtotal),
		Quitacoma(tempSoporteservicio.Subtotal19),
		Quitacoma(tempSoporteservicio.Subtotal5),
		Quitacoma(tempSoporteservicio.Subtotal0),
		Quitacoma(tempSoporteservicio.Subtotaliva19),
		Quitacoma(tempSoporteservicio.Subtotaliva5),
		Quitacoma(tempSoporteservicio.Subtotaliva0),
		Quitacoma(tempSoporteservicio.Subtotalbase19),
		Quitacoma(tempSoporteservicio.Subtotalbase5),
		Quitacoma(tempSoporteservicio.Subtotalbase0),
		Quitacoma(tempSoporteservicio.TotalIva),
		Quitacoma(tempSoporteservicio.Total),
		Quitacoma(tempSoporteservicio.Neto),
		tempSoporteservicio.Items,
		Quitacoma(tempSoporteservicio.PorcentajeRetencionFuente),
		Quitacoma(tempSoporteservicio.TotalRetencionFuente),
		Quitacoma(tempSoporteservicio.PorcentajeRetencionIca),
		Quitacoma(tempSoporteservicio.TotalRetencionIca),
		tempSoporteservicio.Formadepago,
		tempSoporteservicio.Mediodepago,
		tempSoporteservicio.Tercero,
		tempSoporteservicio.Almacenista,
		tempSoporteservicio.Pedidosoporteservicio,
		tempSoporteservicio.Centro,
		tempSoporteservicio.Tipo,
		tempSoporteservicio.Soporteserviciocuenta0,
		tempSoporteservicio.Soporteservicionombre0,
		tempSoporteservicio.Soporteserviciocuentaretfte,
		tempSoporteservicio.Soporteservicionombreretfte)
	if err != nil {
		panic(err)
	}

	// INSERTAR COMPROBANTE CONTABILIDAD
	var tempComprobante comprobante
	var tempComprobanteDetalle comprobantedetalle
	tempComprobante.Documento = "10"
	tempComprobante.Numero = tempSoporteservicio.Codigo
	tempComprobante.Fecha = tempSoporteservicio.Fecha
	tempComprobante.Fechaconsignacion = tempSoporteservicio.Fecha
	tempComprobante.Debito = tempSoporteservicio.Neto + ".00"
	tempComprobante.Credito = tempSoporteservicio.Neto + ".00"
	tempComprobante.Periodo = ""
	tempComprobante.Licencia = ""
	tempComprobante.Usuario = ""
	tempComprobante.Estado = ""

	// borra detalle anterior
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

	// INSERTAR CUENTA DEBITO SOPORTE SERVICIO 0%
	if tempSoporteservicio.Subtotal0 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempSoporteservicio.Subtotal0) - Flotante(tempSoporteservicio.Subtotaldescuento0)
		tempComprobanteDetalle.Cuenta = tempSoporteservicio.Soporteserviciocuenta0
		tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(tempSoporteservicio.Subtotal0) - Flotante(tempSoporteservicio.Subtotaldescuento0))
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteSoporteservicio(tempComprobanteDetalle, tempComprobante, tempSoporteservicio)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA CREDITO DESCUENTO
	//if (tempSoporteservicio.Descuento!="0")	{
	//	fila=fila+1
	//	tempComprobanteDetalle.Fila=strconv.Itoa(fila)
	//	totalCredito+=Flotante(tempSoporteservicio.Descuento)
	//	tempComprobanteDetalle.Cuenta  = parametrosinventario.Soporteserviciocuentadescuento
	//	tempComprobanteDetalle.Debito = ""
	//	tempComprobanteDetalle.Credito = tempSoporteservicio.Descuento
	//	InsertaDetalleComprobanteSoporteservicio(tempComprobanteDetalle,tempComprobante,tempSoporteservicio)
	//	log.Println("credito linea" + fmt.Sprintf("%.2f",totalCredito))
	//}

	// INSERTAR CUENTA CREDITO RET. FTE.
	if tempSoporteservicio.TotalRetencionFuente != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempSoporteservicio.TotalRetencionFuente)
		tempComprobanteDetalle.Cuenta = tempSoporteservicio.Soporteserviciocuentaretfte
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = tempSoporteservicio.TotalRetencionFuente
		InsertaDetalleComprobanteSoporteservicio(tempComprobanteDetalle, tempComprobante, tempSoporteservicio)
		log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito))
	}

	// INSERTAR CUENTA CREDITO RET. ICA.
	if tempSoporteservicio.TotalRetencionIca != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempSoporteservicio.TotalRetencionIca)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Soporteserviciocuentaretica
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = tempSoporteservicio.TotalRetencionIca
		InsertaDetalleComprobanteSoporteservicio(tempComprobanteDetalle, tempComprobante, tempSoporteservicio)
		log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito))
	}

	// INSERTAR CUENTA CREDITO PROVEEDOR
	if tempSoporteservicio.Neto != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempSoporteservicio.Neto)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Soporteserviciocuentaproveedor
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = tempSoporteservicio.Neto
		InsertaDetalleComprobanteSoporteservicio(tempComprobanteDetalle, tempComprobante, tempSoporteservicio)
		log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito))
	}

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
	q += " ) values("
	q += parametros(10)
	q += " )"

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

// INICIA SOPORTE SERVICIO EXISTE
func SoporteservicioExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM soporteservicio  WHERE codigo=$1", Codigo)
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

// INICIA SOPORTE SERVICIO EDITAR
func SoporteservicioEditar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio soporteservicio editar" + Codigo)
	db := dbConn()

	// traer SOPORTE SERVICIO
	v := soporteservicio{}
	err := db.Get(&v, "SELECT * FROM soporteservicio where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []soporteserviciodetalleeditar{}

	err2 := db.Select(&det, SoporteservicioConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
	v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
	v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
	v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
	v.CentroDetalle = TraerCentroConsulta(v.Centro)

	parametros := map[string]interface{}{
		"resolucionsoporte": ListaResolucionsoporte(),
		"soporteservicio":   v,
		"detalle":           det,
		"hosting":           ruta,
		"miperiodo":         periodoSesion(r),
		"retfte":            TraerParametrosInventario().Soportecuentaporcentajeretfte,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/soporteservicio/soporteservicioEditar.html",
		"vista/autocompleta/autocompletaPlandecuentaempresa.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/soporteservicio/soporteservicioScript.html")
	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error soporteservicio nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA SOPORTE SERVICIO BORRAR
func SoporteservicioBorrar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio soporteservicio editar" + Codigo)

	db := dbConn()

	// traer SOPORTE SERVICIO
	v := soporteservicio{}
	err := db.Get(&v, "SELECT * FROM soporteservicio where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []soporteserviciodetalleeditar{}
	err2 := db.Select(&det, SoporteservicioConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
	v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
	v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
	v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
	v.CentroDetalle = TraerCentroConsulta(v.Centro)

	parametros := map[string]interface{}{
		"resolucionsoporte": ListaResolucionsoporte(),
		"soporteservicio":   v,
		"detalle":           det,
		"hosting":           ruta,
		"retfte":            TraerParametrosInventario().Soportecuentaporcentajeretfte,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/soporteservicio/soporteservicioBorrar.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/soporteservicio/soporteservicioScript.html")

	log.Println("Error soporteservicio nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA SOPORTE SERVICIO ELIMINAR
func SoporteservicioEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	codigo := mux.Vars(r)["codigo"]

	// borrar SOPORTE SERVICIO
	delForm, err := db.Prepare("DELETE from soporteservicio WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(codigo)

	// borar detalle
	delForm1, err := db.Prepare("DELETE from soporteserviciodetalle WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm1.Exec(codigo)

	// borar detalle invenario
	Borrarinventario(codigo, "Soporteservicio")

	// borra detalle anterior
	delForm, err = db.Prepare("DELETE from comprobantedetalle WHERE documento=$1 and numero=$2")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec("10", codigo)

	// borra cabecera anterior

	delForm1, err = db.Prepare("DELETE from comprobante WHERE documento=$1 and numero=$2 ")
	if err != nil {
		panic(err.Error())
	}
	delForm1.Exec("10", codigo)

	log.Println("Registro Eliminado" + codigo)
	http.Redirect(w, r, "/SoporteservicioLista", 301)
}

// TRAER PEDIDO SOPORTE SERVICIO
func Datospedidosoporteservicio(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	Tercero := mux.Vars(r)["tercero"]
	log.Println("inicio pedido editar" + Codigo)
	db := dbConn()
	var res []pedidosoporteservicio

	// traer PEDIDO
	v := pedidosoporteservicio{}
	err := db.Get(&v, "SELECT * FROM pedidosoporteservicio where codigo=$1 and Tercero=$2", Codigo, Tercero)
	var valida bool
	valida = true

	switch err {
	case nil:
		log.Printf("pedido existe: %+v\n", v)
	case sql.ErrNoRows:
		log.Println("pedidosoporteservicio NO Existe")
		valida = false
	default:
		log.Printf("pedidosoporteservicio error: %s\n", err)
	}
	det := []pedidosoporteserviciodetalleeditar{}
	t := tercero{}

	// trae datos si existe pedido
	if valida == true {
		err2 := db.Select(&det, PedidosoporteservicioConsultaDetalle(), Codigo)
		if err2 != nil {
			fmt.Println(err2)
		}
		v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
		v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
		v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
		v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
		v.CentroDetalle = TraerCentroConsulta(v.Centro)

		v.DetalleEditar = det
		res = append(res, v)
	}

	log.Println("codigo nombre99" + t.Codigo + t.Nombre)
	data, err := json.Marshal(res)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// INICIA SOPORTE SERVICIO PDF
func SoporteservicioPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]

	// traer SOPORTE SERVICIO
	miSoporteservicio := soporteservicio{}
	err := db.Get(&miSoporteservicio, "SELECT * FROM soporteservicio where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}

	// traer detalle
	miDetalle := []soporteserviciodetalleeditar{}
	err2 := db.Select(&miDetalle, SoporteservicioConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	// traer tercero
	miTercero := tercero{}
	err3 := db.Get(&miTercero, "SELECT * FROM tercero where codigo=$1", miSoporteservicio.Tercero)
	if err3 != nil {
		log.Fatalln(err3)
	}

	// traer almacenista
	miAlmacenista := almacenista{}
	err4 := db.Get(&miAlmacenista, "SELECT * FROM almacenista where codigo=$1", miSoporteservicio.Almacenista)
	if err4 != nil {
		log.Fatalln(err4)
	}

	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	var re Resolucionsoporte = TraerResolucionsoporte(miSoporteservicio.Resolucionsoporte)
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

		// RESOLUCION
		pdf.SetFont("Arial", "", 8)
		pdf.SetY(20)
		pdf.SetX(80)
		pdf.CellFormat(190, 10, "Resolucion No. "+re.Numero, "0", 0, "C",
			false, 0, "")
		pdf.Ln(3)
		pdf.SetX(80)
		pdf.CellFormat(190, 10, "Del Numero "+re.Prefijo+" "+re.NumeroInicial+" al "+re.Prefijo+" "+re.NumeroFinal, "0", 0, "C",
			false, 0, "")
		pdf.Ln(3)
		pdf.SetX(80)
		pdf.CellFormat(190, 10, "Vigencia del "+re.FechaInicial.Format("02/01/2006")+" al "+re.FechaFinal.Format("02/01/2006"), "0", 0, "C",
			false, 0, "")

		// SOPORTE NUMERO
		pdf.Ln(4)
		pdf.SetX(80)
		pdf.SetFont("Arial", "", 10)
		pdf.CellFormat(190, 10, "DOCUMENTO SOPORTE EN", "0", 0, "C",
			false, 0, "")
		pdf.Ln(4)
		pdf.SetX(80)
		pdf.CellFormat(190, 10, "ADQUISICIONES EFECTUADAS A", "0", 0, "C",
			false, 0, "")
		pdf.Ln(4)
		pdf.SetX(80)
		pdf.CellFormat(190, 10, "SUJETOS NO OBLIGADOS A EXPEDIR", "0", 0, "C",
			false, 0, "")
		pdf.Ln(4)
		pdf.SetX(80)
		pdf.CellFormat(190, 10, "FACTURA DE VENTA No. "+re.Prefijo+" "+Codigo, "0", 0, "C",
			false, 0, "")
		pdf.Ln(6)
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
	pdf.AddPage()
	pdf.Ln(1)
	SoporteservicioCabecera(pdf, miTercero, miSoporteservicio, miAlmacenista)

	var filas = len(miDetalle)
	// UNA PAGINA
	if filas <= 32 {
		for i, miFila := range miDetalle {
			var a = i + 1
			SoporteservicioFilaDetalle(pdf, miFila, a)
		}
		SoporteservicioPieDePagina(pdf, miTercero, miSoporteservicio)
	} else {
		// mas de 32 y menos de 73
		if filas > 32 && filas <= 73 {
			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 41 {
					SoporteservicioFilaDetalle(pdf, miFila, a)
				}
			}
			SoporteservicioLinea(pdf)
			// segunda pagina
			pdf.AddPage()
			SoporteservicioCabecera(pdf, miTercero, miSoporteservicio, miAlmacenista)
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 41 {
					SoporteservicioFilaDetalle(pdf, miFila, a)
				}
			}

			SoporteservicioPieDePagina(pdf, miTercero, miSoporteservicio)
		} else {
			// mas de 80

			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 41 {
					SoporteservicioFilaDetalle(pdf, miFila, a)
				}
			}
			SoporteservicioLinea(pdf)
			pdf.AddPage()
			SoporteservicioCabecera(pdf, miTercero, miSoporteservicio, miAlmacenista)
			// segunda pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 41 && a <= 82 {
					SoporteservicioFilaDetalle(pdf, miFila, a)
				}
			}
			SoporteservicioLinea(pdf)
			pdf.AddPage()
			SoporteservicioCabecera(pdf, miTercero, miSoporteservicio, miAlmacenista)
			// tercera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 82 {
					SoporteservicioFilaDetalle(pdf, miFila, a)
				}
			}

			SoporteservicioPieDePagina(pdf, miTercero, miSoporteservicio)
		}
	}

	// genera pdf
	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

func SoporteservicioCabecera(pdf *gofpdf.Fpdf, miTercero tercero, miSoporteservicio soporteservicio, miAlmacenista almacenista) {
	ene := pdf.UnicodeTranslatorFromDescriptor("")
	// RELLENO TITULO
	pdf.SetFont("Arial", "", 10)
	pdf.SetY(44)
	pdf.SetX(20)
	pdf.SetFillColor(224, 231, 239)
	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(8)
	pdf.SetX(20)
	pdf.CellFormat(184, 6, "DATOS DEL PROVEEDOR", "0", 0,
		"L", true, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(94, 6, "LUGAR DE ENTREGA O SERVICIO", "0", 0,
		"L", true, 0, "")
	pdf.Ln(8)
	pdf.SetX(20)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(40, 4, "Nit. No.", "", 0,
		"L", false, 0, "")
	pdf.SetX(40)
	pdf.CellFormat(40, 4, Coma(miTercero.Codigo)+" - "+miTercero.Dv, "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Nit. No. ", "", 0,
		"L", false, 0, "")
	pdf.SetX(130)
	pdf.CellFormat(40, 4, Coma(miTercero.Codigo)+" - "+miTercero.Dv, "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Nombre", "", 0,
		"L", false, 0, "")
	pdf.SetX(40)
	pdf.CellFormat(40, 4, ene(miTercero.Nombre), "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Nombre", "", 0,
		"L", false, 0, "")
	pdf.SetX(130)
	pdf.CellFormat(40, 4, ene(miTercero.Nombre), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Direccion", "", 0,
		"L", false, 0, "")
	pdf.SetX(40)
	pdf.CellFormat(40, 4, ene(miTercero.Direccion), "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Direccion", "", 0,
		"L", false, 0, "")
	pdf.SetX(130)
	pdf.CellFormat(40, 4, ene(miTercero.Direccion), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Forma de Pago", "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, Titulo(TraerFormadepago(miSoporteservicio.Formadepago)), "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Fecha de Expedicion", "", 0,
		"L", false, 0, "")
	pdf.SetX(150)
	pdf.CellFormat(40, 4, miSoporteservicio.Fecha.Format("02/01/2006")+" "+Titulo(miSoporteservicio.Hora), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)

	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Medio de Pago", "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, TraerMediodepago(miSoporteservicio.Mediodepago), "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Fecha de Vencimiento", "", 0,
		"L", false, 0, "")
	pdf.SetX(150)
	pdf.CellFormat(40, 4, miSoporteservicio.Vence.Format("02/01/2006"), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Pedido No.", "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, miSoporteservicio.Pedidosoporteservicio, "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Almacenista", "", 0,
		"L", false, 0, "")
	pdf.SetX(135)
	pdf.CellFormat(40, 4, ene(miAlmacenista.Nombre), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)

	// RELLENO TITULO
	pdf.SetFont("Arial", "", 10)
	pdf.SetY(86)
	pdf.SetX(20)
	pdf.SetFillColor(224, 231, 239)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetX(20)

	pdf.CellFormat(184, 6, "ITEM", "0", 0,
		"L", true, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(190, 6, "CODIGO", "0", 0,
		"L", false, 0, "")
	pdf.SetX(55)
	pdf.CellFormat(190, 6, "DESCRIPCION", "0", 0,
		"L", false, 0, "")
	pdf.SetX(125)
	pdf.CellFormat(40, 6, "UNIDAD", "0", 0,
		"L", false, 0, "")
	pdf.SetX(141)
	pdf.CellFormat(40, 6, "CANTIDAD", "0", 0,
		"L", false, 0, "")
	pdf.SetX(162)
	pdf.CellFormat(40, 6, "P. UNITARIO", "0", 0,
		"L", false, 0, "")
	pdf.SetX(190)
	pdf.CellFormat(40, 6, "TOTAL", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}
func SoporteservicioFilaDetalle(pdf *gofpdf.Fpdf, miFila soporteserviciodetalleeditar, a int) {
	ene := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetX(21)
	var yinicial float64
	yinicial = pdf.GetY()
	pdf.CellFormat(183, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(40, 4, Subcadena(miFila.Codigoservicio, 0, 4), "", 0,
		"L", false, 0, "")
	var y float64
	y = pdf.GetY()
	pdf.SetX(42)
	pdf.MultiCell(85, 4, ene(Mayuscula(miFila.Nombreservicio)), "", "L", false)
	var yfinal float64
	yfinal = pdf.GetY()
	pdf.SetY(y)
	pdf.SetX(100)
	pdf.CellFormat(40, 4, ene(Titulo(Subcadena(miFila.Unidadservicio, 0, 6))), "", 0,
		"R", false, 0, "")
	pdf.SetX(118)
	pdf.CellFormat(40, 4, Coma(miFila.Cantidad), "", 0,
		"R", false, 0, "")
	pdf.SetX(141)
	pdf.CellFormat(40, 4, Coma(miFila.Precio), "", 0,
		"R", false, 0, "")
	pdf.SetX(164)
	pdf.CellFormat(40, 4, Coma(miFila.Subtotal), "", 0,
		"R", false, 0, "")
	pdf.Ln(yfinal - yinicial + 3)
}

func SoporteservicioPieDePagina(pdf *gofpdf.Fpdf, miTercero tercero, miSoporteservicio soporteservicio) {

	Totalletras, err := IntLetra(Cadenaentero(miSoporteservicio.Neto))
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

	pdf.SetX(1)
	pdf.CellFormat(190, 4, "Esta factura es un titulo valor para su emisor o poseedor en caso", "0", 0,
		"C", false, 0, "")
	pdf.Ln(3)
	pdf.SetX(1)
	pdf.CellFormat(190, 4, "de endoso, presta merito ejecutivo y cumple con los requisitos del", "0", 0,
		"C", false, 0, "")
	pdf.Ln(3)
	pdf.SetX(1)
	pdf.CellFormat(190, 4, "art. 617 del E. T. y art. 773 y 774 del Codigo de Comercio", "0", 0,
		"C", false, 0, "")
	pdf.Ln(6)
	pdf.SetX(1)
	pdf.Line(55, 249, 140, 249)
	pdf.Ln(8)
	pdf.SetX(80)
	pdf.CellFormat(40, 4, "A C E P T A D A ", "0", 0, "C",
		false, 0, "")

	// PRESENTA DATOS CON VALORES //
	pdf.SetFont("Arial", "", 9)
	var separador float64
	var altoseparador float64
	separador = 243
	altoseparador = -4

	// INICIA DATOS FACTURA
	if miSoporteservicio.Neto != "0" {
		pdf.SetFont("Arial", "", 9)
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "TOTAL", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miSoporteservicio.Neto), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miSoporteservicio.TotalRetencionIca != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "RET. ICA.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, miSoporteservicio.PorcentajeRetencionIca+"%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miSoporteservicio.TotalRetencionIca), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miSoporteservicio.TotalRetencionFuente != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "RET. FTE.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, miSoporteservicio.PorcentajeRetencionFuente+"%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miSoporteservicio.TotalRetencionFuente), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miSoporteservicio.Subtotal0 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "SUBTOTAL", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miSoporteservicio.Subtotal0), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	pdf.Image(imageFile("QR.jpg"), 20, 229, 25, 0, false,
		"", 0, "")
	pdf.SetFont("Arial", "", 8)
	pdf.SetY(255)
	pdf.SetX(20)
	pdf.CellFormat(184, 4, "Cufexxxxxxxxxxxxxxxxxxxxxxxxxxxxx", "", 0,
		"L", false, 0, "")
}

func SoporteservicioLinea(pdf *gofpdf.Fpdf) {
	pdf.SetY(259)
	pdf.Line(20, 259, 204, 259)
}

// INICIA SOPORTE SERVICIO TODOS PDF
func SoporteservicioTodosCabecera(pdf *gofpdf.Fpdf) {
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
	pdf.CellFormat(190, 6, "Fecha", "0", 0,
		"L", false, 0, "")
	pdf.SetX(75)
	pdf.CellFormat(190, 6, "Proveedor", "0", 0,
		"L", false, 0, "")
	pdf.SetX(130)
	pdf.CellFormat(190, 6, "Almacenista", "0", 0,
		"L", false, 0, "")
	pdf.SetX(193)
	pdf.CellFormat(190, 6, "Total", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}

func SoporteservicioTodosDetalle(pdf *gofpdf.Fpdf, miFila soporteservicioLista, a int) {
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(21)
	pdf.CellFormat(181, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(32)
	pdf.CellFormat(40, 4, Subcadena(miFila.Codigo, 0, 12), "", 0,
		"L", false, 0, "")
	pdf.SetX(55)
	pdf.CellFormat(40, 4, miFila.Fecha.Format("02/01/2006"), "", 0, "L", false, 0, "")
	pdf.SetX(75)
	pdf.CellFormat(40, 4, Subcadena(miFila.TerceroNombre, 0, 29),
		"", 0, "L", false, 0, "")
	pdf.SetX(130)
	pdf.CellFormat(40, 4, Subcadena(miFila.AlmacenistaNombre, 0, 31), "", 0,
		"L", false, 0, "")
	pdf.SetX(164)
	pdf.CellFormat(40, 4, Coma(miFila.Neto), "", 0,
		"R", false, 0, "")
	pdf.Ln(4)
}

func SoporteservicioTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string
	consulta = "  SELECT Almacenista.nombre as AlmacenistaNombre, " +
		"centro.nombre as CentroNombre,soporteservicio.neto," +
		"soporteservicio.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	consulta += " FROM soporteservicio "
	consulta += " inner join tercero on tercero.codigo=soporteservicio.tercero "
	consulta += " inner join centro on centro.codigo=soporteservicio.centro "
	consulta += " inner join almacenista on almacenista.codigo=soporteservicio.almacenista "
	consulta += " ORDER BY cast(soporteservicio.codigo as integer) ASC"

	t := []soporteservicioLista{}
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
		pdf.CellFormat(190, 10, "DATOS SOPORTE SERVICIOS", "0", 0,
			"C", false, 0, "")
		pdf.Ln(10)
	})

	pdf.SetFooterFunc(func() {
		pdf.SetTextColor(0, 0, 0)
		pdf.SetY(253)
		pdf.SetFont("Arial", "", 9)
		pdf.SetX(20)

		// LINEA
		pdf.Line(20, 260, 204, 260)
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

	SoporteservicioTodosCabecera(pdf)
	// tercera pagina

	for i, miFila := range t {
		var a = i + 1
		if math.Mod(float64(a), 49) == 0 {
			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			CotizacionTodosCabecera(pdf)
		}
		SoporteservicioTodosDetalle(pdf, miFila, a)
	}

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

// TERMINA VENTA TODOS PDF

// SOPORTE SERVICIOS EXCEL
func SoporteservicioExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string
	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre,centro.nombre " +
		"as CentroNombre,soporteservicio.neto,soporteservicio.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	consulta += " FROM soporteservicio "
	consulta += " inner join tercero on tercero.codigo=soporteservicio.tercero "
	consulta += " inner join centro on centro.codigo=soporteservicio.centro "
	consulta += " inner join almacenista on almacenista.codigo=soporteservicio.almacenista "
	consulta += " ORDER BY cast(soporteservicio.codigo as integer) ASC"
	t := []soporteservicioLista{}
	err := db.Select(&t, consulta)
	if err != nil {
		fmt.Println(err)
		return
	}

	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	f := excelize.NewFile()

	// FUNCION ANCHO DE LA COLUMNA
	if err = f.SetColWidth("Sheet1", "A", "A", 13); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.SetColWidth("Sheet1", "B", "B", 13); err != nil {
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
	f.SetCellValue("Sheet1", "A9", "LISTADO DE SOPORTE SERVICIOS")
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
	f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel), "Fecha")
	f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel), "Proveedor")
	f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel), "Almacenista")
	f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel), "Total")

	f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel), "A"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel), "B"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel), "C"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel), "D"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "E"+strconv.Itoa(filaExcel), "E"+strconv.Itoa(filaExcel), estiloCabecera)

	filaExcel++

	for i, miFila := range t {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel+i), Flotante(miFila.Codigo))
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel+i), miFila.Fecha.Format("02/01/2006"))
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel+i), miFila.TerceroNombre)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel+i), miFila.AlmacenistaNombre)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel+i), Flotante(miFila.Neto))

		f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel+i), "A"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel+i), "B"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel+i), "D"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel+i), "E"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "E"+strconv.Itoa(filaExcel+i), "F"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
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

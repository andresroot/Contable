package main

// INICIA DEVOLUCION FACTURA GASTO IMPORTAR PAQUETES
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

// INICIA DEVOLUCION FACTURA GASTO ESTRUCTURA JSON
type devolucionfacturagastoJson struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Nombre string `json:"nombre"`
}

// INICIA DEVOLUCION FACTURA GASTO ESTRUCTURA
type devolucionfacturagastoLista struct {
	Codigo            string
	Fecha             time.Time
	Neto              string
	Tercero           string
	TerceroNombre     string
	CentroNombre      string
	AlmacenistaNombre string
}

// INICIA DEVOLUCION FACTURA GASTO ESTRUCTURA
type devolucionfacturagasto struct {
	Codigo                    string
	Fecha                     time.Time
	Vence                     time.Time
	Hora                      string
	Descuento                 string
	Subtotaldescuento19       string
	Subtotaldescuento5        string
	Subtotaldescuento0        string
	Subtotal                  string
	Subtotal19                string
	Subtotal5                 string
	Subtotal0                 string
	Subtotaliva19             string
	Subtotaliva5              string
	Subtotaliva0              string
	Subtotalbase19            string
	Subtotalbase5             string
	Subtotalbase0             string
	TotalIva                  string
	Total                     string
	PorcentajeRetencionFuente string
	TotalRetencionFuente      string
	PorcentajeRetencionIca    string
	TotalRetencionIca         string
	Neto                      string
	Items                     string
	Formadepago               string
	Mediodepago               string
	Tercero                   string
	Almacenista               string
	Accion                    string
	Detalle                   []devolucionfacturagastodetalle       `json:"Detalle"`
	DetalleEditar             []devolucionfacturagastodetalleeditar `json:"DetalleEditar"`
	TerceroDetalle            tercero
	Facturagasto              string
	Tipo                      string
	Centro                    string
	Facturagastocuenta0       string
	Facturagastonombre0       string
	Facturagastocuentaretfte  string
	Facturagastonombreretfte  string
	Facturagastocuentaiva     string
	Facturagastonombreiva     string
	Facturagastoporcentajeiva string
	FormadepagoDetalle        formadepago
	MediodepagoDetalle        mediodepago
	AlmacenistaDetalle        almacenista
	CentroDetalle             centro
}

// INICIA SOPORTE SERVICIODETALLE ESTRUCTURA
type devolucionfacturagastodetalle struct {
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

// INICIA DEVOLUCION FACTURA GASTO DETALLE EDITAR
type devolucionfacturagastodetalleeditar struct {
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

// INICIA DEVOLUCION FACTURA GASTO CONSULTA DETALLE
func DevolucionfacturagastoConsultaDetalle() string {
	var consulta = ""
	consulta = "select "
	consulta += "devolucionfacturagastodetalle.Id as id ,"
	consulta += "devolucionfacturagastodetalle.Codigo as codigo,"
	consulta += "devolucionfacturagastodetalle.Fila as fila,"
	consulta += "devolucionfacturagastodetalle.Cantidad as cantidad,"
	consulta += "devolucionfacturagastodetalle.Precio as precio,"
	consulta += "devolucionfacturagastodetalle.Descuento as descuento,"
	consulta += "devolucionfacturagastodetalle.Montodescuento as montodescuento,"
	consulta += "devolucionfacturagastodetalle.Sigratis as sigratis,"
	consulta += "devolucionfacturagastodetalle.Subtotal as subtotal,"
	consulta += "devolucionfacturagastodetalle.Subtotaldescuento  as subtotaldescuento,"
	consulta += "devolucionfacturagastodetalle.Pagina as pagina ,"
	consulta += "devolucionfacturagastodetalle.Bodega as bodega,"
	consulta += "devolucionfacturagastodetalle.Fecha as fecha,"
	consulta += "devolucionfacturagastodetalle.Nombreservicio,"
	consulta += "devolucionfacturagastodetalle.Unidadservicio,"
	consulta += "devolucionfacturagastodetalle.Codigoservicio"
	consulta += " from devolucionfacturagastodetalle "
	consulta += " where devolucionfacturagastodetalle.codigo=$1"
	log.Println(consulta)
	return consulta
}

// INICIA DEVOLUCION FACTURA GASTO LISTA
func DevolucionfacturagastoLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/devolucionfacturagasto/devolucionfacturagastoLista.html")
	log.Println("Error devolucionfacturagasto 0")
	var consulta string
	var miperiodo = periodoSesion(r)
	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre,centro.nombre as CentroNombre,devolucionfacturagasto.neto,devolucionfacturagasto.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	consulta += " FROM devolucionfacturagasto "
	consulta += " inner join tercero on tercero.codigo=devolucionfacturagasto.tercero "
	consulta += " inner join centro on centro.codigo=devolucionfacturagasto.centro "
	consulta += " inner join almacenista on almacenista.codigo=devolucionfacturagasto.almacenista "
	consulta += " where  extract(year from fecha)=" + miperiodo + " "

	consulta += " ORDER BY cast(devolucionfacturagasto.codigo as integer) ASC"

	db := dbConn()
	res := []devolucionfacturagastoLista{}

	err := db.Select(&res, consulta)

	if err != nil {
		fmt.Println(err)
		return
	}

	varmap := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
	}
	log.Println("Error facturagasto888")
	tmp.Execute(w, varmap)
}

// INICIA DEVOLUCION FACTURA GASTO NUEVO
func DevolucionfacturagastoNuevo(w http.ResponseWriter, r *http.Request) {

	// TRAE COPIA DE EDITAR
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio devolucionfacturagasto editar" + Codigo)

	db := dbConn()
	v := devolucionfacturagasto{}
	det := []devolucionfacturagastodetalleeditar{}
	if Codigo == "False" {

	} else {

		err := db.Get(&v, "SELECT * FROM devolucionfacturagasto where codigo=$1", Codigo)
		if err != nil {
			log.Fatalln(err)
		}

		v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
		v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
		v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
		v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
		v.CentroDetalle = TraerCentroConsulta(v.Centro)

		err2 := db.Select(&det, DevolucionfacturagastoConsultaDetalle(), Codigo)
		if err2 != nil {
			fmt.Println(err2)
		}
	}

	parametros := map[string]interface{}{
		"codigo":                 Codigo,
		"almacenista":            PrimerAlmacenista(),
		"devolucionfacturagasto": v,
		"detalle":                det,
		"hosting":                ruta,
		"miperiodo":              periodoSesion(r),
	}
	//TERMINA TRAE COPIA DE EDITAR

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/devolucionfacturagasto/devolucionfacturagastoNuevo.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/devolucionfacturagasto/devolucionfacturagastoScript.html",
		"vista/autocompleta/autocompletaPlandecuentaempresa.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error devolucionfacturagasto nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA INSERTAR COMPROBANTE DE DEVOLUCION FACTURA GASTO
func InsertaDetalleComprobanteDevolucionfacturagasto(miFilaComprobante comprobantedetalle, miComprobante comprobante, miDevolucionfacturagasto devolucionfacturagasto) {
	db := dbConn()

	// traer tercero
	miTercero := tercero{}
	err3 := db.Get(&miTercero, "SELECT * FROM tercero where codigo=$1", miDevolucionfacturagasto.Tercero)
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
	q += ") values("
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
		miDevolucionfacturagasto.Tercero,
		miDevolucionfacturagasto.Centro,
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

// INICIA DEVOLUCION FACTURA GASTO INSERTAR AJAX
func DevolucionfacturagastoAgregar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var tempDevolucionfacturagasto devolucionfacturagasto

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// carga informacion de la DEVOLUCION FACTURA GASTO
	err = json.Unmarshal(b, &tempDevolucionfacturagasto)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	parametrosinventario := configuracioninventario{}
	err = db.Get(&parametrosinventario, "SELECT * FROM configuracioninventario")
	if err != nil {
		panic(err.Error())
	}

	//var DocumentoContable string
	//DocumentoContable="25"

	if tempDevolucionfacturagasto.Accion == "Actualizar" {

		// borra detalle anterior
		delForm, err := db.Prepare("DELETE from devolucionfacturagastodetalle WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(tempDevolucionfacturagasto.Codigo)

		// borra cabecera anterior
		delForm1, err := db.Prepare("DELETE from devolucionfacturagasto WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm1.Exec(tempDevolucionfacturagasto.Codigo)
	}

	// INSERTA DETALLE
	for i, x := range tempDevolucionfacturagasto.Detalle {
		var a = i
		var q string
		q = "insert into devolucionfacturagastodetalle ("
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

		// TERMINA DEVOLUCION FACTURA GASTO GRABAR INSERTAR
		_, err = insForm.Exec(
			x.Id,
			x.Codigo,
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

	// INICIA INSERTAR DEVOLUCION FACTURA GASTO
	log.Println("Got %s age %s club %s\n", tempDevolucionfacturagasto.Codigo, tempDevolucionfacturagasto.Tercero, tempDevolucionfacturagasto.Subtotal)
	var q string
	q = "insert into devolucionfacturagasto ("
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
	q += "Facturagasto,"
	q += "Centro,"
	q += "Tipo,"
	q += "Facturagastocuenta0,"
	q += "Facturagastonombre0,"
	q += "Facturagastocuentaretfte,"
	q += "Facturagastonombreretfte,"
	q += "Facturagastocuentaiva,"
	q += "Facturagastonombreiva,"
	q += "Facturagastoporcentajeiva"
	q += " ) values("
	q += parametros(40)
	q += " ) "

	log.Println("Cadena SQL " + q)
	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	layout := "2006-01-02"
	log.Println("Hora", tempDevolucionfacturagasto.Fecha.Format("02/01/2006"))
	currentTime := time.Now()

	_, err = insForm.Exec(
		tempDevolucionfacturagasto.Codigo,
		tempDevolucionfacturagasto.Fecha.Format(layout),
		tempDevolucionfacturagasto.Vence.Format(layout),
		currentTime.Format("3:4:5 PM"),
		Quitacoma(tempDevolucionfacturagasto.Descuento),
		Quitacoma(tempDevolucionfacturagasto.Subtotaldescuento19),
		Quitacoma(tempDevolucionfacturagasto.Subtotaldescuento5),
		Quitacoma(tempDevolucionfacturagasto.Subtotaldescuento0),
		Quitacoma(tempDevolucionfacturagasto.Subtotal),
		Quitacoma(tempDevolucionfacturagasto.Subtotal19),
		Quitacoma(tempDevolucionfacturagasto.Subtotal5),
		Quitacoma(tempDevolucionfacturagasto.Subtotal0),
		Quitacoma(tempDevolucionfacturagasto.Subtotaliva19),
		Quitacoma(tempDevolucionfacturagasto.Subtotaliva5),
		Quitacoma(tempDevolucionfacturagasto.Subtotaliva0),
		Quitacoma(tempDevolucionfacturagasto.Subtotalbase19),
		Quitacoma(tempDevolucionfacturagasto.Subtotalbase5),
		Quitacoma(tempDevolucionfacturagasto.Subtotalbase0),
		Quitacoma(tempDevolucionfacturagasto.TotalIva),
		Quitacoma(tempDevolucionfacturagasto.Total),
		Quitacoma(tempDevolucionfacturagasto.Neto),
		tempDevolucionfacturagasto.Items,
		Quitacoma(tempDevolucionfacturagasto.PorcentajeRetencionFuente),
		Quitacoma(tempDevolucionfacturagasto.TotalRetencionFuente),
		Quitacoma(tempDevolucionfacturagasto.PorcentajeRetencionIca),
		Quitacoma(tempDevolucionfacturagasto.TotalRetencionIca),
		tempDevolucionfacturagasto.Formadepago,
		tempDevolucionfacturagasto.Mediodepago,
		tempDevolucionfacturagasto.Tercero,
		tempDevolucionfacturagasto.Almacenista,
		tempDevolucionfacturagasto.Facturagasto,
		tempDevolucionfacturagasto.Centro,
		tempDevolucionfacturagasto.Tipo,
		tempDevolucionfacturagasto.Facturagastocuenta0,
		tempDevolucionfacturagasto.Facturagastonombre0,
		tempDevolucionfacturagasto.Facturagastocuentaretfte,
		tempDevolucionfacturagasto.Facturagastonombreretfte,
		tempDevolucionfacturagasto.Facturagastocuentaiva,
		tempDevolucionfacturagasto.Facturagastonombreiva,
		Quitacoma(tempDevolucionfacturagasto.Facturagastoporcentajeiva))

	if err != nil {
		panic(err)
	}

	// INSERTAR COMPROBANTE CONTABILIDAD
	var tempComprobante comprobante
	var tempComprobanteDetalle comprobantedetalle
	//tempComprobante.Documento=DocumentoContable
	tempComprobante.Documento = "25"
	tempComprobante.Numero = tempDevolucionfacturagasto.Codigo
	tempComprobante.Fecha = tempDevolucionfacturagasto.Fecha
	tempComprobante.Fechaconsignacion = tempDevolucionfacturagasto.Fecha
	tempComprobante.Debito = tempDevolucionfacturagasto.Neto + ".00"
	tempComprobante.Credito = tempDevolucionfacturagasto.Neto + ".00"
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

	// INSERTAR CUENTA DEBITO PROVEEDOR
	if tempDevolucionfacturagasto.Neto != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempDevolucionfacturagasto.Neto)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Facturagastocuentaproveedor
		tempComprobanteDetalle.Debito = tempDevolucionfacturagasto.Neto
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteDevolucionfacturagasto(tempComprobanteDetalle, tempComprobante, tempDevolucionfacturagasto)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA CREDITO RET. FTE.
	if tempDevolucionfacturagasto.TotalRetencionFuente != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempDevolucionfacturagasto.TotalRetencionFuente)
		tempComprobanteDetalle.Cuenta = tempDevolucionfacturagasto.Facturagastocuentaretfte
		tempComprobanteDetalle.Debito = tempDevolucionfacturagasto.TotalRetencionFuente
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteDevolucionfacturagasto(tempComprobanteDetalle, tempComprobante, tempDevolucionfacturagasto)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA CREDITO RET. ICA.
	if tempDevolucionfacturagasto.TotalRetencionIca != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempDevolucionfacturagasto.TotalRetencionIca)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Facturagastocuentaretica
		tempComprobanteDetalle.Debito = tempDevolucionfacturagasto.TotalRetencionIca
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteDevolucionfacturagasto(tempComprobanteDetalle, tempComprobante, tempDevolucionfacturagasto)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR TOTAL IVA
	if tempDevolucionfacturagasto.TotalIva != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempDevolucionfacturagasto.TotalIva)
		tempComprobanteDetalle.Cuenta = tempDevolucionfacturagasto.Facturagastocuentaiva
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = tempDevolucionfacturagasto.TotalIva
		InsertaDetalleComprobanteDevolucionfacturagasto(tempComprobanteDetalle, tempComprobante, tempDevolucionfacturagasto)
		log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito))
	}

	// INSERTAR CUENTA DEBITO DEVOLUCION FACTURA GASTO 0%
	if tempDevolucionfacturagasto.Subtotal0 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempDevolucionfacturagasto.Subtotal0) - Flotante(tempDevolucionfacturagasto.Subtotaldescuento0)
		tempComprobanteDetalle.Cuenta = tempDevolucionfacturagasto.Facturagastocuenta0
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(tempDevolucionfacturagasto.Subtotal0) - Flotante(tempDevolucionfacturagasto.Subtotaldescuento0))
		InsertaDetalleComprobanteDevolucionfacturagasto(tempComprobanteDetalle, tempComprobante, tempDevolucionfacturagasto)
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

// INICIA DEVOLUCION FACTURA GASTO EXISTE
func DevolucionfacturagastoExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM devolucionfacturagasto  WHERE codigo=$1", Codigo)
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

// INICIA DEVOLUCION FACTURA GASTO EDITAR
func DevolucionfacturagastoEditar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio devolucionfacturagasto editar" + Codigo)
	db := dbConn()

	// TRAE DEVOLUCION FACTURA GASTO
	v := devolucionfacturagasto{}
	err := db.Get(&v, "SELECT * FROM devolucionfacturagasto where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []devolucionfacturagastodetalleeditar{}

	err2 := db.Select(&det, DevolucionfacturagastoConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
	v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
	v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
	v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
	v.CentroDetalle = TraerCentroConsulta(v.Centro)

	parametros := map[string]interface{}{
		"devolucionfacturagasto": v,
		"detalle":                det,
		"hosting":                ruta,
		"miperiodo":              periodoSesion(r),
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/devolucionfacturagasto/devolucionfacturagastoEditar.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/devolucionfacturagasto/devolucionfacturagastoScript.html",
		"vista/autocompleta/autocompletaPlandecuentaempresa.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error devolucionfacturagasto nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA DEVOLUCION FACTURA GASTO BORRAR
func DevolucionfacturagastoBorrar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio devolucionfacturagasto editar" + Codigo)
	db := dbConn()

	// TRAE DEVOLUCION FACTURA GASTO
	v := devolucionfacturagasto{}
	err := db.Get(&v, "SELECT * FROM devolucionfacturagasto where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []devolucionfacturagastodetalleeditar{}
	err2 := db.Select(&det, DevolucionfacturagastoConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
	v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
	v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
	v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
	v.CentroDetalle = TraerCentroConsulta(v.Centro)

	parametros := map[string]interface{}{
		"devolucionfacturagasto": v,
		"detalle":                det,
		"hosting":                ruta,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/devolucionfacturagasto/devolucionfacturagastoBorrar.html",
		"vista/devolucionfacturagasto/devolucionfacturagastoScript.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaPlandecuentaempresa.html")

	log.Println("Error devolucionfacturagasto nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA DEVOLUCION FACTURA GASTO ELIMINAR
func DevolucionfacturagastoEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	codigo := mux.Vars(r)["codigo"]

	// borrar DEVOLUCION FACTURA GASTO
	delForm, err := db.Prepare("DELETE from devolucionfacturagasto WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(codigo)

	// borar detalle
	delForm1, err := db.Prepare("DELETE from devolucionfacturagastodetalle WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm1.Exec(codigo)

	// borra detalle anterior
	delForm, err = db.Prepare("DELETE from comprobantedetalle WHERE documento=$1 and numero=$2")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec("25", codigo)

	// borra cabecera anterior

	delForm1, err = db.Prepare("DELETE from comprobante WHERE documento=$1 and numero=$2 ")
	if err != nil {
		panic(err.Error())
	}
	delForm1.Exec("25", codigo)

	log.Println("Registro Eliminado" + codigo)
	http.Redirect(w, r, "/DevolucionfacturagastoLista", 301)
}

// TRAER PEDIDO DEVOLUCION FACTURA GASTO
func Datosfacturagasto(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	Tercero := mux.Vars(r)["tercero"]
	log.Println("inicio factura gasto editar" + Codigo)
	db := dbConn()
	var res []facturagasto

	// TRAE FACTURA GASTO
	v := facturagasto{}
	err := db.Get(&v, "SELECT * FROM facturagasto where codigo=$1 and Tercero=$2", Codigo, Tercero)
	var valida bool
	valida = true

	switch err {
	case nil:
		log.Printf(" facturagasto existe: %+v\n", v)
	case sql.ErrNoRows:
		log.Println(" facturagasto NO Existe")
		valida = false
	default:
		log.Printf(" facturagasto error: %s\n", err)
	}
	det := []facturagastodetalleeditar{}
	t := tercero{}

	// trae datos si existe factura de gastos
	if valida == true {
		err2 := db.Select(&det, FacturagastoConsultaDetalle(), Codigo)
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

// INICIA DEVOLUCION FACTURA GASTO PDF
func DevolucionfacturagastoPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]

	// traer DEVOLUCION FACTURA GASTO
	miDevolucionfacturagasto := devolucionfacturagasto{}
	err := db.Get(&miDevolucionfacturagasto, "SELECT * FROM devolucionfacturagasto where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}

	// traer detalle
	miDetalle := []devolucionfacturagastodetalleeditar{}
	err2 := db.Select(&miDetalle, DevolucionfacturagastoConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	// traer tercero
	miTercero := tercero{}
	err3 := db.Get(&miTercero, "SELECT * FROM tercero where codigo=$1", miDevolucionfacturagasto.Tercero)
	if err3 != nil {
		log.Fatalln(err3)
	}

	// traer almacenista
	miAlmacenista := almacenista{}
	err4 := db.Get(&miAlmacenista, "SELECT * FROM almacenista where codigo=$1", miDevolucionfacturagasto.Almacenista)
	if err4 != nil {
		log.Fatalln(err4)
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
		pdf.CellFormat(190, 10, e.Telefono1+" - "+e.Telefono2, "0", 0, "C", false, 0, "")
		pdf.Ln(4)
		pdf.CellFormat(190, 10, ene(c.NombreCiudad+" - "+c.NombreDepartamento), "0", 0, "C", false, 0,
			"")

		// DEVOLUCION FACTURA GASTO NUMERO
		pdf.SetY(20)
		pdf.SetX(80)
		pdf.Ln(8)
		pdf.SetX(75)
		pdf.SetFont("Arial", "", 11)
		pdf.CellFormat(190, 10, "DEVOLUCION", "0", 0, "C",
			false, 0, "")
		pdf.Ln(5)
		pdf.SetX(75)
		pdf.CellFormat(190, 10, "FACTURA GASTO", "0", 0, "C",
			false, 0, "")
		pdf.Ln(5)
		pdf.SetX(75)
		pdf.CellFormat(190, 10, " No. "+Codigo, "0", 0, "C",
			false, 0, "")
		pdf.Ln(6)
	})

	pdf.SetFooterFunc(func() {
		pdf.SetFont("Arial", "", 8)
		pdf.SetY(259)
		pdf.SetX(20)
		pdf.CellFormat(80, 10, "www.Sadconf.com.co", "",
			0, "L", false, 0, "")
		pdf.SetX(129)
		pdf.CellFormat(80, 10, fmt.Sprintf(" %d de {nb}", pdf.PageNo()), "",
			0, "R", false, 0, "")
	})

	// inicia suma de paginas
	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.Ln(1)
	DevolucionfacturagastoCabecera(pdf, miTercero, miDevolucionfacturagasto, miAlmacenista)

	var filas = len(miDetalle)
	// UNA PAGINA
	if filas <= 15 {
		for i, miFila := range miDetalle {
			var a = i + 1
			DevolucionfacturagastoFilaDetalle(pdf, miFila, a)
		}
		DevolucionfacturagastoPieDePagina(pdf, miTercero, miDevolucionfacturagasto)
	} else {
		// mas de 15 y menos de 19 DOS PAGINAS
		if filas > 15 && filas <= 34 {
			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 19 {
					DevolucionfacturagastoFilaDetalle(pdf, miFila, a)
				}
			}
			DevolucionfacturagastoLinea(pdf, miTercero, miDevolucionfacturagasto)
			// segunda pagina
			pdf.AddPage()
			DevolucionfacturagastoCabecera(pdf, miTercero, miDevolucionfacturagasto, miAlmacenista)
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 19 {
					DevolucionfacturagastoFilaDetalle(pdf, miFila, a)
				}
			}

			DevolucionfacturagastoPieDePagina(pdf, miTercero, miDevolucionfacturagasto)
		} else {
			// mas de TRES PAGINAS

			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 19 {
					DevolucionfacturagastoFilaDetalle(pdf, miFila, a)
				}
			}
			DevolucionfacturagastoLinea(pdf, miTercero, miDevolucionfacturagasto)
			pdf.AddPage()
			DevolucionfacturagastoCabecera(pdf, miTercero, miDevolucionfacturagasto, miAlmacenista)
			// segunda pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 19 && a <= 38 {
					DevolucionfacturagastoFilaDetalle(pdf, miFila, a)
				}
			}
			DevolucionfacturagastoLinea(pdf, miTercero, miDevolucionfacturagasto)
			pdf.AddPage()
			DevolucionfacturagastoCabecera(pdf, miTercero, miDevolucionfacturagasto, miAlmacenista)
			// tercera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 82 {
					DevolucionfacturagastoFilaDetalle(pdf, miFila, a)
				}
			}

			DevolucionfacturagastoPieDePagina(pdf, miTercero, miDevolucionfacturagasto)
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

func DevolucionfacturagastoCabecera(pdf *gofpdf.Fpdf, miTercero tercero, miDevolucionfacturagasto devolucionfacturagasto, miAlmacenista almacenista) {
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
	pdf.CellFormat(40, 4, Titulo(TraerFormadepago(miDevolucionfacturagasto.Formadepago)), "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Fecha de Expedicion", "", 0,
		"L", false, 0, "")
	pdf.SetX(150)
	pdf.CellFormat(40, 4, miDevolucionfacturagasto.Fecha.Format("02/01/2006")+" "+Titulo(miDevolucionfacturagasto.Hora), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)

	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Medio de Pago", "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, TraerMediodepago(miDevolucionfacturagasto.Mediodepago), "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Fecha de Vencimiento", "", 0,
		"L", false, 0, "")
	pdf.SetX(150)
	pdf.CellFormat(40, 4, miDevolucionfacturagasto.Vence.Format("02/01/2006"), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Factura Gasto No.", "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, miDevolucionfacturagasto.Facturagasto, "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Almacenista", "", 0,
		"L", false, 0, "")
	pdf.SetX(135)
	pdf.CellFormat(40, 4, miAlmacenista.Nombre, "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)

	// RELLENO TITULO
	pdf.SetFont("Arial", "", 10)
	pdf.SetY(87)
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
	pdf.SetX(124)
	pdf.CellFormat(190, 6, "UNIDAD", "0", 0,
		"L", false, 0, "")
	pdf.SetX(141)
	pdf.CellFormat(190, 6, "CANTIDAD", "0", 0,
		"L", false, 0, "")
	pdf.SetX(162)
	pdf.CellFormat(190, 6, "P. UNITARIO", "0", 0,
		"L", false, 0, "")
	pdf.SetX(190)
	pdf.CellFormat(190, 6, "TOTAL", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}
func DevolucionfacturagastoFilaDetalle(pdf *gofpdf.Fpdf, miFila devolucionfacturagastodetalleeditar, a int) {
	ene := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetX(21)
	var yinicial float64
	yinicial = pdf.GetY()
	pdf.CellFormat(183, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(40, 4, miFila.Codigoservicio, "", 0,
		"L", false, 0, "")
	var y float64
	y = pdf.GetY()
	pdf.SetX(42)
	pdf.MultiCell(80, 4, ene(Mayuscula(miFila.Nombreservicio)), "", "L", false)
	var yfinal float64
	yfinal = pdf.GetY()
	pdf.SetY(y)
	pdf.SetX(100)
	pdf.CellFormat(40, 4, ene(Titulo(miFila.Unidadservicio)), "", 0,
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

func DevolucionfacturagastoPieDePagina(pdf *gofpdf.Fpdf, miTercero tercero, miDevolucionfacturagasto devolucionfacturagasto) {

	Totalletras, err := IntLetra(Cadenaentero(miDevolucionfacturagasto.Neto))
	if err != nil {
		fmt.Println(err)
	}

	ene := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetFont("Arial", "", 8)
	pdf.SetY(233)
	pdf.SetX(20)
	pdf.MultiCell(184, 4, "SON: "+ene(Mayuscula(Totalletras))+" PESOS MDA. CTE.", "0", "L", false)
	pdf.SetFont("Arial", "", 9)
	pdf.SetY(254)
	pdf.Line(45, 254, 125, 254)
	pdf.Ln(1)
	pdf.SetX(65)
	pdf.CellFormat(40, 4, "FIRMA RESPONSABLE", "0", 0, "C",
		false, 0, "")

	// PRESENTA DATOS CON VALORES //
	pdf.SetFont("Arial", "", 9)
	var separador float64
	var altoseparador float64
	separador = 254
	altoseparador = -4

	// INICIA DATOS FACTURA
	if miDevolucionfacturagasto.Neto != "0" {
		pdf.SetFont("Arial", "", 9)
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "TOTAL", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miDevolucionfacturagasto.Neto), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucionfacturagasto.TotalRetencionIca != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "RET. ICA.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, miDevolucionfacturagasto.PorcentajeRetencionIca+"%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miDevolucionfacturagasto.TotalRetencionIca), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucionfacturagasto.TotalRetencionFuente != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "RET. FTE.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, miDevolucionfacturagasto.PorcentajeRetencionFuente+"%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miDevolucionfacturagasto.TotalRetencionFuente), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucionfacturagasto.TotalIva != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "TOTAL IVA", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, miDevolucionfacturagasto.Facturagastoporcentajeiva+"%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miDevolucionfacturagasto.TotalIva), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucionfacturagasto.Subtotal0 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "SUBTOTAL", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miDevolucionfacturagasto.Subtotal0), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}
}

func DevolucionfacturagastoLinea(pdf *gofpdf.Fpdf, miTercero tercero, miDevolucionfacturagasto devolucionfacturagasto) {
	pdf.SetY(259)
	pdf.Line(20, 259, 204, 259)
}

// INICIA DEVOLUCION FACTURA GASTO TODOS PDF
func DevolucionfacturagastoTodosCabecera(pdf *gofpdf.Fpdf) {
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

func DevolucionfacturagastoTodosDetalle(pdf *gofpdf.Fpdf, miFila devolucionfacturagastoLista, a int) {
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

func DevolucionfacturagastoTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string
	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre,centro.nombre as CentroNombre,devolucionfacturagasto.neto,devolucionfacturagasto.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	consulta += " FROM devolucionfacturagasto "
	consulta += " inner join tercero on tercero.codigo=devolucionfacturagasto.tercero "
	consulta += " inner join centro on centro.codigo=devolucionfacturagasto.centro "
	consulta += " inner join almacenista on almacenista.codigo=devolucionfacturagasto.almacenista "
	consulta += " ORDER BY cast(devolucionfacturagasto.codigo as integer) ASC"

	t := []devolucionfacturagastoLista{}
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
		pdf.CellFormat(190, 10, "DATOS DEVOLUCION FACTURA GASTO", "0", 0,
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

	DevolucionfacturagastoTodosCabecera(pdf)
	// tercera pagina

	for i, miFila := range t {
		var a = i + 1
		if math.Mod(float64(a), 49) == 0 {
			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			DevolucionfacturagastoTodosCabecera(pdf)
		}
		DevolucionfacturagastoTodosDetalle(pdf, miFila, a)
	}

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

// TERMINA DEVOLUCION FACTURA GASTO TODOS PDF

// FACTURA GASTO  EXCEL
func DevolucionfacturagastoExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string
	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre,centro.nombre as CentroNombre,devolucionfacturagasto.neto,devolucionfacturagasto.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	consulta += " FROM devolucionfacturagasto "
	consulta += " inner join tercero on tercero.codigo=devolucionfacturagasto.tercero "
	consulta += " inner join centro on centro.codigo=devolucionfacturagasto.centro "
	consulta += " inner join almacenista on almacenista.codigo=devolucionfacturagasto.almacenista "
	consulta += " ORDER BY cast(devolucionfacturagasto.codigo as integer) ASC"
	t := []devolucionfacturagastoLista{}
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
	f.SetCellValue("Sheet1", "A9", "LISTADO DE DEVOLUCION FACTURA GASTO")
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

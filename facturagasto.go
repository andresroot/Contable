package main

// INICIA FACTURA GASTO IMPORTAR PAQUETES
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

// INICIA FACTURA GASTO ESTRUCTURA JSON
type facturagastoJson struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Nombre string `json:"nombre"`
}

// INICIA FACTURA GASTO ESTRUCTURA
type facturagastoLista struct {
	Codigo            string
	Fecha             time.Time
	Neto              string
	Tercero           string
	TerceroNombre     string
	CentroNombre      string
	AlmacenistaNombre string
}

// INICIA FACTURA GASTO ESTRUCTURA
type facturagasto struct {
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
	Detalle                   []facturagastodetalle       `json:"Detalle"`
	DetalleEditar             []facturagastodetalleeditar `json:"DetalleEditar"`
	TerceroDetalle            tercero
	Pedidofacturagasto        string
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
type facturagastodetalle struct {
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

// INICIA FACTURA GASTO DETALLE EDITAR
type facturagastodetalleeditar struct {
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

// INICIA FACTURA GASTO CONSULTA DETALLE
func FacturagastoConsultaDetalle() string {
	var consulta = ""
	consulta = "select "
	consulta += "facturagastodetalle.Id as id ,"
	consulta += "facturagastodetalle.Codigo as codigo,"
	consulta += "facturagastodetalle.Fila as fila,"
	consulta += "facturagastodetalle.Cantidad as cantidad,"
	consulta += "facturagastodetalle.Precio as precio,"
	consulta += "facturagastodetalle.Descuento as descuento,"
	consulta += "facturagastodetalle.Montodescuento as montodescuento,"
	consulta += "facturagastodetalle.Sigratis as sigratis,"
	consulta += "facturagastodetalle.Subtotal as subtotal,"
	consulta += "facturagastodetalle.Subtotaldescuento  as subtotaldescuento,"
	consulta += "facturagastodetalle.Pagina as pagina ,"
	consulta += "facturagastodetalle.Bodega as bodega,"
	consulta += "facturagastodetalle.Fecha as fecha,"
	consulta += "facturagastodetalle.Nombreservicio,"
	consulta += "facturagastodetalle.Unidadservicio,"
	consulta += "facturagastodetalle.Codigoservicio"
	consulta += " from facturagastodetalle "
	consulta += " where facturagastodetalle.codigo=$1"
	log.Println(consulta)
	return consulta
}

// INICIA FACTURA GASTO LISTA
func FacturagastoLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/facturagasto/facturagastoLista.html")
	log.Println("Error facturagasto 0")
	var consulta string
	var miperiodo = periodoSesion(r)

	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre,centro.nombre as CentroNombre,facturagasto.neto,facturagasto.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	consulta += " FROM facturagasto "
	consulta += " inner join tercero on tercero.codigo=facturagasto.tercero "
	consulta += " inner join centro on centro.codigo=facturagasto.centro "
	consulta += " inner join almacenista on almacenista.codigo=facturagasto.almacenista "
	consulta += " where  extract(year from fecha)=" + miperiodo + " "
	consulta += " ORDER BY fecha ASC"

	db := dbConn()
	res := []facturagastoLista{}

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

// INICIA FACTURA GASTO NUEVO
func FacturagastoNuevo(w http.ResponseWriter, r *http.Request) {

	// TRAE COPIA DE EDITAR
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio facturagasto editar" + Codigo)

	db := dbConn()
	v := facturagasto{}
	det := []facturagastodetalleeditar{}
	if Codigo == "False" {

	} else {

		err := db.Get(&v, "SELECT * FROM facturagasto where codigo=$1", Codigo)
		if err != nil {
			log.Fatalln(err)
		}

		v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
		v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
		v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
		v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
		v.CentroDetalle = TraerCentroConsulta(v.Centro)

		err2 := db.Select(&det, FacturagastoConsultaDetalle(), Codigo)
		if err2 != nil {
			fmt.Println(err2)
		}
	}

	parametros := map[string]interface{}{
		"codigo":       Codigo,
		"facturagasto": v,
		"almacenista":  PrimerAlmacenista(),
		"detalle":      det,
		"hosting":      ruta,
		"miperiodo":    periodoSesion(r),
	}
	//TERMINA TRAE COPIA DE EDITAR

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/facturagasto/facturagastoNuevo.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/facturagasto/facturagastoScript.html",
		"vista/autocompleta/autocompletaPlandecuentaempresa.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error facturagasto nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA INSERTAR COMPROBANTE DE FACTURA GASTO
func InsertaDetalleComprobanteFacturagasto(miFilaComprobante comprobantedetalle, miComprobante comprobante, miFacturagasto facturagasto) {
	db := dbConn()

	// traer tercero
	miTercero := tercero{}
	err3 := db.Get(&miTercero, "SELECT * FROM tercero where codigo=$1", miFacturagasto.Tercero)
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
	q += " )"

	log.Println("Cadena SQL " + q)
	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	if len(miFilaComprobante.Debito) > 0 {
		miFilaComprobante.Debito = miFilaComprobante.Debito
		//+ ".00"
	}

	if len(miFilaComprobante.Credito) > 0 {
		miFilaComprobante.Credito = miFilaComprobante.Credito
		//+ ".00"
	}

	// TERMINA COMPROBANTE GRABAR INSERTAR
	_, err = insForm.Exec(
		miFilaComprobante.Fila,
		miFilaComprobante.Cuenta,
		miFacturagasto.Tercero,
		miFacturagasto.Centro,
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

// INICIA FACTURA GASTO INSERTAR AJAX
func FacturagastoAgregar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var tempFacturagasto facturagasto

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// carga informacion de la FACTURA GASTO
	err = json.Unmarshal(b, &tempFacturagasto)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	parametrosinventario := configuracioninventario{}
	err = db.Get(&parametrosinventario, "SELECT * FROM configuracioninventario")
	if err != nil {
		panic(err.Error())
	}

	var DocumentoContable string
	DocumentoContable = "9"

	if tempFacturagasto.Accion == "Actualizar" {

		// borra detalle anterior
		delForm, err := db.Prepare("DELETE from facturagastodetalle WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(tempFacturagasto.Codigo)

		// borra cabecera anterior
		delForm1, err := db.Prepare("DELETE from facturagasto WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm1.Exec(tempFacturagasto.Codigo)
	}

	// INSERTA DETALLE
	for i, x := range tempFacturagasto.Detalle {
		var a = i
		var q string
		q = "insert into facturagastodetalle ("
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

		// TERMINA FACTURA GASTO GRABAR INSERTAR
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

	// INICIA INSERTAR FACTURA GASTO
	log.Println("Got %s age %s club %s\n", tempFacturagasto.Codigo, tempFacturagasto.Tercero, tempFacturagasto.Subtotal)
	var q string
	q = "insert into facturagasto ("
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
	q += "Pedidofacturagasto,"
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
	log.Println("Hora", tempFacturagasto.Fecha.Format("02/01/2006"))
	currentTime := time.Now()

	_, err = insForm.Exec(
		tempFacturagasto.Codigo,
		tempFacturagasto.Fecha.Format(layout),
		tempFacturagasto.Vence.Format(layout),
		currentTime.Format("3:4:5 PM"),
		Quitacoma(tempFacturagasto.Descuento),
		Quitacoma(tempFacturagasto.Subtotaldescuento19),
		Quitacoma(tempFacturagasto.Subtotaldescuento5),
		Quitacoma(tempFacturagasto.Subtotaldescuento0),
		Quitacoma(tempFacturagasto.Subtotal),
		Quitacoma(tempFacturagasto.Subtotal19),
		Quitacoma(tempFacturagasto.Subtotal5),
		Quitacoma(tempFacturagasto.Subtotal0),
		Quitacoma(tempFacturagasto.Subtotaliva19),
		Quitacoma(tempFacturagasto.Subtotaliva5),
		Quitacoma(tempFacturagasto.Subtotaliva0),
		Quitacoma(tempFacturagasto.Subtotalbase19),
		Quitacoma(tempFacturagasto.Subtotalbase5),
		Quitacoma(tempFacturagasto.Subtotalbase0),
		Quitacoma(tempFacturagasto.TotalIva),
		Quitacoma(tempFacturagasto.Total),
		Quitacoma(tempFacturagasto.Neto),
		tempFacturagasto.Items,
		Quitacoma(tempFacturagasto.PorcentajeRetencionFuente),
		Quitacoma(tempFacturagasto.TotalRetencionFuente),
		Quitacoma(tempFacturagasto.PorcentajeRetencionIca),
		Quitacoma(tempFacturagasto.TotalRetencionIca),
		tempFacturagasto.Formadepago,
		tempFacturagasto.Mediodepago,
		tempFacturagasto.Tercero,
		tempFacturagasto.Almacenista,
		tempFacturagasto.Pedidofacturagasto,
		tempFacturagasto.Centro,
		tempFacturagasto.Tipo,
		tempFacturagasto.Facturagastocuenta0,
		tempFacturagasto.Facturagastonombre0,
		tempFacturagasto.Facturagastocuentaretfte,
		tempFacturagasto.Facturagastonombreretfte,
		tempFacturagasto.Facturagastocuentaiva,
		tempFacturagasto.Facturagastonombreiva,
		Quitacoma(tempFacturagasto.Facturagastoporcentajeiva))

	if err != nil {
		panic(err)
	}

	// INSERTAR COMPROBANTE CONTABILIDAD
	var tempComprobante comprobante
	var tempComprobanteDetalle comprobantedetalle
	tempComprobante.Documento = DocumentoContable
	tempComprobante.Numero = tempFacturagasto.Codigo
	tempComprobante.Fecha = tempFacturagasto.Fecha
	tempComprobante.Fechaconsignacion = tempFacturagasto.Fecha
	tempComprobante.Debito = tempFacturagasto.Neto + ".00"
	tempComprobante.Credito = tempFacturagasto.Neto + ".00"
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

	// INSERTAR CUENTA DEBITO FACTURA GASTO 0%
	if tempFacturagasto.Subtotal0 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempFacturagasto.Subtotal0) - Flotante(tempFacturagasto.Subtotaldescuento0)
		tempComprobanteDetalle.Cuenta = tempFacturagasto.Facturagastocuenta0
		tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(tempFacturagasto.Subtotal0) - Flotante(tempFacturagasto.Subtotaldescuento0))
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteFacturagasto(tempComprobanteDetalle, tempComprobante, tempFacturagasto)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA CREDITO DESCUENTO
	//if (tempFacturagasto.Descuento!="0")	{
	//	fila=fila+1
	//	tempComprobanteDetalle.Fila=strconv.Itoa(fila)
	//	totalCredito+=Flotante(tempFacturagasto.Descuento)
	//	tempComprobanteDetalle.Cuenta  = parametrosinventario.Facturagastocuentadescuento
	//	tempComprobanteDetalle.Debito = ""
	//	tempComprobanteDetalle.Credito = tempFacturagasto.Descuento
	//	InsertaDetalleComprobanteFacturagasto(tempComprobanteDetalle,tempComprobante,tempFacturagasto)
	//	log.Println("credito linea" + fmt.Sprintf("%.2f",totalCredito))
	//}

	// INSERTAR total Iva
	if tempFacturagasto.TotalIva != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempFacturagasto.TotalIva)
		tempComprobanteDetalle.Cuenta = tempFacturagasto.Facturagastocuentaiva
		tempComprobanteDetalle.Debito = tempFacturagasto.TotalIva
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteFacturagasto(tempComprobanteDetalle, tempComprobante, tempFacturagasto)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA CREDITO RET. FTE.
	if tempFacturagasto.TotalRetencionFuente != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempFacturagasto.TotalRetencionFuente)
		tempComprobanteDetalle.Cuenta = tempFacturagasto.Facturagastocuentaretfte
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = tempFacturagasto.TotalRetencionFuente
		InsertaDetalleComprobanteFacturagasto(tempComprobanteDetalle, tempComprobante, tempFacturagasto)
		log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito))
	}

	// INSERTAR CUENTA CREDITO RET. ICA.
	if tempFacturagasto.TotalRetencionIca != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempFacturagasto.TotalRetencionIca)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Facturagastocuentaretica
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = tempFacturagasto.TotalRetencionIca
		InsertaDetalleComprobanteFacturagasto(tempComprobanteDetalle, tempComprobante, tempFacturagasto)
		log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito))
	}

	// INSERTAR CUENTA CREDITO PROVEEDOR
	if tempFacturagasto.Neto != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempFacturagasto.Neto)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Facturagastocuentaproveedor
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = tempFacturagasto.Neto
		InsertaDetalleComprobanteFacturagasto(tempComprobanteDetalle, tempComprobante, tempFacturagasto)
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

// INICIA FACTURA GASTO EXISTE
func FacturagastoExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM facturagasto  WHERE codigo=$1", Codigo)
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

// INICIA FACTURA GASTO EDITAR
func FacturagastoEditar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio facturagasto editar" + Codigo)
	db := dbConn()

	// traer FACTURA GASTO
	v := facturagasto{}
	err := db.Get(&v, "SELECT * FROM facturagasto where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []facturagastodetalleeditar{}

	err2 := db.Select(&det, FacturagastoConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
	v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
	v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
	v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
	v.CentroDetalle = TraerCentroConsulta(v.Centro)

	parametros := map[string]interface{}{
		"facturagasto": v,
		"detalle":      det,
		"hosting":      ruta,
		"miperiodo":    periodoSesion(r),
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/facturagasto/facturagastoEditar.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/facturagasto/facturagastoScript.html",
		"vista/autocompleta/autocompletaPlandecuentaempresa.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error facturagasto nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA FACTURA GASTO BORRAR
func FacturagastoBorrar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio facturagasto editar" + Codigo)
	db := dbConn()

	// traer FACTURA GASTO
	v := facturagasto{}
	err := db.Get(&v, "SELECT * FROM facturagasto where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []facturagastodetalleeditar{}
	err2 := db.Select(&det, FacturagastoConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
	v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
	v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
	v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
	v.CentroDetalle = TraerCentroConsulta(v.Centro)

	parametros := map[string]interface{}{
		"facturagasto": v,
		"detalle":      det,
		"hosting":      ruta,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/facturagasto/facturagastoBorrar.html",
		"vista/facturagasto/facturagastoScript.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaPlandecuentaempresa.html")

	log.Println("Error facturagasto nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA FACTURA GASTO ELIMINAR
func FacturagastoEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	codigo := mux.Vars(r)["codigo"]

	// borrar FACTURA GASTO
	delForm, err := db.Prepare("DELETE from facturagasto WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(codigo)

	// borar detalle
	delForm1, err := db.Prepare("DELETE from facturagastodetalle WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm1.Exec(codigo)

	// borra detalle anterior
	delForm, err = db.Prepare("DELETE from comprobantedetalle WHERE documento=$1 and numero=$2")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec("9", codigo)

	// borra cabecera anterior

	delForm1, err = db.Prepare("DELETE from comprobante WHERE documento=$1 and numero=$2 ")
	if err != nil {
		panic(err.Error())
	}
	delForm1.Exec("9", codigo)

	log.Println("Registro Eliminado" + codigo)
	http.Redirect(w, r, "/FacturagastoLista", 301)
}

// TRAER PEDIDO FACTURA GASTO
func Datospedidofacturagasto(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	Tercero := mux.Vars(r)["tercero"]
	log.Println("inicio pedido editar" + Codigo)
	db := dbConn()
	var res []pedidofacturagasto

	// traer PEDIDO
	v := pedidofacturagasto{}
	err := db.Get(&v, "SELECT * FROM pedidofacturagasto where codigo=$1 and Tercero=$2", Codigo, Tercero)
	var valida bool
	valida = true

	switch err {
	case nil:
		log.Printf("pedido existe: %+v\n", v)
	case sql.ErrNoRows:
		log.Println("pedidofacturagasto NO Existe")
		valida = false
	default:
		log.Printf("pedidofacturagasto error: %s\n", err)
	}
	det := []pedidofacturagastodetalleeditar{}
	t := tercero{}

	// trae datos si existe pedido
	if valida == true {
		err2 := db.Select(&det, PedidofacturagastoConsultaDetalle(), Codigo)
		if err2 != nil {
			fmt.Println(err2)
		}

		v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
		v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
		v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
		v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
		v.CentroDetalle = TraerCentroConsulta(v.Centro)

		//v.TerceroDetalle=t;
		v.DetalleEditar = det
		res = append(res, v)
	}

	log.Println("codigo nombre99" + t.Codigo + t.Nombre)
	data, err := json.Marshal(res)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// INICIA FACTURA GASTO PDF
func FacturagastoPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]

	// TRAER FACTURA GASTO
	miFacturagasto := facturagasto{}
	err := db.Get(&miFacturagasto, "SELECT * FROM facturagasto where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}

	// traer detalle
	miDetalle := []facturagastodetalleeditar{}
	err2 := db.Select(&miDetalle, FacturagastoConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	// traer tercero
	miTercero := tercero{}
	err3 := db.Get(&miTercero, "SELECT * FROM tercero where codigo=$1", miFacturagasto.Tercero)
	if err3 != nil {
		log.Fatalln(err3)
	}

	// traer almacenista
	miAlmacenista := almacenista{}
	err4 := db.Get(&miAlmacenista, "SELECT * FROM almacenista where codigo=$1", miFacturagasto.Almacenista)
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

		// FACTURA GASTO NUMERO
		pdf.SetY(20)
		pdf.SetX(80)
		pdf.Ln(8)
		pdf.SetX(75)
		pdf.SetFont("Arial", "", 11)
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
	FacturagastoCabecera(pdf, miTercero, miFacturagasto, miAlmacenista)

	var filas = len(miDetalle)
	// pagina igual a 15
	if filas <= 15 {
		for i, miFila := range miDetalle {
			var a = i + 1
			FacturagastoFilaDetalle(pdf, miFila, a)
		}
		FacturagastoPieDePagina(pdf, miTercero, miFacturagasto)
	} else {
		// mas de 15 y menos de 19 dos paginas
		if filas > 15 && filas <= 34 {
			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 19 {
					FacturagastoFilaDetalle(pdf, miFila, a)
				}
			}
			FacturagastoLinea(pdf, miTercero, miFacturagasto)
			// segunda pagina
			pdf.AddPage()
			FacturagastoCabecera(pdf, miTercero, miFacturagasto, miAlmacenista)
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 19 {
					FacturagastoFilaDetalle(pdf, miFila, a)
				}
			}

			FacturagastoPieDePagina(pdf, miTercero, miFacturagasto)
		} else {
			// mas de tres paginas

			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 19 {
					FacturagastoFilaDetalle(pdf, miFila, a)
				}
			}
			FacturagastoLinea(pdf, miTercero, miFacturagasto)
			pdf.AddPage()
			FacturagastoCabecera(pdf, miTercero, miFacturagasto, miAlmacenista)
			// segunda pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 19 && a <= 38 {
					FacturagastoFilaDetalle(pdf, miFila, a)
				}
			}
			FacturagastoLinea(pdf, miTercero, miFacturagasto)
			pdf.AddPage()
			FacturagastoCabecera(pdf, miTercero, miFacturagasto, miAlmacenista)
			// tercera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 38 {
					FacturagastoFilaDetalle(pdf, miFila, a)
				}
			}

			FacturagastoPieDePagina(pdf, miTercero, miFacturagasto)
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

func FacturagastoCabecera(pdf *gofpdf.Fpdf, miTercero tercero, miFacturagasto facturagasto, miAlmacenista almacenista) {
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
	pdf.CellFormat(40, 4, Titulo(TraerFormadepago(miFacturagasto.Formadepago)), "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Fecha de Expedicion", "", 0,
		"L", false, 0, "")
	pdf.SetX(150)
	pdf.CellFormat(40, 4, miFacturagasto.Fecha.Format("02/01/2006")+" "+Titulo(miFacturagasto.Hora), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)

	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Medio de Pago", "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, TraerMediodepago(miFacturagasto.Mediodepago), "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Fecha de Vencimiento", "", 0,
		"L", false, 0, "")
	pdf.SetX(150)
	pdf.CellFormat(40, 4, miFacturagasto.Vence.Format("02/01/2006"), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Pedido No.", "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, miFacturagasto.Pedidofacturagasto, "", 0,
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
	pdf.CellFormat(190, 6, "TOTAL", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}
func FacturagastoFilaDetalle(pdf *gofpdf.Fpdf, miFila facturagastodetalleeditar, a int) {
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

func FacturagastoPieDePagina(pdf *gofpdf.Fpdf, miTercero tercero, miFacturagasto facturagasto) {

	Totalletras, err := IntLetra(Cadenaentero(miFacturagasto.Neto))
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
	if miFacturagasto.Neto != "0" {
		pdf.SetFont("Arial", "", 9)
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "TOTAL", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miFacturagasto.Neto), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miFacturagasto.TotalRetencionIca != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "RET. ICA.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, miFacturagasto.PorcentajeRetencionIca+"%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miFacturagasto.TotalRetencionIca), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miFacturagasto.TotalRetencionFuente != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "RET. FTE.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, miFacturagasto.PorcentajeRetencionFuente+"%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miFacturagasto.TotalRetencionFuente), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miFacturagasto.TotalIva != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "TOTAL IVA", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, miFacturagasto.Facturagastoporcentajeiva+"%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miFacturagasto.TotalIva), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miFacturagasto.Subtotal0 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "SUBTOTAL", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miFacturagasto.Subtotal0), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}
}

func FacturagastoLinea(pdf *gofpdf.Fpdf, miTercero tercero, miFacturagasto facturagasto) {
	pdf.SetY(259)
	pdf.Line(20, 259, 204, 259)
}

// INICIA FACTURA GASTO TODOS PDF
func FacturagastoTodosCabecera(pdf *gofpdf.Fpdf) {
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

func FacturagastoTodosDetalle(pdf *gofpdf.Fpdf, miFila facturagastoLista, a int) {
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

func FacturagastoTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string
	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre,centro.nombre as CentroNombre,facturagasto.neto,facturagasto.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	consulta += " FROM facturagasto "
	consulta += " inner join tercero on tercero.codigo=facturagasto.tercero "
	consulta += " inner join centro on centro.codigo=facturagasto.centro "
	consulta += " inner join almacenista on almacenista.codigo=facturagasto.almacenista "
	consulta += " ORDER BY cast(facturagasto.codigo as integer) ASC"

	t := []facturagastoLista{}
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
		pdf.CellFormat(190, 10, "DATOS FACTURA GASTO", "0", 0,
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

	FacturagastoTodosCabecera(pdf)
	// tercera pagina

	for i, miFila := range t {
		var a = i + 1
		if math.Mod(float64(a), 49) == 0 {
			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			FacturagastoTodosCabecera(pdf)
		}
		FacturagastoTodosDetalle(pdf, miFila, a)
	}

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

// TERMINA FACTURA GASTO TODOS PDF

// FACTURA GASTO  EXCEL
func FacturagastoExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string
	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre,centro.nombre as CentroNombre,facturagasto.neto,facturagasto.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	consulta += " FROM facturagasto "
	consulta += " inner join tercero on tercero.codigo=facturagasto.tercero "
	consulta += " inner join centro on centro.codigo=facturagasto.centro "
	consulta += " inner join almacenista on almacenista.codigo=facturagasto.almacenista "
	consulta += " ORDER BY cast(facturagasto.codigo as integer) ASC"
	t := []facturagastoLista{}
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
	f.SetCellValue("Sheet1", "A9", "LISTADO DE FACTURA GASTO")
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

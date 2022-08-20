package main

// INICIA DEVOLUCIONCOMPRA IMPORTAR PAQUETES
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

// INICIA DEVOLUCIONCOMPRA ESTRUCTURA JSON
type DevolucioncompraJson struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Nombre string `json:"nombre"`
}

// INICIA DEVOLUCIONCOMPRA ESTRUCTURA
type devolucioncompralista struct {
	Codigo            string
	Fecha             time.Time
	Neto              string
	Tercero           string
	TerceroNombre     string
	CentroNombre      string
	AlmacenistaNombre string
}

// INICIA DEVOLUCIONCOMPRA ESTRUCTURA
type devolucioncompra struct {
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
	Detalle                   []devolucioncompradetalle       `json:"Detalle"`
	DetalleEditar             []devolucioncompradetalleeditar `json:"DetalleEditar"`
	TerceroDetalle            tercero
	Compra                    string
	Tipo                      string
	Centro                    string
	FormadepagoDetalle        formadepago
	MediodepagoDetalle        mediodepago
	AlmacenistaDetalle        almacenista
	CentroDetalle             centro
}

// INICIA DEVOLUCIONCOMPRADETALLE ESTRUCTURA
type devolucioncompradetalle struct {
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
	Producto          string
	Tipo              string
	Fecha             time.Time
}

// estructura para editar
type devolucioncompradetalleeditar struct {
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
	BodegaNombre      string
	Producto          string
	ProductoNombre    string
	ProductoIva       string
	ProductoUnidad    string
	Tipo              string
	Fecha             time.Time
}

// INICIA COMPRA CONSULTA DETALLE
func DevolucioncompraConsultaDetalle() string {
	var consulta = ""
	consulta = "select "
	consulta += "devolucioncompradetalle.Id as id ,"
	consulta += "devolucioncompradetalle.Codigo as codigo,"
	consulta += "devolucioncompradetalle.Fila as fila,"
	consulta += "devolucioncompradetalle.Cantidad as cantidad,"
	consulta += "devolucioncompradetalle.Precio as precio,"
	consulta += "devolucioncompradetalle.Descuento as descuento,"
	consulta += "devolucioncompradetalle.Montodescuento as montodescuento,"
	consulta += "devolucioncompradetalle.Sigratis as sigratis,"
	consulta += "devolucioncompradetalle.Subtotal as subtotal,"
	consulta += "devolucioncompradetalle.Subtotaldescuento  as subtotaldescuento,"
	consulta += "devolucioncompradetalle.Pagina as pagina ,"
	consulta += "devolucioncompradetalle.Bodega as bodega,"
	consulta += "bodega.Nombre as bodeganombre,"
	consulta += "devolucioncompradetalle.Producto as producto,"
	consulta += "devolucioncompradetalle.Fecha as fecha,"
	consulta += "producto.nombre as ProductoNombre, "
	consulta += "producto.iva as ProductoIva, "
	consulta += "producto.unidad as ProductoUnidad "
	consulta += "from devolucioncompradetalle "
	consulta += "inner join producto on producto.codigo=devolucioncompradetalle.producto "
	consulta += "inner join bodega on bodega.codigo=devolucioncompradetalle.bodega "
	consulta += " where devolucioncompradetalle.codigo=$1 ORDER BY fila"
	log.Println(consulta)
	return consulta
}

// INICIA DEVOLUCIONCOMPRA LISTA
func DevolucioncompraLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/devolucioncompra/devolucioncompraLista.html")
	log.Println("Error devolucioncompra 0")
	var consulta string
	var miperiodo = periodoSesion(r)

	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre,centro.nombre as CentroNombre,devolucioncompra.neto,devolucioncompra.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	consulta += " FROM devolucioncompra "
	consulta += " inner join tercero on tercero.codigo=devolucioncompra.tercero "
	consulta += " inner join centro on centro.codigo=devolucioncompra.centro "
	consulta += " inner join almacenista on almacenista.codigo=devolucioncompra.almacenista "
	consulta += " where  extract(year from fecha)=" + miperiodo + " "

	consulta += " ORDER BY cast(devolucioncompra.codigo as integer) ASC"

	db := dbConn()
	res := []devolucioncompralista{}

	err := db.Select(&res, consulta)

	if err != nil {
		fmt.Println(err)
	}

	varmap := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
	}
	log.Println("Error devolucioncompra888")
	tmp.Execute(w, varmap)
}

// INICIA DEVOLUCIONCOMPRA NUEVO
func DevolucioncompraNuevo(w http.ResponseWriter, r *http.Request) {

	// TRAE COPIA DE EDITAR
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio devolucioncompra editar" + Codigo)

	db := dbConn()
	v := devolucioncompra{}
	det := []devolucioncompradetalleeditar{}
	if Codigo == "False" {

	} else {

		err := db.Get(&v, "SELECT * FROM devolucioncompra where codigo=$1", Codigo)
		if err != nil {
			log.Fatalln(err)
		}

		v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
		v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
		v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
		v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
		v.CentroDetalle = TraerCentroConsulta(v.Centro)
		err2 := db.Select(&det, DevolucioncompraConsultaDetalle(), Codigo)
		if err2 != nil {
			fmt.Println(err2)
		}
	}

	parametros := map[string]interface{}{
		"miperiodo":        periodoSesion(r),
		"codigo":           Codigo,
		"almacenista":      PrimerAlmacenista(),
		"devolucioncompra": v,
		"detalle":          det,
		"hosting":          ruta,
		"retfte":           TraerParametrosInventario().Compracuentaporcentajeretfte,
	}
	//TERMINA TRAE COPIA DE EDITAR

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/devolucioncompra/devolucioncompraNuevo.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaBodega.html",
		"vista/autocompleta/autocompletaTercero.html",
		"vista/devolucioncompra/devolucioncompraScript.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error devolucioncompra nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA INSERTAR COMPROBANTE DE DEVOLUCION EN COMPRA
func InsertaDetalleComprobanteDevolucioncompra(miFilaComprobante comprobantedetalle, miComprobante comprobante, miDevolucioncompra devolucioncompra) {
	db := dbConn()

	// traer tercero
	miTercero := tercero{}
	err3 := db.Get(&miTercero, "SELECT * FROM tercero where codigo=$1", miDevolucioncompra.Tercero)
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
	q += " ) "
	log.Println("Cadena SQL inserta linea detalle comprobante")
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
		miDevolucioncompra.Tercero,
		miDevolucioncompra.Centro,
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

// INICIA DEVOLUCIONCOMPRA INSERTAR AJAX
func DevolucioncompraAgregar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var tempDevolucioncompra devolucioncompra

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// carga informacion de la DEVOLUCIONCOMPRA
	err = json.Unmarshal(b, &tempDevolucioncompra)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	parametrosinventario := configuracioninventario{}
	err = db.Get(&parametrosinventario, "SELECT * FROM configuracioninventario")
	if err != nil {
		panic(err.Error())
	}

	if tempDevolucioncompra.Accion == "Actualizar" {

		// borra detalle anterior
		delForm, err := db.Prepare("DELETE from devolucioncompradetalle WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(tempDevolucioncompra.Codigo)

		// borra inventario
		Borrarinventario(tempDevolucioncompra.Codigo, "Devolucioncompra")

		// borra cabecera anterior
		delForm1, err := db.Prepare("DELETE from devolucioncompra WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm1.Exec(tempDevolucioncompra.Codigo)
	}

	// INSERTA DETALLE
	for i, x := range tempDevolucioncompra.Detalle {
		var a = i
		var q string
		q = "insert into devolucioncompradetalle ("
		q += "Id,"
		q += "Codigo,"
		q += "Fila,"
		q += "Cantidad,"
		q += "Precio,"
		q += "Subtotal,"
		q += "Pagina,"
		q += "Bodega,"
		q += "Producto,"
		q += "Descuento,"
		q += "Montodescuento,"
		q += "Sigratis,"
		q += "Subtotaldescuento,"
		q += "Tipo,"
		q += "Fecha"
		q += " ) values("
		q += parametros(15)
		q += ")"
		log.Println("Cadena SQL " + q)
		insForm, err := db.Prepare(q)
		if err != nil {
			panic(err.Error())
		}

		// TERMINA DEVOLUCIONCOMPRA GRABAR INSERTAR
		_, err = insForm.Exec(
			x.Id,
			x.Codigo,
			x.Fila,
			Quitacoma(x.Cantidad),
			Quitacoma(x.Precio),
			Quitacoma(x.Subtotal),
			x.Pagina,
			x.Bodega,
			x.Producto,
			Quitacoma(x.Descuento),
			Quitacoma(x.Montodescuento),
			x.Sigratis,
			Quitacoma(x.Subtotaldescuento),
			x.Tipo,
			x.Fecha)
		if err != nil {
			panic(err)
		}

		log.Println("Insertar Producto \n", x.Producto, a)
	}

	// inserta inventario
	miInventario := []inventario{}

	// INSERTA DETALLE INVENTARIO
	for _, x := range tempDevolucioncompra.Detalle {
		miInventario = append(miInventario, inventario{
			x.Fecha, x.Tipo, x.Codigo, x.Bodega,
			x.Producto,
			Quitacoma(x.Cantidad),
			Quitacoma(x.Precio),
			operacionDevolucionCompra})
	}
	// insertar inventario

	InsertaInventario(miInventario)

	// INICIA INSERTAR DEVOLUCIONCOMPRAS
	log.Println("Got %s age %s club %s\n", tempDevolucioncompra.Codigo, tempDevolucioncompra.Tercero, tempDevolucioncompra.Total)
	var q string
	q = "insert into devolucioncompra ("
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
	q += "Compra,"
	q += "Centro,"
	q += "Tipo"
	q += " ) values("
	q += parametros(33)
	q += " ) "

	log.Println("Cadena SQL " + q)
	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	layout := "2006-01-02"
	log.Println("Hora", tempDevolucioncompra.Fecha.Format("02/01/2006"))
	currentTime := time.Now()

	_, err = insForm.Exec(
		tempDevolucioncompra.Codigo,
		tempDevolucioncompra.Fecha.Format(layout),
		tempDevolucioncompra.Vence.Format(layout),
		currentTime.Format("3:4:5 PM"),
		Quitacoma(tempDevolucioncompra.Descuento),
		Quitacoma(tempDevolucioncompra.Subtotaldescuento19),
		Quitacoma(tempDevolucioncompra.Subtotaldescuento5),
		Quitacoma(tempDevolucioncompra.Subtotaldescuento0),
		Quitacoma(tempDevolucioncompra.Subtotal),
		Quitacoma(tempDevolucioncompra.Subtotal19),
		Quitacoma(tempDevolucioncompra.Subtotal5),
		Quitacoma(tempDevolucioncompra.Subtotal0),
		Quitacoma(tempDevolucioncompra.Subtotaliva19),
		Quitacoma(tempDevolucioncompra.Subtotaliva5),
		Quitacoma(tempDevolucioncompra.Subtotaliva0),
		Quitacoma(tempDevolucioncompra.Subtotalbase19),
		Quitacoma(tempDevolucioncompra.Subtotalbase5),
		Quitacoma(tempDevolucioncompra.Subtotalbase0),
		Quitacoma(tempDevolucioncompra.TotalIva),
		Quitacoma(tempDevolucioncompra.Total),
		Quitacoma(tempDevolucioncompra.Neto),
		tempDevolucioncompra.Items,
		Quitacoma(tempDevolucioncompra.PorcentajeRetencionFuente),
		Quitacoma(tempDevolucioncompra.TotalRetencionFuente),
		Quitacoma(tempDevolucioncompra.PorcentajeRetencionIca),
		Quitacoma(tempDevolucioncompra.TotalRetencionIca),
		tempDevolucioncompra.Formadepago,
		tempDevolucioncompra.Mediodepago,
		tempDevolucioncompra.Tercero,
		tempDevolucioncompra.Almacenista,
		tempDevolucioncompra.Compra,
		tempDevolucioncompra.Centro,
		tempDevolucioncompra.Tipo)

	if err != nil {
		panic(err)
	}

	// INSERTAR COMPROBANTE CONTABILIDAD
	var tempComprobante comprobante
	var tempComprobanteDetalle comprobantedetalle
	tempComprobante.Documento = "12"
	tempComprobante.Numero = tempDevolucioncompra.Codigo
	tempComprobante.Fecha = tempDevolucioncompra.Fecha
	tempComprobante.Fechaconsignacion = tempDevolucioncompra.Fecha
	tempComprobante.Debito = tempDevolucioncompra.Neto + ".00"
	tempComprobante.Credito = tempDevolucioncompra.Neto + ".00"
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
	if tempDevolucioncompra.Neto != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempDevolucioncompra.Neto)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Compracuentaproveedor
		tempComprobanteDetalle.Debito = tempDevolucioncompra.Neto
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteDevolucioncompra(tempComprobanteDetalle, tempComprobante, tempDevolucioncompra)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA DEBITO DESCUENTO
	//if (tempDevolucioncompra.Descuento!="0")	{
	//	fila=fila+1
	//	tempComprobanteDetalle.Fila=strconv.Itoa(fila)
	//	totalDebito+=Flotante(tempDevolucioncompra.Descuento)
	//	tempComprobanteDetalle.Cuenta  = parametrosinventario.Compradevolucioncuentadescuento
	//	tempComprobanteDetalle.Debito = tempDevolucioncompra.Descuento
	//	tempComprobanteDetalle.Credito = ""
	//	InsertaDetalleComprobanteDevolucioncompra(tempComprobanteDetalle,tempComprobante,tempDevolucioncompra)
	//	log.Println("debito linea" + fmt.Sprintf("%.2f",totalDebito))
	//}

	// INSERTAR CUENTA DEBITO RET. FTE.
	if tempDevolucioncompra.TotalRetencionFuente != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempDevolucioncompra.TotalRetencionFuente)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Compradevolucioncuentaretfte
		tempComprobanteDetalle.Debito = tempDevolucioncompra.TotalRetencionFuente
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteDevolucioncompra(tempComprobanteDetalle, tempComprobante, tempDevolucioncompra)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA DEBITO RET. ICA.
	if tempDevolucioncompra.TotalRetencionIca != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempDevolucioncompra.TotalRetencionIca)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Compradevolucioncuentaretica
		tempComprobanteDetalle.Debito = tempDevolucioncompra.TotalRetencionIca
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteDevolucioncompra(tempComprobanteDetalle, tempComprobante, tempDevolucioncompra)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA CREDITO IVA 19%
	if tempDevolucioncompra.Subtotaliva19 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempDevolucioncompra.Subtotaliva19)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Compradevolucioniva19
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = tempDevolucioncompra.Subtotaliva19
		InsertaDetalleComprobanteDevolucioncompra(tempComprobanteDetalle, tempComprobante, tempDevolucioncompra)
		log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito))
	}

	// INSERTAR CUENTA CREDITO IVA 5%
	if tempDevolucioncompra.Subtotaliva5 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempDevolucioncompra.Subtotaliva5)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Compradevolucioniva5
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = tempDevolucioncompra.Subtotaliva5
		InsertaDetalleComprobanteDevolucioncompra(tempComprobanteDetalle, tempComprobante, tempDevolucioncompra)
		log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito))
	}

	// INSERTAR CUENTA CREDITO COMPRAS 19%
	if tempDevolucioncompra.Subtotal19 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempDevolucioncompra.Subtotal19) - Flotante(tempDevolucioncompra.Subtotaldescuento19)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Compradevolucioncuenta19
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(tempDevolucioncompra.Subtotal19) - Flotante(tempDevolucioncompra.Subtotaldescuento19))
		InsertaDetalleComprobanteDevolucioncompra(tempComprobanteDetalle, tempComprobante, tempDevolucioncompra)
		log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito))
	}

	// INSERTAR CUENTA CREDITO COMPRAS 5%
	if tempDevolucioncompra.Subtotal5 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempDevolucioncompra.Subtotal5) - Flotante(tempDevolucioncompra.Subtotaldescuento5)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Compradevolucioncuenta5
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(tempDevolucioncompra.Subtotal5) - Flotante(tempDevolucioncompra.Subtotaldescuento5))
		InsertaDetalleComprobanteDevolucioncompra(tempComprobanteDetalle, tempComprobante, tempDevolucioncompra)
		log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito))
	}

	// INSERTAR CUENTA CREDITO COMPRAS 0%
	if tempDevolucioncompra.Subtotal0 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempDevolucioncompra.Subtotal0) - Flotante(tempDevolucioncompra.Subtotaldescuento0)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Compradevolucioncuenta0
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(tempDevolucioncompra.Subtotal0) - Flotante(tempDevolucioncompra.Subtotaldescuento0))
		InsertaDetalleComprobanteDevolucioncompra(tempComprobanteDetalle, tempComprobante, tempDevolucioncompra)
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
	q += " ) "

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

// INICIA DEVOLUCIONCOMPRA EXISTE
func DevolucioncompraExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM devolucioncompra  WHERE codigo=$1", Codigo)
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

// INICIA DEVOLUCION COMPRA EDITAR
func DevolucioncompraEditar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio compra editar" + Codigo)
	db := dbConn()

	// TRAER DEVOLUCION COMPRA
	v := devolucioncompra{}
	err := db.Get(&v, "SELECT * FROM devolucioncompra where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []devolucioncompradetalleeditar{}

	err2 := db.Select(&det, DevolucioncompraConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
	v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
	v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
	v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
	v.CentroDetalle = TraerCentroConsulta(v.Centro)

	parametros := map[string]interface{}{
		"miperiodo":        periodoSesion(r),
		"devolucioncompra": v,
		"detalle":          det,
		"hosting":          ruta,
		"retfte":           TraerParametrosInventario().Compracuentaporcentajeretfte,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/devolucioncompra/devolucioncompraEditar.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaBodega.html",
		"vista/autocompleta/autocompletaTercero.html",
		"vista/devolucioncompra/devolucioncompraScript.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error devolucioncompra nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA DEVOLUCION COMPRA BORRAR
func DevolucioncompraBorrar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio devolucioncompra editar" + Codigo)

	db := dbConn()

	// traer DEVOLUCION COMPRA
	v := devolucioncompra{}
	err := db.Get(&v, "SELECT * FROM devolucioncompra where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []devolucioncompradetalleeditar{}
	err2 := db.Select(&det, DevolucioncompraConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
	v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
	v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
	v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
	v.CentroDetalle = TraerCentroConsulta(v.Centro)

	parametros := map[string]interface{}{
		"devolucioncompra": v,
		"detalle":          det,
		"tercero":          TraerTerceroConsulta(v.Tercero),
		"hosting":          ruta,
		"retfte":           TraerParametrosInventario().Compracuentaporcentajeretfte,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/devolucioncompra/devolucioncompraBorrar.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaBodega.html",
		"vista/devolucioncompra/devolucioncompraScript.html")

	log.Println("Error devolucioncompra nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA DEVOLUCION COMPRA ELIMINAR
func DevolucioncompraEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	codigo := mux.Vars(r)["codigo"]

	// borrar COMPRA
	delForm, err := db.Prepare("DELETE from devolucioncompra WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(codigo)

	// borar detalle
	delForm1, err := db.Prepare("DELETE from devolucioncompradetalle WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm1.Exec(codigo)

	// borar detalle inventario
	Borrarinventario(codigo, "Devolucioncompra")

	// borra detalle anterior
	delForm, err = db.Prepare("DELETE from comprobantedetalle WHERE documento=$1 and numero=$2")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec("12", codigo)

	// borra cabecera anterior

	delForm1, err = db.Prepare("DELETE from comprobante WHERE documento=$1 and numero=$2 ")
	if err != nil {
		panic(err.Error())
	}
	delForm1.Exec("12", codigo)

	log.Println("Registro Eliminado" + codigo)
	http.Redirect(w, r, "/DevolucioncompraLista", 301)
}

// TRAER COMPRA EN LA DEVOLUCION
func DatosCompra(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	Tercero := mux.Vars(r)["tercero"]

	log.Println("inicio compra editar" + Codigo)
	db := dbConn()
	var res []compra

	// TRAER COMPRA
	v := compra{}
	err := db.Get(&v, "SELECT * FROM compra where codigo=$1 and tercero=$2", Codigo, Tercero)
	var valida bool
	valida = true

	switch err {
	case nil:
		log.Printf("tercero found: %+v\n", v)
	case sql.ErrNoRows:
		log.Println("tercero NO Existe")
		valida = false
	default:
		log.Printf("tercero error: %s\n", err)
	}
	det := []compradetalleeditar{}
	t := tercero{}

	// trae datos si existe factura
	if valida == true {
		err2 := db.Select(&det, CompraConsultaDetalle(), Codigo)
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

// INICIA DEVOLUCION COMPRA PDF
func DevolucioncompraPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]

	// traer COMPRA
	miDevolucioncompra := devolucioncompra{}
	err := db.Get(&miDevolucioncompra, "SELECT * FROM devolucioncompra where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}

	// traer detalle
	miDetalle := []devolucioncompradetalleeditar{}
	err2 := db.Select(&miDetalle, DevolucioncompraConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	// traer tercero
	miTercero := tercero{}
	err3 := db.Get(&miTercero, "SELECT * FROM tercero where codigo=$1", miDevolucioncompra.Tercero)
	if err3 != nil {
		log.Fatalln(err3)
	}

	// traer almacenista
	miAlmacenista := almacenista{}
	err4 := db.Get(&miAlmacenista, "SELECT * FROM almacenista where codigo=$1", miDevolucioncompra.Almacenista)
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

		// DEVOLUCION COMPRA NUMERO
		pdf.SetY(20)
		pdf.SetX(80)
		pdf.Ln(8)
		pdf.SetX(75)
		pdf.SetFont("Arial", "", 11)
		pdf.CellFormat(190, 10, "DEVOLUCION COMPRA", "0", 0, "C", false, 0, "")
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
	DevolucioncompraCabecera(pdf, miTercero, miDevolucioncompra, miAlmacenista)

	var filas = len(miDetalle)
	// menos de 32
	if filas <= 32 {
		for i, miFila := range miDetalle {
			var a = i + 1
			DevolucioncompraFilaDetalle(pdf, miFila, a)
		}
		DevolucioncompraPieDePagina(pdf, miTercero, miDevolucioncompra)
	} else {
		// mas de 32 y menos de 73
		if filas > 32 && filas <= 73 {
			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 41 {
					DevolucioncompraFilaDetalle(pdf, miFila, a)
				}
			}
			DevolucioncompraLinea(pdf, miTercero, miDevolucioncompra)
			// segunda pagina
			pdf.AddPage()
			DevolucioncompraCabecera(pdf, miTercero, miDevolucioncompra, miAlmacenista)
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 41 {
					DevolucioncompraFilaDetalle(pdf, miFila, a)
				}
			}

			DevolucioncompraPieDePagina(pdf, miTercero, miDevolucioncompra)
		} else {
			// mas de tres paginas

			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 41 {
					DevolucioncompraFilaDetalle(pdf, miFila, a)
				}
			}
			DevolucioncompraLinea(pdf, miTercero, miDevolucioncompra)
			pdf.AddPage()
			DevolucioncompraCabecera(pdf, miTercero, miDevolucioncompra, miAlmacenista)
			// segunda pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 41 && a <= 82 {
					DevolucioncompraFilaDetalle(pdf, miFila, a)
				}
			}
			DevolucioncompraLinea(pdf, miTercero, miDevolucioncompra)
			pdf.AddPage()
			DevolucioncompraCabecera(pdf, miTercero, miDevolucioncompra, miAlmacenista)
			// tercera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 82 {
					DevolucioncompraFilaDetalle(pdf, miFila, a)
				}
			}

			DevolucioncompraPieDePagina(pdf, miTercero, miDevolucioncompra)
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

func DevolucioncompraCabecera(pdf *gofpdf.Fpdf, miTercero tercero, miDevolucioncompra devolucioncompra, miAlmacenista almacenista) {
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
	pdf.CellFormat(40, 4, Titulo(TraerFormadepago(miDevolucioncompra.Formadepago)), "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Fecha de Expedicion", "", 0,
		"L", false, 0, "")
	pdf.SetX(150)
	pdf.CellFormat(40, 4, miDevolucioncompra.Fecha.Format("02/01/2006")+" "+Titulo(miDevolucioncompra.Hora), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)

	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Medio de Pago", "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, TraerMediodepago(miDevolucioncompra.Mediodepago), "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Fecha de Vencimiento", "", 0,
		"L", false, 0, "")
	pdf.SetX(150)
	pdf.CellFormat(40, 4, miDevolucioncompra.Vence.Format("02/01/2006"), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Factura No.", "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, miDevolucioncompra.Compra, "", 0,
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
	pdf.CellFormat(190, 6, "PRODUCTO", "0", 0,
		"L", false, 0, "")
	pdf.SetX(106)
	pdf.CellFormat(40, 6, "UNIDAD", "0", 0,
		"L", false, 0, "")
	pdf.SetX(122)
	pdf.CellFormat(40, 6, "IVA", "0", 0,
		"L", false, 0, "")
	pdf.SetX(135)
	pdf.CellFormat(40, 6, "CANTIDAD", "0", 0,
		"L", false, 0, "")
	pdf.SetX(160)
	pdf.CellFormat(40, 6, "P. UNITARIO", "0", 0,
		"L", false, 0, "")
	pdf.SetX(190)
	pdf.CellFormat(40, 6, "TOTAL", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}
func DevolucioncompraFilaDetalle(pdf *gofpdf.Fpdf, miFila devolucioncompradetalleeditar, a int) {
	ene := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetX(21)
	pdf.CellFormat(183, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(40, 4, Subcadena(miFila.Producto, 0, 12), "", 0,
		"L", false, 0, "")
	pdf.SetX(55)
	pdf.CellFormat(40, 4, ene(Subcadena(miFila.ProductoNombre, 0, 32)), "", 0,
		"L", false, 0, "")
	pdf.SetX(82)
	pdf.CellFormat(40, 4, miFila.ProductoUnidad, "", 0,
		"R", false, 0, "")
	pdf.SetX(89)
	pdf.CellFormat(40, 4, miFila.ProductoIva, "", 0,
		"R", false, 0, "")
	pdf.SetX(115)
	pdf.CellFormat(40, 4, Coma(miFila.Cantidad), "", 0,
		"R", false, 0, "")
	pdf.SetX(140)
	pdf.CellFormat(40, 4, Coma(miFila.Precio), "", 0,
		"R", false, 0, "")
	pdf.SetX(164)
	pdf.CellFormat(40, 4, Coma(miFila.Subtotal), "", 0,
		"R", false, 0, "")
	pdf.Ln(4)
}

func DevolucioncompraPieDePagina(pdf *gofpdf.Fpdf, miTercero tercero, miDevolucioncompra devolucioncompra) {

	Totalletras, err := IntLetra(Cadenaentero(miDevolucioncompra.Neto))
	if err != nil {
		fmt.Println(err)
	}

	ene := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetFont("Arial", "", 8)
	pdf.SetY(221)
	pdf.SetX(20)
	pdf.MultiCell(184, 4, "SON: "+ene(Mayuscula(Totalletras))+" PESOS MDA. CTE.", "0", "L", false)
	pdf.SetFont("Arial", "", 9)
	pdf.SetY(254)
	pdf.Line(45, 254, 125, 254)
	pdf.Ln(1)
	pdf.SetX(65)
	pdf.CellFormat(40, 4, "FIRMA RESPONSABLE", "0", 0, "C",
		false, 0, "")

	// DATOS CON VALORES //
	pdf.SetFont("Arial", "", 9)
	var separador float64
	var altoseparador float64
	separador = 254
	altoseparador = -4

	// INICIA DATOS FACTURA
	if miDevolucioncompra.Neto != "0" {
		pdf.SetFont("Arial", "", 9)
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "TOTAL", "0", 0, "L",
			false, 0, "")
		pdf.SetX(13)
		pdf.CellFormat(190, 4, Coma(miDevolucioncompra.Neto), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucioncompra.TotalRetencionIca != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "RET. ICA.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, miDevolucioncompra.PorcentajeRetencionIca+"%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(13)
		pdf.CellFormat(190, 4, Coma(miDevolucioncompra.TotalRetencionIca), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucioncompra.TotalRetencionFuente != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "RET. FTE.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, miDevolucioncompra.PorcentajeRetencionFuente+"%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(13)
		pdf.CellFormat(190, 4, Coma(miDevolucioncompra.TotalRetencionFuente), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucioncompra.Subtotaliva5 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "I. V. A.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, "5%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(13)
		pdf.CellFormat(190, 4, Coma(miDevolucioncompra.Subtotaliva5), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucioncompra.Subtotaliva19 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "I. V. A.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, "19%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(13)
		pdf.CellFormat(190, 4, Coma(miDevolucioncompra.Subtotaliva19), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucioncompra.Subtotalbase5 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "GRAVABLE", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, "5%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(13)
		pdf.CellFormat(190, 4, Coma(miDevolucioncompra.Subtotalbase5), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucioncompra.Subtotalbase19 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "GRAVABLE", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, "19%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(13)
		pdf.CellFormat(190, 4, Coma(miDevolucioncompra.Subtotalbase19), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucioncompra.Subtotalbase0 != "0" {

		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "NO GRAVABLE", "0", 0, "L",
			false, 0, "")
		pdf.SetX(13)
		pdf.CellFormat(190, 4, Coma(miDevolucioncompra.Subtotalbase0), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}
}

func DevolucioncompraLinea(pdf *gofpdf.Fpdf, miTercero tercero, miDevolucioncompra devolucioncompra) {
	pdf.SetY(259)
	pdf.Line(20, 259, 204, 259)
}

// INICIA DEVOLUCION COMPRA TODOS PDF
func DevolucioncompraTodosCabecera(pdf *gofpdf.Fpdf) {
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

func DevolucioncompraTodosDetalle(pdf *gofpdf.Fpdf, miFila devolucioncompralista, a int) {
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

func DevolucioncompraTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string
	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre,centro.nombre as CentroNombre,devolucioncompra.neto,devolucioncompra.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	consulta += " FROM devolucioncompra "
	consulta += " inner join tercero on tercero.codigo=devolucioncompra.tercero "
	consulta += " inner join centro on centro.codigo=devolucioncompra.centro "
	consulta += " inner join almacenista on almacenista.codigo=devolucioncompra.almacenista "
	consulta += " ORDER BY cast(devolucioncompra.codigo as integer) ASC"

	t := []devolucioncompralista{}
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
		pdf.CellFormat(190, 10, "DATOS DEVOLUCION COMPRA", "0", 0,
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

	DevolucioncompraTodosCabecera(pdf)
	// tercera pagina

	for i, miFila := range t {
		var a = i + 1
		if math.Mod(float64(a), 49) == 0 {
			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			DevolucioncompraTodosCabecera(pdf)
		}
		DevolucioncompraTodosDetalle(pdf, miFila, a)
	}

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

// TERMINA COMPRA TODOS PDF

// COMPRA EXCEL
func DevolucioncompraExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string
	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre,centro.nombre as CentroNombre,devolucioncompra.neto,devolucioncompra.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	consulta += " FROM devolucioncompra "
	consulta += " inner join tercero on tercero.codigo=devolucioncompra.tercero "
	consulta += " inner join centro on centro.codigo=devolucioncompra.centro "
	consulta += " inner join almacenista on almacenista.codigo=devolucioncompra.almacenista "
	consulta += " ORDER BY cast(devolucioncompra.codigo as integer) ASC"
	t := []devolucioncompralista{}
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
	f.SetCellValue("Sheet1", "A9", "LISTADO DE DEVOLUCION COMPRAS")
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

package main

// INICIA DEVOLUCIONVENTA IMPORTAR PAQUETES
import (
	"bytes"
	"database/sql"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	_ "github.com/bitly/go-simplejson"
	"github.com/dustin/go-humanize"
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

// INICIA DEVOLUCIONVENTA ESTRUCTURA JSON
type devolucionventaJson struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Nombre string `json:"nombre"`
}

// INICIA DEVOLUCIONVENTA ESTRUCTURA
type devolucionventaLista struct {
	Codigo         string
	Fecha          time.Time
	Total          string
	Tercero        string
	TerceroNombre  string
	CentroNombre   string
	VendedorNombre string
}

// INICIA DEVOLUCIONVENTA ESTRUCTURA
type devolucionventa struct {
	Resolucion          string
	Codigo              string
	Fecha               time.Time
	Vence               time.Time
	Hora                string
	Descuento           string
	Subtotaldescuento19 string
	Subtotaldescuento5  string
	Subtotaldescuento0  string
	Subtotal            string
	Subtotal19          string
	Subtotal5           string
	Subtotal0           string
	Subtotaliva19       string
	Subtotaliva5        string
	Subtotaliva0        string
	Subtotalbase19      string
	Subtotalbase5       string
	Subtotalbase0       string
	TotalIva            string
	Total               string
	Neto                string
	Items               string
	Formadepago         string
	Mediodepago         string
	Tercero             string
	Vendedor            string
	Accion              string
	Detalle             []devolucionventadetalle       `json:"Detalle"`
	DetalleEditar       []devolucionventadetalleeditar `json:"DetalleEditar"`
	TerceroDetalle      tercero
	Venta               string
	Tipo                string
	Ret2201             string
	Centro              string
	FormadepagoDetalle  formadepago
	MediodepagoDetalle  mediodepago
	VendedorDetalle     vendedor
	CentroDetalle       centro
}

// INICIA DEVOLUCIONVENTADETALLE ESTRUCTURA
type devolucionventadetalle struct {
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

// INICIA COMPRA DETALLE EDITARr
type devolucionventadetalleeditar struct {
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
func DevolucionventaConsultaDetalle() string {
	var consulta = ""
	consulta = "select "
	consulta += "devolucionventadetalle.Id as id ,"
	consulta += "devolucionventadetalle.Codigo as codigo,"
	consulta += "devolucionventadetalle.Fila as fila,"
	consulta += "devolucionventadetalle.Cantidad as cantidad,"
	consulta += "devolucionventadetalle.Precio as precio,"
	consulta += "devolucionventadetalle.Descuento as descuento,"
	consulta += "devolucionventadetalle.Montodescuento as montodescuento,"
	consulta += "devolucionventadetalle.Sigratis as sigratis,"
	consulta += "devolucionventadetalle.Subtotal as subtotal,"
	consulta += "devolucionventadetalle.Subtotaldescuento  as subtotaldescuento,"
	consulta += "devolucionventadetalle.Pagina as pagina ,"
	consulta += "devolucionventadetalle.Bodega as bodega,"
	consulta += "bodega.Nombre as bodeganombre,"
	consulta += "devolucionventadetalle.Producto as producto,"
	consulta += "devolucionventadetalle.Fecha as fecha,"
	consulta += "producto.nombre as ProductoNombre, "
	consulta += "producto.iva as ProductoIva, "
	consulta += "producto.unidad as ProductoUnidad "
	consulta += "from devolucionventadetalle "
	consulta += "inner join producto on producto.codigo=devolucionventadetalle.producto "
	consulta += "inner join bodega on bodega.codigo=devolucionventadetalle.bodega "
	consulta += " where devolucionventadetalle.codigo=$1"
	log.Println(consulta)
	return consulta
}

// INICIA DEVOLUCIONVENTA LISTA
func DevolucionventaLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/devolucionventa/devolucionventaLista.html")
	log.Println("Error devolucionventa 0")
	var consulta string
	var miperiodo = periodoSesion(r)

	consulta = "  SELECT vendedor.nombre as VendedorNombre,centro.nombre as CentroNombre,total,devolucionventa.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	consulta += " FROM devolucionventa "
	consulta += " inner join tercero on tercero.codigo=devolucionventa.tercero "
	consulta += " inner join centro on centro.codigo=devolucionventa.centro "
	consulta += " inner join vendedor on vendedor.codigo=devolucionventa.vendedor "
	consulta += " where  extract(year from fecha)=" + miperiodo + " "
	consulta += " ORDER BY devolucionventa.codigo ASC"

	db := dbConn()
	res := []devolucionventaLista{}

	err := db.Select(&res, consulta)

	if err != nil {
		fmt.Println(err)
	}

	varmap := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
	}
	log.Println("Error devolucionventa888")
	tmp.Execute(w, varmap)
}

// INICIA DEVOLUCIONVENTA NUEVO
func DevolucionventaNuevo(w http.ResponseWriter, r *http.Request) {

	// TRAE COPIA DE EDITAR
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio devolucionventa editar" + Codigo)

	db := dbConn()
	v := devolucionventa{}
	det := []devolucionventadetalleeditar{}
	if Codigo == "False" {

	} else {

		err := db.Get(&v, "SELECT * FROM devolucionventa where codigo=$1", Codigo)
		if err != nil {
			log.Fatalln(err)
		}

		v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
		v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
		v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
		v.VendedorDetalle = TraerVendedorConsulta(v.Vendedor)
		v.CentroDetalle = TraerCentroConsulta(v.Centro)
		err2 := db.Select(&det, DevolucionventaConsultaDetalle(), Codigo)
		if err2 != nil {
			fmt.Println(err2)
		}
	}

	parametros := map[string]interface{}{
		"miperiodo":       periodoSesion(r),
		"resolucionventa": ListaResolucionventa(),
		"codigo":          Codigo,
		"vendedor":        PrimerVendedor(),
		"devolucionventa": v,
		"detalle":         det,
		"hosting":         ruta,
		"ventatipoiva":    TraerParametrosInventario().Ventatipoiva,
		"autoret2201":     TraerParametrosInventario().Ventacuentaporcentajeret2201,
	}
	//TERMINA TRAE COPIA DE EDITAR

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/devolucionventa/devolucionventaNuevo.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaVendedor.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaBodega.html",
		"vista/autocompleta/autocompletaTercero.html",
		"vista/devolucionventa/devolucionventaScript.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error devolucionventa nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA INSERTAR COMPROBANTE DE DEVOLUCION VENTA
func InsertaDetalleComprobanteDevolucionventaVenta(miFilaComprobante comprobantedetalle, miComprobante comprobante, miDevolucionventa devolucionventa) {
	db := dbConn()

	// traer tercero
	miTercero := tercero{}
	err3 := db.Get(&miTercero, "SELECT * FROM tercero where codigo=$1", miDevolucionventa.Tercero)
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
		miFilaComprobante.Debito = miFilaComprobante.Debito + ".00"
	}

	if len(miFilaComprobante.Credito) > 0 {
		miFilaComprobante.Credito = miFilaComprobante.Credito + ".00"
	}

	// TERMINA COMPROBANTE GRABAR INSERTAR
	_, err = insForm.Exec(
		miFilaComprobante.Fila,
		miFilaComprobante.Cuenta,
		miDevolucionventa.Tercero,
		miDevolucionventa.Centro,
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

// INICIA DEVOLUCIONVENTA INSERTAR AJAX
func DevolucionventaAgregar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var tempDevolucionventa devolucionventa

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// carga informacion de la Devolucionventa
	err = json.Unmarshal(b, &tempDevolucionventa)
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
	Codigoactual = tempDevolucionventa.Codigo
	if tempDevolucionventa.Accion == "Actualizar" {

		// borra detalle anterior
		delForm, err := db.Prepare("DELETE from devolucionventadetalle WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(tempDevolucionventa.Codigo)

		// borra detalle inventario
		Borrarinventario(tempDevolucionventa.Codigo, "Devolucionventa")

		// borra cabecera anterior
		delForm1, err := db.Prepare("DELETE from devolucionventa WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm1.Exec(tempDevolucionventa.Codigo)
	}

	// INSERTA DETALLE
	for i, x := range tempDevolucionventa.Detalle {
		var a = i
		var q string
		q = "insert into devolucionventadetalle ("
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

		// TERMINA DEVOLUCIONVENTA GRABAR INSERTAR
		_, err = insForm.Exec(
			x.Id,
			Codigoactual,
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
	for _, x := range tempDevolucionventa.Detalle {
		miInventario = append(miInventario, inventario{
			x.Fecha, x.Tipo, x.Codigo, x.Bodega,
			x.Producto,
			Quitacoma(x.Cantidad),
			Quitacoma(x.Precio),
			operacionDevolucionVenta})
	}
	// insertar inventario

	InsertaInventario(miInventario)

	// INICIA INSERTAR VENTAS
	log.Println("Got %s age %s club %s\n", tempDevolucionventa.Codigo, tempDevolucionventa.Tercero, tempDevolucionventa.Total)
	var q string
	q = "insert into devolucionventa ("
	q += "Resolucion,"
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
	q += "Ret2201,"
	q += "Total,"
	q += "Neto,"
	q += "Items,"
	q += "Formadepago,"
	q += "Mediodepago,"
	q += "Tercero,"
	q += "Vendedor,"
	q += "Venta,"
	q += "Centro,"
	q += "Tipo"
	q += " ) values("
	q += parametros(31)
	q += " ) "

	log.Println("Cadena SQL " + q)
	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	layout := "2006-01-02"
	log.Println("Hora", tempDevolucionventa.Fecha.Format("02/01/2006"))
	currentTime := time.Now()

	_, err = insForm.Exec(
		tempDevolucionventa.Resolucion,
		tempDevolucionventa.Codigo,
		tempDevolucionventa.Fecha.Format(layout),
		tempDevolucionventa.Vence.Format(layout),
		currentTime.Format("3:4:5 PM"),
		Quitacoma(tempDevolucionventa.Descuento),
		Quitacoma(tempDevolucionventa.Subtotaldescuento19),
		Quitacoma(tempDevolucionventa.Subtotaldescuento5),
		Quitacoma(tempDevolucionventa.Subtotaldescuento0),
		Quitacoma(tempDevolucionventa.Subtotal),
		Quitacoma(tempDevolucionventa.Subtotal19),
		Quitacoma(tempDevolucionventa.Subtotal5),
		Quitacoma(tempDevolucionventa.Subtotal0),
		Quitacoma(tempDevolucionventa.Subtotaliva19),
		Quitacoma(tempDevolucionventa.Subtotaliva5),
		Quitacoma(tempDevolucionventa.Subtotaliva0),
		Quitacoma(tempDevolucionventa.Subtotalbase19),
		Quitacoma(tempDevolucionventa.Subtotalbase5),
		Quitacoma(tempDevolucionventa.Subtotalbase0),
		Quitacoma(tempDevolucionventa.TotalIva),
		Quitacoma(tempDevolucionventa.Ret2201),
		Quitacoma(tempDevolucionventa.Total),
		Quitacoma(tempDevolucionventa.Neto),
		tempDevolucionventa.Items,
		tempDevolucionventa.Formadepago,
		tempDevolucionventa.Mediodepago,
		tempDevolucionventa.Tercero,
		tempDevolucionventa.Vendedor,
		tempDevolucionventa.Venta,
		tempDevolucionventa.Centro,
		tempDevolucionventa.Tipo)

	if err != nil {
		panic(err)
	}

	// INSERTAR COMPROBANTE CONTABILIDAD
	var tempComprobante comprobante
	var tempComprobanteDetalle comprobantedetalle
	tempComprobante.Documento = "11"
	tempComprobante.Numero = tempDevolucionventa.Codigo
	tempComprobante.Fecha = tempDevolucionventa.Fecha
	tempComprobante.Fechaconsignacion = tempDevolucionventa.Fecha
	tempComprobante.Debito = tempDevolucionventa.Neto + ".00"
	tempComprobante.Credito = tempDevolucionventa.Neto + ".00"
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

	// INSERTAR CUENTA DEBITO RET. 2201
	if tempDevolucionventa.Ret2201 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempDevolucionventa.Ret2201)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Ventacontracuentaret2201
		tempComprobanteDetalle.Debito = tempDevolucionventa.Ret2201
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteDevolucionventaVenta(tempComprobanteDetalle, tempComprobante, tempDevolucionventa)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA CREDITO IVA 19%
	if tempDevolucionventa.Subtotaliva19 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempDevolucionventa.Subtotaliva19)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Ventadevolucioniva19
		tempComprobanteDetalle.Debito = tempDevolucionventa.Subtotaliva19
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteDevolucionventaVenta(tempComprobanteDetalle, tempComprobante, tempDevolucionventa)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA CREDITO IVA 5%
	if tempDevolucionventa.Subtotaliva5 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempDevolucionventa.Subtotaliva5)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Ventadevolucioniva5
		tempComprobanteDetalle.Debito = tempDevolucionventa.Subtotaliva5
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteDevolucionventaVenta(tempComprobanteDetalle, tempComprobante, tempDevolucionventa)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA CREDITO VENTA IVA 19%
	if tempDevolucionventa.Subtotalbase19 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempDevolucionventa.Subtotalbase19)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Ventadevolucioncuenta19
		tempComprobanteDetalle.Debito = humanize.Comma(int64(Flotante(tempDevolucionventa.Subtotalbase19)))
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteDevolucionventaVenta(tempComprobanteDetalle, tempComprobante, tempDevolucionventa)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA CREDITO VENTA IVA 5%
	if tempDevolucionventa.Subtotalbase5 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempDevolucionventa.Subtotalbase5)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Ventadevolucioncuenta5
		tempComprobanteDetalle.Debito = humanize.Comma(int64(Flotante(tempDevolucionventa.Subtotalbase5)))
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteDevolucionventaVenta(tempComprobanteDetalle, tempComprobante, tempDevolucionventa)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA CREDITO VENTA IVA 0%
	if tempDevolucionventa.Subtotalbase0 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempDevolucionventa.Subtotalbase0)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Ventadevolucioncuenta0
		tempComprobanteDetalle.Debito = humanize.Comma(int64(Flotante(tempDevolucionventa.Subtotalbase0)))
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteDevolucionventaVenta(tempComprobanteDetalle, tempComprobante, tempDevolucionventa)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA DEBITO RET. 2201
	if tempDevolucionventa.Ret2201 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempDevolucionventa.Ret2201)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Ventacuentaret2201
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = tempDevolucionventa.Ret2201
		InsertaDetalleComprobanteDevolucionventaVenta(tempComprobanteDetalle, tempComprobante, tempDevolucionventa)
		log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito))
	}

	// INSERTAR CUENTA DEBITO DESCUENTO
	//if (tempDevolucionventa.Descuento!="0")	{
	//	fila=fila+1
	//	tempComprobanteDetalle.Fila=strconv.Itoa(fila)
	//	totalCredito+=Flotante(tempDevolucionventa.Descuento)
	//	tempComprobanteDetalle.Cuenta  = parametrosinventario.Ventadevolucioncuentadescuento
	//	tempComprobanteDetalle.Debito = ""
	//	tempComprobanteDetalle.Credito = tempDevolucionventa.Descuento
	//	InsertaDetalleComprobanteDevolucionventaVenta(tempComprobanteDetalle,tempComprobante,tempDevolucionventa)
	//	log.Println("credito linea" + fmt.Sprintf("%.2f",totalCredito))
	//}

	// INSERTAR CUENTA DEBITO CLIENTE
	if tempDevolucionventa.Neto != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempDevolucionventa.Neto)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Ventacuentacliente
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = tempDevolucionventa.Neto
		InsertaDetalleComprobanteDevolucionventaVenta(tempComprobanteDetalle, tempComprobante, tempDevolucionventa)
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

// INICIA DEVOLUCIONVENTA EXISTE
func DevolucionventaExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM devolucionventa  WHERE codigo=$1", Codigo)
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

// INICIA DEVOLUCIONVENTA EDITAR
func DevolucionventaEditar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio devolucionventa editar" + Codigo)
	db := dbConn()

	// traer devolucionventa
	v := devolucionventa{}
	err := db.Get(&v, "SELECT * FROM devolucionventa where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []devolucionventadetalleeditar{}

	err2 := db.Select(&det, DevolucionventaConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
	v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
	v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
	v.VendedorDetalle = TraerVendedorConsulta(v.Vendedor)
	v.CentroDetalle = TraerCentroConsulta(v.Centro)

	parametros := map[string]interface{}{
		"miperiodo":       periodoSesion(r),
		"resolucionventa": ListaResolucionventa(),
		"devolucionventa": v,
		"detalle":         det,
		"hosting":         ruta,
		"ventatipoiva":    TraerParametrosInventario().Ventatipoiva,
		"autoret2201":     TraerParametrosInventario().Ventacuentaporcentajeret2201,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/devolucionventa/devolucionventaEditar.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaVendedor.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaBodega.html",
		"vista/autocompleta/autocompletaTercero.html",
		"vista/devolucionventa/devolucionventaScript.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error devolucionventa nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA DEVOLUCIONVENTA BORRAR
func DevolucionventaBorrar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio devolucionventa editar" + Codigo)

	db := dbConn()

	// traer devolucionventa
	v := devolucionventa{}
	err := db.Get(&v, "SELECT * FROM devolucionventa where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []devolucionventadetalleeditar{}
	err2 := db.Select(&det, DevolucionventaConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
	v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
	v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
	v.VendedorDetalle = TraerVendedorConsulta(v.Vendedor)
	v.CentroDetalle = TraerCentroConsulta(v.Centro)

	parametros := map[string]interface{}{
		"resolucionventa": ListaResolucionventa(),
		"devolucionventa": v,
		"detalle":         det,
		"hosting":         ruta,
		"ventatipoiva":    TraerParametrosInventario().Ventatipoiva,
		"autoret2201":     TraerParametrosInventario().Ventacuentaporcentajeret2201,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/devolucionventa/devolucionventaBorrar.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaBodega.html",
		"vista/devolucionventa/devolucionventaScript.html")

	log.Println("Error devolucionventa nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA DEVOLUCIONVENTA ELIMINAR
func DevolucionventaEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	codigo := mux.Vars(r)["codigo"]

	// borrar Devolucionventa
	delForm, err := db.Prepare("DELETE from devolucionventa WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(codigo)

	// borar detalle
	delForm1, err := db.Prepare("DELETE from devolucionventadetalle WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm1.Exec(codigo)

	// borar detalle inventario
	Borrarinventario(codigo, "Devolucionventa")

	// borra detalle anterior
	delForm, err = db.Prepare("DELETE from comprobantedetalle WHERE documento=$1 and numero=$2")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec("11", codigo)

	// borra cabecera anterior

	delForm1, err = db.Prepare("DELETE from comprobante WHERE documento=$1 and numero=$2 ")
	if err != nil {
		panic(err.Error())
	}
	delForm1.Exec("11", codigo)

	log.Println("Registro Eliminado" + codigo)
	http.Redirect(w, r, "/DevolucionventaLista", 301)
}

// TRAER VENTA EN LA DEVOLUCION
func DatosVenta(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	Tercero := mux.Vars(r)["tercero"]
	log.Println("inicio venta editar" + Codigo)
	db := dbConn()
	var res []venta

	// traer VENTA
	v := venta{}
	err := db.Get(&v, "SELECT * FROM venta where codigo=$1 and Tercero = $2", Codigo, Tercero)
	var valida bool
	valida = true

	switch err {
	case nil:
		log.Printf("venta no existe: %+v\n", v)
	case sql.ErrNoRows:
		log.Println("venta No Existe")
		valida = false
	default:
		log.Printf("venta error: %s\n", err)
	}
	det := []ventadetalleeditar{}
	t := tercero{}

	// trae datos si existe factura
	if valida == true {
		err2 := db.Select(&det, VentaConsultaDetalle(), Codigo)
		if err2 != nil {
			fmt.Println(err2)
		}
		v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
		v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
		v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
		v.VendedorDetalle = TraerVendedorConsulta(v.Vendedor)
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

// INICIA DEVOLUCION EN VENTA PDF
func DevolucionventaPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]

	// traer Devolucionventa
	miDevolucionventa := devolucionventa{}
	err := db.Get(&miDevolucionventa, "SELECT * FROM devolucionventa where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}

	// traer detalle
	miDetalle := []devolucionventadetalleeditar{}
	err2 := db.Select(&miDetalle, DevolucionventaConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	// traer tercero
	miTercero := tercero{}
	err3 := db.Get(&miTercero, "SELECT * FROM tercero where codigo=$1", miDevolucionventa.Tercero)
	if err3 != nil {
		log.Fatalln(err3)
	}

	// traer Vendedor
	miVendedor := vendedor{}
	err4 := db.Get(&miVendedor, "SELECT * FROM vendedor where codigo=$1", miDevolucionventa.Vendedor)
	if err4 != nil {
		log.Fatalln(err4)
	}

	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	var re Resolucionventa = TraerResolucionventa(miDevolucionventa.Resolucion)
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
		pdf.CellFormat(190, 10, "Nit No. "+Coma(e.Codigo)+" - "+e.Dv, "0", 0, "C", false, 0, "")
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
		pdf.SetFont("Arial", "", 9)
		pdf.SetY(20)
		pdf.SetX(80)
		pdf.CellFormat(190, 10, "Resolucion No. "+re.Numero, "0", 0, "C",
			false, 0, "")
		pdf.Ln(4)
		pdf.SetX(80)
		pdf.CellFormat(190, 10, "Del Numero "+re.Prefijo+" "+re.NumeroInicial+" al "+re.Prefijo+" "+re.NumeroFinal, "0", 0, "C",
			false, 0, "")
		pdf.Ln(4)
		pdf.SetX(80)
		pdf.CellFormat(190, 10, "Vigencia del "+re.FechaInicial.Format("02/01/2006")+" al "+re.FechaFinal.Format("02/01/2006"), "0", 0, "C",
			false, 0, "")
		pdf.Ln(8)
		pdf.SetX(80)
		pdf.SetFont("Arial", "", 11)
		pdf.CellFormat(190, 10, "DEVOLUCION FACTURA", "0", 0, "C", false, 0, "")
		pdf.Ln(5)
		pdf.SetX(80)
		pdf.CellFormat(190, 10, " DE VENTA No. "+re.Prefijo+" "+miDevolucionventa.Codigo, "0", 0, "C",
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
	DevolucionventaCabecera(pdf, miTercero, miDevolucionventa, miVendedor)

	var filas = len(miDetalle)
	// UNA PAGINA
	if filas <= 32 {
		for i, miFila := range miDetalle {
			var a = i + 1
			DevolucionventaFilaDetalle(pdf, miFila, a)
		}
		DevolucionventaPieDePagina(pdf, miTercero, miDevolucionventa)
	} else {
		// mas de 32 y menos de 73
		if filas > 32 && filas <= 73 {
			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 41 {
					DevolucionventaFilaDetalle(pdf, miFila, a)
				}
			}
			DevolucionventaLinea(pdf)
			// segunda pagina
			pdf.AddPage()
			DevolucionventaCabecera(pdf, miTercero, miDevolucionventa, miVendedor)
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 41 {
					DevolucionventaFilaDetalle(pdf, miFila, a)
				}
			}

			DevolucionventaPieDePagina(pdf, miTercero, miDevolucionventa)
		} else {
			// mas de tres paginas

			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 41 {
					DevolucionventaFilaDetalle(pdf, miFila, a)
				}
			}
			DevolucionventaLinea(pdf)
			pdf.AddPage()
			DevolucionventaCabecera(pdf, miTercero, miDevolucionventa, miVendedor)
			// segunda pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 41 && a <= 82 {
					DevolucionventaFilaDetalle(pdf, miFila, a)
				}
			}
			DevolucionventaLinea(pdf)
			pdf.AddPage()
			DevolucionventaCabecera(pdf, miTercero, miDevolucionventa, miVendedor)
			// tercera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 82 {
					DevolucionventaFilaDetalle(pdf, miFila, a)
				}
			}
			DevolucionventaPieDePagina(pdf, miTercero, miDevolucionventa)
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

func DevolucionventaCabecera(pdf *gofpdf.Fpdf, miTercero tercero, miDevolucionventa devolucionventa, miVendedor vendedor) {
	ene := pdf.UnicodeTranslatorFromDescriptor("")
	// RELLENO TITULO
	pdf.SetFont("Arial", "", 10)
	pdf.SetY(44)
	pdf.SetX(20)
	pdf.SetFillColor(224, 231, 239)
	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(8)
	pdf.SetX(20)
	pdf.CellFormat(184, 6, "DATOS DEL ADQUIRIENTE", "0", 0,
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
	pdf.CellFormat(40, 4, Titulo(TraerFormadepago(miDevolucionventa.Formadepago)), "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Fecha de Expedicion", "", 0,
		"L", false, 0, "")
	pdf.SetX(150)
	pdf.CellFormat(40, 4, miDevolucionventa.Fecha.Format("02/01/2006")+" "+Titulo(miDevolucionventa.Hora), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)

	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Medio de Pago", "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, TraerMediodepago(miDevolucionventa.Mediodepago), "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Fecha de Vencimiento", "", 0,
		"L", false, 0, "")
	pdf.SetX(150)
	pdf.CellFormat(40, 4, miDevolucionventa.Vence.Format("02/01/2006"), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Factura No.", "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, miDevolucionventa.Venta, "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Vendedor", "", 0,
		"L", false, 0, "")
	pdf.SetX(130)
	pdf.CellFormat(40, 4, ene(miVendedor.Nombre), "", 0,
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
func DevolucionventaFilaDetalle(pdf *gofpdf.Fpdf, miFila devolucionventadetalleeditar, a int) {
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
	pdf.SetX(80)
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

func DevolucionventaPieDePagina(pdf *gofpdf.Fpdf, miTercero tercero, miDevolucionventa devolucionventa) {

	Totalletras, err := IntLetra(Cadenaentero(miDevolucionventa.Total))
	if err != nil {
		fmt.Println(err)
	}

	ene := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetFont("Arial", "", 8)
	pdf.SetY(221)
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
	pdf.Line(55, 250, 140, 250)
	pdf.Ln(9)
	pdf.SetX(80)
	pdf.CellFormat(40, 4, "A C E P T A D A ", "0", 0, "C",
		false, 0, "")

	// PRESENTA DATOS CON VALORES //
	pdf.SetFont("Arial", "", 9)
	var separador float64
	var altoseparador float64
	separador = 250
	altoseparador = -4

	// INICIA DATOS FACTURA
	if miDevolucionventa.Neto != "0" {
		pdf.SetFont("Arial", "", 9)
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "TOTAL", "0", 0, "L",
			false, 0, "")
		pdf.SetX(13)
		pdf.CellFormat(190, 4, Coma(miDevolucionventa.Neto), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucionventa.Subtotaliva5 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "I. V. A.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, "5%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(13)
		pdf.CellFormat(190, 4, Coma(miDevolucionventa.Subtotaliva5), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucionventa.Subtotaliva19 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "I. V. A.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, "19%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(13)
		pdf.CellFormat(190, 4, Coma(miDevolucionventa.Subtotaliva19), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucionventa.Subtotalbase5 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "GRAVABLE", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, "5%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(13)
		pdf.CellFormat(190, 4, Coma(miDevolucionventa.Subtotalbase5), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucionventa.Subtotalbase19 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "GRAVABLE", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, "19%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(13)
		pdf.CellFormat(190, 4, Coma(miDevolucionventa.Subtotalbase19), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucionventa.Subtotalbase0 != "0" {

		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "NO GRAVABLE", "0", 0, "L",
			false, 0, "")
		pdf.SetX(13)
		pdf.CellFormat(190, 4, Coma(miDevolucionventa.Subtotalbase0), "0", 0, "R",
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

func DevolucionventaLinea(pdf *gofpdf.Fpdf) {
	pdf.SetY(259)
	pdf.Line(20, 259, 204, 259)
}

// INICIA DEVOLUCION VENTA TODOS PDF
func DevolucionventaTodosCabecera(pdf *gofpdf.Fpdf) {
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
	pdf.CellFormat(190, 6, "Cliente", "0", 0,
		"L", false, 0, "")
	pdf.SetX(130)
	pdf.CellFormat(190, 6, "Vendedor", "0", 0,
		"L", false, 0, "")
	pdf.SetX(193)
	pdf.CellFormat(190, 6, "Total", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}

func DevolucionventaTodosDetalle(pdf *gofpdf.Fpdf, miFila devolucionventaLista, a int) {
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
	pdf.CellFormat(40, 4, Subcadena(miFila.VendedorNombre, 0, 31), "", 0,
		"L", false, 0, "")
	pdf.SetX(164)
	pdf.CellFormat(40, 4, Coma(miFila.Total), "", 0,
		"R", false, 0, "")
	pdf.Ln(4)
}

func DevolucionventaTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string
	consulta = "  SELECT vendedor.nombre as VendedorNombre,centro.nombre as CentroNombre,devolucionventa.total,devolucionventa.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	consulta += " FROM devolucionventa "
	consulta += " inner join tercero on tercero.codigo=devolucionventa.tercero "
	consulta += " inner join centro on centro.codigo=devolucionventa.centro "
	consulta += " inner join vendedor on vendedor.codigo=devolucionventa.vendedor "
	consulta += " ORDER BY cast(devolucionventa.codigo as integer) ASC"

	t := []devolucionventaLista{}
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
		pdf.CellFormat(190, 10, "DATOS DEVOLUCION VENTAS", "0", 0,
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

	CotizacionTodosCabecera(pdf)
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
		DevolucionventaTodosDetalle(pdf, miFila, a)
	}

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

// TERMINA DEVOLUCION VENTA TODOS PDF

// DEVOLUCION VENTA EXCEL
func DevolucionventaExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string
	consulta = "  SELECT vendedor.nombre as VendedorNombre,centro.nombre as CentroNombre,devolucionventa.total,devolucionventa.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	consulta += " FROM devolucionventa "
	consulta += " inner join tercero on tercero.codigo=devolucionventa.tercero "
	consulta += " inner join centro on centro.codigo=devolucionventa.centro "
	consulta += " inner join vendedor on vendedor.codigo=devolucionventa.vendedor "
	consulta += " ORDER BY cast(devolucionventa.codigo as integer) ASC"
	t := []devolucionventaLista{}
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
	f.SetCellValue("Sheet1", "A9", "LISTADO DE DEVOLUCION EN VENTAS")
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
	f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel), "Cliente")
	f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel), "Vendedor")
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
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel+i), miFila.VendedorNombre)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel+i), Flotante(miFila.Total))

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

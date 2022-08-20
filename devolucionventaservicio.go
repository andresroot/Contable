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
type devolucionventaservicioJson struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Nombre string `json:"nombre"`
}

// INICIA DEVOLUCIONVENTASERVICIO ESTRUCTURA
type devolucionventaservicioLista struct {
	Codigo         string
	Fecha          time.Time
	Total          string
	Tercero        string
	TerceroNombre  string
	CentroNombre   string
	VendedorNombre string
}

// INICIA DEVOLUCIONVENTASERVICIO ESTRUCTURA
type devolucionventaservicio struct {
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
	Detalle             []devolucionventaserviciodetalle       `json:"Detalle"`
	DetalleEditar       []devolucionventaserviciodetalleeditar `json:"DetalleEditar"`
	TerceroDetalle      tercero
	Ventaservicio       string
	Tipo                string
	Ret2201             string
	Centro              string
	FormadepagoDetalle  formadepago
	MediodepagoDetalle  mediodepago
	VendedorDetalle     vendedor
	CentroDetalle       centro
}

// INICIA DEVOLUCIONVENTASERVICIODETALLE ESTRUCTURA
type devolucionventaserviciodetalle struct {
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
	Nombreservicio    string
	Unidadservicio    string
	Codigoservicio    string
	Ivaservicio       string
	Tipo              string
	Fecha             time.Time
}

// INICIA COMPRA DETALLE EDITAR
type devolucionventaserviciodetalleeditar struct {
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
	Nombreservicio    string
	Unidadservicio    string
	Codigoservicio    string
	Ivaservicio       string
	Ivanombreservicio string
	Tipo              string
	Fecha             time.Time
}

// INICIA COMPRA CONSULTA DETALLE
func DevolucionventaservicioConsultaDetalle() string {
	var consulta = ""
	consulta = "select "
	consulta += "devolucionventaserviciodetalle.Id as id ,"
	consulta += "devolucionventaserviciodetalle.Codigo as codigo,"
	consulta += "devolucionventaserviciodetalle.Fila as fila,"
	consulta += "devolucionventaserviciodetalle.Cantidad as cantidad,"
	consulta += "devolucionventaserviciodetalle.Precio as precio,"
	consulta += "devolucionventaserviciodetalle.Descuento as descuento,"
	consulta += "devolucionventaserviciodetalle.Montodescuento as montodescuento,"
	consulta += "devolucionventaserviciodetalle.Sigratis as sigratis,"
	consulta += "devolucionventaserviciodetalle.Subtotal as subtotal,"
	consulta += "devolucionventaserviciodetalle.Subtotaldescuento  as subtotaldescuento,"
	consulta += "devolucionventaserviciodetalle.Pagina as pagina ,"
	consulta += "devolucionventaserviciodetalle.Fecha as fecha,"
	consulta += "devolucionventaserviciodetalle.Nombreservicio as Nombreservicio, "
	consulta += "devolucionventaserviciodetalle.Unidadservicio as Unidadservicio, "
	consulta += "devolucionventaserviciodetalle.Codigoservicio as Codigoservicio, "
	consulta += "devolucionventaserviciodetalle.Ivaservicio as Ivaservicio, "
	consulta += "iva.nombre as Ivanombreservicio "
	consulta += "from devolucionventaserviciodetalle "
	consulta += "inner join iva on iva.codigo=devolucionventaserviciodetalle.Ivaservicio "
	consulta += " where devolucionventaserviciodetalle.codigo=$1"
	log.Println(consulta)
	return consulta
}

// INICIA DEVOLUCIONVENTASERVICIO LISTA
func DevolucionventaservicioLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/devolucionventaservicio/devolucionventaservicioLista.html")
	log.Println("Error devolucionventaservicio 0")
	var consulta string
	var miperiodo = periodoSesion(r)
	consulta = "  SELECT vendedor.nombre as VendedorNombre,centro.nombre as CentroNombre,total,devolucionventaservicio.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	consulta += " FROM devolucionventaservicio "
	consulta += " inner join tercero on tercero.codigo=devolucionventaservicio.tercero "
	consulta += " inner join centro on centro.codigo=devolucionventaservicio.centro "
	consulta += " inner join vendedor on vendedor.codigo=devolucionventaservicio.vendedor "
	consulta += " where  extract(year from fecha)=" + miperiodo + " "
	consulta += " ORDER BY devolucionventaservicio.codigo ASC"

	db := dbConn()
	res := []devolucionventaservicioLista{}

	err := db.Select(&res, consulta)

	if err != nil {
		fmt.Println(err)
		return
	}

	varmap := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
	}
	log.Println("Error devolucionventaservicio888")
	tmp.Execute(w, varmap)
}

// INICIA DEVOLUCIONVENTASERVICIO NUEVO
func DevolucionventaservicioNuevo(w http.ResponseWriter, r *http.Request) {

	// TRAE COPIA DE EDITAR
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio ventaservicio editar" + Codigo)

	db := dbConn()
	v := devolucionventaservicio{}
	det := []devolucionventaserviciodetalleeditar{}
	if Codigo == "False" {

	} else {

		err := db.Get(&v, "SELECT * FROM devolucionventaservicio where codigo=$1", Codigo)
		if err != nil {
			log.Fatalln(err)
		}

		v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
		v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
		v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
		v.VendedorDetalle = TraerVendedorConsulta(v.Vendedor)
		v.CentroDetalle = TraerCentroConsulta(v.Centro)

		err2 := db.Select(&det, DevolucionventaservicioConsultaDetalle(), Codigo)
		if err2 != nil {
			fmt.Println(err2)
		}
	}

	parametros := map[string]interface{}{
		"miperiodo":            periodoSesion(r),
		"resolucionventa":      ListaResolucionventa(),
		"codigo":               Codigo,
		"vendedor":             PrimerVendedor(),
		"ventaservicio":        v,
		"detalle":              det,
		"hosting":              ruta,
		"ventaserviciotipoiva": TraerParametrosInventario().Ventaserviciotipoiva,
		"autoret2201":          TraerParametrosInventario().Ventaserviciocuentaporcentajeret2201,
	}
	//TERMINA TRAE COPIA DE EDITAR

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/devolucionventaservicio/devolucionventaservicioNuevo.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaVendedor.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaIva.html",
		"vista/devolucionventaservicio/devolucionventaservicioScript.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error devolucionventaservicio nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA INSERTAR COMPROBANTE DE DEVOLUCION VENTA
func InsertaDetalleComprobanteDevolucionventaservicioVenta(miFilaComprobante comprobantedetalle, miComprobante comprobante, miDevolucionventaservicio devolucionventaservicio) {
	db := dbConn()

	// traer tercero
	miTercero := tercero{}
	err3 := db.Get(&miTercero, "SELECT * FROM tercero where codigo=$1", miDevolucionventaservicio.Tercero)
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
		miDevolucionventaservicio.Tercero,
		miDevolucionventaservicio.Centro,
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

// INICIA DEVOLUCIONVENTASERVICIO INSERTAR AJAX
func DevolucionventaservicioAgregar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var tempDevolucionventaservicio devolucionventaservicio

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// carga informacion de la Devolucionventaservicio
	err = json.Unmarshal(b, &tempDevolucionventaservicio)
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
	//if tempDevolucionventaservicio.Accion == "Nuevo" {
	//	log.Println("Resolucion " + tempDevolucionventaservicio.Resolucion)
	//	Codigoactual = Numeroventa(tempDevolucionventaservicio.Resolucion)
	//	tempDevolucionventaservicio.Codigo = Codigoactual
	//} else {
	Codigoactual = tempDevolucionventaservicio.Codigo
	//}

	if tempDevolucionventaservicio.Accion == "Actualizar" {

		// borra detalle anterior
		delForm, err := db.Prepare("DELETE from devolucionventaserviciodetalle WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(tempDevolucionventaservicio.Codigo)

		// borra cabecera anterior
		delForm1, err := db.Prepare("DELETE from devolucionventaservicio WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm1.Exec(tempDevolucionventaservicio.Codigo)
	}

	// INSERTA DETALLE
	for i, x := range tempDevolucionventaservicio.Detalle {
		var a = i
		var q string
		q = "insert into devolucionventaserviciodetalle ("
		q += "Id,"
		q += "Codigo,"
		q += "Fila,"
		q += "Cantidad,"
		q += "Precio,"
		q += "Subtotal,"
		q += "Pagina,"
		q += "Nombreservicio,"
		q += "Unidadservicio,"
		q += "Codigoservicio,"
		q += "Ivaservicio,"
		q += "Descuento,"
		q += "Montodescuento,"
		q += "Sigratis,"
		q += "Subtotaldescuento,"
		q += "Tipo,"
		q += "Fecha"
		q += " ) values("
		q += parametros(17)
		q += ")"
		log.Println("Cadena SQL " + q)
		insForm, err := db.Prepare(q)
		if err != nil {
			panic(err.Error())
		}

		// TERMINA DEVOLUCIONVENTASERVICIO GRABAR INSERTAR
		_, err = insForm.Exec(
			x.Id,
			Codigoactual,
			x.Fila,
			Quitacoma(x.Cantidad),
			Quitacoma(x.Precio),
			Quitacoma(x.Subtotal),
			x.Pagina,
			Mayuscula(x.Nombreservicio),
			Titulo(x.Unidadservicio),
			x.Codigoservicio,
			Quitacoma(x.Ivaservicio),
			Quitacoma(x.Descuento),
			Quitacoma(x.Montodescuento),
			x.Sigratis,
			Quitacoma(x.Subtotaldescuento),
			x.Tipo,
			x.Fecha)
		if err != nil {
			panic(err)
		}

		log.Println("Insertar Producto \n", x.Nombreservicio, a)
	}

	// INICIA INSERTAR VENTAS
	log.Println("Got %s age %s club %s\n", tempDevolucionventaservicio.Codigo, tempDevolucionventaservicio.Tercero, tempDevolucionventaservicio.Total)
	var q string
	q = "insert into devolucionventaservicio ("
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
	q += "Ventaservicio,"
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
	log.Println("Hora", tempDevolucionventaservicio.Fecha.Format("02/01/2006"))
	currentTime := time.Now()

	_, err = insForm.Exec(
		tempDevolucionventaservicio.Resolucion,
		tempDevolucionventaservicio.Codigo,
		tempDevolucionventaservicio.Fecha.Format(layout),
		tempDevolucionventaservicio.Vence.Format(layout),
		currentTime.Format("3:4:5 PM"),
		Quitacoma(tempDevolucionventaservicio.Descuento),
		Quitacoma(tempDevolucionventaservicio.Subtotaldescuento19),
		Quitacoma(tempDevolucionventaservicio.Subtotaldescuento5),
		Quitacoma(tempDevolucionventaservicio.Subtotaldescuento0),
		Quitacoma(tempDevolucionventaservicio.Subtotal),
		Quitacoma(tempDevolucionventaservicio.Subtotal19),
		Quitacoma(tempDevolucionventaservicio.Subtotal5),
		Quitacoma(tempDevolucionventaservicio.Subtotal0),
		Quitacoma(tempDevolucionventaservicio.Subtotaliva19),
		Quitacoma(tempDevolucionventaservicio.Subtotaliva5),
		Quitacoma(tempDevolucionventaservicio.Subtotaliva0),
		Quitacoma(tempDevolucionventaservicio.Subtotalbase19),
		Quitacoma(tempDevolucionventaservicio.Subtotalbase5),
		Quitacoma(tempDevolucionventaservicio.Subtotalbase0),
		Quitacoma(tempDevolucionventaservicio.TotalIva),
		Quitacoma(tempDevolucionventaservicio.Ret2201),
		Quitacoma(tempDevolucionventaservicio.Total),
		Quitacoma(tempDevolucionventaservicio.Neto),
		tempDevolucionventaservicio.Items,
		tempDevolucionventaservicio.Formadepago,
		tempDevolucionventaservicio.Mediodepago,
		tempDevolucionventaservicio.Tercero,
		tempDevolucionventaservicio.Vendedor,
		tempDevolucionventaservicio.Ventaservicio,
		tempDevolucionventaservicio.Centro,
		tempDevolucionventaservicio.Tipo)

	if err != nil {
		panic(err)
	}

	// INSERTAR COMPROBANTE CONTABILIDAD
	var tempComprobante comprobante
	var tempComprobanteDetalle comprobantedetalle
	tempComprobante.Documento = "11"
	tempComprobante.Numero = tempDevolucionventaservicio.Codigo
	tempComprobante.Fecha = tempDevolucionventaservicio.Fecha
	tempComprobante.Fechaconsignacion = tempDevolucionventaservicio.Fecha
	tempComprobante.Debito = tempDevolucionventaservicio.Neto + ".00"
	tempComprobante.Credito = tempDevolucionventaservicio.Neto + ".00"
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
	if tempDevolucionventaservicio.Ret2201 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempDevolucionventaservicio.Ret2201)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Ventaserviciocontracuentaret2201
		tempComprobanteDetalle.Debito = tempDevolucionventaservicio.Ret2201
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteDevolucionventaservicioVenta(tempComprobanteDetalle, tempComprobante, tempDevolucionventaservicio)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA CREDITO IVA 19%
	if tempDevolucionventaservicio.Subtotaliva19 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempDevolucionventaservicio.Subtotaliva19)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Ventaserviciodevolucioniva19
		tempComprobanteDetalle.Debito = tempDevolucionventaservicio.Subtotaliva19
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteDevolucionventaservicioVenta(tempComprobanteDetalle, tempComprobante, tempDevolucionventaservicio)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA CREDITO IVA 5%
	if tempDevolucionventaservicio.Subtotaliva5 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempDevolucionventaservicio.Subtotaliva5)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Ventaserviciodevolucioniva5
		tempComprobanteDetalle.Debito = tempDevolucionventaservicio.Subtotaliva5
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteDevolucionventaservicioVenta(tempComprobanteDetalle, tempComprobante, tempDevolucionventaservicio)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA CREDITO VENTA IVA 19%
	if tempDevolucionventaservicio.Subtotalbase19 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempDevolucionventaservicio.Subtotalbase19)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Ventaserviciodevolucioncuenta19
		tempComprobanteDetalle.Debito = humanize.Comma(int64(Flotante(tempDevolucionventaservicio.Subtotalbase19)))
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteDevolucionventaservicioVenta(tempComprobanteDetalle, tempComprobante, tempDevolucionventaservicio)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA CREDITO VENTA IVA 5%
	if tempDevolucionventaservicio.Subtotalbase5 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempDevolucionventaservicio.Subtotalbase5)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Ventaserviciodevolucioncuenta5
		tempComprobanteDetalle.Debito = humanize.Comma(int64(Flotante(tempDevolucionventaservicio.Subtotalbase5)))
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteDevolucionventaservicioVenta(tempComprobanteDetalle, tempComprobante, tempDevolucionventaservicio)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA CREDITO VENTA IVA 0%
	if tempDevolucionventaservicio.Subtotalbase0 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempDevolucionventaservicio.Subtotalbase0)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Ventaserviciodevolucioncuenta0
		tempComprobanteDetalle.Debito = humanize.Comma(int64(Flotante(tempDevolucionventaservicio.Subtotalbase0)))
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteDevolucionventaservicioVenta(tempComprobanteDetalle, tempComprobante, tempDevolucionventaservicio)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA DEBITO RET. 2201
	if tempDevolucionventaservicio.Ret2201 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempDevolucionventaservicio.Ret2201)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Ventaserviciocuentaret2201
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = tempDevolucionventaservicio.Ret2201
		InsertaDetalleComprobanteDevolucionventaservicioVenta(tempComprobanteDetalle, tempComprobante, tempDevolucionventaservicio)
		log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito))
	}

	// INSERTAR CUENTA DEBITO CLIENTE
	if tempDevolucionventaservicio.Neto != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempDevolucionventaservicio.Neto)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Ventaserviciocuentacliente
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = tempDevolucionventaservicio.Neto
		InsertaDetalleComprobanteDevolucionventaservicioVenta(tempComprobanteDetalle, tempComprobante, tempDevolucionventaservicio)
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

// INICIA DEVOLUCIONVENTASERVICIO EXISTE
func DevolucionventaservicioExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM devolucionventaservicio  WHERE codigo=$1", Codigo)
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

// INICIA DEVOLUCIONVENTASERVICIO EDITAR
func DevolucionventaservicioEditar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio devolucionventaservicio editar" + Codigo)
	db := dbConn()

	// traer devolucionventaservicio
	v := devolucionventaservicio{}
	err := db.Get(&v, "SELECT * FROM devolucionventaservicio where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []devolucionventaserviciodetalleeditar{}

	err2 := db.Select(&det, DevolucionventaservicioConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
	v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
	v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
	v.VendedorDetalle = TraerVendedorConsulta(v.Vendedor)
	v.CentroDetalle = TraerCentroConsulta(v.Centro)

	parametros := map[string]interface{}{
		"miperiodo":               periodoSesion(r),
		"resolucionventa":         ListaResolucionventa(),
		"devolucionventaservicio": v,
		"detalle":                 det,
		"hosting":                 ruta,
		"ventaserviciotipoiva":    TraerParametrosInventario().Ventaserviciotipoiva,
		"autoret2201":             TraerParametrosInventario().Ventaserviciocuentaporcentajeret2201,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/devolucionventaservicio/devolucionventaservicioEditar.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaVendedor.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaIva.html",
		"vista/devolucionventaservicio/devolucionventaservicioScript.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error devolucionventaservicio nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA DEVOLUCIONVENTASERVICIO BORRAR
func DevolucionventaservicioBorrar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio devolucionventaservicio editar" + Codigo)

	db := dbConn()

	// traer devolucionventaservicio
	v := devolucionventaservicio{}
	err := db.Get(&v, "SELECT * FROM devolucionventaservicio where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []devolucionventaserviciodetalleeditar{}
	err2 := db.Select(&det, DevolucionventaservicioConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
	v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
	v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
	v.VendedorDetalle = TraerVendedorConsulta(v.Vendedor)
	v.CentroDetalle = TraerCentroConsulta(v.Centro)

	parametros := map[string]interface{}{
		"resolucionventa":         ListaResolucionventa(),
		"devolucionventaservicio": v,
		"detalle":                 det,
		"hosting":                 ruta,
		"ventaserviciotipoiva":    TraerParametrosInventario().Ventaserviciotipoiva,
		"autoret2201":             TraerParametrosInventario().Ventacuentaporcentajeret2201,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/devolucionventaservicio/devolucionventaservicioBorrar.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaIva.html",
		"vista/devolucionventaservicio/devolucionventaservicioScript.html")

	log.Println("Error devolucionventaservicio nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA DEVOLUCIONVENTASERVICIO ELIMINAR
func DevolucionventaservicioEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	codigo := mux.Vars(r)["codigo"]

	// borrar Devolucionventaservicio
	delForm, err := db.Prepare("DELETE from devolucionventaservicio WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(codigo)

	// borar detalle
	delForm1, err := db.Prepare("DELETE from devolucionventaserviciodetalle WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm1.Exec(codigo)

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
	http.Redirect(w, r, "/DevolucionventaservicioLista", 301)
}

// TRAER VENTA SERVICIO EN LA DEVOLUCION
func DatosVentaservicio(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	Tercero := mux.Vars(r)["tercero"]
	log.Println("inicio venta editar" + Codigo)
	db := dbConn()
	var res []ventaservicio

	// traer VENTA
	v := ventaservicio{}
	err := db.Get(&v, "SELECT * FROM ventaservicio where codigo=$1 and Tercero = $2", Codigo, Tercero)
	var valida bool
	valida = true

	switch err {
	case nil:
		log.Printf("tercero found: %+v\n", v)
	case sql.ErrNoRows:
		log.Println("tercero NOT found, no error")
		valida = false
	default:
		log.Printf("tercero error: %s\n", err)
	}
	det := []ventaserviciodetalleeditar{}
	t := tercero{}

	// trae datos si existe factura
	if valida == true {
		err2 := db.Select(&det, VentaservicioConsultaDetalle(), Codigo)
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

// INICIA PDF
func DevolucionventaservicioPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]

	// traer Devolucionventaservicio
	miDevolucionventaservicio := devolucionventaservicio{}
	err := db.Get(&miDevolucionventaservicio, "SELECT * FROM devolucionventaservicio where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}

	// traer detalle
	miDetalle := []devolucionventaserviciodetalleeditar{}
	err2 := db.Select(&miDetalle, DevolucionventaservicioConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	// traer tercero
	miTercero := tercero{}
	err3 := db.Get(&miTercero, "SELECT * FROM tercero where codigo=$1", miDevolucionventaservicio.Tercero)
	if err3 != nil {
		log.Fatalln(err3)
	}

	// traer Vendedor
	miVendedor := vendedor{}
	err4 := db.Get(&miVendedor, "SELECT * FROM vendedor where codigo=$1", miDevolucionventaservicio.Vendedor)
	if err4 != nil {
		log.Fatalln(err4)
	}

	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	var re Resolucionventa = TraerResolucionventa(miDevolucionventaservicio.Resolucion)
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
		pdf.CellFormat(190, 10, "DEVOLUCION VENTA ", "0", 0, "C", false, 0, "")
		pdf.Ln(5)
		pdf.SetX(80)
		pdf.CellFormat(190, 10, " DE SERVICIO No. "+re.Prefijo+" "+Codigo, "0", 0, "C",
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
	DevolucionventaservicioCabecera(pdf, miTercero, miDevolucionventaservicio, miVendedor)

	var filas = len(miDetalle)
	// UNA PAGINA
	if filas <= 32 {
		for i, miFila := range miDetalle {
			var a = i + 1
			DevolucionventaservicioFilaDetalle(pdf, miFila, a)
		}
		DevolucionventaservicioPieDePagina(pdf, miTercero, miDevolucionventaservicio)
	} else {
		// mas de 32 y menos de 73
		if filas > 32 && filas <= 73 {
			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 41 {
					DevolucionventaservicioFilaDetalle(pdf, miFila, a)
				}
			}
			DevolucionventaservicioLinea(pdf)
			// segunda pagina
			pdf.AddPage()
			DevolucionventaservicioCabecera(pdf, miTercero, miDevolucionventaservicio, miVendedor)
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 41 {
					DevolucionventaservicioFilaDetalle(pdf, miFila, a)
				}
			}

			DevolucionventaservicioPieDePagina(pdf, miTercero, miDevolucionventaservicio)
		} else {
			// mas de 80

			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 41 {
					DevolucionventaservicioFilaDetalle(pdf, miFila, a)
				}
			}
			DevolucionventaservicioLinea(pdf)
			pdf.AddPage()
			DevolucionventaservicioCabecera(pdf, miTercero, miDevolucionventaservicio, miVendedor)
			// segunda pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 41 && a <= 82 {
					DevolucionventaservicioFilaDetalle(pdf, miFila, a)
				}
			}
			DevolucionventaservicioLinea(pdf)
			pdf.AddPage()
			DevolucionventaservicioCabecera(pdf, miTercero, miDevolucionventaservicio, miVendedor)
			// tercera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 82 {
					DevolucionventaservicioFilaDetalle(pdf, miFila, a)
				}
			}
			DevolucionventaservicioPieDePagina(pdf, miTercero, miDevolucionventaservicio)
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

func DevolucionventaservicioCabecera(pdf *gofpdf.Fpdf, miTercero tercero, miDevolucionventaservicio devolucionventaservicio, miVendedor vendedor) {
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
	pdf.CellFormat(40, 4, Titulo(TraerFormadepago(miDevolucionventaservicio.Formadepago)), "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Fecha de Expedicion", "", 0,
		"L", false, 0, "")
	pdf.SetX(150)
	pdf.CellFormat(40, 4, miDevolucionventaservicio.Fecha.Format("02/01/2006")+" "+Titulo(miDevolucionventaservicio.Hora), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)

	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Medio de Pago", "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, TraerMediodepago(miDevolucionventaservicio.Mediodepago), "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Fecha de Vencimiento", "", 0,
		"L", false, 0, "")
	pdf.SetX(150)
	pdf.CellFormat(40, 4, miDevolucionventaservicio.Vence.Format("02/01/2006"), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Factura No.", "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, miDevolucionventaservicio.Ventaservicio, "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Vendedor", "", 0,
		"L", false, 0, "")
	pdf.SetX(130)
	pdf.CellFormat(40, 4, ene(miVendedor.Nombre), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)

	// CABECERA PRODUCTO
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
	pdf.SetX(112)
	pdf.CellFormat(40, 6, "UNIDAD", "0", 0,
		"L", false, 0, "")
	pdf.SetX(129)
	pdf.CellFormat(40, 6, "IVA", "0", 0,
		"L", false, 0, "")
	pdf.SetX(140)
	pdf.CellFormat(40, 6, "CANTIDAD", "0", 0,
		"L", false, 0, "")
	pdf.SetX(163)
	pdf.CellFormat(40, 6, "P. UNITARIO", "0", 0,
		"L", false, 0, "")
	pdf.SetX(190)
	pdf.CellFormat(40, 6, "TOTAL", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}
func DevolucionventaservicioFilaDetalle(pdf *gofpdf.Fpdf, miFila devolucionventaserviciodetalleeditar, a int) {
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
	pdf.MultiCell(72, 4, ene(Mayuscula(miFila.Nombreservicio)), "", "L", false)
	var yfinal float64
	yfinal = pdf.GetY()
	pdf.SetY(y)
	pdf.SetX(87)
	pdf.CellFormat(40, 4, ene(Subcadena(miFila.Unidadservicio, 0, 6)), "", 0,
		"R", false, 0, "")
	pdf.SetX(97)
	pdf.CellFormat(40, 4, miFila.Ivanombreservicio+"%", "", 0,
		"R", false, 0, "")
	pdf.SetX(120)
	pdf.CellFormat(40, 4, Coma(miFila.Cantidad), "", 0,
		"R", false, 0, "")
	pdf.SetX(141)
	pdf.CellFormat(40, 4, Coma(miFila.Precio), "", 0,
		"R", false, 0, "")
	pdf.SetX(165)
	pdf.CellFormat(40, 4, Coma(miFila.Subtotal), "", 0,
		"R", false, 0, "")
	pdf.Ln(yfinal - yinicial + 3)
}

func DevolucionventaservicioPieDePagina(pdf *gofpdf.Fpdf, miTercero tercero, miDevolucionventaservicio devolucionventaservicio) {

	Totalletras, err := IntLetra(Cadenaentero(miDevolucionventaservicio.Total))
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
	if miDevolucionventaservicio.Neto != "0" {
		pdf.SetFont("Arial", "", 9)
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "TOTAL", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miDevolucionventaservicio.Neto), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucionventaservicio.Subtotaliva5 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "I. V. A.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(170)
		pdf.CellFormat(190, 4, "5%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miDevolucionventaservicio.Subtotaliva5), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucionventaservicio.Subtotaliva19 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "I. V. A.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(170)
		pdf.CellFormat(190, 4, "19%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miDevolucionventaservicio.Subtotaliva19), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucionventaservicio.Subtotalbase5 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "GRAVABLE", "0", 0, "L",
			false, 0, "")
		pdf.SetX(170)
		pdf.CellFormat(190, 4, "5%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miDevolucionventaservicio.Subtotalbase5), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucionventaservicio.Subtotalbase19 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "GRAVABLE", "0", 0, "L",
			false, 0, "")
		pdf.SetX(170)
		pdf.CellFormat(190, 4, "19%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miDevolucionventaservicio.Subtotalbase19), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucionventaservicio.Subtotalbase0 != "0" {

		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "NO GRAVABLE", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miDevolucionventaservicio.Subtotalbase0), "0", 0, "R",
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

func DevolucionventaservicioLinea(pdf *gofpdf.Fpdf) {
	pdf.SetY(259)
	pdf.Line(20, 259, 204, 259)
}

// INICIA DEVOLUCION VENTA SERVICIO TODOS PDF
func DevolucionventaservicioTodosCabecera(pdf *gofpdf.Fpdf) {
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

func DevolucionventaservicioTodosDetalle(pdf *gofpdf.Fpdf, miFila devolucionventaservicioLista, a int) {
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

func DevolucionventaservicioTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string
	consulta = "  SELECT vendedor.nombre as VendedorNombre," +
		"centro.nombre as CentroNombre,devolucionventaservicio.total," +
		"devolucionventaservicio.codigo,fecha,tercero, tercero.nombre as TerceroNombre "
	consulta += " FROM devolucionventaservicio "
	consulta += " inner join tercero on tercero.codigo=devolucionventaservicio.tercero "
	consulta += " inner join centro on centro.codigo=devolucionventaservicio.centro "
	consulta += " inner join vendedor on vendedor.codigo=devolucionventaservicio.vendedor "
	consulta += " ORDER BY cast(devolucionventaservicio.codigo as integer) ASC"

	t := []devolucionventaservicioLista{}
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
		pdf.CellFormat(190, 10, "DATOS DEVOLUCION VENTA SERVICIOS", "0", 0,
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

	VentaservicioTodosCabecera(pdf)
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
		DevolucionventaservicioTodosDetalle(pdf, miFila, a)
	}

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

// TERMINA VENTA TODOS PDF

// VENTA EXCEL
func DevolucionventaservicioExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string
	consulta = "  SELECT vendedor.nombre as VendedorNombre," +
		"centro.nombre as CentroNombre,devolucionventaservicio.total," +
		"devolucionventaservicio.codigo,fecha,tercero," +
		"tercero.nombre as TerceroNombre "
	consulta += " FROM devolucionventaservicio "
	consulta += " inner join tercero on tercero." +
		"codigo=devolucionventaservicio.tercero "
	consulta += " inner join centro on centro.codigo=devolucionventaservicio." +
		"centro "
	consulta += " inner join vendedor on vendedor." +
		"codigo=devolucionventaservicio." +
		"vendedor "
	consulta += " ORDER BY cast(devolucionventaservicio.codigo as integer) ASC"
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
	f.SetCellValue("Sheet1", "A9", "LISTADO DEVOLUCION VENTA DE SERVICIOS")
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

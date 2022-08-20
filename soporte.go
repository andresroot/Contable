package main

// INICIA SOPORTE IMPORTAR PAQUETES
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

// INICIA SOPORTE ESTRUCTURA JSON
type soporteJson struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Nombre string `json:"nombre"`
}

// INICIA SOPORTE ESTRUCTURA
type soporteLista struct {
	Codigo            string
	Fecha             time.Time
	Neto              string
	Tercero           string
	TerceroNombre     string
	CentroNombre      string
	AlmacenistaNombre string
}

// INICIA SOPORTE ESTRUCTURA
type soporte struct {
	Resolucionsoporte         string
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
	Detalle                   []soportedetalle       `json:"Detalle"`
	DetalleEditar             []soportedetalleeditar `json:"DetalleEditar"`
	TerceroDetalle            tercero
	Pedidosoporte             string
	Tipo                      string
	Centro                    string
	FormadepagoDetalle        formadepago
	MediodepagoDetalle        mediodepago
	AlmacenistaDetalle        almacenista
	CentroDetalle             centro
}

// INICIA SOPORTEDETALLE ESTRUCTURA
type soportedetalle struct {
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

// INICIA SOPORTE DETALLE EDITARr
type soportedetalleeditar struct {
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

// INICIA SOPORTE CONSULTA DETALLE
func SoporteConsultaDetalle() string {
	var consulta = ""
	consulta = "select "
	consulta += "soportedetalle.Id as id ,"
	consulta += "soportedetalle.Codigo as codigo,"
	consulta += "soportedetalle.Fila as fila,"
	consulta += "soportedetalle.Cantidad as cantidad,"
	consulta += "soportedetalle.Precio as precio,"
	consulta += "soportedetalle.Descuento as descuento,"
	consulta += "soportedetalle.Montodescuento as montodescuento,"
	consulta += "soportedetalle.Sigratis as sigratis,"
	consulta += "soportedetalle.Subtotal as subtotal,"
	consulta += "soportedetalle.Subtotaldescuento  as subtotaldescuento,"
	consulta += "soportedetalle.Pagina as pagina ,"
	consulta += "soportedetalle.Bodega as bodega,"
	consulta += "bodega.Nombre as bodeganombre,"
	consulta += "soportedetalle.Producto as producto,"
	consulta += "soportedetalle.Fecha as fecha,"
	consulta += "producto.nombre as ProductoNombre, "
	consulta += "producto.iva as ProductoIva, "
	consulta += "producto.unidad as ProductoUnidad "
	consulta += "from soportedetalle "
	consulta += "inner join producto on producto.codigo=soportedetalle.producto "
	consulta += "inner join bodega on bodega.codigo=soportedetalle.bodega "
	consulta += " where soportedetalle.codigo=$1"
	log.Println(consulta)
	return consulta
}

// INICIA SOPORTE LISTA
func SoporteLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/soporte/soporteLista.html")
	log.Println("Error soporte 0")
	var consulta string
	var miperiodo = periodoSesion(r)
	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre,centro.nombre as CentroNombre,soporte.neto,soporte.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	consulta += " FROM soporte "
	consulta += " inner join tercero on tercero.codigo=soporte.tercero "
	consulta += " inner join centro on centro.codigo=soporte.centro "
	consulta += " inner join almacenista on almacenista.codigo=soporte.almacenista "
	consulta += " where  extract(year from fecha)=" + miperiodo + " "
	consulta += " ORDER BY cast(soporte.codigo as integer) ASC"

	db := dbConn()
	res := []soporteLista{}

	err := db.Select(&res, consulta)

	if err != nil {
		fmt.Println(err)
	}

	varmap := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
	}
	log.Println("Error soporte888")
	tmp.Execute(w, varmap)
}

// INICIA SOPORTE NUEVO
func SoporteNuevo(w http.ResponseWriter, r *http.Request) {

	// TRAE COPIA DE EDITAR
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio soporte editar" + Codigo)

	db := dbConn()
	v := soporte{}
	det := []soportedetalleeditar{}
	if Codigo == "False" {

	} else {

		err := db.Get(&v, "SELECT * FROM soporte where codigo=$1", Codigo)
		if err != nil {
			log.Fatalln(err)
		}

		v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
		v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
		v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
		v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
		v.CentroDetalle = TraerCentroConsulta(v.Centro)
		err2 := db.Select(&det, SoporteConsultaDetalle(), Codigo)
		if err2 != nil {
			fmt.Println(err2)
		}
	}

	parametros := map[string]interface{}{
		"miperiodo":         periodoSesion(r),
		"resolucionsoporte": ListaResolucionsoporte(),
		"codigo":            Codigo,
		"soporte":           v,
		"almacenista":       PrimerAlmacenista(),
		"detalle":           det,
		"hosting":           ruta,
		"retfte":            TraerParametrosInventario().Soportecuentaporcentajeretfte,
	}
	//TERMINA TRAE COPIA DE EDITAR

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/soporte/soporteNuevo.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaBodega.html",
		"vista/autocompleta/autocompletaTercero.html",
		"vista/soporte/soporteScript.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error soporte nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA INSERTAR COMPROBANTE DE SOPORTE
func InsertaDetalleComprobanteSoporte(miFilaComprobante comprobantedetalle, miComprobante comprobante, miSoporte soporte) {
	db := dbConn()

	// traer tercero
	miTercero := tercero{}
	err3 := db.Get(&miTercero, "SELECT * FROM tercero where codigo=$1", miSoporte.Tercero)
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
		miSoporte.Tercero,
		miSoporte.Centro,
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

// INICIA SOPORTE INSERTAR AJAX
func SoporteAgregar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var tempSoporte soporte

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// carga informacion de la SOPORTE
	err = json.Unmarshal(b, &tempSoporte)
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
	if tempSoporte.Accion == "Nuevo" {
		log.Println("Resolucion " + tempSoporte.Resolucionsoporte)
		Codigoactual = Numerosoporte(tempSoporte.Resolucionsoporte)
		tempSoporte.Codigo = Codigoactual
	} else {
		Codigoactual = tempSoporte.Codigo
	}

	if tempSoporte.Accion == "Actualizar" {

		// borra detalle anterior
		delForm, err := db.Prepare("DELETE from soportedetalle WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(tempSoporte.Codigo)

		// borra detalle inventario
		Borrarinventario(Codigoactual, "Soporte")

		// borra cabecera anterior
		delForm1, err := db.Prepare("DELETE from soporte WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm1.Exec(tempSoporte.Codigo)
	}

	// INSERTA DETALLE
	for i, x := range tempSoporte.Detalle {
		var a = i
		var q string
		q = "insert into soportedetalle ("
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

		// TERMINA SOPORTE GRABAR INSERTAR
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
	for _, x := range tempSoporte.Detalle {
		miInventario = append(miInventario, inventario{
			x.Fecha, x.Tipo, Codigoactual, x.Bodega,
			x.Producto,
			Quitacoma(x.Cantidad),
			Quitacoma(x.Precio),
			operacionSoporte})
	}
	// insertar inventario

	InsertaInventario(miInventario)

	// INICIA INSERTAR SOPORTES
	log.Println("Got %s age %s club %s\n", tempSoporte.Codigo, tempSoporte.Tercero, tempSoporte.Subtotal)
	var q string
	q = "insert into soporte ("
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
	q += "Pedidosoporte,"
	q += "Centro,"
	q += "Tipo"
	q += " ) values("
	q += parametros(34)
	q += " ) "

	log.Println("Cadena SQL " + q)
	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	layout := "2006-01-02"
	log.Println("Hora", tempSoporte.Fecha.Format("02/01/2006"))
	currentTime := time.Now()

	_, err = insForm.Exec(
		tempSoporte.Resolucionsoporte,
		tempSoporte.Codigo,
		tempSoporte.Fecha.Format(layout),
		tempSoporte.Vence.Format(layout),
		currentTime.Format("3:4:5 PM"),
		Quitacoma(tempSoporte.Descuento),
		Quitacoma(tempSoporte.Subtotaldescuento19),
		Quitacoma(tempSoporte.Subtotaldescuento5),
		Quitacoma(tempSoporte.Subtotaldescuento0),
		Quitacoma(tempSoporte.Subtotal),
		Quitacoma(tempSoporte.Subtotal19),
		Quitacoma(tempSoporte.Subtotal5),
		Quitacoma(tempSoporte.Subtotal0),
		Quitacoma(tempSoporte.Subtotaliva19),
		Quitacoma(tempSoporte.Subtotaliva5),
		Quitacoma(tempSoporte.Subtotaliva0),
		Quitacoma(tempSoporte.Subtotalbase19),
		Quitacoma(tempSoporte.Subtotalbase5),
		Quitacoma(tempSoporte.Subtotalbase0),
		Quitacoma(tempSoporte.TotalIva),
		Quitacoma(tempSoporte.Total),
		Quitacoma(tempSoporte.Neto),
		tempSoporte.Items,
		Quitacoma(tempSoporte.PorcentajeRetencionFuente),
		Quitacoma(tempSoporte.TotalRetencionFuente),
		Quitacoma(tempSoporte.PorcentajeRetencionIca),
		Quitacoma(tempSoporte.TotalRetencionIca),
		tempSoporte.Formadepago,
		tempSoporte.Mediodepago,
		tempSoporte.Tercero,
		tempSoporte.Almacenista,
		tempSoporte.Pedidosoporte,
		tempSoporte.Centro,
		tempSoporte.Tipo)

	if err != nil {
		panic(err)
	}

	// INSERTAR COMPROBANTE CONTABILIDAD
	var tempComprobante comprobante
	var tempComprobanteDetalle comprobantedetalle
	tempComprobante.Documento = "10"
	tempComprobante.Numero = tempSoporte.Codigo
	tempComprobante.Fecha = tempSoporte.Fecha
	tempComprobante.Fechaconsignacion = tempSoporte.Fecha
	tempComprobante.Debito = tempSoporte.Neto
	tempComprobante.Credito = tempSoporte.Neto
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

	// INSERTAR CUENTA DEBITO SOPORTE 19%
	if tempSoporte.Subtotal19 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempSoporte.Subtotal19) - Flotante(tempSoporte.Subtotaldescuento19)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Soportecuenta19
		tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(tempSoporte.Subtotal19) - Flotante(tempSoporte.Subtotaldescuento19))
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteSoporte(tempComprobanteDetalle, tempComprobante, tempSoporte)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA DEBITO SOPORTE 5%
	if tempSoporte.Subtotal5 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempSoporte.Subtotal5) - Flotante(tempSoporte.Subtotaldescuento5)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Soportecuenta5
		tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(tempSoporte.Subtotal5) - Flotante(tempSoporte.Subtotaldescuento5))
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteSoporte(tempComprobanteDetalle, tempComprobante, tempSoporte)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA DEBITO SOPORTE 0%
	if tempSoporte.Subtotal0 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempSoporte.Subtotal0) - Flotante(tempSoporte.Subtotaldescuento0)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Soportecuenta0
		tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(tempSoporte.Subtotal0) - Flotante(tempSoporte.Subtotaldescuento0))
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteSoporte(tempComprobanteDetalle, tempComprobante, tempSoporte)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA CREDITO DESCUENTO
	//if (tempSoporte.Descuento!="0")	{
	//	fila=fila+1
	//	tempComprobanteDetalle.Fila=strconv.Itoa(fila)
	//	totalCredito+=Flotante(tempSoporte.Descuento)
	//	tempComprobanteDetalle.Cuenta  = parametrosinventario.Soportecuentadescuento
	//	tempComprobanteDetalle.Debito = ""
	//	tempComprobanteDetalle.Credito = tempSoporte.Descuento
	//	InsertaDetalleComprobanteSoporte(tempComprobanteDetalle,tempComprobante,tempSoporte)
	//	log.Println("credito linea" + fmt.Sprintf("%.2f",totalCredito))
	//}

	// INSERTAR CUENTA CREDITO RET. FTE.
	if tempSoporte.TotalRetencionFuente != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempSoporte.TotalRetencionFuente)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Soportecuentaretfte
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = tempSoporte.TotalRetencionFuente
		InsertaDetalleComprobanteSoporte(tempComprobanteDetalle, tempComprobante, tempSoporte)
		log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito))
	}

	// INSERTAR CUENTA CREDITO RET. ICA.
	if tempSoporte.TotalRetencionIca != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempSoporte.TotalRetencionIca)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Soportecuentaretica
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = tempSoporte.TotalRetencionIca
		InsertaDetalleComprobanteSoporte(tempComprobanteDetalle, tempComprobante, tempSoporte)
		log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito))
	}

	// INSERTAR CUENTA CREDITO PROVEEDOR
	if tempSoporte.Neto != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempSoporte.Neto)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Soportecuentaproveedor
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = tempSoporte.Neto
		InsertaDetalleComprobanteSoporte(tempComprobanteDetalle, tempComprobante, tempSoporte)
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

// INICIA SOPORTE EXISTE
func SoporteExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM soporte  WHERE codigo=$1", Codigo)
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

// INICIA SOPORTE EDITAR
func SoporteEditar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio soporte editar" + Codigo)
	db := dbConn()

	// traer SOPORTE
	v := soporte{}
	err := db.Get(&v, "SELECT * FROM soporte where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []soportedetalleeditar{}

	err2 := db.Select(&det, SoporteConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
	v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
	v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
	v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
	v.CentroDetalle = TraerCentroConsulta(v.Centro)

	parametros := map[string]interface{}{
		"miperiodo":         periodoSesion(r),
		"resolucionsoporte": ListaResolucionsoporte(),
		"soporte":           v,
		"detalle":           det,
		"hosting":           ruta,
		"retfte":            TraerParametrosInventario().Soportecuentaporcentajeretfte,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/soporte/soporteEditar.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaBodega.html",
		"vista/autocompleta/autocompletaTercero.html",
		"vista/soporte/soporteScript.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error soporte nuevo 444")
	miTemplate.Execute(w, parametros)
}

// INICIA SOPORTE BORRAR
func SoporteBorrar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio soporte editar" + Codigo)

	db := dbConn()

	// traer SOPORTE
	v := soporte{}
	err := db.Get(&v, "SELECT * FROM soporte where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []soportedetalleeditar{}
	err2 := db.Select(&det, SoporteConsultaDetalle(), Codigo)
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
		"soporte":           v,
		"detalle":           det,
		"hosting":           ruta,
		"retfte":            TraerParametrosInventario().Soportecuentaporcentajeretfte,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/soporte/soporteBorrar.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaBodega.html",
		"vista/soporte/soporteScript.html")

	log.Println("Error soporte nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA SOPORTE ELIMINAR
func SoporteEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	codigo := mux.Vars(r)["codigo"]

	// borrar SOPORTE
	delForm, err := db.Prepare("DELETE from soporte WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(codigo)

	// borar detalle
	delForm1, err := db.Prepare("DELETE from soportedetalle WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm1.Exec(codigo)

	// borar detalle invenario
	Borrarinventario(codigo, "Soporte")

	// borra detalle anterior
	delForm, err = db.Prepare("DELETE from comprobantedetalle WHERE documento=$1 and numero=$2")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec("8", codigo)

	// borra cabecera anterior

	delForm1, err = db.Prepare("DELETE from comprobante WHERE documento=$1 and numero=$2 ")
	if err != nil {
		panic(err.Error())
	}
	delForm1.Exec("8", codigo)

	log.Println("Registro Eliminado" + codigo)
	http.Redirect(w, r, "/SoporteLista", 301)
}

// TRAER PEDIDO
func Datospedidosoporte(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	Tercero := mux.Vars(r)["tercero"]
	log.Println("inicio pedido editar" + Codigo)
	db := dbConn()
	var res []pedidosoporte

	// traer PEDIDO
	v := pedidosoporte{}
	err := db.Get(&v, "SELECT * FROM pedidosoporte where codigo=$1 and Tercero=$2", Codigo, Tercero)
	var valida bool
	valida = true

	switch err {
	case nil:
		log.Printf("pedidosoporte existe: %+v\n", v)
	case sql.ErrNoRows:
		log.Println("pedidosoporte NO Existe")
		valida = false
	default:
		log.Printf("pedidosoporte error: %s\n", err)
	}
	det := []pedidosoportedetalleeditar{}
	t := tercero{}

	// trae datos si existe pedido
	if valida == true {
		err2 := db.Select(&det, PedidosoporteConsultaDetalle(), Codigo)
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

// INICIA SOPORTE PDF
func SoportePdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]

	// traer SOPORTE
	miSoporte := soporte{}
	err := db.Get(&miSoporte, "SELECT * FROM soporte where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}

	// traer detalle
	miDetalle := []soportedetalleeditar{}
	err2 := db.Select(&miDetalle, SoporteConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	// traer tercero
	miTercero := tercero{}
	err3 := db.Get(&miTercero, "SELECT * FROM tercero where codigo=$1", miSoporte.Tercero)
	if err3 != nil {
		log.Fatalln(err3)
	}

	// traer almacenista
	miAlmacenista := almacenista{}
	err4 := db.Get(&miAlmacenista, "SELECT * FROM almacenista where codigo=$1", miSoporte.Almacenista)
	if err4 != nil {
		log.Fatalln(err4)
	}

	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	var re Resolucionsoporte = TraerResolucionsoporte(miSoporte.Resolucionsoporte)
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
	SoporteCabecera(pdf, miTercero, miSoporte, miAlmacenista)

	var filas = len(miDetalle)
	// menos de 32
	if filas <= 32 {
		for i, miFila := range miDetalle {
			var a = i + 1
			SoporteFilaDetalle(pdf, miFila, a)
		}
		SoportePieDePagina(pdf, miTercero, miSoporte)
	} else {
		// mas de 32 y menos de 73
		if filas > 32 && filas <= 73 {
			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 41 {
					SoporteFilaDetalle(pdf, miFila, a)
				}
			}
			SoporteLinea(pdf)
			// segunda pagina
			pdf.AddPage()
			SoporteCabecera(pdf, miTercero, miSoporte, miAlmacenista)
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 41 {
					SoporteFilaDetalle(pdf, miFila, a)
				}
			}

			SoportePieDePagina(pdf, miTercero, miSoporte)
		} else {
			// mas de 80

			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 41 {
					SoporteFilaDetalle(pdf, miFila, a)
				}
			}
			PedidosoporteLinea(pdf)
			pdf.AddPage()
			SoporteCabecera(pdf, miTercero, miSoporte, miAlmacenista)
			// segunda pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 41 && a <= 82 {
					SoporteFilaDetalle(pdf, miFila, a)
				}
			}
			PedidosoporteLinea(pdf)
			pdf.AddPage()
			SoporteCabecera(pdf, miTercero, miSoporte, miAlmacenista)
			// tercera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 82 {
					SoporteFilaDetalle(pdf, miFila, a)
				}
			}
			SoportePieDePagina(pdf, miTercero, miSoporte)
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

func SoporteCabecera(pdf *gofpdf.Fpdf, miTercero tercero, miSoporte soporte, miAlmacenista almacenista) {
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
	pdf.CellFormat(40, 4, Titulo(TraerFormadepago(miSoporte.Formadepago)), "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Fecha de Expedicion", "", 0,
		"L", false, 0, "")
	pdf.SetX(150)
	pdf.CellFormat(40, 4, miSoporte.Fecha.Format("02/01/2006")+" "+Titulo(miSoporte.Hora), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)

	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Medio de Pago", "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, TraerMediodepago(miSoporte.Mediodepago), "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Fecha de Vencimiento", "", 0,
		"L", false, 0, "")
	pdf.SetX(150)
	pdf.CellFormat(40, 4, miSoporte.Vence.Format("02/01/2006"), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Pedido No.", "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, miSoporte.Pedidosoporte, "", 0,
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
func SoporteFilaDetalle(pdf *gofpdf.Fpdf, miFila soportedetalleeditar, a int) {
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
	pdf.SetX(91)
	pdf.CellFormat(40, 4, miFila.ProductoIva+"%", "", 0,
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

func SoportePieDePagina(pdf *gofpdf.Fpdf, miTercero tercero, miSoporte soporte) {

	Totalletras, err := IntLetra(Cadenaentero(miSoporte.Neto))
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
	separador = 250
	altoseparador = -4

	// INICIA DATOS FACTURA
	if miSoporte.Neto != "0" {
		pdf.SetFont("Arial", "", 9)
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "TOTAL", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miSoporte.Neto), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miSoporte.TotalRetencionIca != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "RET. ICA.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, miSoporte.PorcentajeRetencionIca+"%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miSoporte.TotalRetencionIca), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miSoporte.TotalRetencionFuente != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "RET. FTE.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, miSoporte.PorcentajeRetencionFuente+"%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miSoporte.TotalRetencionFuente), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miSoporte.Subtotalbase5 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "GRAVABLE", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, "5%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miSoporte.Subtotalbase5), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miSoporte.Subtotalbase19 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "GRAVABLE", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, "19%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miSoporte.Subtotalbase19), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miSoporte.Subtotalbase0 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "NO GRAVABLE", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miSoporte.Subtotalbase0), "0", 0, "R",
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

func SoporteLinea(pdf *gofpdf.Fpdf) {
	pdf.SetY(259)
	pdf.Line(20, 259, 204, 259)
}

// INICIA SOPORTE TODOS PDF
func SoporteTodosCabecera(pdf *gofpdf.Fpdf) {
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

func SoporteTodosDetalle(pdf *gofpdf.Fpdf, miFila soporteLista, a int) {
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

func SoporteTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string
	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre," +
		"centro.nombre as CentroNombre,soporte.neto," +
		"soporte.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	consulta += " FROM soporte "
	consulta += " inner join tercero on tercero.codigo=soporte.tercero "
	consulta += " inner join centro on centro.codigo=soporte.centro "
	consulta += " inner join almacenista on almacenista.codigo=soporte.almacenista "
	consulta += " ORDER BY cast(soporte.codigo as integer) ASC"

	t := []soporteLista{}
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
		pdf.CellFormat(190, 10, "DATOS DOCUMENTO SOPORTE", "0", 0,
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
		SoporteTodosDetalle(pdf, miFila, a)
	}

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

// TERMINA PEDIDO TODOS PDF

// SOPORTE EXCEL
func SoporteExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string
	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre," +
		"centro.nombre as CentroNombre,soporte.neto," +
		"soporte.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	consulta += " FROM soporte "
	consulta += " inner join tercero on tercero.codigo=soporte.tercero "
	consulta += " inner join centro on centro.codigo=soporte.centro "
	consulta += " inner join almacenista on almacenista.codigo=soporte." +
		"almacenista "
	consulta += " ORDER BY cast(soporte.codigo as integer) ASC"
	t := []soporteLista{}
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
	f.SetCellValue("Sheet1", "A9", "LISTADO DE DOCUMENTOS SOPORTE")
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

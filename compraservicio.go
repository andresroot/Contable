package main

// INICIA COMPRASERVICIO IMPORTAR PAQUETES
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

// INICIA COMPRASERVICIO ESTRUCTURA JSON
type compraservicioJson struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Nombre string `json:"nombre"`
}

// INICIA COMPRASERVICIO ESTRUCTURA
type compraservicioLista struct {
	Codigo            string
	Fecha             time.Time
	Total             string
	Tercero           string
	TerceroNombre     string
	CentroNombre      string
	AlmacenistaNombre string
}

func DatosPedidocompraServicio(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	Tercero := mux.Vars(r)["tercero"]
	log.Println("inicio compra editar" + Codigo)
	db := dbConn()
	var res []pedidoservicio

	// traer PEDIDO
	v := pedidoservicio{}
	err := db.Get(&v, "SELECT * FROM pedidoservicio where codigo=$1 and Tercero =$2", Codigo, Tercero)
	var valida bool
	valida = true

	switch err {
	case nil:
		log.Printf("compra existe: %+v\n", v)
	case sql.ErrNoRows:
		log.Println("compra NO Existe")
		valida = false
	default:
		log.Printf("compra error: %s\n", err)
	}
	det := []pedidoserviciodetalleeditar{}
	t := tercero{}

	// trae datos si existe compra
	if valida == true {
		err2 := db.Select(&det, PedidoservicioConsultaDetalle(), Codigo)
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

// INICIA COMPRASERVICIO ESTRUCTURA
type compraservicio struct {
	Codigo                    string
	Pedido                    string
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
	Neto                      string
	Items                     string
	Formadepago               string
	Mediodepago               string
	Tercero                   string
	Almacenista               string
	Accion                    string
	Detalle                   []compraserviciodetalle       `json:"Detalle"`
	DetalleEditar             []compraserviciodetalleeditar `json:"DetalleEditar"`
	TerceroDetalle            tercero
	Tipo                      string
	Ret2201                   string
	Centro                    string
	FormadepagoDetalle        formadepago
	MediodepagoDetalle        mediodepago
	AlmacenistaDetalle        almacenista
	CentroDetalle             centro
	PorcentajeRetencionFuente string
	TotalRetencionFuente      string
	PorcentajeRetencionIca    string
	TotalRetencionIca         string
}

// INICIA COMPRASERVICIODETALLE ESTRUCTURA
type compraserviciodetalle struct {
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

// estructura para editar
type compraserviciodetalleeditar struct {
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
func CompraservicioConsultaDetalle() string {
	var consulta = ""
	consulta = "select "
	consulta += "compraserviciodetalle.Id as id ,"
	consulta += "compraserviciodetalle.Codigo as codigo,"
	consulta += "compraserviciodetalle.Fila as fila,"
	consulta += "compraserviciodetalle.Cantidad as cantidad,"
	consulta += "compraserviciodetalle.Precio as precio,"
	consulta += "compraserviciodetalle.Descuento as descuento,"
	consulta += "compraserviciodetalle.Montodescuento as montodescuento,"
	consulta += "compraserviciodetalle.Sigratis as sigratis,"
	consulta += "compraserviciodetalle.Subtotal as subtotal,"
	consulta += "compraserviciodetalle.Subtotaldescuento  as subtotaldescuento,"
	consulta += "compraserviciodetalle.Pagina as pagina ,"
	consulta += "compraserviciodetalle.Fecha as fecha,"
	consulta += "compraserviciodetalle.Nombreservicio as Nombreservicio,"
	consulta += "compraserviciodetalle.Unidadservicio as Unidadservicio, "
	consulta += "compraserviciodetalle.Codigoservicio as Codigoservicio, "
	consulta += "compraserviciodetalle.Ivaservicio as Ivaservicio,"
	consulta += "iva.nombre as Ivanombreservicio "
	consulta += "from compraserviciodetalle "
	consulta += "inner join iva on iva.codigo=compraserviciodetalle.Ivaservicio "
	consulta += " where compraserviciodetalle.codigo=$1"
	log.Println(consulta)
	return consulta
}

// INICIA COMPRASERVICIO LISTA
func CompraservicioLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/compraservicio/compraservicioLista.html")
	log.Println("Error compraservicio 0")
	var consulta string
	var miperiodo = periodoSesion(r)
	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre,centro.nombre as CentroNombre,total,compraservicio.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	consulta += " FROM compraservicio "
	consulta += " inner join tercero on tercero.codigo=compraservicio.tercero "
	consulta += " inner join centro on centro.codigo=compraservicio.centro "
	consulta += " inner join almacenista on almacenista.codigo=compraservicio.almacenista "
	consulta += " where  extract(year from fecha)=" + miperiodo + " "
	consulta += " ORDER BY compraservicio.codigo ASC"

	db := dbConn()
	res := []compraservicioLista{}

	err := db.Select(&res, consulta)

	if err != nil {
		fmt.Println(err)
		return
	}

	varmap := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
	}
	log.Println("Error compraservicio888")
	tmp.Execute(w, varmap)
}

// INICIA COMPRASERVICIO NUEVO
func CompraservicioNuevo(w http.ResponseWriter, r *http.Request) {

	// TRAE COPIA DE EDITAR
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio compraservicio editar" + Codigo)

	db := dbConn()
	v := compraservicio{}
	det := []compraserviciodetalleeditar{}
	if Codigo == "False" {

	} else {

		err := db.Get(&v, "SELECT * FROM compraservicio where codigo=$1", Codigo)
		if err != nil {
			log.Fatalln(err)
		}

		v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
		v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
		v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
		v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
		v.CentroDetalle = TraerCentroConsulta(v.Centro)

		err2 := db.Select(&det, CompraservicioConsultaDetalle(), Codigo)
		if err2 != nil {
			fmt.Println(err2)
		}
	}

	parametros := map[string]interface{}{
		"miperiodo":            periodoSesion(r),
		"resolucionventa":      ListaResolucionventa(),
		"codigo":               Codigo,
		"almacenista":          PrimerAlmacenista(),
		"compraservicio":       v,
		"detalle":              det,
		"hosting":              ruta,
		"retfte":               TraerParametrosInventario().Compracuentaporcentajeretfte,
		"ventaserviciotipoiva": TraerParametrosInventario().Ventaserviciotipoiva,
		"autoret2201":          TraerParametrosInventario().Ventacuentaporcentajeret2201,
	}
	//TERMINA TRAE COPIA DE EDITAR

	miTemplete, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/compraservicio/compraservicioNuevo.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaIva.html",
		"vista/compraservicio/compraservicioScript.html")

	fmt.Printf("%v, %v", miTemplete, err)
	log.Println("Error compraservicio nuevo 3")
	miTemplete.Execute(w, parametros)
}

// INICIA COMPRASERVICIO INSERTAR AJAX
func CompraservicioAgregar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var tempCompra compraservicio

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// carga informacion de la compraservicio
	err = json.Unmarshal(b, &tempCompra)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	parametrosinventario := configuracioninventario{}
	err = db.Get(&parametrosinventario, "SELECT * FROM configuracioninventario")
	if err != nil {
		panic(err.Error())
	}

	if tempCompra.Accion == "Actualizar" {

		// borra detalle anterior
		delForm, err := db.Prepare("DELETE from compraserviciodetalle WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(tempCompra.Codigo)

		// borra cabecera anterior
		delForm1, err := db.Prepare("DELETE from compraservicio WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm1.Exec(tempCompra.Codigo)
	}

	// INSERTA DETALLE
	for i, x := range tempCompra.Detalle {
		var a = i
		var q string
		q = "insert into compraserviciodetalle ("
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

		// TERMINA COMPRASERVICIO GRABAR INSERTAR
		_, err = insForm.Exec(
			x.Id,
			x.Codigo,
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

	// INICIA INSERTAR COMPRASERVICIO
	log.Println("Got %s age %s club %s\n", tempCompra.Codigo, tempCompra.Tercero, tempCompra.Total)
	var q string
	q = "insert into compraservicio ("
	q += "Codigo,"
	q += "Pedido,"
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
	q += "Almacenista,"
	q += "Centro,"
	q += "Tipo,"
	q += "PorcentajeRetencionFuente,"
	q += "TotalRetencionFuente,"
	q += "PorcentajeRetencionIca,"
	q += "TotalRetencionIca"
	q += " ) values("
	q += parametros(34)
	q += " ) "

	log.Println("Cadena SQL " + q)
	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	layout := "2006-01-02"
	log.Println("Hora", tempCompra.Fecha.Format("02/01/2006"))
	currentTime := time.Now()

	_, err = insForm.Exec(
		tempCompra.Codigo,
		tempCompra.Pedido,
		tempCompra.Fecha.Format(layout),
		tempCompra.Vence.Format(layout),
		currentTime.Format("3:4:5 PM"),
		Quitacoma(tempCompra.Descuento),
		Quitacoma(tempCompra.Subtotaldescuento19),
		Quitacoma(tempCompra.Subtotaldescuento5),
		Quitacoma(tempCompra.Subtotaldescuento0),
		Quitacoma(tempCompra.Subtotal),
		Quitacoma(tempCompra.Subtotal19),
		Quitacoma(tempCompra.Subtotal5),
		Quitacoma(tempCompra.Subtotal0),
		Quitacoma(tempCompra.Subtotaliva19),
		Quitacoma(tempCompra.Subtotaliva5),
		Quitacoma(tempCompra.Subtotaliva0),
		Quitacoma(tempCompra.Subtotalbase19),
		Quitacoma(tempCompra.Subtotalbase5),
		Quitacoma(tempCompra.Subtotalbase0),
		Quitacoma(tempCompra.TotalIva),
		Quitacoma(tempCompra.Ret2201),
		Quitacoma(tempCompra.Total),
		Quitacoma(tempCompra.Neto),
		tempCompra.Items,
		tempCompra.Formadepago,
		tempCompra.Mediodepago,
		tempCompra.Tercero,
		tempCompra.Almacenista,
		tempCompra.Centro,
		tempCompra.Tipo,
		tempCompra.PorcentajeRetencionFuente,
		Quitacoma(tempCompra.TotalRetencionFuente),
		tempCompra.PorcentajeRetencionIca,
		Quitacoma(tempCompra.TotalRetencionIca))

	if err != nil {
		panic(err)
	}

	// INSERTAR COMPROBANTE CONTABILIDAD
	var tempComprobante comprobante
	var tempComprobanteDetalle comprobantedetalle
	tempComprobante.Documento = "8"
	tempComprobante.Numero = tempCompra.Codigo
	tempComprobante.Fecha = tempCompra.Fecha
	tempComprobante.Fechaconsignacion = tempCompra.Fecha
	tempComprobante.Debito = tempCompra.Neto + ".00"
	tempComprobante.Credito = tempCompra.Neto + ".00"
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

	// INSERTAR CUENTA DEBITO COMPRAS 19%
	if tempCompra.Subtotal19 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempCompra.Subtotal19) - Flotante(tempCompra.Subtotaldescuento19)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Cscuenta19
		tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(tempCompra.Subtotal19) - Flotante(tempCompra.Subtotaldescuento19))
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteCompraservicio(tempComprobanteDetalle, tempComprobante, tempCompra)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA DEBITO COMPRAS 5%
	if tempCompra.Subtotal5 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempCompra.Subtotal5) - Flotante(tempCompra.Subtotaldescuento5)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Cscuenta5
		tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(tempCompra.Subtotal5) - Flotante(tempCompra.Subtotaldescuento5))
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteCompraservicio(tempComprobanteDetalle, tempComprobante, tempCompra)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA DEBITO COMPRAS 0%
	if tempCompra.Subtotal0 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempCompra.Subtotal0) - Flotante(tempCompra.Subtotaldescuento0)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Cscuenta0
		tempComprobanteDetalle.Debito = FormatoFlotanteComprobante(Flotante(tempCompra.Subtotal0) - Flotante(tempCompra.Subtotaldescuento0))
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteCompraservicio(tempComprobanteDetalle, tempComprobante, tempCompra)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA DEBITO IVA 19%
	if tempCompra.Subtotaliva19 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempCompra.Subtotaliva19)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Csiva19
		tempComprobanteDetalle.Debito = tempCompra.Subtotaliva19
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteCompraservicio(tempComprobanteDetalle, tempComprobante, tempCompra)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA DEBITO IVA 5%
	if tempCompra.Subtotaliva5 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempCompra.Subtotaliva5)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Csiva5
		tempComprobanteDetalle.Debito = tempCompra.Subtotaliva5
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteCompraservicio(tempComprobanteDetalle, tempComprobante, tempCompra)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA CREDITO DESCUENTO
	//if (tempCompra.Descuento!="0")	{
	//	fila=fila+1
	//	tempComprobanteDetalle.Fila=strconv.Itoa(fila)
	//	totalCredito+=Flotante(tempCompra.Descuento)
	//	tempComprobanteDetalle.Cuenta  = parametrosinventario.Compracuentadescuento
	//	tempComprobanteDetalle.Debito = ""
	//	tempComprobanteDetalle.Credito = tempCompra.Descuento
	//	InsertaDetalleComprobanteCompraservicio(tempComprobanteDetalle,tempComprobante,tempCompra)
	//	log.Println("credito linea" + fmt.Sprintf("%.2f",totalCredito))
	//}

	// INSERTAR CUENTA CREDITO RET. FTE.
	if tempCompra.TotalRetencionFuente != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempCompra.TotalRetencionFuente)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Cscuentaretfte
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = tempCompra.TotalRetencionFuente
		InsertaDetalleComprobanteCompraservicio(tempComprobanteDetalle, tempComprobante, tempCompra)
		log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito))
	}

	// INSERTAR CUENTA CREDITO RET. ICA.
	if tempCompra.TotalRetencionIca != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempCompra.TotalRetencionIca)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Cscuentaretica
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = tempCompra.TotalRetencionIca
		InsertaDetalleComprobanteCompraservicio(tempComprobanteDetalle, tempComprobante, tempCompra)
		log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito))
	}

	// INSERTAR CUENTA CREDITO PROVEEDOR
	if tempCompra.Neto != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempCompra.Neto)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Cscuentaproveedor
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = tempCompra.Neto
		InsertaDetalleComprobanteCompraservicio(tempComprobanteDetalle, tempComprobante, tempCompra)
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
func InsertaDetalleComprobanteCompraservicio(miFilaComprobante comprobantedetalle, miComprobante comprobante, miCompra compraservicio) {
	db := dbConn()

	// traer tercero
	miTercero := tercero{}
	err3 := db.Get(&miTercero, "SELECT * FROM tercero where codigo=$1", miCompra.Tercero)
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
		miCompra.Tercero,
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

// INICIA COMPRASERVICIO EXISTE
func CompraservicioExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM compraservicio  WHERE codigo=$1", Codigo)
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

// INICIA COMPRASERVICIO EDITAR
func CompraservicioEditar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio compraservicio editar" + Codigo)
	db := dbConn()

	// traer compraservicio
	v := compraservicio{}
	err := db.Get(&v, "SELECT * FROM compraservicio where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []compraserviciodetalleeditar{}

	err2 := db.Select(&det, CompraservicioConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
	v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
	v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
	v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
	v.CentroDetalle = TraerCentroConsulta(v.Centro)

	parametros := map[string]interface{}{
		"miperiodo":            periodoSesion(r),
		"compraservicio":       v,
		"detalle":              det,
		"hosting":              ruta,
		"retfte":               TraerParametrosInventario().Compracuentaporcentajeretfte,
		"ventaserviciotipoiva": TraerParametrosInventario().Ventaserviciotipoiva,
		"autoret2201":          TraerParametrosInventario().Ventacuentaporcentajeret2201,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/compraservicio/compraservicioEditar.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaIva.html",
		"vista/compraservicio/compraservicioScript.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error compraservicio nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA COMPRASERVICIO BORRAR
func CompraservicioBorrar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio compraservicio borrar" + Codigo)

	db := dbConn()

	// traer compraservicio
	v := compraservicio{}
	err := db.Get(&v, "SELECT * FROM compraservicio where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []compraserviciodetalleeditar{}
	err2 := db.Select(&det, CompraservicioConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
	v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
	v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
	v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
	v.CentroDetalle = TraerCentroConsulta(v.Centro)

	parametros := map[string]interface{}{
		"miperiodo":            periodoSesion(r),
		"Compraservicio":       v,
		"detalle":              det,
		"hosting":              ruta,
		"retfte":               TraerParametrosInventario().Compracuentaporcentajeretfte,
		"ventaserviciotipoiva": TraerParametrosInventario().Ventaserviciotipoiva,
		"autoret2201":          TraerParametrosInventario().Ventacuentaporcentajeret2201,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/compraservicio/compraservicioBorrar.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaIva.html",
		"vista/compraservicio/compraservicioScript.html")
	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error compraservicio nuevo 3")
	log.Println("Error compraservicio nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA COMPRA SERVICIO ELIMINAR
func CompraservicioEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	codigo := mux.Vars(r)["codigo"]

	// borrar compraservicio
	delForm, err := db.Prepare("DELETE from compraservicio WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(codigo)

	// borar detalle
	delForm1, err := db.Prepare("DELETE from compraserviciodetalle WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm1.Exec(codigo)

	log.Println("Registro Eliminado" + codigo)
	http.Redirect(w, r, "/CompraservicioLista", 301)
}

// INICIA COMPRA SERVICIO PDF
func CompraservicioPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]

	// traer compraservicio
	miCompraservicio := compraservicio{}
	err := db.Get(&miCompraservicio, "SELECT * FROM compraservicio where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}

	// traer detalle
	miDetalle := []compraserviciodetalleeditar{}
	err2 := db.Select(&miDetalle, CompraservicioConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	// traer tercero
	miTercero := tercero{}
	err3 := db.Get(&miTercero, "SELECT * FROM tercero where codigo=$1", miCompraservicio.Tercero)
	if err3 != nil {
		log.Fatalln(err3)
	}

	// traer Almacenista
	miAlmacenista := almacenista{}
	err4 := db.Get(&miAlmacenista, "SELECT * FROM almacenista where codigo=$1", miCompraservicio.Almacenista)
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

		// COMPRA NUMERO
		pdf.SetY(20)
		pdf.SetX(80)
		pdf.Ln(8)
		pdf.SetX(75)
		pdf.SetFont("Arial", "", 11)
		pdf.CellFormat(190, 10, "COMPRA SERVICIO", "0", 0, "C", false, 0, "")
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
	CompraservicioCabecera(pdf, miTercero, miCompraservicio, miAlmacenista)

	var filas = len(miDetalle)
	// menos de 32
	if filas <= 32 {
		for i, miFila := range miDetalle {
			var a = i + 1
			CompraservicioFilaDetalle(pdf, miFila, a)
		}
		CompraservicioPieDePagina(pdf, miTercero, miCompraservicio)
	} else {
		// mas de 32 y menos de 73
		if filas > 32 && filas <= 73 {
			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 41 {
					CompraservicioFilaDetalle(pdf, miFila, a)
				}
			}
			CompraLinea(pdf)
			// segunda pagina
			pdf.AddPage()
			CompraservicioCabecera(pdf, miTercero, miCompraservicio, miAlmacenista)
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 41 {
					CompraservicioFilaDetalle(pdf, miFila, a)
				}
			}

			CompraservicioPieDePagina(pdf, miTercero, miCompraservicio)
		} else {
			// mas de tres paginas

			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 41 {
					CompraservicioFilaDetalle(pdf, miFila, a)
				}
			}
			CompraLinea(pdf)
			pdf.AddPage()
			CompraservicioCabecera(pdf, miTercero, miCompraservicio, miAlmacenista)
			// segunda pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 41 && a <= 82 {
					CompraservicioFilaDetalle(pdf, miFila, a)
				}
			}
			CompraLinea(pdf)
			pdf.AddPage()
			CompraservicioCabecera(pdf, miTercero, miCompraservicio, miAlmacenista)
			// tercera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 82 {
					CompraservicioFilaDetalle(pdf, miFila, a)
				}
			}
			CompraservicioPieDePagina(pdf, miTercero, miCompraservicio)
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

func CompraservicioCabecera(pdf *gofpdf.Fpdf, miTercero tercero, miCompraservicio compraservicio, miAlmacenista almacenista) {
	ene := pdf.UnicodeTranslatorFromDescriptor("")
	// RELLENO TITULO
	pdf.SetFont("Arial", "", 10)
	pdf.SetY(44)
	pdf.SetX(20)
	pdf.SetFillColor(224, 231, 239)
	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(8)
	pdf.SetX(20)
	pdf.CellFormat(184, 6, "DATOS DEL COMPRADOR", "0", 0,
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
	pdf.CellFormat(40, 4, Titulo(TraerFormadepago(miCompraservicio.Formadepago)), "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Fecha de Expedicion", "", 0,
		"L", false, 0, "")
	pdf.SetX(150)
	pdf.CellFormat(40, 4, miCompraservicio.Fecha.Format("02/01/2006")+" "+Titulo(miCompraservicio.Hora), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)

	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Medio de Pago", "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, TraerMediodepago(miCompraservicio.Mediodepago), "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Fecha de Vencimiento", "", 0,
		"L", false, 0, "")
	pdf.SetX(150)
	pdf.CellFormat(40, 4, miCompraservicio.Vence.Format("02/01/2006"), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Condiciones", "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, "", "", 0,
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
func CompraservicioFilaDetalle(pdf *gofpdf.Fpdf, miFila compraserviciodetalleeditar, a int) {
	ene := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetX(21)
	var yinicial float64
	yinicial = pdf.GetY()
	pdf.CellFormat(40, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(40, 4, Subcadena(miFila.Codigoservicio, 0, 4), "", 0,
		"L", false, 0, "")
	var y float64
	y = pdf.GetY()
	pdf.SetX(42)
	pdf.MultiCell(68, 4, Mayuscula(miFila.Nombreservicio), "", "L", false)
	var yfinal float64
	yfinal = pdf.GetY()
	pdf.SetY(y)
	pdf.SetX(87)
	pdf.CellFormat(40, 4, ene(Titulo(Subcadena((miFila.Unidadservicio), 0, 6))), "", 0,
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

func CompraservicioPieDePagina(pdf *gofpdf.Fpdf, miTercero tercero, miCompraservicio compraservicio) {

	Totalletras, err := IntLetra(Cadenaentero(miCompraservicio.Total))
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

	// PRESENTA DATOS CON VALORES //
	pdf.SetFont("Arial", "", 9)
	var separador float64
	var altoseparador float64
	separador = 254
	altoseparador = -4

	// INICIA DATOS FACTURA
	if miCompraservicio.Neto != "0" {
		pdf.SetFont("Arial", "", 9)
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "TOTAL", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miCompraservicio.Neto), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miCompraservicio.TotalRetencionIca != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "Ret. Ica.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(170)
		pdf.CellFormat(190, 4, (miCompraservicio.PorcentajeRetencionIca)+"%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miCompraservicio.TotalRetencionIca), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miCompraservicio.TotalRetencionFuente != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "Ret. Fte.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(170)
		pdf.CellFormat(190, 4, (miCompraservicio.PorcentajeRetencionFuente)+"%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miCompraservicio.TotalRetencionFuente), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miCompraservicio.Subtotaliva5 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "I. V. A.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(170)
		pdf.CellFormat(190, 4, "5%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miCompraservicio.Subtotaliva5), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miCompraservicio.Subtotaliva19 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "I. V. A.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(170)
		pdf.CellFormat(190, 4, "19%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miCompraservicio.Subtotaliva19), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miCompraservicio.Subtotalbase5 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "GRAVABLE", "0", 0, "L",
			false, 0, "")
		pdf.SetX(170)
		pdf.CellFormat(190, 4, "5%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miCompraservicio.Subtotalbase5), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miCompraservicio.Subtotalbase19 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "GRAVABLE", "0", 0, "L",
			false, 0, "")
		pdf.SetX(170)
		pdf.CellFormat(190, 4, "19%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miCompraservicio.Subtotalbase19), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miCompraservicio.Subtotalbase0 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "NO GRAVABLE", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miCompraservicio.Subtotalbase0), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}
}
func CompraserviLinea(pdf *gofpdf.Fpdf) {
	pdf.SetY(259)
	pdf.Line(20, 259, 204, 259)
}

// INICIA COMPRA SERVICIO TODOS PDF
func CompraservicioTodosCabecera(pdf *gofpdf.Fpdf) {
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
	pdf.CellFormat(190, 6, "Almacenista", "0", 0,
		"L", false, 0, "")
	pdf.SetX(193)
	pdf.CellFormat(190, 6, "Total", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}

func CompraservicioTodosDetalle(pdf *gofpdf.Fpdf, miFila compraservicioLista, a int) {
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
	pdf.CellFormat(40, 4, Coma(miFila.Total), "", 0,
		"R", false, 0, "")
	pdf.Ln(4)
}

func CompraservicioTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string
	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre," +
		"centro.nombre as CentroNombre,compraservicio.total," +
		"compraservicio.codigo,fecha,tercero," +
		"tercero.nombre as TerceroNombre "
	consulta += " FROM compraservicio "
	consulta += " inner join tercero on tercero.codigo=compraservicio." +
		"tercero "
	consulta += " inner join centro on centro.codigo=compraservicio.centro "
	consulta += " inner join almacenista on almacenista.codigo=compraservicio." +
		"almacenista "
	consulta += " ORDER BY cast(compraservicio.codigo as integer) ASC"

	t := []compraservicioLista{}
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
		pdf.CellFormat(190, 10, "DATOS COMPRA SERVICIOS", "0", 0,
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

	CompraTodosCabecera(pdf)
	// tercera pagina

	for i, miFila := range t {
		var a = i + 1
		if math.Mod(float64(a), 49) == 0 {
			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			CompraservicioTodosCabecera(pdf)
		}
		CompraservicioTodosDetalle(pdf, miFila, a)
	}

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

// TERMINA COMPRA TODOS PDF

// COMPRA SERVICIO EXCEL
func CompraservicioExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string
	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre," +
		"centro.nombre as CentroNombre,compraservicio.total," +
		"compraservicio.codigo,fecha,tercero," +
		"tercero.nombre as TerceroNombre "
	consulta += " FROM compraservicio "
	consulta += " inner join tercero on tercero.codigo=compraservicio." +
		"tercero "
	consulta += " inner join centro on centro.codigo=compraservicio.centro "
	consulta += " inner join almacenista on almacenista.codigo=compraservicio." +
		"almacenista "
	consulta += " ORDER BY cast(compraservicio.codigo as integer) ASC"
	t := []compraservicioLista{}
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
	f.SetCellValue("Sheet1", "A9", "LISTADO DE COMPRA SERVICIOS")
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

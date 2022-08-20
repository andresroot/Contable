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
type devolucioncompraservicioJson struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Nombre string `json:"nombre"`
}

// INICIA COMPRASERVICIO ESTRUCTURA
type devolucioncompraservicioLista struct {
	Codigo            string
	Fecha             time.Time
	Total             string
	Tercero           string
	TerceroNombre     string
	CentroNombre      string
	AlmacenistaNombre string
}

func DatosdevolucioncompraServicio(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio devolucioncompra editar" + Codigo)
	db := dbConn()
	var res []devolucioncompraservicio

	// TRAER COMPRA SERVICIO
	v := devolucioncompraservicio{}
	err := db.Get(&v, "SELECT * FROM devolucioncompraservicio where codigo=$1", Codigo)
	var valida bool
	valida = true

	switch err {
	case nil:
		log.Printf("devolucioncompraservicio existe: %+v\n", v)
	case sql.ErrNoRows:
		log.Println("devolucion compra NO Existe")
		valida = false
	default:
		log.Printf("compra error: %s\n", err)
	}
	det := []compraserviciodetalleeditar{}
	t := tercero{}

	// trae datos si existe compra
	if valida == true {
		err2 := db.Select(&det, DevolucioncompraservicioConsultaDetalle(), Codigo)
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

func DatoscompraServicio(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	Tercero := mux.Vars(r)["tercero"]
	log.Println("inicio compra editar" + Codigo)
	db := dbConn()
	var res []compraservicio

	// TRAER COMPRA SERVICIO
	v := compraservicio{}
	err := db.Get(&v, "SELECT * FROM compraservicio where codigo=$1 and Tercero = $2", Codigo, Tercero)
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
	det := []compraserviciodetalleeditar{}
	t := tercero{}

	// trae datos si existe compra
	if valida == true {
		err2 := db.Select(&det, CompraservicioConsultaDetalle(), Codigo)
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
type devolucioncompraservicio struct {
	Codigo                    string
	Compra                    string
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
type devolucioncompraserviciodetalle struct {
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
type devolucioncompraserviciodetalleeditar struct {
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
func DevolucioncompraservicioConsultaDetalle() string {
	var consulta = ""
	consulta = "select "
	consulta += "devolucioncompraserviciodetalle.Id as id ,"
	consulta += "devolucioncompraserviciodetalle.Codigo as codigo,"
	consulta += "devolucioncompraserviciodetalle.Fila as fila,"
	consulta += "devolucioncompraserviciodetalle.Cantidad as cantidad,"
	consulta += "devolucioncompraserviciodetalle.Precio as precio,"
	consulta += "devolucioncompraserviciodetalle.Descuento as descuento,"
	consulta += "devolucioncompraserviciodetalle.Montodescuento as montodescuento,"
	consulta += "devolucioncompraserviciodetalle.Sigratis as sigratis,"
	consulta += "devolucioncompraserviciodetalle.Subtotal as subtotal,"
	consulta += "devolucioncompraserviciodetalle.Subtotaldescuento  as subtotaldescuento,"
	consulta += "devolucioncompraserviciodetalle.Pagina as pagina ,"
	consulta += "devolucioncompraserviciodetalle.Fecha as fecha,"
	consulta += "devolucioncompraserviciodetalle.Nombreservicio as Nombreservicio,"
	consulta += "devolucioncompraserviciodetalle.Unidadservicio as Unidadservicio, "
	consulta += "devolucioncompraserviciodetalle.Codigoservicio as Codigoservicio, "
	consulta += "devolucioncompraserviciodetalle.Ivaservicio as Ivaservicio, "
	consulta += "iva.nombre as Ivanombreservicio "
	consulta += "from devolucioncompraserviciodetalle "
	consulta += "inner join iva on iva.codigo=devolucioncompraserviciodetalle.Ivaservicio "
	consulta += " where devolucioncompraserviciodetalle.codigo=$1"
	log.Println(consulta)
	return consulta
}

// INICIA COMPRASERVICIO LISTA
func DevolucioncompraservicioLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/devolucioncompraservicio/devolucioncompraservicioLista.html")
	log.Println("Error devolucioncompraservicio 0")
	var consulta string
	var miperiodo = periodoSesion(r)
	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre,centro.nombre as CentroNombre,total,devolucioncompraservicio.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	consulta += " FROM devolucioncompraservicio "
	consulta += " inner join tercero on tercero.codigo=devolucioncompraservicio.tercero "
	consulta += " inner join centro on centro.codigo=devolucioncompraservicio.centro "
	consulta += " inner join almacenista on almacenista.codigo=devolucioncompraservicio.almacenista "
	consulta += " where  extract(year from fecha)=" + miperiodo + " "
	consulta += " ORDER BY devolucioncompraservicio.codigo ASC"

	db := dbConn()
	res := []devolucioncompraservicioLista{}

	err := db.Select(&res, consulta)

	if err != nil {
		fmt.Println(err)
		return
	}

	varmap := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
	}
	log.Println("Error devolucioncompraservicio888")
	tmp.Execute(w, varmap)
}

// INICIA COMPRASERVICIO NUEVO
func DevolucioncompraservicioNuevo(w http.ResponseWriter, r *http.Request) {

	// TRAE COPIA DE EDITAR
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio devolucioncompraservicio editar" + Codigo)

	db := dbConn()
	v := devolucioncompraservicio{}
	det := []devolucioncompraserviciodetalleeditar{}
	if Codigo == "False" {

	} else {

		err := db.Get(&v, "SELECT * FROM devolucioncompraservicio where codigo=$1", Codigo)
		if err != nil {
			log.Fatalln(err)
		}

		v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
		v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
		v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
		v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
		v.CentroDetalle = TraerCentroConsulta(v.Centro)

		err2 := db.Select(&det, DevolucioncompraservicioConsultaDetalle(), Codigo)
		if err2 != nil {
			fmt.Println(err2)
		}
	}

	parametros := map[string]interface{}{
		"miperiodo":                periodoSesion(r),
		"codigo":                   Codigo,
		"almacenista":              PrimerAlmacenista(),
		"devolucioncompraservicio": v,
		"detalle":                  det,
		"hosting":                  ruta,
		"retfte":                   TraerParametrosInventario().Compracuentaporcentajeretfte,
		"ventaserviciotipoiva":     TraerParametrosInventario().Ventaserviciotipoiva,
		"autoret2201":              TraerParametrosInventario().Ventacuentaporcentajeret2201,
	}
	//TERMINA TRAE COPIA DE EDITAR

	miTemplete, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/devolucioncompraservicio/devolucioncompraservicioNuevo.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaIva.html",
		"vista/devolucioncompraservicio/devolucioncompraservicioScript.html")
	if err != nil {
		log.Println(err.Error())
	}
	fmt.Printf("%v, %v", miTemplete, err)
	log.Println("Error devolucioncompraservicio nuevo 3")
	miTemplete.Execute(w, parametros)
}

// INICIA COMPRASERVICIO INSERTAR AJAX
func DevolucioncompraservicioAgregar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var tempDevolucioncompra devolucioncompraservicio

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// carga informacion de la compraservicio
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
		delForm, err := db.Prepare("DELETE from devolucioncompraserviciodetalle WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(tempDevolucioncompra.Codigo)

		// borra cabecera anterior
		delForm1, err := db.Prepare("DELETE from devolucioncompraservicio WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm1.Exec(tempDevolucioncompra.Codigo)
	}

	// INSERTA DETALLE
	for i, x := range tempDevolucioncompra.Detalle {
		var a = i
		var q string
		q = "insert into devolucioncompraserviciodetalle ("
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

	// INICIA INSERTAR DEVOLUCION COMPRASERVICIO
	log.Println("Got %s age %s club %s\n", tempDevolucioncompra.Codigo, tempDevolucioncompra.Tercero, tempDevolucioncompra.Total)
	var q string
	q = "insert into devolucioncompraservicio ("
	q += "Codigo,"
	q += "Compra,"
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
	log.Println("Hora", tempDevolucioncompra.Fecha.Format("02/01/2006"))
	currentTime := time.Now()

	_, err = insForm.Exec(
		tempDevolucioncompra.Codigo,
		tempDevolucioncompra.Compra,
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
		Quitacoma(tempDevolucioncompra.Ret2201),
		Quitacoma(tempDevolucioncompra.Total),
		Quitacoma(tempDevolucioncompra.Neto),
		tempDevolucioncompra.Items,
		tempDevolucioncompra.Formadepago,
		tempDevolucioncompra.Mediodepago,
		tempDevolucioncompra.Tercero,
		tempDevolucioncompra.Almacenista,
		tempDevolucioncompra.Centro,
		tempDevolucioncompra.Tipo,
		tempDevolucioncompra.PorcentajeRetencionFuente,
		Quitacoma(tempDevolucioncompra.TotalRetencionFuente),
		tempDevolucioncompra.PorcentajeRetencionIca,
		Quitacoma(tempDevolucioncompra.TotalRetencionIca))

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
		tempComprobanteDetalle.Cuenta = parametrosinventario.Cscuentaproveedor
		tempComprobanteDetalle.Debito = tempDevolucioncompra.Neto
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteDevolucioncompraservicio(tempComprobanteDetalle, tempComprobante, tempDevolucioncompra)
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
	//	InsertaDetalleComprobanteDevolucioncompraservicio(tempComprobanteDetalle,tempComprobante,tempDevolucioncompra)
	//	log.Println("debito linea" + fmt.Sprintf("%.2f",totalDebito))
	//}

	// INSERTAR CUENTA DEBITO RET. FTE.
	if tempDevolucioncompra.TotalRetencionFuente != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempDevolucioncompra.TotalRetencionFuente)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Csdevolucioncuentaretfte
		tempComprobanteDetalle.Debito = tempDevolucioncompra.TotalRetencionFuente
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteDevolucioncompraservicio(tempComprobanteDetalle, tempComprobante, tempDevolucioncompra)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA DEBITO RET. ICA.
	if tempDevolucioncompra.TotalRetencionIca != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalDebito += Flotante(tempDevolucioncompra.TotalRetencionIca)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Csdevolucioncuentaretica
		tempComprobanteDetalle.Debito = tempDevolucioncompra.TotalRetencionIca
		tempComprobanteDetalle.Credito = ""
		InsertaDetalleComprobanteDevolucioncompraservicio(tempComprobanteDetalle, tempComprobante, tempDevolucioncompra)
		log.Println("debito linea" + fmt.Sprintf("%.2f", totalDebito))
	}

	// INSERTAR CUENTA CREDITO IVA 19%
	if tempDevolucioncompra.Subtotaliva19 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempDevolucioncompra.Subtotaliva19)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Csdevolucioniva19
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = tempDevolucioncompra.Subtotaliva19
		InsertaDetalleComprobanteDevolucioncompraservicio(tempComprobanteDetalle, tempComprobante, tempDevolucioncompra)
		log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito))
	}

	// INSERTAR CUENTA CREDITO IVA 5%
	if tempDevolucioncompra.Subtotaliva5 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempDevolucioncompra.Subtotaliva5)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Csdevolucioniva5
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = tempDevolucioncompra.Subtotaliva5
		InsertaDetalleComprobanteDevolucioncompraservicio(tempComprobanteDetalle, tempComprobante, tempDevolucioncompra)
		log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito))
	}

	// INSERTAR CUENTA CREDITO COMPRAS 19%
	if tempDevolucioncompra.Subtotal19 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempDevolucioncompra.Subtotal19) - Flotante(tempDevolucioncompra.Subtotaldescuento19)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Csdevolucioncuenta19
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(tempDevolucioncompra.Subtotal19) - Flotante(tempDevolucioncompra.Subtotaldescuento19))
		InsertaDetalleComprobanteDevolucioncompraservicio(tempComprobanteDetalle, tempComprobante, tempDevolucioncompra)
		log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito))
	}

	// INSERTAR CUENTA CREDITO COMPRAS 5%
	if tempDevolucioncompra.Subtotal5 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempDevolucioncompra.Subtotal5) - Flotante(tempDevolucioncompra.Subtotaldescuento5)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Csdevolucioncuenta5
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(tempDevolucioncompra.Subtotal5) - Flotante(tempDevolucioncompra.Subtotaldescuento5))
		InsertaDetalleComprobanteDevolucioncompraservicio(tempComprobanteDetalle, tempComprobante, tempDevolucioncompra)
		log.Println("credito linea" + fmt.Sprintf("%.2f", totalCredito))
	}

	// INSERTAR CUENTA CREDITO COMPRAS 0%
	if tempDevolucioncompra.Subtotal0 != "0" {
		fila = fila + 1
		tempComprobanteDetalle.Fila = strconv.Itoa(fila)
		totalCredito += Flotante(tempDevolucioncompra.Subtotal0) - Flotante(tempDevolucioncompra.Subtotaldescuento0)
		tempComprobanteDetalle.Cuenta = parametrosinventario.Csdevolucioncuenta0
		tempComprobanteDetalle.Debito = ""
		tempComprobanteDetalle.Credito = FormatoFlotanteComprobante(Flotante(tempDevolucioncompra.Subtotal0) - Flotante(tempDevolucioncompra.Subtotaldescuento0))
		InsertaDetalleComprobanteDevolucioncompraservicio(tempComprobanteDetalle, tempComprobante, tempDevolucioncompra)
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

// INICIA INSERTAR COMPROBANTE DE DEVOLUCION EN COMPRA
func InsertaDetalleComprobanteDevolucioncompraservicio(miFilaComprobante comprobantedetalle, miComprobante comprobante, miDevolucioncompra devolucioncompraservicio) {
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

// INICIA DEVOLUCION COMPRASERVICIO EXISTE
func DevolucioncompraservicioExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM devolucioncompraservicio  WHERE codigo=$1", Codigo)
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
func DevolucioncompraservicioEditar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio devolucioncompraservicio editar" + Codigo)
	db := dbConn()

	// traer devolucioncompraservicio
	v := devolucioncompraservicio{}
	err := db.Get(&v, "SELECT * FROM devolucioncompraservicio where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []devolucioncompraserviciodetalleeditar{}

	err2 := db.Select(&det, DevolucioncompraservicioConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
	v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
	v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
	v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
	v.CentroDetalle = TraerCentroConsulta(v.Centro)

	parametros := map[string]interface{}{
		"miperiodo":                periodoSesion(r),
		"devolucioncompraservicio": v,
		"detalle":                  det,
		"hosting":                  ruta,
		"retfte":                   TraerParametrosInventario().Compracuentaporcentajeretfte,
		"ventaserviciotipoiva":     TraerParametrosInventario().Ventaserviciotipoiva,
		"autoret2201":              TraerParametrosInventario().Ventacuentaporcentajeret2201,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/devolucioncompraservicio/devolucioncompraservicioEditar.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaIva.html",
		"vista/devolucioncompraservicio/devolucioncompraservicioScript.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error devolucioncompraservicio nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA DEVOLUCION COMPRASERVICIO BORRAR
func DevolucioncompraservicioBorrar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio devolucioncompraservicio borrar" + Codigo)

	db := dbConn()

	// traer devolucioncompraservicio
	v := devolucioncompraservicio{}
	err := db.Get(&v, "SELECT * FROM devolucioncompraservicio where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []devolucioncompraserviciodetalleeditar{}
	err2 := db.Select(&det, DevolucioncompraservicioConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
	v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
	v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
	v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
	v.CentroDetalle = TraerCentroConsulta(v.Centro)

	parametros := map[string]interface{}{
		"miperiodo":                periodoSesion(r),
		"Devolucioncompraservicio": v,
		"detalle":                  det,
		"hosting":                  ruta,
		"retfte":                   TraerParametrosInventario().Compracuentaporcentajeretfte,
		"ventaserviciotipoiva":     TraerParametrosInventario().Ventaserviciotipoiva,
		"autoret2201":              TraerParametrosInventario().Ventacuentaporcentajeret2201,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/devolucioncompraservicio/devolucioncompraservicioBorrar.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaIva.html",
		"vista/devolucioncompraservicio/devolucioncompraservicioScript.html")
	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error devolucioncompraservicio nuevo 3")
	log.Println("Error devolucioncompraservicio nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA DEVOLUCION COMPRA SERVICIO ELIMINAR
func DevolucioncompraservicioEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	codigo := mux.Vars(r)["codigo"]

	// borrar devolucioncompraservicio
	delForm, err := db.Prepare("DELETE from devolucioncompraservicio WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(codigo)

	// borar detalle
	delForm1, err := db.Prepare("DELETE from devolucioncompraserviciodetalle WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm1.Exec(codigo)

	log.Println("Registro Eliminado" + codigo)
	http.Redirect(w, r, "/DevolucioncompraservicioLista", 301)
}

// INICIA COMPRA SERVICIO PDF
func DevolucioncompraservicioPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]

	// traer devolucioncompraservicio
	miDevolucioncompraservicio := devolucioncompraservicio{}
	err := db.Get(&miDevolucioncompraservicio, "SELECT * FROM devolucioncompraservicio where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}

	// traer detalle
	miDetalle := []devolucioncompraserviciodetalleeditar{}
	err2 := db.Select(&miDetalle, DevolucioncompraservicioConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	// traer tercero
	miTercero := tercero{}
	err3 := db.Get(&miTercero, "SELECT * FROM tercero where codigo=$1", miDevolucioncompraservicio.Tercero)
	if err3 != nil {
		log.Fatalln(err3)
	}

	// traer Almacenista
	miAlmacenista := almacenista{}
	err4 := db.Get(&miAlmacenista, "SELECT * FROM almacenista where codigo=$1", miDevolucioncompraservicio.Almacenista)
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
	DevolucioncompraservicioCabecera(pdf, miTercero, miDevolucioncompraservicio, miAlmacenista)

	var filas = len(miDetalle)
	// menos de 32
	if filas <= 32 {
		for i, miFila := range miDetalle {
			var a = i + 1
			DevolucioncompraservicioFilaDetalle(pdf, miFila, a)
		}
		DevolucioncompraservicioPieDePagina(pdf, miTercero, miDevolucioncompraservicio)
	} else {
		// mas de 32 y menos de 73
		if filas > 32 && filas <= 73 {
			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 41 {
					DevolucioncompraservicioFilaDetalle(pdf, miFila, a)
				}
			}
			DevolucioncompraservicioLinea(pdf)
			// segunda pagina
			pdf.AddPage()
			DevolucioncompraservicioCabecera(pdf, miTercero, miDevolucioncompraservicio, miAlmacenista)
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 41 {
					DevolucioncompraservicioFilaDetalle(pdf, miFila, a)
				}
			}

			DevolucioncompraservicioPieDePagina(pdf, miTercero, miDevolucioncompraservicio)
		} else {
			// mas de tres paginas

			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 41 {
					DevolucioncompraservicioFilaDetalle(pdf, miFila, a)
				}
			}
			DevolucioncompraservicioLinea(pdf)
			pdf.AddPage()
			DevolucioncompraservicioCabecera(pdf, miTercero, miDevolucioncompraservicio, miAlmacenista)
			// segunda pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 41 && a <= 82 {
					DevolucioncompraservicioFilaDetalle(pdf, miFila, a)
				}
			}
			DevolucioncompraservicioLinea(pdf)
			pdf.AddPage()
			DevolucioncompraservicioCabecera(pdf, miTercero, miDevolucioncompraservicio, miAlmacenista)
			// tercera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 82 {
					DevolucioncompraservicioFilaDetalle(pdf, miFila, a)
				}
			}
			DevolucioncompraservicioPieDePagina(pdf, miTercero, miDevolucioncompraservicio)
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

func DevolucioncompraservicioCabecera(pdf *gofpdf.Fpdf, miTercero tercero, miDevolucioncompraservicio devolucioncompraservicio, miAlmacenista almacenista) {
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
	pdf.CellFormat(40, 4, Titulo(TraerFormadepago(miDevolucioncompraservicio.Formadepago)), "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Fecha de Expedicion", "", 0,
		"L", false, 0, "")
	pdf.SetX(150)
	pdf.CellFormat(40, 4, miDevolucioncompraservicio.Fecha.Format("02/01/2006")+" "+Titulo(miDevolucioncompraservicio.Hora), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)

	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Medio de Pago", "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, TraerMediodepago(miDevolucioncompraservicio.Mediodepago), "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Fecha de Vencimiento", "", 0,
		"L", false, 0, "")
	pdf.SetX(150)
	pdf.CellFormat(40, 4, miDevolucioncompraservicio.Vence.Format("02/01/2006"), "", 0,
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
func DevolucioncompraservicioFilaDetalle(pdf *gofpdf.Fpdf, miFila devolucioncompraserviciodetalleeditar, a int) {
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

func DevolucioncompraservicioPieDePagina(pdf *gofpdf.Fpdf, miTercero tercero, miDevolucioncompraservicio devolucioncompraservicio) {

	Totalletras, err := IntLetra(Cadenaentero(miDevolucioncompraservicio.Total))
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
	if miDevolucioncompraservicio.Neto != "0" {
		pdf.SetFont("Arial", "", 9)
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "TOTAL", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miDevolucioncompraservicio.Neto), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucioncompraservicio.TotalRetencionIca != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "Ret. Ica.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(170)
		pdf.CellFormat(190, 4, (miDevolucioncompraservicio.PorcentajeRetencionIca)+"%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miDevolucioncompraservicio.TotalRetencionIca), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucioncompraservicio.TotalRetencionFuente != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "Ret. Fte.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(170)
		pdf.CellFormat(190, 4, (miDevolucioncompraservicio.PorcentajeRetencionFuente)+"%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miDevolucioncompraservicio.TotalRetencionFuente), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucioncompraservicio.Subtotaliva5 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "I. V. A.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(170)
		pdf.CellFormat(190, 4, "5%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miDevolucioncompraservicio.Subtotaliva5), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucioncompraservicio.Subtotaliva19 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "I. V. A.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(170)
		pdf.CellFormat(190, 4, "19%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miDevolucioncompraservicio.Subtotaliva19), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucioncompraservicio.Subtotalbase5 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "GRAVABLE", "0", 0, "L",
			false, 0, "")
		pdf.SetX(170)
		pdf.CellFormat(190, 4, "5%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miDevolucioncompraservicio.Subtotalbase5), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucioncompraservicio.Subtotalbase19 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "GRAVABLE", "0", 0, "L",
			false, 0, "")
		pdf.SetX(170)
		pdf.CellFormat(190, 4, "19%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miDevolucioncompraservicio.Subtotalbase19), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miDevolucioncompraservicio.Subtotalbase0 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "NO GRAVABLE", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miDevolucioncompraservicio.Subtotalbase0), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}
}
func DevolucioncompraservicioLinea(pdf *gofpdf.Fpdf) {
	pdf.SetY(259)
	pdf.Line(20, 259, 204, 259)
}

// INICIA COMPRA SERVICIO TODOS PDF
func DevolucioncompraservicioTodosCabecera(pdf *gofpdf.Fpdf) {
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

func DevolucioncompraservicioTodosDetalle(pdf *gofpdf.Fpdf, miFila devolucioncompraservicioLista, a int) {
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

func DevolucioncompraservicioTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string
	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre," +
		"centro.nombre as CentroNombre,devolucioncompraservicio.total," +
		"devolucioncompraservicio.codigo,fecha,tercero," +
		"tercero.nombre as TerceroNombre "
	consulta += " FROM devolucioncompraservicio "
	consulta += " inner join tercero on tercero.codigo=devolucioncompraservicio." +
		"tercero "
	consulta += " inner join centro on centro.codigo=devolucioncompraservicio.centro "
	consulta += " inner join almacenista on almacenista.codigo=devolucioncompraservicio." +
		"almacenista "
	consulta += " ORDER BY cast(devolucioncompraservicio.codigo as integer) ASC"

	t := []devolucioncompraservicioLista{}
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

	DevolucioncompraTodosCabecera(pdf)
	// tercera pagina

	for i, miFila := range t {
		var a = i + 1
		if math.Mod(float64(a), 49) == 0 {
			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			DevolucioncompraservicioTodosCabecera(pdf)
		}
		DevolucioncompraservicioTodosDetalle(pdf, miFila, a)
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
func DevolucioncompraservicioExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string
	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre," +
		"centro.nombre as CentroNombre,devolucioncompraservicio.total," +
		"devolucioncompraservicio.codigo,fecha,tercero," +
		"tercero.nombre as TerceroNombre "
	consulta += " FROM devolucioncompraservicio "
	consulta += " inner join tercero on tercero.codigo=devolucioncompraservicio." +
		"tercero "
	consulta += " inner join centro on centro.codigo=devolucioncompraservicio.centro "
	consulta += " inner join almacenista on almacenista.codigo=devolucioncompraservicio." +
		"almacenista "
	consulta += " ORDER BY cast(devolucioncompraservicio.codigo as integer) ASC"
	t := []devolucioncompraservicioLista{}
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

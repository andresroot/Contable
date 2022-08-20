package main

// INICIA PEDIDO SOPORTE IMPORTAR PAQUETES
import (
	"bytes"
	"github.com/360EntSecGroup-Skylar/excelize"
	"math"
	//"database/sql"
	"encoding/json"
	_ "encoding/json"
	"fmt"
	_ "github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
	_ "github.com/lib/pq"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// INICIA PEDIDO SOPORTE ESTRUCTURA JSON
type pedidosoporteJson struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Nombre string `json:"nombre"`
}

// INICIA PEDIDO SOPORTE ESTRUCTURA
type pedidosoporteLista struct {
	Codigo            string
	Fecha             time.Time
	Neto              string
	Tercero           string
	TerceroNombre     string
	CentroNombre      string
	AlmacenistaNombre string
}

// INICIA PEDIDO SOPORTE ESTRUCTURA
type pedidosoporte struct {
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
	Detalle                   []pedidosoportedetalle       `json:"Detalle"`
	DetalleEditar             []pedidosoportedetalleeditar `json:"DetalleEditar"`
	TerceroDetalle            tercero
	Tipo                      string
	Centro                    string
	FormadepagoDetalle        formadepago
	MediodepagoDetalle        mediodepago
	AlmacenistaDetalle        almacenista
	CentroDetalle             centro
}

// INICIA PEDIDO SOPORTEDETALLE ESTRUCTURA
type pedidosoportedetalle struct {
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
type pedidosoportedetalleeditar struct {
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
func PedidosoporteConsultaDetalle() string {
	var consulta = ""
	consulta = "select "
	consulta += "pedidosoportedetalle.Id as id ,"
	consulta += "pedidosoportedetalle.Codigo as codigo,"
	consulta += "pedidosoportedetalle.Fila as fila,"
	consulta += "pedidosoportedetalle.Cantidad as cantidad,"
	consulta += "pedidosoportedetalle.Precio as precio,"
	consulta += "pedidosoportedetalle.Descuento as descuento,"
	consulta += "pedidosoportedetalle.Montodescuento as montodescuento,"
	consulta += "pedidosoportedetalle.Sigratis as sigratis,"
	consulta += "pedidosoportedetalle.Subtotal as subtotal,"
	consulta += "pedidosoportedetalle.Subtotaldescuento  as subtotaldescuento,"
	consulta += "pedidosoportedetalle.Pagina as pagina ,"
	consulta += "pedidosoportedetalle.Bodega as bodega,"
	consulta += "bodega.Nombre as bodeganombre,"
	consulta += "pedidosoportedetalle.Producto as producto,"
	consulta += "pedidosoportedetalle.Fecha as fecha,"
	consulta += "producto.nombre as ProductoNombre, "
	consulta += "producto.iva as ProductoIva, "
	consulta += "producto.unidad as ProductoUnidad "
	consulta += "from pedidosoportedetalle "
	consulta += "inner join producto on producto.codigo=pedidosoportedetalle.producto "
	consulta += "inner join bodega on bodega.codigo=pedidosoportedetalle.bodega "
	consulta += " where pedidosoportedetalle.codigo=$1 ORDER BY fila"
	log.Println(consulta)
	return consulta
}

// INICIA PEDIDO SOPORTE LISTA
func PedidosoporteLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/pedidosoporte/pedidosoporteLista.html")
	log.Println("Error pedidosoporte 0")
	var consulta string
	var miperiodo = periodoSesion(r)
	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre,centro.nombre as CentroNombre,pedidosoporte.neto,pedidosoporte.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	consulta += " FROM pedidosoporte "
	consulta += " inner join tercero on tercero.codigo=pedidosoporte.tercero "
	consulta += " inner join centro on centro.codigo=pedidosoporte.centro "
	consulta += " inner join almacenista on almacenista.codigo=pedidosoporte.almacenista "
	consulta += " where  extract(year from fecha)=" + miperiodo + " "
	consulta += " ORDER BY cast(pedidosoporte.codigo as integer) ASC"

	db := dbConn()
	res := []pedidosoporteLista{}

	err := db.Select(&res, consulta)

	if err != nil {
		fmt.Println(err)
	}

	varmap := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
	}
	log.Println("Error pedidosoporte888")
	tmp.Execute(w, varmap)
}

// INICIA PEDIDO SOPORTE NUEVO
func PedidosoporteNuevo(w http.ResponseWriter, r *http.Request) {

	// TRAE COPIA DE EDITAR
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio pedidosoporte editar" + Codigo)

	db := dbConn()
	v := pedidosoporte{}
	det := []pedidosoportedetalleeditar{}
	if Codigo == "False" {

	} else {

		err := db.Get(&v, "SELECT * FROM pedidosoporte where codigo=$1", Codigo)
		if err != nil {
			log.Fatalln(err)
		}

		v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
		v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
		v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
		v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
		v.CentroDetalle = TraerCentroConsulta(v.Centro)
		err2 := db.Select(&det, PedidosoporteConsultaDetalle(), Codigo)
		if err2 != nil {
			fmt.Println(err2)
		}
	}

	parametros := map[string]interface{}{
		"codigo":        Codigo,
		"pedidosoporte": v,
		"almacenista":   PrimerAlmacenista(),
		"detalle":       det,
		"hosting":       ruta,
		"miperiodo":     periodoSesion(r),
		"retfte":        TraerParametrosInventario().Compracuentaporcentajeretfte,
	}
	//TERMINA TRAE COPIA DE EDITAR

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/pedidosoporte/pedidosoporteNuevo.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaBodega.html",
		"vista/autocompleta/autocompletaTercero.html",
		"vista/pedidosoporte/pedidosoporteScript.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error pedidosoporte nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA PEDIDO SOPORTE INSERTAR AJAX
func PedidosoporteAgregar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var tempPedidosoporte pedidosoporte

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// carga informacion de la PEDIDO SOPORTE
	err = json.Unmarshal(b, &tempPedidosoporte)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	parametrosinventario := configuracioninventario{}
	err = db.Get(&parametrosinventario, "SELECT * FROM configuracioninventario")
	if err != nil {
		panic(err.Error())
	}

	if tempPedidosoporte.Accion == "Actualizar" {

		// borra detalle anterior
		delForm, err := db.Prepare("DELETE from pedidosoportedetalle WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(tempPedidosoporte.Codigo)

		// borra cabecera anterior
		delForm1, err := db.Prepare("DELETE from pedidosoporte WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm1.Exec(tempPedidosoporte.Codigo)
	}

	// INSERTA DETALLE
	for i, x := range tempPedidosoporte.Detalle {
		var a = i
		var q string
		q = "insert into pedidosoportedetalle ("
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

		// TERMINA PEDIDO SOPORTE GRABAR INSERTAR
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

	// INICIA INSERTAR PEDIDO SOPORTE
	log.Println("Got %s age %s club %s\n", tempPedidosoporte.Codigo, tempPedidosoporte.Tercero, tempPedidosoporte.Neto)
	var q string
	q = "insert into pedidosoporte ("
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
	q += "PorcentajeRetencionFuente,"
	q += "TotalRetencionFuente,"
	q += "PorcentajeRetencionIca,"
	q += "TotalRetencionIca,"
	q += "Neto,"
	q += "Items,"
	q += "Formadepago,"
	q += "Mediodepago,"
	q += "Tercero,"
	q += "Almacenista,"
	q += "Centro,"
	q += "Tipo"
	q += " ) values("
	q += parametros(32)
	q += " ) "

	log.Println("Cadena SQL " + q)
	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	layout := "2006-01-02"
	log.Println("Hora", tempPedidosoporte.Fecha.Format("02/01/2006"))
	currentTime := time.Now()

	_, err = insForm.Exec(
		tempPedidosoporte.Codigo,
		tempPedidosoporte.Fecha.Format(layout),
		tempPedidosoporte.Vence.Format(layout),
		currentTime.Format("3:4:5 PM"),
		Quitacoma(tempPedidosoporte.Descuento),
		Quitacoma(tempPedidosoporte.Subtotaldescuento19),
		Quitacoma(tempPedidosoporte.Subtotaldescuento5),
		Quitacoma(tempPedidosoporte.Subtotaldescuento0),
		Quitacoma(tempPedidosoporte.Subtotal),
		Quitacoma(tempPedidosoporte.Subtotal19),
		Quitacoma(tempPedidosoporte.Subtotal5),
		Quitacoma(tempPedidosoporte.Subtotal0),
		Quitacoma(tempPedidosoporte.Subtotaliva19),
		Quitacoma(tempPedidosoporte.Subtotaliva5),
		Quitacoma(tempPedidosoporte.Subtotaliva0),
		Quitacoma(tempPedidosoporte.Subtotalbase19),
		Quitacoma(tempPedidosoporte.Subtotalbase5),
		Quitacoma(tempPedidosoporte.Subtotalbase0),
		Quitacoma(tempPedidosoporte.TotalIva),
		Quitacoma(tempPedidosoporte.Total),
		Quitacoma(tempPedidosoporte.PorcentajeRetencionFuente),
		Quitacoma(tempPedidosoporte.TotalRetencionFuente),
		Quitacoma(tempPedidosoporte.PorcentajeRetencionIca),
		Quitacoma(tempPedidosoporte.TotalRetencionIca),
		Quitacoma(tempPedidosoporte.Neto),
		tempPedidosoporte.Items,
		tempPedidosoporte.Formadepago,
		tempPedidosoporte.Mediodepago,
		tempPedidosoporte.Tercero,
		tempPedidosoporte.Almacenista,
		tempPedidosoporte.Centro,
		tempPedidosoporte.Tipo)

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

// INICIA PEDIDO SOPORTE EXISTE
func PedidosoporteExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM pedidosoporte  WHERE codigo=$1", Codigo)
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

// INICIA PEDIDO SOPORTE EDITAR
func PedidosoporteEditar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio pedidosoporte editar" + Codigo)
	db := dbConn()

	// traer PEDIDO SOPORTE
	v := pedidosoporte{}
	err := db.Get(&v, "SELECT * FROM pedidosoporte where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []pedidosoportedetalleeditar{}

	err2 := db.Select(&det, PedidosoporteConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
	v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
	v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
	v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
	v.CentroDetalle = TraerCentroConsulta(v.Centro)

	parametros := map[string]interface{}{
		"pedidosoporte": v,
		"detalle":       det,
		"hosting":       ruta,
		"retfte":        TraerParametrosInventario().Compracuentaporcentajeretfte,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/pedidosoporte/pedidosoporteEditar.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaBodega.html",
		"vista/autocompleta/autocompletaTercero.html",
		"vista/pedidosoporte/pedidosoporteScript.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error pedidosoporte nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA PEDIDO SOPORTE BORRAR
func PedidosoporteBorrar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio pedidosoporte editar" + Codigo)

	db := dbConn()

	// traer PEDIDO SOPORTE
	v := pedidosoporte{}
	err := db.Get(&v, "SELECT * FROM pedidosoporte where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []pedidosoportedetalleeditar{}
	err2 := db.Select(&det, PedidosoporteConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
	v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
	v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
	v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
	v.CentroDetalle = TraerCentroConsulta(v.Centro)

	parametros := map[string]interface{}{
		"pedidosoporte": v,
		"detalle":       det,
		"tercero":       TraerTerceroConsulta(v.Tercero),
		"hosting":       ruta,
		"retfte":        TraerParametrosInventario().Compracuentaporcentajeretfte,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/pedidosoporte/pedidosoporteBorrar.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaBodega.html",
		"vista/pedidosoporte/pedidosoporteScript.html")

	log.Println("Error pedidosoporte nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA PEDIDO SOPORTE ELIMINAR
func PedidosoporteEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	codigo := mux.Vars(r)["codigo"]

	// borrar PEDIDO SOPORTE
	delForm, err := db.Prepare("DELETE from pedidosoporte WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(codigo)

	// borar detalle
	delForm1, err := db.Prepare("DELETE from pedidosoportedetalle WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm1.Exec(codigo)

	log.Println("Registro Eliminado" + codigo)
	http.Redirect(w, r, "/PedidosoporteLista", 301)
}

// INICIA PEDIDO SOPORTE PDF
func PedidosoportePdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]

	// traer PEDIDO SOPORTE
	miPedidosoporte := pedidosoporte{}
	err := db.Get(&miPedidosoporte, "SELECT * FROM pedidosoporte where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}

	// traer detalle
	miDetalle := []pedidosoportedetalleeditar{}
	err2 := db.Select(&miDetalle, PedidosoporteConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	// traer tercero
	miTercero := tercero{}
	err3 := db.Get(&miTercero, "SELECT * FROM tercero where codigo=$1", miPedidosoporte.Tercero)
	if err3 != nil {
		log.Fatalln(err3)
	}

	// traer almacenista
	miAlmacenista := almacenista{}
	err4 := db.Get(&miAlmacenista, "SELECT * FROM almacenista where codigo=$1", miPedidosoporte.Almacenista)
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

		// PEDIDO NUMERO
		pdf.SetY(20)
		pdf.SetX(80)
		pdf.Ln(8)
		pdf.SetX(75)
		pdf.SetFont("Arial", "", 11)
		pdf.CellFormat(190, 10, "PEDIDO SOPORTE", "0", 0, "C", false, 0, "")
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
	PedidosoporteCabecera(pdf, miTercero, miPedidosoporte, miAlmacenista)

	var filas = len(miDetalle)
	// UNA PAGINA
	if filas <= 32 {
		for i, miFila := range miDetalle {
			var a = i + 1
			PedidosoporteFilaDetalle(pdf, miFila, a)
		}
		PedidosoportePieDePagina(pdf, miTercero, miPedidosoporte)
	} else {
		// mas de 32 y menos de 73
		if filas > 32 && filas <= 73 {
			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 41 {
					PedidosoporteFilaDetalle(pdf, miFila, a)
				}
			}
			PedidosoporteLinea(pdf)
			// segunda pagina
			pdf.AddPage()
			PedidosoporteCabecera(pdf, miTercero, miPedidosoporte, miAlmacenista)
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 41 {
					PedidosoporteFilaDetalle(pdf, miFila, a)
				}
			}

			PedidosoportePieDePagina(pdf, miTercero, miPedidosoporte)
		} else {
			// mas de 80

			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 41 {
					PedidosoporteFilaDetalle(pdf, miFila, a)
				}
			}
			PedidosoporteLinea(pdf)
			pdf.AddPage()
			PedidosoporteCabecera(pdf, miTercero, miPedidosoporte, miAlmacenista)
			// segunda pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 41 && a <= 82 {
					PedidosoporteFilaDetalle(pdf, miFila, a)
				}
			}
			PedidosoporteLinea(pdf)
			pdf.AddPage()
			PedidosoporteCabecera(pdf, miTercero, miPedidosoporte, miAlmacenista)
			// tercera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 82 {
					PedidosoporteFilaDetalle(pdf, miFila, a)
				}
			}
			PedidosoportePieDePagina(pdf, miTercero, miPedidosoporte)
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

func PedidosoporteCabecera(pdf *gofpdf.Fpdf, miTercero tercero, miPedidosoporte pedidosoporte, miAlmacenista almacenista) {
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
	pdf.CellFormat(40, 4, Titulo(TraerFormadepago(miPedidosoporte.Formadepago)), "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Fecha de Expedicion", "", 0,
		"L", false, 0, "")
	pdf.SetX(150)
	pdf.CellFormat(40, 4, miPedidosoporte.Fecha.Format("02/01/2006")+" "+Titulo(miPedidosoporte.Hora), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)

	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Medio de Pago", "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, TraerMediodepago(miPedidosoporte.Mediodepago), "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Fecha de Vencimiento", "", 0,
		"L", false, 0, "")
	pdf.SetX(150)
	pdf.CellFormat(40, 4, miPedidosoporte.Vence.Format("02/01/2006"), "", 0,
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
func PedidosoporteFilaDetalle(pdf *gofpdf.Fpdf, miFila pedidosoportedetalleeditar, a int) {
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

func PedidosoportePieDePagina(pdf *gofpdf.Fpdf, miTercero tercero, miPedidosoporte pedidosoporte) {

	Totalletras, err := IntLetra(Cadenaentero(miPedidosoporte.Neto))
	if err != nil {
		fmt.Println(err)
	}

	ene := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetFont("Arial", "", 8)
	pdf.SetY(229)
	pdf.SetX(20)
	pdf.MultiCell(184, 4, "SON: "+ene(Mayuscula(Totalletras))+" PESOS MDA. CTE.", "0", "L", false)
	pdf.SetFont("Arial", "", 9)
	pdf.SetY(254)
	pdf.Line(45, 254, 125, 254)
	pdf.Ln(1)
	pdf.SetX(65)
	pdf.CellFormat(40, 4, "FIRMA RESPONSABLE", "0", 0, "C", false, 0, "")

	// PRESENTA DATOS CON VALORES //
	pdf.SetFont("Arial", "", 9)
	var separador float64
	var altoseparador float64
	separador = 254
	altoseparador = -4

	// INICIA DATOS FACTURA
	if miPedidosoporte.Neto != "0" {
		pdf.SetFont("Arial", "", 9)
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "TOTAL", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miPedidosoporte.Neto), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miPedidosoporte.TotalRetencionIca != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "RET. ICA.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, miPedidosoporte.PorcentajeRetencionIca+"%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miPedidosoporte.TotalRetencionIca), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miPedidosoporte.TotalRetencionFuente != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "RET. FTE.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, miPedidosoporte.PorcentajeRetencionFuente+"%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miPedidosoporte.TotalRetencionFuente), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miPedidosoporte.Subtotalbase5 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "GRAVABLE", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, "5%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miPedidosoporte.Subtotalbase5), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miPedidosoporte.Subtotalbase19 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "GRAVABLE", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, "19%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miPedidosoporte.Subtotalbase19), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miPedidosoporte.Subtotalbase0 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "NO GRAVABLE", "0", 0, "L",
			false, 0, "")
		pdf.SetX(15)
		pdf.CellFormat(190, 4, Coma(miPedidosoporte.Subtotalbase0), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}
}

func PedidosoporteLinea(pdf *gofpdf.Fpdf) {
	pdf.SetY(259)
	pdf.Line(20, 259, 204, 259)
}

// INICIA PEDIDO SOPORTE TODOS PDF
func PedidosoporteTodosCabecera(pdf *gofpdf.Fpdf) {
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

func PedidosoporteTodosDetalle(pdf *gofpdf.Fpdf, miFila pedidosoporteLista,
	a int) {
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

func PedidosoporteTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string
	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre," +
		"centro.nombre as CentroNombre,pedidosoporte.neto," +
		"pedidosoporte.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	consulta += " FROM pedidosoporte "
	consulta += " inner join tercero on tercero.codigo=pedidosoporte.tercero "
	consulta += " inner join centro on centro.codigo=pedidosoporte.centro "
	consulta += " inner join almacenista on almacenista.codigo=pedidosoporte." +
		"almacenista "
	consulta += " ORDER BY cast(pedidosoporte.codigo as integer) ASC"

	t := []pedidosoporteLista{}
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
		pdf.CellFormat(190, 10, "DATOS PEDIDO SOPORTE", "0", 0,
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
		PedidosoporteTodosDetalle(pdf, miFila, a)
	}

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

// TERMINA PEDIDO TODOS PDF

// PEDIDO SOPORTE EXCEL
func PedidosoporteExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string
	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre," +
		"centro.nombre as CentroNombre,pedidosoporte.neto," +
		"pedidosoporte.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	consulta += " FROM pedidosoporte "
	consulta += " inner join tercero on tercero.codigo=pedidosoporte.tercero "
	consulta += " inner join centro on centro.codigo=pedidosoporte.centro "
	consulta += " inner join almacenista on almacenista.codigo=pedidosoporte." +
		"almacenista "
	consulta += " ORDER BY cast(pedidosoporte.codigo as integer) ASC"
	t := []pedidoLista{}
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

	// FUNCION PARA UNIR DOS CELDAS
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
	f.SetCellValue("Sheet1", "A9", "LISTADO DE PEDIDOS SOPORTE")
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

package main

// INICIA PEDIDO IMPORTAR PAQUETES
import (
	"bytes"
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

// INICIA PEDIDO ESTRUCTURA JSON
type pedidoJson struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Nombre string `json:"nombre"`
}

// INICIA PEDIDO ESTRUCTURA
type pedidoLista struct {
	Codigo            string
	Fecha             time.Time
	Neto              string
	Tercero           string
	TerceroNombre     string
	CentroNombre      string
	AlmacenistaNombre string
}

// INICIA PEDIDO ESTRUCTURA
type pedido struct {
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
	Detalle                   []pedidodetalle       `json:"Detalle"`
	DetalleEditar             []pedidodetalleeditar `json:"DetalleEditar"`
	TerceroDetalle            tercero
	Tipo                      string
	Centro                    string
	FormadepagoDetalle        formadepago
	MediodepagoDetalle        mediodepago
	AlmacenistaDetalle        almacenista
	CentroDetalle             centro
}

// INICIA PEDIDODETALLE ESTRUCTURA
type pedidodetalle struct {
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
type pedidodetalleeditar struct {
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
func PedidoConsultaDetalle() string {
	var consulta = ""
	consulta = "select "
	consulta += "pedidodetalle.Id as id ,"
	consulta += "pedidodetalle.Codigo as codigo,"
	consulta += "pedidodetalle.Fila as fila,"
	consulta += "pedidodetalle.Cantidad as cantidad,"
	consulta += "pedidodetalle.Precio as precio,"
	consulta += "pedidodetalle.Descuento as descuento,"
	consulta += "pedidodetalle.Montodescuento as montodescuento,"
	consulta += "pedidodetalle.Sigratis as sigratis,"
	consulta += "pedidodetalle.Subtotal as subtotal,"
	consulta += "pedidodetalle.Subtotaldescuento  as subtotaldescuento,"
	consulta += "pedidodetalle.Pagina as pagina ,"
	consulta += "pedidodetalle.Bodega as bodega,"
	consulta += "bodega.Nombre as bodeganombre,"
	consulta += "pedidodetalle.Producto as producto,"
	consulta += "pedidodetalle.Fecha as fecha,"
	consulta += "producto.nombre as ProductoNombre, "
	consulta += "producto.iva as ProductoIva, "
	consulta += "producto.unidad as ProductoUnidad "
	consulta += "from pedidodetalle "
	consulta += "inner join producto on producto.codigo=pedidodetalle.producto "
	consulta += "inner join bodega on bodega.codigo=pedidodetalle.bodega "
	consulta += " where pedidodetalle.codigo=$1 ORDER BY fila"
	log.Println(consulta)
	return consulta
}

// INICIA PEDIDO LISTA
func PedidoLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/pedido/pedidoLista.html")
	log.Println("Error pedido 0")
	var consulta string
	var miperiodo = periodoSesion(r)
	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre,centro.nombre as CentroNombre,pedido.neto,pedido.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	consulta += " FROM pedido "
	consulta += " inner join tercero on tercero.codigo=pedido.tercero "
	consulta += " inner join centro on centro.codigo=pedido.centro "
	consulta += " inner join almacenista on almacenista.codigo=pedido.almacenista "
	consulta += " where  extract(year from fecha)=" + miperiodo + " "

	consulta += " ORDER BY cast(pedido.codigo as integer) ASC"

	db := dbConn()
	res := []pedidoLista{}

	err := db.Select(&res, consulta)

	if err != nil {
		fmt.Println(err)
	}

	varmap := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
	}
	log.Println("Error pedido888")
	tmp.Execute(w, varmap)
}

// INICIA PEDIDO NUEVO
func PedidoNuevo(w http.ResponseWriter, r *http.Request) {

	// TRAE COPIA DE EDITAR
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio pedido editar" + Codigo)

	db := dbConn()
	v := pedido{}
	det := []pedidodetalleeditar{}
	if Codigo == "False" {

	} else {

		err := db.Get(&v, "SELECT * FROM pedido where codigo=$1", Codigo)
		if err != nil {
			log.Fatalln(err)
		}

		v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
		v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
		v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
		v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
		v.CentroDetalle = TraerCentroConsulta(v.Centro)
		err2 := db.Select(&det, PedidoConsultaDetalle(), Codigo)
		if err2 != nil {
			fmt.Println(err2)
		}
	}

	parametros := map[string]interface{}{
		"miperiodo":   periodoSesion(r),
		"almacenista": PrimerAlmacenista(),
		"codigo":      Codigo,
		"pedido":      v,
		"detalle":     det,
		"hosting":     ruta,
		"retfte":      TraerParametrosInventario().Compracuentaporcentajeretfte,
	}
	//TERMINA TRAE COPIA DE EDITAR

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/pedido/pedidoNuevo.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaBodega.html",
		"vista/autocompleta/autocompletaTercero.html",
		"vista/pedido/pedidoScript.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error pedido nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA PEDIDO INSERTAR AJAX
func PedidoAgregar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var tempPedido pedido

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// carga informacion de la PEDIDO
	err = json.Unmarshal(b, &tempPedido)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	parametrosinventario := configuracioninventario{}
	err = db.Get(&parametrosinventario, "SELECT * FROM configuracioninventario")
	if err != nil {
		panic(err.Error())
	}

	if tempPedido.Accion == "Actualizar" {

		// borra detalle anterior
		delForm, err := db.Prepare("DELETE from pedidodetalle WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(tempPedido.Codigo)

		// borra cabecera anterior
		delForm1, err := db.Prepare("DELETE from pedido WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm1.Exec(tempPedido.Codigo)
	}

	// INSERTA DETALLE
	for i, x := range tempPedido.Detalle {
		var a = i
		var q string
		q = "insert into pedidodetalle ("
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

		// TERMINA PEDIDO GRABAR INSERTAR
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

	// INICIA INSERTAR PEDIDO
	log.Println("Got %s age %s club %s\n", tempPedido.Codigo, tempPedido.Tercero, tempPedido.Neto)
	var q string
	q = "insert into pedido ("
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
	log.Println("Hora", tempPedido.Fecha.Format("02/01/2006"))
	currentTime := time.Now()

	_, err = insForm.Exec(
		tempPedido.Codigo,
		tempPedido.Fecha.Format(layout),
		tempPedido.Vence.Format(layout),
		currentTime.Format("3:4:5 PM"),
		Quitacoma(tempPedido.Descuento),
		Quitacoma(tempPedido.Subtotaldescuento19),
		Quitacoma(tempPedido.Subtotaldescuento5),
		Quitacoma(tempPedido.Subtotaldescuento0),
		Quitacoma(tempPedido.Subtotal),
		Quitacoma(tempPedido.Subtotal19),
		Quitacoma(tempPedido.Subtotal5),
		Quitacoma(tempPedido.Subtotal0),
		Quitacoma(tempPedido.Subtotaliva19),
		Quitacoma(tempPedido.Subtotaliva5),
		Quitacoma(tempPedido.Subtotaliva0),
		Quitacoma(tempPedido.Subtotalbase19),
		Quitacoma(tempPedido.Subtotalbase5),
		Quitacoma(tempPedido.Subtotalbase0),
		Quitacoma(tempPedido.TotalIva),
		Quitacoma(tempPedido.Total),
		Quitacoma(tempPedido.Neto),
		tempPedido.Items,
		Quitacoma(tempPedido.PorcentajeRetencionFuente),
		Quitacoma(tempPedido.TotalRetencionFuente),
		Quitacoma(tempPedido.PorcentajeRetencionIca),
		Quitacoma(tempPedido.TotalRetencionIca),
		tempPedido.Formadepago,
		tempPedido.Mediodepago,
		tempPedido.Tercero,
		tempPedido.Almacenista,
		tempPedido.Centro,
		tempPedido.Tipo)

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

// INICIA PEDIDO EXISTE
func PedidoExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM pedido  WHERE codigo=$1", Codigo)
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

// INICIA PEDIDO EDITAR
func PedidoEditar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio pedido editar" + Codigo)
	db := dbConn()

	// traer PEDIDO
	v := pedido{}
	err := db.Get(&v, "SELECT * FROM pedido where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []pedidodetalleeditar{}

	err2 := db.Select(&det, PedidoConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
	v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
	v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
	v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
	v.CentroDetalle = TraerCentroConsulta(v.Centro)

	parametros := map[string]interface{}{
		"miperiodo": periodoSesion(r),
		"pedido":    v,
		"detalle":   det,
		"hosting":   ruta,
		"retfte":    TraerParametrosInventario().Compracuentaporcentajeretfte,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/pedido/pedidoEditar.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaBodega.html",
		"vista/autocompleta/autocompletaTercero.html",
		"vista/pedido/pedidoScript.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error pedido nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA PEDIDO BORRAR
func PedidoBorrar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio pedido editar" + Codigo)

	db := dbConn()

	// traer PEDIDO
	v := pedido{}
	err := db.Get(&v, "SELECT * FROM pedido where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []pedidodetalleeditar{}
	err2 := db.Select(&det, PedidoConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	v.TerceroDetalle = TraerTerceroConsulta(v.Tercero)
	v.FormadepagoDetalle = TraerFormadepagoConsulta(v.Formadepago)
	v.MediodepagoDetalle = TraerMediodepagoConsulta(v.Mediodepago)
	v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
	v.CentroDetalle = TraerCentroConsulta(v.Centro)

	parametros := map[string]interface{}{
		"pedido":  v,
		"detalle": det,
		"tercero": TraerTerceroConsulta(v.Tercero),
		"hosting": ruta,
		"retfte":  TraerParametrosInventario().Compracuentaporcentajeretfte,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/pedido/pedidoBorrar.html",
		"vista/autocompleta/autocompletaFormadepago.html",
		"vista/autocompleta/autocompletaMediodepago.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaBodega.html",
		"vista/pedido/pedidoScript.html")

	log.Println("Error pedido nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA PEDIDO ELIMINAR
func PedidoEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	codigo := mux.Vars(r)["codigo"]

	// borrar PEDIDO
	delForm, err := db.Prepare("DELETE from pedido WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(codigo)

	// borar detalle
	delForm1, err := db.Prepare("DELETE from pedidodetalle WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm1.Exec(codigo)

	log.Println("Registro Eliminado" + codigo)
	http.Redirect(w, r, "/PedidoLista", 301)
}

// INICIA PEDIDO PDF
func PedidoPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]

	// TRAER PEDIDO
	miPedido := pedido{}
	err := db.Get(&miPedido, "SELECT * FROM pedido where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}

	// traer detalle
	miDetalle := []pedidodetalleeditar{}
	err2 := db.Select(&miDetalle, PedidoConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	// traer tercero
	miTercero := tercero{}
	err3 := db.Get(&miTercero, "SELECT * FROM tercero where codigo=$1", miPedido.Tercero)
	if err3 != nil {
		log.Fatalln(err3)
	}

	// traer almacenista
	miAlmacenista := almacenista{}
	err4 := db.Get(&miAlmacenista, "SELECT * FROM almacenista where codigo=$1", miPedido.Almacenista)
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
		pdf.CellFormat(190, 10, "PEDIDO", "0", 0, "C", false, 0, "")
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
	PedidoCabecera(pdf, miTercero, miPedido, miAlmacenista)

	var filas = len(miDetalle)
	// UNA PAGINA
	if filas <= 32 {
		for i, miFila := range miDetalle {
			var a = i + 1
			PedidoFilaDetalle(pdf, miFila, a)
		}
		PedidoPieDePagina(pdf, miTercero, miPedido)
	} else {
		// mas de 32 y menos de 73
		if filas > 32 && filas <= 73 {
			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 41 {
					PedidoFilaDetalle(pdf, miFila, a)
				}
			}
			PedidoLinea(pdf)
			// segunda pagina
			pdf.AddPage()
			PedidoCabecera(pdf, miTercero, miPedido, miAlmacenista)
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 41 {
					PedidoFilaDetalle(pdf, miFila, a)
				}
			}

			PedidoPieDePagina(pdf, miTercero, miPedido)
		} else {
			// mas de tres paginas

			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 41 {
					PedidoFilaDetalle(pdf, miFila, a)
				}
			}
			PedidoLinea(pdf)
			pdf.AddPage()
			PedidoCabecera(pdf, miTercero, miPedido, miAlmacenista)
			// segunda pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 41 && a <= 82 {
					PedidoFilaDetalle(pdf, miFila, a)
				}
			}
			PedidoLinea(pdf)
			pdf.AddPage()
			PedidoCabecera(pdf, miTercero, miPedido, miAlmacenista)
			// tercera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 82 {
					PedidoFilaDetalle(pdf, miFila, a)
				}
			}

			PedidoPieDePagina(pdf, miTercero, miPedido)
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

func PedidoCabecera(pdf *gofpdf.Fpdf, miTercero tercero, miPedido pedido, miAlmacenista almacenista) {
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
	pdf.CellFormat(40, 4, Titulo(TraerFormadepago(miPedido.Formadepago)), "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Fecha de Expedicion", "", 0,
		"L", false, 0, "")
	pdf.SetX(150)
	pdf.CellFormat(40, 4, miPedido.Fecha.Format("02/01/2006")+" "+Titulo(miPedido.Hora), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)

	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Medio de Pago", "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, ene(TraerMediodepago(miPedido.Mediodepago)), "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Fecha de Vencimiento", "", 0,
		"L", false, 0, "")
	pdf.SetX(150)
	pdf.CellFormat(40, 4, miPedido.Vence.Format("02/01/2006"), "", 0,
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
func PedidoFilaDetalle(pdf *gofpdf.Fpdf, miFila pedidodetalleeditar, a int) {
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

func PedidoPieDePagina(pdf *gofpdf.Fpdf, miTercero tercero, miPedido pedido) {

	Totalletras, err := IntLetra(Cadenaentero(miPedido.Neto))
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
	if miPedido.Neto != "0" {
		pdf.SetFont("Arial", "", 9)
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "TOTAL", "0", 0, "L",
			false, 0, "")
		pdf.SetX(13)
		pdf.CellFormat(190, 4, Coma(miPedido.Neto), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miPedido.TotalRetencionIca != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "RET. ICA.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, miPedido.PorcentajeRetencionIca+"%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(13)
		pdf.CellFormat(190, 4, Coma(miPedido.TotalRetencionIca), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miPedido.TotalRetencionFuente != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "RET. FTE.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, miPedido.PorcentajeRetencionFuente+"%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(13)
		pdf.CellFormat(190, 4, Coma(miPedido.TotalRetencionFuente), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miPedido.Subtotaliva5 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "I. V. A.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, "5%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(13)
		pdf.CellFormat(190, 4, Coma(miPedido.Subtotaliva5), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miPedido.Subtotaliva19 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "I. V. A.", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, "19%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(13)
		pdf.CellFormat(190, 4, Coma(miPedido.Subtotaliva19), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miPedido.Subtotalbase5 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "GRAVABLE", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, "5%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(13)
		pdf.CellFormat(190, 4, Coma(miPedido.Subtotalbase5), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miPedido.Subtotalbase19 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "GRAVABLE", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, "19%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(13)
		pdf.CellFormat(190, 4, Coma(miPedido.Subtotalbase19), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miPedido.Subtotalbase0 != "0" {

		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "NO GRAVABLE", "0", 0, "L",
			false, 0, "")
		pdf.SetX(13)
		pdf.CellFormat(190, 4, Coma(miPedido.Subtotalbase0), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}
}

func PedidoLinea(pdf *gofpdf.Fpdf) {
	pdf.SetY(259)
	pdf.Line(20, 259, 204, 259)
}

// INICIA PEDIDO TODOS PDF
func PedidoTodosCabecera(pdf *gofpdf.Fpdf) {
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

func PedidoTodosDetalle(pdf *gofpdf.Fpdf, miFila pedidoLista, a int) {
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

func PedidoTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string
	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre,centro.nombre as CentroNombre,pedido.neto,pedido.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	consulta += " FROM pedido "
	consulta += " inner join tercero on tercero.codigo=pedido.tercero "
	consulta += " inner join centro on centro.codigo=pedido.centro "
	consulta += " inner join almacenista on almacenista.codigo=pedido.almacenista "
	consulta += " ORDER BY cast(pedido.codigo as integer) ASC"

	t := []pedidoLista{}
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
		pdf.CellFormat(190, 10, "DATOS PEDIDO", "0", 0,
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
		PedidoTodosDetalle(pdf, miFila, a)
	}

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

// TERMINA PEDIDO TODOS PDF

// PEDIDO EXCEL
func PedidoExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string
	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre,centro.nombre as CentroNombre,pedido.neto,pedido.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	consulta += " FROM pedido "
	consulta += " inner join tercero on tercero.codigo=pedido.tercero "
	consulta += " inner join centro on centro.codigo=pedido.centro "
	consulta += " inner join almacenista on almacenista.codigo=pedido.almacenista "
	consulta += " ORDER BY cast(pedido.codigo as integer) ASC"
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
	f.SetCellValue("Sheet1", "A9", "LISTADO DE PEDIDOS")
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

package main

// INICIA INVENTARIO INICIAL IMPORTAR PAQUETES
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

// INICIA INVENTARIO INICIAL ESTRUCTURA JSON
type inventarioinicialJson struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Nombre string `json:"nombre"`
}

// INICIA INVENTARIO INICIAL ESTRUCTURA
type inventarioinicialLista struct {
	Codigo            string
	Fecha             time.Time
	Total             string
	Almacenista       string
	AlmacenistaNombre string
	Centronombre      string
}

// INICIA INVENTARIO INICIAL ESTRUCTURA
type inventarioinicial struct {
	Items              string
	Codigo             string
	Fecha              time.Time
	Almacenista        string
	Subtotalbase19     string
	Subtotalbase5      string
	Subtotalbase0      string
	Total              string
	Accion             string
	Detalle            []inventarioinicialdetalle       `json:"Detalle"`
	DetalleEditar      []inventarioinicialdetalleeditar `json:"DetalleEditar"`
	Tipo               string
	Centro             string
	AlmacenistaDetalle almacenista
	CentroDetalle      centro
}

// INICIA INVENTARIO INICIALDETALLE ESTRUCTURA
type inventarioinicialdetalle struct {
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
type inventarioinicialdetalleeditar struct {
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
func InventarioinicialConsultaDetalle() string {
	var consulta = ""
	consulta = "select "
	consulta += "inventarioinicialdetalle.Id as id ,"
	consulta += "inventarioinicialdetalle.Codigo as codigo,"
	consulta += "inventarioinicialdetalle.Fila as fila,"
	consulta += "inventarioinicialdetalle.Cantidad as cantidad,"
	consulta += "inventarioinicialdetalle.Precio as precio,"
	consulta += "inventarioinicialdetalle.Descuento as descuento,"
	consulta += "inventarioinicialdetalle.Montodescuento as montodescuento,"
	consulta += "inventarioinicialdetalle.Sigratis as sigratis,"
	consulta += "inventarioinicialdetalle.Subtotal as subtotal,"
	consulta += "inventarioinicialdetalle.Subtotaldescuento  as subtotaldescuento,"
	consulta += "inventarioinicialdetalle.Pagina as pagina ,"
	consulta += "inventarioinicialdetalle.Bodega as bodega,"
	consulta += "bodega.Nombre as bodeganombre,"
	consulta += "inventarioinicialdetalle.Producto as producto,"
	consulta += "inventarioinicialdetalle.Fecha as fecha,"
	consulta += "producto.nombre as ProductoNombre, "
	consulta += "producto.iva as ProductoIva, "
	consulta += "producto.unidad as ProductoUnidad "
	consulta += "from inventarioinicialdetalle "
	consulta += "inner join producto on producto.codigo=inventarioinicialdetalle.producto "
	consulta += "inner join bodega on bodega.codigo=inventarioinicialdetalle.bodega "
	consulta += " where inventarioinicialdetalle.codigo=$1 ORDER BY fila"
	log.Println(consulta)
	return consulta
}

// INICIA INVENTARIO INICIAL LISTA
func InventarioinicialLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/inventarioinicial/inventarioinicialLista.html")
	log.Println("Error inventarioinicial 0")
	var consulta string
	var miperiodo = periodoSesion(r)

	consulta = "  SELECT inventarioinicial.almacenista,inventarioinicial.total,inventarioinicial.codigo,fecha,almacenista.nombre as almacenistanombre"
	consulta += " FROM inventarioinicial "
	consulta += " inner join almacenista on almacenista.codigo=inventarioinicial.almacenista "
	consulta += " inner join centro on centro.codigo=inventarioinicial.centro "
	consulta += " where  extract(year from fecha)=" + miperiodo + " "
	consulta += " ORDER BY inventarioinicial.codigo ASC"

	db := dbConn()
	res := []inventarioinicialLista{}

	err := db.Select(&res, consulta)

	if err != nil {
		fmt.Println(err)
	}

	varmap := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
	}
	log.Println("Error inventarioinicial888")
	tmp.Execute(w, varmap)
}

// INICIA INVENTARIO INICIAL NUEVO
func InventarioinicialNuevo(w http.ResponseWriter, r *http.Request) {

	// TRAE COPIA DE EDITAR
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio inventarioinicial editar" + Codigo)

	db := dbConn()
	v := inventarioinicial{}
	det := []inventarioinicialdetalleeditar{}
	if Codigo == "False" {

	} else {

		err := db.Get(&v, "SELECT * FROM inventarioinicial where codigo=$1", Codigo)
		if err != nil {
			log.Fatalln(err)
		}

		v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
		v.CentroDetalle = TraerCentroConsulta(v.Centro)

		err2 := db.Select(&det, InventarioinicialConsultaDetalle(), Codigo)
		if err2 != nil {
			fmt.Println(err2)
		}
	}

	parametros := map[string]interface{}{
		"codigo":            Codigo,
		"inventarioinicial": v,
		"detalle":           det,
		"hosting":           ruta,
		"almacenista":       ListaAlmacenista(),
		"bodega":            ListaBodega(),
		"centro":            ListaCentro(),
		"miperiodo":         periodoSesion(r),
	}
	//TERMINA TRAE COPIA DE EDITAR

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/inventarioinicial/inventarioinicialNuevo.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaBodega.html",
		"vista/inventarioinicial/inventarioinicialScript.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error inventarioinicial nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA INVENTARIO INICIAL INSERTAR AJAX
func InventarioinicialAgregar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var tempInventarioinicial inventarioinicial

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// carga informacion de la Inventarioinicial
	err = json.Unmarshal(b, &tempInventarioinicial)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	parametrosinventario := configuracioninventario{}
	err = db.Get(&parametrosinventario, "SELECT * FROM configuracioninventario")
	if err != nil {
		panic(err.Error())
	}

	if tempInventarioinicial.Accion == "Actualizar" {

		// borra detalle anterior
		delForm, err := db.Prepare("DELETE from inventarioinicialdetalle WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(tempInventarioinicial.Codigo)

		// borra inventario
		Borrarinventario(tempInventarioinicial.Codigo, "Inventarioinicial")

		// borra cabecera anterior
		delForm1, err := db.Prepare("DELETE from inventarioinicial WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm1.Exec(tempInventarioinicial.Codigo)
	}

	// INSERTA DETALLE
	for i, x := range tempInventarioinicial.Detalle {
		var a = i
		var q string
		q = "insert into inventarioinicialdetalle ("
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

		// TERMINA INVENTARIO INICIAL GRABAR INSERTAR
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
	for _, x := range tempInventarioinicial.Detalle {
		miInventario = append(miInventario, inventario{
			x.Fecha, x.Tipo, x.Codigo, x.Bodega,
			x.Producto,
			Quitacoma(x.Cantidad),
			Quitacoma(x.Precio),
			operacionInventarioInicial})
	}
	// insertar inventario

	InsertaInventario(miInventario)

	// INICIA INSERTAR INVENTARIO INICIAL
	log.Println("Got %s age %s club %s\n", tempInventarioinicial.Codigo, tempInventarioinicial.Total)
	var q string
	q = "insert into inventarioinicial ("
	q += "Codigo,"
	q += "Fecha,"
	q += "Subtotalbase19,"
	q += "Subtotalbase5,"
	q += "Subtotalbase0,"
	q += "Total,"
	q += "Items,"
	q += "Almacenista,"
	q += "Centro"
	q += " ) values("
	q += parametros(9)
	q += " ) "

	log.Println("Cadena SQL " + q)
	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	layout := "2006-01-02"

	_, err = insForm.Exec(
		tempInventarioinicial.Codigo,
		tempInventarioinicial.Fecha.Format(layout),
		Quitacoma(tempInventarioinicial.Subtotalbase19),
		Quitacoma(tempInventarioinicial.Subtotalbase5),
		Quitacoma(tempInventarioinicial.Subtotalbase0),
		Quitacoma(tempInventarioinicial.Total),
		tempInventarioinicial.Items,
		tempInventarioinicial.Almacenista,
		tempInventarioinicial.Centro)

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

// INICIA INVENTARIO INICIAL EXISTE
func InventarioinicialExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM inventarioinicial  WHERE codigo=$1", Codigo)
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

// INICIA INVENTARIO INICIAL EDITAR
func InventarioinicialEditar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio inventarioinicial editar" + Codigo)
	db := dbConn()

	// traer INVENTARIO INICIAL
	v := inventarioinicial{}
	err := db.Get(&v, "SELECT * FROM inventarioinicial where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []inventarioinicialdetalleeditar{}

	err2 := db.Select(&det, InventarioinicialConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
	v.CentroDetalle = TraerCentroConsulta(v.Centro)

	parametros := map[string]interface{}{
		"inventarioinicial": v,
		"detalle":           det,
		"hosting":           ruta,
		"almacenista":       ListaAlmacenista(),
		"bodega":            ListaBodega(),
		"centro":            ListaCentro(),
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/inventarioinicial/inventarioinicialEditar.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaBodega.html",
		"vista/inventarioinicial/inventarioinicialScript.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error inventarioinicial nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA INVENTARIO INICIAL BORRAR
func InventarioinicialBorrar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio inventarioinicial editar" + Codigo)

	db := dbConn()

	// traer INVENTARIO INICIAL
	v := inventarioinicial{}
	err := db.Get(&v, "SELECT * FROM inventarioinicial where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []inventarioinicialdetalleeditar{}
	err2 := db.Select(&det, InventarioinicialConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)
	v.CentroDetalle = TraerCentroConsulta(v.Centro)

	parametros := map[string]interface{}{
		"inventarioinicial": v,
		"detalle":           det,
		"hosting":           ruta,
		"almacenista":       ListaAlmacenista(),
		"bodega":            ListaBodega(),
		"centro":            ListaCentro(),
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/inventarioinicial/inventarioinicialBorrar.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaBodega.html",
		"vista/inventarioinicial/inventarioinicialScript.html")

	log.Println("Error inventarioinicial nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA INVENTARIO INICIAL ELIMINAR
func InventarioinicialEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	codigo := mux.Vars(r)["codigo"]

	// borrar INVENTARIO INICIAL
	delForm, err := db.Prepare("DELETE from inventarioinicial WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(codigo)

	// borar detalle
	delForm1, err := db.Prepare("DELETE from inventarioinicialdetalle WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm1.Exec(codigo)

	// borar detalle invenario
	Borrarinventario(codigo, "Inventarioinicial")

	log.Println("Registro Eliminado" + codigo)
	http.Redirect(w, r, "/InventarioinicialLista", 301)
}

// INICIA INVENTARIO INICIAL PDF
func InventarioinicialPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]

	// traer PEDIDO FACTURA GASTO
	miInventarioinicial := inventarioinicial{}
	err := db.Get(&miInventarioinicial, "SELECT * FROM inventarioinicial where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}

	// traer detalle
	miDetalle := []inventarioinicialdetalleeditar{}
	err2 := db.Select(&miDetalle, InventarioinicialConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	// traer almacenista
	miAlmacenista := almacenista{}
	err4 := db.Get(&miAlmacenista, "SELECT * FROM almacenista where codigo=$1", miInventarioinicial.Almacenista)
	if err4 != nil {
		log.Fatalln(err4)
	}

	miCentro := centro{}
	err5 := db.Get(&miCentro, "SELECT * FROM centro where codigo=$1", miInventarioinicial.Centro)
	if err5 != nil {
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

		// FACTURA GASTO NUMERO
		pdf.SetY(20)
		pdf.SetX(80)
		pdf.Ln(8)
		pdf.SetX(75)
		pdf.SetFont("Arial", "", 11)
		pdf.CellFormat(190, 10, "INVENTARIO INICIAL", "0", 0, "C", false, 0, "")
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
	InventarioinicialCabecera(pdf, miCentro, miInventarioinicial, miAlmacenista)

	var filas = len(miDetalle)
	// UNA PAGINA
	if filas <= 32 {
		for i, miFila := range miDetalle {
			var a = i + 1
			InventarioinicialFilaDetalle(pdf, miFila, a)
		}
		InventarioinicialPieDePagina(pdf, miInventarioinicial)
	} else {
		// mas de 15 y menos de 19 dos paginas
		if filas > 32 && filas <= 73 {
			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 41 {
					InventarioinicialFilaDetalle(pdf, miFila, a)
				}
			}
			InventarioinicialLinea(pdf, miInventarioinicial)
			// segunda pagina
			pdf.AddPage()
			InventarioinicialCabecera(pdf, miCentro, miInventarioinicial, miAlmacenista)
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 41 {
					InventarioinicialFilaDetalle(pdf, miFila, a)
				}
			}

			InventarioinicialPieDePagina(pdf, miInventarioinicial)
		} else {
			// mas de tres paginas

			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 41 {
					InventarioinicialFilaDetalle(pdf, miFila, a)
				}
			}
			InventarioinicialLinea(pdf, miInventarioinicial)
			pdf.AddPage()
			InventarioinicialCabecera(pdf, miCentro, miInventarioinicial, miAlmacenista)
			// segunda pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 41 && a <= 82 {
					InventarioinicialFilaDetalle(pdf, miFila, a)
				}
			}
			InventarioinicialLinea(pdf, miInventarioinicial)
			pdf.AddPage()
			InventarioinicialCabecera(pdf, miCentro, miInventarioinicial, miAlmacenista)
			// tercera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 82 {
					InventarioinicialFilaDetalle(pdf, miFila, a)
				}
			}

			InventarioinicialPieDePagina(pdf, miInventarioinicial)
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

func InventarioinicialCabecera(pdf *gofpdf.Fpdf, miCentro centro, miInventarioinicial inventarioinicial, miAlmacenista almacenista) {
	ene := pdf.UnicodeTranslatorFromDescriptor("")
	// RELLENO TITULO
	pdf.SetFont("Arial", "", 10)
	pdf.SetY(44)
	pdf.SetX(20)
	pdf.SetFillColor(224, 231, 239)
	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(8)
	pdf.SetX(20)
	pdf.CellFormat(184, 6, "DATOS INVENTRIO INICIAL", "0", 0,
		"C", true, 0, "")
	pdf.Ln(8)
	pdf.SetX(20)
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(40, 4, "Documento No.", "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, Coma(miInventarioinicial.Codigo), "", 0,
		"L", false, 0, "")
	pdf.SetX(60)
	pdf.CellFormat(40, 4, "Fecha:", "", 0,
		"L", false, 0, "")
	pdf.SetX(75)
	pdf.CellFormat(40, 4, miInventarioinicial.Fecha.Format("02/01/2006"), "", 0,
		"L", false, 0, "")
	pdf.SetX(100)
	pdf.CellFormat(40, 4, "Almacenista:", "", 0,
		"L", false, 0, "")
	pdf.SetX(125)
	pdf.CellFormat(40, 4, ene(miAlmacenista.Nombre), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Centro Costo:", "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, miCentro.Nombre, "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)

	// RELLENO TITULO
	pdf.SetFont("Arial", "", 10)
	pdf.SetY(70)
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
	pdf.SetX(116)
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
func InventarioinicialFilaDetalle(pdf *gofpdf.Fpdf, miFila inventarioinicialdetalleeditar, a int) {
	ene := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetX(21)
	pdf.CellFormat(183, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(40, 4, miFila.Producto, "", 0,
		"L", false, 0, "")
	pdf.SetX(55)
	pdf.CellFormat(40, 4, miFila.ProductoNombre, "", 0,
		"L", false, 0, "")
	pdf.SetX(92)
	pdf.CellFormat(40, 4, Titulo(ene(miFila.ProductoUnidad)), "", 0,
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
	pdf.Ln(4)
}

func InventarioinicialPieDePagina(pdf *gofpdf.Fpdf, miInventarioinicial inventarioinicial) {

	Totalletras, err := IntLetra(Cadenaentero(miInventarioinicial.Total))
	if err != nil {
		fmt.Println(err)
	}

	ene := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetFont("Arial", "", 8)
	pdf.SetY(237)
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
	if miInventarioinicial.Total != "0" {
		pdf.SetFont("Arial", "", 9)
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "TOTAL", "0", 0, "L",
			false, 0, "")
		pdf.SetX(13)
		pdf.CellFormat(190, 4, Coma(miInventarioinicial.Total), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miInventarioinicial.Subtotalbase5 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "GRAVABLE", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, "5%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(13)
		pdf.CellFormat(190, 4, Coma(miInventarioinicial.Subtotalbase5), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miInventarioinicial.Subtotalbase19 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "GRAVABLE", "0", 0, "L",
			false, 0, "")
		pdf.SetX(168)
		pdf.CellFormat(190, 4, "19%", "0", 0, "L",
			false, 0, "")
		pdf.SetX(13)
		pdf.CellFormat(190, 4, Coma(miInventarioinicial.Subtotalbase19), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}

	if miInventarioinicial.Subtotalbase0 != "0" {
		pdf.SetY(separador)
		pdf.SetX(145)
		pdf.CellFormat(190, 4, "NO GRAVABLE", "0", 0, "L",
			false, 0, "")
		pdf.SetX(13)
		pdf.CellFormat(190, 4, Coma(miInventarioinicial.Subtotalbase0), "0", 0, "R",
			false, 0, "")
		separador += altoseparador
	}
}

func InventarioinicialLinea(pdf *gofpdf.Fpdf, miInventarioinicial inventarioinicial) {
	pdf.SetY(259)
	pdf.Line(20, 259, 204, 259)
}

// INICIA PEDIDO FACTURA GASTO TODOS PDF
func InventarioinicialTodosCabecera(pdf *gofpdf.Fpdf) {
	pdf.SetFont("Arial", "", 10)
	// RELLENO TITULO
	pdf.SetY(50)
	pdf.SetFillColor(224, 231, 239)
	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(7)
	pdf.SetX(20)
	pdf.CellFormat(184, 6, "No", "0", 0,
		"L", true, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(190, 6, "Documento", "0", 0,
		"L", false, 0, "")
	pdf.SetX(60)
	pdf.CellFormat(190, 6, "Nombre Almacenista", "0", 0,
		"L", false, 0, "")
	pdf.SetX(120)
	pdf.CellFormat(190, 6, "Centro", "0", 0,
		"L", false, 0, "")
	pdf.SetX(160)
	pdf.CellFormat(190, 6, "Fecha", "0", 0,
		"L", false, 0, "")
	pdf.SetX(193)
	pdf.CellFormat(190, 6, "Total", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}

func InventarioinicialTodosDetalle(pdf *gofpdf.Fpdf, miFila inventarioinicialLista, a int) {
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(21)
	pdf.CellFormat(181, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(10, 4, miFila.Codigo, "", 0,
		"L", false, 0, "")
	pdf.SetX(60)
	pdf.CellFormat(40, 4, miFila.AlmacenistaNombre, "", 0,
		"L", false, 0, "")
	pdf.SetX(120)
	pdf.CellFormat(40, 4, miFila.Centronombre, "", 0,
		"L", false, 0, "")
	pdf.SetX(160)
	pdf.CellFormat(40, 4, miFila.Fecha.Format("02/01/2006"), "", 0,
		"L", false, 0, "")
	pdf.SetX(164)
	pdf.CellFormat(40, 4, Coma(miFila.Total), "", 0,
		"R", false, 0, "")
	pdf.Ln(4)
}

func InventarioinicialTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string
	consulta = "  SELECT Inventarioinicial.almacenista, almacenista.nombre as AlmacenistaNombre,centro.nombre as CentroNombre,inventarioinicial.total,inventarioinicial.codigo,fecha "
	consulta += " FROM inventarioinicial "
	consulta += " inner join centro on centro.codigo=inventarioinicial.centro "
	consulta += " inner join almacenista on almacenista.codigo=inventarioinicial.almacenista "
	consulta += " ORDER BY cast(inventarioinicial.codigo as integer) ASC"

	t := []inventarioinicialLista{}
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
		pdf.CellFormat(190, 10, "DATOS INVENTARIO INICIAL", "0", 0,
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

	InventarioinicialTodosCabecera(pdf)
	// tercera pagina

	for i, miFila := range t {
		var a = i + 1
		if math.Mod(float64(a), 49) == 0 {
			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			InventarioinicialTodosCabecera(pdf)
		}
		InventarioinicialTodosDetalle(pdf, miFila, a)
	}

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

// TERMINA PEDIDO FACTURA GASTO TODOS PDF

// PEDIDO FACTURA GASTO  EXCEL
func InventarioinicialExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string
	consulta = "  SELECT inventarioinicial.almacenista,almacenista.nombre as AlmacenistaNombre,centro.nombre as CentroNombre,inventarioinicial.Total,inventarioinicial.codigo,fecha "
	consulta += " FROM inventarioinicial "
	consulta += " inner join centro on centro.codigo=inventarioinicial.centro "
	consulta += " inner join almacenista on almacenista.codigo=inventarioinicial.almacenista "
	consulta += " ORDER BY cast(inventarioinicial.codigo as integer) ASC"
	t := []inventarioinicialLista{}
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

	if err = f.SetColWidth("Sheet1", "F", "F", 15); err != nil {
		fmt.Println(err)
		return
	}

	// FUNCION PARA UNIR DOS CELDAS
	if err = f.MergeCell("Sheet1", "A1", "F1"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A2", "F2"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A3", "F3"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A4", "F4"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A5", "F5"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A6", "F6"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A7", "F7"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A8", "F8"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A9", "F9"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A10", "F10"); err != nil {
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
	f.SetCellValue("Sheet1", "A9", "LISTADO DE INVENTARIO INICIAL")
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
	f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel), "Documento")
	f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel), "Almacenista")
	f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel), "Nombre")
	f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel), "Centro")
	f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel), "Fecha")
	f.SetCellValue("Sheet1", "F"+strconv.Itoa(filaExcel), "Total")

	f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel), "A"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel), "B"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel), "C"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel), "D"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "E"+strconv.Itoa(filaExcel), "E"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "F"+strconv.Itoa(filaExcel), "F"+strconv.Itoa(filaExcel), estiloCabecera)
	filaExcel++

	for i, miFila := range t {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel+i), Entero(miFila.Codigo))
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel+i), Entero(miFila.Almacenista))
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel+i), miFila.AlmacenistaNombre)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel+i), miFila.Centronombre)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel+i), miFila.Fecha.Format("02/01/2006"))
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(filaExcel+i), Flotante(miFila.Total))

		f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel+i), "A"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel+i), "B"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel+i), "C"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel+i), "D"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "E"+strconv.Itoa(filaExcel+i), "E"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "F"+strconv.Itoa(filaExcel+i), "F"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
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

package main

// INICIA TRASLADO IMPORTAR PAQUETES
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

// INICIA TRASLADO ESTRUCTURA JSON
type trasladoJson struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Nombre string `json:"nombre"`
}

// INICIA TRASLADO ESTRUCTURA
type trasladoLista struct {
	Codigo            string
	Fecha             time.Time
	AlmacenistaNombre string
}

// INICIA TRASLADO ESTRUCTURA
type traslado struct {
	Codigo             string
	Fecha              time.Time
	Items              string
	Almacenista        string
	Accion             string
	Detalle            []trasladodetalle       `json:"Detalle"`
	DetalleEditar      []trasladodetalleeditar `json:"DetalleEditar"`
	Tipo               string
	AlmacenistaDetalle almacenista
}

// INICIA TRASLADODETALLE ESTRUCTURA
type trasladodetalle struct {
	Id       string
	Codigo   string
	Fila     string
	Bodega   string
	Producto string
	Tipo     string
	Entra    string
	Sale     string
	Fecha    time.Time
}

// INICIA TRASLADO DETALLE EDITAR
type trasladodetalleeditar struct {
	Id             string
	Codigo         string
	Fila           string
	Bodega         string
	Producto       string
	Fecha          time.Time
	BodegaNombre   string
	ProductoNombre string
	ProductoIva    string
	ProductoUnidad string
	Tipo           string
	Entra          string
	Sale           string
}

// INICIA TRASLADO CONSULTA DETALLE
func TrasladoConsultaDetalle() string {
	var consulta = ""
	consulta = "select "
	consulta += "trasladodetalle.Id as id ,"
	consulta += "trasladodetalle.Codigo as codigo,"
	consulta += "trasladodetalle.Entra as entra,"
	consulta += "trasladodetalle.Sale as sale,"
	consulta += "trasladodetalle.Fila as fila,"
	consulta += "trasladodetalle.Bodega as bodega,"
	consulta += "trasladodetalle.Producto as producto,"
	consulta += "trasladodetalle.Fecha as fecha,"
	consulta += "bodega.nombre as BodegaNombre, "
	consulta += "producto.nombre as ProductoNombre, "
	consulta += "producto.iva as ProductoIva, "
	consulta += "producto.unidad as ProductoUnidad "
	consulta += "from trasladodetalle "
	consulta += "inner join producto on producto.codigo=trasladodetalle.producto "
	consulta += "inner join bodega on bodega.codigo=trasladodetalle.bodega "
	consulta += " where trasladodetalle.codigo=$1"
	log.Println(consulta)
	return consulta
}

// INICIA TRASLADO LISTA
func TrasladoLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/traslado/trasladoLista.html")
	var miperiodo = periodoSesion(r)
	log.Println("Error traslado 0")
	var consulta string

	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre,traslado.codigo,fecha"
	consulta += " FROM traslado "
	consulta += " inner join almacenista on almacenista.codigo=traslado.almacenista "
	consulta += " where  extract(year from fecha)=" + miperiodo + " "
	consulta += " ORDER BY traslado.codigo ASC"

	db := dbConn()
	res := []trasladoLista{}

	err := db.Select(&res, consulta)

	if err != nil {
		fmt.Println(err)
		return
	}

	varmap := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
	}
	log.Println("Error traslado888")
	tmp.Execute(w, varmap)
}

// INICIA TRASLADO NUEVO
func TrasladoNuevo(w http.ResponseWriter, r *http.Request) {

	// TRAE COPIA DE EDITAR
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio traslado editar" + Codigo)

	db := dbConn()
	v := traslado{}
	det := []trasladodetalleeditar{}
	if Codigo == "False" {

	} else {

		err := db.Get(&v, "SELECT * FROM traslado where codigo=$1", Codigo)
		if err != nil {
			log.Fatalln(err)
		}

		v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)

		err2 := db.Select(&det, TrasladoConsultaDetalle(), Codigo)
		if err2 != nil {
			fmt.Println(err2)
		}
	}

	parametros := map[string]interface{}{
		"codigo":      Codigo,
		"traslado":    v,
		"detalle":     det,
		"hosting":     ruta,
		"almacenista": ListaAlmacenista(),
		"bodega":      ListaBodega(),
		"miperiodo":   periodoSesion(r),
	}
	//TERMINA TRAE COPIA DE EDITAR

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/traslado/trasladoNuevo.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaBodega.html",
		"vista/traslado/trasladoScript.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error traslado nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA TRASLADO INSERTAR AJAX
func TrasladoAgregar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var tempTraslado traslado

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// carga informacion de la TRASLADO
	err = json.Unmarshal(b, &tempTraslado)
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
	//if tempTraslado.Accion == "Nuevo" {
	//	log.Println("Resolucion " + tempTraslado.Resoluciontraslado)
	//	Codigoactual=Numerotraslado(tempTraslado.Resoluciontraslado)
	//	tempTraslado.Codigo=Codigoactual
	//}else{
	Codigoactual = tempTraslado.Codigo
	//}

	if tempTraslado.Accion == "Actualizar" {

		// borra detalle anterior
		delForm, err := db.Prepare("DELETE from trasladodetalle WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(tempTraslado.Codigo)

		// borra detalle inventario

		Borrarinventario(tempTraslado.Codigo, "Traslado")

		// borra cabecera anterior
		delForm1, err := db.Prepare("DELETE from traslado WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm1.Exec(tempTraslado.Codigo)
	}

	// INSERTA DETALLE
	for i, x := range tempTraslado.Detalle {
		var a = i
		var q string
		q = "insert into trasladodetalle ("
		q += "Id,"
		q += "Codigo,"
		q += "Fila,"
		q += "Bodega,"
		q += "Producto,"
		q += "Tipo,"
		q += "Entra,"
		q += "Sale,"
		q += "Fecha"
		q += " ) values("
		q += parametros(9)
		q += " ) "

		log.Println("Cadena SQL " + q)
		insForm, err := db.Prepare(q)
		if err != nil {
			panic(err.Error())
		}

		// TERMINA TRASLADO GRABAR INSERTAR
		_, err = insForm.Exec(
			x.Id,
			Codigoactual,
			x.Fila,
			x.Bodega,
			x.Producto,
			x.Tipo,
			Quitacoma(x.Entra),
			Quitacoma(x.Sale),
			x.Fecha)
		if err != nil {
			panic(err)
		}

		log.Println("Insertar Producto \n", x.Producto, a)
	}

	// INSERTA DETALLE INVENTARIO

	// inserta inventario
	miInventario := []inventario{}

	// INSERTA DETALLE INVENTARIO
	for _, x := range tempTraslado.Detalle {
		var tipooperacion string
		var cantidadtraslado string

		if x.Sale == "" {
			cantidadtraslado = x.Entra
			tipooperacion = operacionTrasladoEntrada
		} else {
			cantidadtraslado = x.Sale
			tipooperacion = operacionTrasladoSalida

		}
		miInventario = append(miInventario, inventario{
			x.Fecha, x.Tipo, x.Codigo, x.Bodega,
			x.Producto,
			Quitacoma(cantidadtraslado),
			"0",
			tipooperacion})
	}
	// insertar inventario

	InsertaInventario(miInventario)

	//for i, x := range tempTraslado.Detalle {
	//	var a = i
	//	var q string
	//	q = "insert into inventario ("
	//	q += "Fecha,"
	//	q += "Tipo,"
	//	q += "Codigo,"
	//	q += "Bodega,"
	//	q += "Producto,"
	//	q += "Cantidad,"
	//	q += "Precio,"
	//	q += "Operacion"
	//	q += " ) values("
	//	q += parametros(8)
	//	q += " ) "
	//
	//	log.Println("Cadena SQL " + q)
	//	insForm, err := db.Prepare(q)
	//	if err != nil {
	//		panic(err.Error())
	//	}
	//
	//	var tipooperacion string
	//	var cantidadtraslado string
	//
	//	if x.Sale == "" {
	//		cantidadtraslado = x.Entra
	//
	//		tipooperacion = operacionTrasladoEntrada
	//	} else {
	//		cantidadtraslado = x.Sale
	//		tipooperacion = operacionTrasladoSalida
	//
	//	}
	//
	//
	//	// TERMINA TRASLADO GRABAR INSERTAR
	//	_, err = insForm.Exec(
	//		x.Fecha,
	//		x.Tipo,
	//		Codigoactual,
	//		x.Bodega,
	//		x.Producto,
	//		cantidadtraslado,
	//		"0",
	//		tipooperacion)
	//	if err != nil {
	//		panic(err)
	//	}
	//	log.Println("Insertar Producto \n", x.Producto, a)
	//}

	// INICIA INSERTAR TRASLADOS
	log.Println("Got %s age %s club %s\n", tempTraslado.Codigo)
	var q string
	q = "insert into traslado ("
	q += "Codigo,"
	q += "Fecha,"
	q += "Items,"
	q += "Almacenista,"
	q += "Tipo"
	q += " ) values("
	q += parametros(5)
	q += " ) "

	log.Println("Cadena SQL " + q)
	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	layout := "2006-01-02"
	log.Println("Hora", tempTraslado.Fecha.Format("02/01/2006"))

	_, err = insForm.Exec(
		tempTraslado.Codigo,
		tempTraslado.Fecha.Format(layout),
		tempTraslado.Items,
		tempTraslado.Almacenista,
		tempTraslado.Tipo)

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

// INICIA TRASLADO EXISTE
func TrasladoExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM traslado  WHERE codigo=$1", Codigo)
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

// INICIA TRASLADO EDITAR
func TrasladoEditar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio traslado editar" + Codigo)
	db := dbConn()

	// traer TRASLADO
	v := traslado{}
	err := db.Get(&v, "SELECT * FROM traslado where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []trasladodetalleeditar{}

	err2 := db.Select(&det, TrasladoConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)

	parametros := map[string]interface{}{
		"traslado":    v,
		"detalle":     det,
		"hosting":     ruta,
		"almacenista": ListaAlmacenista(),
		"bodega":      ListaBodega(),
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/traslado/trasladoEditar.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaBodega.html",
		"vista/traslado/trasladoScript.html")
	fmt.Printf("%v, %v", err)
	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error traslado nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA TRASLADO BORRAR
func TrasladoBorrar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio traslado editar" + Codigo)

	db := dbConn()

	// traer TRASLADO
	v := traslado{}
	err := db.Get(&v, "SELECT * FROM traslado where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	// traer detalle

	det := []trasladodetalleeditar{}
	err2 := db.Select(&det, TrasladoConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	v.AlmacenistaDetalle = TraerAlmacenistaConsulta(v.Almacenista)

	parametros := map[string]interface{}{
		"traslado":    v,
		"detalle":     det,
		"hosting":     ruta,
		"almacenista": ListaAlmacenista(),
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/traslado/trasladoBorrar.html",
		"vista/autocompleta/autocompletaAlmacenista.html",
		"vista/autocompleta/autocompletaBodega.html",
		"vista/traslado/trasladoScript.html")

	log.Println("Error traslado nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA TRASLADO ELIMINAR
func TrasladoEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	codigo := mux.Vars(r)["codigo"]

	// borrar TRASLADO
	delForm, err := db.Prepare("DELETE from traslado WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(codigo)

	// borar detalle
	delForm1, err := db.Prepare("DELETE from trasladodetalle WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm1.Exec(codigo)

	// borar detalle invenario
	Borrarinventario(codigo, "Traslado")

	log.Println("Registro Eliminado" + codigo)
	http.Redirect(w, r, "/TrasladoLista", 301)
}

// INICIA TRASLADO PDF
func TrasladoPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]

	// traer TRASLADO
	miTraslado := traslado{}
	err := db.Get(&miTraslado, "SELECT * FROM traslado where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}

	// traer detalle
	miDetalle := []trasladodetalleeditar{}
	err2 := db.Select(&miDetalle, TrasladoConsultaDetalle(), Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	// traer tercero

	// traer almacenista
	miAlmacenista := almacenista{}
	err4 := db.Get(&miAlmacenista, "SELECT * FROM almacenista where codigo=$1", miTraslado.Almacenista)
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
		pdf.CellFormat(190, 10, e.Telefono1+" - "+e.Telefono2, "0", 0, "C", false, 0,
			"")
		pdf.Ln(4)
		pdf.CellFormat(190, 10, ene(c.NombreCiudad+" - "+c.NombreDepartamento), "0", 0, "C", false, 0,
			"")
		pdf.Ln(10)

	})

	pdf.SetFooterFunc(func() {
		pdf.SetTextColor(0, 0, 0)
		pdf.SetY(252)
		pdf.SetFont("Arial", "", 9)
		pdf.SetX(20)

		pdf.Ln(6)
		pdf.SetX(20)
		pdf.CellFormat(90, 10, "Sadconf.com", "", 0, "L", false, 0, "")
		pdf.SetX(129)
		pdf.CellFormat(80, 10, fmt.Sprintf(" %d de {nb}", pdf.PageNo()), "",
			0, "R", false, 0, "")
	})

	// inicia suma de paginas
	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.Ln(1)
	TrasladoCabecera(pdf, miTraslado, miAlmacenista)

	var filas = len(miDetalle)
	// menos de 32
	if filas <= 32 {
		for i, miFila := range miDetalle {
			var a = i + 1
			TrasladoFilaDetalle(pdf, miFila, a)
		}
		TrasladoPieDePagina(pdf, miTraslado)
	} else {
		// mas de 32 y menos de 73
		if filas > 32 && filas <= 73 {
			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 41 {
					TrasladoFilaDetalle(pdf, miFila, a)
				}
			}
			// segunda pagina
			pdf.AddPage()
			TrasladoCabecera(pdf, miTraslado, miAlmacenista)
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 41 {
					TrasladoFilaDetalle(pdf, miFila, a)
				}
			}

			TrasladoPieDePagina(pdf, miTraslado)
		} else {
			// mas de 80

			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 41 {
					TrasladoFilaDetalle(pdf, miFila, a)
				}
			}

			pdf.AddPage()
			TrasladoCabecera(pdf, miTraslado, miAlmacenista)
			// segunda pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 41 && a <= 82 {
					TrasladoFilaDetalle(pdf, miFila, a)
				}
			}

			pdf.AddPage()
			TrasladoCabecera(pdf, miTraslado, miAlmacenista)
			// tercera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 82 {
					TrasladoFilaDetalle(pdf, miFila, a)
				}
			}

			TrasladoPieDePagina(pdf, miTraslado)
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
func TrasladoCabecera(pdf *gofpdf.Fpdf, miTraslado traslado, miAlmacenista almacenista) {
	pdf.SetFont("Arial", "", 10)
	// RELLENO TITULO
	pdf.SetY(50)
	pdf.SetFillColor(224, 231, 239)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetX(20)
	pdf.CellFormat(184, 6, "TRASLADOS Y AJUSTES", "0", 0,
		"C", true, 0, "")

	//pdf.SetFont("Arial", "", 9)
	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(8)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Traslado No. "+miTraslado.Codigo, "0", 0, "L",
		false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, "Fecha de Expedicion", "", 0,
		"L", false, 0, "")
	pdf.SetX(90)
	pdf.CellFormat(40, 4, miTraslado.Fecha.Format("02/01/2006"), "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Almacenista", "", 0,
		"L", false, 0, "")
	pdf.SetX(135)
	pdf.CellFormat(40, 4, miAlmacenista.Nombre, "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)

	pdf.SetFont("Arial", "", 10)
	pdf.SetY(64)

	pdf.SetFillColor(224, 231, 239)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetX(20)

	pdf.CellFormat(184, 6, "ITEM", "0", 0,
		"L", true, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(40, 6, "CODIGO", "0", 0,
		"L", false, 0, "")
	pdf.SetX(55)
	pdf.CellFormat(40, 6, "PRODUCTO", "0", 0,
		"L", false, 0, "")
	pdf.SetX(105)
	pdf.CellFormat(40, 6, "UNIDAD", "0", 0,
		"L", false, 0, "")
	pdf.SetX(95)
	pdf.CellFormat(40, 6, "IVA", "0", 0,
		"R", false, 0, "")
	pdf.SetX(120)
	pdf.CellFormat(40, 6, "BODEGA", "0", 0,
		"R", false, 0, "")
	pdf.SetX(144)
	pdf.CellFormat(40, 6, "ENTRADA", "0", 0,
		"R", false, 0, "")
	pdf.SetX(164)
	pdf.CellFormat(40, 6, "SALIDA", "0", 0,
		"R", false, 0, "")
	pdf.Ln(8)
}
func TrasladoFilaDetalle(pdf *gofpdf.Fpdf, miFila trasladodetalleeditar, a int) {
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(21)
	pdf.CellFormat(183, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(40, 4, Subcadena(miFila.Producto, 0, 12), "", 0,
		"L", false, 0, "")
	pdf.SetX(55)
	pdf.CellFormat(40, 4, Subcadena(miFila.ProductoNombre, 0, 32), "", 0,
		"L", false, 0, "")
	pdf.SetX(105)
	pdf.CellFormat(40, 4, miFila.ProductoUnidad, "", 0,
		"L", false, 0, "")
	pdf.SetX(93)
	pdf.CellFormat(40, 4, miFila.ProductoIva, "", 0,
		"R", false, 0, "")
	pdf.SetX(120)
	pdf.CellFormat(40, 4, miFila.Bodega, "", 0,
		"R", false, 0, "")
	pdf.SetX(144)
	pdf.CellFormat(40, 4, miFila.Entra, "", 0,
		"R", false, 0, "")
	pdf.SetX(164)
	pdf.CellFormat(40, 4, miFila.Sale, "", 0,
		"R", false, 0, "")
	pdf.Ln(4)
}

func TrasladoPieDePagina(pdf *gofpdf.Fpdf, miTraslado traslado) {
	pdf.SetFont("Arial", "", 9)
	pdf.SetY(232)

	// LINEA
	pdf.Line(85, 250, 140, 250)

	pdf.Ln(16)
	pdf.SetX(95)
	pdf.CellFormat(40, 10, "FIRMA RESPONSABLE ", "0", 0, "C",
		false, 0, "")

}

// INICIA TRASLADO TODOS PDF
func TrasladoTodosCabecera(pdf *gofpdf.Fpdf) {
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
	pdf.CellFormat(190, 6, "Traslado No.", "0", 0,
		"L", false, 0, "")
	pdf.SetX(55)
	pdf.CellFormat(190, 6, "Fecha", "0", 0,
		"L", false, 0, "")
	pdf.SetX(90)
	pdf.CellFormat(190, 6, "Almacenista", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}

func TrasladoTodosDetalle(pdf *gofpdf.Fpdf, miFila trasladoLista, a int) {
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(21)
	pdf.CellFormat(181, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(40, 4, Subcadena(miFila.Codigo, 0, 12), "", 0,
		"L", false, 0, "")
	pdf.SetX(55)
	pdf.CellFormat(40, 4, miFila.Fecha.Format("02/01/2006"), "", 0, "L", false, 0, "")
	pdf.SetX(90)
	pdf.CellFormat(40, 4, miFila.AlmacenistaNombre, "", 0,
		"L", false, 0, "")
	pdf.Ln(4)
}

func TrasladoTodosPdf(w http.ResponseWriter, r *http.Request) {
	var consulta string

	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre,traslado.codigo,fecha"
	consulta += " FROM traslado "
	consulta += " inner join almacenista on almacenista.codigo=traslado.almacenista "
	consulta += " ORDER BY traslado.codigo ASC"

	db := dbConn()
	t := []trasladoLista{}
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	err := db.Select(&t, consulta)
	if err != nil {
		log.Fatalln(err)
	}
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
		pdf.CellFormat(190, 10, "DATOS TRASLADO", "0", 0,
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

	TrasladoTodosCabecera(pdf)
	// tercera pagina

	for i, miFila := range t {
		var a = i + 1
		if math.Mod(float64(a), 49) == 0 {
			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			TrasladoTodosCabecera(pdf)
		}
		TrasladoTodosDetalle(pdf, miFila, a)
	}

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

// TERMINA TRASLADO TODOS PDF

// DOCUMENTO EXCEL
func TrasladoExcel(w http.ResponseWriter, r *http.Request) {
	var consulta string

	consulta = "  SELECT almacenista.nombre as AlmacenistaNombre,traslado.codigo,fecha"
	consulta += " FROM traslado "
	consulta += " inner join almacenista on almacenista.codigo=traslado.almacenista "
	consulta += " ORDER BY traslado.codigo ASC"

	db := dbConn()
	t := []trasladoLista{}
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	err := db.Select(&t, consulta)
	if err != nil {
		log.Fatalln(err)
	}

	f := excelize.NewFile()

	// FUNCION ANCHO DE LA COLUMNA
	if err = f.SetColWidth("Sheet1", "A", "A", 15); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.SetColWidth("Sheet1", "B", "B", 15); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.SetColWidth("Sheet1", "C", "C", 50); err != nil {
		fmt.Println(err)
		return
	}

	// FUNCION PARA UNIR DOS CELDAS
	if err = f.MergeCell("Sheet1", "A1", "C1"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A2", "C2"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A3", "C3"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A4", "C4"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A5", "C5"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A6", "C6"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A7", "C7"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A8", "C8"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A9", "C9"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A10", "C10"); err != nil {
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
	f.SetCellValue("Sheet1", "A9", "LISTADO DE TRASLADOS")
	f.SetCellValue("Sheet1", "A10", "")

	f.SetCellStyle("Sheet1", "A1", "A1", estiloTitulo)
	f.SetCellStyle("Sheet1", "A2", "A2", estiloTitulo)
	f.SetCellStyle("Sheet1", "A3", "A3", estiloTitulo)
	f.SetCellStyle("Sheet1", "A4", "A4", estiloTitulo)
	f.SetCellStyle("Sheet1", "A5", "A5", estiloTitulo)
	f.SetCellStyle("Sheet1", "A6", "A6", estiloTitulo)
	f.SetCellStyle("Sheet1", "A7", "A7", estiloTitulo)
	f.SetCellStyle("Sheet1", "A8", "A8", estiloTitulo)
	f.SetCellStyle("Sheet1", "A9", "A7", estiloTitulo)
	f.SetCellStyle("Sheet1", "A10", "A8", estiloTitulo)

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
	f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel), "Traslado No.")
	f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel), "Fecha")
	f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel), "Almacenista")

	f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel), "A"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel), "B"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel), "C"+strconv.Itoa(filaExcel), estiloCabecera)

	filaExcel++

	for i, miFila := range t {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel+i), Entero(miFila.Codigo))
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel+i), miFila.Fecha.Format("02/01/2006"))
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel+i), miFila.AlmacenistaNombre)

		f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel+i), "A"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel+i), "B"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel+i), "C"+strconv.Itoa(filaExcel+i), estiloTexto)
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

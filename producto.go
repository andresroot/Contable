package main

// INICIA PRODUCTO IMPORTAR PAQUETES
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
	"log"
	"math"
	"net/http"
	"strconv"
)

// INICIA PRODUCTO ESTRUCTURA JSON
type productoJson struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Nombre string `json:"nombre"`
	Iva    string `json:"Iva"`
	Unidad string `json:"Unidad"`
	Precio string `json:"Precio"`
}

// INICIA PRODUCTO ESTRUCTURA
type producto struct {
	Codigo         string
	Nombre         string
	Iva            string
	Unidad         string
	Subgrupo       string
	SubgrupoNombre string
	Tipo           string
	Precio         string
	Costo          string
	Cantidad       string
	Total          string
	Utilidad       string
}

// INICIA PRODUCTO LISTA
func ProductoLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/producto/productoLista.html")
	log.Println("Error producto 0")
	db := dbConn()
	res := []producto{}
	db.Select(&res, "SELECT * FROM producto ORDER BY codigo ASC")
	varmap := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
	}
	tmp.Execute(w, varmap)
}

// INICIA PRODUCTO NUEVO
func ProductoNuevo(w http.ResponseWriter, r *http.Request) {
	// TRAER COPIA DE EDITAR
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	Panel := mux.Vars(r)["panel"]
	Elemento := mux.Vars(r)["elemento"]

	emp := producto{}
	if Codigo == "False" {
	} else {
		if Elemento == "False" {
			err := db.Get(&emp, "SELECT * FROM producto WHERE codigo=$1", Codigo)
			if err != nil {
				log.Fatalln(err)
			}
		}

	}

	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/producto/autocompletaProductocrear.html",
		"vista/autocompleta/autocompletaSubgrupo.html",
		"vista/producto/ProductoNuevo.html")

	parametros := map[string]interface{}{
		// INICIA PRODUCTO EDITAR AUTOCOMPLETADO
		"Codigo":         Codigo,
		"Panel":          Panel,
		"Elemento":       Elemento,
		"emp":            emp,
		"hosting":        ruta,
		"subgrupo":       ListaSubgrupo(),
		"unidaddemedida": ListaUnidaddemedida(),
		"codigocopia":    Codigo,
		"copiar":         "False",
	}
	tmp.Execute(w, parametros)
	// TERMINA TRAER COPIA DE EDITAR
}
func ProductoNuevoCopiar(w http.ResponseWriter, r *http.Request) {
	// TRAER COPIA DE EDITAR
	db := dbConn()
	Codigo := "False"
	Panel := "False"
	Elemento := "False"
	log.Println("prueba1")
	Codigocopia := mux.Vars(r)["codigocopia"]
	emp := producto{}
	if Codigocopia == "False" {
	} else {

		log.Println("prueba codigo" + Codigocopia)

		err := db.Get(&emp, "SELECT * FROM producto WHERE codigo=$1", Codigocopia)
		if err != nil {
			log.Println("prueba error" + err.Error())
			//	log.Fatalln(err)
		}
	}

	log.Println("prueba2 ")
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/producto/autocompletaProductocrear.html",
		"vista/autocompleta/autocompletaSubgrupo.html",
		"vista/producto/ProductoNuevo.html")

	//tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
	//	"vista/autocompleta/autocompletaSubgrupo.html",
	//	"vista/producto/ProductoNuevo.html")

	parametros := map[string]interface{}{
		// INICIA PRODUCTO EDITAR AUTOCOMPLETADO
		"Codigo":         Codigo,
		"Panel":          Panel,
		"Elemento":       Elemento,
		"emp":            emp,
		"hosting":        ruta,
		"subgrupo":       ListaSubgrupo(),
		"unidaddemedida": ListaUnidaddemedida(),
		"codigocopia":    Codigo,
		"copiar":         "True",
	}

	tmp.Execute(w, parametros)
	// TERMINA TRAER COPIA DE EDITAR
}

// INICIA PRODUCTO INSERTAR
func ProductoInsertar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	err := r.ParseForm()
	if err != nil {
		panic(err.Error())
	}
	var t producto
	err = decoder.Decode(&t, r.PostForm)
	if err != nil {
		panic(err.Error())
	}
	var q string
	q = "insert into producto ("
	q += "Codigo,"
	q += "Nombre,"
	q += "Iva,"
	q += "Unidad,"
	q += "Subgrupo,"
	q += "SubgrupoNombre,"
	q += "Tipo,"
	q += "Precio,"
	q += "Costo,"
	q += "Cantidad,"
	q += "Total,"
	q += "Utilidad"
	q += " ) values("
	q += parametros(12)
	q += " ) "

	log.Println("Cadena SQL " + q)
	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	// INICIA GRABAR PRODUCTO INSERTAR
	t.Nombre = Titulo(t.Nombre)
	// TERMINA PRODUCTO GRABAR INSERTAR
	_, err = insForm.Exec(
		t.Codigo,
		t.Nombre,
		t.Iva,
		t.Unidad,
		t.Subgrupo,
		t.SubgrupoNombre,
		t.Tipo,
		Quitacoma(t.Precio),
		Quitacoma(t.Costo),
		Quitacoma(t.Cantidad),
		Quitacoma(t.Total),
		Quitacoma(t.Utilidad))

	if err != nil {
		panic(err)
	}

	// crea saldo bodega
	// crea bodega en los productos
	//db := dbConn()
	res := []bodega{}
	db.Select(&res, "SELECT * FROM bodega ORDER BY codigo ASC")

	for _, miFila := range res {
		q = "insert into saldo ("
		q += "Producto,"
		q += "Bodega,"
		q += "Cantidad"
		q += " ) values("
		q += parametros(3)
		q += " ) "
		log.Println("Cadena SQL nuevo saldo  " + q)

		insForm, err = db.Prepare(q)

		if err != nil {
			panic(err.Error())
		}

		// TERMINA COMPRA GRABAR INSERTAR
		_, err = insForm.Exec(
			t.Codigo,
			miFila.Codigo,
			0)

		if err != nil {
			panic(err)
		}

	}
	http.Redirect(w, r, "/ProductoLista", 301)
}

// INICIA PRODUCTO BUSCAR
func ProductoBuscar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	selDB, err := db.Query("SELECT codigo,"+
		"nombre,unidad,iva,precio FROM producto where codigo LIKE '%' || $1 || '%'  or  nombre LIKE '%' || $1 || '%' ORDER BY"+
		" codigo DESC", Codigo)
	if err != nil {
		panic(err.Error())
	}
	var resJson []productoJson
	var contar int
	contar = 0
	for selDB.Next() {
		contar++
		var id string
		var label string
		var value string
		var nombre string
		var unidad string
		var iva string
		var precio1 string

		err = selDB.Scan(&id, &nombre, &iva, &unidad, &precio1)
		if err != nil {
			panic(err.Error())
		}
		value = id
		label = id + "  -  " + nombre
		resJson = append(resJson, productoJson{id, label, value, nombre, unidad, iva, precio1})
	}
	if err := selDB.Err(); err != nil { // make sure that there was no issue during the process
		log.Println(err)
		return
	}
	if contar == 0 {
		var slice []string
		slice = make([]string, 0)
		data, _ := json.Marshal(slice)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	} else {
		data, _ := json.Marshal(resJson)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

// PRODUCTO BUSCAR NOMBRE
func ProductoBuscarCrear(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	var tabla = "producto"

	selDB, err := db.Query(" select codigo,nombre from "+tabla+
		"  where codigo LIKE '%' || $1 || '%'"+
		"  or  Upper(nombre) LIKE '%' || Upper($1) || '%' "+
		" ORDER BY"+
		" codigo ", Codigo)
	if err != nil {
		panic(err.Error())
	}
	var resJson []documentoJson
	var contar int
	contar = 0
	for selDB.Next() {
		contar++
		var id string
		var label string
		var value string
		var nombre string
		err = selDB.Scan(&id, &nombre)
		if err != nil {
			panic(err.Error())
		}
		value = id
		label = id + "  -  " + nombre
		resJson = append(resJson, documentoJson{id, label, value, nombre})
	}
	if err := selDB.Err(); err != nil { // make sure that there was no issue during the process
		log.Println(err)
		return
	}
	if contar == 0 {
		var slice []string
		slice = make([]string, 0)
		data, _ := json.Marshal(slice)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	} else {
		data, _ := json.Marshal(resJson)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

// INICIA PRODUCTO EXISTE
func ProductoExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM producto  WHERE codigo=$1", Codigo)
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

// INICIA PRODUCTO ACTUAL
func ProductoActual(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	t := producto{}
	var res []producto
	err := db.Get(&t, "SELECT * FROM producto where codigo=$1", Codigo)

	switch err {
	case nil:
		log.Printf("user found: %+v\n", t)
	case sql.ErrNoRows:
		log.Println("user NOT found, no error")
	default:
		log.Printf("error: %s\n", err)
	}

	log.Println("codigo nombre99" + t.Codigo + t.Nombre)
	res = append(res, t)
	data, err := json.Marshal(res)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// INICIA PRODUCTO EDITAR
func ProductoEditar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio producto editar" + Codigo)
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/autocompleta/autocompletaSubgrupo.html",
		"vista/producto/ProductoEditar.html")
	db := dbConn()
	t := producto{}
	err := db.Get(&t, "SELECT * FROM producto where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}

	varmap := map[string]interface{}{
		// INICIA PRODUCTO EDITAR AUTOCOMPLETADO
		"emp":            t,
		"hosting":        ruta,
		"subgrupo":       ListaSubgrupo(),
		"unidaddemedida": ListaUnidaddemedida(),
		// TERMINA PRODUCTO EDITAR AUTOCOMPLETADO
	}
	tmp.Execute(w, varmap)
}

// INICIA PRODUCTO ACTUALIZAR
func ProductoActualizar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	err := r.ParseForm()
	if err != nil {
		panic(err.Error())
		// Handle error
	}
	var t producto
	err = decoder.Decode(&t, r.PostForm)
	if err != nil {
		// Handle error
		panic(err.Error())
	}
	var q string
	q = "UPDATE producto set "
	q += "Nombre=$2,"
	q += "Iva=$3,"
	q += "Unidad=$4,"
	q += "Subgrupo=$5,"
	q += "Subgruponombre=$6,"
	q += "Tipo=$7,"
	q += "Precio=$8,"
	q += "Costo=$9,"
	q += "Cantidad=$10,"
	q += "Total=$11,"
	q += "Utilidad=$12"
	q += " where "
	q += "Codigo=$1"

	log.Println("cadena" + q)
	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	// INICIA GRABAR PRODUCTO ACTUALIZAR
	t.Nombre = Titulo(t.Nombre)
	// TERMINA GRABAR PRODUCTO ACTUALIZAR
	_, err = insForm.Exec(
		t.Codigo,
		t.Nombre,
		t.Iva,
		t.Unidad,
		t.Subgrupo,
		t.SubgrupoNombre,
		t.Tipo,
		Quitacoma(t.Precio),
		Quitacoma(t.Costo),
		Quitacoma(t.Cantidad),
		Quitacoma(t.Total),
		Quitacoma(t.Utilidad))
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/ProductoLista", 301)

}

// INICIA PRODUCTO BORRAR
func ProductoBorrar(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/Producto/ProductoBorrar.html")
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio Producto borrar" + Codigo)
	db := dbConn()
	t := producto{}
	err := db.Get(&t, "SELECT * FROM Producto where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("codigo nombre99 borrar" + t.Codigo + t.Nombre)
	varmap := map[string]interface{}{
		// INICIA PRODUCTO BORRAR AUTOCOMPLETADO
		"emp":            t,
		"hosting":        ruta,
		"subgrupo":       ListaSubgrupo(),
		"unidaddemedida": ListaUnidaddemedida(),
		// TERMINA PRODUCTO BORRAR AUTOCOMPLETADO
	}
	tmp.Execute(w, varmap)
}

// INICIA PRODUCTO ELIMINAR
func ProductoEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	var q string
	emp := mux.Vars(r)["codigo"]
	delForm, err := db.Prepare("DELETE from producto WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("Registro Eliminado" + emp)

	// borrar bodega del saldo
	q = "delete from saldo where producto=$1"

	log.Println("Cadena SQL nuevo saldo  " + q)

	insForm, err := db.Prepare(q)

	if err != nil {
		panic(err.Error())
	}

	// TERMINA COMPRA GRABAR INSERTAR
	_, err = insForm.Exec(
		emp)

	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/ProductoLista", 301)
}

// INICIA PRODUCTO PDF
func ProductoPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	t := producto{}
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	err := db.Get(&t, "SELECT * FROM producto where codigo=$1", Codigo)
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
		pdf.CellFormat(190, 10, e.Telefono1+" - "+e.Telefono2, "0", 0, "C", false, 0,
			"")
		pdf.Ln(4)
		pdf.CellFormat(190, 10, ene(c.NombreCiudad+" - "+c.NombreDepartamento), "0", 0, "C", false, 0,
			"")
		pdf.Ln(10)

		// RELLENO TITULO
		pdf.SetX(20)
		pdf.SetFillColor(224, 231, 239)
		pdf.SetTextColor(0, 0, 0)

		pdf.SetX(20)
		pdf.CellFormat(184, 6, "DATOS PRODUCTO", "0", 0,
			"C", true, 0, "")
		pdf.Ln(8)
	})

	pdf.SetTextColor(0, 0, 0)
	pdf.SetX(21)
	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.SetX(21)

	pdf.CellFormat(40, 4, "Codigo", "", 0,
		"", false, 0, "")
	pdf.CellFormat(142, 4, t.Codigo, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Nombre:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Nombre, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Iva:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Iva, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Unidad:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, TraerUnidaddemedida(t.Unidad), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Subgrupo:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, TraerSubgrupo(t.Subgrupo), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Tipo:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Tipo, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Precio Venta:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Coma(t.Precio), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Costo:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Coma(t.Costo), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Cantidad:", "0", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Coma(t.Cantidad), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Total:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Coma(t.Total), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Utilidad:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Coma(t.Utilidad), "", 0,
		"", false, 0, "")

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

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

// TERMINA PRODUCTO PDF

// INICIA PRODUCTO TODOS PDF
func ProductoTodosCabecera(pdf *gofpdf.Fpdf) {
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
	pdf.CellFormat(40, 6, "Codigo", "0", 0,
		"L", false, 0, "")
	pdf.SetX(60)
	pdf.CellFormat(40, 6, "Nombre", "0", 0,
		"L", false, 0, "")
	pdf.SetX(87)
	pdf.CellFormat(40, 6, "Iva", "0", 0,
		"R", false, 0, "")
	pdf.SetX(102)
	pdf.CellFormat(40, 6, "Saldo", "0", 0,
		"R", false, 0, "")
	pdf.SetX(123)
	pdf.CellFormat(40, 6, "Venta", "0", 0,
		"R", false, 0, "")
	pdf.SetX(143)
	pdf.CellFormat(40, 6, "Costo", "0", 0,
		"R", false, 0, "")
	pdf.SetX(163)
	pdf.CellFormat(40, 6, "Utilidad", "0", 0,
		"R", false, 0, "")
	pdf.Ln(8)
}

func ProductoTodosDetalle(pdf *gofpdf.Fpdf, miFila producto, a int) {
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(21)
	pdf.CellFormat(181, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(32)
	pdf.CellFormat(40, 4, Subcadena(miFila.Codigo, 0, 12), "", 0,
		"L", false, 0, "")
	pdf.SetX(60)
	pdf.CellFormat(40, 4, miFila.Nombre, "", 0, "L", false, 0, "")
	pdf.SetX(85)
	pdf.CellFormat(40, 4, miFila.Iva, "", 0,
		"R", false, 0, "")
	pdf.SetX(102)
	pdf.CellFormat(40, 4, Coma(miFila.Cantidad), "", 0,
		"R", false, 0, "")
	pdf.SetX(124)
	pdf.CellFormat(40, 4, Coma(miFila.Precio), "", 0,
		"R", false, 0, "")
	pdf.SetX(144)
	pdf.CellFormat(40, 4, Coma(miFila.Costo), "", 0,
		"R", false, 0, "")
	pdf.SetX(164)
	pdf.CellFormat(40, 4, Coma(miFila.Utilidad), "", 0,
		"R", false, 0, "")
	pdf.Ln(4)
}

func ProductoTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	t := []producto{}
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	err := db.Select(&t, "SELECT * FROM producto ORDER BY cast(codigo as integer) ")
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
		pdf.CellFormat(190, 10, "DATOS PRODUCTO", "0", 0,
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

	ProductoTodosCabecera(pdf)
	// tercera pagina

	for i, miFila := range t {
		var a = i + 1
		if math.Mod(float64(a), 49) == 0 {
			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			ProductoTodosCabecera(pdf)
		}
		ProductoTodosDetalle(pdf, miFila, a)
	}

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

// TERMINA PRODUCTO TODOS PDF

// PRODUCTO EXCEL
func ProductoExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	t := []producto{}
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	err := db.Select(&t, "SELECT * FROM producto ORDER BY cast(codigo as integer) ")
	if err != nil {
		log.Fatalln(err)
	}

	f := excelize.NewFile()

	// FUNCION ANCHO DE LA COLUMNA
	if err = f.SetColWidth("Sheet1", "A", "A", 13); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.SetColWidth("Sheet1", "B", "B", 50); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.SetColWidth("Sheet1", "C", "C", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "D", "D", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "E", "E", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "F", "F", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "G", "G", 13); err != nil {
		fmt.Println(err)
		return
	}

	// FUNCION PARA UNIR DOS CELDAS
	if err = f.MergeCell("Sheet1", "A1", "G1"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A2", "G2"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A3", "G3"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A4", "G4"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A5", "G5"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A6", "G6"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A7", "G7"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A8", "G8"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A9", "G9"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A10", "G10"); err != nil {
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
	f.SetCellValue("Sheet1", "A9", "LISTADO DE PRODUCTOS")
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
	f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel), "Codigo")
	f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel), "Nombre")
	f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel), "Iva")
	f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel), "Saldo")
	f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel), "Precio")
	f.SetCellValue("Sheet1", "F"+strconv.Itoa(filaExcel), "Costo")
	f.SetCellValue("Sheet1", "G"+strconv.Itoa(filaExcel), "Utilidad")

	f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel), "A"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel), "B"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel), "C"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel), "D"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "E"+strconv.Itoa(filaExcel), "E"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "F"+strconv.Itoa(filaExcel), "F"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "G"+strconv.Itoa(filaExcel), "G"+strconv.Itoa(filaExcel), estiloCabecera)

	filaExcel++

	for i, miFila := range t {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel+i), miFila.Codigo)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel+i), miFila.Nombre)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel+i), Entero(miFila.Iva))
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel+i), Flotante(miFila.Cantidad))
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel+i), Flotante(miFila.Precio))
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(filaExcel+i), Flotante(miFila.Costo))
		f.SetCellValue("Sheet1", "G"+strconv.Itoa(filaExcel+i), Flotante(miFila.Utilidad))

		f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel+i), "A"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel+i), "B"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel+i), "C"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel+i), "D"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "E"+strconv.Itoa(filaExcel+i), "E"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "F"+strconv.Itoa(filaExcel+i), "F"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "G"+strconv.Itoa(filaExcel+i), "G"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
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

package main

// INICIA COMPROBANTE IMPORTAR PAQUETES
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

// INICIA COMPROBANTE ESTRUCTURA JSON
type comprobanteJson struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Nombre string `json:"nombre"`
}

// TERMINA COMPROBANTE ESTRUCTURA JSON

// INICIA COMPROBANTE ESTRUCTURA
type comprobanteLista struct {
	Documento       string
	Documentonombre string
	Numero          string
	Fecha           time.Time
	Debito          string
	Credito         string
}

// INICIA COMPROBANTE ESTRUCTURA
type comprobante struct {
	Documento         string
	Numero            string
	Fecha             time.Time
	Fechaconsignacion time.Time
	Periodo           string
	Licencia          string
	Usuario           string
	Estado            string
	Debito            string
	Credito           string
	Accion            string
	Detalle           []comprobantedetalle       `json:"Detalle"`
	DetalleEditar     []comprobantedetalleeditar `json:"DetalleEditar"`
}

// INICIA COMPROBANTE DETALLE ESTRUCTURA
type comprobantedetalle struct {
	Fila              string
	Cuenta            string
	Tercero           string
	Centro            string
	Concepto          string
	Factura           string
	Debito            string
	Credito           string
	Documento         string
	Numero            string
	Fecha             time.Time
	Fechaconsignacion time.Time
	Banco             string
	MesConciliacion   string
}

// estructura para editar
type comprobantedetalleeditar struct {
	Fila              string
	Cuenta            string
	CuentaNombre      string
	Tercero           string
	TerceroNombre     string
	Centro            string
	CentroNombre      string
	Concepto          string
	Factura           string
	Debito            string
	Credito           string
	Documento         string
	Numero            string
	Fecha             time.Time
	Fechaconsignacion time.Time
	MesConciliacion   string
}

// TERMINA COMPRA DETALLE EDITAR

// INICIA COMPRA CONSULTA DETALLE
func ComprobanteConsultaDetalle(miperiodo string) string {
	var consulta = ""
	consulta = "select "
	consulta += "comprobantedetalle.Fila as fila, "
	consulta += "comprobantedetalle.Cuenta as cuenta, "
	consulta += "comprobantedetalle.Tercero as tercero, "
	consulta += "comprobantedetalle.Centro as centro, "
	consulta += "comprobantedetalle.Concepto as concepto, "
	consulta += "comprobantedetalle.Factura as factura, "
	consulta += "comprobantedetalle.Debito as debito, "
	consulta += "comprobantedetalle.Credito as credito, "
	consulta += "comprobantedetalle.Documento as documento, "
	consulta += "comprobantedetalle.numero as numero, "
	consulta += "comprobantedetalle.Fecha as fecha, "
	consulta += "comprobantedetalle.Fechaconsignacion as fechaconsignacion, "
	consulta += "comprobantedetalle.Mesconciliacion, "
	consulta += "plandecuentaempresa.nombre as cuentanombre, "
	consulta += "tercero.nombre as terceronombre, "
	consulta += "centro.nombre as centronombre "
	consulta += "from comprobantedetalle "
	consulta += "inner join plandecuentaempresa on plandecuentaempresa.codigo=comprobantedetalle.cuenta "
	consulta += "inner join tercero  on tercero.codigo=comprobantedetalle.tercero "
	consulta += "inner join centro  on centro.codigo=comprobantedetalle.centro "
	consulta += " where (comprobantedetalle.documento=$1 and comprobantedetalle.numero=$2) "
	consulta += " and  extract(year from comprobantedetalle.fecha)=" + miperiodo + " "
	consulta += " order by CAST(comprobantedetalle.fila AS INTEGER)"
	log.Println(consulta)
	return consulta
}

// TERMINA COMPROBANTEDETALLE ESTRUCTURA

// INICIA COMPROBANTE LISTA
func ComprobanteLista(w http.ResponseWriter, r *http.Request) {
	var miperiodo = periodoSesion(r)
	//session, _ := cookie.Get(r, "Golang-session")
	//var miperiodo string
	//miperiodo = "SIN PERIODO"
	//miperiodo = session.Values["periodo"].(string)
	//println("periodo actual" + miperiodo)
	//miempresa=ListaEmpresa()

	//var iperiodo interface{} = session.Values["periodo"]

	//if iperiodo != nil {
	//	miperiodo = session.Values["periodo"].(string)
	//}
	//var miperiodo = ListaEmpresa().Periodo

	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/comprobante/comprobanteLista.html")
	log.Println("Error comprobante 0")
	var consulta string

	consulta = "  SELECT comprobante.documento,comprobante.numero,fecha,documento.nombre as Documentonombre,comprobante.debito,comprobante.credito "
	consulta += " FROM comprobante  "
	consulta += " inner join documento on documento.codigo=comprobante.documento "
	consulta += " where extract(year from fecha)=$1 "
	//	consulta += " ORDER BY fecha  "

	consulta += " ORDER BY cast(comprobante.documento as int ) ASC, fecha, NUMERO  ASC"

	db := dbConn()
	res := []comprobanteLista{}
	//db.Select(&res, consulta)

	//error1 = db.Select(&res, consulta)
	err := db.Select(&res, consulta, miperiodo)

	if err != nil {
		fmt.Println(err)
		return
	}

	varmap := map[string]interface{}{
		"res":       res,
		"hosting":   ruta,
		"miperiodo": miperiodo,
	}
	log.Println("Error comprobante888")
	tmp.Execute(w, varmap)
}

// INICIA COMPROBANTE NUEVO
func ComprobanteNuevo(w http.ResponseWriter, r *http.Request) {
	var miperiodo = periodoSesion(r)
	// TRAE COPIA DE EDITAR
	Documento := mux.Vars(r)["documento"]
	Numero := mux.Vars(r)["numero"]
	log.Println("inicio comprobante editar" + Documento)
	db := dbConn()
	v := comprobante{}
	det := []comprobantedetalleeditar{}
	if Documento == "False" {

	} else {

		err := db.Get(&v, "SELECT * FROM comprobante WHERE documento=$1 and Numero=$2", Documento, Numero)
		if err != nil {
			log.Fatalln(err)
		}

		err2 := db.Select(&det, ComprobanteConsultaDetalle(miperiodo), Documento, Numero)
		if err2 != nil {
			fmt.Println(err2)
		}
	}
	//	log.Println("detalle comprobante)
	parametros := map[string]interface{}{
		"comprobante":        v,
		"detalle":            det,
		"hosting":            ruta,
		"documentoparametro": Documento,
		"documento":          ListaDocumento(),
		"miperiodo":          periodoSesion(r),
	}

	//TERMINA TRAE COPIA DE EDITAR

	log.Println("Error comprobante nuevo 1")
	log.Println("Error comprobante nuevo 2")

	t, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/comprobante/comprobanteNuevo.html",
		"vista/comprobante/comprobanteScript.html",
		"vista/comprobante/autocompletaplandecuentaempresa.html",
		"vista/comprobante/autocompletatercero.html",
		"vista/comprobante/autocompletacentro.html")
	fmt.Printf("%v, %v", t, err)
	log.Println("Error comprobante nuevo 3")
	t.Execute(w, parametros)
}

func ComprobanteAgregarGenerar(tempComprobante comprobante) {
	db := dbConn()

	if tempComprobante.Accion == "Actualizar" {
		// borra detalle anterior
		delForm, err := db.Prepare("DELETE from comprobantedetalle WHERE documento=$1 and numero=$2 and EXTRACT(YEAR FROM  fecha)=$3")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(tempComprobante.Documento, tempComprobante.Numero, tempComprobante.Periodo)

		// borra cabecera anterior

		delForm1, err := db.Prepare("DELETE from comprobante WHERE documento=$1 and numero=$2 and EXTRACT(YEAR FROM  fecha)=$3")
		if err != nil {
			panic(err.Error())
		}
		delForm1.Exec(tempComprobante.Documento, tempComprobante.Numero, tempComprobante.Periodo)
	}

	// INSERTA DETALLE
	for i, x := range tempComprobante.Detalle {
		var a = i
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
		log.Println("Cadena SQL " + q)
		insForm, err := db.Prepare(q)
		if err != nil {
			panic(err.Error())
		}

		// TERMINA COMPROBANTE GRABAR INSERTAR
		_, err = insForm.Exec(
			x.Fila,
			x.Cuenta,
			x.Tercero,
			x.Centro,
			Titulo(x.Concepto),
			x.Factura,
			Flotantedatabase(x.Debito),
			Flotantedatabase(x.Credito),
			x.Documento,
			x.Numero,
			x.Fecha,
			x.Fechaconsignacion)
		if err != nil {
			panic(err)
		}

		log.Println("Insertar Cuenta Detalle fila  \n", x.Cuenta, a)
	}

	log.Println("Got %s age %s club %s\n", tempComprobante.Documento, tempComprobante.Numero)
	var q string
	q += "insert into comprobante ("
	q += "Documento,"
	q += "Numero,"
	q += "Fecha,"
	q += "Fechaconsignacion,"
	q += "Periodo,"
	q += "Licencia,"
	q += "Usuario,"
	q += "Estado,"
	q += "Debito,"
	q += "Credito"
	q += " ) values("
	var coma string
	coma = ""
	for i := 1; i <= 10; i++ {
		q += coma + "$" + strconv.Itoa(i)
		if i == 1 {
			coma = ","
		}
	}
	q += " ) "

	log.Println("Cadena SQL " + q)
	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}
	layout := "2006-01-02"

	log.Println("Hora", tempComprobante.Fecha.Format("02/01/2006"))

	// TERMINA COMPROBANTE GRABAR INSERTAR
	_, err = insForm.Exec(
		tempComprobante.Documento,
		tempComprobante.Numero,
		tempComprobante.Fecha.Format(layout),
		tempComprobante.Fechaconsignacion.Format(layout),
		tempComprobante.Periodo,
		tempComprobante.Licencia,
		tempComprobante.Usuario,
		tempComprobante.Estado,
		Quitacoma(tempComprobante.Debito),
		Quitacoma(tempComprobante.Credito))
	if err != nil {
		panic(err)
	}

	//var resultado bool
	//resultado = true
	//
	//js, err := json.Marshal(SomeStruct{resultado})
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}
	//w.Header().Set("Content-Type", "application/json")
	//w.Write(js)

	//http.Redirect(w, r, "/COMPROBANTELista", 301)
}

// INICIA COMPROBANTE INSERTAR AJAX
func ComprobanteAgregar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var tempComprobante comprobante

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// carga informacion de la COMPROBANTE
	err = json.Unmarshal(b, &tempComprobante)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	if tempComprobante.Accion == "Actualizar" {
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
	}

	// INSERTA DETALLE
	for i, x := range tempComprobante.Detalle {
		var a = i
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
		q += "Fechaconsignacion,"
		q += "Mesconciliacion"
		q += " ) values("
		q += parametros(13)
		q += " ) "
		log.Println("Cadena SQL " + q)
		insForm, err := db.Prepare(q)
		if err != nil {
			panic(err.Error())
		}
		//log.Println("debito " + Flotantedatabase(x.Debito))
		// TERMINA COMPROBANTE GRABAR INSERTAR
		_, err = insForm.Exec(
			x.Fila,
			x.Cuenta,
			x.Tercero,
			x.Centro,
			Titulo(x.Concepto),
			x.Factura,
			Flotantedatabase(x.Debito),
			Flotantedatabase(x.Credito),
			x.Documento,
			x.Numero,
			x.Fecha,
			x.Fechaconsignacion,
			x.MesConciliacion)
		if err != nil {
			panic(err)
		}

		log.Println("Insertar Detalle \n", x.Cuenta, a)
	}

	log.Println("Got %s age %s club %s\n", tempComprobante.Documento, tempComprobante.Numero)
	var q string
	q += "insert into comprobante ("
	q += "Documento,"
	q += "Numero,"
	q += "Fecha,"
	q += "Fechaconsignacion,"
	q += "Periodo,"
	q += "Licencia,"
	q += "Usuario,"
	q += "Estado,"
	q += "Debito,"
	q += "Credito"
	q += " ) values("
	q += parametros(10)
	q += " ) "

	log.Println("Cadena SQL " + q)
	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}
	layout := "2006-01-02"

	log.Println("Hora", tempComprobante.Fecha.Format("02/01/2006"))

	// TERMINA COMPROBANTE GRABAR INSERTAR
	_, err = insForm.Exec(
		tempComprobante.Documento,
		tempComprobante.Numero,
		tempComprobante.Fecha.Format(layout),
		tempComprobante.Fechaconsignacion.Format(layout),
		tempComprobante.Periodo,
		tempComprobante.Licencia,
		tempComprobante.Usuario,
		tempComprobante.Estado,
		Quitacoma(tempComprobante.Debito),
		Quitacoma(tempComprobante.Credito))
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

	//http.Redirect(w, r, "/COMPROBANTELista", 301)
}

// INICIA COMPROBANTE EXISTE
func ComprobanteExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Documento := mux.Vars(r)["documento"]
	Numero := mux.Vars(r)["numero"]
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM comprobante  WHERE documento=$1 and Numero=$2", Documento, Numero)
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

// INICIA COMPROBANTE EDITAR
func ComprobanteEditar(w http.ResponseWriter, r *http.Request) {
	var miperiodo = periodoSesion(r)
	//INICIA COPIA A NUEVO
	Documento := mux.Vars(r)["documento"]
	Numero := mux.Vars(r)["numero"]
	log.Println("inicio comprobante editar" + Documento)
	db := dbConn()

	// traer comprobante
	v := comprobante{}
	err := db.Get(&v, "SELECT * FROM comprobante WHERE (documento=$1 and Numero=$2 ) and  extract(year from comprobante.fecha)="+miperiodo, Documento, Numero)
	if err != nil {
		log.Fatalln(err)
	}

	// traer detalle
	det := []comprobantedetalleeditar{}
	err2 := db.Select(&det, ComprobanteConsultaDetalle(miperiodo), Documento, Numero)
	if err2 != nil {
		fmt.Println(err2)
	}

	//	log.Println("detalle comprobante)
	parametros := map[string]interface{}{
		"comprobante": v,
		"detalle":     det,
		"hosting":     ruta,
		"documento":   ListaDocumento(),
		"miperiodo":   periodoSesion(r),
	}
	//TERMINA COPIA A NUEVO
	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/comprobante/comprobanteEditar.html",
		"vista/comprobante/comprobanteScript.html",
		"vista/comprobante/autocompletaplandecuentaempresa.html",
		"vista/comprobante/autocompletatercero.html",
		"vista/comprobante/autocompletacentro.html")
	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error comprobante nuevo 3")
	miTemplate.Execute(w, parametros)

	//tmp.Execute(w, parametros)
}

// INICIA COMPROBANTE BORRAR
func ComprobanteBorrar(w http.ResponseWriter, r *http.Request) {
	var miperiodo = periodoSesion(r)
	Documento := mux.Vars(r)["documento"]
	Numero := mux.Vars(r)["numero"]
	log.Println("inicio comprobante editar" + Documento)
	db := dbConn()

	// traer COMPROBANTE
	v := comprobante{}
	err := db.Get(&v, "SELECT * FROM comprobante WHERE documento=$1 and Numero=$2", Documento, Numero)
	if err != nil {
		log.Fatalln(err)
	}

	// traer detalle
	det := []comprobantedetalleeditar{}
	err2 := db.Select(&det, ComprobanteConsultaDetalle(miperiodo), Documento, Numero)
	if err2 != nil {
		fmt.Println(err2)
	}

	//	log.Println("detalle producto" + det.Producto+det.ProductoNombre)
	parametros := map[string]interface{}{
		"comprobante": v,
		"detalle":     det,
		"hosting":     ruta,
		"documento":   ListaDocumento(),
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/comprobante/comprobanteBorrar.html",
		"vista/comprobante/comprobanteScript.html",
		"vista/comprobante/autocompletaplandecuentaempresa.html",
		"vista/comprobante/autocompletatercero.html",
		"vista/comprobante/autocompletacentro.html")

	log.Println("Error comprobante nuevo 3")
	miTemplate.Execute(w, parametros)

	//tmp.Execute(w, parametros)
}

// INICIA COMPROBANTE ELIMINAR
func ComprobanteEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	Documento := mux.Vars(r)["documento"]
	Numero := mux.Vars(r)["numero"]

	// borrar COMPROBANTE
	delForm, err := db.Prepare("DELETE from comprobante WHERE documento=$1 and Numero=$2")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(Documento, Numero)

	// borar detalle
	delForm1, err := db.Prepare("DELETE from comprobantedetalle WHERE documento=$1 and Numero=$2")
	if err != nil {
		panic(err.Error())
	}
	delForm1.Exec(Documento, Numero)

	log.Println("Registro Eliminado" + Documento)
	http.Redirect(w, r, "/ComprobanteLista", 301)
}

// INICIA COMPROBANTE PDF
func ComprobantePdf(w http.ResponseWriter, r *http.Request) {
	var miperiodo = periodoSesion(r)
	db := dbConn()
	Documento := mux.Vars(r)["documento"]
	Numero := mux.Vars(r)["numero"]
	// traer COMPROBANTE
	miComprobante := comprobante{}
	err := db.Get(&miComprobante, "SELECT * FROM comprobante where documento=$1 and numero=$2", Documento, Numero)
	if err != nil {
		log.Fatalln(err)
	}

	// traer detalle
	miDetalle := []comprobantedetalleeditar{}
	err2 := db.Select(&miDetalle, ComprobanteConsultaDetalle(miperiodo), Documento, Numero)
	if err2 != nil {
		fmt.Println(err2)
	}
	var miDocumento documento = TraerDocumento(miComprobante.Documento)
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)

	var buf bytes.Buffer
	var err1 error
	pdf := gofpdf.New("P", "mm", "Letter", cnFontDir)
	ene := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetHeaderFunc(func() {
		pdf.Image(imageFile("logo.png"), 20, 25, 40, 0, false,
			"", 0, "")
		pdf.SetY(17)
		pdf.SetFont("Arial", "", 10)
		pdf.CellFormat(190, 10, e.Nombre, "0", 0,
			"C", false, 0, "")
		pdf.Ln(4)
		pdf.CellFormat(190, 10, "Nit No. "+Coma(e.Codigo)+" - "+e.Dv, "0", 0, "C",
			false, 0, "")
		pdf.Ln(4)
		pdf.CellFormat(190, 10, e.Telefono1+" - "+e.Telefono2, "0", 0, "C", false, 0,
			"")
		pdf.Ln(4)
		pdf.CellFormat(190, 10, ene(c.NombreCiudad+" - "+c.NombreDepartamento), "0", 0, "C", false, 0,
			"")
		pdf.Ln(4)

		// NOMBRE DEL DOCUMENTO
		pdf.SetY(20)
		pdf.SetX(80)
		pdf.Ln(5)
		pdf.SetX(80)
		pdf.SetFont("Arial", "", 11)
		pdf.CellFormat(190, 10, miDocumento.Nombre, "0", 0, "C",
			false, 0, "")
		pdf.Ln(6)
		pdf.SetX(80)
		pdf.CellFormat(190, 10, " No.  "+miComprobante.Numero, "0", 0, "C",
			false, 0, "")
		pdf.Ln(10)

		// PIE DE PAGINA
		pdf.SetFooterFunc(func() {
			pdf.SetY(120)
			pdf.SetX(160)
			pdf.SetFont("Arial", "", 8)
			pdf.CellFormat(40, 10, "www.Sadconf.com.co", "",
				0, "L", false, 0, "")
			pdf.SetX(179)
			pdf.CellFormat(30, 10, fmt.Sprintf(" %d de {nb}", pdf.PageNo()), "",
				0, "R", false, 0, "")
		})
	})

	// inicia suma de paginas
	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.Ln(1)

	ComprobanteCabecera(pdf, miComprobante, miDetalle)
	var filas = len(miDetalle)
	// menos de 12
	if filas <= 12 {
		for i, miFila := range miDetalle {
			var a = i + 1
			ComprobanteFilaDetalle(pdf, miFila, a)
		}
		ComprobantePieDePagina(pdf, miComprobante)
	} else {
		// mas de 12 y menos de 27
		if filas > 12 && filas <= 27 {
			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 15 {
					ComprobanteFilaDetalle(pdf, miFila, a)
				}
			}
			// segunda pagina
			pdf.AddPage()
			ComprobanteCabecera(pdf, miComprobante, miDetalle)
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 15 {
					ComprobanteFilaDetalle(pdf, miFila, a)
				}
			}

			ComprobantePieDePagina(pdf, miComprobante)
		} else {
			// mas de 42
			if filas < 43 {
				// primera pagina
				for i, miFila := range miDetalle {
					var a = i + 1
					if a <= 15 {
						ComprobanteFilaDetalle(pdf, miFila, a)
					}
				}

				pdf.AddPage()
				ComprobanteCabecera(pdf, miComprobante, miDetalle)
				// segunda pagina
				for i, miFila := range miDetalle {
					var a = i + 1
					if a > 15 && a <= 30 {
						ComprobanteFilaDetalle(pdf, miFila, a)
					}
				}

				pdf.AddPage()
				ComprobanteCabecera(pdf, miComprobante, miDetalle)
				// tercera pagina
				for i, miFila := range miDetalle {
					var a = i + 1
					if a > 30 {
						ComprobanteFilaDetalle(pdf, miFila, a)
					}
				}

				ComprobantePieDePagina(pdf, miComprobante)
			} else {
				//57

				for i, miFila := range miDetalle {
					var a = i + 1
					if a <= 15 {
						ComprobanteFilaDetalle(pdf, miFila, a)
					}
				}

				pdf.AddPage()
				ComprobanteCabecera(pdf, miComprobante, miDetalle)
				// segunda pagina
				for i, miFila := range miDetalle {
					var a = i + 1
					if a > 15 && a <= 30 {
						ComprobanteFilaDetalle(pdf, miFila, a)
					}
				}

				pdf.AddPage()
				ComprobanteCabecera(pdf, miComprobante, miDetalle)
				// tercera pagina
				for i, miFila := range miDetalle {
					var a = i + 1
					if a > 30 && a <= 45 {
						ComprobanteFilaDetalle(pdf, miFila, a)
					}
				}
				//4 paginoa
				pdf.AddPage()
				ComprobanteCabecera(pdf, miComprobante, miDetalle)
				// tercera pagina
				for i, miFila := range miDetalle {
					var a = i + 1
					if a > 45 {
						ComprobanteFilaDetalle(pdf, miFila, a)
					}
				}

				ComprobantePieDePagina(pdf, miComprobante)

			}
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

// CABECERA PAGINA
func ComprobanteCabecera(pdf *gofpdf.Fpdf, miComprobante comprobante, miDetalle []comprobantedetalleeditar) {
	var utercero string
	for i, miFila := range miDetalle {
		var a = i + 1
		utercero = miFila.Tercero
		a = a + 1
	}

	miTercero := tercero{}
	err3 := db.Get(&miTercero, "SELECT * FROM tercero where codigo=$1", utercero)
	if err3 != nil {
		log.Fatalln(err3)
	}
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Fecha", "", 0,
		"L", false, 0, "")
	pdf.SetX(40)
	pdf.CellFormat(40, 4, miComprobante.Fecha.Format("02/01/2006"), "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Valor $", "", 0,
		"L", false, 0, "")
	pdf.SetX(130)
	pdf.CellFormat(40, 4, FormatoNumeroComprobante(miComprobante.Credito), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Nombre", "", 0,
		"L", false, 0, "")
	pdf.SetX(40)
	pdf.CellFormat(40, 4, miTercero.Nombre, "", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, "Nit. No.", "", 0,
		"L", false, 0, "")
	pdf.SetX(130)
	pdf.CellFormat(40, 4, Coma(miTercero.Codigo), "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)

	// RELLENO TITULO
	pdf.SetX(20)
	pdf.SetFillColor(224, 231, 239)
	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(3)
	pdf.SetX(20)
	pdf.CellFormat(184, 6, "No.", "0", 0,
		"L", true, 0, "")
	pdf.SetX(28)
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(190, 6, "CUENTA", "0", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(190, 6, "NOMBRE", "0", 0,
		"L", false, 0, "")
	pdf.SetX(79)
	pdf.CellFormat(190, 6, "NIT. No.", "0", 0,
		"L", false, 0, "")
	pdf.SetX(100)
	pdf.CellFormat(190, 6, "C.C.", "0", 0,
		"L", false, 0, "")
	pdf.SetX(112)
	pdf.CellFormat(190, 6, "CONCEPTO", "0", 0,
		"L", false, 0, "")
	pdf.SetX(161)
	pdf.CellFormat(190, 6, "DEBITO", "0", 0,
		"L", false, 0, "")
	pdf.SetX(186)
	pdf.CellFormat(190, 6, "CREDITO", "0", 0,
		"L", false, 0, "")
	pdf.Ln(7)
}

// DETALLE DEL COMPROBANTE PDF
func ComprobanteFilaDetalle(pdf *gofpdf.Fpdf, miFila comprobantedetalleeditar, a int) {
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(28)
	pdf.CellFormat(60, 4, miFila.Cuenta, "", 0,
		"L", false, 0, "")
	pdf.SetX(45)
	pdf.CellFormat(60, 4, Subcadena(miFila.CuentaNombre, 0, 20), "", 0,
		"L", false, 0, "")
	pdf.SetX(79)
	pdf.CellFormat(50, 4, Coma(miFila.Tercero), "", 0,
		"L", false, 0, "")
	pdf.SetX(103)
	pdf.CellFormat(30, 4, miFila.Centro, "", 0,
		"L", false, 0, "")
	pdf.SetX(108)
	pdf.CellFormat(40, 4, Subcadena(miFila.Concepto, 0, 25), "", 0,
		"L", false, 0, "")
	pdf.SetX(136)
	pdf.CellFormat(40, 4, FormatoNumeroComprobante(miFila.Debito), "", 0,
		"R", false, 0, "")
	pdf.SetX(165)
	pdf.CellFormat(40, 4, FormatoNumeroComprobante(miFila.Credito), "", 0,
		"R", false, 0, "")
	pdf.Ln(4)
}

func ComprobantePieDePagina(pdf *gofpdf.Fpdf, miComprobante comprobante) {
	// LINEAS
	pdf.SetFont("Arial", "", 9)
	pdf.Line(20, 122, 55, 122)
	pdf.SetY(117)
	pdf.Ln(4)
	pdf.SetX(20)
	pdf.CellFormat(40, 10, "Elaboro", "0", 0, "C",
		false, 0, "")
	pdf.Line(60, 122, 95, 122)
	pdf.SetY(117)
	pdf.Ln(4)
	pdf.SetX(58)
	pdf.CellFormat(40, 10, "Reviso", "0", 0, "C",
		false, 0, "")
	pdf.Line(100, 122, 148, 122)
	pdf.SetY(117)
	pdf.Ln(4)
	pdf.SetX(77)
	pdf.CellFormat(60, 10, "C. C. No,", "0", 0, "C",
		false, 0, "")

	// TOTALES
	pdf.SetFillColor(224, 231, 239)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetY(117)
	pdf.SetX(136)
	pdf.CellFormat(40, 6, FormatoNumeroComprobante(miComprobante.Debito), "0", 0, "R",
		true, 0, "")
	pdf.SetY(117)
	pdf.SetX(165 + 5 + 5)
	pdf.CellFormat(30, 6, FormatoNumeroComprobante(miComprobante.Credito), "0", 0, "R",
		true, 0, "")
}

// INICIA COMPROBANTE TODOS PDF
func ComprobanteTodosCabecera(pdf *gofpdf.Fpdf) {

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
	pdf.CellFormat(190, 6, "Codigo", "0", 0,
		"L", false, 0, "")
	pdf.SetX(55)
	pdf.CellFormat(190, 6, "Nombre", "0", 0,
		"L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(190, 6, "Numero", "0", 0,
		"L", false, 0, "")
	pdf.SetX(137)
	pdf.CellFormat(190, 6, "Fecha", "0", 0,
		"L", false, 0, "")
	pdf.SetX(180)
	pdf.CellFormat(190, 6, "Total", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}
func ComprobanteTodosDetalle(pdf *gofpdf.Fpdf, miFila comprobanteLista, a int) {

	pdf.SetFont("Arial", "", 9)
	pdf.SetX(21)
	pdf.CellFormat(181, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(40, 4, Subcadena(miFila.Documento, 0, 12), "", 0,
		"L", false, 0, "")
	pdf.SetX(46)
	pdf.CellFormat(40, 4, miFila.Documentonombre, "", 0, "L", false, 0, "")
	pdf.SetX(116)
	pdf.CellFormat(40, 4, miFila.Numero, "", 0,
		"L", false, 0, "")
	pdf.SetX(137)
	pdf.CellFormat(40, 4, miFila.Fecha.Format("02/01/2006"), "", 0,
		"L", false, 0, "")
	pdf.SetX(150)
	pdf.CellFormat(40, 4, Coma(miFila.Debito), "", 0,
		"R", false, 0, "")
	pdf.Ln(4)
}
func ComprobanteTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	//	Codigo := mux.Vars(r)["codigo"]
	var consulta string

	consulta = "  SELECT comprobante.documento,comprobante.numero,fecha,documento.nombre as Documentonombre,comprobante.debito,comprobante.credito "
	consulta += " FROM comprobante "
	consulta += " inner join documento on documento.codigo=comprobante.documento "
	consulta += " ORDER BY cast(documento.codigo as integer) ASC"

	t := []comprobanteLista{}
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
		pdf.CellFormat(190, 10, "LISTADO COMPROBANTES", "0", 0,
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

	ComprobanteTodosCabecera(pdf)
	// tercera pagina

	for i, miFila := range t {
		ComprobanteTodosDetalle(pdf, miFila, i+1)

		if math.Mod(float64(i+1), 48) == 0 {
			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			ComprobanteTodosCabecera(pdf)
		}

	}
	//BalancePieDePagina(pdf)

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

// TERMINA COMPROBANTE TODOS PDF

// COMPROBANTE EXCEL
func ComprobanteExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	//	Codigo := mux.Vars(r)["codigo"]
	var consulta string

	consulta = "  SELECT comprobante.documento,comprobante.numero,fecha,documento.nombre as Documentonombre,comprobante.debito,comprobante.credito "
	consulta += " FROM comprobante "
	consulta += " inner join documento on documento.codigo=comprobante.documento "
	consulta += " ORDER BY comprobante.documento ASC"

	t := []comprobanteLista{}
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	err := db.Select(&t, consulta)
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
	f.SetCellValue("Sheet1", "A9", "LISTADO DE COMPROBANTES")
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
	f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel), "Numero")
	f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel), "Fecha")
	f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel), "Total")

	f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel), "A"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel), "B"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel), "C"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel), "D"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "E"+strconv.Itoa(filaExcel), "E"+strconv.Itoa(filaExcel), estiloCabecera)

	filaExcel++

	for i, miFila := range t {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel+i), Flotante(miFila.Documento))
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel+i), miFila.Documentonombre)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel+i), Flotante(miFila.Numero))
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel+i), miFila.Fecha.Format("02/01/2006"))
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel+i), Flotante(miFila.Debito))

		f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel+i), "A"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel+i), "B"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel+i), "C"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel+i), "D"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "E"+strconv.Itoa(filaExcel+i), "E"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		//van=i
	}

	// Set the headers necessary to get browsers to interpret the downloadable file
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment;filename=Libro.xlsx")
	w.Header().Set("File-Name", "Libro.xlsx")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Expires", "0")
	err = f.Write(w)
	if err != nil {
		panic(err.Error())
	}
}

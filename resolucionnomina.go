package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)



// RESOLUCION TABLA
type Resolucionnomina struct {
	Codigo        string
	Numero        string
	Prefijo       string
	Tipo          string
	FechaInicial  time.Time
	FechaFinal    time.Time
	NumeroInicial string
	NumeroFinal   string
	NumeroActual  string
	Local         string
	Direccion     string
	Ciudad        string
	Telefono      string
	Informe       string
	Clavetecnica  string
	Idesoftware	  string
	Testid        string
	Pin           string
	Ambiente      string
}

// INICIA RESOLUCION LISTA
func ResolucionnominaLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/resolucionnomina/resolucionnominaLista.html")
	log.Println("Error resolucionnomina 0")
	db := dbConn()
	res := []Resolucionnomina{}
	db.Select(&res, "SELECT * FROM resolucionnomina ORDER BY cast(codigo as integer) ASC")
	varmap := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
	}
	log.Println("Error resolucionnomina888")
	tmp.Execute(w, varmap)
}

// RESOLUCION NUEVO
func ResolucionnominaNuevo(w http.ResponseWriter, r *http.Request) {
	// TRAER COPIA DE EDITAR
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	emp := Resolucionnomina{}
	if Codigo == "False"{
	} else {
		err := db.Get(&emp, "SELECT * FROM resolucionnomina where codigo=$1", Codigo)
		if err != nil {
			log.Fatalln(err)
		}
	}

	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/resolucionnomina/resolucionnominaNuevo.html")

	parametros := map[string]interface{}{
		"emp":      emp,
		"hosting":  ruta,
		"ciudad":   ListaCiudad(),
		"codigo":   Codigo,
	}
	tmp.Execute(w, parametros)
	// TERMINA TRAER COPIA DE EDITAR
}

// RESOLUCION INSERTAR
func ResolucionnominaInsertar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		Codigo := r.FormValue("Codigo")
		Numero := r.FormValue("Numero")
		Prefijo := r.FormValue("Prefijo")
		Tipo := r.FormValue("Tipo")
		FechaInicial := r.FormValue("FechaInicial")
		FechaFinal := r.FormValue("FechaFinal")
		NumeroInicial := r.FormValue("NumeroInicial")
		NumeroFinal := r.FormValue("NumeroFinal")
		NumeroActual := r.FormValue("NumeroActual")
		Local := r.FormValue("Local")
		Direccion := r.FormValue("Direccion")
		Ciudad := r.FormValue("Ciudad")
		Telefono := r.FormValue("Telefono")
		Informe := r.FormValue("Informe")
		Clavetecnica := r.FormValue("Clavetecnica")
		Idesoftware := r.FormValue("Idesoftware")
		Testid := r.FormValue("Testid")
		Pin := r.FormValue("Pin")
		Ambiente := r.FormValue("Ambiente")

		Prefijo = Mayuscula(Prefijo)
		Local = Titulo(Local)
		Direccion = Titulo(Direccion)

		var q = "INSERT INTO resolucionnomina(" +
			"codigo," +
			"numero," +
			"prefijo," +
			"tipo," +
			"fechainicial," +
			"fechafinal," +
			"numeroinicial," +
			"numerofinal, " +
			"numeroactual, " +
			"local, " +
			"direccion, " +
			"ciudad, " +
			"telefono, " +
			"informe, " +
			"clavetecnica, " +
			"idesoftware, " +
			"testid, " +
			"pin, " +
			"ambiente " +
			")" +
			"VALUES("
		q += parametros(19)
		q += ")"

		insForm, err := db.Prepare(q)
		if err != nil {
			panic(err.Error())

		}

		log.Println("fechafinal:" + FechaFinal)

		_, err =insForm.Exec(Codigo,
			Numero,
			Prefijo,
			Tipo,
			FechaInicial,
			FechaFinal,
			Quitacoma(NumeroInicial),
			Quitacoma(NumeroFinal),
			Quitacoma(NumeroActual),
			Local,
			Direccion,
			Ciudad,
			Telefono,
			Informe,
			Clavetecnica,
			Idesoftware,
			Testid,
			Pin,
			Ambiente)

		if err != nil {
			panic(err)
		}
		log.Println("Nuevo Registro:" + Codigo + "," + Numero)
	}
	http.Redirect(w, r, "/ResolucionnominaLista", 301)
}

// RESOLUCION EXISTE
func ResolucionnominaExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM resolucionnomina  WHERE codigo=$1", Codigo)
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

// INICIA RESOLUCION EDITAR
func ResolucionnominaEditar(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	log.Println("inicio resolucionnomina editar" + Codigo)
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/resolucionnomina/resolucionnominaEditar.html")
	db := dbConn()
	t := Resolucionnomina{}
	err := db.Get(&t, "SELECT * FROM resolucionnomina where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("codigo Numero99" + t.Codigo + t.Numero)
	varmap := map[string]interface{}{
		// INICIA TERCERO EDITAR AUTOCOMPLETADO
		"emp":                     t,
		"hosting":                 ruta,
		"ciudad":                  ListaCiudad(),

		// TERMINA TERCERO EDITAR AUTOCOMPLETADO
	}
	tmp.Execute(w, varmap)
}

// INICIA RESOLUCION ACTUALIZAR
func ResolucionnominaActualizar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	err := r.ParseForm()
	if err != nil {
		panic(err.Error())
		// Handle error
	}
	var t Resolucionnomina
	err = decoder.Decode(&t, r.PostForm)
	decoder.RegisterConverter(time.Time{}, timeConverter)

	if err != nil {
		// Handle error
		panic(err.Error())
	}
	var q string
	q = "UPDATE resolucionnomina set "
	q += "Numero=$2,"
	q += "Prefijo=$3,"
	q += "Tipo=$4,"
	q += "FechaInicial=$5,"
	q += "FechaFinal=$6,"
	q += "NumeroInicial=$7,"
	q += "NumeroFinal=$8,"
	q += "NumeroActual=$9,"
	q += "Local=$10,"
	q += "Direccion=$11,"
	q += "Ciudad=$12,"
	q += "Telefono=$13,"
	q += "Informe=$14,"
	q += "Clavetecnica=$15,"
	q += "Idesoftware=$16,"
	q += "Testid=$17,"
	q += "Pin=$18,"
	q += "Ambiente=$19"
	q += " where "
	q += "Codigo=$1"

	log.Println("cadena" + q)
	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	// INICIA GRABAR RESOLUCION ACTUALIZAR
	t.Prefijo = Mayuscula(t.Prefijo)
	t.Local = Titulo(t.Local)
	t.Direccion = Titulo(t.Direccion)
	// TERMINA GRABAR RESOLUCION ACTUALIZAR
	_, err = insForm.Exec(
		t.Codigo,
		t.Numero,
		t.Prefijo,
		t.Tipo,
		t.FechaInicial,
		t.FechaFinal,
		Quitacoma(t.NumeroInicial),
		Quitacoma(t.NumeroFinal),
		Quitacoma(t.NumeroActual),
		t.Local,
		t.Direccion,
		t.Ciudad,
		t.Telefono,
		t.Informe,
		t.Clavetecnica,
		t.Idesoftware,
		t.Testid,
		t.Pin,
		t.Ambiente)

	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/ResolucionnominaLista", 301)

}

// RESOLUCION BORRAR
func ResolucionnominaBorrar(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/resolucionnomina/resolucionnominaBorrar.html")
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	t := Resolucionnomina{}
	err := db.Get(&t, "SELECT * FROM resolucionnomina where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	varmap := map[string]interface{}{
		"emp":    t,
		"hosting": ruta,
		"ciudad":  ListaCiudad(),
	}
	tmp.Execute(w, varmap)
}

// RESOLUCION ELIMINAR
func ResolucionnominaEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	emp := mux.Vars(r)["codigo"]
	delForm, err := db.Prepare("DELETE from resolucionnomina WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("Registro Eliminado" + emp)
	http.Redirect(w, r, "/ResolucionnominaLista", 301)
}

// INICIA RESOLUCION PDF
func ResolucionnominaPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	t := Resolucionnomina{}
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	err := db.Get(&t, "SELECT * FROM resolucionnomina where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	var ciudadresolucionnomina ciudad = TraerCiudad(t.Ciudad)
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
		pdf.CellFormat(190, 10, "Nit No. " +Coma(e.Codigo)+ " - "+e.Dv, "0", 0, "C",
			false, 0, "")
		pdf.Ln(4)
		pdf.CellFormat(190, 10, e.Iva+ " - "+e.ReteIva, "0", 0, "C", false, 0,
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
		pdf.CellFormat(190, 10, ene(c.NombreCiudad+ " - "+c.NombreDepartamento), "0", 0, "C", false, 0,
			"")
		pdf.Ln(10)

		// RELLENO TITULO
		pdf.SetX(20)
		pdf.SetFillColor(224,231,239)
		pdf.SetTextColor(0,0,0)

		pdf.SetX(20)
		pdf.CellFormat(184, 6, "DATOS RESOLUCION NOMINA", "0", 0,
			"C", true, 0, "")
		pdf.Ln(8)
	})

	pdf.SetTextColor(0,0,0)
	pdf.SetX(21)
	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.SetX(21)

	pdf.CellFormat(40, 4, "Codigo", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Codigo, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Resolucion No.:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Numero, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Prefijo", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Prefijo, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Tipo", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Tipo, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Fecha Inicial", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.FechaInicial.Format("02/01/2006"), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Fecha Final", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.FechaFinal.Format("02/01/2006"), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Numero Inicial", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Coma(t.NumeroInicial), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Numero Final", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Coma(t.NumeroFinal), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Numero Actual", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Coma(t.NumeroActual), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Nombre Local", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Local, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Direccion", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Direccion, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Ciudad", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, ciudadresolucionnomina.NombreCiudad, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Telefono", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Telefono, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Informe", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Informe, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Clave Tecnica", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Clavetecnica, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Ide Software", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Idesoftware, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Test Id", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Testid, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Pin", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Pin, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Ambiente", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Ambiente, "", 0,
		"", false, 0, "")

	pdf.SetFooterFunc(func() {
		pdf.SetTextColor(0,0,0)
		pdf.SetY(252)
		pdf.SetFont("Arial", "", 9)
		pdf.SetX(20)

		// LINEA
		pdf.Line(20,259,204,259)
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
// TERMINA RESOLUCION PDF

// INICIA RESOLUCION NOMINA TODOS PDF
func ResolucionnominaTodosCabecera(pdf *gofpdf.Fpdf){
	pdf.SetFont("Arial", "", 10)
	// RELLENO TITULO
	pdf.SetY(50)
	pdf.SetFillColor(224,231,239)
	pdf.SetTextColor(0,0,0)
	pdf.Ln(7)
	pdf.SetX(20)
	pdf.CellFormat(184, 6, "No", "0", 0,
		"L", true, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(190, 6, "Codigo", "0", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(190, 6, "Numero", "0", 0,
		"L", false, 0, "")
	pdf.SetX(75)
	pdf.CellFormat(190, 6, "Pre-Fijo", "0", 0,
		"L", false, 0, "")
	pdf.SetX(95)
	pdf.CellFormat(190, 6, "F. Inicial", "0", 0,
		"L", false, 0, "")
	pdf.SetX(120)
	pdf.CellFormat(190, 6, "F. Final", "0", 0,
		"L", false, 0, "")
	pdf.SetX(145)
	pdf.CellFormat(190, 6, "N. Inicial", "0", 0,
		"L", false, 0, "")
	pdf.SetX(165)
	pdf.CellFormat(190, 6, "N. Final", "0", 0,
		"L", false, 0, "")
	pdf.SetX(185)
	pdf.CellFormat(190, 6, "N. Actual", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}

func ResolucionnominaTodosDetalle(pdf *gofpdf.Fpdf,miFila Resolucionnomina, a int ){
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(21)
	pdf.CellFormat(181, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(40, 4, miFila.Codigo, "", 0,
		"L", false, 0, "")
	pdf.SetX(40)
	pdf.CellFormat(40, 4, miFila.Numero, "", 0,"L", false, 0, "")
	pdf.SetX(75)
	pdf.CellFormat(40, 4, miFila.Prefijo, "", 0,
		"L", false, 0, "")
	pdf.SetX(95)
	pdf.CellFormat(40, 4, miFila.FechaInicial.Format("02/01/2006"), "", 0,
		"L", false, 0, "")
	pdf.SetX(120)
	pdf.CellFormat(40, 4, miFila.FechaFinal.Format("02/01/2006"), "", 0,
		"L", false, 0, "")
	pdf.SetX(120)
	pdf.CellFormat(40, 4, Coma(miFila.NumeroInicial), "", 0,
		"R", false, 0, "")
	pdf.SetX(140)
	pdf.CellFormat(40, 4, Coma(miFila.NumeroFinal), "", 0,
		"R", false, 0, "")
	pdf.SetX(161)
	pdf.CellFormat(40, 4, Coma(miFila.NumeroActual), "", 0,
		"R", false, 0, "")
	pdf.Ln(4)
}

func ResolucionnominaTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	t := []Resolucionnomina{}
	var e  empresa=ListaEmpresa()
	var c  ciudad=TraerCiudad(e.Ciudad)
	err := db.Select(&t, "SELECT * FROM resolucionnomina ORDER BY cast(codigo as integer) ")
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
		pdf.CellFormat(190, 10, ene(c.NombreCiudad+ " - "+c.NombreDepartamento), "0", 0, "C", false, 0,
			"")
		pdf.Ln(6)
		pdf.CellFormat(190, 10, "DATOS RESOLUCION NOMINA", "0", 0,
			"C", false, 0, "")
		pdf.Ln(10)
	})

	pdf.SetFooterFunc(func() {
		pdf.SetTextColor(0,0,0)
		pdf.SetY(252)
		pdf.SetFont("Arial", "", 9)
		pdf.SetX(20)

		// LINEA
		pdf.Line(20,259,204,259)
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

	ResolucionnominaTodosCabecera(pdf)
	// tercera pagina

	for i, miFila := range t {
		var a = i + 1
		if math.Mod(float64(a),49)==0 {
			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			ResolucionnominaTodosCabecera(pdf)
		}
		ResolucionnominaTodosDetalle(pdf,miFila,a)
	}

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}
// TERMINA RESOLUCION NOMINA TODOS PDF

// RESOLUCION NOMINA EXCEL
func ResolucionnominaExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	t := []Resolucionnomina{}
	var e  empresa=ListaEmpresa()
	var c  ciudad=TraerCiudad(e.Ciudad)
	err := db.Select(&t, "SELECT * FROM resolucionnomina ORDER BY cast(codigo as integer) ")
	if err != nil {
		log.Fatalln(err)
	}

	f := excelize.NewFile()

	// FUNCION ANCHO DE LA COLUMNA
	if err =f.SetColWidth("Sheet1", "A", "A", 13); err != nil {
		fmt.Println(err)
		return
	}
	if err =f.SetColWidth("Sheet1", "B", "B", 50); err != nil {
		fmt.Println(err)
		return
	}
	if err =f.SetColWidth("Sheet1", "C", "C", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err =f.SetColWidth("Sheet1", "D", "D", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err =f.SetColWidth("Sheet1", "E", "E", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err =f.SetColWidth("Sheet1", "F", "F", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err =f.SetColWidth("Sheet1", "G", "G", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err =f.SetColWidth("Sheet1", "H", "H", 13); err != nil {
		fmt.Println(err)
		return
	}
	// FUNCION PARA UNIR DOS CELDAS
	if err = f.MergeCell("Sheet1", "A1", "H1"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A2", "H2"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A3", "H3"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A4", "H4"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A5", "H5"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A6", "H6"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A7", "H7"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A8", "H8"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A9", "H9"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A10", "H10"); err != nil {
		fmt.Println(err)
		return
	}
	estiloTitulo, err := f.NewStyle(`{  "alignment":{"horizontal": "center"},"font":{"bold":false,"italic":false,"family":"Arial","size":10,"color":"##000000"}}`)

	// titulo
	f.SetCellValue("Sheet1", "A1", e.Nombre)
	f.SetCellValue("Sheet1", "A2","Nit No. "+Coma(e.Codigo)+" - "+e.Dv)
	f.SetCellValue("Sheet1", "A3",e.Iva+" - "+e.ReteIva)
	f.SetCellValue("Sheet1", "A4","Actividad Ica - "+e.ActividadIca)
	f.SetCellValue("Sheet1", "A5",e.Direccion)
	f.SetCellValue("Sheet1", "A6",(e.Telefono1+" - "+e.Telefono2))
	f.SetCellValue("Sheet1", "A7",(c.NombreCiudad+" - "+c.NombreDepartamento))
	f.SetCellValue("Sheet1", "A8","")
	f.SetCellValue("Sheet1", "A9","LISTADO DE RESOLUCION NOMINA")
	f.SetCellValue("Sheet1", "A10","")

	f.SetCellStyle("Sheet1","A1","A1",estiloTitulo)
	f.SetCellStyle("Sheet1","A2","A2",estiloTitulo)
	f.SetCellStyle("Sheet1","A3","A3",estiloTitulo)
	f.SetCellStyle("Sheet1","A4","A4",estiloTitulo)
	f.SetCellStyle("Sheet1","A5","A5",estiloTitulo)
	f.SetCellStyle("Sheet1","A6","A6",estiloTitulo)
	f.SetCellStyle("Sheet1","A7","A7",estiloTitulo)
	f.SetCellStyle("Sheet1","A8","A8",estiloTitulo)
	f.SetCellStyle("Sheet1","A9","A7",estiloTitulo)
	f.SetCellStyle("Sheet1","A10","A8",estiloTitulo)

	var filaExcel=11

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

	estiloNumerosindecimales, err := f.NewStyle(`{"alignment":{"horizontal":"left"}, "number_format": 1,"font":{"bold":false,"italic":false,"family":"Arial","size":10,"color":"##000000"}}`)

	if err != nil {
		fmt.Println(err)
	}
	//cabecera
	f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel),"Codigo")
	f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel), "Numero")
	f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel), "Pre-Fijo")
	f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel), "F. Inicial")
	f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel), "F. Final")
	f.SetCellValue("Sheet1", "F"+strconv.Itoa(filaExcel), "N. Inicial")
	f.SetCellValue("Sheet1", "G"+strconv.Itoa(filaExcel), "N. Final")
	f.SetCellValue("Sheet1", "H"+strconv.Itoa(filaExcel), "N. Actual")

	f.SetCellStyle("Sheet1","A"+strconv.Itoa(filaExcel),"A"+strconv.Itoa(filaExcel),estiloCabecera)
	f.SetCellStyle("Sheet1","B"+strconv.Itoa(filaExcel),"B"+strconv.Itoa(filaExcel),estiloCabecera)
	f.SetCellStyle("Sheet1","C"+strconv.Itoa(filaExcel),"C"+strconv.Itoa(filaExcel),estiloCabecera)
	f.SetCellStyle("Sheet1","D"+strconv.Itoa(filaExcel),"D"+strconv.Itoa(filaExcel),estiloCabecera)
	f.SetCellStyle("Sheet1","E"+strconv.Itoa(filaExcel),"E"+strconv.Itoa(filaExcel),estiloCabecera)
	f.SetCellStyle("Sheet1","F"+strconv.Itoa(filaExcel),"F"+strconv.Itoa(filaExcel),estiloCabecera)
	f.SetCellStyle("Sheet1","G"+strconv.Itoa(filaExcel),"G"+strconv.Itoa(filaExcel),estiloCabecera)
	f.SetCellStyle("Sheet1","H"+strconv.Itoa(filaExcel),"H"+strconv.Itoa(filaExcel),estiloCabecera)
	filaExcel++

	for i, miFila := range t{
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel+i), Entero(miFila.Codigo))
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel+i), Flotante(miFila.Numero))
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel+i), miFila.Prefijo)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel+i), miFila.FechaInicial.Format("02/01/2006"))
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel+i), miFila.FechaFinal.Format("02/01/2006"))
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(filaExcel+i), Flotante(miFila.NumeroInicial))
		f.SetCellValue("Sheet1", "G"+strconv.Itoa(filaExcel+i), Flotante(miFila.NumeroFinal))
		f.SetCellValue("Sheet1", "H"+strconv.Itoa(filaExcel+i), Flotante(miFila.NumeroActual))

		f.SetCellStyle("Sheet1","A"+strconv.Itoa(filaExcel+i),"A"+strconv.Itoa(filaExcel+i),estiloNumeroDetalle)
		f.SetCellStyle("Sheet1","B"+strconv.Itoa(filaExcel+i),"B"+strconv.Itoa(filaExcel+i),estiloNumerosindecimales)
		f.SetCellStyle("Sheet1","C"+strconv.Itoa(filaExcel+i),"C"+strconv.Itoa(filaExcel+i),estiloTexto)
		f.SetCellStyle("Sheet1","D"+strconv.Itoa(filaExcel+i),"D"+strconv.Itoa(filaExcel+i),estiloTexto)
		f.SetCellStyle("Sheet1","E"+strconv.Itoa(filaExcel+i),"E"+strconv.Itoa(filaExcel+i),estiloTexto)
		f.SetCellStyle("Sheet1","F"+strconv.Itoa(filaExcel+i),"F"+strconv.Itoa(filaExcel+i),estiloNumeroDetalle)
		f.SetCellStyle("Sheet1","G"+strconv.Itoa(filaExcel+i),"G"+strconv.Itoa(filaExcel+i),estiloNumeroDetalle)
		f.SetCellStyle("Sheet1","H"+strconv.Itoa(filaExcel+i),"H"+strconv.Itoa(filaExcel+i),estiloNumeroDetalle)
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

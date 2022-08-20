package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
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

// CENTRO TABLA
type ConciliacionDatos struct {
	Inicialbanco string         `json:"Inicialbanco"`
	Iniciallibro string         `json:"Iniciallibro"`
	Libromes     string         `json:"Libromes"`
	Detalle      []Conciliacion `json:"Detalle"`
}
type Conciliacion struct {
	Filas     string  `json:"Filas"`
	Fecha     string  `json:"Fecha"`
	Documento string  `json:"Documento"`
	Fila      string  `json:"Fila"`
	Numero    string  `json:"Numero"`
	Concepto  string  `json:"Concepto"`
	Debito    string  `json:"Debito"`
	Credito   string  `json:"Credito"`
	Banco     float64 `json:"Banco"`
	Mc        string  `json:"Mc"`
}

type ConciliacionGuardar struct {
	Fila            string  `json:"Fila"`
	Documento       string  `json:"Documento"`
	Numero          string  `json:"Numero"`
	Debito          string  `json:"Debito"`
	Credito         string  `json:"Credito"`
	Banco           float64 `json:"Banco"`
	MesConciliacion string  `json:"MesConciliacion"`
}

func SumaLibroMes(CodigoCuenta string, Mes string) float64 {
	var consulta string
	listadoDatosDetalle := []datosdetalle{}
	consulta = ""
	consulta = "select  Cuenta,Tercero,Centro,Concepto,Factura ,Debito ,Credito,Documento,Numero,Fecha,Fechaconsignacion  from comprobantedetalle "
	consulta += " where  "
	consulta += " cuenta=$1 and (EXTRACT(MONTH FROM  fecha)=$2 )"

	err2 := db.Select(&listadoDatosDetalle, consulta,
		CodigoCuenta, Mes)

	if err2 != nil {
		panic(err2.Error())
	}
	var debito float64
	var credito float64

	debito = 0
	credito = 0
	// sumar el resultado
	for _, x := range listadoDatosDetalle {
		log.Println("suma propiedad acumulado9999" + FormatoFlotanteEntero(x.Credito))
		debito += x.Debito
		credito += x.Credito
	}

	return (debito - credito)

}
func SumaLibro(CodigoCuenta string, Mes string) float64 {
	var consulta string
	listadoDatosDetalle := []datosdetalle{}
	consulta = ""
	consulta = "select  Cuenta,Tercero,Centro,Concepto,Factura ,Debito ,Credito,Documento,Numero,Fecha,Fechaconsignacion  from comprobantedetalle "
	consulta += " where  "
	consulta += " cuenta=$1 and (EXTRACT(MONTH FROM  fecha)<$2 )"

	err2 := db.Select(&listadoDatosDetalle, consulta,
		CodigoCuenta, Mes)

	if err2 != nil {
		panic(err2.Error())
	}
	var debito float64
	var credito float64

	debito = 0
	credito = 0
	// sumar el resultado
	for _, x := range listadoDatosDetalle {
		log.Println("suma propiedad acumulado9999" + FormatoFlotanteEntero(x.Credito))
		debito += x.Debito
		credito += x.Credito
	}

	return (debito - credito)

}

func SumaBanco(CodigoCuenta string, Mes string) float64 {
	var consulta string
	listadoDatosDetalle := []datosdetalle{}
	consulta = ""
	consulta = "select  Cuenta,Tercero,Centro,Concepto,Factura ,Debito ,Credito,Documento,Numero,Fecha,Fechaconsignacion  from comprobantedetalle "
	consulta += " where  "
	consulta += " cuenta=$1 and ((EXTRACT(MONTH FROM  fecha)<$2 and banco=0) or (EXTRACT(MONTH FROM  fecha)<$2 and cast(mesconciliacion as integer )>=$2 and banco>0))"

	err2 := db.Select(&listadoDatosDetalle, consulta,
		CodigoCuenta, Mes)

	if err2 != nil {
		panic(err2.Error())
	}
	var debito float64
	var credito float64

	debito = 0
	credito = 0
	// sumar el resultado
	for _, x := range listadoDatosDetalle {
		log.Println("suma propiedad acumulado9999" + FormatoFlotanteEntero(x.Credito))
		debito += x.Debito
		credito += x.Credito
	}
	log.Println("Datos bancao")
	log.Println(FormatoFlotante(debito - credito))

	return (debito - credito)

}
func generaConciliacion(CuentaParametro string, MesParametro string) ConciliacionDatos {
	db := dbConn()
	log.Printf(CuentaParametro)
	log.Printf(MesParametro)

	var consulta = " select '' as filas,fila,documento,numero,concepto,fecha,debito,credito,banco,mesconciliacion as mc "
	consulta += " from comprobantedetalle "
	//consulta+=" inner join documento on comprobantedetalle.documento=documento.codigo "
	consulta += "  where  "
	consulta += " cuenta=$3 and ((EXTRACT(MONTH FROM  fecha)<=$1 and banco=0)  or mesconciliacion=$2 or (EXTRACT(MONTH FROM  fecha)=$1) or (EXTRACT(MONTH FROM  fecha)<=$1 and banco>0 and mesconciliacion>=$2 )) "
	consulta += " order by fecha,documento,numero,fila"

	res := []Conciliacion{}

	datos := ConciliacionDatos{}

	res1 := []Conciliacion{}

	err := db.Select(&res, consulta, MesParametro, MesParametro, CuentaParametro)
	switch err {
	//resltadvaa
	case nil:
		log.Printf("Datos existe")
	//	siexiste = true
	case sql.ErrNoRows:
		log.Println("Datos no encontrados")
	default:
		log.Printf("datos error: %s\n", err)
	}
	log.Println("Datos consulta")

	// las marcadas en mes posterior las pone banco 0
	for i, x := range res {
		x.Filas = strconv.Itoa(i + 1)

		var meslista int32
		if meslista, err1 := strconv.Atoi(x.Mc); err1 == nil {
			fmt.Printf("%T, %v", meslista, meslista)
		}

		var mesparametro int32
		if mesparametro, err2 := strconv.Atoi(MesParametro); err2 == nil {
			fmt.Printf("%T, %v", mesparametro, mesparametro)
		}

		if meslista > mesparametro && x.Banco > 0 {
			x.Banco = 0
			x.Filas = "M"
			log.Println("Datos bancarios posterior")
			log.Println(FormatoFlotante(x.Banco))

		}

		t, err := time.Parse("2006-01-02", Subcadena(x.Fecha, 0, 10))
		if err != nil {
			fmt.Println(err)
		}
		x.Fecha = t.Format("02/01/2006")

		res1 = append(res1, x)
	}

	// agrupa datos
	datos.Inicialbanco = FormatoFlotante(SumaLibro(CuentaParametro, MesParametro) - SumaBanco(CuentaParametro, MesParametro))
	datos.Iniciallibro = FormatoFlotante(SumaLibro(CuentaParametro, MesParametro))
	datos.Libromes = FormatoFlotante(SumaLibroMes(CuentaParametro, MesParametro))

	datos.Detalle = res1
	return datos
}

func ConciliacionDato(w http.ResponseWriter, r *http.Request) {
	//db := dbConn()
	CuentaParametro := mux.Vars(r)["cuenta"]
	MesParametro := mux.Vars(r)["mes"]
	log.Printf(CuentaParametro)
	log.Printf(MesParametro)
	datos := ConciliacionDatos{}

	datos = generaConciliacion(CuentaParametro, MesParametro)

	if len(datos.Detalle) == 0 {
		var slice []string
		slice = make([]string, 0)
		data, _ := json.Marshal(slice)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)

	} else {
		data, _ := json.Marshal(datos)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}

}

func ConciliacionInsertar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	//var tempCuentadecobro cuentadecobro
	listacuentadecobroDato := []ConciliacionGuardar{}
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// carga informacion de la CUENTADECOBRO
	err = json.Unmarshal(b, &listacuentadecobroDato)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// INSERTA DETALLE
	for _, x := range listacuentadecobroDato {
		var q string
		var miMes string
		if x.Banco == 0 {
			miMes = ""
		} else {
			miMes = x.MesConciliacion
		}
		q = " update comprobantedetalle set "
		q += " banco= $4 ,"
		q += " mesconciliacion = $5 "
		q += " where Documento = $1 and Numero=$2 and Fila=$3"

		log.Println("Cadena SQL " + q)
		insForm, err := db.Prepare(q)
		if err != nil {
			panic(err.Error())
		}

		// TERMINA CUENTADECOBRO GRABAR INSERTAR
		_, err = insForm.Exec(
			x.Documento, x.Numero, x.Fila, x.Banco, miMes)

		if err != nil {
			panic(err)
		}

		log.Println("Insertar Detalle \n")
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

	//http.Redirect(w, r, "/CUENTADECOBROLista", 301)
}

func ConciliacionLista(w http.ResponseWriter, r *http.Request) {
	parametros := map[string]interface{}{
		//"res":     listadokardex,
		"hosting": ruta,
		"cuenta":  ListaCuentaBanco(),
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/conciliacion/conciliacionLista.html",
		"vista/autocompleta/autocompletaPlandecuentaempresa.html",
		"vista/conciliacion/conciliacionScript.html")
	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error comprobante nuevo 3")
	miTemplate.Execute(w, parametros)
}

// INICIA COMPROBANTE TODOS PDF
func ConciliacionTodosCabecera(pdf *gofpdf.Fpdf) {
	pdf.SetFont("Arial", "", 10)
	// RELLENO TITULO
	pdf.SetY(73)
	pdf.SetFillColor(224, 231, 239)
	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(7)
	pdf.SetX(20)
	pdf.CellFormat(184, 6, "No", "0", 0,
		"L", true, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(190, 6, "Fecha", "0", 0,
		"L", false, 0, "")
	pdf.SetX(52)
	pdf.CellFormat(190, 6, "Documento", "0", 0,
		"L", false, 0, "")
	pdf.SetX(78)
	pdf.CellFormat(190, 6, "Mc", "0", 0,
		"L", false, 0, "")
	pdf.SetX(90)
	pdf.CellFormat(190, 6, "Concepto", "0", 0,
		"L", false, 0, "")
	pdf.SetX(161)
	pdf.CellFormat(40, 6, "Debito", "0", 0,
		"L", false, 0, "")
	pdf.SetX(190)
	pdf.CellFormat(40, 6, "Credito", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}

func ConciliacionTodosDetalle(pdf *gofpdf.Fpdf, miFila Conciliacion, a int) {
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(21)
	pdf.CellFormat(15, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(15, 4, miFila.Fecha, "", 0,
		"L", false, 0, "")
	pdf.SetX(52)
	pdf.CellFormat(10, 4, miFila.Documento+" - "+miFila.Numero, "", 0, "L", false, 0, "")
	pdf.SetX(78)
	pdf.CellFormat(40, 4, miFila.Mc, "", 0,
		"L", false, 0, "")
	pdf.SetX(90)
	pdf.CellFormat(40, 4, miFila.Concepto, "", 0,
		"L", false, 0, "")
	pdf.SetX(144)
	pdf.CellFormat(30, 4, FormatoFlotante(Flotante(miFila.Debito)), "", 0,
		"R", false, 0, "")
	pdf.SetX(174)
	pdf.CellFormat(30, 4, FormatoFlotante(Flotante(miFila.Credito)), "", 0,
		"R", false, 0, "")
	pdf.Ln(4)
}

func ConciliacionTodosPdf(w http.ResponseWriter, r *http.Request) {

	var periodo = periodoSesion(r)
	CuentaParametro := mux.Vars(r)["cuenta"]
	MesParametro := mux.Vars(r)["mes"]
	log.Printf(CuentaParametro)
	log.Printf(MesParametro)
	datos1 := ConciliacionDatos{}
	datosAnterior1 := ConciliacionDatos{}
	var finallibro float64
	var finalbanco float64
	finalbanco = 0
	var totaldebito float64
	var totalcredito float64
	var totaldebitopendiente float64
	var totalcreditopendiente float64
	var totalpendientemes float64

	totaldebito = 0
	totalcredito = 0
	totaldebitopendiente = 0
	totalcreditopendiente = 0

	var totalpendienteanterior float64
	datos1 = generaConciliacion(CuentaParametro, MesParametro)
	if MesParametro == "1" {
		totalpendienteanterior = 0
	} else {
		//var mesanerior int64
		var mesanteriorcadena string

		mesanterior, _ := strconv.Atoi(MesParametro)
		mesanteriorcadena = strconv.Itoa(mesanterior - 1)

		datosAnterior1 = generaConciliacion(CuentaParametro, mesanteriorcadena)

		for _, miFila := range datosAnterior1.Detalle {
			if miFila.Debito == "0.00" {
				totalcredito += Flotante(miFila.Credito)
			} else {
				totaldebito += Flotante(miFila.Debito)
			}
			//mimes := miFila.Mc
			mimes, _ := strconv.Atoi(miFila.Mc)

			//mesparametro := miFila.Mc
			mesparametro, _ := strconv.Atoi(MesParametro)

			if miFila.Mc == "" || mimes >= mesparametro {
				if miFila.Debito == "0.00" {
					totalcreditopendiente += Flotante(miFila.Credito)
				} else {
					totaldebitopendiente += Flotante(miFila.Debito)
				}
			}
		}
		totalpendienteanterior = totaldebitopendiente - totalcreditopendiente

	}

	datos1 = generaConciliacion(CuentaParametro, MesParametro)

	db := dbConn()
	miCuenta := plandecuentaempresa{}
	err := db.Get(&miCuenta, "SELECT * FROM plandecuentaempresa where codigo=$1", CuentaParametro)
	if err != nil {
		log.Fatalln(err)
	}

	totaldebito = 0
	totalcredito = 0
	totaldebitopendiente = 0
	totalcreditopendiente = 0
	totalpendientemes = 0
	for _, miFila := range datos1.Detalle {
		if miFila.Debito == "0.00" {
			totalcredito += Flotante(miFila.Credito)
		} else {
			totaldebito += Flotante(miFila.Debito)
		}
		//mimes := miFila.Mc
		mimes, _ := strconv.Atoi(miFila.Mc)

		//mesparametro := miFila.Mc
		mesparametro, _ := strconv.Atoi(MesParametro)

		if miFila.Mc == "" || mimes > mesparametro {
			if miFila.Debito == "0.00" {
				totalcreditopendiente += Flotante(miFila.Credito)
			} else {
				totaldebitopendiente += Flotante(miFila.Debito)
			}
		}
	}

	totalpendientemes = (totaldebitopendiente - totalcreditopendiente)
	finallibro = Flotante(datos1.Iniciallibro) + Flotante(datos1.Libromes)
	finalbanco = finallibro - (totaldebitopendiente - totalcreditopendiente)
	log.Println("traer conciliacion 1")

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
		pdf.Ln(10)

		// RELLENO TITULO
		pdf.SetX(20)
		pdf.SetFillColor(224, 231, 239)
		pdf.SetTextColor(0, 0, 0)
		pdf.CellFormat(184, 6, "CONCILIACION BANCARIA DEL MES DE"+" "+Mayuscula(mesLetras(MesParametro))+" DE "+" "+periodo, "0", 0,
			"C", true, 0, "")
		pdf.Ln(6)
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

	// DATOS BANCOS
	pdf.SetX(20)
	pdf.CellFormat(40, 10, "Cuenta:", "0", 0,
		"L", false, 0, "")
	pdf.SetX(40)
	pdf.CellFormat(40, 10, CuentaParametro, "0", 0,
		"L", false, 0, "")
	pdf.SetX(60)
	pdf.CellFormat(40, 10, "Nombre:", "0", 0,
		"L", false, 0, "")
	pdf.SetX(80)
	pdf.CellFormat(40, 10, miCuenta.Nombre, "0", 0,
		"L", false, 0, "")
	pdf.Ln(4)
	pdf.SetX(20)
	pdf.CellFormat(40, 10, "Inicial Libros:", "0", 0,
		"L", false, 0, "")
	pdf.SetX(35)
	pdf.CellFormat(40, 10, datos1.Iniciallibro, "0", 0,
		"R", false, 0, "")
	pdf.SetX(80)
	pdf.CellFormat(40, 10, "Inicial Banco:", "0", 0,
		"L", false, 0, "")
	pdf.SetX(95)
	pdf.CellFormat(40, 10, datos1.Inicialbanco, "0", 0,
		"R", false, 0, "")
	pdf.SetX(140)
	pdf.CellFormat(40, 10, "Inicial Pendiente:", "0", 0,
		"L", false, 0, "")
	pdf.SetX(165)
	pdf.CellFormat(40, 10, FormatoFlotante(totalpendienteanterior), "0", 0,
		"R", false, 0, "")
	pdf.Ln(4)
	pdf.SetX(20)
	pdf.CellFormat(40, 10, "Final libros:", "0", 0,
		"L", false, 0, "")
	pdf.SetX(35)
	pdf.CellFormat(40, 10, FormatoFlotante(finallibro), "0", 0,
		"R", false, 0, "")
	pdf.SetX(80)
	pdf.CellFormat(40, 10, "Final Banco:", "0", 0,
		"L", false, 0, "")
	pdf.SetX(95)
	pdf.CellFormat(40, 10, FormatoFlotante(finalbanco), "0", 0,
		"R", false, 0, "")
	pdf.SetX(140)
	pdf.CellFormat(40, 10, "Final Pendiente:", "0", 0,
		"L", false, 0, "")
	pdf.SetX(165)
	pdf.CellFormat(40, 10, FormatoFlotante(totalpendientemes), "0", 0,
		"R", false, 0, "")
	pdf.Ln(6)
	pdf.SetX(20)
	pdf.CellFormat(184, 10, "Documentos Pendientes:", "0", 0,
		"C", false, 0, "")
	pdf.Ln(6)

	ConciliacionTodosCabecera(pdf)

	log.Println("traer conciliacion 2")
	var consecutivo int
	consecutivo = 0
	for i, miFila := range datos1.Detalle {

		mesactual, _ := strconv.Atoi(MesParametro)
		mesfila, _ := strconv.Atoi(miFila.Mc)

		if miFila.Banco == 0 || mesfila > mesactual {
			consecutivo++
			ConciliacionTodosDetalle(pdf, miFila, consecutivo)

			if math.Mod(float64(i+1), 48) == 0 {
				pdf.AliasNbPages("")
				pdf.AddPage()
				pdf.SetFont("Arial", "", 10)
				pdf.SetX(30)
				ConciliacionTodosCabecera(pdf)
			}
		}
	}
	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

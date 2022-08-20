package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"text/template"
	"time"
)

type CertificadoretencionParametro struct {
	FechaInicial    string `json:"FechaInicial"`
	FechaFinal      string `json:"FechaFinal"`
	FechaExpedicion string `json:"FechaExpedicion"`
	TerceroInicial  string `json:"TerceroInicial"`
	TerceroFinal    string `json:"TerceroFinal"`
	CuentaInicial   string `json:"CuentaInicial"`
	CuentaFinal     string `json:"CuentaFinal"`
}
type Certificadoretencionlista struct {
	Fila   string
	Codigo string
	Nombre string
	Accion string
}
type cuentaSaldoRetencion struct {
	Cuenta       string
	CuentaNombre string
	Debito       string
	Credito      string
	Saldo        string
}

func CertificadoretencionDatos(w http.ResponseWriter, r *http.Request) {
	var tempParametro CertificadoretencionParametro
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Println("error 001")
		http.Error(w, err.Error(), 500)
		return
	}
	// carga informacion de la venta
	err = json.Unmarshal(b, &tempParametro)
	if err != nil {
		fmt.Println("error 002")
		http.Error(w, err.Error(), 500)
		return
	}

	res := []Certificadoretencionlista{}
	final := []Certificadoretencionlista{}

	var consulta string
	consulta = "select "
	consulta += " 	distinct tercero.codigo as codigo ,tercero.nombre as nombre from comprobantedetalle "
	consulta += " inner join tercero on tercero.codigo=comprobantedetalle.tercero "
	consulta += " 	where  (comprobantedetalle.tercero>=$1 and comprobantedetalle.tercero<=$2) "
	consulta += " 		and  (comprobantedetalle.cuenta>=$3 and comprobantedetalle.cuenta<=$4) "
	consulta += " 	and (comprobantedetalle.fecha>=$5 and comprobantedetalle.fecha<=$6) "

	err = db.Select(&res, consulta, tempParametro.TerceroInicial, tempParametro.TerceroFinal, tempParametro.CuentaInicial, tempParametro.CuentaFinal, tempParametro.FechaInicial, tempParametro.FechaFinal)

	if err != nil {
		fmt.Println(err)
	}

	for a, miFila := range res {
		miFila.Fila = strconv.Itoa(a + 1)
		final = append(final, Certificadoretencionlista{miFila.Fila, miFila.Codigo, miFila.Nombre, ""})
	}

	data, _ := json.Marshal(final)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func CertificadoretencionLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/certificadoretencion/CertificadoretencionLista.html",
		"vista/certificadoretencion/Autocompletaplandecuentaempresa.html",
		"vista/certificadoretencion/Autocompletatercero.html")

	Codigo := mux.Vars(r)["codigo"]
	if Codigo == "False" {
	} else {
	}

	varmap := map[string]interface{}{
		"hosting":  ruta,
		"bodega":   ListaBodega(),
		"producto": ListaProducto(),
	}
	tmp.Execute(w, varmap)
}

func saldoRetencion(mitercero string, cuentaInicial string, cuentaFinal string, fechaInicial time.Time, fechaFinal time.Time) []cuentaSaldoRetencion {

	var consultasaldoAnterior = "SELECT distinct comprobantedetalle.cuenta as cuenta,"
	consultasaldoAnterior += "  plandecuentaempresa.nombre as cuentanombre,sum(debito) as debito,sum(credito ) as credito "
	consultasaldoAnterior += "  , sum(credito-debito) as saldo "
	consultasaldoAnterior += " from comprobantedetalle "
	consultasaldoAnterior += " Inner Join plandecuentaempresa on plandecuentaempresa.codigo=comprobantedetalle.cuenta "
	consultasaldoAnterior += "  where"
	consultasaldoAnterior += " comprobantedetalle.tercero=$1  "
	consultasaldoAnterior += " and (fecha>=$2 and fecha<=$3) "
	consultasaldoAnterior += " and (comprobantedetalle.cuenta>=$4 and comprobantedetalle.cuenta<=$5) "
	consultasaldoAnterior += " group by comprobantedetalle.cuenta,plandecuentaempresa.nombre"

	listadoSaldoTercero := []cuentaSaldoRetencion{}
	err := db.Select(&listadoSaldoTercero, consultasaldoAnterior, mitercero, fechaInicial, fechaFinal, cuentaInicial, cuentaFinal)

	if err != nil {
		fmt.Println(err)
	}

	var totalSaldo float64
	totalSaldo = 0
	for _, miSaldo := range listadoSaldoTercero {
		log.Println("suma saldo RETENCIONFUENTE ")
		log.Println(miSaldo.Debito)
		log.Println(miSaldo.Credito)

		//var tipocuenta = Subcadena(miSaldo.Cuenta, 0, 1)
		//log.Println("tipo" + tipocuenta)
		//if tipocuenta == "1" || tipocuenta == "5" || tipocuenta == "6" || tipocuenta == "7" || tipocuenta == "8" {
		totalSaldo += Flotante(miSaldo.Credito)
		//} else {
		//	totalSaldo += Flotante(miSaldo.Credito) - Flotante(miSaldo.Debito)
		//}
	}
	return listadoSaldoTercero

}

func baseRetencion(mitercero string, fechaInicial time.Time, fechaFinal time.Time) float64 {

	var consultasaldoAnterior = "SELECT cuenta,debito,credito, debito-credito as saldo FROM comprobantedetalle  where"
	consultasaldoAnterior += " tercero=$1  and (fecha>=$2 and fecha<=$3)  and "
	consultasaldoAnterior += " (substring(cuenta,1,2)='14'  or "
	consultasaldoAnterior += " substring(cuenta,1,2)='15'  or "
	consultasaldoAnterior += " substring(cuenta,1,1)='5'  or "
	consultasaldoAnterior += " substring(cuenta,1,1)='6'  or "
	consultasaldoAnterior += " substring(cuenta,1,1)='7'  )"

	//
	listadoSaldoTercero := []cuentaSaldoRetencion{}
	err := db.Select(&listadoSaldoTercero, consultasaldoAnterior, mitercero, fechaInicial, fechaFinal)

	if err != nil {
		fmt.Println(err)
	}

	// fechaInicial, fechaFinal,
	//	log.Println("calculo cuenta anterior " + micuenta)
	var totalSaldo float64
	totalSaldo = 0
	for _, miSaldo := range listadoSaldoTercero {
		//var c=k
		log.Println("suma saldo base retencion")

		log.Println(miSaldo.Debito)
		log.Println(miSaldo.Credito)
		totalSaldo += Flotante(miSaldo.Saldo)

	}
	return totalSaldo

}

func CertificadoretencionPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	CuentaInicial := mux.Vars(r)["cuentainicial"]
	CuentaFinal := mux.Vars(r)["cuentafinal"]
	dateinicial, err := time.Parse("2006-01-02", mux.Vars(r)["fechainicial"])
	datefinal, err := time.Parse("2006-01-02", mux.Vars(r)["fechafinal"])
	datecertificado, err := time.Parse("2006-01-02", mux.Vars(r)["fechaexpedicion"])

	miTercero := tercero{}
	var miEmpresa empresa = ListaEmpresa()
	var miCiudad ciudad = TraerCiudad(miEmpresa.Ciudad)
	err = db.Get(&miTercero, "SELECT * FROM tercero where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	var buf bytes.Buffer
	var err1 error
	pdf := gofpdf.New("P", "mm", "Letter", cnFontDir)
	ene := pdf.UnicodeTranslatorFromDescriptor("")

	listadoCuentaRetencion := saldoRetencion(Codigo, CuentaInicial, CuentaFinal, dateinicial, datefinal)
	pdf.SetFooterFunc(func() {
		pdf.Ln(3)
		pdf.SetFont("Arial", "", 8)
		pdf.SetX(21)
		pdf.SetY(121)
		pdf.CellFormat(40, 4, "Sadconf.com", "", 0,
			"C", false, 0, "")
		pdf.SetX(165)
		pdf.CellFormat(40, 4, fmt.Sprintf(" %d de {nb}", pdf.PageNo()), "",
			0, "R", false, 0, "")
	})

	for _, miCuenta := range listadoCuentaRetencion {

		var nombreCuenta = miCuenta.CuentaNombre
		var valorRetencion = miCuenta.Saldo
		var valorBase = baseRetencion(Codigo, dateinicial, datefinal)

		pdf.SetX(21)
		pdf.AliasNbPages("")
		pdf.AddPage()

		// VERTICAL
		pdf.SetDrawColor(84, 153, 199)
		pdf.Line(21, 14, 21, 120)
		pdf.Line(201, 14, 201, 120)
		pdf.Line(21, 120, 201, 120)
		pdf.Line(110, 33, 110, 93)

		// LINEA HORIZONTAL
		pdf.SetDrawColor(84, 153, 199)
		pdf.Line(21, 14, 201, 14)
		pdf.Line(21, 33, 201, 33)
		pdf.Line(21, 39, 201, 39)
		pdf.Line(21, 45, 201, 45)
		pdf.Line(21, 51, 201, 51)
		pdf.Line(21, 57, 201, 57)
		pdf.Line(21, 63, 201, 63)
		pdf.Line(21, 69, 201, 69)
		pdf.Line(21, 75, 201, 75)
		pdf.Line(21, 81, 201, 81)
		pdf.Line(21, 87, 201, 87)
		pdf.Line(21, 93, 201, 93)
		pdf.Line(21, 99, 201, 99)
		pdf.SetFont("Arial", "B", 10)
		pdf.SetY(17)
		pdf.SetX(40)
		pdf.CellFormat(140, 4, ene("Certificado de Retenciones Por Otros Conceptos"), "0", 0, "C", false, 0, "")
		pdf.SetY(22)
		pdf.SetX(40)
		pdf.CellFormat(140, 4, ene("Articulo 381 del E. T."), "0", 0, "C", false, 0, "")
		pdf.SetY(27)
		pdf.SetX(40)
		pdf.CellFormat(140, 4, ene("Año Gravable 2021"), "0", 0, "C", false, 0, "")

		pdf.SetFont("Arial", "", 8)
		pdf.Ln(2)
		pdf.Ln(-1)
		pdf.SetX(21)
		pdf.SetFillColor(225, 232, 239)
		pdf.CellFormat(89, 6, "CIUDAD DONDE SE CONSIGNO LA RETENCION", "1", 0,
			"L", true, 0, "")
		pdf.SetFont("Arial", "", 8)
		pdf.CellFormat(91, 6, ene(miCiudad.Nombre), "1", 0,
			"L", true, 0, "")
		pdf.Ln(-1)
		pdf.SetX(21)
		pdf.CellFormat(89, 6, "NOMBRE DEL AGENTE RETENEDOR", "", 0,
			"L", false, 0, "")
		pdf.CellFormat(91, 6, miEmpresa.Nombre, "", 0,
			"L", false, 0, "")
		pdf.Ln(-1)
		pdf.SetX(21)
		pdf.CellFormat(89, 6, "NUMERO DE IDENTIFICACION TRIBUTARIA", "1", 0,
			"L", true, 0, "")
		//pdf.SetFont("Arial", "", 8)
		pdf.CellFormat(91, 6, Coma(miEmpresa.Codigo)+" - "+miEmpresa.Dv, "1", 0,
			"L", true, 0, "")
		pdf.Ln(-1)
		pdf.SetX(21)
		pdf.CellFormat(89, 6, "DIRECCION DEL AGENTE RETENEDOR", "", 0,
			"L", false, 0, "")
		pdf.CellFormat(91, 6, miEmpresa.Direccion, "", 0,
			"L", false, 0, "")
		pdf.Ln(-1)
		pdf.SetX(21)
		pdf.CellFormat(89, 6, "APELLIDOS Y NOMBRES O RAZON SOCIAL BENEFICIARIO", "1", 0,
			"L", true, 0, "")
		pdf.CellFormat(91, 6, miTercero.Nombre, "1", 0,
			"L", true, 0, "")
		pdf.Ln(-1)
		pdf.SetX(21)
		pdf.CellFormat(89, 6, "NIT DEL BENEFICIARIO", "", 0,
			"L", false, 0, "")
		pdf.CellFormat(91, 6, Coma(miTercero.Codigo)+" - "+miTercero.Dv, "", 0,
			"L", false, 0, "")
		pdf.Ln(-1)
		pdf.SetX(21)
		pdf.CellFormat(89, 6, "MONTO TOTAL PAGADO", "1", 0,
			"L", true, 0, "")
		pdf.CellFormat(91, 6, "$ "+Coma(FormatoFlotanteEntero(valorBase)), "1", 0,
			"L", true, 0, "")
		pdf.Ln(-1)
		pdf.SetX(21)
		pdf.CellFormat(89, 6, "CONCEPTO DEL PAGO SUJETO A RETENCION", "", 0,
			"L", false, 0, "")
		pdf.CellFormat(91, 6, ene(nombreCuenta), "", 0,
			"L", false, 0, "")
		pdf.Ln(-1)
		pdf.SetX(21)
		pdf.CellFormat(89, 6, "VALOR RETENIDO", "1", 0,
			"L", true, 0, "")
		pdf.CellFormat(91, 6, "$ "+Coma(FormatoFlotanteEntero(Flotante(valorRetencion))), "1", 0,
			"L", true, 0, "")
		pdf.Ln(-1)
		pdf.SetX(21)
		pdf.CellFormat(89, 6, "PERIODO DEL CERTIFICADO", "", 0,
			"L", false, 0, "")
		pdf.CellFormat(91, 6, "Del "+dateinicial.Format("02/01/2006")+" Al "+datefinal.Format("02/01/2006"), "", 0,
			"L", false, 0, "")
		pdf.Ln(-1)
		pdf.SetX(21)
		pdf.CellFormat(89, 6, "FECHA DE EXPEDICION", "1", 0,
			"L", true, 0, "")
		pdf.CellFormat(91, 6, fechaLetras(datecertificado), "1", 0,
			"L", true, 0, "")
		pdf.Ln(-1)

		pdf.Ln(2)
		pdf.SetX(21)
		pdf.SetFont("Arial", "", 10)
		pdf.CellFormat(89, 6, "Se omite firma autografa segun Articulo "+
			"10 del D. R. 836 de 1991, Recopilado en el articulo 1.6.1.12.12 del",
			"0", 0,
			"L", false, 0, "")
		pdf.Ln(-1)
		pdf.SetX(21)
		pdf.CellFormat(89, 6, " Decreto Unico Tributario 1625 del 11 de Octubre de 2016 ", "0", 0,
			"L", false, 0, "")
	}

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

// listado

func CertificadoRetencionTodosDetalle(pdf *gofpdf.Fpdf, miFila Certificadoretencionlista, a int) {
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(21)
	pdf.CellFormat(181, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(32)
	pdf.CellFormat(40, 4, Subcadena(miFila.Codigo, 0, 12), "", 0,
		"L", false, 0, "")
	pdf.SetX(55)
	pdf.CellFormat(40, 4, miFila.Nombre, "", 0, "L", false, 0, "")
	pdf.SetX(75)
	//pdf.CellFormat(40, 4, Subcadena(miFila.TerceroNombre, 0, 29),
	//	"", 0, "L", false, 0, "")
	//pdf.SetX(130)
	//pdf.CellFormat(40, 4, Subcadena(miFila.VendedorNombre, 0, 31), "", 0,
	//	"L", false, 0, "")
	//pdf.SetX(164)
	//pdf.CellFormat(40, 4, Coma(miFila.Total), "", 0,
	//	"R", false, 0, "")
	pdf.Ln(4)
}
func CertificadoRetencionTodosCabecera(pdf *gofpdf.Fpdf) {
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
	pdf.CellFormat(190, 6, "Codigo", "0", 0,
		"L", false, 0, "")
	pdf.SetX(55)
	pdf.CellFormat(190, 6, "Nombre", "0", 0,
		"L", false, 0, "")
	pdf.SetX(75)
	pdf.CellFormat(190, 6, "", "0", 0,
		"L", false, 0, "")
	pdf.SetX(130)
	pdf.CellFormat(190, 6, "", "0", 0,
		"L", false, 0, "")
	pdf.SetX(193)
	//pdf.CellFormat(190, 6, "Total", "0", 0,
	//	"L", false, 0, "")
	pdf.Ln(8)
}

func CertificadoRetencionTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	//CuentaInicial := mux.Vars(r)["cuentainicial"]
	//CuentaFinal := mux.Vars(r)["cuentafinal"]
	//dateinicial, err := time.Parse("2006-01-02", mux.Vars(r)["fechainicial"])
	//datefinal, err := time.Parse("2006-01-02", mux.Vars(r)["fechafinal"])
	//datecertificado, err := time.Parse("2006-01-02", mux.Vars(r)["fechaexpedicion"])

	var tempParametro CertificadoretencionParametro
	tempParametro.CuentaInicial = mux.Vars(r)["cuentainicial"]
	tempParametro.CuentaFinal = mux.Vars(r)["cuentafinal"]
	tempParametro.FechaInicial = mux.Vars(r)["fechainicial"]
	tempParametro.FechaFinal = mux.Vars(r)["fechafinal"]
	tempParametro.FechaExpedicion = mux.Vars(r)["fechaexpedicion"]
	tempParametro.TerceroInicial = mux.Vars(r)["terceroinicial"]
	tempParametro.TerceroFinal = mux.Vars(r)["tercerofinal"]

	listaRetencion := []Certificadoretencionlista{}

	var consulta string
	consulta = "select "
	consulta += " 	distinct tercero.codigo as codigo ,tercero.nombre as nombre from comprobantedetalle "
	consulta += " inner join tercero on tercero.codigo=comprobantedetalle.tercero "
	consulta += " 	where  (comprobantedetalle.tercero>=$1 and comprobantedetalle.tercero<=$2) "
	consulta += " 		and  (comprobantedetalle.cuenta>=$3 and comprobantedetalle.cuenta<=$4) "
	consulta += " 	and (comprobantedetalle.fecha>=$5 and comprobantedetalle.fecha<=$6) "

	err := db.Select(&listaRetencion, consulta, tempParametro.TerceroInicial, tempParametro.TerceroFinal, tempParametro.CuentaInicial, tempParametro.CuentaFinal, tempParametro.FechaInicial, tempParametro.FechaFinal)
	//err := db.Select(&res, consulta)

	if err != nil {
		fmt.Println(err)
	}

	//var consulta string
	//consulta = "  SELECT vendedor.nombre as VendedorNombre,centro.nombre as CentroNombre,venta.total,venta.codigo,fecha,tercero,tercero.nombre as TerceroNombre "
	//consulta += " FROM venta "
	//consulta += " inner join tercero on tercero.codigo=venta.tercero "
	//consulta += " inner join centro on centro.codigo=venta.centro "
	//consulta += " inner join vendedor on vendedor.codigo=venta.vendedor "
	//consulta += " ORDER BY cast(venta.codigo as integer) ASC"
	//
	//t := []ventaLista{}
	//err := db.Select(&t, consulta)

	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

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
		pdf.CellFormat(190, 10, "Certificado De Retencion", "0", 0,
			"C", false, 0, "")
		pdf.Ln(10)
	})

	pdf.SetFooterFunc(func() {
		pdf.SetTextColor(0, 0, 0)
		pdf.SetY(253)
		pdf.SetFont("Arial", "", 9)
		pdf.SetX(20)

		// LINEA
		pdf.Line(20, 260, 204, 260)
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

	CertificadoRetencionTodosCabecera(pdf)
	// tercera pagina

	for i, miFila := range listaRetencion {
		var a = i + 1
		if math.Mod(float64(a), 49) == 0 {
			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			CertificadoRetencionTodosCabecera(pdf)
		}
		CertificadoRetencionTodosDetalle(pdf, miFila, a)
	}

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/jung-kurt/gofpdf"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"
)

// PROPIEDAD TABLA
type propiedad struct {
	Codigo              string
	Nombre              string
	Tercero             string
	Terceronombre       string
	Cuenta              string
	Ubicacion           string
	Fecha               time.Time
	Inicia              time.Time
	Valor               string
	Vresidual           string
	Vidautil            string
	Totalmes            string
	Acumulado           string
	Saldo               string
	Libros              string
	Cuentagasto         string
	Cuentacontra        string
	Centro              string
	CuentaDetalle       plandecuentaempresa
	CuentagastoDetalle  plandecuentaempresa
	CuentacontraDetalle plandecuentaempresa
	CentroDetalle       centro
	TerceroDetalle      tercero
}

// lista propiedad
type listaPropiedad struct {
	Codigo      string
	Nombre      string
	Fecha       time.Time
	Vidautil    string
	Valor       string
	Vresidual   string
	Totalmes    string
	Acumulado   string
	Saldo       string
	Libros      string
	Cuentagasto string
}

var consultalistapropiedad = "select codigo,nombre,fecha,vidautil,valor,vresidual,totalmes," +
	"acumulado,saldo,libros,cuentagasto from propiedad"

// CUENTA JSON
type propiedadJson struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Nombre string `json:"nombre"`
}

type Propiedadresultado struct {
	Documento string `json:"Documento"`
	Numero    string `json:"Numero"`
}

// PROPIEDAD BUSCAR
func PropiedadBuscar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	selDB, err := db.Query("SELECT codigo,"+
		"nombre FROM propiedad where codigo LIKE '%' || $1 || '%' ORDER BY"+
		" codigo DESC", Codigo)
	if err != nil {
		panic(err.Error())
	}
	var resJson []propiedadJson
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
		resJson = append(resJson, propiedadJson{id, label, value, nombre})
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

func SumarPropiedad(CodigoCuenta string, Codigopropiedad string) string {
	var consulta string
	listadoDatosDetalle := []datosdetalle{}
	consulta = ""
	consulta = "select  Cuenta,Tercero,Centro,Concepto,Factura ,Debito ,Credito,Documento,Numero,Fecha,Fechaconsignacion  from comprobantedetalle "
	consulta += " where  "
	consulta += " cuenta=$1 and factura=$2"

	err2 := db.Select(&listadoDatosDetalle, consulta,
		CodigoCuenta, Codigopropiedad)

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

	return FormatoFlotanteEntero(debito)

}

// PROPIEDAD LISTA
func PropiedadLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/propiedad/propiedadLista.html")
	db := dbConn()

	res := []listaPropiedad{}
	res2 := []listaPropiedad{}
	err := db.Select(&res, consultalistapropiedad)

	var totalsaldo float64
	totalsaldo = 0

	var totallibros float64
	totallibros = 0

	var totalacumulado float64
	totalacumulado = 0

	var totaltotalmes float64
	totaltotalmes = 0

	var totalvresidual float64
	totalvresidual = 0

	var totalvalor float64
	totalvalor = 0

	for _, x := range res {
		x.Acumulado = Quitacoma(SumarPropiedad(x.Cuentagasto, x.Codigo))
		log.Println("suma propiedad acumulado" + x.Acumulado)
		x.Libros = Quitacoma(FormatoFlotanteEntero(Flotante(x.Valor) - Flotante(x.Acumulado)))
		x.Saldo = Quitacoma(FormatoFlotanteEntero(Flotante(x.Valor) - Flotante(x.Vresidual) - Flotante(x.Acumulado)))
		res2 = append(res2, x)

		// sumatoria
		totalsaldo = totalsaldo + Flotante(x.Saldo)
		totallibros = totallibros + Flotante(x.Libros)
		totalacumulado = totalacumulado + Flotante(x.Acumulado)
		totaltotalmes = totaltotalmes + Flotante(x.Totalmes)
		totalvresidual = totalvresidual + Flotante(x.Vresidual)
		totalvalor = totalvalor + Flotante(x.Valor)
	}

	res2 = append(res2, listaPropiedad{"", "TOTAL", time.Now(), "", Quitacoma(FormatoFlotante(totalvalor)), Quitacoma(FormatoFlotante(totalvresidual)), Quitacoma(FormatoFlotante(totaltotalmes)), Quitacoma(FormatoFlotante(totalacumulado)), Quitacoma(FormatoFlotante(totalsaldo)), Quitacoma(FormatoFlotante(totallibros)), ""})
	//valor =

	if err != nil {
		panic(err.Error())
	}
	varmap := map[string]interface{}{
		"res":     res2,
		"hosting": ruta,
	}
	tmp.Execute(w, varmap)
}

// PROPIEDAD NUEVO
func PropiedadNuevo(w http.ResponseWriter, r *http.Request) {
	// TRAER COPIA DE EDITAR
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	emp := propiedad{}
	if Codigo == "False" {
	} else {
		err := db.Get(&emp, "SELECT * FROM propiedad where codigo=$1", Codigo)
		if err != nil {
			log.Fatalln(err)
		}

		emp.TerceroDetalle = TraerTerceroConsulta(emp.Tercero)
		emp.CentroDetalle = TraerCentroConsulta(emp.Centro)
		emp.CuentaDetalle = TraerCuentaConsulta(emp.Cuenta)
		emp.CuentagastoDetalle = TraerCuentaConsulta(emp.Cuentagasto)
		emp.CuentacontraDetalle = TraerCuentaConsulta(emp.Cuentacontra)
	}

	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/propiedad/propiedadNuevo.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaPlandecuentaempresa.html",
		"vista/autocompleta/autocompletaTercero.html")

	parametros := map[string]interface{}{
		"codigo":    Codigo,
		"emp":       emp,
		"hosting":   ruta,
		"cuenta":    ListaCuentaAuxiliar(),
		"parametro": TraerParametrosContabilidad(),
		"centro":    ListaCentro(),
	}
	tmp.Execute(w, parametros)
	// TERMINA TRAER COPIA DE EDITAR
}

// generar mes
func PropiedadGenerar(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/propiedad/propiedadGenerar.html")

	varmap := map[string]interface{}{
		"hosting": ruta,
		"centro":  ListaCentro(),
	}
	tmp.Execute(w, varmap)
}

func PropiedadMes(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	mes := mux.Vars(r)["mes"]
	miResultado := Propiedadresultado{}
	log.Println("Generar mes 1")
	var miTercero string
	var Documentocontable string
	Documentocontable = "22"
	var NumeroComprobante = NumeroDocumento(Documentocontable)

	miResultado.Numero = NumeroComprobante
	miResultado.Documento = DocumentoNombre(Documentocontable)

	//	*mes
	var fechaString string
	var miperiodo = periodoSesion(r)
	fechaString = fechaInicial(miperiodo, mes)
	const (
		layoutISO = "2006-01-02"
	)
	fechaDate, _ := time.Parse(layoutISO, fechaString)

	var totalDebito float64
	var totalCredito float64
	totalDebito = 0
	totalCredito = 0
	log.Println("Generar mes 2")
	// BORRA MOVIMIENTOS
	//var consultaborracomprobante = "delete from comprobante where documento=$1 and EXTRACT(MONTH FROM  fecha)>=$2 AND EXTRACT(YEAR FROM  fecha)=$3"
	//db.Exec(consultaborracomprobante, Documentocontable, mes, miperiodo)
	//
	//var consultaborracomprobantedetalle = "delete from comprobantedetalle where documento=$1 and  EXTRACT(MONTH FROM  fecha)>=$2 AND EXTRACT(YEAR FROM  fecha)=$3"
	//db.Exec(consultaborracomprobantedetalle, Documentocontable, mes, miperiodo)

	// borra datos anteriores
	listadopropiedad := []propiedad{}
	miComprobanteDetalle := []comprobantedetalle{}

	var consultaborra = "select * from propiedad where  $1 >=EXTRACT(MONTH FROM  inicia) order by codigo"
	db.Select(&listadopropiedad, consultaborra, mes)
	var miFilaComprobante int
	miFilaComprobante = 0
	log.Println("Generar mes 3")
	for _, miPropiedad := range listadopropiedad {
		log.Println("Generar movimiento")
		// inserta fila cuenta
		miTercero = miPropiedad.Tercero
		miFilaComprobante++
		miComprobanteDetalle = append(miComprobanteDetalle,
			comprobantedetalle{strconv.Itoa(miFilaComprobante),
				miPropiedad.Cuentagasto,
				miTercero,
				miPropiedad.Centro,
				"Depreciacion Del Mes De " + mesLetras(mes),
				miPropiedad.Codigo,
				(miPropiedad.Totalmes) + ".00",
				"",
				Documentocontable,
				NumeroComprobante,
				fechaDate,
				fechaDate, "", ""})
		// Inserta Fila contra
		miFilaComprobante++
		miComprobanteDetalle = append(miComprobanteDetalle,
			comprobantedetalle{strconv.Itoa(miFilaComprobante),
				miPropiedad.Cuentacontra,
				miTercero,
				miPropiedad.Centro,
				"Depreciacion Del Mes De " + mesLetras(mes),
				miPropiedad.Codigo,
				"",
				(miPropiedad.Totalmes) + ".00",
				Documentocontable,
				NumeroComprobante,
				fechaDate,
				fechaDate, "", ""})

		totalDebito += Flotante(miPropiedad.Totalmes)
		totalCredito += Flotante(miPropiedad.Totalmes)

	}
	log.Println("Generar mes 4")

	// crea comprobante
	if totalDebito > 0 {

		ComprobanteAgregarGenerar(comprobante{Documentocontable,
			NumeroComprobante, fechaDate,
			fechaDate,
			miperiodo,
			"",
			"",
			"",
			FormatoFlotante(totalDebito),
			FormatoFlotante(totalCredito),
			"Actualizar",
			miComprobanteDetalle, nil})

	}
	//if simueve == false {
	//var slice []string
	//slice = make([]string, 0)
	data, _ := json.Marshal(miResultado)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// PROPIEDAD INSERTAR
func PropiedadInsertar(w http.ResponseWriter, r *http.Request) {

	db := dbConn()

	t := propiedad{}
	r.ParseForm()
	defer r.Body.Close()
	decoder := schema.NewDecoder()
	decoder.RegisterConverter(time.Time{}, timeConverter)

	if err := decoder.Decode(&t, r.Form); err != nil {
		fmt.Println(err)
	}
	var q string

	q = "insert into propiedad ("
	q += "Codigo,"
	q += "Nombre,"
	q += "Cuenta,"
	q += "Ubicacion,"
	q += "Fecha,"
	q += "Inicia,"
	q += "Valor,"
	q += "Vresidual,"
	q += "Vidautil,"
	q += "Totalmes,"
	q += "Acumulado,"
	q += "Saldo,"
	q += "Libros,"
	q += "Cuentagasto,"
	q += "Cuentacontra,"
	q += "Centro,"
	q += "Tercero,"
	q += "Terceronombre"

	q += " ) values("
	q += parametros(18)
	q += " ) "

	log.Println("Cadena SQL " + q)
	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}
	_, err = insForm.Exec(
		t.Codigo,
		Titulo(t.Nombre),
		t.Cuenta,
		Titulo(t.Ubicacion),
		t.Fecha,
		t.Inicia,
		Quitacoma(t.Valor),
		Quitacoma(t.Vresidual),
		Quitacoma(t.Vidautil),
		Quitacoma(t.Totalmes),
		Quitacoma(t.Acumulado),
		Quitacoma(t.Saldo),
		Quitacoma(t.Libros),
		t.Cuentagasto,
		t.Cuentacontra,
		t.Centro,
		t.Tercero,
		t.Terceronombre)

	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/PropiedadLista", 301)
}

// PROPIEDAD EXISTE
func PropiedadExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM propiedad  WHERE codigo=$1", Codigo)
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

// PROPIEDAD EDITAR
func PropiedadEditar(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/propiedad/propiedadEditar.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/autocompleta/autocompletaPlandecuentaempresa.html",
		"vista/autocompleta/autocompletaTercero.html")
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]

	//db := dbConn()
	emp := propiedad{}
	err := db.Get(&emp, "SELECT * FROM propiedad where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}

	emp.TerceroDetalle = TraerTerceroConsulta(emp.Tercero)
	emp.CentroDetalle = TraerCentroConsulta(emp.Centro)
	emp.CuentaDetalle = TraerCuentaConsulta(emp.Cuenta)
	emp.CuentagastoDetalle = TraerCuentaConsulta(emp.Cuentagasto)
	emp.CuentacontraDetalle = TraerCuentaConsulta(emp.Cuentacontra)

	emp.Acumulado = SumarPropiedad(emp.Cuentagasto, emp.Codigo)
	log.Println("suma propiedad acumulado" + emp.Acumulado)
	emp.Libros = FormatoFlotanteEntero(Flotante(emp.Valor) - Flotante(emp.Acumulado))
	emp.Saldo = FormatoFlotanteEntero(Flotante(emp.Valor) - Flotante(emp.Vresidual) - Flotante(emp.Acumulado))
	varmap := map[string]interface{}{
		"emp":       emp,
		"hosting":   ruta,
		"cuenta":    ListaCuentaAuxiliar(),
		"parametro": TraerParametrosContabilidad(),
		"centro":    ListaCentro(),
	}
	//vistaPropiedad.ExecuteTemplate(w, "PropiedadEditar", varmap)
	tmp.Execute(w, varmap)
}

// PROPIEDAD ACTUALIZAR
func PropiedadActualizar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	t := propiedad{}
	r.ParseForm()
	defer r.Body.Close()
	decoder := schema.NewDecoder()
	decoder.RegisterConverter(time.Time{}, timeConverter)

	if err := decoder.Decode(&t, r.Form); err != nil {
		fmt.Println(err)
	}
	var q string
	q = "UPDATE propiedad set "
	q += " Nombre=$2,"
	q += " Cuenta=$3,"
	q += " Ubicacion=$4,"
	q += " Fecha=$5,"
	q += " Inicia=$6,"
	q += " Valor=$7,"
	q += " Vresidual=$8,"
	q += " Vidautil=$9,"
	q += " Totalmes=$10,"
	q += " Acumulado=$11,"
	q += " Saldo=$12,"
	q += " Libros=$13,"
	q += " Cuentagasto=$14,"
	q += " Cuentacontra=$15,"
	q += " Centro=$16,"
	q += " Tercero=$17,"
	q += " Terceronombre=$18"
	q += " where codigo=$1"

	log.Println("cadena" + q)

	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	// TERMINA GRABAR TERCERO ACTUALIZAR

	_, err = insForm.Exec(
		t.Codigo,
		Titulo(t.Nombre),
		t.Cuenta,
		Titulo(t.Ubicacion),
		t.Fecha,
		t.Inicia,
		Quitacoma(t.Valor),
		Quitacoma(t.Vresidual),
		Quitacoma(t.Vidautil),
		Quitacoma(t.Totalmes),
		Quitacoma(t.Acumulado),
		Quitacoma(t.Saldo),
		Quitacoma(t.Libros),
		t.Cuentagasto,
		t.Cuentacontra,
		t.Centro,
		t.Tercero,
		t.Terceronombre)

	if err != nil {
		panic(err)
	}

	http.Redirect(w, r, "/PropiedadLista", 301)
}

// PROPIEDAD BORRAR
func PropiedadBorrar(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/propiedad/propiedadBorrar.html")
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]

	//db := dbConn()
	emp := propiedad{}
	err := db.Get(&emp, "SELECT * FROM propiedad where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}

	emp.TerceroDetalle = TraerTerceroConsulta(emp.Tercero)
	emp.CentroDetalle = TraerCentroConsulta(emp.Centro)
	emp.CuentaDetalle = TraerCuentaConsulta(emp.Cuenta)
	emp.CuentagastoDetalle = TraerCuentaConsulta(emp.Cuentagasto)
	emp.CuentacontraDetalle = TraerCuentaConsulta(emp.Cuentacontra)

	emp.Acumulado = SumarPropiedad(emp.Cuentagasto, emp.Codigo)
	log.Println("suma propiedad acumulado" + emp.Acumulado)
	emp.Libros = FormatoFlotanteEntero(Flotante(emp.Valor) - Flotante(emp.Acumulado))
	emp.Saldo = FormatoFlotanteEntero(Flotante(emp.Valor) - Flotante(emp.Vresidual) - Flotante(emp.Acumulado))
	varmap := map[string]interface{}{

		"emp":       emp,
		"hosting":   ruta,
		"cuenta":    ListaCuentaAuxiliar(),
		"parametro": TraerParametrosContabilidad(),
		"centro":    ListaCentro(),
	}
	tmp.Execute(w, varmap)
}

// PROPIEDAD ELIMINAR
func PropiedadEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	emp := mux.Vars(r)["codigo"]
	//Codigo, _ := strconv.ParseInt(emp, 10, 0)
	delForm, err := db.Prepare("DELETE from propiedad WHERE codigo=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(emp)
	log.Println("Registro Eliminado" + emp)
	http.Redirect(w, r, "/PropiedadLista", 301)
}

// PROPIEDAD ACTUAL
func PropiedadActual(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	selDB, err := db.Query("SELECT * FROM propiedad where codigo=$1", Codigo)
	if err != nil {
		panic(err.Error())
	}
	emp := propiedad{}
	var res []propiedad
	for selDB.Next() {
		var codigo string
		var nombre string
		var Cuenta string
		var Ubicacion string
		var Fecha time.Time
		var Inicia time.Time
		var Valor string
		var Vresidual string
		var Vidautil string
		var Totalmes string
		var Acumulado string
		var Saldo string
		var Libros string
		var Cuentagasto string
		var Cuentacontra string
		var Centro string
		var Tercero string
		var Terceronombre string

		err = selDB.Scan(&codigo, &nombre, &Cuenta, &Ubicacion, &Fecha, &Inicia,
			&Valor, &Vresidual, &Vidautil, &Totalmes, &Acumulado, &Saldo, &Libros,
			Cuentagasto, &Cuentacontra, &Centro, &Tercero, &Terceronombre)
		if err != nil {
			panic(err.Error())
		}
		emp.Codigo = codigo
		emp.Nombre = nombre
		emp.Cuenta = Cuenta
		emp.Ubicacion = Ubicacion
		emp.Fecha = Fecha
		emp.Inicia = Inicia
		emp.Valor = Valor
		emp.Vresidual = Vresidual
		emp.Vidautil = Vidautil
		emp.Totalmes = Totalmes
		emp.Acumulado = Acumulado
		emp.Saldo = Saldo
		emp.Libros = Libros
		emp.Cuentagasto = Cuentagasto
		emp.Cuentacontra = Cuentacontra
		emp.Centro = Centro
		emp.Tercero = Tercero
		emp.Terceronombre = Terceronombre

		res = append(res, emp)
	}
	if err := selDB.Err(); err != nil { // make sure that there was no issue during the process
		log.Println(err)
		return
	}
	data, err := json.Marshal(res)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

// INICIA PROPIEDAD PDF
func PropiedadPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	t := propiedad{}
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	err := db.Get(&t, "SELECT * FROM propiedad where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	t.Acumulado = SumarPropiedad(t.Cuentagasto, t.Codigo)
	log.Println("suma propiedad acumulado" + t.Acumulado)
	t.Libros = FormatoFlotanteEntero(Flotante(t.Valor) - Flotante(t.Acumulado))
	t.Saldo = FormatoFlotanteEntero(Flotante(t.Valor) - Flotante(t.Vresidual) - Flotante(t.Acumulado))
	var buf bytes.Buffer
	var err1 error
	pdf := gofpdf.New("P", "mm", "Letter", cnFontDir)
	ene := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetHeaderFunc(func() {
		pdf.Image(imageFile("logo.png"), 20, 20, 40, 0, false,
			"", 0, "")
		pdf.SetY(15)
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
		pdf.CellFormat(190, 10, e.Telefono1+"  "+e.Telefono2, "0", 0, "C", false, 0,
			"")
		pdf.Ln(4)
		pdf.CellFormat(190, 10, ene(c.NombreCiudad+" - "+c.NombreDepartamento), "0", 0, "C", false, 0,
			"")
		pdf.Ln(6)

		// RELLENO TITULO
		pdf.SetY(50)
		pdf.SetFillColor(224, 231, 239)
		pdf.SetTextColor(0, 0, 0)
		pdf.SetX(20)
		pdf.CellFormat(184, 6, "PROPIEDAD PLANTA Y EQUIPO", "0", 0,
			"C", true, 0, "")
		pdf.Ln(8)
	})
	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Codigo", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Codigo, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Nombre:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Nombre, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Cuenta:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Cuenta, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Ubicacion:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Ubicacion, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Centro:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Centro+" - "+TraerCentro(t.Centro), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Fecha:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Fecha.Format("02/01/2006"), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Inicia:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Inicia.Format("02/01/2006"), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Valor:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Coma(t.Valor), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "V. Residual:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Coma(t.Vresidual), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Meses:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Coma(t.Vidautil), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Total Mes:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Coma(t.Totalmes), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Acumulado:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Coma(t.Acumulado), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Saldo:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Coma(t.Saldo), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Libros:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Coma(t.Libros), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Proveedor:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, Coma(t.Tercero)+"  "+t.Terceronombre, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Cuenta Gasto:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Cuentagasto, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Cuenta Contra:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Cuentacontra, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)

	pdf.SetFooterFunc(func() {
		pdf.SetTextColor(0, 0, 0)
		pdf.SetY(253)
		pdf.SetFont("Arial", "", 9)
		pdf.SetX(20)
		// LINEA
		pdf.Line(20, 260, 205, 260)
		pdf.Ln(6)
		pdf.SetX(20)
		pdf.CellFormat(90, 10, "Sadconf.com", "", 0,
			"L", false, 0, "")
		pdf.SetX(130)
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

// TERMINA PROPIEDAD PDF

// INICIA PROPIEDAD TODOS PDF
func PropiedadTodosCabecera(pdf *gofpdf.Fpdf) {
	pdf.SetFont("Arial", "", 10)
	// RELLENO TITULO
	pdf.SetY(50)
	pdf.SetFillColor(224, 231, 239)
	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(6)
	pdf.SetX(15)
	pdf.CellFormat(251, 6, "No", "0", 0,
		"L", true, 0, "")
	pdf.SetX(25)
	pdf.CellFormat(190, 6, "Codigo", "0", 0,
		"L", false, 0, "")
	pdf.SetX(45)
	pdf.CellFormat(190, 6, "Nombre", "0", 0,
		"L", false, 0, "")
	pdf.SetX(95)
	pdf.CellFormat(190, 6, "Fecha", "0", 0,
		"L", false, 0, "")
	pdf.SetX(114)
	pdf.CellFormat(190, 6, "Meses", "0", 0,
		"L", false, 0, "")
	pdf.SetX(132)
	pdf.CellFormat(190, 6, "Valor", "0", 0,
		"L", false, 0, "")
	pdf.SetX(147)
	pdf.CellFormat(190, 6, "V. Residual", "0", 0,
		"L", false, 0, "")
	pdf.SetX(174)
	pdf.CellFormat(190, 6, "Valor Mes", "0", 0,
		"L", false, 0, "")
	pdf.SetX(197)
	pdf.CellFormat(190, 6, "Acumulado", "0", 0,
		"L", false, 0, "")
	pdf.SetX(230)
	pdf.CellFormat(190, 6, "Saldo", "0", 0,
		"L", false, 0, "")
	pdf.SetX(255)
	pdf.CellFormat(190, 6, "Libros", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}

func PropiedadTodosDetalle(pdf *gofpdf.Fpdf, miFila propiedad, a int) {

	pdf.SetFont("Arial", "", 9)

	pdf.SetX(15)
	pdf.CellFormat(180, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(25)
	pdf.CellFormat(40, 4, miFila.Codigo, "", 0,
		"L", false, 0, "")
	pdf.SetX(45)
	pdf.CellFormat(40, 4, miFila.Nombre, "", 0, "L", false, 0, "")
	pdf.SetX(95)
	pdf.CellFormat(20, 4, miFila.Fecha.Format("02/01/2006"), "", 0, "L", false, 0, "")
	pdf.SetX(102)
	pdf.CellFormat(20, 4, Coma(miFila.Vidautil), "", 0, "R", false, 0, "")
	pdf.SetX(122)
	pdf.CellFormat(20, 4, Coma(miFila.Valor), "", 0, "R", false, 0, "")
	pdf.SetX(147)
	pdf.CellFormat(20, 4, Coma(miFila.Vresidual), "", 0, "R", false, 0, "")
	pdf.SetX(172)
	pdf.CellFormat(20, 4, Coma(miFila.Totalmes), "", 0, "R", false, 0, "")
	pdf.SetX(197)

	miFila.Acumulado = SumarPropiedad(miFila.Cuentagasto, miFila.Codigo)
	log.Println("suma propiedad acumulado" + miFila.Acumulado)
	miFila.Libros = FormatoFlotanteEntero(Flotante(miFila.Valor) - Flotante(miFila.Acumulado))
	miFila.Saldo = FormatoFlotanteEntero(Flotante(miFila.Valor) - Flotante(miFila.Vresidual) - Flotante(miFila.Acumulado))

	pdf.CellFormat(20, 4, miFila.Acumulado, "", 0, "R", false, 0, "")
	pdf.SetX(222)
	pdf.CellFormat(20, 4, miFila.Saldo, "", 0, "R", false, 0, "")
	pdf.SetX(247)
	pdf.CellFormat(20, 4, miFila.Libros, "", 0, "R", false, 0, "")
	pdf.Ln(4)
}

func PropiedadTodosDetalleTotal(pdf *gofpdf.Fpdf, miFila listaPropiedad, a int) {
	pdf.SetFillColor(224, 231, 239)
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(15)
	pdf.CellFormat(251, 5, "", "", 0,
		"L", true, 0, "")
	pdf.SetX(33)
	pdf.SetX(55)
	pdf.CellFormat(40, 4, miFila.Nombre, "", 0, "L", false, 0, "")
	pdf.SetX(114)
	pdf.SetX(124)
	pdf.SetX(122)
	pdf.CellFormat(20, 4, Coma(miFila.Valor), "", 0, "R", false, 0, "")
	pdf.SetX(147)
	pdf.CellFormat(20, 4, Coma(miFila.Vresidual), "", 0, "R", false, 0, "")
	pdf.SetX(172)
	pdf.CellFormat(20, 4, Coma(miFila.Totalmes), "", 0, "R", false, 0, "")
	pdf.SetX(197)
	log.Println("suma propiedad acumulado" + miFila.Acumulado)
	pdf.CellFormat(20, 4, miFila.Acumulado, "", 0, "R", false, 0, "")
	pdf.SetX(222)
	pdf.CellFormat(20, 4, miFila.Saldo, "", 0, "R", false, 0, "")
	pdf.SetX(247)
	pdf.CellFormat(20, 4, miFila.Libros, "", 0, "R", false, 0, "")
	pdf.Ln(4)
}

func PropiedadTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	//	Codigo := mux.Vars(r)["codigo"]

	t := []propiedad{}
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	err := db.Select(&t, "SELECT * FROM propiedad ORDER BY cast(codigo as integer) ")
	if err != nil {
		log.Fatalln(err)
	}
	var buf bytes.Buffer
	var err1 error
	pdf := gofpdf.New("L", "mm", "Letter", cnFontDir)
	ene := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetHeaderFunc(func() {
		pdf.Image(imageFile("logo.png"), 20, 20, 40, 0, false,
			"", 0, "")
		pdf.SetY(15)
		pdf.SetX(55)
		pdf.SetFont("Arial", "", 10)
		pdf.CellFormat(190, 10, e.Nombre, "0", 0,
			"C", false, 0, "")
		pdf.Ln(4)
		pdf.SetX(55)
		pdf.CellFormat(190, 10, "Nit No. "+Coma(e.Codigo)+" - "+e.Dv, "0", 0, "C",
			false, 0, "")
		pdf.Ln(4)
		pdf.SetX(55)
		pdf.CellFormat(190, 10, e.Iva+" - "+e.ReteIva, "0", 0, "C", false, 0,
			"")
		pdf.Ln(4)
		pdf.SetX(55)
		pdf.CellFormat(190, 10, "Actividad Ica - "+e.ActividadIca, "0",
			0, "C", false, 0, "")
		pdf.Ln(4)
		pdf.SetX(55)
		pdf.CellFormat(190, 10, e.Direccion, "0", 0, "C", false, 0,
			"")
		pdf.Ln(4)
		pdf.SetX(55)
		pdf.CellFormat(190, 10, e.Telefono1+" "+e.Telefono2, "0", 0, "C", false, 0,
			"")
		pdf.Ln(4)
		pdf.SetX(55)
		pdf.CellFormat(190, 10, ene(c.NombreCiudad+" - "+c.NombreDepartamento), "0", 0, "C", false, 0,
			"")
		pdf.Ln(6)
		pdf.SetX(55)
		pdf.CellFormat(190, 10, "DATOS PROPIEDAD, PLANTA Y EQUIPO", "0", 0,
			"C", false, 0, "")
		pdf.Ln(10)
	})

	pdf.SetFooterFunc(func() {
		pdf.SetFont("Arial", "", 9)
		pdf.SetTextColor(0, 0, 0)
		pdf.SetY(192)
		pdf.SetX(23)
		// LINEA LARGA
		pdf.Line(23, 198, 268, 198)
		pdf.Ln(6)
		pdf.SetY(198)
		pdf.SetX(23)
		pdf.CellFormat(90, 10, "Sadconf.com", "", 0,
			"L", false, 0, "")
		pdf.CellFormat(161, 10, fmt.Sprintf(" %d de {nb}", pdf.PageNo()), "",
			0, "R", false, 0, "")
	})

	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)
	pdf.SetX(30)

	var totalsaldo float64
	totalsaldo = 0

	var totallibros float64
	totallibros = 0

	var totalacumulado float64
	totalacumulado = 0

	var totaltotalmes float64
	totaltotalmes = 0

	var totalvresidual float64
	totalvresidual = 0

	var totalvalor float64
	totalvalor = 0

	PropiedadTodosCabecera(pdf)
	// tercera pagina
	var a = 1
	for i, miFila := range t {
		var a = i + 1
		if math.Mod(float64(a), 49) == 0 {
			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			DocumentoTodosCabecera(pdf)
		}
		miFila.Acumulado = SumarPropiedad(miFila.Cuentagasto, miFila.Codigo)
		miFila.Libros = FormatoFlotanteEntero(Flotante(miFila.Valor) - Flotante(miFila.Acumulado))
		miFila.Saldo = FormatoFlotanteEntero(Flotante(miFila.Valor) - Flotante(miFila.Vresidual) - Flotante(miFila.Acumulado))
		// sumatoria
		totalsaldo = totalsaldo + Flotante(miFila.Saldo)
		totallibros = totallibros + Flotante(miFila.Libros)
		totalacumulado = totalacumulado + Flotante(miFila.Acumulado)
		totaltotalmes = totaltotalmes + Flotante(miFila.Totalmes)
		totalvresidual = totalvresidual + Flotante(miFila.Vresidual)
		totalvalor = totalvalor + Flotante(miFila.Valor)

		PropiedadTodosDetalle(pdf, miFila, a)
	}
	totales := listaPropiedad{"", "TOTALES", time.Now(), "", FormatoFlotanteEntero(totalvalor), FormatoFlotanteEntero(totalvresidual), FormatoFlotanteEntero(totaltotalmes), FormatoFlotanteEntero(totalacumulado), FormatoFlotanteEntero(totalsaldo), FormatoFlotanteEntero(totallibros), ""}
	PropiedadTodosDetalleTotal(pdf, totales, a)

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

// TERMINA PROPIEDAD TODOS PDF

// PROPIEDAD EXCEL
func PropiedadExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	t := []propiedad{}
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	err := db.Select(&t, "SELECT * FROM propiedad ORDER BY cast(codigo as integer) ")
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
	if err = f.SetColWidth("Sheet1", "C", "C", 10); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.SetColWidth("Sheet1", "D", "D", 10); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.SetColWidth("Sheet1", "E", "E", 10); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.SetColWidth("Sheet1", "F", "F", 10); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.SetColWidth("Sheet1", "G", "G", 10); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.SetColWidth("Sheet1", "H", "H", 10); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.SetColWidth("Sheet1", "I", "I", 10); err != nil {
		fmt.Println(err)
		return
	}

	// FUNCION PARA UNIR DOS CELDAS TITULO
	if err = f.MergeCell("Sheet1", "A1", "I1"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A2", "I2"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A3", "I3"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A4", "I4"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A5", "I5"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A6", "I6"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A7", "I7"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A8", "I8"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A9", "I9"); err != nil {
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
	f.SetCellValue("Sheet1", "A9", "PROPIEDAD PLANTA Y EQUIPO")
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
	f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel), "Fecha")
	f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel), "Meses")
	f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel), "Valor")
	f.SetCellValue("Sheet1", "F"+strconv.Itoa(filaExcel), "V. Residual")
	f.SetCellValue("Sheet1", "G"+strconv.Itoa(filaExcel), "Valor Mes")
	f.SetCellValue("Sheet1", "H"+strconv.Itoa(filaExcel), "Acumulado")
	f.SetCellValue("Sheet1", "I"+strconv.Itoa(filaExcel), "Saldo")
	f.SetCellValue("Sheet1", "J"+strconv.Itoa(filaExcel), "Libros")

	f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel), "A"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel), "B"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel), "C"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel), "D"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "E"+strconv.Itoa(filaExcel), "E"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "F"+strconv.Itoa(filaExcel), "F"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "G"+strconv.Itoa(filaExcel), "G"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "H"+strconv.Itoa(filaExcel), "H"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "I"+strconv.Itoa(filaExcel), "I"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "J"+strconv.Itoa(filaExcel), "J"+strconv.Itoa(filaExcel), estiloCabecera)
	filaExcel++

	for i, miFila := range t {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel+i), miFila.Codigo)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel+i), miFila.Nombre)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel+i), miFila.Fecha.Format("02/01/2006"))
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel+i), Flotante(miFila.Vidautil))
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel+i), Flotante(miFila.Valor))
		f.SetCellValue("Sheet1", "F"+strconv.Itoa(filaExcel+i), Flotante(miFila.Vresidual))
		f.SetCellValue("Sheet1", "G"+strconv.Itoa(filaExcel+i), Flotante(miFila.Totalmes))
		f.SetCellValue("Sheet1", "H"+strconv.Itoa(filaExcel+i), Flotante(miFila.Acumulado))
		f.SetCellValue("Sheet1", "I"+strconv.Itoa(filaExcel+i), (Flotante(miFila.Valor) - Flotante(miFila.Acumulado)))
		f.SetCellValue("Sheet1", "J"+strconv.Itoa(filaExcel+i), (Flotante(miFila.Valor) - Flotante(miFila.Vresidual) - Flotante(miFila.Acumulado)))

		f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel+i), "A"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel+i), "B"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel+i), "B"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel+i), "B"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "E"+strconv.Itoa(filaExcel+i), "B"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "F"+strconv.Itoa(filaExcel+i), "B"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "G"+strconv.Itoa(filaExcel+i), "B"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "H"+strconv.Itoa(filaExcel+i), "B"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "I"+strconv.Itoa(filaExcel+i), "B"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "J"+strconv.Itoa(filaExcel+i), "B"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		//van=i
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

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
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type ResultadoBorrar struct {
	Siborrar bool   `json:"Siborrar"`
	Mensaje  string `json:"Mensaje"`
}

// CUENTA JSON
type plandecuentaempresaJson struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Nombre string `json:"nombre"`
}

// CUENTA ESTRUCTURA
type Resultado struct {
	NivelAnterior bool `json:"NivelAnterior"`
	ExisteCuenta  bool `json:"ExisteCuenta"`
	Resultado     bool `json:"Resultado"`
}

type plandecuentaempresaLista struct {
	Codigo string
	Nombre string
	Auto   string
	Nivel  string
}

// CENTRO TABLA
type plandecuentaempresa struct {
	Codigo          string
	Nombre          string
	Auto            string
	Nivel           string
	Tercero         string
	Centro          string
	Factura         string
	Financiero      string
	Contra          string
	Interes         string
	Cuota           string
	Cuentaintereses string
	Tipo            string
	Grupo           string
	ContraNombre    string
	Nombreintereses string
}

// CENTRO LISTA
func PlandecuentaempresaLista(w http.ResponseWriter, r *http.Request) {
	panel := mux.Vars(r)["panel"]
	codigo := mux.Vars(r)["codigo"]
	elemento := mux.Vars(r)["elemento"]

	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/plandecuentaempresa/plandecuentaempresaLista.html")
	db := dbConn()
	var consulta = " select codigo,nombre,auto, nivel from plandecuentaempresa order by codigo"
	////+
	//	"  union" +
	///	"   select codigo,nombre,auto, nivel from cuenta order by codigo ;"
	selDB, err := db.Query(consulta)
	if err != nil {
		panic(err.Error())
	}
	res := []plandecuentaempresaLista{}
	for selDB.Next() {
		var Codigo string
		var Nombre string
		var Auto string
		var Nivel string
		err = selDB.Scan(&Codigo, &Nombre, &Auto, &Nivel)
		if err != nil {
			panic(err.Error())
		}
		res = append(res, plandecuentaempresaLista{Codigo, Nombre, Auto, Nivel})
	}
	varmap := map[string]interface{}{
		"res":      res,
		"hosting":  ruta,
		"panel":    panel,
		"codigo":   codigo,
		"elemento": elemento,
		"modulo":   ListaEmpresa().Modulo,
	}
	tmp.Execute(w, varmap)
}

// CUENTA BUSCAR
func PlandecuentaempresaBuscar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	selDB, err := db.Query("SELECT codigo,"+
		"nombre FROM plandecuentaempresa where (codigo LIKE '%' || $1 || '%') or (Upper(nombre) LIKE '%' || $1 || '%') ORDER BY"+
		" codigo ", Mayuscula(Codigo))
	if err != nil {
		panic(err.Error())
	}
	var resJson []plandecuentaempresaJson
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
		resJson = append(resJson, plandecuentaempresaJson{id, label, value, nombre})
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

//FUNCION EDITAR REGISTROS//
func PlandecuentaempresaEditar(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/autocompleta/autocompletaPlandecuentaempresa."+
			"html",
		"vista/plandecuentaempresa/plandecuentaempresaEditar.html")
	Codigo := mux.Vars(r)["codigo"]
	db := dbConn()
	emp := plandecuentaempresa{}
	err := db.Get(&emp, "SELECT * FROM plandecuentaempresa where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}

	varmap := map[string]interface{}{
		"emp":          emp,
		"hosting":      ruta,
		"financiero":   ListaFinanciero(),
		"nombremodulo": NombreModulo(),
	}
	tmp.Execute(w, varmap)
}

// INICIA PLAN DE CUENTA EMPRESA NUEVO
func PlandecuentaempresaNuevo(w http.ResponseWriter, r *http.Request) {
	Codigo := mux.Vars(r)["codigo"]
	Panel := mux.Vars(r)["panel"]
	Elemento := mux.Vars(r)["elemento"]

	emp := plandecuentaempresa{}
	parametros := map[string]interface{}{
		"Codigo":       Codigo,
		"Panel":        Panel,
		"Elemento":     Elemento,
		"hosting":      ruta,
		"financiero":   ListaFinanciero(),
		"nombremodulo": NombreModulo(),
		"copiar":       "False",
		"emp":          emp,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/autocompleta/autocompletaPlandecuentaempresa.html", "vista/plandecuentaempresa/autocompletaPlandecuentaempresacrear.html",
		"vista/plandecuentaempresa/PlandecuentaempresaNuevo.html")
	//fmt.Printf("%v, %v", t, err)
	fmt.Printf("%v, %v", miTemplate, err)
	miTemplate.Execute(w, parametros)
}

// INICIA PLAN DE CUENTA EMPRESA DUPLICAR
func PlandecuentaempresaNuevoCopia(w http.ResponseWriter, r *http.Request) {
	Codigo := "False"
	Panel := "False"
	Elemento := "False"

	copiaCodigo := mux.Vars(r)["copiacodigo"]
	db := dbConn()
	emp := plandecuentaempresa{}
	err := db.Get(&emp, "SELECT * FROM plandecuentaempresa where codigo=$1", copiaCodigo)
	if err != nil {
		log.Fatalln(err)
	}

	parametros := map[string]interface{}{
		"emp":          emp,
		"hosting":      ruta,
		"financiero":   ListaFinanciero(),
		"nombremodulo": NombreModulo(),
		"copiar":       "True",
		"Codigo":       Codigo,
		"Panel":        Panel,
		"Elemento":     Elemento,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/autocompleta/autocompletaPlandecuentaempresa.html",
		"vista/plandecuentaempresa/PlandecuentaempresaNuevo.html",
		"vista/plandecuentaempresa/autocompletaPlandecuentaempresacrear.html")
	//fmt.Printf("%v, %v", t, err)
	fmt.Printf("%v, %v", miTemplate, err)
	miTemplate.Execute(w, parametros)
}

// CUENTA INSERTAR
func PlandecuentaempresaInsertar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		Codigo := r.FormValue("Codigo")
		Nombre := r.FormValue("Nombre")
		Auto := r.FormValue("Auto")
		Nivel := r.FormValue("Nivel")
		Tercero := r.FormValue("Tercero")
		Centro := r.FormValue("Centro")
		Factura := r.FormValue("Factura")
		Financiero := r.FormValue("Financiero")
		Contra := r.FormValue("Contra")
		Interes := r.FormValue("Interes")
		Cuota := r.FormValue("Cuota")
		Cuentaintereses := r.FormValue("Cuentaintereses")
		Tipo := r.FormValue("Tipo")
		Grupo := r.FormValue("Grupo")
		Contranombre := r.FormValue("Contranombre")
		Nombreintereses := r.FormValue("Nombreintereses")
		Nombre = Titulo(Nombre)
		Auto = Mayuscula(Auto)
		Contranombre = Titulo(Contranombre)
		Nombreintereses = Titulo(Nombreintereses)

		log.Println("Datos Emvoadps:" + Codigo + "," + Nombre + "," +
			"" + Auto + "," + Nivel + "," + Tercero + "," + Centro + "," +
			"" + Factura + "," + Financiero + "," + Contra + "," +
			"" + Interes + "," + Cuota + "," + Cuentaintereses + "," +
			"" + Tipo + "," + Grupo + "," + Contranombre + "," +
			"" + Nombreintereses)
		var q = ""

		q += "INSERT INTO plandecuentaempresa("
		q += "codigo,"
		q += "nombre,"
		q += "auto,"
		q += "nivel,"
		q += "tercero,"
		q += "centro,"
		q += "factura,"
		q += "financiero,"
		q += "contra,"
		q += "interes,"
		q += "cuota,"
		q += "cuentaintereses,"
		q += "tipo,"
		q += "grupo,"
		q += "contranombre,"
		q += "nombreintereses)"
		q += "VALUES("
		q += parametros(16)
		q += ")"
		insForm, err := db.Prepare(q)

		if err != nil {
			panic(err.Error())
		}
		_, err1 := insForm.Exec(Codigo, Nombre, Auto, Nivel, Tercero, Centro, Factura,
			Financiero, Contra, Interes, Cuota, Cuentaintereses, Tipo, Grupo, Contranombre, Nombreintereses)
		if err1 != nil {
			panic(err1.Error())
		}

		log.Println("Registro Actualizado:" + Codigo + "," + Nombre + "," +
			"" + Auto + "," + Nivel + "," + Tercero + "," + Centro + "," +
			"" + Factura + "," + Financiero + "," + Contra + "," +
			"" + Interes + "," + Cuota + "," + Cuentaintereses + "," +
			"" + Tipo + "," + Grupo + "," + Contranombre + "," +
			"" + Nombreintereses)
	}
	http.Redirect(w, r, "/PlandecuentaempresaLista/False/False/False", 301)
}

// CUENTA EXISTE
func PlandecuentaempresaExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM plandecuentaempresa  WHERE codigo=$1", Codigo)
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

// CUENTA ACTUAL
func PlandecuentaempresaActual(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	selDB, err := db.Query("SELECT * FROM plandecuentaempresa where codigo=$1", Codigo)
	if err != nil {
		panic(err.Error())
	}
	emp := plandecuentaempresa{}
	var res []plandecuentaempresa
	for selDB.Next() {
		var codigo string
		var nombre string
		var auto string
		var nivel string
		var tercero string
		var centro string
		var factura string
		var financiero string
		var contra string
		var interes string
		var cuota string
		var cuentaintereses string
		var tipo string
		var grupo string
		var contranombre string
		var interesnombre string
		err = selDB.Scan(&codigo, &nombre, &auto, &nivel, &tercero,
			&centro, &factura, &financiero, &contra, &interes, &cuota,
			&cuentaintereses, &tipo, &grupo, &contranombre, &interesnombre)
		if err != nil {
			panic(err.Error())
		}
		emp.Codigo = codigo
		emp.Nombre = nombre
		emp.Auto = auto
		emp.Nivel = nivel
		emp.Tercero = tercero
		emp.Centro = centro
		emp.Factura = factura
		emp.Financiero = financiero
		emp.Contra = contra
		emp.Interes = interes
		emp.Cuota = cuota
		emp.Cuentaintereses = cuentaintereses
		emp.Tipo = tipo
		emp.Grupo = grupo
		emp.ContraNombre = contranombre
		emp.Nombreintereses = interesnombre
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

// CUENTA ACTUALIZAR
func PlandecuentaempresaActualizar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		codigo := r.FormValue("Codigo")
		nombre := r.FormValue("Nombre")
		auto := r.FormValue("Auto")
		nivel := r.FormValue("Nivel")
		tercero := r.FormValue("Tercero")
		centro := r.FormValue("Centro")
		factura := r.FormValue("Factura")
		financiero := r.FormValue("Financiero")
		contra := r.FormValue("Contra")
		interes := r.FormValue("Interes")
		cuota := r.FormValue("Cuota")
		cuentaintereses := r.FormValue("Cuentaintereses")
		tipo := r.FormValue("Tipo")
		grupo := r.FormValue("Grupo")
		contranombre := r.FormValue("ContraNombre")
		nombreintereses := r.FormValue("Nombreintereses")
		nombre = Titulo(nombre)
		auto = Mayuscula(auto)
		log.Println("Datos Emvoadps:" + codigo + "," + nombre + "," +
			"" + auto + "," + nivel + "," + tercero + "," + centro + "," +
			"" + factura + "," + financiero + "," + contra + "," +
			"" + interes + "," + cuota + "," + cuentaintereses + "," +
			"" + tipo + "," + grupo + "," + contranombre + "," +
			"" + nombreintereses)
		insForm, err := db.Prepare("UPDATE plandecuentaempresa set " +
			" nombre=$2, " +
			" auto=$3," +
			" nivel=$4," +
			" tercero=$5, " +
			" centro=$6," +
			" factura=$7," +
			" financiero=$8, " +
			" contra=$9," +
			" interes=$10," +
			" cuota=$11, " +
			" cuentaintereses=$12," +
			" tipo=$13," +
			" grupo=$14, " +
			" contranombre=$15," +
			" nombreintereses=$16" +
			" WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		_, err1 := insForm.Exec(codigo, nombre, auto, nivel, tercero, centro, factura,
			financiero, contra, interes, cuota, cuentaintereses, tipo, grupo, contranombre, nombreintereses)
		if err1 != nil {
			panic(err1.Error())
		}

		log.Println("Registro Actualizado:" + codigo + "," + nombre + "," +
			"" + auto + "," + nivel + "," + tercero + "," + centro + "," +
			"" + factura + "," + financiero + "," + contra + "," +
			"" + interes + "," + cuota + "," + cuentaintereses + "," +
			"" + tipo + "," + grupo + "," + contranombre + "," +
			"" + nombreintereses)
	}
	http.Redirect(w, r, "/PlandecuentaempresaLista/False/False/False", 301)
	//	http.Redirect(w, r, "/PlandecuentaempresaLista", 301)
}

// AGREGA CUENTAS DESDE EL PLAN NIIF
func PlandecuentaempresaAgregar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var miplandecuentasniif plandecuentaniif

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// carga informacion de la COMPRA
	err = json.Unmarshal(b, &miplandecuentasniif)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	var ancho = len(miplandecuentasniif.Codigo)
	var micuenta = miplandecuentasniif.Codigo
	var minivel = miplandecuentasniif.Nivel
	var minombre = miplandecuentasniif.Nombre
	var verificar string
	var nivelverificar string
	var siespecial bool

	switch ancho {
	case 6:
		verificar = Subcadena(micuenta, 0, 4)
		nivelverificar = "3"
		fmt.Println("nivel Auxiliar nivel 4" + verificar)
		fmt.Println("nivel Auxiliar nivel 4" + nivelverificar)
		siespecial = false
	case 4:
		verificar = Subcadena(micuenta, 0, 2)
		nivelverificar = "2"
		fmt.Println("nivel Auxiliar nivel 3" + verificar)
		fmt.Println("nivel Auxiliar nivel 3" + nivelverificar)
		siespecial = false
	case 2:
		verificar = Subcadena(micuenta, 0, 1)
		nivelverificar = "1"
		fmt.Println("cuenta2 Auxiliar" + verificar)
		fmt.Println("nivel Auxiliar" + nivelverificar)
		siespecial = false

	case 1:
		verificar = Subcadena(micuenta, 0, 1)
		nivelverificar = "1"
		fmt.Println("cuenta1 Auxiliar" + verificar)
		fmt.Println("nivel Auxiliar" + nivelverificar)
		siespecial = true
	default:
		fmt.Println("Too far away.")
	}

	var total int
	var sinivelanterior bool
	if siespecial == false {

		var total int
		row := db.QueryRow("SELECT COUNT(*) FROM plandecuentaempresa  WHERE codigo=$1 and nivel=$2", verificar, nivelverificar)
		err1 := row.Scan(&total)
		if err1 != nil {
			log.Fatal(err)
		}

		if total > 0 {
			sinivelanterior = true
		} else {
			sinivelanterior = false
		}
	} else {

		sinivelanterior = true
	}

	log.Println("nivelanterior ")
	log.Println(sinivelanterior)

	// si existe nivel anterior
	var existecuenta bool
	existecuenta = true

	if sinivelanterior == true {

		row := db.QueryRow("SELECT COUNT(*) FROM plandecuentaempresa WHERE codigo=$1 and nivel=$2", micuenta, minivel)
		err := row.Scan(&total)
		if err != nil {
			log.Fatal(err)
		}

		if total == 0 {
			existecuenta = false
		}

	}

	var resultado bool
	resultado = false

	if sinivelanterior == true && existecuenta == false {

		var q string
		q = "insert into plandecuentaempresa ("
		q += "Codigo,"
		q += "Nombre,"
		q += "Auto,"
		q += "Nivel,"
		q += "Tercero,"
		q += "Centro,"
		q += "Factura,"
		q += "Financiero,"
		q += "Contra,"
		q += "Interes,"
		q += "Cuota,"
		q += "Cuentaintereses,"
		q += "Tipo,"
		q += "Grupo,"
		q += "Contranombre,"
		q += "Nombreintereses"
		q += " ) values("
		q += parametros(16)
		q += " )"

		log.Println("Cadena SQL " + q)
		insForm, err := db.Prepare(q)
		if err != nil {
			panic(err.Error())
		}

		// TERMINA COMPRA GRABAR INSERTAR
		_, err = insForm.Exec(micuenta,
			minombre,
			"",
			minivel,
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"",
			"")

		if err != nil {
			panic(err)
		}

		log.Println("Insertar cuenta \n", micuenta)

		resultado = true
	}

	js, err := json.Marshal(Resultado{sinivelanterior, existecuenta, resultado})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	//http.Redirect(w, r, "/CompraLista", 301)
}

// CUENTAHORIZONTAL BORRAR
func PlandecuentaempresaBorrar(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/plandecuentaempresa/plandecuentaempresaBorrar.html")
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	selDB, err := db.Query("SELECT * FROM plandecuentaempresa WHERE codigo=$1", Codigo)
	if err != nil {
		panic(err.Error())
	}
	emp := plandecuentaempresa{}
	for selDB.Next() {
		var codigo string
		var nombre string
		var auto string
		var nivel string
		var tercero string
		var centro string
		var factura string
		var financiero string
		var contra string
		var interes string
		var cuota string
		var cuentaintereses string
		var tipo string
		var grupo string
		var contranombre string
		var nombreintereses string
		err = selDB.Scan(&codigo, &nombre, &auto, &nivel, &tercero,
			&centro, &factura, &financiero, &contra, &interes, &cuota,
			&cuentaintereses, &tipo, &grupo, &contranombre, &nombreintereses)
		if err != nil {
			panic(err.Error())
		}
		emp.Codigo = codigo
		emp.Nombre = nombre
		emp.Auto = auto
		emp.Nivel = nivel
		emp.Tercero = tercero
		emp.Centro = centro
		emp.Factura = factura
		emp.Financiero = financiero
		emp.Contra = contra
		emp.Interes = interes
		emp.Cuota = cuota
		emp.Cuentaintereses = cuentaintereses
		emp.Tipo = tipo
		emp.Grupo = grupo
		emp.ContraNombre = contranombre
		emp.Nombreintereses = nombreintereses
	}
	varmap := map[string]interface{}{
		"emp":          emp,
		"hosting":      ruta,
		"financiero":   ListaFinanciero(),
		"nombremodulo": NombreModulo(),
	}
	tmp.Execute(w, varmap)
}

// CUENTAHORIZONTAL ELIMINAR
func PlandecuentaempresaEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	emp := mux.Vars(r)["codigo"]

	//veeificar
	miCuetnaempresa := plandecuentaempresa{}
	//var res []tercero
	err := db.Get(&miCuetnaempresa, "SELECT * FROM tercero where codigo=$1", emp)

	var ancho = len(emp)
	var verificar string
	var siauxiliar bool
	siauxiliar = false

	switch ancho {
	case 8:
		// cuenta auxiliar
		siauxiliar = true

	case 6:
		//nivel 4
		verificar = "select count(*) from plandecuentaempresa where nivel='A' and SUBSTRING(codigo, 1, 8)=$1"
		fmt.Println("nivel Auxiliar nivel 6" + verificar)

	case 4:
		// nivel 3
		verificar = "select count(*) from plandecuentaempresa where nivel='4' and SUBSTRING(codigo, 1, 4)=$1"
		fmt.Println("nivel Auxiliar nivel 3" + verificar)
	case 2:
		verificar = "select count(*) from plandecuentaempresa where nivel='3' and SUBSTRING(codigo, 1, 2)=$1"
		fmt.Println("cuenta2 Auxiliar" + verificar)
	case 1:
		verificar = "select count(*) from plandecuentaempresa where nivel='2' and SUBSTRING(codigo, 1, 1)=$1"

		fmt.Println("cuenta1 Auxiliar" + verificar)
	default:
		fmt.Println("Too far away.")
	}

	var siborrar = false
	var mensaje string
	mensaje = "Borrado Correctamente"

	if siauxiliar == true {
		var total int
		row := db.QueryRow("SELECT COUNT(*) FROM comprobantedetalle  WHERE cuenta=$1", emp)
		err := row.Scan(&total)
		if err != nil {
			log.Fatal(err)
		}
		if total > 0 {
			siborrar = false
			mensaje = "Cuenta Auxiliar Con Movimiento"
		} else {
			siborrar = true
		}

	} else {
		var total int
		row := db.QueryRow(verificar, emp)
		err := row.Scan(&total)
		if err != nil {
			log.Fatal(err)
		}
		if total > 0 {
			mensaje = "Por Favor Borrar Cuentas Superiores"
			siborrar = false
		} else {
			siborrar = true
		}
	}

	if siborrar == true {
		delForm, err := db.Prepare("DELETE from plandecuentaempresa WHERE codigo=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(emp)
		log.Println("Registro Eliminado" + emp)
	}

	js, err := json.Marshal(ResultadoBorrar{siborrar, mensaje})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

}

// INICIA CUENTA BUSCAR
func PlandecuentaempresaBuscarAuxiliar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	Codigo = strings.ToLower(Codigo)
	selDB, err := db.Query(" select codigo,nombre from plandecuentaempresa"+
		"  where codigo LIKE '%' || $1 || '%'"+
		"  or  lower(nombre) LIKE '%' || $1 || '%' "+
		" ORDER BY"+
		" codigo ", Codigo)
	if err != nil {
		panic(err.Error())
	}
	var resJson []plandecuentaempresaJson
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
		resJson = append(resJson, plandecuentaempresaJson{id, label, value, nombre})
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

// INICIA PLAN DE CUENTA EMPRESA PDF
func PlandecuentaempresaPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	var miModulo = ListaEmpresa().Modulo
	Codigo := mux.Vars(r)["codigo"]
	t := plandecuentaempresa{}
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	err := db.Get(&t, "SELECT * FROM plandecuentaempresa where codigo=$1", Codigo)
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
		pdf.CellFormat(184, 6, "PLAN DE CUENTA EMPRESA", "0", 0,
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
	pdf.CellFormat(40, 4, "Auto", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Auto, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Nivel:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Nivel, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Financiero:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, FinancieroNombre(t.Financiero), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Tercero:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Tercero, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Centro:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Centro, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(40, 4, "Factura:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Factura, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)

	if miModulo == "Conjuntos" {
		pdf.CellFormat(40, 4, "Cuota:", "", 0,
			"", false, 0, "")
		pdf.CellFormat(40, 4, t.Cuota, "", 0,
			"", false, 0, "")
		pdf.Ln(-1)
		pdf.SetX(21)
		pdf.CellFormat(40, 4, "Grupo:", "", 0,
			"", false, 0, "")
		pdf.CellFormat(40, 4, t.Grupo, "", 0,
			"", false, 0, "")
		pdf.Ln(-1)
		pdf.SetX(21)
		pdf.CellFormat(40, 4, "Tipo:", "", 0,
			"", false, 0, "")
		pdf.CellFormat(40, 4, t.Tipo, "", 0,
			"", false, 0, "")
		pdf.Ln(-1)
		pdf.SetX(21)
		pdf.CellFormat(40, 4, "Cta. Contra:", "", 0,
			"", false, 0, "")
		pdf.CellFormat(40, 4, t.Contra, "", 0,
			"", false, 0, "")
		pdf.Ln(-1)
		pdf.SetX(21)
		pdf.CellFormat(40, 4, "Cta. Nombre:", "", 0,
			"", false, 0, "")
		pdf.CellFormat(40, 4, t.ContraNombre, "", 0,
			"", false, 0, "")
		pdf.Ln(-1)
		pdf.SetX(21)
		pdf.CellFormat(40, 4, "Intereses:", "", 0,
			"", false, 0, "")
		pdf.CellFormat(40, 4, t.Cuentaintereses, "", 0,
			"", false, 0, "")
		pdf.Ln(-1)
		pdf.SetX(21)
		pdf.CellFormat(40, 4, "Cta. Interes:", "", 0,
			"", false, 0, "")
		pdf.CellFormat(40, 4, t.Interes, "", 0,
			"", false, 0, "")
		pdf.Ln(-1)
		pdf.SetX(21)
		pdf.CellFormat(40, 4, "Cta. Nombre:", "", 0,
			"", false, 0, "")
		pdf.CellFormat(40, 4, t.Nombreintereses, "", 0,
			"", false, 0, "")
	}

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

// INICIA PLAN DE CUENTA EMPRESA TODOS PDF
func PlandecuentaempresaTodosCabecera(pdf *gofpdf.Fpdf) {
	pdf.SetFont("Arial", "", 10)
	// RELLENO TITULO
	pdf.SetY(52)
	pdf.SetFillColor(224, 231, 239)
	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(6)
	pdf.SetX(20)
	pdf.CellFormat(184, 6, "No", "0", 0,
		"L", true, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(40, 6, "Codigo", "0", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 6, "Nombre", "0", 0,
		"L", false, 0, "")
	pdf.SetX(175)
	pdf.CellFormat(40, 6, "Auto", "0", 0,
		"L", false, 0, "")
	pdf.SetX(192)
	pdf.CellFormat(40, 6, "Nivel", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}
func PlandecuentaempresaTodosDetalle(pdf *gofpdf.Fpdf, miFila plandecuentaempresa, a int) {
	pdf.SetFont("Arial", "", 9)

	ene := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetTextColor(0, 0, 0)
	// fila normal
	pdf.SetX(21)
	pdf.CellFormat(183, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(40, 4, miFila.Codigo, "", 0,
		"L", false, 0, "")
	pdf.SetX(50)
	pdf.CellFormat(40, 4, ene(miFila.Nombre), "", 0, "L", false, 0, "")
	pdf.SetX(175)
	pdf.CellFormat(20, 4, miFila.Auto, "", 0,
		"L", false, 0, "")
	pdf.SetX(195)
	pdf.CellFormat(20, 4, miFila.Nivel, "", 0,
		"L", false, 0, "")
	pdf.Ln(4)
}

func PlandecuentaempresaTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	//	Codigo := mux.Vars(r)["codigo"]

	t := []plandecuentaempresa{}
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	err := db.Select(&t, "SELECT * FROM plandecuentaempresa ORDER BY codigo ")
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
		pdf.CellFormat(190, 10, "PLAN DE CUENTA EMPRESA", "0", 0,
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

	PlandecuentaempresaTodosCabecera(pdf)
	// tercera pagina

	for i, miFila := range t {
		PlandecuentaempresaTodosDetalle(pdf, miFila, i+1)
		if math.Mod(float64(i+1), 48) == 0 {

			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			PlandecuentaempresaTodosCabecera(pdf)
		}
	}

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

// PLANDECUENTAEMPRESA EXCEL
func PlandecuentaempresaExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	t := []plandecuentaempresa{}
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	err := db.Select(&t, "SELECT * FROM plandecuentaempresa ORDER BY codigo")
	if err != nil {
		log.Fatalln(err)
	}

	f := excelize.NewFile()

	// FUNCION ANCHO DE LA COLUMNA
	if err = f.SetColWidth("Sheet1", "A", "A", 13); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.SetColWidth("Sheet1", "B", "B", 60); err != nil {
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

	// FUNCION PARA UNIR DOS CELDAS
	if err = f.MergeCell("Sheet1", "A1", "D1"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A2", "D2"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A3", "D3"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A4", "D4"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A5", "D5"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A6", "D6"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A7", "D7"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A8", "D8"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A9", "D9"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A10", "D10"); err != nil {
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
	f.SetCellValue("Sheet1", "A9", "PLAN DE CUENTA EMPRESA")
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
	f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel), "Auto")
	f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel), "Nivel")

	f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel), "A"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel), "B"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel), "C"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel), "D"+strconv.Itoa(filaExcel), estiloCabecera)
	filaExcel++

	for i, miFila := range t {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel+i), miFila.Codigo)
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel+i), miFila.Nombre)
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel+i), miFila.Auto)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel+i), miFila.Nivel)

		f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel+i), "A"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel+i), "B"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel+i), "C"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel+i), "D"+strconv.Itoa(filaExcel+i), estiloTexto)

		//van=i
	}

	// LIENA FINAL
	//a=strconv.Itoa(van+1+filaExcel)
	// Set the headers necessary to get browsers to interpret the downloadable file
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

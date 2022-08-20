package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
)

type ExogenaFormato struct {
	Codigo string
	Nombre string
}

func ListaExogenaFormato() []ExogenaFormato {
	log.Println("lista formato")
	db := dbConn()
	res := []ExogenaFormato{}
	err := db.Select(&res, "SELECT * FROM exogenaformato order by codigo")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("lista formato")
	return res
}

func listaConceptoFormato() []ExogenaConcepto {
	log.Println("lista concepto")
	db := dbConn()
	res := []ExogenaConcepto{}
	err := db.Select(&res, "SELECT * FROM exogenaconcepto  order by formato,concepto")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("lista concepto")
	return res
}

func listaColumnaFormato() []ExogenaColumna {
	log.Println("lista concepto")
	db := dbConn()
	res := []ExogenaColumna{}
	err := db.Select(&res, "SELECT columna as codigo,columna,formato,nombre FROM exogenacolumna  order by formato,columna")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("lista concepto")
	return res
}

type ExogenaConcepto struct {
	Formato  string
	Concepto string
	Nombre   string
}

func ListaExogenaConcepto(codigo string) []ExogenaConcepto {
	log.Println("lista formato")
	db := dbConn()
	res := []ExogenaConcepto{}
	err := db.Select(&res, "SELECT * FROM exogenaConcepto where formato=$1 order by concepto", codigo)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("lista formato")
	return res
}

type ExogenaColumna struct {
	Codigo  string
	Nombre  string
	Formato string
	Columna string
}

func ListaExogenaColumna(codigo string) []ExogenaColumna {
	log.Println("lista formato")
	db := dbConn()
	res := []ExogenaColumna{}
	db.Select(&res, "SELECT codigo FROM exogenaColumna where formato=$1 order by columna", codigo)
	log.Println("lista formato")
	return res
}

type ExogenaDatos struct {
	Cuenta   string
	Formato  string
	Concepto string
	Columna  string
	Valor    string
}
type Exogena struct {
	Codigo  string
	Detalle []ExogenaDatos
}

type exogenaeditar struct {
	//Fila          string
	Cuenta         string `json:"Cuenta"`
	Formato        string
	FormatoNombre  string
	Concepto       string
	ConceptoNombre string
	Columna        string
	ColumnaNombre  string
	Valor          string
}

type formatolista struct {
	//Fila          string
	Formato string
}

func ExogenaConsultaDetalle() string {
	var consulta = ""
	consulta = " select "
	consulta += "exogena.cuenta, "
	consulta += "exogena.formato, "
	consulta += "exogenaformato.codigo||'-'||exogenaformato.nombre as formatonombre, "
	consulta += "exogena.concepto, "
	consulta += "'' as conceptonombre, "
	consulta += "exogena.columna, "
	consulta += "exogenacolumna.columna as columnanombre, "

	consulta += "exogena.valor "
	consulta += "from exogena "
	consulta += " inner join exogenaformato "
	consulta += " on exogena.formato= exogenaformato.codigo"
	//	consulta += " inner join exogenaconcepto "
	//	consulta += " on exogena.formato=exogenaconcepto.formato and exogena.concepto= exogenaconcepto.concepto"
	consulta += " inner join exogenacolumna "
	consulta += " on exogena.columna=exogenacolumna.columna and exogena.formato= exogenacolumna.formato"
	consulta += " where exogena.cuenta=$1 "
	log.Println(consulta)
	return consulta
}

func ExogenaAgregar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var tmpExogena = Exogena{}

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// carga informacion de la CUENTADECOBRO
	err = json.Unmarshal(b, &tmpExogena)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// borra detalle anterior
	delForm, err := db.Prepare("DELETE from exogena WHERE cuenta=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(tmpExogena.Codigo)

	for i, x := range tmpExogena.Detalle {
		var a = i
		var q string

		q = "insert into exogena ("
		q += "Cuenta,"
		q += "Formato,"
		q += "Concepto,"
		q += "Columna,"
		q += "Valor"
		q += " ) values("
		q += parametros(5)
		q += ")"

		log.Println("Cadena SQL " + q)
		insForm, err := db.Prepare(q)
		if err != nil {
			panic(err.Error())
		}

		// TERMINA CUENTADECOBRO GRABAR INSERTAR
		_, err = insForm.Exec(
			x.Cuenta,
			x.Formato,
			x.Concepto,
			x.Columna,
			x.Valor)

		if err != nil {
			panic(err)
		}

		log.Println("Insertar Detalle \n", x.Cuenta, a)
	}

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

	//http.Redirect(w, r, "/CUENTADECOBROLista", 301)
}

func ExogenaEditar(w http.ResponseWriter, r *http.Request) {

	Numero := mux.Vars(r)["numero"]
	log.Println("inicio exogena Cuenta" + Numero)
	db := dbConn()

	// traer cuentadecobro
	v := plandecuentaempresa{}
	err := db.Get(&v, "SELECT * FROM plandecuentaempresa WHERE  codigo=$1", Numero)
	if err != nil {
		log.Fatalln(err)
	}

	// traer detalle
	det := []exogenaeditar{}
	err2 := db.Select(&det, ExogenaConsultaDetalle(), Numero)
	if err2 != nil {
		fmt.Println(err2)
	}

	// traer tercero

	var panel = "False"

	//	log.Println("detalle cuentadecobro)
	parametros := map[string]interface{}{
		"panel":    panel,
		"cuenta":   v,
		"detalle":  det,
		"hosting":  ruta,
		"centro":   ListaCentro(),
		"formato":  ListaExogenaFormato(),
		"concepto": listaConceptoFormato(),
		"columna":  listaColumnaFormato(),
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/exogena/exogenaEditar.html",
		"vista/exogena/autocompletaFormato.html",
		"vista/exogena/autocompletaConcepto.html",
		"vista/exogena/autocompletaColumna.html",
		"vista/exogena/autocompletaplandecuentaempresa.html",
		"vista/exogena/exogenaScript.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error cuentadecobro nuevo 3")
	miTemplate.Execute(w, parametros)

	//tmp.Execute(w, parametros)
}
func traeformato(cuenta string) string {
	// traer cuentadecobro
	listaformato1 := []formatolista{}
	err := db.Select(&listaformato1, "SELECT distinct formato from exogena where cuenta=$1", cuenta)
	if err != nil {
		log.Fatalln(err)
	}
	var formatos = ""

	for _, x := range listaformato1 {
		formatos += "(  " + x.Formato + "  )"

	}
	return formatos
}

func movimiento(Codigo string) string {

	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM comprobantedetalle  WHERE cuenta=$1", Codigo)
	err := row.Scan(&total)
	if err != nil {
		log.Fatal(err)
	}
	var resultado string
	resultado = ""
	if total > 0 {
		log.Println("si encontro")
		resultado = "SI"
	} else {
		resultado = "NO"
		log.Println("no encontro")
	}
	return resultado
}

func ExogenaLista(w http.ResponseWriter, r *http.Request) {
	panel := mux.Vars(r)["panel"]
	codigo := mux.Vars(r)["codigo"]
	elemento := mux.Vars(r)["elemento"]

	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/exogena/ExogenaLista.html")
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
		Auto = ""
		// traer formato
		if Nivel == "A" {
			//Nombre += traeformato(Codigo)
			Auto = traeformato(Codigo)
			if movimiento(Codigo) == "SI" {
				res = append(res, plandecuentaempresaLista{Codigo, Nombre, Auto, Nivel})

			}
		} else {
			res = append(res, plandecuentaempresaLista{Codigo, Nombre, Auto, Nivel})

		}

	}
	varmap := map[string]interface{}{
		"res":      res,
		"hosting":  ruta,
		"panel":    panel,
		"codigo":   codigo,
		"elemento": elemento,
	}
	tmp.Execute(w, varmap)
}

// TERMINA CUENTADECOBRODETALLE ESTRUCTURA
func ExogenaGenerar(w http.ResponseWriter, r *http.Request) {

	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/exogena/ExogenaGenerar.html")
	db := dbConn()
	res := []ExogenaFormato{}
	err2 := db.Select(&res, "select codigo,nombre FROM exogenaformato ")
	if err2 != nil {
		fmt.Println(err2)
	}
	log.Println("lista formato")

	panel := mux.Vars(r)["panel"]
	varmap := map[string]interface{}{
		"hosting": ruta,
		"centro":  ListaCentro(),
		"panel":   panel,
		"formato": res,
	}
	tmp.Execute(w, varmap)
}

// consultar concepto
func TraerConceptoExogena(w http.ResponseWriter, r *http.Request) {
	//db := dbConn()
	Codigo := mux.Vars(r)["formato"]

	var resJson []ExogenaConcepto
	resJson = ListaExogenaConcepto(Codigo)

	data, _ := json.Marshal(resJson)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}

func TraerFormatoExogena(w http.ResponseWriter, r *http.Request) {
	//db := dbConn()
	//Codigo := mux.Vars(r)["formato"]

	db := dbConn()
	res := []ExogenaFormato{}
	err2 := db.Select(&res, "select codigo,nombre FROM exogenaformato ")
	if err2 != nil {
		fmt.Println(err2)
	}

	data, _ := json.Marshal(res)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}

func TraerColumnaExogena(w http.ResponseWriter, r *http.Request) {
	//db := dbConn()
	Codigo := mux.Vars(r)["formato"]

	db := dbConn()
	res := []ExogenaColumna{}
	err2 := db.Select(&res, "select columna as codigo,nombre,formato FROM exogenacolumna where formato=$1", Codigo)
	if err2 != nil {
		fmt.Println(err2)
	}

	data, _ := json.Marshal(res)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}

func FormatoExogenaBuscar(w http.ResponseWriter, r *http.Request) {

	var tabla = "exogenaformato"
	db := dbConn()
	Codigo := Mayuscula(mux.Vars(r)["codigo"])
	selDB, err := db.Query("SELECT codigo,"+
		"codigo||'-'||nombre as nombre FROM "+tabla+" where codigo LIKE '%' || $1 || '%'  or  UPPER(codigo||'-'||nombre) LIKE '%' || $1 || '%'  ORDER BY"+
		" codigo ASC", Codigo)
	if err != nil {
		log.Println(err.Error())
	}
	var resJson []buscarJson
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
		value = nombre
		label = nombre
		resJson = append(resJson, buscarJson{id, label, value, nombre})
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

type buscarParametroConcepto struct {
	Codigo  string `json:"Codigo"`
	Nombre  string `json:"Nombre"`
	Formato string `json:"Formato"`
}

// CENTRO ACTUAL
func FormatoExogenaActual(w http.ResponseWriter, r *http.Request) {
	var tabla = "exogenaformato"
	emp := buscarParametro{}
	var res []buscarParametro
	db := dbConn()
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var tempParametro buscarParametro
	err = json.Unmarshal(b, &tempParametro)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	log.Printf("codigo: %+v\n", tempParametro.Codigo)
	log.Printf("Nombre: %+v\n", tempParametro.Codigo)

	err = db.Get(&emp, "SELECT codigo,codigo||'-'||nombre as nombre FROM "+tabla+" where ( codigo=$1 or UPPER(nombre)=$2)", tempParametro.Codigo, Mayuscula(tempParametro.Nombre))
	switch err {
	case nil:
		log.Printf("tercero found: %+v\n", emp)
		res = append(res, emp)
		data, err := json.Marshal(res)
		if err != nil {
			log.Fatalln(err)
		}
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)

	case sql.ErrNoRows:
		log.Println("tercero NOT found, no error")
		var slice []string
		slice = make([]string, 0)
		data, _ := json.Marshal(slice)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	default:
		log.Printf("tercero error: %s\n", err)
	}
}

func ConceptoExogenaBuscar(w http.ResponseWriter, r *http.Request) {

	var tabla = "exogenaConcepto"
	db := dbConn()
	Codigo := Mayuscula(mux.Vars(r)["codigo"])
	Formato := Mayuscula(mux.Vars(r)["formato"])
	var consulta = ""
	consulta = "SELECT concepto as codigo,"
	consulta += "concepto||'-'||nombre as nombre FROM " + tabla + " where formato=$1 and ( concepto LIKE '%' || $2 || '%'  or  UPPER(nombre) LIKE '%' || $2 || '%')  ORDER BY"
	consulta += " concepto ASC"

	selDB, err := db.Query(consulta, Formato, Codigo)

	if err != nil {
		log.Println(err.Error())
	}
	var resJson []buscarJson
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
		value = nombre
		label = nombre
		resJson = append(resJson, buscarJson{id, label, value, nombre})
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

// CENTRO ACTUAL
func ConceptoExogenaActual(w http.ResponseWriter, r *http.Request) {
	var tabla = "exogenaconcepto"
	emp := buscarParametro{}
	var res []buscarParametro
	db := dbConn()
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var tempParametro buscarParametroConcepto
	err = json.Unmarshal(b, &tempParametro)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	log.Printf("codigo: %+v\n", tempParametro.Codigo)
	log.Printf("Nombre: %+v\n", tempParametro.Codigo)

	err = db.Get(&emp, "SELECT concepto as codigo,concepto||'-'||nombre as nombre FROM "+tabla+" where concepto=$1 or UPPER(concepto||'-'||nombre)=$2", tempParametro.Codigo, Mayuscula(tempParametro.Nombre))
	switch err {
	case nil:
		log.Printf("tercero found: %+v\n", emp)
		res = append(res, emp)
		data, err := json.Marshal(res)
		if err != nil {
			log.Fatalln(err)
		}
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)

	case sql.ErrNoRows:
		log.Println("tercero NOT found, no error")
		var slice []string
		slice = make([]string, 0)
		data, _ := json.Marshal(slice)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	default:
		log.Printf("tercero error: %s\n", err)
	}
}

/// columna

func ColumnaExogenaBuscar(w http.ResponseWriter, r *http.Request) {

	var tabla = "exogenaColumna"
	db := dbConn()
	Codigo := Mayuscula(mux.Vars(r)["codigo"])
	Formato := Mayuscula(mux.Vars(r)["formato"])
	var consulta = ""
	consulta = "SELECT columna as codigo,"
	consulta += "columna||'-'||nombre as nombre FROM " + tabla + " where formato=$1 and ( columna LIKE '%' || $2 || '%'  or  UPPER(nombre) LIKE '%' || $2 || '%')  ORDER BY"
	consulta += " columna ASC"

	selDB, err := db.Query(consulta, Formato, Codigo)

	if err != nil {
		log.Println(err.Error())
	}
	var resJson []buscarJson
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
		value = nombre
		label = nombre
		resJson = append(resJson, buscarJson{id, label, value, nombre})
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

// CENTRO ACTUAL
func ColumnaExogenaActual(w http.ResponseWriter, r *http.Request) {
	var tabla = "exogenacolumna"
	emp := buscarParametro{}
	var res []buscarParametro
	db := dbConn()
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	var tempParametro buscarParametroConcepto
	err = json.Unmarshal(b, &tempParametro)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	log.Printf("codigo: %+v\n", tempParametro.Codigo)
	log.Printf("Nombre: %+v\n", tempParametro.Codigo)

	err = db.Get(&emp, "SELECT columna as codigo,columna||'-'||nombre as nombre FROM "+tabla+" where columna=$1 or UPPER(columna||'-'||nombre)=$2", tempParametro.Codigo, Mayuscula(tempParametro.Nombre))
	switch err {
	case nil:
		log.Printf("tercero found: %+v\n", emp)
		res = append(res, emp)
		data, err := json.Marshal(res)
		if err != nil {
			log.Fatalln(err)
		}
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)

	case sql.ErrNoRows:
		log.Println("tercero NOT found, no error")
		var slice []string
		slice = make([]string, 0)
		data, _ := json.Marshal(slice)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	default:
		log.Printf("tercero error: %s\n", err)
	}
}

//datos
func ExogenaListaDatos(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Numero := Mayuscula(mux.Vars(r)["numero"])
	// traer detalle
	det := []exogenaeditar{}
	//var det exogenaeditar []
	err2 := db.Select(&det, ExogenaConsultaDetalle(), Numero)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	var resJson []exogenaeditar
	for _, x := range det {
		resJson = append(resJson, exogenaeditar{x.Cuenta, x.Formato, x.FormatoNombre, x.Concepto, x.ConceptoNombre, x.Columna, x.ColumnaNombre, x.Valor})

	}
	data, err := json.Marshal(resJson)
	if err != nil {
		log.Fatalln(err)
	}
	//data, _ := json.Marshal(det)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)

}

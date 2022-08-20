package main

import (
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net"
	"net/http"
	"reflect"
	"sync"
	"time"
)

// CONEXION A BASE DE DATOS POSTGRES
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Murc4505"
	dbname   = "vhsa0306"
)

var mapaRuta = map[string]interface{}{
	"hosting": ruta,
}
var (
	once sync.Once
	db   *sqlx.DB
)

func dbConn() *sqlx.DB {
	var err error
	once.Do(func() {
		psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname)
		//db, err = sql.Open("postgres", psqlInfo)
		db, err = sqlx.Open("postgres", psqlInfo)

		log.Println("conectando...")
		if err != nil {
			panic(err.Error())
			log.Println(err.Error())
		}
		log.Println("Base de Datos Conectada")
	})
	return db
}
func periodoSesion(r *http.Request) string {
	session, _ := cookie.Get(r, "Golang-session")
	var miperiodo string
	miperiodo = "SIN PERIODO"
	var iperiodo interface{} = session.Values["periodo"]
	if iperiodo != nil {
		miperiodo = session.Values["periodo"].(string)
	}
	return miperiodo

}

// LISTA DE REGISTROS//
func IndexEmpresa(w http.ResponseWriter, r *http.Request) {
	//var miperiodo = periodoSesion(r)
	miperiodo := mux.Vars(r)["periodo"]
	////periodo := r.FormValue("periodo")
	log.Println("miperiodo cokies" + miperiodo)
	session, _ := cookie.Get(r, "Golang-session")
	session.Values["periodo"] = miperiodo
	err := session.Save(r, w)
	if err != nil {
		log.Println("failed to save session", err)
	}

	// trae año sesion
	//session, _ := cookie.Get(r, "Golang-session")
	//var miperiodo string
	//miperiodo = "SIN PERIODO"
	//var iperiodo interface{} = session.Values["periodo"]
	//if iperiodo != nil {
	//	miperiodo = session.Values["periodo"].(string)
	//}
	// año sesion

	////**miperiodo := session.Values["periodo"].(string)
	//miperiodo = miperiodo + " sesion "
	//println(miperiodo)
	tmp, _ := template.ParseFiles("vista/inicio/Index.html",
		"vista/inicio/appInicio.html")
	db := dbConn()
	Empresa := empresa{}
	err = db.Get(&Empresa, "SELECT * FROM empresa limit 1")
	if err != nil {
		//panic(err.Error())
	}

	Empresa.Nombre = Titulo(Empresa.Nombre)
	Empresa.Codigo = Coma(Empresa.Codigo)
	log.Println("inicio listado")

	varmap := map[string]interface{}{
		"hosting":   ruta,
		"empresa":   Empresa,
		"miperiodo": miperiodo,
	}

	// verifica tabla saldos
	//db := dbConn()
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM saldo ")
	err = row.Scan(&total)
	if err != nil {
		log.Fatal(err)
	}
	//var resultado bool
	if total > 0 {
		//resultado = true
	} else {
		log.Println("consultar configuracion saldo ")
		GeneraTablaSaldo()
		//Datosinicialesconfiguracioninventario()
		//resultado = false
	}

	log.Println("Error empresa888")
	tmp.Execute(w, varmap)
}

type EmpresaListaPeriodo struct {
	Codigo    string
	Nombre    string
	Anualidad string
}

func Index(w http.ResponseWriter, r *http.Request) {
	//tmp, _ := template.ParseFiles("vista/inicio/IndexPeriodo.html",
	//	"vista/inicio/appInicio.html")
	db := dbConn()
	Empresa := empresa{}
	err := db.Get(&Empresa, "SELECT * FROM empresa limit 1")
	if err != nil {
		//panic(err.Error())
	}

	Empresalista1 := []EmpresaListaPeriodo{}
	err = db.Select(&Empresalista1, "select empresa.codigo,empresa.nombre,periodo.anualidad from periodo,empresa order by periodo.anualidad")
	if err != nil {
		println(err.Error())
	}

	Empresa.Nombre = Titulo(Empresa.Nombre)
	Empresa.Codigo = Coma(Empresa.Codigo)
	log.Println("inicio listado")

	parametros := map[string]interface{}{
		"hosting":      ruta,
		"empresa":      Empresa,
		"periodo":      ListaPeriodo(),
		"empresalista": Empresalista1,
	}

	// verifica tabla saldos
	//db := dbConn()
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM saldo ")
	err = row.Scan(&total)
	if err != nil {
		log.Fatal(err)
	}
	//var resultado bool
	if total > 0 {
		//resultado = true
	} else {
		log.Println("consultar configuracion saldo ")
		GeneraTablaSaldo()
		//Datosinicialesconfiguracioninventario()
		//resultado = false
	}
	miTemplate, err := template.ParseFiles("vista/inicio/IndexPeriodo.html",
		"vista/inicio/appInicio.html")
	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error comprobante nuevo 3")
	miTemplate.Execute(w, parametros)

	log.Println("Error empresa888")
	//	tmp.Execute(w, varmap)
}
func IndexPeriodo(w http.ResponseWriter, r *http.Request) {
	log.Println("indexperiodo")

	db := dbConn()
	miperiodo := mux.Vars(r)["periodo"]
	////periodo := r.FormValue("periodo")
	log.Println("miperiodo cokies" + miperiodo)
	session, _ := cookie.Get(r, "Golang-session")
	session.Values["periodo"] = miperiodo
	err := session.Save(r, w)
	if err != nil {
		log.Println("failed to save session", err)
	}
	var q string
	q = "UPDATE empresa set "
	q += "Periodo=$1"

	log.Println("cadena actualizaicion" + q)

	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	_, err = insForm.Exec(miperiodo)
	if err != nil {
		panic(err)
	}
	http.Redirect(w, r, "/IndexEmpresa555", http.StatusSeeOther)
}

// RUTAS GENERALES
var router = mux.NewRouter()
var ruta = "http://localhost:9002/"

//var ruta="http://192.168.1.3:9002/"
//var ruta="/"

func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
func favicon(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s\n", r.RequestURI)
	w.Header().Set("Content-Type", "image/x-icon")
	w.Header().Set("Cache-Control", "public, max-age=7776000")
	fmt.Fprintln(w, "data:image/x-icon;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQEAYAAABPYyMiAAAABmJLR0T///////8JWPfcAAAACXBIWXMAAABIAAAASABGyWs+AAAAF0lEQVRIx2NgGAWjYBSMglEwCkbBSAcACBAAAeaR9cIAAAAASUVORK5CYII=\n")
}
func CacheControlWrapper(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Cache-Control", "max-age=2592000") // 30 days
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")

		h.ServeHTTP(w, r)
	})
}

var cookie *sessions.CookieStore

func main() {
	cookie = sessions.NewCookieStore([]byte("Golang-Blogs"))
	//router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
	//http.FileServer(http.Dir("static"))))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/",
		CacheControlWrapper(http.FileServer(http.Dir("static")))))

	router.Path("/").HandlerFunc(Index).Name("Index")

	router.Path("/IndexEmpresa/{periodo}").HandlerFunc(IndexEmpresa).Name("IndexEmpresa")
	router.Path("/IndexPeriodo/{periodo}").HandlerFunc(IndexPeriodo).Name("IndexPeriodo")

	// favicon
	http.HandleFunc("/favicon.ico", favicon)

	// ARCHIVO FORMA DE PAGO
	router.Path("/FormadepagoNuevo/{panel}/{codigo}/{elemento}").HandlerFunc(FormadepagoNuevo).Name("FormadepagoNuevo")
	router.Path("/FormadepagoLista").HandlerFunc(FormadepagoLista).Name("FormadepagoLista")
	router.Path("/FormadepagoExiste/{codigo:[0-9]+}").HandlerFunc(FormadepagoExiste).
		Name("FormadepagoExiste")
	router.Path("/FormadepagoInsertar").HandlerFunc(FormadepagoInsertar).Name(
		"FormadepagoInsertar")
	router.Path("/FormadepagoActualizar/{codigo:[0-9]+}").HandlerFunc(FormadepagoActualizar).Name(
		"FormadepagoActualizar")
	router.Path("/FormadepagoBorrar/{codigo:[0-9]+}").HandlerFunc(FormadepagoBorrar).Name(
		"FormadepagoBorrar")
	router.Path("/FormadepagoEliminar/{codigo:[0-9]+}").HandlerFunc(
		FormadepagoEliminar).Name("FormadepagoEliminar")
	router.Path("/FormadepagoEditar/{codigo:[0-9]+}").HandlerFunc(FormadepagoEditar).Name(
		"FormadepagoEditar")
	router.Path("/FormadepagoPdf/{codigo:[0-9]+}").HandlerFunc(FormadepagoPdf).Name(
		"FormadepagoPdf")
	router.Path("/FormadepagoBuscar/{codigo}").HandlerFunc(FormadepagoBuscar).
		Name("FormadepagoBuscar")
	router.Path("/FormadepagoActual").HandlerFunc(FormadepagoActual).
		Name("FormadepagoActual")

	router.Path("/FormadepagoTodosPdf").HandlerFunc(FormadepagoTodosPdf).
		Name("FormadepagoTodosPdf")
	router.Path("/FormadepagoExcel").HandlerFunc(FormadepagoExcel).
		Name("FormadepagoExcel")

	// ARCHIVO CENTRO
	router.Path("/CentroNuevo/{panel}/{codigo}/{elemento}").HandlerFunc(CentroNuevo).Name("CentroNuevo")
	router.Path("/CentroLista").HandlerFunc(CentroLista).Name("CentroLista")
	router.Path("/CentroExiste/{codigo:[0-9]+}").HandlerFunc(CentroExiste).
		Name("CentroExiste")
	router.Path("/CentroInsertar").HandlerFunc(CentroInsertar).Name(
		"CentroInsertar")
	router.Path("/CentroActualizar/{codigo:[0-9]+}").HandlerFunc(CentroActualizar).Name(
		"CentroActualizar")
	router.Path("/CentroBorrar/{codigo:[0-9]+}").HandlerFunc(CentroBorrar).Name(
		"CentroBorrar")
	router.Path("/CentroEliminar/{codigo:[0-9]+}").HandlerFunc(
		CentroEliminar).Name("CentroEliminar")
	router.Path("/CentroEditar/{codigo:[0-9]+}").HandlerFunc(CentroEditar).Name(
		"CentroEditar")
	router.Path("/CentroPdf/{codigo:[0-9]+}").HandlerFunc(CentroPdf).Name(
		"CentroPdf")

	router.Path("/CentroBuscar/{codigo}").HandlerFunc(CentroBuscar).
		Name("CentroBuscar")

	router.Path("/CentroBuscarCrear/{codigo}").HandlerFunc(CentroBuscarCrear).
		Name("CentroBuscarCrear")

	router.Path("/CentroActual").HandlerFunc(CentroActual).
		Name("CentroActual")

	router.Path("/CentroActual/{codigo}").HandlerFunc(CentroActualConsulta).
		Name("CentroActualConsulta")
	router.Path("/CentroTodosPdf").HandlerFunc(CentroTodosPdf).
		Name("CentroTodosPdf")
	router.Path("/CentroExcel").HandlerFunc(CentroExcel).
		Name("CentroExcel")

	// ARCHIVO IVA
	router.Path("/IvaNuevo/{panel}/{codigo}/{elemento}").HandlerFunc(IvaNuevo).Name("IvaNuevo")
	router.Path("/IvaLista").HandlerFunc(IvaLista).Name("IvaLista")
	router.Path("/IvaExiste/{codigo:[0-9]+}").HandlerFunc(IvaExiste).
		Name("IvaExiste")
	router.Path("/IvaInsertar").HandlerFunc(IvaInsertar).Name(
		"IvaInsertar")
	router.Path("/IvaActualizar/{codigo:[0-9]+}").HandlerFunc(IvaActualizar).Name(
		"IvaActualizar")
	router.Path("/IvaBorrar/{codigo:[0-9]+}").HandlerFunc(IvaBorrar).Name(
		"IvaBorrar")
	router.Path("/IvaEliminar/{codigo:[0-9]+}").HandlerFunc(
		IvaEliminar).Name("IvaEliminar")
	router.Path("/IvaEditar/{codigo:[0-9]+}").HandlerFunc(IvaEditar).Name(
		"IvaEditar")
	router.Path("/IvaPdf/{codigo:[0-9]+}").HandlerFunc(IvaPdf).Name(
		"IvaPdf")

	router.Path("/IvaBuscar/{codigo}").HandlerFunc(IvaBuscar).
		Name("IvaBuscar")

	router.Path("/IvaBuscarCrear/{codigo}").HandlerFunc(IvaBuscarCrear).
		Name("IvaBuscarCrear")

	router.Path("/IvaActual").HandlerFunc(IvaActual).
		Name("IvaActual")

	router.Path("/IvaActual/{codigo}").HandlerFunc(IvaActualConsulta).
		Name("IvaActualConsulta")
	router.Path("/IvaTodosPdf").HandlerFunc(IvaTodosPdf).
		Name("IvaTodosPdf")
	router.Path("/IvaExcel").HandlerFunc(IvaExcel).
		Name("IvaExcel")

	// ARCHIVO RESPONSABILIDAD FISCAL
	router.Path("/ResponsabilidadfiscalNuevo/{panel}/{codigo}/{elemento}").HandlerFunc(ResponsabilidadfiscalNuevo).Name("ResponsabilidadfiscalNuevo")
	router.Path("/ResponsabilidadfiscalBuscar/{codigo}").HandlerFunc(ResponsabilidadfiscalBuscar).
		Name("ResponsabilidadfiscalBuscar")
	router.Path("/ResponsabilidadfiscalActual/{codigo}").HandlerFunc(ResponsabilidadfiscalActual).
		Name("ResponsabilidadfiscalActual")
	router.Path("/ResponsabilidadfiscalLista").HandlerFunc(ResponsabilidadfiscalLista).Name("ResponsabilidadfiscalLista")
	router.Path("/ResponsabilidadfiscalExiste/{codigo}").HandlerFunc(ResponsabilidadfiscalExiste).
		Name("ResponsabilidadfiscalExiste")
	router.Path("/ResponsabilidadfiscalInsertar").HandlerFunc(ResponsabilidadfiscalInsertar).Name(
		"ResponsabilidadfiscalInsertar")
	router.Path("/ResponsabilidadfiscalActualizar/{codigo}").HandlerFunc(ResponsabilidadfiscalActualizar).Name(
		"ResponsabilidadfiscalActualizar")
	router.Path("/ResponsabilidadfiscalBorrar/{codigo}").HandlerFunc(ResponsabilidadfiscalBorrar).Name(
		"ResponsabilidadfiscalBorrar")
	router.Path("/ResponsabilidadfiscalEliminar/{codigo}").HandlerFunc(
		ResponsabilidadfiscalEliminar).Name("ResponsabilidadfiscalEliminar")
	router.Path("/ResponsabilidadfiscalEditar/{codigo}").HandlerFunc(ResponsabilidadfiscalEditar).Name(
		"ResponsabilidadfiscalEditar")
	router.Path("/ResponsabilidadfiscalPdf/{codigo}").HandlerFunc(ResponsabilidadfiscalPdf).Name(
		"ResponsabilidadfiscalPdf")

	router.Path("/ResponsabilidadfiscalTodosPdf").HandlerFunc(ResponsabilidadfiscalTodosPdf).
		Name("ResponsabilidadfiscalTodosPdf")
	router.Path("/ResponsabilidadfiscalExcel").HandlerFunc(ResponsabilidadfiscalExcel).
		Name("ResponsabilidadfiscalExcel")

	// ARCHIVO DOCUMENTO IDENTIFICACION
	router.Path("/DocumentoidentificacionNuevo/{panel}/{codigo}/{elemento}").HandlerFunc(DocumentoidentificacionNuevo).Name("DocumentoidentificacionNuevo")
	router.Path("/DocumentoidentificacionBuscar/{codigo}").HandlerFunc(DocumentoidentificacionBuscar).
		Name("DocumentoidentificacionBuscar")
	router.Path("/DocumentoidentificacionActual/{codigo}").HandlerFunc(DocumentoidentificacionActual).
		Name("DocumentoidentificacionActual")
	router.Path("/DocumentoidentificacionLista").HandlerFunc(DocumentoidentificacionLista).Name("DocumentoidentificacionLista")
	router.Path("/DocumentoidentificacionExiste/{codigo:[0-9]+}").HandlerFunc(DocumentoidentificacionExiste).
		Name("DocumentoidentificacionExiste")
	router.Path("/DocumentoidentificacionInsertar").HandlerFunc(DocumentoidentificacionInsertar).Name(
		"DocumentoidentificacionInsertar")
	router.Path("/DocumentoidentificacionActualizar/{codigo:[0-9]+}").HandlerFunc(DocumentoidentificacionActualizar).Name(
		"DocumentoidentificacionActualizar")
	router.Path("/DocumentoidentificacionBorrar/{codigo:[0-9]+}").HandlerFunc(DocumentoidentificacionBorrar).Name(
		"DocumentoidentificacionBorrar")
	router.Path("/DocumentoidentificacionEliminar/{codigo:[0-9]+}").HandlerFunc(
		DocumentoidentificacionEliminar).Name("DocumentoidentificacionEliminar")
	router.Path("/DocumentoidentificacionEditar/{codigo:[0-9]+}").HandlerFunc(DocumentoidentificacionEditar).Name(
		"DocumentoidentificacionEditar")
	router.Path("/DocumentoidentificacionPdf/{codigo:[0-9]+}").HandlerFunc(DocumentoidentificacionPdf).Name(
		"DocumentoidentificacionPdf")

	router.Path("/DocumentoidentificacionTodosPdf").HandlerFunc(DocumentoidentificacionTodosPdf).
		Name("DocumentoidentificacionTodosPdf")
	router.Path("/DocumentoidentificacionExcel").HandlerFunc(DocumentoidentificacionExcel).
		Name("DocumentoidentificacionExcel")

	// ARCHIVO REGIMEN FISCAL
	router.Path("/RegimenfiscalNuevo/{panel}/{codigo}/{elemento}").HandlerFunc(RegimenfiscalNuevo).Name("RegimenfiscalNuevo")
	router.Path("/RegimenfiscalBuscar/{codigo}").HandlerFunc(RegimenfiscalBuscar).
		Name("RegimenfiscalBuscar")
	router.Path("/RegimenfiscalActual/{codigo}").HandlerFunc(RegimenfiscalActual).
		Name("RegimenfiscalActual")
	router.Path("/RegimenfiscalLista").HandlerFunc(RegimenfiscalLista).Name("RegimenfiscalLista")
	router.Path("/RegimenfiscalExiste/{codigo:[0-9]+}").HandlerFunc(RegimenfiscalExiste).
		Name("RegimenfiscalExiste")
	router.Path("/RegimenfiscalInsertar").HandlerFunc(RegimenfiscalInsertar).Name(
		"RegimenfiscalInsertar")
	router.Path("/RegimenfiscalActualizar/{codigo:[0-9]+}").HandlerFunc(RegimenfiscalActualizar).Name(
		"RegimenfiscalActualizar")
	router.Path("/RegimenfiscalBorrar/{codigo:[0-9]+}").HandlerFunc(RegimenfiscalBorrar).Name(
		"RegimenfiscalBorrar")
	router.Path("/RegimenfiscalEliminar/{codigo:[0-9]+}").HandlerFunc(
		RegimenfiscalEliminar).Name("RegimenfiscalEliminar")
	router.Path("/RegimenfiscalEditar/{codigo:[0-9]+}").HandlerFunc(RegimenfiscalEditar).Name(
		"RegimenfiscalEditar")
	router.Path("/RegimenfiscalPdf/{codigo:[0-9]+}").HandlerFunc(RegimenfiscalPdf).Name(
		"RegimenfiscalPdf")

	router.Path("/RegimenfiscalTodosPdf").HandlerFunc(RegimenfiscalTodosPdf).
		Name("RegimenfiscalTodosPdf")
	router.Path("/RegimenfiscalExcel").HandlerFunc(RegimenfiscalExcel).
		Name("RegimenfiscalExcel")

	// ARCHIVO MEDIO DE PAGO
	router.Path("/MediodepagoNuevo/{panel}/{codigo}/{elemento}").HandlerFunc(MediodepagoNuevo).Name("MediodepagoNuevo")
	router.Path("/MediodepagoLista").HandlerFunc(MediodepagoLista).Name("MediodepagoLista")
	router.Path("/MediodepagoExiste/{codigo:[0-9]+}").HandlerFunc(MediodepagoExiste).
		Name("MediodepagoExiste")
	router.Path("/MediodepagoInsertar").HandlerFunc(MediodepagoInsertar).Name(
		"MediodepagoInsertar")
	router.Path("/MediodepagoActualizar/{codigo:[0-9]+}").HandlerFunc(MediodepagoActualizar).Name(
		"MediodepagoActualizar")
	router.Path("/MediodepagoBorrar/{codigo:[0-9]+}").HandlerFunc(MediodepagoBorrar).Name(
		"MediodepagoBorrar")
	router.Path("/MediodepagoEliminar/{codigo:[0-9]+}").HandlerFunc(
		MediodepagoEliminar).Name("MediodepagoEliminar")
	router.Path("/MediodepagoEditar/{codigo:[0-9]+}").HandlerFunc(MediodepagoEditar).Name(
		"MediodepagoEditar")
	router.Path("/MediodepagoPdf/{codigo:[0-9]+}").HandlerFunc(MediodepagoPdf).Name(
		"MediodepagoPdf")
	router.Path("/MediodepagoBuscar/{codigo}").HandlerFunc(MediodepagoBuscar).
		Name("MediodepagoBuscar")
	router.Path("/MediodepagoActual").HandlerFunc(MediodepagoActual).
		Name("MediodepagoActual")

	router.Path("/MediodepagoTodosPdf").HandlerFunc(MediodepagoTodosPdf).
		Name("MediodepagoTodosPdf")
	router.Path("/MediodepagoExcel").HandlerFunc(MediodepagoExcel).
		Name("MediodepagoExcel")

	// ARCHIVO CARGO
	router.Path("/CargoNuevo/{codigo}").HandlerFunc(CargoNuevo).Name("CargoNuevo")
	router.Path("/CargoBuscar/{codigo}").HandlerFunc(CargoBuscar).
		Name("CargoBuscar")
	router.Path("/CargoActual/{codigo}").HandlerFunc(CargoActual).
		Name("CargoActual")
	router.Path("/CargoLista").HandlerFunc(CargoLista).Name("CargoLista")
	router.Path("/CargoExiste/{codigo:[0-9]+}").HandlerFunc(CargoExiste).
		Name("CargoExiste")
	router.Path("/CargoInsertar").HandlerFunc(CargoInsertar).Name(
		"CargoInsertar")
	router.Path("/CargoActualizar/{codigo:[0-9]+}").HandlerFunc(CargoActualizar).Name(
		"CargoActualizar")
	router.Path("/CargoBorrar/{codigo:[0-9]+}").HandlerFunc(CargoBorrar).Name(
		"CargoBorrar")
	router.Path("/CargoEliminar/{codigo:[0-9]+}").HandlerFunc(
		CargoEliminar).Name("CargoEliminar")
	router.Path("/CargoEditar/{codigo:[0-9]+}").HandlerFunc(CargoEditar).Name(
		"CargoEditar")
	router.Path("/CargoPdf/{codigo:[0-9]+}").HandlerFunc(CargoPdf).Name(
		"CargoPdf")
	router.Path("/CargoTodosPdf").HandlerFunc(CargoTodosPdf).
		Name("CargoTodosPdf")
	router.Path("/CargoExcel").HandlerFunc(CargoExcel).
		Name("CargoExcel")

	// DOCUMENTO
	router.Path("/DocumentoNuevo/{panel}/{codigo}/{elemento}").HandlerFunc(DocumentoNuevo).Name("DocumentoNuevo")
	router.Path("/DocumentoLista").HandlerFunc(DocumentoLista).Name("DocumentoLista")
	router.Path("/DocumentoExiste/{codigo:[0-9]+}").HandlerFunc(DocumentoExiste).
		Name("DocumentoExiste")
	router.Path("/DocumentoActual").HandlerFunc(DocumentoActual).
		Name("DocumentoActual")
	router.Path("/DocumentoActualConsulta/{codigo}").HandlerFunc(DocumentoActualConsulta).
		Name("DocumentoActualConsulta")
	router.Path("/DocumentoBuscar/{codigo}").HandlerFunc(DocumentoBuscar).
		Name("DocumentoBuscar")
	router.Path("/DocumentoBuscarTesoreria/{codigo}").HandlerFunc(DocumentoBuscarTesoreria).
		Name("DocumentoBuscarTesoreria")

	router.Path("/DocumentoBuscarConsulta/{codigo}").HandlerFunc(DocumentoBuscarConsulta).
		Name("DocumentoBuscarConsulta")
	router.Path("/DocumentoInsertar").HandlerFunc(DocumentoInsertar).Name(
		"DocumentoInsertar")
	router.Path("/DocumentoActualizar/{codigo:[0-9]+}").HandlerFunc(DocumentoActualizar).Name(
		"DocumentoActualizar")
	router.Path("/DocumentoBorrar/{codigo:[0-9]+}").HandlerFunc(DocumentoBorrar).Name(
		"DocumentoBorrar")
	router.Path("/DocumentoEliminar/{codigo:[0-9]+}").HandlerFunc(
		DocumentoEliminar).Name("DocumentoEliminar")
	router.Path("/DocumentoEditar/{codigo:[0-9]+}").HandlerFunc(DocumentoEditar).Name(
		"DocumentoEditar")
	router.Path("/DocumentoPdf/{codigo:[0-9]+}").HandlerFunc(DocumentoPdf).
		Name("DocumentoPdf")

	router.Path("/DocumentoTodosPdf").HandlerFunc(DocumentoTodosPdf).
		Name("DocumentoTodosPdf")
	router.Path("/DocumentoExcel").HandlerFunc(DocumentoExcel).
		Name("DocumentoExcel")

	// DATOS PERIODO
	router.Path("/PeriodoNuevo/{panel}/{codigo}/{elemento}").HandlerFunc(PeriodoNuevo).Name("PeriodoNuevo")
	router.Path("/PeriodoLista").HandlerFunc(PeriodoLista).Name("PeriodoLista")
	router.Path("/PeriodoExiste/{codigo:[0-9]+}").HandlerFunc(PeriodoExiste).
		Name("PeriodoExiste")
	router.Path("/PeriodoActual").HandlerFunc(PeriodoActual).
		Name("PeriodoActual")
	router.Path("/PeriodoActualConsulta/{codigo}").HandlerFunc(PeriodoActualConsulta).
		Name("PeriodoActualConsulta")
	router.Path("/PeriodoBuscar/{codigo}").HandlerFunc(PeriodoBuscar).
		Name("PeriodoBuscar")
	router.Path("/PeriodoBuscarTesoreria/{codigo}").HandlerFunc(PeriodoBuscarTesoreria).
		Name("PeriodoBuscarTesoreria")

	router.Path("/PeriodoBuscarConsulta/{codigo}").HandlerFunc(PeriodoBuscarConsulta).
		Name("PeriodoBuscarConsulta")
	router.Path("/PeriodoInsertar").HandlerFunc(PeriodoInsertar).Name(
		"PeriodoInsertar")
	router.Path("/PeriodoActualizar/{codigo:[0-9]+}").HandlerFunc(PeriodoActualizar).Name(
		"PeriodoActualizar")
	router.Path("/PeriodoBorrar/{codigo:[0-9]+}").HandlerFunc(PeriodoBorrar).Name(
		"PeriodoBorrar")
	router.Path("/PeriodoEliminar/{codigo:[0-9]+}").HandlerFunc(
		PeriodoEliminar).Name("PeriodoEliminar")
	router.Path("/PeriodoEditar/{codigo:[0-9]+}").HandlerFunc(PeriodoEditar).Name(
		"PeriodoEditar")
	router.Path("/PeriodoPdf/{codigo:[0-9]+}").HandlerFunc(PeriodoPdf).
		Name("PeriodoPdf")

	router.Path("/PeriodoTodosPdf").HandlerFunc(PeriodoTodosPdf).
		Name("PeriodoTodosPdf")
	router.Path("/PeriodoExcel").HandlerFunc(PeriodoExcel).
		Name("PeriodoExcel")

	// ARCHIVO BODEGA
	router.Path("/BodegaNuevo/{panel}/{codigo}/{elemento}").HandlerFunc(BodegaNuevo).Name("BodegaNuevo")
	router.Path("/BodegaMovimiento/{codigo}").HandlerFunc(BodegaMovimiento).Name("BodegaMovimiento")

	router.Path("/BodegaLista").HandlerFunc(BodegaLista).Name("BodegaLista")
	router.Path("/BodegaExiste/{codigo}").HandlerFunc(BodegaExiste).
		Name("BodegaExiste")
	router.Path("/BodegaInsertar").HandlerFunc(BodegaInsertar).Name(
		"BodegaInsertar")
	router.Path("/BodegaActualizar/{codigo:[0-9]+}").HandlerFunc(BodegaActualizar).Name(
		"BodegaActualizar")
	router.Path("/BodegaBorrar/{codigo:[0-9]+}").HandlerFunc(BodegaBorrar).Name(
		"BodegaBorrar")
	router.Path("/BodegaEliminar/{codigo:[0-9]+}").HandlerFunc(
		BodegaEliminar).Name("BodegaEliminar")
	router.Path("/BodegaEditar/{codigo:[0-9]+}").HandlerFunc(BodegaEditar).Name(
		"BodegaEditar")
	router.Path("/BodegaPdf/{codigo:[0-9]+}").HandlerFunc(BodegaPdf).Name(
		"BodegaPdf")
	router.Path("/BodegaBuscar/{codigo}").HandlerFunc(BodegaBuscar).
		Name("BodegaBuscar")
	router.Path("/BodegaActual").HandlerFunc(BodegaActual).
		Name("BodegaActual")

	router.Path("/BodegaBuscarCrear/{codigo}").HandlerFunc(BodegaBuscarCrear).
		Name("BodegaBuscarCrear")

	router.Path("/BodegaTodosPdf").HandlerFunc(BodegaTodosPdf).
		Name("BodegaTodosPdf")
	router.Path("/BodegaExcel").HandlerFunc(BodegaExcel).
		Name("BodegaExcel")

	// ARCHIVO GRUPO
	router.Path("/GrupoNuevo/{panel}/{codigo}/{elemento}").HandlerFunc(GrupoNuevo).Name("GrupoNuevo")
	router.Path("/GrupoLista").HandlerFunc(GrupoLista).Name("GrupoLista")
	router.Path("/GrupoExiste/{codigo:[0-9]+}").HandlerFunc(GrupoExiste).
		Name("GrupoExiste")
	router.Path("/GrupoInsertar").HandlerFunc(GrupoInsertar).Name(
		"GrupoInsertar")
	router.Path("/GrupoActualizar/{codigo:[0-9]+}").HandlerFunc(GrupoActualizar).Name(
		"GrupoActualizar")
	router.Path("/GrupoBorrar/{codigo:[0-9]+}").HandlerFunc(GrupoBorrar).Name(
		"GrupoBorrar")
	router.Path("/GrupoEliminar/{codigo:[0-9]+}").HandlerFunc(
		GrupoEliminar).Name("GrupoEliminar")
	router.Path("/GrupoEditar/{codigo:[0-9]+}").HandlerFunc(GrupoEditar).Name(
		"GrupoEditar")
	router.Path("/GrupoPdf/{codigo:[0-9]+}").HandlerFunc(GrupoPdf).Name(
		"GrupoPdf")

	router.Path("/GrupoBuscar/{codigo}").HandlerFunc(GrupoBuscar).
		Name("GrupoBuscar")

	router.Path("/GrupoActual").HandlerFunc(GrupoActual).
		Name("GrupoActual")

	router.Path("/GrupoBuscarCrear/{codigo}").HandlerFunc(GrupoBuscarCrear).
		Name("GrupoBuscarCrear")

	router.Path("/GrupoTodosPdf").HandlerFunc(GrupoTodosPdf).
		Name("GrupoTodosPdf")
	router.Path("/GrupoExcel").HandlerFunc(GrupoExcel).
		Name("GrupoExcel")

	// TERCERO
	router.Path("/TerceroLista").HandlerFunc(TerceroLista).Name("TerceroLista")
	router.Path("/TerceroNuevo/{panel}/{codigo}/{elemento}").HandlerFunc(TerceroNuevo).Name("TerceroNuevo")
	router.Path("/TerceroNuevoCopia/{copiacodigo}").HandlerFunc(TerceroNuevoCopia).Name("TerceroNuevoCopia")
	router.Path("/TerceroBuscar/{codigo}").HandlerFunc(TerceroBuscar).
		Name("TerceroBuscar")
	router.Path("/TerceroActual/{codigo}").HandlerFunc(TerceroActual).
		Name("TerceroActual")

	router.Path("/TerceroActualConsulta").HandlerFunc(TerceroActualConsulta).
		Name("TerceroActualConsulta")

	router.Path("/TerceroActualBanco/{codigo}").HandlerFunc(TerceroActualBanco).
		Name("TerceroActualBanco")
	router.Path("/TerceroExiste/{codigo}").HandlerFunc(TerceroExiste).
		Name("TerceroExiste")
	router.Path("/TerceroEditar/{codigo:[0-9]+}").HandlerFunc(TerceroEditar).
		Name("TerceroEditar")
	router.Path("/TerceroActualizar/{codigo:[0-9]+}").HandlerFunc(
		TerceroActualizar).Name("TerceroActualizar")
	router.Path("/TerceroInsertar").HandlerFunc(TerceroInsertar).Name(
		"TerceroInsertar")
	router.Path("/TerceroBorrar/{codigo:[0-9]+}").HandlerFunc(TerceroBorrar).
		Name("TerceroBorrar")
	router.Path("/TerceroEliminar/{codigo:[0-9]+}").HandlerFunc(
		TerceroEliminar).Name("TerceroEliminar")
	router.Path("/TerceroPdf/{codigo}").HandlerFunc(TerceroPdf).Name(
		"TerceroPdf")
	router.Path("/TerceroTodosPdf").HandlerFunc(TerceroTodosPdf).
		Name("TerceroTodosPdf")
	router.Path("/TerceroExcel").HandlerFunc(TerceroExcel).
		Name("TerceroExcel")

	// DATOS EMPLEADOS
	router.Path("/EmpleadoLista").HandlerFunc(EmpleadoLista).Name("EmpleadoLista")
	router.Path("/EmpleadoNuevo/{panel}/{codigo}/{elemento}").HandlerFunc(EmpleadoNuevo).Name("EmpleadoNuevo")
	router.Path("/EmpleadoNuevoCopia/{copiacodigo}").HandlerFunc(EmpleadoNuevoCopia).Name("EmpleadoNuevoCopia")
	router.Path("/EmpleadoBuscar/{codigo}").HandlerFunc(EmpleadoBuscar).
		Name("EmpleadoBuscar")
	router.Path("/EmpleadoActual/{codigo}").HandlerFunc(EmpleadoActual).
		Name("EmpleadoActual")
	//router.Path("/EmpleadoActualBanco/{codigo}").HandlerFunc(EmpleadoActualBanco).
	//Name("EmpleadoActualBanco")
	router.Path("/EmpleadoExiste/{codigo}").HandlerFunc(EmpleadoExiste).
		Name("EmpleadoExiste")
	router.Path("/EmpleadoEditar/{codigo}").HandlerFunc(EmpleadoEditar).
		Name("EmpleadoEditar")
	router.Path("/EmpleadoActualizar/{codigo:[0-9]+}").HandlerFunc(
		EmpleadoActualizar).Name("EmpleadoActualizar")
	router.Path("/EmpleadoInsertar").HandlerFunc(EmpleadoInsertar).Name(
		"EmpleadoInsertar")
	router.Path("/EmpleadoBorrar/{codigo:[0-9]+}").HandlerFunc(EmpleadoBorrar).
		Name("EmpleadoBorrar")
	router.Path("/EmpleadoEliminar/{codigo:[0-9]+}").HandlerFunc(
		EmpleadoEliminar).Name("EmpleadoEliminar")
	router.Path("/EmpleadoPdf/{codigo}").HandlerFunc(EmpleadoPdf).Name(
		"EmpleadoPdf")

	router.Path("/EmpleadoTodosPdf").HandlerFunc(EmpleadoTodosPdf).
		Name("EmpleadoTodosPdf")
	router.Path("/EmpleadoExcel").HandlerFunc(EmpleadoExcel).
		Name("EmpleadoExcel")

	// DATOS RESIDENTES
	router.Path("/ResidenteLista").HandlerFunc(ResidenteLista).Name("ResidenteLista")
	router.Path("/ResidenteNuevo/{panel}/{codigo}/{elemento}").HandlerFunc(ResidenteNuevo).Name("ResidenteNuevo")
	router.Path("/ResidenteNuevoCopia/{copiacodigo}").HandlerFunc(ResidenteNuevoCopia).Name("ResidenteNuevoCopia")
	router.Path("/ResidenteBuscar/{codigo}").HandlerFunc(ResidenteBuscar).
		Name("ResidenteBuscar")
	router.Path("/ResidenteActual/{codigo}").HandlerFunc(ResidenteActual).
		Name("ResidenteActual")
	//router.Path("/ResidenteActualBanco/{codigo}").HandlerFunc(ResidenteActualBanco).
	//Name("ResidenteActualBanco")
	router.Path("/ResidenteExiste/{codigo}").HandlerFunc(ResidenteExiste).
		Name("ResidenteExiste")
	router.Path("/ResidenteEditar/{codigo}").HandlerFunc(ResidenteEditar).
		Name("ResidenteEditar")
	router.Path("/ResidenteActualizar").HandlerFunc(
		ResidenteActualizar).Name("ResidenteActualizar")
	router.Path("/ResidenteInsertar").HandlerFunc(ResidenteInsertar).Name(
		"ResidenteInsertar")
	router.Path("/ResidenteBorrar/{codigo:[0-9]+}").HandlerFunc(ResidenteBorrar).
		Name("ResidenteBorrar")
	router.Path("/ResidenteEliminar/{codigo:[0-9]+}").HandlerFunc(
		ResidenteEliminar).Name("ResidenteEliminar")

	router.Path("/ResidentePdf/{codigo}").HandlerFunc(ResidentePdf).Name(
		"ResidentePdf")
	router.Path("/ResidenteTodosPdf").HandlerFunc(ResidenteTodosPdf).
		Name("ResidenteTodosPdf")
	router.Path("/ResidenteExcel").HandlerFunc(ResidenteExcel).
		Name("ResidenteExcel")

	// NOMINA DATOS
	router.Path("/NominaLista").HandlerFunc(NominaLista).Name("NominaLista")
	router.Path("/NominaNuevo/{codigo}").HandlerFunc(NominaNuevo).Name("NominaNuevo")
	router.Path("/NominaAgregar").HandlerFunc(NominaAgregar).Name("NominaAgregar")
	router.Path("/NominaEditar/{codigo}").HandlerFunc(NominaEditar).
		Name("NominaEditar")
	router.Path("/NominaBorrar/{codigo}").HandlerFunc(NominaBorrar).
		Name("NominaBorrar")
	router.Path("/NominaEliminar/{codigo}").HandlerFunc(
		NominaEliminar).Name("NominaEliminar")
	router.Path("/NominaPdf/{numero}").HandlerFunc(NominaPdf).Name(
		"NominaPdf")
	router.Path("/NominaTodosPdf").HandlerFunc(NominaTodosPdf).
		Name("NominaTodosPdf")
	router.Path("/NominaIndividualExcel/{codigo}").HandlerFunc(NominaIndividualExcel).
		Name("NominaIndividualExcel")
	router.Path("/NominaExcel").HandlerFunc(NominaExcel).
		Name("NominaExcel")

	// PRESTACIONES SOCIALES DATOS
	router.Path("/NominaprestacionesLista").HandlerFunc(NominaprestacionesLista).Name("NominaprestacionesLista")
	router.Path("/NominaprestacionesNuevo/{codigo}").HandlerFunc(NominaprestacionesNuevo).Name("NominaprestacionesNuevo")
	router.Path("/NominaprestacionesAgregar").HandlerFunc(NominaprestacionesAgregar).Name("NominaprestacionesAgregar")
	router.Path("/NominaprestacionesEditar/{codigo}").HandlerFunc(NominaprestacionesEditar).
		Name("NominaprestacionesEditar")
	router.Path("/NominaprestacionesBorrar/{codigo}").HandlerFunc(NominaprestacionesBorrar).
		Name("NominaprestacionesBorrar")
	router.Path("/NominaprestacionesEliminar/{codigo}").HandlerFunc(
		NominaprestacionesEliminar).Name("NominaprestacionesEliminar")
	router.Path("/NominaprestacionesPdf/{numero}").HandlerFunc(NominaprestacionesPdf).Name(
		"NominaprestacionesPdf")
	router.Path("/NominaprestacionesTodosPdf").HandlerFunc(NominaprestacionesTodosPdf).
		Name("NominaprestacionesTodosPdf")
	router.Path("/NominaprestacionesIndividualExcel/{codigo}").HandlerFunc(NominaprestacionesIndividualExcel).
		Name("NominaprestacionesIndividualExcel")

	// NOMINA APORTES DATOS
	router.Path("/NominaaportesLista").HandlerFunc(NominaaportesLista).Name("NominaaportesLista")
	router.Path("/NominaaportesNuevo/{codigo}").HandlerFunc(NominaaportesNuevo).Name("NominaaportesNuevo")
	router.Path("/NominaaportesAgregar").HandlerFunc(NominaaportesAgregar).Name("NominaaportesAgregar")
	router.Path("/NominaaportesEditar/{codigo}").HandlerFunc(NominaaportesEditar).
		Name("NominaaportesEditar")
	router.Path("/NominaaportesBorrar/{codigo}").HandlerFunc(NominaaportesBorrar).
		Name("NominaaportesBorrar")
	router.Path("/NominaaportesEliminar/{codigo}").HandlerFunc(
		NominaaportesEliminar).Name("NominaaportesEliminar")
	router.Path("/NominaaportesPdf/{numero}").HandlerFunc(NominaaportesPdf).Name(
		"NominaaportesPdf")
	router.Path("/NominaaportesTodosPdf").HandlerFunc(NominaaportesTodosPdf).
		Name("NominaaportesTodosPdf")
	router.Path("/NominaaportesIndividualExcel/{codigo}").HandlerFunc(NominaaportesIndividualExcel).
		Name("NominaaportesIndividualExcel")

	// CERTIFICADO EMPLEADOS
	router.Path("/CertificadoempleadoLista").HandlerFunc(CertificadoempleadoLista).Name("CertificadoempleadoLista")
	router.Path("/CertificadoempleadoDatos").HandlerFunc(CertificadoempleadoDatos).Name("CertificadoempleadoDatos")
	router.Path("/Certificadoempleado/{codigo}/{fechaexpedicion}/{fechainicial}/{fechafinal}").HandlerFunc(Certificadoempleado).Name(
		"Certificadoempleado")

	// CERTIFICADO RETENCION
	router.Path("/CertificadoretencionLista").HandlerFunc(CertificadoretencionLista).Name("CertificadoretencionLista")
	router.Path("/CertificadoretencionDatos").HandlerFunc(CertificadoretencionDatos).Name("CertificadoretencionDatos")
	router.Path("/Certificadoretencion/{codigo}/{cuentainicial}/{cuentafinal}/{fechaexpedicion}/{fechainicial}/{fechafinal}").HandlerFunc(CertificadoretencionPdf).Name(
		"Certificadoretencion")
	router.Path("/CertificadoRetencionTodosPdf/{cuentainicial}/{cuentafinal}/{fechaexpedicion}/{fechainicial}/{fechafinal}/{terceroinicial}/{tercerofinal}").HandlerFunc(CertificadoRetencionTodosPdf).Name("CertificadoretencionDatos")

	// CERTIFICADO SALARIOS
	router.Path("/CertificadosalariosLista").HandlerFunc(CertificadosalariosLista).Name("CertificadosalariosLista")
	router.Path("/CertificadosalariosDatos").HandlerFunc(CertificadosalariosDatos).Name("CertificadosalariosDatos")
	router.Path("/Certificadosalarios/{codigo}/{cuentainicial}/{cuentafinal}/{fechaexpedicion}/{fechainicial}/{fechafinal}").HandlerFunc(CertificadosalariosPdf).Name(
		"Certificadosalarios")
	router.Path("/CertificadoSalariosTodosPdf/{cuentainicial}/{cuentafinal}/{fechaexpedicion}/{fechainicial}/{fechafinal}/{terceroinicial}/{tercerofinal}").HandlerFunc(CertificadoSalariosTodosPdf).Name("CertificadosalariosDatos")

	// cierrecontable

	router.Path("/CierrecontableLista").HandlerFunc(CierrecontableLista).Name("CierrecontableLista")
	router.Path("/CierrecontableDatos/{mes}/{centro}").HandlerFunc(CierrecontableDatos).Name("CierrecontableLista")

	//router.Path("/CertificadoretencionDatos").HandlerFunc(CertificadoretencionDatos).Name("CertificadoretencionDatos")

	// EXOGENA
	router.Path("/ExogenaLista/{panel}/{codigo}/{elemento}").HandlerFunc(ExogenaLista).Name("ExogenaLista")
	router.Path("/ExogenaAgregar").HandlerFunc(ExogenaAgregar).Name(
		"ExogenaAgregar")
	router.Path("/ExogenaEditar/{numero}").HandlerFunc(ExogenaEditar).
		Name("ExogenaEditar")
	// datos
	router.Path("/ExogenaListaDatos/{numero}").HandlerFunc(ExogenaListaDatos).
		Name("ExogenaListaDatos")

	router.Path("/TraerConceptoExogena/{formato}").HandlerFunc(TraerConceptoExogena).
		Name("TraerConceptoExogena")

	router.Path("/TraerColumnaExogena/{formato}").HandlerFunc(TraerColumnaExogena).
		Name("TraerColumnaExogena")

	router.Path("/TraerFormatoExogena").HandlerFunc(TraerFormatoExogena).
		Name("ExogenaConcepto")

	router.Path("/FormatoExogenaBuscar/{codigo}").HandlerFunc(FormatoExogenaBuscar).
		Name("FormatoExogenaBuscar")

	router.Path("/FormatoExogenaActual").HandlerFunc(FormatoExogenaActual).
		Name("FormatoExogenaActual")

	router.Path("/ConceptoExogenaBuscar/{formato}/{codigo}").HandlerFunc(ConceptoExogenaBuscar).
		Name("ConceptoExogenaBuscar")

	router.Path("/ConceptoExogenaActual").HandlerFunc(ConceptoExogenaActual).
		Name("ConceptoExogenaActual")

	router.Path("/ColumnaExogenaBuscar/{formato}/{codigo}").HandlerFunc(ColumnaExogenaBuscar).
		Name("ColumnaExogenaBuscar")

	router.Path("/ColumnaExogenaActual").HandlerFunc(ColumnaExogenaActual).
		Name("ColumnaExogenaActual")

	//FORMATOS EXOGENA
	router.Path("/ExogenaGenerar/{panel}").HandlerFunc(ExogenaGenerar).Name("ExogenaGenerar")
	router.Path("/Formato1001").HandlerFunc(Formato1001).Name("ExogenaGenerarFormato")
	router.Path("/Formato1003").HandlerFunc(Formato1003).Name("ExogenaGenerarFormato")
	router.Path("/Formato1004").HandlerFunc(Formato1004).Name("ExogenaGenerarFormato")
	router.Path("/Formato1005").HandlerFunc(Formato1005).Name("ExogenaGenerarFormato")
	router.Path("/Formato1006").HandlerFunc(Formato1006).Name("ExogenaGenerarFormato")
	router.Path("/Formato1007").HandlerFunc(Formato1007).Name("ExogenaGenerarFormato")
	router.Path("/Formato1008").HandlerFunc(Formato1008).Name("ExogenaGenerarFormato")
	router.Path("/Formato1009").HandlerFunc(Formato1009).Name("ExogenaGenerarFormato")
	router.Path("/Formato1010").HandlerFunc(Formato1010).Name("ExogenaGenerarFormato")
	router.Path("/Formato1011").HandlerFunc(Formato1011).Name("ExogenaGenerarFormato")
	router.Path("/Formato1012").HandlerFunc(Formato1012).Name("ExogenaGenerarFormato")
	router.Path("/Formato1647").HandlerFunc(Formato1647).Name("ExogenaGenerarFormato")
	router.Path("/Formato2275").HandlerFunc(Formato2275).Name("ExogenaGenerarFormato")
	router.Path("/Formato2276").HandlerFunc(Formato2276).Name("ExogenaGenerarFormato")

	// EMPRESA
	router.Path("/EmpresaLista").HandlerFunc(EmpresaLista).Name("EmpresaLista")
	router.Path("/EmpresaNuevo/{codigo}").HandlerFunc(EmpresaNuevo).Name("EmpresaNuevo")
	router.Path("/EmpresaBuscar/{codigo}").HandlerFunc(EmpresaBuscar).
		Name("EmpresaBuscar")
	router.Path("/EmpresaActual/{codigo}").HandlerFunc(EmpresaActual).
		Name("EmpresaActual")
	router.Path("/EmpresaExiste/{codigo:[0-9]+}").HandlerFunc(EmpresaExiste).
		Name("EmpresaExiste")
	router.Path("/EmpresaEditar/{codigo:[0-9]+}").HandlerFunc(EmpresaEditar).
		Name("EmpresaEditar")
	router.Path("/EmpresaActualizar/{codigo:[0-9]+}").HandlerFunc(
		EmpresaActualizar).Name("EmpresaActualizar")
	router.Path("/EmpresaInsertar").HandlerFunc(EmpresaInsertar).Name(
		"EmpresaInsertar")
	router.Path("/EmpresaBorrar/{codigo:[0-9]+}").HandlerFunc(EmpresaBorrar).
		Name("EmpresaBorrar")
	router.Path("/EmpresaEliminar/{codigo:[0-9]+}").HandlerFunc(
		EmpresaEliminar).Name("EmpresaEliminar")
	router.Path("/EmpresaPdf/{codigo:[0-9]+}").HandlerFunc(EmpresaPdf).Name(
		"EmpresaPdf")

	router.Path("/EmpresaTodosPdf").HandlerFunc(EmpresaTodosPdf).
		Name("EmpresaTodosPdf")
	router.Path("/EmpresaExcel").HandlerFunc(EmpresaExcel).
		Name("EmpresaExcel")

	// USUARIO
	router.Path("/UsuarioLista").HandlerFunc(UsuarioLista).Name("UsuarioLista")
	router.Path("/UsuarioNuevo/{codigo}").HandlerFunc(UsuarioNuevo).Name("UsuarioNuevo")
	router.Path("/UsuarioBuscar/{codigo}").HandlerFunc(UsuarioBuscar).
		Name("UsuarioBuscar")
	router.Path("/UsuarioActual/{codigo}").HandlerFunc(UsuarioActual).
		Name("UsuarioActual")
	router.Path("/UsuarioExiste/{codigo:[0-9]+}").HandlerFunc(UsuarioExiste).
		Name("UsuarioExiste")
	router.Path("/UsuarioEditar/{codigo:[0-9]+}").HandlerFunc(UsuarioEditar).
		Name("UsuarioEditar")
	router.Path("/UsuarioActualizar/{codigo:[0-9]+}").HandlerFunc(
		UsuarioActualizar).Name("UsuarioActualizar")
	router.Path("/UsuarioInsertar").HandlerFunc(UsuarioInsertar).Name(
		"UsuarioInsertar")
	router.Path("/UsuarioBorrar/{codigo:[0-9]+}").HandlerFunc(UsuarioBorrar).
		Name("UsuarioBorrar")
	router.Path("/UsuarioEliminar/{codigo:[0-9]+}").HandlerFunc(
		UsuarioEliminar).Name("UsuarioEliminar")
	router.Path("/UsuarioPdf/{codigo}").HandlerFunc(UsuarioPdf).Name(
		"UsuarioPdf")

	router.Path("/UsuarioTodosPdf").HandlerFunc(UsuarioTodosPdf).
		Name("UsuarioTodosPdf")
	router.Path("/UsuarioExcel").HandlerFunc(UsuarioExcel).
		Name("UsuarioaExcel")

	// ARCHIVO SUBGRUPO
	router.Path("/SubgrupoNuevo/{panel}/{codigo}/{elemento}").HandlerFunc(SubgrupoNuevo).Name("SubgrupoNuevo")
	router.Path("/SubgrupoLista").HandlerFunc(SubgrupoLista).Name("SubgrupoLista")
	router.Path("/SubgrupoExiste/{codigo:[0-9]+}").HandlerFunc(SubgrupoExiste).
		Name("SubgrupoExiste")
	router.Path("/SubgrupoActual").HandlerFunc(SubgrupoActual).
		Name("SubgrupoActual")

	router.Path("/SubgrupoInsertar").HandlerFunc(SubgrupoInsertar).Name(
		"SubgrupoInsertar")
	router.Path("/SubgrupoActualizar/{codigo:[0-9]+}").HandlerFunc(SubgrupoActualizar).Name(
		"SubgrupoActualizar")
	router.Path("/SubgrupoBorrar/{codigo:[0-9]+}").HandlerFunc(SubgrupoBorrar).Name(
		"SubgrupoBorrar")
	router.Path("/SubgrupoEliminar/{codigo:[0-9]+}").HandlerFunc(
		SubgrupoEliminar).Name("SubgrupoEliminar")
	router.Path("/SubgrupoEditar/{codigo:[0-9]+}").HandlerFunc(SubgrupoEditar).Name(
		"SubgrupoEditar")
	router.Path("/SubgrupoPdf/{codigo:[0-9]+}").HandlerFunc(SubgrupoPdf).Name(
		"SubgrupoPdf")

	router.Path("/SubgrupoBuscar/{codigo}").HandlerFunc(SubgrupoBuscar).
		Name("SubgrupoBuscar")

	router.Path("/SubgrupoTodosPdf").HandlerFunc(SubgrupoTodosPdf).
		Name("SubgrupoTodosPdf")
	router.Path("/SubgrupoExcel").HandlerFunc(SubgrupoExcel).
		Name("SubgrupoExcel")

	// VENDEDOR
	router.Path("/VendedorLista").HandlerFunc(VendedorLista).Name("VendedorLista")
	router.Path("/VendedorNuevo/{panel}/{codigo}/{elemento}").HandlerFunc(VendedorNuevo).Name("VendedorNuevo")
	router.Path("/VendedorBuscar/{codigo}").HandlerFunc(VendedorBuscar).
		Name("VendedorBuscar")
	router.Path("/VendedorActual").HandlerFunc(VendedorActual).
		Name("VendedorActual")
	router.Path("/VendedorExiste/{codigo:[0-9]+}").HandlerFunc(VendedorExiste).
		Name("VendedorExiste")
	router.Path("/VendedorEditar/{codigo:[0-9]+}").HandlerFunc(VendedorEditar).
		Name("VendedorEditar")
	router.Path("/VendedorActualizar/{codigo:[0-9]+}").HandlerFunc(
		VendedorActualizar).Name("VendedorActualizar")
	router.Path("/VendedorInsertar").HandlerFunc(VendedorInsertar).Name(
		"VendedorInsertar")
	router.Path("/VendedorBorrar/{codigo:[0-9]+}").HandlerFunc(VendedorBorrar).
		Name("VendedorBorrar")
	router.Path("/VendedorEliminar/{codigo:[0-9]+}").HandlerFunc(
		VendedorEliminar).Name("VendedorEliminar")
	router.Path("/VendedorPdf/{codigo}").HandlerFunc(VendedorPdf).Name(
		"VendedorPdf")

	router.Path("/VendedorBuscarCrear/{codigo}").HandlerFunc(VendedorBuscarCrear).
		Name("VendedorBuscarCrear")

	router.Path("/VendedorTodosPdf").HandlerFunc(VendedorTodosPdf).
		Name("VendedorTodosPdf")
	router.Path("/VendedorExcel").HandlerFunc(VendedorExcel).
		Name("VendedorExcel")

	// ALMACENISTA
	router.Path("/AlmacenistaLista").HandlerFunc(AlmacenistaLista).Name("AlmacenistaLista")
	router.Path("/AlmacenistaNuevo/{panel}/{codigo}/{elemento}").HandlerFunc(AlmacenistaNuevo).Name("AlmacenistaNuevo")
	router.Path("/AlmacenistaBuscar/{codigo}").HandlerFunc(AlmacenistaBuscar).
		Name("AlmacenistaBuscar")
	router.Path("/AlmacenistaActual").HandlerFunc(AlmacenistaActual).
		Name("AlmacenistaActual")
	router.Path("/AlmacenistaExiste/{codigo:[0-9]+}").HandlerFunc(AlmacenistaExiste).
		Name("AlmacenistaExiste")
	router.Path("/AlmacenistaEditar/{codigo:[0-9]+}").HandlerFunc(AlmacenistaEditar).
		Name("AlmacenistaEditar")
	router.Path("/AlmacenistaActualizar/{codigo:[0-9]+}").HandlerFunc(
		AlmacenistaActualizar).Name("AlmacenistaActualizar")
	router.Path("/AlmacenistaInsertar").HandlerFunc(AlmacenistaInsertar).Name(
		"AlmacenistaInsertar")
	router.Path("/AlmacenistaBorrar/{codigo:[0-9]+}").HandlerFunc(AlmacenistaBorrar).
		Name("AlmacenistaBorrar")
	router.Path("/AlmacenistaEliminar/{codigo:[0-9]+}").HandlerFunc(
		AlmacenistaEliminar).Name("AlmacenistaEliminar")
	router.Path("/AlmacenistaPdf/{codigo}").HandlerFunc(AlmacenistaPdf).Name(
		"AlmacenistaPdf")

	router.Path("/AlmacenistaBuscarCrear/{codigo}").HandlerFunc(AlmacenistaBuscarCrear).
		Name("AlmacenistaBuscarCrear")

	router.Path("/AlmacenistaTodosPdf").HandlerFunc(AlmacenistaTodosPdf).
		Name("AlmacenistaTodosPdf")
	router.Path("/AlmacenistaExcel").HandlerFunc(AlmacenistaExcel).
		Name("AlmacenistaExcel")

	// PRODUCTO
	router.Path("/ProductoLista").HandlerFunc(ProductoLista).Name("ProductoLista")
	router.Path("/ProductoNuevo/{panel}/{codigo}/{elemento}").HandlerFunc(ProductoNuevo).Name("ProductoNuevo")

	router.Path("/ProductoNuevoCopiar/{codigocopia}").HandlerFunc(ProductoNuevoCopiar).Name("ProductoNuevo")

	router.Path("/ProductoBuscar/{codigo}").HandlerFunc(ProductoBuscar).
		Name("ProductoBuscar")
	router.Path("/ProductoActual/{codigo}").HandlerFunc(ProductoActual).
		Name("ProductoActual")
	router.Path("/ProductoExiste/{codigo:[0-9]+}").HandlerFunc(ProductoExiste).
		Name("ProductoExiste")
	router.Path("/ProductoEditar/{codigo:[0-9]+}").HandlerFunc(ProductoEditar).
		Name("ProductoEditar")
	router.Path("/ProductoActualizar/{codigo:[0-9]+}").HandlerFunc(
		ProductoActualizar).Name("ProductoActualizar")
	router.Path("/ProductoInsertar").HandlerFunc(ProductoInsertar).Name(
		"ProductoInsertar")
	router.Path("/ProductoBorrar/{codigo:[0-9]+}").HandlerFunc(ProductoBorrar).
		Name("ProductoBorrar")
	router.Path("/ProductoEliminar/{codigo:[0-9]+}").HandlerFunc(
		ProductoEliminar).Name("ProductoEliminar")
	router.Path("/ProductoPdf/{codigo}").HandlerFunc(ProductoPdf).Name(
		"ProductoPdf")

	router.Path("/ProductoBuscarCrear/{codigo}").HandlerFunc(ProductoBuscarCrear).
		Name("ProductoBuscarCrear")

	router.Path("/ProductoTodosPdf").HandlerFunc(ProductoTodosPdf).
		Name("ProductoTodosPdf")
	router.Path("/ProductoExcel").HandlerFunc(ProductoExcel).
		Name("ProductoExcel")

	// ARCHIVO RESOLUCION VENTA
	router.Path("/ResolucionventaNuevo/{codigo}").HandlerFunc(ResolucionventaNuevo).Name("ResolucionventaNuevo")
	router.Path("/ResolucionventaLista").HandlerFunc(ResolucionventaLista).Name("ResolucionventaLista")
	router.Path("/ResolucionventaExiste/{codigo:[0-9]+}").HandlerFunc(ResolucionventaExiste).
		Name("ResolucionventaExiste")
	router.Path("/ResolucionventaInsertar").HandlerFunc(ResolucionventaInsertar).Name(
		"ResolucionventaInsertar")
	router.Path("/ResolucionventaActualizar/{codigo:[0-9]+}").HandlerFunc(ResolucionventaActualizar).Name(
		"ResolucionventaActualizar")
	router.Path("/ResolucionventaBorrar/{codigo:[0-9]+}").HandlerFunc(ResolucionventaBorrar).Name(
		"ResolucionventaBorrar")
	router.Path("/ResolucionventaEliminar/{codigo:[0-9]+}").HandlerFunc(
		ResolucionventaEliminar).Name("ResolucionventaEliminar")
	router.Path("/ResolucionventaEditar/{codigo:[0-9]+}").HandlerFunc(ResolucionventaEditar).Name(
		"ResolucionEditar")
	router.Path("/ResolucionventaPdf/{codigo:[0-9]+}").HandlerFunc(ResolucionventaPdf).Name(
		"ResolucionventaPdf")

	router.Path("/ResolucionventaTodosPdf").HandlerFunc(ResolucionventaTodosPdf).
		Name("ResolucionventaTodosPdf")
	router.Path("/ResolucionventaExcel").HandlerFunc(ResolucionventaExcel).
		Name("ResolucionventaExcel")

	// ARCHIVO RESOLUCION SOPORTE
	router.Path("/ResolucionsoporteNuevo/{codigo}").HandlerFunc(ResolucionsoporteNuevo).Name("ResolucionsoporteNuevo")
	router.Path("/ResolucionsoporteLista").HandlerFunc(ResolucionsoporteLista).Name("ResolucionsoporteLista")
	router.Path("/ResolucionsoporteExiste/{codigo:[0-9]+}").HandlerFunc(ResolucionsoporteExiste).
		Name("ResolucionsoporteExiste")
	router.Path("/ResolucionsoporteInsertar").HandlerFunc(ResolucionsoporteInsertar).Name(
		"ResolucionsoporteInsertar")
	router.Path("/ResolucionsoporteActualizar/{codigo:[0-9]+}").HandlerFunc(ResolucionsoporteActualizar).Name(
		"ResolucionsoporteActualizar")
	router.Path("/ResolucionsoporteBorrar/{codigo:[0-9]+}").HandlerFunc(ResolucionsoporteBorrar).Name(
		"ResolucionsoporteBorrar")
	router.Path("/ResolucionsoporteEliminar/{codigo:[0-9]+}").HandlerFunc(
		ResolucionsoporteEliminar).Name("ResolucionsoporteEliminar")
	router.Path("/ResolucionsoporteEditar/{codigo:[0-9]+}").HandlerFunc(ResolucionsoporteEditar).Name(
		"ResolucionEditar")
	router.Path("/ResolucionsoportePdf/{codigo:[0-9]+}").HandlerFunc(ResolucionsoportePdf).Name(
		"ResolucionsoportePdf")

	router.Path("/ResolucionsoporteTodosPdf").HandlerFunc(ResolucionsoporteTodosPdf).
		Name("ResolucionsoporteTodosPdf")
	router.Path("/ResolucionsoporteExcel").HandlerFunc(ResolucionsoporteExcel).
		Name("ResolucionsoporteExcel")

	// ARCHIVO RESOLUCION NOMINA
	router.Path("/ResolucionnominaNuevo/{codigo}").HandlerFunc(ResolucionnominaNuevo).Name("ResolucionnominaNuevo")
	router.Path("/ResolucionnominaLista").HandlerFunc(ResolucionnominaLista).Name("ResolucionnominaLista")
	router.Path("/ResolucionnominaExiste/{codigo:[0-9]+}").HandlerFunc(ResolucionnominaExiste).
		Name("ResolucionnominaExiste")
	router.Path("/ResolucionnominaInsertar").HandlerFunc(ResolucionnominaInsertar).Name(
		"ResolucionnominaInsertar")
	router.Path("/ResolucionnominaActualizar/{codigo:[0-9]+}").HandlerFunc(ResolucionnominaActualizar).Name(
		"ResolucionnominaActualizar")
	router.Path("/ResolucionnominaBorrar/{codigo:[0-9]+}").HandlerFunc(ResolucionnominaBorrar).Name(
		"ResolucionnominaBorrar")
	router.Path("/ResolucionnominaEliminar/{codigo:[0-9]+}").HandlerFunc(
		ResolucionnominaEliminar).Name("ResolucionnominaEliminar")
	router.Path("/ResolucionnominaEditar/{codigo:[0-9]+}").HandlerFunc(ResolucionnominaEditar).Name(
		"ResolucionEditar")
	router.Path("/ResolucionnominaPdf/{codigo:[0-9]+}").HandlerFunc(ResolucionnominaPdf).Name(
		"ResolucionnominaPdf")

	router.Path("/ResolucionnominaTodosPdf").HandlerFunc(ResolucionnominaTodosPdf).
		Name("ResolucionnominaTodosPdf")
	router.Path("/ResolucionnominaExcel").HandlerFunc(ResolucionnominaExcel).
		Name("ResolucionnominaExcel")

	// ARCHIVO RESOLUCION NOMINA
	router.Path("/ResolucionnominaNuevo/{codigo}").HandlerFunc(ResolucionnominaNuevo).Name("ResolucionnominaNuevo")
	router.Path("/ResolucionnominaLista").HandlerFunc(ResolucionnominaLista).Name("ResolucionnominaLista")
	router.Path("/ResolucionnominaExiste/{codigo:[0-9]+}").HandlerFunc(ResolucionnominaExiste).
		Name("ResolucionnominaExiste")
	router.Path("/ResolucionnominaInsertar").HandlerFunc(ResolucionnominaInsertar).Name(
		"ResolucionnominaInsertar")
	router.Path("/ResolucionnominaActualizar/{codigo:[0-9]+}").HandlerFunc(ResolucionnominaActualizar).Name(
		"ResolucionnominaActualizar")
	router.Path("/ResolucionnominaBorrar/{codigo:[0-9]+}").HandlerFunc(ResolucionnominaBorrar).Name(
		"ResolucionnominaBorrar")
	router.Path("/ResolucionnominaEliminar/{codigo:[0-9]+}").HandlerFunc(
		ResolucionnominaEliminar).Name("ResolucionnominaEliminar")
	router.Path("/ResolucionnominaEditar/{codigo:[0-9]+}").HandlerFunc(ResolucionnominaEditar).Name(
		"ResolucionEditar")
	router.Path("/ResolucionnominaPdf/{codigo:[0-9]+}").HandlerFunc(ResolucionnominaPdf).Name(
		"ResolucionnominaPdf")

	router.Path("/ResolucionnominaTodosPdf").HandlerFunc(ResolucionnominaTodosPdf).
		Name("ResolucionnominaTodosPdf")
	router.Path("/ResolucionnominaExcel").HandlerFunc(ResolucionnominaExcel).
		Name("ResolucionnominaExcel")

	// COTIZACION
	router.Path("/CotizacionLista").HandlerFunc(CotizacionLista).Name("CotizacionLista")
	router.Path("/CotizacionNuevo/{codigo}").HandlerFunc(CotizacionNuevo).Name("CotizacionNuevo")
	router.Path("/BodegaLlenar").HandlerFunc(BodegaLlenar).Name("BodegaLlenar")
	router.Path("/CotizacionExiste/{codigo}").HandlerFunc(CotizacionExiste).
		Name("CotizacionExiste")
	router.Path("/CotizacionEditar/{codigo}").HandlerFunc(CotizacionEditar).
		Name("CotizacionEditar")
	router.Path("/CotizacionAgregar").HandlerFunc(CotizacionAgregar).Name(
		"CotizacionAgregar")
	router.Path("/CotizacionBorrar/{codigo}").HandlerFunc(CotizacionBorrar).
		Name("CotizacionBorrar")
	router.Path("/CotizacionEliminar/{codigo}").HandlerFunc(
		CotizacionEliminar).Name("CotizacionEliminar")
	router.Path("/CotizacionPdf/{codigo}").HandlerFunc(CotizacionPdf).Name(
		"CotizacionPdf")

	router.Path("/CotizacionTodosPdf").HandlerFunc(CotizacionTodosPdf).
		Name("CotizacionTodosPdf")
	router.Path("/CotizacionExcel").HandlerFunc(CotizacionExcel).
		Name("CotizacionExcel")

	// COTIZACION SERVICIO
	router.Path("/CotizacionservicioLista").HandlerFunc(CotizacionservicioLista).Name("CotizacionservicioLista")
	router.Path("/CotizacionservicioNuevo/{codigo}").HandlerFunc(CotizacionservicioNuevo).Name("CotizacionservicioNuevo")
	router.Path("/BodegaLlenar").HandlerFunc(BodegaLlenar).Name("BodegaLlenar")
	router.Path("/CotizacionservicioExiste/{codigo}").HandlerFunc(CotizacionservicioExiste).
		Name("CotizacionservicioExiste")
	router.Path("/CotizacionservicioEditar/{codigo}").HandlerFunc(CotizacionservicioEditar).
		Name("CotizacionservicioEditar")
	router.Path("/CotizacionservicioAgregar").HandlerFunc(CotizacionservicioAgregar).Name(
		"CotizacionservicioAgregar")
	router.Path("/CotizacionservicioBorrar/{codigo}").HandlerFunc(CotizacionservicioBorrar).
		Name("CotizacionservicioBorrar")
	router.Path("/CotizacionservicioEliminar/{codigo}").HandlerFunc(
		CotizacionservicioEliminar).Name("CotizacionservicioEliminar")
	router.Path("/CotizacionservicioPdf/{codigo}").HandlerFunc(CotizacionservicioPdf).Name(
		"CotizacionservicioPdf")

	router.Path("/CotizacionservicioTodosPdf").HandlerFunc(CotizacionservicioTodosPdf).
		Name("CotizacionservicioTodosPdf")
	router.Path("/CotizacionservicioExcel").HandlerFunc(CotizacionservicioExcel).
		Name("CotizacionservicioExcel")

	// PEDIDO SERVICIO
	router.Path("/PedidoservicioLista").HandlerFunc(PedidoservicioLista).Name("PedidoservicioLista")
	router.Path("/PedidoservicioNuevo/{codigo}").HandlerFunc(PedidoservicioNuevo).Name("PedidoservicioNuevo")
	router.Path("/PedidoservicioExiste/{codigo}").HandlerFunc(PedidoservicioExiste).
		Name("PedidoservicioExiste")
	router.Path("/PedidoservicioEditar/{codigo}").HandlerFunc(PedidoservicioEditar).
		Name("PedidoservicioEditar")
	router.Path("/PedidoservicioAgregar").HandlerFunc(PedidoservicioAgregar).Name(
		"PedidoservicioAgregar")
	router.Path("/PedidoservicioBorrar/{codigo}").HandlerFunc(PedidoservicioBorrar).
		Name("PedidoservicioBorrar")
	router.Path("/PedidoservicioEliminar/{codigo}").HandlerFunc(
		PedidoservicioEliminar).Name("PedidoservicioEliminar")
	router.Path("/PedidoservicioPdf/{codigo}").HandlerFunc(PedidoservicioPdf).Name(
		"PedidoservicioPdf")

	router.Path("/PedidoservicioTodosPdf").HandlerFunc(PedidoservicioTodosPdf).
		Name("PedidoservicioTodosPdf")
	router.Path("/PedidoservicioExcel").HandlerFunc(PedidoservicioExcel).
		Name("PedidoservicioExcel")

	// COMPRA SERVICIO
	router.Path("/CompraservicioLista").HandlerFunc(CompraservicioLista).Name("CompraservicioLista")
	router.Path("/CompraservicioNuevo/{codigo}").HandlerFunc(CompraservicioNuevo).Name("CompraservicioNuevo")
	router.Path("/CompraservicioExiste/{codigo}").HandlerFunc(CompraservicioExiste).
		Name("CompraservicioExiste")
	router.Path("/CompraservicioEditar/{codigo}").HandlerFunc(CompraservicioEditar).
		Name("CompraservicioEditar")
	router.Path("/CompraservicioAgregar").HandlerFunc(CompraservicioAgregar).Name(
		"CompraservicioAgregar")
	router.Path("/CompraservicioBorrar/{codigo}").HandlerFunc(CompraservicioBorrar).
		Name("CompraservicioBorrar")
	router.Path("/CompraservicioEliminar/{codigo}").HandlerFunc(
		CompraservicioEliminar).Name("CompraservicioEliminar")
	router.Path("/CompraservicioPdf/{codigo}").HandlerFunc(CompraservicioPdf).Name(
		"CompraservicioPdf")

	router.Path("/CompraservicioTodosPdf").HandlerFunc(CompraservicioTodosPdf).
		Name("CompraservicioTodosPdf")
	router.Path("/CompraservicioExcel").HandlerFunc(CompraservicioExcel).
		Name("CompraservicioExcel")

	router.Path("/DatosPedidocompraServicio/{codigo}/{tercero}").HandlerFunc(DatosPedidocompraServicio).
		Name("DatosPedidocompraServicio")
	router.Path("/DatoscompraServicio/{codigo}/{tercero}").HandlerFunc(DatoscompraServicio).
		Name("DatoscompraServicio")

	// DEVOLUCION COMPRA SERVICIO
	router.Path("/DevolucioncompraservicioLista").HandlerFunc(DevolucioncompraservicioLista).Name("DevolucioncompraservicioLista")
	router.Path("/DevolucioncompraservicioNuevo/{codigo}").HandlerFunc(DevolucioncompraservicioNuevo).Name("DevolucioncompraservicioNuevo")
	router.Path("/DevolucioncompraservicioExiste/{codigo}").HandlerFunc(DevolucioncompraservicioExiste).
		Name("DevolucioncompraservicioExiste")
	router.Path("/DevolucioncompraservicioEditar/{codigo}").HandlerFunc(DevolucioncompraservicioEditar).
		Name("DevolucioncompraservicioEditar")
	router.Path("/DevolucioncompraservicioAgregar").HandlerFunc(DevolucioncompraservicioAgregar).Name(
		"DevolucioncompraservicioAgregar")
	router.Path("/DevolucioncompraservicioBorrar/{codigo}").HandlerFunc(DevolucioncompraservicioBorrar).
		Name("DevolucioncompraservicioBorrar")
	router.Path("/DevolucioncompraservicioEliminar/{codigo}").HandlerFunc(
		DevolucioncompraservicioEliminar).Name("DevolucioncompraservicioEliminar")
	router.Path("/DevolucioncompraservicioPdf/{codigo}").HandlerFunc(DevolucioncompraservicioPdf).Name(
		"DevolucioncompraservicioPdf")

	router.Path("/DevolucioncompraservicioTodosPdf").HandlerFunc(DevolucioncompraservicioTodosPdf).
		Name("DevolucioncompraservicioTodosPdf")
	router.Path("/DevolucioncompraservicioExcel").HandlerFunc(DevolucioncompraservicioExcel).
		Name("DevolucioncompraservicioExcel")
	router.Path("/DatosPedidocompraServicio/{codigo}").HandlerFunc(DatosPedidocompraServicio).
		Name("DatosPedidocompraServicio")
	router.Path("/DatoscompraServicio/{codigo}").HandlerFunc(DatoscompraServicio).
		Name("DatoscompraServicio")

	// VENTA
	router.Path("/VentaLista").HandlerFunc(VentaLista).Name("VentaLista")
	router.Path("/VentaNuevo/{codigo}").HandlerFunc(VentaNuevo).Name("VentaNuevo")
	router.Path("/BodegaLlenar").HandlerFunc(BodegaLlenar).Name("BodegaLlenar")
	router.Path("/VentaExiste/{codigo}").HandlerFunc(VentaExiste).
		Name("VentaExiste")
	router.Path("/VentaEditar/{codigo}").HandlerFunc(VentaEditar).
		Name("VentaEditar")
	router.Path("/VentaAgregar").HandlerFunc(VentaAgregar).Name(
		"VentaAgregar")
	router.Path("/VentaBorrar/{codigo}").HandlerFunc(VentaBorrar).
		Name("VentaBorrar")
	router.Path("/VentaEliminar/{codigo}").HandlerFunc(
		VentaEliminar).Name("VentaEliminar")
	router.Path("/VentaPdf/{codigo}").HandlerFunc(VentaPdf).Name(
		"VentaPdf")

	router.Path("/VentaTodosPdf").HandlerFunc(VentaTodosPdf).
		Name("VentaTodosPdf")
	router.Path("/VentaExcel").HandlerFunc(VentaExcel).
		Name("VentaExcel")

	// VENTA SERVICIO
	router.Path("/VentaservicioLista").HandlerFunc(VentaservicioLista).Name("VentaservicioLista")
	router.Path("/VentaservicioNuevo/{codigo}").HandlerFunc(VentaservicioNuevo).Name("VentaservicioNuevo")
	router.Path("/BodegaLlenar").HandlerFunc(BodegaLlenar).Name("BodegaLlenar")
	router.Path("/VentaservicioExiste/{codigo}").HandlerFunc(VentaservicioExiste).
		Name("VentaservicioExiste")
	router.Path("/VentaservicioEditar/{codigo}").HandlerFunc(VentaservicioEditar).
		Name("VentaservicioEditar")
	router.Path("/VentaservicioAgregar").HandlerFunc(VentaservicioAgregar).Name(
		"VentaservicioAgregar")
	router.Path("/VentaservicioBorrar/{codigo}").HandlerFunc(VentaservicioBorrar).
		Name("VentaservicioBorrar")
	router.Path("/VentaservicioEliminar/{codigo}").HandlerFunc(
		VentaservicioEliminar).Name("VentaservicioEliminar")
	router.Path("/VentaservicioPdf/{codigo}").HandlerFunc(VentaservicioPdf).Name(
		"VentaservicioPdf")

	router.Path("/VentaservicioTodosPdf").HandlerFunc(VentaservicioTodosPdf).
		Name("VentaservicioTodosPdf")
	router.Path("/VentaservicioExcel").HandlerFunc(VentaservicioExcel).
		Name("VentaservicioExcel")

	// COMPRA
	router.Path("/CompraLista").HandlerFunc(CompraLista).Name("CompraLista")
	router.Path("/CompraNuevo/{codigo}").HandlerFunc(CompraNuevo).Name("CompraNuevo")
	router.Path("/BodegaLlenar").HandlerFunc(BodegaLlenar).Name("BodegaLlenar")
	router.Path("/CompraExiste/{codigo}").HandlerFunc(CompraExiste).
		Name("CompraExiste")
	router.Path("/CompraEditar/{codigo}").HandlerFunc(CompraEditar).
		Name("CompraEditar")
	router.Path("/CompraAgregar").HandlerFunc(CompraAgregar).Name(
		"CompraAgregar")
	router.Path("/CompraBorrar/{codigo}").HandlerFunc(CompraBorrar).
		Name("CompraBorrar")
	router.Path("/CompraEliminar/{codigo}").HandlerFunc(
		CompraEliminar).Name("CompraEliminar")
	router.Path("/CompraPdf/{codigo}").HandlerFunc(CompraPdf).Name(
		"CompraPdf")

	router.Path("/CompraTodosPdf").HandlerFunc(CompraTodosPdf).
		Name("CompraTodosPdf")
	router.Path("/CompraExcel").HandlerFunc(CompraExcel).
		Name("CompraExcel")

	// TRAE EL PEDIDO EN LA COMPRA
	router.Path("/DatosPedido/{codigo}/{tercero}").HandlerFunc(DatosPedido).
		Name("DatosPedido")

	// TRAE EL PEDIDO EN LA COMPRA SERVICIO
	router.Path("/DatosPedidoservicio/{codigo}/{tercero}").HandlerFunc(DatosPedido).
		Name("DatosPedidoservicio")

	// TRAE EL PEDIDO SOPORTE EN EL SOPORTE
	router.Path("/DatosPedidosoporte/{codigo}/{tercero}").HandlerFunc(Datospedidosoporte).
		Name("DatosPedidosoporte")

	// TRAE EL PEDIDO SOPORTE SERVICIO EN EL SOPORTE SERVICIO
	router.Path("/DatosPedidosoporteservicio/{codigo}/{tercero}").HandlerFunc(Datospedidosoporteservicio).
		Name("DatosPedidosoporteservicio")

	// TRAE EL PEDIDO FACTURA GASTO LA FACTURA GASTO
	router.Path("/DatosPedidofacturagasto/{codigo}/{tercero}").HandlerFunc(Datospedidofacturagasto).
		Name("DatosPedidofacturagasto")

	// TRAE LA COMPRA EN DEVOLUCION
	router.Path("/DatosCompra/{codigo}/{tercero}").HandlerFunc(DatosCompra).
		Name("DatosCompra")

	// TRAE LA COTIZACION EN LA VENTA
	router.Path("/DatosCotizacion/{codigo}/{tercero}").HandlerFunc(DatosCotizacion).
		Name("DatosCotizacion")

	// TRAE LA COTIZACION SERVICIO EN LA VENTA SERVICIO
	router.Path("/DatosCotizacionservicio/{codigo}/{tercero}").HandlerFunc(DatosCotizacionservicio).
		Name("DatosCotizacionservicio")

	// TRAE LA VENTA EN DEVOLUCION
	router.Path("/DatosVenta/{codigo}/{tercero}").HandlerFunc(DatosVenta).
		Name("DatosVenta")

	// TRAE LA VENTA SERVICIO EN DEVOLUCION
	router.Path("/DatosVentaservicio/{codigo}/{tercero}").HandlerFunc(DatosVentaservicio).
		Name("DatosVentaservicio")

	// TRAE EL SOPORTE EN LA DEVOLUCION
	router.Path("/DatosSoporte/{codigo}/{tercero}").HandlerFunc(DatosSoporte).
		Name("DatosSoporte")

	// TRAE EL SOPORTE SERVICIO  EN LA DEVOLUCION SOPORTE SERVICIO
	router.Path("/Datossoporteservicio/{codigo}/{{tercero}").HandlerFunc(Datossoporteservicio).
		Name("Datossoporteservicio")

	// TRAE LA FACTURA GASTO  EN LA DEVOLUCION FACTURA GASTO
	router.Path("/Datosfacturagasto/{codigo}/{tercero}").HandlerFunc(Datosfacturagasto).
		Name("Datosfacturagasto")

	// PEDIDO
	router.Path("/PedidoLista").HandlerFunc(PedidoLista).Name("PedidoLista")
	router.Path("/PedidoNuevo/{codigo}").HandlerFunc(PedidoNuevo).Name("PedidoNuevo")
	router.Path("/BodegaLlenar").HandlerFunc(BodegaLlenar).Name("BodegaLlenar")
	router.Path("/PedidoExiste/{codigo}").HandlerFunc(PedidoExiste).
		Name("PedidoExiste")
	router.Path("/PedidoEditar/{codigo}").HandlerFunc(PedidoEditar).
		Name("PedidoEditar")
	router.Path("/PedidoAgregar").HandlerFunc(PedidoAgregar).Name(
		"PedidoAgregar")
	router.Path("/PedidoBorrar/{codigo}").HandlerFunc(PedidoBorrar).
		Name("PedidoBorrar")
	router.Path("/PedidoEliminar/{codigo}").HandlerFunc(
		PedidoEliminar).Name("PedidoEliminar")
	router.Path("/PedidoPdf/{codigo}").HandlerFunc(PedidoPdf).Name(
		"PedidoPdf")

	router.Path("/PedidoTodosPdf").HandlerFunc(PedidoTodosPdf).
		Name("PedidoTodosPdf")
	router.Path("/PedidoExcel").HandlerFunc(PedidoExcel).
		Name("PedidoExcel")

	// PEDIDO SOPORTE
	router.Path("/PedidosoporteLista").HandlerFunc(PedidosoporteLista).Name("PedidosoporteLista")
	router.Path("/PedidosoporteNuevo/{codigo}").HandlerFunc(PedidosoporteNuevo).Name("PedidosoporteNuevo")
	router.Path("/BodegaLlenar").HandlerFunc(BodegaLlenar).Name("BodegaLlenar")
	router.Path("/PedidosoporteExiste/{codigo}").HandlerFunc(PedidosoporteExiste).
		Name("PedidosoporteExiste")
	router.Path("/PedidosoporteEditar/{codigo}").HandlerFunc(PedidosoporteEditar).
		Name("PedidosoporteEditar")
	router.Path("/PedidosoporteAgregar").HandlerFunc(PedidosoporteAgregar).Name(
		"PedidosoporteAgregar")
	router.Path("/PedidosoporteBorrar/{codigo}").HandlerFunc(PedidosoporteBorrar).
		Name("PedidosoporteBorrar")
	router.Path("/PedidosoporteEliminar/{codigo}").HandlerFunc(
		PedidosoporteEliminar).Name("PedidosoporteEliminar")
	router.Path("/PedidosoportePdf/{codigo}").HandlerFunc(PedidosoportePdf).Name(
		"PedidosoportePdf")

	router.Path("/PedidosoporteTodosPdf").HandlerFunc(PedidosoporteTodosPdf).
		Name("PedidosoporteTodosPdf")
	router.Path("/PedidosoporteExcel").HandlerFunc(PedidosoporteExcel).
		Name("PedidosoporteExcel")

	// PEDIDO SOPORTE SERVICIO
	router.Path("/PedidosoporteservicioLista").HandlerFunc(PedidosoporteservicioLista).Name("PedidosoporteservicioLista")
	router.Path("/PedidosoporteservicioNuevo/{codigo}").HandlerFunc(PedidosoporteservicioNuevo).Name("PedidosoporteservicioNuevo")
	router.Path("/BodegaLlenar").HandlerFunc(BodegaLlenar).Name("BodegaLlenar")
	router.Path("/PedidosoporteservicioExiste/{codigo}").HandlerFunc(PedidosoporteservicioExiste).
		Name("PedidosoporteservicioExiste")
	router.Path("/PedidosoporteservicioEditar/{codigo}").HandlerFunc(PedidosoporteservicioEditar).
		Name("PedidosoporteservicioEditar")
	router.Path("/PedidosoporteservicioAgregar").HandlerFunc(PedidosoporteservicioAgregar).Name(
		"PedidosoporteservicioAgregar")
	router.Path("/PedidosoporteservicioBorrar/{codigo}").HandlerFunc(PedidosoporteservicioBorrar).
		Name("PedidosoporteservicioBorrar")
	router.Path("/PedidosoporteservicioEliminar/{codigo}").HandlerFunc(
		PedidosoporteservicioEliminar).Name("PedidosoporteservicioEliminar")
	router.Path("/PedidosoporteservicioPdf/{codigo}").HandlerFunc(PedidosoporteservicioPdf).Name(
		"PedidosoporteservicioPdf")

	router.Path("/PedidosoporteservicioTodosPdf").HandlerFunc(PedidosoporteservicioTodosPdf).
		Name("PedidosoporteservicioTodosPdf")
	router.Path("/PedidosoporteservicioExcel").HandlerFunc(PedidosoporteservicioExcel).
		Name("PedidosoporteservicioExcel")

	// DEVOLUCION COMPRA
	router.Path("/DevolucioncompraLista").HandlerFunc(DevolucioncompraLista).Name("DevolucioncompraLista")
	router.Path("/DevolucioncompraNuevo/{codigo}").HandlerFunc(DevolucioncompraNuevo).Name("DevolucioncompraNuevo")
	router.Path("/BodegaLlenar").HandlerFunc(BodegaLlenar).Name("BodegaLlenar")
	router.Path("/DevolucioncompraExiste/{codigo}").HandlerFunc(DevolucioncompraExiste).
		Name("DevolucioncompraExiste")
	router.Path("/DevolucioncompraEditar/{codigo}").HandlerFunc(DevolucioncompraEditar).
		Name("DevolucioncompraEditar")
	router.Path("/DevolucioncompraAgregar").HandlerFunc(DevolucioncompraAgregar).Name(
		"DevolucioncompraAgregar")
	router.Path("/DevolucioncompraBorrar/{codigo}").HandlerFunc(DevolucioncompraBorrar).
		Name("DevolucioncompraBorrar")
	router.Path("/DevolucioncompraEliminar/{codigo}").HandlerFunc(
		DevolucioncompraEliminar).Name("DevolucioncompraEliminar")
	router.Path("/DevolucioncompraPdf/{codigo}").HandlerFunc(DevolucioncompraPdf).Name(
		"DevolucioncompraPdf")

	router.Path("/DevolucioncompraTodosPdf").HandlerFunc(DevolucioncompraTodosPdf).
		Name("DevolucioncompraTodosPdf")
	router.Path("/DevolucioncompraExcel").HandlerFunc(DevolucioncompraExcel).
		Name("DevolucioncompraExcel")

	// DEVOLUCION VENTA
	router.Path("/DevolucionventaLista").HandlerFunc(DevolucionventaLista).Name("DevolucionventaLista")
	router.Path("/DevolucionventaNuevo/{codigo}").HandlerFunc(DevolucionventaNuevo).Name("DevolucionventaNuevo")
	router.Path("/BodegaLlenar").HandlerFunc(BodegaLlenar).Name("BodegaLlenar")
	router.Path("/DevolucionventaExiste/{codigo}").HandlerFunc(DevolucionventaExiste).
		Name("DevolucionventaExiste")
	router.Path("/DevolucionventaEditar/{codigo}").HandlerFunc(DevolucionventaEditar).
		Name("DevolucionventaEditar")
	router.Path("/DevolucionventaAgregar").HandlerFunc(DevolucionventaAgregar).Name(
		"DevolucionventaAgregar")
	router.Path("/DevolucionventaBorrar/{codigo}").HandlerFunc(DevolucionventaBorrar).
		Name("DevolucionventaBorrar")
	router.Path("/DevolucionventaEliminar/{codigo}").HandlerFunc(
		DevolucionventaEliminar).Name("DevolucionventaEliminar")
	router.Path("/DevolucionventaPdf/{codigo}").HandlerFunc(DevolucionventaPdf).Name(
		"DevolucionventaPdf")

	router.Path("/DevolucionventaTodosPdf").HandlerFunc(DevolucionventaTodosPdf).
		Name("DevolucionventaTodosPdf")
	router.Path("/DevolucionventaExcel").HandlerFunc(DevolucionventaExcel).
		Name("DevolucionventaExcel")

	// DEVOLUCION VENTA SERVICIO
	router.Path("/DevolucionventaservicioLista").HandlerFunc(DevolucionventaservicioLista).Name("DevolucionventaservicioLista")
	router.Path("/DevolucionventaservicioNuevo/{codigo}").HandlerFunc(DevolucionventaservicioNuevo).Name("DevolucionventaservicioNuevo")
	router.Path("/BodegaLlenar").HandlerFunc(BodegaLlenar).Name("BodegaLlenar")
	router.Path("/DevolucionventaservicioExiste/{codigo}").HandlerFunc(DevolucionventaservicioExiste).
		Name("DevolucionventaservicioExiste")
	router.Path("/DevolucionventaservicioEditar/{codigo}").HandlerFunc(DevolucionventaservicioEditar).
		Name("DevolucionventaservicioEditar")
	router.Path("/DevolucionventaservicioAgregar").HandlerFunc(DevolucionventaservicioAgregar).Name(
		"DevolucionventaservicioAgregar")
	router.Path("/DevolucionventaservicioBorrar/{codigo}").HandlerFunc(DevolucionventaservicioBorrar).
		Name("DevolucionventaservicioBorrar")
	router.Path("/DevolucionventaservicioEliminar/{codigo}").HandlerFunc(
		DevolucionventaservicioEliminar).Name("DevolucionventaservicioEliminar")
	router.Path("/DevolucionventaservicioPdf/{codigo}").HandlerFunc(DevolucionventaservicioPdf).Name(
		"DevolucionventaservicioPdf")

	router.Path("/DevolucionventaservicioTodosPdf").HandlerFunc(
		DevolucionventaservicioTodosPdf).
		Name("DevolucionventaservicioTodosPdf")
	router.Path("/DevolucionventaservicioExcel").HandlerFunc(
		DevolucionventaservicioExcel).
		Name("DevolucionventaservicioExcel")

	// INVENTARIO INICIAL
	router.Path("/InventarioinicialLista").HandlerFunc(InventarioinicialLista).Name("InventarioinicialLista")
	router.Path("/InventarioinicialNuevo/{codigo}").HandlerFunc(InventarioinicialNuevo).Name("InventarioinicialNuevo")
	router.Path("/BodegaLlenar").HandlerFunc(BodegaLlenar).Name("BodegaLlenar")
	router.Path("/InventarioinicialExiste/{codigo}").HandlerFunc(InventarioinicialExiste).
		Name("InventarioinicialExiste")
	router.Path("/InventarioinicialEditar/{codigo}").HandlerFunc(InventarioinicialEditar).
		Name("InventarioinicialEditar")
	router.Path("/InventarioinicialAgregar").HandlerFunc(InventarioinicialAgregar).Name(
		"InventarioinicialAgregar")
	router.Path("/InventarioinicialBorrar/{codigo}").HandlerFunc(InventarioinicialBorrar).
		Name("InventarioinicialBorrar")
	router.Path("/InventarioinicialEliminar/{codigo}").HandlerFunc(
		InventarioinicialEliminar).Name("InventarioinicialEliminar")
	router.Path("/InventarioinicialPdf/{codigo}").HandlerFunc(InventarioinicialPdf).Name(
		"InventarioinicialPdf")

	router.Path("/InventarioinicialTodosPdf").HandlerFunc(InventarioinicialTodosPdf).
		Name("InventarioinicialTodosPdf")
	router.Path("/InventarioinicialExcel").HandlerFunc(InventarioinicialExcel).
		Name("InventarioinicialExcel")

	// ARCHIVO CONFIGURACION INVENTARIO
	router.Path("/ConfiguracioninventarioNuevo/{panel}").HandlerFunc(ConfiguracioninventarioNuevo).Name("ConfiguracioninventarioNuevo")
	router.Path("/ConfiguracioninventarioInsertar").HandlerFunc(ConfiguracioninventarioInsertar).Name(
		"ConfiguracioninventarioInsertar")
	router.Path("/ConfiguracioninventarioPdf/{codigo:[0-9]+}").HandlerFunc(ConfiguracioninventarioPdf).Name(
		"ConfiguracioninventarioPdf")

	// ARCHIVO CONFIGURACION CONTABILIDAD
	router.Path("/ConfiguracioncontabilidadNuevo/{panel}").HandlerFunc(ConfiguracioncontabilidadNuevo).Name("ConfiguracioninventarioNuevo")
	router.Path("/ConfiguracioncontabilidadInsertar").HandlerFunc(ConfiguracioncontabilidadInsertar).Name(
		"ConfiguracioncontabilidadInsertar")
	router.Path("/ConfiguracioncontabilidadPdf/{codigo:[0-9]+}").HandlerFunc(ConfiguracioncontabilidadPdf).Name(
		"ConfiguracioncontabilidadPdf")

	// ARCHIVO CONFIGURACION NOMINA
	router.Path("/ConfiguracionnominaNuevo/{panel}").HandlerFunc(ConfiguracionnominaNuevo).Name("ConfiguracionnominaNuevo")
	router.Path("/ConfiguracionnominaInsertar").HandlerFunc(ConfiguracionnominaInsertar).Name(
		"ConfiguracionnominaInsertar")
	router.Path("/ConfiguracionnominaPdf/{codigo:[0-9]+}").HandlerFunc(ConfiguracionnominaPdf).Name(
		"ConfiguracionnominaPdf")

	// ARCHIVO PLAN DE CUENTAS NIIF
	router.Path("/PlandecuentaniifLista").HandlerFunc(PlandecuentaniifLista).Name("PlandecuentaniifLista")
	router.Path("/PlandecuentaniifListaBuscar/{panel}/{codigo}/{elemento}").HandlerFunc(PlandecuentaniifListaBuscar).Name("PlandecuentaniifLista")

	router.Path("/PlandecuentaniifPdf/{codigo:[0-9]+}").HandlerFunc(PlandecuentaniifPdf).Name(
		"PlandecuentaniifPdf")

	// ARCHIVO PLAN DE CUENTAS PUC
	router.Path("/PlandecuentapucLista").HandlerFunc(PlandecuentapucLista).Name("PlandecuentapucLista")
	router.Path("/PlandecuentapucPdf/{codigo:[0-9]+}").HandlerFunc(PlandecuentapucPdf).Name(
		"PlandecuentapucPdf")

	// ARCHIVO PLAN DE CUENTAS EMPRESA
	router.Path("/PlandecuentaempresaBuscar/{codigo}").HandlerFunc(PlandecuentaempresaBuscar).
		Name("PlandecuentaempresaBuscar")
	router.Path("/PlandecuentaempresaBuscarAuxiliar/{codigo}").HandlerFunc(PlandecuentaempresaBuscarAuxiliar).
		Name("PlandecuentaempresaBuscarAuxiliar")
	router.Path("/PlandecuentaempresaNuevo/{panel}/{codigo}/{elemento}").HandlerFunc(
		PlandecuentaempresaNuevo).Name("PlandecuentaempresaNuevo")
	router.Path("/PlandecuentaempresaNuevoCopia/{copiacodigo}").HandlerFunc(
		PlandecuentaempresaNuevoCopia).Name("PlandecuentaempresaNuevoCopia")
	router.Path("/PlandecuentaempresaInsertar").HandlerFunc(PlandecuentaempresaInsertar).Name(
		"PlandecuentaempresaInsertar")
	router.Path("/PlandecuentaempresaAgregar").HandlerFunc(PlandecuentaempresaAgregar).Name(
		"PlandecuentaempresaAgregar")
	router.Path("/PlandecuentaempresaActual/{codigo}").HandlerFunc(PlandecuentaempresaActual).
		Name("PlandecuentaempresaActual")
	router.Path("/PlandecuentaempresaLista/{panel}/{codigo}/{elemento}").HandlerFunc(PlandecuentaempresaLista).Name("PlandecuentaempresaLista")
	router.Path("/PlandecuentaempresaExiste/{codigo:[0-9]+}").HandlerFunc(PlandecuentaempresaExiste).
		Name("PlandecuentaempresaExiste")
	router.Path("/PlandecuentaempresaEditar/{codigo:[0-9]+}").HandlerFunc(PlandecuentaempresaEditar).Name(
		"PlandecuentaempresaEditar")
	router.Path("/PlandecuentaempresaActualizar/{codigo}").HandlerFunc(PlandecuentaempresaActualizar).Name(
		"PlandecuentaempresaActualizar")
	router.Path("/PlandecuentaempresaBorrar/{codigo}").HandlerFunc(PlandecuentaempresaBorrar).Name(
		"PlandecuentaempresaBorrar")
	router.Path("/PlandecuentaempresaEliminar/{codigo}").HandlerFunc(PlandecuentaempresaEliminar).Name("PlandecuentaempresaEliminar")
	router.Path("/PlandecuentaempresaPdf/{codigo:[0-9]+}").HandlerFunc(PlandecuentaempresaPdf).Name(
		"PlandecuentaempresaPdf")
	router.Path("/PlandecuentaempresaTodosPdf").HandlerFunc(PlandecuentaempresaTodosPdf).
		Name("PlandecuentaempresaTodosPdf")
	router.Path("/PlandecuentaempresaExcel").HandlerFunc(PlandecuentaempresaExcel).
		Name("PlandecuentaempresaExcel")

	// ARCHIVO PLAN DE CUENTAS CONJUNTOS
	router.Path("/PlandecuentaempresaphNuevo/{panel}/{codigo}/{elemento}").HandlerFunc(
		PlandecuentaempresaNuevo).Name("PlandecuentaempresaphNuevo")
	router.Path("/PlandecuentaempresaphNuevoCopia/{copiacodigo}").HandlerFunc(
		PlandecuentaempresaNuevoCopia).Name("PlandecuentaempresaphNuevoCopia")
	router.Path("/PlandecuentaempresaphInsertar").HandlerFunc(PlandecuentaempresaInsertar).Name(
		"PlandecuentaempresaphInsertar")
	router.Path("/PlandecuentaempresaphAgregar").HandlerFunc(PlandecuentaempresaAgregar).Name(
		"PlandecuentaempresaphAgregar")
	router.Path("/PlandecuentaempresaphEditar/{codigo:[0-9]+}").HandlerFunc(PlandecuentaempresaEditar).Name(
		"PlandecuentaempresaphEditar")
	router.Path("/PlandecuentaempresaphActualizar/{codigo}").HandlerFunc(PlandecuentaempresaActualizar).Name(
		"PlandecuentaempresaphActualizar")
	router.Path("/PlandecuentaempresaphBorrar/{codigo}").HandlerFunc(PlandecuentaempresaBorrar).Name(
		"PlandecuentaempresaphBorrar")
	router.Path("/PlandecuentaempresaphEliminar/{codigo}").HandlerFunc(PlandecuentaempresaEliminar).Name("PlandecuentaempresaphEliminar")
	router.Path("/PlandecuentaempresaphPdf/{codigo:[0-9]+}").HandlerFunc(PlandecuentaempresaPdf).Name(
		"PlandecuentaempresaphPdf")
	router.Path("/PlandecuentaempresaphTodosPdf").HandlerFunc(PlandecuentaempresaTodosPdf).
		Name("PlandecuentaempresaphTodosPdf")
	router.Path("/PlandecuentaempresaphExcel").HandlerFunc(PlandecuentaempresaExcel).
		Name("PlandecuentaempresaphExcel")

	// ARCHIVO RETENCION EN LA FUENTE
	router.Path("/RetencionenlafuenteLista").HandlerFunc(RetencionenlafuenteLista).Name("RetencionenlafuenteLista")
	router.Path("/RetencionenlafuentePdf/{codigo:[0-9]+}").HandlerFunc(RetencionenlafuentePdf).Name(
		"RetencionenlafuentePdf")

	// ARCHIVO DEPRECIACION
	router.Path("/DepreciacionLista").HandlerFunc(DepreciacionLista).Name("DepreciacionLista")
	router.Path("/DepreciacionPdf/{codigo:[0-9]+}").HandlerFunc(DepreciacionPdf).Name(
		"DepreciacionPdf")

	// COMPROBANTE
	router.Path("/ComprobanteLista").HandlerFunc(ComprobanteLista).Name("ComprobanteLista")
	router.Path("/ComprobanteNuevo/{documento}/{numero}").HandlerFunc(ComprobanteNuevo).Name("ComprobanteNuevo")
	router.Path("/ComprobanteExiste/{documento}/{numero}").HandlerFunc(ComprobanteExiste).
		Name("ComprobanteExiste")
	router.Path("/ComprobanteAgregar").HandlerFunc(ComprobanteAgregar).Name(
		"ComprobanteAgregar")
	router.Path("/ComprobanteEditar/{documento}/{numero}").HandlerFunc(ComprobanteEditar).
		Name("ComprobanteEditar")
	router.Path("/ComprobanteBorrar/{documento}/{numero}").HandlerFunc(ComprobanteBorrar).
		Name("ComprobanteBorrar")
	router.Path("/ComprobanteEliminar/{documento}/{numero}").HandlerFunc(
		ComprobanteEliminar).Name("ComprobanteEliminar")
	router.Path("/ComprobantePdf/{documento}/{numero}").HandlerFunc(ComprobantePdf).Name(
		"ComprobantePdf")

	router.Path("/ComprobanteTodosPdf").HandlerFunc(ComprobanteTodosPdf).
		Name("ComprobanteTodosPdf")
	router.Path("/ComprobanteExcel").HandlerFunc(ComprobanteExcel).
		Name("ComprobanteExcel")

	// CUENTADECOBRO
	router.Path("/CuentadecobroGenerar/{panel}").HandlerFunc(CuentadecobroGenerar).Name("CuentadecobroGenerar")
	router.Path("/CuentadecobroGenerarMes/{mes}/{ano}/{centro}/{porcentaje}").HandlerFunc(CuentadecobroGenerarMes).Name("CuentadecobroGenerarMes")

	router.Path("/CuentadecobroLista/{panel}").HandlerFunc(CuentadecobroLista).Name("CuentadecobroLista")
	router.Path("/CuentadecobroNuevo").HandlerFunc(CuentadecobroNuevo).Name("CuentadecobroNuevo")
	router.Path("/CuentadecobroExiste/{numero}").HandlerFunc(CuentadecobroExiste).
		Name("CuentadecobroExiste")
	router.Path("/CuentadecobroAgregar").HandlerFunc(CuentadecobroAgregar).Name(
		"CuentadecobroAgregar")
	router.Path("/CuentadecobroEditar/{panel}/{numero}").HandlerFunc(CuentadecobroEditar).
		Name("CuentadecobroEditar")
	router.Path("/CuentadecobroBorrar/{panel}/{numero}").HandlerFunc(CuentadecobroBorrar).
		Name("CuentadecobroBorrar")
	router.Path("/CuentadecobroEliminar/{panel}/{numero}").HandlerFunc(
		CuentadecobroEliminar).Name("CuentadecobroEliminar")
	router.Path("/CuentadecobroPdf/{numero}").HandlerFunc(CuentadecobroPdf).Name(
		"CuentadecobroPdf")

	router.Path("/CuentadecobroTodosPdf").HandlerFunc(CuentadecobroTodosPdf).
		Name("CuentadecobroTodosPdf")
	router.Path("/CuentadecobroExcel").HandlerFunc(CuentadecobroExcel).
		Name("CuentadecobroExcel")

	// CUENTA DE COBRO DATO
	router.Path("/CuentadecobroDato/{panel}").HandlerFunc(CuentadecobroDato).Name("CuentadecobroGenerar")
	router.Path("/CuentadecobroDatoAgregar").HandlerFunc(CuentadecobroDatoAgregar).Name("CuentadecobroGenerarAgregar")

	// ARCHIVO INVENTARIO
	router.Path("/InventarioLista").HandlerFunc(InventarioLista).Name("InventarioLista")

	// CONCILIACION

	router.Path("/ConciliacionLista").HandlerFunc(ConciliacionLista).Name("ConciliacionLista")
	router.Path("/ConciliacionDatos/{cuenta}/{mes}").HandlerFunc(ConciliacionDato).Name("ConciliacionDato")
	router.Path("/ConciliacionInsertar").HandlerFunc(ConciliacionInsertar).Name("ConciliacionInsertar")
	router.Path("/ConciliacionTodosPdf/{cuenta}/{mes}").HandlerFunc(ConciliacionTodosPdf).Name("ConciliacionCuentaTodosPdf")

	// BANCO
	router.Path("/BancoLista").HandlerFunc(BancoLista).Name("BancoLista")
	router.Path("/BancoDatos/{tercero}/{documento}/{fecha}").HandlerFunc(BancoDato).Name("BancoDato")
	router.Path("/BancoDatoAgregar").HandlerFunc(BancoDatoAgregar).Name("BancoDatoAgregar")

	// SALDO
	router.Path("/SaldoLista").HandlerFunc(SaldoLista).Name("SaldoLista")
	router.Path("/SaldoDatosTodos").HandlerFunc(SaldoDatosTodos).Name("SaldoDatosTodos")
	router.Path("/SaldoDatosPdf/{codigo}/{fechainicial}/{fechafinal}/{bodega}/{tipo}/{discriminar}").HandlerFunc(SaldoDatosPdf).Name("SaldoDatosTodos")
	router.Path("/SaldoDatosExcel/{codigo}/{fechainicial}/{fechafinal}/{bodega}/{tipo}/{discriminar}").HandlerFunc(SaldoDatosExcel).Name("SaldoDatosTodos")

	// saldo bodega json
	router.Path("/SaldoDatosProducto/{codigo}").HandlerFunc(SaldoDatosProducto).Name("SaldoDatosTodos")
	router.Path("/SaldoDatosProductoBodega/{codigo}/{bodega}").HandlerFunc(SaldoDatosProductoBodega).Name("SaldoDatosTodos")

	// SALDO BODEGA

	router.Path("/SaldoBodegaLista").HandlerFunc(SaldoBodegaLista).Name("SaldoLista")
	router.Path("/SaldoBodegaDatosTodos").HandlerFunc(SaldoBodegaDatosTodos).Name("SaldoDatosTodos")
	router.Path("/SaldoBodegaDatosPdf/{codigo}/{fechainicial}/{fechafinal}/{bodega}/{tipo}/{discriminar}").HandlerFunc(SaldoBodegaDatosPdf).Name("SaldoDatosTodos")
	router.Path("/SaldoBodegaDatosExcel/{codigo}/{fechainicial}/{fechafinal}/{bodega}/{tipo}/{discriminar}").HandlerFunc(SaldoBodegaDatosExcel).Name("SaldoDatosTodos")

	// ARCHIVO KARDEX
	router.Path("/KardexLista").HandlerFunc(KardexLista).Name("KardexLista")
	router.Path("/KardexDatosTodos/{codigo}/{fechainicial}/{fechafinal}/{bodega}/{tipo}/{discriminar}").HandlerFunc(KardexDatosTodos).Name("KardexDatosTodos")
	router.Path("/KardexDatosPdf/{codigo}/{fechainicial}/{fechafinal}/{bodega}/{tipo}/{discriminar}").HandlerFunc(KardexDatosPdf).Name("KardexDatosTodos")
	router.Path("/KardexDatosExcel/{codigo}/{fechainicial}/{fechafinal}/{bodega}/{tipo}/{discriminar}").HandlerFunc(KardexDatosExcel).Name("KardexDatosTodos")

	// KARDEX COSTO
	router.Path("/KardexCostoLista").HandlerFunc(KardexCostoLista).Name("KardexCostoLista")
	router.Path("/KardexCostoDatos/{mes}/{centro}").HandlerFunc(KardexCostoDatos).Name("KardexCostoGenerar")

	// BALANCE DE PRUEBA
	router.Path("/BalancedepruebaLista").HandlerFunc(BalancedepruebaLista).Name("BalancedepruebaLista")
	router.Path("/BalancedepruebaDatos").HandlerFunc(BalancedepruebaDatos).Name("BalancedepruebaDatos")
	//router.Path("/BalancedepruebaPdf").HandlerFunc(BalancedepruebaPdf).Name("BalancedepruebaPdf")
	router.Path("/BalancedepruebaPdf/{FechaInicial}/{FechaFinal}/{CuentaInicial}/{CuentaFinal}/{TerceroInicial}/{TerceroFinal}/{CentroInicial}/{CentroFinal}/{DocumentoInicial}/{DocumentoFinal}/{NumeroInicial}/{NumeroFinal}/{Detalle}/{Nivel}/{Activa}/{Subtotal}").HandlerFunc(BalancedepruebaPdf).Name("BalancedepruebaPdf")
	router.Path("/BalancedepruebaExcel/{FechaInicial}/{FechaFinal}/{CuentaInicial}/{CuentaFinal}/{TerceroInicial}/{TerceroFinal}/{CentroInicial}/{CentroFinal}/{DocumentoInicial}/{DocumentoFinal}/{NumeroInicial}/{NumeroFinal}/{Detalle}/{Nivel}/{Activa}/{Subtotal}").HandlerFunc(BalancedepruebaExcel).Name("BalancedepruebaExcel")

	//{FechaInicial}/{FechaFinal}/{CuentaInicial}/{CuentaFinal}/{TerceroInicial}/{TerceroFinal}/{CentroInicial}/{CentroFinal}/{DocumentoInicial}/{DocumentoFinal}/{NumeroInicial}/{NumeroFinal}/{Detalle}/{Nivel}/{Activa}/{Subtotal}

	// ESTADO DE RESULTADO FINANCIERO
	router.Path("/SituacionFinancieraLista").HandlerFunc(SituacionFinancieraLista).Name("SituacionFinancieraLista")
	router.Path("/SituacionFinancieraDatos").HandlerFunc(SituacionFinancieraDatos).Name("BalancedepruebaDatos")

	// estado resultados

	router.Path("/EstadoResultadoLista").HandlerFunc(EstadoResultadoLista).Name("SituacionFinancieraLista")
	router.Path("/EstadoResultadoDatos").HandlerFunc(EstadoResultadoDatos).Name("BalancedepruebaDatos")
	router.Path("/EstadoResultadoPdf/{FechaInicial}/{FechaFinal}/{Nivel}").HandlerFunc(EstadoResultadoPdf).Name("EstadoResultadoPdf")
	router.Path("/EstadoResultadoExcel/{FechaInicial}/{FechaFinal}/{Nivel}").HandlerFunc(EstadoResultadoExcel).Name("EstadoResultadoExcel")

	// ARCHIVO CONCEPTO
	router.Path("/ConceptoNuevo/{codigo}").HandlerFunc(ConceptoNuevo).Name("ConceptoNuevo")
	router.Path("/ConceptoBuscar/{codigo}").HandlerFunc(ConceptoBuscar).
		Name("ConceptoBuscar")
	router.Path("/ConceptoActual/{codigo}").HandlerFunc(ConceptoActual).
		Name("ConceptoActual")
	router.Path("/ConceptoLista").HandlerFunc(ConceptoLista).Name("ConceptoLista")
	router.Path("/ConceptoExiste/{codigo:[0-9]+}").HandlerFunc(ConceptoExiste).
		Name("ConceptoExiste")
	router.Path("/ConceptoInsertar").HandlerFunc(ConceptoInsertar).Name(
		"ConceptoInsertar")
	router.Path("/ConceptoActualizar/{codigo:[0-9]+}").HandlerFunc(ConceptoActualizar).Name(
		"ConceptoActualizar")
	router.Path("/ConceptoBorrar/{codigo:[0-9]+}").HandlerFunc(ConceptoBorrar).Name(
		"ConceptoBorrar")
	router.Path("/ConceptoEliminar/{codigo:[0-9]+}").HandlerFunc(
		ConceptoEliminar).Name("ConceptoEliminar")
	router.Path("/ConceptoEditar/{codigo:[0-9]+}").HandlerFunc(ConceptoEditar).Name(
		"ConceptoEditar")
	router.Path("/ConceptoPdf/{codigo:[0-9]+}").HandlerFunc(ConceptoPdf).Name(
		"ConceptoPdf")

	router.Path("/ConceptoTodosPdf").HandlerFunc(ConceptoTodosPdf).
		Name("ConceptoTodosPdf")
	router.Path("/ConceptoExcel").HandlerFunc(ConceptoExcel).
		Name("ConceptoExcel")

	// SOPORTE
	router.Path("/SoporteLista").HandlerFunc(SoporteLista).Name("SoporteLista")
	router.Path("/SoporteNuevo/{codigo}").HandlerFunc(SoporteNuevo).Name("SoporteNuevo")
	router.Path("/BodegaLlenar").HandlerFunc(BodegaLlenar).Name("BodegaLlenar")
	router.Path("/SoporteExiste/{codigo}").HandlerFunc(SoporteExiste).
		Name("SoporteExiste")
	router.Path("/SoporteEditar/{codigo}").HandlerFunc(SoporteEditar).
		Name("SoporteEditar")
	router.Path("/SoporteAgregar").HandlerFunc(SoporteAgregar).Name(
		"SoporteAgregar")
	router.Path("/SoporteBorrar/{codigo}").HandlerFunc(SoporteBorrar).
		Name("SoporteBorrar")
	router.Path("/SoporteEliminar/{codigo}").HandlerFunc(
		SoporteEliminar).Name("SoporteEliminar")
	router.Path("/SoportePdf/{codigo}").HandlerFunc(SoportePdf).Name(
		"SoportePdf")

	router.Path("/SoporteTodosPdf").HandlerFunc(SoporteTodosPdf).
		Name("SoporteTodosPdf")
	router.Path("/SoporteExcel").HandlerFunc(SoporteExcel).
		Name("SoporteExcel")

	// SOPORTE SERVICIO
	router.Path("/SoporteservicioLista").HandlerFunc(SoporteservicioLista).Name("SoporteservicioLista")
	router.Path("/SoporteservicioNuevo/{codigo}").HandlerFunc(SoporteservicioNuevo).Name("SoporteservicioNuevo")
	router.Path("/BodegaLlenar").HandlerFunc(BodegaLlenar).Name("BodegaLlenar")
	router.Path("/SoporteservicioExiste/{codigo}").HandlerFunc(SoporteservicioExiste).
		Name("SoporteservicioExiste")
	router.Path("/SoporteservicioEditar/{codigo}").HandlerFunc(SoporteservicioEditar).
		Name("SoporteservicioEditar")
	router.Path("/SoporteservicioAgregar").HandlerFunc(SoporteservicioAgregar).Name(
		"SoporteservicioAgregar")
	router.Path("/SoporteservicioBorrar/{codigo}").HandlerFunc(SoporteservicioBorrar).
		Name("SoporteservicioBorrar")
	router.Path("/SoporteservicioEliminar/{codigo}").HandlerFunc(
		SoporteservicioEliminar).Name("SoporteservicioEliminar")
	router.Path("/SoporteservicioPdf/{codigo}").HandlerFunc(SoporteservicioPdf).Name(
		"SoporteservicioPdf")

	router.Path("/SoporteservicioTodosPdf").HandlerFunc(SoporteservicioTodosPdf).
		Name("SoporteservicioTodosPdf")
	router.Path("/SoporteservicioExcel").HandlerFunc(SoporteservicioExcel).
		Name("SoporteservicioExcel")

	// DEVOLUCION SOPORTE
	router.Path("/DevolucionsoporteLista").HandlerFunc(DevolucionsoporteLista).Name("DevolucionsoporteLista")
	router.Path("/DevolucionsoporteNuevo/{codigo}").HandlerFunc(DevolucionsoporteNuevo).Name("DevolucionsoporteNuevo")
	router.Path("/BodegaLlenar").HandlerFunc(BodegaLlenar).Name("BodegaLlenar")
	router.Path("/DevolucionsoporteExiste/{codigo}").HandlerFunc(DevolucionsoporteExiste).
		Name("DevolucionsoporteExiste")
	router.Path("/DevolucionsoporteEditar/{codigo}").HandlerFunc(DevolucionsoporteEditar).
		Name("DevolucionsoporteEditar")
	router.Path("/DevolucionsoporteAgregar").HandlerFunc(DevolucionsoporteAgregar).Name(
		"DevolucionsoporteAgregar")
	router.Path("/DevolucionsoporteBorrar/{codigo}").HandlerFunc(DevolucionsoporteBorrar).
		Name("DevolucionsoporteBorrar")
	router.Path("/DevolucionsoporteEliminar/{codigo}").HandlerFunc(
		DevolucionsoporteEliminar).Name("DevolucionsoporteEliminar")
	router.Path("/DevolucionsoportePdf/{codigo}").HandlerFunc(DevolucionsoportePdf).Name(
		"DevolucionsoportePdf")

	router.Path("/DevolucionsoporteTodosPdf").HandlerFunc(DevolucionsoporteTodosPdf).
		Name("DesolucionsoporteTodosPdf")
	router.Path("/DevolucionesSoporteExcel").HandlerFunc(DevolucionsoporteExcel).
		Name("DevolucionsoporteExcel")

	// DEVOLUCION SOPORTE SERVICIO
	router.Path("/DevolucionsoporteservicioLista").HandlerFunc(DevolucionsoporteservicioLista).Name("DevolucionsoporteservicioLista")
	router.Path("/DevolucionsoporteservicioNuevo/{codigo}").HandlerFunc(DevolucionsoporteservicioNuevo).Name("DevolucionsoporteservicioNuevo")
	router.Path("/BodegaLlenar").HandlerFunc(BodegaLlenar).Name("BodegaLlenar")
	router.Path("/DevolucionsoporteservicioExiste/{codigo}").HandlerFunc(DevolucionsoporteservicioExiste).
		Name("DevolucionsoporteservicioExiste")
	router.Path("/DevolucionsoporteservicioEditar/{codigo}").HandlerFunc(DevolucionsoporteservicioEditar).
		Name("DevolucionsoporteservicioEditar")
	router.Path("/DevolucionsoporteservicioAgregar").HandlerFunc(DevolucionsoporteservicioAgregar).Name(
		"DevolucionsoporteservicioAgregar")
	router.Path("/DevolucionsoporteservicioBorrar/{codigo}").HandlerFunc(DevolucionsoporteservicioBorrar).
		Name("DevolucionsoporteservicioBorrar")
	router.Path("/DevolucionsoporteservicioEliminar/{codigo}").HandlerFunc(
		DevolucionsoporteservicioEliminar).Name("DevolucionsoporteservicioEliminar")
	router.Path("/DevolucionsoporteservicioPdf/{codigo}").HandlerFunc(DevolucionsoporteservicioPdf).Name(
		"DevolucionsoporteservicioPdf")

	router.Path("/DevolucionsoporteservicioTodosPdf").HandlerFunc(DevolucionsoporteservicioTodosPdf).
		Name("DevolucionsoporteservicioTodosPdf")
	router.Path("/DevolucionsoporteservicioExcel").HandlerFunc(DevolucionsoporteservicioExcel).
		Name("DevolucionsoporteservicioExcel")

	// PEDIDO FACTURA GASTO
	router.Path("/PedidofacturagastoLista").HandlerFunc(PedidofacturagastoLista).Name("PedidofacturagastoLista")
	router.Path("/PedidofacturagastoNuevo/{codigo}").HandlerFunc(PedidofacturagastoNuevo).Name("PedidofacturagastoNuevo")
	router.Path("/PedidofacturagastoExiste/{codigo}").HandlerFunc(PedidofacturagastoExiste).
		Name("PedidofacturagastoExiste")
	router.Path("/PedidofacturagastoEditar/{codigo}").HandlerFunc(PedidofacturagastoEditar).
		Name("PedidofacturagastoEditar")
	router.Path("/PedidofacturagastoAgregar").HandlerFunc(PedidofacturagastoAgregar).Name(
		"PedidofacturagastoAgregar")
	router.Path("/PedidofacturagastoBorrar/{codigo}").HandlerFunc(PedidofacturagastoBorrar).
		Name("PedidofacturagastoBorrar")
	router.Path("/PedidofacturagastoEliminar/{codigo}").HandlerFunc(
		PedidofacturagastoEliminar).Name("PedidofacturagastoEliminar")
	router.Path("/PedidofacturagastoPdf/{codigo}").HandlerFunc(PedidofacturagastoPdf).Name(
		"PedidofacturagastoPdf")

	router.Path("/PedidofacturagastoTodosPdf").HandlerFunc(PedidofacturagastoTodosPdf).
		Name("PedidofacturagastoTodosPdf")
	router.Path("/PedidofacturagastoExcel").HandlerFunc(PedidofacturagastoExcel).
		Name("PedidofacturagastoExcel")

	// FACTURA GASTO
	router.Path("/FacturagastoLista").HandlerFunc(FacturagastoLista).Name("FacturagastoLista")
	router.Path("/FacturagastoNuevo/{codigo}").HandlerFunc(FacturagastoNuevo).Name("FacturagastoNuevo")
	router.Path("/FacturagastoExiste/{codigo}").HandlerFunc(FacturagastoExiste).
		Name("FacturagastoExiste")
	router.Path("/FacturagastoEditar/{codigo}").HandlerFunc(FacturagastoEditar).
		Name("FacturagastoEditar")
	router.Path("/FacturagastoAgregar").HandlerFunc(FacturagastoAgregar).Name(
		"FacturagastoAgregar")
	router.Path("/FacturagastoBorrar/{codigo}").HandlerFunc(FacturagastoBorrar).
		Name("FacturagastoBorrar")
	router.Path("/FacturagastoEliminar/{codigo}").HandlerFunc(
		FacturagastoEliminar).Name("FacturagastoEliminar")
	router.Path("/FacturagastoPdf/{codigo}").HandlerFunc(FacturagastoPdf).Name(
		"FacturagastoPdf")

	router.Path("/FacturagastoTodosPdf").HandlerFunc(FacturagastoTodosPdf).
		Name("FacturagastoTodosPdf")
	router.Path("/FacturagastoExcel").HandlerFunc(FacturagastoExcel).
		Name("FacturagastoExcel")

	// DEVOLUCION FACTURA GASTO
	router.Path("/DevolucionfacturagastoLista").HandlerFunc(DevolucionfacturagastoLista).Name("DevolucionfacturagastoLista")
	router.Path("/DevolucionfacturagastoNuevo/{codigo}").HandlerFunc(DevolucionfacturagastoNuevo).Name("DevolucionfacturagastoNuevo")
	router.Path("/DevolucionfacturagastoExiste/{codigo}").HandlerFunc(DevolucionfacturagastoExiste).
		Name("DevolucionfacturagastoExiste")
	router.Path("/DevolucionfacturagastoEditar/{codigo}").HandlerFunc(DevolucionfacturagastoEditar).
		Name("DevolucionfacturagastoEditar")
	router.Path("/DevolucionfacturagastoAgregar").HandlerFunc(DevolucionfacturagastoAgregar).Name(
		"DevolucionfacturagastoAgregar")
	router.Path("/DevolucionfacturagastoBorrar/{codigo}").HandlerFunc(DevolucionfacturagastoBorrar).
		Name("DevolucionfacturagastoBorrar")
	router.Path("/DevolucionfacturagastoEliminar/{codigo}").HandlerFunc(
		DevolucionfacturagastoEliminar).Name("DevolucionfacturagastoEliminar")
	router.Path("/DevolucionfacturagastoPdf/{codigo}").HandlerFunc(DevolucionfacturagastoPdf).Name(
		"DevolucionfacturagastoPdf")

	router.Path("/DevolucionfacturagastoTodosPdf").HandlerFunc(DevolucionfacturagastoTodosPdf).
		Name("DevolucionfacturagastoTodosPdf")
	router.Path("/DevolucionfacturagastoExcel").HandlerFunc(DevolucionfacturagastoExcel).
		Name("DevolucionfacturagastoExcel")

	// TRASLADO AJUSTES
	router.Path("/TrasladoLista").HandlerFunc(TrasladoLista).Name("TrasladoLista")
	router.Path("/TrasladoNuevo/{codigo}").HandlerFunc(TrasladoNuevo).Name("TrasladoNuevo")
	router.Path("/BodegaLlenar").HandlerFunc(BodegaLlenar).Name("BodegaLlenar")
	router.Path("/TrasladoExiste/{codigo}").HandlerFunc(TrasladoExiste).
		Name("TrasladoExiste")
	router.Path("/TrasladoEditar/{codigo}").HandlerFunc(TrasladoEditar).
		Name("TrasladoEditar")
	router.Path("/TrasladoAgregar").HandlerFunc(TrasladoAgregar).Name(
		"TrasladoAgregar")
	router.Path("/TrasladoBorrar/{codigo}").HandlerFunc(TrasladoBorrar).
		Name("TrasladoBorrar")
	router.Path("/TrasladoEliminar/{codigo}").HandlerFunc(
		TrasladoEliminar).Name("TrasladoEliminar")
	router.Path("/TrasladoPdf/{codigo}").HandlerFunc(TrasladoPdf).Name(
		"TrasladoPdf")

	router.Path("/TrasladoTodosPdf").HandlerFunc(TrasladoTodosPdf).
		Name("TrasladoTodosPdf")
	router.Path("/TrasladoExcel").HandlerFunc(TrasladoExcel).
		Name("TrasladoExcel")

	// ARCHIVO FINANCIERO
	router.Path("/FinancieroNuevo/{codigo}").HandlerFunc(FinancieroNuevo).Name("FinancieroNuevo")
	router.Path("/FinancieroBuscar/{codigo}").HandlerFunc(FinancieroBuscar).
		Name("FinancieroBuscar")
	router.Path("/FinancieroActual/{codigo}").HandlerFunc(FinancieroActual).
		Name("FinancieroActual")
	router.Path("/FinancieroLista").HandlerFunc(FinancieroLista).Name("FinancieroLista")
	router.Path("/FinancieroExiste/{codigo:[0-9]+}").HandlerFunc(FinancieroExiste).
		Name("FinancieroExiste")
	router.Path("/FinancieroInsertar").HandlerFunc(FinancieroInsertar).Name(
		"FinancieroInsertar")
	router.Path("/FinancieroActualizar/{codigo:[0-9]+}").HandlerFunc(FinancieroActualizar).Name(
		"FinancieroActualizar")
	router.Path("/FinancieroBorrar/{codigo:[0-9]+}").HandlerFunc(FinancieroBorrar).Name(
		"FinancieroBorrar")
	router.Path("/FinancieroEliminar/{codigo:[0-9]+}").HandlerFunc(
		FinancieroEliminar).Name("FinancieroEliminar")
	router.Path("/FinancieroEditar/{codigo:[0-9]+}").HandlerFunc(FinancieroEditar).Name(
		"FinancieroEditar")
	router.Path("/FinancieroPdf/{codigo:[0-9]+}").HandlerFunc(FinancieroPdf).Name(
		"FinancieroPdf")

	router.Path("/FinancieroTodosPdf").HandlerFunc(FinancieroTodosPdf).
		Name("FinancieroTodosPdf")
	router.Path("/FinancieroExcel").HandlerFunc(FinancieroExcel).
		Name("FinancieroExcel")

	// ARCHIVO PROPIEDAD
	router.Path("/PropiedadNuevo/{codigo}").HandlerFunc(PropiedadNuevo).Name("PropiedadNuevo")
	router.Path("/PropiedadBuscar/{codigo}").HandlerFunc(PropiedadBuscar).
		Name("PropiedadBuscar")
	router.Path("/PropiedadActual/{codigo}").HandlerFunc(PropiedadActual).
		Name("PropiedadActual")
	router.Path("/PropiedadLista").HandlerFunc(PropiedadLista).Name("PropiedadLista")
	router.Path("/PropiedadExiste/{codigo:[0-9]+}").HandlerFunc(PropiedadExiste).
		Name("PropiedadExiste")
	router.Path("/PropiedadInsertar").HandlerFunc(PropiedadInsertar).Name(
		"PropiedadInsertar")
	router.Path("/PropiedadActualizar/{codigo:[0-9]+}").HandlerFunc(PropiedadActualizar).Name(
		"PropiedadActualizar")
	router.Path("/PropiedadBorrar/{codigo:[0-9]+}").HandlerFunc(PropiedadBorrar).Name(
		"PropiedadBorrar")
	router.Path("/PropiedadEliminar/{codigo:[0-9]+}").HandlerFunc(
		PropiedadEliminar).Name("PropiedadEliminar")
	router.Path("/PropiedadEditar/{codigo:[0-9]+}").HandlerFunc(PropiedadEditar).Name(
		"PropiedadEditar")
	router.Path("/PropiedadPdf/{codigo:[0-9]+}").HandlerFunc(PropiedadPdf).Name(
		"PropiedadPdf")

	router.Path("/PropiedadTodosPdf").HandlerFunc(PropiedadTodosPdf).
		Name("PropiedadTodosPdf")
	router.Path("/PropiedadExcel").HandlerFunc(PropiedadExcel).
		Name("PropiedadExcel")

	router.Path("/PropiedadGenerar").HandlerFunc(PropiedadGenerar).
		Name("PropiedadGenerar")

	router.Path("/PropiedadMes/{mes}").HandlerFunc(PropiedadMes).
		Name("PropiedadMes")

	// ARCHIVO DIFERIDO
	router.Path("/DiferidoNuevo/{codigo}").HandlerFunc(DiferidoNuevo).Name("DiferidoNuevo")
	router.Path("/DiferidoBuscar/{codigo}").HandlerFunc(DiferidoBuscar).
		Name("DiferidoBuscar")
	router.Path("/DiferidoActual/{codigo}").HandlerFunc(DiferidoActual).
		Name("DiferidoActual")
	router.Path("/DiferidoLista").HandlerFunc(DiferidoLista).Name("DiferidoLista")
	router.Path("/DiferidoExiste/{codigo:[0-9]+}").HandlerFunc(DiferidoExiste).
		Name("DiferidoExiste")
	router.Path("/DiferidoInsertar").HandlerFunc(DiferidoInsertar).Name(
		"DiferidoInsertar")
	router.Path("/DiferidoActualizar/{codigo:[0-9]+}").HandlerFunc(DiferidoActualizar).Name(
		"DiferidoActualizar")
	router.Path("/DiferidoBorrar/{codigo:[0-9]+}").HandlerFunc(DiferidoBorrar).Name(
		"DiferidoBorrar")
	router.Path("/DiferidoEliminar/{codigo:[0-9]+}").HandlerFunc(
		DiferidoEliminar).Name("DiferidoEliminar")
	router.Path("/DiferidoEditar/{codigo:[0-9]+}").HandlerFunc(DiferidoEditar).Name(
		"DiferidoEditar")
	router.Path("/DiferidoPdf/{codigo:[0-9]+}").HandlerFunc(DiferidoPdf).Name(
		"DiferidoPdf")

	router.Path("/DiferidoTodosPdf").HandlerFunc(DiferidoTodosPdf).
		Name("DiferidoTodosPdf")
	router.Path("/DiferidoExcel").HandlerFunc(DiferidoExcel).
		Name("DiferidoExcel")

	router.Path("/DiferidoGenerar").HandlerFunc(DiferidoGenerar).
		Name("DiferidoGenerar")

	router.Path("/DiferidoMes/{mes}").HandlerFunc(DiferidoMes).
		Name("DiferidoMes")

	// ARCHIVO CUOTA
	router.Path("/CuotaLista").HandlerFunc(CuotaLista).Name("CuotaLista")
	router.Path("/CuotaDatos/{monto}/{plazo}/{intereses}/{fechainicial}").HandlerFunc(CuotaDatos).Name("CuotaDatos")

	router.Path("/CuotaTodosPdf/{monto}/{plazo}/{intereses}/{fechainicial}").HandlerFunc(CuotaTodosPdf).
		Name("CuotaTodosPdf")
	router.Path("/CuotaExcel/{monto}/{plazo}/{intereses}/{fechainicial}").HandlerFunc(CuotaExcel).
		Name("CuotaExcel")

	// copia de base de datos
	router.Path("/CopiaGenerar/{panel}").HandlerFunc(CopiaGenerar).
		Name("CopiaGenerar")

	router.Path("/CopiaGenerarArchivo").HandlerFunc(CopiaGenerarArchivo).
		Name("CopiaGenerarArchivo")

	// LOCAL HOST 9002
	log.Println("Servidor Corriendo en " + ruta)
	if err := http.ListenAndServe(":9002", router); err != nil {
		log.Fatal(err)
	}
}

// FORMATO DE FECHA
var decoder = schema.NewDecoder()
var timeConverter = func(value string) reflect.Value {
	if v, err := time.Parse("2006-01-02", value); err == nil {

		return reflect.ValueOf(v)
	}
	return reflect.Value{} // this is the same as the private const invalidType
}

package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os/exec"
	"text/template"
	"time"
)

//C:\WINDOWS\system32>psql  -U postgres vivc0286 < "D://vivc0286-20220627215023.bak"

//psql --set ON_ERROR_STOP=on -U postgres -d vivc0286 -1 -f "D://vivc0286-20220627215023.bak"

// para copia se deben hacer dos proesos
// 1-crear archivo de passwords
//127.0.0.1:5432:*:postgres:Murc4505
//C:\Users\EOE\AppData\Roaming\postgresql\pgpass.conf
// 2- agregar al path ruta de pg_dump

func CopiaGenerar(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/copia/copiaGenerar.html")

	panel := mux.Vars(r)["panel"]
	varmap := map[string]interface{}{
		"hosting": ruta,
		"centro":  ListaCentro(),
		"panel":   panel,
	}
	tmp.Execute(w, varmap)
}

// generar archivo
func CopiaGenerarArchivo(w http.ResponseWriter, r *http.Request) {

	log.Println("Banco 0")
	var archivo string
	var fechahora = time.Now().Format("20060102150405")
	var ruta = "C://CrudGoOriginal//static//copiadb//"
	var nombrearchivo = dbname + "-" + fechahora + ".bak"

	archivo = ruta + nombrearchivo

	cmd := exec.Command("pg_dump", "-U", "postgres", "-h", "localhost", "-f", archivo, dbname)

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Output:\n%s\n", string(output))

	js, err := json.Marshal(MensajeBanco{nombrearchivo})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	//http.Redirect(w, r, "/CUENTADECOBROLista", 301)
}

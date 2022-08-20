package main

import (
	"bytes"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
)

// CONFIGURACION CONTABILIDAD TABLA
type configuracioncontabilidad struct {
	// INICIA ESTRUCTURA
	Cuentaefectivo             string
	Cuentanombreefectivo       string
	Cuentacliente              string
	Cuentanombrecliente        string
	Cuentatarjetadebito        string
	Cuentanombretarjetadebito  string
	Cuentatarjetacredito       string
	Cuentanombretarjetacredito string
	Cuentatransferencia        string
	Cuentanombretransferencia  string
	Cuentaproveedor            string
	Cuentanombreproveedor      string
	Cuentaajuste               string
	Cuentanombreajuste         string
	Phinicial                  string
	Textodescuento1            string
	Textodescuento2            string
	Textodescuento3            string
	Descuento1diainicial       string
	Descuento1diafinal         string
	Descuento2diainicial       string
	Descuento2diafinal         string
	Descuento3diainicial       string
	Descuento3diafinal         string
	Textoaviso1                string
	Textoaviso2                string
	Textoaviso3                string
	Textoaviso4                string
	Cuentautilidad             string
	Cuentautilidadnombre       string
	Trabajado                  string
	Trabajadonombre            string
	Transporte                 string
	Transportenombre           string
	Cesantias                  string
	Cesantiasnombre            string
	Intereses                  string
	Interesesnombre            string
	Prima                      string
	Primanombre                string
	Vacaciones                 string
	Vacacionesnombre           string
	Viaticos                   string
	Viaticosnombre             string
	Horasextras                string
	Horasextrasnombre          string
	Incapacidades              string
	Incapacidadesnombre        string
	Licencias                  string
	Licenciasnombre            string
	Bonificaciones             string
	Bonificacionesnombre       string
	Auxilios                   string
	Auxiliosnombre             string
	Huelgas                    string
	Huelgasnombre              string
	Conceptos                  string
	Conceptosnombre            string
	Compensaciones             string
	Compensacionesnombre       string
	Bonos                      string
	Bonosnombre                string
	Comisiones                 string
	Comisionesnombre           string
	Dotaciones                 string
	Dotacionesnombre           string
	Sostenimiento              string
	Sostenimientonombre        string
	Teletrabajo                string
	Teletrabajonombre          string
	Indemnizaciones            string
	Indemnizacionesnombre      string
	Salud                      string
	Saludnombre                string
	Pension                    string
	Pensionnombre              string
	Solidaridad                string
	Solidaridadnombre          string
	Subsistencia               string
	Subsistencianombre         string
	Sindicatos                 string
	Sindicatosnombre           string
	Sanciones                  string
	Sancionesnombre            string
	Libranzas                  string
	Libranzasnombre            string
	Terceros                   string
	Tercerosnombre             string
	Anticipos                  string
	Anticiposnombre            string
	Otras                      string
	Otrasnombre                string
	Retencion                  string
	Retencionnombre            string
	Afc                        string
	Afcnombre                  string
	Embargos                   string
	Embargosnombre             string
	Educacion                  string
	Educacionnombre            string
	Deuda                      string
	Deudanombre                string
	Neto                       string
	Netonombre                 string
	Cesantiascxp               string
	Cesantiascxpnombre         string
	Interesescxp               string
	Interesescxpnombre         string
	Primacxp                   string
	Primacxpnombre             string
	Vacacionescxp              string
	Vacacionescxpnombre        string
	Dotacionescxp              string
	Dotacionescxpnombre        string
	Riesgoscxp                 string
	Riesgoscxpnombre           string
	Icbfcxp                    string
	Icbfcxpnombre              string
	Senacxp                    string
	Senacxpnombre              string
	Cajacxp                    string
	Cajacxpnombre              string
	Riesgos                    string
	Riesgosnombre              string
	Caja                       string
	Cajanombre                 string
	Icbf                       string
	Icbfnombre                 string
	Sena                       string
	Senanombre                 string
	Saludgasto                 string
	Saludgastonombre           string
	Pensiongasto               string
	Pensiongastonombre         string
	Honorariogasto             string
	Honorariogastonombre       string
	Serviciogasto              string
	Serviciogastonombre        string
	Honorariocxpgasto          string
	Honorariocxpgastonombre    string
	Serviciocxpgasto           string
	Serviciocxpgastonombre     string
	Rethonorario               string
	Rethonorarionombre         string
	Retservicio                string
	Retservicionombre          string
	Pensionrais                string
	Pensionraisnombre          string
	Pensionvoluntaria          string
	Pensionvoluntarianombre    string
}

// CONFIGURACIONINVENTARIO NUEVO
func ConfiguracioncontabilidadNuevo(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/configuracioncontabilidad/configuracioncontabilidadNuevo.html",
		"vista/autocompleta/autocompletaPlandecuentaempresa."+
			"html")
	db := dbConn()
	panel := mux.Vars(r)["panel"]

	t := configuracioncontabilidad{}
	err := db.Get(&t, "SELECT * FROM configuracioncontabilidad ")
	switch err {
	case nil:
	case sql.ErrNoRows:
		log.Println("tercero NOT found, no error")
	default:
		log.Printf("tercero error: %s\n", err)
	}

	varmap := map[string]interface{}{
		"parametro": t,
		"hosting":   ruta,
		"panel":     panel,
	}

	tmp.Execute(w, varmap)

}

// CONFIGURACIONINVENTARIO INSERTAR
func ConfiguracioncontabilidadInsertar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	panel := r.FormValue("panel")
	if r.Method == "POST" {
		// INICIA INSERTAR CONFIGURACIONCONTABILIDAD
		Cuentaefectivo := r.FormValue("Cuentaefectivo")
		Cuentanombreefectivo := r.FormValue("Cuentanombreefectivo")
		Cuentatransferencia := r.FormValue("Cuentatransferencia")
		Cuentanombretransferencia := r.FormValue("Cuentanombretransferencia")
		Cuentacliente := r.FormValue("Cuentacliente")
		Cuentanombrecliente := r.FormValue("Cuentanombrecliente")
		Cuentatarjetadebito := r.FormValue("Cuentatarjetadebito")
		Cuentanombretarjetadebito := r.FormValue("Cuentanombretarjetadebito")
		Cuentatarjetacredito := r.FormValue("Cuentatarjetacredito")
		Cuentanombretarjetacredito := r.FormValue("Cuentanombretarjetacredito")
		Cuentaproveedor := r.FormValue("Cuentaproveedor")
		Cuentanombreproveedor := r.FormValue("Cuentanombreproveedor")
		Cuentaajuste := r.FormValue("Cuentaajuste")
		Cuentanombreajuste := r.FormValue("Cuentanombreajuste")
		Phinicial := r.FormValue("Phinicial")
		Textodescuento1 := r.FormValue("Textodescuento1")
		Textodescuento2 := r.FormValue("Textodescuento2")
		Textodescuento3 := r.FormValue("Textodescuento3")

		Descuento1diainicial := r.FormValue("Descuento1diainicial")
		Descuento1diafinal := r.FormValue("Descuento1diafinal")
		Descuento2diainicial := r.FormValue("Descuento2diainicial")
		Descuento2diafinal := r.FormValue("Descuento2diafinal")
		Descuento3diainicial := r.FormValue("Descuento3diainicial")
		Descuento3diafinal := r.FormValue("Descuento3diafinal")

		Textoaviso1 := r.FormValue("Textoaviso1")
		Textoaviso2 := r.FormValue("Textoaviso2")
		Textoaviso3 := r.FormValue("Textoaviso3")
		Textoaviso4 := r.FormValue("Textoaviso4")
		Cuentanombreefectivo = Titulo(Cuentanombreefectivo)
		Cuentanombrecliente = Titulo(Cuentanombrecliente)
		Cuentanombretarjetadebito = Titulo(Cuentanombretarjetadebito)
		Cuentanombretarjetacredito = Titulo(Cuentanombretarjetacredito)
		Cuentanombretransferencia = Titulo(Cuentanombretransferencia)
		Cuentanombreproveedor = Titulo(Cuentanombreproveedor)
		Cuentanombreajuste = Titulo(Cuentanombreajuste)
		Textodescuento1 = Titulo(Textodescuento1)
		Textodescuento2 = Titulo(Textodescuento2)
		Textodescuento3 = Titulo(Textodescuento3)
		Textoaviso1 = Mayuscula(Textoaviso1)
		Textoaviso2 = Mayuscula(Textoaviso2)
		Textoaviso3 = Mayuscula(Textoaviso3)
		Textoaviso4 = Mayuscula(Textoaviso4)

		Cuentautilidad := r.FormValue("Cuentautilidad")
		Cuentautilidadnombre := r.FormValue("Cuentautilidadnombre")

		Trabajado := r.FormValue("Trabajado")
		Trabajadonombre := r.FormValue("Trabajadonombre")
		Transporte := r.FormValue("Transporte")
		Transportenombre := r.FormValue("Transportenombre")
		Cesantias := r.FormValue("Cesantias")
		Cesantiasnombre := r.FormValue("Cesantiasnombre")
		Intereses := r.FormValue("Intereses")
		Interesesnombre := r.FormValue("Interesesnombre")
		Prima := r.FormValue("Prima")
		Primanombre := r.FormValue("Primanombre")
		Vacaciones := r.FormValue("Vacaciones")
		Vacacionesnombre := r.FormValue("Vacacionesnombre")
		Viaticos := r.FormValue("Viaticos")
		Viaticosnombre := r.FormValue("Viaticosnombre")
		Horasextras := r.FormValue("Horasextras")
		Horasextrasnombre := r.FormValue("Horasextrasnombre")
		Incapacidades := r.FormValue("Incapacidades")
		Incapacidadesnombre := r.FormValue("Incapacidadesnombre")
		Licencias := r.FormValue("Licencias")
		Licenciasnombre := r.FormValue("Licenciasnombre")
		Bonificaciones := r.FormValue("Bonificaciones")
		Bonificacionesnombre := r.FormValue("Bonificacionesnombre")
		Auxilios := r.FormValue("Auxilios")
		Auxiliosnombre := r.FormValue("Auxiliosnombre")
		Huelgas := r.FormValue("Huelgas")
		Huelgasnombre := r.FormValue("Huelgasnombre")
		Conceptos := r.FormValue("Conceptos")
		Conceptosnombre := r.FormValue("Conceptosnombre")
		Compensaciones := r.FormValue("Compensaciones")
		Compensacionesnombre := r.FormValue("Compensacionesnombre")
		Bonos := r.FormValue("Bonos")
		Bonosnombre := r.FormValue("Bonosnombre")
		Comisiones := r.FormValue("Comisiones")
		Comisionesnombre := r.FormValue("Comisionesnombre")
		Dotaciones := r.FormValue("Dotaciones")
		Dotacionesnombre := r.FormValue("Dotacionesnombre")
		Sostenimiento := r.FormValue("Sostenimiento")
		Sostenimientonombre := r.FormValue("Sostenimientonombre")
		Teletrabajo := r.FormValue("Teletrabajo")
		Teletrabajonombre := r.FormValue("Teletrabajonombre")
		Indemnizaciones := r.FormValue("Indemnizaciones")
		Indemnizacionesnombre := r.FormValue("Indemnizacionesnombre")
		Salud := r.FormValue("Salud")
		Saludnombre := r.FormValue("Saludnombre")
		Pension := r.FormValue("Pension")
		Pensionnombre := r.FormValue("Pensionnombre")
		Solidaridad := r.FormValue("Solidaridad")
		Solidaridadnombre := r.FormValue("Solidaridadnombre")
		Subsistencia := r.FormValue("Subsistencia")
		Subsistencianombre := r.FormValue("Subsistencianombre")
		Sindicatos := r.FormValue("Sindicatos")
		Sindicatosnombre := r.FormValue("Sindicatosnombre")
		Sanciones := r.FormValue("Sanciones")
		Sancionesnombre := r.FormValue("Sancionesnombre")
		Libranzas := r.FormValue("Libranzas")
		Libranzasnombre := r.FormValue("Libranzasnombre")
		Terceros := r.FormValue("Terceros")
		Tercerosnombre := r.FormValue("Tercerosnombre")
		Anticipos := r.FormValue("Anticipos")
		Anticiposnombre := r.FormValue("Anticiposnombre")
		Otras := r.FormValue("Otras")
		Otrasnombre := r.FormValue("Otrasnombre")
		Retencion := r.FormValue("Retencion")
		Retencionnombre := r.FormValue("Retencionnombre")
		Afc := r.FormValue("Afc")
		Afcnombre := r.FormValue("Afcnombre")
		Embargos := r.FormValue("Embargos")
		Embargosnombre := r.FormValue("Embargosnombre")
		Educacion := r.FormValue("Educacion")
		Educacionnombre := r.FormValue("Educacionnombre")
		Deuda := r.FormValue("Deuda")
		Deudanombre := r.FormValue("Deudanombre")
		Neto := r.FormValue("Neto")
		Netonombre := r.FormValue("Netonombre")
		Cesantiascxp := r.FormValue("Cesantiascxp")
		Cesantiascxpnombre := r.FormValue("Cesantiascxpnombre")
		Interesescxp := r.FormValue("Interesescxp")
		Interesescxpnombre := r.FormValue("Interesesnombre")
		Primacxp := r.FormValue("Primacxp")
		Primacxpnombre := r.FormValue("Primacxpnombre")
		Vacacionescxp := r.FormValue("Vacacionescxp")
		Vacacionescxpnombre := r.FormValue("Vacacionescxpnombre")
		Dotacionescxp := r.FormValue("Dotacionescxp")
		Dotacionescxpnombre := r.FormValue("Dotacionescxpnombre")
		Riesgoscxp := r.FormValue("Riesgoscxp")
		Riesgoscxpnombre := r.FormValue("Riesgoscxpnombre")
		Icbfcxp := r.FormValue("Icbfcxp")
		Icbfcxpnombre := r.FormValue("Icbfcxpnombre")
		Senacxp := r.FormValue("Senacxp")
		Senacxpnombre := r.FormValue("Senacxpnombre")
		Cajacxp := r.FormValue("Cajacxp")
		Cajacxpnombre := r.FormValue("Cajacxpnombre")
		Riesgos := r.FormValue("Riesgos")
		Riesgosnombre := r.FormValue("Riesgosnombre")
		Caja := r.FormValue("Caja")
		Cajanombre := r.FormValue("Cajanombre")
		Icbf := r.FormValue("Icbf")
		Icbfnombre := r.FormValue("Icbfnombre")
		Sena := r.FormValue("Sena")
		Senanombre := r.FormValue("Senanombre")
		Saludgasto := r.FormValue("Saludgasto")
		Saludgastonombre := r.FormValue("Saludgastonombre")
		Pensiongasto := r.FormValue("Pensiongasto")
		Pensiongastonombre := r.FormValue("Pensiongastonombre")
		Honorariogasto := r.FormValue("Honorariogasto")
		Honorariogastonombre := r.FormValue("Honorariogastonombre")
		Serviciogasto := r.FormValue("Serviciogasto")
		Serviciogastonombre := r.FormValue("Serviciogastonombre")
		Honorariocxpgasto := r.FormValue("Honorariocxpgasto")
		Honorariocxpgastonombre := r.FormValue("Honorariocxpgastonombre")
		Serviciocxpgasto := r.FormValue("Serviciocxpgasto")
		Serviciocxpgastonombre := r.FormValue("Serviciocxpgastonombre")
		Rethonorario := r.FormValue("Rethonorario")
		Rethonorarionombre := r.FormValue("Rethonorarionombre")
		Retservicio := r.FormValue("Retservicio")
		Retservicionombre := r.FormValue("Retservicionombre")
		Pensionrais := r.FormValue("Pensionrais")
		Pensionraisnombre := r.FormValue("Pensionraisnombre")
		Pensionvoluntaria := r.FormValue("Pensionvoluntaria")
		Pensionvoluntarianombre := r.FormValue("Pensionvoluntarianombre")

		var consulta = "INSERT INTO configuracioncontabilidad("
		// INICIA CONSULTA COMPRA
		consulta += "Cuentaefectivo,"
		consulta += "Cuentanombreefectivo,"
		consulta += "Cuentacliente,"
		consulta += "Cuentanombrecliente,"
		consulta += "Cuentatarjetadebito,"
		consulta += "Cuentanombretarjetadebito,"
		consulta += "Cuentatarjetacredito,"
		consulta += "Cuentanombretarjetacredito,"
		consulta += "Cuentatransferencia,"
		consulta += "Cuentanombretransferencia,"
		consulta += "Cuentaproveedor,"
		consulta += "Cuentanombreproveedor,"
		consulta += "Cuentaajuste,"
		consulta += "Cuentanombreajuste,"
		consulta += "Phinicial,"
		consulta += "Textodescuento1,"
		consulta += "Textodescuento2,"
		consulta += "Textodescuento3,"
		consulta += "Descuento1diainicial,"
		consulta += "Descuento1diafinal,"
		consulta += "Descuento2diainicial,"
		consulta += "Descuento2diafinal,"
		consulta += "Descuento3diainicial,"
		consulta += "Descuento3diafinal,"
		consulta += "Textoaviso1,"
		consulta += "Textoaviso2,"
		consulta += "Textoaviso3,"
		consulta += "Textoaviso4,"
		consulta += "Cuentautilidad,"
		consulta += "Cuentautilidadnombre,"
		consulta += "Trabajado,"
		consulta += "Trabajadonombre,"
		consulta += "Transporte,"
		consulta += "Transportenombre,"
		consulta += "Cesantias,"
		consulta += "Cesantiasnombre,"
		consulta += "Intereses,"
		consulta += "Interesesnombre,"
		consulta += "Prima,"
		consulta += "Primanombre,"
		consulta += "Vacaciones,"
		consulta += "Vacacionesnombre,"
		consulta += "Viaticos,"
		consulta += "Viaticosnombre,"
		consulta += "Horasextras,"
		consulta += "Horasextrasnombre,"
		consulta += "Incapacidades,"
		consulta += "Incapacidadesnombre,"
		consulta += "Licencias,"
		consulta += "Licenciasnombre,"
		consulta += "Bonificaciones,"
		consulta += "Bonificacionesnombre,"
		consulta += "Auxilios,"
		consulta += "Auxiliosnombre,"
		consulta += "Huelgas,"
		consulta += "Huelgasnombre,"
		consulta += "Conceptos,"
		consulta += "Conceptosnombre,"
		consulta += "Compensaciones,"
		consulta += "Compensacionesnombre,"
		consulta += "Bonos,"
		consulta += "Bonosnombre,"
		consulta += "Comisiones,"
		consulta += "Comisionesnombre,"
		consulta += "Dotaciones,"
		consulta += "Dotacionesnombre,"
		consulta += "Sostenimiento,"
		consulta += "Sostenimientonombre,"
		consulta += "Teletrabajo,"
		consulta += "Teletrabajonombre,"
		consulta += "Indemnizaciones,"
		consulta += "Indemnizacionesnombre,"
		consulta += "Salud,"
		consulta += "Saludnombre,"
		consulta += "Pension,"
		consulta += "Pensionnombre,"
		consulta += "Solidaridad,"
		consulta += "Solidaridadnombre,"
		consulta += "Subsistencia,"
		consulta += "Subsistencianombre,"
		consulta += "Sindicatos,"
		consulta += "Sindicatosnombre,"
		consulta += "Sanciones,"
		consulta += "Sancionesnombre,"
		consulta += "Libranzas,"
		consulta += "Libranzasnombre,"
		consulta += "Terceros,"
		consulta += "Tercerosnombre,"
		consulta += "Anticipos,"
		consulta += "Anticiposnombre,"
		consulta += "Otras,"
		consulta += "Otrasnombre,"
		consulta += "Retencion,"
		consulta += "Retencionnombre,"
		consulta += "Afc,"
		consulta += "Afcnombre,"
		consulta += "Embargos,"
		consulta += "Embargosnombre,"
		consulta += "Educacion,"
		consulta += "Educacionnombre,"
		consulta += "Deuda,"
		consulta += "Deudanombre,"
		consulta += "Neto,"
		consulta += "Netonombre,"
		consulta += "Cesantiascxp,"
		consulta += "Cesantiascxpnombre,"
		consulta += "Interesescxp,"
		consulta += "Interesescxpnombre,"
		consulta += "Primacxp,"
		consulta += "Primacxpnombre,"
		consulta += "Vacacionescxp,"
		consulta += "Vacacionescxpnombre,"
		consulta += "Dotacionescxp,"
		consulta += "Dotacionescxpnombre,"
		consulta += "Riesgoscxp,"
		consulta += "Riesgoscxpnombre,"
		consulta += "Icbfcxp,"
		consulta += "Icbfcxpnombre,"
		consulta += "Senacxp,"
		consulta += "Senacxpnombre,"
		consulta += "Cajacxp,"
		consulta += "Cajacxpnombre,"
		consulta += "Riesgos,"
		consulta += "Riesgosnombre,"
		consulta += "Caja,"
		consulta += "Cajanombre,"
		consulta += "Icbf,"
		consulta += "Icbfnombre,"
		consulta += "Sena,"
		consulta += "Senanombre,"
		consulta += "Saludgasto,"
		consulta += "Saludgastonombre,"
		consulta += "Pensiongasto,"
		consulta += "Pensiongastonombre,"
		consulta += "Honorariogasto,"
		consulta += "Honorariogastonombre,"
		consulta += "Serviciogasto,"
		consulta += "Serviciogastonombre,"
		consulta += "Honorariocxpgasto,"
		consulta += "Honorariocxpgastonombre,"
		consulta += "Serviciocxpgasto,"
		consulta += "Serviciocxpgastonombre,"
		consulta += "Rethonorario,"
		consulta += "Rethonorarionombre,"
		consulta += "Retservicio,"
		consulta += "Retservicionombre,"
		consulta += "Pensionrais,"
		consulta += "Pensionraisnombre,"
		consulta += "Pensionvoluntaria,"
		consulta += "Pensionvoluntarianombre"
		consulta += ")VALUES("
		consulta += parametros(150)
		consulta += ")"

		delForm, err := db.Prepare("DELETE from configuracioncontabilidad")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec()
		insForm, err := db.Prepare(consulta)
		if err != nil {
			log.Println(err.Error())
			panic(err.Error())
		}
		_, err = insForm.Exec(
			// INICIA BORRAR COMPRA
			Cuentaefectivo,
			Cuentanombreefectivo,
			Cuentacliente,
			Cuentanombrecliente,
			Cuentatarjetadebito,
			Cuentanombretarjetadebito,
			Cuentatarjetacredito,
			Cuentanombretarjetacredito,
			Cuentatransferencia,
			Cuentanombretransferencia,
			Cuentaproveedor,
			Cuentanombreproveedor,
			Cuentaajuste,
			Cuentanombreajuste,
			Phinicial,
			Textodescuento1,
			Textodescuento2,
			Textodescuento3,
			Descuento1diainicial,
			Descuento1diafinal,
			Descuento2diainicial,
			Descuento2diafinal,
			Descuento3diainicial,
			Descuento3diafinal,
			Textoaviso1,
			Textoaviso2,
			Textoaviso3,
			Textoaviso4,
			Cuentautilidad,
			Cuentautilidadnombre,
			Trabajado,
			Trabajadonombre,
			Transporte,
			Transportenombre,
			Cesantias,
			Cesantiasnombre,
			Intereses,
			Interesesnombre,
			Prima,
			Primanombre,
			Vacaciones,
			Vacacionesnombre,
			Viaticos,
			Viaticosnombre,
			Horasextras,
			Horasextrasnombre,
			Incapacidades,
			Incapacidadesnombre,
			Licencias,
			Licenciasnombre,
			Bonificaciones,
			Bonificacionesnombre,
			Auxilios,
			Auxiliosnombre,
			Huelgas,
			Huelgasnombre,
			Conceptos,
			Conceptosnombre,
			Compensaciones,
			Compensacionesnombre,
			Bonos,
			Bonosnombre,
			Comisiones,
			Comisionesnombre,
			Dotaciones,
			Dotacionesnombre,
			Sostenimiento,
			Sostenimientonombre,
			Teletrabajo,
			Teletrabajonombre,
			Indemnizaciones,
			Indemnizacionesnombre,
			Salud,
			Saludnombre,
			Pension,
			Pensionnombre,
			Solidaridad,
			Solidaridadnombre,
			Subsistencia,
			Subsistencianombre,
			Sindicatos,
			Sindicatosnombre,
			Sanciones,
			Sancionesnombre,
			Libranzas,
			Libranzasnombre,
			Terceros,
			Tercerosnombre,
			Anticipos,
			Anticiposnombre,
			Otras,
			Otrasnombre,
			Retencion,
			Retencionnombre,
			Afc,
			Afcnombre,
			Embargos,
			Embargosnombre,
			Educacion,
			Educacionnombre,
			Deuda,
			Deudanombre,
			Neto,
			Netonombre,
			Cesantiascxp,
			Cesantiascxpnombre,
			Interesescxp,
			Interesescxpnombre,
			Primacxp,
			Primacxpnombre,
			Vacacionescxp,
			Vacacionescxpnombre,
			Dotacionescxp,
			Dotacionescxpnombre,
			Riesgoscxp,
			Riesgoscxpnombre,
			Icbfcxp,
			Icbfcxpnombre,
			Senacxp,
			Senacxpnombre,
			Cajacxp,
			Cajacxpnombre,
			Riesgos,
			Riesgosnombre,
			Caja,
			Cajanombre,
			Icbf,
			Icbfnombre,
			Sena,
			Senanombre,
			Saludgasto,
			Saludgastonombre,
			Pensiongasto,
			Pensiongastonombre,
			Honorariogasto,
			Honorariogastonombre,
			Serviciogasto,
			Serviciogastonombre,
			Honorariocxpgasto,
			Honorariocxpgastonombre,
			Serviciocxpgasto,
			Serviciocxpgastonombre,
			Rethonorario,
			Rethonorarionombre,
			Retservicio,
			Retservicionombre,
			Pensionrais,
			Pensionraisnombre,
			Pensionvoluntaria,
			Pensionvoluntarianombre)

		// TERMINA BORRAR SERVICIO
		if err != nil {
			log.Println(err.Error())
			panic(err)
		}
		log.Println("Nuevo Registro:" + Cuentaefectivo + "," + Cuentanombreefectivo)
	}
	http.Redirect(w, r, "/ConfiguracioncontabilidadNuevo/"+panel, 301)
}

// INICIA CONFIGURACIONINVENTARIO PDF
func ConfiguracioncontabilidadPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Pagocuentaefectivo := mux.Vars(r)["pagocuentaefectivo"]
	t := configuracioncontabilidad{}
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	err := db.Get(&t, "SELECT * FROM configuracioncontabilidad where pagocuentaefectivo=$1", Pagocuentaefectivo)
	if err != nil {
		log.Fatalln(err)
	}
	var buf bytes.Buffer
	var err1 error
	pdf := gofpdf.New("P", "mm", "Letter", cnFontDir)
	ene := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetHeaderFunc(func() {
		pdf.Image(imageFile("logo.png"), 20, 20, 40, 0, false,
			"", 0, "")
		pdf.SetY(15)
		//pdf.AddFont("Helvetica", "", "cp1251.map")
		pdf.SetFont("Helvetica", "", 10)
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
		log.Println("tercero 3")
		pdf.CellFormat(190, 10, ene(c.NombreCiudad+" - "+c.NombreDepartamento), "0", 0, "C", false, 0,
			"")
		log.Println("tercero 4")
		pdf.Ln(10)
		pdf.CellFormat(190, 10, "Datos Centro de Costos", "0", 0,
			"C", false, 0, "")
		pdf.Ln(10)
	})
	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.SetFont("Arial", "", 10)
	pdf.SetX(30)
	pdf.CellFormat(40, 4, "Cuenta", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Cuentaefectivo, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(30)
	pdf.CellFormat(40, 4, "Nombre:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Cuentanombreefectivo, "", 0,
		"", false, 0, "")

	pdf.SetFooterFunc(func() {
		pdf.SetY(-20)
		pdf.SetFont("Arial", "", 9)
		pdf.SetX(30)
		pdf.CellFormat(90, 10, "Sadconf.com", "", 0,
			"L", false, 0, "")
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

// TERMINA CONFIGURACIONINVENTARIO PDF

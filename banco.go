package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
	//"strconv"
)

type MensajeBanco struct {
	Mensaje string `json:"Mensaje"`
}

// CONCILIACION TABLA
type Facturas struct {
	Fila    string  `json:"Fila"`
	Cuenta  string  `json:"Cuenta"`
	Nombre  string  `json:"Nombre"`
	Fecha   string  `json:"Fecha"`
	Tipo    string  `json:"Tipo"`
	Factura string  `json:"Factura"`
	Neto    float64 `json:"Neto"`
	Saldo   float64 `json:"Saldo"`
	Abono   float64 `json:"Abono"`
	Total   float64 `json:"Total"`
}
type PagoBanco struct {
	Documento            string             `json:"Documento"`
	Numero               string             `json:"Numero"`
	Centro               string             `json:"Centro"`
	Fecha                string             `json:"Fecha"`
	Tercero              string             `json:"Tercero"`
	Consignacion         string             `json:"Consignacion"`
	Valorefectivo        string             `json:"Valorefectivo"`
	Cuentaefectivo       string             `json:"Cuentaefectivo"`
	Valortarjetadebito   string             `json:"Valortarjetadebito"`
	Cuentatarjetadebito  string             `json:"Cuentatarjetadebito"`
	Valortarjetacredito  string             `json:"Valortarjetacredito"`
	Cuentatarjetacredito string             `json:"Cuentatarjetacredito"`
	Valorcliente         string             `json:"Valorcliente"`
	Cuentacliente        string             `json:"Cuentacliente"`
	Cuentatransferencia  string             `json:"Cuentatransferencia"`
	Valortransferencia   string             `json:"Valortransferencia"`
	Cuentaajuste         string             `json:"Cuentaajuste"`
	Valorajuste          string             `json:"Valorajuste"`
	Detalle              []PagoBancoDetalle `json:"Detalle"`
}

type PagoBancoDetalle struct {
	Factura string `json:"Factura"`
	Abono   string `json:"Abono"`
	Cuenta  string `json:"Cuenta"`
}

type ListaDeuda struct {
	Tipo   string    `json:"Tipo"`
	Codigo string    `json:"Codigo"`
	Fecha  time.Time `json:"Fecha"`
	Neto   string    `json:"Neto"`
	Cuenta string    `json:"Cuenta"`
	Nombre string    `json:"Nombre"`
}

type ListaPago struct {
	Factura string  `json:"Codigo"`
	Avance  float64 `json:"Avance"`
	Cuenta  string  `json:"Cuenta"`
}

type tercerobanco struct {
	Codigo string
	Dv     string
	Nombre string
	Saldo  float64
}

type nominaDatosPersonalesbanco struct {
	Codigo string
	Dv     string
	Nombre string
	Saldo  float64
}

type tercerosaldo struct {
	Saldo float64
}

type nominaDatosPersonalessaldo struct {
	Saldo float64
}

func SaldoTerceroBanco(codigotercero string) float64 {
	parametroscontabilidad := configuracioncontabilidad{}
	err := db.Get(&parametroscontabilidad, "SELECT * FROM configuracioncontabilidad")
	if err != nil {
		panic(err.Error())
	}

	var consulta string
	consulta = ""
	consulta = "select distinct  coalesce(sum(credito-debito),0) as saldo  from comprobantedetalle where tercero=$1 and cuenta=$2"

	listadoDatos := tercerosaldo{}
	err1 := db.Get(&listadoDatos, consulta, codigotercero, parametroscontabilidad.Cuentacliente)
	if err1 != nil {
		panic(err1.Error())
	}
	log.Println("saldo favor" + FormatoFlotante(listadoDatos.Saldo))
	return listadoDatos.Saldo

}

func SaldoNominaDatosPersonalesBanco(codigotercero string) float64 {
	parametroscontabilidad := configuracioncontabilidad{}
	err := db.Get(&parametroscontabilidad, "SELECT * FROM configuracioncontabilidad")
	if err != nil {
		panic(err.Error())
	}

	var consulta string
	consulta = ""
	consulta = "select distinct  coalesce(sum(credito-debito),0) as saldo  from comprobantedetalle where tercero=$1 and cuenta=$2"

	listadoDatos := nominaDatosPersonalessaldo{}
	err1 := db.Get(&listadoDatos, consulta, codigotercero, parametroscontabilidad.Cuentacliente)
	if err1 != nil {
		panic(err1.Error())
	}
	log.Println("saldo favor" + FormatoFlotante(listadoDatos.Saldo))
	return listadoDatos.Saldo

}

func TerceroActualBanco(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	Codigo = Quitacoma(Codigo)

	t := tercerobanco{}
	var res []tercerobanco
	err := db.Get(&t, "SELECT codigo,dv,nombre,999999999 as saldo FROM tercero where codigo=$1", Codigo)

	switch err {
	case nil:
		log.Printf("tercero found: %+v\n", t)
	case sql.ErrNoRows:
		log.Println("tercero NOT found, no error")
	default:
		log.Printf("tercero error: %s\n", err)
	}

	t.Saldo = SaldoTerceroBanco(Codigo)

	log.Println("codigo nombre99" + t.Codigo + t.Nombre)
	res = append(res, t)
	data, err := json.Marshal(res)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func NominaDatosPersonalesActualBanco(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	Codigo = Quitacoma(Codigo)

	t := nominaDatosPersonalesbanco{}
	var res []nominaDatosPersonalesbanco
	err := db.Get(&t, "SELECT codigo,dv,nombre,999999999 as saldo FROM nominaDatosPersonales where codigo=$1", Codigo)

	switch err {
	case nil:
		log.Printf("tercero found: %+v\n", t)
	case sql.ErrNoRows:
		log.Println("tercero NOT found, no error")
	default:
		log.Printf("tercero error: %s\n", err)
	}

	t.Saldo = SaldoNominaDatosPersonalesBanco(Codigo)

	log.Println("codigo nombre99" + t.Codigo + t.Nombre)
	res = append(res, t)
	data, err := json.Marshal(res)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func BancoDato(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var miperiodo = periodoSesion(r)
	terceroParametro := mux.Vars(r)["tercero"]
	documentoParametro := mux.Vars(r)["documento"]
	fecha := mux.Vars(r)["fecha"]
	log.Printf(terceroParametro)

	// parametros inventario
	parametrosinventario := configuracioninventario{}
	err := db.Get(&parametrosinventario, "SELECT * FROM configuracioninventario")
	if err != nil {
		panic(err.Error())
	}

	var consulta = ""
	// listado de deudas
	listadeuda := []ListaDeuda{}
	if documentoParametro == "1" {
		consulta = " select cuenta,plandecuentaempresa.nombre as nombre, "
		consulta += " documento.nombre as tipo, "
		consulta += " (debito) as neto,fecha,numero as codigo from comprobantedetalle  "
		consulta += " inner join documento on  documento.codigo=comprobantedetalle.documento  "
		consulta += " inner join plandecuentaempresa on  plandecuentaempresa.codigo=comprobantedetalle.cuenta  "
		consulta += " where "
		consulta += " extract(year from fecha)<=" + miperiodo + " "
		consulta += " and  fecha<=$2 "
		consulta += " and (substring(cuenta,1,2)='13' and comprobantedetalle.tercero=$1 and debito>0 )"
	} else {

		consulta = "  select cuenta,plandecuentaempresa.nombre as nombre, "
		consulta += " documento.nombre as tipo, "
		consulta += " (credito) as neto,fecha,numero as codigo from comprobantedetalle  "
		consulta += " inner join documento on  documento.codigo=comprobantedetalle.documento  "
		consulta += " inner join plandecuentaempresa on  plandecuentaempresa.codigo=comprobantedetalle.cuenta  "
		consulta += " where "
		consulta += " extract(year from fecha)<=" + miperiodo + " "
		consulta += " and fecha<=$2 "
		consulta += " and (substring(cuenta,1,2)='23' and comprobantedetalle.tercero=$1 and credito>0 )"

	}

	consulta += " order by fecha"

	log.Println("Datos Facturación")

	var siexiste bool
	err = db.Select(&listadeuda, consulta, terceroParametro, fecha)
	switch err {
	//resltadvaa
	case nil:
		log.Printf("Datos existe Facuras")
		siexiste = true
	case sql.ErrNoRows:
		log.Println("Datos no encontrados Facturas")
	default:
		log.Printf("datos error: %s\n", err)
	}
	listafactura := []Facturas{}

	// listado de pagos
	listapago := []ListaPago{}
	//	var cuentaDeuda string
	//cuentaDeuda=""
	if documentoParametro == "1" {
		//	consulta="select sum(credito) as avance,factura from comprobantedetalle where cuenta=$2 and tercero=$1 group by factura"
		consulta = "select cuenta,sum(credito) as avance,factura from comprobantedetalle where "
		consulta += " substring(cuenta,1,2)='13' and tercero=$1 group by factura,cuenta"

		//cuentaDeuda=parametrosinventario.Ventacuentacliente
	} else {
		//consulta="select sum(debito) as avance,factura from comprobantedetalle where cuenta=$2 and tercero=$1 group by factura"
		consulta = "select cuenta,sum(debito) as avance,factura from comprobantedetalle where "
		consulta += " substring(cuenta,1,2)='23' and tercero=$1 group by factura,cuenta "
		//cuentaDeuda=parametrosinventario.Compracuentaproveedor
	}
	//var siexistepago bool
	err = db.Select(&listapago, consulta, terceroParametro)
	switch err {
	//resltadvaa
	case nil:
		log.Printf("Datos existe Pagos")
		//siexistepago = true
	case sql.ErrNoRows:
		log.Println("Datos no encontrados pagos")
	default:
		log.Printf("datos error111: %s\n", err)
	}

	// recorro facturas
	log.Println("Datos consulta")
	var totalPago float64
	var totalSaldo float64
	for _, miDeuda := range listadeuda {

		totalPago = 0

		for _, miPago := range listapago {
			if miPago.Factura == miDeuda.Codigo && miPago.Cuenta == miDeuda.Cuenta {
				totalPago += miPago.Avance
			}
		}

		totalSaldo = Flotante(miDeuda.Neto) - totalPago
		if totalSaldo > 0 {
			//strconv.Itoa(i)
			listafactura = append(listafactura, Facturas{"", miDeuda.Cuenta, miDeuda.Nombre, miDeuda.Fecha.Format("02/01/2006"), miDeuda.Tipo, miDeuda.Codigo, Flotante(miDeuda.Neto), totalSaldo, 0, totalSaldo})

		}
	}

	if siexiste == false {
		var slice []string
		slice = make([]string, 0)
		data, _ := json.Marshal(slice)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	} else {
		data, _ := json.Marshal(listafactura)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

func BancoDatoAgregar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	log.Println("Banco 0")
	//var tempCuentadecobro cuentadecobro
	parametrosinventario := configuracioninventario{}
	err := db.Get(&parametrosinventario, "SELECT * FROM configuracioninventario")
	if err != nil {
		panic(err.Error())
	}

	log.Println("Banco 1")

	listaBanco := PagoBanco{}
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// carga informacion de la CUENTADECOBRO
	err = json.Unmarshal(b, &listaBanco)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	const (
		layoutISO = "2006-01-02"
	)
	fechaDate, _ := time.Parse(layoutISO, listaBanco.Fecha)
	consignacionDate, _ := time.Parse(layoutISO, listaBanco.Consignacion)
	//var CuentaFactura string
	//
	//// cuenta abono
	//if listaBanco.Documento=="1"{
	//	CuentaFactura=parametrosinventario.Ventacuentacliente
	//}else {
	//	CuentaFactura=parametrosinventario.Compracuentaproveedor
	//}

	log.Println("Banco 2")

	// traer tercero
	t := tercero{}
	err1 := db.Get(&t, "SELECT * FROM tercero where codigo=$1", listaBanco.Tercero)
	if err1 != nil {
		log.Fatalln(err1)
	}

	var totalDebito float64
	var totalCredito float64
	totalDebito = 0
	totalCredito = 0
	miComprobanteDetalle := []comprobantedetalle{}
	miComprobanteDetalleDebito := []comprobantedetalle{}
	miComprobanteDetalleCredito := []comprobantedetalle{}
	miComprobanteDetalleFinal := []comprobantedetalle{}

	var miFilaComprobante int
	miFilaComprobante = 0
	// INSERTA DETALLE
	var miCentro = listaBanco.Centro

	listaBanco.Numero = NumeroDocumento(listaBanco.Documento)

	var debito string
	var credito string

	log.Println("Banco 3")
	// efectivo
	if Flotante(listaBanco.Valorefectivo) > 0 {
		miFilaComprobante++
		if listaBanco.Documento == "1" {
			debito = listaBanco.Valorefectivo
			credito = ""
		} else {
			debito = ""
			credito = listaBanco.Valorefectivo
		}
		miComprobanteDetalle = append(miComprobanteDetalle,
			comprobantedetalle{strconv.Itoa(miFilaComprobante),
				listaBanco.Cuentaefectivo,
				listaBanco.Tercero,
				miCentro,
				t.Nombre,
				"",
				debito,
				credito,
				listaBanco.Documento,
				listaBanco.Numero,
				fechaDate,
				fechaDate, "", ""})

		totalDebito += Flotante(debito)
		totalCredito += Flotante(credito)

	}
	log.Println("Banco 4")
	// transferencia
	if Flotante(listaBanco.Valortransferencia) > 0 {
		miFilaComprobante++
		if listaBanco.Documento == "1" {
			debito = listaBanco.Valortransferencia
			credito = ""
		} else {
			debito = ""
			credito = listaBanco.Valortransferencia
		}
		miComprobanteDetalle = append(miComprobanteDetalle,
			comprobantedetalle{strconv.Itoa(miFilaComprobante),
				listaBanco.Cuentatransferencia,
				listaBanco.Tercero,
				miCentro,
				t.Nombre,
				"",
				debito,
				credito,
				listaBanco.Documento,
				listaBanco.Numero,
				fechaDate,
				fechaDate, "", ""})
		totalDebito += Flotante(debito)
		totalCredito += Flotante(credito)
	}
	log.Println("Banco 5")
	// TarjetaDebito
	if Flotante(listaBanco.Valortarjetadebito) > 0 {
		miFilaComprobante++
		if listaBanco.Documento == "1" {
			debito = listaBanco.Valortarjetadebito
			credito = ""
		} else {
			debito = ""
			credito = listaBanco.Valortarjetadebito
		}
		miComprobanteDetalle = append(miComprobanteDetalle,
			comprobantedetalle{strconv.Itoa(miFilaComprobante),
				listaBanco.Cuentatarjetadebito,
				listaBanco.Tercero,
				miCentro,
				t.Nombre,
				"",
				debito,
				credito,
				listaBanco.Documento,
				listaBanco.Numero,
				fechaDate,
				fechaDate, "", ""})
		totalDebito += Flotante(debito)
		totalCredito += Flotante(credito)
	}
	log.Println("Banco 7")
	// TarjetaCredito
	if Flotante(listaBanco.Valortarjetacredito) > 0 {
		miFilaComprobante++
		if listaBanco.Documento == "1" {
			debito = listaBanco.Valortarjetacredito
			credito = ""
		} else {
			debito = ""
			credito = listaBanco.Valortarjetacredito
		}
		miComprobanteDetalle = append(miComprobanteDetalle,
			comprobantedetalle{strconv.Itoa(miFilaComprobante),
				listaBanco.Cuentatarjetacredito,
				listaBanco.Tercero,
				miCentro,
				t.Nombre,
				"",
				debito,
				credito,
				listaBanco.Documento,
				listaBanco.Numero,
				fechaDate,
				fechaDate, "", ""})
		totalDebito += Flotante(debito)
		totalCredito += Flotante(credito)
	}
	log.Println("Banco 8")
	// saldofavor
	log.Println("Banco 9999saldo favor")
	log.Println(listaBanco.Valorcliente)

	if Flotante(listaBanco.Valorcliente) > 0 {
		log.Println("Banco 9999")
		log.Println(listaBanco.Valorcliente)

		miFilaComprobante++
		if listaBanco.Documento == "1" {
			debito = listaBanco.Valorcliente
			credito = ""
		} else {
			debito = ""
			credito = listaBanco.Valorcliente
		}
		miComprobanteDetalle = append(miComprobanteDetalle,
			comprobantedetalle{strconv.Itoa(miFilaComprobante),
				listaBanco.Cuentacliente,
				listaBanco.Tercero,
				miCentro,
				t.Nombre,
				"",
				debito,
				credito,
				listaBanco.Documento,
				listaBanco.Numero,
				fechaDate,
				fechaDate, "", ""})
		totalDebito += Flotante(debito)
		totalCredito += Flotante(credito)

	}

	// valorajuste
	log.Println("Banco 9999saldo favor")
	log.Println(listaBanco.Valorajuste)

	if Flotante(listaBanco.Valorajuste) > 0 {
		log.Println("Banco 9999")
		log.Println(listaBanco.Valorajuste)

		miFilaComprobante++
		if listaBanco.Documento == "1" {
			debito = listaBanco.Valorajuste
			credito = ""
		} else {
			debito = ""
			credito = listaBanco.Valorajuste
		}
		miComprobanteDetalle = append(miComprobanteDetalle,
			comprobantedetalle{strconv.Itoa(miFilaComprobante),
				listaBanco.Cuentaajuste,
				listaBanco.Tercero,
				miCentro,
				t.Nombre,
				"",
				debito,
				credito,
				listaBanco.Documento,
				listaBanco.Numero,
				fechaDate,
				fechaDate, "", ""})
		totalDebito += Flotante(debito)
		totalCredito += Flotante(credito)
	}

	log.Println("Banco 9")
	for _, x := range listaBanco.Detalle {

		//var q string

		miFilaComprobante++

		if listaBanco.Documento == "1" {
			debito = ""
			credito = x.Abono
		} else {
			debito = x.Abono
			credito = ""

		}

		miComprobanteDetalle = append(miComprobanteDetalle,
			comprobantedetalle{strconv.Itoa(miFilaComprobante),
				x.Cuenta,
				listaBanco.Tercero,
				miCentro,
				t.Nombre,
				x.Factura,
				debito,
				credito,
				listaBanco.Documento,
				listaBanco.Numero,
				fechaDate,
				consignacionDate, "", ""})

		totalDebito += Flotante(debito)
		totalCredito += Flotante(credito)
		//	log.Println("Insertar Detalle \n"+ x.Tercero)
	}
	for _, miFila := range miComprobanteDetalle {
		if miFila.Debito == "" {
			miComprobanteDetalleCredito = append(miComprobanteDetalleCredito, miFila)
		} else {
			miComprobanteDetalleDebito = append(miComprobanteDetalleDebito, miFila)
		}

	}
	var filavan = 1
	for i, miFila := range miComprobanteDetalleDebito {
		filavan = i + 1
		miFila.Fila = strconv.Itoa(filavan)
		miComprobanteDetalleFinal = append(miComprobanteDetalleFinal, miFila)
	}
	filavan++
	for i, miFila := range miComprobanteDetalleCredito {
		miFila.Fila = strconv.Itoa(i + filavan)
		miComprobanteDetalleFinal = append(miComprobanteDetalleFinal, miFila)
	}
	log.Println("Banco 10")
	// crea comprobante
	var miperiodo = periodoSesion(r)

	ComprobanteAgregarGenerar(comprobante{listaBanco.Documento,
		listaBanco.Numero, fechaDate,
		consignacionDate,
		miperiodo,
		"",
		"",
		"",
		FormatoFlotante(totalDebito),
		FormatoFlotante(totalCredito),
		"Actualizar",
		miComprobanteDetalleFinal,
		nil})

	//var resultado bool
	//resultado = true
	var miDocumento = documento{}

	miDocumento = TraerDocumento(listaBanco.Documento)

	js, err := json.Marshal(MensajeBanco{"" + miDocumento.Nombre + " Numero  " + listaBanco.Numero + " Generado Correctamente "})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(js)

	//http.Redirect(w, r, "/CUENTADECOBROLista", 301)
}

// BANCO LISTA
func BancoLista(w http.ResponseWriter, r *http.Request) {

	parametros := map[string]interface{}{
		//"res":     listadokardex,
		"hosting":     ruta,
		"cuentabanco": ListaCuentaBanco(),
		"cuenta":      ListaCuentaAuxiliar(),
		"parametro":   TraerParametrosContabilidad(),
		"documento":   ListaDocumentoBanco(),
		"centro":      ListaCentro(),
		"miperiodo":   periodoSesion(r),
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/banco/bancoLista.html",
		"vista/banco/bancoScript.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/banco/autocompletaDocumento.html",
		"vista/autocompleta/autocompletaPlandecuentaempresa.html",
		"vista/banco/autocompletaTercero.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error comprobante nuevo 3")
	miTemplate.Execute(w, parametros)

	//tmp.Execute(w, varmap)
}

// INICIA CENTRO PDF

// TERMINA CENTRO PDF

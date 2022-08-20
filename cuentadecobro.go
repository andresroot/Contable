package main

// INICIA CUENTADECOBRO IMPORTAR PAQUETES
import (
	"bytes"
	"github.com/360EntSecGroup-Skylar/excelize"
	"math"
	"strings"

	"encoding/json"
	_ "encoding/json"
	"fmt"
	_ "github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
	_ "github.com/lib/pq"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// TERMINA CUENTADECOBRO IMPORTAR PAQUETES

// INICIA CUENTADECOBRO ESTRUCTURA JSON
type cuentadecobroJson struct {
	Id     string `json:"id"`
	Label  string `json:"label"`
	Value  string `json:"value"`
	Nombre string `json:"nombre"`
}

// TERMINA CUENTADECOBRO ESTRUCTURA JSON

// INICIA CUENTADECOBRO ESTRUCTURA
type cuentadecobroLista struct {
	Numero          string
	Fecha           time.Time
	Residente       string
	Residentenombre string
	Total           string
}

type cuentadecobroDato struct {
	Residente  string
	Nombre     string
	Descuento1 string
	Descuento2 string
	Cuotap     string
	Cuota1     string
	Cuota2     string
	Direccion  string
}

type cuentaResidente struct {
	Cuenta string
}

type cuentaSaldo struct {
	Debito  string
	Credito string
}

type cuentaSaldoFecha struct {
	Debito  string
	Credito string
	Fecha   string
}

// INICIA CUENTADECOBRO ESTRUCTURA
type cuentadecobro struct {
	Numero        string
	Fecha         time.Time
	Centro        string
	Residente     string
	Totalanterior string
	Totalactual   string
	Total         string
	Accion        string
	Detalle       []cuentadecobrodetalle       `json:"Detalle"`
	DetalleEditar []cuentadecobrodetalleeditar `json:"DetalleEditar"`
}

// TERMINA CUENTADECOBRO ESTRUCTURA

// INICIA CUENTADECOBRO DETALLE ESTRUCTURA
type cuentadecobrodetalle struct {
	Fila     string
	Numero   string
	Cuenta   string
	Anterior string
	Actual   string
	Total    string
}

type cuentadecobrodetalleGenerar struct {
	Numero          string
	Residente       string
	ResidenteNombre string
	Totalanterior   string
	Totalactual     string
	Total           string
}

// estructura para editar
type cuentadecobrodetalleeditar struct {
	Fila         string
	Numero       string
	Cuenta       string
	Cuentanombre string
	Anterior     string
	Actual       string
	Total        string
}

// TERMINA CUENTA DE COBRO DETALLE EDITAR

// INICIA CUENTA DE COBRO CONSULTA DETALLE
func CuentadecobroConsultaDetalle() string {
	var consulta = ""
	consulta = "select "
	consulta += "cuentadecobrodetalle.Fila as fila, "
	consulta += "cuentadecobrodetalle.Cuenta as cuenta, "
	consulta += "plandecuentaempresa.nombre as cuentanombre, "
	consulta += "cuentadecobrodetalle.Anterior as anterior, "
	consulta += "cuentadecobrodetalle.Actual as actual, "
	consulta += "cuentadecobrodetalle.Total as total "
	consulta += "from cuentadecobrodetalle "
	consulta += " inner join plandecuentaempresa on Plandecuentaempresa.codigo = Cuentadecobrodetalle.cuenta "
	consulta += " where cuentadecobrodetalle.numero=$1 "
	consulta += " order by fila "
	log.Println(consulta)
	return consulta
}
func saldoAnterior(residente string, fechaFinal time.Time, micuenta string) float64 {
	var consultasaldoAnterior = "SELECT fecha,debito,credito FROM comprobantedetalle  where"
	consultasaldoAnterior += " tercero=$1 and fecha<=$2 and cuenta=$3"

	log.Println("Fecha final anterior" + fechaFinal.Format("02/01/2006"))

	listadoSaldoResidente := []cuentaSaldoFecha{}
	db.Select(&listadoSaldoResidente, consultasaldoAnterior, residente, fechaFinal, micuenta)

	var totalSaldo float64
	totalSaldo = 0
	for _, miSaldo := range listadoSaldoResidente {
		//var c=k
		log.Println(miSaldo.Debito)
		log.Println(miSaldo.Credito)
		log.Println(miSaldo.Fecha)
		totalSaldo += Flotante(miSaldo.Debito) - Flotante(miSaldo.Credito)
	}

	log.Println("calculo cuenta anterior " + micuenta)
	log.Println("total anterior " + CadenaFlotante(totalSaldo))

	return totalSaldo

}

// cuenta de cobro todos
func CuentadecobroGenerarMes(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var periodoActual = periodoSesion(r)

	mes := mux.Vars(r)["mes"]
	miCentro := mux.Vars(r)["centro"]
	miPorcentaje := mux.Vars(r)["porcentaje"]

	log.Println(miPorcentaje)

	var Nitcontable string
	Nitcontable = "30"
	var miPorcentajeNumero float64
	miPorcentajeNumero, err := strconv.ParseFloat(miPorcentaje, 64)
	if err != nil {
		fmt.Println(err)
		return
	}
	parametroscontabilidad := configuracioncontabilidad{}
	parametroscontabilidad = TraerParametrosContabilidad()

	// borra datos anteriores
	listadoCuentaCobroBorrar := []cuentadecobro{}

	var consultaborra = "select * from cuentadecobro where EXTRACT(MONTH FROM  fecha)>=$1"
	db.Select(&listadoCuentaCobroBorrar, consultaborra, mes)

	for _, miCuentaBorra := range listadoCuentaCobroBorrar {

		var consultaborradetalle = "delete from cuentadecobrodetalle where numero=$1"
		db.Exec(consultaborradetalle, miCuentaBorra.Numero)

		consultaborradetalle = "delete from cuentadecobro where numero=$1"
		db.Exec(consultaborradetalle, miCuentaBorra.Numero)

		// BORRA MOVIMIENTOS
		var consultaborracomprobante = "delete from comprobante where documento=$2 and numero=$1"
		db.Exec(consultaborracomprobante, miCuentaBorra.Numero, Nitcontable)

		var consultaborracomprobantedetalle = "delete from comprobantedetalle where documento=$2 and  numero=$1"
		db.Exec(consultaborracomprobantedetalle, miCuentaBorra.Numero, Nitcontable)

	}

	var ultimo int

	if parametroscontabilidad.Phinicial == "0" {

	} else {
		db := dbConn()
		Numero := parametroscontabilidad.Phinicial
		var total int
		row := db.QueryRow("SELECT COUNT(*) FROM cuentadecobro  WHERE  Numero=$1", Numero)
		err := row.Scan(&total)
		if err != nil {
			log.Fatal(err)
		}
		//var resultado bool
		if total > 0 {
			// BUSCAR ULTIMO
			row := db.QueryRow("SELECT MAX(CAST(NUMERO AS INTEGER)) AS NUMERO FROM cuentadecobro")
			err := row.Scan(&ultimo)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			ultimo, _ = strconv.Atoi(parametroscontabilidad.Phinicial)
			ultimo--
		}
	}
	log.Println("Ultimo numero")

	var fechaString string
	fechaString = fechaInicial(periodoActual, mes)
	const (
		layoutISO = "2006-01-02"
	)

	fechaDate, _ := time.Parse(layoutISO, fechaString)

	log.Println(mes)
	log.Println(fechaString)
	cuentaInteresCodigo := plandecuentaempresa{}
	cuentaInteres := plandecuentaempresa{}

	db.Get(&cuentaInteresCodigo, "select distinct cuentaintereses as codigo from plandecuentaempresa where interes='SI'")
	db.Get(&cuentaInteres, "select * from plandecuentaempresa where codigo=$1", cuentaInteresCodigo.Codigo)

	var cuentasPropiedad = ListaEmpresa().Modulo
	var longitudPropiedad = strconv.Itoa(len(cuentasPropiedad))

	var consultacuentaAnterior = " SELECT distinct cuenta FROM comprobantedetalle "
	consultacuentaAnterior += " inner join plandecuentaempresa on "
	consultacuentaAnterior += " plandecuentaempresa.codigo=comprobantedetalle.cuenta "
	consultacuentaAnterior += " where "
	consultacuentaAnterior += " (plandecuentaempresa.cuota='NO' or plandecuentaempresa.cuota='' ) and "
	consultacuentaAnterior += " comprobantedetalle.tercero=$1 and "
	consultacuentaAnterior += " comprobantedetalle.fecha<=$2 and "
	consultacuentaAnterior += " substring(comprobantedetalle.cuenta,1," + longitudPropiedad + ")='" + cuentasPropiedad + "' and comprobantedetalle.cuenta<>'" + cuentaInteres.Codigo + "' group by cuenta"

	fmt.Println(consultacuentaAnterior)

	var consultasaldoAnterior = "SELECT debito,credito FROM comprobantedetalle  where"
	consultasaldoAnterior += " tercero=$1 and fecha<=$2 and cuenta=$3"

	//listadoGenerar := []cuentadecobrodetalleGenerar{}
	listadoCuentaResidente := []cuentaResidente{}
	//res := []residente{}
	listaResidente := []residente{}

	err = db.Select(&listaResidente, "SELECT * FROM residente where  not (cuotap='0' and cuota1='0' and cuota2='0') order by codigo")
	if err != nil {
		fmt.Println(err)
	}

	listadoSaldoResidente := []cuentaSaldo{}

	//FacturaInicial=""
	for i, miResidente := range listaResidente {
		//var miUltimo int
		//miUltimo=ultimo+1
		var numeroFactura string
		numeroFactura = strconv.Itoa(ultimo + i + 1)
		var miFila int
		miFila = 0
		var miFilaComprobante int
		miFilaComprobante = 0

		miCuentaResidentedetalle := []cuentadecobrodetalle{}
		miComprobanteDetalleDebito := []comprobantedetalle{}
		miComprobanteDetalleCredito := []comprobantedetalle{}
		miComprobanteDetalle := []comprobantedetalle{}

		var totalCuotaActual float64
		var totalCuotaAnterior float64
		var totalCuota float64
		var totalBaseInteres float64

		totalCuotaActual = 0
		totalCuotaAnterior = 0
		totalCuota = 0
		totalBaseInteres = 0
		var totalDebito float64
		var totalCredito float64

		totalDebito = 0
		totalCredito = 0

		// sumar saldos anteriores
		log.Println("Residente"+miResidente.Nit, miResidente.Nombre)
		db.Select(&listadoCuentaResidente, consultacuentaAnterior, miResidente.Nit, fechaString)
		for _, miCuentaSaldo := range listadoCuentaResidente {
			//var b=j
			db.Select(&listadoSaldoResidente, consultasaldoAnterior, miResidente.Nit, fechaString, miCuentaSaldo.Cuenta)
			log.Println("Cuenta anterior 111" + miCuentaSaldo.Cuenta)
			var totalSaldo float64
			totalSaldo = 0
			for _, miSaldo := range listadoSaldoResidente {
				//var c=k
				log.Println(miSaldo.Debito)
				log.Println(miSaldo.Credito)
				totalSaldo = Flotante(miSaldo.Debito) - Flotante(miSaldo.Credito)
			}
			if totalSaldo == 0 {
			} else {
				// nueva fila anterior
				miFila++
				miCuentaResidentedetalle = append(miCuentaResidentedetalle,
					cuentadecobrodetalle{strconv.Itoa(miFila),
						numeroFactura, miCuentaSaldo.Cuenta, FormatoFlotanteEntero(totalSaldo), "0", FormatoFlotanteEntero(totalSaldo)})
				totalCuotaAnterior += totalSaldo
				totalBaseInteres += totalSaldo
				totalCuota += totalSaldo
			}

			log.Println("Total Cuenta Anterior " + miCuentaSaldo.Cuenta)
			log.Println(FormatoFlotanteEntero(totalSaldo))
		}
		log.Println("Cuentap")

		// saldo actual CUENTAp
		CuentaP := plandecuentaempresa{}
		db.Get(&CuentaP, "SELECT * FROM plandecuentaempresa where cuota='SI' and tipo='CuotaP'")
		var totalSaldoAnterior float64
		totalSaldoAnterior = saldoAnterior(miResidente.Nit, fechaDate, CuentaP.Codigo)
		if (miResidente.CuotaP == "" || miResidente.CuotaP == "0") && totalSaldoAnterior == 0 {
		} else {
			miFila++
			var cuotapActual float64
			cuotapActual = Flotante(miResidente.CuotaP)
			miCuentaResidentedetalle = append(miCuentaResidentedetalle,
				cuentadecobrodetalle{strconv.Itoa(miFila),
					numeroFactura, CuentaP.Codigo,
					FormatoFlotanteEntero(totalSaldoAnterior),
					FormatoFlotanteEntero(Flotante(miResidente.CuotaP)),
					FormatoFlotanteEntero(totalSaldoAnterior + cuotapActual)})

			// sumatorias Interes
			if CuentaP.Interes == "SI" {
				totalBaseInteres += totalSaldoAnterior
			}

			// inserta fila cuenta
			miFilaComprobante++
			miComprobanteDetalleDebito = append(miComprobanteDetalleDebito,
				comprobantedetalle{strconv.Itoa(miFilaComprobante),
					CuentaP.Codigo,
					miResidente.Nit,
					miCentro,
					strings.TrimSpace(CuentaP.Nombre) + " " + mesLetras(mes),
					"",
					FormatoFlotante(cuotapActual),
					"",
					Nitcontable,
					numeroFactura,
					fechaDate,
					fechaDate, "", ""})
			// Inserta Fila contra
			miFilaComprobante++
			miComprobanteDetalleCredito = append(miComprobanteDetalleCredito,
				comprobantedetalle{strconv.Itoa(miFilaComprobante),
					CuentaP.Contra,
					miResidente.Nit,
					miCentro,
					strings.TrimSpace(CuentaP.Nombre) + " " + mesLetras(mes),
					"",
					"",
					FormatoFlotante(cuotapActual),
					Nitcontable,
					numeroFactura,
					fechaDate,
					fechaDate, "", ""})

			totalDebito += cuotapActual
			totalCredito += cuotapActual

			totalCuotaAnterior += totalSaldoAnterior
			totalCuotaActual += cuotapActual
			totalCuota += totalSaldoAnterior + cuotapActual
		}

		// cuenta 1
		Cuenta1 := plandecuentaempresa{}
		db.Get(&Cuenta1, "SELECT * FROM plandecuentaempresa where cuota='SI' and tipo='Cuota1'")
		totalSaldoAnterior = saldoAnterior(miResidente.Nit, fechaDate, Cuenta1.Codigo)
		if (miResidente.Cuota1 == "" || miResidente.Cuota1 == "0") && totalSaldoAnterior == 0 {
		} else {
			miFila++
			var cuota1Actual float64
			cuota1Actual = Flotante(miResidente.Cuota1)
			miCuentaResidentedetalle = append(miCuentaResidentedetalle,
				cuentadecobrodetalle{strconv.Itoa(miFila),
					numeroFactura, Cuenta1.Codigo, FormatoFlotanteEntero(totalSaldoAnterior),
					FormatoFlotanteEntero(Flotante(miResidente.Cuota1)),
					FormatoFlotanteEntero(totalSaldoAnterior + cuota1Actual)})

			// sumatorias Interes
			if Cuenta1.Interes == "SI" {
				totalBaseInteres += totalSaldoAnterior
			}
			// inserta fila cuenta1
			miFilaComprobante++
			miComprobanteDetalleDebito = append(miComprobanteDetalleDebito,
				comprobantedetalle{strconv.Itoa(miFilaComprobante),
					Cuenta1.Codigo,
					miResidente.Nit,
					miCentro,
					strings.TrimSpace(Cuenta1.Nombre) + " " + mesLetras(mes),
					"",
					FormatoFlotante(cuota1Actual),
					"",
					Nitcontable,
					numeroFactura,
					fechaDate,
					fechaDate, "", ""})
			// Inserta Fila contra
			miFilaComprobante++
			miComprobanteDetalleCredito = append(miComprobanteDetalleCredito,
				comprobantedetalle{strconv.Itoa(miFilaComprobante),
					Cuenta1.Contra,
					miResidente.Nit,
					miCentro,
					strings.TrimSpace(Cuenta1.Nombre) + " " + mesLetras(mes),
					"",
					"",
					FormatoFlotante(cuota1Actual),
					Nitcontable,
					numeroFactura,
					fechaDate,
					fechaDate, "", ""})

			totalDebito += cuota1Actual
			totalCredito += cuota1Actual

			// sumatorias
			totalCuotaAnterior += totalSaldoAnterior
			totalCuotaActual += cuota1Actual
			totalCuota += totalSaldoAnterior + cuota1Actual
		}

		//cuenta 2
		Cuenta2 := plandecuentaempresa{}
		db.Get(&Cuenta2, "SELECT * FROM plandecuentaempresa where cuota='SI' and tipo='Cuota2'")
		totalSaldoAnterior = saldoAnterior(miResidente.Nit, fechaDate, Cuenta2.Codigo)
		if (miResidente.Cuota2 == "" || miResidente.Cuota2 == "0") && totalSaldoAnterior == 0 {
		} else {
			miFila++
			var cuota2Actual float64
			cuota2Actual = Flotante(miResidente.Cuota2)
			miCuentaResidentedetalle = append(miCuentaResidentedetalle,
				cuentadecobrodetalle{strconv.Itoa(miFila),
					numeroFactura, Cuenta2.Codigo, FormatoFlotanteEntero(totalSaldoAnterior), FormatoFlotanteEntero(Flotante(miResidente.Cuota2)), FormatoFlotanteEntero(totalSaldoAnterior + cuota2Actual)})

			// sumatorias Interes
			if Cuenta2.Interes == "SI" {
				totalBaseInteres += totalSaldoAnterior
			}

			// inserta fila cuenta1
			miFilaComprobante++
			miComprobanteDetalleDebito = append(miComprobanteDetalleDebito,
				comprobantedetalle{strconv.Itoa(miFilaComprobante),
					Cuenta2.Codigo,
					miResidente.Nit,
					miCentro,
					strings.TrimSpace(Cuenta2.Nombre) + " " + mesLetras(mes),
					"",
					FormatoFlotante(cuota2Actual),
					"",
					Nitcontable,
					numeroFactura,
					fechaDate,
					fechaDate, "", ""})
			// Inserta Fila contra
			miFilaComprobante++
			miComprobanteDetalleCredito = append(miComprobanteDetalleCredito,
				comprobantedetalle{strconv.Itoa(miFilaComprobante),
					Cuenta2.Contra,
					miResidente.Nit,
					miCentro,
					strings.TrimSpace(Cuenta2.Nombre) + " " + mesLetras(mes),
					"",
					"",
					FormatoFlotante(cuota2Actual),
					Nitcontable,
					numeroFactura,
					fechaDate,
					fechaDate, "", ""})

			totalDebito += cuota2Actual
			totalCredito += cuota2Actual
			// sumatorias
			totalCuotaAnterior += totalSaldoAnterior
			totalCuotaActual += cuota2Actual
			totalCuota += totalSaldoAnterior + cuota2Actual
		}
		// cuenta 3

		//Cuenta3 := plandecuentaempresa{}
		//db.Get(&Cuenta3, "SELECT * FROM plandecuentaempresa where cuota='SI' and tipo='Cuota3'")
		//totalSaldoAnterior=saldoAnterior(miResidente.Nit,fechaDate,Cuenta3.Codigo)
		//if (miResidente.Cuota3=="" ||miResidente.Cuota3=="0") && totalSaldoAnterior==0{
		//} else {
		//	miFila++
		//	var cuota3Actual float64
		//	cuota3Actual=Flotante(miResidente.Cuota3)
		//	miCuentaResidentedetalle=append(miCuentaResidentedetalle,
		//		cuentadecobrodetalle{strconv.Itoa(miFila),
		//			numeroFactura,Cuenta3.Codigo,FormatoFlotanteEntero(totalSaldoAnterior),FormatoFlotanteEntero(Flotante(miResidente.Cuota3)),FormatoFlotanteEntero(totalSaldoAnterior+cuota3Actual)})
		//	// sumatorias
		//
		//	// sumatorias Interes
		//	if Cuenta3.Interes=="SI"{
		//		totalBaseInteres+=totalSaldoAnterior
		//	}
		//
		//
		//	// inserta fila cuenta1
		//	miFilaComprobante++;
		//	miComprobanteDetalleDebito=append(miComprobanteDetalleDebito,
		//		comprobantedetalle{strconv.Itoa(miFilaComprobante),
		//			Cuenta3.Codigo,
		//			miResidente.Nit,
		//			miCentro,
		//			strings.TrimSpace(Cuenta3.Nombre)+" "+mesLetras(mes),
		//			"",
		//			FormatoFlotante(cuota3Actual)	,
		//			"",
		//			Nitcontable,
		//			numeroFactura,
		//			fechaDate,
		//			fechaDate})
		//	// Inserta Fila contra
		//	miFilaComprobante++;
		//	miComprobanteDetalleCredito=append(miComprobanteDetalleCredito,
		//		comprobantedetalle{strconv.Itoa(miFilaComprobante),
		//			Cuenta3.Contra,
		//			miResidente.Nit,
		//			miCentro,
		//			strings.TrimSpace(Cuenta3.Nombre)+" "+mesLetras(mes),
		//			"",
		//			"",
		//			FormatoFlotante(cuota3Actual),
		//			Nitcontable,
		//			numeroFactura,
		//			fechaDate,
		//			fechaDate})
		//
		//	totalDebito+=cuota3Actual;
		//	totalCredito+=cuota3Actual;
		//
		//	totalCuotaAnterior+=totalSaldoAnterior
		//	totalCuotaActual+=cuota3Actual
		//	totalCuota+=totalSaldoAnterior+cuota3Actual
		//}

		// cuenta interes

		totalSaldoAnterior = 0
		//saldoAnterior(miResidente.Nit,fechaFinal,cuentaInteres.Codigo)

		totalSaldoAnterior = 0
		totalSaldoAnterior = saldoAnterior(miResidente.Nit, fechaDate, cuentaInteres.Codigo)

		log.Println("saldo interes anterior ")
		log.Println(FormatoFlotante(totalSaldoAnterior))

		log.Println("base interes ")
		log.Println(FormatoFlotante(totalBaseInteres))

		if totalBaseInteres == 0 && totalSaldoAnterior == 0 {
		} else {
			miFila++
			var cuotaInteresActual float64
			cuotaInteresActual = (totalBaseInteres) * (miPorcentajeNumero / 100)

			miCuentaResidentedetalle = append(miCuentaResidentedetalle,
				cuentadecobrodetalle{strconv.Itoa(miFila),
					numeroFactura, cuentaInteres.Codigo, FormatoFlotanteEntero(totalSaldoAnterior), FormatoFlotanteEntero(cuotaInteresActual), FormatoFlotanteEntero(totalSaldoAnterior + cuotaInteresActual)})
			// sumatorias
			if cuotaInteresActual == 0 {

			} else {

				// inserta fila cuenta1
				miFilaComprobante++
				miComprobanteDetalleDebito = append(miComprobanteDetalleDebito,
					comprobantedetalle{strconv.Itoa(miFilaComprobante),
						cuentaInteres.Codigo,
						miResidente.Nit,
						miCentro,
						strings.TrimSpace(cuentaInteres.Nombre) + " " + mesLetras(mes),
						"",
						FormatoFlotante(cuotaInteresActual),
						"",
						Nitcontable,
						numeroFactura,
						fechaDate,
						fechaDate, "", ""})
				// Inserta Fila contra
				miFilaComprobante++
				miComprobanteDetalleCredito = append(miComprobanteDetalleCredito,
					comprobantedetalle{strconv.Itoa(miFilaComprobante),
						cuentaInteres.Contra,
						miResidente.Nit,
						miCentro,
						strings.TrimSpace(cuentaInteres.Nombre) + " " + mesLetras(mes),
						"",
						"",
						FormatoFlotante(cuotaInteresActual),
						Nitcontable,
						numeroFactura,
						fechaDate,
						fechaDate, "", ""})
				totalDebito += cuotaInteresActual
				totalCredito += cuotaInteresActual
			}

			totalCuotaAnterior += totalSaldoAnterior
			totalCuotaActual += cuotaInteresActual
			totalCuota += totalSaldoAnterior + cuotaInteresActual
		}

		// genera ceunta de cobro
		CuentadecobroNuevaGenerar(cuentadecobro{numeroFactura,
			fechaDate, miCentro,
			miResidente.Codigo,
			FormatoFlotanteEntero(totalCuotaAnterior),
			FormatoFlotanteEntero(totalCuotaActual),
			FormatoFlotanteEntero(totalCuota),
			"Nueva",
			miCuentaResidentedetalle,
			nil})

		// agrega lineas debito
		var filavan = 1
		for i, midetalle := range miComprobanteDetalleDebito {
			filavan = i + 1
			midetalle.Fila = strconv.Itoa(filavan)
			miComprobanteDetalle = append(miComprobanteDetalle, midetalle)
			log.Println("fila" + midetalle.Fila)
			log.Println(midetalle.Cuenta)
			log.Println(midetalle.Debito)
		}

		filavan++
		for i, midetalle := range miComprobanteDetalleCredito {
			midetalle.Fila = strconv.Itoa(i + filavan)
			miComprobanteDetalle = append(miComprobanteDetalle, midetalle)
			log.Println("fila" + midetalle.Fila)
			log.Println(midetalle.Cuenta)
			log.Println(midetalle.Debito)
		}

		// crea comprobante
		ComprobanteAgregarGenerar(comprobante{Nitcontable,
			numeroFactura, fechaDate,
			fechaDate,
			periodoActual,
			"",
			"",
			"",
			FormatoFlotante(totalDebito),
			FormatoFlotante(totalCredito),
			"Actualizar",
			miComprobanteDetalle, nil})

		// fin residente
	}

	var consulta string

	consulta = "select Numero,Residente,residente.nombre as ResidenteNombre,"
	consulta += " Totalanterior,Totalactual,Total from cuentadecobro "
	consulta += " inner join residente on residente.codigo=cuentadecobro.residente"
	consulta += " where EXTRACT(MONTH FROM  cuentadecobro.fecha)=$1"
	log.Println(consulta)
	log.Println(mes)
	listacuentadecobro := []cuentadecobrodetalleGenerar{}

	err1 := db.Select(&listacuentadecobro, consulta, mes)
	if err != nil {
		fmt.Println(err1)
		return
	}

	//if simueve == false {
	//	var slice []string
	//	slice = make([]string, 0)
	//	data, _ := json.Marshal(slice)
	//	w.WriteHeader(200)
	//	w.Header().Set("Content-Type", "application/json")
	//	w.Write(data)
	//} else {
	data, _ := json.Marshal(listacuentadecobro)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	//}
}

// TERMINA CUENTADECOBRODETALLE ESTRUCTURA
func CuentadecobroGenerar(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/autocompleta/autocompletacentro.html",
		"vista/cuentadecobro/cuentadecobroGenerar.html")

	panel := mux.Vars(r)["panel"]
	varmap := map[string]interface{}{
		"hosting": ruta,
		"centro":  ListaCentro(),
		"panel":   panel,
	}
	tmp.Execute(w, varmap)
}

// INICIA CUENTADECOBRO LISTA
func CuentadecobroLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/cuentadecobro/cuentadecobroLista.html")
	log.Println("Error cuentadecobro 0")
	var consulta string
	panel := mux.Vars(r)["panel"]

	consulta = "  SELECT cuentadecobro.numero,fecha,residente.codigo as residente,residente.nombre as Residentenombre,cuentadecobro.total "
	consulta += " FROM cuentadecobro "
	consulta += " inner join residente on residente.codigo=cuentadecobro.residente "
	consulta += " ORDER BY cuentadecobro.numero asc,cuentadecobro.residente asc"

	fmt.Println(consulta)

	db := dbConn()
	res := []cuentadecobroLista{}
	//db.Select(&res, consulta)

	//error1 = db.Select(&res, consulta)
	err := db.Select(&res, consulta)

	if err != nil {
		fmt.Println(err)
		return
	}

	varmap := map[string]interface{}{
		"res":     res,
		"hosting": ruta,
		"panel":   panel,
	}
	log.Println("Error cuentadecobro888")
	tmp.Execute(w, varmap)
}

// TERMINA CUENTADECOBRO LISTA

// INICIA CUENTADECOBRO NUEVO
func CuentadecobroNuevo(w http.ResponseWriter, r *http.Request) {
	log.Println("Error cuentadecobro nuevo 1")
	log.Println("Error cuentadecobro nuevo 2")
	parametros := map[string]interface{}{
		"hosting": ruta,
		"centro":  ListaCentro(),
	}

	t, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/cuentadecobro/cuentadecobroNuevo.html",
		"vista/cuentadecobro/cuentadecobroScript.html",
		"vista/cuentadecobro/autocompletaplandecuentaempresa.html",
		"vista/cuentadecobro/autocompletaresidente.html")
	fmt.Printf("%v, %v", t, err)
	log.Println("Error cuentadecobro nuevo 3")
	t.Execute(w, parametros)
}

// proceso que crea la cuenta de cobro desde objeto
func CuentadecobroNuevaGenerar(tempCuentadecobro cuentadecobro) {
	//	var periodoActual="2022"

	db := dbConn()
	//var tempCuentadecobro cuentadecobro
	//
	//b, err := ioutil.ReadAll(r.Body)
	//defer r.Body.Close()
	//if err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}
	//
	//// carga informacion de la CUENTADECOBRO
	//err = json.Unmarshal(b, &tempCuentadecobro)
	//if err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}

	if tempCuentadecobro.Accion == "Actualizar" {
		// borra detalle anterior
		delForm, err := db.Prepare("DELETE from cuentadecobrodetalle WHERE numero=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(tempCuentadecobro.Numero)

		// borra cabecera anterior

		delForm1, err := db.Prepare("DELETE from cuentadecobro WHERE numero=$1 ")
		if err != nil {
			panic(err.Error())
		}
		delForm1.Exec(tempCuentadecobro.Numero)
	}

	// INSERTA DETALLE
	for i, x := range tempCuentadecobro.Detalle {
		var a = i
		var q string

		q = "insert into cuentadecobrodetalle ("
		q += "Fila,"
		q += "Numero,"
		q += "Cuenta,"
		q += "Anterior,"
		q += "Actual,"
		q += "Total"
		q += " ) values("
		q += parametros(6)
		q += ")"

		log.Println("Cadena SQL " + q)
		insForm, err := db.Prepare(q)
		if err != nil {
			panic(err.Error())
		}

		// TERMINA CUENTADECOBRO GRABAR INSERTAR
		_, err = insForm.Exec(
			x.Fila,
			x.Numero,
			x.Cuenta,
			Quitacoma(x.Anterior),
			Quitacoma(x.Actual),
			Quitacoma(x.Total))
		if err != nil {
			panic(err)
		}

		// crea detalle

		log.Println("Insertar Detalle \n", x.Cuenta, a)
	}

	log.Println("Got %s age %s club %s\n", tempCuentadecobro.Numero)
	var q string
	q += "insert into cuentadecobro ("
	q += "Numero,"
	q += "Fecha,"
	q += "Centro,"
	q += "Residente,"
	q += "Totalanterior,"
	q += "Totalactual,"
	q += "Total"
	q += " ) values("
	q += parametros(7)
	q += " ) "

	log.Println("Cadena SQL " + q)
	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}
	layout := "2006-01-02"

	log.Println("Hora", tempCuentadecobro.Fecha.Format("02/01/2006"))

	// TERMINA CUENTADECOBRO GRABAR INSERTAR
	_, err = insForm.Exec(
		tempCuentadecobro.Numero,
		tempCuentadecobro.Fecha.Format(layout),
		tempCuentadecobro.Centro,
		tempCuentadecobro.Residente,
		Quitacoma(tempCuentadecobro.Totalanterior),
		Quitacoma(tempCuentadecobro.Totalactual),
		Quitacoma(tempCuentadecobro.Total))
	if err != nil {
		panic(err)
	}

	//http.Redirect(w, r, "/CUENTADECOBROLista", 301)
}

// INICIA CUENTADECOBRO INSERTAR AJAX
func CuentadecobroAgregar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	var tempCuentadecobro cuentadecobro

	var Nitcontable string
	Nitcontable = "30"

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// carga informacion de la CUENTADECOBRO
	err = json.Unmarshal(b, &tempCuentadecobro)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	var periodoActual = strconv.Itoa(tempCuentadecobro.Fecha.Year())

	if tempCuentadecobro.Accion == "Actualizar" {
		// borra detalle anterior
		delForm, err := db.Prepare("DELETE from cuentadecobrodetalle WHERE numero=$1")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec(tempCuentadecobro.Numero)

		// borra cabecera anterior

		delForm1, err := db.Prepare("DELETE from cuentadecobro WHERE numero=$1 ")
		if err != nil {
			panic(err.Error())
		}
		delForm1.Exec(tempCuentadecobro.Numero)

	}
	var miFilaComprobante = 0
	miComprobanteDetalleDebito := []comprobantedetalle{}
	miComprobanteDetalleCredito := []comprobantedetalle{}
	miComprobanteDetalle := []comprobantedetalle{}
	// INSERTA DETALLE
	var totalDebito float64
	var totalCredito float64
	totalDebito = 0
	totalCredito = 0

	for i, x := range tempCuentadecobro.Detalle {
		var a = i
		var q string

		q = "insert into cuentadecobrodetalle ("
		q += "Fila,"
		q += "Numero,"
		q += "Cuenta,"
		q += "Anterior,"
		q += "Actual,"
		q += "Total"
		q += " ) values("
		q += parametros(6)
		q += ")"

		log.Println("Cadena SQL " + q)
		insForm, err := db.Prepare(q)
		if err != nil {
			panic(err.Error())
		}

		// TERMINA CUENTADECOBRO GRABAR INSERTAR
		_, err = insForm.Exec(
			x.Fila,
			x.Numero,
			x.Cuenta,
			Quitacoma(x.Anterior),
			Quitacoma(x.Actual),
			Quitacoma(x.Total))
		if err != nil {
			panic(err)
		}

		CuentaP := plandecuentaempresa{}
		db.Get(&CuentaP, "SELECT * FROM plandecuentaempresa where codigo=$1", x.Cuenta)

		var cuotaActual = Flotante(x.Actual)

		log.Println("cuenta debito " + x.Cuenta)
		miFilaComprobante++
		miComprobanteDetalleDebito = append(miComprobanteDetalleDebito,
			comprobantedetalle{strconv.Itoa(miFilaComprobante),
				CuentaP.Codigo,
				tempCuentadecobro.Residente,
				tempCuentadecobro.Centro,
				strings.TrimSpace(CuentaP.Nombre) + " " + mesLetras(strconv.Itoa(int(tempCuentadecobro.Fecha.Month()))),
				"",
				FormatoFlotante(cuotaActual),
				"",
				Nitcontable,
				tempCuentadecobro.Numero,
				tempCuentadecobro.Fecha,
				tempCuentadecobro.Fecha, "", ""})
		// Inserta Fila contra

		miFilaComprobante++
		miComprobanteDetalleCredito = append(miComprobanteDetalleCredito,
			comprobantedetalle{strconv.Itoa(miFilaComprobante),
				CuentaP.Contra,
				tempCuentadecobro.Residente,
				tempCuentadecobro.Centro,
				strings.TrimSpace(CuentaP.Nombre) + " " + mesLetras(strconv.Itoa(int(tempCuentadecobro.Fecha.Month()))),
				"",
				"",
				FormatoFlotante(cuotaActual),
				Nitcontable,
				tempCuentadecobro.Numero,
				tempCuentadecobro.Fecha,
				tempCuentadecobro.Fecha, "", ""})

		totalDebito += cuotaActual
		totalCredito += cuotaActual

		log.Println("Insertar Detalle \n", x.Cuenta, a)
	}

	// agrega lineas debito
	var filavan = 1
	for i, midetalle := range miComprobanteDetalleDebito {
		filavan = i + 1
		midetalle.Fila = strconv.Itoa(filavan)
		miComprobanteDetalle = append(miComprobanteDetalle, midetalle)
		log.Println("fila" + midetalle.Fila)
		log.Println(midetalle.Cuenta)
		log.Println(midetalle.Debito)
	}

	filavan++
	for i, midetalle := range miComprobanteDetalleCredito {
		midetalle.Fila = strconv.Itoa(i + filavan)
		miComprobanteDetalle = append(miComprobanteDetalle, midetalle)
		log.Println("fila" + midetalle.Fila)
		log.Println(midetalle.Cuenta)
		log.Println(midetalle.Debito)
	}

	// crea comprobante
	ComprobanteAgregarGenerar(comprobante{Nitcontable,
		tempCuentadecobro.Numero, tempCuentadecobro.Fecha,
		tempCuentadecobro.Fecha,
		periodoActual,
		"",
		"",
		"",
		FormatoFlotante(totalDebito),
		FormatoFlotante(totalCredito),
		"Actualizar",
		miComprobanteDetalle,
		nil})

	log.Println("Got %s age %s club %s\n", tempCuentadecobro.Numero)
	var q string
	q += "insert into cuentadecobro ("
	q += "Numero,"
	q += "Fecha,"
	q += "Centro,"
	q += "Residente,"
	q += "Totalanterior,"
	q += "Totalactual,"
	q += "Total"
	q += " ) values("
	q += parametros(7)
	q += " ) "

	log.Println("Cadena SQL " + q)
	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}
	layout := "2006-01-02"

	log.Println("Hora", tempCuentadecobro.Fecha.Format("02/01/2006"))

	// TERMINA CUENTADECOBRO GRABAR INSERTAR
	_, err = insForm.Exec(
		tempCuentadecobro.Numero,
		tempCuentadecobro.Fecha.Format(layout),
		tempCuentadecobro.Centro,
		tempCuentadecobro.Residente,
		Quitacoma(tempCuentadecobro.Totalanterior),
		Quitacoma(tempCuentadecobro.Totalactual),
		Quitacoma(tempCuentadecobro.Total))
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

func CuentadecobroDatoAgregar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	//var tempCuentadecobro cuentadecobro
	listacuentadecobroDato := []cuentadecobroDato{}
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

		q = "update residente set "
		q += "descuento1 = $2 ,"
		q += "descuento2 = $3, "
		q += "cuotap = $4, "
		q += "cuota1 = $5, "
		q += "cuota2 = $6, "
		q += "direccion = $7 "
		q += "where codigo = $1 "

		log.Println("Cadena SQL " + q)
		insForm, err := db.Prepare(q)
		if err != nil {
			panic(err.Error())
		}

		// TERMINA CUENTADECOBRO GRABAR INSERTAR
		_, err = insForm.Exec(
			x.Residente,
			Quitacoma(x.Descuento1),
			Quitacoma(x.Descuento2),
			Quitacoma(x.Cuotap),
			Quitacoma(x.Cuota1),
			Quitacoma(x.Cuota2),
			x.Direccion)

		if err != nil {
			panic(err)
		}

		log.Println("Insertar Detalle \n" + x.Residente)
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

// INICIA CUENTADECOBRO EXISTE
func CuentadecobroExiste(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Numero := mux.Vars(r)["numero"]
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM cuentadecobro  WHERE  Numero=$1", Numero)
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

// INICIA CUENTADECOBRO EDITAR
func CuentadecobroDato(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta = ""
	consulta = "select "
	consulta += "residente.Codigo as Residente, "
	consulta += "residente.Nombre as Nombre, "
	consulta += "residente.Descuento1 as Descuento1, "
	consulta += "residente.Descuento2 as Descuento2, "
	consulta += "residente.CuotaP as CuotaP, "
	consulta += "residente.Cuota1 as Cuota1, "
	consulta += "residente.Cuota2 as Cuota2, "
	consulta += "residente.Direccion as Direccion "
	consulta += "from residente "
	consulta += " order by cast(codigo as integer) "

	log.Println("Cadena SQL " + consulta)

	// traer detalle
	det := []cuentadecobroDato{}
	err2 := db.Select(&det, consulta)
	if err2 != nil {
		fmt.Println(err2)
	}
	panel := mux.Vars(r)["panel"]

	//	log.Println("detalle cuentadecobro)
	parametros := map[string]interface{}{
		"cuentadecobroDato": det,
		"hosting":           ruta,
		"panel":             panel,
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/cuentadecobro/cuentadecobroDato.html",
		"vista/cuentadecobro/cuentadecobroDatoScript.html",
	)

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error cuentadecobro nuevo 3")
	miTemplate.Execute(w, parametros)

	//tmp.Execute(w, parametros)
}

// INICIA CUENTADECOBRO EDITAR
func CuentadecobroEditar(w http.ResponseWriter, r *http.Request) {

	Numero := mux.Vars(r)["numero"]
	//log.Println("inicio cuentadecobro editar" + Nit)
	db := dbConn()

	// traer cuentadecobro
	v := cuentadecobro{}
	err := db.Get(&v, "SELECT * FROM cuentadecobro WHERE  Numero=$1", Numero)
	if err != nil {
		log.Fatalln(err)
	}

	// traer detalle
	det := []cuentadecobrodetalleeditar{}
	err2 := db.Select(&det, CuentadecobroConsultaDetalle(), Numero)
	if err2 != nil {
		fmt.Println(err2)
	}
	// traer residente
	t := residente{}
	err1 := db.Get(&t, "SELECT * FROM residente where codigo=$1", v.Residente)
	if err1 != nil {
		log.Fatalln(err1)
	}
	panel := mux.Vars(r)["panel"]

	//	log.Println("detalle cuentadecobro)
	parametros := map[string]interface{}{
		"panel":         panel,
		"cuentadecobro": v,
		"detalle":       det,
		"hosting":       ruta,
		"tercero":       t,
		"centro":        ListaCentro(),
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/cuentadecobro/cuentadecobroEditar.html",
		"vista/cuentadecobro/cuentadecobroScript.html",
		"vista/cuentadecobro/autocompletaplandecuentaempresa.html",
		"vista/cuentadecobro/autocompletatercero.html")

	fmt.Printf("%v, %v", miTemplate, err)
	log.Println("Error cuentadecobro nuevo 3")
	miTemplate.Execute(w, parametros)

	//tmp.Execute(w, parametros)
}

// INICIA CUENTADECOBRO BORRAR
func CuentadecobroBorrar(w http.ResponseWriter, r *http.Request) {
	Numero := mux.Vars(r)["numero"]
	log.Println("inicio cuentadecobro editar" + Numero)
	db := dbConn()

	// traer CUENTADECOBRO
	v := cuentadecobro{}
	err := db.Get(&v, "SELECT * FROM cuentadecobro WHERE Numero=$1", Numero)
	if err != nil {
		log.Fatalln(err)
	}

	// traer detalle
	det := []cuentadecobrodetalleeditar{}
	err2 := db.Select(&det, CuentadecobroConsultaDetalle(), Numero)
	if err2 != nil {
		fmt.Println(err2)
	}

	// traer residente
	t := residente{}
	err1 := db.Get(&t, "SELECT * FROM residente where codigo=$1", v.Residente)
	if err1 != nil {
		log.Fatalln(err1)
	}
	//	log.Println("detalle producto" + det.Producto+det.ProductoNombre)
	panel := mux.Vars(r)["panel"]

	//	log.Println("detalle cuentadecobro)
	parametros := map[string]interface{}{
		"panel":         panel,
		"cuentadecobro": v,
		"detalle":       det,
		"hosting":       ruta,
		"tercero":       t,
		"centro":        ListaCentro(),
	}

	miTemplate, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/cuentadecobro/cuentadecobroBorrar.html",
		"vista/cuentadecobro/cuentadecobroScript.html",
		"vista/cuentadecobro/autocompletaplandecuentaempresa.html",
		"vista/cuentadecobro/autocompletatercero.html",
		"vista/cuentadecobro/autocompletacentro.html")

	log.Println("Error cuentadecobro nuevo 3")
	miTemplate.Execute(w, parametros)

	//tmp.Execute(w, parametros)
}

// INICIA CUENTADECOBRO ELIMINAR
func CuentadecobroEliminar(w http.ResponseWriter, r *http.Request) {
	log.Println("Inicio Eliminar")
	db := dbConn()
	Numero := mux.Vars(r)["numero"]
	panel := mux.Vars(r)["panel"]

	// borrar CUENTADECOBRO
	delForm, err := db.Prepare("DELETE from cuentadecobro WHERE  Numero=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(Numero)

	// borar detalle
	delForm1, err := db.Prepare("DELETE from cuentadecobrodetalle WHERE  Numero=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm1.Exec(Numero)

	log.Println("Registro Eliminado" + Numero)
	http.Redirect(w, r, "/CuentadecobroLista/"+panel, 301)
}

// TERMINA CUENTADECOBRO ELIMINAR

// INICIA CUENTA DE COBRO PDF
func CuentadecobroPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Numero := mux.Vars(r)["numero"]
	// TRAER CUENTADECOBRO
	miCuentadecobro := cuentadecobro{}
	err := db.Get(&miCuentadecobro, "SELECT * FROM cuentadecobro where numero=$1", Numero)
	if err != nil {
		log.Fatalln(err)
	}

	// TRAER DETALLE
	miDetalle := []cuentadecobrodetalleeditar{}
	err2 := db.Select(&miDetalle, CuentadecobroConsultaDetalle(), Numero)
	if err2 != nil {
		fmt.Println(err2)
	}

	//var e empresa = ListaEmpresa()

	var buf bytes.Buffer
	var err1 error

	pdf := gofpdf.New("P", "mm", "Letter", cnFontDir)

	CuentadecobroHeader(pdf, miCuentadecobro)
	CuentadecobroFooter(pdf)

	// inicia suma de paginas
	pdf.AliasNbPages("")
	pdf.AddPage()
	pdf.Ln(1)
	CuentadecobroCabecera(pdf, miCuentadecobro, miDetalle)

	var filas = len(miDetalle)
	// menos de 32
	if filas <= 32 {
		for i, miFila := range miDetalle {
			var a = i + 1
			CuentadecobroFilaDetalle(pdf, miFila, a)
		}
		CuentadecobroPieDePagina(pdf, miCuentadecobro)
	} else {
		// mas de 32 y menos de 73
		if filas > 32 && filas <= 73 {
			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 41 {
					CuentadecobroFilaDetalle(pdf, miFila, a)
				}
			}
			// segunda pagina
			pdf.AddPage()
			CuentadecobroCabecera(pdf, miCuentadecobro, miDetalle)
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 41 {
					CuentadecobroFilaDetalle(pdf, miFila, a)
				}
			}

			CuentadecobroPieDePagina(pdf, miCuentadecobro)
		} else {
			// mas de 80

			// primera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a <= 41 {
					CuentadecobroFilaDetalle(pdf, miFila, a)
				}
			}

			pdf.AddPage()
			CuentadecobroCabecera(pdf, miCuentadecobro, miDetalle)
			// segunda pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 41 && a <= 82 {
					CuentadecobroFilaDetalle(pdf, miFila, a)
				}
			}

			pdf.AddPage()
			CuentadecobroCabecera(pdf, miCuentadecobro, miDetalle)
			// tercera pagina
			for i, miFila := range miDetalle {
				var a = i + 1
				if a > 82 {
					CuentadecobroFilaDetalle(pdf, miFila, a)
				}
			}

			CuentadecobroPieDePagina(pdf, miCuentadecobro)
		}
	}

	// genera pdf
	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

// TERMINA CUENTADECOBRO PDF

// INICIA EMPRESA CUENTA DE COBRO PDF
func CuentadecobroHeader(pdf *gofpdf.Fpdf, miCuentadecobro cuentadecobro) {
	// ENCABEZADO
	var e empresa = ListaEmpresa()
	pdf.SetHeaderFunc(func() {
		pdf.SetFont("Arial", "", 10)
		pdf.Image(imageFile("logo.png"), 20, 25, 40, 0, false,
			"", 0, "")

		// EMPRESA

		pdf.SetY(20)
		pdf.SetX(15)
		//pdf.CellFormat(185, 32, "", "0", 0, "C",
		//false, 0, "")

		//pdf.SetY(20)
		pdf.CellFormat(190, 4, e.Nombre, "0", 0,
			"C", false, 0, "")
		pdf.Ln(4)
		pdf.SetX(15)
		pdf.CellFormat(190, 4, "Nit No. "+Coma(e.Codigo)+" - "+e.Dv, "0", 0, "C",
			false, 0, "")
		pdf.Ln(4)
		pdf.SetX(15)
		pdf.CellFormat(190, 4, e.Direccion, "0", 0, "C", false, 0,
			"")
		pdf.Ln(4)
		pdf.SetX(15)
		pdf.CellFormat(190, 4, e.Telefono1, "0", 0, "C", false, 0,
			"")

		// NOMBRE DEL DOCUMENTO
		pdf.SetFont("Arial", "", 11)
		//pdf.Ln(10)
		pdf.SetY(24)
		pdf.SetX(85)
		pdf.CellFormat(183, 4, "CUENTA DE COBRO", "0", 0, "C",
			false, 0, "")
		pdf.Ln(6)
		pdf.SetX(85)
		pdf.CellFormat(190, 4, " No.  "+miCuentadecobro.Numero, "0", 0, "C",
			false, 0, "")
	})
}

// INICIA CABECERA
func CuentadecobroCabecera(pdf *gofpdf.Fpdf, miCuentadecobro cuentadecobro, miDetalle []cuentadecobrodetalleeditar) {
	miResidente := residente{}
	err3 := db.Get(&miResidente, "SELECT * FROM residente where codigo=$1", miCuentadecobro.Residente)
	if err3 != nil {
		log.Fatalln(err3)
	}

	var mesletras string
	mesletras = mesLetras(strconv.Itoa(int(miCuentadecobro.Fecha.Month())))
	pdf.SetFont("Arial", "", 10)
	pdf.Ln(10)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Codigo", "", 0,
		"L", false, 0, "")
	pdf.SetX(40)
	pdf.CellFormat(40, 4, miResidente.Codigo, "", 0,
		"L", false, 0, "")
	pdf.SetX(60)
	pdf.CellFormat(40, 4, "Nombre", "", 0,
		"L", false, 0, "")
	pdf.SetX(80)
	pdf.CellFormat(40, 4, miResidente.Nombre, "", 0,
		"L", false, 0, "")
	pdf.SetX(158)
	pdf.CellFormat(40, 4, "Nit. No.", "", 0,
		"L", false, 0, "")
	pdf.SetX(175)
	pdf.CellFormat(40, 4, Coma(miResidente.Nit), "", 0,
		"L", false, 0, "")

	pdf.Ln(5)
	pdf.SetX(20)
	pdf.CellFormat(40, 4, "Fecha", "", 0,
		"L", false, 0, "")
	pdf.SetX(40)
	pdf.CellFormat(40, 4, miCuentadecobro.Fecha.Format("02/01/2006"), "", 0,
		"L", false, 0, "")
	pdf.SetX(60)
	pdf.CellFormat(40, 4, "Direccion", "", 0,
		"L", false, 0, "")
	pdf.SetX(80)
	pdf.CellFormat(40, 4, miResidente.Direccion, "", 0,
		"L", false, 0, "")
	pdf.SetX(158)
	pdf.CellFormat(40, 4, "Telefono", "", 0,
		"L", false, 0, "")
	pdf.SetX(175)
	pdf.CellFormat(40, 4, miResidente.Telefono1, "", 0,
		"L", false, 0, "")
	pdf.Ln(-1)

	// CUADRO TITULO
	pdf.SetFont("Arial", "", 10)
	pdf.SetY(52)
	pdf.SetX(20)
	pdf.SetFillColor(224, 231, 239)
	pdf.SetTextColor(0, 0, 0)

	pdf.CellFormat(183, 6, "", "0", 0, "C",
		true, 0, "")
	pdf.SetX(20)
	pdf.CellFormat(184, 6, "No.", "0", 0,
		"L", false, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(190, 6, "CUENTA", "0", 0,
		"L", false, 0, "")
	pdf.SetX(55)
	pdf.CellFormat(190, 6, "NOMBRE", "0", 0,
		"L", false, 0, "")
	pdf.SetX(126)
	pdf.CellFormat(190, 6, "ANTERIOR", "0", 0,
		"L", false, 0, "")
	pdf.SetX(165)
	pdf.CellFormat(4, 6, Mayuscula(mesletras), "0", 0,
		"C", false, 0, "")
	pdf.SetX(190)
	pdf.CellFormat(190, 6, "TOTAL", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}

// TERMINA CABECERA

// INICIA DETALLE CUENTADECOBRO PDF
func CuentadecobroFilaDetalle(pdf *gofpdf.Fpdf, miFila cuentadecobrodetalleeditar, a int) {
	pdf.SetFont("Arial", "", 10)
	pdf.SetX(21)

	pdf.CellFormat(40, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(60, 4, miFila.Cuenta, "", 0,
		"L", false, 0, "")
	pdf.SetX(55)
	pdf.CellFormat(60, 4, Subcadena(miFila.Cuentanombre, 0, 40), "", 0,
		"L", false, 0, "")
	pdf.SetX(115)
	pdf.CellFormat(30, 4, Coma(miFila.Anterior), "", 0,
		"R", false, 0, "")
	pdf.SetX(145)
	pdf.CellFormat(30, 4, Coma(miFila.Actual), "", 0,
		"R", false, 0, "")
	pdf.SetX(174)
	pdf.CellFormat(30, 4, Coma(miFila.Total), "", 0,
		"R", false, 0, "")
	pdf.Ln(4)
}

// TERMINA DETALLE CUENTADECOBRO

// INICIA FINAL DE PAGINA
func CuentadecobroPieDePagina(pdf *gofpdf.Fpdf, miCuentadecobro cuentadecobro) {
	miResidente := residente{}
	err3 := db.Get(&miResidente, "SELECT * FROM residente where codigo=$1", miCuentadecobro.Residente)
	if err3 != nil {
		log.Fatalln(err3)
	}

	var totaldescuento1 string
	var totaldescuento2 string

	totaldescuento1 = FormatoFlotanteEntero(Flotante(miCuentadecobro.Total) - Flotante(miResidente.Descuento1))
	totaldescuento2 = FormatoFlotanteEntero(Flotante(miCuentadecobro.Total) - Flotante(miResidente.Descuento2))

	parametroscontabilidad := configuracioncontabilidad{}
	parametroscontabilidad = TraerParametrosContabilidad()

	pdf.SetFont("Arial", "", 9)
	pdf.Ln(10)
	pdf.SetY(114)
	pdf.SetX(20)
	pdf.CellFormat(40, 10, "", "0", 0,
		"C", false, 0, "")
	pdf.Ln(4)
	pdf.SetX(22)
	pdf.CellFormat(40, 10, "Se omite firma autografo Art. 10 D. R. 836/1991 Art. 1.6.1.12.12 DUR 1625 de 2016", "0", 0, "L",
		false, 0, "")

	pdf.SetFont("Arial", "", 10)

	// RELLENO TOTALES
	pdf.SetY(94)
	pdf.SetX(20)
	pdf.SetFillColor(224, 231, 239)
	pdf.SetTextColor(0, 0, 0)

	//pdf.CellFormat(50, 6, "", "0", 0, "L",
	//	true, 0, "")
	pdf.CellFormat(183, 6, "", "0", 0, "L",
		true, 0, "")
	pdf.SetX(55)
	pdf.CellFormat(40, 6, "TOTALES", "0", 0, "L",
		false, 0, "")
	//pdf.SetY(92)
	pdf.SetX(105)
	pdf.CellFormat(40, 6, Coma(miCuentadecobro.Totalanterior), "0", 0, "R",
		false, 0, "")
	//pdf.SetY(92)
	pdf.SetX(135)
	pdf.CellFormat(40, 6, Coma(miCuentadecobro.Totalactual), "0", 0, "R",
		false, 0, "")
	//pdf.SetY(92)
	pdf.SetX(164)
	pdf.CellFormat(40, 6, Coma(miCuentadecobro.Total), "0", 0, "R",
		false, 0, "")

	// CUADRO DETALLE
	//pdf.SetTextColor(0,0,0)

	//pdf.SetY(57)
	//pdf.SetX(20)

	//pdf.CellFormat(183, 43, "", "0", 0, "C",
	//	false, 0, "")

	// CUADRO TOTALES
	//pdf.SetY(94)
	//pdf.SetX(20)

	//pdf.CellFormat(183, 6, "", "0", 0, "C",
	//	false, 0, "")

	// CUADRO PIE
	//pdf.SetY(100)
	//pdf.SetX(145)
	//pdf.SetDrawColor(0,82,165)
	//pdf.CellFormat(60, 26, "", "0", 0, "C",false, 0, "")

	// CUADRO AVISO
	//pdf.SetY(100)
	//pdf.SetX(20)

	//pdf.CellFormat(125, 26, "", "0", 0, "C",false, 0, "")

	pdf.SetFont("Arial", "", 8)
	pdf.SetY(102)
	pdf.SetX(22)
	pdf.CellFormat(121, 6, parametroscontabilidad.Textoaviso1, "0", 0, "C",
		false, 0, "")

	pdf.SetY(106)
	pdf.SetX(22)
	pdf.CellFormat(121, 6, parametroscontabilidad.Textoaviso2, "0", 0, "C",
		false, 0, "")

	pdf.SetY(110)
	pdf.SetX(22)
	pdf.CellFormat(121, 6, parametroscontabilidad.Textoaviso3, "0", 0, "C",
		false, 0, "")

	pdf.SetY(112)
	pdf.SetX(22)
	pdf.CellFormat(121, 6, parametroscontabilidad.Textoaviso4, "0", 0, "C",
		false, 0, "")

	pdf.SetFont("Arial", "", 10)
	pdf.SetY(100)
	pdf.SetX(147)
	pdf.CellFormat(40, 10, parametroscontabilidad.Textodescuento1, "0", 0, "L",
		false, 0, "")
	pdf.SetY(100)
	pdf.SetX(164)
	pdf.CellFormat(40, 10, Coma(totaldescuento1), "0", 0, "R",
		false, 0, "")

	pdf.SetY(104)
	pdf.SetX(147)
	pdf.CellFormat(40, 10, parametroscontabilidad.Textodescuento2, "0", 0, "L",
		false, 0, "")
	pdf.SetY(104)
	pdf.SetX(164)
	pdf.CellFormat(40, 10, Coma(totaldescuento2), "0", 0, "R",
		false, 0, "")

	pdf.SetY(108)
	pdf.SetX(147)
	pdf.CellFormat(40, 10, parametroscontabilidad.Textodescuento3, "0", 0, "L",
		false, 0, "")

	pdf.SetY(108)
	pdf.SetX(164)
	pdf.CellFormat(40, 10, Coma(miCuentadecobro.Total), "0", 0, "R",
		false, 0, "")
}

// TERMINA FINAL DE PAGINA

// INICIA PIE DE PAGINA
func CuentadecobroFooter(pdf *gofpdf.Fpdf) {

	pdf.SetFooterFunc(func() {
		pdf.SetY(118)
		pdf.SetX(147)
		pdf.SetFont("Arial", "", 8)
		pdf.CellFormat(40, 10, "www.Sadconf.com.co", "",
			0, "L", false, 0, "")
		pdf.SetX(177)
		pdf.CellFormat(30, 10, fmt.Sprintf(" %d de {nb}", pdf.PageNo()), "",
			0, "R", false, 0, "")
	})

}

// TERMINA PIE DE PAGINA

// INICIA CUENTA DE COBRO TODOS PDF
func CuentadecobroTodosCabecera(pdf *gofpdf.Fpdf) {
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
	pdf.CellFormat(190, 6, "Numero", "0", 0,
		"L", false, 0, "")
	pdf.SetX(55)
	pdf.CellFormat(190, 6, "Fecha", "0", 0,
		"L", false, 0, "")
	pdf.SetX(75)
	pdf.CellFormat(190, 6, "Codigo", "0", 0,
		"L", false, 0, "")
	pdf.SetX(130)
	pdf.CellFormat(190, 6, "Nombre", "0", 0,
		"L", false, 0, "")
	pdf.SetX(193)
	pdf.CellFormat(190, 6, "Total", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}

func CuentadecobroTodosDetalle(pdf *gofpdf.Fpdf, miFila cuentadecobroLista, a int) {
	pdf.SetFont("Arial", "", 9)
	pdf.SetX(21)
	pdf.CellFormat(181, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(32)
	pdf.CellFormat(40, 4, miFila.Numero, "", 0,
		"L", false, 0, "")
	pdf.SetX(55)
	pdf.CellFormat(40, 4, miFila.Fecha.Format("02/01/2006"), "", 0, "L", false, 0, "")
	pdf.SetX(75)
	pdf.CellFormat(40, 4, Coma(miFila.Residente),
		"", 0, "L", false, 0, "")
	pdf.SetX(130)
	pdf.CellFormat(40, 4, Subcadena(miFila.Residentenombre, 0, 31), "", 0,
		"L", false, 0, "")
	pdf.SetX(164)
	pdf.CellFormat(40, 4, Coma(miFila.Total), "", 0,
		"R", false, 0, "")
	pdf.Ln(4)
}

func CuentadecobroTodosPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string

	consulta = "  SELECT cuentadecobro.numero,fecha,residente,residente.nombre as Residentenombre,cuentadecobro.total "
	consulta += " FROM cuentadecobro "
	consulta += " inner join residente on residente.codigo=cuentadecobro.residente "
	consulta += " ORDER BY cast(cuentadecobro.numero as integer) ASC"

	t := []cuentadecobroLista{}

	err := db.Select(&t, consulta)

	if err != nil {
		fmt.Println(err)
		return
	}

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
		pdf.CellFormat(190, 10, "DATOS CUENTAS DE COBRO", "0", 0,
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

	CuentadecobroTodosCabecera(pdf)
	// tercera pagina

	for i, miFila := range t {
		var a = i + 1
		if math.Mod(float64(a), 49) == 0 {
			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			CuentadecobroTodosCabecera(pdf)
		}
		CuentadecobroTodosDetalle(pdf, miFila, a)
	}

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

// TERMINA CUENTA DE COBRO TODOS PDF

// CUENTA DE COBRO EXCEL
func CuentadecobroExcel(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var consulta string

	consulta = "  SELECT cuentadecobro.numero,fecha,residente,residente.nombre as Residentenombre,cuentadecobro.total "
	consulta += " FROM cuentadecobro "
	consulta += " inner join residente on residente.codigo=cuentadecobro.residente "
	consulta += " ORDER BY cast(cuentadecobro.numero as integer) ASC"

	t := []cuentadecobroLista{}

	err := db.Select(&t, consulta)

	if err != nil {
		fmt.Println(err)
		return
	}

	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	f := excelize.NewFile()

	// FUNCION ANCHO DE LA COLUMNA
	if err = f.SetColWidth("Sheet1", "A", "A", 13); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.SetColWidth("Sheet1", "B", "B", 13); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.SetColWidth("Sheet1", "C", "C", 30); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "D", "D", 30); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.SetColWidth("Sheet1", "E", "E", 15); err != nil {
		fmt.Println(err)
		return
	}

	// FUNCION PARA UNIR DOS CELDAS
	if err = f.MergeCell("Sheet1", "A1", "E1"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A2", "E2"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A3", "E3"); err != nil {
		fmt.Println(err)
		return
	}
	if err = f.MergeCell("Sheet1", "A4", "E4"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A5", "E5"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A6", "E6"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A7", "E7"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A8", "E8"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A9", "E9"); err != nil {
		fmt.Println(err)
		return
	}

	if err = f.MergeCell("Sheet1", "A10", "E10"); err != nil {
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
	f.SetCellValue("Sheet1", "A9", "LISTADO DE CUENTAS DE COBRO")
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
	f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel), "Numero")
	f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel), "Fecha")
	f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel), "Codigo")
	f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel), "Nombre")
	f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel), "Total")

	f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel), "A"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel), "B"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel), "C"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel), "D"+strconv.Itoa(filaExcel), estiloCabecera)
	f.SetCellStyle("Sheet1", "E"+strconv.Itoa(filaExcel), "E"+strconv.Itoa(filaExcel), estiloCabecera)

	filaExcel++

	for i, miFila := range t {
		f.SetCellValue("Sheet1", "A"+strconv.Itoa(filaExcel+i), Flotante(miFila.Numero))
		f.SetCellValue("Sheet1", "B"+strconv.Itoa(filaExcel+i), miFila.Fecha.Format("02/01/2006"))
		f.SetCellValue("Sheet1", "C"+strconv.Itoa(filaExcel+i), miFila.Residente)
		f.SetCellValue("Sheet1", "D"+strconv.Itoa(filaExcel+i), miFila.Residentenombre)
		f.SetCellValue("Sheet1", "E"+strconv.Itoa(filaExcel+i), Flotante(miFila.Total))

		f.SetCellStyle("Sheet1", "A"+strconv.Itoa(filaExcel+i), "A"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "B"+strconv.Itoa(filaExcel+i), "B"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "C"+strconv.Itoa(filaExcel+i), "D"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "D"+strconv.Itoa(filaExcel+i), "E"+strconv.Itoa(filaExcel+i), estiloTexto)
		f.SetCellStyle("Sheet1", "E"+strconv.Itoa(filaExcel+i), "F"+strconv.Itoa(filaExcel+i), estiloNumeroDetalle)
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

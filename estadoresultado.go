package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type estadoresultadoparametro struct {
	FechaInicial        string `json:"FechaInicial"`
	FechaFinal          string `json:"FechaFinal"`
	Nivel               string `json:"Nivel"`
	ComparaFechaInicial string `json:"ComparaFechaInicial"`
	ComparaFechaFinal   string `json:"ComparaFechaFinal"`
	Comparativo         string `json:"Comparativo"`
}

func sumaCuentaEstadoResultado(cuenta plandecuentaempresa, datos []datosresumenfinanciero, fechaInicial string, fechaFinal string, minivel string) balancedeprueba {
	//	inicioperiodo, err := time.Parse("2006-01-02", "2021-01-01")
	dateinicial, err := time.Parse("2006-01-02", fechaInicial)
	//datefinal, err := time.Parse("2006-01-02", fechaFinal)

	if err == nil {
		fmt.Println("Fecha Inicial suma" + dateinicial.String())
	}

	var totalanterior float64
	var debitoanterior float64
	var creditoanterior float64
	var debito float64
	var credito float64
	var saldo float64

	debitoanterior = 0
	creditoanterior = 0
	debito = 0
	credito = 0
	saldo = 0

	for _, x := range datos {

		log.Println("cuentadatos : " + x.Cuenta)
		log.Println("fecha movimiento : " + x.Fecha.String())
		log.Println("cuenta parametro : " + cuenta.Codigo)
		log.Println("cuenta cortada : " + x.Cuenta[0:len(cuenta.Codigo)])

		if cuenta.Codigo == x.Cuenta[0:len(cuenta.Codigo)] {

			if x.Fecha.Before(dateinicial) {
				debitoanterior += x.Debito
				creditoanterior += x.Credito
				log.Println("movimiento anterior  : " + x.Fecha.String())
			} else {
				debito += x.Debito
				credito += x.Credito
				log.Println("movimiento mes  : " + x.Fecha.String())
			}

		}

		//listadobalancedeprueba=append(listadobalancedeprueba, balancedeprueba{x.Fecha,strconv.Itoa(i),x.Cuenta, })
	}

	if "1" == cuenta.Codigo[0:1] || "5" == cuenta.Codigo[0:1] || "6" == cuenta.Codigo[0:1] || "7" == cuenta.Codigo[0:1] || "8" == cuenta.Codigo[0:1] {
		totalanterior = debitoanterior - creditoanterior
		saldo = totalanterior + debito - credito
	} else {

		totalanterior = creditoanterior - debitoanterior
		saldo = totalanterior + credito - debito

	}

	log.Println("total anterior cuenta  : " + FormatoFlotante(totalanterior))

	if cuenta.Nivel == minivel {
		return balancedeprueba{"", cuenta.Codigo,
			cuenta.Nombre, FormatoFlotante(totalanterior),
			FormatoFlotante(debito),
			FormatoFlotante(credito),
			FormatoFlotante(saldo), cuenta.Nivel, "NO"}

	} else {
		return balancedeprueba{"", cuenta.Codigo,
			cuenta.Nombre, "",
			"",
			"",
			"", cuenta.Nivel, "NO"}

	}

}

func saldoComparaEstado(micuenta string, parametro estadoresultadoparametro) float64 {
	if parametro.Comparativo == "No" {
		return 0
	}

	var consultasaldoAnterior = "SELECT debito,credito FROM comprobantedetalle  where"
	consultasaldoAnterior += "  (fecha>=$1 and fecha<=$2) and cuenta=$3"
	//
	listadoSaldoTercero := []cuentaSaldo{}
	db.Select(&listadoSaldoTercero, consultasaldoAnterior, parametro.ComparaFechaInicial, parametro.ComparaFechaFinal, micuenta)
	// fechaInicial, fechaFinal,
	log.Println("calculo cuenta anterior " + micuenta)
	var totalSaldo float64
	totalSaldo = 0
	for _, miSaldo := range listadoSaldoTercero {
		//var c=k
		log.Println("suma saldo")
		log.Println("suma saldo")

		log.Println(miSaldo.Debito)
		log.Println(miSaldo.Credito)

		var tipocuenta = Subcadena(micuenta, 0, 1)
		log.Println("tipo" + tipocuenta)
		if tipocuenta == "1" || tipocuenta == "5" || tipocuenta == "6" || tipocuenta == "7" || tipocuenta == "8" {
			totalSaldo += Flotante(miSaldo.Debito) - Flotante(miSaldo.Credito)
		} else {
			totalSaldo += Flotante(miSaldo.Credito) - Flotante(miSaldo.Debito)
		}
	}
	return totalSaldo

}

func EstadoResultadoEstado(migrupo string, tempParametro estadoresultadoparametro, miperiodo string) ([]balancedeprueba, float64) {
	db := dbConn()
	var consulta string
	consulta = ""
	consulta = "select distinct plandecuentaempresa.financiero as financiero, fecha,cuenta,sum(debito)as debito,sum(credito) as credito from comprobantedetalle "
	consulta += " inner join plandecuentaempresa on plandecuentaempresa.codigo=comprobantedetalle.cuenta "
	consulta += " where (fecha<=$1 and plandecuentaempresa.financiero=$2 "
	consulta += " and  extract(year from fecha)=$3 )"
	consulta += " and not  (comprobantedetalle.documento='99' and comprobantedetalle.numero='13')"
	consulta += " group by fecha,cuenta,financiero   "

	listadoDatos := []datosresumenfinanciero{}
	listadobalancedeprueba := []balancedeprueba{}

	err1 := db.Select(&listadoDatos, consulta,
		tempParametro.FechaFinal, migrupo, miperiodo)

	if err1 != nil {
		panic(err1.Error())
	}

	listadoCuenta := []plandecuentaempresa{}
	err2 := db.Select(&listadoCuenta, "select * from plandecuentaempresa where nivel=$1 order by codigo", tempParametro.Nivel)
	if err2 != nil {
		panic(err2.Error())
	}

	// nivel inferior A
	if tempParametro.Nivel == "A" {
		var q string

		var primero bool
		primero = true
		q = ""
		for _, x := range listadoDatos {
			if primero == true {
				q += "  codigo=substring('" + x.Cuenta + "',1,2) or codigo=substring('" + x.Cuenta + "',1,4)"
				primero = false
			} else {
				q += " or  codigo=substring('" + x.Cuenta + "',1,2) or codigo=substring('" + x.Cuenta + "',1,4)"

			}
		}
		if len(listadoDatos) > 0 {
			q = " where " + q

			fmt.Println("consulta cuentas " + q)

			listadoCuentaa := []plandecuentaempresa{}
			err2 := db.Select(&listadoCuentaa, "select * from plandecuentaempresa  "+q+" order by codigo")
			if err2 != nil {
				panic(err2.Error())
			}

			for _, x := range listadoCuentaa {
				listadobalancedeprueba = append(listadobalancedeprueba, balancedeprueba{"", x.Codigo, x.Nombre, "0", "0", "0", "0", x.Nivel, "NO"})
			}
		}
	}

	// nivel 3
	if tempParametro.Nivel == "3" {
		var q string
		var primero bool
		primero = true
		q = ""
		for _, x := range listadoDatos {
			if primero == true {
				q += " codigo=substring('" + x.Cuenta + "',1,2)  "
				fmt.Println("consulta cuentas 3333" + q)
				primero = false
			} else {
				q += " or codigo=substring('" + x.Cuenta + "',1,2) "

			}
		}
		if len(listadoDatos) > 0 {
			q = " where " + q

			fmt.Println("consulta cuentas444 " + q)
			listadoCuenta3 := []plandecuentaempresa{}
			err2 := db.Select(&listadoCuenta3, "select * from plandecuentaempresa "+q+" order by codigo")
			if err2 != nil {
				panic(err2.Error())
			}

			for _, x := range listadoCuenta3 {
				listadobalancedeprueba = append(listadobalancedeprueba, balancedeprueba{"", x.Codigo, x.Nombre, "0", "0", "0", "0", x.Nivel, "NO"})
			}
		}
	}

	var nivelInicial = ""
	var Anterior1 float64
	var Anterior2 float64
	var Anterior3 float64
	var Anterior4 float64
	var Anterior5 float64
	var Anterior6 float64
	var Anterior7 float64
	var Anterior8 float64
	var Anterior9 float64

	var Saldo1 float64
	var Saldo2 float64
	var Saldo3 float64
	var Saldo4 float64
	var Saldo5 float64
	var Saldo6 float64
	var Saldo7 float64
	var Saldo8 float64
	var Saldo9 float64

	var Debito1 float64
	var Debito2 float64
	var Debito3 float64
	var Debito4 float64
	var Debito5 float64
	var Debito6 float64
	var Debito7 float64
	var Debito8 float64
	var Debito9 float64

	var Credito1 float64
	var Credito2 float64
	var Credito3 float64
	var Credito4 float64
	var Credito5 float64
	var Credito6 float64
	var Credito7 float64
	var Credito8 float64
	var Credito9 float64
	var miBalanceFinal balancedeprueba
	for _, x := range listadoCuenta {
		fmt.Println("cuenta listado" + x.Codigo)
		//var a = i
		if nivelInicial == "" {
			nivelInicial = x.Nivel
		}

		var miBalance balancedeprueba

		miBalance = sumaCuentaEstadoResultado(x, listadoDatos, tempParametro.FechaInicial, tempParametro.FechaFinal, tempParametro.Nivel)

		//	miBalanceCompara = sumaCuentaEstadoResultado(x, listadoDatos, tempParametro.FechaInicial, tempParametro.ComparaFechaFinal, tempParametro.Nivel)

		if miBalance.Anterior == "0.00" && miBalance.Debito == "0.00" && miBalance.Credito == "0.00" && miBalance.Saldo == "0.00" {
		} else {
			listadobalancedeprueba = append(listadobalancedeprueba, miBalance)
		}
		// suma totales finales
		if x.Nivel == nivelInicial {

			switch x.Codigo[0:1] {
			case "1":
				Anterior1 += Flotante(miBalance.Anterior)
				Saldo1 += Flotante(miBalance.Saldo)
				Debito1 += Flotante(miBalance.Debito)
				Credito1 += Flotante(miBalance.Credito)
			case "2":
				Anterior2 += Flotante(miBalance.Anterior)
				Saldo2 += Flotante(miBalance.Saldo)
				Debito2 += Flotante(miBalance.Debito)
				Credito2 += Flotante(miBalance.Credito)
			case "3":
				Anterior3 += Flotante(miBalance.Anterior)
				Saldo3 += Flotante(miBalance.Saldo)
				Debito3 += Flotante(miBalance.Debito)
				Credito3 += Flotante(miBalance.Credito)
			case "4":
				Anterior4 += Flotante(miBalance.Anterior)
				Saldo4 += Flotante(miBalance.Saldo)
				Debito4 += Flotante(miBalance.Debito)
				Credito4 += Flotante(miBalance.Credito)
			case "5":
				Anterior5 += Flotante(miBalance.Anterior)
				Saldo5 += Flotante(miBalance.Saldo)
				Debito5 += Flotante(miBalance.Debito)
				Credito5 += Flotante(miBalance.Credito)
			case "6":
				Anterior6 += Flotante(miBalance.Anterior)
				Saldo6 += Flotante(miBalance.Saldo)
				Debito6 += Flotante(miBalance.Debito)
				Credito6 += Flotante(miBalance.Credito)
			case "7":
				Anterior7 += Flotante(miBalance.Anterior)
				Saldo7 += Flotante(miBalance.Saldo)
				Debito7 += Flotante(miBalance.Debito)
				Credito7 += Flotante(miBalance.Credito)
			case "8":
				Anterior8 += Flotante(miBalance.Anterior)
				Saldo8 += Flotante(miBalance.Saldo)
				Debito8 += Flotante(miBalance.Debito)
				Credito8 += Flotante(miBalance.Credito)
			case "9":
				Anterior9 += Flotante(miBalance.Anterior)
				Saldo9 += Flotante(miBalance.Saldo)
				Debito9 += Flotante(miBalance.Debito)
				Credito9 += Flotante(miBalance.Credito)
			default:
				fmt.Println("Too far away.")
			}
		}

	}
	var anteriorFinal float64
	var debitoFinal float64
	var creditoFinal float64
	var saldoFinal float64

	// total final
	if (Anterior1 + Anterior5 + Anterior6 + Anterior7 + Anterior8) == 0 {
		anteriorFinal = (Anterior2 + Anterior3 + Anterior4 + Anterior9)
	} else {
		anteriorFinal = (Anterior1 + Anterior5 + Anterior6 + Anterior7 + Anterior8) - (Anterior2 + Anterior3 + Anterior4 + Anterior9)
	}

	if (Saldo1 + Saldo5 + Saldo6 + Saldo7 + Saldo8) == 0 {
		saldoFinal = Saldo2 + Saldo3 + Saldo4 + Saldo9
	} else {
		saldoFinal = (Saldo1 + Saldo5 + Saldo6 + Saldo7 + Saldo8) - (Saldo2 + Saldo3 + Saldo4 + Saldo9)
	}

	debitoFinal = Debito1 + Debito2 + Debito3 + Debito4 + Debito5 + Debito6 + Debito7 + Debito8 + Debito9
	creditoFinal = Credito1 + Credito2 + Credito3 + Credito4 + Credito5 + Credito6 + Credito7 + Credito8 + Credito9

	miBalanceFinal.Descripcion = "TOTALES"
	miBalanceFinal.SiFinal = "SI"
	miBalanceFinal.Anterior = FormatoFlotante(anteriorFinal)
	miBalanceFinal.Debito = FormatoFlotante(debitoFinal)
	miBalanceFinal.Credito = FormatoFlotante(creditoFinal)
	miBalanceFinal.Saldo = FormatoFlotante(saldoFinal)

	//listadobalancedeprueba=append(listadobalancedeprueba,miBalanceFinal)

	sort.Slice(listadobalancedeprueba, func(i, j int) bool {
		switch strings.Compare(listadobalancedeprueba[i].Codigo, listadobalancedeprueba[j].Codigo) {
		case -1:
			return true
		case 1:
			return false
		}
		return listadobalancedeprueba[i].Codigo > listadobalancedeprueba[j].Codigo
	})
	return listadobalancedeprueba, saldoFinal
}

func GenerarEstadoResultado(tempParametro estadoresultadoparametro, miperiodo string) []situacionfinanciero {
	listadosituacion := []situacionfinanciero{}
	listadosituacionfinal := []situacionfinanciero{}

	// ingreso ordinarios
	var Total_Io float64
	listadoIngresosOrdinarios := []balancedeprueba{}
	listadoIngresosOrdinarios, Total_Io = EstadoResultadoEstado("6", tempParametro, miperiodo)
	listadosituacion = append(listadosituacion, situacionfinanciero{"", "", "INGRESOS ORDINARIOS", "0", "", "", "", "NO"})
	var saldoCompara float64
	saldoCompara = 0

	for _, miFila := range listadoIngresosOrdinarios {

		saldoCompara = saldoComparaEstado(miFila.Codigo, tempParametro)
		listadosituacion = append(listadosituacion, situacionfinanciero{"", miFila.Codigo, miFila.Descripcion, FormatoFlotante(saldoCompara), "", miFila.Saldo, "0", "NO"})
	}
	listadosituacion = append(listadosituacion, situacionfinanciero{"", "", "TOTAL INGRESOS ORDINARIOS", "", "", "0", FormatoFlotante(Total_Io), "NO"})

	// costo de venta
	var Total_Cv float64
	listadoCostoDeVenta := []balancedeprueba{}
	listadoCostoDeVenta, Total_Cv = EstadoResultadoEstado("14", tempParametro, miperiodo)
	listadosituacion = append(listadosituacion, situacionfinanciero{"", "", "COSTO DE VENTA", "", "", "0", "", "NO"})
	for _, miFila := range listadoCostoDeVenta {
		listadosituacion = append(listadosituacion, situacionfinanciero{"", miFila.Codigo, miFila.Descripcion, "", "", miFila.Saldo, "0", "NO"})
	}
	listadosituacion = append(listadosituacion, situacionfinanciero{"", "", "TOTAL COSTO DE VENTA", "", "", "0", FormatoFlotante(Total_Cv), "NO"})
	// titulo
	listadosituacion = append(listadosituacion, situacionfinanciero{"", "", "UTILIDAD BRUTA EN VENTAS", "", "", "0", FormatoFlotante(Total_Io - Total_Cv), "NO"})

	// gastos de administraci[on
	var Total_Ga float64
	listadoGastosAdministracion := []balancedeprueba{}
	listadoGastosAdministracion, Total_Ga = EstadoResultadoEstado("9", tempParametro, miperiodo)
	listadosituacion = append(listadosituacion, situacionfinanciero{"", "", "GASTOS", "", "", "0", "", "NO"})
	listadosituacion = append(listadosituacion, situacionfinanciero{"", "", "GASTOS DE ADMINISTRACION", "", "", "0", "", "NO"})
	for _, miFila := range listadoGastosAdministracion {
		listadosituacion = append(listadosituacion, situacionfinanciero{"", miFila.Codigo, miFila.Descripcion, "", "", miFila.Saldo, "0", "NO"})
	}
	listadosituacion = append(listadosituacion, situacionfinanciero{"", "", "TOTAL GASTOS DE ADINISTRACION", "", "", "0", FormatoFlotante(Total_Ga), "NO"})

	// gastos de ventas

	var Total_Gv float64
	listadoGastosVentas := []balancedeprueba{}
	listadoGastosVentas, Total_Gv = EstadoResultadoEstado("10", tempParametro, miperiodo)

	if Total_Gv != 0 {
		listadosituacion = append(listadosituacion, situacionfinanciero{"", "", "GASTOS DE VENTAS", "", "", "0", "", "NO"})
		for _, miFila := range listadoGastosVentas {
			listadosituacion = append(listadosituacion, situacionfinanciero{"", miFila.Codigo, miFila.Descripcion, "", "", miFila.Saldo, "0", "NO"})
		}
		listadosituacion = append(listadosituacion, situacionfinanciero{"", "", "TOTAL GASTOS EN VENTAS", "", "", "0", FormatoFlotante(Total_Gv), "NO"})
	}
	// titulo
	listadosituacion = append(listadosituacion, situacionfinanciero{"", "", "UTILIDAD OPERACIONAL", "", "", "0", FormatoFlotante((Total_Io - Total_Cv) - (Total_Gv + Total_Ga)), "NO"})

	// otros ingresos

	var Total_Oi float64
	listadoOtrosIngresos := []balancedeprueba{}
	listadoOtrosIngresos, Total_Oi = EstadoResultadoEstado("8", tempParametro, miperiodo)
	if Total_Gv != 0 {
		listadosituacion = append(listadosituacion, situacionfinanciero{"", "", "OTROS INGRESOS", "", "", "0", "", "NO"})
		for _, miFila := range listadoOtrosIngresos {
			listadosituacion = append(listadosituacion, situacionfinanciero{"", miFila.Codigo, miFila.Descripcion, "", "", miFila.Saldo, "0", "NO"})
		}
		listadosituacion = append(listadosituacion, situacionfinanciero{"", "", "TOTAL OTROS INGRESOS", "", "", "0", FormatoFlotante(Total_Oi), "NO"})

	}

	// otros Egresos

	var Total_Oe float64
	listadoOtrosEgresos := []balancedeprueba{}
	listadoOtrosEgresos, Total_Oe = EstadoResultadoEstado("12", tempParametro, miperiodo)
	listadosituacion = append(listadosituacion, situacionfinanciero{"", "", "OTROS EGRESOS", "", "", "0", "", "NO"})
	for _, miFila := range listadoOtrosEgresos {
		listadosituacion = append(listadosituacion, situacionfinanciero{"", miFila.Codigo, miFila.Descripcion, "", "", miFila.Saldo, "0", "NO"})
	}
	listadosituacion = append(listadosituacion, situacionfinanciero{"", "", "TOTAL OTROS EGRESOS", "", "", "0", FormatoFlotante(Total_Oe), "NO"})

	// titulo
	listadosituacion = append(listadosituacion, situacionfinanciero{"", "", "UTILIDAD ANTES DE IMPUESTOS", "", "", "0", FormatoFlotante(((Total_Io - Total_Cv) - (Total_Gv + Total_Ga)) + Total_Oi - Total_Oe), "NO"})

	// Impuesto de renta

	var Total_Ir float64
	listadoImpuestoRenta := []balancedeprueba{}
	listadoImpuestoRenta, Total_Ir = EstadoResultadoEstado("13", tempParametro, miperiodo)

	if Total_Ir != 0 {

		listadosituacion = append(listadosituacion, situacionfinanciero{"", "", "IMPUESTO A LA RENTA", "", "", "0", "", "NO"})
		for _, miFila := range listadoImpuestoRenta {
			listadosituacion = append(listadosituacion, situacionfinanciero{"", miFila.Codigo, miFila.Descripcion, "", "", miFila.Saldo, "0", "NO"})
		}
		listadosituacion = append(listadosituacion, situacionfinanciero{"", "", "TOTAL IMPUESTO A LA RENTA", "", "", "0", FormatoFlotante(Total_Ir), "NO"})

	}

	// titulo
	//listadosituacion=append(listadosituacion,situacionfinanciero{"","","UTILIDAD NETA DEL EJERCICIO","0",FormatoFlotante(((Total_Io-Total_Cv)-(Total_Gv+Total_Ga))+Total_Oi-Total_Oe),"NO"})
	listadosituacion = append(listadosituacion, situacionfinanciero{"", "", "UTILIDAD NETA DEL EJERCICIO", "", "", "0", FormatoFlotante((((Total_Io - Total_Cv) - (Total_Gv + Total_Ga)) + Total_Oi - Total_Oe) - Total_Ir), "NO"})

	//var parcial string
	//var total string
	for _, miFila := range listadosituacion {
		if miFila.Parcial == "0" {
			miFila.Parcial = ""
		}

		if miFila.Total == "0" {
			miFila.Total = ""
		}

		//parcial = miFila.Parcial
		listadosituacionfinal = append(listadosituacionfinal, situacionfinanciero{"", miFila.Codigo, miFila.Descripcion, "", "", miFila.Parcial, miFila.Total, miFila.SiFinal})
	}
	return listadosituacionfinal

}

func EstadoResultadoDatos(w http.ResponseWriter, r *http.Request) {
	//db := dbConn()
	var miperiodo = periodoSesion(r)

	var tempParametro estadoresultadoparametro
	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	// carga informacion de la venta
	err = json.Unmarshal(b, &tempParametro)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	listadosituacion := []situacionfinanciero{}

	listadosituacion = GenerarEstadoResultado(tempParametro, miperiodo)
	//listadobalancedeprueba := []balancedeprueba{}

	//	var miBalancetotal balancedeprueba

	//TOTAL_IO=GENERA_GRUPO("IO","INGRESOS ORDINARIOS")
	//total_costo=GENERA_GRUPO("CV","COSTO DE VENTAS")
	//CREA_TITULO("UTILIDAD BRUTA EN VENTA", TOTAL_IO-total_costo)
	//TOTAL_GA=GENERA_GRUPO("GA","GASTOS ADMINISTRACION")
	//TOTAL_GV=GENERA_GRUPO("GV","GASTOS DE VENTAS")
	//CREA_TITULO("UTILIDAD OPERACIONAL", (TOTAL_IO-total_costo)-(TOTAL_GV+TOTAL_GA))
	//TOTAL_OI=GENERA_GRUPO("OI","OTROS INGRESOS")
	//TOTAL_OE=GENERA_GRUPO("OE","OTROS EGRESOS")
	//CREA_TITULO("UTILIDAD ANTES DE IMPUESTOS",( (TOTAL_IO-total_costo)-(TOTAL_GV+TOTAL_GA))+TOTAL_OI-TOTAL_OE)
	//TOTAL_IR=GENERA_GRUPO("IR","IMPUESTO DE RENTA")
	//CREA_TITULO("UTILIDAD NETA DEL EJERCICIO",(( (TOTAL_IO-total_costo)-(TOTAL_GV+TOTAL_GA))+TOTAL_OI-TOTAL_OE)-TOTAL_IR)

	//1,Activo Corriente
	//2,Activo No Corriente
	//3,Pasivo Corriente
	//4,Pasivo No Corriente
	//5,Patrimonio
	//6,Ingresos Ordinarios
	//8,Otros Ingresos
	//7,Ingresos Financieros
	//9,Gastos De Administracion
	//10,Gastos De Ventas
	//11,Gastos Financieros
	//12,Otros Egresos
	//13,Impuesto A La Renta
	//14 Costo de venta

	//var cadena string
	var siexiste bool
	siexiste = true

	if siexiste == false {
		var slice []string
		slice = make([]string, 0)
		data, _ := json.Marshal(slice)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	} else {
		data, _ := json.Marshal(listadosituacion)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

// CENTRO BALANCE DE PRUEBA
func EstadoResultadoLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/estadoresultado/EstadoResultadoLista.html",
		"vista/estadoresultado/Autocompletaplandecuentaempresa.html",
		"vista/estadoresultado/Autocompletatercero.html",
		"vista/estadoresultado/Autocompletacentro.html",
		"vista/estadoresultado/Autocompletadocumento.html")
	//	db := dbConn()

	varmap := map[string]interface{}{
		//"res":     listadobalancedeprueba,
		"hosting":   ruta,
		"bodega":    ListaBodega(),
		"producto":  ListaProducto(),
		"miperiodo": periodoSesion(r),
	}
	tmp.Execute(w, varmap)
}

// INICIA BALANCE DE PRUEBA TODOS PDF
func EstadoResultadoPdf(w http.ResponseWriter, r *http.Request) {
	//	db := dbConn()
	var miperiodo = periodoSesion(r)
	var tempParametro estadoresultadoparametro
	tempParametro.FechaInicial = mux.Vars(r)["FechaInicial"]
	tempParametro.FechaFinal = mux.Vars(r)["FechaFinal"]
	tempParametro.Nivel = mux.Vars(r)["Nivel"]
	//b, err := ioutil.ReadAll(r.Body)
	//defer r.Body.Close()
	//if err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}
	//// carga informacion de la venta
	//err = json.Unmarshal(b, &tempParametro)
	//if err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}

	listadobalancedeprueba := []situacionfinanciero{}

	listadobalancedeprueba = GenerarEstadoResultado(tempParametro, miperiodo)

	DateInicial, _ := time.Parse("2006-01-02", tempParametro.FechaInicial)
	DateFinal, _ := time.Parse("2006-01-02", tempParametro.FechaFinal)

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
		pdf.CellFormat(190, 10, e.Telefono1+" - "+e.Telefono2, "0", 0, "C", false, 0,
			"")
		pdf.Ln(4)
		pdf.CellFormat(190, 10, ene(c.NombreCiudad+" - "+c.NombreDepartamento), "0", 0, "C", false, 0,
			"")
		pdf.Ln(6)
		pdf.CellFormat(190, 10, "ESTADO DE RESULTADOS INTEGRAL", "0",
			0, "C", false, 0, "")
		pdf.Ln(6)
		pdf.CellFormat(190, 10, "Del "+DateInicial.Format("02/01/2006")+" Al "+DateFinal.Format("02/01/2006"), "0", 0,
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

	EstadoResultadoCabecera(pdf, tempParametro)

	for i, miFila := range listadobalancedeprueba {
		EstadoResultadoFilaDetalle(pdf, miFila, i+1)

		if math.Mod(float64(i+1), 47) == 0 {
			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			EstadoResultadoCabecera(pdf, tempParametro)
		}

	}
	//EstadoResultadoPieDePagina(pdf,listadobalancedeprueba)
	//BalancePieDePagina(pdf,miBalancetotal)

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

func EstadoResultadoCabecera(pdf *gofpdf.Fpdf, tempParametro estadoresultadoparametro) {
	pdf.SetFont("Arial", "", 10)
	// RELLENO TITULO
	pdf.SetY(52)
	pdf.SetFillColor(224, 231, 239)
	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(10)
	pdf.SetX(20)
	pdf.CellFormat(184, 6, "Codigo", "0", 0,
		"L", true, 0, "")
	pdf.SetX(46)
	pdf.CellFormat(190, 6, "Cuenta", "0", 0,
		"L", false, 0, "")
	pdf.SetX(106)
	pdf.CellFormat(190, 6, "Anterior", "0", 0,
		"L", false, 0, "")
	pdf.SetX(136)
	pdf.CellFormat(190, 6, "", "0", 0,
		"L", false, 0, "")
	pdf.SetX(163)
	pdf.CellFormat(190, 6, "Actual", "0", 0,
		"L", false, 0, "")
	pdf.SetX(193)
	pdf.CellFormat(190, 6, "", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}
func EstadoResultadoFilaDetalle(pdf *gofpdf.Fpdf, miFila situacionfinanciero, a int) {
	pdf.SetFont("Arial", "", 9)

	ene := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetTextColor(0, 0, 0)
	// fila normal
	pdf.SetX(20)
	pdf.CellFormat(40, 4, Subcadena(miFila.Codigo, 0, 12), "", 0,
		"L", false, 0, "")
	pdf.SetX(37)
	pdf.CellFormat(40, 4, Subcadena(ene(miFila.Descripcion), 0, 30), "", 0,
		"L", false, 0, "")
	pdf.SetX(81)
	pdf.CellFormat(40, 4, miFila.Parcial, "", 0,
		"R", false, 0, "")
	pdf.SetX(109)
	pdf.CellFormat(40, 4, miFila.Total, "", 0,
		"R", false, 0, "")
	//pdf.SetX(137)
	//pdf.CellFormat(40, 4, miFila.Credito, "", 0,
	//	"R", false, 0, "")
	//pdf.SetX(165)
	//pdf.CellFormat(40, 4, miFila.Saldo, "", 0,
	//	"R", false, 0, "")
	pdf.SetX(141)
	pdf.Ln(4)
}

func EstadoResultadoPieDePagina(pdf *gofpdf.Fpdf, miFila situacionfinanciero) {
	pdf.SetFillColor(224, 231, 239)
	pdf.SetTextColor(0, 0, 0)

	pdf.SetX(20)
	pdf.CellFormat(20, 6, "", "", 0,
		"L", true, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(40, 6, Subcadena(miFila.Codigo, 0, 12), "", 0,
		"L", true, 0, "")
	pdf.SetX(46)
	pdf.CellFormat(47, 6, Subcadena((miFila.Descripcion), 0, 25), "", 0,
		"L", true, 0, "")
	pdf.SetX(93)
	pdf.CellFormat(28, 6, miFila.Parcial, "", 0,
		"R", true, 0, "")
	pdf.SetX(121)
	pdf.CellFormat(28, 6, miFila.Total, "", 0,
		"R", true, 0, "")
	//pdf.SetX(149)
	//pdf.CellFormat(28, 6, miFila.Credito, "", 0,
	//	"R", true, 0, "")
	//pdf.SetX(177)
	//pdf.CellFormat(28, 6, miFila.Saldo, "", 0,
	//	"R", true, 0, "")
	//pdf.SetX(141)
	pdf.Ln(4)
}

// BALANCE DE PRUEBA EXCEL
func EstadoResultadoExcel(w http.ResponseWriter, r *http.Request) {
	//db := dbConn()
	var miperiodo = periodoSesion(r)
	var tempParametro estadoresultadoparametro
	tempParametro.FechaInicial = mux.Vars(r)["FechaInicial"]
	tempParametro.FechaFinal = mux.Vars(r)["FechaFinal"]
	tempParametro.Nivel = mux.Vars(r)["Nivel"]

	listadobalancedeprueba := []situacionfinanciero{}

	listadobalancedeprueba = GenerarEstadoResultado(tempParametro, miperiodo)

	DateInicial, _ := time.Parse("2006-01-02", tempParametro.FechaInicial)
	DateFinal, _ := time.Parse("2006-01-02", tempParametro.FechaFinal)

	//t := inventario{}
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	//err := db.Get(&t, "SELECT * FROM inventario ")

	f := excelize.NewFile()
	if err := f.MergeCell("Sheet1", "A1", "F1"); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "B", "B", 24); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.SetColWidth("Sheet1", "C", "C", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth("Sheet1", "D", "D", 13); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.SetColWidth("Sheet1", "E", "E", 13); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.SetColWidth("Sheet1", "F", "F", 13); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.MergeCell("Sheet1", "A2", "F2"); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.MergeCell("Sheet1", "A3", "F3"); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.MergeCell("Sheet1", "A4", "F4"); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.MergeCell("Sheet1", "A5", "F5"); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.MergeCell("Sheet1", "A6", "F6"); err != nil {
		fmt.Println(err)
		return
	}

	estiloTitulo, err := f.NewStyle(`{  "alignment":{"horizontal": "center"},"font":{"bold":false,"italic":false,"family":"Arial","size":8,"color":"##000000"}}`)

	// titulo
	f.SetCellValue("Sheet1", "A1", e.Nombre)
	f.SetCellValue("Sheet1", "A2", "Nit No. "+Coma(e.Codigo)+" - "+e.Dv)
	f.SetCellValue("Sheet1", "A3", e.Iva+" - "+e.ReteIva)
	f.SetCellValue("Sheet1", "A4", "Actividad Ica - "+e.ActividadIca)
	f.SetCellValue("Sheet1", "A5", e.Direccion)
	f.SetCellValue("Sheet1", "A6", (c.NombreCiudad + " - " + c.NombreDepartamento))
	f.SetCellValue("Sheet1", "A6", "BALANCE DE PRUEBA DEL "+DateInicial.Format("02/01/2006")+" AL "+DateFinal.Format("02/01/2006"))
	f.SetCellStyle("Sheet1", "A1", "A1", estiloTitulo)
	f.SetCellStyle("Sheet1", "A2", "A2", estiloTitulo)
	f.SetCellStyle("Sheet1", "A3", "A3", estiloTitulo)
	f.SetCellStyle("Sheet1", "A4", "A4", estiloTitulo)
	f.SetCellStyle("Sheet1", "A5", "A5", estiloTitulo)
	f.SetCellStyle("Sheet1", "A6", "A6", estiloTitulo)

	var filaExcel = 8
	//var a string
	//a = ""

	estiloTexto, err := f.NewStyle(`{"font":{"bold":false,"italic":false,"family":"Arial","size":8,"color":"##000000"}}`)

	estiloNumeroDetalle, err := f.NewStyle(`{"number_format": 4,"font":{"bold":false,"italic":false,"family":"Arial","size":8,"color":"##000000"}}`)

	if err != nil {
		fmt.Println(err)
	}

	for i, miFila := range listadobalancedeprueba {
		//	var van int
		var a = strconv.Itoa(filaExcel + i)
		f.SetCellValue("Sheet1", "A"+a, miFila.Codigo)
		f.SetCellValue("Sheet1", "B"+a, miFila.Descripcion)
		//f.SetCellValue("Sheet1", "C"+a, Flotante(miFila.Anterior))

		//f.SetCellValue("Sheet1", "D"+a, Flotante(miFila.Debito))
		f.SetCellValue("Sheet1", "E"+a, Flotante(miFila.Parcial))
		f.SetCellValue("Sheet1", "F"+a, Flotante(miFila.Total))

		f.SetCellStyle("Sheet1", "A"+a, "B"+a, estiloTexto)
		f.SetCellStyle("Sheet1", "B"+a, "B"+a, estiloTexto)
		f.SetCellStyle("Sheet1", "C"+a, "C"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "D"+a, "D"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "E"+a, "E"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "F"+a, "F"+a, estiloNumeroDetalle)
		//	van = i
	}

	// LIENA FINAL
	//a = strconv.Itoa(van + 1 + filaExcel)
	//f.SetCellValue("Sheet1", "A"+a, miBalancetotal.Codigo)
	//f.SetCellValue("Sheet1", "B"+a, miBalancetotal.Descripcion)
	//f.SetCellValue("Sheet1", "C"+a, Flotante(miBalancetotal.Anterior))
	//f.SetCellValue("Sheet1", "D"+a, Flotante(miBalancetotal.Debito))
	//f.SetCellValue("Sheet1", "E"+a, Flotante(miBalancetotal.Credito))
	//f.SetCellValue("Sheet1", "F"+a, Flotante(miBalancetotal.Saldo))
	//
	//// aplica formato
	//f.SetCellStyle("Sheet1", "A"+a, "B"+a, estiloTexto)
	//f.SetCellStyle("Sheet1", "B"+a, "B"+a, estiloTexto)
	//f.SetCellStyle("Sheet1", "C"+a, "C"+a, estiloNumeroDetalle)
	//f.SetCellStyle("Sheet1", "D"+a, "D"+a, estiloNumeroDetalle)
	//f.SetCellStyle("Sheet1", "E"+a, "E"+a, estiloNumeroDetalle)
	//f.SetCellStyle("Sheet1", "F"+a, "F"+a, estiloNumeroDetalle)

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

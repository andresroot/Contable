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
	"regexp"
	"strconv"
	"time"
)

type balancedeprueba struct {
	Filas       string `json:"Filas"`
	Codigo      string `json:"Codigo"`
	Descripcion string `json:"Descripcion"`
	Anterior    string `json:"Anterior"`
	Debito      string `json:"Debito"`
	Credito     string `json:"Credito"`
	Saldo       string `json:"Saldo"`
	Nivel       string `json:"Nivel"`
	SiFinal     string `json:"SiFinal"`
}

type balancetercero struct {
	tercero string
}

type balancecuenta struct {
	cuenta string
}

type datosresumen struct {
	Fecha   time.Time `json:"Fecha"`
	Debito  float64   `json:"Debito"`
	Credito float64   `json:"Credito"`
	Cuenta  string    `json:"Cuenta"`
}

type datosresumentercero struct {
	Fecha   time.Time `json:"Fecha"`
	Debito  float64   `json:"Debito"`
	Credito float64   `json:"Credito"`
	Cuenta  string    `json:"Cuenta"`
	Tercero string    `json:"Tercero"`
}

type datostercero struct {
	Tercero       string `json:"Tercero"`
	TerceroNombre string `json:"TerceroNombre"`
}

type datosdetalle struct {
	Cuenta            string
	Tercero           string
	Centro            string
	Concepto          string
	Factura           string
	Debito            float64
	Credito           float64
	Documento         string
	Numero            string
	Fecha             time.Time
	Fechaconsignacion time.Time
}

type balanceparametro struct {
	FechaInicial     string `json:"FechaInicial"`
	FechaFinal       string `json:"FechaFinal"`
	CuentaInicial    string `json:"CuentaInicial"`
	CuentaFinal      string `json:"CuentaFinal"`
	TerceroInicial   string `json:"TerceroInicial"`
	TerceroFinal     string `json:"TerceroFinal"`
	CentroInicial    string `json:"CentroInicial"`
	CentroFinal      string `json:"CentroFinal"`
	DocumentoInicial string `json:"DocumentoInicial"`
	DocumentoFinal   string `json:"DocumentoFinal"`
	NumeroInicial    string `json:"NumeroInicial"`
	NumeroFinal      string `json:"NumeroFinal"`
	Detalle          string `json:"Detalle"`
	Nivel            string `json:"Nivel"`
	Activa           string `json:"Activa"`
	Subtotal         string `json:"Subtotal"`
}

func sumaCuenta(cuenta plandecuentaempresa, datosmovimiento []datosresumen, fechaInicial string, fechaFinal string) balancedeprueba {
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

	for _, x := range datosmovimiento {
		if cuenta.Codigo == x.Cuenta[0:len(cuenta.Codigo)] {
			//log.Println("cuentadatos : " + x.Cuenta)
			//log.Println("fecha movimiento : " + x.Fecha.String())
			//log.Println("cuenta parametro : " + cuenta.Codigo)
			//
			//log.Println("cuenta cortada : " + x.Cuenta[0:len(cuenta.Codigo)])

			if x.Fecha.Before(dateinicial) {
				debitoanterior += x.Debito
				creditoanterior += x.Credito
				//			log.Println("movimiento anterior  : " + x.Fecha.String())
			} else {
				debito += x.Debito
				credito += x.Credito
				//	log.Println("movimiento mes  : " + x.Fecha.String())
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

	return balancedeprueba{"", cuenta.Codigo,
		cuenta.Nombre, FormatoFlotante(totalanterior),
		FormatoFlotante(debito),
		FormatoFlotante(credito),
		FormatoFlotante(saldo), cuenta.Nivel, "NO"}
}

func sumaCuentaTercerp(cuenta plandecuentaempresa, datosmovimiento []datosresumen, fechaInicial string, fechaFinal string) balancedeprueba {
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

	for _, x := range datosmovimiento {
		if cuenta.Codigo == x.Cuenta[0:len(cuenta.Codigo)] {

			//log.Println("fecha movimiento : " + x.Fecha.String())
			//log.Println("cuenta parametro : " + cuenta.Codigo)
			//
			//log.Println("cuenta cortada : " + x.Cuenta[0:len(cuenta.Codigo)])

			if x.Fecha.Before(dateinicial) {
				debitoanterior += x.Debito
				creditoanterior += x.Credito
				//log.Println("cuentadatostercero  debitocredito anterior: " + x.Tercero)
				//			log.Println("movimiento anterior  : " + x.Fecha.String())
			} else {
				//log.Println("cuentadatostercero  debitocredito: " + x.Tercero)

				debito += x.Debito
				credito += x.Credito
				//	log.Println("movimiento mes  : " + x.Fecha.String())
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

	//log.Println("total anterior cuenta  : " + FormatoFlotante(totalanterior))

	return balancedeprueba{"", cuenta.Codigo,
		cuenta.Nombre + "", FormatoFlotante(totalanterior),
		FormatoFlotante(debito),
		FormatoFlotante(credito),
		FormatoFlotante(saldo), cuenta.Nivel, "NO"}
}

func InformeDetalle(tempParametro balanceparametro) []balancedeprueba {
	log.Println("Inicia Balance Detalle")
	layout := "2006-01-02"

	dateInicial, _ := time.Parse(layout, tempParametro.FechaInicial)

	var consulta string
	consulta = ""
	consulta = "select distinct fecha,cuenta,sum(debito)as debito,sum(credito) as credito from comprobantedetalle "
	consulta += " where fecha<=$1 "
	//consulta += " and  (extract(year from fecha)>=" + tempParametro.FechaInicial[0:4] + " "
	consulta += " and  extract(year from fecha)<=" + tempParametro.FechaFinal[0:4] + " "
	consulta += " and (Cuenta>=$2 and cuenta<= $3) "
	consulta += " and (documento>=$4 and documento<= $5) "
	consulta += " and (centro>=$6 and centro<= $7) "
	consulta += " and (tercero>=$8 and tercero<= $9)"
	consulta += " and (numero>=$10 and numero<= $11)"
	consulta += "group by fecha,cuenta"

	consulta += ""
	listadoDatos := []datosresumen{}
	listadoDatosDetalle := []datosdetalle{}
	log.Println("Inicia Detalle 2" + consulta)
	listadobalancedeprueba := []balancedeprueba{}

	err1 := db.Select(&listadoDatos, consulta,
		tempParametro.FechaFinal,
		tempParametro.CuentaInicial,
		tempParametro.CuentaFinal,
		tempParametro.DocumentoInicial,
		tempParametro.DocumentoFinal,
		tempParametro.CentroInicial,
		tempParametro.CentroFinal,
		tempParametro.TerceroInicial,
		tempParametro.TerceroFinal,
		tempParametro.NumeroInicial,
		tempParametro.NumeroFinal)
	if err1 != nil {
		println(err1.Error())
	}

	log.Println("Inicia Detalle3")

	consulta = ""
	consulta = "select Cuenta,Tercero,Centro,Concepto,Factura ,Debito ,Credito,Documento,Numero,Fecha,Fechaconsignacion  from comprobantedetalle "
	consulta += " where fecha<=$1 "
	consulta += " and (Cuenta>=$2 and cuenta<= $3) "
	consulta += " and (documento>=$4 and documento<= $5) "
	consulta += " and (centro>=$6 and centro<= $7) "
	consulta += " and (tercero>=$8 and tercero<= $9)"
	consulta += " and (numero>=$10 and numero<= $11)"
	consulta += "order by fecha"

	err2 := db.Select(&listadoDatosDetalle, consulta,
		tempParametro.FechaFinal,
		tempParametro.CuentaInicial,
		tempParametro.CuentaFinal,
		tempParametro.DocumentoInicial,
		tempParametro.DocumentoFinal,
		tempParametro.CentroInicial,
		tempParametro.CentroFinal,
		tempParametro.TerceroInicial,
		tempParametro.TerceroFinal,
		tempParametro.NumeroInicial,
		tempParametro.NumeroFinal)
	if err2 != nil {
		panic(err1.Error())
	}

	listadoCuenta := []plandecuentaempresa{}
	err3 := db.Select(&listadoCuenta, "select * from plandecuentaempresa where nivel='A' order by codigo")
	if err3 != nil {
		println(err2.Error())
	}

	for _, x := range listadoCuenta {
		//var a = i
		var miBalance balancedeprueba
		miBalance = sumaCuenta(x, listadoDatos, tempParametro.FechaInicial, tempParametro.FechaFinal)
		if miBalance.Anterior == "0.00" && miBalance.Debito == "0.00" && miBalance.Credito == "0.00" && miBalance.Saldo == "0.00" {
		} else {

			for _, d := range listadoDatosDetalle {
				if x.Codigo == d.Cuenta[0:len(x.Codigo)] {

					if !d.Fecha.Before(dateInicial) {

						//	var codigo=d.Fecha.Format("02/01")+" "+d.Centro
						var concepto = d.Fecha.Format("02/01") + " " + d.Centro + " " + d.Documento + " " + d.Numero + " " + d.Concepto

						listadobalancedeprueba = append(listadobalancedeprueba, balancedeprueba{"",
							"",
							concepto,
							"",
							FormatoFlotante(d.Debito),
							FormatoFlotante(d.Credito),
							"",
							"",
							"NO",
						})

					}
				}
			}
			listadobalancedeprueba = append(listadobalancedeprueba, miBalance)
		}
	}
	return listadobalancedeprueba
}

func InformeDetalleTercero(tempParametro balanceparametro) []balancedeprueba {
	log.Println("Inicia Balance Detalle")
	layout := "2006-01-02"
	db := dbConn()
	dateInicial, _ := time.Parse(layout, tempParametro.FechaInicial)

	var consulta string
	consulta = ""
	consulta = "select distinct fecha,cuenta,sum(debito)as debito,sum(credito) as credito from comprobantedetalle "
	consulta += " where fecha<=$1 "
	//consulta += " and  (extract(year from fecha)>=" + tempParametro.FechaInicial[0:4] + " "
	consulta += " and  extract(year from fecha)<=" + tempParametro.FechaFinal[0:4] + " "
	consulta += " and (Cuenta>=$2 and cuenta<= $3) "
	consulta += " and (documento>=$4 and documento<= $5) "
	consulta += " and (centro>=$6 and centro<= $7) "
	consulta += " and (tercero>=$8 and tercero<= $9)"
	consulta += " and (numero>=$10 and numero<= $11)"
	consulta += "group by fecha,cuenta"

	consulta += ""
	listadoDatos := []datosresumen{}
	listadoDatosDetalle := []datosdetalle{}
	log.Println("Inicia Detalle 2" + consulta)
	listadobalancedeprueba := []balancedeprueba{}

	err1 := db.Select(&listadoDatos, consulta,
		tempParametro.FechaFinal,
		tempParametro.CuentaInicial,
		tempParametro.CuentaFinal,
		tempParametro.DocumentoInicial,
		tempParametro.DocumentoFinal,
		tempParametro.CentroInicial,
		tempParametro.CentroFinal,
		tempParametro.TerceroInicial,
		tempParametro.TerceroFinal,
		tempParametro.NumeroInicial,
		tempParametro.NumeroFinal)

	if err1 != nil {
		println(err1.Error())
	}

	log.Println("Inicia Detalle3")

	consulta = ""
	consulta = "select Cuenta,Tercero,Centro,Concepto,Factura ,Debito ,Credito,Documento,Numero,Fecha,Fechaconsignacion  from comprobantedetalle "
	consulta += " where fecha<=$1 "
	consulta += " and  extract(year from fecha)<=" + tempParametro.FechaFinal[0:4] + " "
	consulta += " and (Cuenta>=$2 and cuenta<= $3) "
	consulta += " and (documento>=$4 and documento<= $5) "
	consulta += " and (centro>=$6 and centro<= $7) "
	consulta += " and (tercero>=$8 and tercero<= $9)"
	consulta += " and (numero>=$10 and numero<= $11)"
	consulta += "order by fecha"

	err2 := db.Select(&listadoDatosDetalle, consulta,
		tempParametro.FechaFinal,
		tempParametro.CuentaInicial,
		tempParametro.CuentaFinal,
		tempParametro.DocumentoInicial,
		tempParametro.DocumentoFinal,
		tempParametro.CentroInicial,
		tempParametro.CentroFinal,
		tempParametro.TerceroInicial,
		tempParametro.TerceroFinal,
		tempParametro.NumeroInicial,
		tempParametro.NumeroFinal)
	if err2 != nil {
		panic(err2.Error())
	}

	// listado de tercero s
	listadoTercero := []datostercero{}
	consulta = ""
	consulta = "  select distinct comprobantedetalle.tercero as tercero,tercero.nombre as terceronombre from comprobantedetalle "
	consulta += "  inner join tercero on tercero.codigo=comprobantedetalle.tercero "
	consulta += " where comprobantedetalle.fecha<=$1 "
	consulta += " and  extract(year from fecha)<=" + tempParametro.FechaFinal[0:4] + " "
	consulta += " and (comprobantedetalle.Cuenta>=$2 and comprobantedetalle.cuenta<= $3) "
	consulta += " and (comprobantedetalle.documento>=$4 and comprobantedetalle.documento<= $5) "
	consulta += " and (comprobantedetalle.centro>=$6 and comprobantedetalle.centro<= $7) "
	consulta += " and (comprobantedetalle.tercero>=$8 and comprobantedetalle.tercero<= $9)"
	consulta += " and (comprobantedetalle.numero>=$10 and comprobantedetalle.numero<= $11)"
	consulta += " group by comprobantedetalle.tercero,tercero.nombre"

	//	consulta += "order by tercero"
	log.Println("" + consulta)

	err4 := db.Select(&listadoTercero, consulta,
		tempParametro.FechaFinal,
		tempParametro.CuentaInicial,
		tempParametro.CuentaFinal,
		tempParametro.DocumentoInicial,
		tempParametro.DocumentoFinal,
		tempParametro.CentroInicial,
		tempParametro.CentroFinal,
		tempParametro.TerceroInicial,
		tempParametro.TerceroFinal,
		tempParametro.NumeroInicial,
		tempParametro.NumeroFinal)
	if err4 != nil {
		panic(err4.Error())
	}

	listadoCuenta := []plandecuentaempresa{}
	err3 := db.Select(&listadoCuenta, "select * from plandecuentaempresa where nivel='A' order by codigo")
	if err3 != nil {
		println(err2.Error())
	}
	var debitotercero float64
	var creditotercero float64
	var debitoanteriortercero float64
	var creditoanteriortercero float64

	var saldo float64
	var totalanterior float64

	for _, x := range listadoCuenta {
		//log.Println("recorre cuenta" + x.Codigo)
		//var a = i
		var miBalance balancedeprueba
		miBalance = sumaCuentaTercerp(x, listadoDatos, tempParametro.FechaInicial, tempParametro.FechaFinal)
		if miBalance.Anterior == "0.00" && miBalance.Debito == "0.00" && miBalance.Credito == "0.00" && miBalance.Saldo == "0.00" {
		} else {

			for _, mitercero := range listadoTercero {
				//log.Println("recorre tercero" + mitercero.Tercero)
				debitotercero = 0
				creditotercero = 0
				debitoanteriortercero = 0
				creditoanteriortercero = 0
				saldo = 0
				totalanterior = 0
				for _, d := range listadoDatosDetalle {
					//	log.Println("tercero detalle detalle " + mitercero.Tercero)
					//	log.Println("tercero detalle cuenta " + x.Codigo)

					if x.Codigo == d.Cuenta && d.Tercero == mitercero.Tercero {

						if !d.Fecha.Before(dateInicial) {
							debitotercero += d.Debito
							creditotercero += d.Credito

							//	log.Println("tercero moviimento " + mitercero.Tercero)

							//	var codigo=d.Fecha.Format("02/01")+" "+d.Centro
							var concepto = d.Fecha.Format("02/01") + " " + d.Centro + " " + d.Documento + " " + d.Numero + " " + d.Concepto

							listadobalancedeprueba = append(listadobalancedeprueba, balancedeprueba{"",
								"",
								concepto,
								"",
								FormatoFlotante(d.Debito),
								FormatoFlotante(d.Credito),
								"",
								"",
								"NO",
							})
						}

					} else {
						debitoanteriortercero += d.Debito
						creditoanteriortercero += d.Credito
					}
					// fin detalle
				}

				if "1" == x.Codigo[0:1] || "5" == x.Codigo[0:1] || "6" == x.Codigo[0:1] || "7" == x.Codigo[0:1] || "8" == x.Codigo[0:1] {
					totalanterior = debitoanteriortercero - creditoanteriortercero
					saldo = totalanterior + debitotercero - creditotercero
				} else {

					totalanterior = creditoanteriortercero - debitoanteriortercero
					saldo = totalanterior + creditotercero - debitotercero

				}
				if FormatoFlotante(saldo) == "0.00" && FormatoFlotante(totalanterior) == "0.00" && FormatoFlotante(debitotercero) == "0.00" && FormatoFlotante(creditotercero) == "0.00" {

				} else {
					// fin tercero
					listadobalancedeprueba = append(listadobalancedeprueba, balancedeprueba{"", enterocoma((mitercero.Tercero)), mitercero.TerceroNombre, FormatoFlotante(totalanterior), FormatoFlotante(debitotercero), FormatoFlotante(creditotercero), FormatoFlotante(saldo), "", "NO"})
				}

			}
			// fin cuenta
			listadobalancedeprueba = append(listadobalancedeprueba, miBalance)
		}

	}
	return listadobalancedeprueba
}
func enterocoma(minumero string) string {
	var num = Entero(minumero)
	str := fmt.Sprintf("%d", num)
	re := regexp.MustCompile("(\\d+)(\\d{3})")
	for n := ""; n != str; {
		n = str
		str = re.ReplaceAllString(str, "$1,$2")
	}
	return str
}
func InformeResumen(tempParametro balanceparametro, miperiodo string) ([]balancedeprueba, balancedeprueba) {
	var consulta string
	log.Print("iNICIO CONSULTA BALANCE" + tempParametro.FechaFinal[0:4])
	consulta = ""
	consulta = "select distinct fecha,cuenta,sum(debito)as debito,sum(credito) as credito from comprobantedetalle "
	consulta += " where fecha<=$1 "
	//consulta += " and  (extract(year from fecha)>=" + tempParametro.FechaInicial[0:4] + " "
	consulta += " and  extract(year from fecha)<=" + tempParametro.FechaFinal[0:4] + " "

	consulta += " and (Cuenta>=$2 and cuenta<= $3) "
	consulta += " and (documento>=$4 and documento<= $5) "
	consulta += " and (centro>=$6 and centro<= $7) "
	consulta += " and (tercero>=$8 and tercero<= $9)"
	consulta += " and (numero>=$10 and numero<= $11)"
	consulta += "group by fecha,cuenta"
	consulta += ""
	listadoMovimiento := []datosresumen{}
	listadobalancedeprueba := []balancedeprueba{}
	err1 := db.Select(&listadoMovimiento, consulta,
		tempParametro.FechaFinal,
		tempParametro.CuentaInicial,
		tempParametro.CuentaFinal,
		tempParametro.DocumentoInicial,
		tempParametro.DocumentoFinal,
		tempParametro.CentroInicial,
		tempParametro.CentroFinal,
		tempParametro.TerceroInicial,
		tempParametro.TerceroFinal,
		tempParametro.NumeroInicial,
		tempParametro.NumeroFinal)
	if err1 != nil {
		println(err1.Error())
	}
	log.Print("FIN  CONSULTA BALANCE")
	listadoCuenta := []plandecuentaempresa{}
	err2 := db.Select(&listadoCuenta, "select * from plandecuentaempresa where nivel<=$1 order by codigo", tempParametro.Nivel)
	if err2 != nil {
		println(err2.Error())
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
		//var a = i
		if nivelInicial == "" {
			nivelInicial = x.Nivel
		}

		var miBalance balancedeprueba
		//println("iNICIO SALDO CUENTA")
		miBalance = sumaCuenta(x, listadoMovimiento, tempParametro.FechaInicial, tempParametro.FechaFinal)
		//	println("FIN SALDO CUENTA ")
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

	return listadobalancedeprueba, miBalanceFinal
}

func BalancedepruebaDatos(w http.ResponseWriter, r *http.Request) {
	//db := dbConn()
	var miperiodo = periodoSesion(r)
	var tempParametro balanceparametro
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
	listadobalancedeprueba := []balancedeprueba{}
	var miBalancetotal balancedeprueba
	if tempParametro.Detalle == "NO" {

		listadobalancedeprueba, miBalancetotal = InformeResumen(tempParametro, miperiodo)
	} else {
		if tempParametro.Subtotal == "Cuenta" {
			listadobalancedeprueba = InformeDetalle(tempParametro)
		} else {
			listadobalancedeprueba = InformeDetalleTercero(tempParametro)
		}
	}
	listadobalancedeprueba = append(listadobalancedeprueba, miBalancetotal)

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
		data, _ := json.Marshal(listadobalancedeprueba)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

// CENTRO BALANCE DE PRUEBA
func BalancedepruebaLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/balancedeprueba/balancedepruebaLista.html",
		"vista/balancedeprueba/Autocompletaplandecuentaempresa.html",
		"vista/balancedeprueba/Autocompletatercero.html",
		"vista/balancedeprueba/Autocompletacentro.html",
		"vista/balancedeprueba/Autocompletadocumento.html")
	//	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	//res := []inventario{}
	//listadobalancedeprueba := []balancedeprueba{}

	if Codigo == "False" {

	} else {

		//	FechaInicial := mux.Vars(r)["fechainicial"]

	}

	varmap := map[string]interface{}{
		"miperiodo": periodoSesion(r),
		"hosting":   ruta,
		"bodega":    ListaBodega(),
		"producto":  ListaProducto(),
	}
	tmp.Execute(w, varmap)
}

// INICIA BALANCE DE PRUEBA TODOS PDF
func BalancedepruebaPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	var miperiodo = periodoSesion(r)
	var tempParametro balanceparametro
	tempParametro.FechaInicial = mux.Vars(r)["FechaInicial"]
	tempParametro.FechaFinal = mux.Vars(r)["FechaFinal"]
	tempParametro.CuentaInicial = mux.Vars(r)["CuentaInicial"]
	tempParametro.CuentaFinal = mux.Vars(r)["CuentaFinal"]
	tempParametro.TerceroInicial = mux.Vars(r)["TerceroInicial"]
	tempParametro.TerceroFinal = mux.Vars(r)["TerceroFinal"]
	tempParametro.CentroInicial = mux.Vars(r)["CentroInicial"]
	tempParametro.CentroFinal = mux.Vars(r)["CentroFinal"]
	tempParametro.DocumentoInicial = mux.Vars(r)["DocumentoInicial"]
	tempParametro.DocumentoFinal = mux.Vars(r)["DocumentoFinal"]
	tempParametro.NumeroInicial = mux.Vars(r)["NumeroInicial"]
	tempParametro.NumeroFinal = mux.Vars(r)["NumeroFinal"]
	tempParametro.Detalle = mux.Vars(r)["Detalle"]
	tempParametro.Nivel = mux.Vars(r)["Nivel"]
	tempParametro.Activa = mux.Vars(r)["Activa"]
	tempParametro.Subtotal = mux.Vars(r)["Subtotal"]

	DateInicial, _ := time.Parse("2006-01-02", tempParametro.FechaInicial)
	DateFinal, _ := time.Parse("2006-01-02", tempParametro.FechaFinal)

	listadobalancedeprueba := []balancedeprueba{}
	var miBalancetotal balancedeprueba
	if tempParametro.Detalle == "NO" {
		listadobalancedeprueba, miBalancetotal = InformeResumen(tempParametro, miperiodo)
	} else {
		listadobalancedeprueba = InformeDetalle(tempParametro)

	}

	t := inventario{}
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	log.Println("Imprime balance pdf 1 ")
	err := db.Get(&t, "SELECT * FROM inventario ")
	if err != nil {
		log.Println(err)
	}
	log.Println("Imprime balance pdf 2 ")
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
		pdf.CellFormat(190, 10, "Balance De Prueba del "+DateInicial.Format("02/01/2006")+" AL "+DateFinal.Format("02/01/2006"), "0", 0,
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

	BalancePruebaCabecera(pdf, tempParametro)

	for i, miFila := range listadobalancedeprueba {
		BalancePruebaFilaDetalle(pdf, miFila, i+1)

		if math.Mod(float64(i+1), 48) == 0 {
			pdf.AliasNbPages("")
			pdf.AddPage()
			pdf.SetFont("Arial", "", 10)
			pdf.SetX(30)
			BalancePruebaCabecera(pdf, tempParametro)
		}

	}
	BalancePieDePagina(pdf, miBalancetotal)
	//BalancePieDePagina(pdf,miBalancetotal)

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

func BalancePruebaCabecera(pdf *gofpdf.Fpdf, tempParametro balanceparametro) {
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
	pdf.CellFormat(190, 6, "Codigo", "0", 0,
		"L", false, 0, "")
	pdf.SetX(55)
	pdf.CellFormat(190, 6, "Descripcion", "0", 0,
		"L", false, 0, "")
	pdf.SetX(106)
	pdf.CellFormat(190, 6, "Anterior", "0", 0,
		"L", false, 0, "")
	pdf.SetX(136)
	pdf.CellFormat(190, 6, "Debito", "0", 0,
		"L", false, 0, "")
	pdf.SetX(163)
	pdf.CellFormat(190, 6, "Credito", "0", 0,
		"L", false, 0, "")
	pdf.SetX(193)
	pdf.CellFormat(190, 6, "Saldo", "0", 0,
		"L", false, 0, "")
	pdf.Ln(8)
}
func BalancePruebaFilaDetalle(pdf *gofpdf.Fpdf, miFila balancedeprueba, a int) {
	pdf.SetFont("Arial", "", 9)

	ene := pdf.UnicodeTranslatorFromDescriptor("")
	pdf.SetTextColor(0, 0, 0)
	// fila normal
	pdf.SetX(21)
	pdf.CellFormat(183, 4, strconv.Itoa(a), "", 0,
		"L", false, 0, "")
	pdf.SetX(30)
	pdf.CellFormat(40, 4, Subcadena(miFila.Codigo, 0, 12), "", 0,
		"L", false, 0, "")
	pdf.SetX(46)
	pdf.CellFormat(40, 4, Subcadena(ene(miFila.Descripcion), 0, 25), "", 0,
		"L", false, 0, "")
	pdf.SetX(81)
	pdf.CellFormat(40, 4, miFila.Anterior, "", 0,
		"R", false, 0, "")
	pdf.SetX(109)
	pdf.CellFormat(40, 4, miFila.Debito, "", 0,
		"R", false, 0, "")
	pdf.SetX(137)
	pdf.CellFormat(40, 4, miFila.Credito, "", 0,
		"R", false, 0, "")
	pdf.SetX(165)
	pdf.CellFormat(40, 4, miFila.Saldo, "", 0,
		"R", false, 0, "")
	pdf.SetX(141)
	pdf.Ln(4)
}

func BalancePieDePagina(pdf *gofpdf.Fpdf, miFila balancedeprueba) {
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
	pdf.CellFormat(28, 6, miFila.Anterior, "", 0,
		"R", true, 0, "")
	pdf.SetX(121)
	pdf.CellFormat(28, 6, miFila.Debito, "", 0,
		"R", true, 0, "")
	pdf.SetX(149)
	pdf.CellFormat(28, 6, miFila.Credito, "", 0,
		"R", true, 0, "")
	pdf.SetX(177)
	pdf.CellFormat(28, 6, miFila.Saldo, "", 0,
		"R", true, 0, "")
	pdf.SetX(141)
	pdf.Ln(4)
}

// BALANCE DE PRUEBA EXCEL
func BalancedepruebaExcel(w http.ResponseWriter, r *http.Request) {
	//db := dbConn()
	var miperiodo = periodoSesion(r)
	var tempParametro balanceparametro
	tempParametro.FechaInicial = mux.Vars(r)["FechaInicial"]
	tempParametro.FechaFinal = mux.Vars(r)["FechaFinal"]
	tempParametro.CuentaInicial = mux.Vars(r)["CuentaInicial"]
	tempParametro.CuentaFinal = mux.Vars(r)["CuentaFinal"]
	tempParametro.TerceroInicial = mux.Vars(r)["TerceroInicial"]
	tempParametro.TerceroFinal = mux.Vars(r)["TerceroFinal"]
	tempParametro.CentroInicial = mux.Vars(r)["CentroInicial"]
	tempParametro.CentroFinal = mux.Vars(r)["CentroFinal"]
	tempParametro.DocumentoInicial = mux.Vars(r)["DocumentoInicial"]
	tempParametro.DocumentoFinal = mux.Vars(r)["DocumentoFinal"]
	tempParametro.NumeroInicial = mux.Vars(r)["NumeroInicial"]
	tempParametro.NumeroFinal = mux.Vars(r)["NumeroFinal"]
	tempParametro.Detalle = mux.Vars(r)["Detalle"]
	tempParametro.Nivel = mux.Vars(r)["Nivel"]
	tempParametro.Activa = mux.Vars(r)["Activa"]
	tempParametro.Subtotal = mux.Vars(r)["Subtotal"]

	DateInicial, _ := time.Parse("2006-01-02", tempParametro.FechaInicial)
	DateFinal, _ := time.Parse("2006-01-02", tempParametro.FechaFinal)

	listadobalancedeprueba := []balancedeprueba{}
	var miBalancetotal balancedeprueba
	if tempParametro.Detalle == "NO" {
		listadobalancedeprueba, miBalancetotal = InformeResumen(tempParametro, miperiodo)
	} else {
		listadobalancedeprueba = InformeDetalle(tempParametro)

	}

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
	var a string
	a = ""
	var van int
	estiloTexto, err := f.NewStyle(`{"font":{"bold":false,"italic":false,"family":"Arial","size":8,"color":"##000000"}}`)

	estiloNumeroDetalle, err := f.NewStyle(`{"number_format": 4,"font":{"bold":false,"italic":false,"family":"Arial","size":8,"color":"##000000"}}`)

	if err != nil {
		fmt.Println(err)
	}

	for i, miFila := range listadobalancedeprueba {
		a := strconv.Itoa(filaExcel + i)
		f.SetCellValue("Sheet1", "A"+a, miFila.Codigo)
		f.SetCellValue("Sheet1", "B"+a, miFila.Descripcion)
		f.SetCellValue("Sheet1", "C"+a, Flotante(miFila.Anterior))

		f.SetCellValue("Sheet1", "D"+a, Flotante(miFila.Debito))
		f.SetCellValue("Sheet1", "E"+a, Flotante(miFila.Credito))
		f.SetCellValue("Sheet1", "F"+a, Flotante(miFila.Saldo))

		f.SetCellStyle("Sheet1", "A"+a, "B"+a, estiloTexto)
		f.SetCellStyle("Sheet1", "B"+a, "B"+a, estiloTexto)
		f.SetCellStyle("Sheet1", "C"+a, "C"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "D"+a, "D"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "E"+a, "E"+a, estiloNumeroDetalle)
		f.SetCellStyle("Sheet1", "F"+a, "F"+a, estiloNumeroDetalle)
		van = i
	}

	// LIENA FINAL
	a = strconv.Itoa(van + 1 + filaExcel)
	f.SetCellValue("Sheet1", "A"+a, miBalancetotal.Codigo)
	f.SetCellValue("Sheet1", "B"+a, miBalancetotal.Descripcion)
	f.SetCellValue("Sheet1", "C"+a, Flotante(miBalancetotal.Anterior))
	f.SetCellValue("Sheet1", "D"+a, Flotante(miBalancetotal.Debito))
	f.SetCellValue("Sheet1", "E"+a, Flotante(miBalancetotal.Credito))
	f.SetCellValue("Sheet1", "F"+a, Flotante(miBalancetotal.Saldo))

	// aplica formato
	f.SetCellStyle("Sheet1", "A"+a, "B"+a, estiloTexto)
	f.SetCellStyle("Sheet1", "B"+a, "B"+a, estiloTexto)
	f.SetCellStyle("Sheet1", "C"+a, "C"+a, estiloNumeroDetalle)
	f.SetCellStyle("Sheet1", "D"+a, "D"+a, estiloNumeroDetalle)
	f.SetCellStyle("Sheet1", "E"+a, "E"+a, estiloNumeroDetalle)
	f.SetCellStyle("Sheet1", "F"+a, "F"+a, estiloNumeroDetalle)

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

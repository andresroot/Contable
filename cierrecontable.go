package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"
)

type cuentaSaldoCierre struct {
	Cuenta  string
	Debito  string
	Credito string
	Saldo   string
}

type listadoCierre struct {
	Fila        string
	Codigo      string
	Descripcion string
	Debito      string
	Credito     string
}

func CierrecontableLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/autocompleta/autocompletaCentro.html",
		"vista/cierrecontable/cierrecontableLista.html",
	)

	Codigo := mux.Vars(r)["codigo"]
	if Codigo == "False" {
	} else {
	}

	varmap := map[string]interface{}{
		"hosting":   ruta,
		"bodega":    ListaBodega(),
		"producto":  ListaProducto(),
		"miperiodo": periodoSesion(r),
	}
	tmp.Execute(w, varmap)
}

func InformeResumenCierre(FechaInicial string, FechaFinal string, miperiodo string) ([]balancedeprueba, balancedeprueba) {
	fmt.Println("FECHA FINAL " + FechaFinal)

	var consulta string
	consulta = ""
	consulta = "select distinct fecha,cuenta,sum(debito)as debito,sum(credito) as credito from comprobantedetalle "
	consulta += " where fecha<=$1 "
	consulta += " and extract(year from fecha)=" + miperiodo + " and  "
	consulta += " (substring(cuenta,1,1)='4'  or "
	consulta += " substring(cuenta,1,1)='5'  or "
	consulta += " substring(cuenta,1,1)='6'  or "
	consulta += " substring(cuenta,1,1)='7')"
	consulta += "  group by fecha,cuenta  "

	consulta += ""
	listadoDatos := []datosresumen{}
	listadobalancedeprueba := []balancedeprueba{}

	err1 := db.Select(&listadoDatos, consulta,
		FechaFinal)
	if err1 != nil {
		panic(err1.Error())
	}

	listadoCuenta := []plandecuentaempresa{}
	err2 := db.Select(&listadoCuenta, "select * from plandecuentaempresa where nivel=$1 order by codigo", "A")
	if err2 != nil {
		panic(err2.Error())
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
		fmt.Println("TCUENTA BALANCE " + x.Codigo)

		if nivelInicial == "" {
			nivelInicial = x.Nivel
		}

		var miBalance balancedeprueba
		miBalance = sumaCuenta(x, listadoDatos, FechaInicial, FechaFinal)
		if miBalance.Anterior == "0.00" && miBalance.Debito == "0.00" && miBalance.Credito == "0.00" && miBalance.Saldo == "0.00" {
			fmt.Println("TCUENTA BALANCE VACIA " + x.Codigo)
		} else {
			fmt.Println("TCUENTA BALANCE LLENA" + x.Codigo)

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

	if Saldo4-Saldo5-Saldo6-Saldo7 > 0 {
		//Replace mcrecom With Abs(saldo1-saldo2-saldo6-saldo7)
		miBalanceFinal.Credito = FormatoFlotante(math.Abs(Saldo4 - Saldo5 - Saldo6 - Saldo7))
	} else {
		miBalanceFinal.Debito = FormatoFlotante(math.Abs(Saldo4 - Saldo5 - Saldo6 - Saldo7))

	}

	miBalanceFinal.Descripcion = "TOTALES"
	miBalanceFinal.SiFinal = "SI"
	miBalanceFinal.Anterior = FormatoFlotante(anteriorFinal)
	//miBalanceFinal.Debito = FormatoFlotante(debitoFinal)
	//miBalanceFinal.Credito = FormatoFlotante(creditoFinal)
	miBalanceFinal.Saldo = FormatoFlotante(saldoFinal)

	//listadobalancedeprueba = append(listadobalancedeprueba, miBalanceFinal)

	return listadobalancedeprueba, miBalanceFinal
}

func ultimodia(year, month int) time.Time {
	if month++; month > 12 {
		month = 1
	}
	t := time.Date(year+1, time.Month(month), 0, 0, 0, 0, 0, time.UTC)
	println("conversion fecha final " + t.Format("2006-01-02"))

	return t
}

func CierrecontableDatos(w http.ResponseWriter, r *http.Request) {
	//db := dbConn()
	var miperiodo = periodoSesion(r)
	mes := mux.Vars(r)["mes"]
	var FechaInicial = miperiodo + "-01-01"
	cuentacierre := TraerParametrosContabilidad().Cuentautilidad

	var FechaFinal time.Time

	if mes == "13" {
		FechaFinal = ultimodia(Entero(miperiodo), Entero("12"))
	} else {
		FechaFinal = ultimodia(Entero(miperiodo), Entero(mes))
	}

	var DocumentoContable string

	DocumentoContable = "99"

	var NumeroContable = mes

	var TerceroContable = ListaEmpresa().Codigo
	var CentroContable = mux.Vars(r)["centro"]

	// borrar comprobante anterior
	//delForm, err := db.Prepare("DELETE from comprobantedetalle WHERE documento=$1 and numero=$2")
	//if err != nil {
	//	panic(err.Error())
	//}
	//delForm.Exec(DocumentoContable, NumeroContable)
	//
	//// borra cabecera anterior
	//
	//delForm1, err := db.Prepare("DELETE from comprobante WHERE documento=$1 and numero=$2 ")
	//if err != nil {
	//	panic(err.Error())
	//}
	//delForm1.Exec(DocumentoContable, NumeroContable)

	listadocuentacierre := []balancedeprueba{}
	listadocuentacierrefinal := balancedeprueba{}
	listadocierrefinal := []listadoCierre{}

	listadocuentacierre, listadocuentacierrefinal = InformeResumenCierre(FechaInicial, FechaFinal.Format("2006-01-02"), miperiodo)
	println("terminada 1")

	var miFilaComprobante int
	miFilaComprobante = 0
	var debitofila = ""
	var creditofila = ""
	miComprobanteDetalle := []comprobantedetalle{}

	var totaldebito float64
	var totalcredito float64
	totaldebito = 0
	totalcredito = 0

	//FacturaInicial=""
	for _, mibalance := range listadocuentacierre {
		//var miUltimo int
		//miUltimo=ultimo+1

		totaldebito += Flotante(mibalance.Debito)
		totalcredito += Flotante(mibalance.Credito)
		creditofila = ""
		debitofila = ""
		if mibalance.Codigo[0:1] == "4" {

			if Flotante(mibalance.Saldo) < 0 {
				creditofila = FormatoFlotante(math.Abs(Flotante(mibalance.Saldo)))
			} else {
				debitofila = FormatoFlotante(math.Abs(Flotante(mibalance.Saldo)))
			}

		} else {
			if Flotante(mibalance.Saldo) < 0 {
				debitofila = FormatoFlotante(math.Abs(Flotante(mibalance.Saldo)))
			} else {
				creditofila = FormatoFlotante(math.Abs(Flotante(mibalance.Saldo)))
			}
		}
		listadocierrefinal = append(listadocierrefinal, listadoCierre{strconv.Itoa(miFilaComprobante + 1), mibalance.Codigo, mibalance.Descripcion, debitofila, creditofila})

		miFilaComprobante++
		miComprobanteDetalle = append(miComprobanteDetalle,
			comprobantedetalle{strconv.Itoa(miFilaComprobante),
				mibalance.Codigo,
				TerceroContable,
				CentroContable,
				strings.TrimSpace(mibalance.Descripcion),
				"",
				debitofila,
				creditofila,
				DocumentoContable,
				NumeroContable,
				FechaFinal,
				FechaFinal, "", ""})
		// Inserta Fila contra

	}
	println("terminada 2")

	miFilaComprobante++
	miComprobanteDetalle = append(miComprobanteDetalle,
		comprobantedetalle{strconv.Itoa(miFilaComprobante),
			cuentacierre,
			TerceroContable,
			CentroContable,
			TraerCuentaConsulta(cuentacierre).Nombre,
			"",
			listadocuentacierrefinal.Debito,
			listadocuentacierrefinal.Credito,
			DocumentoContable,
			NumeroContable,
			FechaFinal,
			FechaFinal, "", ""})

	listadocierrefinal = append(listadocierrefinal, listadoCierre{strconv.Itoa(miFilaComprobante + 1), cuentacierre, TraerCuentaConsulta(cuentacierre).Nombre, listadocuentacierrefinal.Debito, listadocuentacierrefinal.Credito})

	totaldebito += Flotante(listadocuentacierrefinal.Debito)
	totalcredito += Flotante(listadocuentacierrefinal.Credito)
	// cabecera comprobante
	//var miperiodo = periodoSesion(r)
	ComprobanteAgregarGenerar(comprobante{DocumentoContable,
		NumeroContable, FechaFinal,
		FechaFinal,
		miperiodo,
		"",
		"",
		"",
		FormatoFlotante(totaldebito),
		FormatoFlotante(totalcredito),
		"Actualizar",
		miComprobanteDetalle, nil})
	println("terminada 3")

	data, _ := json.Marshal(listadocierrefinal)
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	//}
}

// INICIA INSERTAR COMPROBANTE DE VENTA
func InsertaDetalleCierreContable(miFilaComprobante comprobantedetalle, miComprobante comprobante, miVenta venta) {
	db := dbConn()

	// traer tercero
	miTercero := tercero{}
	err3 := db.Get(&miTercero, "SELECT * FROM tercero where codigo=$1", miVenta.Tercero)
	if err3 != nil {
		log.Fatalln(err3)
	}

	var q string
	q = "insert into comprobantedetalle ("
	q += "Fila,"
	q += "Cuenta,"
	q += "Tercero,"
	q += "Centro,"
	q += "Concepto,"
	q += "Factura,"
	q += "Debito,"
	q += "Credito,"
	q += "Documento,"
	q += "Numero,"
	q += "Fecha,"
	q += "Fechaconsignacion"
	q += " ) values("
	q += parametros(12)
	q += ")"
	log.Println("Cadena SQL " + q)
	insForm, err := db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}

	if len(miFilaComprobante.Debito) > 0 {
		miFilaComprobante.Debito = miFilaComprobante.Debito + ".00"
	}

	if len(miFilaComprobante.Credito) > 0 {
		miFilaComprobante.Credito = miFilaComprobante.Credito + ".00"
	}

	// TERMINA COMPROBANTE GRABAR INSERTAR
	_, err = insForm.Exec(
		miFilaComprobante.Fila,
		miFilaComprobante.Cuenta,
		miVenta.Tercero,
		miVenta.Centro,
		miTercero.Nombre,
		"",
		Flotantedatabase(miFilaComprobante.Debito),
		Flotantedatabase(miFilaComprobante.Credito),
		miComprobante.Documento,
		miComprobante.Numero,
		miComprobante.Fecha,
		miComprobante.Fechaconsignacion)
	if err != nil {
		panic(err)
	}
}

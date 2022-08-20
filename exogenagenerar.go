package main

import (
	"fmt"
	"log"
	"time"
)

type FormatoNit struct {
	Tercero string
}

type FormatoCuenta struct {
	Cuenta string
}

type FormatoConcepto struct {
	Formato  string
	Concepto string
}

type FormatoColumna struct {
	Cuenta  string
	Columna string
	Valor   string
}

type FormatoTotal struct {
	Debito  string
	Credito string
	Saldo   string
	Neto    string
}
type miFormatoExcel struct {
	Concepto        string
	Documento       string
	Codigo          string
	Dv              string
	PrimerApllido   string
	SegundoApellido string
	PrimerNombre    string
	SegundoNombre   string
	Juridica        string
	Direccion       string
	Ciudad          string
	Departamento    string
	Pais            string
	Correo          string
	Columna1        string
	Columna2        string
	Columna3        string
	Columna4        string
	Columna5        string
	Columna6        string
	Columna7        string
	Columna8        string
	Columna9        string
	Columna10       string
	Columna11       string
	Columna12       string
	Columna13       string
	Columna14       string
	Columna15       string
	Columna16       string
	Columna17       string
	Columna18       string
	Columna19       string
	Columna20       string
	Columna21       string
	Columna22       string
	Columna23       string
	Columna24       string
	Columna25       string
	Columna26       string
}

func saldoTerceroFormato(tercero string, fechaInicial string, fechaFinal string, miCuenta string, valor string) float64 {
	layout := "2006-01-02"

	dateFinal, _ := time.Parse(layout, fechaFinal)
	dateinicial, _ := time.Parse("2006-01-02", fechaInicial)

	listadoDatos := []datosresumen{}
	var consulta string
	consulta = ""
	consulta = "select distinct fecha,cuenta,sum(debito)as debito,sum(credito) as credito from comprobantedetalle "
	consulta += " where (fecha>=$1 and fecha<= $2)  "
	consulta += " and Cuenta=$3  "
	consulta += " and tercero=$4 "
	consulta += "group by fecha,cuenta"
	err1 := db.Select(&listadoDatos, consulta,
		dateinicial, dateFinal, miCuenta, tercero)

	log.Println(consulta)

	if err1 != nil {
		panic(err1.Error())
	}

	log.Println("Fecha final" + dateFinal.Format("02/01/2006"))

	var totalanterior float64
	var debitoanterior float64
	var creditoanterior float64
	var debito float64
	var credito float64
	var saldo float64
	var neto float64

	debitoanterior = 0
	creditoanterior = 0
	debito = 0
	credito = 0
	saldo = 0
	neto = 0

	for _, x := range listadoDatos {
		log.Println("cuentadatos : " + miCuenta)
		log.Println("fecha movimiento : " + x.Fecha.String())
		log.Println("cuenta parametro : " + miCuenta)

		log.Println("cuenta cortada : " + miCuenta[0:len(miCuenta)])

		//	if cuenta.Codigo == miCuenta[0:len(cuenta.Codigo)] {

		if x.Fecha.Before(dateinicial) {
			debitoanterior += x.Debito
			creditoanterior += x.Credito
			log.Println("movimiento anterior  : " + x.Fecha.String())
		} else {
			debito += x.Debito
			credito += x.Credito
			log.Println("movimiento mes  : " + x.Fecha.String())
		}
		//}

		//listadobalancedeprueba=append(listadobalancedeprueba, balancedeprueba{x.Fecha,strconv.Itoa(i),x.Cuenta, })
	}

	if "1" == miCuenta[0:1] || "5" == miCuenta[0:1] || "6" == miCuenta[0:1] || "7" == miCuenta[0:1] || "8" == miCuenta[0:1] {
		totalanterior = debitoanterior - creditoanterior
		saldo = totalanterior + debito - credito
	} else {

		totalanterior = creditoanterior - debitoanterior
		saldo = totalanterior + credito - debito

	}

	if "1" == miCuenta[0:1] || "5" == miCuenta[0:1] || "6" == miCuenta[0:1] || "7" == miCuenta[0:1] || "8" == miCuenta[0:1] {

		neto = debito - credito
	} else {

		neto = credito - debito

	}

	log.Println("total anterior cuenta  : " + FormatoFlotante(totalanterior))

	log.Println("calculo cuenta anterior " + miCuenta)
	log.Println("total anterior " + CadenaFlotante(saldo))
	log.Println("valor " + valor)
	log.Println("debito " + CadenaFlotante(debito))

	switch valor {
	case "DEBITO":
		//	fmt.Println("one")
		return debito
	case "CREDITO":
		return credito
	case "SALDO":
		return saldo
	case "NETO":
		return neto
		//fmt.Println("three")

	}
	return 0
	//return totalSaldo

}

// sumar columna por tercero
func saldoTercero(tercero string, fechaInicial string, fechaFinal string, concepto string, formato string, columna string) float64 {

	//fmt.Println(	cadenacuenta)
	var consultaNit = "select cuenta,columna,valor from exogena "
	consultaNit += " where formato=$1 and concepto=$2 and columna=$3"

	listacuentavalor := []FormatoColumna{}
	err2 := db.Select(&listacuentavalor, consultaNit, formato, concepto, columna)
	if err2 != nil {
		fmt.Println(err2)
	}
	var valorcolumna float64
	valorcolumna = 0

	for _, mivalor := range listacuentavalor {
		println("lista valor cuenta  " + mivalor.Cuenta + " " + mivalor.Valor)

		println("lista valor fechas  " + fechaInicial + " " + fechaFinal)
		valorcolumna += saldoTerceroFormato(tercero, fechaInicial, fechaFinal, mivalor.Cuenta, mivalor.Valor)

	}
	return valorcolumna
}

func generarformato(formato string, SiConcepto bool) []miFormatoExcel {

	var periodo = ListaEmpresa().Periodo
	var fechaInicial = periodo + "-01-01"
	var fechaFinal = periodo + "-12-31"
	// formato concepto
	listaconcepto := []FormatoConcepto{}
	var consulta = "select distinct concepto,formato from exogena where formato=$1"
	err2 := db.Select(&listaconcepto, consulta, formato)
	if err2 != nil {
		fmt.Println(err2)
	}
	// cuentas usadas
	listaformato := []miFormatoExcel{}
	listacuenta := []exogenaeditar{}
	fmt.Println("Inicio error")
	fmt.Println(formato)

	//var consulta = ""
	consulta = " select "
	consulta += "exogena.cuenta, "
	consulta += "exogena.formato, "
	consulta += "exogenaformato.codigo||'-'||exogenaformato.nombre as formatonombre, "

	if SiConcepto == false {
		consulta += "   exogena.concepto, "
		consulta += "'' as conceptonombre, "

	} else {
		consulta += "exogena.concepto, "
		consulta += "exogenaconcepto.concepto||'-'||exogenaconcepto.nombre as conceptonombre, "

	}

	//	consulta += "exogena.columna, "
	//	consulta += "exogenacolumna.columna||'-'||exogenaconcepto.nombre as columnanombre, "
	consulta += "exogena.valor "
	consulta += "from exogena "

	consulta += " inner join exogenaformato "
	consulta += " on exogena.formato= exogenaformato.codigo"
	if SiConcepto == false {

	} else {

		consulta += " inner join exogenaconcepto "
		consulta += " on exogena.formato=exogenaconcepto.formato and exogena.concepto= exogenaconcepto.concepto"
	}

	consulta += " inner join exogenacolumna "
	consulta += " on exogena.columna=exogenacolumna.columna and exogena.formato= exogenacolumna.formato"
	consulta += " where exogena.formato=$1 "
	fmt.Println(consulta)

	err2 = db.Select(&listacuenta, consulta, formato)

	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println("fin erro ")
	var cadenacuenta = ""
	// terceros
	for _, x := range listacuenta {

		if cadenacuenta == "" {
			cadenacuenta += " (comprobantedetalle.cuenta='" + x.Cuenta + "' "
		} else {
			cadenacuenta += " or comprobantedetalle.cuenta='" + x.Cuenta + "' "
		}
	}
	cadenacuenta = cadenacuenta + " ) "

	fmt.Println(cadenacuenta)
	var consultaNit = "select distinct tercero from comprobantedetalle "
	consultaNit += " where " + cadenacuenta + " and (fecha>=$1 and fecha<= $2)"

	listanit := []FormatoNit{}
	err2 = db.Select(&listanit, consultaNit, fechaInicial, fechaFinal)
	fmt.Println(consultaNit)
	if err2 != nil {
		fmt.Println(err2)
	}

	var columna1 float64
	var columna2 float64
	var columna3 float64
	var columna4 float64
	var columna5 float64
	var columna6 float64
	var columna7 float64
	var columna8 float64
	var columna9 float64
	var columna10 float64
	var columna11 float64
	var columna12 float64
	var columna13 float64
	var columna14 float64
	var columna15 float64
	var columna16 float64
	var columna17 float64
	var columna18 float64
	var columna19 float64
	var columna20 float64
	var columna21 float64
	var columna22 float64
	var columna23 float64
	var columna24 float64
	var columna25 float64
	var columna26 float64

	columna1 = 0
	columna2 = 0
	columna3 = 0
	columna4 = 0
	columna5 = 0
	columna6 = 0
	columna7 = 0
	columna8 = 0
	columna9 = 0
	columna10 = 0
	columna11 = 0
	columna12 = 0
	columna13 = 0
	columna14 = 0
	columna15 = 0
	columna16 = 0
	columna17 = 0
	columna18 = 0
	columna19 = 0
	columna20 = 0
	columna21 = 0
	columna22 = 0
	columna23 = 0
	columna24 = 0
	columna25 = 0
	columna26 = 0

	mitercero := tercero{}

	for _, miConcepto := range listaconcepto {
		println("concepto ---- " + miConcepto.Concepto)
		for _, miTercerolista := range listanit {
			mitercero = TraerTercero(miTercerolista.Tercero)
			var miDepartamento = Subcadena(mitercero.Ciudad, 0, 2)
			var miCiudad = Subcadena(mitercero.Ciudad, 2, 5)

			println("Tercero ------ " + miTercerolista.Tercero)
			columna1 = saldoTercero(miTercerolista.Tercero, fechaInicial, fechaFinal, miConcepto.Concepto, formato, "1")
			columna2 = saldoTercero(miTercerolista.Tercero, fechaInicial, fechaFinal, miConcepto.Concepto, formato, "2")
			columna3 = saldoTercero(miTercerolista.Tercero, fechaInicial, fechaFinal, miConcepto.Concepto, formato, "3")
			columna4 = saldoTercero(miTercerolista.Tercero, fechaInicial, fechaFinal, miConcepto.Concepto, formato, "4")
			columna5 = saldoTercero(miTercerolista.Tercero, fechaInicial, fechaFinal, miConcepto.Concepto, formato, "5")
			columna6 = saldoTercero(miTercerolista.Tercero, fechaInicial, fechaFinal, miConcepto.Concepto, formato, "6")
			columna7 = saldoTercero(miTercerolista.Tercero, fechaInicial, fechaFinal, miConcepto.Concepto, formato, "7")
			columna8 = saldoTercero(miTercerolista.Tercero, fechaInicial, fechaFinal, miConcepto.Concepto, formato, "8")
			columna9 = saldoTercero(miTercerolista.Tercero, fechaInicial, fechaFinal, miConcepto.Concepto, formato, "9")
			columna10 = saldoTercero(miTercerolista.Tercero, fechaInicial, fechaFinal, miConcepto.Concepto, formato, "10")
			columna11 = saldoTercero(miTercerolista.Tercero, fechaInicial, fechaFinal, miConcepto.Concepto, formato, "11")
			columna12 = saldoTercero(miTercerolista.Tercero, fechaInicial, fechaFinal, miConcepto.Concepto, formato, "12")
			columna13 = saldoTercero(miTercerolista.Tercero, fechaInicial, fechaFinal, miConcepto.Concepto, formato, "13")
			columna14 = saldoTercero(miTercerolista.Tercero, fechaInicial, fechaFinal, miConcepto.Concepto, formato, "14")
			columna15 = saldoTercero(miTercerolista.Tercero, fechaInicial, fechaFinal, miConcepto.Concepto, formato, "15")
			columna16 = saldoTercero(miTercerolista.Tercero, fechaInicial, fechaFinal, miConcepto.Concepto, formato, "16")
			columna17 = saldoTercero(miTercerolista.Tercero, fechaInicial, fechaFinal, miConcepto.Concepto, formato, "17")
			columna18 = saldoTercero(miTercerolista.Tercero, fechaInicial, fechaFinal, miConcepto.Concepto, formato, "18")
			columna19 = saldoTercero(miTercerolista.Tercero, fechaInicial, fechaFinal, miConcepto.Concepto, formato, "19")
			columna20 = saldoTercero(miTercerolista.Tercero, fechaInicial, fechaFinal, miConcepto.Concepto, formato, "20")
			columna21 = saldoTercero(miTercerolista.Tercero, fechaInicial, fechaFinal, miConcepto.Concepto, formato, "21")
			columna22 = saldoTercero(miTercerolista.Tercero, fechaInicial, fechaFinal, miConcepto.Concepto, formato, "22")
			columna23 = saldoTercero(miTercerolista.Tercero, fechaInicial, fechaFinal, miConcepto.Concepto, formato, "23")
			columna24 = saldoTercero(miTercerolista.Tercero, fechaInicial, fechaFinal, miConcepto.Concepto, formato, "24")
			columna25 = saldoTercero(miTercerolista.Tercero, fechaInicial, fechaFinal, miConcepto.Concepto, formato, "25")
			columna26 = saldoTercero(miTercerolista.Tercero, fechaInicial, fechaFinal, miConcepto.Concepto, formato, "26")

			//	columna1 = 1
			listaformato = append(listaformato, miFormatoExcel{
				miConcepto.Concepto,
				mitercero.Documento,
				mitercero.Codigo,
				mitercero.Dv,
				mitercero.PrimerApellido,
				mitercero.SegundoApellido,
				mitercero.PrimerNombre,
				mitercero.SegundoNombre,
				mitercero.Juridica,
				mitercero.Direccion,
				miCiudad,
				miDepartamento,
				"169",
				mitercero.Email1,
				FormatoFlotanteEntero(columna1),
				FormatoFlotanteEntero(columna2),
				FormatoFlotanteEntero(columna3),
				FormatoFlotanteEntero(columna4),
				FormatoFlotanteEntero(columna5),
				FormatoFlotanteEntero(columna6),
				FormatoFlotanteEntero(columna7),
				FormatoFlotanteEntero(columna8),
				FormatoFlotanteEntero(columna9),
				FormatoFlotanteEntero(columna10),
				FormatoFlotanteEntero(columna11),
				FormatoFlotanteEntero(columna12),
				FormatoFlotanteEntero(columna13),
				FormatoFlotanteEntero(columna14),
				FormatoFlotanteEntero(columna15),
				FormatoFlotanteEntero(columna16),
				FormatoFlotanteEntero(columna17),
				FormatoFlotanteEntero(columna18),
				FormatoFlotanteEntero(columna19),
				FormatoFlotanteEntero(columna20),
				FormatoFlotanteEntero(columna21),
				FormatoFlotanteEntero(columna22),
				FormatoFlotanteEntero(columna23),
				FormatoFlotanteEntero(columna24),
				FormatoFlotanteEntero(columna25),
				FormatoFlotanteEntero(columna26),
			})
		}

	}
	return listaformato
}

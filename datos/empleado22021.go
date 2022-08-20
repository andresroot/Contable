package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
	"log"
	"net/http"
	"time"
)

// INICIA CERTIFICADO DE PAGO 220 PDF
func Empleado22021Pdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	dateinicial, err := time.Parse("2006-01-02", mux.Vars(r)["fechainicial"])
	datefinal, err := time.Parse("2006-01-02", mux.Vars(r)["fechafinal"])
	datecertificado, err := time.Parse("2006-01-02", mux.Vars(r)["fechaexpedicion"])

	miEmpleado := empleado{}
	var miEmpresa empresa = ListaEmpresa()
	//var c  ciudad=TraerCiudad(e.Ciudad)
	err = db.Get(&miEmpleado, "SELECT * FROM empleado where codigo=$1", Codigo)
	if err != nil {
		log.Fatalln(err)
	}
	var buf bytes.Buffer
	var err1 error
	pdf := gofpdf.New("P", "mm", "Letter", cnFontDir)
	ene := pdf.UnicodeTranslatorFromDescriptor("")

	pdf.SetX(21)
	pdf.AliasNbPages("")
	pdf.AddPage()

	pdf.Ln(2)
	pdf.SetFont("Arial", "", 20)
	//pdf.SetDrawColor(84,153,199)
	pdf.SetDrawColor(95, 119, 146)
	pdf.SetTextColor(0, 0, 0)

	// LINEA HORIZONTAL
	pdf.Line(21, 14, 201, 14)
	pdf.Line(21, 26, 201, 26)
	pdf.Line(21, 33, 201, 33)
	pdf.Line(21, 55, 201, 55)
	pdf.Line(21, 69, 201, 69)
	pdf.Line(30, 45, 201, 45)

	// LINEA DE PRUEBA
	//pdf.SetDrawColor(255, 0, 0)
	//pdf.Line(130, 69, 130, 187)
	//pdf.SetTextColor(0, 0, 0)
	// TERMINA LINEA DE PRUEBA

	pdf.SetDrawColor(84, 153, 199)
	pdf.Line(47, 14, 47, 26)
	pdf.Line(175, 14, 175, 26)
	pdf.Line(21, 14, 21, 79)
	pdf.Line(201, 14, 201, 79)
	pdf.Line(111, 26, 111, 33)
	pdf.Line(30, 33, 30, 55)
	pdf.Line(30, 55, 30, 69)
	pdf.Line(74, 69, 74, 79)
	pdf.Line(105, 69, 105, 79)
	pdf.Line(163, 69, 163, 79)
	pdf.Line(178, 69, 178, 79)

	// RETENEDOR
	pdf.Line(85, 40, 85, 45)
	pdf.Line(100, 40, 100, 45)
	pdf.Line(125, 40, 125, 45)
	pdf.Line(150, 40, 150, 45)
	pdf.Line(175, 40, 175, 45)

	// TRABAJADOR
	pdf.Line(43, 64, 43, 69)
	pdf.Line(100, 64, 100, 69)
	pdf.Line(125, 64, 125, 69)
	pdf.Line(150, 64, 150, 69)
	pdf.Line(175, 64, 175, 69)

	pdf.SetY(15)
	pdf.SetX(185)
	pdf.Image(imageFile("Dian.png"), 25, 15, 20, 10, false,
		"", 0, "")
	pdf.SetFont("Arial", "B", 10)
	pdf.SetY(15)
	pdf.SetX(40)
	pdf.CellFormat(140, 4, ene("Certificado de Ingresos y Retenciones Por Rentas de Trabajo y Pensiones"), "0", 0, "C", false, 0, "")
	pdf.SetY(20)
	pdf.SetX(40)
	pdf.CellFormat(140, 4, ene("Año Gravable 2021"), "0", 0, "C", false, 0, "")
	pdf.SetY(15)
	pdf.SetX(177)
	pdf.SetFont("Arial", "", 20)

	pdf.SetTextColor(253, 254, 254)
	pdf.SetFillColor(56, 100, 146)
	pdf.CellFormat(22, 10, "220", "0", 0,
		"C", true, 0, "")
	pdf.SetTextColor(0, 0, 0)
	pdf.Ln(-1)
	pdf.SetX(23)
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(95, 8, ene("Antes de diligenciar este formulario lea las instrucciones"), "0", 0,
		"", false, 0, "")
	pdf.CellFormat(89, 8, " 4. Formulario No."+" "+miEmpleado.Codigo, "0", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.SetFont("Arial", "", 7)
	pdf.TransformBegin()
	pdf.TransformRotate(90, 28.5, 42.5)
	pdf.CellFormat(15, 15, "Retenedor", "0", 0,
		"", false, 0, "")
	pdf.TransformEnd()
	pdf.SetX(50)
	pdf.SetY(35)
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(100, 4, " 5. Numero de Identificacion Tributario", "0", 0, "C",
		false, 0, "")
	pdf.SetX(80)
	pdf.CellFormat(20, 4, "6. Dv", "0", 0, "C",
		false, 0, "")
	pdf.SetY(40)
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(70, 4, Coma(miEmpresa.Codigo), "0", 0, "C",
		false, 0, "")
	pdf.SetX(35)
	pdf.CellFormat(110, 4, miEmpresa.Dv, "0", 0, "C",
		false, 0, "")
	pdf.SetFont("Arial", "", 8)
	pdf.SetY(35)
	pdf.CellFormat(280, 4, " 7. Primer Apellido  8. Segundo Apellido  9. Primer Nombre  10. Otros Nombres", "0", 0, "C",
		false, 0, "")
	pdf.SetY(40)
	pdf.SetX(101)
	pdf.SetFont("Arial", "", 10)
	if miEmpresa.Tipo == "2" {
		pdf.CellFormat(40, 4, miEmpresa.Primernombre, "0", 0,
			"", false, 0, "")
		pdf.SetX(126)
		pdf.CellFormat(40, 4, miEmpresa.Segundonombre, "0", 0,
			"", false, 0, "")
		pdf.SetX(151)
		pdf.SetFont("Arial", "", 10)
		pdf.CellFormat(40, 4, miEmpresa.Primerapellido, "0", 0,
			"", false, 0, "")
		pdf.SetX(176)
		pdf.CellFormat(40, 4, miEmpresa.Segundoapellido, "0", 0,
			"", false, 0, "")

	}

	pdf.SetFont("Arial", "", 8)
	pdf.SetY(45)
	pdf.SetX(37)
	pdf.CellFormat(80, 4, " 11. Razon Social", "0", 0, "",
		false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.SetY(50)
	pdf.SetX(40)

	if miEmpresa.Tipo == "1" {
		pdf.CellFormat(120, 4, miEmpresa.Juridica, "0", 0,
			"", false, 0, "")
	}

	pdf.SetX(21)
	pdf.SetFont("Arial", "", 7)
	pdf.TransformBegin()
	pdf.TransformRotate(90, 29.5, 60.5)
	pdf.CellFormat(15, 15, "Trabajador", "0", 0,
		"", false, 0, "")
	pdf.TransformEnd()
	pdf.SetY(57)
	pdf.SetX(32)
	pdf.CellFormat(10, 4, " 24. Tipo", "0", 0, "R",
		false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.SetY(62)
	pdf.SetX(30)
	pdf.CellFormat(10, 4, "13", "0", 0, "C",
		false, 0, "")
	pdf.SetX(150)
	pdf.SetY(57)
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(120, 4, " 25. Numero de Identificacion Tributario", "0", 0, "C",
		false, 0, "")
	pdf.SetY(62)
	pdf.SetX(21)
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(70, 4, Coma(miEmpleado.Codigo), "0", 0, "C",
		false, 0, "")
	pdf.SetFont("Arial", "", 8)
	pdf.SetY(57)
	pdf.CellFormat(280, 4, "26. Primer Apellido  27. Segundo Apellido  28. Primer Nombre  29. Otros Nombres", "0", 0, "C",
		false, 0, "")
	pdf.SetY(62)
	pdf.SetX(101)
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(40, 4, miEmpleado.PrimerNombre, "0", 0,
		"", false, 0, "")
	pdf.SetX(126)
	pdf.CellFormat(40, 4, miEmpleado.SegundoNombre, "0", 0,
		"", false, 0, "")
	pdf.SetX(151)
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(40, 4, miEmpleado.PrimerApellido, "0", 0,
		"", false, 0, "")
	pdf.SetX(176)
	pdf.CellFormat(40, 4, miEmpleado.SegundoApellido, "0", 0,
		"", false, 0, "")

	pdf.SetY(70)
	pdf.SetX(23)
	pdf.SetFont("Arial", "", 8)
	pdf.CellFormat(40, 4, " Periodo de Certificacion", "0", 0,
		"", false, 0, "")
	pdf.SetX(75)
	pdf.CellFormat(40, 4, "32. Fecha Expedicion", "0", 0,
		"", false, 0, "")
	pdf.SetX(107)
	pdf.CellFormat(40, 4, "33. Lugar Donde Se Practico la Retencion", "0", 0,
		"", false, 0, "")
	pdf.SetX(165)
	pdf.CellFormat(40, 4, "34. Cod.", "0", 0,
		"", false, 0, "")
	pdf.SetX(178)
	pdf.CellFormat(40, 4, "35. Cod. Ciudad", "0", 0,
		"", false, 0, "")

	pdf.SetY(75)
	pdf.SetX(21)
	pdf.SetFont("Arial", "", 9)
	pdf.CellFormat(40, 4, "30. de "+dateinicial.Format("02/01/2006")+" A 31. "+datefinal.Format("02/01/2006"), "0", 0,
		"", false, 0, "")
	pdf.SetX(70)
	pdf.CellFormat(40, 4, datecertificado.Format("02/01/2006"), "0", 0,
		"C", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(40, 4, TraerCiudad(miEmpleado.Ciudad).NombreCiudad+"-"+TraerCiudad(miEmpleado.Ciudad).NombreDepartamento, "0", 0,
		"", false, 0, "")
	pdf.SetX(167)
	pdf.CellFormat(40, 4, TraerCiudad(miEmpleado.Ciudad).CodigoDepartamento, "0", 0,
		"", false, 0, "")
	pdf.SetX(185)
	pdf.CellFormat(40, 4, TraerCiudad(miEmpleado.Ciudad).CodigoCiudad, "0", 0,
		"", false, 0, "")

	pdf.SetFont("Arial", "B", 8)
	pdf.SetFillColor(225, 232, 239)
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(142, 4, "Concepto de los Ingresos", "1", 0,
		"C", true, 0, "")
	pdf.CellFormat(38, 4, "Valor", "1", 0,
		"C", true, 0, "")
	pdf.SetFont("Arial", "", 8)
	pdf.Ln(-1)
	pdf.SetX(21)

	var saldoTrabajado float64
	saldoTrabajado = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Trabajado)

	var saldoTransporte float64
	saldoTransporte = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Transporte)

	var saldoHorasextras float64
	saldoHorasextras = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Horasextras)

	var saldoIncapacidades float64
	saldoIncapacidades = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Incapacidades)

	var saldoLicencias float64
	saldoLicencias = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Licencias)

	var saldoBonificaciones float64
	saldoBonificaciones = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Bonificaciones)

	var saldoAuxilios float64
	saldoAuxilios = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Auxilios)

	var saldoHuelgas float64
	saldoHuelgas = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Huelgas)

	var saldoConceptos float64
	saldoConceptos = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Conceptos)

	var saldoCompensaciones float64
	saldoCompensaciones = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Compensaciones)

	var saldoSostenimiento float64
	saldoSostenimiento = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Sostenimiento)

	var saldoTeletrabajo float64
	saldoTeletrabajo = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Teletrabajo)

	var saldoIndemnizaciones float64
	saldoIndemnizaciones = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Indemnizaciones)

	var saldo36 string
	saldo36 = FormatoFlotanteEntero(saldoTrabajado + saldoTransporte + saldoHorasextras +
		saldoHorasextras + saldoIncapacidades + saldoLicencias + saldoBonificaciones +
		saldoAuxilios + saldoHuelgas + saldoConceptos + saldoCompensaciones + saldoSostenimiento +
		saldoTeletrabajo + saldoIndemnizaciones)

	if saldo36 == "0" {
		saldo36 = ""
	}

	pdf.CellFormat(142, 4, "Pagos por salarios o emolumentos eclesiasticos", "", 0,
		"", false, 0, "")
	pdf.CellFormat(5, 4, "36", "", 0,
		"", false, 0, "")
	pdf.CellFormat(33, 4, saldo36, "", 0,
		"R", false, 0, "")
	pdf.SetFillColor(225, 232, 239)
	pdf.Ln(-1)
	pdf.SetX(21)

	var saldoBonos float64
	saldoBonos = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Bonos)

	var saldo37 string
	saldo37 = FormatoFlotanteEntero(saldoBonos)

	if saldo37 == "0" {
		saldo37 = ""
	}

	pdf.CellFormat(142, 4, "Pagos realizados con bonos electronicos o de papel de servicio, cheques, tarjetas, vales, etc", "", 0,
		"", true, 0, "")
	pdf.CellFormat(5, 4, "37", "", 0,
		"", true, 0, "")
	pdf.CellFormat(33, 4, saldo37, "", 0,
		"R", true, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)

	var saldoHonorarios float64
	saldoHonorarios = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Honorariogasto)

	var saldo38 string
	saldo38 = FormatoFlotanteEntero(saldoHonorarios)

	if saldo38 == "0" {
		saldo38 = ""
	}

	pdf.CellFormat(142, 4, "Pagos por Honorarios", "", 0,
		"", false, 0, "")
	pdf.CellFormat(5, 4, "38", "", 0,
		"", false, 0, "")
	pdf.CellFormat(33, 4, saldo38, "", 0,
		"R", false, 0, "")
	pdf.SetFillColor(225, 232, 239)
	pdf.Ln(-1)
	pdf.SetX(21)

	var saldoServicios float64
	saldoServicios = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Serviciogasto)

	var saldo39 string
	saldo39 = FormatoFlotanteEntero(saldoServicios)

	if saldo39 == "0" {
		saldo39 = ""
	}

	pdf.CellFormat(142, 4, "Pagos por Servicios", "", 0,
		"", true, 0, "")
	pdf.CellFormat(5, 4, "39", "", 0,
		"", true, 0, "")
	pdf.CellFormat(33, 4, saldo39, "", 0,
		"R", true, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)

	var saldoComisiones float64
	saldoComisiones = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Comisiones)

	var saldo40 string
	saldo40 = FormatoFlotanteEntero(saldoComisiones)

	if saldo40 == "0" {
		saldo40 = ""
	}

	pdf.CellFormat(142, 4, "Pagos por Comisiones", "", 0,
		"", false, 0, "")
	pdf.CellFormat(5, 4, "40", "", 0,
		"", false, 0, "")
	pdf.CellFormat(33, 4, saldo40, "", 0,
		"R", false, 0, "")
	pdf.SetFillColor(225, 232, 239)
	pdf.Ln(-1)
	pdf.SetX(21)

	var saldoPrima float64
	saldoPrima = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Prima)

	var saldoVacaciones float64
	saldoVacaciones = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Vacaciones)

	var saldoDotaciones float64
	saldoDotaciones = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Dotaciones)

	var saldo41 string
	saldo41 = FormatoFlotanteEntero(saldoPrima + saldoVacaciones + saldoDotaciones)

	if saldo41 == "0" {
		saldo41 = ""
	}

	pdf.CellFormat(142, 4, "Pagos Por Prestaciones Sociales ", "", 0,
		"", true, 0, "")
	pdf.CellFormat(5, 4, "41", "", 0,
		"", true, 0, "")
	pdf.CellFormat(33, 4, saldo41, "", 0,
		"R", true, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)

	var saldoViaticos float64
	saldoViaticos = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Viaticos)

	var saldo42 string
	saldo42 = FormatoFlotanteEntero(saldoViaticos)

	if saldo42 == "0" {
		saldo42 = ""
	}

	pdf.CellFormat(142, 4, "Pagos por Viaticos", "", 0,
		"", false, 0, "")
	pdf.CellFormat(5, 4, "42", "", 0,
		"", false, 0, "")
	pdf.CellFormat(33, 4, "", "", 0,
		"R", false, 0, "")
	pdf.SetFillColor(225, 232, 239)
	pdf.Ln(-1)
	pdf.SetX(21)

	var saldo43 string

	if saldo43 == "0" {
		saldo43 = ""
	}

	pdf.CellFormat(142, 4, "Pagos por Gastos de Representacion", "", 0,
		"", true, 0, "")
	pdf.CellFormat(5, 4, "43", "", 0,
		"", true, 0, "")
	pdf.CellFormat(33, 4, saldo43, "", 0,
		"R", true, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)

	var saldo44 string

	if saldo44 == "0" {
		saldo44 = ""
	}

	pdf.CellFormat(142, 4, "Pagos por Compensaciones por el Trabajo Asociado Cooperativo", "", 0,
		"", false, 0, "")
	pdf.CellFormat(5, 4, "44", "", 0,
		"", false, 0, "")
	pdf.CellFormat(33, 4, saldo44, "", 0,
		"R", false, 0, "")
	pdf.SetFillColor(225, 232, 239)
	pdf.Ln(-1)
	pdf.SetX(21)

	var saldo45 string

	if saldo45 == "0" {
		saldo45 = ""
	}

	pdf.CellFormat(142, 4, "Otros Pagos", "", 0,
		"", true, 0, "")
	pdf.CellFormat(5, 4, "45", "", 0,
		"", true, 0, "")
	pdf.CellFormat(33, 4, saldo45, "", 0,
		"R", true, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)

	var saldoCesantias float64
	saldoCesantias = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Cesantias)

	var saldoIntereses float64
	saldoIntereses = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Intereses)

	var saldo46 string
	saldo46 = FormatoFlotanteEntero(saldoCesantias + saldoIntereses)

	if saldo46 == "0" {
		saldo46 = ""
	}

	pdf.CellFormat(142, 4, "Cesantias e Intereses de Cesantias Efectivamente Pagados al Empleado", "", 0,
		"", false, 0, "")
	pdf.CellFormat(5, 4, "46", "", 0,
		"", false, 0, "")
	pdf.CellFormat(33, 4, saldo46, "", 0,
		"R", false, 0, "")
	pdf.SetFillColor(225, 232, 239)
	pdf.Ln(-1)
	pdf.SetX(21)

	var saldoFondocesantias float64
	saldoFondocesantias = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Cesantiascxp)

	var saldo47 string
	saldo47 = FormatoFlotanteEntero(saldoFondocesantias)

	if saldo47 == "0" {
		saldo47 = ""
	}

	pdf.CellFormat(142, 4, "Cesantias Consignadas al Fondo de Cesantias", "", 0,
		"", true, 0, "")
	pdf.CellFormat(5, 4, "47", "", 0,
		"", true, 0, "")
	pdf.CellFormat(33, 4, saldo47, "", 0,
		"R", true, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)

	var saldo48 string

	if saldo48 == "0" {
		saldo48 = ""
	}

	pdf.CellFormat(142, 4, "Pension de Jubilacion, Vejes o Invalidez", "", 0,
		"", false, 0, "")
	pdf.CellFormat(5, 4, "48", "", 0,
		"", false, 0, "")
	pdf.CellFormat(33, 4, saldo48, "", 0,
		"R", false, 0, "")
	pdf.SetFillColor(225, 232, 239)
	pdf.Ln(-1)
	pdf.SetX(21)

	var saldo49 string
	saldo49 = FormatoFlotanteEntero(Flotante(saldo36) + Flotante(saldo37) + Flotante(saldo38) +
		Flotante(saldo39) + Flotante(saldo40) + Flotante(saldo41) + Flotante(saldo42) +
		Flotante(saldo43) + Flotante(saldo44) + Flotante(saldo45) + Flotante(saldo46) +
		Flotante(saldo47) + Flotante(saldo48))

	if saldo49 == "0" {
		saldo49 = ""
	}

	pdf.CellFormat(142, 4, "Total de Ingresos Brutos (Suma Casillas 36 a 48)", "", 0,
		"", true, 0, "")
	pdf.CellFormat(5, 4, "49", "", 0,
		"", true, 0, "")
	pdf.CellFormat(33, 4, saldo49, "", 0,
		"R", true, 0, "")
	pdf.SetFillColor(225, 232, 239)
	pdf.SetFont("Arial", "B", 8)
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(142, 4, "Concepto de los Aportes", "1", 0,
		"C", true, 0, "")
	pdf.CellFormat(38, 4, "Valor", "1", 0,
		"C", true, 0, "")
	pdf.SetFont("Arial", "", 8)
	pdf.Ln(-1)
	pdf.SetX(21)

	var saldoSalud float64
	saldoSalud = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Salud)

	var saldo50 string
	saldo50 = FormatoFlotanteEntero(saldoSalud)

	if saldo50 == "0" {
		saldo50 = ""
	}

	pdf.CellFormat(142, 4, "Aportes Obligatorios por Salud a Cargo del Trabajador", "", 0,
		"", false, 0, "")
	pdf.CellFormat(5, 4, "50", "", 0,
		"", false, 0, "")
	pdf.CellFormat(33, 4, saldo50, "", 0,
		"R", false, 0, "")
	pdf.SetFillColor(225, 232, 239)
	pdf.Ln(-1)
	pdf.SetX(21)

	var saldoPension float64
	saldoPension = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Pension)

	var saldoSolidaridad float64
	saldoSolidaridad = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Solidaridad)

	var saldo51 string
	saldo51 = FormatoFlotanteEntero(saldoPension + saldoSolidaridad)

	if saldo51 == "0" {
		saldo51 = ""
	}

	pdf.CellFormat(142, 4, "Aportes Obligatorios a Fondo de Pensiones y Solidaridad Pensional a Cargo del Trabajador", "", 0,
		"", true, 0, "")
	pdf.CellFormat(5, 4, "51", "", 0,
		"", true, 0, "")
	pdf.CellFormat(33, 4, saldo51, "", 0,
		"R", true, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)

	var saldoPensionrais float64
	saldoPensionrais = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Pensionrais)

	var saldo52 string
	saldo52 = FormatoFlotanteEntero(saldoPensionrais)

	if saldo52 == "0" {
		saldo52 = ""
	}

	pdf.CellFormat(142, 4, "Aportes Obligatorios a Fondo de Pensiones y Solidaridad Pensional y Apaortes Voluntarios - RAIS", "", 0,
		"", false, 0, "")
	pdf.CellFormat(5, 4, "52", "", 0,
		"", false, 0, "")
	pdf.CellFormat(33, 4, saldo52, "", 0,
		"R", false, 0, "")
	pdf.SetFillColor(225, 232, 239)
	pdf.Ln(-1)
	pdf.SetX(21)

	var saldoPensionvoluntaria float64
	saldoPensionvoluntaria = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Pensionvoluntaria)

	var saldo53 string
	saldo53 = FormatoFlotanteEntero(saldoPensionvoluntaria)

	if saldo53 == "0" {
		saldo53 = ""
	}

	pdf.CellFormat(142, 4, "Aportes Voluntarios al Fondo de Pensiones", "", 0,
		"", true, 0, "")
	pdf.CellFormat(5, 4, "53", "", 0,
		"", true, 0, "")
	pdf.CellFormat(33, 4, saldo53, "", 0,
		"R", true, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)

	var saldoAfc float64
	saldoAfc = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Afc)

	var saldo54 string
	saldo54 = FormatoFlotanteEntero(saldoAfc)

	if saldo54 == "0" {
		saldo54 = ""
	}

	pdf.CellFormat(142, 4, "Aportes a Cuentas AFC", "", 0,
		"", false, 0, "")
	pdf.CellFormat(5, 4, "54", "", 0,
		"", false, 0, "")
	pdf.CellFormat(33, 4, saldo54, "", 0,
		"R", false, 0, "")
	pdf.SetTextColor(253, 254, 254)
	pdf.SetFillColor(56, 100, 146)
	pdf.Ln(-1)
	pdf.SetX(21)

	var saldoRetencion float64
	saldoRetencion = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Retencion)

	var saldoRethonorario float64
	saldoRethonorario = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Rethonorario)

	var saldoRetservicio float64
	saldoRetservicio = saldoEmpleado(Codigo, dateinicial, datefinal, TraerParametrosContabilidad().Retservicio)

	var saldo55 string
	saldo55 = FormatoFlotanteEntero(saldoRetencion + saldoRethonorario + saldoRetservicio)

	if saldo55 == "0" {
		saldo55 = ""
	}

	pdf.CellFormat(142, 4, "Valor de la Retencion en la Fuente Por Renta de Trabajo y Pensiones", "1", 0,
		"", true, 0, "")
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(5, 4, "55", "1", 0,
		"", false, 0, "")
	pdf.CellFormat(33, 4, saldo55, "1", 0,
		"R", false, 0, "")

	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(180, 4, "Nombre del Pagador o Agente Retenedor:", "0", 0,
		"", false, 0, "")
	pdf.SetX(91)
	pdf.CellFormat(180, 4, miEmpresa.RepresentanteNombre, "0", 0,
		"", false, 0, "")
	pdf.SetTextColor(253, 254, 254)
	pdf.SetFillColor(56, 100, 146)
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(180, 4, "Datos a Cargo del Trabajador o Pensionado", "", 0,
		"C", true, 0, "")
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFillColor(225, 232, 239)
	pdf.SetFont("Arial", "B", 8)
	pdf.Ln(-1)
	pdf.SetX(21)

	pdf.CellFormat(110, 4, "Concepto de Otros Ingresos", "1", 0,
		"C", true, 0, "")
	pdf.CellFormat(32, 4, "Valor Recibido", "1", 0,
		"C", true, 0, "")
	pdf.CellFormat(38, 4, "Valor Retenido", "1", 0,
		"C", true, 0, "")
	pdf.SetFont("Arial", "", 8)
	pdf.Ln(-1)
	pdf.SetX(21)

	var saldo56 string

	if saldo56 == "0" {
		saldo56 = ""
	}

	var saldo63 string

	if saldo63 == "0" {
		saldo63 = ""
	}

	pdf.CellFormat(110, 4, "Arrendamientos", "", 0,
		"", false, 0, "")
	pdf.CellFormat(5, 4, "56", "1", 0,
		"", false, 0, "")
	pdf.CellFormat(27, 4, saldo56, "", 0,
		"R", false, 0, "")
	pdf.CellFormat(5, 4, "63", "1", 0,
		"", false, 0, "")
	pdf.CellFormat(33, 4, saldo63, "", 0,
		"R", false, 0, "")
	pdf.SetFillColor(225, 232, 239)
	pdf.Ln(-1)
	pdf.SetX(21)

	var saldo57 string

	if saldo57 == "0" {
		saldo57 = ""
	}

	var saldo64 string

	if saldo64 == "0" {
		saldo64 = ""
	}

	pdf.CellFormat(110, 4, "Honorarios, Comisiones y Servicios.", "", 0,
		"", true, 0, "")
	pdf.CellFormat(5, 4, "57", "", 0,
		"", true, 0, "")
	pdf.CellFormat(27, 4, saldo57, "", 0,
		"R", true, 0, "")
	pdf.CellFormat(5, 4, "64", "", 0,
		"", true, 0, "")
	pdf.CellFormat(33, 4, saldo64, "", 0,
		"R", true, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)

	var saldo58 string

	if saldo58 == "0" {
		saldo58 = ""
	}

	var saldo65 string

	if saldo65 == "0" {
		saldo65 = ""
	}

	pdf.CellFormat(110, 4, "Intereses y Rendimientos Financieros", "", 0,
		"", false, 0, "")
	pdf.CellFormat(5, 4, "58", "", 0,
		"", false, 0, "")
	pdf.CellFormat(27, 4, saldo58, "", 0,
		"R", false, 0, "")
	pdf.CellFormat(5, 4, "65", "", 0,
		"", false, 0, "")
	pdf.CellFormat(33, 4, saldo65, "", 0,
		"R", false, 0, "")
	pdf.SetFillColor(225, 232, 239)
	pdf.Ln(-1)
	pdf.SetX(21)

	var saldo59 string

	if saldo59 == "0" {
		saldo59 = ""
	}

	var saldo66 string

	if saldo66 == "0" {
		saldo66 = ""
	}

	pdf.CellFormat(110, 4, "Enajenacion de Activos Fijos", "", 0,
		"", true, 0, "")
	pdf.CellFormat(5, 4, "59", "", 0,
		"", true, 0, "")
	pdf.CellFormat(27, 4, saldo59, "", 0,
		"R", true, 0, "")
	pdf.CellFormat(5, 4, "66", "", 0,
		"", true, 0, "")
	pdf.CellFormat(33, 4, saldo66, "", 0,
		"R", true, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)

	var saldo60 string

	if saldo60 == "0" {
		saldo60 = ""
	}

	var saldo67 string

	if saldo67 == "0" {
		saldo67 = ""
	}

	pdf.CellFormat(110, 4, "Loterias, Rifas, Apuestas y Similares", "", 0,
		"", false, 0, "")
	pdf.CellFormat(5, 4, "60", "", 0,
		"", false, 0, "")
	pdf.CellFormat(27, 4, saldo60, "", 0,
		"R", false, 0, "")
	pdf.CellFormat(5, 4, "67", "", 0,
		"", false, 0, "")
	pdf.CellFormat(33, 4, saldo67, "", 0,
		"R", false, 0, "")
	pdf.SetFillColor(225, 232, 239)
	pdf.Ln(-1)
	pdf.SetX(21)

	var saldo61 string

	if saldo61 == "0" {
		saldo61 = ""
	}

	var saldo68 string

	if saldo68 == "0" {
		saldo68 = ""
	}

	pdf.CellFormat(110, 4, "Otros", "", 0,
		"", true, 0, "")
	pdf.CellFormat(5, 4, "61", "", 0,
		"", true, 0, "")
	pdf.CellFormat(27, 4, saldo61, "", 0,
		"R", true, 0, "")
	pdf.CellFormat(5, 4, "68", "", 0,
		"", true, 0, "")
	pdf.CellFormat(33, 4, saldo68, "", 0,
		"R", true, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)

	var saldo62 string

	if saldo62 == "0" {
		saldo62 = ""
	}

	var saldo69 string

	if saldo69 == "0" {
		saldo69 = ""
	}

	pdf.CellFormat(110, 4, "Totales (Valores Recibidos suma 56 a 61) (Valores Retenidos suma 63 a 68", "1", 0,
		"", false, 0, "")
	pdf.CellFormat(5, 4, "62", "1", 0,
		"", false, 0, "")
	pdf.CellFormat(27, 4, saldo62, "1", 0,
		"R", false, 0, "")
	pdf.CellFormat(5, 4, "69", "1", 0,
		"", false, 0, "")
	pdf.CellFormat(33, 4, saldo69, "1", 0,
		"R", false, 0, "")

	pdf.Ln(-1)
	pdf.SetX(21)

	var saldo70 string

	if saldo70 == "0" {
		saldo70 = ""
	}

	pdf.CellFormat(142, 4, ene("Total Retenciones Año Gravable 2021 Suma casilla 55 + 69"), "1", 0,
		"", false, 0, "")
	pdf.CellFormat(5, 4, "70", "1", 0,
		"", false, 0, "")
	pdf.CellFormat(33, 4, saldo70, "1", 0,
		"R", false, 0, "")
	pdf.SetTextColor(0, 0, 0)
	pdf.SetFillColor(225, 232, 239)
	pdf.SetFont("Arial", "B", 8)

	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(7, 4, "Item", "1", 0,
		"C", true, 0, "")
	pdf.CellFormat(135, 4, "71. Identificacion de los Bienes Poseidos", "1", 0,
		"C", true, 0, "")
	pdf.CellFormat(38, 4, "72. Valor Patrimonial", "1", 0,
		"C", true, 0, "")
	pdf.SetFont("Arial", "", 8)
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(7, 4, "1", "1", 0,
		"C", false, 0, "")
	pdf.CellFormat(135, 4, "", "1", 0,
		"C", false, 0, "")
	pdf.CellFormat(38, 4, "", "1", 0,
		"C", false, 0, "")
	pdf.SetFillColor(225, 232, 239)
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(7, 4, "2", "1", 0,
		"C", true, 0, "")
	pdf.CellFormat(135, 4, "", "1", 0,
		"C", true, 0, "")
	pdf.CellFormat(38, 4, "", "1", 0,
		"C", true, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(7, 4, "3", "1", 0,
		"C", false, 0, "")
	pdf.CellFormat(135, 4, "", "1", 0,
		"C", false, 0, "")
	pdf.CellFormat(38, 4, "", "1", 0,
		"C", false, 0, "")
	pdf.SetFillColor(225, 232, 239)
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(7, 4, "4", "1", 0,
		"C", true, 0, "")
	pdf.CellFormat(135, 4, "", "1", 0,
		"C", true, 0, "")
	pdf.CellFormat(38, 4, "", "1", 0,
		"C", true, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(7, 4, "5", "1", 0,
		"C", false, 0, "")
	pdf.CellFormat(135, 4, "", "1", 0,
		"C", false, 0, "")
	pdf.CellFormat(38, 4, "", "1", 0,
		"C", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(142, 4, "Deudas Vigentes a Diciembre 31 de 2021", "1", 0,
		"L", true, 0, "")
	pdf.SetTextColor(0, 0, 0)
	pdf.CellFormat(5, 4, "73", "1", 0,
		"C", false, 0, "")
	pdf.CellFormat(33, 4, "", "1", 0,
		"C", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(180, 4, "Identificacion del Dependiente Economico de acuerdo al paragrafo 2 del Articulo 387 del E. T.", "1", 0,
		"C", false, 0, "")
	pdf.SetFillColor(225, 232, 239)
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(45, 4, "74. Tipo Doc 75. Numero Doc", "1", 0,
		"", true, 0, "")
	pdf.CellFormat(97, 4, " 76. Apellidos y Nombres", "1", 0,
		"", true, 0, "")
	pdf.CellFormat(38, 4, " 77. Parentesco", "1", 0,
		"", true, 0, "")
	pdf.Ln(-1)
	pdf.SetX(21)
	pdf.CellFormat(45, 4, "", "1", 0,
		"", false, 0, "")
	pdf.CellFormat(97, 4, "", "1", 0,
		"", false, 0, "")
	pdf.CellFormat(38, 4, "", "1", 0,
		"", false, 0, "")

	pdf.SetFooterFunc(func() {

		pdf.SetY(252)
		pdf.SetX(25)
		pdf.SetFont("Arial", "", 5)
		pdf.CellFormat(142, 2, ene("Certifico que Durante el Año Gravable 2021:"), "0", 0,
			"", false, 0, "")

		pdf.SetY(252)
		pdf.SetX(160)
		pdf.SetFont("Arial", "", 6)
		pdf.CellFormat(42, 2, ene("    Firma del Trabajador o Pensionado"), "0", 0,
			"C", false, 0, "")

		pdf.SetFont("Arial", "", 5)
		pdf.Ln(-1)
		pdf.SetX(25)
		pdf.CellFormat(140, 2, ene("1. Mi Patrimonio no Excedio de 4.500 UVT $ 163.386.000"), "0", 0,
			"", false, 0, "")

		pdf.Ln(-1)
		pdf.SetX(25)
		pdf.SetFont("Arial", "", 5)
		pdf.CellFormat(140, 2, ene("2. Mis Ingresos Fueron Inferiores a 1.400 UVT $ 50.831.000"), "0", 0,
			"", false, 0, "")
		pdf.Ln(-1)
		pdf.SetX(25)
		pdf.CellFormat(140, 2, ene("3. No Fui Responsable del Impuesto a Las Ventas"), "0", 0,
			"", false, 0, "")

		pdf.Ln(-1)
		pdf.SetX(25)
		pdf.CellFormat(140, 2, ene("4. Mis Consumos Mediante Tarjeta de Credito No Excedieron La Suma de 1.400 UVT $ 50.831.000"), "0", 0,
			"", false, 0, "")

		pdf.Ln(-1)
		pdf.SetX(25)
		pdf.CellFormat(140, 2, ene("5. Que el Total de Mis Compras y Consumos No Superaron la Suma de 1.400 UVT $ 50.831.000"), "0", 0,
			"", false, 0, "")
		pdf.Ln(-1)
		pdf.SetX(25)
		pdf.CellFormat(140, 2, ene("6. Que el Valor Total de Mis Consignaciones Bancarias, Depositos o Inversiones Financieras No Excedieron los 1.400 UVT $ 50.831.000"), "0", 0,
			"", false, 0, "")
		pdf.Ln(-1)
		pdf.SetX(25)
		pdf.CellFormat(140, 2, ene("Por lo Tanto, Manifiesto Que No Extoy Obligado a Presentar Declaracion de Renta y Complementarios por El Año Gravable 2021"), "0", 0,
			"", false, 0, "")

		// LINEA HORIZONTAL
		pdf.Line(21, 269, 201, 269)

		// LINEA VERTICAL
		pdf.Line(21, 50, 21, 269)
		pdf.Line(201, 50, 201, 269)
		pdf.Line(163, 250, 163, 269)
		pdf.Line(163, 70, 163, 167)

		pdf.Line(168, 83, 168, 139)
		pdf.Line(168, 143, 168, 167)

		pdf.Line(168, 183, 168, 211)
		pdf.Line(163, 183, 163, 211)

		pdf.Line(131, 183, 131, 207)
		pdf.Line(136, 183, 136, 207)

		pdf.Ln(3)
		pdf.SetFont("Arial", "", 8)
		pdf.SetX(21)
		pdf.SetY(269)
		pdf.CellFormat(40, 4, "Sadconf.com", "", 0,
			"C", false, 0, "")
		pdf.SetX(165)
		pdf.CellFormat(40, 4, fmt.Sprintf(" %d de {nb}", pdf.PageNo()), "",
			0, "R", false, 0, "")
	})

	err1 = pdf.Output(&buf)
	if err1 != nil {
		panic(err1.Error())
	}
	w.Header().Set("Content-Type", "application/pdf; charset=utf-8")
	w.Write(buf.Bytes())
}

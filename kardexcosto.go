package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

//const operacionInventarioInicial = "1"
//const operacionCompra = "2"
//const operacionSoporte = "3"
//const operacionDevolucionVenta = "4"
//const operacionTrasladoEntrada = "5"
//const operacionDevolucionCompra = "6"
//const operacionDevolucionSoporte = "7"
//const operacionVenta = "8"
//const operacionTrasladoSalida = "9"

type KardexCosto struct {
	Filas string  `json:"Filas"`
	Tipo  string  `json:"Tipo"`
	Total float64 `json:"Total"`
}

// CENTRO TABLA
type kardex55 struct {
	Fecha string `json:"Fecha"`
	//time.Time `json:"Fecha"`
	Filas           string  `json:"Filas"`
	Producto        string  `json:"Producto"`
	Tipo            string  `json:"Tipo"`
	Operacion       string  `json:"Operacion"`
	Codigo          string  `json:"Documento"`
	Bodega          string  `json:"Bodega"`
	Cantidadentrada float64 `json:"CantidadE"`
	Precioentrada   float64 `json:"PrecioE"`
	Totalentrada    float64 `json:"TotalE"`
	Cantidadsalida  float64 `json:"Cantidad"`
	Preciosalida    float64 `json:"Precio"`
	Totalsalida     float64 `json:"Total"`
	Cantidadsaldo   float64 `json:"CantidadT"`
	Preciosaldo     float64 `json:"PrecioT"`
	Totalsaldo      float64 `json:"TotalT"`
}

func KardexCostoDatos(w http.ResponseWriter, r *http.Request) {
	db := dbConn()

	//CodigoParamerto := mux.Vars(r)["codigo"]
	miMes := mux.Vars(r)["mes"]
	Centro := mux.Vars(r)["centro"]
	args := []interface{}{}

	res := []inventario{}
	listadokardex := []kardex{}
	//listadokardexfinal := []kardex{}
	listadocosto := []KardexCosto{}

	var cadena string
	var cadenaProducto string

	cadenaProducto = "SELECT  distinct inventario.producto, producto.nombre,producto.iva FROM inventario " +
		" inner join producto on producto.codigo=inventario.producto where" +
		"  EXTRACT(MONTH FROM  fecha)<=$1" +
		" group by inventario.producto, producto.nombre,producto.iva order by inventario.producto"

	args = append(args, &miMes)

	var siexisteproducto bool
	selDB1, err := db.Query(cadenaProducto, args...)
	switch err {
	//resltadvaa
	case nil:
		log.Printf("Datos Kardex existe")
		siexisteproducto = true
	case sql.ErrNoRows:
		log.Println("Datos Kardex no encontrados")
	default:
		log.Printf("tercero error: %s\n", err)
	}

	var simueve bool
	simueve = false

	if siexisteproducto == true {
		var codigoProducto string
		var nombreProducto string
		var ivaProduco string
		args = nil

		for selDB1.Next() {
			// recorrer productos
			args = nil
			simueve = true

			err = selDB1.Scan(&codigoProducto, &nombreProducto, &ivaProduco)

			log.Println("Datos Kardex producto" + codigoProducto)
			if err != nil {
				panic(err.Error())
			}

			cadena = "SELECT fecha,EXTRACT(MONTH FROM  fecha) as mes, tipo, codigo, bodega, producto," +
				" cantidad, precio, operacion FROM inventario where" +
				" producto=$1 and " +
				"  EXTRACT(MONTH FROM  fecha)<=$2" +
				" ORDER BY producto,Fecha,operacion "
			args = append(args, &codigoProducto)
			args = append(args, &miMes)

			selDB, err := db.Query(cadena, args...)
			log.Println("sql : " + cadena)
			var siexiste bool
			siexiste = false
			switch err {
			//resltadvaa
			case nil:
				log.Printf("Datos Kardex existe")
				siexiste = true
			case sql.ErrNoRows:
				log.Println("Datos Kardex no encontrados")
			default:
				log.Printf("tercero error: %s\n", err)
			}

			if siexiste == true {
				var saldo float64
				var costo float64
				//var precosto float64
				var total float64
				saldo = 0
				costo = 0
				total = 0
				var primero bool
				primero = false
				var totalanterior float64
				totalanterior = 0
				var Cantidadsaldo float64
				var Preciosaldo float64
				var Totalsaldo float64
				var Producto string

				for selDB.Next() {
					var Fecha time.Time
					var Tipo string

					var Mes string
					var Codigo string
					var Bodega string

					var Cantidad string
					var Precio string
					var Operacion string
					var Cantidadentrada float64
					var Precioentrada float64
					var Totalentrada float64
					var Cantidadsalida float64
					var Preciosalida float64
					var Totalsalida float64

					Cantidadentrada = 0
					Precioentrada = 0
					Totalentrada = 0
					Cantidadsalida = 0
					Preciosalida = 0
					Totalsalida = 0
					Cantidadsaldo = 0
					Preciosaldo = 0
					Totalsaldo = 0

					err = selDB.Scan(&Fecha, &Mes, &Tipo, &Codigo, &Bodega, &Producto, &Cantidad,
						&Precio, &Operacion)
					if err != nil {
						panic(err.Error())
					}

					// OPERACION DE ASIGNACION AL COSTO O PRECIO DE ENTRADAS

					if Operacion == operacionInventarioInicial || Operacion == operacionCompra || Operacion == operacionSoporte || Operacion == operacionDevolucionVenta || Operacion == operacionTrasladoEntrada {

						Cantidadentrada = Flotante(Cantidad)
						Precioentrada = Flotante(Precio)
						Totalentrada = Flotante(Cantidad) * Flotante(Precio)

						if Operacion == operacionDevolucionVenta || Operacion == operacionTrasladoEntrada {
							Cantidadentrada = Flotante(Cantidad)
							Precioentrada = costo
							Totalentrada = Flotante(Cantidad) * costo
						} else {
							Cantidadentrada = Flotante(Cantidad)
							Precioentrada = Flotante(Precio)
							Totalentrada = Flotante(Cantidad) * Flotante(Precio)
						}

					} else {

						if Operacion == operacionDevolucionCompra || Operacion == operacionDevolucionSoporte {
							Cantidadsalida = Flotante(Cantidad)
							Preciosalida = Flotante(Precio)
							Totalsalida = Flotante(Cantidad) * Flotante(Precio)
						} else {

							Cantidadsalida = Flotante(Cantidad)
							Preciosalida = costo
							Totalsalida = Flotante(Cantidad) * costo
						}

					}

					if primero == false || saldo == 0 {
						saldo = Flotante(Cantidad)
						costo = Flotante(Precio)
						total = Flotante(Cantidad) * Flotante(Precio)
						primero = true
					} else {
						log.Println("cantidad")
						log.Println(Cantidadentrada + saldo)
						log.Println("total")
						log.Println(((Flotante(Cantidad) * Flotante(Precio)) + total))

						// OPERACION DE ASIGNACION CALCULOS DE ENTRADA Y SALIDA
						if Operacion == operacionInventarioInicial || Operacion == operacionCompra || Operacion == operacionSoporte || Operacion == operacionDevolucionVenta || Operacion == operacionTrasladoEntrada {

							if Operacion == operacionDevolucionVenta || Operacion == operacionTrasladoEntrada {
								costo = ((Flotante(Cantidad) * costo) + total) / (Cantidadentrada - Cantidadsalida + saldo)
							} else {
								costo = ((Flotante(Cantidad) * Flotante(Precio)) + total) / (Cantidadentrada - Cantidadsalida + saldo)
							}

						} else {
							costo = (total - (Flotante(Cantidad) * (costo))) / (Cantidadentrada - Cantidadsalida + saldo)

							if Operacion == operacionDevolucionCompra || Operacion == operacionDevolucionSoporte {

								costo = (total - (Flotante(Cantidad) * Flotante(Precio))) / (Cantidadentrada - Cantidadsalida + saldo)
							} else {

								costo = (total - (Flotante(Cantidad) * (costo))) / (Cantidadentrada - Cantidadsalida + saldo)

							}
						}
						log.Println("costo")
						log.Println(costo)
						saldo = saldo + Cantidadentrada - Cantidadsalida
						total = saldo * costo
						costo = Redondear(costo, 2)

					}
					if saldo == 0 {
						//totalanterior=0
						Cantidadsaldo = 0
						Preciosaldo = 0
						Totalsaldo = 0

						if Cantidadentrada > 0 {
							total = Totalentrada
						} else {
							total = Totalsalida
							Totalsalida = totalanterior
						}
						totalanterior = 0
						//total
						Cantidadsaldo = saldo
						Preciosaldo = 0
						//costo
						Totalsaldo = 0

					} else {

						if Cantidadentrada > 0 {
							total = totalanterior + Totalentrada
						} else {
							total = totalanterior - Totalsalida
						}
						totalanterior = total
						Cantidadsaldo = saldo
						Preciosaldo = costo
						Totalsaldo = total
					}
					//if saldo == 0{
					//	totalanterior=0
					//	Cantidadsaldo=0
					//	Preciosaldo=0
					//	Totalsaldo=0
					//
					//}else{
					//
					//	if Cantidadentrada>0{
					//		total=totalanterior+Totalentrada
					//	}else{
					//		total=totalanterior-Totalsalida
					//	}
					//	totalanterior=total
					//	Cantidadsaldo = saldo
					//	Preciosaldo = costo
					//	Totalsaldo = total
					//}

					// operaciones
					if Cantidadsaldo < 0 {

						log.Println("Saldo Negativo " + Producto + " " + Operacion + " " + Tipo + " " + Codigo)
					}

					res = append(res, inventario{Fecha, Tipo, Codigo,
						Bodega, Producto, Cantidad, Precio, Operacion})

					listadokardex = append(listadokardex, kardex{Fecha.Format("02/01/06"),
						Mes,
						"",
						Producto,
						Tipo,
						Operacion,
						Codigo,
						Bodega,
						Cantidadentrada,
						Precioentrada,
						Totalentrada,
						Cantidadsalida,
						Preciosalida,
						Totalsalida,
						Cantidadsaldo,
						Preciosaldo,
						Totalsaldo})

				}

			}
		}
	}

	if simueve == false {
		var slice []string
		slice = make([]string, 0)
		data, _ := json.Marshal(slice)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	} else {

		var totalventa float64
		var totaldevolucionventa float64
		var totaltrasladoentrada float64
		var totaltrasladosalida float64

		for _, miFila := range listadokardex {
			if miFila.Mes == miMes {

				if miFila.Tipo == "Venta" {
					totalventa += miFila.Totalsalida
					//listadokardexfinal=append(listadokardexfinal,miFila)
				}

				if miFila.Tipo == "Devolucionventa" {
					totaldevolucionventa += miFila.Totalentrada
					//listadokardexfinal=append(listadokardexfinal,miFila)
				}

				if miFila.Tipo == "Traslado" {
					totaltrasladoentrada += miFila.Totalentrada
					totaltrasladosalida += miFila.Totalsalida
					//listadokardexfinal=append(listadokardexfinal,miFila)
				}
			}
		}
		var totaldetalle float64

		totaldetalle = totalventa - totaldevolucionventa - totaltrasladoentrada + totaltrasladosalida

		listadocosto = append(listadocosto, KardexCosto{"", "Venta", totalventa})
		listadocosto = append(listadocosto, KardexCosto{"", "Devolucion Venta", totaldevolucionventa})
		listadocosto = append(listadocosto, KardexCosto{"", "Traslado Entrada", totaltrasladoentrada})
		listadocosto = append(listadocosto, KardexCosto{"", "Traslado Salida", totaltrasladosalida})
		listadocosto = append(listadocosto, KardexCosto{"", "Costo De ventas", totaldetalle})

		parametrosinventario := configuracioninventario{}
		err = db.Get(&parametrosinventario, "SELECT * FROM configuracioninventario")
		if err != nil {
			panic(err.Error())
		}

		var Documentocontable = "20"
		var numeroFactura = miMes

		// BORRA MOVIMIENTOS
		var consultaborracomprobante = "delete from comprobante where documento=$2 and numero=$1"
		db.Exec(consultaborracomprobante, numeroFactura, Documentocontable)

		var consultaborracomprobantedetalle = "delete from comprobantedetalle where documento=$2 and  numero=$1"
		db.Exec(consultaborracomprobantedetalle, numeroFactura, Documentocontable)
		var miTercero = "1"
		var miFilaComprobante int
		miFilaComprobante = 0
		var fechaString string
		var miperiodo = periodoSesion(r)
		fechaString = fechaInicial(miperiodo, miMes)
		const (
			layoutISO = "2006-01-02"
		)
		fechaDate, _ := time.Parse(layoutISO, fechaString)

		miFilaComprobante++
		miComprobanteDetalle := []comprobantedetalle{}

		miComprobanteDetalle = append(miComprobanteDetalle,
			comprobantedetalle{strconv.Itoa(miFilaComprobante),
				parametrosinventario.Cuentacosto,
				miTercero,
				Centro,
				" Costo de Ventas " + mesLetras(miMes),
				"",
				FormatoFlotante(totaldetalle),
				"",
				Documentocontable,
				numeroFactura,
				fechaDate,
				fechaDate, "", ""})
		// Inserta Fila contra
		miFilaComprobante++
		miComprobanteDetalle = append(miComprobanteDetalle,
			comprobantedetalle{strconv.Itoa(miFilaComprobante),
				parametrosinventario.Cuentacostocontra,
				miTercero,
				Centro,
				" Costo de Ventas " + mesLetras(miMes),
				"",
				"",
				FormatoFlotante(totaldetalle),
				Documentocontable,
				numeroFactura,
				fechaDate,
				fechaDate, "", ""})

		// crea comprobante
		ComprobanteAgregarGenerar(comprobante{Documentocontable,
			numeroFactura, fechaDate,
			fechaDate,
			miperiodo,
			"",
			"",
			"",
			FormatoFlotante(totaldetalle),
			FormatoFlotante(totaldetalle),
			"Actualizar",
			miComprobanteDetalle,
			nil})
		//Venta
		//Devolucionventa
		//Traslado E
		//Traslado Salida
		//if TipoParametro=="Todos"{
		//	listadokardexfinal=listadokardex
		//}else{
		//	for _, miFila := range listadokardex {
		//		if miFila.Operacion==TipoParametro{
		//			listadokardexfinal=append(listadokardexfinal,miFila)
		//		}
		//
		//	}
		//}
		//if TipoParametro=="Todos"{
		//listadokardexfinal=listadokardex
		//}else{
		//	var totalEntrada float64
		//	var totalSalida float64
		//
		//	totalEntrada=0
		//	totalSalida=0
		//
		//
		//	for _, miFila := range listadokardex {
		//
		//		if miFila.Operacion==TipoParametro{
		//			totalEntrada+=miFila.Totalentrada
		//			totalSalida+=miFila.Totalsalida
		//			listadokardexfinal=append(listadokardexfinal,miFila)
		//
		//		}
		//
		//	}
		//
		//	// suma totales
		//	listadokardexfinal=append(listadokardexfinal, kardex{
		//		"",
		//		"",
		//		"",
		//		"Total",
		//		"",
		//		"",
		//		"",
		//		0,
		//		0,
		//		totalEntrada,
		//		0,
		//		0,
		//		totalSalida,
		//		0,
		//		0,
		//		0})
		//}
		data, _ := json.Marshal(listadocosto)
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	}
}

// CENTRO KARDEX
func KardexCostoLista(w http.ResponseWriter, r *http.Request) {
	tmp, _ := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/kardex/kardexCostoLista.html")
	//	db := dbConn()
	Codigo := mux.Vars(r)["codigo"]
	//res := []inventario{}
	//listadokardex := []kardex{}

	if Codigo == "False" {

	} else {

		//	FechaInicial := mux.Vars(r)["fechainicial"]

	}

	varmap := map[string]interface{}{
		//"res":     listadokardex,
		"hosting":   ruta,
		"bodega":    ListaBodega(),
		"producto":  ListaProducto(),
		"centro":    ListaCentro(),
		"miperiodo": periodoSesion(r),
	}
	tmp.Execute(w, varmap)
}

// INICIA CENTRO PDF
func kardexPdf111(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Fecha := mux.Vars(r)["Fecha"]
	t := inventario{}
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	err := db.Get(&t, "SELECT * FROM inventario where fecha=$1", Fecha)
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
	pdf.CellFormat(40, 4, "Fecha", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Fecha.Format("02/01/2006"), "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(30)
	pdf.CellFormat(40, 4, "Tipo:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Tipo, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(30)
	pdf.CellFormat(40, 4, "Documento No.:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Codigo, "", 0,
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

// TERMINA CENTRO PDF

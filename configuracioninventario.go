package main

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
)

// CONFIGURACIONINVENTARIO TABLA
type configuracioninventario struct {
	// INICIA ESTRUCTURA COMPRA
	Compracuenta19                  string
	Compranombre19                  string
	Compracuenta5                   string
	Compranombre5                   string
	Compracuenta0                   string
	Compranombre0                   string
	Compraiva19                     string
	Compranombreiva19               string
	Compraiva5                      string
	Compranombreiva5                string
	Compracuentaretfte              string
	Compranombreretfte              string
	Compracuentaretica              string
	Compranombreretica              string
	Compracuentaretotra             string
	Compranombreretotra             string
	Compracuentadescuento           string
	Compranombredescuento           string
	Compradevolucioncuenta19        string
	Compradevolucionnombre19        string
	Compradevolucioncuenta5         string
	Compradevolucionnombre5         string
	Compradevolucioncuenta0         string
	Compradevolucionnombre0         string
	Compradevolucioniva19           string
	Compradevolucionnombreiva19     string
	Compradevolucioniva5            string
	Compradevolucionnombreiva5      string
	Compradevolucioncuentaretfte    string
	Compradevolucionnombreretfte    string
	Compradevolucioncuentaretica    string
	Compradevolucionnombreretica    string
	Compradevolucioncuentaretotra   string
	Compradevolucionnombreretotra   string
	Compradevolucioncuentadescuento string
	Compradevolucionnombredescuento string
	Compracuentaproveedor           string
	Compranombreproveedor           string
	Compracuentabase                string
	Compracuentaporcentajeretfte    string
	Compracuentaporcentajeretotra   string
	// TERMINA ESTRUCTURA COMPRA

	// INICIA ESTRUCTURA COMPRA
	Cscuenta19                  string
	Csnombre19                  string
	Cscuenta5                   string
	Csnombre5                   string
	Cscuenta0                   string
	Csnombre0                   string
	Csiva19                     string
	Csnombreiva19               string
	Csiva5                      string
	Csnombreiva5                string
	Cscuentaretfte              string
	Csnombreretfte              string
	Cscuentaretica              string
	Csnombreretica              string
	Cscuentaretotra             string
	Csnombreretotra             string
	Cscuentadescuento           string
	Csnombredescuento           string
	Csdevolucioncuenta19        string
	Csdevolucionnombre19        string
	Csdevolucioncuenta5         string
	Csdevolucionnombre5         string
	Csdevolucioncuenta0         string
	Csdevolucionnombre0         string
	Csdevolucioniva19           string
	Csdevolucionnombreiva19     string
	Csdevolucioniva5            string
	Csdevolucionnombreiva5      string
	Csdevolucioncuentaretfte    string
	Csdevolucionnombreretfte    string
	Csdevolucioncuentaretica    string
	Csdevolucionnombreretica    string
	Csdevolucioncuentaretotra   string
	Csdevolucionnombreretotra   string
	Csdevolucioncuentadescuento string
	Csdevolucionnombredescuento string
	Cscuentaproveedor           string
	Csnombreproveedor           string
	Cscuentabase                string
	Cscuentaporcentajeretfte    string
	Cscuentaporcentajeretotra   string
	// TERMINA ESTRUCTURA COMPRA

	// INICIA ESTRUCTURA VENTA
	Ventacuenta19                      string
	Ventanombre19                      string
	Ventacuenta5                       string
	Ventanombre5                       string
	Ventacuenta0                       string
	Ventanombre0                       string
	Ventaiva19                         string
	Ventanombreiva19                   string
	Ventaiva5                          string
	Ventanombreiva5                    string
	Ventacuentaret2201                 string
	Ventanombreret2201                 string
	Ventacuentadescuento               string
	Ventanombredescuento               string
	Ventadevolucioncuenta19            string
	Ventadevolucionnombre19            string
	Ventadevolucioncuenta5             string
	Ventadevolucionnombre5             string
	Ventadevolucioncuenta0             string
	Ventadevolucionnombre0             string
	Ventadevolucioniva19               string
	Ventadevolucionnombreiva19         string
	Ventadevolucioniva5                string
	Ventadevolucionnombreiva5          string
	Ventadevolucioncuentaret2201       string
	Ventadevolucionnombreret2201       string
	Ventadevolucioncuentadescuento     string
	Ventadevolucionnombredescuento     string
	Ventacuentacliente                 string
	Ventanombrecliente                 string
	Ventacontracuentaret2201           string
	Ventacontranombreret2201           string
	Ventadevolucioncontracuentaret2201 string
	Ventadevolucioncontranombreret2201 string
	Ventacuentaporcentajeret2201       string
	Ventatipoiva                       string
	// TERMINA ESTRUCTURA VENTA

	// INICIA ESTRUCTURA SERVICIO
	Ventaserviciocuenta19                      string
	Ventaservicionombre19                      string
	Ventaserviciocuenta5                       string
	Ventaservicionombre5                       string
	Ventaserviciocuenta0                       string
	Ventaservicionombre0                       string
	Ventaservicioiva19                         string
	Ventaservicionombreiva19                   string
	Ventaservicioiva5                          string
	Ventaservicionombreiva5                    string
	Ventaserviciocuentaret2201                 string
	Ventaservicionombreret2201                 string
	Ventaserviciocuentadescuento               string
	Ventaservicionombredescuento               string
	Ventaserviciodevolucioncuenta19            string
	Ventaserviciodevolucionnombre19            string
	Ventaserviciodevolucioncuenta5             string
	Ventaserviciodevolucionnombre5             string
	Ventaserviciodevolucioncuenta0             string
	Ventaserviciodevolucionnombre0             string
	Ventaserviciodevolucioniva19               string
	Ventaserviciodevolucionnombreiva19         string
	Ventaserviciodevolucioniva5                string
	Ventaserviciodevolucionnombreiva5          string
	Ventaserviciodevolucioncuentaret2201       string
	Ventaserviciodevolucionnombreret2201       string
	Ventaserviciodevolucioncuentadescuento     string
	Ventaserviciodevolucionnombredescuento     string
	Ventaserviciocuentacliente                 string
	Ventaservicionombrecliente                 string
	Ventaserviciocontracuentaret2201           string
	Ventaserviciocontranombreret2201           string
	Ventaserviciodevolucioncontracuentaret2201 string
	Ventaserviciodevolucioncontranombreret2201 string
	Ventaserviciocuentaporcentajeret2201       string
	Ventaserviciotipoiva                       string
	// TERMINA ESTRUCTURA SERVICIO

	// INICIA ESTRUCTURA SOPORTE
	Soportecuenta19                  string
	Soportenombre19                  string
	Soportecuenta5                   string
	Soportenombre5                   string
	Soportecuenta0                   string
	Soportenombre0                   string
	Soportecuentaretfte              string
	Soportenombreretfte              string
	Soportecuentaretica              string
	Soportenombreretica              string
	Soportecuentaretotra             string
	Soportenombreretotra             string
	Soportecuentadescuento           string
	Soportenombredescuento           string
	Soportedevolucioncuenta19        string
	Soportedevolucionnombre19        string
	Soportedevolucioncuenta5         string
	Soportedevolucionnombre5         string
	Soportedevolucioncuenta0         string
	Soportedevolucionnombre0         string
	Soportedevolucioncuentaretfte    string
	Soportedevolucionnombreretfte    string
	Soportedevolucioncuentaretica    string
	Soportedevolucionnombreretica    string
	Soportedevolucioncuentaretotra   string
	Soportedevolucionnombreretotra   string
	Soportedevolucioncuentadescuento string
	Soportedevolucionnombredescuento string
	Soportecuentaproveedor           string
	Soportenombreproveedor           string
	Soportecuentabase                string
	Soportecuentaporcentajeretfte    string
	Soportecuentaporcentajeretotra   string
	// TERMINA ESTRUCTURA SOPORTE

	// INICIA ESTRUCTURA SOPORTE SERVICIO
	Soporteserviciocuenta19                  string
	Soporteservicionombre19                  string
	Soporteserviciocuenta5                   string
	Soporteservicionombre5                   string
	Soporteserviciocuenta0                   string
	Soporteservicionombre0                   string
	Soporteserviciocuentaretfte              string
	Soporteservicionombreretfte              string
	Soporteserviciocuentaretica              string
	Soporteservicionombreretica              string
	Soporteserviciocuentaretotra             string
	Soporteservicionombreretotra             string
	Soporteserviciocuentadescuento           string
	Soporteservicionombredescuento           string
	Soporteserviciodevolucioncuenta          string
	Soporteserviciodevolucionnombre          string
	Soporteserviciodevolucioncuentaretfte    string
	Soporteserviciodevolucionnombreretfte    string
	Soporteserviciodevolucioncuentaretica    string
	Soporteserviciodevolucionnombreretica    string
	Soporteserviciodevolucioncuentaretotra   string
	Soporteserviciodevolucionnombreretotra   string
	Soporteserviciodevolucioncuentadescuento string
	Soporteserviciodevolucionnombredescuento string
	Soporteserviciocuentaproveedor           string
	Soporteservicionombreproveedor           string
	Soporteserviciocuentabase                string
	Soporteserviciocuentaporcentajeretfte    string
	Soporteserviciocuentaporcentajeretotra   string
	// TERMINA ESTRUCTURA SOPORTE SERVICIO

	// INICIA ESTRUCTURA FACTURA GASTO
	Facturagastocuenta19                  string
	Facturagastonombre19                  string
	Facturagastocuenta5                   string
	Facturagastonombre5                   string
	Facturagastocuenta0                   string
	Facturagastonombre0                   string
	Facturagastocuentaretfte              string
	Facturagastonombreretfte              string
	Facturagastocuentaretica              string
	Facturagastonombreretica              string
	Facturagastocuentaretotra             string
	Facturagastonombreretotra             string
	Facturagastocuentadescuento           string
	Facturagastonombredescuento           string
	Facturagastodevolucioncuenta          string
	Facturagastodevolucionnombre          string
	Facturagastodevolucioncuentaretfte    string
	Facturagastodevolucionnombreretfte    string
	Facturagastodevolucioncuentaretica    string
	Facturagastodevolucionnombreretica    string
	Facturagastodevolucioncuentaretotra   string
	Facturagastodevolucionnombreretotra   string
	Facturagastodevolucioncuentadescuento string
	Facturagastodevolucionnombredescuento string
	Facturagastocuentaproveedor           string
	Facturagastonombreproveedor           string
	Facturagastocuentabase                string
	Facturagastocuentaporcentajeretfte    string
	Facturagastocuentaporcentajeretotra   string
	Facturagastocuentaiva                 string
	Facturagastonombreiva                 string
	Facturagastoporcentajeiva             string
	// TERMINA ESTRUCTURA SOPORTE

	// INICIA ESTRUCTURA COSTO DE VENTAS
	Cuentacosto             string
	Cuentacostonombre       string
	Cuentacostocontra       string
	Cuentacostocontranombre string
	Decimales               string
	// TERMINA ESTRUCTURA COSTO DE VENTAS

}

func Datosinicialesconfiguracioninventario() {
	var q string
	q = "INSERT INTO configuracioninventario("
	q += "	compracuenta19	 , "
	q += "	compranombre19	 , "
	q += "	compracuenta5	 , "
	q += "	compranombre5	 , "
	q += "	compracuenta0	 , "
	q += "	compranombre0	 , "
	q += "	compraiva19	 , "
	q += "	compranombreiva19	 , "
	q += "	compraiva5	 , "
	q += "	compranombreiva5	 , "
	q += "	compracuentaretfte	 , "
	q += "	compranombreretfte	 , "
	q += "	compracuentaretica	 , "
	q += "	compranombreretica	 , "
	q += "	compracuentaretotra	 , "
	q += "	compranombreretotra	 , "
	q += "	compracuentadescuento	 , "
	q += "	compranombredescuento	 , "
	q += "	compradevolucioncuenta19	 , "
	q += "	compradevolucionnombre19	 , "
	q += "	compradevolucioncuenta5	 , "
	q += "	compradevolucionnombre5	 , "
	q += "	compradevolucioncuenta0	 , "
	q += "	compradevolucionnombre0	 , "
	q += "	compradevolucioniva19	 , "
	q += "	compradevolucionnombreiva19	 , "
	q += "	compradevolucioniva5	 , "
	q += "	compradevolucionnombreiva5	 , "
	q += "	compradevolucioncuentaretfte	 , "
	q += "	compradevolucionnombreretfte	 , "
	q += "	compradevolucioncuentaretica	 , "
	q += "	compradevolucionnombreretica	 , "
	q += "	compradevolucioncuentaretotra	 , "
	q += "	compradevolucionnombreretotra	 , "
	q += "	compradevolucioncuentadescuento	 , "
	q += "	compradevolucionnombredescuento	 , "
	q += "	compracuentaproveedor	 , "
	q += "	compranombreproveedor	 , "
	q += "	compracuentabase	 , "
	q += "	compracuentaporcentajeretfte	 , "
	q += "	compracuentaporcentajeretotra	 , "
	q += "	cscuenta19	 , "
	q += "	csnombre19	 , "
	q += "	cscuenta5	 , "
	q += "	csnombre5	 , "
	q += "	cscuenta0	 , "
	q += "	csnombre0	 , "
	q += "	csiva19	 , "
	q += "	csnombreiva19	 , "
	q += "	csiva5	 , "
	q += "	csnombreiva5	 , "
	q += "	cscuentaretfte	 , "
	q += "	csnombreretfte	 , "
	q += "	cscuentaretica	 , "
	q += "	csnombreretica	 , "
	q += "	cscuentaretotra	 , "
	q += "	csnombreretotra	 , "
	q += "	cscuentadescuento	 , "
	q += "	csnombredescuento	 , "
	q += "	csdevolucioncuenta19	 , "
	q += "	csdevolucionnombre19	 , "
	q += "	csdevolucioncuenta5	 , "
	q += "	csdevolucionnombre5	 , "
	q += "	csdevolucioncuenta0	 , "
	q += "	csdevolucionnombre0	 , "
	q += "	csdevolucioniva19	 , "
	q += "	csdevolucionnombreiva19	 , "
	q += "	csdevolucioniva5	 , "
	q += "	csdevolucionnombreiva5	 , "
	q += "	csdevolucioncuentaretfte	 , "
	q += "	csdevolucionnombreretfte	 , "
	q += "	csdevolucioncuentaretica	 , "
	q += "	csdevolucionnombreretica	 , "
	q += "	csdevolucioncuentaretotra	 , "
	q += "	csdevolucionnombreretotra	 , "
	q += "	csdevolucioncuentadescuento	 , "
	q += "	csdevolucionnombredescuento	 , "
	q += "	cscuentaproveedor	 , "
	q += "	csnombreproveedor	 , "
	q += "	cscuentabase	 , "
	q += "	cscuentaporcentajeretfte	 , "
	q += "	cscuentaporcentajeretotra	 , "
	q += "	ventacuenta19	 , "
	q += "	ventanombre19	 , "
	q += "	ventacuenta5	 , "
	q += "	ventanombre5	 , "
	q += "	ventacuenta0	 , "
	q += "	ventanombre0	 , "
	q += "	ventaiva19	 , "
	q += "	ventanombreiva19	 , "
	q += "	ventaiva5	 , "
	q += "	ventanombreiva5	 , "
	q += "	ventacuentaret2201	 , "
	q += "	ventanombreret2201	 , "
	q += "	ventacuentadescuento	 , "
	q += "	ventanombredescuento	 , "
	q += "	ventadevolucioncuenta19	 , "
	q += "	ventadevolucionnombre19	 , "
	q += "	ventadevolucioncuenta5	 , "
	q += "	ventadevolucionnombre5	 , "
	q += "	ventadevolucioncuenta0	 , "
	q += "	ventadevolucionnombre0	 , "
	q += "	ventadevolucioniva19	 , "
	q += "	ventadevolucionnombreiva19	 , "
	q += "	ventadevolucioniva5	 , "
	q += "	ventadevolucionnombreiva5	 , "
	q += "	ventadevolucioncuentaret2201	 , "
	q += "	ventadevolucionnombreret2201	 , "
	q += "	ventadevolucioncuentadescuento	 , "
	q += "	ventadevolucionnombredescuento	 , "
	q += "	ventacuentacliente	 , "
	q += "	ventanombrecliente	 , "
	q += "	ventacontracuentaret2201	 , "
	q += "	ventacontranombreret2201	 , "
	q += "	ventadevolucioncontracuentaret2201	 , "
	q += "	ventadevolucioncontranombreret2201	 , "
	q += "	ventacuentaporcentajeret2201	 , "
	q += "	ventatipoiva	 , "
	q += "	ventaserviciocuenta19	 , "
	q += "	ventaservicionombre19	 , "
	q += "	ventaserviciocuenta5	 , "
	q += "	ventaservicionombre5	 , "
	q += "	ventaserviciocuenta0	 , "
	q += "	ventaservicionombre0	 , "
	q += "	ventaservicioiva19	 , "
	q += "	ventaservicionombreiva19	 , "
	q += "	ventaservicioiva5	 , "
	q += "	ventaservicionombreiva5	 , "
	q += "	ventaserviciocuentaret2201	 , "
	q += "	ventaservicionombreret2201	 , "
	q += "	ventaserviciocuentadescuento	 , "
	q += "	ventaservicionombredescuento	 , "
	q += "	ventaserviciodevolucioncuenta19	 , "
	q += "	ventaserviciodevolucionnombre19	 , "
	q += "	ventaserviciodevolucioncuenta5	 , "
	q += "	ventaserviciodevolucionnombre5	 , "
	q += "	ventaserviciodevolucioncuenta0	 , "
	q += "	ventaserviciodevolucionnombre0	 , "
	q += "	ventaserviciodevolucioniva19	 , "
	q += "	ventaserviciodevolucionnombreiva19	 , "
	q += "	ventaserviciodevolucioniva5	 , "
	q += "	ventaserviciodevolucionnombreiva5	 , "
	q += "	ventaserviciodevolucioncuentaret2201	 , "
	q += "	ventaserviciodevolucionnombreret2201	 , "
	q += "	ventaserviciodevolucioncuentadescuento	 , "
	q += "	ventaserviciodevolucionnombredescuento	 , "
	q += "	ventaserviciocuentacliente	 , "
	q += "	ventaservicionombrecliente	 , "
	q += "	ventaserviciocontracuentaret2201	 , "
	q += "	ventaserviciocontranombreret2201	 , "
	q += "	ventaserviciodevolucioncontracuentaret2201	 , "
	q += "	ventaserviciodevolucioncontranombreret2201	 , "
	q += "	ventaserviciocuentaporcentajeret2201	 , "
	q += "	ventaserviciotipoiva	 , "
	q += "	soportecuenta19	 , "
	q += "	soportenombre19	 , "
	q += "	soportecuenta5	 , "
	q += "	soportenombre5	 , "
	q += "	soportecuenta0	 , "
	q += "	soportenombre0	 , "
	q += "	soportecuentaretfte	 , "
	q += "	soportenombreretfte	 , "
	q += "	soportecuentaretica	 , "
	q += "	soportenombreretica	 , "
	q += "	soportecuentaretotra	 , "
	q += "	soportenombreretotra	 , "
	q += "	soportecuentadescuento	 , "
	q += "	soportenombredescuento	 , "
	q += "	soportedevolucioncuenta19	 , "
	q += "	soportedevolucionnombre19	 , "
	q += "	soportedevolucioncuenta5	 , "
	q += "	soportedevolucionnombre5	 , "
	q += "	soportedevolucioncuenta0	 , "
	q += "	soportedevolucionnombre0	 , "
	q += "	soportedevolucioncuentaretfte	 , "
	q += "	soportedevolucionnombreretfte	 , "
	q += "	soportedevolucioncuentaretica	 , "
	q += "	soportedevolucionnombreretica	 , "
	q += "	soportedevolucioncuentaretotra	 , "
	q += "	soportedevolucionnombreretotra	 , "
	q += "	soportedevolucioncuentadescuento	 , "
	q += "	soportedevolucionnombredescuento	 , "
	q += "	soportecuentaproveedor	 , "
	q += "	soportenombreproveedor	 , "
	q += "	soportecuentabase	 , "
	q += "	soportecuentaporcentajeretfte	 , "
	q += "	soportecuentaporcentajeretotra	 , "
	q += "	soporteserviciocuenta19	 , "
	q += "	soporteservicionombre19	 , "
	q += "	soporteserviciocuenta5	 , "
	q += "	soporteservicionombre5	 , "
	q += "	soporteserviciocuenta0	 , "
	q += "	soporteservicionombre0	 , "
	q += "	soporteserviciocuentaretfte	 , "
	q += "	soporteservicionombreretfte	 , "
	q += "	soporteserviciocuentaretica	 , "
	q += "	soporteservicionombreretica	 , "
	q += "	soporteserviciocuentaretotra	 , "
	q += "	soporteservicionombreretotra	 , "
	q += "	soporteserviciocuentadescuento	 , "
	q += "	soporteservicionombredescuento	 , "
	q += "	soporteserviciodevolucioncuenta	 , "
	q += "	soporteserviciodevolucionnombre	 , "
	q += "	soporteserviciodevolucioncuentaretfte	 , "
	q += "	soporteserviciodevolucionnombreretfte	 , "
	q += "	soporteserviciodevolucioncuentaretica	 , "
	q += "	soporteserviciodevolucionnombreretica	 , "
	q += "	soporteserviciodevolucioncuentaretotra	 , "
	q += "	soporteserviciodevolucionnombreretotra	 , "
	q += "	soporteserviciodevolucioncuentadescuento	 , "
	q += "	soporteserviciodevolucionnombredescuento	 , "
	q += "	soporteserviciocuentaproveedor	 , "
	q += "	soporteservicionombreproveedor	 , "
	q += "	soporteserviciocuentabase	 , "
	q += "	soporteserviciocuentaporcentajeretfte	 , "
	q += "	soporteserviciocuentaporcentajeretotra	 , "
	q += "	facturagastocuenta19	 , "
	q += "	facturagastonombre19	 , "
	q += "	facturagastocuenta5	 , "
	q += "	facturagastonombre5	 , "
	q += "	facturagastocuenta0	 , "
	q += "	facturagastonombre0	 , "
	q += "	facturagastocuentaretfte	 , "
	q += "	facturagastonombreretfte	 , "
	q += "	facturagastocuentaretica	 , "
	q += "	facturagastonombreretica	 , "
	q += "	facturagastocuentaretotra	 , "
	q += "	facturagastonombreretotra	 , "
	q += "	facturagastocuentadescuento	 , "
	q += "	facturagastonombredescuento	 , "
	q += "	facturagastodevolucioncuenta	 , "
	q += "	facturagastodevolucionnombre	 , "
	q += "	facturagastodevolucioncuentaretfte	 , "
	q += "	facturagastodevolucionnombreretfte	 , "
	q += "	facturagastodevolucioncuentaretica	 , "
	q += "	facturagastodevolucionnombreretica	 , "
	q += "	facturagastodevolucioncuentaretotra	 , "
	q += "	facturagastodevolucionnombreretotra	 , "
	q += "	facturagastodevolucioncuentadescuento	 , "
	q += "	facturagastodevolucionnombredescuento	 , "
	q += "	facturagastocuentaproveedor	 , "
	q += "	facturagastonombreproveedor	 , "
	q += "	facturagastocuentabase	 , "
	q += "	facturagastocuentaporcentajeretfte	 , "
	q += "	facturagastocuentaporcentajeretotra	 , "
	q += "	facturagastocuentaiva	 , "
	q += "	facturagastonombreiva	 , "
	q += "	facturagastoporcentajeiva	 , "
	q += "	cuentacosto	 , "
	q += "	cuentacostonombre	 , "
	q += "	cuentacostocontra	 , "
	q += "	cuentacostocontranombre  ,"
	q += "	decimales"
	q += " ) values ("
	q += "'14350501' ,"
	q += "'Compras Iva 19%' ,"
	q += "'14350502' ,"
	q += "'Compras Iva 5%' ,"
	q += "'14350503' ,"
	q += "'Compras Iva 0%' ,"
	q += "'24030501' ,"
	q += "'Iva 19% Compras' ,"
	q += "'24030302' ,"
	q += "'Iva 5% Compras' ,"
	q += "'24020701' ,"
	q += "'Ret. Fte. Compras 2.5%' ,"
	q += "'24040201' ,"
	q += "'I. C. A. Retenido' ,"
	q += "'24051201' ,"
	q += "'Otras Retenciones' ,"
	q += "'53053001' ,"
	q += "'Descuentos Concedidos Por Pronto Pago' ,"
	q += "'14350511' ,"
	q += "'Devolucion Compra 19%' ,"
	q += "'14350512' ,"
	q += "'Devolucion Compra 5%' ,"
	q += "'14350513' ,"
	q += "'Devolucion Compra 0%' ,"
	q += "'24030301' ,"
	q += "'Iva 19% Devoluciones Compras' ,"
	q += "'24030302' ,"
	q += "'Iva 5% Devoluciones Compras' ,"
	q += "'24020701' ,"
	q += "'Ret. Fte. Compras 2.5%' ,"
	q += "'24040201' ,"
	q += "'I. C. A. Retenido' ,"
	q += "'24051201' ,"
	q += "'Otras Retenciones' ,"
	q += "'53053001' ,"
	q += "'Descuentos Concedidos Por Pronto Pago' ,"
	q += "'23051001' ,"
	q += "'Proveedores Nacionales' ,"
	q += "'1,000,000' ,"
	q += "'2.50' ,"
	q += "'0' ,"
	q += "'62350501' ,"
	q += "'Compras Iva 19%' ,"
	q += "'62350502' ,"
	q += "'Compras Iva 5%' ,"
	q += "'62350503' ,"
	q += "'Compras Iva 0%' ,"
	q += "'24030601' ,"
	q += "'Iva 19% Compras' ,"
	q += "'24030602' ,"
	q += "'Iva 5% Compras' ,"
	q += "'24020401' ,"
	q += "'Ret. Fte. Servicios 4.0%' ,"
	q += "'24040201' ,"
	q += "'I. C. A. Retenido' ,"
	q += "'24051201' ,"
	q += "'Otras Retenciones' ,"
	q += "'62751001' ,"
	q += "'Descuentos Compras' ,"
	q += "'62750501' ,"
	q += "'Devolucion Compra 19%' ,"
	q += "'62750502' ,"
	q += "'Devolucion Compra 5%' ,"
	q += "'62750503' ,"
	q += "'Devolucion Compra 0%' ,"
	q += "'24030401' ,"
	q += "'Iva 19% Devoluciones Compras' ,"
	q += "'24030402' ,"
	q += "'Iva 5% Devoluciones Compras' ,"
	q += "'24020401' ,"
	q += "'Ret. Fte. Servicios 4.0%' ,"
	q += "'24040201' ,"
	q += "'I. C. A. Retenido' ,"
	q += "'24051201' ,"
	q += "'Otras Retenciones' ,"
	q += "'62751001' ,"
	q += "'Descuentos Compras' ,"
	q += "'23051001' ,"
	q += "'Proveedores Nacionales' ,"
	q += "'1,000,000' ,"
	q += "'4.00' ,"
	q += "'0' ,"
	q += "'41350501' ,"
	q += "'Ventas Iva 19%' ,"
	q += "'41350502' ,"
	q += "'Ventas Iva 5%' ,"
	q += "'41350503' ,"
	q += "'Ventas Iva 0%' ,"
	q += "'24030101' ,"
	q += "'Iva 19% Ventas' ,"
	q += "'24030102' ,"
	q += "'Iva 5% Ventas' ,"
	q += "'17050102' ,"
	q += "'Autorretencion 0.80% 2201 (Db)' ,"
	q += "'41751001' ,"
	q += "'Descuento En Ventas 19%' ,"
	q += "'41750501' ,"
	q += "'Devolucion En Ventas 19%' ,"
	q += "'41750502' ,"
	q += "'Devolucion En Ventas 5%' ,"
	q += "'41750503' ,"
	q += "'Devolucion En Ventas 0%' ,"
	q += "'24030701' ,"
	q += "'Iva 19% Devoluciones Ventas' ,"
	q += "'24030702' ,"
	q += "'Iva 5% Devoluciones Ventas' ,"
	q += "'17050102' ,"
	q += "'Autorretencion 0.80% 2201 (Db)' ,"
	q += "'41751001' ,"
	q += "'Descuento En Ventas 19%' ,"
	q += "'13060601' ,"
	q += "'Clientes Nacionales' ,"
	q += "'24021001' ,"
	q += "'Autorretencion 0.80% 2201' ,"
	q += "'24021001' ,"
	q += "'Autorretencion 0.80% 2201' ,"
	q += "'0.80' ,"
	q += "'' ,"
	q += "'42150501' ,"
	q += "'Venta De Servicios 19%' ,"
	q += "'42150502' ,"
	q += "'Venta De Servicios 5%' ,"
	q += "'42150503' ,"
	q += "'Venta De Servicios 0%' ,"
	q += "'24030201' ,"
	q += "'Iva 19% Venta Servicios' ,"
	q += "'24030202' ,"
	q += "'Iva 5% Venta Servicios' ,"
	q += "'17050102' ,"
	q += "'Autorretencion 0.80% 2201 (Db)' ,"
	q += "'42751001' ,"
	q += "'Descuento En Venta De Servicios 19%' ,"
	q += "'42750501' ,"
	q += "'Devolucion En Venta De Servicios 19%' ,"
	q += "'42750502' ,"
	q += "'Devolucion En Venta De Servicios 5%' ,"
	q += "'42750503' ,"
	q += "'Devolucion En Venta De Servicio 0%' ,"
	q += "'24030801' ,"
	q += "'Iva 19% Devolucion Venta Servicios' ,"
	q += "'24030802' ,"
	q += "'Iva 5% Devoluciones Venta Servicios' ,"
	q += "'17050102' ,"
	q += "'Autorretencion 0.80% 2201 (Db)' ,"
	q += "'42751001' ,"
	q += "'Descuento En Venta De Servicios 19%' ,"
	q += "'13060601' ,"
	q += "'Clientes Nacionales' ,"
	q += "'24021001' ,"
	q += "'Autorretencion 0.80% 2201' ,"
	q += "'24021001' ,"
	q += "'Autorretencion 0.80% 2201' ,"
	q += "'0.80' ,"
	q += "'' ,"
	q += "'14350501' ,"
	q += "'Compras Iva 19%' ,"
	q += "'14350502' ,"
	q += "'Compras Iva 5%' ,"
	q += "'14350503' ,"
	q += "'Compras Iva 0%' ,"
	q += "'24020701' ,"
	q += "'Ret. Fte. Compras 2.5%' ,"
	q += "'24040201' ,"
	q += "'I. C. A. Retenido' ,"
	q += "'24051201' ,"
	q += "'Otras Retenciones' ,"
	q += "'53053001' ,"
	q += "'Descuentos Concedidos Por Pronto Pago' ,"
	q += "'14350511' ,"
	q += "'Devolucion Compra 19%' ,"
	q += "'14350512' ,"
	q += "'Devolucion Compra 5%' ,"
	q += "'14350513' ,"
	q += "'Devolucion Compra 0%' ,"
	q += "'24020701' ,"
	q += "'Ret. Fte. Compras 2.5%' ,"
	q += "'24040201' ,"
	q += "'I. C. A. Retenido' ,"
	q += "'24051201' ,"
	q += "'Otras Retenciones' ,"
	q += "'53053001' ,"
	q += "'Descuentos Concedidos Por Pronto Pago' ,"
	q += "'23051001' ,"
	q += "'Proveedores Nacionales' ,"
	q += "'1,000,000' ,"
	q += "'2.50' ,"
	q += "'0' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'24040201' ,"
	q += "'I. C. A. Retenido' ,"
	q += "'24051201' ,"
	q += "'Otras Retenciones' ,"
	q += "'53053001' ,"
	q += "'Descuentos Concedidos Por Pronto Pago' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'24040201' ,"
	q += "'I. C. A. Retenido' ,"
	q += "'24051201' ,"
	q += "'Otras Retenciones' ,"
	q += "'53053001' ,"
	q += "'Descuentos Concedidos Por Pronto Pago' ,"
	q += "'23051001' ,"
	q += "'Proveedores Nacionales' ,"
	q += "'1,000,000' ,"
	q += "'' ,"
	q += "'0' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'' ,"
	q += "'293' , '0')"

	log.Println("comando inicial " + q)

	db := dbConn()
	_, err := db.Exec(q)
	if err != nil {
		panic(err.Error())
	}

}

// CONFIGURACIONINVENTARIO NUEVO
func ConfiguracioninventarioNuevo(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles("vista/inicio/Modulo.html",
		"vista/configuracioninventario/configuracioninventarioNuevo.html",
		"vista/autocompleta/autocompletaPlandecuentaempresa.html",
		"vista/configuracioninventario/configuracioninventariocompra.html",
		"vista/configuracioninventario/configuracioninventariocompraservicio.html",
		"vista/configuracioninventario/configuracioninventarioventa.html",
		"vista/configuracioninventario/configuracioninventarioventaservicio.html",
		"vista/configuracioninventario/configuracioninventariosoporte.html",
		"vista/configuracioninventario/configuracioninventariosoporteservicio.html",
		"vista/configuracioninventario/configuracioninventariofacturagasto.html",
		"vista/configuracioninventario/configuracioninventariogeneral.html",
	)

	fmt.Printf("%v, %v", tmp, err)
	log.Println("Error comprobante nuevo 3")

	db := dbConn()
	var total int
	row := db.QueryRow("SELECT COUNT(*) FROM configuracioninventario  ")
	err = row.Scan(&total)
	if err != nil {
		//	log.Fatal(err)
	}
	//var resultado bool
	if total > 0 {
		//resultado = true
	} else {
		log.Println("consultar configuracion inventario4444 ")
		Datosinicialesconfiguracioninventario()
		//resultado = false
	}

	//Compracuenta19 := mux.Vars(r)["compracuenta19"]
	panel := mux.Vars(r)["panel"]

	log.Println("consultar configuracion inventario0111122 ")

	//TRAE ESTRUCTURA DE LA TABLA
	emp := configuracioninventario{}
	err1 := db.Get(&emp, "SELECT * FROM configuracioninventario")
	if err1 != nil {
		log.Println("Error consulta " + err1.Error())
		log.Fatalln(err1)
	}

	log.Println("consultar configuracion inventario ")
	emp.Compracuentabase = Coma(emp.Compracuentabase)
	emp.Cscuentabase = Coma(emp.Cscuentabase)
	emp.Soportecuentabase = Coma(emp.Soportecuentabase)
	emp.Soporteserviciocuentabase = Coma(emp.Soporteserviciocuentabase)
	emp.Facturagastocuentabase = Coma(emp.Facturagastocuentabase)

	varmap := map[string]interface{}{
		"parametro": emp,
		"hosting":   ruta,
		"panel":     panel,
	}

	tmp.Execute(w, varmap)

}

// CONFIGURACIONINVENTARIO INSERTAR
func ConfiguracioninventarioInsertar(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	panel := r.FormValue("panel")
	if r.Method == "POST" {

		// INICIA INSERTAR COMPRA
		Compracuenta19 := r.FormValue("Compracuenta19")
		Compranombre19 := r.FormValue("Compranombre19")
		Compracuenta5 := r.FormValue("Compracuenta5")
		Compranombre5 := r.FormValue("Compranombre5")
		Compracuenta0 := r.FormValue("Compracuenta0")
		Compranombre0 := r.FormValue("Compranombre0")
		Compraiva19 := r.FormValue("Compraiva19")
		Compranombreiva19 := r.FormValue("Compranombreiva19")
		Compraiva5 := r.FormValue("Compraiva5")
		Compranombreiva5 := r.FormValue("Compranombreiva5")
		Compracuentaretfte := r.FormValue("Compracuentaretfte")
		Compranombreretfte := r.FormValue("Compranombreretfte")
		Compracuentaretica := r.FormValue("Compracuentaretica")
		Compranombreretica := r.FormValue("Compranombreretica")
		Compracuentaretotra := r.FormValue("Compracuentaretotra")
		Compranombreretotra := r.FormValue("Compranombreretotra")
		Compracuentadescuento := r.FormValue("Compracuentadescuento")
		Compranombredescuento := r.FormValue("Compranombredescuento")
		Compradevolucioncuenta19 := r.FormValue("Compradevolucioncuenta19")
		Compradevolucionnombre19 := r.FormValue("Compradevolucionnombre19")
		Compradevolucioncuenta5 := r.FormValue("Compradevolucioncuenta5")
		Compradevolucionnombre5 := r.FormValue("Compradevolucionnombre5")
		Compradevolucioncuenta0 := r.FormValue("Compradevolucioncuenta0")
		Compradevolucionnombre0 := r.FormValue("Compradevolucionnombre0")
		Compradevolucioniva19 := r.FormValue("Compradevolucioniva19")
		Compradevolucionnombreiva19 := r.FormValue("Compradevolucionnombreiva19")
		Compradevolucioniva5 := r.FormValue("Compradevolucioniva5")
		Compradevolucionnombreiva5 := r.FormValue("Compradevolucionnombreiva5")
		Compradevolucioncuentaretfte := r.FormValue("Compradevolucioncuentaretfte")
		Compradevolucionnombreretfte := r.FormValue("Compradevolucionnombreretfte")
		Compradevolucioncuentaretica := r.FormValue("Compradevolucioncuentaretica")
		Compradevolucionnombreretica := r.FormValue("Compradevolucionnombreretica")
		Compradevolucioncuentaretotra := r.FormValue("Compradevolucioncuentaretotra")
		Compradevolucionnombreretotra := r.FormValue("Compradevolucionnombreretotra")
		Compradevolucioncuentadescuento := r.FormValue("Compradevolucioncuentadescuento")
		Compradevolucionnombredescuento := r.FormValue("Compradevolucionnombredescuento")
		Compracuentaproveedor := r.FormValue("Compracuentaproveedor")
		Compranombreproveedor := r.FormValue("Compranombreproveedor")
		Compracuentabase := r.FormValue("Compracuentabase")
		Compracuentaporcentajeretfte := r.FormValue("Compracuentaporcentajeretfte")
		Compracuentaporcentajeretotra := r.FormValue("Compracuentaporcentajeretotra")
		Compranombre19 = Titulo(Compranombre19)
		Compranombre5 = Titulo(Compranombre5)
		Compranombre0 = Titulo(Compranombre0)
		Compranombreiva19 = Titulo(Compranombreiva19)
		Compranombreiva5 = Titulo(Compranombreiva5)
		Compranombreretfte = Titulo(Compranombreretfte)
		Compranombreretica = Titulo(Compranombreretica)
		Compranombreretotra = Titulo(Compranombreretotra)
		Compranombredescuento = Titulo(Compranombredescuento)
		Compradevolucionnombre19 = Titulo(Compradevolucionnombre19)
		Compradevolucionnombre5 = Titulo(Compradevolucionnombre5)
		Compradevolucionnombre0 = Titulo(Compradevolucionnombre0)
		Compradevolucionnombreiva19 = Titulo(Compradevolucionnombreiva19)
		Compradevolucionnombreiva5 = Titulo(Compradevolucionnombreiva5)
		Compradevolucionnombreretfte = Titulo(Compradevolucionnombreretfte)
		Compradevolucionnombreretica = Titulo(Compradevolucionnombreretica)
		Compradevolucionnombreretotra = Titulo(Compradevolucionnombreretotra)
		Compradevolucionnombredescuento = Titulo(Compradevolucionnombredescuento)
		Compranombreproveedor = Titulo(Compranombreproveedor)
		Compracuentabase = Quitacoma(Compracuentabase)
		// TERMINA INSERTAR COMPRA

		// INICIA INSERTAR COMPRA SERVICIOS
		Cscuenta19 := r.FormValue("Cscuenta19")
		Csnombre19 := r.FormValue("Csnombre19")
		Cscuenta5 := r.FormValue("Cscuenta5")
		Csnombre5 := r.FormValue("Csnombre5")
		Cscuenta0 := r.FormValue("Cscuenta0")
		Csnombre0 := r.FormValue("Csnombre0")
		Csiva19 := r.FormValue("Csiva19")
		Csnombreiva19 := r.FormValue("Csnombreiva19")
		Csiva5 := r.FormValue("Csiva5")
		Csnombreiva5 := r.FormValue("Csnombreiva5")
		Cscuentaretfte := r.FormValue("Cscuentaretfte")
		Csnombreretfte := r.FormValue("Csnombreretfte")
		Cscuentaretica := r.FormValue("Cscuentaretica")
		Csnombreretica := r.FormValue("Csnombreretica")
		Cscuentaretotra := r.FormValue("Cscuentaretotra")
		Csnombreretotra := r.FormValue("Csnombreretotra")
		Cscuentadescuento := r.FormValue("Cscuentadescuento")
		Csnombredescuento := r.FormValue("Csnombredescuento")
		Csdevolucioncuenta19 := r.FormValue("Csdevolucioncuenta19")
		Csdevolucionnombre19 := r.FormValue("Csdevolucionnombre19")
		Csdevolucioncuenta5 := r.FormValue("Csdevolucioncuenta5")
		Csdevolucionnombre5 := r.FormValue("Csdevolucionnombre5")
		Csdevolucioncuenta0 := r.FormValue("Csdevolucioncuenta0")
		Csdevolucionnombre0 := r.FormValue("Csdevolucionnombre0")
		Csdevolucioniva19 := r.FormValue("Csdevolucioniva19")
		Csdevolucionnombreiva19 := r.FormValue("Csdevolucionnombreiva19")
		Csdevolucioniva5 := r.FormValue("Csdevolucioniva5")
		Csdevolucionnombreiva5 := r.FormValue("Csdevolucionnombreiva5")
		Csdevolucioncuentaretfte := r.FormValue("Csdevolucioncuentaretfte")
		Csdevolucionnombreretfte := r.FormValue("Csdevolucionnombreretfte")
		Csdevolucioncuentaretica := r.FormValue("Csdevolucioncuentaretica")
		Csdevolucionnombreretica := r.FormValue("Csdevolucionnombreretica")
		Csdevolucioncuentaretotra := r.FormValue("Csdevolucioncuentaretotra")
		Csdevolucionnombreretotra := r.FormValue("Csdevolucionnombreretotra")
		Csdevolucioncuentadescuento := r.FormValue("Csdevolucioncuentadescuento")
		Csdevolucionnombredescuento := r.FormValue("Csdevolucionnombredescuento")
		Cscuentaproveedor := r.FormValue("Cscuentaproveedor")
		Csnombreproveedor := r.FormValue("Csnombreproveedor")
		Cscuentabase := r.FormValue("Cscuentabase")
		Cscuentaporcentajeretfte := r.FormValue("Cscuentaporcentajeretfte")
		Cscuentaporcentajeretotra := r.FormValue("Cscuentaporcentajeretotra")
		Csnombre19 = Titulo(Csnombre19)
		Csnombre5 = Titulo(Csnombre5)
		Csnombre0 = Titulo(Csnombre0)
		Csnombreiva19 = Titulo(Csnombreiva19)
		Csnombreiva5 = Titulo(Csnombreiva5)
		Csnombreretfte = Titulo(Csnombreretfte)
		Csnombreretica = Titulo(Csnombreretica)
		Csnombreretotra = Titulo(Csnombreretotra)
		Csnombredescuento = Titulo(Csnombredescuento)
		Csdevolucionnombre19 = Titulo(Csdevolucionnombre19)
		Csdevolucionnombre5 = Titulo(Csdevolucionnombre5)
		Csdevolucionnombre0 = Titulo(Csdevolucionnombre0)
		Csdevolucionnombreiva19 = Titulo(Csdevolucionnombreiva19)
		Csdevolucionnombreiva5 = Titulo(Csdevolucionnombreiva5)
		Csdevolucionnombreretfte = Titulo(Csdevolucionnombreretfte)
		Csdevolucionnombreretica = Titulo(Csdevolucionnombreretica)
		Csdevolucionnombreretotra = Titulo(Csdevolucionnombreretotra)
		Csdevolucionnombredescuento = Titulo(Csdevolucionnombredescuento)
		Csnombreproveedor = Titulo(Csnombreproveedor)
		Cscuentabase = Quitacoma(Cscuentabase)
		// TERMINA INSERTAR COMPRA SERVICIOS

		// INICIA INSERTAR VENTA
		Ventacuenta19 := r.FormValue("Ventacuenta19")
		Ventanombre19 := r.FormValue("Ventanombre19")
		Ventacuenta5 := r.FormValue("Ventacuenta5")
		Ventanombre5 := r.FormValue("Ventanombre5")
		Ventacuenta0 := r.FormValue("Ventacuenta0")
		Ventanombre0 := r.FormValue("Ventanombre0")
		Ventaiva19 := r.FormValue("Ventaiva19")
		Ventanombreiva19 := r.FormValue("Ventanombreiva19")
		Ventaiva5 := r.FormValue("Ventaiva5")
		Ventanombreiva5 := r.FormValue("Ventanombreiva5")
		Ventacuentaret2201 := r.FormValue("Ventacuentaret2201")
		Ventanombreret2201 := r.FormValue("Ventanombreret2201")
		Ventacuentadescuento := r.FormValue("Ventacuentadescuento")
		Ventanombredescuento := r.FormValue("Ventanombredescuento")
		Ventadevolucioncuenta19 := r.FormValue("Ventadevolucioncuenta19")
		Ventadevolucionnombre19 := r.FormValue("Ventadevolucionnombre19")
		Ventadevolucioncuenta5 := r.FormValue("Ventadevolucioncuenta5")
		Ventadevolucionnombre5 := r.FormValue("Ventadevolucionnombre5")
		Ventadevolucioncuenta0 := r.FormValue("Ventadevolucioncuenta0")
		Ventadevolucionnombre0 := r.FormValue("Ventadevolucionnombre0")
		Ventadevolucioniva19 := r.FormValue("Ventadevolucioniva19")
		Ventadevolucionnombreiva19 := r.FormValue("Ventadevolucionnombreiva19")
		Ventadevolucioniva5 := r.FormValue("Ventadevolucioniva5")
		Ventadevolucionnombreiva5 := r.FormValue("Ventadevolucionnombreiva5")
		Ventadevolucioncuentaret2201 := r.FormValue("Ventadevolucioncuentaret2201")
		Ventadevolucionnombreret2201 := r.FormValue("Ventadevolucionnombreret2201")
		Ventadevolucioncuentadescuento := r.FormValue("Ventadevolucioncuentadescuento")
		Ventadevolucionnombredescuento := r.FormValue("Ventadevolucionnombredescuento")
		Ventacuentacliente := r.FormValue("Ventacuentacliente")
		Ventanombrecliente := r.FormValue("Ventanombrecliente")
		Ventacontracuentaret2201 := r.FormValue("Ventacontracuentaret2201")
		Ventacontranombreret2201 := r.FormValue("Ventacontranombreret2201")
		Ventadevolucioncontracuentaret2201 := r.FormValue("Ventadevolucioncontracuentaret2201")
		Ventadevolucioncontranombreret2201 := r.FormValue("Ventadevolucioncontranombreret2201")
		Ventacuentaporcentajeret2201 := r.FormValue("Ventacuentaporcentajeret2201")
		Ventatipoiva := r.FormValue("Ventatipoiva")
		Ventanombre19 = Titulo(Ventanombre19)
		Ventanombre5 = Titulo(Ventanombre5)
		Ventanombre0 = Titulo(Ventanombre0)
		Ventanombreiva19 = Titulo(Ventanombreiva19)
		Ventanombreiva5 = Titulo(Ventanombreiva5)
		Ventanombreret2201 = Titulo(Ventanombreret2201)
		Ventanombredescuento = Titulo(Ventanombredescuento)
		Ventadevolucionnombre19 = Titulo(Ventadevolucionnombre19)
		Ventadevolucionnombre5 = Titulo(Ventadevolucionnombre5)
		Ventadevolucionnombre0 = Titulo(Ventadevolucionnombre0)
		Ventadevolucionnombreiva19 = Titulo(Ventadevolucionnombreiva19)
		Ventadevolucionnombreiva5 = Titulo(Ventadevolucionnombreiva5)
		Ventadevolucionnombreret2201 = Titulo(Ventadevolucionnombreret2201)
		Ventadevolucionnombredescuento = Titulo(Ventadevolucionnombredescuento)
		Ventanombrecliente = Titulo(Ventanombrecliente)
		Ventacontranombreret2201 = Titulo(Ventacontranombreret2201)
		Ventadevolucioncontranombreret2201 = Titulo(Ventadevolucioncontranombreret2201)
		// TERMINA INSERTAR VENTA

		// INICIA INSERTAR SERIVICIO
		Ventaserviciocuenta19 := r.FormValue("Ventaserviciocuenta19")
		Ventaservicionombre19 := r.FormValue("Ventaservicionombre19")
		Ventaserviciocuenta5 := r.FormValue("Ventaserviciocuenta5")
		Ventaservicionombre5 := r.FormValue("Ventaservicionombre5")
		Ventaserviciocuenta0 := r.FormValue("Ventaserviciocuenta0")
		Ventaservicionombre0 := r.FormValue("Ventaservicionombre0")
		Ventaservicioiva19 := r.FormValue("Ventaservicioiva19")
		Ventaservicionombreiva19 := r.FormValue("Ventaservicionombreiva19")
		Ventaservicioiva5 := r.FormValue("Ventaservicioiva5")
		Ventaservicionombreiva5 := r.FormValue("Ventaservicionombreiva5")
		Ventaserviciocuentaret2201 := r.FormValue("Ventaserviciocuentaret2201")
		Ventaservicionombreret2201 := r.FormValue("Ventaservicionombreret2201")
		Ventaserviciocuentadescuento := r.FormValue("Ventaserviciocuentadescuento")
		Ventaservicionombredescuento := r.FormValue("Ventaservicionombredescuento")
		Ventaserviciodevolucioncuenta19 := r.FormValue("Ventaserviciodevolucioncuenta19")
		Ventaserviciodevolucionnombre19 := r.FormValue("Ventaserviciodevolucionnombre19")
		Ventaserviciodevolucioncuenta5 := r.FormValue("Ventaserviciodevolucioncuenta5")
		Ventaserviciodevolucionnombre5 := r.FormValue("Ventaserviciodevolucionnombre5")
		Ventaserviciodevolucioncuenta0 := r.FormValue("Ventaserviciodevolucioncuenta0")
		Ventaserviciodevolucionnombre0 := r.FormValue("Ventaserviciodevolucionnombre0")
		Ventaserviciodevolucioniva19 := r.FormValue("Ventaserviciodevolucioniva19")
		Ventaserviciodevolucionnombreiva19 := r.FormValue("Ventaserviciodevolucionnombreiva19")
		Ventaserviciodevolucioniva5 := r.FormValue("Ventaserviciodevolucioniva5")
		Ventaserviciodevolucionnombreiva5 := r.FormValue("Ventaserviciodevolucionnombreiva5")
		Ventaserviciodevolucioncuentaret2201 := r.FormValue("Ventaserviciodevolucioncuentaret2201")
		Ventaserviciodevolucionnombreret2201 := r.FormValue("Ventaserviciodevolucionnombreret2201")
		Ventaserviciodevolucioncuentadescuento := r.FormValue("Ventaserviciodevolucioncuentadescuento")
		Ventaserviciodevolucionnombredescuento := r.FormValue("Ventaserviciodevolucionnombredescuento")
		Ventaserviciocuentacliente := r.FormValue("Ventaserviciocuentacliente")
		Ventaservicionombrecliente := r.FormValue("Ventaservicionombrecliente")
		Ventaserviciocontracuentaret2201 := r.FormValue("Ventaserviciocontracuentaret2201")
		Ventaserviciocontranombreret2201 := r.FormValue("Ventaserviciocontranombreret2201")
		Ventaserviciodevolucioncontracuentaret2201 := r.FormValue("Ventaserviciodevolucioncontracuentaret2201")
		Ventaserviciodevolucioncontranombreret2201 := r.FormValue("Ventaserviciodevolucioncontranombreret2201")
		Ventaserviciocuentaporcentajeret2201 := r.FormValue("Ventaserviciocuentaporcentajeret2201")
		Ventaserviciotipoiva := r.FormValue("Ventaserviciotipoiva")
		Ventaservicionombre19 = Titulo(Ventaservicionombre19)
		Ventaservicionombre5 = Titulo(Ventaservicionombre5)
		Ventaservicionombre0 = Titulo(Ventaservicionombre0)
		Ventaservicionombreiva19 = Titulo(Ventaservicionombreiva19)
		Ventaservicionombreiva5 = Titulo(Ventaservicionombreiva5)
		Ventaservicionombreret2201 = Titulo(Ventaservicionombreret2201)
		Ventaservicionombredescuento = Titulo(Ventaservicionombredescuento)
		Ventaserviciodevolucionnombre19 = Titulo(Ventaserviciodevolucionnombre19)
		Ventaserviciodevolucionnombre5 = Titulo(Ventaserviciodevolucionnombre5)
		Ventaserviciodevolucionnombre0 = Titulo(Ventaserviciodevolucionnombre0)
		Ventaserviciodevolucionnombreiva19 = Titulo(Ventaserviciodevolucionnombreiva19)
		Ventaserviciodevolucionnombreiva5 = Titulo(Ventaserviciodevolucionnombreiva5)
		Ventaserviciodevolucionnombreret2201 = Titulo(Ventaserviciodevolucionnombreret2201)
		Ventaserviciodevolucionnombredescuento = Titulo(Ventaserviciodevolucionnombredescuento)
		Ventaservicionombrecliente = Titulo(Ventaservicionombrecliente)
		Ventaserviciocontranombreret2201 = Titulo(Ventaserviciocontranombreret2201)
		Ventaserviciodevolucioncontranombreret2201 = Titulo(Ventaserviciodevolucioncontranombreret2201)
		// TERMINA INSERTAR SERVICIO

		// INICIA INSERTAR SOPORTE
		Soportecuenta19 := r.FormValue("Soportecuenta19")
		Soportenombre19 := r.FormValue("Soportenombre19")
		Soportecuenta5 := r.FormValue("Soportecuenta5")
		Soportenombre5 := r.FormValue("Soportenombre5")
		Soportecuenta0 := r.FormValue("Soportecuenta0")
		Soportenombre0 := r.FormValue("Soportenombre0")
		Soportecuentaretfte := r.FormValue("Soportecuentaretfte")
		Soportenombreretfte := r.FormValue("Soportenombreretfte")
		Soportecuentaretica := r.FormValue("Soportecuentaretica")
		Soportenombreretica := r.FormValue("Soportenombreretica")
		Soportecuentaretotra := r.FormValue("Soportecuentaretotra")
		Soportenombreretotra := r.FormValue("Soportenombreretotra")
		Soportecuentadescuento := r.FormValue("Soportecuentadescuento")
		Soportenombredescuento := r.FormValue("Soportenombredescuento")
		Soportedevolucioncuenta19 := r.FormValue("Soportedevolucioncuenta19")
		Soportedevolucionnombre19 := r.FormValue("Soportedevolucionnombre19")
		Soportedevolucioncuenta5 := r.FormValue("Soportedevolucioncuenta5")
		Soportedevolucionnombre5 := r.FormValue("Soportedevolucionnombre5")
		Soportedevolucioncuenta0 := r.FormValue("Soportedevolucioncuenta0")
		Soportedevolucionnombre0 := r.FormValue("Soportedevolucionnombre0")
		Soportedevolucioncuentaretfte := r.FormValue("Soportedevolucioncuentaretfte")
		Soportedevolucionnombreretfte := r.FormValue("Soportedevolucionnombreretfte")
		Soportedevolucioncuentaretica := r.FormValue("Soportedevolucioncuentaretica")
		Soportedevolucionnombreretica := r.FormValue("Soportedevolucionnombreretica")
		Soportedevolucioncuentaretotra := r.FormValue("Soportedevolucioncuentaretotra")
		Soportedevolucionnombreretotra := r.FormValue("Soportedevolucionnombreretotra")
		Soportedevolucioncuentadescuento := r.FormValue("Soportedevolucioncuentadescuento")
		Soportedevolucionnombredescuento := r.FormValue("Soportedevolucionnombredescuento")
		Soportecuentaproveedor := r.FormValue("Soportecuentaproveedor")
		Soportenombreproveedor := r.FormValue("Soportenombreproveedor")
		Soportecuentabase := r.FormValue("Soportecuentabase")
		Soportecuentaporcentajeretfte := r.FormValue("Soportecuentaporcentajeretfte")
		Soportecuentaporcentajeretotra := r.FormValue("Soportecuentaporcentajeretotra")
		Soportenombre19 = Titulo(Soportenombre19)
		Soportenombre5 = Titulo(Soportenombre5)
		Soportenombre0 = Titulo(Soportenombre0)
		Soportenombreretfte = Titulo(Soportenombreretfte)
		Soportenombreretica = Titulo(Soportenombreretica)
		Soportenombreretotra = Titulo(Soportenombreretotra)
		Soportenombredescuento = Titulo(Soportenombredescuento)
		Soportedevolucionnombre19 = Titulo(Soportedevolucionnombre19)
		Soportedevolucionnombre5 = Titulo(Soportedevolucionnombre5)
		Soportedevolucionnombre0 = Titulo(Soportedevolucionnombre0)
		Soportedevolucionnombreretfte = Titulo(Soportedevolucionnombreretfte)
		Soportedevolucionnombreretica = Titulo(Soportedevolucionnombreretica)
		Soportedevolucionnombreretotra = Titulo(Soportedevolucionnombreretotra)
		Soportedevolucionnombredescuento = Titulo(Soportedevolucionnombredescuento)
		Soportenombreproveedor = Titulo(Soportenombreproveedor)
		Soportecuentabase = Quitacoma(Soportecuentabase)
		// TERMINA INSERTAR SOPORTE

		// INICIA INSERTAR SOPORTE SERVICIOS
		Soporteserviciocuenta19 := r.FormValue("Soporteserviciocuenta19")
		Soporteservicionombre19 := r.FormValue("Soporteservicionombre19")
		Soporteserviciocuenta5 := r.FormValue("Soporteserviciocuenta5")
		Soporteservicionombre5 := r.FormValue("Soporteservicionombre5")
		Soporteserviciocuenta0 := r.FormValue("Soporteserviciocuenta0")
		Soporteservicionombre0 := r.FormValue("Soporteservicionombre0")
		Soporteserviciocuentaretfte := r.FormValue("Soporteserviciocuentaretfte")
		Soporteservicionombreretfte := r.FormValue("Soporteservicionombreretfte")
		Soporteserviciocuentaretica := r.FormValue("Soporteserviciocuentaretica")
		Soporteservicionombreretica := r.FormValue("Soporteservicionombreretica")
		Soporteserviciocuentaretotra := r.FormValue("Soporteserviciocuentaretotra")
		Soporteservicionombreretotra := r.FormValue("Soporteservicionombreretotra")
		Soporteserviciocuentadescuento := r.FormValue("Soporteserviciocuentadescuento")
		Soporteservicionombredescuento := r.FormValue("Soporteservicionombredescuento")
		Soporteserviciodevolucioncuenta := r.FormValue("Soporteserviciodevolucioncuenta")
		Soporteserviciodevolucionnombre := r.FormValue("Soporteserviciodevolucionnombre")
		Soporteserviciodevolucioncuentaretfte := r.FormValue("Soporteserviciodevolucioncuentaretfte")
		Soporteserviciodevolucionnombreretfte := r.FormValue("Soporteserviciodevolucionnombreretfte")
		Soporteserviciodevolucioncuentaretica := r.FormValue("Soporteserviciodevolucioncuentaretica")
		Soporteserviciodevolucionnombreretica := r.FormValue("Soporteserviciodevolucionnombreretica")
		Soporteserviciodevolucioncuentaretotra := r.FormValue("Soporteserviciodevolucioncuentaretotra")
		Soporteserviciodevolucionnombreretotra := r.FormValue("Soporteserviciodevolucionnombreretotra")
		Soporteserviciodevolucioncuentadescuento := r.FormValue("Soporteserviciodevolucioncuentadescuento")
		Soporteserviciodevolucionnombredescuento := r.FormValue("Soporteserviciodevolucionnombredescuento")
		Soporteserviciocuentaproveedor := r.FormValue("Soporteserviciocuentaproveedor")
		Soporteservicionombreproveedor := r.FormValue("Soporteservicionombreproveedor")
		Soporteserviciocuentabase := r.FormValue("Soporteserviciocuentabase")
		Soporteserviciocuentaporcentajeretfte := r.FormValue("Soporteserviciocuentaporcentajeretfte")
		Soporteserviciocuentaporcentajeretotra := r.FormValue("Soporteserviciocuentaporcentajeretotra")
		Soporteservicionombre19 = Titulo(Soporteservicionombre19)
		Soporteservicionombre5 = Titulo(Soporteservicionombre5)
		Soporteservicionombre0 = Titulo(Soporteservicionombre0)
		Soporteservicionombreretfte = Titulo(Soporteservicionombreretfte)
		Soporteservicionombreretica = Titulo(Soporteservicionombreretica)
		Soporteservicionombreretotra = Titulo(Soporteservicionombreretotra)
		Soporteservicionombredescuento = Titulo(Soporteservicionombredescuento)
		Soporteserviciodevolucionnombre = Titulo(Soporteserviciodevolucionnombre)
		Soporteserviciodevolucionnombreretfte = Titulo(Soporteserviciodevolucionnombreretfte)
		Soporteserviciodevolucionnombreretica = Titulo(Soporteserviciodevolucionnombreretica)
		Soporteserviciodevolucionnombreretotra = Titulo(Soporteserviciodevolucionnombreretotra)
		Soporteserviciodevolucionnombredescuento = Titulo(Soporteserviciodevolucionnombredescuento)
		Soporteservicionombreproveedor = Titulo(Soporteservicionombreproveedor)
		Soporteserviciocuentabase = Quitacoma(Soporteserviciocuentabase)
		// TERMINA INSERTAR SOPORTE SERVICIO

		// INICIA INSERTAR FACTURA GASTO
		Facturagastocuenta19 := r.FormValue("Facturagastocuenta19")
		Facturagastonombre19 := r.FormValue("Facturagastonombre19")
		Facturagastocuenta5 := r.FormValue("Facturagastocuenta5")
		Facturagastonombre5 := r.FormValue("Facturagastonombre5")
		Facturagastocuenta0 := r.FormValue("Facturagastocuenta0")
		Facturagastonombre0 := r.FormValue("Facturagastonombre0")
		Facturagastocuentaretfte := r.FormValue("Facturagastocuentaretfte")
		Facturagastonombreretfte := r.FormValue("Facturagastonombreretfte")
		Facturagastocuentaretica := r.FormValue("Facturagastocuentaretica")
		Facturagastonombreretica := r.FormValue("Facturagastonombreretica")
		Facturagastocuentaretotra := r.FormValue("Facturagastocuentaretotra")
		Facturagastonombreretotra := r.FormValue("Facturagastonombreretotra")
		Facturagastocuentadescuento := r.FormValue("Facturagastocuentadescuento")
		Facturagastonombredescuento := r.FormValue("Facturagastonombredescuento")
		Facturagastodevolucioncuenta := r.FormValue("Facturagastodevolucioncuenta")
		Facturagastodevolucionnombre := r.FormValue("Facturagastodevolucionnombre")
		Facturagastodevolucioncuentaretfte := r.FormValue("Facturagastodevolucioncuentaretfte")
		Facturagastodevolucionnombreretfte := r.FormValue("Facturagastodevolucionnombreretfte")
		Facturagastodevolucioncuentaretica := r.FormValue("Facturagastodevolucioncuentaretica")
		Facturagastodevolucionnombreretica := r.FormValue("Facturagastodevolucionnombreretica")
		Facturagastodevolucioncuentaretotra := r.FormValue("Facturagastodevolucioncuentaretotra")
		Facturagastodevolucionnombreretotra := r.FormValue("Facturagastodevolucionnombreretotra")
		Facturagastodevolucioncuentadescuento := r.FormValue("Facturagastodevolucioncuentadescuento")
		Facturagastodevolucionnombredescuento := r.FormValue("Facturagastodevolucionnombredescuento")
		Facturagastocuentaproveedor := r.FormValue("Facturagastocuentaproveedor")
		Facturagastonombreproveedor := r.FormValue("Facturagastonombreproveedor")
		Facturagastocuentabase := r.FormValue("Facturagastocuentabase")
		Facturagastocuentaporcentajeretfte := r.FormValue("Facturagastocuentaporcentajeretfte")
		Facturagastocuentaporcentajeretotra := r.FormValue("Facturagastocuentaporcentajeretotra")
		Facturagastonombre19 = Titulo(Facturagastonombre19)
		Facturagastonombre5 = Titulo(Facturagastonombre5)
		Facturagastonombre0 = Titulo(Facturagastonombre0)
		Facturagastonombreretfte = Titulo(Facturagastonombreretfte)
		Facturagastonombreretica = Titulo(Facturagastonombreretica)
		Facturagastonombreretotra = Titulo(Facturagastonombreretotra)
		Facturagastonombredescuento = Titulo(Facturagastonombredescuento)
		Facturagastodevolucionnombre = Titulo(Facturagastodevolucionnombre)
		Facturagastodevolucionnombreretfte = Titulo(Facturagastodevolucionnombreretfte)
		Facturagastodevolucionnombreretica = Titulo(Facturagastodevolucionnombreretica)
		Facturagastodevolucionnombreretotra = Titulo(Facturagastodevolucionnombreretotra)
		Facturagastodevolucionnombredescuento = Titulo(Facturagastodevolucionnombredescuento)
		Facturagastonombreproveedor = Titulo(Facturagastonombreproveedor)
		Facturagastocuentaiva := r.FormValue("Facturagastocuentaiva")
		Facturagastonombreiva := r.FormValue("Facturagastonombreiva")
		Facturagastoporcentajeiva := r.FormValue("Facturagastoporcentajeiva")
		Facturagastocuentabase = Quitacoma(Facturagastocuentabase)
		// TERMINA INSERTAR SOPORTE SERVICIO

		// INICIA INSERTAR COSTO DE VENTAS
		Cuentacosto := r.FormValue("Cuentacosto")
		Cuentacostonombre := r.FormValue("Cuentacostonombre")
		Cuentacostocontra := r.FormValue("Cuentacostocontra")
		Cuentacostocontranombre := r.FormValue("Cuentacostocontranombre")
		Decimales := r.FormValue("Decimales")
		// TERMINA INSERTAR COSTO DE VENTAS

		var consulta = "INSERT INTO configuracioninventario("

		// INICIA CONSULTA COMPRA-
		consulta += "Compracuenta19,"
		consulta += "Compranombre19,"
		consulta += "Compracuenta5,"
		consulta += "Compranombre5,"
		consulta += "Compracuenta0,"
		consulta += "Compranombre0,"
		consulta += "Compraiva19,"
		consulta += "Compranombreiva19,"
		consulta += "Compraiva5,"
		consulta += "Compranombreiva5,"
		consulta += "Compracuentaretfte,"
		consulta += "Compranombreretfte,"
		consulta += "Compracuentaretica,"
		consulta += "Compranombreretica,"
		consulta += "Compracuentaretotra,"
		consulta += "Compranombreretotra,"
		consulta += "Compracuentadescuento,"
		consulta += "Compranombredescuento,"
		consulta += "Compradevolucioncuenta19,"
		consulta += "Compradevolucionnombre19,"
		consulta += "Compradevolucioncuenta5,"
		consulta += "Compradevolucionnombre5,"
		consulta += "Compradevolucioncuenta0,"
		consulta += "Compradevolucionnombre0,"
		consulta += "Compradevolucioniva19,"
		consulta += "Compradevolucionnombreiva19,"
		consulta += "Compradevolucioniva5,"
		consulta += "Compradevolucionnombreiva5,"
		consulta += "Compradevolucioncuentaretfte,"
		consulta += "Compradevolucionnombreretfte,"
		consulta += "Compradevolucioncuentaretica,"
		consulta += "Compradevolucionnombreretica,"
		consulta += "Compradevolucioncuentaretotra,"
		consulta += "Compradevolucionnombreretotra,"
		consulta += "Compradevolucioncuentadescuento,"
		consulta += "Compradevolucionnombredescuento,"
		consulta += "Compracuentaproveedor,"
		consulta += "Compranombreproveedor,"
		consulta += "Compracuentabase,"
		consulta += "Compracuentaporcentajeretfte,"
		consulta += "Compracuentaporcentajeretotra,"
		// TERMINA CONSULTA COMPRA

		// INICIA CONSULTA COMPRA SERVICIOS
		consulta += "Cscuenta19,"
		consulta += "Csnombre19,"
		consulta += "Cscuenta5,"
		consulta += "Csnombre5,"
		consulta += "Cscuenta0,"
		consulta += "Csnombre0,"
		consulta += "Csiva19,"
		consulta += "Csnombreiva19,"
		consulta += "Csiva5,"
		consulta += "Csnombreiva5,"
		consulta += "Cscuentaretfte,"
		consulta += "Csnombreretfte,"
		consulta += "Cscuentaretica,"
		consulta += "Csnombreretica,"
		consulta += "Cscuentaretotra,"
		consulta += "Csnombreretotra,"
		consulta += "Cscuentadescuento,"
		consulta += "Csnombredescuento,"
		consulta += "Csdevolucioncuenta19,"
		consulta += "Csdevolucionnombre19,"
		consulta += "Csdevolucioncuenta5,"
		consulta += "Csdevolucionnombre5,"
		consulta += "Csdevolucioncuenta0,"
		consulta += "Csdevolucionnombre0,"
		consulta += "Csdevolucioniva19,"
		consulta += "Csdevolucionnombreiva19,"
		consulta += "Csdevolucioniva5,"
		consulta += "Csdevolucionnombreiva5,"
		consulta += "Csdevolucioncuentaretfte,"
		consulta += "Csdevolucionnombreretfte,"
		consulta += "Csdevolucioncuentaretica,"
		consulta += "Csdevolucionnombreretica,"
		consulta += "Csdevolucioncuentaretotra,"
		consulta += "Csdevolucionnombreretotra,"
		consulta += "Csdevolucioncuentadescuento,"
		consulta += "Csdevolucionnombredescuento,"
		consulta += "Cscuentaproveedor,"
		consulta += "Csnombreproveedor,"
		consulta += "Cscuentabase,"
		consulta += "Cscuentaporcentajeretfte,"
		consulta += "Cscuentaporcentajeretotra,"
		// TERMINA CONSULTA COMPRA SERVICIOS

		// INICIA CONSULTA VENTA
		consulta += "Ventacuenta19,"
		consulta += "Ventanombre19,"
		consulta += "Ventacuenta5,"
		consulta += "Ventanombre5,"
		consulta += "Ventacuenta0,"
		consulta += "Ventanombre0,"
		consulta += "Ventaiva19,"
		consulta += "Ventanombreiva19,"
		consulta += "Ventaiva5,"
		consulta += "Ventanombreiva5,"
		consulta += "Ventacuentaret2201,"
		consulta += "Ventanombreret2201,"
		consulta += "Ventacuentadescuento,"
		consulta += "Ventanombredescuento,"
		consulta += "Ventadevolucioncuenta19,"
		consulta += "Ventadevolucionnombre19,"
		consulta += "Ventadevolucioncuenta5,"
		consulta += "Ventadevolucionnombre5,"
		consulta += "Ventadevolucioncuenta0,"
		consulta += "Ventadevolucionnombre0,"
		consulta += "Ventadevolucioniva19,"
		consulta += "Ventadevolucionnombreiva19,"
		consulta += "Ventadevolucioniva5,"
		consulta += "Ventadevolucionnombreiva5,"
		consulta += "Ventadevolucioncuentaret2201,"
		consulta += "Ventadevolucionnombreret2201,"
		consulta += "Ventadevolucioncuentadescuento,"
		consulta += "Ventadevolucionnombredescuento,"
		consulta += "Ventacuentacliente,"
		consulta += "Ventanombrecliente,"
		consulta += "Ventacontracuentaret2201,"
		consulta += "Ventacontranombreret2201,"
		consulta += "Ventadevolucioncontracuentaret2201,"
		consulta += "Ventadevolucioncontranombreret2201,"
		consulta += "Ventacuentaporcentajeret2201,"
		consulta += "Ventatipoiva,"

		// TERMINA CONSULTA VENTA

		// INICIA CONSULTA SERVICIO
		consulta += "Ventaserviciocuenta19,"
		consulta += "Ventaservicionombre19,"
		consulta += "Ventaserviciocuenta5,"
		consulta += "Ventaservicionombre5,"
		consulta += "Ventaserviciocuenta0,"
		consulta += "Ventaservicionombre0,"
		consulta += "Ventaservicioiva19,"
		consulta += "Ventaservicionombreiva19,"
		consulta += "Ventaservicioiva5,"
		consulta += "Ventaservicionombreiva5,"
		consulta += "Ventaserviciocuentaret2201,"
		consulta += "Ventaservicionombreret2201,"
		consulta += "Ventaserviciocuentadescuento,"
		consulta += "Ventaservicionombredescuento,"
		consulta += "Ventaserviciodevolucioncuenta19,"
		consulta += "Ventaserviciodevolucionnombre19,"
		consulta += "Ventaserviciodevolucioncuenta5,"
		consulta += "Ventaserviciodevolucionnombre5,"
		consulta += "Ventaserviciodevolucioncuenta0,"
		consulta += "Ventaserviciodevolucionnombre0,"
		consulta += "Ventaserviciodevolucioniva19,"
		consulta += "Ventaserviciodevolucionnombreiva19,"
		consulta += "Ventaserviciodevolucioniva5,"
		consulta += "Ventaserviciodevolucionnombreiva5,"
		consulta += "Ventaserviciodevolucioncuentaret2201,"
		consulta += "Ventaserviciodevolucionnombreret2201,"
		consulta += "Ventaserviciodevolucioncuentadescuento,"
		consulta += "Ventaserviciodevolucionnombredescuento,"
		consulta += "Ventaserviciocuentacliente,"
		consulta += "Ventaservicionombrecliente,"
		consulta += "Ventaserviciocontracuentaret2201,"
		consulta += "Ventaserviciocontranombreret2201,"
		consulta += "Ventaserviciodevolucioncontracuentaret2201,"
		consulta += "Ventaserviciodevolucioncontranombreret2201,"
		consulta += "Ventaserviciocuentaporcentajeret2201,"
		consulta += "Ventaserviciotipoiva,"
		// INICIA CONSULTA SERVICIO

		// INICIA CONSULTA SOPORTE
		consulta += "Soportecuenta19,"
		consulta += "Soportenombre19,"
		consulta += "Soportecuenta5,"
		consulta += "Soportenombre5,"
		consulta += "Soportecuenta0,"
		consulta += "Soportenombre0,"
		consulta += "Soportecuentaretfte,"
		consulta += "Soportenombreretfte,"
		consulta += "Soportecuentaretica,"
		consulta += "Soportenombreretica,"
		consulta += "Soportecuentaretotra,"
		consulta += "Soportenombreretotra,"
		consulta += "Soportecuentadescuento,"
		consulta += "Soportenombredescuento,"
		consulta += "Soportedevolucioncuenta19,"
		consulta += "Soportedevolucionnombre19,"
		consulta += "Soportedevolucioncuenta5,"
		consulta += "Soportedevolucionnombre5,"
		consulta += "Soportedevolucioncuenta0,"
		consulta += "Soportedevolucionnombre0,"
		consulta += "Soportedevolucioncuentaretfte,"
		consulta += "Soportedevolucionnombreretfte,"
		consulta += "Soportedevolucioncuentaretica,"
		consulta += "Soportedevolucionnombreretica,"
		consulta += "Soportedevolucioncuentaretotra,"
		consulta += "Soportedevolucionnombreretotra,"
		consulta += "Soportedevolucioncuentadescuento,"
		consulta += "Soportedevolucionnombredescuento,"
		consulta += "Soportecuentaproveedor,"
		consulta += "Soportenombreproveedor,"
		consulta += "Soportecuentabase,"
		consulta += "Soportecuentaporcentajeretfte,"
		consulta += "Soportecuentaporcentajeretotra,"
		// TERMINA CONSULTA SOPORTE

		// INICIA CONSULTA SOPORTE
		consulta += "Soporteserviciocuenta19,"
		consulta += "Soporteservicionombre19,"
		consulta += "Soporteserviciocuenta5,"
		consulta += "Soporteservicionombre5,"
		consulta += "Soporteserviciocuenta0,"
		consulta += "Soporteservicionombre0,"
		consulta += "Soporteserviciocuentaretfte,"
		consulta += "Soporteservicionombreretfte,"
		consulta += "Soporteserviciocuentaretica,"
		consulta += "Soporteservicionombreretica,"
		consulta += "Soporteserviciocuentaretotra,"
		consulta += "Soporteservicionombreretotra,"
		consulta += "Soporteserviciocuentadescuento,"
		consulta += "Soporteservicionombredescuento,"
		consulta += "Soporteserviciodevolucioncuenta,"
		consulta += "Soporteserviciodevolucionnombre,"
		consulta += "Soporteserviciodevolucioncuentaretfte,"
		consulta += "Soporteserviciodevolucionnombreretfte,"
		consulta += "Soporteserviciodevolucioncuentaretica,"
		consulta += "Soporteserviciodevolucionnombreretica,"
		consulta += "Soporteserviciodevolucioncuentaretotra,"
		consulta += "Soporteserviciodevolucionnombreretotra,"
		consulta += "Soporteserviciodevolucioncuentadescuento,"
		consulta += "Soporteserviciodevolucionnombredescuento,"
		consulta += "Soporteserviciocuentaproveedor,"
		consulta += "Soporteservicionombreproveedor,"
		consulta += "Soporteserviciocuentabase,"
		consulta += "Soporteserviciocuentaporcentajeretfte,"
		consulta += "Soporteserviciocuentaporcentajeretotra,"

		// TERMINA CONSULTA SOPORTE SERVICIO

		// INICIA CONSULTA FACTURA GASTO
		consulta += "Facturagastocuenta19,"
		consulta += "Facturagastonombre19,"
		consulta += "Facturagastocuenta5,"
		consulta += "Facturagastonombre5,"
		consulta += "Facturagastocuenta0,"
		consulta += "Facturagastonombre0,"
		consulta += "Facturagastocuentaretfte,"
		consulta += "Facturagastonombreretfte,"
		consulta += "Facturagastocuentaretica,"
		consulta += "Facturagastonombreretica,"
		consulta += "Facturagastocuentaretotra,"
		consulta += "Facturagastonombreretotra,"
		consulta += "Facturagastocuentadescuento,"
		consulta += "Facturagastonombredescuento,"
		consulta += "Facturagastodevolucioncuenta,"
		consulta += "Facturagastodevolucionnombre,"
		consulta += "Facturagastodevolucioncuentaretfte,"
		consulta += "Facturagastodevolucionnombreretfte,"
		consulta += "Facturagastodevolucioncuentaretica,"
		consulta += "Facturagastodevolucionnombreretica,"
		consulta += "Facturagastodevolucioncuentaretotra,"
		consulta += "Facturagastodevolucionnombreretotra,"
		consulta += "Facturagastodevolucioncuentadescuento,"
		consulta += "Facturagastodevolucionnombredescuento,"
		consulta += "Facturagastocuentaproveedor,"
		consulta += "Facturagastonombreproveedor,"
		consulta += "Facturagastocuentabase,"
		consulta += "Facturagastocuentaporcentajeretfte,"
		consulta += "Facturagastocuentaporcentajeretotra,"
		consulta += "Facturagastocuentaiva,"
		consulta += "Facturagastonombreiva,"
		consulta += "Facturagastoporcentajeiva,"
		// TERMINA CONSULTA SOPORTE

		// INICIA CONSULTA COSTO DE VENTAS
		consulta += "Cuentacosto,"
		consulta += "Cuentacostonombre,"
		consulta += "Cuentacostocontra,"
		consulta += "Cuentacostocontranombre,"
		consulta += "Decimales"

		// INICIA VALORES COMPRA
		consulta += ")VALUES("
		consulta += parametros(253)
		consulta += ")"

		delForm, err := db.Prepare("DELETE from configuracioninventario")
		if err != nil {
			panic(err.Error())
		}
		delForm.Exec()
		insForm, err := db.Prepare(consulta)
		if err != nil {
			panic(err.Error())
		}
		_, err = insForm.Exec(

			// INICIA BORRAR COMPRA
			Compracuenta19,
			Compranombre19,
			Compracuenta5,
			Compranombre5,
			Compracuenta0,
			Compranombre0,
			Compraiva19,
			Compranombreiva19,
			Compraiva5,
			Compranombreiva5,
			Compracuentaretfte,
			Compranombreretfte,
			Compracuentaretica,
			Compranombreretica,
			Compracuentaretotra,
			Compranombreretotra,
			Compracuentadescuento,
			Compranombredescuento,
			Compradevolucioncuenta19,
			Compradevolucionnombre19,
			Compradevolucioncuenta5,
			Compradevolucionnombre5,
			Compradevolucioncuenta0,
			Compradevolucionnombre0,
			Compradevolucioniva19,
			Compradevolucionnombreiva19,
			Compradevolucioniva5,
			Compradevolucionnombreiva5,
			Compradevolucioncuentaretfte,
			Compradevolucionnombreretfte,
			Compradevolucioncuentaretica,
			Compradevolucionnombreretica,
			Compradevolucioncuentaretotra,
			Compradevolucionnombreretotra,
			Compradevolucioncuentadescuento,
			Compradevolucionnombredescuento,
			Compracuentaproveedor,
			Compranombreproveedor,
			Compracuentabase,
			Compracuentaporcentajeretfte,
			Compracuentaporcentajeretotra,
			// TERMINA BORRAR COMPRA

			// INICIA BORRAR COMPRA SERVICIOS
			Cscuenta19,
			Csnombre19,
			Cscuenta5,
			Csnombre5,
			Cscuenta0,
			Csnombre0,
			Csiva19,
			Csnombreiva19,
			Csiva5,
			Csnombreiva5,
			Cscuentaretfte,
			Csnombreretfte,
			Cscuentaretica,
			Csnombreretica,
			Cscuentaretotra,
			Csnombreretotra,
			Cscuentadescuento,
			Csnombredescuento,
			Csdevolucioncuenta19,
			Csdevolucionnombre19,
			Csdevolucioncuenta5,
			Csdevolucionnombre5,
			Csdevolucioncuenta0,
			Csdevolucionnombre0,
			Csdevolucioniva19,
			Csdevolucionnombreiva19,
			Csdevolucioniva5,
			Csdevolucionnombreiva5,
			Csdevolucioncuentaretfte,
			Csdevolucionnombreretfte,
			Csdevolucioncuentaretica,
			Csdevolucionnombreretica,
			Csdevolucioncuentaretotra,
			Csdevolucionnombreretotra,
			Csdevolucioncuentadescuento,
			Csdevolucionnombredescuento,
			Cscuentaproveedor,
			Csnombreproveedor,
			Cscuentabase,
			Cscuentaporcentajeretfte,
			Cscuentaporcentajeretotra,
			// TERMINA BORRAR COMPRA SERVICIOS

			// INICIA BORRAR VENTA
			Ventacuenta19,
			Ventanombre19,
			Ventacuenta5,
			Ventanombre5,
			Ventacuenta0,
			Ventanombre0,
			Ventaiva19,
			Ventanombreiva19,
			Ventaiva5,
			Ventanombreiva5,
			Ventacuentaret2201,
			Ventanombreret2201,
			Ventacuentadescuento,
			Ventanombredescuento,
			Ventadevolucioncuenta19,
			Ventadevolucionnombre19,
			Ventadevolucioncuenta5,
			Ventadevolucionnombre5,
			Ventadevolucioncuenta0,
			Ventadevolucionnombre0,
			Ventadevolucioniva19,
			Ventadevolucionnombreiva19,
			Ventadevolucioniva5,
			Ventadevolucionnombreiva5,
			Ventadevolucioncuentaret2201,
			Ventadevolucionnombreret2201,
			Ventadevolucioncuentadescuento,
			Ventadevolucionnombredescuento,
			Ventacuentacliente,
			Ventanombrecliente,
			Ventacontracuentaret2201,
			Ventacontranombreret2201,
			Ventadevolucioncontracuentaret2201,
			Ventadevolucioncontranombreret2201,
			Ventacuentaporcentajeret2201,
			Ventatipoiva,
			// TERMINA BORRAR VENTA

			// INICIA BORRAR SERVICIO
			Ventaserviciocuenta19,
			Ventaservicionombre19,
			Ventaserviciocuenta5,
			Ventaservicionombre5,
			Ventaserviciocuenta0,
			Ventaservicionombre0,
			Ventaservicioiva19,
			Ventaservicionombreiva19,
			Ventaservicioiva5,
			Ventaservicionombreiva5,
			Ventaserviciocuentaret2201,
			Ventaservicionombreret2201,
			Ventaserviciocuentadescuento,
			Ventaservicionombredescuento,
			Ventaserviciodevolucioncuenta19,
			Ventaserviciodevolucionnombre19,
			Ventaserviciodevolucioncuenta5,
			Ventaserviciodevolucionnombre5,
			Ventaserviciodevolucioncuenta0,
			Ventaserviciodevolucionnombre0,
			Ventaserviciodevolucioniva19,
			Ventaserviciodevolucionnombreiva19,
			Ventaserviciodevolucioniva5,
			Ventaserviciodevolucionnombreiva5,
			Ventaserviciodevolucioncuentaret2201,
			Ventaserviciodevolucionnombreret2201,
			Ventaserviciodevolucioncuentadescuento,
			Ventaserviciodevolucionnombredescuento,
			Ventaserviciocuentacliente,
			Ventaservicionombrecliente,
			Ventaserviciocontracuentaret2201,
			Ventaserviciocontranombreret2201,
			Ventaserviciodevolucioncontracuentaret2201,
			Ventaserviciodevolucioncontranombreret2201,
			Ventaserviciocuentaporcentajeret2201,
			Ventaserviciotipoiva,
			// TERMINA BORRAR SERVICIO

			// INICIA BORRAR SOPORTE
			Soportecuenta19,
			Soportenombre19,
			Soportecuenta5,
			Soportenombre5,
			Soportecuenta0,
			Soportenombre0,
			Soportecuentaretfte,
			Soportenombreretfte,
			Soportecuentaretica,
			Soportenombreretica,
			Soportecuentaretotra,
			Soportenombreretotra,
			Soportecuentadescuento,
			Soportenombredescuento,
			Soportedevolucioncuenta19,
			Soportedevolucionnombre19,
			Soportedevolucioncuenta5,
			Soportedevolucionnombre5,
			Soportedevolucioncuenta0,
			Soportedevolucionnombre0,
			Soportedevolucioncuentaretfte,
			Soportedevolucionnombreretfte,
			Soportedevolucioncuentaretica,
			Soportedevolucionnombreretica,
			Soportedevolucioncuentaretotra,
			Soportedevolucionnombreretotra,
			Soportedevolucioncuentadescuento,
			Soportedevolucionnombredescuento,
			Soportecuentaproveedor,
			Soportenombreproveedor,
			Soportecuentabase,
			Soportecuentaporcentajeretfte,
			Soportecuentaporcentajeretotra,
			// TERMINA BORRAR SOPORTE

			// INICIA BORRAR SOPORTE SERVICIO
			Soporteserviciocuenta19,
			Soporteservicionombre19,
			Soporteserviciocuenta5,
			Soporteservicionombre5,
			Soporteserviciocuenta0,
			Soporteservicionombre0,
			Soporteserviciocuentaretfte,
			Soporteservicionombreretfte,
			Soporteserviciocuentaretica,
			Soporteservicionombreretica,
			Soporteserviciocuentaretotra,
			Soporteservicionombreretotra,
			Soporteserviciocuentadescuento,
			Soporteservicionombredescuento,
			Soporteserviciodevolucioncuenta,
			Soporteserviciodevolucionnombre,
			Soporteserviciodevolucioncuentaretfte,
			Soporteserviciodevolucionnombreretfte,
			Soporteserviciodevolucioncuentaretica,
			Soporteserviciodevolucionnombreretica,
			Soporteserviciodevolucioncuentaretotra,
			Soporteserviciodevolucionnombreretotra,
			Soporteserviciodevolucioncuentadescuento,
			Soporteserviciodevolucionnombredescuento,
			Soporteserviciocuentaproveedor,
			Soporteservicionombreproveedor,
			Soporteserviciocuentabase,
			Soporteserviciocuentaporcentajeretfte,
			Soporteserviciocuentaporcentajeretotra,
			// TERMINA BORRAR SOPORTE

			// INICIA BORRAR FACTURA GASTO
			Facturagastocuenta19,
			Facturagastonombre19,
			Facturagastocuenta5,
			Facturagastonombre5,
			Facturagastocuenta0,
			Facturagastonombre0,
			Facturagastocuentaretfte,
			Facturagastonombreretfte,
			Facturagastocuentaretica,
			Facturagastonombreretica,
			Facturagastocuentaretotra,
			Facturagastonombreretotra,
			Facturagastocuentadescuento,
			Facturagastonombredescuento,
			Facturagastodevolucioncuenta,
			Facturagastodevolucionnombre,
			Facturagastodevolucioncuentaretfte,
			Facturagastodevolucionnombreretfte,
			Facturagastodevolucioncuentaretica,
			Facturagastodevolucionnombreretica,
			Facturagastodevolucioncuentaretotra,
			Facturagastodevolucionnombreretotra,
			Facturagastodevolucioncuentadescuento,
			Facturagastodevolucionnombredescuento,
			Facturagastocuentaproveedor,
			Facturagastonombreproveedor,
			Facturagastocuentabase,
			Facturagastocuentaporcentajeretfte,
			Facturagastocuentaporcentajeretotra,
			Facturagastocuentaiva,
			Facturagastonombreiva,
			Facturagastoporcentajeiva,
			// TERMINA BORRAR FACTURA GASTO

			// INICIA BORRAR COSTO
			Cuentacosto,
			Cuentacostonombre,
			Cuentacostocontra,
			Cuentacostocontranombre,
			Decimales)

		if err != nil {
			panic(err)
		}
		log.Println("Nuevo Registro:" + Compracuenta19 + "," + Compranombre19)
	}
	http.Redirect(w, r, "/ConfiguracioninventarioNuevo/"+panel, 301)
}

// INICIA CONFIGURACIONINVENTARIO PDF
func ConfiguracioninventarioPdf(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	Compracuenta19 := mux.Vars(r)["compracuenta19"]
	t := configuracioninventario{}
	var e empresa = ListaEmpresa()
	var c ciudad = TraerCiudad(e.Ciudad)
	err := db.Get(&t, "SELECT * FROM configuracioninventario where compracuenta19=$1", Compracuenta19)
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
	pdf.CellFormat(40, 4, t.Compracuenta19, "", 0,
		"", false, 0, "")
	pdf.Ln(-1)
	pdf.SetX(30)
	pdf.CellFormat(40, 4, "Nombre:", "", 0,
		"", false, 0, "")
	pdf.CellFormat(40, 4, t.Compranombre19, "", 0,
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

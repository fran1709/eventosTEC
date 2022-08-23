package main

/**
@author Francisco Ovares Rojas
@author Samantha
@startDate 21/08/2022
@endDate   --/09/2022
*/

import (
	"fmt"
	"mimodulo/com"
)

type Asientos = map[int32]com.Asiento
type Clientes = map[int32]com.Cliente
type Facturas = map[int32]com.Factura

var dAsientos Asientos //Diccionario de asientos
var dClientes Clientes //Diccionario de clientes
var dFacturas Facturas //Diccionario de facturas

var idAsiento int32 = 0 // id serial
var idCliente int32 = 0 // id serial
var idFactura int32 = 0 // id serial

/*
* Agrega un nuevo elemento "cliente" al diccionario de Clientes->dClientes
 */
func agregarCliente(pId int32, pNombre string, pApellido1 string, pApellido2 string) {
	dClientes[pId] = com.Cliente{IdCliente: pId, Nombre: pNombre, Apellido1: pApellido1, Apellido2: pApellido2}
}

/*
* Agrega un nuevo elemento "cliente" al diccionario de Factura->dFactura
 */
func agregarFactura(pId int32, pCliente com.Cliente, pAsiento com.Asiento, pPrecio int32) {
	dFacturas[pId] = com.Factura{IdFactura: pId, Cliente: pCliente, Asiento: pAsiento, Precio: pPrecio}
}

/*
*
Agrega un nuevo elemento "asiento" al diccionario de Asientos->dAsientos
*/
func agregarAsiento(pCategoria string, pZonas string, pNumero int32, pFila int16, pColum int16) {
	dAsientos[pNumero] = com.Asiento{Categoria: pCategoria, Zona: pZonas, Numero: pNumero, Fila: pFila, Columna: pColum}
}

func clientesData() {
	agregarCliente(idCliente, "Francisco", "Ovares", "Rojas")
	idCliente++
	agregarCliente(idCliente, "Josué", "Ovares", "Rojas")
	idCliente++
	agregarCliente(idCliente, "Thomas", "Ovares", "Molina")
	idCliente++
}
func facturasData() {
	//falta agregar facturas
}
func asientosData() {
	//falta agregar clientes
}
func cargarDatos() {
	clientesData()
	facturasData()
	asientosData()
}

func main() {
	cargarDatos()
	fmt.Println("Eventos Luna")
}

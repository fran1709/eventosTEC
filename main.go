package main

/**
@author Francisco Ovares Rojas
@author Samantha Acuña Montero
@startDate 21/08/2022
@endDate   --/09/2022
*/

import (
	"fmt"
	"mimodulo/com"
)

var dAsientos = make(map[int32]com.Asiento)
var dClientes = make(map[int32]com.Cliente)
var dFacturas = make(map[int32]com.Factura)

/*
*
var dAsientos Asientos //Diccionario de asientos
var dClientes Clientes //Diccionario de clientes
var dFacturas Facturas //Diccionario de facturas
*/
var numAsiento int32 = 0 // id serial
var idCliente int32 = 0  // id serial
var idFactura int32 = 0  // id serial

/*
Buscar cliente
*/
func buscarCliente(idC int32) bool {
	var encontrado bool
	for _, element := range dClientes { // Recorre el map por el valor de la key
		if element.IdCliente == idC {
			fmt.Println("Encontrado")
			encontrado = true
		} else {
			fmt.Println("No encontrado")
			encontrado = false
		}
	}
	return encontrado
}

/*
* Agrega un nuevo elemento "cliente" al diccionario de Clientes->dClientes
 */
func agregarCliente(pId int32, pNombre string, pApellido1 string, pApellido2 string) {
	cliente := buscarCliente(pId)
	if !cliente {
		dClientes[pId] = com.Cliente{IdCliente: pId, Nombre: pNombre, Apellido1: pApellido1, Apellido2: pApellido2}
		fmt.Print("Cliente agregado \n")
	} else {
		fmt.Print("Cliente no agregado, porque ya existe \n")
	}
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
	//agregarCliente(1, "Francisco", "Ovares", "Rojas")
	//idCliente++
	//agregarCliente(1, "Josué", "Ovares", "Rojas")
	//idCliente++
	//agregarCliente(idCliente, "Thomas", "Ovares", "Molina")
	//idCliente++
}

func asientosData() {
	agregarAsiento("VIP", "Palco", numAsiento, 1, 1)
	numAsiento++
	agregarAsiento("VIP", "Palco", numAsiento, 1, 2)
	numAsiento++
	agregarAsiento("Regular", "Sombra", numAsiento, 1, 1)
	numAsiento++
	agregarAsiento("Regular", "Sombra", numAsiento, 1, 2)
	numAsiento++
	agregarAsiento("Premiun", "Gramilla", numAsiento, 1, 1)
	numAsiento++
	agregarAsiento("Premiun", "Gramilla", numAsiento, 1, 2)
	numAsiento++
}

func facturasData() {
	agregarFactura(idFactura, dClientes[0], dAsientos[0], 40000)
	idFactura++
	agregarFactura(idFactura, dClientes[1], dAsientos[2], 30000)
	idFactura++
	agregarFactura(idFactura, dClientes[2], dAsientos[3], 20000)
	idFactura++
}

func cargarDatos() {
	clientesData()
	//asientosData()
	//facturasData()
}

func main() {
	/*cargarDatos()
	fmt.Println("Eventos Luna")
	fmt.Println("----Clientes----")
	fmt.Println(dClientes)
	fmt.Println("----Asientos----")
	fmt.Println(dAsientos)
	fmt.Println("----Facturas----")
	fmt.Println(dFacturas)*/
	agregarCliente(idCliente, "Sam", "Acunna", "Montero")
	idCliente++
	agregarCliente(idCliente, "Sam", "Acunna", "Montero")
	idCliente++
	fmt.Println(dClientes)
}

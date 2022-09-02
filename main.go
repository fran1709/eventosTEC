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

var dClientes = make(map[int32]com.Cliente)
var dFacturas = make(map[int32]com.Factura)
var dCategorias = make(map[string]com.Categoria)

var idCliente int32 = 0 // id serial
var idFactura int32 = 0 // id serial

/**
--------------------------------------------------------------------------------
							CLIENTES FUNCTIONS
--------------------------------------------------------------------------------
*/

/*
Buscar cliente
*/
func buscarCliente(idC int32) bool {
	var encontrado bool
	for _, element := range dClientes { // Recorre el map por el valor de la key
		if element.IdCliente == idC {
			//fmt.Println("Encontrado")
			encontrado = true
		} else {
			//fmt.Println("No encontrado")
			encontrado = false
		}
	}
	return encontrado
}

/*
Agrega un nuevo elemento "cliente" al diccionario de Clientes->dClientes
*/
func agregarCliente(pId int32, pNombre string, pApellido1 string, pApellido2 string) {
	cliente := buscarCliente(pId)
	if !cliente {
		dClientes[pId] = com.Cliente{IdCliente: pId, Nombre: pNombre, Apellido1: pApellido1, Apellido2: pApellido2}
		//fmt.Print("Cliente agregado \n")
	} else {
		//fmt.Print("Cliente no agregado, porque ya existe \n")
	}
}

func clientesData() {
	agregarCliente(idCliente, "Francisco", "Ovares", "Rojas")
	idCliente++
	agregarCliente(idCliente, "Josué", "Ovares", "Rojas")
	idCliente++
	agregarCliente(idCliente, "Thomas", "Ovares", "Molina")
	idCliente++
}

/**
--------------------------------------------------------------------------------
							FACTURACION FUNCTIONS
--------------------------------------------------------------------------------
*/
/*
Buscar factura
*/
func buscarFactura(idF int32) bool {
	var encontrado bool
	for _, element := range dFacturas { // Recorre el map por el valor de la key
		if element.IdFactura == idF {
			//fmt.Println("Encontrado")
			encontrado = true
		} else {
			//fmt.Println("No encontrado")
			encontrado = false
		}
	}
	return encontrado
}

/*
Agrega un nuevo elemento "cliente" al diccionario de Factura->dFactura
*/
func agregarFactura(pId int32, pCliente com.Cliente, pAsiento com.Asiento, pPrecio int32) {
	factura := buscarFactura(pId)
	if !factura {
		dFacturas[pId] = com.Factura{IdFactura: pId, Cliente: pCliente, Asiento: pAsiento, Precio: pPrecio}
		//fmt.Print("Factura agregada \n")
	} else {
		//fmt.Print("Factura existente no agregado \n")
	}
}

func facturasData() {
	/*
		agregarFactura(idFactura, dClientes[0], dAsientos[0], 40000)
		idFactura++
		agregarFactura(idFactura, dClientes[1], dAsientos[2], 30000)
		idFactura++
		agregarFactura(idFactura, dClientes[2], dAsientos[3], 20000)
		idFactura++*/
}

/**
--------------------------------------------------------------------------------
							ASIENTOS FUNCTIONS
--------------------------------------------------------------------------------
*/
/*
Buscar asiento
*/
func buscarAsiento(Categoria string, Zona string, fila int16, columna int16) bool {
	var encontrado bool
	if (dCategorias[Categoria].ZonaA[fila][columna].Categoria == Categoria) && (dCategorias[Categoria].ZonaA[fila][columna].Zona == Zona) { // Asegura la categoria y la zona del asiento
		if (dCategorias[Categoria].ZonaA[fila][columna].Fila == fila) && (dCategorias[Categoria].ZonaA[fila][columna].Columna == columna) { // Asegura la fila y columna
			fmt.Print("\nAsiento: ", dCategorias[Categoria].ZonaA[fila][columna].Fila, dCategorias[Categoria].ZonaA[fila][columna].Columna, " Zona: ", dCategorias[Categoria].ZonaA[fila][columna].Zona)
			encontrado = true
		} else {
			fmt.Println("Asiento no encontrado")
			encontrado = false
		}
	} else if (dCategorias[Categoria].ZonaB[fila][columna].Categoria == Categoria) && (dCategorias[Categoria].ZonaB[fila][columna].Zona == Zona) { // Asegura la categoria y la zona del asiento
		if (dCategorias[Categoria].ZonaB[fila][columna].Fila == fila) && (dCategorias[Categoria].ZonaB[fila][columna].Columna == columna) { // Asegura la fila y columna
			fmt.Print("\nAsiento: ", dCategorias[Categoria].ZonaB[fila][columna].Fila, dCategorias[Categoria].ZonaB[fila][columna].Columna, " Zona: ", dCategorias[Categoria].ZonaB[fila][columna].Zona)
			encontrado = true
		} else {
			fmt.Println("Asiento no encontrado")
			encontrado = false
		}
	} else if (dCategorias[Categoria].ZonaC[fila][columna].Categoria == Categoria) && (dCategorias[Categoria].ZonaC[fila][columna].Zona == Zona) { // Asegura la categoria y la zona del asiento
		if (dCategorias[Categoria].ZonaC[fila][columna].Fila == fila) && (dCategorias[Categoria].ZonaC[fila][columna].Columna == columna) { // Asegura la fila y columna
			fmt.Print("\nAsiento: ", dCategorias[Categoria].ZonaC[fila][columna].Fila, dCategorias[Categoria].ZonaC[fila][columna].Columna, " Zona: ", dCategorias[Categoria].ZonaC[fila][columna].Zona)
			encontrado = true
		} else {
			fmt.Println("Asiento no encontrado")
			encontrado = false
		}
	}
	return encontrado
}

/*
Disponibilidad del asiento:

	1 - Disponible
	0 - Reservado
	-1 - Comprado
*/
func disponibilidad(Categoria string, Zona string, fila int16, columna int16) bool {
	var disponibilidadAsiento bool
	asiento := buscarAsiento(Categoria, Zona, fila, columna)
	if asiento {
		if dCategorias[Categoria].ZonaB[fila][columna].Estado == 1 {
			fmt.Println("\nDisponible")
			disponibilidadAsiento = true
		} else {
			fmt.Println("\nEl asiento se encuentra reservado o comprado")
			disponibilidadAsiento = false
		}
	}
	return disponibilidadAsiento
}

/*
Crea una matriz de asientos según categoria y zona parametreadas con una tamaño constante.
*/
func crearAsientos(pCat string, pZona string) [4][5]com.Asiento {
	const filas = 4
	const colums = 5
	var asientos [filas][colums]com.Asiento
	for i := 0; i < filas; i++ {
		for j := 0; j < colums; j++ {
			asientos[i][j] = com.Asiento{Categoria: pCat, Zona: pZona, Fila: int16(i), Columna: int16(j), Estado: 1}
		}
	}
	return asientos
}

/*
*
--------------------------------------------------------------------------------

	CATEGORIAS FUNCTIONS

--------------------------------------------------------------------------------
*/
func agregarCategoria(pCate string, pZonaA [4][5]com.Asiento, pZonaB [4][5]com.Asiento, pZonaC [4][5]com.Asiento) {
	dCategorias[pCate] = com.Categoria{ZonaA: pZonaA, ZonaB: pZonaB, ZonaC: pZonaC}
}

func categoriaData() {
	agregarCategoria("VIP", crearAsientos("VIP", "A"), crearAsientos("VIP", "B"), crearAsientos("VIP", "C"))
	agregarCategoria("GRAMILLA", crearAsientos("GRAMILLA", "A"), crearAsientos("GRAMILLA", "B"), crearAsientos("GRAMILLA", "C"))
	agregarCategoria("PALCO", crearAsientos("PALCO", "A"), crearAsientos("PALCO", "B"), crearAsientos("PALCO", "C"))
	agregarCategoria("SOMBRA", crearAsientos("SOMBRA", "A"), crearAsientos("SOMBRA", "B"), crearAsientos("SOMBRA", "C"))
}

/*
*
--------------------------------------------------------------------------------

	MAIN FUNCTIONS

--------------------------------------------------------------------------------
*/
func motorDeBusqueda() {

}

func cargarDatos() {
	clientesData()
	facturasData()
	categoriaData()
}

func main() {
	cargarDatos()
	/*fmt.Println("Eventos Luna")
	fmt.Println("----Clientes----")
	fmt.Println(dClientes)
	fmt.Println("----Facturas----")
	fmt.Println(dFacturas)
	fmt.Println("----Categorias----")
	fmt.Println(dCategorias)*/
	// dCategorias["nombreCategoria"].zona[fila][#asiento]]
	fmt.Println(disponibilidad("VIP", "A", 0, 3))

}

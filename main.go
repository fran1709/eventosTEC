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
	"strings"
)

var dClientes = make(map[int32]com.Cliente)
var dFacturas = make(map[int32]com.Factura)
var dCategorias = make(map[string]com.Categoria)

var idCliente int32 = 0 // id serial
var idFactura int32 = 0 // id serial

type ListaAsientos []com.Asiento

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

/*
*
--------------------------------------------------------------------------------

	ASIENTOS FUNCTIONS

--------------------------------------------------------------------------------
*/
func getSizeList(pList *[][]com.Asiento) int {
	var cont = 0
	for i := range *pList {
		cont += len((*pList)[i])
	}
	//fmt.Println(cont)
	return cont
}

/*
*
Mapea todos los asientos vip de una categoría y los devuelve en una sola lista de asientos.
*/
func mapAsientos(pList *[][]com.Asiento, f func(pAsiento com.Asiento) com.Asiento) *[]com.Asiento {
	size := getSizeList(pList)
	mapped := make([]com.Asiento, size)
	for i := range *pList {
		for _, e := range (*pList)[i] {
			//fmt.Println(e)
			size--
			mapped[size] = f(e)

		}
	}
	//fmt.Println("mapped\n", mapped)
	return &mapped
}

/*
Retorna los asientos solicitados, que se encuentren disponibles.
*/
func filterAsientos(pLista *[]com.Asiento, pCantidad int) ListaAsientos {
	var mejoresOptions ListaAsientos
	var conta = 0
	for _, e := range *pLista {
		if conta == pCantidad {
			break
		}
		if e.Estado == 1 {
			mejoresOptions = append(mejoresOptions, e)
			conta++
		}
	}
	return mejoresOptions
}

/*
*
Retorna la matriz de los asientos de una categoria & zona respectiva.
*/
func obtenerAsientos(pCategoria string, pZona string) *[][]com.Asiento {
	switch {
	case "a" == strings.ToLower(pZona):
		return dCategorias[strings.ToUpper(pCategoria)].ZonaA
	case "b" == strings.ToLower(pZona):
		return dCategorias[strings.ToUpper(pCategoria)].ZonaB
	case "c" == strings.ToLower(pZona):
		return dCategorias[strings.ToUpper(pCategoria)].ZonaC
	}
	return &[][]com.Asiento{}
}

func mejoresOpciones(pCategoria string, pZona string, pCantidad int) {
	var matrixAsientos = obtenerAsientos(pCategoria, pZona)
	var asientos = mapAsientos(matrixAsientos, func(p com.Asiento) com.Asiento {
		if p.Zona == strings.ToUpper(pZona) {
			return p
		} else {
			return com.Asiento{}
		}
	})
	asientosfiltrados := filterAsientos(asientos, pCantidad)
	//fmt.Println("asientos mapeados:\n", asientos)
	fmt.Println("asientos filtrados:\n", asientosfiltrados)
}

/*
Disponibilidad del asiento:

	1 - Disponible
	0 - Reservado
	-1 - Comprado
*/

/*
Cambia la disponibilidad del asiento segun se requiera
*/
func cambiarDisponibilidad(pCate string, pZona string, pFila int, pAsiento int, pDispo int8) {
	switch {
	case "a" == strings.ToLower(pZona):
		(*dCategorias[strings.ToUpper(pCate)].ZonaA)[pFila][pAsiento].Estado = pDispo
	case "b" == strings.ToLower(pZona):
		(*dCategorias[strings.ToUpper(pCate)].ZonaB)[pFila][pAsiento].Estado = pDispo
	case "c" == strings.ToLower(pZona):
		(*dCategorias[strings.ToUpper(pCate)].ZonaC)[pFila][pAsiento].Estado = pDispo
	}
}

/*
Retorna la diponibilidad del asiento
*/
func disponibilidadAsiento(pCate string, pZona string, pFila int, pAsiento int) int8 {
	switch {
	case "a" == strings.ToLower(pZona):
		return (*dCategorias[strings.ToUpper(pCate)].ZonaA)[pFila][pAsiento].Estado
	case "b" == strings.ToLower(pZona):
		return (*dCategorias[strings.ToUpper(pCate)].ZonaB)[pFila][pAsiento].Estado
	case "c" == strings.ToLower(pZona):
		return (*dCategorias[strings.ToUpper(pCate)].ZonaC)[pFila][pAsiento].Estado
	}
	return -2
}

func comprarAsiento() {

}

/*
Crea una matriz de asientos según categoria y zona parametreadas con una tamaño constante.
*/
func crearAsientos(pCat string, pZona string, pFilas int, pAsientos int) *[][]com.Asiento {

	var asientos = make([][]com.Asiento, pFilas)
	//fmt.Println("salidos del horno\n", asientos)
	for i := 0; i < pFilas; i++ {
		asientos[i] = make([]com.Asiento, pAsientos)
	}
	//fmt.Println(asientos)
	for i := 0; i < pFilas; i++ {
		for j := 0; j < pAsientos; j++ {
			asientos[i][j] = com.Asiento{Categoria: pCat, Zona: pZona, Fila: int16(i), Columna: int16(j), Estado: 1}
		}
	}
	//fmt.Println(asientos)
	return &asientos
}

/*
*
--------------------------------------------------------------------------------

	CATEGORIAS FUNCTIONS

--------------------------------------------------------------------------------
*/
func agregarCategoria(pCate string, pZonaA *[][]com.Asiento, pZonaB *[][]com.Asiento, pZonaC *[][]com.Asiento) {
	dCategorias[pCate] = com.Categoria{ZonaA: pZonaA, ZonaB: pZonaB, ZonaC: pZonaC}
}

func categoriaData() {
	agregarCategoria("VIP", crearAsientos("VIP", "A", 2, 10), crearAsientos("VIP", "B", 4, 20), crearAsientos("VIP", "C", 5, 30))
	agregarCategoria("GRAMILLA", crearAsientos("GRAMILLA", "A", 2, 10), crearAsientos("GRAMILLA", "B", 4, 20), crearAsientos("GRAMILLA", "C", 5, 30))
	agregarCategoria("PALCO", crearAsientos("PALCO", "A", 2, 10), crearAsientos("PALCO", "B", 4, 20), crearAsientos("PALCO", "C", 5, 30))
	agregarCategoria("SOMBRA", crearAsientos("SOMBRA", "A", 2, 10), crearAsientos("SOMBRA", "B", 4, 20), crearAsientos("SOMBRA", "C", 5, 30))

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
	fmt.Println("----Categorias----")*/
	//fmt.Println(dCategorias)
	// dCategorias["nombreCategoria"].zona[fila][#asiento]]
	//fmt.Println(disponibilidad("VIP", "A", 0, 3))
	mejoresOpciones("vip", "a", 4)
	cambiarDisponibilidad("vip", "a", 1, 9, -1)
	cambiarDisponibilidad("vip", "a", 1, 7, -1)
	mejoresOpciones("vip", "a", 4)
}

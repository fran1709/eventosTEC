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
	"net"
	"os"
	"strconv"
	"strings"
)

const (
	SERVER_HOST = "127.0.0.1"
	SERVER_PORT = "3000"
	SERVER_TYPE = "tcp"
)

var dClientes = make(map[string]com.Cliente)
var dFacturas = make(map[int32]com.Factura)
var dCategorias = make(map[string]com.Categoria)

var idFactura int32 = 0 // id serial

type ListaAsientos []com.Asiento

/**
--------------------------------------------------------------------------------
							CLIENTES FUNCTIONS
--------------------------------------------------------------------------------
*/

/*
Buscar cliente
return true El cliente existe en sistema
return false El cliente no existe en sistema.
*/
func isCliente(pCedula string) bool {
	if _, isFound := dClientes[pCedula]; isFound {
		//fmt.Println("El cliente ", el, ", existe en sistema: ", isFound)
		return isFound
	}
	return false
}

func datosCliente(pCedula string) string {
	cadena := ""

	return cadena
}

/*
Agrega un nuevo elemento "cliente" al diccionario de Clientes->dClientes
*/
func agregarCliente(pCedula string, pNombreCompleto string, pCorreo string) {
	cliente := isCliente(pCedula)
	if !cliente {
		dClientes[pCedula] = com.Cliente{Cedula: pCedula, NombreCompleto: pNombreCompleto, Correo: pCorreo}
		//fmt.Print("Cliente agregado \n")
	} else {
		//fmt.Print("Cliente no agregado, porque ya existe \n")
	}
}

func clientesData() {
	agregarCliente("207710202", "Francisco Ovares Rojas", "fran1709@estudiantec.cr")
	agregarCliente("207710203", "Josue Ovares Rojas", "josue1908@estudiantec.cr")
	agregarCliente("207710205", "Thomas Ovares Molina", "thom1105@estudiantec.cr")
	agregarCliente("207710205", "Thomas Ovares Molina", "thom1105@estudiantec.cr")
}

/**
--------------------------------------------------------------------------------
							FACTURACION FUNCTIONS
--------------------------------------------------------------------------------
*/
/*
Buscar factura
return true La factura existe en sistema
return false La factura no existe en sistema.
*/
func isFactura(pId int32) bool {
	if _, isFound := dFacturas[pId]; isFound {
		return isFound
	}
	return false
}

/*
Agrega un nuevo elemento "cliente" al diccionario de Factura->dFactura
*/
func agregarFactura(pId int32, pCliente com.Cliente, pAsiento com.Asiento, pPrecio int32) {
	factura := isFactura(pId)
	if !factura {
		dFacturas[pId] = com.Factura{IdFactura: pId, Cliente: pCliente, Asiento: pAsiento, Precio: pPrecio}
		//fmt.Print("Factura agregada \n")
	} else {
		//fmt.Print("Factura existente no agregado \n")
	}
}

func comprarAsiento(pConnection net.Conn, pCategory string, pZona string, pCantidad int) {
	fmt.Println("Comprando tickets...")
	buffer := make([]byte, 1024)

	pConnection.Write([]byte("\n1.Sugerencias.\n2.Listado Disponible.\n"))
	read, _ := pConnection.Read(buffer)
	var choice = string(buffer[:read])
	switch choice {
	case "1":
		pConnection.Write([]byte(mejoresOpciones(pCategory, pZona, pCantidad)))
	case "2":
		pConnection.Write([]byte(mostrarAsientos(pCategory, pZona)))
	}
	return
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

func mejoresOpciones(pCategoria string, pZona string, pCantidad int) string {
	var cadena = ""
	var zona = ""
	switch pZona {
	case "1":
		zona = "a"
	case "2":
		zona = "b"
	case "3":
		zona = "c"
	}
	var matrixAsientos = obtenerAsientos(pCategoria, zona)
	var asientos = mapAsientos(matrixAsientos, func(p com.Asiento) com.Asiento {
		if p.Zona == strings.ToUpper(zona) {
			return p
		} else {
			return com.Asiento{}
		}
	})
	asientosfiltrados := filterAsientos(asientos, pCantidad)
	//fmt.Println("asientos mapeados:\n", asientos)
	//fmt.Println("asientos filtrados:\n", asientosfiltrados)
	for fila, asiento := range asientosfiltrados {
		cadena += "\nCategoria: " + strings.ToUpper(pCategoria) + ", Zona: " + strings.ToUpper(zona) + ", Fila: " + strconv.Itoa(fila) + ", Asiento: " + strconv.Itoa(int(asiento.Columna))
	}
	fmt.Println("Mejores Opciones enviando...\n", cadena)
	fmt.Println(asientos)
	cadena += "\n"
	return cadena
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
func cargarDatos() {
	clientesData()
	facturasData()
	categoriaData()
}

func mostrarClientes() string {
	fmt.Println("Solicitando al servidor clientes...")
	var answer string = ""
	answer += "----Clientes----\n"
	for _, client := range dClientes {
		answer += client.NombreCompleto
		answer += "\n"
	}
	return answer
}

func mostrarFacturas() string {
	fmt.Println("Solicitando al servidor facturas...")
	var cadena = ""
	/*fmt.Println("----Facturas----")
	fmt.Println(dFacturas)*/
	return cadena
}

func monstrarCategorias() string {
	var answer = ""
	fmt.Println("Solicitando al servidor categorias...")
	answer += "----Categorias----\n"
	for category := range dCategorias {
		answer += category
		answer += "\n"
	}
	return answer
}

func mostrarZonas() string {
	var answer = ""
	fmt.Println("Solicitando al servidor zonas...")
	answer += "----Zonas----\n"
	answer += "1. Zona A.\n"
	answer += "2. Zona B.\n"
	answer += "3. Zona C.\n"
	return answer
}

func mostrarAsientos(pCategoria string, pZona string) string {
	var cadena = ""
	var zona = ""
	switch pZona {
	case "1":
		zona = "a"
	case "2":
		zona = "b"
	case "3":
		zona = "c"
	}
	var matrixAsientos = obtenerAsientos(pCategoria, zona)
	var asientos = mapAsientos(matrixAsientos, func(p com.Asiento) com.Asiento {
		if p.Zona == strings.ToUpper(zona) {
			return p
		} else {
			return com.Asiento{}
		}
	})
	asientosfiltrados := filterAsientos(asientos, -1)
	for fila, asiento := range asientosfiltrados {
		cadena += "\nCategoria: " + strings.ToUpper(pCategoria) + ", Zona: " + strings.ToUpper(zona) + ", Fila: " + strconv.Itoa(fila) + ", Asiento: " + strconv.Itoa(int(asiento.Columna))
	}
	cadena += "\n"
	fmt.Println(asientos)
	return cadena
}

func mostrarInfoPrincipal() string {
	var cadena = ""
	cadena += "\n-----Eventos Luna-----\n"
	cadena += "1. Comprar Asientos.\n"
	cadena += "2. Buscar Factura.\n"
	cadena += "3. Buscar Cliente.\n"
	cadena += "4. Registrarse\n"
	cadena += "5. Salir.\n"
	cadena += "6. Clientes.\n"
	cadena += "7. Facturas.\n"
	//fmt.Println(disponibilidad("VIP", "A", 0, 3))
	//mejoresOpciones("vip", "a", 4)
	//cambiarDisponibilidad("vip", "a", 1, 9, -1)
	//cambiarDisponibilidad("vip", "a", 1, 7, -1)
	//mejoresOpciones("vip", "a", 4)
	return cadena
}

func logInTask(pConnection net.Conn) string {
	buffer := make([]byte, 1024)
	var cadena string = ""

	pConnection.Write([]byte("Cedula:"))
	read, err := pConnection.Read(buffer)
	var cedula string = string(buffer[:read])
	fmt.Println("cedula:", cedula)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	pConnection.Write([]byte("Nombre:"))
	read2, err := pConnection.Read(buffer)
	var nombre string = string(buffer[:read2])
	fmt.Println("name:", nombre)

	pConnection.Write([]byte("Email:"))
	read3, err := pConnection.Read(buffer)
	var email string = string(buffer[:read3])
	fmt.Println("email:", email)

	agregarCliente(cedula, nombre, email)

	cadena += "Succesfull registration!\n"
	return cadena
}

func buyTask(pConnection net.Conn) {
	buffer := make([]byte, 1024)

	pConnection.Write([]byte(monstrarCategorias() + "\nQue categoria desea comprar?\n"))
	read, _ := pConnection.Read(buffer)
	var category = string(buffer[:read])

	pConnection.Write([]byte(mostrarZonas() + "\nQue zona desea comprar?\n"))
	read1, _ := pConnection.Read(buffer)
	var zona = string(buffer[:read1])

	pConnection.Write([]byte("\nCuantos tickets desea comprar?\n"))
	read2, _ := pConnection.Read(buffer)
	var cantidad = string(buffer[:read2])
	cant, _ := strconv.Atoi(cantidad)

	if strings.ToUpper(category) != "VIP" && strings.ToUpper(category) != "PALCO" &&
		strings.ToUpper(category) != "SOMBRA" && strings.ToUpper(category) != "GRAMILLA" {
		return
	}

	switch strings.ToUpper(category) {
	case "VIP":
		switch zona {
		case "1":
			comprarAsiento(pConnection, category, zona, cant) // categoria string, zona string, fila int, columna/asiento int
		case "2":
			comprarAsiento(pConnection, category, zona, cant)
		case "3":
			comprarAsiento(pConnection, category, zona, cant)
		}
	case "GRAMILLA":
		switch zona {
		case "1":
			comprarAsiento(pConnection, category, zona, cant)
		case "2":
			comprarAsiento(pConnection, category, zona, cant)
		case "3":
			comprarAsiento(pConnection, category, zona, cant)
		}
	case "PALCO":
		switch zona {
		case "1":
			comprarAsiento(pConnection, category, zona, cant)
		case "2":
			comprarAsiento(pConnection, category, zona, cant)
		case "3":
			comprarAsiento(pConnection, category, zona, cant)
		}
	case "SOMBRA":
		switch zona {
		case "1":
			comprarAsiento(pConnection, category, zona, cant)
		case "2":
			comprarAsiento(pConnection, category, zona, cant)
		case "3":
			comprarAsiento(pConnection, category, zona, cant)
		}
	}
}

func taskFlow(pConnection net.Conn, pTask string) string {
	var cadena = "\n-----Answer-----\n"

	switch pTask {
	case "1": // compra de asientos
		buyTask(pConnection)
	case "2": // buscar factura
	case "3": // buscar cliente
	case "4": // registrarse
		answer := logInTask(pConnection)
		cadena += answer
	case "6":
		pConnection.Write([]byte(mostrarClientes()))
	}

	return cadena
}

func processClient(connection net.Conn) {
	buffer := make([]byte, 1024)
	for {
		connection.Write([]byte(mostrarInfoPrincipal()))
		mLen, err := connection.Read(buffer)
		if err != nil {
			fmt.Println("Error reading:", err.Error())
			return
		}
		fmt.Println("Received: ", string(buffer[:mLen]))
		if string(buffer[:mLen]) != "" {
			taskFlow(connection, string(buffer[:mLen]))
			//_, err = connection.Write([]byte(answer))
		}
	}
}

func runServer() {
	fmt.Println("Server Running...")
	server, err := net.Listen(SERVER_TYPE, SERVER_HOST+":"+SERVER_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()
	fmt.Println("Listening on " + SERVER_HOST + ":" + SERVER_PORT)
	fmt.Println("Waiting for client...")
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		fmt.Println("Client connected")
		go processClient(connection)
	}
}

func main() {
	cargarDatos()
	//mejoresOpciones("vip", "3", 3)
	//mostrarAsientos("vip", "3")
	runServer()
}

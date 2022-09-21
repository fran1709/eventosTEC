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

func datosCliente(pCedula string) *com.Cliente {
	var cliente = dClientes[pCedula]
	return &cliente
}

/*
Agrega un nuevo elemento "cliente" al diccionario de Clientes->dClientes
*/
func agregarCliente(pCedula string, pNombreCompleto string, pCorreo string) {
	cliente := isCliente(pCedula)
	if !cliente {
		dClientes[pCedula] = com.Cliente{Cedula: pCedula, NombreCompleto: pNombreCompleto, Correo: pCorreo}
		fmt.Print("Cliente agregado \n")
	} else {
		fmt.Print("Cliente no agregado, porque ya existe \n")
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
func agregarFactura(pId int32, pCliente *com.Cliente, pAsiento *com.Asiento, pPrecio int32) {
	factura := isFactura(pId)
	if !factura {
		dFacturas[pId] = com.Factura{IdFactura: pId, Cliente: pCliente, Asiento: pAsiento, Precio: pPrecio}
		fmt.Print("Factura agregada -> ", dFacturas[pId], "\n")
	} else {
		fmt.Print("Factura existente no agregado. \n")
	}
}

/*
Si el cliente decide comprar las mejores opciones, se ejecuta esta funcion que les cambia el estado a comprado.
Además crea y agrega la factura en el sistema con los datos del cliente(debe estar registrado para comprar)
*/
func comprarSugerencias(pListaAsientos ListaAsientos, pCedula string) {
	for _, asiento := range pListaAsientos {
		cambiarDisponibilidad(asiento.Categoria, asiento.Zona, int(asiento.Fila), int(asiento.Columna), -1)
		//facturar a nombre del cliente
		agregarFactura(idFactura, datosCliente(pCedula), &asiento, int32(asiento.Costo))
		idFactura++
	}
}

/*
*
Funcion encargada de la compra de asientos
*/
func comprarAsiento(pConnection net.Conn, pCategory string, pZona string, pCantidad int, pBuffer []byte) {
	fmt.Println("Comprando tickets...")

	asientos, pLista := mejoresOpciones(pCategory, pZona, pCantidad)

	pConnection.Write([]byte("\nIngrese su cedula cliente:"))
	read2, _ := pConnection.Read(pBuffer)
	var cedula = string(pBuffer[:read2])
	fmt.Println("Recibe cedula: " + cedula)

	pConnection.Write([]byte("\n1.Sugerencias.\n2.Listado Disponible.\n"))
	read, _ := pConnection.Read(pBuffer)
	var choice = string(pBuffer[:read])
	fmt.Println("Recibe decision: " + choice)

	switch choice {
	case "1":
		pConnection.Write([]byte(asientos + "\n1. Comprar Sugerencias.\n2. Salir."))
		read1, _ := pConnection.Read(pBuffer)
		var choice1 = string(pBuffer[:read1])
		fmt.Println("Recibe: " + choice1 + "\n")
		fmt.Println("Recibe decision:" + choice + "\n")

		switch choice1 {
		case "1":
			if isCliente(cedula) {
				comprarSugerencias(pLista, cedula)
			} else {
				pConnection.Write([]byte("\nNo existe un cliente asociado a la cedula -> " + cedula))
			}
		case "2":
			return
		}
	case "2":
		for i := 0; i < pCantidad; i++ {
			pConnection.Write([]byte(mostrarAsientos(pCategory, pZona) + "Que fila desea comprar?"))
			read3, _ := pConnection.Read(pBuffer)
			var fila = string(pBuffer[:read3])
			filaNum, _ := strconv.Atoi(fila)

			pConnection.Write([]byte("Que asiento desea comprar?"))
			read4, _ := pConnection.Read(pBuffer)
			var colum = string(pBuffer[:read4])
			columNum, _ := strconv.Atoi(colum)

			agregarFactura(idFactura, datosCliente(cedula), datosAsiento(pCategory, pZona, filaNum, columNum), int32(datosAsiento(pCategory, pZona, filaNum, columNum).Costo))
			cambiarDisponibilidad(pCategory, pZona, filaNum, columNum, -1)
			idFactura++
		}
	}
}

/*
*
--------------------------------------------------------------------------------

	ASIENTOS FUNCTIONS

--------------------------------------------------------------------------------
*/

func filtrarAsiento(pFila int, pColum int, pLista *[]com.Asiento) *com.Asiento {
	var asiento *com.Asiento
	for _, asient := range *pLista {
		if asient.Fila == int16(pFila) && asient.Columna == int16(pColum) {
			asiento = &asient
			break
		}
	}
	return asiento
}

func datosAsiento(pCategoria string, pZona string, pFila int, pColum int) *com.Asiento {
	matrixAsientos := obtenerAsientos(pCategoria, pZona)
	var asientos = mapAsientos(matrixAsientos, func(p com.Asiento) com.Asiento {
		if p.Zona == strings.ToUpper(pZona) {
			return p
		} else {
			return com.Asiento{}
		}
	})
	var asiento = filtrarAsiento(pFila, pColum, asientos)
	return asiento
}

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

func mejoresOpciones(pCategoria string, pZona string, pCantidad int) (string, ListaAsientos) {
	var cadena = ""
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
	//fmt.Println("asientos filtrados:\n", asientosfiltrados)
	for _, asiento := range asientosfiltrados {
		cadena += "-" + strings.ToUpper(pCategoria) + "-" + strings.ToUpper(pZona) + "-" + strconv.Itoa(int(asiento.Fila)) + "-" + strconv.Itoa(int(asiento.Columna)) + "-"
	}
	fmt.Println("Mejores Opciones enviando...\n", cadena)
	//fmt.Println(asientos)
	return cadena, asientosfiltrados
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
func crearAsientos(pCat string, pZona string, pFilas int, pAsientos int, pCosto int) *[][]com.Asiento {

	var asientos = make([][]com.Asiento, pFilas)
	//fmt.Println("salidos del horno\n", asientos)
	for i := 0; i < pFilas; i++ {
		asientos[i] = make([]com.Asiento, pAsientos)
	}
	//fmt.Println(asientos)
	for i := 0; i < pFilas; i++ {
		for j := 0; j < pAsientos; j++ {
			asientos[i][j] = com.Asiento{Categoria: pCat, Zona: pZona, Fila: int16(i), Columna: int16(j), Estado: 1, Costo: pCosto}
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
	agregarCategoria("VIP", crearAsientos("VIP", "A", 2, 5, 100), crearAsientos("VIP", "B", 4, 4, 75), crearAsientos("VIP", "C", 5, 3, 50))
	agregarCategoria("GRAMILLA", crearAsientos("GRAMILLA", "A", 2, 5, 40), crearAsientos("GRAMILLA", "B", 4, 4, 35), crearAsientos("GRAMILLA", "C", 5, 3, 30))
	agregarCategoria("PALCO", crearAsientos("PALCO", "A", 2, 5, 25), crearAsientos("PALCO", "B", 4, 20, 4), crearAsientos("PALCO", "C", 5, 3, 15))
	agregarCategoria("SOMBRA", crearAsientos("SOMBRA", "A", 2, 5, 10), crearAsientos("SOMBRA", "B", 4, 4, 5), crearAsientos("SOMBRA", "C", 5, 3, 2))

}

/*
*
--------------------------------------------------------------------------------

	MAIN FUNCTIONS

--------------------------------------------------------------------------------
*/
func cargarDatos() {
	clientesData()
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
	for factura := range dFacturas {
		cadena += "|--------------- FACTURA ---------------|\n"
		cadena += "Id -> " + strconv.Itoa(int(dFacturas[factura].IdFactura))
		cadena += "         ---ASIENTO---         \n"
		cadena += "Categoria -> " + dFacturas[factura].Asiento.Categoria + "\n"
		cadena += "Zona -> " + dFacturas[factura].Asiento.Zona + "\n"
		cadena += "Fila -> " + strconv.Itoa(int(dFacturas[factura].Asiento.Fila)) + "\n"
		cadena += "Asiento -> " + strconv.Itoa(int(dFacturas[factura].Asiento.Columna)) + "\n"
		cadena += "         ---CLIENTE---         \n"
		cadena += "Nombre -> " + dFacturas[factura].Cliente.NombreCompleto + "\n"
		cadena += "Cedula -> " + dFacturas[factura].Cliente.Cedula + "\n"
		cadena += "Email -> " + dFacturas[factura].Cliente.Correo + "\n"
		cadena += "|----------------- END -----------------|\n"
		cadena += "\n"
	}
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
	var matrixAsientos = obtenerAsientos(pCategoria, pZona)
	var asientos = mapAsientos(matrixAsientos, func(p com.Asiento) com.Asiento {
		if p.Zona == strings.ToUpper(pZona) {
			return p
		} else {
			return com.Asiento{}
		}
	})
	asientosfiltrados := filterAsientos(asientos, -1)
	for _, asiento := range asientosfiltrados {
		cadena += "\nCategoria: " + strings.ToUpper(pCategoria) + ", Zona: " + strings.ToUpper(pZona) + ", Fila: " + strconv.Itoa(int(asiento.Fila)) + ", Asiento: " + strconv.Itoa(int(asiento.Columna))
	}
	cadena += "\n"
	fmt.Println("Enviando listado disponible total...")
	//fmt.Println(asientos)
	return cadena
}

func mostrarInfoPrincipal() string {
	var cadena = ""
	cadena += "\n-----Eventos Luna-----\n"
	cadena += "1. Comprar Asientos.\n"
	cadena += "2. Buscar Cliente.\n"
	cadena += "3. Registrarse.\n"
	cadena += "4. Clientes.\n"
	cadena += "5. Facturas.\n"
	cadena += "6. Asientos.\n"
	cadena += "7. Salir.\n"
	return cadena
}

/*
*
Recibe datos respectivos para la creación de un usuario desde el cliente haskell.
*/
func logInTask(pConnection net.Conn, pBuffer []byte) {

	pConnection.Write([]byte("Cedula:"))
	read, _ := pConnection.Read(pBuffer)
	var cedula string = string(pBuffer[:read])
	fmt.Println("cedula:", cedula)

	pConnection.Write([]byte("Nombre:"))
	read2, _ := pConnection.Read(pBuffer)
	var nombre string = string(pBuffer[:read2])
	fmt.Println("name:", nombre)

	pConnection.Write([]byte("Email:"))
	read3, _ := pConnection.Read(pBuffer)
	var email string = string(pBuffer[:read3])
	fmt.Println("email:", email)

	agregarCliente(cedula, nombre, email)
}

func buyTask(pConnection net.Conn, pBuffer []byte) {
	pConnection.Write([]byte(monstrarCategorias() + "\nQue categoria desea comprar?\n"))
	read, _ := pConnection.Read(pBuffer)
	var category = string(pBuffer[:read])
	fmt.Println("Recibe categoria: " + category)

	pConnection.Write([]byte(mostrarZonas() + "\nQue zona desea comprar?\n"))
	read1, _ := pConnection.Read(pBuffer)
	var zona = string(pBuffer[:read1])
	var pZona = ""
	switch zona {
	case "1":
		pZona = "a"
	case "2":
		pZona = "b"
	case "3":
		pZona = "c"
	}
	fmt.Println("Recibe zona: " + pZona)

	pConnection.Write([]byte("\nCuantos tickets desea comprar?\n"))
	read2, _ := pConnection.Read(pBuffer)
	var cantidad = string(pBuffer[:read2])
	cant, _ := strconv.Atoi(cantidad)
	fmt.Println("Recibe cantidad: " + cantidad)

	if strings.ToUpper(category) != "VIP" && strings.ToUpper(category) != "PALCO" &&
		strings.ToUpper(category) != "SOMBRA" && strings.ToUpper(category) != "GRAMILLA" {
		return
	}
	if pZona != "a" && pZona != "b" && pZona != "c" {
		return
	}

	switch strings.ToUpper(category) {
	case "VIP":
		switch pZona {
		case "a":
			comprarAsiento(pConnection, category, pZona, cant, pBuffer) // categoria string, zona string, fila int, columna/asiento int
		case "b":
			comprarAsiento(pConnection, category, pZona, cant, pBuffer)
		case "c":
			comprarAsiento(pConnection, category, pZona, cant, pBuffer)
		}
	case "GRAMILLA":
		switch pZona {
		case "a":
			comprarAsiento(pConnection, category, pZona, cant, pBuffer)
		case "b":
			comprarAsiento(pConnection, category, pZona, cant, pBuffer)
		case "c":
			comprarAsiento(pConnection, category, pZona, cant, pBuffer)
		}
	case "PALCO":
		switch pZona {
		case "a":
			comprarAsiento(pConnection, category, pZona, cant, pBuffer)
		case "b":
			comprarAsiento(pConnection, category, pZona, cant, pBuffer)
		case "c":
			comprarAsiento(pConnection, category, pZona, cant, pBuffer)
		}
	case "SOMBRA":
		switch pZona {
		case "a":
			comprarAsiento(pConnection, category, pZona, cant, pBuffer)
		case "b":
			comprarAsiento(pConnection, category, pZona, cant, pBuffer)
		case "c":
			comprarAsiento(pConnection, category, pZona, cant, pBuffer)
		}
	}

}

/*
*
Funcion que maneja el flujo de trabajo(acciones) disponibles en el servidor segun las decisiones del cliente.
*/
func taskFlow(pConnection net.Conn, pTask string, pBuffer []byte) {
	switch pTask {
	case "1": // compra de asientos
		buyTask(pConnection, pBuffer)
	case "2": // buscar cliente (byCedula)
	case "3": // registrarse (cliente)
		logInTask(pConnection, pBuffer)
	case "4": // clientes en sistema
		pConnection.Write([]byte(mostrarClientes()))
	case "5": // facturas
		pConnection.Write([]byte(mostrarFacturas()))
	case "6":
		//pConnection.Write([]byte(showAll()))
	case "7": //salir
		fmt.Println("Client Dessconnected!")
	}
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
			taskFlow(connection, string(buffer[:mLen]), buffer)
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
	//fmt.Println(disponibilidad("VIP", "A", 0, 3))
	//mejoresOpciones("vip", "a", 4)
	//cambiarDisponibilidad("vip", "a", 1, 9, -1)
	//cambiarDisponibilidad("vip", "a", 1, 7, -1)
	//mejoresOpciones("vip", "a", 4)
	cargarDatos()
	//mejoresOpciones("vip", "3", 3)
	//mostrarAsientos("vip", "3")
	//_, lista := mejoresOpciones("vip", "3", 2)
	//comprarSugerencias(lista)
	//mejoresOpciones("vip", "3", 2)
	//mejoresOpciones("vip", "3", 2)
	runServer()
}

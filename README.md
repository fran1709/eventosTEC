# eventosTEC

-------------------------------------- MENÚ PRINCIPAL --------------------------------------
* Al hacer la conexión se presenta el menú mostrado anteriormente, y se espera una respuesta por parte del cliente.
Está programado para que reciba por respuesta los valores numéricos asociados llámesen 1,2,3,4,5,6 respectivamente.

-----  Eventos Luna  -----
1. Comprar Asientos.  
2. Buscar Cliente.    
3. Registrarse.       
4. Clientes.
5. Facturas.
6. Salir.

Mensaje respuesta al servidor: ....

-------------------------------------- Comprar Asientos Sugeridos --------------------------------------
* Habiendo seleccionado 1(Comprar Asientos), se muestra el siguiente menú, primero se muestran las categorias disponibles, se espera se escriba la categoria correspondiente(string) luego se muestra un menú de zonas en el cuál se espera recibir valores numéricos asociados a la consulta deseada por el cliente, luego se solicitan la cantidad de tiquets(asientos) a comprar, posteriormente se solicita la cedula cliente (solo se podrá comprar si el cliente(usuario) está registrado y porporciona una cédula válida). 

Mensaje respuesta al servidor:
1
----Categorias----
GRAMILLA
PALCO
SOMBRA
VIP

Que categoria desea comprar?

Mensaje respuesta al servidor: ...


Mensaje respuesta al servidor:
vip
----Zonas----
1. Zona A.
2. Zona B.
3. Zona C.

Que zona desea comprar?

Mensaje respuesta al servidor: 
1

Cuantos tickets desea comprar?

Mensaje respuesta al servidor:
3

Ingrese su cedula cliente:
Mensaje respuesta al servidor:
207710202

1.Sugerencias.
2.Listado Disponible.

Mensaje respuesta al servidor:
1
-VIP-A-1-4--VIP-A-1-3--VIP-A-1-2-
1. Comprar Sugerencias.
2. Salir.
Mensaje respuesta al servidor:

* --- VISTA EN EL SERVIDOR DE LOS ASIENTOS MODIFICADOS A COMPRADOS ---
Estado de Asientos {DISPONIBLE = 1, RESERVADO = 0, COMPRADO = -1}
com.Asiento = {Categoria Zona Fila Asiento Estado}

{VIP A 1 4 1 100}
Factura agregada -> {0 0xc0000a4060 0xc0000a4090 100}
&[[{VIP A 0 0 1 100} {VIP A 0 1 1 100} {VIP A 0 2 1 100} {VIP A 0 3 1 100} {VIP A 0 4 1 100}] [{VIP A 1 0 1 100} {VIP A 1 1 1 100} {VIP A 1 2 1 100} {VIP A 1 3 1 100} {VIP A 1 4 -1 100}]]

{VIP A 1 3 1 100}
Factura agregada -> {1 0xc0000a42d0 0xc0000a4300 100}
&[[{VIP A 0 0 1 100} {VIP A 0 1 1 100} {VIP A 0 2 1 100} {VIP A 0 3 1 100} {VIP A 0 4 1 100}] [{VIP A 1 0 1 100} {VIP A 1 1 1 100} {VIP A 1 2 1 100} {VIP A 1 3 -1 100} {VIP A 1 4 -1 100}]]

* Vista en FACTURAS opcion 5 del menú principal

|--------------- FACTURA ---------------|
 Id -> 0                ---ASIENTO---    
 Categoria -> VIP
 Zona -> A
 Fila -> 1
 Asiento -> 4
 Precio -> 100
                ---CLIENTE---
 Nombre -> Francisco Ovares Rojas        
 Cedula -> 207710202
 Email -> fran1709@estudiantec.cr        
|----------------- END -----------------|

|--------------- FACTURA ---------------|
 Id -> 1                ---ASIENTO---
 Categoria -> VIP
 Zona -> A
 Fila -> 1
 Asiento -> 3
 Precio -> 100
                ---CLIENTE---
 Nombre -> Francisco Ovares Rojas
 Cedula -> 207710202
 Email -> fran1709@estudiantec.cr
|----------------- END -----------------|


-------------------------------------- Comprar Asientos Específicos --------------------------------------
* SOMBRA - ZONA A - 3 Asientos seleccionados manualmente

Ingrese su cedula cliente:
Mensaje respuesta al servidor:
207710205

1.Sugerencias.
2.Listado Disponible.

Mensaje respuesta al servidor:
2

Categoria: SOMBRA, Zona: A, Fila: 1, Asiento: 4
Categoria: SOMBRA, Zona: A, Fila: 1, Asiento: 3
Categoria: SOMBRA, Zona: A, Fila: 1, Asiento: 2
Categoria: SOMBRA, Zona: A, Fila: 1, Asiento: 1
Categoria: SOMBRA, Zona: A, Fila: 1, Asiento: 0
Categoria: SOMBRA, Zona: A, Fila: 0, Asiento: 4
Categoria: SOMBRA, Zona: A, Fila: 0, Asiento: 3
Categoria: SOMBRA, Zona: A, Fila: 0, Asiento: 2
Categoria: SOMBRA, Zona: A, Fila: 0, Asiento: 1
Categoria: SOMBRA, Zona: A, Fila: 0, Asiento: 0
Que fila desea comprar?
Mensaje respuesta al servidor:
1
Que asiento desea comprar?
Mensaje respuesta al servidor:
2

Categoria: SOMBRA, Zona: A, Fila: 1, Asiento: 4
Categoria: SOMBRA, Zona: A, Fila: 1, Asiento: 3
Categoria: SOMBRA, Zona: A, Fila: 1, Asiento: 1
Categoria: SOMBRA, Zona: A, Fila: 1, Asiento: 0
Categoria: SOMBRA, Zona: A, Fila: 0, Asiento: 4
Categoria: SOMBRA, Zona: A, Fila: 0, Asiento: 3
Categoria: SOMBRA, Zona: A, Fila: 0, Asiento: 2
Categoria: SOMBRA, Zona: A, Fila: 0, Asiento: 1
Categoria: SOMBRA, Zona: A, Fila: 0, Asiento: 0
Que fila desea comprar?
Mensaje respuesta al servidor:
1
Que asiento desea comprar?
Mensaje respuesta al servidor:
3

Categoria: SOMBRA, Zona: A, Fila: 1, Asiento: 4
Categoria: SOMBRA, Zona: A, Fila: 1, Asiento: 1
Categoria: SOMBRA, Zona: A, Fila: 1, Asiento: 0
Categoria: SOMBRA, Zona: A, Fila: 0, Asiento: 4
Categoria: SOMBRA, Zona: A, Fila: 0, Asiento: 3
Categoria: SOMBRA, Zona: A, Fila: 0, Asiento: 2
Categoria: SOMBRA, Zona: A, Fila: 0, Asiento: 1
Categoria: SOMBRA, Zona: A, Fila: 0, Asiento: 0
Que fila desea comprar?
Mensaje respuesta al servidor:
1
Que asiento desea comprar?
Mensaje respuesta al servidor:
1

* CONSULTANDO EN FACTURAS LAS COMPRAS 

|--------------- FACTURA ---------------|
 Id -> 3                ---ASIENTO---
 Categoria -> SOMBRA
 Zona -> A
 Fila -> 1
 Asiento -> 3
 Precio -> 10
                ---CLIENTE---
 Nombre -> Thomas Ovares Molina
 Cedula -> 207710205
 Email -> thom1105@estudiantec.cr
|----------------- END -----------------|

|--------------- FACTURA ---------------|
 Id -> 4                ---ASIENTO---
 Categoria -> SOMBRA
 Zona -> A
 Fila -> 1
 Asiento -> 1
 Precio -> 10
                ---CLIENTE---
 Nombre -> Thomas Ovares Molina
 Cedula -> 207710205
 Email -> thom1105@estudiantec.cr
|----------------- END -----------------|

|--------------- FACTURA ---------------|
 Id -> 0                ---ASIENTO---
 Categoria -> VIP
 Zona -> A
 Fila -> 1
 Asiento -> 4
 Precio -> 100
                ---CLIENTE---
 Nombre -> Francisco Ovares Rojas
 Cedula -> 207710202
 Email -> fran1709@estudiantec.cr
|----------------- END -----------------|

|--------------- FACTURA ---------------|
 Id -> 1                ---ASIENTO---
 Categoria -> VIP
 Zona -> A
 Fila -> 1
 Asiento -> 3
 Precio -> 100
                ---CLIENTE---
 Nombre -> Francisco Ovares Rojas
 Cedula -> 207710202
 Email -> fran1709@estudiantec.cr
|----------------- END -----------------|

|--------------- FACTURA ---------------|
 Id -> 2                ---ASIENTO---
 Categoria -> SOMBRA
 Zona -> A
 Fila -> 1
 Asiento -> 2
 Precio -> 10
                ---CLIENTE---
 Nombre -> Thomas Ovares Molina
 Cedula -> 207710205
 Email -> thom1105@estudiantec.cr
|----------------- END -----------------|


-------------------------------------- BUSCANDO CLIENTE POR CÉDULA --------------------------------------
* Ejemplicación del proceso con una cédula no existente y una que exista.

-----Eventos Luna-----        
1. Comprar Asientos.
2. Buscar Cliente.
3. Registrarse.
4. Clientes.
5. Facturas.
6. Salir.

Mensaje respuesta al servidor:
2

Ingrese cedula cliente a consultar

* CÉDULA INEXISTENTE
Mensaje respuesta al servidor:    
1234567
No existe un cliente asociado a la cedula -> 1234567

* CÉDULA EXISTENTE
Mensaje respuesta al servidor:
207710202
----- CLIENTE -----
Francisco Ovares Rojas
207710202
fran1709@estudiantec.cr

-------------------------------------- CREANDO CLIENTE --------------------------------------
-----Eventos Luna-----
1. Comprar Asientos.
2. Buscar Cliente.
3. Registrarse.
4. Clientes.
5. Facturas.
6. Salir.

Mensaje respuesta al servidor:
3
Cedula:
Mensaje respuesta al servidor:
1234567
Nombre:
Mensaje respuesta al servidor:
clienteNuevo
Email:
Mensaje respuesta al servidor:
12345@123123.com

-----Eventos Luna-----
1. Comprar Asientos.
2. Buscar Cliente.
3. Registrarse.
4. Clientes.
5. Facturas.
6. Salir.

Mensaje respuesta al servidor:
4
----Clientes----
Josue Ovares Rojas 207710203 josue1908@estudiantec.cr

Thomas Ovares Molina 207710205 thom1105@estudiantec.cr

clienteNuevo 1234567 12345@123123.com

Francisco Ovares Rojas 207710202 fran1709@estudiantec.cr

* Consultando cliente una vez creado.
Ingrese cedula cliente a consultar

Mensaje respuesta al servidor:
1234567
----- CLIENTE -----
clienteNuevo
1234567
12345@123123.com

-------------------------------------- SALIENDO DEL CLIENTE --------------------------------------
-----Eventos Luna-----
1. Comprar Asientos.
2. Buscar Cliente.
3. Registrarse.
4. Clientes.
5. Facturas.
6. Salir.

Mensaje respuesta al servidor:
6
Disconnected!S

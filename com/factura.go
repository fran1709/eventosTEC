package com

type Factura struct {
	IdFactura int32
	Cliente   *Cliente
	Asiento   *Asiento
	Precio    int32
}

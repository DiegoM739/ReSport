package models

// MetodoPago es una INTERFACE que define el comportamiento de cualquier método de pago.
// Esto permite POLIMORFISMO: distintos tipos de pago se procesan distinto,
// pero se usan a través de la misma interface.
type MetodoPago interface {
	Procesar(monto float64) (bool, string)
	Validar() bool
	NombreMetodo() string
}

// ===== Implementación 1: Pago con Tarjeta =====

type PagoTarjeta struct {
	NumeroTarjeta   string
	Titular         string
	FechaExpiracion string
	CVV             string
}

func (p PagoTarjeta) Validar() bool {
	return len(p.NumeroTarjeta) == 16 && len(p.CVV) == 3
}

func (p PagoTarjeta) Procesar(monto float64) (bool, string) {
	if !p.Validar() {
		return false, "Datos de tarjeta inválidos"
	}
	return true, "Pago con tarjeta procesado correctamente"
}

func (p PagoTarjeta) NombreMetodo() string {
	return "Tarjeta"
}

// ===== Implementación 2: Pago por Transferencia =====

type PagoTransferencia struct {
	Banco        string
	NumeroCuenta string
	Comprobante  string
}

func (p PagoTransferencia) Validar() bool {
	return p.Comprobante != "" && p.NumeroCuenta != ""
}

func (p PagoTransferencia) Procesar(monto float64) (bool, string) {
	if !p.Validar() {
		return false, "Datos de transferencia incompletos"
	}
	return true, "Transferencia registrada, pendiente de verificación"
}

func (p PagoTransferencia) NombreMetodo() string {
	return "Transferencia"
}
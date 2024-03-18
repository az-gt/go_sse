package utils

import "time"

// Función para determinar el tiempo de retraso según el plan del cliente
func GetDelayForPlan(plan string) time.Duration {
	switch plan {
	case "start":
		return 15 * time.Second
	case "premium":
		// No hay retraso para los clientes premium
		return 0 * time.Second
	default:
		// Usar un valor predeterminado de retraso si el plan no es reconocido
		return 30 * time.Second
	}
}

package medidas //Creamos el paquete llamado medidas

// Definimos tres constantes para realizar conversiones
const KmPorMilla = 1.60934
const MetrosPorPie = 0.3048
const CentimetrosPorPulgada = 2.54

// Creamos tres funciones para la conversion de unidades del sistema inglés al sistema decimal
func MillasEnKm(s float64) float64{
	return s * KmPorMilla
}

func PiesEnMetros(s float64) float64{
	return s * MetrosPorPie
}

func PulgadasEnCentimetros( s float64) float64{
	return s * CentimetrosPorPulgada
}
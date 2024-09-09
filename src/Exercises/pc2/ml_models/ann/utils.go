package ann

import "math"

// Función de activación Sigmoid
func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

// Derivada de la función de activación Sigmoid
func sigmoidDerivative(x float64) float64 {
	return x * (1.0 - x)
}
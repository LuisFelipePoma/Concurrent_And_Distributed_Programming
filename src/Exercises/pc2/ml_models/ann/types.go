package ann

import "sync"

// Estructura de la Red Neuronal
type MLPC struct {
	inputNodes   int
	hiddenNodes  int
	outputNodes  int
	learningRate float64

	weightsInputHidden  [][]float64
	weightsHiddenOutput [][]float64

	mu sync.Mutex // Mutex para proteger los pesos durante las actualizaciones concurrentes
}

// Estructura de la Red Neuronal
type MLP struct {
	inputNodes   int
	hiddenNodes  int
	outputNodes  int
	learningRate float64

	weightsInputHidden  [][]float64
	weightsHiddenOutput [][]float64
}

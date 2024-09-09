package ann

import (
	"math/rand"
)

// Inicializa una nueva red MLP
func NewMLP(inputNodes, hiddenNodes, outputNodes int, learningRate float64) *MLP {
	mlp := &MLP{
		inputNodes:   inputNodes,
		hiddenNodes:  hiddenNodes,
		outputNodes:  outputNodes,
		learningRate: learningRate,
	}

	// Inicialización de pesos con valores aleatorios
	mlp.weightsInputHidden = make([][]float64, inputNodes)
	for i := range mlp.weightsInputHidden {
		mlp.weightsInputHidden[i] = make([]float64, hiddenNodes)
		for j := range mlp.weightsInputHidden[i] {
			mlp.weightsInputHidden[i][j] = rand.Float64()
		}
	}

	mlp.weightsHiddenOutput = make([][]float64, hiddenNodes)
	for i := range mlp.weightsHiddenOutput {
		mlp.weightsHiddenOutput[i] = make([]float64, outputNodes)
		for j := range mlp.weightsHiddenOutput[i] {
			mlp.weightsHiddenOutput[i][j] = rand.Float64()
		}
	}

	return mlp
}

// Entrena la red MLP con un conjunto de entrada y salida
func (mlp *MLP) Train(inputData [][]float64, targetData []float64) {
	for i := range inputData {
		// Forward pass
		hiddenInputs := make([]float64, mlp.hiddenNodes)
		for j := 0; j < mlp.hiddenNodes; j++ {
			for k := 0; k < mlp.inputNodes; k++ {
				hiddenInputs[j] += inputData[i][k] * mlp.weightsInputHidden[k][j]
			}
			hiddenInputs[j] = sigmoid(hiddenInputs[j])
		}

		finalOutputs := make([]float64, mlp.outputNodes)
		for j := 0; j < mlp.outputNodes; j++ {
			for k := 0; k < mlp.hiddenNodes; k++ {
				finalOutputs[j] += hiddenInputs[k] * mlp.weightsHiddenOutput[k][j]
			}
			finalOutputs[j] = sigmoid(finalOutputs[j])
		}

		// Calcula el error
		outputErrors := make([]float64, mlp.outputNodes)
		for j := 0; j < mlp.outputNodes; j++ {
			outputErrors[j] = targetData[i] - finalOutputs[j]
		}

		// Backward pass
		hiddenErrors := make([]float64, mlp.hiddenNodes)
		for j := 0; j < mlp.hiddenNodes; j++ {
			for k := 0; k < mlp.outputNodes; k++ {
				hiddenErrors[j] += outputErrors[k] * mlp.weightsHiddenOutput[j][k]
			}
		}

		// Actualiza los pesos entre la capa oculta y de salida
		for j := 0; j < mlp.hiddenNodes; j++ {
			for k := 0; k < mlp.outputNodes; k++ {
				mlp.weightsHiddenOutput[j][k] += mlp.learningRate * outputErrors[k] * sigmoidDerivative(finalOutputs[k]) * hiddenInputs[j]
			}
		}

		// Actualiza los pesos entre la capa de entrada y la capa oculta
		for j := 0; j < mlp.inputNodes; j++ {
			for k := 0; k < mlp.hiddenNodes; k++ {
				mlp.weightsInputHidden[j][k] += mlp.learningRate * hiddenErrors[k] * sigmoidDerivative(hiddenInputs[k]) * inputData[i][j]
			}
		}
	}
}

// Predice la salida para un conjunto de entradas
func (mlp *MLP) Predict(inputData [][]float64) []float64 {
	finalOutputs := make([]float64, len(inputData))

	for i := range inputData {
		hiddenInputs := make([]float64, mlp.hiddenNodes)
		for j := 0; j < mlp.hiddenNodes; j++ {
			for k := 0; k < mlp.inputNodes; k++ {
				hiddenInputs[j] += inputData[i][k] * mlp.weightsInputHidden[k][j]
			}
			hiddenInputs[j] = sigmoid(hiddenInputs[j])
		}

		output := 0.0
		for j := 0; j < mlp.outputNodes; j++ {
			for k := 0; k < mlp.hiddenNodes; k++ {
				output += hiddenInputs[k] * mlp.weightsHiddenOutput[k][j]
			}
			output = sigmoid(output)
		}

		// Aplicar el umbral para convertir la probabilidad en 1 o 0
		if output >= 0.5 {
			finalOutputs[i] = 1
		} else {
			finalOutputs[i] = 0
		}
	}

	return finalOutputs
}

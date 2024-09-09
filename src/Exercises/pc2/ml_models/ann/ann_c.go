package ann

import (
	"math/rand"
	"sync"
)



// NewMLP inicializa una nueva red MLP
func NewMLPC(inputNodes, hiddenNodes, outputNodes int, learningRate float64) *MLP {
	mlp := &MLP{
		inputNodes:      inputNodes,
		hiddenNodes:     hiddenNodes,
		outputNodes:     outputNodes,
		learningRate:    learningRate,
		weightsInputHidden: make([][]float64, inputNodes),
		weightsHiddenOutput: make([][]float64, hiddenNodes),
	}

	// Inicialización de pesos con valores aleatorios
	for i := range mlp.weightsInputHidden {
		mlp.weightsInputHidden[i] = make([]float64, hiddenNodes)
		for j := range mlp.weightsInputHidden[i] {
			mlp.weightsInputHidden[i][j] = rand.Float64()
		}
	}

	for i := range mlp.weightsHiddenOutput {
		mlp.weightsHiddenOutput[i] = make([]float64, outputNodes)
		for j := range mlp.weightsHiddenOutput[i] {
			mlp.weightsHiddenOutput[i][j] = rand.Float64()
		}
	}

	return mlp
}

// Train entrena la red MLP con un conjunto de entrada y salida utilizando concurrencia
func (mlp *MLPC) Train(inputData [][]float64, targetData []float64) {
	var wg sync.WaitGroup

	for i := 0; i < len(inputData); i += 10 { // Procesamiento en lotes de 10 entradas
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			for j := i; j < i+10 && j < len(inputData); j++ {
				hiddenInputs := make([]float64, mlp.hiddenNodes)
				for k := 0; k < mlp.hiddenNodes; k++ {
					for l := 0; l < mlp.inputNodes; l++ {
						hiddenInputs[k] += inputData[j][l] * mlp.weightsInputHidden[l][k]
					}
					hiddenInputs[k] = sigmoid(hiddenInputs[k])
				}

				finalOutputs := make([]float64, mlp.outputNodes)
				for k := 0; k < mlp.outputNodes; k++ {
					for l := 0; l < mlp.hiddenNodes; l++ {
						finalOutputs[k] += hiddenInputs[l] * mlp.weightsHiddenOutput[l][k]
					}
					finalOutputs[k] = sigmoid(finalOutputs[k])
				}

				// Calcula el error
				outputErrors := make([]float64, mlp.outputNodes)
				for k := 0; k < mlp.outputNodes; k++ {
					outputErrors[k] = targetData[j] - finalOutputs[k]
				}

				// Backward pass
				hiddenErrors := make([]float64, mlp.hiddenNodes)
				for k := 0; k < mlp.hiddenNodes; k++ {
					for l := 0; l < mlp.outputNodes; l++ {
						hiddenErrors[k] += outputErrors[l] * mlp.weightsHiddenOutput[k][l]
					}
				}

				// Protege las actualizaciones de pesos con un Mutex
				mlp.mu.Lock()
				// Actualiza los pesos entre la capa oculta y de salida
				for k := 0; k < mlp.hiddenNodes; k++ {
					for l := 0; l < mlp.outputNodes; l++ {
						mlp.weightsHiddenOutput[k][l] += mlp.learningRate * outputErrors[l] * sigmoidDerivative(finalOutputs[l]) * hiddenInputs[k]
					}
				}

				// Actualiza los pesos entre la capa de entrada y la capa oculta
				for k := 0; k < mlp.inputNodes; k++ {
					for l := 0; l < mlp.hiddenNodes; l++ {
						mlp.weightsInputHidden[k][l] += mlp.learningRate * hiddenErrors[l] * sigmoidDerivative(hiddenInputs[l]) * inputData[j][k]
					}
				}
				mlp.mu.Unlock()
			}
		}(i)
	}

	wg.Wait()
}

// Predict predice la salida para un conjunto de entradas utilizando concurrencia
func (mlp *MLPC) Predict(inputData [][]float64) []float64 {
	finalOutputs := make([]float64, len(inputData))
	var wg sync.WaitGroup

	for i := 0; i < len(inputData); i += 10 { // Procesamiento en lotes de 10 entradas
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			for j := i; j < i+10 && j < len(inputData); j++ {
				hiddenInputs := make([]float64, mlp.hiddenNodes)
				for k := 0; k < mlp.hiddenNodes; k++ {
					for l := 0; l < mlp.inputNodes; l++ {
						hiddenInputs[k] += inputData[j][l] * mlp.weightsInputHidden[l][k]
					}
					hiddenInputs[k] = sigmoid(hiddenInputs[k])
				}

				output := 0.0
				for k := 0; k < mlp.outputNodes; k++ {
					for l := 0; l < mlp.hiddenNodes; l++ {
						output += hiddenInputs[l] * mlp.weightsHiddenOutput[l][k]
					}
					output = sigmoid(output)
				}

				finalOutputs[j] = output
			}
		}(i)
	}

	wg.Wait()
	return finalOutputs
}
package svm

import (
	"sync"
)

// SVM es una estructura que contiene los parámetros del SVM.
type SVMC struct {
	weights        []float64
	bias           float64
	learningRate   float64
	regularization float64
	epochs         int
	mu             sync.Mutex // Mutex para sincronización de acceso concurrente
}

// NewSVM crea un nuevo clasificador SVM con los parámetros iniciales.
func NewSVMC(learningRate, regularization float64, epochs int) *SVM {
	return &SVM{
		bias:           0.0,
		learningRate:   learningRate,
		regularization: regularization,
		epochs:         epochs,
	}
}


// Train entrena el clasificador SVM usando el descenso por gradiente en paralelo.
func (svm *SVMC) Train(X [][]float64, y []float64) {
	numSamples := len(X)
	numFeatures := len(X[0])
	svm.weights = make([]float64, numFeatures)

	var wg sync.WaitGroup
	for epoch := 0; epoch < svm.epochs; epoch++ {
		for i := 0; i < numSamples; i++ {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				x := X[i]
				label := y[i]
				prediction := svm.predict(x)
				if label*prediction < 1 {
					svm.mu.Lock()
					for j := 0; j < numFeatures; j++ {
						svm.weights[j] += svm.learningRate * (label*x[j] - svm.regularization*svm.weights[j])
					}
					svm.bias += svm.learningRate * label
					svm.mu.Unlock()
				}
			}(i)
		}
		wg.Wait()
	}
}

// predict realiza una predicción sobre un dato de entrada.
func (svm *SVMC) predict(x []float64) float64 {
	prediction := svm.bias
	for i, value := range x {
		prediction += svm.weights[i] * value
	}
	return sigmoid(prediction)
}

// Predict clasifica un conjunto de datos de entrada y devuelve las predicciones como un slice de float64.
func (svm *SVMC) Predict(X [][]float64) []float64 {
	predictions := make([]float64, len(X))
	var wg sync.WaitGroup
	for i, x := range X {
		wg.Add(1)
		go func(i int, x []float64) {
			defer wg.Done()
			predictions[i] = svm.predict(x)
		}(i, x)
	}
	wg.Wait()
	return predictions
}

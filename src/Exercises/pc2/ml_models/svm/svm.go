package svm

import (
	"math"
)

// SVM es una estructura que contiene los parámetros del SVM.
type SVM struct {
	weights        []float64
	bias           float64
	learningRate   float64
	regularization float64
	epochs         int
}

// NewSVMClassifier crea un nuevo clasificador SVM con los parámetros iniciales.
func NewSVM(learningRate, regularization float64, epochs int) *SVM {
	return &SVM{
		bias:           0.0,
		learningRate:   learningRate,
		regularization: regularization,
		epochs:         epochs,
	}
}

// Sigmoid es una función de activación que transforma el valor en una probabilidad.
func sigmoid(z float64) float64 {
	return 1 / (1 + math.Exp(-z))
}

// Train entrena el clasificador SVM usando el descenso por gradiente.
func (svm *SVM) Train(X [][]float64, y []float64) {
	numSamples := len(X)
	numFeatures := len(X[0])
	svm.weights  = make([]float64, numFeatures)

	for epoch := 0; epoch < svm.epochs; epoch++ {
		for i := 0; i < numSamples; i++ {
			x := X[i]
			label := y[i]
			prediction := svm.predict(x)
			if label*prediction < 1 {
				for j := 0; j < numFeatures; j++ {
					svm.weights[j] += svm.learningRate * (label*x[j] - svm.regularization*svm.weights[j])
				}
				svm.bias += svm.learningRate * label
			}
		}
	}
}

// predict realiza una predicción sobre un dato de entrada.
func (svm *SVM) predict(x []float64) float64 {
	prediction := svm.bias
	for i, value := range x {
		prediction += svm.weights[i] * value
	}
	return sigmoid(prediction)
}

// Predict clasifica un conjunto de datos de entrada y devuelve las predicciones como un slice de float64.
func (svm *SVM) Predict(X [][]float64) []float64 {
	predictions := make([]float64, len(X))
	for i, x := range X {
		predictions[i] = svm.predict(x)
	}
	return predictions
}
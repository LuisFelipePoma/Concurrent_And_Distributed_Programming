package metrics

import (
	"fmt"
	"pc2/ml_models/fc"
	"time"
)

// Interfaz que define los métodos que deben implementar los modelos
type MLModel interface {
	Train(X [][]float64, y []float64)
	Predict(X [][]float64) []float64
}

// Función para medir el tiempo de ejecución de un método
func measureExecutionTime(description string, f func()) time.Duration {
	start := time.Now()
	f()
	elapsed := time.Since(start)
	fmt.Printf("%s took %v\n", description, elapsed)
	return elapsed
}

// Función para calcular la precisión
func calculateAccuracy(predictions, labels []float64) float64 {
	if len(predictions) != len(labels) {
		return 0
	}
	correct := 0
	for i := range predictions {
		if predictions[i] == labels[i] {
			correct++
		}
	}
	return float64(correct) / float64(len(labels))
}

// Función para comparar los tiempos de ejecución y precisión entre dos modelos
func Comparar(modelo1, modelo2 MLModel, nombre1, nombre2 string, X [][]float64, y []float64) {
	// Medir el tiempo de entrenamiento del primer modelo
	trainTime1 := measureExecutionTime(fmt.Sprintf("%s Training", nombre1), func() {
		modelo1.Train(X, y)
	})

	// Medir el tiempo de predicción del primer modelo
	var predictions1 []float64
	predictTime1 := measureExecutionTime(fmt.Sprintf("%s Prediction", nombre1), func() {
		predictions1 = modelo1.Predict(X)
	})

	// Medir el tiempo de entrenamiento del segundo modelo
	trainTime2 := measureExecutionTime(fmt.Sprintf("%s Training", nombre2), func() {
		modelo2.Train(X, y)
	})

	// Medir el tiempo de predicción del segundo modelo
	var predictions2 []float64
	predictTime2 := measureExecutionTime(fmt.Sprintf("%s Prediction", nombre2), func() {
		predictions2 = modelo2.Predict(X)
	})

	// Calcular la precisión de ambos modelos
	accuracy1 := calculateAccuracy(predictions1, y)
	accuracy2 := calculateAccuracy(predictions2, y)

	// Comparar y mostrar resultados
	fmt.Println("\nComparison of Execution Times and Accuracy:")
	fmt.Printf("Training Time: %s vs %s: %v vs %v\n", nombre1, nombre2, trainTime1, trainTime2)
	fmt.Printf("Prediction Time: %s vs %s: %v vs %v\n", nombre1, nombre2, predictTime1, predictTime2)
	fmt.Printf("Accuracy: %s: %.2f%% vs %s: %.2f%%\n", nombre1, accuracy1*100, nombre2, accuracy2*100)
}

// CompararFC compara los tiempos de entrenamiento de fc y fc_c
func CompararFC(users []fc.User, targetUser int, k int) {
	// Medir el tiempo de entrenamiento de fc
	startFC := time.Now()
	recommendationsFC := fc.RecommendItems(users, targetUser, k)
	durationFC := time.Since(startFC)

	// Medir el tiempo de entrenamiento de fc_c
	startFCC := time.Now()
	recommendationsFCC := fc.RecommendItemsC(users, targetUser, k)
	durationFCC := time.Since(startFCC)

	// Imprimir los resultados
	fmt.Printf("Tiempo de entrenamiento de fc: %v\n", durationFC)
	fmt.Printf("Tiempo de entrenamiento de fc_c: %v\n", durationFCC)

	// Comparar las recomendaciones (opcional)
	compareRecommendations(recommendationsFC, recommendationsFCC)
}

// compareRecommendations compara las recomendaciones de fc y fc_c
func compareRecommendations(recommendationsFC, recommendationsFCC []int) {
	if len(recommendationsFC) != len(recommendationsFCC) {
		fmt.Println("Las longitudes de las recomendaciones no coinciden")
		return
	}

	var correct int
	for i := 0; i < len(recommendationsFC); i++ {
		if recommendationsFC[i] == recommendationsFCC[i] {
			correct++
		}
	}

	accuracy := float64(correct) / float64(len(recommendationsFC))
	fmt.Printf("Precisión de las recomendaciones: %f\n", accuracy)
}

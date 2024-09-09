package main

import (
	"fmt"
	"log"
	"pc2/metrics"
	"pc2/ml_models/ann"
	"pc2/ml_models/fc"
	// "pc2/ml_models/rf"
	"pc2/ml_models/svm"
	"pc2/panditas"
)

func main() {
	// ?test csv
	// Leer archivo
	fmt.Println("Reading CSV...")
	df, err := panditas.ReadCSV("dataset/sample_100.csv")
	if err != nil {
		log.Fatalf("Error reading CSV: %v", err)
	}

	// Obtener características y etiquetas
	features, labels, err := df.GetFeaturesAndLabels("target")
	if err != nil {
		log.Fatalf("Error getting features and labels: %v", err)
	}

	// !RANDOM FOREST
	// Crear instancias de los modelos secuencial y concurrente
	// rfSeq := rf.NewRandomForest(10, 2, 1, 1, 1)
	// rfConc := rf.NewRandomForestC(10, 2, 1, 1, 1)

	// !SVM
	svmSeq := svm.NewSVM(0.01, 0.1, 50)

	svmConc := svm.NewSVMC(0.01, 0.1, 50)

	// !ANN
	modelSeq := ann.NewMLP(25, 10, 2, 0.1)
	modelCon := ann.NewMLPC(25, 10, 2, 0.1)

	// !Filtrado Colaborativo
	ratings, err := fc.ReadRatingsFromCSV("dataset/rating.csv")
	if err != nil {
		log.Fatalf("Error reading ratings from CSV: %v", err)
	}

	// ?Comparar
	// metrics.Comparar(rfSeq, rfConc, "Sequential RF", "Concurrent RF", features, labels)
	metrics.Comparar(svmSeq, svmConc, "Sequential SVM", "Concurrent SVM", features, labels)
	metrics.Comparar(modelSeq, modelCon, "Sequential ANN", "Concurrent ANN", features, labels)
	metrics.CompararFC(ratings, 1, 5)
}
package rf

import (
	"fmt"
	"sync"
)

// RandomForest representa un bosque aleatorio.
type RandomForestC struct {
	Trees       []*TreeNode
	MaxDepth    int
	MinSize     int
	SampleSize  float64
	NumFeatures int
	NumTrees    int
	mu          sync.Mutex // Mutex para proteger el acceso concurrente a Trees
}

// NewRandomForest inicializa un modelo Random Forest con los parámetros de configuración.
func NewRandomForestC(numTrees int, maxDepth int, minSize int, sampleSize float64, numFeatures int) *RandomForestC {
	return &RandomForestC{
		Trees:       make([]*TreeNode, 0, numTrees),
		NumTrees:    numTrees,
		MaxDepth:    maxDepth,
		MinSize:     minSize,
		SampleSize:  sampleSize,
		NumFeatures: numFeatures,
	}
}

// Train entrena el modelo Random Forest.
func (rf *RandomForestC) Train(features [][]float64, labels []float64) {
	data := make([]DataPoint, len(features))
	for i := range features {
		data[i] = DataPoint{Features: features[i], Label: labels[i]}
	}

	fmt.Println("Numero de arboles: ", rf.NumTrees)

	var wg sync.WaitGroup
	for i := 0; i < rf.NumTrees; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tree := buildTree(data, rf.MaxDepth, rf.MinSize, rf.NumFeatures)
			if tree == nil {
				panic("Failed to build a tree")
			}
			rf.mu.Lock()
			rf.Trees = append(rf.Trees, tree)
			rf.mu.Unlock()
		}()
	}
	wg.Wait()
}

// Predict realiza la predicción utilizando un bosque aleatorio para múltiples conjuntos de características.
func (rf *RandomForestC) Predict(features [][]float64) []float64 {
	predictions := make([]float64, len(features))

	for i, featureSet := range features {
		if len(featureSet) == 0 {
			panic("Feature set is empty")
		}
		votes := make(map[float64]int)
		for _, tree := range rf.Trees {
			prediction := predict(tree, featureSet)
			votes[prediction]++
		}

		maxVotes := 0
		var finalLabel float64
		for label, count := range votes {
			if count > maxVotes {
				maxVotes = count
				finalLabel = label
			}
		}
		predictions[i] = finalLabel
	}

	return predictions
}
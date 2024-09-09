package rf

// DataPoint representa un punto de datos con características y etiqueta.
type DataPoint struct {
	Features []float64
	Label    float64
}

// TreeNode representa un nodo en el árbol de decisión.
type TreeNode struct {
	Left       *TreeNode
	Right      *TreeNode
	SplitValue float64
	Feature    int
	Label      float64
}

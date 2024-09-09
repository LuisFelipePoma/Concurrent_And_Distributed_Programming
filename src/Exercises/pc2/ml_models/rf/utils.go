package rf

import (
    "encoding/csv"
    "math"
    "math/rand"
    "os"
    "strconv"
)

// LoadCSV carga los datos desde un archivo CSV.
func LoadCSV(filename string) ([]DataPoint, error) {
    file, err := os.Open(filename)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    rawData, err := reader.ReadAll()
    if err != nil {
        return nil, err
    }

    var data []DataPoint
    for _, row := range rawData[1:] { // Suponiendo que la primera fila es el encabezado
        features := make([]float64, len(row)-1)
        for i, val := range row[:len(row)-1] {
            f, err := strconv.ParseFloat(val, 64)
            if err != nil {
                return nil, err
            }
            features[i] = f
        }
        label, err := strconv.ParseFloat(row[len(row)-1], 64)
        if err != nil {
            return nil, err
        }
        data = append(data, DataPoint{Features: features, Label: label})
    }
    return data, nil
}

// buildTree construye el árbol de decisión.
func buildTree(data []DataPoint, maxDepth int, minSize int, numFeatures int) *TreeNode {
    return splitNode(data, maxDepth, minSize, numFeatures, 1)
}

// splitNode divide un nodo del árbol.
func splitNode(data []DataPoint, maxDepth int, minSize int, numFeatures int, depth int) *TreeNode {
    labels := make(map[float64]int)
    for _, point := range data {
        labels[point.Label]++
    }

    if len(labels) == 1 || depth >= maxDepth || len(data) <= minSize {
        return &TreeNode{Label: majorityLabel(data)}
    }

    splitFeature, splitValue, left, right := bestSplit(data, numFeatures)
    if len(left) == 0 || len(right) == 0 {
        return &TreeNode{Label: majorityLabel(data)}
    }

    leftNode := splitNode(left, maxDepth, minSize, numFeatures, depth+1)
    rightNode := splitNode(right, maxDepth, minSize, numFeatures, depth+1)
    return &TreeNode{Left: leftNode, Right: rightNode, SplitValue: splitValue, Feature: splitFeature}
}

// bestSplit encuentra la mejor división para un nodo.
func bestSplit(data []DataPoint, numFeatures int) (int, float64, []DataPoint, []DataPoint) {
    var bestFeature int
    var bestValue float64
    var bestGini = math.MaxFloat64
    var bestLeft, bestRight []DataPoint

    features := rand.Perm(len(data[0].Features))[:numFeatures]

    for _, feature := range features {
        for _, point := range data {
            left, right := splitData(data, feature, point.Features[feature])
            gini := giniIndex(left, right)
            if gini < bestGini {
                bestGini = gini
                bestFeature = feature
                bestValue = point.Features[feature]
                bestLeft = left
                bestRight = right
            }
        }
    }
    return bestFeature, bestValue, bestLeft, bestRight
}

// splitData divide los datos según una característica.
func splitData(data []DataPoint, feature int, value float64) ([]DataPoint, []DataPoint) {
    left := make([]DataPoint, 0, len(data)/2)
    right := make([]DataPoint, 0, len(data)/2)
    for _, point := range data {
        if point.Features[feature] < value {
            left = append(left, point)
        } else {
            right = append(right, point)
        }
    }
    return left, right
}

// giniIndex calcula el índice Gini para una división.
func giniIndex(left, right []DataPoint) float64 {
    totalSize := len(left) + len(right)
    gini := 0.0

    for _, group := range [][]DataPoint{left, right} {
        size := len(group)
        if size == 0 {
            continue
        }
        score := 0.0
        labelCounts := make(map[float64]int)
        for _, point := range group {
            labelCounts[point.Label]++
        }
        for _, count := range labelCounts {
            proportion := float64(count) / float64(size)
            score += proportion * proportion
        }
        gini += (1.0 - score) * (float64(size) / float64(totalSize))
    }
    return gini
}

// majorityLabel encuentra la etiqueta mayoritaria.
func majorityLabel(data []DataPoint) float64 {
    labelCounts := make(map[float64]int)
    for _, point := range data {
        labelCounts[point.Label]++
    }
    var majorityLabel float64
    maxCount := 0
    for label, count := range labelCounts {
        if count > maxCount {
            maxCount = count
            majorityLabel = label
        }
    }
    return majorityLabel
}

// predict predice la etiqueta de un punto de datos.
func predict(node *TreeNode, features []float64) float64 {
    if node == nil {
        panic("Attempted to predict with a nil node")
    }
    if node.Left == nil && node.Right == nil {
        return node.Label
    }
    if node.Feature >= len(features) {
        panic("Feature index out of range")
    }
    if features[node.Feature] < node.SplitValue {
        return predict(node.Left, features)
    }
    return predict(node.Right, features)
}
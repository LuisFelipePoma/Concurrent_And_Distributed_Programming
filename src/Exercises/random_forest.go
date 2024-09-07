package exercises

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// TreeNode represents a node in the decision tree
type TreeNode struct {
	FeatureIndex int
	Threshold    float64
	Left         *TreeNode
	Right        *TreeNode
	Label        int
}

// RandomForest represents the random forest model
type RandomForest struct {
	Trees []*TreeNode
}

// splitDataSet splits the dataset based on a feature and threshold
func splitDataSet(data [][]float64, labels []int, featureIndex int, threshold float64) ([][]float64, []int, [][]float64, []int) {
	leftData, rightData := [][]float64{}, [][]float64{}
	leftLabels, rightLabels := []int{}, []int{}
	for i, row := range data {
		if row[featureIndex] <= threshold {
			leftData = append(leftData, row)
			leftLabels = append(leftLabels, labels[i])
		} else {
			rightData = append(rightData, row)
			rightLabels = append(rightLabels, labels[i])
		}
	}
	return leftData, leftLabels, rightData, rightLabels
}

// calculateGini calculates the Gini impurity for a dataset
func calculateGini(labels []int) float64 {
	labelCounts := make(map[int]int)
	for _, label := range labels {
		labelCounts[label]++
	}
	gini := 1.0
	for _, count := range labelCounts {
		prob := float64(count) / float64(len(labels))
		gini -= prob * prob
	}
	return gini
}

// buildTree builds a decision tree using the given data and labels
func buildTree(data [][]float64, labels []int, maxDepth, minSize int) *TreeNode {
	if len(data) == 0 || maxDepth == 0 || len(data) <= minSize {
		return &TreeNode{Label: majorityLabel(labels)}
	}

	bestFeature, bestThreshold, bestGini := -1, 0.0, math.MaxFloat64
	for featureIndex := range data[0] {
		thresholds := uniqueThresholds(data, featureIndex)
		for _, threshold := range thresholds {
			leftData, leftLabels, rightData, rightLabels := splitDataSet(data, labels, featureIndex, threshold)
			gini := (calculateGini(leftLabels)*float64(len(leftLabels)) + calculateGini(rightLabels)*float64(len(rightLabels))) / float64(len(labels))
			if gini < bestGini {
				bestFeature, bestThreshold, bestGini = featureIndex, threshold, gini
			}
		}
	}

	leftData, leftLabels, rightData, rightLabels := splitDataSet(data, labels, bestFeature, bestThreshold)
	leftNode := buildTree(leftData, leftLabels, maxDepth-1, minSize)
	rightNode := buildTree(rightData, rightLabels, maxDepth-1, minSize)
	return &TreeNode{FeatureIndex: bestFeature, Threshold: bestThreshold, Left: leftNode, Right: rightNode}
}

// majorityLabel returns the majority label in the dataset
func majorityLabel(labels []int) int {
	labelCounts := make(map[int]int)
	for _, label := range labels {
		labelCounts[label]++
	}
	majorityLabel, maxCount := -1, 0
	for label, count := range labelCounts {
		if count > maxCount {
			majorityLabel, maxCount = label, count
		}
	}
	return majorityLabel
}

// uniqueThresholds returns unique thresholds for a feature
func uniqueThresholds(data [][]float64, featureIndex int) []float64 {
	thresholds := make(map[float64]struct{})
	for _, row := range data {
		thresholds[row[featureIndex]] = struct{}{}
	}
	uniqueThresholds := make([]float64, 0, len(thresholds))
	for threshold := range thresholds {
		uniqueThresholds = append(uniqueThresholds, threshold)
	}
	return uniqueThresholds
}

// fit trains the random forest model
func (rf *RandomForest) fit(data [][]float64, labels []int, numTrees, maxDepth, minSize int) {
	rf.Trees = make([]*TreeNode, numTrees)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < numTrees; i++ {
		sampleData, sampleLabels := bootstrapSample(data, labels)
		rf.Trees[i] = buildTree(sampleData, sampleLabels, maxDepth, minSize)
	}
}

// bootstrapSample generates a bootstrap sample of the dataset
func bootstrapSample(data [][]float64, labels []int) ([][]float64, []int) {
	sampleData, sampleLabels := make([][]float64, len(data)), make([]int, len(labels))
	for i := range data {
		index := rand.Intn(len(data))
		sampleData[i] = data[index]
		sampleLabels[i] = labels[index]
	}
	return sampleData, sampleLabels
}

// predict makes a prediction for a single data point
func (rf *RandomForest) predict(data []float64) int {
	labelCounts := make(map[int]int)
	for _, tree := range rf.Trees {
		label := predictTree(tree, data)
		labelCounts[label]++
	}
	return majorityLabelFromCounts(labelCounts)
}

// predictTree makes a prediction using a single decision tree
func predictTree(node *TreeNode, data []float64) int {
	if node.Left == nil && node.Right == nil {
		return node.Label
	}
	if data[node.FeatureIndex] <= node.Threshold {
		return predictTree(node.Left, data)
	}
	return predictTree(node.Right, data)
}

// majorityLabelFromCounts returns the majority label from label counts
func majorityLabelFromCounts(labelCounts map[int]int) int {
	majorityLabel, maxCount := -1, 0
	for label, count := range labelCounts {
		if count > maxCount {
			majorityLabel, maxCount = label, count
		}
	}
	return majorityLabel
}

// accuracy calculates the accuracy of the model

func (rf *RandomForest) accuracy(data [][]float64, labels []int) float64 {
	correct := 0
	for i, row := range data {
		if rf.predict(row) == labels[i] {
			correct++
		}
	}
	return float64(correct) / float64(len(data))
}

// NewRandomForest creates a new random
func NewRandomForest() *RandomForest {
	return &RandomForest{}
}

func main(){
	// Load the dataset
	data, labels := LoadData("data.csv")

	// Split the dataset into training and testing sets
	trainData, trainLabels, testData, testLabels := TrainTestSplit(data, labels, 0.8)

	// Create a random forest model
	rf := NewRandomForest()

	// Train the model
	rf.fit(trainData, trainLabels, 100, 10, 1)

	// Make predictions on the test set
	accuracy := rf.accuracy(testData, testLabels)
	fmt.Println("Accuracy:", accuracy)
}
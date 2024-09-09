package panditas

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// DataFrame representa una estructura de datos similar a un DataFrame de Pandas.
type DataFrame struct {
	Headers       []string
	Rows          [][]float64
	columnIndices map[string]int
}

// ReadCSV lee un archivo CSV y retorna un DataFrame.
func ReadCSV(filename string) (*DataFrame, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open file %s: %v", filename, err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not read file %s: %v", filename, err)
	}

	if len(records) == 0 {
		return &DataFrame{}, nil
	}

	headers := records[0]
	var rows [][]float64
	columnIndices := make(map[string]int)

	for i, header := range headers {
		columnIndices[header] = i
	}

	for _, record := range records[1:] {
		row := make([]float64, len(record))
		for i, value := range record {
			floatValue, err := strconv.ParseFloat(value, 64)
			if err != nil {
				return nil, fmt.Errorf("could not parse value %s: %v", value, err)
			}
			row[i] = floatValue
		}
		rows = append(rows, row)
	}

	// print the nyumbers of rows and cols
	fmt.Println("Rows:", len(rows))
	fmt.Println("Columns:", len(rows[0]))

	return &DataFrame{
		Headers:       headers,
		Rows:          rows,
		columnIndices: columnIndices,
	}, nil
}

// Print muestra el contenido del DataFrame.
func (df *DataFrame) Print() {
	fmt.Println("Headers:", df.Headers)
	for i, row := range df.Rows {
		fmt.Printf("Row %d: %v\n", i, row)
	}
}

// GetColumn devuelve los valores de una columna dada por su nombre.
func (df *DataFrame) GetColumn(columnName string) ([]float64, error) {
	index, exists := df.columnIndices[columnName]
	if !exists {
		return nil, fmt.Errorf("column %s does not exist", columnName)
	}

	column := make([]float64, len(df.Rows))
	for i, row := range df.Rows {
		column[i] = row[index]
	}

	return column, nil
}

// GetFeaturesAndLabels separa las características y las etiquetas.
func (df *DataFrame) GetFeaturesAndLabels(labelColumn string) ([][]float64, []float64, error) {
	labelIndex, exists := df.columnIndices[labelColumn]
	if !exists {
		return nil, nil, fmt.Errorf("label column %s does not exist", labelColumn)
	}

	var features [][]float64
	var labels []float64

	for _, row := range df.Rows {
		label := row[labelIndex]
		featureRow := append([]float64{}, row[:labelIndex]...)
		features = append(features, featureRow)
		labels = append(labels, label)
	}

	return features, labels, nil
}


// Function to save the csv
func (df *DataFrame) SaveCSV(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create file %s: %v", filename, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write headers
	if err := writer.Write(df.Headers); err != nil {
		return fmt.Errorf("could not write headers: %v", err)
	}

	// Write rows
	for _, row := range df.Rows {
		strRow := make([]string, len(row))
		for i, value := range row {
			strRow[i] = strconv.FormatFloat(value, 'f', -1, 64)
		}
		if err := writer.Write(strRow); err != nil {
			return fmt.Errorf("could not write row: %v", err)
		}
	}

	return nil
}
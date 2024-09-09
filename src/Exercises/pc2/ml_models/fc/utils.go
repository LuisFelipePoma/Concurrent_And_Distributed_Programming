package fc

import (
	// "encoding/csv"
	// "fmt"
	// "os"
	// "strconv"
)

// // ReadRatingsFromCSV lee las calificaciones de los usuarios desde un archivo CSV
// func ReadRatingsFromCSV(filename string) ([]User, error) {
//     file, err := os.Open(filename)
//     if err != nil {
//         return nil, err
//     }
//     defer file.Close()

//     reader := csv.NewReader(file)
//     records, err := reader.ReadAll()
//     if err != nil {
//         return nil, err
//     }

//     userMap := make(map[int]map[int]float64)
//     for _, record := range records[1:] { // Skip header
//         userID, _ := strconv.Atoi(record[0])
//         itemID, _ := strconv.Atoi(record[1])
//         score, _ := strconv.ParseFloat(record[2], 64)

//         if _, exists := userMap[userID]; !exists {
//             userMap[userID] = make(map[int]float64)
//         }
//         userMap[userID][itemID] = score
//     }

//     var users []User
//     for userID, ratings := range userMap {
//         users = append(users, User{ID: userID, Ratings: ratings})
//     }
// 	// Print users len, csv len
// 	fmt.Println("Users:", len(users))
// 	fmt.Println("Total reviews:", len(records))

//     return users, nil
// }

package fc

import (
	"sort"
	"sync"
)

// Función para encontrar los usuarios más similares a un usuario dado
func mostSimilarUsersC(users []User, userIndex int) []int {
	similarities := make(map[int]float64)
	var wg sync.WaitGroup
	mu := &sync.Mutex{}

	for i := 0; i < len(users); i++ {
		if i != userIndex {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				similarity := cosineSimilarity(users[userIndex].Ratings, users[i].Ratings)
				mu.Lock()
				similarities[i] = similarity
				mu.Unlock()
			}(i)
		}
	}

	wg.Wait()

	// Ordenar los usuarios por similitud
	type kv struct {
		Key   int
		Value float64
	}
	var sortedSimilarities []kv
	for k, v := range similarities {
		sortedSimilarities = append(sortedSimilarities, kv{k, v})
	}
	// Ordenar en orden descendente
	sort.Slice(sortedSimilarities, func(i, j int) bool {
		return sortedSimilarities[i].Value > sortedSimilarities[j].Value
	})

	// Devolver los índices de los usuarios más similares
	var mostSimilar []int
	for _, kv := range sortedSimilarities {
		mostSimilar = append(mostSimilar, kv.Key)
	}
	return mostSimilar
}

// Función para recomendar ítems a un usuario basado en usuarios similares
func RecommendItemsC(users []User, userIndex int, numRecommendations int) []int {
	similarUsers := mostSimilarUsersC(users, userIndex)

	recommendations := make(map[int]float64)
	var wg sync.WaitGroup
	mu := &sync.Mutex{}

	for _, similarUser := range similarUsers {
		wg.Add(1)
		go func(similarUser int) {
			defer wg.Done()
			for itemID, rating := range users[similarUser].Ratings {
				// Si el usuario no ha calificado este ítem
				if _, exists := users[userIndex].Ratings[itemID]; !exists {
					mu.Lock()
					recommendations[itemID] += rating
					mu.Unlock()
				}
			}
		}(similarUser)
	}

	wg.Wait()

	// Ordenar las recomendaciones por las calificaciones acumuladas
	type kv struct {
		Key   int
		Value float64
	}
	var sortedRecommendations []kv
	for k, v := range recommendations {
		sortedRecommendations = append(sortedRecommendations, kv{k, v})
	}
	// Ordenar en orden descendente
	sort.Slice(sortedRecommendations, func(i, j int) bool {
		return sortedRecommendations[i].Value > sortedRecommendations[j].Value
	})

	// Devolver los índices de los ítems recomendados
	var recommendedItems []int
	for i := 0; i < numRecommendations && i < len(sortedRecommendations); i++ {
		recommendedItems = append(recommendedItems, sortedRecommendations[i].Key)
	}

	return recommendedItems
}

package main

func MapToSlice(mapData map[string]string) [][]string {
	var result [][]string
	for key, value := range mapData {
		result = append(result, []string{key, value})
	}
	return result
}

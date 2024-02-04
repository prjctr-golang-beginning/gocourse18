package helpers

func SplitMap(where map[string]any) ([]string, []any) {
	var keys []string
	var values []any

	for key, value := range where {
		keys = append(keys, key)
		values = append(values, value)
	}

	return keys, values
}

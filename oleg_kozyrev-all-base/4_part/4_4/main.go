package main

import (
	"fmt"
)

func main() {
	oldMap := map[string][]string{
		"group1": {"apple", "banana"},
		"group2": {"carrot"},
	}

	fmt.Println(oldMap)

	newValues := []string{"banana", "cherry", "cherry"}

	key := "group1"

	MergeToMap(newValues, key, oldMap)
	MergeToMap(newValues, "group3", oldMap)

	fmt.Println(oldMap)
}

func MergeToMap(values []string, key string, data map[string][]string) {
	uniqueExists := make(map[string]struct{})

	for _, v := range data[key] {
		uniqueExists[v] = struct{}{}
	}

	for _, v := range values {
		if _, ok := uniqueExists[v]; !ok {
			data[key] = append(data[key], v)
			uniqueExists[v] = struct{}{}
		}
	}
}

package main

import (
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]interface{} {
	var result []map[string]interface{}

	if len(data) == 0 {
		return []map[string]interface{}{}
	}

	for _, entry := range data {
		fields := strings.Split(entry, ";")
		person := make(map[string]interface{})

		if len(fields) > 0 {
			person["name"] = fields[0]
		}
		if len(fields) > 1 {
			age, err := strconv.Atoi(fields[1])
			if err == nil {
				person["age"] = age
			}
		}
		if len(fields) > 2 {
			person["address"] = fields[2]
		}
		if len(fields) > 3 && fields[3] != "" {
			height, err := strconv.ParseFloat(fields[3], 64)
			if err == nil {
				person["height"] = height
			}
		}
		if len(fields) > 4 && fields[4] != "" {
			isMarried, err := strconv.ParseBool(fields[4])
			if err == nil {
				person["isMarried"] = isMarried
			}
		}

		result = append(result, person)
	}

	return result
}

package utils

import (
	"log/slog"
	"net/url"
	"strconv"
	"strings"
)

func ProccessQuery(query url.Values, paramsDefaultValues map[string]string) map[string]string {
	proccessedQuery := make(map[string]string)

	for param, defaultValue := range paramsDefaultValues {

		if !query.Has(param) {
			proccessedQuery[param] = defaultValue
			continue
		}

		proccessedQuery[param] = query.Get(param)
	}

	return proccessedQuery
}

func PaginateText(text string) map[int]string {
	paginatedText := make(map[int]string)

	for i, verse := range strings.Split(text, "\n\n") {
		paginatedText[i+1] = verse
	}

	return paginatedText
}

func ConvertToInt(n string) int {
	i, err := strconv.Atoi(n)

	if err != nil {
		slog.Error("Failed to parse int.", slog.Any("error", err))
		return -1
	}

	return i
}

func AddPercentageToString(s string) string {
	return "%" + s + "%"
}

package utils

import (
	"strconv"
)

/*** page, limit */
func GetPaginatorSql(h map[string][]string) (int, int) {
	var page = 0
	var limit = 5

	if h["Limit"] != nil {
		Li, err := strconv.Atoi(h["Limit"][0])
		if err == nil {
			limit = Li
		}
	}

	if h["Page"] != nil {
		P, err := strconv.Atoi(h["Page"][0])
		if err == nil {
			page = (P - 1) * limit
		}
	}
	return page, limit
}

/*** page, limit */
func GetPaginatorAwsSql(h map[string]string) (int, int) {
	var page = 0
	var limit = 5

	if l, ok := h["limit"]; ok {
		Li, err := strconv.Atoi(l)
		if err == nil {
			limit = Li
		}
	}

	if p, ok := h["Page"]; ok {
		P, err := strconv.Atoi(p)
		if err == nil {
			page = (P - 1) * limit
		}
	}

	return page, limit
}

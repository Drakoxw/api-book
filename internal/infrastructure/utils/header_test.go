package utils_test

import (
	"api-book/internal/infrastructure/utils"
	"testing"
)

func TestGetPaginatorAwsSql(t *testing.T) {
	h := make(map[string]string)
	pageExpect := 0
	limitExpect := 10
	h["page"] = "1"
	h["limit"] = "10"
	page, limit := utils.GetPaginatorAwsSql(h)
	if page == pageExpect && limit == limitExpect {
		t.Log("Page y Limit son correctos")
	} else {
		t.Errorf("Se esperaban un Page: %v y un Limit: %v y se optuvieron Page: %v y un Limit: %v", pageExpect, limitExpect, page, limit)
		t.Fail()
	}
}
func TestGetPaginatorSql(t *testing.T) {
	h := make(map[string][]string)
	pageExpect := 0
	limitExpect := 10
	h["Page"] = []string{"1"}
	h["Limit"] = []string{"10"}
	page, limit := utils.GetPaginatorSql(h)
	if page == pageExpect && limit == limitExpect {
		t.Log("Page y Limit son correctos")
	} else {
		t.Errorf("Se esperaban un Page: %v y un Limit: %v y se optuvieron Page: %v y un Limit: %v", pageExpect, limitExpect, page, limit)
		t.Fail()
	}
}

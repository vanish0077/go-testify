package main

import (
	"net/http"
	"strconv"
	"strings"
)

var cafeList = map[string][]string{
	"moscow": []string{"Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"},
}

func mainHandle(w http.ResponseWriter, req *http.Request) {
	countStr := req.URL.Query().Get("count")
	if countStr == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("count missing"))
		return
	}

	count, err := strconv.Atoi(countStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong count value"))
		return
	}

	city := req.URL.Query().Get("city")

	cafe, ok := cafeList[city]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("wrong city value"))
		return
	}

	if count > len(cafe) {
		count = len(cafe)
	}

	answer := strings.Join(cafe[:count], ",")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(answer))
}

/*
func TestMainHandlerWhenCountMoreThanTotal(t *testing.T) {
	totalCount := 4
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=5&city=moscow", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.NotEmpty(t, responseRecorder.Code)

	status := responseRecorder.Code
	require.Equal(t, http.StatusOK, status)

	body := responseRecorder.Body.String()
	assert.NotEmpty(t, body)

	list := strings.Split(body, ",")
	assert.Len(t, list, totalCount)

}

func TestMainHandlerSuccess(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=3&city=moscow", nil)
	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)

	require.NotEmpty(t, responseRecorder.Code)

	assert.NotNil(t, responseRecorder.Body)

	status := responseRecorder.Code
	require.Equal(t, http.StatusOK, status)
}

func TestMainHandlerWrongCity(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/cafe?count=4&city=UnExistsCity", nil)

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(mainHandle)
	handler.ServeHTTP(responseRecorder, req)
	require.Equal(t, http.StatusBadRequest, responseRecorder.Code)

	body := responseRecorder.Body.String()
	wrongCity := "wrong city value"
	assert.Equal(t, wrongCity, body)
}
*/

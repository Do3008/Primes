package main

import (
	"testing"
	// "fmt"
	"net/http"
	"net/http/httptest"
	"github.com/stretchr/testify/assert"
	"github.com/gorilla/mux"
	"encoding/json"
)


func Router()  *mux.Router {
	router :=mux.NewRouter()
	router.HandleFunc("/primes/sum/{num}", sumPrimes)
	return router
}

func TestPrimeNormalCase(t *testing.T)  {
	request, _ := http.NewRequest("GET", "/primes/sum/10", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code)
// 	expected := `{"Sum":17}
// `
	var data map[string]interface{}
    err := json.Unmarshal([]byte(response.Body.String()), &data)
    if err != nil {
        panic(err)
    }
	// fmt.Println(data["Sum"])
	// assert.Equal(t, 17, data["Sum"])
	assert.Equal(t, 17, int(data["Sum"].(float64)))
	// fmt.Println(json.NewDecoder(response.Body))
}
func TestPrimeNormalCase1(t *testing.T)  {
	request, _ := http.NewRequest("GET", "/primes/sum/3000000", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code)
	var data map[string]interface{}
    err := json.Unmarshal([]byte(response.Body.String()), &data)
    if err != nil {
        panic(err)
    }
	assert.Equal(t, 312471072265, int(data["Sum"].(float64)))
}

func TestPrimeAbnormalCase(t *testing.T)  {
	request, _ := http.NewRequest("GET", "/primes/sum/-1", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code)
	var data map[string]interface{}
    err := json.Unmarshal([]byte(response.Body.String()), &data)
    if err != nil {
        panic(err)
	}
	assert.Equal(t, "Invalid param", data["Message"])
}

func TestPrimeAbnormalCase1(t *testing.T)  {
	request, _ := http.NewRequest("GET", "/primes/sum/sdf", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code)
	var data map[string]interface{}
    err := json.Unmarshal([]byte(response.Body.String()), &data)
    if err != nil {
        panic(err)
	}
	assert.Equal(t, "Invalid param", data["Message"])
}
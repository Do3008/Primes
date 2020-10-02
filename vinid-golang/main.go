package main

import (
 "strconv"
 "fmt"
 "log"
 "net/http"
 "encoding/json"
 "github.com/gorilla/mux"
)

type Prime struct {
	Sum int `json:"Sum"`
}

type Error struct {
	Message string `json:"Message"`
}

func sumPrimes(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)


    // string to int
    numInt, err := strconv.Atoi(params["num"])
    if err != nil {
        // handle error
		// fmt.Println(err)
		error := &Error{Message: "Invalid param"}
		json.NewEncoder(w).Encode(error)
		return
	}
	if numInt < 2 {
		error := &Error{Message: "Invalid param"}
		json.NewEncoder(w).Encode(error)
		return
	}
    // fmt.Println(params["num"], numInt)

//
	sieve := make([]bool, numInt)
	sum := 0

	for i := 2; i < numInt; i++ {
		if (!sieve[i]) {
			sum += i;
			for j := i << 1; j < numInt; j += i {
				sieve[j] = true;
			}
		}
	}
//
	prime := &Prime{Sum: sum}

	// fmt.Println("Endpoint primes")
	// fmt.Println(params)

	json.NewEncoder(w).Encode(prime)
}


func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, `<p>Go to <a href="https://github.com/Do3008/Primes">Mulapese</a>  for guidelines</p>`)
}

func handleRequests(){
	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/primes/sum/{num}", sumPrimes)
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func main(){
	handleRequests()
}
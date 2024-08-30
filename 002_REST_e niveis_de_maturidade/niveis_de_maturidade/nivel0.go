//The Swamp of POX: A API trata uma requisição como uma chamada de método com parâmetros, sem considerar os recursos HTTP.
//RPC: O código em Go simula uma chamada RPC simples onde o cliente passa um método e parâmetros em um payload JSON.

package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func rpcHandler(w http.ResponseWriter, r *http.Request) {
	var params map[string]interface{}
	json.NewDecoder(r.Body).Decode(&params)

	method := params["method"].(string)
	if method == "add" {
		a := params["a"].(float64)
		b := params["b"].(float64)
		result := a + b
		json.NewEncoder(w).Encode(map[string]float64{"result": result})
	} else {
		http.Error(w, "Unknown method", http.StatusBadRequest)
	}
}

func main() {
	http.HandleFunc("/rpc", rpcHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

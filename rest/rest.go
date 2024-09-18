package rest

import (
	"advancedGolang/blockchain"
	"advancedGolang/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var port string

type url string

func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}

type urlDescription struct {
	URL         url    `json:"url"`
	Method      string `json:"method"`
	Description string `json:"description"`
}

type addBlocklBody struct {
	Message string `json:"message"`
}

type errorResponse struct {
	errorResponse string `json:"errorMessage"`
}

//func (u urlDescription) String() string {
//	return "url"
//}

func documentation(rw http.ResponseWriter, req *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      http.MethodGet,
			Description: `See Documentation`,
		},
		{
			URL:         url("/block"),
			Method:      http.MethodPost,
			Description: `Add A Block`,
		},
		{
			URL:         url("/block/{hash}"),
			Method:      http.MethodGet,
			Description: `See A Block`,
		},
	}
	fmt.Println(data)

	// go to json
	//b, err := json.Marshal(data)
	//utils.HandleError(err)
	//fmt.Fprintf(rw, "%s", b)
	json.NewEncoder(rw).Encode(data)
}

func blocks(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:

		json.NewEncoder(rw).Encode(blockchain.Blockchain().Blocks())
	case http.MethodPost:
		var addBlockBody addBlocklBody
		utils.HandleError(json.NewDecoder(req.Body).Decode(&addBlockBody))
		blockchain.Blockchain().Blocks()
		rw.WriteHeader(http.StatusCreated)
	}
}
func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	hash := vars["hash"]

	block, err := blockchain.FindBlock(hash)
	encoder := json.NewEncoder(rw)

	if err == blockchain.ErrNotFound {
		encoder.Encode(errorResponse{fmt.Sprint(err)})
	} else {
		encoder.Encode(block)
	}

}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func Start(aPort int) {
	port = fmt.Sprintf(":%d", aPort)
	router := mux.NewRouter()
	router.Use(jsonContentTypeMiddleware)

	router.HandleFunc("/", documentation).Methods(http.MethodGet)
	router.HandleFunc("/blocks", blocks).Methods(http.MethodPost, http.MethodGet)
	router.HandleFunc("/blocks/{hash:[a-f0-9]+}", block).Methods(http.MethodGet)
	fmt.Printf("Listening on http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, router))
}

package controller

import(
	"ct-feedback-manager/manager"
	"net/http"
	"encoding/json"
	"io"
	"io/ioutil"
	"github.com/gorilla/mux"
)

/*
* GET request to get all evolutions
*
* @param http.ResponseWriter w
* @param http.Request r
*/
func GetEvolutions(w http.ResponseWriter, r *http.Request) {
    evolutions := manager.GetEvolutions()

  	w.Header().Set("Access-Control-Allow-Origin", "*")

  	if err := json.NewEncoder(w).Encode(evolutions); err != nil {
      panic(err)
    }
}

/*
* GET request to get a evolution by its ID
*
* @param http.ResponseWriter w
* @param http.Request r
*/
func GetEvolution(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

  evolution := manager.GetEvolution(vars["id"])
  if evolution == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(evolution); err != nil {
    panic(err)
  }
}

/*
* POST request to create a new Evolution object
*/
func CreateEvolution(w http.ResponseWriter, r *http.Request) {
  var body []byte
  var err error
	if body, err = ioutil.ReadAll(io.LimitReader(r.Body, 1048576)); err != nil {
    panic(err)
  }
	if err = r.Body.Close(); err != nil {
    panic(err)
  }
  var data map[string]interface{}
	if err = json.Unmarshal(body, &data); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err = json.NewEncoder(w).Encode(err); err != nil {
  			panic(err)
		}
	}

  evolution := manager.CreateEvolution(
		data["title"].(string),
		data["description"].(string),
		data["status"].(string),
		data["author"].(map[string]interface{}),
	)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(&evolution); err != nil {
    panic(err)
  }
}

/*
* PUT request to update an evolution by its ID
*/
func UpdateEvolution(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var body []byte
  var err error
	if body, err = ioutil.ReadAll(io.LimitReader(r.Body, 1048576)); err != nil {
    panic(err)
  }
	if err = r.Body.Close(); err != nil {
    panic(err)
  }
  var data map[string]interface{}
	if err = json.Unmarshal(body, &data); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err = json.NewEncoder(w).Encode(err); err != nil {
  			panic(err)
		}
	}

  evolution := manager.UpdateEvolution(vars["id"], data)
  if evolution == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(&evolution); err != nil {
    panic(err)
  }
}

/*
* DELETE request to delete an evolution by its ID
*/
func DeleteEvolution(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

  if !manager.DeleteEvolution(vars["id"]) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

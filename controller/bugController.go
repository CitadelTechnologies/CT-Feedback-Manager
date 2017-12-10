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
* GET request to get all bugs
*/
func GetBugs(w http.ResponseWriter, r *http.Request) {
    bugs := manager.GetBugs()

  	w.Header().Set("Access-Control-Allow-Origin", "*")

  	if err := json.NewEncoder(w).Encode(bugs); err != nil {
      panic(err)
    }
}

/*
* GET request to get a bug by its ID
*/
func GetBug(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

  bug := manager.GetBug(vars["id"])
  if bug == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(bug); err != nil {
    panic(err)
  }
}

/*
* POST request to create a new Bug object
*/
func CreateBug(w http.ResponseWriter, r *http.Request) {
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

  bug := manager.CreateBug(
		data["title"].(string),
		data["description"].(string),
		data["status"].(string),
		data["author"].(map[string]interface{}),
	)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(&bug); err != nil {
    panic(err)
  }
}

/*
* PUT request to update a bug by its ID
*/
func UpdateBug(w http.ResponseWriter, r *http.Request) {
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

  bug := manager.UpdateBug(vars["id"], data)
  if bug == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	if err := json.NewEncoder(w).Encode(&bug); err != nil {
    panic(err)
  }
}

/*
* DELETE request to delete a bug by its ID
*/
func DeleteBug(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

  if !manager.DeleteBug(vars["id"]) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

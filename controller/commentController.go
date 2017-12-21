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
* GET request to get all feedback comments
*/
func GetFeedbackComments(w http.ResponseWriter, r *http.Request, collectionName string) {
  	vars := mux.Vars(r)

    comments := manager.GetFeedbackComments(vars["id"], collectionName)

  	w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

  	if err := json.NewEncoder(w).Encode(comments); err != nil {
      panic(err)
    }
}

/*
* POST request to create a new Evolution object
*/
func CreateComment(w http.ResponseWriter, r *http.Request, collectionName string) {
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

  comment := manager.CreateComment(
    vars["id"],
		data["content"].(string),
		data["author"].(map[string]interface{}),
    collectionName,
	)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(&comment); err != nil {
    panic(err)
  }
}

/*
* PUT request to update a comment by its ID
*/
func UpdateComment(w http.ResponseWriter, r *http.Request, collectionName string) {
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

  comment := manager.UpdateComment(vars["id"], vars["comment_id"], data, collectionName)
  if comment == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(&comment); err != nil {
    panic(err)
  }
}

/*
* DELETE request to delete a comment by its ID
*/
func DeleteComment(w http.ResponseWriter, r *http.Request, collectionName string) {
	vars := mux.Vars(r)

  if !manager.DeleteComment(vars["id"], vars["comment_id"], collectionName) {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

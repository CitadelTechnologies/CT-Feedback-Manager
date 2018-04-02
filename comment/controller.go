package comment

import(
	"net/http"
	"encoding/json"
	"io"
	"io/ioutil"
	"github.com/gorilla/mux"
)

func GetFeedbackCommentsAction(w http.ResponseWriter, r *http.Request) {
  	vars := mux.Vars(r)

    comments := GetFeedbackComments(vars["id"])

  	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

  	if err := json.NewEncoder(w).Encode(comments); err != nil {
      	panic(err)
    }
}

func CreateCommentAction(w http.ResponseWriter, r *http.Request) {
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

	comment, err := CreateComment(
		vars["id"],
		data["content"].(string),
		data["author"].(map[string]interface{}),
	)
	if err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		if err = json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(&comment); err != nil {
		panic(err)
	}
}

func UpdateCommentAction(w http.ResponseWriter, r *http.Request) {
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

	comment := UpdateComment(vars["id"], vars["comment_id"], data)
	if comment == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(&comment); err != nil {
		panic(err)
	}
}

func DeleteCommentAction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if err := DeleteComment(vars["id"], vars["comment_id"]); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

package label

import(
	"net/http"
	"encoding/json"
	"io"
	"io/ioutil"
	"github.com/gorilla/mux"
)

func GetLabelsAction(w http.ResponseWriter, r *http.Request) {
    labels := GetLabels()

  	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

  	if err := json.NewEncoder(w).Encode(labels); err != nil {
		panic(err)
    }
}

func CreateLabelAction(w http.ResponseWriter, r *http.Request) {
	var body []byte
	var err error
	if body, err = ioutil.ReadAll(io.LimitReader(r.Body, 1048576)); err != nil {
		panic(err)
	}
	if err = r.Body.Close(); err != nil {
		panic(err)
	}
	var data map[string]string
	if err = json.Unmarshal(body, &data); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err = json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	label := CreateLabel(
		data["name"],
		data["color"],
	)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(&label); err != nil {
		panic(err)
	}
}

func UpdateLabelAction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	var body []byte
	var err error
	if body, err = ioutil.ReadAll(io.LimitReader(r.Body, 1048576)); err != nil {
		panic(err)
	}
	if err = r.Body.Close(); err != nil {
		panic(err)
	}
	var data map[string]string
	if err = json.Unmarshal(body, &data); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err = json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	label := UpdateLabel(vars["id"], data)
	if label == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(&label); err != nil {
		panic(err)
	}
}

func DeleteLabelAction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if err := DeleteLabel(vars["id"]); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

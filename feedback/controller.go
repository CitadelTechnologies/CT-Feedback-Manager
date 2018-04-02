package feedback

import(
    "ct-feedback-manager/label"
    "net/http"
    "encoding/json"
    "io"
    "io/ioutil"
    "github.com/gorilla/mux"
)

func GetFeedbacksAction(w http.ResponseWriter, r *http.Request) {
    feedbacks := GetFeedbacks()

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.Header().Set("Access-Control-Allow-Origin", "*")

    if err := json.NewEncoder(w).Encode(feedbacks); err != nil {
        panic(err)
    }
}

func GetFeedbackAction(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)

    feedback := GetFeedback(vars["id"])
    if feedback == nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    if err := json.NewEncoder(w).Encode(feedback); err != nil {
        panic(err)
    }
}

func CreateFeedbackAction(w http.ResponseWriter, r *http.Request) {
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

    feedback, err := CreateFeedback(
        data["title"].(string),
        data["type"].(string),
        data["description"].(string),
        data["status"].(string),
        data["author"].(map[string]interface{}),
    )
    if err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(err)
    }

    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusCreated)
    if err = json.NewEncoder(w).Encode(&feedback); err != nil {
        panic(err)
    }
}

func UpdateFeedbackAction(w http.ResponseWriter, r *http.Request) {
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

    feedback := UpdateFeedback(vars["id"], data)
    if feedback == nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    if err != nil {
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(http.StatusBadRequest)
        json.NewEncoder(w).Encode(err)
    }
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    if err := json.NewEncoder(w).Encode(&feedback); err != nil {
        panic(err)
    }
}

func AddLabelToFeedbackAction(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)

    label := label.GetLabel(vars["label_id"])
    if label == nil {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    feedback := AddLabelToFeedback(vars["feedback_id"], label)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    if err := json.NewEncoder(w).Encode(&feedback); err != nil {
        panic(err)
    }
}

func RemoveLabelFromFeedbackAction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	label := label.GetLabel(vars["label_id"])
	if label == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	feedback := RemoveLabelFromFeedback(vars["feedback_id"], label)
	if err := json.NewEncoder(w).Encode(&feedback); err != nil {
        panic(err)
    }
}

func DeleteFeedbackAction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

    if err := DeleteFeedback(vars["id"]); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

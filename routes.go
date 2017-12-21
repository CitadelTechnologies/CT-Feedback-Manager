package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"ct-feedback-manager/controller"
)

type(
	Route struct {
		Name        string
		Method      string
		Pattern     string
		HandlerFunc http.HandlerFunc
	}
	Routes []Route
)

func NewRouter() *mux.Router {

    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
			router.HandleFunc(route.Pattern, route.HandlerFunc).Methods(route.Method)
    }
    return router
}

var routes = Routes{
		Route{
				"Create Evolution",
				"POST",
				"/evolutions",
				controller.CreateEvolution,
		},
		Route{
				"Update Evolution",
				"PUT",
				"/evolutions/{id}",
				controller.UpdateEvolution,
		},
		Route{
				"Delete Evolution",
				"DELETE",
				"/evolutions/{id}",
				controller.DeleteEvolution,
		},
		Route{
				"Get Evolutions",
				"GET",
				"/evolutions",
				controller.GetEvolutions,
		},
		Route{
				"Get Evolution",
				"GET",
				"/evolutions/{id}",
				controller.GetEvolution,
		},
    Route{
        "Bugs",
        "GET",
        "/bugs",
        controller.GetBugs,
    },
    Route{
        "Create Bug",
        "POST",
        "/bugs",
        controller.CreateBug,
    },
    Route{
        "Update Bug",
        "PUT",
        "/bugs/{id}",
        controller.UpdateBug,
    },
    Route{
        "Add Label to Bug",
        "POST",
        "/bugs/{feedback_id}/labels/{label_id}",
        controller.AddLabelToBug,
    },
    Route{
        "Remove Label from Bug",
        "DELETE",
        "/bugs/{feedback_id}/labels/{label_id}",
        controller.RemoveLabelFromBug,
    },
    Route{
        "Delete Bug",
        "DELETE",
        "/bugs/{id}",
        controller.DeleteBug,
    },
    Route{
        "Get Bug",
        "GET",
        "/bugs/{id}",
        controller.GetBug,
    },
    Route{
        "Create Bug Comment",
        "POST",
        "/bugs/{id}/comments",
        func (w http.ResponseWriter, r *http.Request) {
					controller.CreateComment(w, r, "bugs")
				},
    },
    Route{
        "Update Bug Comment",
        "PUT",
        "/bugs/{id}/comments/{comment_id}",
        func (w http.ResponseWriter, r *http.Request) {
					controller.UpdateComment(w, r, "bugs")
				},
    },
    Route{
        "Delete Bug Comment",
        "DELETE",
        "/bugs/{id}/comments/{comment_id}",
        func (w http.ResponseWriter, r *http.Request) {
					controller.DeleteComment(w, r, "bugs")
				},
    },
    Route{
        "Get Bug Comments",
        "GET",
        "/bugs/{id}/comments",
        func (w http.ResponseWriter, r *http.Request) {
					controller.GetFeedbackComments(w, r, "bugs")
				},
    },
    Route{
        "Create Evolution Comment",
        "POST",
        "/evolutions/{id}/comments",
        func (w http.ResponseWriter, r *http.Request) {
					controller.CreateComment(w, r, "evolutions")
				},
    },
    Route{
        "Update Evolution Comment",
        "PUT",
        "/evolutions/{id}/comments/{comment_id}",
        func (w http.ResponseWriter, r *http.Request) {
					controller.UpdateComment(w, r, "evolutions")
				},
    },
    Route{
        "Add Label to Evolution",
        "POST",
        "/evolutions/{feedback_id}/labels/{label_id}",
        controller.AddLabelToEvolution,
    },
    Route{
        "Remove Label from Evolution",
        "DELETE",
        "/evolutions/{feedback_id}/labels/{label_id}",
        controller.RemoveLabelFromEvolution,
    },
    Route{
        "Delete Evolution Comment",
        "DELETE",
        "/evolutions/{id}/comments/{comment_id}",
        func (w http.ResponseWriter, r *http.Request) {
					controller.DeleteComment(w, r, "evolutions")
				},
    },
    Route{
        "Get Evolution Comments",
        "GET",
        "/evolutions/{id}/comments",
        func (w http.ResponseWriter, r *http.Request) {
					controller.GetFeedbackComments(w, r, "evolutions")
				},
    },
    Route{
        "Create Label",
        "POST",
        "/labels",
        func (w http.ResponseWriter, r *http.Request) {
					controller.CreateLabel(w, r)
				},
    },
    Route{
        "Update Label",
        "PUT",
        "/labels/{id}",
        func (w http.ResponseWriter, r *http.Request) {
					controller.UpdateLabel(w, r)
				},
    },
    Route{
        "Delete Label",
        "DELETE",
        "/labels/{id}",
        func (w http.ResponseWriter, r *http.Request) {
					controller.DeleteLabel(w, r)
				},
    },
    Route{
        "Get Labels",
        "GET",
        "/labels",
        func (w http.ResponseWriter, r *http.Request) {
					controller.GetLabels(w, r)
				},
    },
}

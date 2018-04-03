package main

import (
	"net/http"

	"github.com/gorilla/mux"

	"ct-feedback-manager/comment"
	"ct-feedback-manager/feedback"
	"ct-feedback-manager/label"
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
		"Create Feedback",
		"POST",
		"/feedbacks",
		feedback.CreateFeedbackAction,
	},
	Route{
		"Update Feedback",
		"PUT",
		"/feedbacks/{id}",
		feedback.UpdateFeedbackAction,
	},
	Route{
		"Delete Feedback",
		"DELETE",
		"/feedbacks/{id}",
		feedback.DeleteFeedbackAction,
	},
	Route{
		"Get Feedbacks",
		"GET",
		"/feedbacks",
		feedback.GetFeedbacksAction,
	},
	Route{
		"Search Feedbacks",
		"POST",
		"/feedbacks/search",
		feedback.SearchFeedbacksAction,
	},
	Route{
		"Get Feedback",
		"GET",
		"/feedbacks/{id}",
		feedback.GetFeedbackAction,
	},
    Route{
        "Add Label to Feedback",
        "POST",
        "/feedbacks/{feedback_id}/labels/{label_id}",
        feedback.AddLabelToFeedbackAction,
    },
    Route{
        "Remove Label from Feedback",
        "DELETE",
        "/bugs/{feedback_id}/labels/{label_id}",
        feedback.RemoveLabelFromFeedbackAction,
    },
    Route{
        "Create Feedback Comment",
        "POST",
        "/feedbacks/{id}/comments",
        comment.CreateCommentAction,
    },
    Route{
        "Update Feedback Comment",
        "PUT",
        "/feedbacks/{id}/comments/{comment_id}",
        comment.UpdateCommentAction,
    },
    Route{
        "Delete Feedback Comment",
        "DELETE",
        "/feedbacks/{id}/comments/{comment_id}",
        comment.DeleteCommentAction,
    },
    Route{
        "Get Feedback Comments",
        "GET",
        "/feedbacks/{id}/comments",
        comment.GetFeedbackCommentsAction,
    },
    Route{
        "Create Label",
        "POST",
        "/labels",
        label.CreateLabelAction,
    },
    Route{
        "Update Label",
        "PUT",
        "/labels/{id}",
        label.UpdateLabelAction,
    },
    Route{
        "Delete Label",
        "DELETE",
        "/labels/{id}",
        label.DeleteLabelAction,
    },
    Route{
        "Get Labels",
        "GET",
        "/labels",
        label.GetLabelsAction,
    },
}

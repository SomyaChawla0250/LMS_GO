package controller

import (
	"net/http"

	"mvc/pkg/models"
	"mvc/pkg/views"
)

func BrowseBooks(writer http.ResponseWriter, request *http.Request) {
	
	
	booksList,err := models.FetchBooks()
	if err != nil {
		http.Error(writer, "Database error", http.StatusInternalServerError)
		return
	}
	t := views.BrowsePage()
	writer.WriteHeader(http.StatusOK)
	t.Execute(writer, booksList)
}



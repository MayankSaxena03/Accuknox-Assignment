package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/MayamkSaxena03/Accuknox/models"
	"github.com/MayamkSaxena03/Accuknox/utility"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CreateNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var requestBody models.NoteDataRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid request body")
		return
	}

	if requestBody.SID == "" || requestBody.Note == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid request body")
		return
	}

	userId, err := utility.GetUserIDFromToken(requestBody.SID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid token")
		return
	}

	note := models.NoteData{
		UserID: userId,
		Note:   requestBody.Note,
	}

	noteCollection := models.GetNoteCollection()
	result, err := noteCollection.InsertOne(r.Context(), note)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bson.M{"Id": result.InsertedID})
}

func GetNotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var requestBody models.NoteDataRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid request body")
		return
	}

	if requestBody.SID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Token")
		return
	}

	userId, err := utility.GetUserIDFromToken(requestBody.SID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid token")
		return
	}

	noteCollection := models.GetNoteCollection()
	var notes []models.NoteData
	cursor, err := noteCollection.Find(r.Context(), bson.M{"userid": userId}, &options.FindOptions{Projection: bson.M{"userid": 0}})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = cursor.All(r.Context(), &notes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	type Response struct {
		Notes []models.AllNotes `json:"notes"`
	}

	var response Response
	for _, note := range notes {
		response.Notes = append(response.Notes, models.AllNotes{ID: note.ID, Note: note.Note})
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func DeleteNote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var requestBody models.NoteDataRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid request body")
		return
	}

	if requestBody.SID == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid Token")
		return
	}

	noteCollection := models.GetNoteCollection()
	userId, err := utility.GetUserIDFromToken(requestBody.SID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Invalid token")
		return
	}

	result, err := noteCollection.DeleteOne(r.Context(), bson.M{"_id": requestBody.ID, "userid": userId})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	if result.DeletedCount == 0 {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Note not found")
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Note deleted successfully")
}

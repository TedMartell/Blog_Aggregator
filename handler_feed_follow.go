package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/TedMartell/Blog_Aggregator/internal/database"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerFeedFollowsCreate(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	feedFollow, err := cfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed_follow")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedFollowToFeedFollow(feedFollow))
}

func (cfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {

	feedFollows, err := cfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create feed_follow")
		return
	}

	respondWithJSON(w, http.StatusOK, databaseFeedFollowsToFeedFollows(feedFollows))
}

func (cfg *apiConfig) handlerFeedFollowDelete(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollowIDStr := r.PathValue("feedFollowID")
	feedFollowID, err := uuid.Parse(feedFollowIDStr)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid feed follow ID")
		return
	}

	err = cfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		ID:     feedFollowID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't delete feed follow")
		return
	}

	respondWithJSON(w, http.StatusOK, struct{}{})
}

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	internal "github.com/sakthi-lucia0567/Commonplace-Book/internal/database"
)

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	userUUID := uuid.New()
	generatedType := pgtype.UUID{Bytes: userUUID, Valid: true}
	hashPassword, _ := HashPassword(params.Password)

	user, err := apiCfg.DB.CreateUser(r.Context(), internal.CreateUserParams{
		ID:        generatedType,
		Username:  params.Username,
		Password:  hashPassword,
		Email:     params.Email,
		CreatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
		UpdatedAt: pgtype.Timestamp{Time: time.Now(), Valid: true},
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create User: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseUserToUser(user))
}

package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "profile-service/dao"
    "profile-service/model"
    "fmt"
)

// our main function
func main() {
    router := mux.NewRouter()
    router.HandleFunc("/profiles", GetProfiles).Methods("GET")
    router.HandleFunc("/profile/{id}", GetProfile).Methods("GET")
    router.HandleFunc("/profile/{id}", CreateProfile).Methods("POST")
    router.HandleFunc("/profile/{id}", DeleteProfile).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":8000", router))
}

func GetProfiles(w http.ResponseWriter, r *http.Request) {
  profiles, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, movies)
}

func GetProfile(w http.ResponseWriter, r *http.Request) {
  params := mux.Vars(r)
	profile, err := dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Movie ID")
		return
	}
	respondWithJson(w, http.StatusOK, profile)
}

func CreateProfile(w http.ResponseWriter, r *http.Request) {
  defer r.Body.Close()
	var profile Profile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	movie.ID = bson.NewObjectId()
	if err := dao.Insert(profile); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, profile)
}

func DeleteProfile(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "not implemented yet !")
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
  defer r.Body.Close()
	var profile Profile
	if err := json.NewDecoder(r.Body).Decode(&profile); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := dao.Update(profile); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

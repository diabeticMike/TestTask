package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/TestTask/logger"
	"github.com/TestTask/model"
	"github.com/TestTask/repository"
)

// NewUserController return UserController realization
func NewUserController(log logger.Logger, userRepo repository.UserRepository) *UserController {
	return &UserController{log: log, userRepo: userRepo}
}

// UserController implements user logic
type UserController struct {
	log      logger.Logger
	userRepo repository.UserRepository
}

func (uc *UserController) GetUserData(w http.ResponseWriter, r *http.Request) {
	if _, ok := r.URL.Query()["username"]; !ok {

		return
	}
	username := r.URL.Query()["username"][0]
	user, err := uc.userRepo.GetUserByUsername(username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			uc.log.Errorf("User with username %s isn't exist\n", username)
			http.Error(w, fmt.Sprintf("User with username %s isn't exist\n", username), http.StatusBadRequest)
			return
		}
		uc.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userData, err := uc.userRepo.GetUserDataByID(user.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		uc.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	userProfile, err := uc.userRepo.GetUserProfileByID(user.ID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		uc.log.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	wholeUser := newWholeUser(*user, *userData, *userProfile)

	responce, err := json.Marshal(wholeUser)
	if err != nil {
		uc.log.Errorln(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(responce); err != nil {
		uc.log.Errorln(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func newWholeUser(user model.User, userData model.UserData, userProfile model.UserProfile) model.WholeUser {
	wholeUser := model.WholeUser{}
	wholeUser.ID = user.ID
	wholeUser.Username = user.Username
	wholeUser.FirstName = userProfile.FirstName
	wholeUser.LastName = userProfile.LastName
	wholeUser.Phone = userProfile.Phone
	wholeUser.Address = userProfile.Address
	wholeUser.City = userProfile.City
	wholeUser.School = userData.School
	return wholeUser
}

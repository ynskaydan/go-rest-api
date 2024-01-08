package routes

import (
	"UserManagementProject/api/handlers"
	"github.com/gorilla/mux"
)

func SetupRoutes(userHandler *handlers.UserHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", userHandler.GetAllUsers).Methods("GET")
	r.HandleFunc("/createUser", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", userHandler.GetUserById).Methods("GET")
	r.HandleFunc("/updateUser/{id}", userHandler.UpdateUser).Methods("PUT")
	r.HandleFunc("/deleteUser/{id}", userHandler.DeleteUser).Methods("DELETE")

	return r
}

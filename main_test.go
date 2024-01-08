package main

import (
	"UserManagementProject/api/handlers"
	"UserManagementProject/internal/repository"
	"UserManagementProject/internal/service"
	"UserManagementProject/pkg/db"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAll(t *testing.T) {
	database, err := db.InitializeDB("./pkg/mydatabase.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.GetAllUsers)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `[{"id":115,"first_name":"Lorem","last_name":"Ipsum","email":"lorem@ipsum.com"},{"id":150,"first_name":"asd","last_name":"dasdsadsa","email":"dasd@gmal.com"},{"id":530,"first_name":"Oguz Kaan","last_name":"Pehlivan","email":"oguzkaan3096@gmail.com"},{"id":627,"first_name":"bb","last_name":"aa","email":"asd2@gmail.com"},{"id":716,"first_name":"Yunus","last_name":"Kaydan","email":"ynskaydan@gmail.com"}]
`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestGetById(t *testing.T) {
	database, err := db.InitializeDB("./pkg/mydatabase.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	// test with id := 1
	req, err := http.NewRequest("GET", "/user/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.GetUserById)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "{\"id\":1,\"first_name\":\"Example\",\"last_name\":\"Example\",\"email\":\"example@example.com\"}\n"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestUpdate(t *testing.T) {
	database, err := db.InitializeDB("./pkg/mydatabase.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	// test with id := 1
	req, err := http.NewRequest("PUT", "/updateUser/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.UpdateUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := ""
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestDelete(t *testing.T) {
	database, err := db.InitializeDB("./pkg/mydatabase.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	// test with id := 1
	req, err := http.NewRequest("DELETE", "/deleteUser/1", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.DeleteUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := ""
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
func TestCreate(t *testing.T) {
	database, err := db.InitializeDB("./pkg/mydatabase.db")
	if err != nil {
		log.Fatal(err)
	}
	defer database.Close()

	userRepo := repository.NewUserRepository(database)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)
	// test with id := 1
	req, err := http.NewRequest("POST", "/createUser", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(userHandler.CreateUser)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := "null"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

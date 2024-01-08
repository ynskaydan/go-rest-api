
# User Management REST-API



## Overview
This User Management REST API, written in Go, provides a set of endpoints for creating, retrieving, updating, and deleting user information. It is designed to offer a straightforward and efficient way to manage user data for applications.

## Getting Started
Ensure you have Go installed on your system.
Any additional setup requirements (e.g., environment setup).

### Prerequisites

•	Go

•	SQLite3

•	A good text editor or IDE (e.g., VSCode, Goland)





## Installation

#### 1. Clone the repository
```sh
git clone <repository-url>
cd <project-name>
go build
```
#### 2. Run project
```sh
go run main.go
```

## API Endpoints
Details about api endpoints

### '/users'
Method: GET

Description: Returns all users.

Payload: JSON object containing user array with their details(first-name,last-name email).
### '/user/{id}'
Method: GET

Description: Returns a user with id.

Payload: JSON object containing user details (first-name,last-name email).
### '/createUser'
Method: POST

Description: Creates a new user.

Payload: JSON object containing user details (first-name,last-name email).

### '/updateUser/{id}'
Method: PUT

Description: Updates a user with provided id.


### '/deleteUser/{id}'
Method: DELETE

Description: Deletes a user with provided id.


### Running the Tests
```sh
go test
```


## Authors

- Yunus Kaydan [@ynskaydan](https://www.github.com/ynskaydan)


## License

This project is licensed under the MIT License - see the LICENSE.md file for details.




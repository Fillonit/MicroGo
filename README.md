# MicroGo
MicroGo is a REST API built with Go and MongoDB. It is a simple API that allows you to create, read, update and delete users.

## Installation
1. Clone the repository
2. Install the dependencies
3. Run the server

```bash
git clone
cd microgo
go get
go run main.go
```

## Usage
### Create a user
```bash
curl -X POST -H "Content-Type: application/json" -d '{"name":"Fillonit","location":"Kosovo", "title": "Software Engineer"}' http://localhost:8080/users
```

### Get all users
```bash
curl -X GET http://localhost:8080/users
```

### Get a user
```bash
curl -X GET http://localhost:8080/users/{id}
```

### Update a user
```bash
curl -X PUT -H "Content-Type: application/json" -d '{"name":"Filloniti","location":"Kosovo", "title": "Software Engineer"}' http://localhost:8080/users/{id}
```

### Delete a user
```bash
curl -X DELETE http://localhost:8080/users/{id}
```

## License
[MIT](https://choosealicense.com/licenses/mit/)

## Technologies

![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white) ![MongoDB](https://img.shields.io/badge/MongoDB-%234ea94b.svg?style=for-the-badge&logo=mongodb&logoColor=white)
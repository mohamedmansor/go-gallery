# go-gallery
go-gallery 

# Api go simple
It is a just simple RESTful API with Go using:
1. **Gin Framework**
2. **Gorm** 

## Installation & Run
```bash
# Run this project
$ docker build -t go-gallery .
$ docker-compose up --build
```

## Structure
```
├── api 
│   ├── controller
│   │   └── image.go
│   ├── repository
│   │   └── image.go
│   ├── routes
│   │   └── image.go
│   └── service
│       └── image.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── infrastructure
│   ├── db.go
│   ├── env.go
│   └── routes.go
├── main
├── main.go
├── models
│   └── image.go
└── util
    └── response.go

```

## API

#### /images
* `GET` : Get all images
* `POST` : Create a new image

#### /images/:id
* `GET` : Get a image
* `PUT` : Update a image
* `DELETE` : Delete a image

#Post Params
```
{
	"Title": "Op Super John Doe Bilw",
	"ImageUrl": "Implementation Golang",
}
```

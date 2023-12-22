package main

import (
	"log"

	"github.com/bzhtux/tsa/internal/utils"
	"github.com/bzhtux/tsa/models"
)

func main() {
	utils.DefaultDBFile = "data/db/sqlite.d"
	log.Printf("Tanzu Sample Application Loader is running ...")
	config := utils.GetConfig()
	log.Printf("Loading SQLite file: %s", config)
	db := utils.ConnectDB(config)
	err := db.AutoMigrate(&models.HttpStatusCode{})
	if err != nil {
		log.Printf("Error Migration Model: %s", err.Error())
	}
	code101 := models.HttpStatusCode{
		Code:    101,
		Name:    "Continue",
		Desc:    "The client should continue the request or ignore the response if the request is already finished",
		Picture: "/picture/101.jpg",
	}
	db.Create(&code101)
	code200 := models.HttpStatusCode{
		Code:    200,
		Name:    "Ok",
		Desc:    "The request succeeded",
		Picture: "/picture/200.jpg",
	}
	db.Create(&code200)
	code201 := models.HttpStatusCode{
		Code:    201,
		Name:    "Created",
		Desc:    "The request succeeded and a new resource was created",
		Picture: "/picture/201.jpg",
	}
	db.Create(&code201)
	code202 := models.HttpStatusCode{
		Code:    202,
		Name:    "Accepted",
		Desc:    "The request has been accepted for processing but this is noncommital",
		Picture: "/picture/202.jpg",
	}
	db.Create(&code202)
	code301 := models.HttpStatusCode{
		Code:    301,
		Name:    "Moved permanently",
		Desc:    "The URL of the requested resource has been changed permanently. The new URL is given in the response",
		Picture: "/picture/301.jpg",
	}
	db.Create(&code301)
	code302 := models.HttpStatusCode{
		Code:    302,
		Name:    "Found",
		Desc:    "The URI of requested resource has been changed temporarily",
		Picture: "/picture/302.jpg",
	}
	db.Create(&code302)
	code400 := models.HttpStatusCode{
		Code:    400,
		Name:    "Bad request",
		Desc:    "The server cannot or will not process the request due to something that is perceived to be a client error",
		Picture: "/picture/400.jpg",
	}
	db.Create(&code400)
	code401 := models.HttpStatusCode{
		Code:    401,
		Name:    "Unauthorized",
		Desc:    "Although the HTTP standard specifies 'unauthorized', semantically this response means 'unauthenticated'. That is, the client must authenticate itself to get the requested response",
		Picture: "/picture/401.jpg",
	}
	db.Create(&code401)
	code403 := models.HttpStatusCode{
		Code:    403,
		Name:    "Forbidden",
		Desc:    "The client does not have access rights to the content; that is, it is unauthorized, so the server is refusing to give the requested resource",
		Picture: "/picture/403.jpg",
	}
	db.Create(&code403)
	code404 := models.HttpStatusCode{
		Code:    404,
		Name:    "Not found",
		Desc:    "The server cannot find the requested resource",
		Picture: "/picture/404.jpg",
	}
	db.Create(&code404)
	code418 := models.HttpStatusCode{
		Code:    418,
		Name:    "I'm a teapot",
		Desc:    "The server refuses to brew coffee because it is, permanently, a teapot",
		Picture: "/picture/418.jpg",
	}
	db.Create(&code418)
	code500 := models.HttpStatusCode{
		Code:    500,
		Name:    "Internal server error",
		Desc:    "The server has encountered a situation it does not know how to handle",
		Picture: "/picture/500.jpg",
	}
	db.Create(&code500)
	code501 := models.HttpStatusCode{
		Code:    501,
		Name:    "Not implemented",
		Desc:    "The request method is not supported by the server and cannot be handled",
		Picture: "/picture/501.jpg",
	}
	db.Create(&code501)
}

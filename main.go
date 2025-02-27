package main

import (
	"log"
	"myapi/controllers"
	"myapi/routers"
	"myapi/services"
	"net/http"
)

func main() {
	db, err := services.ConnectDB()
	if err != nil {
		log.Println("failed to connect db")
		return
	}
	ser := services.NewMyAppService(db)
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)

	r := routers.NewRouter(aCon, cCon)

	log.Print("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

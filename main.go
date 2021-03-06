package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/ilhamtubagus/urlShortener/api"
	"github.com/ilhamtubagus/urlShortener/lib"
	"github.com/kamva/mgm/v3"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	_, err := time.LoadLocation("Jakarta")
	if err != nil {
		fmt.Println(err)
	}
	if lib.JakartaTime, err = time.LoadLocation("Asia/Jakarta"); err != nil {
		panic("Error loading '" + "Asia/Jakarta" + "' as timezone location: " + err.Error())
	}
	//uncomment line below in production stage
	lib.LoadEnv(".env")
	// Setup the mgm default config
	err = mgm.SetDefaultConfig(nil, "url-shortener", options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		log.Fatal("Error while initializing database connections " + err.Error())
	}

}
func main() {
	//Create new echo instance
	e := echo.New()
	api.StartApp(e)
	p := os.Getenv("PORT")
	port, err := strconv.Atoi(p)
	if err != nil {
		e.Logger.Fatal(err)
	}
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", port)))

}

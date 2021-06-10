package main

import (
	"fmt"
	"net/http"

	TransportHTTP "github.com/NguyenTanTuan/go-rest-api/internal/transport/http"
)

//App - the struct that containt pointer
// to database connection
type App struct{}

// Run - handles the startup of our application
func (app *App) Run() error {
	fmt.Println("Setting Up Our App")
	handler := TransportHTTP.NewHandler()
	handler.SetUphandler()
	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to set up Server")
		return err
	}
	return nil
}

// Our main entrypoint for the application
func main() {
	fmt.Println("Go Rest API Course")
	app := App{}
	if err := app.Run(); err != nil {
		fmt.Println("Error Starting Up")
		fmt.Println(err)
	}
}

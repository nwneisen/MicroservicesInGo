package app

import (
	"fmt"
	"net/http"

	"github.com/nwneisen/MicroservicesInGo/Section2-MVC/mvc/controllers"
)

func StartApp() {
	http.HandleFunc("/users", controllers.GetUser)

	fmt.Println("running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

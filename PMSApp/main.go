package main

import (
	"PMSApp/app"
)

func main() {
	server := app.BuildInjector()

	_ = server.ListenAndServe()
}

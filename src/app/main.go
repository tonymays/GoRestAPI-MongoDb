package main

// ---- main go routine ----
func main() {
	// establish a nil App structure
	a := App{}

	// initialize the App struct for production and panic if it fails
	err := a.Init("prod")
	if err != nil {
		panic(err)
	}

	// run the app
	a.Run()
}

package main

// ---- main function ----
func main() {
	// let's get a nil app structure for our service
	a := App{}

	// use dependency injection to fill it in
	err := a.Init("prod")

	// if Init fails with error ...
	if err != nil {
		// ... then, stop the train right here.
		panic(err)
	}

	// ... otherwise, all aboard ... let's rock some tracks.
	a.Run()
}

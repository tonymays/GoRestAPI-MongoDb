package main

func main() {
	a := App{}
	err := a.Init("prod")
	if err != nil {
		panic(err)
	}
	a.Run()
}

package main

import "demo/bootstrap"

func main() {

	app := bootstrap.App{}
	app.Prepare()
	app.Run()

}

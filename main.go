package main

import "fmt"

func main () {
	app := newApp()
	app.start()
	fmt.Printf("Exiting %s....\n",app.cfg.APP_NAME)
}


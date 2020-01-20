package main



import "flag"

func main () {
	reset := flag.Bool("reset", false, "Varning, removing all the apps from the system...")
	flag.Parse()
	system := NewServer()
	system.Start(*reset)
}


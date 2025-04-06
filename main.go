package main

import "log"

func main() {
	store, err := NewPostgressDB()

	if err != nil {
		log.Fatal(err)
	}

	apiServer := NewAPIServer(":3002", store)
	err = apiServer.Run()

	if err != nil {
		log.Fatal(err)
	}

}

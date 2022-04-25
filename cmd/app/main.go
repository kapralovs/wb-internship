package main

import "github.com/kapralovs/wb0/internal/server"

func main() {
	s := server.New()

	if err := s.Run; err != nil {
		panic(err)
	}
}

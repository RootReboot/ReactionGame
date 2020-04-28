package main

import (
	"flag"
	"log"

	grpchighscore "github.com/RootReboot/ReactionGame/m-highscore/internal/server/grpc"
)

func main() {
	var addressPtr = flag.String("address", ":50051", "address where you can connect with m-highscore service")
	flag.Parse()
	s := grpchighscore.NewServer(*addressPtr)

	err := s.ListenAndServe()
	if err != nil {
		log.Fatal("Failed to start grpc server of m-highscore ", err)
	}
}

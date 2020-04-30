package main

import (
	"context"
	"flag"

	"github.com/rs/zerolog/log"

	"google.golang.org/grpc"

	game "github.com/RootReboot/ReactionGame/m-apis"
)

func main() {
	var addressPtr = flag.String("address", "localhost:50051", "address to connect")
	flag.Parse()

	conn, err := grpc.Dial(*addressPtr, grpc.WithInsecure())

	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to get a response")
	}

	defer func() {
		err := conn.Close()
		if err != nil {
			log.Error().Err(err).Str("address", *addressPtr).Msg("Failed to close connection")
		}
	}()

	c := game.NewGameClient(conn)

	if c == nil {
		log.Printf("Client nil")
	}

	r, err := c.GetHighScore(context.Background(), &game.GetHighScoreRequest{})

	if err != nil {
		log.Fatal().Err(err).Str("address", *addressPtr).Msg("failed to get a response")
	}

	if r != nil {
		log.Info().Interface("highscore", r.GetHighScore()).Msg("Highscore from m-highscore microservice")
	} else {
		log.Error().Msg("Could not get highscore")
	}
}

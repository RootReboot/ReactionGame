package grpc

import (
	"context"

	game "github.com/RootReboot/ReactionGame/m-apis"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type Grpc struct {
	address string
	srv     *grpc.Server
}

var HighScore = 999999999.0

func (g *Grpc) SetHighScore(ctx context.Context, input *game.SetHighScoreRequest) (*game.SetHighScoreResponse, error) {
	log.Info().Msg("SetHighScore in m-highscore is called")
	HighScore = input.HighScore
	return &game.SetHighScoreResponse{
		Set: true,
	}, nil
}

func (g *Grpc) GetHighScore(ctx context.Context, input *game.GetHighScoreRequest) (*game.GetHighScoreResponse, error) {
	log.Info().Msg("GetHighScore in m-highscore is called")
	return &game.GetHighScoreResponse{
		HighScore: HighScore,
	}, nil
}

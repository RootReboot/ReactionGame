package grpchighscore

import (
	"context"
	"net"

	game "github.com/RootReboot/ReactionGame/m-apis"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

type Grpc struct {
	address string
	srv     *grpc.Server
}

var HighScore = 999999999.0

func NewServer(address string) *Grpc {
	return &Grpc{
		address: address,
	}
}

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

func (g *Grpc) ListenAndServe() error {
	lis, err := net.Listen("tcp", g.address)
	if err != nil {
		return errors.Wrap(err, "failed to open tct port")
	}
	serverOpts := []grpc.ServerOption{}

	g.srv = grpc.NewServer(serverOpts...)

	game.RegisterGameServer(g.srv, g)

	log.Info().Str("address", g.address).Msg("starting gRPC server for m-highscore microservice")

	err = g.srv.Serve(lis)

	if err != nil {
		return errors.Wrap(err, "failed to start gRPC server for m-highscore microservice")
	}

	return nil
}

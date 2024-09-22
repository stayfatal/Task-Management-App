package main

import (
	"fmt"
	"net/http"
	"os"
	"tma/services/account/config"
	"tma/services/account/internal/repository"
	"tma/services/account/internal/service"
	"tma/services/account/internal/transport"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

func main() {
	err := config.LoadConfig()
	if err != nil {
		log.Fatal().Stack().Err(errors.Wrap(err, "loading config")).Msg("")
	}

	db, err := config.NewDb()
	if err != nil {
		log.Fatal().Stack().Err(errors.Wrap(err, "creating db")).Msg("")
	}

	repo := repository.New(db)

	svc := service.New(repo)

	srv := transport.NewServer(svc)

	log.Info().Msg(fmt.Sprintf("Starting server on port %d", config.Cfg.Port))
	defer log.Info().Msg("Server ended")
	if err := http.ListenAndServe(fmt.Sprintf(":%d", config.Cfg.Port), srv); err != nil {
		log.Fatal().Err(errors.Wrap(err, "starting server")).Msg("")
	}
}

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
}

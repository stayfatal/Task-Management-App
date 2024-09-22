package middleware

import (
	"context"
	"time"
	"tma/services/account/internal/auth"
	"tma/services/account/internal/contextkeys"
	"tma/services/account/internal/utils"

	"github.com/go-kit/kit/endpoint"
	"github.com/rs/zerolog/log"
)

func Authentication() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			token, err := utils.ParseToken(request)
			if err != nil {
				return nil, err
			}

			claims, err := auth.ValidateToken(token)
			if err != nil {
				return nil, err
			}

			ctx = context.WithValue(ctx, contextkeys.UserIdKey{}, claims.Id)

			resp, err := next(ctx, request)
			if err != nil {
				return nil, err
			}

			return resp, err
		}
	}
}

func Logger() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			startTime := time.Now()
			resp, err := next(ctx, request)
			dur := time.Since(startTime)
			if err != nil {
				log.Error().Stack().Err(err).Str("time", dur.String()).Msg("")
			}

			return resp, err
		}
	}
}

func Recoverer() endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			defer func() {
				if r := recover(); r != nil {
					log.Error().Stack().Msgf("Recovered from panic: %v", r)
				}
			}()
			return next(ctx, request)
		}
	}
}

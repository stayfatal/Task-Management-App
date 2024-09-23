package middleware

import "github.com/go-kit/kit/endpoint"

func DefaultChain() endpoint.Middleware {
	return endpoint.Chain(Recoverer(), Logger())
}

func DefaultChainWithAuth() endpoint.Middleware {
	return endpoint.Chain(Recoverer(), Logger(), Authentication())
}

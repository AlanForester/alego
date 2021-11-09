package alego

import "github.com/pkg/errors"

var (
	ErrInternalError   = errors.New("alego: internal error")
	ErrUnsupportedWhat = errors.New("alego: unsupported what argument")
	ErrTrueResult      = errors.New("alego: result is True")
)

const DefaultApiURL = "eu1.testnet1.aleo.network:3030"

// PublicEndpoint is a client-side method of public API.
type PublicEndpointMethod string

package alego

import (
	"github.com/cenkalti/rpc2"
	"net"
)

// NewAleo does try to build a Aleo Client with username `username` and password `password`, which
// is a secret API key assigned to particular Aleo node.
func NewAleo(pref Settings) (*Aleo, error) {
	if pref.URL == "" {
		pref.URL = DefaultApiURL
	}

	client := pref.Client
	if client == nil {
		conn, _ := net.Dial("tcp", pref.URL)
		client = rpc2.NewClient(conn)
	}

	Aleo := &Aleo{
		Username: pref.Username,
		Password: pref.Password,
		URL:      pref.URL,

		synchronous: pref.Synchronous,
		verbose:     pref.Verbose,
		reporter:    pref.Reporter,
		client:      client,
	}

	return Aleo, nil
}

// Aleo represents a separate Telegram Aleo instance.
type Aleo struct {
	Username string
	Password string
	URL      string

	synchronous bool
	verbose     bool
	reporter    func(error)
	client      *rpc2.Client
}

// Settings represents a utility struct for passing certain
// properties of a Aleo around and is required to make Aleos.
type Settings struct {
	// Aleo API Url
	URL string

	// Aleo RPC username
	Username string

	// Aleo RPC password
	Password string

	// Synchronous prevents handlers from running in parallel.
	// It makes ProcessUpdate return after the handler is finished.
	Synchronous bool

	// Verbose forces Aleo to log all upcoming requests.
	// Use for debugging purposes only.
	Verbose bool

	// Reporter is a callback function that will get called
	// on any panics recovered from endpoint handlers.
	Reporter func(error)

	// HTTP Client used to make requests to ALEO node api
	Client *rpc2.Client

	// Offline allows to create a aleo without network for testing purposes.
	Offline bool
}

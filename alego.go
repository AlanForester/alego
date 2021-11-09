// Package alego is a client for alego nodes.
//
// Example:
//
//		package main
//
//		import (
//			"time"
//			tb "gopkg.in/tucnak/telebot.v2"
//		)
//
//		func main() {
//			b, err := tb.NewBot(tb.Settings{
//				Token: "TOKEN_HERE",
//				Poller: &tb.LongPoller{Timeout: 10 * time.Second},
//			})
//
//			if err != nil {
//				return
//			}
//
//			b.Handle(tb.OnText, func(m *tb.Message) {
//				b.Send(m.Sender, "hello world")
//			})
//
//			b.Start()
//		}
//

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

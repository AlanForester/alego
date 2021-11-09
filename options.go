package alego

// Option is a shortcut flag type for certain message features
// (so-called options). It means that instead of passing
// fully-fledged SendOptions* to Send(), you can use these
// flags instead.
//
// Supported options are defined as iota-constants.
type Option int

const (
	Silent Option = iota
)

// SendOptions has most complete control over in what way the message
// must be sent, providing an API-complete set of custom properties
// and options.
//
// Despite its power, SendOptions is rather inconvenient to use all
// the way through bot logic, so you might want to consider storing
// and re-using it somewhere or be using Option flags instead.
type SendOptions struct {
	// The block hash of the requested block
	BlockHash string

	// The block height of the requested block hash
	BlockHeight int

	// The transaction id of the requested transaction hex
	TransactionID string

	// The raw transaction hex to broadcast
	TransactionBytes string
}

func (og *SendOptions) copy() *SendOptions {
	cp := *og
	return &cp
}

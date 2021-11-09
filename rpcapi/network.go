package rpcapi

type NetworkEndpoint struct{}

// Block methods represents a Oleo block.
type Network struct {
	NodeCount             int     `json:"node_count"`
	ConnectionCount       int     `json:"connection_count"`
	Density               float64 `json:"density"`
	AlgebraicConnectivity float64 `json:"algebraic_connectivity"`
	DegreeCentralityDelta float64 `json:"degree_centrality_delta"`
	Edges                 map[int]struct {
		Source string `json:"source"`
		Target string `json:"target"`
	} `json:"edges"`
	Vertices map[int]struct {
		Addr                  string  `json:"addr"`
		IsBootNode            bool    `json:"is_bootnode"`
		DegreeCentrality      uint16  `json:"degree_centrality"`
		EigenvectorCentrality float64 `json:"eigenvector_centrality"`
		FiedlerValue          float64 `json:"fiedler_value"`
	} `json:"vertices"`
}

func (b NetworkEndpoint) ConnectionCount() int {
	return 0
}

func (NetworkEndpoint) Graph() *Network {
	return &Network{}
}

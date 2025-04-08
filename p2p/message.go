package p2p

//Message holds any arbitrary data that is being sent over
// each transport netween two nodes in the network
type RPC struct {
	From    string
	Payload []byte
}

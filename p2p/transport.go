package p2p

// Peer is an inteface that represents the remote node
type Peer interface {
}

// Transport is anything that handles the communication
// between the nodes in the network. This can be of the form
// (TCP, UCP, websockets,...)
type Transport interface {
	ListenAndAccept() error
}

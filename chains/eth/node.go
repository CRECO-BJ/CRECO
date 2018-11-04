package eth

// Node maintains the core 
type Node struct {
	serverConfig p2p.Config
	server       *p2p.Server // Currently running P2P networking layer

	lock sync.RWMutex
	stop chan struct{} // Channel to wait for termination notifications
}
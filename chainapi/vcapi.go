package chainapi

// Chain ...
type Chain interface {
	Node()
}

// Token ...
type Token interface {
	Transfer()
	Balance()
	Approve()
}

// ChainAPI given the common interface opened to the backend core
type ChainAPI interface {
	Token
	Chain
	ValidatingTransaction()
	RecvTransactions()
}

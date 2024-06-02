package p2p

type Transport interface {
	ListenAndAccept() error
}
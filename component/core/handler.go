package core

type Handler interface {
	ServeMsg(*Context)
}

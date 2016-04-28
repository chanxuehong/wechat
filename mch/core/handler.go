package core

type Handler interface {
	ServeMsg(*Context)
}

// HandlerChain --------------------------------------------------------------------------------------------------------

const maxHandlerChainSize = 64

var _ Handler = (HandlerChain)(nil)

type HandlerChain []Handler

// ServeMsg 实现 Handler 接口
func (chain HandlerChain) ServeMsg(ctx *Context) {
	ctx.handlers = chain
	ctx.Next()
}

func (chain *HandlerChain) AppendHandlerFunc(handlers ...func(*Context)) {
	if len(handlers) == 0 {
		return
	}
	for _, h := range handlers {
		if h == nil {
			panic("handler can not be nil")
		}
	}
	handlers2 := make(HandlerChain, len(handlers))
	for i := 0; i < len(handlers); i++ {
		handlers2[i] = HandlerFunc(handlers[i])
	}
	chain.AppendHandler(handlers2)
}

func (chain *HandlerChain) AppendHandler(handlers ...Handler) {
	if len(handlers) == 0 {
		return
	}
	for _, h := range handlers {
		if h == nil {
			panic("handler can not be nil")
		}
	}
	*chain = combineHandlerChain(*chain, handlers)
}

func combineHandlerChain(middlewares, handlers HandlerChain) HandlerChain {
	if len(middlewares)+len(handlers) > maxHandlerChainSize {
		panic("too many handlers")
	}
	return append(middlewares, handlers...)
}

// HandlerFunc ---------------------------------------------------------------------------------------------------------

var _ Handler = HandlerFunc(nil)

type HandlerFunc func(*Context)

// ServeMsg 实现 Handler 接口
func (fn HandlerFunc) ServeMsg(ctx *Context) { fn(ctx) }

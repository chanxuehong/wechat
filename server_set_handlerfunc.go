package wechat

// NOTE: Server 投入使用后就不要再设置了, 不是并发安全的.
func (s *Server) SetInvalidRequestHandler(handler InvalidRequestHandlerFunc) {
	if handler != nil {
		s.invalidRequestHandler = handler
	}
}

// NOTE: Server 投入使用后就不要再设置了, 不是并发安全的.
func (s *Server) SetUnknownRequestHandler(handler UnknownRequestHandlerFunc) {
	if handler != nil {
		s.unknownRequestHandler = handler
	}
}

// NOTE: Server 投入使用后就不要再设置了, 不是并发安全的.
func (s *Server) SetTextRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.textRequestHandler = handler
	}
}

// NOTE: Server 投入使用后就不要再设置了, 不是并发安全的.
func (s *Server) SetImageRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.imageRequestHandler = handler
	}
}

// NOTE: Server 投入使用后就不要再设置了, 不是并发安全的.
func (s *Server) SetVoiceRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.voiceRequestHandler = handler
	}
}

// NOTE: Server 投入使用后就不要再设置了, 不是并发安全的.
func (s *Server) SetVoiceRecognitionRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.voiceRecognitionRequestHandler = handler
	}
}

// NOTE: Server 投入使用后就不要再设置了, 不是并发安全的.
func (s *Server) SetVideoRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.videoRequestHandler = handler
	}
}

// NOTE: Server 投入使用后就不要再设置了, 不是并发安全的.
func (s *Server) SetLocationRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.locationRequestHandler = handler
	}
}

// NOTE: Server 投入使用后就不要再设置了, 不是并发安全的.
func (s *Server) SetLinkRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.linkRequestHandler = handler
	}
}

// NOTE: Server 投入使用后就不要再设置了, 不是并发安全的.
func (s *Server) SetSubscribeEventRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.subscribeEventRequestHandler = handler
	}
}

// NOTE: Server 投入使用后就不要再设置了, 不是并发安全的.
func (s *Server) SetSubscribeEventByScanRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.subscribeEventByScanRequestHandler = handler
	}
}

// NOTE: Server 投入使用后就不要再设置了, 不是并发安全的.
func (s *Server) SetUnsubscribeEventRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.unsubscribeEventRequestHandler = handler
	}
}

// NOTE: Server 投入使用后就不要再设置了, 不是并发安全的.
func (s *Server) SetScanEventRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.scanEventRequestHandler = handler
	}
}

// NOTE: Server 投入使用后就不要再设置了, 不是并发安全的.
func (s *Server) SetLocationEventRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.locationEventRequestHandler = handler
	}
}

// NOTE: Server 投入使用后就不要再设置了, 不是并发安全的.
func (s *Server) SetClickEventRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.clickEventRequestHandler = handler
	}
}

// NOTE: Server 投入使用后就不要再设置了, 不是并发安全的.
func (s *Server) SetViewEventRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.viewEventRequestHandler = handler
	}
}

// NOTE: Server 投入使用后就不要再设置了, 不是并发安全的.
func (s *Server) SetMasssendjobfinishEventRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.masssendjobfinishEventRequestHandler = handler
	}
}

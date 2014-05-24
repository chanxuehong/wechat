package wechat

func (s *Server) SetInvalidRequestHandler(handler InvalidRequestHandlerFunc) {
	if handler != nil {
		s.invalidRequestHandler = handler
	}
}

func (s *Server) SetUnknownRequestHandler(handler UnknownRequestHandlerFunc) {
	if handler != nil {
		s.unknownRequestHandler = handler
	}
}

func (s *Server) SetTextRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.textRequestHandler = handler
	}
}

func (s *Server) SetImageRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.imageRequestHandler = handler
	}
}

func (s *Server) SetVoiceRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.voiceRequestHandler = handler
	}
}

func (s *Server) SetVoiceRecognitionRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.voiceRecognitionRequestHandler = handler
	}
}

func (s *Server) SetVideoRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.videoRequestHandler = handler
	}
}

func (s *Server) SetLocationRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.locationRequestHandler = handler
	}
}

func (s *Server) SetLinkRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.linkRequestHandler = handler
	}
}

func (s *Server) SetSubscribeEventRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.subscribeEventRequestHandler = handler
	}
}

func (s *Server) SetSubscribeEventByScanRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.subscribeEventByScanRequestHandler = handler
	}
}

func (s *Server) SetUnsubscribeEventRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.unsubscribeEventRequestHandler = handler
	}
}

func (s *Server) SetScanEventRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.scanEventRequestHandler = handler
	}
}

func (s *Server) SetLocationEventRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.locationEventRequestHandler = handler
	}
}

func (s *Server) SetClickEventRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.clickEventRequestHandler = handler
	}
}

func (s *Server) SetViewEventRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.viewEventRequestHandler = handler
	}
}

func (s *Server) SetMasssendjobfinishEventRequestHandler(handler RequestHandlerFunc) {
	if handler != nil {
		s.masssendjobfinishEventRequestHandler = handler
	}
}

package server
func (s *EchoServer) registerRoutes() {
	categoryRoutes(s)
}

func categoryRoutes(s *EchoServer) {
	categoryGroup := s.echo.Group("/category")
	categoryGroup.POST("", s.AddCategory)
}
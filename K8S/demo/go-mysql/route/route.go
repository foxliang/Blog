package route

func (s Server) registeRoute() {

	// data
	data := s.engine.Group("/data")
	data.Use()
	{
		data.GET("/get", s.get)
		data.GET("/list", s.list)
		data.GET("/create", s.create)
		data.GET("/update", s.update)
		data.GET("/del", s.del)
	}

}

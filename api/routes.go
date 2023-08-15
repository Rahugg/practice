package api

func (s *Server) RegisterRoutes() {
	s.Gin.GET("/people/:name", s.HandleGetByName)
	s.Gin.POST("/people", s.HandleAddPerson)
	s.Gin.GET("/people", s.HandleGetNames)
	s.Gin.GET("/person", s.HandleSearchPersonByName)
	s.Gin.GET("/person_by_surname", s.HandleSearchPersonBySurname)
	s.Gin.GET("/person_by_address", s.HandleSearchPersonByAddress)
	s.Gin.GET("/sorted_people_asc", s.SortedPeopleASC)
	s.Gin.GET("/sorted_people_desc", s.SortedPeopleDESC)
	s.Gin.GET("/filter", s.HandleFilterByAge)
}

package internal

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	repo Repo
}

func (s *Server) Run() error {
	var err error
	s.repo, err = newSqliteRepo()

	if err != nil {
		return err
	}

	router := gin.Default()

	router.POST("/new", s.CreateNewShortUrl)
	router.GET("/:id", s.RedirectTarget)

	return router.Run()
}

func (s *Server) RedirectTarget(c *gin.Context) {
	id := c.Param("id")

	shorturl, err := s.repo.Get(id)
	if err != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.Redirect(http.StatusMovedPermanently, shorturl.Target)
}

func (s *Server) CreateNewShortUrl(c *gin.Context) {
	var request struct {
		Target string `json:"target"`
	}

	if err := c.Bind(&request); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	shorturl, err := s.repo.Create(request.Target)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, &shorturl)
}

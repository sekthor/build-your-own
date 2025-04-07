package internal

import (
	"fmt"
	"net/http"
	"net/url"

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
	router.LoadHTMLGlob("../templates/*")

	router.POST("/new", s.CreateNewShortUrl)
	router.GET("/:id", s.RedirectTarget)
	router.GET("/", s.Home)

	return router.Run()
}

func (s *Server) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{})
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
		Target string `json:"target" form:"target"`
	}

	if err := c.Bind(&request); err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	// validate url
	_, err := url.ParseRequestURI(request.Target)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	shorturl, err := s.repo.Create(request.Target)
	if err != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("http://%s/%s", c.Request.Host, shorturl.ID)

	c.HTML(http.StatusOK, "created.html", gin.H{
		"Url": url,
	})
}

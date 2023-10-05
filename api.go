package main

import "github.com/gin-gonic/gin"

type APIServer struct {
	listAddr string
	store    Store
}

func NewAPIServer(listAddr string, store Store) *APIServer {
	return &APIServer{
		listAddr: listAddr,
		store:    store,
	}
}

func (api *APIServer) handleHomePage(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func (api *APIServer) handleCreateURL(c *gin.Context) {
	url := c.PostForm("url")
	var hash string

	if c.PostForm("hash") != "" {
		hash = c.PostForm("hash")
	} else {
		hash = RanHash()
	}

	err := api.store.InsertURL(CreateURLRequest{
		url:  url,
		hash: hash,
	})
	if err != nil {
		c.String(500, "Error creating URL")
		return
	}

	c.String(200, "<a href='%s'>/%s</a>", hash, hash)
}

func (api *APIServer) handleRedirect(c *gin.Context) {
	url, err := api.store.GetURLByHash(c.Param("hash"))
	if err != nil {
		c.String(500, "Error retrieving URL")
		return
	}

	c.Redirect(302, url)
}

func (api *APIServer) Start() error {
	r := gin.Default()

	// disable if behind a proxy or in production
	r.SetTrustedProxies(nil)

	r.Static("/assets", "./static")
	r.LoadHTMLGlob("templates/*")

	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/", api.handleHomePage)
	r.POST("/create", api.handleCreateURL)
	r.GET("/:hash", api.handleRedirect)

	err := r.Run(api.listAddr)
	if err != nil {
		return err
	}

	return nil
}

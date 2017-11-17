package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
	"net/http"
	"strings"
)

type Server struct {
	clientset *kubernetes.Clientset
	builder   *Builder
	deployer  *Deployer
	engine    *gin.Engine
}

func NewServer(clientset *kubernetes.Clientset, builder *Builder, deployer *Deployer) *Server {
	engine := gin.Default()
	// Set lower memory limit for uploads
	engine.MaxMultipartMemory = 8 << 20 // 8 MiB

	return &Server{
		clientset: clientset,
		builder:   builder,
		deployer:  deployer,
		engine:    engine,
	}
}

func (s *Server) Start() error {
	s.setRoutes()

	return s.engine.Run(":8080")
}

func (s *Server) setRoutes() {
	s.engine.GET("/ping", s.pingRoute)
	s.engine.POST("/deploy", s.deployRoute)
}

func (s *Server) pingRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func (s *Server) deployRoute(c *gin.Context) {
	// TODO: don't use filename, but a random name (avoid clashing)
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("error with file: %s", err.Error()),
		})
		return
	}

	project := strings.Replace(file.Filename, ".tar.gz", "", 1)
	filePath := fmt.Sprintf("/public/%s", project)

	if err := c.SaveUploadedFile(file, filePath); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("error uploading file: %s", err.Error()),
		})
		return
	}

	if err := s.builder.BuildFromFile(project, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error on build: %s", err.Error()),
		})
		return
	}

	if err := s.deployer.Deploy(project); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error on deploy: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Deployment finished",
		"project": project,
	})
}

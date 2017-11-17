package main

import (
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
)

type Server struct {
	clientset *kubernetes.Clientset
	engine    *gin.Engine
}

func NewServer(clientset *kubernetes.Clientset) *Server {
	return &Server{
		clientset: clientset,
		engine:    gin.Default(),
	}
}

func (s *Server) Start() error {
	s.setRoutes()

	return s.engine.Run()
}

func (s *Server) setRoutes() {
	s.engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
}

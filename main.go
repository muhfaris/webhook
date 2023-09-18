package main

import (
	"os"

	"github.com/muhfaris/webhook/cmd"
	"github.com/sirupsen/logrus"
)

func init() {
	// Initialize Logrus
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
}

func main() {
	cmd.Run()
}

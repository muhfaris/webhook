package cmd

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/webhook/configs"
	"github.com/muhfaris/webhook/handlers"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var restCommand = &cobra.Command{
	Use: "api",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		err := configs.ReadConfig(cfgFile)
		if err != nil {
			logrus.Fatalf("rest config: %v", err)
		}

		return err
	},
	Run: func(cmd *cobra.Command, args []string) {
		app := fiber.New(fiber.Config{AppName: "wekbook"})
		app.Post("/webhook", handlers.HandleWebhook)
		app.Post("/hash", handlers.HandleHash)

		go func() {
			log.Println("webhook server listening on :8080")
			if err := app.Listen(fmt.Sprintf(":%d", configs.App.Port)); err != nil {
				log.Fatalf("server error: %v", err)
			}
		}()

		// Set up a signal channel to listen for SIGINT and SIGTERM signals
		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

		// Block until a signal is received
		<-signalChan
		logrus.Println("shutting down the server...")

		// Shutdown the server gracefully
		if err := app.Shutdown(); err != nil {
			logrus.Fatalf("server shutdown error: %v", err)
		}

		logrus.Println("server shutdown complete")
	},
}

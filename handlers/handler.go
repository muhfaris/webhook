package handlers

import (
	"fmt"
	"strings"
	"sync"

	"github.com/Jeffail/gabs"
	"github.com/gofiber/fiber/v2"
	"github.com/muhfaris/webhook/configs"
	"github.com/muhfaris/webhook/pkg"
)

type WebhookPayload struct {
	ID      string      `json:"id"`
	Token   string      `json:"token"`
	Payload interface{} `json:"data"`
}

var configMutex sync.Mutex

func HandleWebhook(c *fiber.Ctx) error {
	if c.Method() != fiber.MethodPost {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("method not allowed")
	}

	configMutex.Lock()
	defer configMutex.Unlock()

	var payload WebhookPayload
	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"message": "failed to parse JSON"})
	}

	webhook, err := configs.App.Webhooks.ByID(payload.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(map[string]interface{}{"message": "webhook id not found"})
	}

	hashToken := pkg.HashSha1(webhook.Token)
	if payload.Token != hashToken {
		return c.Status(fiber.StatusUnauthorized).JSON(map[string]interface{}{"message": "invalid token"})
	}

	body := c.Body()
	jsonParsed, err := gabs.ParseJSON(body)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "error executing command",
			"error":   err.Error(),
		})
	}

	var args []string
	for index, argument := range webhook.CommandArguments {
		// add space
		if index == 0 {
			argument.Name = fmt.Sprintf(" %s", argument.Name)
		}

		value := jsonParsed.Path(argument.Source).Data()
		if value == nil {
			args = append(args, argument.Name)
			continue
		}
		args = append(args, argument.Name, fmt.Sprintf("%v", value))
	}

	argAll := strings.Join(args, " ")
	command := webhook.ExecuteCommand + argAll

	tmpCommand := strings.TrimSpace(command)
	if command == "" || tmpCommand == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": fmt.Sprintf("the id %s don't have command execution", payload.ID),
		})
	}

	if err := pkg.ExecuteCommand(command, webhook.Workdir); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(map[string]interface{}{
			"message": "error executing command",
			"error":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{"message": "webhook received and command executed successfully"})
}

func HandleHash(c *fiber.Ctx) error {
	if c.Method() != fiber.MethodPost {
		return c.Status(fiber.StatusMethodNotAllowed).SendString("method not allowed")
	}

	var payload = struct {
		Message string `json:"message"`
	}{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(map[string]interface{}{"message": "failed to parse JSON"})
	}

	hash := pkg.HashSha1(payload.Message)
	return c.Status(fiber.StatusOK).JSON(map[string]interface{}{"message": hash})
}

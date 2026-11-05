package middleware

import (
	"github.com/gofiber/fiber/v3"
)

func KunciApi(kunci string) fiber.Handler {
	return func(c fiber.Ctx) error {
		if c.Get("X-Api-Key") != kunci {
			return fiber.NewError(fiber.StatusUnauthorized, "Kunci API tidak valid")
		}

		return c.Next()
	}
}

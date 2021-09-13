package handlers

import (
	"github.com/erikyvanov/pruebas-demos/services"
	"github.com/gofiber/fiber/v2"
)

func GetCache(ctx *fiber.Ctx) error {
	cacheService := services.GetCacheService()

	return ctx.JSON(cacheService.GetCache())
}

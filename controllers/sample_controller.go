package controllers

import (
	"demo-gofiber/services"
	"github.com/gofiber/fiber/v2"
)

type SampleController struct {
	service services.SampleServicer
}

func NewSampleController(s services.SampleServicer) *SampleController {
	return &SampleController{service: s}
}

func (c *SampleController) SampleFunc(ctx *fiber.Ctx) error {
	result := c.service.SampleFunc(ctx.UserContext())
	return ctx.Status(200).JSON(result)
}

func (c SampleController) RenderHtml(ctx *fiber.Ctx) error {
	return ctx.Render("index", fiber.Map{
		"Title": "Hello, World!",
	})
}

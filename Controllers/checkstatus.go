package controllers

import "github.com/gofiber/fiber"

func CheckStatus(ctx *fiber.Ctx) {	
    ctx.Send("uwu")
}
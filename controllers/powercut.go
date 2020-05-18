package controllers

import "github.com/gofiber/fiber"

// GetDefault get default powercut data
func GetDefault(ctx *fiber.Ctx) {
	ctx.Send("Default PowerCut Data")
}

// GetByDivision get powercuts by division name
func GetByDivision(ctx *fiber.Ctx) {
	ctx.Send("Powercuts by Division")
}

// GetByDivisionAndSubDivision get powercuts by division and sub-division names
func GetByDivisionAndSubDivision(ctx *fiber.Ctx) {
	ctx.Send("Powercuts by division and subdivision")
}

// GetByYear get powercuts by year number (YYYY)
func GetByYear(ctx *fiber.Ctx) {
	ctx.Send("Power cuts by Year")
}

// GetByMonth get powercuts by month name
func GetByMonth(ctx *fiber.Ctx) {
	ctx.Send("Power cuts by Year")
}

// GetByYearandMonth get powercuts by year number (YYYY) and month name
func GetByYearandMonth(ctx *fiber.Ctx) {
	ctx.Send("Power cuts by Year")
}

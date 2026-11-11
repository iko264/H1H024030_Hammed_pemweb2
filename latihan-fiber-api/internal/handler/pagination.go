package handler

import (
	"github.com/gofiber/fiber/v3"
)

func normalisasiHalaman(halaman, perHalaman int) (int, int) {
	if halaman < 1 {
		halaman = 1
	}
	if perHalaman < 1 || perHalaman > 100 {
		perHalaman = 10
	}

	return halaman, perHalaman
}

func metaHalaman(halaman, perHalaman int, total int64) fiber.Map {
	totalHalaman := int(total) / perHalaman
	if int(total)%perHalaman != 0 {
		totalHalaman = totalHalaman + 1
	}

	return fiber.Map{
		"halaman":       halaman,
		"per_halaman":   perHalaman,
		"total":         total,
		"total_halaman": totalHalaman,
	}
}

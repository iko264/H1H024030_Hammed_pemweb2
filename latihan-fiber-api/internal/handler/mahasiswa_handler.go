package handler

import (
	"sort"
	"strconv"

	"github.com/gofiber/fiber/v3"

	"latihan-fiber-api/internal/model"
)

func DaftarMahasiswa(c fiber.Ctx) error {
	angkatan := c.Query("angkatan", "")
	urut := c.Query("urut", "")

	hasil := make([]model.Mahasiswa, 0, len(model.Daftar))
	if angkatan == "" {
		hasil = append(hasil, model.Daftar...)
	} else {
		nilaiAngkatan, err := strconv.Atoi(angkatan)
		if err != nil {
			return fiber.NewError(fiber.StatusBadRequest, "Parameter angkatan harus berupa angka")
		}

		for _, mahasiswa := range model.Daftar {
			if mahasiswa.Angkatan == nilaiAngkatan {
				hasil = append(hasil, mahasiswa)
			}
		}
	}

	switch urut {
	case "":
	case "nama":
		sort.Slice(hasil, func(i, j int) bool {
			return hasil[i].Nama < hasil[j].Nama
		})
	case "ipk":
		sort.Slice(hasil, func(i, j int) bool {
			return hasil[i].Ipk > hasil[j].Ipk
		})
	default:
		return fiber.NewError(fiber.StatusBadRequest, "Parameter urut harus bernilai nama atau ipk")
	}

	return c.JSON(fiber.Map{
		"sukses": true,
		"data":   hasil,
	})
}

func DetailMahasiswa(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "ID harus berupa angka")
	}

	for i := 0; i < len(model.Daftar); i++ {
		if model.Daftar[i].ID == id {
			return c.JSON(fiber.Map{
				"sukses": true,
				"data":   model.Daftar[i],
			})
		}
	}

	return fiber.NewError(fiber.StatusNotFound, "Mahasiswa tidak ditemukan")
}

func CariMahasiswaByNIM(c fiber.Ctx) error {
	nim := c.Params("nim")

	for i := 0; i < len(model.Daftar); i++ {
		if model.Daftar[i].Nim == nim {
			return c.JSON(fiber.Map{
				"sukses": true,
				"data":   model.Daftar[i],
			})
		}
	}

	return fiber.NewError(fiber.StatusNotFound, "Mahasiswa tidak ditemukan")
}

package auth

import (
	"errors"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"

	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/model"
)

type Klaim struct {
	Peran string `json:"peran"`
	jwt.RegisteredClaims
}

func TerbitkanToken(idPengguna uint, peran string, rahasia string, ttlMenit int) (string, error) {
	sekarang := time.Now()
	klaim := Klaim{
		Peran: peran,
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.FormatUint(uint64(idPengguna), 10),
			IssuedAt:  jwt.NewNumericDate(sekarang),
			ExpiresAt: jwt.NewNumericDate(sekarang.Add(time.Duration(ttlMenit) * time.Minute)),
			Issuer:    "pemweb2-fiber-api",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, klaim)
	return token.SignedString([]byte(rahasia))
}

func PeriksaToken(nilaiToken string, rahasia string) (*Klaim, error) {
	token, err := jwt.ParseWithClaims(nilaiToken, &Klaim{}, func(t *jwt.Token) (any, error) {
		if _, cocok := t.Method.(*jwt.SigningMethodHMAC); cocok == false {
			return nil, errors.New("metode penandatanganan tidak didukung")
		}
		return []byte(rahasia), nil
	})

	if err != nil {
		return nil, err
	}

	klaim, cocok := token.Claims.(*Klaim)
	if cocok == false || token.Valid == false {
		return nil, errors.New("token tidak valid")
	}

	if DalamDaftarHitam(nilaiToken) {
		return nil, errors.New("token sudah dicabut")
	}

	return klaim, nil
}

func DalamDaftarHitam(nilaiToken string) bool {
	var jumlah int64
	config.DB.Model(&model.TokenDaftarHitam{}).
		Where("token = ?", nilaiToken).
		Count(&jumlah)
	return jumlah > 0
}

func CabutToken(nilaiToken string, kedaluwarsa time.Time) error {
	return config.DB.Create(&model.TokenDaftarHitam{
		Token:       nilaiToken,
		Kedaluwarsa: kedaluwarsa,
	}).Error
}

func BersihkanDaftarHitam() error {
	return config.DB.Where("kedaluwarsa < ?", time.Now()).
		Delete(&model.TokenDaftarHitam{}).Error
}
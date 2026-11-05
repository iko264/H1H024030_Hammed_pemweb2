package main

import (
	"errors"
	"fmt"
)

type mahasiswa struct {
	Nama  string
	Nim   string
	Angkatan int
	Ipk   float64
}

func cariMahasiswa(daftar []mahasiswa, nim string) (mahasiswa, error) {
	for i := 0; i < len(daftar); i++ {
		if daftar[i].Nim == nim {
			return daftar[i], nil
		}	
	}
	return mahasiswa{}, errors.New("mahasiswa tidak ditemukan")
}

func rerataIpk(daftar []mahasiswa) float64 {
	if len(daftar) == 0 {
		return 0
	}
	var total float64
	for i := 0; i < len(daftar); i++ {
		total += daftar[i].Ipk
	}
	return total / float64(len(daftar))
}

func main() {
	daftar := []mahasiswa{
		{Nim: "H1A123001", Nama: "Andi Prasetyo", Angkatan: 2023, Ipk: 3.45},
		{Nim: "H1A123002", Nama: "Bunga Lestari", Angkatan: 2023, Ipk: 3.80},
		{Nim: "H1A124003", Nama: "Citra Ramadhani", Angkatan: 2024, Ipk: 3.20},
	}
	ditemukan, err := cariMahasiswa(daftar, "H1A123002")
	if err != nil {
		fmt.Println("galat:", err)
		return
	}	

	fmt.Println("Mahasiswa ditemukan:", ditemukan.Nama)
	fmt.Println("dengan IPK:", ditemukan.Ipk)
	jumlahPerAngkatan := map[int]int{}

	for i := 0; i < len(daftar); i++ {
		jumlahPerAngkatan[daftar[i].Angkatan] =
		jumlahPerAngkatan[daftar[i].Angkatan] + 1
	}
	fmt.Println("Jumlah mahasiswa per angkatan:", jumlahPerAngkatan)
}
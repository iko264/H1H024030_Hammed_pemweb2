package repository

import (
	"fmt"

	"latihan-fiber-api/internal/config"
	"latihan-fiber-api/internal/model"
)

func AmbilProgramStudi() ([]model.ProgramStudiDenganJumlah, error) {
	var daftar []model.ProgramStudiDenganJumlah

	tabelProgramStudi := config.DB.NamingStrategy.TableName("ProgramStudi")
	tabelMahasiswa := config.DB.NamingStrategy.TableName("Mahasiswa")

	err := config.DB.Model(&model.ProgramStudi{}).
		Select(fmt.Sprintf("%s.*, COUNT(%s.id) AS jumlah_mahasiswa", tabelProgramStudi, tabelMahasiswa)).
		Joins(fmt.Sprintf(
			"LEFT JOIN %s ON %s.program_studi_id = %s.id AND %s.deleted_at IS NULL",
			tabelMahasiswa, tabelMahasiswa, tabelProgramStudi, tabelMahasiswa,
		)).
		Group(tabelProgramStudi + ".id").
		Order(tabelProgramStudi + ".nama ASC").
		Scan(&daftar).Error

	return daftar, err
}

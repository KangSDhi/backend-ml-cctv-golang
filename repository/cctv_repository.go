package repository

import (
	"backend-ml-cctv-golang/config"
	"backend-ml-cctv-golang/entity"
)

func SaveCCTVData(cctv entity.CCTV) (entity.CCTV, error) {
	err := config.DB.Create(&cctv).Error
	if err != nil {
		return entity.CCTV{}, err
	}
	return cctv, err
}

func GetLatestCCTVData() ([]entity.CCTV, error) {
	var cctvs []entity.CCTV

	subQuery := config.DB.Table("cctvs").
		Select("nama_cctv, MAX(created_at) AS latest").
		Group("nama_cctv")

	err := config.DB.Table("cctvs AS c").
		Joins("JOIN (?) AS recent ON c.nama_cctv = recent.nama_cctv AND c.created_at = recent.latest", subQuery).
		Find(&cctvs).Error

	if err != nil {
		return nil, err
	}

	return cctvs, err
}

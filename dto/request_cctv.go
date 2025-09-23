package dto

type RequestCCTV struct {
	NamaCCTV string `json:"nama_cctv" validate:"required"`
	Objek    *uint  `json:"objek" validate:"required"`
}

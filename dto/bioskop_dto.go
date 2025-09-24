package dto

type CreateRequest struct {
	Name   string  `json:"name" validate:"required,min=1"`
	Lokasi string  `json:"lokasi" validate:"required,min=4"`
	Rating float64 `json:"rating"`
}

type BioskopResponse struct {
	ID     uint    `json:"id"`
	Name   string  `json:"name"`
	Lokasi string  `json:"lokasi"`
	Rating float64 `json:"rating"`
}

type GetResponse struct {
	Data []*BioskopResponse `json:"data"`
}

type UpdateRequest struct {
	Name   string  `json:"name,omitempty" validate:"omitempty,min=1"`
	Lokasi string  `json:"lokasi,omitempty" validate:"omitempty,min=4"`
	Rating float64 `json:"rating,omitempty" validate:"omitempty"`
}

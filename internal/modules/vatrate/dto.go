package vatrate

type CreateVatRateRequest struct {
	Rate     float64 `json:"rate" validate:"required,min=0,max=100"`
	Year     int     `json:"year" validate:"required,min=2000"`
	Month    int     `json:"month" validate:"required,min=1,max=12"`
	IsActive bool    `json:"is_active"`
}

type UpdateVatRateRequest struct {
	Rate     float64 `json:"rate" validate:"required,min=0,max=100"`
	Year     int     `json:"year" validate:"required,min=2000"`
	Month    int     `json:"month" validate:"required,min=1,max=12"`
	IsActive bool    `json:"is_active"`
}

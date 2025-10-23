package vatrate

type CreateVatRateRequest struct {
	Name     string  `json:"name" validate:"required,min=3,max=8"`
	Rate     float32 `json:"rate" validate:"required,min=0,max=100"`
	Year     int     `json:"year" validate:"required,min=2000"`
	Month    int8    `json:"month" validate:"required,min=1,max=12"`
	IsActive bool    `json:"is_active"`
}

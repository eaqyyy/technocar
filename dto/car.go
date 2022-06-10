package dto

type Car struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Brand string `json:"brand"`
	IsElectric bool `json:"isElectric"`
	HorsePower int64 `json:"horsePower"`
	BasePrice float64 `json:"basePrice"` // for currency use BigRat
	StorePrice float64 `json:"StorePrice"` 
}

// receiver function -> function yg nempel sama obyek
func (c *Car) CalculatePrice() {
	if c.IsElectric {
		c.StorePrice = c.BasePrice + (float64(c.HorsePower) / 4)
	} else {
		c.StorePrice = c.BasePrice + (float64(c.HorsePower) / 2) + 100
	}
}
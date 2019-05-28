package products

type Product struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Qunatity int    `json:"qty,omitempty"`
}

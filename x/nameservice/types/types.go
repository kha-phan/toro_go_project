type Product struct {
	ProductID   string         `json:"productID"`
	Description string         `json:"description"`
	Owner       sdk.AccAddress `json:"owner"` // kiểu địa chỉ được định nghĩa trong cosmos sdk ví dụ: cosmos1ug35j0s0mfn6hah5sk076yfjqwxlh4gtfvdfpa
	Price       sdk.Coins      `json:"price"` // kiểu tiền tệ được định nghĩa bởi cosmos sdk ví dụ: [{"denom": "nametoken","amount":"10"}]
}

func NewProduct() Product {
	return Product{}
}
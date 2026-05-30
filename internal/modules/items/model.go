package items

type Item struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

var Items = []Item{
	{
		ID:          1,
		Name:        "Item exemplo",
		Description: "Primeiro item do starter",
	},
}

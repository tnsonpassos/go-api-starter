package items

import "strings"

func ValidateItem(item Item) string {
	if strings.TrimSpace(item.Name) == "" {
		return "O campo name é obrigatório"
	}

	return ""
}

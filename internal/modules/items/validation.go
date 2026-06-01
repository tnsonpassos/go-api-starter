package items

import "strings"

func ValidateItem(item Item) map[string]string {
	errors := make(map[string]string)

	if strings.TrimSpace(item.Name) == "" {
		errors["name"] = "O campo name é obrigatório"
	}

	return errors
}

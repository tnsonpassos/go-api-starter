package items

import "errors"

func FindAll() []Item {
	return Items
}

func FindByID(id int) (*Item, error) {
	for _, item := range Items {
		if item.ID == id {
			return &item, nil
		}
	}

	return nil, errors.New("item não encontrado")
}

func Create(item Item) Item {
	item.ID = len(Items) + 1
	Items = append(Items, item)

	return item
}

func Update(id int, updatedItem Item) (*Item, error) {
	for index, item := range Items {
		if item.ID == id {
			updatedItem.ID = id
			Items[index] = updatedItem

			return &updatedItem, nil
		}
	}

	return nil, errors.New("item não encontrado")
}

func Delete(id int) error {
	for index, item := range Items {
		if item.ID == id {
			Items = append(Items[:index], Items[index+1:]...)
			return nil
		}
	}

	return errors.New("item não encontrado")
}

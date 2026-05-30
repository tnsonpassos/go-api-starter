package items

func ListItems() []Item {
	return FindAll()
}

func GetItem(id int) (*Item, error) {
	return FindByID(id)
}

func CreateItem(item Item) Item {
	return Create(item)
}

func UpdateItem(id int, item Item) (*Item, error) {
	return Update(id, item)
}

func DeleteItem(id int) error {
	return Delete(id)
}

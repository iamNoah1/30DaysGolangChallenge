package internal

type TogoItem struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func New(id int, description string) *TogoItem {
	return &TogoItem{
		ID:          id,
		Description: description,
		Completed:   false,
	}
}

var store = make(map[int]*TogoItem)

func Add(item *TogoItem) {
	store[item.ID] = item
}

func Get(id int) *TogoItem {
	return store[id]
}

package structure

type (
	CollectionAware interface {
		GetId() string
		Append(...interface{}) CollectionAware
		GetItems() []interface{}
		GetLength() uint64
	}

	Collection struct {
		id     string
		items  []interface{}
		length uint64
		CollectionAware
	}
)

func NewCollection(id string) *Collection {
	return &Collection{
		id: id,
	}
}

func (self *Collection) GetId() string {
	return self.id
}

func (self *Collection) Append(items ...interface{}) *Collection {
	for _, item := range items {
		self.items = append(self.items, item)
		self.length = uint64(len(self.items))
	}

	return self
}

func (self *Collection) GetItems() []interface{} {
	return self.items
}

func (self *Collection) GetLength() uint64 {
	return self.length
}

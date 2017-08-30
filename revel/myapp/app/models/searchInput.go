package models

type SearchInput struct {
	Name string
}

func (input *SearchInput) SetInput(name string) {
	input.Name = name
}

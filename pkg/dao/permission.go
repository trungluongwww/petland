package dao

type Permission interface {
}

type permission struct {
}

func NewPermission() Permission {
	return &permission{}
}

package common

type Page[T any] struct {
	Data  []T `json:"data"`
	Total int `json:"total"`
	Page  int `json:"page"`
	Size  int `json:"size"`
}

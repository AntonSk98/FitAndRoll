package common

// Page represents a paginated response containing a slice of data of any type.
// T is a generic type parameter representing the data type.
type Page[T any] struct {
	Data  []T `json:"data"`
	Total int `json:"total"`
	Page  int `json:"page"`
	Size  int `json:"size"`
}

// PageParams represents pagination parameters for requests.
type PageParams struct {
	Page int `json:"currentPage"`
	Size int `json:"pageSize"`
}

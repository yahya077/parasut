package parasut

type HasPagination struct {
	SortField  string `url:"sort,omitempty"`
	PageNumber int    `url:"page[number],omitempty"`
	PageSize   int    `url:"page[size],omitempty"`
}

type HasInclude struct {
	Include string `url:"include,omitempty"`
}

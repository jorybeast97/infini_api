package domain

type PageQuery struct {
    Page  int
    Limit int
}

type SortQuery struct {
    Field string
    Dir   string
}

package domain

type Location struct {
    Lat  float64 `json:"lat"`
    Lng  float64 `json:"lng"`
    Name string  `json:"name"`
}

type AppProject struct {
    ID          string   `json:"id"`
    Name        string   `json:"name"`
    Description string   `json:"description"`
    Icon        string   `json:"icon"`
    URL         string   `json:"url"`
    Tags        []string `json:"tags"`
}

type BlogPost struct {
    ID       string    `json:"id"`
    Title    string    `json:"title"`
    Excerpt  string    `json:"excerpt"`
    Content  string    `json:"content"`
    Date     string    `json:"date"`
    Location *Location `json:"location"`
    ReadTime string    `json:"readTime"`
    Status   string    `json:"status,omitempty"`
    Partners []string  `json:"partners,omitempty"`
}

type Photo struct {
    ID       string   `json:"id"`
    URL      string   `json:"url"`
    Caption  string   `json:"caption"`
    Location Location `json:"location"`
    Date     string   `json:"date"`
}

type Social struct {
    Github   *string `json:"github,omitempty"`
    Twitter  *string `json:"twitter,omitempty"`
    Linkedin *string `json:"linkedin,omitempty"`
}

type Author struct {
    ID     string `json:"id"`
    Name   string `json:"name"`
    Role   string `json:"role"`
    Avatar string `json:"avatar"`
    Bio    string `json:"bio"`
    Social Social `json:"social"`
}

type Meta struct {
    Page  int `json:"page"`
    Limit int `json:"limit"`
    Total int `json:"total"`
}

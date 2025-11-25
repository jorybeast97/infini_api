package domain

type Location struct {
	Lat  float64 `json:"lat"`
	Lng  float64 `json:"lng"`
	Name string  `json:"name"`
}

type AppProject struct {
	ID             string      `json:"id" gorm:"primaryKey;size:64"`
	Name           string      `json:"name" gorm:"size:120;not null"`
	Description    string      `json:"description" gorm:"size:400"`
	Icon           string      `json:"icon" gorm:"size:40"`
	URL            string      `json:"url" gorm:"size:255"`
    Tags           StringArray `json:"tags" gorm:"type:jsonb"`
	CreateByUserID string      `json:"createByUserId" gorm:"size:64;index"`
	UpdateByUserID string      `json:"updateByUserId" gorm:"size:64;index"`
	CreatedAt      int64       `json:"createdAt" gorm:"type:bigint"`
	UpdatedAt      int64       `json:"updatedAt" gorm:"type:bigint"`
}

type BlogPost struct {
	ID             string      `json:"id" gorm:"primaryKey;size:64"`
	Title          string      `json:"title" gorm:"size:120;not null"`
	Excerpt        string      `json:"excerpt" gorm:"size:280"`
    Content        string      `json:"content" gorm:"type:text"`
	Date           int64       `json:"date" gorm:"type:bigint"`
    Location       *Location   `json:"location" gorm:"type:jsonb"`
	ReadTime       string      `json:"readTime" gorm:"size:40"`
	Status         string      `json:"status,omitempty" gorm:"size:20"`
    Partners       StringArray `json:"partners,omitempty" gorm:"type:jsonb"`
	CreateByUserID string      `json:"createByUserId" gorm:"size:64;index"`
	UpdateByUserID string      `json:"updateByUserId" gorm:"size:64;index"`
	CreatedAt      int64       `json:"createdAt" gorm:"type:bigint"`
	UpdatedAt      int64       `json:"updatedAt" gorm:"type:bigint"`
}

type Photo struct {
	ID             string   `json:"id" gorm:"primaryKey;size:64"`
	URL            string   `json:"url" gorm:"size:255"`
	Caption        string   `json:"caption" gorm:"size:140"`
    Location       Location `json:"location" gorm:"type:jsonb"`
	Date           int64    `json:"date" gorm:"type:bigint"`
	CreateByUserID string   `json:"createByUserId" gorm:"size:64;index"`
	UpdateByUserID string   `json:"updateByUserId" gorm:"size:64;index"`
	CreatedAt      int64    `json:"createdAt" gorm:"type:bigint"`
	UpdatedAt      int64    `json:"updatedAt" gorm:"type:bigint"`
}

type Social struct {
	Github   *string `json:"github,omitempty"`
	Twitter  *string `json:"twitter,omitempty"`
	Linkedin *string `json:"linkedin,omitempty"`
}

type Author struct {
	ID             string `json:"id" gorm:"primaryKey;size:64"`
	Name           string `json:"name" gorm:"size:80;not null"`
	Role           string `json:"role" gorm:"size:60;not null"`
	Avatar         string `json:"avatar" gorm:"size:255"`
	Bio            string `json:"bio" gorm:"size:400"`
    Social         Social `json:"social" gorm:"type:jsonb"`
	CreateByUserID string `json:"createByUserId" gorm:"size:64;index"`
	UpdateByUserID string `json:"updateByUserId" gorm:"size:64;index"`
	CreatedAt      int64  `json:"createdAt" gorm:"type:bigint"`
	UpdatedAt      int64  `json:"updatedAt" gorm:"type:bigint"`
}

type Meta struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
	Total int `json:"total"`
}

type User struct {
	ID           string `json:"id" gorm:"primaryKey;size:64"`
	UserName     string `json:"userName" gorm:"size:80;uniqueIndex;not null"`
	NickName     string `json:"nickName" gorm:"size:120"`
	PasswordHash string `json:"-" gorm:"size:255;not null"`
	Avatar       string `json:"avatar" gorm:"size:255"`
	Bio          string `json:"bio" gorm:"size:400"`
	Role         string `json:"role" gorm:"size:40"`
	Status       string `json:"status" gorm:"size:40"`
	CreatedAt    int64  `json:"createdAt" gorm:"type:bigint"`
	UpdatedAt    int64  `json:"updatedAt" gorm:"type:bigint"`
}

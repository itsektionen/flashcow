package internal

type User struct {
	ID                  int64  `json:"id" db:"id"`
	FullName            string `json:"full_name" db:"full_name"`
	ChapterEmailAddress string `json:"chapter_email" db:"chapter_email"`
}

type Committee struct {
	ID        int64  `json:"id"`
	ShortName string `json:"short_name"`
	Name      string `json:"name"`
}

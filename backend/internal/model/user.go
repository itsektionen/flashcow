package model

type User struct {
	ID                  int64  `json:"id" db:"id"`
	FullName            string `json:"full_name" db:"full_name"`
	ChapterEmailAddress string `json:"chapter_email" db:"chapter_email"`
}

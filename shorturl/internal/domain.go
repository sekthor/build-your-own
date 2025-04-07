package internal

type ShortURL struct {
	ID     string `gorm:"primaryKey" json:"id"`
	Target string `json:"target"`
}

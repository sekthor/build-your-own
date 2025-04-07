package internal

type ShortURL struct {
	ID     string `gorm:"primaryKey"`
	Target string
}

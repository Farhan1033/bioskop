package entity

type Bioskop struct {
	ID     uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name   string  `gorm:"type:varchar(255);not null" json:"name"`
	Lokasi string  `gorm:"type:varchar(100);not null" json:"lokasi"`
	Rating float64 `gorm:"type:decimal;not null" json:"rating"`
}

func (Bioskop) TableName() string {
	return "bioskop"
}

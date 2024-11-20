package models

type User struct {
	ID       uint   `gorm:"primaryKey;column:id_user"` // Especifica que a chave primária é `id_user`
	Name     string `gorm:"not null"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

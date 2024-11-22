package models

type User struct {
	ID       uint   `gorm:"primaryKey;column:id"` // Especifica que a chave primária é `id_user`
	Username     string `gorm:"not null;column:username"`
	Name     string `gorm:"not null;column:full_name"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null;column:password_hash"`
}

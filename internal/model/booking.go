package model

import "time"

// Структура для представления бронирования
type Booking struct {
	ID     int       `json:"id"`      // Уникальный идентификатор бронирования
	UserID int       `json:"user_id"` // Идентификатор пользователя, который сделал бронирование
	DeskID int       `json:"desk_id"` // Идентификатор стола, который забронирован
	Date   time.Time `json:"date"`    // Дата бронирования
}

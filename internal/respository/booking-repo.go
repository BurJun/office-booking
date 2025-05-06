package repository

import (
	"context"                       // Контекст для работы с запросами
	"office-booking/internal/model" // Структура данных Booking

	"github.com/jackc/pgx/v5/pgxpool" // Пул подключений для работы с PostgreSQL
)

// Функция для создания нового бронирования в базе данных
func CreateBooking(pool *pgxpool.Pool, b model.Booking) error {
	// Выполняем SQL-запрос на вставку нового бронирования в таблицу bookings
	_, err := pool.Exec(context.Background(),
		`INSERT INTO bookings (user_id, desk_id, date) VALUES ($1, $2, $3)`,
		b.UserID, b.DeskID, b.Date) // Передаём параметры в запрос через placeholders
	return err // Возвращаем ошибку, если она произошла, или nil, если запрос выполнен успешно
}

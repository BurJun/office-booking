package handler

import (
	"net/http"                                       // Для работы с HTTP статусами
	"office-booking/internal/model"                  // Модели данных (например, Booking)
	repository "office-booking/internal/respository" // Репозитории для работы с БД

	"github.com/gin-gonic/gin"        // Gin для обработки HTTP-запросов
	"github.com/jackc/pgx/v5/pgxpool" // Для работы с PostgreSQL через пул подключений
)

// Функция для обработки запроса на создание нового бронирования
func CreateBookingHandler(db *pgxpool.Pool) gin.HandlerFunc {
	return func(c *gin.Context) {
		var b model.Booking // Создаём переменную для хранения данных бронирования

		// Пробуем привязать JSON из тела запроса к структуре Booking
		if err := c.ShouldBindJSON(&b); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"}) // Если данные некорректны, возвращаем ошибку
			return
		}

		// Пробуем создать новое бронирование в базе данных через репозиторий
		if err := repository.CreateBooking(db, b); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create booking"}) // Если ошибка при сохранении — возвращаем ошибку
			return
		}

		// Возвращаем успешный ответ
		c.JSON(http.StatusCreated, gin.H{"status": "Booking created"})
	}

}

// Функция для регистрации маршрутов и соответствующих обработчиков
func RegisterRoutes(r *gin.Engine, db *pgxpool.Pool) {
	// Регистрируем маршрут для создания нового бронирования
	r.POST("/bookings", CreateBookingHandler(db))
}

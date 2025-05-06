package db

import (
	"context" // Контекст для работы с запросами в Go
	"log"     // Для логирования ошибок
	"os"      // Для работы с переменными окружения

	"github.com/jackc/pgx/v5/pgxpool" // Подключение к PostgreSQL с использованием пула подключений
	"github.com/joho/godotenv"        // Для загрузки переменных из .env файла
)

// Функция для подключения к базе данных PostgreSQL
func ConnectDB() *pgxpool.Pool {
	// Загружаем переменные окружения из файла .env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file") // Ошибка при загрузке .env файла
	}

	// Получаем строку подключения к базе данных из переменной окружения
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:password@localhost:5432/office_booking?sslmode=disable" // Если переменная не задана, используем дефолтную строку
	}

	// Создаём пул подключений к базе данных
	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err) // Если не удалось подключиться, выводим ошибку и завершаем программу
	}

	return pool // Возвращаем пул подключений
}

package main

import (
	"log"                             // Для логирования ошибок
	"office-booking/internal/db"      // Подключение к базе данных
	"office-booking/internal/handler" // Обработчики HTTP-запросов

	"github.com/gin-gonic/gin" // Web-фреймворк для обработки HTTP-запросов
)

func main() {
	// Подключаемся к базе данных с помощью функции из пакета db
	pool := db.ConnectDB()
	defer pool.Close() // Закрываем соединение с БД по завершению работы программы

	r := gin.Default() // Создаём новый роутер Gin

	// Регистрируем маршруты и соответствующие обработчики
	handler.RegisterRoutes(r, pool)

	// Запускаем сервер на порту 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err) // В случае ошибки запуска сервера, выводим её в лог
	}
}

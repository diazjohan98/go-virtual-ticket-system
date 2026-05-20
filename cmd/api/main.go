package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"

	deliveryHttp "github.com/diazjohan98/go-virtual-queue-system/internal/delivery/http"
	"github.com/diazjohan98/go-virtual-queue-system/internal/infrastructure/database"
	redisRepo "github.com/diazjohan98/go-virtual-queue-system/internal/infrastructure/redis"
	"github.com/diazjohan98/go-virtual-queue-system/internal/infrastructure/worker"
	"github.com/diazjohan98/go-virtual-queue-system/internal/usecase"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No se encontró archivo .env local, se usarán las variables del sistema")
	}

	httpPort := os.Getenv("PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf(" Error al preparar MySQL: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Error al conectar a MySQL: %v", err)
	}

	// 4. Conexión dinámica a Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
	})
	defer redisClient.Close()

	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Error al conectar a Redis: %v", err)
	}
	fmt.Println("Conexión segura y exitosa a MySQL y Redis usando variables de entorno")

	// 5. INYECCIÓN DE DEPENDENCIAS
	eventRepo := database.NewMysqlEventRepository(db)
	queueRepo := redisRepo.NewRedisQueueRepository(redisClient)
	enqueueUseCase := usecase.NewEnqueueUserUseCase(eventRepo, queueRepo)
	queueHandler := deliveryHttp.NewQueueHandler(enqueueUseCase)
	queueWorker := worker.NewQueueWorker(queueRepo, eventRepo)

	// 6. ENCENDER EL WORKER EN SEGUNDO PLANO
	eventIDPrueba := "evt-12345"
	go queueWorker.Start(context.Background(), eventIDPrueba)

	// 7. Configurar Rutas HTTP
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("¡El servidor seguro está vivo!"))
	})

	http.HandleFunc("/join-queue", queueHandler.Join)

	// 8. Levantar el servidor de forma dinámica
	fmt.Printf("Servidor seguro corriendo en http://localhost:%s\n", httpPort)
	if err := http.ListenAndServe(":"+httpPort, nil); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}

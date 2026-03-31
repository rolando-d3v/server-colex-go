package config

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

// InitDB inicializa la conexión a PostgreSQL usando pgx
func InitDB() error {
	var err error

	// Obtener DATABASE_URL de las variables de entorno
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		return fmt.Errorf("DATABASE_URL no está configurada")
	}

	// Crear pool de conexiones
	Pool, err = pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		return fmt.Errorf("Unable to create connection pool: %w", err)
	}

	// Verificar la conexión
	err = Pool.Ping(context.Background())
	if err != nil {
		return fmt.Errorf("Unable to ping database: %w", err)
	}

	fmt.Println("✅ Conexión a PostgreSQL exitosa")
	return nil
}

// CloseDB cierra la conexión a la base de datos
func CloseDB() {
	if Pool != nil {
		Pool.Close()
		fmt.Println("❌ Conexión a PostgreSQL cerrada")
	}
}

// Query ejecuta una consulta que retorna múltiples filas
func Query(ctx context.Context, query string, args ...interface{}) (pgx.Rows, error) {
	rows, err := Pool.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("Error en Query: %w", err)
	}
	return rows, nil
}

// QueryRow ejecuta una consulta que retorna una sola fila
func QueryRow(ctx context.Context, query string, args ...interface{}) pgx.Row {
	return Pool.QueryRow(ctx, query, args...)
}

// Exec ejecuta una consulta sin retornar filas (INSERT, UPDATE, DELETE)
func Exec(ctx context.Context, query string, args ...interface{}) (interface{}, error) {
	result, err := Pool.Exec(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("Error en Exec: %w", err)
	}
	return result, nil
}

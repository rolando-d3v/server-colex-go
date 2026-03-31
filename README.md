# 🚀 Server Colex - Go + PostgreSQL + Gin

API REST built con **Go**, **PostgreSQL (pgx)** y **Gin Framework**.

## 📋 Requisitos

- Go 1.21+
- PostgreSQL 13+
- Docker (opcional, para PostgreSQL containerizado)

## 🔧 Instalación

### Opción 1: Con Docker (Recomendado)

```bash
# Levantar PostgreSQL
docker-compose up -d

# Instalar dependencias
go mod download

# Ejecutar servidor
go run main.go
```

### Opción 2: PostgreSQL Local

```bash
# Instalar dependencias
go mod download

# Crear base de datos
createdb -U postgres db_rahemza_colex

# Ejecutar migraciones
psql -U postgres -d db_rahemza_colex < init.sql

# Ejecutar servidor
go run main.go
```

## 📦 Dependencias

- **github.com/gin-gonic/gin** - Web framework
- **github.com/jackc/pgx/v5** - Driver PostgreSQL
- **github.com/joho/godotenv** - Variables de entorno

## 🌐 Endpoints

### Health Check
```
GET /health
```
Verifica estado del servidor.

### Listar Usuarios
```
GET /api/users
```

**Response:**
```json
{
  "success": true,
  "message": "Usuarios obtenidos",
  "data": [
    {
      "Id": "1",
      "persona_id": "P001",
      "colegio_id": 1,
      "email": "usuario@example.com"
    }
  ]
}
```

### Obtener Usuario por ID
```
GET /api/users/:id
```

**Response:**
```json
{
  "success": true,
  "message": "Usuario encontrado",
  "data": {
    "Id": "1",
    "persona_id": "P001",
    "colegio_id": 1,
    "email": "usuario@example.com",
    "password": "hashedpassword"
  }
}
```

### Crear Usuario
```
POST /api/users
```

**Body:**
```json
{
  "persona_id": "P001",
  "colegio_id": 1,
  "email": "nuevo@example.com",
  "password": "securepassword"
}
```

**Response:**
```json
{
  "success": true,
  "message": "Usuario creado",
  "data": {
    "Id": "4",
    "persona_id": "P001",
    "colegio_id": 1,
    "email": "nuevo@example.com",
    "password": "securepassword"
  }
}
```

### Actualizar Usuario
```
PUT /api/users/:id
```

**Body:**
```json
{
  "persona_id": "P002",
  "colegio_id": 2,
  "email": "actualizado@example.com",
  "password": "newpassword"
}
```

### Eliminar Usuario
```
DELETE /api/users/:id
```

**Response:**
```json
{
  "success": true,
  "message": "Usuario eliminado"
}
```

## 📁 Estructura del Proyecto

```
server-colex-go/
├── config/
│   ├── db.go           # Configuración de BD (pgxpool, helpers)
│   └── response.go     # Estructuras de respuesta
├── modules/
│   ├── models.go       # Modelos (User)
│   └── handlers.go     # Handlers CRUD
├── main.go             # Punto de entrada
├── go.mod              # Módulos Go
├── go.sum
├── .env                # Variables de entorno
├── .env.example        # Template de .env
├── init.sql            # Script de migraciones SQL
├── Makefile            # Comandos útiles
├── docker-compose.yml  # PostgreSQL containerizado
└── README.md
```

## 🛠️ Comandos Útiles

```bash
# Ejecutar servidor
go run main.go

# Compilar
go build -o server-colex-go

# Ejecutar tests
go test -v ./...

# Formatear código
go fmt ./...

# Linter
go vet ./...

# Levantar PostgreSQL con Docker
docker-compose up -d

# Detener PostgreSQL
docker-compose down

# Ejecutar migraciones
make db-migrate

# Ver logs de Docker
docker-compose logs -f postgres
```

## 🔌 Configuración de BD

Las funciones helper en `config/db.go` simplifican las queries:

```go
// Query - Múltiples filas
rows, err := Query(ctx, "SELECT * FROM usuario")

// QueryRow - Una sola fila
err := QueryRow(ctx, "SELECT * FROM usuario WHERE id = $1", id).Scan(&user)

// Exec - INSERT, UPDATE, DELETE
result, err := Exec(ctx, "UPDATE usuario SET email = $1 WHERE id = $2", email, id)
```

## ⚙️ Variables de Entorno

```env
PORT=4002
DATABASE_URL=postgresql://usuario:contraseña@localhost:5432/db_rahemza_colex
```

## 📝 Notas

- El pool de conexiones está configurado con:
  - **MaxConns**: 25
  - **MinConns**: 5
  - **MaxConnLifetime**: 1 hora
  - **MaxConnIdleTime**: 10 minutos

- Las queries tienen timeout de 5-10 segundos
- La tabla `usuario` incluye índice en `email` para búsquedas rápidas
- Se implementa graceful shutdown para cerrar conexiones correctamente

## 🐛 Troubleshooting

### "DATABASE_URL no está configurada"
Asegúrate de crear un archivo `.env` con la configuración de conexión.

### "Error conectando a PostgreSQL"
- Verifica que PostgreSQL esté running
- Valida las credenciales en `DATABASE_URL`
- Comprueba que el servidor esté accesible

### "Tabla usuario no existe"
Ejecuta el script de migraciones:
```bash
make db-migrate
```

## 📄 Licencia

MIT

## 📦 Dependencias

```bash
go get github.com/gin-gonic/gin@v1.12.0
go get github.com/jackc/pgx/v5@v5.5.5
go get github.com/joho/godotenv@v1.5.1
```

## ⚙️ Configuración

### 1. Archivo `.env`

```env
PORT=4002
DATABASE_URL=postgres://usuario:password@localhost:5432/colex_db
```

**Formato de DATABASE_URL:**
```
postgres://[usuario]:[contraseña]@[host]:[puerto]/[base_de_datos]
```

### 2. Crear base de datos en PostgreSQL

```sql
-- Conectar a PostgreSQL
psql -U postgres

-- Crear base de datos
CREATE DATABASE colex_db;

-- Conectar a la nueva base de datos
\c colex_db

-- Ejecutar el script de inicialización (dentro de la BD)
\i init.sql
```

**O si prefieres en línea:**
```bash
psql -U postgres -d colex_db < init.sql
```

## 📁 Estructura del Proyecto

```
server-colex-go/
├── config/
│   ├── db.go          # Configuración de pgxpool
│   └── response.go    # Helpers para respuestas JSON
├── modules/
│   ├── models.go      # Definición de modelos (User)
│   └── handlers.go    # Handlers/Rutas (CRUD de usuarios)
├── main.go            # Entry point, rutas principales
├── .env               # Variables de entorno
├── go.mod             # Dependencias
├── go.sum             # Hash de dependencias
└── init.sql           # Script de inicialización de BD
```

## 🏃 Ejecutar el Proyecto

### 1. Descargar dependencias
```bash
go mod download
go mod tidy
```

### 2. Crear las tablas en PostgreSQL
```bash
psql -U postgres -d colex_db < init.sql
```

### 3. Ejecutar el servidor
```bash
go run main.go
```

O con hot-reload (requiere `air`):
```bash
go install github.com/cosmtrek/air@latest
air
```

## 🔗 Endpoints

### Health Check
- `GET /health` - Verificar que el servidor está activo

### Usuarios (CRUD)
- `GET /api/users` - Obtener todos los usuarios
- `GET /api/users/:id` - Obtener usuario por ID
- `POST /api/users` - Crear nuevo usuario
- `PUT /api/users/:id` - Actualizar usuario
- `DELETE /api/users/:id` - Eliminar usuario

### Ejemplos de uso

**Obtener todos los usuarios:**
```bash
curl http://localhost:4002/api/users
```

**Obtener usuario por ID:**
```bash
curl http://localhost:4002/api/users/1
```

**Crear usuario:**
```bash
curl -X POST http://localhost:4002/api/users \
  -H "Content-Type: application/json" \
  -d '{"name": "Luis", "email": "luis@example.com"}'
```

**Actualizar usuario:**
```bash
curl -X PUT http://localhost:4002/api/users/1 \
  -H "Content-Type: application/json" \
  -d '{"name": "Luis Updated", "email": "luis.updated@example.com"}'
```

**Eliminar usuario:**
```bash
curl -X DELETE http://localhost:4002/api/users/1
```

## 🔑 Características Principales

### 1. **Pool de Conexiones (pgxpool)**
- Máximo 25 conexiones
- Mínimo 5 conexiones
- Tiempo de vida máximo: 1 hora
- Timeout de inactividad: 10 minutos

### 2. **Manejo de Contexto**
Todos los handlers usan `context.WithTimeout` para operaciones de base de datos:
```go
ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
defer cancel()
```

### 3. **Respuestas Estandarizadas**
```json
{
  "success": true,
  "message": "Descripción de la operación",
  "data": { ... }
}
```

### 4. **Helpers de Respuesta**
```go
SendSuccess(c, http.StatusOK, "Mensaje", data)
SendError(c, http.StatusBadRequest, "Error")
```

## 🛠️ Agregar Nuevas Rutas

### 1. Crear modelo en `modules/models.go`
```go
type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name" binding:"required"`
    Price float64 `json:"price" binding:"required"`
}
```

### 2. Agregar handlers en `modules/handlers.go`
```go
func GetAllProducts(c *gin.Context) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    
    rows, err := Query(ctx, "SELECT id, name, price FROM products")
    // ... resto del código
}
```

### 3. Registrar rutas en `main.go`
```go
productsGroup := router.Group("/api/products")
{
    productsGroup.GET("", modules.GetAllProducts)
    productsGroup.POST("", modules.CreateProduct)
}
```

## 🐛 Solución de Problemas

### Error: "connection refused"
- Verificar que PostgreSQL está en ejecución
- Validar las credenciales en `.env`
- Asegurar que la base de datos existe

### Error: "DATABASE_URL no está configurada"
- Revisar que `.env` existe en la raíz del proyecto
- Verificar sintaxis de `DATABASE_URL`

### Error de conexión timeout
- Aumentar el timeout en `config/db.go`
- Verificar capacidad del servidor PostgreSQL

## 📚 Referencias

- [pgx Documentation](https://pkg.go.dev/github.com/jackc/pgx/v5)
- [Gin Documentation](https://gin-gonic.com/docs/)
- [PostgreSQL](https://www.postgresql.org/docs/)

## 📝 Licencia

MIT

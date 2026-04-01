package modules

import (
	"context"
	"net/http"
	"strconv"
	"github.com/google/uuid"
	"time"
	"github.com/gin-gonic/gin"
	
	. "server-colex-go/config"

)

//? GetAllUsers obtiene todos los usuarios
//? ***********************************************************************************************/
func GetAllUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	rows, err := Query(ctx, "SELECT id, persona_id, colegio_id, email, is_active, created_at FROM usuario")
	if err != nil {
		SendError(c, http.StatusInternalServerError, "Error consultando usuarios")
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.PersonaID, &user.ColegioID, &user.Email, &user.IsActive, &user.CreatedAt)
		if err != nil {
			SendError(c, http.StatusInternalServerError, "Error leyendo datos")
			return
		}
		users = append(users, user)
	}

	SendSuccess(c, http.StatusOK, "Usuarios obtenidos", users)
}

// GetUserByID obtiene un usuario por ID
//? ***********************************************************************************************/
func GetUserByID(c *gin.Context) {
	idStr := c.Param("id")

	// 🔥 convertir string → UUID
	id, err := uuid.Parse(idStr)
	if err != nil {
		SendError(c, http.StatusBadRequest, "ID inválido")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user User
	err = QueryRow(ctx,
		"SELECT id, persona_id, colegio_id, email, password FROM usuario WHERE id = $1",
		id,
	).Scan(&user.Id, &user.PersonaID, &user.ColegioID, &user.Email, &user.Password)

	if err != nil {
		SendError(c, http.StatusNotFound, "Usuario no encontrado")
		return
	}

	SendSuccess(c, http.StatusOK, "Usuario encontrado", user)
}

// CreateUser crea un nuevo usuario
//? ***********************************************************************************************/
func CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var id uuid.UUID
	err := QueryRow(ctx,
		"INSERT INTO usuario (persona_id, colegio_id, email, password) VALUES ($1, $2, $3, $4) RETURNING id",
		user.PersonaID, user.ColegioID, user.Email, user.Password).Scan(&id)
	if err != nil {
		SendError(c, http.StatusInternalServerError, "Error creando usuario")
		return
	}

	user.Id = id
	SendSuccess(c, http.StatusCreated, "Usuario creado", user)
}

// UpdateUser actualiza un usuario existente
//? ***********************************************************************************************/
func UpdateUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		SendError(c, http.StatusBadRequest, "ID inválido")
		return
	}

	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		SendError(c, http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := Exec(ctx,
		"UPDATE usuario SET persona_id = $1, colegio_id = $2, email = $3, password = $4 WHERE id = $5",
		user.PersonaID, user.ColegioID, user.Email, user.Password, id)
	if err != nil {
		SendError(c, http.StatusInternalServerError, "Error actualizando usuario: "+err.Error())
		return
	}

	if result.(interface{ RowsAffected() int64 }).RowsAffected() == 0 {
		SendError(c, http.StatusNotFound, "Usuario no encontrado")
		return
	}

	SendSuccess(c, http.StatusOK, "Usuario actualizado", user)
}

// DeleteUser elimina un usuario
//? ***********************************************************************************************/
func DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		SendError(c, http.StatusBadRequest, "ID inválido")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := Exec(ctx, "DELETE FROM usuario WHERE id = $1", id)
	if err != nil {
		SendError(c, http.StatusInternalServerError, "Error eliminando usuario")
		return
	}

	// Verificar si se eliminó alguna fila
	if result.(interface{ RowsAffected() int64 }).RowsAffected() == 0 {
		SendError(c, http.StatusNotFound, "Usuario no encontrado")
		return
	}

	SendSuccess(c, http.StatusOK, "Usuario eliminado", nil)
}

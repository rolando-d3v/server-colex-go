package modules

import "github.com/google/uuid"

// Estructura para el login request
type LoginRequest struct {
	Codigo   string `json:"codigo" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Estructura para el usuario Response
type Usuariox struct {
	Id            uuid.UUID `json:"id"`
	PersonaID     string    `json:"persona_id"`
	CodigoUsuario string    `json:"codigo_usuario"`
	ColegioID     int       `json:"colegio_id"`
	Password      string    `json:"-"`
	IsActive      bool      `json:"is_active"`
	Roles        []Role    `json:"roles"`
}
type Role struct {
	Id     int    `json:"id"`
	Nombre string `json:"nombre"`
}

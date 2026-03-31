package modules

// User modelo para la tabla usuarios
type User struct {
	Id    string  `json:"id"`
	PersonaID    string  `json:"persona_id"`
	ColegioID    int16  `json:"colegio_id"`
	Password string  `json:"password,omitempty"`
	Email string  `json:"email"`
}

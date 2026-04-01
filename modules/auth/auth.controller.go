package modules

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	// files
	. "server-colex-go/config"
)

// ? AuthLogin inicia sesión
// ? ***********************************************************************************************/
func AuthLogin(c *gin.Context) {
	// 1. Binding y validación automática
	var body LoginRequest

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msj": "Campos requeridos: codigo y password ❗️"})
		return
	}

	// 2. Buscar usuario activo
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user Usuariox
	var rolesJSON []byte

	query := `
		SELECT 
		u.id,
		u.persona_id,
		u.codigo_usuario,
		u.colegio_id,
		u.password,
		u.is_active,
		json_agg(
			json_build_object(
				'id', r.id,
				'nombre', r.nombre
			)
		) AS roles
	    FROM usuario u
	    INNER JOIN usuario_rol ur ON ur.usuario_id = u.id
	    INNER JOIN rol r ON r.id = ur.rol_id
	    WHERE u.codigo_usuario = $1
	    AND u.is_active = true
	    GROUP BY u.id;
	`
	err := QueryRow(ctx, query, body.Codigo).Scan(
		&user.Id,
		&user.PersonaID,
		&user.CodigoUsuario,
		&user.ColegioID,
		&user.Password,
		&user.IsActive,
		&rolesJSON)

	if err != nil {
		log.Println("Error DB:", err)
		SendError(c, http.StatusNotFound, "Usuario no encontrado")
		return
	}

	// Parsear roles JSON después de obtener los datos del query
	if err := json.Unmarshal(rolesJSON, &user.Roles); err != nil {
		log.Println("Error parse roles:", err)
	}

	log.Printf("User: %+v\n", user)

	// 3. Comparar password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msj": "Credenciales inválidas ❗️"})
		return
	}

	// 4. Generar JWT
	secretToken := os.Getenv("SECRET_TOKEN")
	claims := jwt.MapClaims{
		"id":             user.Id,
		"persona_id":     user.PersonaID,
		"codigo_usuario": user.CodigoUsuario,
		"colegio_id":     user.ColegioID,
		"ok":             true,
		"exp":            time.Now().Add(5 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretToken))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msj": "Error interno del servidor ❗️"})
		return
	}

	// 5. Setear cookie
	c.SetCookie(
		"token",     // nombre
		tokenString, // valor
		60*60*5,     // maxAge en segundos (5 horas)
		"/",         // path
		"",          // domain
		false,       // secure (true en producción con HTTPS)
		true,        // httpOnly 🔒
	)

	// Devolver el usuario y un rol por defecto para que el frontend no falle
	c.JSON(http.StatusOK, gin.H{
		"msj": "Login exitoso ✔️",
		"user": gin.H{
			"id":             user.Id,
			"codigo_usuario": user.CodigoUsuario,
			"colegio_id":     user.ColegioID, // mock temporal
		},
		"roles": user.Roles,
	})
}

// ? Verifica la cookie httpOnly y retorna la sesión
// ? ***********************************************************************************************/
func VerifyAuth(c *gin.Context) {
	// 1. Obtener la cookie "token"
	tokenString, err := c.Cookie("token")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"msj": "No autorizado, falta token ❗️"})
		return
	}

	// 2. Parsear el token
	secretToken := os.Getenv("SECRET_TOKEN")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretToken), nil
	})

	// 3. Verificar validez
	if err != nil || !token.Valid {
		// Limpiar cookie si el token es inválido
		c.SetCookie("token", "", -1, "/", "", false, true)
		c.JSON(http.StatusUnauthorized, gin.H{"msj": "Token inválido o expirado ❗️"})
		return
	}

	// 4. Leer los claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		c.JSON(http.StatusOK, gin.H{
			"ok": true,
			"user": gin.H{
				"id":             claims["id"],
				"codigo_usuario": claims["codigo_usuario"],
				"colegio_id":     claims["colegio_id"],
			},
			"roles": []string{"admin_colegio"}, // de nuevo, mock de tu rol esperado
		})
		return
	}

	c.JSON(http.StatusUnauthorized, gin.H{"msj": "Error leyendo claims ❗️"})
}

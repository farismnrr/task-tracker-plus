/** 
 * Package middleware provides middleware functions for authentication and authorization in web applications.
 * 
 * Functions:
 * 
 * - Auth: Function to create an authentication middleware.
 *   Returns:
 *   - gin.HandlerFunc: A Gin middleware handler function.
 *   Description: This function returns a Gin middleware handler function that performs authentication. 
 *     It checks for the presence of a session token in the request cookie. 
 *     If the token is missing, it returns an unauthorized response or redirects the user to the login page based on the request content type. 
 *     If the token is present, it parses and validates the token using JWT and sets the user's email in the Gin context for further request processing.
 */

package middleware

import (
	"a21hc3NpZ25tZW50/model"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Auth() gin.HandlerFunc {
	return gin.HandlerFunc(func(ctx *gin.Context) {
		authHeader, err := ctx.Cookie("session_token")
		if err != nil {
			if ctx.Request.Header.Get("Content-Type") == "application/json" {
				ctx.JSON(http.StatusUnauthorized, model.NewErrorResponse("Unauthorized"))
				ctx.Abort()
				return
			}
			ctx.Redirect(http.StatusSeeOther, "/login")
			return
		}

		claims := &model.Claims{}

		token, err := jwt.ParseWithClaims(authHeader, claims, func(token *jwt.Token) (interface{}, error) {
			return model.JwtKey, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				ctx.JSON(http.StatusUnauthorized, model.NewErrorResponse("Unauthorized"))
				ctx.Abort()
				return
			}
			ctx.JSON(http.StatusBadRequest, model.NewErrorResponse("Bad Request"))
			ctx.Abort()
			return
		}

		if !token.Valid {
			ctx.JSON(http.StatusUnauthorized, model.NewErrorResponse("Unauthorized"))
			ctx.Abort()
			return
		}

		ctx.Set("email", claims.Email)
		ctx.Next()
	})
}

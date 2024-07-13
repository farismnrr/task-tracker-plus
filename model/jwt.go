/** 
 * Package model provides data models and utility functions for handling authentication in web applications.
 * 
 * Variables:
 * 
 * - JwtKey: Variable containing the secret key used for JWT token generation and validation.
 *   Type: []byte
 *   Description: This variable holds the secret key used for generating and validating JWT tokens.
 * 
 * Structs:
 * 
 * - Claims: Struct representing the JWT claims including user email.
 *   Fields:
 *   - Email: Email address of the user.
 *     Type: string
 *     Description: This field holds the email address of the user included in the JWT claims.
 *   - StandardClaims: Embedded struct containing standard JWT claims.
 *     Type: jwt.StandardClaims
 *     Description: This embedded struct contains standard JWT claims such as expiration time, issuer, and subject.
 */

package model

import "github.com/golang-jwt/jwt"

var JwtKey = []byte("secret-key")

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

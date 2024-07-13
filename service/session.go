/**
 * Package service provides interfaces and implementations for managing sessions.
 * 
 * Interfaces:
 * 
 * - SessionService: Interface defining methods for session management.
 *   Methods:
 *   - GetSessionByEmail: Method to retrieve a session by email.
 * 
 * Structs:
 * 
 * - sessionService: Struct implementing the SessionService interface.
 *   Fields:
 *   - sessionRepo: Instance of repo.SessionRepository for session repository operations.
 *   Methods:
 *   - NewSessionService: Function to create a new instance of sessionService.
 *   - GetSessionByEmail: Method to retrieve a session by email using the session repository.
 */

package service

import (
	"a21hc3NpZ25tZW50/model"
	repo "a21hc3NpZ25tZW50/repository"
)

type SessionService interface {
	GetSessionByEmail(email string) (model.Session, error)
}

type sessionService struct {
	sessionRepo repo.SessionRepository
}

func NewSessionService(sessionRepo repo.SessionRepository) *sessionService {
	return &sessionService{sessionRepo}
}

func (c *sessionService) GetSessionByEmail(email string) (model.Session, error) {
	return c.sessionRepo.SessionAvailEmail(email)
}

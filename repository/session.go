/** 
 * Package repository provides interfaces and implementations for managing user sessions.
 * 
 * Interfaces:
 * 
 * - SessionRepository: Interface defining methods for session data manipulation.
 *   Methods:
 *   - AddSessions: Method to add a new session.
 *   - DeleteSession: Method to delete a session by token.
 *   - UpdateSessions: Method to update an existing session.
 *   - SessionAvailEmail: Method to check if a session is available by email.
 *   - SessionAvailToken: Method to check if a session is available by token.
 *   - TokenValidity: Method to validate a session token.
 *   - TokenExpired: Method to check if a session token has expired.
 *
 * Structs
 * 
 * - sessionsRepo: Struct implementing the SessionRepos
 *   Fields:
 *   - filebasedDb: Instance of filebased.Data for files.
 *   Methods:
 *   - NewSessionsRepo: Function to create a new instance of sessionsRepo.
 *   - AddSessions: Method to add a new session using file-based database operations.
 *   - DeleteSession: Method to delete a session by token using file-based database operations.
 *   - UpdateSessions: Method to update ting session using file-based database operations.
 *   - SessionAvailEmail: Method to checsession is available by email using file-based database operations.
 *   - SessionAvailToken: Method  check ssion is available by token using file-based database operations.
 *   - TokenValidity: Method to vidate a session token and delete it if expired using file-based database operations.
 *   - TokenExpired: Method to check if a session token has expired.
 */

package repository

import (
	"a21hc3NpZ25tZW50/db/filebased"
	"a21hc3NpZ25tZW50/model"
	"time"
)

type SessionRepository interface {
	AddSessions(session model.Session) error
	DeleteSession(token string) error
	UpdateSessions(session model.Session) error
	SessionAvailEmail(email string) (model.Session, error)
	SessionAvailToken(token string) (model.Session, error)
	TokenExpired(session model.Session) bool
}

type sessionsRepo struct {
	filebasedDb *filebased.Data
}

func NewSessionsRepo(filebasedDb *filebased.Data) *sessionsRepo {
	return &sessionsRepo{filebasedDb}
}

func (u *sessionsRepo) AddSessions(session model.Session) error {
	return u.filebasedDb.AddSession(session)
}

func (u *sessionsRepo) DeleteSession(token string) error {
	return u.filebasedDb.DeleteSession(token)
}

func (u *sessionsRepo) UpdateSessions(session model.Session) error {
	return u.filebasedDb.UpdateSession(session)
}

func (u *sessionsRepo) SessionAvailEmail(email string) (model.Session, error) {
	return u.filebasedDb.SessionAvailEmail(email)
}

func (u *sessionsRepo) SessionAvailToken(token string) (model.Session, error) {
	return u.filebasedDb.SessionAvailToken(token)
}

func (u *sessionsRepo) TokenValidity(token string) (model.Session, error) {
	session, err := u.SessionAvailToken(token)
	if err != nil {
		return model.Session{}, err
	}

	if u.TokenExpired(session) {
		err := u.DeleteSession(token)
		if err != nil {
			return model.Session{}, err
		}
		return model.Session{}, err
	}

	return session, nil
}

func (u *sessionsRepo) TokenExpired(session model.Session) bool {
	return session.Expiry.Before(time.Now())
}

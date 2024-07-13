/** 
 * Package model provides data models for representing various entities in the application domain.
 * 
 * Structs:
 * 
 * - Category: Struct representing a task category.
 *   Fields:
 *   - ID: Unique identifier for the category.
 *     Type: int
 *   - Name: Name of the category.
 *     Type: string
 * 
 * - User: Struct representing a user.
 *   Fields:
 *   - ID: Unique identifier for the user.
 *     Type: int
 *   - Fullname: Full name of the user.
 *     Type: string
 *   - Email: Email address of the user.
 *     Type: string
 *   - Password: Password of the user.
 *     Type: string
 *   - CreatedAt: Timestamp indicating the creation time of the user record.
 *     Type: time.Time
 *   - UpdatedAt: Timestamp indicating the last update time of the user record.
 *     Type: time.Time
 * 
 * - UserLogin: Struct representing user login credentials.
 *   Fields:
 *   - Email: Email address of the user.
 *     Type: string
 *   - Password: Password of the user.
 *     Type: string
 * 
 * - UserRegister: Struct representing user registration data.
 *   Fields:
 *   - Fullname: Full name of the user.
 *     Type: string
 *   - Email: Email address of the user.
 *     Type: string
 *   - Password: Password of the user.
 *     Type: string
 * 
 * - Task: Struct representing a task.
 *   Fields:
 *   - ID: Unique identifier for the task.
 *     Type: int
 *   - Title: Title of the task.
 *     Type: string
 *   - Deadline: Deadline of the task.
 *     Type: string
 *   - Priority: Priority level of the task.
 *     Type: int
 *   - Status: Status of the task.
 *     Type: string
 *   - CategoryID: ID of the category to which the task belongs.
 *     Type: int
 *   - UserID: ID of the user who owns the task.
 *     Type: int
 * 
 * - Session: Struct representing a user session.
 *   Fields:
 *   - ID: Unique identifier for the session.
 *     Type: int
 *   - Token: Token associated with the session.
 *     Type: string
 *   - Email: Email address of the user associated with the session.
 *     Type: string
 *   - Expiry: Timestamp indicating the expiry time of the session.
 *     Type: time.Time
 * 
 * - TaskCategory: Struct representing a task category with additional details.
 *   Fields:
 *   - ID: Unique identifier for the category.
 *     Type: int
 *   - Title: Title of the category.
 *     Type: string
 *   - Category: Additional category information.
 *     Type: string
 * 
 * - UserTaskCategory: Struct representing user task details with category information.
 *   Fields:
 *   - ID: Unique identifier for the task.
 *     Type: int
 *   - Fullname: Full name of the user.
 *     Type: string
 *   - Email: Email address of the user.
 *     Type: string
 *   - Task: Title of the task.
 *     Type: string
 *   - Deadline: Deadline of the task.
 *     Type: string
 *   - Priority: Priority level of the task.
 *     Type: int
 *   - Status: Status of the task.
 *     Type: string
 *   - Category: Name of the category to which the task belongs.
 *     Type: string
 * 
 * - Credential: Struct representing database connection credentials.
 *   Fields:
 *   - Host: Database host address.
 *     Type: string
 *   - Username: Username for database authentication.
 *     Type: string
 *   - Password: Password for database authentication.
 *     Type: string
 *   - DatabaseName: Name of the database.
 *     Type: string
 *   - Port: Port number for database connection.
 *     Type: int
 *   - Schema: Database schema name.
 *     Type: string
 */

package model

import "time"

type Category struct {
	ID   int    `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
}

type User struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Fullname  string    `json:"fullname" gorm:"type:varchar(255);"`
	Email     string    `json:"email" gorm:"type:varchar(255);not null"`
	Password  string    `json:"password" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserRegister struct {
	Fullname string `json:"fullname" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Task struct {
	ID         int    `gorm:"primaryKey" json:"id"`
	Title      string `json:"title"`
	Deadline   string `json:"deadline"`
	Priority   int    `json:"priority"`
	Status     string `json:"status"`
	CategoryID int    `json:"category_id"`
	UserID     int    `json:"user_id"`
}

type Session struct {
	ID     int       `gorm:"primaryKey" json:"id"`
	Token  string    `json:"token"`
	Email  string    `json:"email"`
	Expiry time.Time `json:"expiry"`
}

type TaskCategory struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"`
}

type UserTaskCategory struct {
	ID       int    `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Task     string `json:"task"`
	Deadline string `json:"deadline"`
	Priority int    `json:"priority"`
	Status   string `json:"status"`
	Category string `json:"category"`
}

type Credential struct {
	Host         string
	Username     string
	Password     string
	DatabaseName string
	Port         int
	Schema       string
}

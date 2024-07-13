/** 
 * Package model provides data models and utility functions for handling responses in web applications.
 * 
 * Structs:
 * 
 * - ErrorResponse: Struct representing an error response.
 *   Fields:
 *   - Error: Error message.
 *     Type: string
 * 
 * - SuccessResponse: Struct representing a success response.
 *   Fields:
 *   - Message: Success message.
 *     Type: string
 * 
 * Functions:
 * 
 * - NewErrorResponse: Function to create a new ErrorResponse instance.
 *   Parameters:
 *   - msg: Error message.
 *     Type: string
 *   Returns:
 *   - ErrorResponse: New instance of ErrorResponse.
 * 
 * - NewSuccessResponse: Function to create a new SuccessResponse instance.
 *   Parameters:
 *   - msg: Success message.
 *     Type: string
 *   Returns:
 *   - SuccessResponse: New instance of SuccessResponse.
 */

package model

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string `json:"message"`
}

func NewErrorResponse(msg string) ErrorResponse {
	return ErrorResponse{
		Error: msg,
	}
}

func NewSuccessResponse(msg string) SuccessResponse {
	return SuccessResponse{
		Message: msg,
	}
}

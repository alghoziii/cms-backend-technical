package dto

// Auth DTOs
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

// Category DTOs
type CategoryRequest struct {
	Name string `json:"name" binding:"required"`
}

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// News DTOs
type NewsRequest struct {
	Title      string `json:"title" binding:"required"`
	Content    string `json:"content" binding:"required"`
	CategoryID uint   `json:"category_id" binding:"required"`
}

type NewsResponse struct {
	ID         uint   `json:"id"`
	Title      string `json:"title"`
	Content    string `json:"content"`
	CategoryID uint   `json:"category_id"`
	UserID     uint   `json:"user_id"`
}

// CustomPage DTOs
type CustomPageRequest struct {
	URL     string `json:"url" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type CustomPageResponse struct {
	ID      uint   `json:"id"`
	URL     string `json:"url"`
	Content string `json:"content"`
	UserID  uint   `json:"user_id"`
}

// Comment DTOs
type CommentRequest struct {
	Name    string `json:"name" binding:"required"`
	Comment string `json:"comment" binding:"required"`
}

type CommentResponse struct {
	ID      uint   `json:"id"`
	Name    string `json:"name"`
	Comment string `json:"content"`
	NewsID  uint   `json:"news_id"`
}

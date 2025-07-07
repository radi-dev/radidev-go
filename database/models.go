package database

import "time"

type Post struct {
	ID               string    `db:"id"`
	Title            string    `db:"title"`
	IsProject        bool      `db:"is_project"`
	IsPublished      bool      `db:"is_published"`
	Tags             []string  `db:"tags"` // jsonb
	AuthorID         string    `db:"author_id"`
	Text             string    `db:"text"`
	FeaturedImageURL string    `db:"featured_image_url"`
	CreatedAt        time.Time `db:"created_at"`
	UpdatedAt        time.Time `db:"updated_at"`
}

type Comment struct {
	ID          string    `db:"id"`
	Text        string    `db:"text"`
	PostID      *string   `db:"post_id"`    // Nullable
	CommentID   *string   `db:"comment_id"` // Nullable (reply)
	AuthorName  string    `db:"author_name"`
	AuthorEmail *string   `db:"author_email"` // Nullable
	CreatedAt   time.Time `db:"created_at"`
}

type Author struct {
	ID    string `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
	About string `db:"about"`
}

type Upload struct {
	ID   string `db:"upload_id"`
	Name string `db:"name"`
	URL  string `db:"url"`
}

type User struct {
	ID           string `db:"id"`
	Username     string `db:"username"`
	PasswordHash string `db:"password_hash"`
	CreatedAt    string `db:"created_at"`
}

type Tag struct {
	ID    string `db:"id"`
	Name  string `db:"name"`
	Posts []Post `db:"posts"` // Many-to-many relationship with posts
}

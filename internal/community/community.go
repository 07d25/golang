// Package community contiene las reglas y endpoints de la comunidad gamer.
package community

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi/v5"
	"noniserver/m/v2/internal/chat"
)

type Store struct{ db *sql.DB }

func NewStore(db *sql.DB) *Store { return &Store{db: db} }

type User struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Bio         string `json:"bio"`
}
type Post struct {
	ID      int64  `json:"id"`
	UserID  int64  `json:"user_id"`
	Content string `json:"content"`
	Likes   int    `json:"likes"`
}
type Room struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}
type Message struct {
	ID      int64  `json:"id"`
	RoomID  int64  `json:"room_id"`
	UserID  int64  `json:"user_id"`
	Content string `json:"content"`
}
type Comment struct {
	ID      int64  `json:"id"`
	UserID  int64  `json:"user_id"`
	PostID  int64  `json:"post_id"`
	Content string `json:"content"`
}
type userRequest struct {
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
	Bio         string `json:"bio"`
	Password    string `json:"password"`
}
type postRequest struct {
	Content string `json:"content"`
}
type roomRequest struct {
	Name string `json:"name"`
}
type messageRequest struct {
	UserID  int64  `json:"user_id"`
	Content string `json:"content"`
}
type commentRequest struct {
	Content string `json:"content"`
}
type member struct {
	UserID int64 `json:"user_id"`
	RoomID int64 `json:"room_id"`
}

func (s *Store) CreateUser(ctx context.Context, req userRequest) (User, error) {
	if strings.TrimSpace(req.Username) == "" || strings.TrimSpace(req.DisplayName) == "" {
		return User{}, errors.New("username y display_name son obligatorios")
	}
	if len(req.Password) < 8 {
		return User{}, errors.New("password debe tener al menos 8 caracteres")
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return User{}, err
	}
	var u User
	err = s.db.QueryRowContext(ctx, "INSERT INTO users(username,display_name,bio,password_hash) VALUES($1,$2,$3,$4) RETURNING id,username,display_name,bio", strings.TrimSpace(req.Username), strings.TrimSpace(req.DisplayName), strings.TrimSpace(req.Bio), string(hash)).Scan(&u.ID, &u.Username, &u.DisplayName, &u.Bio)
	return u, err
}

func (s *Store) Authenticate(ctx context.Context, username, password string) (int64, error) {
	var id int64
	var hash string
	err := s.db.QueryRowContext(ctx, "SELECT id,password_hash FROM users WHERE username=$1 AND deleted_at IS NULL", strings.TrimSpace(username)).Scan(&id, &hash)
	if err != nil {
		return 0, errors.New("credenciales inválidas")
	}
	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil {
		return 0, errors.New("credenciales inválidas")
	}
	return id, nil
}

func (s *Store) CreatePost(ctx context.Context, userID int64, req postRequest) (Post, error) {
	if strings.TrimSpace(req.Content) == "" {
		return Post{}, errors.New("content es obligatorio")
	}
	if len([]rune(req.Content)) > 500 {
		return Post{}, errors.New("content no puede superar 500 caracteres")
	}
	var p Post
	err := s.db.QueryRowContext(ctx, "INSERT INTO posts(user_id,content) VALUES($1,$2) RETURNING id,user_id,content", userID, strings.TrimSpace(req.Content)).Scan(&p.ID, &p.UserID, &p.Content)
	return p, err
}

func (s *Store) CreateRoom(ctx context.Context, req roomRequest) (Room, error) {
	if strings.TrimSpace(req.Name) == "" {
		return Room{}, errors.New("name es obligatorio")
	}
	var room Room
	err := s.db.QueryRowContext(ctx, "INSERT INTO rooms(name) VALUES($1) RETURNING id,name", strings.TrimSpace(req.Name)).Scan(&room.ID, &room.Name)
	return room, err
}

func (s *Store) CreateMessage(ctx context.Context, roomID int64, req messageRequest) (Message, error) {
	if req.UserID <= 0 || strings.TrimSpace(req.Content) == "" {
		return Message{}, errors.New("user_id positivo y content obligatorio")
	}
	if len([]rune(req.Content)) > 500 {
		return Message{}, errors.New("content no puede superar 500 caracteres")
	}
	member, err := s.IsMember(ctx, roomID, req.UserID)
	if err != nil {
		return Message{}, err
	}
	if !member {
		return Message{}, errors.New("el usuario debe unirse a la sala")
	}
	var m Message
	err = s.db.QueryRowContext(ctx, "INSERT INTO messages(room_id,user_id,content) VALUES($1,$2,$3) RETURNING id,room_id,user_id,content", roomID, req.UserID, strings.TrimSpace(req.Content)).Scan(&m.ID, &m.RoomID, &m.UserID, &m.Content)
	return m, err
}

func (s *Store) JoinRoom(ctx context.Context, roomID, userID int64) error {
	if roomID <= 0 || userID <= 0 {
		return errors.New("room_id y user_id deben ser positivos")
	}
	var exists bool
	if err := s.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM rooms WHERE id=$1)", roomID).Scan(&exists); err != nil {
		return err
	}
	if !exists {
		return errors.New("sala no encontrada")
	}
	if err := s.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM users WHERE id=$1 AND deleted_at IS NULL)", userID).Scan(&exists); err != nil {
		return err
	}
	if !exists {
		return errors.New("usuario no encontrado")
	}
	_, err := s.db.ExecContext(ctx, "INSERT INTO room_members(room_id,user_id) VALUES($1,$2) ON CONFLICT DO NOTHING", roomID, userID)
	return err
}

func (s *Store) LeaveRoom(ctx context.Context, roomID, userID int64) error {
	_, err := s.db.ExecContext(ctx, "DELETE FROM room_members WHERE room_id=$1 AND user_id=$2", roomID, userID)
	return err
}
func (s *Store) IsMember(ctx context.Context, roomID, userID int64) (bool, error) {
	var exists bool
	err := s.db.QueryRowContext(ctx, "SELECT EXISTS(SELECT 1 FROM room_members WHERE room_id=$1 AND user_id=$2)", roomID, userID).Scan(&exists)
	return exists, err
}
func (s *Store) SaveEvent(ctx context.Context, event chat.Event) (chat.Event, error) {
	roomID, err := strconv.ParseInt(event.RoomID, 10, 64)
	if err != nil {
		return event, errors.New("room_id inválido")
	}
	userID, err := strconv.ParseInt(event.UserID, 10, 64)
	if err != nil {
		return event, errors.New("user_id inválido")
	}
	member, err := s.IsMember(ctx, roomID, userID)
	if err != nil {
		return event, err
	}
	if !member {
		return event, errors.New("el usuario no pertenece a la sala")
	}
	message, err := s.CreateMessage(ctx, roomID, messageRequest{UserID: userID, Content: event.Content})
	if err != nil {
		return event, err
	}
	event.Type = "message"
	event.UserID = strconv.FormatInt(message.UserID, 10)
	event.RoomID = strconv.FormatInt(message.RoomID, 10)
	return event, nil
}
func (s *Store) AuthorizeSocket(ctx context.Context, roomID, userID string) error {
	room, err := strconv.ParseInt(roomID, 10, 64)
	if err != nil {
		return errors.New("room_id inválido")
	}
	user, err := strconv.ParseInt(userID, 10, 64)
	if err != nil {
		return errors.New("X-User-ID inválido")
	}
	member, err := s.IsMember(ctx, room, user)
	if err != nil {
		return err
	}
	if !member {
		return errors.New("el usuario debe unirse a la sala")
	}
	return nil
}

func (s *Store) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id,username,display_name,bio FROM users WHERE deleted_at IS NULL ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Username, &u.DisplayName, &u.Bio); err != nil {
			return nil, err
		}
		out = append(out, u)
	}
	return out, rows.Err()
}
func (s *Store) GetUser(ctx context.Context, userID int64) (User, error) {
	var user User
	err := s.db.QueryRowContext(ctx, "SELECT id,username,display_name,bio FROM users WHERE id=$1 AND deleted_at IS NULL", userID).Scan(&user.ID, &user.Username, &user.DisplayName, &user.Bio)
	if errors.Is(err, sql.ErrNoRows) {
		return User{}, errors.New("usuario no encontrado")
	}
	return user, err
}
func (s *Store) ListPosts(ctx context.Context) ([]Post, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT p.id,p.user_id,p.content,COUNT(l.user_id) FROM posts p LEFT JOIN likes l ON l.post_id=p.id WHERE p.deleted_at IS NULL GROUP BY p.id ORDER BY p.id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Post
	for rows.Next() {
		var p Post
		if err := rows.Scan(&p.ID, &p.UserID, &p.Content, &p.Likes); err != nil {
			return nil, err
		}
		out = append(out, p)
	}
	return out, rows.Err()
}
func (s *Store) ListRooms(ctx context.Context) ([]Room, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id,name FROM rooms ORDER BY id DESC")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Room
	for rows.Next() {
		var room Room
		if err := rows.Scan(&room.ID, &room.Name); err != nil {
			return nil, err
		}
		out = append(out, room)
	}
	return out, rows.Err()
}
func (s *Store) ListMessages(ctx context.Context, roomID int64) ([]Message, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id,room_id,user_id,content FROM messages WHERE room_id=$1 ORDER BY id", roomID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Message
	for rows.Next() {
		var m Message
		if err := rows.Scan(&m.ID, &m.RoomID, &m.UserID, &m.Content); err != nil {
			return nil, err
		}
		out = append(out, m)
	}
	return out, rows.Err()
}
func (s *Store) ListComments(ctx context.Context, postID int64) ([]Comment, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT id,user_id,post_id,content FROM comments WHERE post_id=$1 ORDER BY id", postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var out []Comment
	for rows.Next() {
		var value Comment
		if err := rows.Scan(&value.ID, &value.UserID, &value.PostID, &value.Content); err != nil {
			return nil, err
		}
		out = append(out, value)
	}
	return out, rows.Err()
}

func Routes(store *Store, require func(http.Handler) http.Handler) chi.Router {
	r := chi.NewRouter()
	r.Post("/users", store.createUser)
	r.Get("/users", store.listUsers)
	r.Get("/users/{id}", store.getUser)
	r.Get("/posts", store.listPosts)
	r.Get("/rooms", store.listRooms)
	r.Get("/rooms/{id}/messages", store.listMessages)
	r.Get("/rooms/{id}/members", store.listMembers)
	r.Group(func(protected chi.Router) {
		protected.Use(require)
		protected.Post("/posts", store.createPost)
		protected.Post("/posts/{id}/likes", store.likePost)
		protected.Delete("/posts/{id}/likes", store.unlikePost)
		protected.Post("/posts/{id}/comments", store.commentPost)
		protected.Get("/posts/{id}/comments", store.listComments)
		protected.Post("/rooms", store.createRoom)
		protected.Post("/rooms/{id}/messages", store.createMessage)
		protected.Post("/rooms/{id}/join", store.joinRoom)
		protected.Delete("/rooms/{id}/join", store.leaveRoom)
		protected.Post("/users/{id}/follow", store.follow)
		protected.Delete("/users/{id}/follow", store.unfollow)
	})
	return r
}

func decode(r *http.Request, target any) error { return json.NewDecoder(r.Body).Decode(target) }
func write(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]any{"status": status, "data": data})
}
func writeError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]any{"status": status, "error": message})
}
func fail(w http.ResponseWriter, err error) {
	status := http.StatusBadRequest
	message := err.Error()
	if strings.Contains(message, "no pertenece") || strings.Contains(message, "debe unirse") {
		status = http.StatusForbidden
	}
	if strings.Contains(message, "no encontrada") || strings.Contains(message, "no encontrado") {
		status = http.StatusNotFound
	}
	if strings.Contains(message, "duplicate key") || strings.Contains(message, "ya existe") {
		status = http.StatusConflict
	}
	writeError(w, status, message)
}
func id(r *http.Request) (int64, error) { return strconv.ParseInt(chi.URLParam(r, "id"), 10, 64) }
func (s *Store) createUser(w http.ResponseWriter, r *http.Request) {
	var req userRequest
	if err := decode(r, &req); err != nil {
		fail(w, err)
		return
	}
	value, err := s.CreateUser(r.Context(), req)
	if err != nil {
		fail(w, err)
		return
	}
	write(w, http.StatusCreated, value)
}
func (s *Store) createPost(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.Header.Get("X-User-ID"), 10, 64)
	if err != nil {
		fail(w, errors.New("X-User-ID inválido"))
		return
	}
	var req postRequest
	if err = decode(r, &req); err != nil {
		fail(w, err)
		return
	}
	value, err := s.CreatePost(r.Context(), userID, req)
	if err != nil {
		fail(w, err)
		return
	}
	write(w, http.StatusCreated, value)
}
func (s *Store) createRoom(w http.ResponseWriter, r *http.Request) {
	var req roomRequest
	if err := decode(r, &req); err != nil {
		fail(w, err)
		return
	}
	value, err := s.CreateRoom(r.Context(), req)
	if err != nil {
		fail(w, err)
		return
	}
	write(w, http.StatusCreated, value)
}
func (s *Store) createMessage(w http.ResponseWriter, r *http.Request) {
	roomID, err := id(r)
	if err != nil {
		fail(w, err)
		return
	}
	var req messageRequest
	if err = decode(r, &req); err != nil {
		fail(w, err)
		return
	}
	value, err := s.CreateMessage(r.Context(), roomID, req)
	if err != nil {
		fail(w, err)
		return
	}
	write(w, http.StatusCreated, value)
}
func (s *Store) follow(w http.ResponseWriter, r *http.Request) {
	followerID, err := strconv.ParseInt(r.Header.Get("X-User-ID"), 10, 64)
	if err != nil {
		fail(w, errors.New("X-User-ID inválido"))
		return
	}
	followingID, err := id(r)
	if err != nil {
		fail(w, err)
		return
	}
	if followerID == followingID {
		fail(w, errors.New("no puedes seguirte a ti mismo"))
		return
	}
	if _, err = s.db.ExecContext(r.Context(), "INSERT INTO follows(follower_id,following_id) VALUES($1,$2)", followerID, followingID); err != nil {
		fail(w, err)
		return
	}
	write(w, http.StatusCreated, map[string]bool{"following": true})
}

func (s *Store) listUsers(w http.ResponseWriter, r *http.Request) {
	value, err := s.ListUsers(r.Context())
	if err != nil {
		fail(w, err)
		return
	}
	write(w, http.StatusOK, value)
}
func (s *Store) getUser(w http.ResponseWriter, r *http.Request) {
	userID, err := id(r)
	if err != nil {
		fail(w, err)
		return
	}
	value, err := s.GetUser(r.Context(), userID)
	if err != nil {
		fail(w, err)
		return
	}
	write(w, http.StatusOK, value)
}
func (s *Store) listPosts(w http.ResponseWriter, r *http.Request) {
	value, err := s.ListPosts(r.Context())
	if err != nil {
		fail(w, err)
		return
	}
	write(w, http.StatusOK, value)
}
func (s *Store) listRooms(w http.ResponseWriter, r *http.Request) {
	value, err := s.ListRooms(r.Context())
	if err != nil {
		fail(w, err)
		return
	}
	write(w, http.StatusOK, value)
}
func (s *Store) listMessages(w http.ResponseWriter, r *http.Request) {
	roomID, err := id(r)
	if err != nil {
		fail(w, err)
		return
	}
	value, err := s.ListMessages(r.Context(), roomID)
	if err != nil {
		fail(w, err)
		return
	}
	write(w, http.StatusOK, value)
}
func (s *Store) likePost(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.Header.Get("X-User-ID"), 10, 64)
	if err != nil {
		fail(w, errors.New("X-User-ID inválido"))
		return
	}
	postID, err := id(r)
	if err != nil {
		fail(w, err)
		return
	}
	if _, err = s.db.ExecContext(r.Context(), "INSERT INTO likes(user_id,post_id) VALUES($1,$2)", userID, postID); err != nil {
		fail(w, err)
		return
	}
	write(w, http.StatusCreated, map[string]bool{"liked": true})
}
func (s *Store) unlikePost(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.Header.Get("X-User-ID"), 10, 64)
	if err != nil {
		fail(w, errors.New("X-User-ID inválido"))
		return
	}
	postID, err := id(r)
	if err != nil {
		fail(w, err)
		return
	}
	if _, err = s.db.ExecContext(r.Context(), "DELETE FROM likes WHERE user_id=$1 AND post_id=$2", userID, postID); err != nil {
		fail(w, err)
		return
	}
	write(w, http.StatusOK, map[string]bool{"liked": false})
}
func (s *Store) commentPost(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.Header.Get("X-User-ID"), 10, 64)
	if err != nil {
		fail(w, errors.New("X-User-ID inválido"))
		return
	}
	postID, err := id(r)
	if err != nil {
		fail(w, err)
		return
	}
	var req commentRequest
	if err = decode(r, &req); err != nil {
		fail(w, err)
		return
	}
	if strings.TrimSpace(req.Content) == "" {
		fail(w, errors.New("content es obligatorio"))
		return
	}
	if len([]rune(req.Content)) > 500 {
		fail(w, errors.New("content no puede superar 500 caracteres"))
		return
	}
	var comment Comment
	err = s.db.QueryRowContext(r.Context(), "INSERT INTO comments(user_id,post_id,content) VALUES($1,$2,$3) RETURNING id,user_id,post_id,content", userID, postID, strings.TrimSpace(req.Content)).Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Content)
	if err != nil {
		fail(w, err)
		return
	}
	write(w, http.StatusCreated, comment)
}
func (s *Store) listComments(w http.ResponseWriter, r *http.Request) {
	postID, err := id(r)
	if err != nil {
		fail(w, err)
		return
	}
	value, err := s.ListComments(r.Context(), postID)
	if err != nil {
		fail(w, err)
		return
	}
	write(w, http.StatusOK, value)
}
func (s *Store) unfollow(w http.ResponseWriter, r *http.Request) {
	followerID, err := strconv.ParseInt(r.Header.Get("X-User-ID"), 10, 64)
	if err != nil {
		fail(w, errors.New("X-User-ID inválido"))
		return
	}
	followingID, err := id(r)
	if err != nil {
		fail(w, err)
		return
	}
	if _, err = s.db.ExecContext(r.Context(), "DELETE FROM follows WHERE follower_id=$1 AND following_id=$2", followerID, followingID); err != nil {
		fail(w, err)
		return
	}
	write(w, http.StatusOK, map[string]bool{"following": false})
}

func (s *Store) joinRoom(w http.ResponseWriter, r *http.Request) {
	userID, err := currentUser(r)
	if err != nil {
		fail(w, err)
		return
	}
	roomID, err := id(r)
	if err != nil {
		fail(w, err)
		return
	}
	if err = s.JoinRoom(r.Context(), roomID, userID); err != nil {
		fail(w, err)
		return
	}
	write(w, http.StatusCreated, member{UserID: userID, RoomID: roomID})
}
func (s *Store) leaveRoom(w http.ResponseWriter, r *http.Request) {
	userID, err := currentUser(r)
	if err != nil {
		fail(w, err)
		return
	}
	roomID, err := id(r)
	if err != nil {
		fail(w, err)
		return
	}
	if err = s.LeaveRoom(r.Context(), roomID, userID); err != nil {
		fail(w, err)
		return
	}
	write(w, http.StatusOK, map[string]bool{"member": false})
}
func (s *Store) listMembers(w http.ResponseWriter, r *http.Request) {
	roomID, err := id(r)
	if err != nil {
		fail(w, err)
		return
	}
	rows, err := s.db.QueryContext(r.Context(), "SELECT user_id,room_id FROM room_members WHERE room_id=$1 ORDER BY user_id", roomID)
	if err != nil {
		fail(w, err)
		return
	}
	defer rows.Close()
	var out []member
	for rows.Next() {
		var value member
		if err = rows.Scan(&value.UserID, &value.RoomID); err != nil {
			fail(w, err)
			return
		}
		out = append(out, value)
	}
	write(w, http.StatusOK, out)
}
func currentUser(r *http.Request) (int64, error) {
	value, err := strconv.ParseInt(r.Header.Get("X-User-ID"), 10, 64)
	if err != nil || value <= 0 {
		return 0, errors.New("X-User-ID inválido")
	}
	return value, nil
}

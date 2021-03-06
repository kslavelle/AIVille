package router

// CreateGameModel is the expected body when creating
// a new game
type CreateGameModel struct {
	Name string `json:"name"`
}

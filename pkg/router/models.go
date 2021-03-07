package router

// CreateGameModel is the expected body when creating
// a new game
type CreateGameModel struct {
	Name string `json:"name"`
}

// CreateResourcesModel is the expected body when adding
// a resource
type CreateResourceModel struct {
	Name string `json: "gameid"`
}

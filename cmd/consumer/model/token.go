package model

// FirebaseMessagingToken holds data related to Firebase high priority messaging token.
// This token is used when pushing notifications to clients.
type FirebaseMessagingToken struct {
	Token       string `json:"token"`
	SuperheroID string `json:"superheroID"`
	UpdatedAt   string `json:"updatedAt"`
}

package model

// FirebaseMessagingToken holds data related to Firebase high priority messaging token.
// This token is used when pushing notifications to clients.
type FirebaseMessagingToken struct {
	Token       string `db:"firebase_token"`
	SuperheroID string `db:"id"`
	UpdatedAt   string `db:"updated_at"`
}

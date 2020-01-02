package db

import (
	"github.com/consumer-firebase-token/internal/db/model"
)

// UpdateFirebaseToken updates existing Firebase Messaging Token.
func(db *DB) UpdateFirebaseToken (f model.FirebaseMessagingToken) error {
	_, err := db.stmtUpdateFirebaseToken.Exec(
		f.SuperheroID,
		f.Token,
		f.UpdatedAt,
	)
	if err != nil {
		return err
	}

	return nil
}

package reader

import (
	"context"
	"encoding/json"
	"fmt"

	chm "github.com/consumer-firebase-token/internal/cache/model"
	"github.com/consumer-firebase-token/internal/consumer/model"
	dbm "github.com/consumer-firebase-token/internal/db/model"
)

// Read consumes the Kafka topic and updates the Firebase Messaging Token in DB and Cache.
func (r *Reader) Read() error {
	ctx := context.Background()

	for {
		fmt.Print("before FetchMessage")
		m, err := r.Consumer.Consumer.FetchMessage(ctx)
		fmt.Print("after FetchMessage")
		if err != nil {
			err = r.Consumer.Consumer.Close()
			if err != nil {
				return err
			}

			return err
		}

		fmt.Printf(
			"message at topic/partition/offset \n%v/\n%v/\n%v: \n%s = \n%s\n",
			m.Topic,
			m.Partition,
			m.Offset,
			string(m.Key),
			string(m.Value),
		)

		var f model.FirebaseMessagingToken
		if err := json.Unmarshal(m.Value, &f); err != nil {
			_ = r.Consumer.Consumer.Close()
			if err != nil {
				err = r.Consumer.Consumer.Close()
				if err != nil {
					return err
				}

				return err
			}
		}

		err = r.DB.UpdateFirebaseToken(dbm.FirebaseMessagingToken{
			SuperheroID: f.SuperheroID,
			Token:       f.Token,
			UpdatedAt:   f.UpdatedAt,
		}, )
		if err != nil {
			err = r.Consumer.Consumer.Close()
			if err != nil {
				return err
			}

			return err
		}

		err = r.Cache.SetToken(chm.FirebaseMessagingToken{
			SuperheroID: f.SuperheroID,
			Token:       f.Token,
			UpdatedAt:   f.UpdatedAt,
		}, )
		if err != nil {
			err = r.Consumer.Consumer.Close()
			if err != nil {
				return err
			}

			return err
		}

		err = r.Consumer.Consumer.CommitMessages(ctx, m)
		if err != nil {
			err = r.Consumer.Consumer.Close()
			if err != nil {
				return err
			}

			return err
		}
	}
}

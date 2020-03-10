/*
  Copyright (C) 2019 - 2020 MWSOFT
  This program is free software: you can redistribute it and/or modify
  it under the terms of the GNU General Public License as published by
  the Free Software Foundation, either version 3 of the License, or
  (at your option) any later version.
  This program is distributed in the hope that it will be useful,
  but WITHOUT ANY WARRANTY; without even the implied warranty of
  MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
  GNU General Public License for more details.
  You should have received a copy of the GNU General Public License
  along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package reader

import (
	"context"
	"encoding/json"
	"fmt"

	chm "github.com/superhero-match/consumer-firebase-token/internal/cache/model"
	"github.com/superhero-match/consumer-firebase-token/internal/consumer/model"
	dbm "github.com/superhero-match/consumer-firebase-token/internal/db/model"
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

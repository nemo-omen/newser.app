package subscription

import (
	"newser.app/internal/domain/entity"
)

type SubscriptionAggregate struct {
	user *entity.Person
	feed *entity.Feed
}

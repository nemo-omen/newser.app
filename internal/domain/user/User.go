package user

import (
	"newser.app/internal/domain/shared"
	"newser.app/internal/domain/subscription"
)

type User struct {
	user          shared.Person
	subscriptions []*subscription.Subscription
}

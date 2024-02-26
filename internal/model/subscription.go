package model

type Subscription struct{}

type SubscriptionGorm struct{}

func (SubscriptionGorm) TableName() string {
	return "categories"
}

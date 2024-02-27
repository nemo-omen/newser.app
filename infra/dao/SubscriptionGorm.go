package dao

type SubscriptionGorm struct{}

func (SubscriptionGorm) TableName() string {
	return "categories"
}

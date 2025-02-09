package handler

import (
	"birthdayReminder/internal/repository/user"
)

//go:generate mockgen -source=contract.go -destination=mocks/mockRepo.go
type UserRepository interface {
	CreateUser(user *user.User, hashedPassword []byte) error
	GetUserByEmail(email string) (*user.User, error)
	GetAvailableUsersForSubscription(userID int) ([]user.User, error)
	GetUsersWithBirthdayTomorrow() ([]user.User, error)
	GetSubscribers(userID int) ([]user.User, error)
}

type SubscriptionRepository interface {
	CreateSubscription(userID int, relatedUserID int) error
	UnsubscribeUser(userID int, relatedUserID int) error
}

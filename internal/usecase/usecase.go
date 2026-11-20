package usecase

import "log/slog"

type UseCase struct {
	log              *slog.Logger
	subscriptionRepo subscriptionRepo
}

type Dependencies struct {
	Log              *slog.Logger
	SubscriptionRepo subscriptionRepo
}

type subscriptionRepo any

func NewUseCase(d *Dependencies) *UseCase {
	return &UseCase{
		log:              d.Log,
		subscriptionRepo: d.SubscriptionRepo,
	}
}

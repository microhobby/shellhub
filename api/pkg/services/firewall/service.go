package firewall

import (
	"context"

	"github.com/shellhub-io/shellhub/api/pkg/store"
	"github.com/shellhub-io/shellhub/pkg/models"
)

type Service interface {
	CreateRule(ctx context.Context, rule models.FirewallRule) error
	ListRules(ctx context.Context, perPage, page int) ([]models.FirewallRule, int, error)
	GetRule(ctx context.Context, id string) (*models.FirewallRule, error)
	DeleteRule(ctx context.Context, id string) error
}

type service struct {
	store store.Store
}

func NewService(store store.Store) Service {
	return &service{store}
}

func (s *service) CreateRule(ctx context.Context, rule models.FirewallRule) error {
	return s.store.CreateFirewallRule(ctx, rule)
}

func (s *service) ListRules(ctx context.Context, perPage, page int) ([]models.FirewallRule, int, error) {
	return s.store.ListFirewallRules(ctx, perPage, page)
}

func (s *service) GetRule(ctx context.Context, id string) (*models.FirewallRule, error) {
	return s.store.GetFirewallRule(ctx, id)
}

func (s *service) DeleteRule(ctx context.Context, id string) error {
	return s.store.DeleteFirewallRule(ctx, id)
}

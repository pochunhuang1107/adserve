package campaign

import "errors"

type Campaign struct {
	name        string
	budgetCents int
	active      bool
}

func NewCampaign(name string, budgetCents int) (*Campaign, error) {
	if name == "" {
		return nil, ErrInvalidName
	}

	if budgetCents <= 0 {
		return nil, ErrInvalidBudget
	}

	return &Campaign{
		name:        name,
		budgetCents: budgetCents,
		active:      true,
	}, nil
}

func (c Campaign) Name() string {
	return c.name
}

func (c Campaign) BudgetCents() int {
	return c.budgetCents
}

func (c Campaign) IsActive() bool {
	return c.active
}

func (c Campaign) CanServe() bool {
	return c.active && c.budgetCents > 0
}

func (c *Campaign) Deactivate() {
	c.active = false
}

var ErrInvalidBudget = errors.New("campaign budget must be greater than zero")
var ErrInvalidName = errors.New("campaign name must not be an empty string")

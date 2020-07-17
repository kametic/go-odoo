package odoo

import (
	"fmt"
)

// LoyaltyReward represents loyalty.reward model.
type LoyaltyReward struct {
	LastUpdate        *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omptempty"`
	Discount          *Float     `xmlrpc:"discount,omptempty"`
	DiscountProductId *Many2One  `xmlrpc:"discount_product_id,omptempty"`
	DisplayName       *String    `xmlrpc:"display_name,omptempty"`
	GiftProductId     *Many2One  `xmlrpc:"gift_product_id,omptempty"`
	Id                *Int       `xmlrpc:"id,omptempty"`
	LoyaltyProgramId  *Many2One  `xmlrpc:"loyalty_program_id,omptempty"`
	MinimumPoints     *Float     `xmlrpc:"minimum_points,omptempty"`
	Name              *String    `xmlrpc:"name,omptempty"`
	PointCost         *Float     `xmlrpc:"point_cost,omptempty"`
	PointProductId    *Many2One  `xmlrpc:"point_product_id,omptempty"`
	RewardType        *Selection `xmlrpc:"reward_type,omptempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// LoyaltyRewards represents array of loyalty.reward model.
type LoyaltyRewards []LoyaltyReward

// LoyaltyRewardModel is the odoo model name.
const LoyaltyRewardModel = "loyalty.reward"

// Many2One convert LoyaltyReward to *Many2One.
func (lr *LoyaltyReward) Many2One() *Many2One {
	return NewMany2One(lr.Id.Get(), "")
}

// CreateLoyaltyReward creates a new loyalty.reward model and returns its id.
func (c *Client) CreateLoyaltyReward(lr *LoyaltyReward) (int64, error) {
	return c.Create(LoyaltyRewardModel, lr)
}

// UpdateLoyaltyReward updates an existing loyalty.reward record.
func (c *Client) UpdateLoyaltyReward(lr *LoyaltyReward) error {
	return c.UpdateLoyaltyRewards([]int64{lr.Id.Get()}, lr)
}

// UpdateLoyaltyRewards updates existing loyalty.reward records.
// All records (represented by ids) will be updated by lr values.
func (c *Client) UpdateLoyaltyRewards(ids []int64, lr *LoyaltyReward) error {
	return c.Update(LoyaltyRewardModel, ids, lr)
}

// DeleteLoyaltyReward deletes an existing loyalty.reward record.
func (c *Client) DeleteLoyaltyReward(id int64) error {
	return c.DeleteLoyaltyRewards([]int64{id})
}

// DeleteLoyaltyRewards deletes existing loyalty.reward records.
func (c *Client) DeleteLoyaltyRewards(ids []int64) error {
	return c.Delete(LoyaltyRewardModel, ids)
}

// GetLoyaltyReward gets loyalty.reward existing record.
func (c *Client) GetLoyaltyReward(id int64) (*LoyaltyReward, error) {
	lrs, err := c.GetLoyaltyRewards([]int64{id})
	if err != nil {
		return nil, err
	}
	if lrs != nil && len(*lrs) > 0 {
		return &((*lrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of loyalty.reward not found", id)
}

// GetLoyaltyRewards gets loyalty.reward existing records.
func (c *Client) GetLoyaltyRewards(ids []int64) (*LoyaltyRewards, error) {
	lrs := &LoyaltyRewards{}
	if err := c.Read(LoyaltyRewardModel, ids, nil, lrs); err != nil {
		return nil, err
	}
	return lrs, nil
}

// FindLoyaltyReward finds loyalty.reward record by querying it with criteria.
func (c *Client) FindLoyaltyReward(criteria *Criteria) (*LoyaltyReward, error) {
	lrs := &LoyaltyRewards{}
	if err := c.SearchRead(LoyaltyRewardModel, criteria, NewOptions().Limit(1), lrs); err != nil {
		return nil, err
	}
	if lrs != nil && len(*lrs) > 0 {
		return &((*lrs)[0]), nil
	}
	return nil, fmt.Errorf("loyalty.reward was not found")
}

// FindLoyaltyRewards finds loyalty.reward records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLoyaltyRewards(criteria *Criteria, options *Options) (*LoyaltyRewards, error) {
	lrs := &LoyaltyRewards{}
	if err := c.SearchRead(LoyaltyRewardModel, criteria, options, lrs); err != nil {
		return nil, err
	}
	return lrs, nil
}

// FindLoyaltyRewardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLoyaltyRewardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(LoyaltyRewardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindLoyaltyRewardId finds record id by querying it with criteria.
func (c *Client) FindLoyaltyRewardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LoyaltyRewardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("loyalty.reward was not found")
}

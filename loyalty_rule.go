package odoo

import (
	"fmt"
)

// LoyaltyRule represents loyalty.rule model.
type LoyaltyRule struct {
	LastUpdate       *Time      `xmlrpc:"__last_update,omptempty"`
	CategoryId       *Many2One  `xmlrpc:"category_id,omptempty"`
	CreateDate       *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omptempty"`
	Cumulative       *Bool      `xmlrpc:"cumulative,omptempty"`
	DisplayName      *String    `xmlrpc:"display_name,omptempty"`
	Id               *Int       `xmlrpc:"id,omptempty"`
	LoyaltyProgramId *Many2One  `xmlrpc:"loyalty_program_id,omptempty"`
	Name             *String    `xmlrpc:"name,omptempty"`
	PpCurrency       *Float     `xmlrpc:"pp_currency,omptempty"`
	PpProduct        *Float     `xmlrpc:"pp_product,omptempty"`
	ProductId        *Many2One  `xmlrpc:"product_id,omptempty"`
	RuleType         *Selection `xmlrpc:"rule_type,omptempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// LoyaltyRules represents array of loyalty.rule model.
type LoyaltyRules []LoyaltyRule

// LoyaltyRuleModel is the odoo model name.
const LoyaltyRuleModel = "loyalty.rule"

// Many2One convert LoyaltyRule to *Many2One.
func (lr *LoyaltyRule) Many2One() *Many2One {
	return NewMany2One(lr.Id.Get(), "")
}

// CreateLoyaltyRule creates a new loyalty.rule model and returns its id.
func (c *Client) CreateLoyaltyRule(lr *LoyaltyRule) (int64, error) {
	return c.Create(LoyaltyRuleModel, lr)
}

// UpdateLoyaltyRule updates an existing loyalty.rule record.
func (c *Client) UpdateLoyaltyRule(lr *LoyaltyRule) error {
	return c.UpdateLoyaltyRules([]int64{lr.Id.Get()}, lr)
}

// UpdateLoyaltyRules updates existing loyalty.rule records.
// All records (represented by ids) will be updated by lr values.
func (c *Client) UpdateLoyaltyRules(ids []int64, lr *LoyaltyRule) error {
	return c.Update(LoyaltyRuleModel, ids, lr)
}

// DeleteLoyaltyRule deletes an existing loyalty.rule record.
func (c *Client) DeleteLoyaltyRule(id int64) error {
	return c.DeleteLoyaltyRules([]int64{id})
}

// DeleteLoyaltyRules deletes existing loyalty.rule records.
func (c *Client) DeleteLoyaltyRules(ids []int64) error {
	return c.Delete(LoyaltyRuleModel, ids)
}

// GetLoyaltyRule gets loyalty.rule existing record.
func (c *Client) GetLoyaltyRule(id int64) (*LoyaltyRule, error) {
	lrs, err := c.GetLoyaltyRules([]int64{id})
	if err != nil {
		return nil, err
	}
	if lrs != nil && len(*lrs) > 0 {
		return &((*lrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of loyalty.rule not found", id)
}

// GetLoyaltyRules gets loyalty.rule existing records.
func (c *Client) GetLoyaltyRules(ids []int64) (*LoyaltyRules, error) {
	lrs := &LoyaltyRules{}
	if err := c.Read(LoyaltyRuleModel, ids, nil, lrs); err != nil {
		return nil, err
	}
	return lrs, nil
}

// FindLoyaltyRule finds loyalty.rule record by querying it with criteria.
func (c *Client) FindLoyaltyRule(criteria *Criteria) (*LoyaltyRule, error) {
	lrs := &LoyaltyRules{}
	if err := c.SearchRead(LoyaltyRuleModel, criteria, NewOptions().Limit(1), lrs); err != nil {
		return nil, err
	}
	if lrs != nil && len(*lrs) > 0 {
		return &((*lrs)[0]), nil
	}
	return nil, fmt.Errorf("loyalty.rule was not found")
}

// FindLoyaltyRules finds loyalty.rule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLoyaltyRules(criteria *Criteria, options *Options) (*LoyaltyRules, error) {
	lrs := &LoyaltyRules{}
	if err := c.SearchRead(LoyaltyRuleModel, criteria, options, lrs); err != nil {
		return nil, err
	}
	return lrs, nil
}

// FindLoyaltyRuleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLoyaltyRuleIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(LoyaltyRuleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindLoyaltyRuleId finds record id by querying it with criteria.
func (c *Client) FindLoyaltyRuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LoyaltyRuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("loyalty.rule was not found")
}

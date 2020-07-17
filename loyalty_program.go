package odoo

import (
	"fmt"
)

// LoyaltyProgram represents loyalty.program model.
type LoyaltyProgram struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	PpCurrency  *Float    `xmlrpc:"pp_currency,omptempty"`
	PpOrder     *Float    `xmlrpc:"pp_order,omptempty"`
	PpProduct   *Float    `xmlrpc:"pp_product,omptempty"`
	RewardIds   *Relation `xmlrpc:"reward_ids,omptempty"`
	Rounding    *Float    `xmlrpc:"rounding,omptempty"`
	RuleIds     *Relation `xmlrpc:"rule_ids,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// LoyaltyPrograms represents array of loyalty.program model.
type LoyaltyPrograms []LoyaltyProgram

// LoyaltyProgramModel is the odoo model name.
const LoyaltyProgramModel = "loyalty.program"

// Many2One convert LoyaltyProgram to *Many2One.
func (lp *LoyaltyProgram) Many2One() *Many2One {
	return NewMany2One(lp.Id.Get(), "")
}

// CreateLoyaltyProgram creates a new loyalty.program model and returns its id.
func (c *Client) CreateLoyaltyProgram(lp *LoyaltyProgram) (int64, error) {
	return c.Create(LoyaltyProgramModel, lp)
}

// UpdateLoyaltyProgram updates an existing loyalty.program record.
func (c *Client) UpdateLoyaltyProgram(lp *LoyaltyProgram) error {
	return c.UpdateLoyaltyPrograms([]int64{lp.Id.Get()}, lp)
}

// UpdateLoyaltyPrograms updates existing loyalty.program records.
// All records (represented by ids) will be updated by lp values.
func (c *Client) UpdateLoyaltyPrograms(ids []int64, lp *LoyaltyProgram) error {
	return c.Update(LoyaltyProgramModel, ids, lp)
}

// DeleteLoyaltyProgram deletes an existing loyalty.program record.
func (c *Client) DeleteLoyaltyProgram(id int64) error {
	return c.DeleteLoyaltyPrograms([]int64{id})
}

// DeleteLoyaltyPrograms deletes existing loyalty.program records.
func (c *Client) DeleteLoyaltyPrograms(ids []int64) error {
	return c.Delete(LoyaltyProgramModel, ids)
}

// GetLoyaltyProgram gets loyalty.program existing record.
func (c *Client) GetLoyaltyProgram(id int64) (*LoyaltyProgram, error) {
	lps, err := c.GetLoyaltyPrograms([]int64{id})
	if err != nil {
		return nil, err
	}
	if lps != nil && len(*lps) > 0 {
		return &((*lps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of loyalty.program not found", id)
}

// GetLoyaltyPrograms gets loyalty.program existing records.
func (c *Client) GetLoyaltyPrograms(ids []int64) (*LoyaltyPrograms, error) {
	lps := &LoyaltyPrograms{}
	if err := c.Read(LoyaltyProgramModel, ids, nil, lps); err != nil {
		return nil, err
	}
	return lps, nil
}

// FindLoyaltyProgram finds loyalty.program record by querying it with criteria.
func (c *Client) FindLoyaltyProgram(criteria *Criteria) (*LoyaltyProgram, error) {
	lps := &LoyaltyPrograms{}
	if err := c.SearchRead(LoyaltyProgramModel, criteria, NewOptions().Limit(1), lps); err != nil {
		return nil, err
	}
	if lps != nil && len(*lps) > 0 {
		return &((*lps)[0]), nil
	}
	return nil, fmt.Errorf("loyalty.program was not found")
}

// FindLoyaltyPrograms finds loyalty.program records by querying it
// and filtering it with criteria and options.
func (c *Client) FindLoyaltyPrograms(criteria *Criteria, options *Options) (*LoyaltyPrograms, error) {
	lps := &LoyaltyPrograms{}
	if err := c.SearchRead(LoyaltyProgramModel, criteria, options, lps); err != nil {
		return nil, err
	}
	return lps, nil
}

// FindLoyaltyProgramIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindLoyaltyProgramIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(LoyaltyProgramModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindLoyaltyProgramId finds record id by querying it with criteria.
func (c *Client) FindLoyaltyProgramId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(LoyaltyProgramModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("loyalty.program was not found")
}

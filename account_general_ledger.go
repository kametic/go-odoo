package odoo

import (
	"fmt"
)

// AccountGeneralLedger represents account.general.ledger model.
type AccountGeneralLedger struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// AccountGeneralLedgers represents array of account.general.ledger model.
type AccountGeneralLedgers []AccountGeneralLedger

// AccountGeneralLedgerModel is the odoo model name.
const AccountGeneralLedgerModel = "account.general.ledger"

// Many2One convert AccountGeneralLedger to *Many2One.
func (agl *AccountGeneralLedger) Many2One() *Many2One {
	return NewMany2One(agl.Id.Get(), "")
}

// CreateAccountGeneralLedger creates a new account.general.ledger model and returns its id.
func (c *Client) CreateAccountGeneralLedger(agl *AccountGeneralLedger) (int64, error) {
	return c.Create(AccountGeneralLedgerModel, agl)
}

// UpdateAccountGeneralLedger updates an existing account.general.ledger record.
func (c *Client) UpdateAccountGeneralLedger(agl *AccountGeneralLedger) error {
	return c.UpdateAccountGeneralLedgers([]int64{agl.Id.Get()}, agl)
}

// UpdateAccountGeneralLedgers updates existing account.general.ledger records.
// All records (represented by ids) will be updated by agl values.
func (c *Client) UpdateAccountGeneralLedgers(ids []int64, agl *AccountGeneralLedger) error {
	return c.Update(AccountGeneralLedgerModel, ids, agl)
}

// DeleteAccountGeneralLedger deletes an existing account.general.ledger record.
func (c *Client) DeleteAccountGeneralLedger(id int64) error {
	return c.DeleteAccountGeneralLedgers([]int64{id})
}

// DeleteAccountGeneralLedgers deletes existing account.general.ledger records.
func (c *Client) DeleteAccountGeneralLedgers(ids []int64) error {
	return c.Delete(AccountGeneralLedgerModel, ids)
}

// GetAccountGeneralLedger gets account.general.ledger existing record.
func (c *Client) GetAccountGeneralLedger(id int64) (*AccountGeneralLedger, error) {
	agls, err := c.GetAccountGeneralLedgers([]int64{id})
	if err != nil {
		return nil, err
	}
	if agls != nil && len(*agls) > 0 {
		return &((*agls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.general.ledger not found", id)
}

// GetAccountGeneralLedgers gets account.general.ledger existing records.
func (c *Client) GetAccountGeneralLedgers(ids []int64) (*AccountGeneralLedgers, error) {
	agls := &AccountGeneralLedgers{}
	if err := c.Read(AccountGeneralLedgerModel, ids, nil, agls); err != nil {
		return nil, err
	}
	return agls, nil
}

// FindAccountGeneralLedger finds account.general.ledger record by querying it with criteria.
func (c *Client) FindAccountGeneralLedger(criteria *Criteria) (*AccountGeneralLedger, error) {
	agls := &AccountGeneralLedgers{}
	if err := c.SearchRead(AccountGeneralLedgerModel, criteria, NewOptions().Limit(1), agls); err != nil {
		return nil, err
	}
	if agls != nil && len(*agls) > 0 {
		return &((*agls)[0]), nil
	}
	return nil, fmt.Errorf("account.general.ledger was not found")
}

// FindAccountGeneralLedgers finds account.general.ledger records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountGeneralLedgers(criteria *Criteria, options *Options) (*AccountGeneralLedgers, error) {
	agls := &AccountGeneralLedgers{}
	if err := c.SearchRead(AccountGeneralLedgerModel, criteria, options, agls); err != nil {
		return nil, err
	}
	return agls, nil
}

// FindAccountGeneralLedgerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountGeneralLedgerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountGeneralLedgerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountGeneralLedgerId finds record id by querying it with criteria.
func (c *Client) FindAccountGeneralLedgerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountGeneralLedgerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.general.ledger was not found")
}

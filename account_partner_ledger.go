package odoo

import (
	"fmt"
)

// AccountPartnerLedger represents account.partner.ledger model.
type AccountPartnerLedger struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// AccountPartnerLedgers represents array of account.partner.ledger model.
type AccountPartnerLedgers []AccountPartnerLedger

// AccountPartnerLedgerModel is the odoo model name.
const AccountPartnerLedgerModel = "account.partner.ledger"

// Many2One convert AccountPartnerLedger to *Many2One.
func (apl *AccountPartnerLedger) Many2One() *Many2One {
	return NewMany2One(apl.Id.Get(), "")
}

// CreateAccountPartnerLedger creates a new account.partner.ledger model and returns its id.
func (c *Client) CreateAccountPartnerLedger(apl *AccountPartnerLedger) (int64, error) {
	return c.Create(AccountPartnerLedgerModel, apl)
}

// UpdateAccountPartnerLedger updates an existing account.partner.ledger record.
func (c *Client) UpdateAccountPartnerLedger(apl *AccountPartnerLedger) error {
	return c.UpdateAccountPartnerLedgers([]int64{apl.Id.Get()}, apl)
}

// UpdateAccountPartnerLedgers updates existing account.partner.ledger records.
// All records (represented by ids) will be updated by apl values.
func (c *Client) UpdateAccountPartnerLedgers(ids []int64, apl *AccountPartnerLedger) error {
	return c.Update(AccountPartnerLedgerModel, ids, apl)
}

// DeleteAccountPartnerLedger deletes an existing account.partner.ledger record.
func (c *Client) DeleteAccountPartnerLedger(id int64) error {
	return c.DeleteAccountPartnerLedgers([]int64{id})
}

// DeleteAccountPartnerLedgers deletes existing account.partner.ledger records.
func (c *Client) DeleteAccountPartnerLedgers(ids []int64) error {
	return c.Delete(AccountPartnerLedgerModel, ids)
}

// GetAccountPartnerLedger gets account.partner.ledger existing record.
func (c *Client) GetAccountPartnerLedger(id int64) (*AccountPartnerLedger, error) {
	apls, err := c.GetAccountPartnerLedgers([]int64{id})
	if err != nil {
		return nil, err
	}
	if apls != nil && len(*apls) > 0 {
		return &((*apls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.partner.ledger not found", id)
}

// GetAccountPartnerLedgers gets account.partner.ledger existing records.
func (c *Client) GetAccountPartnerLedgers(ids []int64) (*AccountPartnerLedgers, error) {
	apls := &AccountPartnerLedgers{}
	if err := c.Read(AccountPartnerLedgerModel, ids, nil, apls); err != nil {
		return nil, err
	}
	return apls, nil
}

// FindAccountPartnerLedger finds account.partner.ledger record by querying it with criteria.
func (c *Client) FindAccountPartnerLedger(criteria *Criteria) (*AccountPartnerLedger, error) {
	apls := &AccountPartnerLedgers{}
	if err := c.SearchRead(AccountPartnerLedgerModel, criteria, NewOptions().Limit(1), apls); err != nil {
		return nil, err
	}
	if apls != nil && len(*apls) > 0 {
		return &((*apls)[0]), nil
	}
	return nil, fmt.Errorf("account.partner.ledger was not found")
}

// FindAccountPartnerLedgers finds account.partner.ledger records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPartnerLedgers(criteria *Criteria, options *Options) (*AccountPartnerLedgers, error) {
	apls := &AccountPartnerLedgers{}
	if err := c.SearchRead(AccountPartnerLedgerModel, criteria, options, apls); err != nil {
		return nil, err
	}
	return apls, nil
}

// FindAccountPartnerLedgerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPartnerLedgerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountPartnerLedgerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountPartnerLedgerId finds record id by querying it with criteria.
func (c *Client) FindAccountPartnerLedgerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountPartnerLedgerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.partner.ledger was not found")
}

package odoo

import (
	"fmt"
)

// AccountAgedPayable represents account.aged.payable model.
type AccountAgedPayable struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// AccountAgedPayables represents array of account.aged.payable model.
type AccountAgedPayables []AccountAgedPayable

// AccountAgedPayableModel is the odoo model name.
const AccountAgedPayableModel = "account.aged.payable"

// Many2One convert AccountAgedPayable to *Many2One.
func (aap *AccountAgedPayable) Many2One() *Many2One {
	return NewMany2One(aap.Id.Get(), "")
}

// CreateAccountAgedPayable creates a new account.aged.payable model and returns its id.
func (c *Client) CreateAccountAgedPayable(aap *AccountAgedPayable) (int64, error) {
	return c.Create(AccountAgedPayableModel, aap)
}

// UpdateAccountAgedPayable updates an existing account.aged.payable record.
func (c *Client) UpdateAccountAgedPayable(aap *AccountAgedPayable) error {
	return c.UpdateAccountAgedPayables([]int64{aap.Id.Get()}, aap)
}

// UpdateAccountAgedPayables updates existing account.aged.payable records.
// All records (represented by ids) will be updated by aap values.
func (c *Client) UpdateAccountAgedPayables(ids []int64, aap *AccountAgedPayable) error {
	return c.Update(AccountAgedPayableModel, ids, aap)
}

// DeleteAccountAgedPayable deletes an existing account.aged.payable record.
func (c *Client) DeleteAccountAgedPayable(id int64) error {
	return c.DeleteAccountAgedPayables([]int64{id})
}

// DeleteAccountAgedPayables deletes existing account.aged.payable records.
func (c *Client) DeleteAccountAgedPayables(ids []int64) error {
	return c.Delete(AccountAgedPayableModel, ids)
}

// GetAccountAgedPayable gets account.aged.payable existing record.
func (c *Client) GetAccountAgedPayable(id int64) (*AccountAgedPayable, error) {
	aaps, err := c.GetAccountAgedPayables([]int64{id})
	if err != nil {
		return nil, err
	}
	if aaps != nil && len(*aaps) > 0 {
		return &((*aaps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.aged.payable not found", id)
}

// GetAccountAgedPayables gets account.aged.payable existing records.
func (c *Client) GetAccountAgedPayables(ids []int64) (*AccountAgedPayables, error) {
	aaps := &AccountAgedPayables{}
	if err := c.Read(AccountAgedPayableModel, ids, nil, aaps); err != nil {
		return nil, err
	}
	return aaps, nil
}

// FindAccountAgedPayable finds account.aged.payable record by querying it with criteria.
func (c *Client) FindAccountAgedPayable(criteria *Criteria) (*AccountAgedPayable, error) {
	aaps := &AccountAgedPayables{}
	if err := c.SearchRead(AccountAgedPayableModel, criteria, NewOptions().Limit(1), aaps); err != nil {
		return nil, err
	}
	if aaps != nil && len(*aaps) > 0 {
		return &((*aaps)[0]), nil
	}
	return nil, fmt.Errorf("account.aged.payable was not found")
}

// FindAccountAgedPayables finds account.aged.payable records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAgedPayables(criteria *Criteria, options *Options) (*AccountAgedPayables, error) {
	aaps := &AccountAgedPayables{}
	if err := c.SearchRead(AccountAgedPayableModel, criteria, options, aaps); err != nil {
		return nil, err
	}
	return aaps, nil
}

// FindAccountAgedPayableIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAgedPayableIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountAgedPayableModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountAgedPayableId finds record id by querying it with criteria.
func (c *Client) FindAccountAgedPayableId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAgedPayableModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.aged.payable was not found")
}

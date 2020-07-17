package odoo

import (
	"fmt"
)

// AccountAgedPartner represents account.aged.partner model.
type AccountAgedPartner struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// AccountAgedPartners represents array of account.aged.partner model.
type AccountAgedPartners []AccountAgedPartner

// AccountAgedPartnerModel is the odoo model name.
const AccountAgedPartnerModel = "account.aged.partner"

// Many2One convert AccountAgedPartner to *Many2One.
func (aap *AccountAgedPartner) Many2One() *Many2One {
	return NewMany2One(aap.Id.Get(), "")
}

// CreateAccountAgedPartner creates a new account.aged.partner model and returns its id.
func (c *Client) CreateAccountAgedPartner(aap *AccountAgedPartner) (int64, error) {
	return c.Create(AccountAgedPartnerModel, aap)
}

// UpdateAccountAgedPartner updates an existing account.aged.partner record.
func (c *Client) UpdateAccountAgedPartner(aap *AccountAgedPartner) error {
	return c.UpdateAccountAgedPartners([]int64{aap.Id.Get()}, aap)
}

// UpdateAccountAgedPartners updates existing account.aged.partner records.
// All records (represented by ids) will be updated by aap values.
func (c *Client) UpdateAccountAgedPartners(ids []int64, aap *AccountAgedPartner) error {
	return c.Update(AccountAgedPartnerModel, ids, aap)
}

// DeleteAccountAgedPartner deletes an existing account.aged.partner record.
func (c *Client) DeleteAccountAgedPartner(id int64) error {
	return c.DeleteAccountAgedPartners([]int64{id})
}

// DeleteAccountAgedPartners deletes existing account.aged.partner records.
func (c *Client) DeleteAccountAgedPartners(ids []int64) error {
	return c.Delete(AccountAgedPartnerModel, ids)
}

// GetAccountAgedPartner gets account.aged.partner existing record.
func (c *Client) GetAccountAgedPartner(id int64) (*AccountAgedPartner, error) {
	aaps, err := c.GetAccountAgedPartners([]int64{id})
	if err != nil {
		return nil, err
	}
	if aaps != nil && len(*aaps) > 0 {
		return &((*aaps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.aged.partner not found", id)
}

// GetAccountAgedPartners gets account.aged.partner existing records.
func (c *Client) GetAccountAgedPartners(ids []int64) (*AccountAgedPartners, error) {
	aaps := &AccountAgedPartners{}
	if err := c.Read(AccountAgedPartnerModel, ids, nil, aaps); err != nil {
		return nil, err
	}
	return aaps, nil
}

// FindAccountAgedPartner finds account.aged.partner record by querying it with criteria.
func (c *Client) FindAccountAgedPartner(criteria *Criteria) (*AccountAgedPartner, error) {
	aaps := &AccountAgedPartners{}
	if err := c.SearchRead(AccountAgedPartnerModel, criteria, NewOptions().Limit(1), aaps); err != nil {
		return nil, err
	}
	if aaps != nil && len(*aaps) > 0 {
		return &((*aaps)[0]), nil
	}
	return nil, fmt.Errorf("account.aged.partner was not found")
}

// FindAccountAgedPartners finds account.aged.partner records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAgedPartners(criteria *Criteria, options *Options) (*AccountAgedPartners, error) {
	aaps := &AccountAgedPartners{}
	if err := c.SearchRead(AccountAgedPartnerModel, criteria, options, aaps); err != nil {
		return nil, err
	}
	return aaps, nil
}

// FindAccountAgedPartnerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAgedPartnerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountAgedPartnerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountAgedPartnerId finds record id by querying it with criteria.
func (c *Client) FindAccountAgedPartnerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAgedPartnerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.aged.partner was not found")
}

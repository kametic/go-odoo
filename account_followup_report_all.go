package odoo

import (
	"fmt"
)

// AccountFollowupReportAll represents account.followup.report.all model.
type AccountFollowupReportAll struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// AccountFollowupReportAlls represents array of account.followup.report.all model.
type AccountFollowupReportAlls []AccountFollowupReportAll

// AccountFollowupReportAllModel is the odoo model name.
const AccountFollowupReportAllModel = "account.followup.report.all"

// Many2One convert AccountFollowupReportAll to *Many2One.
func (afra *AccountFollowupReportAll) Many2One() *Many2One {
	return NewMany2One(afra.Id.Get(), "")
}

// CreateAccountFollowupReportAll creates a new account.followup.report.all model and returns its id.
func (c *Client) CreateAccountFollowupReportAll(afra *AccountFollowupReportAll) (int64, error) {
	return c.Create(AccountFollowupReportAllModel, afra)
}

// UpdateAccountFollowupReportAll updates an existing account.followup.report.all record.
func (c *Client) UpdateAccountFollowupReportAll(afra *AccountFollowupReportAll) error {
	return c.UpdateAccountFollowupReportAlls([]int64{afra.Id.Get()}, afra)
}

// UpdateAccountFollowupReportAlls updates existing account.followup.report.all records.
// All records (represented by ids) will be updated by afra values.
func (c *Client) UpdateAccountFollowupReportAlls(ids []int64, afra *AccountFollowupReportAll) error {
	return c.Update(AccountFollowupReportAllModel, ids, afra)
}

// DeleteAccountFollowupReportAll deletes an existing account.followup.report.all record.
func (c *Client) DeleteAccountFollowupReportAll(id int64) error {
	return c.DeleteAccountFollowupReportAlls([]int64{id})
}

// DeleteAccountFollowupReportAlls deletes existing account.followup.report.all records.
func (c *Client) DeleteAccountFollowupReportAlls(ids []int64) error {
	return c.Delete(AccountFollowupReportAllModel, ids)
}

// GetAccountFollowupReportAll gets account.followup.report.all existing record.
func (c *Client) GetAccountFollowupReportAll(id int64) (*AccountFollowupReportAll, error) {
	afras, err := c.GetAccountFollowupReportAlls([]int64{id})
	if err != nil {
		return nil, err
	}
	if afras != nil && len(*afras) > 0 {
		return &((*afras)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.followup.report.all not found", id)
}

// GetAccountFollowupReportAlls gets account.followup.report.all existing records.
func (c *Client) GetAccountFollowupReportAlls(ids []int64) (*AccountFollowupReportAlls, error) {
	afras := &AccountFollowupReportAlls{}
	if err := c.Read(AccountFollowupReportAllModel, ids, nil, afras); err != nil {
		return nil, err
	}
	return afras, nil
}

// FindAccountFollowupReportAll finds account.followup.report.all record by querying it with criteria.
func (c *Client) FindAccountFollowupReportAll(criteria *Criteria) (*AccountFollowupReportAll, error) {
	afras := &AccountFollowupReportAlls{}
	if err := c.SearchRead(AccountFollowupReportAllModel, criteria, NewOptions().Limit(1), afras); err != nil {
		return nil, err
	}
	if afras != nil && len(*afras) > 0 {
		return &((*afras)[0]), nil
	}
	return nil, fmt.Errorf("account.followup.report.all was not found")
}

// FindAccountFollowupReportAlls finds account.followup.report.all records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFollowupReportAlls(criteria *Criteria, options *Options) (*AccountFollowupReportAlls, error) {
	afras := &AccountFollowupReportAlls{}
	if err := c.SearchRead(AccountFollowupReportAllModel, criteria, options, afras); err != nil {
		return nil, err
	}
	return afras, nil
}

// FindAccountFollowupReportAllIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFollowupReportAllIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountFollowupReportAllModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountFollowupReportAllId finds record id by querying it with criteria.
func (c *Client) FindAccountFollowupReportAllId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountFollowupReportAllModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.followup.report.all was not found")
}

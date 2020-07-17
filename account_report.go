package odoo

import (
	"fmt"
)

// AccountReport represents account.report model.
type AccountReport struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// AccountReports represents array of account.report model.
type AccountReports []AccountReport

// AccountReportModel is the odoo model name.
const AccountReportModel = "account.report"

// Many2One convert AccountReport to *Many2One.
func (ar *AccountReport) Many2One() *Many2One {
	return NewMany2One(ar.Id.Get(), "")
}

// CreateAccountReport creates a new account.report model and returns its id.
func (c *Client) CreateAccountReport(ar *AccountReport) (int64, error) {
	return c.Create(AccountReportModel, ar)
}

// UpdateAccountReport updates an existing account.report record.
func (c *Client) UpdateAccountReport(ar *AccountReport) error {
	return c.UpdateAccountReports([]int64{ar.Id.Get()}, ar)
}

// UpdateAccountReports updates existing account.report records.
// All records (represented by ids) will be updated by ar values.
func (c *Client) UpdateAccountReports(ids []int64, ar *AccountReport) error {
	return c.Update(AccountReportModel, ids, ar)
}

// DeleteAccountReport deletes an existing account.report record.
func (c *Client) DeleteAccountReport(id int64) error {
	return c.DeleteAccountReports([]int64{id})
}

// DeleteAccountReports deletes existing account.report records.
func (c *Client) DeleteAccountReports(ids []int64) error {
	return c.Delete(AccountReportModel, ids)
}

// GetAccountReport gets account.report existing record.
func (c *Client) GetAccountReport(id int64) (*AccountReport, error) {
	ars, err := c.GetAccountReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if ars != nil && len(*ars) > 0 {
		return &((*ars)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.report not found", id)
}

// GetAccountReports gets account.report existing records.
func (c *Client) GetAccountReports(ids []int64) (*AccountReports, error) {
	ars := &AccountReports{}
	if err := c.Read(AccountReportModel, ids, nil, ars); err != nil {
		return nil, err
	}
	return ars, nil
}

// FindAccountReport finds account.report record by querying it with criteria.
func (c *Client) FindAccountReport(criteria *Criteria) (*AccountReport, error) {
	ars := &AccountReports{}
	if err := c.SearchRead(AccountReportModel, criteria, NewOptions().Limit(1), ars); err != nil {
		return nil, err
	}
	if ars != nil && len(*ars) > 0 {
		return &((*ars)[0]), nil
	}
	return nil, fmt.Errorf("account.report was not found")
}

// FindAccountReports finds account.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReports(criteria *Criteria, options *Options) (*AccountReports, error) {
	ars := &AccountReports{}
	if err := c.SearchRead(AccountReportModel, criteria, options, ars); err != nil {
		return nil, err
	}
	return ars, nil
}

// FindAccountReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountReportId finds record id by querying it with criteria.
func (c *Client) FindAccountReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.report was not found")
}

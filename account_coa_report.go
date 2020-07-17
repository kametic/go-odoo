package odoo

import (
	"fmt"
)

// AccountCoaReport represents account.coa.report model.
type AccountCoaReport struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// AccountCoaReports represents array of account.coa.report model.
type AccountCoaReports []AccountCoaReport

// AccountCoaReportModel is the odoo model name.
const AccountCoaReportModel = "account.coa.report"

// Many2One convert AccountCoaReport to *Many2One.
func (acr *AccountCoaReport) Many2One() *Many2One {
	return NewMany2One(acr.Id.Get(), "")
}

// CreateAccountCoaReport creates a new account.coa.report model and returns its id.
func (c *Client) CreateAccountCoaReport(acr *AccountCoaReport) (int64, error) {
	return c.Create(AccountCoaReportModel, acr)
}

// UpdateAccountCoaReport updates an existing account.coa.report record.
func (c *Client) UpdateAccountCoaReport(acr *AccountCoaReport) error {
	return c.UpdateAccountCoaReports([]int64{acr.Id.Get()}, acr)
}

// UpdateAccountCoaReports updates existing account.coa.report records.
// All records (represented by ids) will be updated by acr values.
func (c *Client) UpdateAccountCoaReports(ids []int64, acr *AccountCoaReport) error {
	return c.Update(AccountCoaReportModel, ids, acr)
}

// DeleteAccountCoaReport deletes an existing account.coa.report record.
func (c *Client) DeleteAccountCoaReport(id int64) error {
	return c.DeleteAccountCoaReports([]int64{id})
}

// DeleteAccountCoaReports deletes existing account.coa.report records.
func (c *Client) DeleteAccountCoaReports(ids []int64) error {
	return c.Delete(AccountCoaReportModel, ids)
}

// GetAccountCoaReport gets account.coa.report existing record.
func (c *Client) GetAccountCoaReport(id int64) (*AccountCoaReport, error) {
	acrs, err := c.GetAccountCoaReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if acrs != nil && len(*acrs) > 0 {
		return &((*acrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.coa.report not found", id)
}

// GetAccountCoaReports gets account.coa.report existing records.
func (c *Client) GetAccountCoaReports(ids []int64) (*AccountCoaReports, error) {
	acrs := &AccountCoaReports{}
	if err := c.Read(AccountCoaReportModel, ids, nil, acrs); err != nil {
		return nil, err
	}
	return acrs, nil
}

// FindAccountCoaReport finds account.coa.report record by querying it with criteria.
func (c *Client) FindAccountCoaReport(criteria *Criteria) (*AccountCoaReport, error) {
	acrs := &AccountCoaReports{}
	if err := c.SearchRead(AccountCoaReportModel, criteria, NewOptions().Limit(1), acrs); err != nil {
		return nil, err
	}
	if acrs != nil && len(*acrs) > 0 {
		return &((*acrs)[0]), nil
	}
	return nil, fmt.Errorf("account.coa.report was not found")
}

// FindAccountCoaReports finds account.coa.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountCoaReports(criteria *Criteria, options *Options) (*AccountCoaReports, error) {
	acrs := &AccountCoaReports{}
	if err := c.SearchRead(AccountCoaReportModel, criteria, options, acrs); err != nil {
		return nil, err
	}
	return acrs, nil
}

// FindAccountCoaReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountCoaReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountCoaReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountCoaReportId finds record id by querying it with criteria.
func (c *Client) FindAccountCoaReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountCoaReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.coa.report was not found")
}

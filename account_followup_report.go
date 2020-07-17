package odoo

import (
	"fmt"
)

// AccountFollowupReport represents account.followup.report model.
type AccountFollowupReport struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// AccountFollowupReports represents array of account.followup.report model.
type AccountFollowupReports []AccountFollowupReport

// AccountFollowupReportModel is the odoo model name.
const AccountFollowupReportModel = "account.followup.report"

// Many2One convert AccountFollowupReport to *Many2One.
func (afr *AccountFollowupReport) Many2One() *Many2One {
	return NewMany2One(afr.Id.Get(), "")
}

// CreateAccountFollowupReport creates a new account.followup.report model and returns its id.
func (c *Client) CreateAccountFollowupReport(afr *AccountFollowupReport) (int64, error) {
	return c.Create(AccountFollowupReportModel, afr)
}

// UpdateAccountFollowupReport updates an existing account.followup.report record.
func (c *Client) UpdateAccountFollowupReport(afr *AccountFollowupReport) error {
	return c.UpdateAccountFollowupReports([]int64{afr.Id.Get()}, afr)
}

// UpdateAccountFollowupReports updates existing account.followup.report records.
// All records (represented by ids) will be updated by afr values.
func (c *Client) UpdateAccountFollowupReports(ids []int64, afr *AccountFollowupReport) error {
	return c.Update(AccountFollowupReportModel, ids, afr)
}

// DeleteAccountFollowupReport deletes an existing account.followup.report record.
func (c *Client) DeleteAccountFollowupReport(id int64) error {
	return c.DeleteAccountFollowupReports([]int64{id})
}

// DeleteAccountFollowupReports deletes existing account.followup.report records.
func (c *Client) DeleteAccountFollowupReports(ids []int64) error {
	return c.Delete(AccountFollowupReportModel, ids)
}

// GetAccountFollowupReport gets account.followup.report existing record.
func (c *Client) GetAccountFollowupReport(id int64) (*AccountFollowupReport, error) {
	afrs, err := c.GetAccountFollowupReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if afrs != nil && len(*afrs) > 0 {
		return &((*afrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.followup.report not found", id)
}

// GetAccountFollowupReports gets account.followup.report existing records.
func (c *Client) GetAccountFollowupReports(ids []int64) (*AccountFollowupReports, error) {
	afrs := &AccountFollowupReports{}
	if err := c.Read(AccountFollowupReportModel, ids, nil, afrs); err != nil {
		return nil, err
	}
	return afrs, nil
}

// FindAccountFollowupReport finds account.followup.report record by querying it with criteria.
func (c *Client) FindAccountFollowupReport(criteria *Criteria) (*AccountFollowupReport, error) {
	afrs := &AccountFollowupReports{}
	if err := c.SearchRead(AccountFollowupReportModel, criteria, NewOptions().Limit(1), afrs); err != nil {
		return nil, err
	}
	if afrs != nil && len(*afrs) > 0 {
		return &((*afrs)[0]), nil
	}
	return nil, fmt.Errorf("account.followup.report was not found")
}

// FindAccountFollowupReports finds account.followup.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFollowupReports(criteria *Criteria, options *Options) (*AccountFollowupReports, error) {
	afrs := &AccountFollowupReports{}
	if err := c.SearchRead(AccountFollowupReportModel, criteria, options, afrs); err != nil {
		return nil, err
	}
	return afrs, nil
}

// FindAccountFollowupReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFollowupReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountFollowupReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountFollowupReportId finds record id by querying it with criteria.
func (c *Client) FindAccountFollowupReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountFollowupReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.followup.report was not found")
}

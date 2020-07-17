package odoo

import (
	"fmt"
)

// AccountGenericTaxReport represents account.generic.tax.report model.
type AccountGenericTaxReport struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// AccountGenericTaxReports represents array of account.generic.tax.report model.
type AccountGenericTaxReports []AccountGenericTaxReport

// AccountGenericTaxReportModel is the odoo model name.
const AccountGenericTaxReportModel = "account.generic.tax.report"

// Many2One convert AccountGenericTaxReport to *Many2One.
func (agtr *AccountGenericTaxReport) Many2One() *Many2One {
	return NewMany2One(agtr.Id.Get(), "")
}

// CreateAccountGenericTaxReport creates a new account.generic.tax.report model and returns its id.
func (c *Client) CreateAccountGenericTaxReport(agtr *AccountGenericTaxReport) (int64, error) {
	return c.Create(AccountGenericTaxReportModel, agtr)
}

// UpdateAccountGenericTaxReport updates an existing account.generic.tax.report record.
func (c *Client) UpdateAccountGenericTaxReport(agtr *AccountGenericTaxReport) error {
	return c.UpdateAccountGenericTaxReports([]int64{agtr.Id.Get()}, agtr)
}

// UpdateAccountGenericTaxReports updates existing account.generic.tax.report records.
// All records (represented by ids) will be updated by agtr values.
func (c *Client) UpdateAccountGenericTaxReports(ids []int64, agtr *AccountGenericTaxReport) error {
	return c.Update(AccountGenericTaxReportModel, ids, agtr)
}

// DeleteAccountGenericTaxReport deletes an existing account.generic.tax.report record.
func (c *Client) DeleteAccountGenericTaxReport(id int64) error {
	return c.DeleteAccountGenericTaxReports([]int64{id})
}

// DeleteAccountGenericTaxReports deletes existing account.generic.tax.report records.
func (c *Client) DeleteAccountGenericTaxReports(ids []int64) error {
	return c.Delete(AccountGenericTaxReportModel, ids)
}

// GetAccountGenericTaxReport gets account.generic.tax.report existing record.
func (c *Client) GetAccountGenericTaxReport(id int64) (*AccountGenericTaxReport, error) {
	agtrs, err := c.GetAccountGenericTaxReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if agtrs != nil && len(*agtrs) > 0 {
		return &((*agtrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.generic.tax.report not found", id)
}

// GetAccountGenericTaxReports gets account.generic.tax.report existing records.
func (c *Client) GetAccountGenericTaxReports(ids []int64) (*AccountGenericTaxReports, error) {
	agtrs := &AccountGenericTaxReports{}
	if err := c.Read(AccountGenericTaxReportModel, ids, nil, agtrs); err != nil {
		return nil, err
	}
	return agtrs, nil
}

// FindAccountGenericTaxReport finds account.generic.tax.report record by querying it with criteria.
func (c *Client) FindAccountGenericTaxReport(criteria *Criteria) (*AccountGenericTaxReport, error) {
	agtrs := &AccountGenericTaxReports{}
	if err := c.SearchRead(AccountGenericTaxReportModel, criteria, NewOptions().Limit(1), agtrs); err != nil {
		return nil, err
	}
	if agtrs != nil && len(*agtrs) > 0 {
		return &((*agtrs)[0]), nil
	}
	return nil, fmt.Errorf("account.generic.tax.report was not found")
}

// FindAccountGenericTaxReports finds account.generic.tax.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountGenericTaxReports(criteria *Criteria, options *Options) (*AccountGenericTaxReports, error) {
	agtrs := &AccountGenericTaxReports{}
	if err := c.SearchRead(AccountGenericTaxReportModel, criteria, options, agtrs); err != nil {
		return nil, err
	}
	return agtrs, nil
}

// FindAccountGenericTaxReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountGenericTaxReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountGenericTaxReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountGenericTaxReportId finds record id by querying it with criteria.
func (c *Client) FindAccountGenericTaxReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountGenericTaxReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.generic.tax.report was not found")
}

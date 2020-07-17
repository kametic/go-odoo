package odoo

import (
	"fmt"
)

// AccountBankReconciliationReport represents account.bank.reconciliation.report model.
type AccountBankReconciliationReport struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// AccountBankReconciliationReports represents array of account.bank.reconciliation.report model.
type AccountBankReconciliationReports []AccountBankReconciliationReport

// AccountBankReconciliationReportModel is the odoo model name.
const AccountBankReconciliationReportModel = "account.bank.reconciliation.report"

// Many2One convert AccountBankReconciliationReport to *Many2One.
func (abrr *AccountBankReconciliationReport) Many2One() *Many2One {
	return NewMany2One(abrr.Id.Get(), "")
}

// CreateAccountBankReconciliationReport creates a new account.bank.reconciliation.report model and returns its id.
func (c *Client) CreateAccountBankReconciliationReport(abrr *AccountBankReconciliationReport) (int64, error) {
	return c.Create(AccountBankReconciliationReportModel, abrr)
}

// UpdateAccountBankReconciliationReport updates an existing account.bank.reconciliation.report record.
func (c *Client) UpdateAccountBankReconciliationReport(abrr *AccountBankReconciliationReport) error {
	return c.UpdateAccountBankReconciliationReports([]int64{abrr.Id.Get()}, abrr)
}

// UpdateAccountBankReconciliationReports updates existing account.bank.reconciliation.report records.
// All records (represented by ids) will be updated by abrr values.
func (c *Client) UpdateAccountBankReconciliationReports(ids []int64, abrr *AccountBankReconciliationReport) error {
	return c.Update(AccountBankReconciliationReportModel, ids, abrr)
}

// DeleteAccountBankReconciliationReport deletes an existing account.bank.reconciliation.report record.
func (c *Client) DeleteAccountBankReconciliationReport(id int64) error {
	return c.DeleteAccountBankReconciliationReports([]int64{id})
}

// DeleteAccountBankReconciliationReports deletes existing account.bank.reconciliation.report records.
func (c *Client) DeleteAccountBankReconciliationReports(ids []int64) error {
	return c.Delete(AccountBankReconciliationReportModel, ids)
}

// GetAccountBankReconciliationReport gets account.bank.reconciliation.report existing record.
func (c *Client) GetAccountBankReconciliationReport(id int64) (*AccountBankReconciliationReport, error) {
	abrrs, err := c.GetAccountBankReconciliationReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if abrrs != nil && len(*abrrs) > 0 {
		return &((*abrrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.bank.reconciliation.report not found", id)
}

// GetAccountBankReconciliationReports gets account.bank.reconciliation.report existing records.
func (c *Client) GetAccountBankReconciliationReports(ids []int64) (*AccountBankReconciliationReports, error) {
	abrrs := &AccountBankReconciliationReports{}
	if err := c.Read(AccountBankReconciliationReportModel, ids, nil, abrrs); err != nil {
		return nil, err
	}
	return abrrs, nil
}

// FindAccountBankReconciliationReport finds account.bank.reconciliation.report record by querying it with criteria.
func (c *Client) FindAccountBankReconciliationReport(criteria *Criteria) (*AccountBankReconciliationReport, error) {
	abrrs := &AccountBankReconciliationReports{}
	if err := c.SearchRead(AccountBankReconciliationReportModel, criteria, NewOptions().Limit(1), abrrs); err != nil {
		return nil, err
	}
	if abrrs != nil && len(*abrrs) > 0 {
		return &((*abrrs)[0]), nil
	}
	return nil, fmt.Errorf("account.bank.reconciliation.report was not found")
}

// FindAccountBankReconciliationReports finds account.bank.reconciliation.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankReconciliationReports(criteria *Criteria, options *Options) (*AccountBankReconciliationReports, error) {
	abrrs := &AccountBankReconciliationReports{}
	if err := c.SearchRead(AccountBankReconciliationReportModel, criteria, options, abrrs); err != nil {
		return nil, err
	}
	return abrrs, nil
}

// FindAccountBankReconciliationReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountBankReconciliationReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountBankReconciliationReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountBankReconciliationReportId finds record id by querying it with criteria.
func (c *Client) FindAccountBankReconciliationReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountBankReconciliationReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.bank.reconciliation.report was not found")
}

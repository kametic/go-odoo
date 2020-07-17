package odoo

import (
	"fmt"
)

// AccountReportManager represents account.report.manager model.
type AccountReportManager struct {
	LastUpdate        *Time     `xmlrpc:"__last_update,omptempty"`
	CompanyId         *Many2One `xmlrpc:"company_id,omptempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName       *String   `xmlrpc:"display_name,omptempty"`
	FinancialReportId *Many2One `xmlrpc:"financial_report_id,omptempty"`
	FootnotesIds      *Relation `xmlrpc:"footnotes_ids,omptempty"`
	Id                *Int      `xmlrpc:"id,omptempty"`
	PartnerId         *Many2One `xmlrpc:"partner_id,omptempty"`
	ReportName        *String   `xmlrpc:"report_name,omptempty"`
	Summary           *String   `xmlrpc:"summary,omptempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountReportManagers represents array of account.report.manager model.
type AccountReportManagers []AccountReportManager

// AccountReportManagerModel is the odoo model name.
const AccountReportManagerModel = "account.report.manager"

// Many2One convert AccountReportManager to *Many2One.
func (arm *AccountReportManager) Many2One() *Many2One {
	return NewMany2One(arm.Id.Get(), "")
}

// CreateAccountReportManager creates a new account.report.manager model and returns its id.
func (c *Client) CreateAccountReportManager(arm *AccountReportManager) (int64, error) {
	return c.Create(AccountReportManagerModel, arm)
}

// UpdateAccountReportManager updates an existing account.report.manager record.
func (c *Client) UpdateAccountReportManager(arm *AccountReportManager) error {
	return c.UpdateAccountReportManagers([]int64{arm.Id.Get()}, arm)
}

// UpdateAccountReportManagers updates existing account.report.manager records.
// All records (represented by ids) will be updated by arm values.
func (c *Client) UpdateAccountReportManagers(ids []int64, arm *AccountReportManager) error {
	return c.Update(AccountReportManagerModel, ids, arm)
}

// DeleteAccountReportManager deletes an existing account.report.manager record.
func (c *Client) DeleteAccountReportManager(id int64) error {
	return c.DeleteAccountReportManagers([]int64{id})
}

// DeleteAccountReportManagers deletes existing account.report.manager records.
func (c *Client) DeleteAccountReportManagers(ids []int64) error {
	return c.Delete(AccountReportManagerModel, ids)
}

// GetAccountReportManager gets account.report.manager existing record.
func (c *Client) GetAccountReportManager(id int64) (*AccountReportManager, error) {
	arms, err := c.GetAccountReportManagers([]int64{id})
	if err != nil {
		return nil, err
	}
	if arms != nil && len(*arms) > 0 {
		return &((*arms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.report.manager not found", id)
}

// GetAccountReportManagers gets account.report.manager existing records.
func (c *Client) GetAccountReportManagers(ids []int64) (*AccountReportManagers, error) {
	arms := &AccountReportManagers{}
	if err := c.Read(AccountReportManagerModel, ids, nil, arms); err != nil {
		return nil, err
	}
	return arms, nil
}

// FindAccountReportManager finds account.report.manager record by querying it with criteria.
func (c *Client) FindAccountReportManager(criteria *Criteria) (*AccountReportManager, error) {
	arms := &AccountReportManagers{}
	if err := c.SearchRead(AccountReportManagerModel, criteria, NewOptions().Limit(1), arms); err != nil {
		return nil, err
	}
	if arms != nil && len(*arms) > 0 {
		return &((*arms)[0]), nil
	}
	return nil, fmt.Errorf("account.report.manager was not found")
}

// FindAccountReportManagers finds account.report.manager records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportManagers(criteria *Criteria, options *Options) (*AccountReportManagers, error) {
	arms := &AccountReportManagers{}
	if err := c.SearchRead(AccountReportManagerModel, criteria, options, arms); err != nil {
		return nil, err
	}
	return arms, nil
}

// FindAccountReportManagerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportManagerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountReportManagerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountReportManagerId finds record id by querying it with criteria.
func (c *Client) FindAccountReportManagerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReportManagerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.report.manager was not found")
}

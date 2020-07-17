package odoo

import (
	"fmt"
)

// AccountFinancialHtmlReport represents account.financial.html.report model.
type AccountFinancialHtmlReport struct {
	LastUpdate        *Time     `xmlrpc:"__last_update,omptempty"`
	Analytic          *Bool     `xmlrpc:"analytic,omptempty"`
	CashBasis         *Bool     `xmlrpc:"cash_basis,omptempty"`
	CompanyId         *Many2One `xmlrpc:"company_id,omptempty"`
	Comparison        *Bool     `xmlrpc:"comparison,omptempty"`
	CreateDate        *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One `xmlrpc:"create_uid,omptempty"`
	DateRange         *Bool     `xmlrpc:"date_range,omptempty"`
	DebitCredit       *Bool     `xmlrpc:"debit_credit,omptempty"`
	DisplayName       *String   `xmlrpc:"display_name,omptempty"`
	GeneratedMenuId   *Many2One `xmlrpc:"generated_menu_id,omptempty"`
	HierarchyOption   *Bool     `xmlrpc:"hierarchy_option,omptempty"`
	Id                *Int      `xmlrpc:"id,omptempty"`
	LineIds           *Relation `xmlrpc:"line_ids,omptempty"`
	Name              *String   `xmlrpc:"name,omptempty"`
	ParentId          *Many2One `xmlrpc:"parent_id,omptempty"`
	ShowJournalFilter *Bool     `xmlrpc:"show_journal_filter,omptempty"`
	TaxReport         *Bool     `xmlrpc:"tax_report,omptempty"`
	UnfoldAllFilter   *Bool     `xmlrpc:"unfold_all_filter,omptempty"`
	WriteDate         *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountFinancialHtmlReports represents array of account.financial.html.report model.
type AccountFinancialHtmlReports []AccountFinancialHtmlReport

// AccountFinancialHtmlReportModel is the odoo model name.
const AccountFinancialHtmlReportModel = "account.financial.html.report"

// Many2One convert AccountFinancialHtmlReport to *Many2One.
func (afhr *AccountFinancialHtmlReport) Many2One() *Many2One {
	return NewMany2One(afhr.Id.Get(), "")
}

// CreateAccountFinancialHtmlReport creates a new account.financial.html.report model and returns its id.
func (c *Client) CreateAccountFinancialHtmlReport(afhr *AccountFinancialHtmlReport) (int64, error) {
	return c.Create(AccountFinancialHtmlReportModel, afhr)
}

// UpdateAccountFinancialHtmlReport updates an existing account.financial.html.report record.
func (c *Client) UpdateAccountFinancialHtmlReport(afhr *AccountFinancialHtmlReport) error {
	return c.UpdateAccountFinancialHtmlReports([]int64{afhr.Id.Get()}, afhr)
}

// UpdateAccountFinancialHtmlReports updates existing account.financial.html.report records.
// All records (represented by ids) will be updated by afhr values.
func (c *Client) UpdateAccountFinancialHtmlReports(ids []int64, afhr *AccountFinancialHtmlReport) error {
	return c.Update(AccountFinancialHtmlReportModel, ids, afhr)
}

// DeleteAccountFinancialHtmlReport deletes an existing account.financial.html.report record.
func (c *Client) DeleteAccountFinancialHtmlReport(id int64) error {
	return c.DeleteAccountFinancialHtmlReports([]int64{id})
}

// DeleteAccountFinancialHtmlReports deletes existing account.financial.html.report records.
func (c *Client) DeleteAccountFinancialHtmlReports(ids []int64) error {
	return c.Delete(AccountFinancialHtmlReportModel, ids)
}

// GetAccountFinancialHtmlReport gets account.financial.html.report existing record.
func (c *Client) GetAccountFinancialHtmlReport(id int64) (*AccountFinancialHtmlReport, error) {
	afhrs, err := c.GetAccountFinancialHtmlReports([]int64{id})
	if err != nil {
		return nil, err
	}
	if afhrs != nil && len(*afhrs) > 0 {
		return &((*afhrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.financial.html.report not found", id)
}

// GetAccountFinancialHtmlReports gets account.financial.html.report existing records.
func (c *Client) GetAccountFinancialHtmlReports(ids []int64) (*AccountFinancialHtmlReports, error) {
	afhrs := &AccountFinancialHtmlReports{}
	if err := c.Read(AccountFinancialHtmlReportModel, ids, nil, afhrs); err != nil {
		return nil, err
	}
	return afhrs, nil
}

// FindAccountFinancialHtmlReport finds account.financial.html.report record by querying it with criteria.
func (c *Client) FindAccountFinancialHtmlReport(criteria *Criteria) (*AccountFinancialHtmlReport, error) {
	afhrs := &AccountFinancialHtmlReports{}
	if err := c.SearchRead(AccountFinancialHtmlReportModel, criteria, NewOptions().Limit(1), afhrs); err != nil {
		return nil, err
	}
	if afhrs != nil && len(*afhrs) > 0 {
		return &((*afhrs)[0]), nil
	}
	return nil, fmt.Errorf("account.financial.html.report was not found")
}

// FindAccountFinancialHtmlReports finds account.financial.html.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFinancialHtmlReports(criteria *Criteria, options *Options) (*AccountFinancialHtmlReports, error) {
	afhrs := &AccountFinancialHtmlReports{}
	if err := c.SearchRead(AccountFinancialHtmlReportModel, criteria, options, afhrs); err != nil {
		return nil, err
	}
	return afhrs, nil
}

// FindAccountFinancialHtmlReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFinancialHtmlReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountFinancialHtmlReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountFinancialHtmlReportId finds record id by querying it with criteria.
func (c *Client) FindAccountFinancialHtmlReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountFinancialHtmlReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.financial.html.report was not found")
}

package odoo

import (
	"fmt"
)

// AccountFinancialHtmlReportLine represents account.financial.html.report.line model.
type AccountFinancialHtmlReportLine struct {
	LastUpdate         *Time      `xmlrpc:"__last_update,omptempty"`
	ActionId           *Many2One  `xmlrpc:"action_id,omptempty"`
	ChildrenIds        *Relation  `xmlrpc:"children_ids,omptempty"`
	Code               *String    `xmlrpc:"code,omptempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName        *String    `xmlrpc:"display_name,omptempty"`
	Domain             *String    `xmlrpc:"domain,omptempty"`
	FigureType         *Selection `xmlrpc:"figure_type,omptempty"`
	FinancialReportId  *Many2One  `xmlrpc:"financial_report_id,omptempty"`
	Formulas           *String    `xmlrpc:"formulas,omptempty"`
	GreenOnPositive    *Bool      `xmlrpc:"green_on_positive,omptempty"`
	Groupby            *String    `xmlrpc:"groupby,omptempty"`
	HideIfZero         *Bool      `xmlrpc:"hide_if_zero,omptempty"`
	Id                 *Int       `xmlrpc:"id,omptempty"`
	Level              *Int       `xmlrpc:"level,omptempty"`
	Name               *String    `xmlrpc:"name,omptempty"`
	ParentId           *Many2One  `xmlrpc:"parent_id,omptempty"`
	Sequence           *Int       `xmlrpc:"sequence,omptempty"`
	ShowDomain         *Selection `xmlrpc:"show_domain,omptempty"`
	SpecialDateChanger *Selection `xmlrpc:"special_date_changer,omptempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountFinancialHtmlReportLines represents array of account.financial.html.report.line model.
type AccountFinancialHtmlReportLines []AccountFinancialHtmlReportLine

// AccountFinancialHtmlReportLineModel is the odoo model name.
const AccountFinancialHtmlReportLineModel = "account.financial.html.report.line"

// Many2One convert AccountFinancialHtmlReportLine to *Many2One.
func (afhrl *AccountFinancialHtmlReportLine) Many2One() *Many2One {
	return NewMany2One(afhrl.Id.Get(), "")
}

// CreateAccountFinancialHtmlReportLine creates a new account.financial.html.report.line model and returns its id.
func (c *Client) CreateAccountFinancialHtmlReportLine(afhrl *AccountFinancialHtmlReportLine) (int64, error) {
	return c.Create(AccountFinancialHtmlReportLineModel, afhrl)
}

// UpdateAccountFinancialHtmlReportLine updates an existing account.financial.html.report.line record.
func (c *Client) UpdateAccountFinancialHtmlReportLine(afhrl *AccountFinancialHtmlReportLine) error {
	return c.UpdateAccountFinancialHtmlReportLines([]int64{afhrl.Id.Get()}, afhrl)
}

// UpdateAccountFinancialHtmlReportLines updates existing account.financial.html.report.line records.
// All records (represented by ids) will be updated by afhrl values.
func (c *Client) UpdateAccountFinancialHtmlReportLines(ids []int64, afhrl *AccountFinancialHtmlReportLine) error {
	return c.Update(AccountFinancialHtmlReportLineModel, ids, afhrl)
}

// DeleteAccountFinancialHtmlReportLine deletes an existing account.financial.html.report.line record.
func (c *Client) DeleteAccountFinancialHtmlReportLine(id int64) error {
	return c.DeleteAccountFinancialHtmlReportLines([]int64{id})
}

// DeleteAccountFinancialHtmlReportLines deletes existing account.financial.html.report.line records.
func (c *Client) DeleteAccountFinancialHtmlReportLines(ids []int64) error {
	return c.Delete(AccountFinancialHtmlReportLineModel, ids)
}

// GetAccountFinancialHtmlReportLine gets account.financial.html.report.line existing record.
func (c *Client) GetAccountFinancialHtmlReportLine(id int64) (*AccountFinancialHtmlReportLine, error) {
	afhrls, err := c.GetAccountFinancialHtmlReportLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if afhrls != nil && len(*afhrls) > 0 {
		return &((*afhrls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.financial.html.report.line not found", id)
}

// GetAccountFinancialHtmlReportLines gets account.financial.html.report.line existing records.
func (c *Client) GetAccountFinancialHtmlReportLines(ids []int64) (*AccountFinancialHtmlReportLines, error) {
	afhrls := &AccountFinancialHtmlReportLines{}
	if err := c.Read(AccountFinancialHtmlReportLineModel, ids, nil, afhrls); err != nil {
		return nil, err
	}
	return afhrls, nil
}

// FindAccountFinancialHtmlReportLine finds account.financial.html.report.line record by querying it with criteria.
func (c *Client) FindAccountFinancialHtmlReportLine(criteria *Criteria) (*AccountFinancialHtmlReportLine, error) {
	afhrls := &AccountFinancialHtmlReportLines{}
	if err := c.SearchRead(AccountFinancialHtmlReportLineModel, criteria, NewOptions().Limit(1), afhrls); err != nil {
		return nil, err
	}
	if afhrls != nil && len(*afhrls) > 0 {
		return &((*afhrls)[0]), nil
	}
	return nil, fmt.Errorf("account.financial.html.report.line was not found")
}

// FindAccountFinancialHtmlReportLines finds account.financial.html.report.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFinancialHtmlReportLines(criteria *Criteria, options *Options) (*AccountFinancialHtmlReportLines, error) {
	afhrls := &AccountFinancialHtmlReportLines{}
	if err := c.SearchRead(AccountFinancialHtmlReportLineModel, criteria, options, afhrls); err != nil {
		return nil, err
	}
	return afhrls, nil
}

// FindAccountFinancialHtmlReportLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountFinancialHtmlReportLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountFinancialHtmlReportLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountFinancialHtmlReportLineId finds record id by querying it with criteria.
func (c *Client) FindAccountFinancialHtmlReportLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountFinancialHtmlReportLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.financial.html.report.line was not found")
}

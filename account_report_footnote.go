package odoo

import (
	"fmt"
)

// AccountReportFootnote represents account.report.footnote model.
type AccountReportFootnote struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Line        *String   `xmlrpc:"line,omptempty"`
	ManagerId   *Many2One `xmlrpc:"manager_id,omptempty"`
	Text        *String   `xmlrpc:"text,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountReportFootnotes represents array of account.report.footnote model.
type AccountReportFootnotes []AccountReportFootnote

// AccountReportFootnoteModel is the odoo model name.
const AccountReportFootnoteModel = "account.report.footnote"

// Many2One convert AccountReportFootnote to *Many2One.
func (arf *AccountReportFootnote) Many2One() *Many2One {
	return NewMany2One(arf.Id.Get(), "")
}

// CreateAccountReportFootnote creates a new account.report.footnote model and returns its id.
func (c *Client) CreateAccountReportFootnote(arf *AccountReportFootnote) (int64, error) {
	return c.Create(AccountReportFootnoteModel, arf)
}

// UpdateAccountReportFootnote updates an existing account.report.footnote record.
func (c *Client) UpdateAccountReportFootnote(arf *AccountReportFootnote) error {
	return c.UpdateAccountReportFootnotes([]int64{arf.Id.Get()}, arf)
}

// UpdateAccountReportFootnotes updates existing account.report.footnote records.
// All records (represented by ids) will be updated by arf values.
func (c *Client) UpdateAccountReportFootnotes(ids []int64, arf *AccountReportFootnote) error {
	return c.Update(AccountReportFootnoteModel, ids, arf)
}

// DeleteAccountReportFootnote deletes an existing account.report.footnote record.
func (c *Client) DeleteAccountReportFootnote(id int64) error {
	return c.DeleteAccountReportFootnotes([]int64{id})
}

// DeleteAccountReportFootnotes deletes existing account.report.footnote records.
func (c *Client) DeleteAccountReportFootnotes(ids []int64) error {
	return c.Delete(AccountReportFootnoteModel, ids)
}

// GetAccountReportFootnote gets account.report.footnote existing record.
func (c *Client) GetAccountReportFootnote(id int64) (*AccountReportFootnote, error) {
	arfs, err := c.GetAccountReportFootnotes([]int64{id})
	if err != nil {
		return nil, err
	}
	if arfs != nil && len(*arfs) > 0 {
		return &((*arfs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.report.footnote not found", id)
}

// GetAccountReportFootnotes gets account.report.footnote existing records.
func (c *Client) GetAccountReportFootnotes(ids []int64) (*AccountReportFootnotes, error) {
	arfs := &AccountReportFootnotes{}
	if err := c.Read(AccountReportFootnoteModel, ids, nil, arfs); err != nil {
		return nil, err
	}
	return arfs, nil
}

// FindAccountReportFootnote finds account.report.footnote record by querying it with criteria.
func (c *Client) FindAccountReportFootnote(criteria *Criteria) (*AccountReportFootnote, error) {
	arfs := &AccountReportFootnotes{}
	if err := c.SearchRead(AccountReportFootnoteModel, criteria, NewOptions().Limit(1), arfs); err != nil {
		return nil, err
	}
	if arfs != nil && len(*arfs) > 0 {
		return &((*arfs)[0]), nil
	}
	return nil, fmt.Errorf("account.report.footnote was not found")
}

// FindAccountReportFootnotes finds account.report.footnote records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportFootnotes(criteria *Criteria, options *Options) (*AccountReportFootnotes, error) {
	arfs := &AccountReportFootnotes{}
	if err := c.SearchRead(AccountReportFootnoteModel, criteria, options, arfs); err != nil {
		return nil, err
	}
	return arfs, nil
}

// FindAccountReportFootnoteIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReportFootnoteIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountReportFootnoteModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountReportFootnoteId finds record id by querying it with criteria.
func (c *Client) FindAccountReportFootnoteId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReportFootnoteModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.report.footnote was not found")
}

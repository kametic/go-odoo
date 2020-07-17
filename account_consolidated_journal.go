package odoo

import (
	"fmt"
)

// AccountConsolidatedJournal represents account.consolidated.journal model.
type AccountConsolidatedJournal struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omptempty"`
	DisplayName *String `xmlrpc:"display_name,omptempty"`
	Id          *Int    `xmlrpc:"id,omptempty"`
}

// AccountConsolidatedJournals represents array of account.consolidated.journal model.
type AccountConsolidatedJournals []AccountConsolidatedJournal

// AccountConsolidatedJournalModel is the odoo model name.
const AccountConsolidatedJournalModel = "account.consolidated.journal"

// Many2One convert AccountConsolidatedJournal to *Many2One.
func (acj *AccountConsolidatedJournal) Many2One() *Many2One {
	return NewMany2One(acj.Id.Get(), "")
}

// CreateAccountConsolidatedJournal creates a new account.consolidated.journal model and returns its id.
func (c *Client) CreateAccountConsolidatedJournal(acj *AccountConsolidatedJournal) (int64, error) {
	return c.Create(AccountConsolidatedJournalModel, acj)
}

// UpdateAccountConsolidatedJournal updates an existing account.consolidated.journal record.
func (c *Client) UpdateAccountConsolidatedJournal(acj *AccountConsolidatedJournal) error {
	return c.UpdateAccountConsolidatedJournals([]int64{acj.Id.Get()}, acj)
}

// UpdateAccountConsolidatedJournals updates existing account.consolidated.journal records.
// All records (represented by ids) will be updated by acj values.
func (c *Client) UpdateAccountConsolidatedJournals(ids []int64, acj *AccountConsolidatedJournal) error {
	return c.Update(AccountConsolidatedJournalModel, ids, acj)
}

// DeleteAccountConsolidatedJournal deletes an existing account.consolidated.journal record.
func (c *Client) DeleteAccountConsolidatedJournal(id int64) error {
	return c.DeleteAccountConsolidatedJournals([]int64{id})
}

// DeleteAccountConsolidatedJournals deletes existing account.consolidated.journal records.
func (c *Client) DeleteAccountConsolidatedJournals(ids []int64) error {
	return c.Delete(AccountConsolidatedJournalModel, ids)
}

// GetAccountConsolidatedJournal gets account.consolidated.journal existing record.
func (c *Client) GetAccountConsolidatedJournal(id int64) (*AccountConsolidatedJournal, error) {
	acjs, err := c.GetAccountConsolidatedJournals([]int64{id})
	if err != nil {
		return nil, err
	}
	if acjs != nil && len(*acjs) > 0 {
		return &((*acjs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.consolidated.journal not found", id)
}

// GetAccountConsolidatedJournals gets account.consolidated.journal existing records.
func (c *Client) GetAccountConsolidatedJournals(ids []int64) (*AccountConsolidatedJournals, error) {
	acjs := &AccountConsolidatedJournals{}
	if err := c.Read(AccountConsolidatedJournalModel, ids, nil, acjs); err != nil {
		return nil, err
	}
	return acjs, nil
}

// FindAccountConsolidatedJournal finds account.consolidated.journal record by querying it with criteria.
func (c *Client) FindAccountConsolidatedJournal(criteria *Criteria) (*AccountConsolidatedJournal, error) {
	acjs := &AccountConsolidatedJournals{}
	if err := c.SearchRead(AccountConsolidatedJournalModel, criteria, NewOptions().Limit(1), acjs); err != nil {
		return nil, err
	}
	if acjs != nil && len(*acjs) > 0 {
		return &((*acjs)[0]), nil
	}
	return nil, fmt.Errorf("account.consolidated.journal was not found")
}

// FindAccountConsolidatedJournals finds account.consolidated.journal records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountConsolidatedJournals(criteria *Criteria, options *Options) (*AccountConsolidatedJournals, error) {
	acjs := &AccountConsolidatedJournals{}
	if err := c.SearchRead(AccountConsolidatedJournalModel, criteria, options, acjs); err != nil {
		return nil, err
	}
	return acjs, nil
}

// FindAccountConsolidatedJournalIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountConsolidatedJournalIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountConsolidatedJournalModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountConsolidatedJournalId finds record id by querying it with criteria.
func (c *Client) FindAccountConsolidatedJournalId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountConsolidatedJournalModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.consolidated.journal was not found")
}

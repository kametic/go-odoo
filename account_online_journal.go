package odoo

import (
	"fmt"
)

// AccountOnlineJournal represents account.online.journal model.
type AccountOnlineJournal struct {
	LastUpdate              *Time     `xmlrpc:"__last_update,omptempty"`
	AccountNumber           *String   `xmlrpc:"account_number,omptempty"`
	AccountOnlineProviderId *Many2One `xmlrpc:"account_online_provider_id,omptempty"`
	Balance                 *Float    `xmlrpc:"balance,omptempty"`
	CreateDate              *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid               *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName             *String   `xmlrpc:"display_name,omptempty"`
	Id                      *Int      `xmlrpc:"id,omptempty"`
	JournalIds              *Relation `xmlrpc:"journal_ids,omptempty"`
	LastSync                *Time     `xmlrpc:"last_sync,omptempty"`
	Name                    *String   `xmlrpc:"name,omptempty"`
	OnlineIdentifier        *String   `xmlrpc:"online_identifier,omptempty"`
	ProviderName            *String   `xmlrpc:"provider_name,omptempty"`
	WriteDate               *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                *Many2One `xmlrpc:"write_uid,omptempty"`
	YodleeAccountStatus     *String   `xmlrpc:"yodlee_account_status,omptempty"`
	YodleeStatusCode        *Int      `xmlrpc:"yodlee_status_code,omptempty"`
}

// AccountOnlineJournals represents array of account.online.journal model.
type AccountOnlineJournals []AccountOnlineJournal

// AccountOnlineJournalModel is the odoo model name.
const AccountOnlineJournalModel = "account.online.journal"

// Many2One convert AccountOnlineJournal to *Many2One.
func (aoj *AccountOnlineJournal) Many2One() *Many2One {
	return NewMany2One(aoj.Id.Get(), "")
}

// CreateAccountOnlineJournal creates a new account.online.journal model and returns its id.
func (c *Client) CreateAccountOnlineJournal(aoj *AccountOnlineJournal) (int64, error) {
	return c.Create(AccountOnlineJournalModel, aoj)
}

// UpdateAccountOnlineJournal updates an existing account.online.journal record.
func (c *Client) UpdateAccountOnlineJournal(aoj *AccountOnlineJournal) error {
	return c.UpdateAccountOnlineJournals([]int64{aoj.Id.Get()}, aoj)
}

// UpdateAccountOnlineJournals updates existing account.online.journal records.
// All records (represented by ids) will be updated by aoj values.
func (c *Client) UpdateAccountOnlineJournals(ids []int64, aoj *AccountOnlineJournal) error {
	return c.Update(AccountOnlineJournalModel, ids, aoj)
}

// DeleteAccountOnlineJournal deletes an existing account.online.journal record.
func (c *Client) DeleteAccountOnlineJournal(id int64) error {
	return c.DeleteAccountOnlineJournals([]int64{id})
}

// DeleteAccountOnlineJournals deletes existing account.online.journal records.
func (c *Client) DeleteAccountOnlineJournals(ids []int64) error {
	return c.Delete(AccountOnlineJournalModel, ids)
}

// GetAccountOnlineJournal gets account.online.journal existing record.
func (c *Client) GetAccountOnlineJournal(id int64) (*AccountOnlineJournal, error) {
	aojs, err := c.GetAccountOnlineJournals([]int64{id})
	if err != nil {
		return nil, err
	}
	if aojs != nil && len(*aojs) > 0 {
		return &((*aojs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.online.journal not found", id)
}

// GetAccountOnlineJournals gets account.online.journal existing records.
func (c *Client) GetAccountOnlineJournals(ids []int64) (*AccountOnlineJournals, error) {
	aojs := &AccountOnlineJournals{}
	if err := c.Read(AccountOnlineJournalModel, ids, nil, aojs); err != nil {
		return nil, err
	}
	return aojs, nil
}

// FindAccountOnlineJournal finds account.online.journal record by querying it with criteria.
func (c *Client) FindAccountOnlineJournal(criteria *Criteria) (*AccountOnlineJournal, error) {
	aojs := &AccountOnlineJournals{}
	if err := c.SearchRead(AccountOnlineJournalModel, criteria, NewOptions().Limit(1), aojs); err != nil {
		return nil, err
	}
	if aojs != nil && len(*aojs) > 0 {
		return &((*aojs)[0]), nil
	}
	return nil, fmt.Errorf("account.online.journal was not found")
}

// FindAccountOnlineJournals finds account.online.journal records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountOnlineJournals(criteria *Criteria, options *Options) (*AccountOnlineJournals, error) {
	aojs := &AccountOnlineJournals{}
	if err := c.SearchRead(AccountOnlineJournalModel, criteria, options, aojs); err != nil {
		return nil, err
	}
	return aojs, nil
}

// FindAccountOnlineJournalIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountOnlineJournalIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountOnlineJournalModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountOnlineJournalId finds record id by querying it with criteria.
func (c *Client) FindAccountOnlineJournalId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountOnlineJournalModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.online.journal was not found")
}

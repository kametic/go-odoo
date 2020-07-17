package odoo

import (
	"fmt"
)

// AccountOnlineWizard represents account.online.wizard model.
type AccountOnlineWizard struct {
	LastUpdate                *Time     `xmlrpc:"__last_update,omptempty"`
	AccountOnlineJournalId    *Many2One `xmlrpc:"account_online_journal_id,omptempty"`
	CountAccountOnlineJournal *Int      `xmlrpc:"count_account_online_journal,omptempty"`
	CreateDate                *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid                 *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName               *String   `xmlrpc:"display_name,omptempty"`
	Id                        *Int      `xmlrpc:"id,omptempty"`
	JournalId                 *Many2One `xmlrpc:"journal_id,omptempty"`
	NumberAdded               *String   `xmlrpc:"number_added,omptempty"`
	SyncDate                  *Time     `xmlrpc:"sync_date,omptempty"`
	WriteDate                 *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid                  *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountOnlineWizards represents array of account.online.wizard model.
type AccountOnlineWizards []AccountOnlineWizard

// AccountOnlineWizardModel is the odoo model name.
const AccountOnlineWizardModel = "account.online.wizard"

// Many2One convert AccountOnlineWizard to *Many2One.
func (aow *AccountOnlineWizard) Many2One() *Many2One {
	return NewMany2One(aow.Id.Get(), "")
}

// CreateAccountOnlineWizard creates a new account.online.wizard model and returns its id.
func (c *Client) CreateAccountOnlineWizard(aow *AccountOnlineWizard) (int64, error) {
	return c.Create(AccountOnlineWizardModel, aow)
}

// UpdateAccountOnlineWizard updates an existing account.online.wizard record.
func (c *Client) UpdateAccountOnlineWizard(aow *AccountOnlineWizard) error {
	return c.UpdateAccountOnlineWizards([]int64{aow.Id.Get()}, aow)
}

// UpdateAccountOnlineWizards updates existing account.online.wizard records.
// All records (represented by ids) will be updated by aow values.
func (c *Client) UpdateAccountOnlineWizards(ids []int64, aow *AccountOnlineWizard) error {
	return c.Update(AccountOnlineWizardModel, ids, aow)
}

// DeleteAccountOnlineWizard deletes an existing account.online.wizard record.
func (c *Client) DeleteAccountOnlineWizard(id int64) error {
	return c.DeleteAccountOnlineWizards([]int64{id})
}

// DeleteAccountOnlineWizards deletes existing account.online.wizard records.
func (c *Client) DeleteAccountOnlineWizards(ids []int64) error {
	return c.Delete(AccountOnlineWizardModel, ids)
}

// GetAccountOnlineWizard gets account.online.wizard existing record.
func (c *Client) GetAccountOnlineWizard(id int64) (*AccountOnlineWizard, error) {
	aows, err := c.GetAccountOnlineWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if aows != nil && len(*aows) > 0 {
		return &((*aows)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.online.wizard not found", id)
}

// GetAccountOnlineWizards gets account.online.wizard existing records.
func (c *Client) GetAccountOnlineWizards(ids []int64) (*AccountOnlineWizards, error) {
	aows := &AccountOnlineWizards{}
	if err := c.Read(AccountOnlineWizardModel, ids, nil, aows); err != nil {
		return nil, err
	}
	return aows, nil
}

// FindAccountOnlineWizard finds account.online.wizard record by querying it with criteria.
func (c *Client) FindAccountOnlineWizard(criteria *Criteria) (*AccountOnlineWizard, error) {
	aows := &AccountOnlineWizards{}
	if err := c.SearchRead(AccountOnlineWizardModel, criteria, NewOptions().Limit(1), aows); err != nil {
		return nil, err
	}
	if aows != nil && len(*aows) > 0 {
		return &((*aows)[0]), nil
	}
	return nil, fmt.Errorf("account.online.wizard was not found")
}

// FindAccountOnlineWizards finds account.online.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountOnlineWizards(criteria *Criteria, options *Options) (*AccountOnlineWizards, error) {
	aows := &AccountOnlineWizards{}
	if err := c.SearchRead(AccountOnlineWizardModel, criteria, options, aows); err != nil {
		return nil, err
	}
	return aows, nil
}

// FindAccountOnlineWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountOnlineWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountOnlineWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountOnlineWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountOnlineWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountOnlineWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.online.wizard was not found")
}

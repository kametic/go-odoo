package odoo

import (
	"fmt"
)

// AccountSepaCreditTransfer represents account.sepa.credit.transfer model.
type AccountSepaCreditTransfer struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	BankAccountId  *Many2One `xmlrpc:"bank_account_id,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	File           *String   `xmlrpc:"file,omptempty"`
	Filename       *String   `xmlrpc:"filename,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	IsGeneric      *Bool     `xmlrpc:"is_generic,omptempty"`
	JournalId      *Many2One `xmlrpc:"journal_id,omptempty"`
	WarningMessage *String   `xmlrpc:"warning_message,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// AccountSepaCreditTransfers represents array of account.sepa.credit.transfer model.
type AccountSepaCreditTransfers []AccountSepaCreditTransfer

// AccountSepaCreditTransferModel is the odoo model name.
const AccountSepaCreditTransferModel = "account.sepa.credit.transfer"

// Many2One convert AccountSepaCreditTransfer to *Many2One.
func (asct *AccountSepaCreditTransfer) Many2One() *Many2One {
	return NewMany2One(asct.Id.Get(), "")
}

// CreateAccountSepaCreditTransfer creates a new account.sepa.credit.transfer model and returns its id.
func (c *Client) CreateAccountSepaCreditTransfer(asct *AccountSepaCreditTransfer) (int64, error) {
	return c.Create(AccountSepaCreditTransferModel, asct)
}

// UpdateAccountSepaCreditTransfer updates an existing account.sepa.credit.transfer record.
func (c *Client) UpdateAccountSepaCreditTransfer(asct *AccountSepaCreditTransfer) error {
	return c.UpdateAccountSepaCreditTransfers([]int64{asct.Id.Get()}, asct)
}

// UpdateAccountSepaCreditTransfers updates existing account.sepa.credit.transfer records.
// All records (represented by ids) will be updated by asct values.
func (c *Client) UpdateAccountSepaCreditTransfers(ids []int64, asct *AccountSepaCreditTransfer) error {
	return c.Update(AccountSepaCreditTransferModel, ids, asct)
}

// DeleteAccountSepaCreditTransfer deletes an existing account.sepa.credit.transfer record.
func (c *Client) DeleteAccountSepaCreditTransfer(id int64) error {
	return c.DeleteAccountSepaCreditTransfers([]int64{id})
}

// DeleteAccountSepaCreditTransfers deletes existing account.sepa.credit.transfer records.
func (c *Client) DeleteAccountSepaCreditTransfers(ids []int64) error {
	return c.Delete(AccountSepaCreditTransferModel, ids)
}

// GetAccountSepaCreditTransfer gets account.sepa.credit.transfer existing record.
func (c *Client) GetAccountSepaCreditTransfer(id int64) (*AccountSepaCreditTransfer, error) {
	ascts, err := c.GetAccountSepaCreditTransfers([]int64{id})
	if err != nil {
		return nil, err
	}
	if ascts != nil && len(*ascts) > 0 {
		return &((*ascts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of account.sepa.credit.transfer not found", id)
}

// GetAccountSepaCreditTransfers gets account.sepa.credit.transfer existing records.
func (c *Client) GetAccountSepaCreditTransfers(ids []int64) (*AccountSepaCreditTransfers, error) {
	ascts := &AccountSepaCreditTransfers{}
	if err := c.Read(AccountSepaCreditTransferModel, ids, nil, ascts); err != nil {
		return nil, err
	}
	return ascts, nil
}

// FindAccountSepaCreditTransfer finds account.sepa.credit.transfer record by querying it with criteria.
func (c *Client) FindAccountSepaCreditTransfer(criteria *Criteria) (*AccountSepaCreditTransfer, error) {
	ascts := &AccountSepaCreditTransfers{}
	if err := c.SearchRead(AccountSepaCreditTransferModel, criteria, NewOptions().Limit(1), ascts); err != nil {
		return nil, err
	}
	if ascts != nil && len(*ascts) > 0 {
		return &((*ascts)[0]), nil
	}
	return nil, fmt.Errorf("account.sepa.credit.transfer was not found")
}

// FindAccountSepaCreditTransfers finds account.sepa.credit.transfer records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountSepaCreditTransfers(criteria *Criteria, options *Options) (*AccountSepaCreditTransfers, error) {
	ascts := &AccountSepaCreditTransfers{}
	if err := c.SearchRead(AccountSepaCreditTransferModel, criteria, options, ascts); err != nil {
		return nil, err
	}
	return ascts, nil
}

// FindAccountSepaCreditTransferIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountSepaCreditTransferIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AccountSepaCreditTransferModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAccountSepaCreditTransferId finds record id by querying it with criteria.
func (c *Client) FindAccountSepaCreditTransferId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountSepaCreditTransferModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("account.sepa.credit.transfer was not found")
}

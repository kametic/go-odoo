package odoo

import (
	"fmt"
)

// PosMercuryMercuryTransaction represents pos_mercury.mercury_transaction model.
type PosMercuryMercuryTransaction struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// PosMercuryMercuryTransactions represents array of pos_mercury.mercury_transaction model.
type PosMercuryMercuryTransactions []PosMercuryMercuryTransaction

// PosMercuryMercuryTransactionModel is the odoo model name.
const PosMercuryMercuryTransactionModel = "pos_mercury.mercury_transaction"

// Many2One convert PosMercuryMercuryTransaction to *Many2One.
func (pm *PosMercuryMercuryTransaction) Many2One() *Many2One {
	return NewMany2One(pm.Id.Get(), "")
}

// CreatePosMercuryMercuryTransaction creates a new pos_mercury.mercury_transaction model and returns its id.
func (c *Client) CreatePosMercuryMercuryTransaction(pm *PosMercuryMercuryTransaction) (int64, error) {
	return c.Create(PosMercuryMercuryTransactionModel, pm)
}

// UpdatePosMercuryMercuryTransaction updates an existing pos_mercury.mercury_transaction record.
func (c *Client) UpdatePosMercuryMercuryTransaction(pm *PosMercuryMercuryTransaction) error {
	return c.UpdatePosMercuryMercuryTransactions([]int64{pm.Id.Get()}, pm)
}

// UpdatePosMercuryMercuryTransactions updates existing pos_mercury.mercury_transaction records.
// All records (represented by ids) will be updated by pm values.
func (c *Client) UpdatePosMercuryMercuryTransactions(ids []int64, pm *PosMercuryMercuryTransaction) error {
	return c.Update(PosMercuryMercuryTransactionModel, ids, pm)
}

// DeletePosMercuryMercuryTransaction deletes an existing pos_mercury.mercury_transaction record.
func (c *Client) DeletePosMercuryMercuryTransaction(id int64) error {
	return c.DeletePosMercuryMercuryTransactions([]int64{id})
}

// DeletePosMercuryMercuryTransactions deletes existing pos_mercury.mercury_transaction records.
func (c *Client) DeletePosMercuryMercuryTransactions(ids []int64) error {
	return c.Delete(PosMercuryMercuryTransactionModel, ids)
}

// GetPosMercuryMercuryTransaction gets pos_mercury.mercury_transaction existing record.
func (c *Client) GetPosMercuryMercuryTransaction(id int64) (*PosMercuryMercuryTransaction, error) {
	pms, err := c.GetPosMercuryMercuryTransactions([]int64{id})
	if err != nil {
		return nil, err
	}
	if pms != nil && len(*pms) > 0 {
		return &((*pms)[0]), nil
	}
	return nil, fmt.Errorf("id %v of pos_mercury.mercury_transaction not found", id)
}

// GetPosMercuryMercuryTransactions gets pos_mercury.mercury_transaction existing records.
func (c *Client) GetPosMercuryMercuryTransactions(ids []int64) (*PosMercuryMercuryTransactions, error) {
	pms := &PosMercuryMercuryTransactions{}
	if err := c.Read(PosMercuryMercuryTransactionModel, ids, nil, pms); err != nil {
		return nil, err
	}
	return pms, nil
}

// FindPosMercuryMercuryTransaction finds pos_mercury.mercury_transaction record by querying it with criteria.
func (c *Client) FindPosMercuryMercuryTransaction(criteria *Criteria) (*PosMercuryMercuryTransaction, error) {
	pms := &PosMercuryMercuryTransactions{}
	if err := c.SearchRead(PosMercuryMercuryTransactionModel, criteria, NewOptions().Limit(1), pms); err != nil {
		return nil, err
	}
	if pms != nil && len(*pms) > 0 {
		return &((*pms)[0]), nil
	}
	return nil, fmt.Errorf("pos_mercury.mercury_transaction was not found")
}

// FindPosMercuryMercuryTransactions finds pos_mercury.mercury_transaction records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosMercuryMercuryTransactions(criteria *Criteria, options *Options) (*PosMercuryMercuryTransactions, error) {
	pms := &PosMercuryMercuryTransactions{}
	if err := c.SearchRead(PosMercuryMercuryTransactionModel, criteria, options, pms); err != nil {
		return nil, err
	}
	return pms, nil
}

// FindPosMercuryMercuryTransactionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosMercuryMercuryTransactionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PosMercuryMercuryTransactionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPosMercuryMercuryTransactionId finds record id by querying it with criteria.
func (c *Client) FindPosMercuryMercuryTransactionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosMercuryMercuryTransactionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("pos_mercury.mercury_transaction was not found")
}

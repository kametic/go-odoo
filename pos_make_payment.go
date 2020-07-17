package odoo

import (
	"fmt"
)

// PosMakePayment represents pos.make.payment model.
type PosMakePayment struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Amount      *Float    `xmlrpc:"amount,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	JournalId   *Many2One `xmlrpc:"journal_id,omptempty"`
	PaymentDate *Time     `xmlrpc:"payment_date,omptempty"`
	PaymentName *String   `xmlrpc:"payment_name,omptempty"`
	SessionId   *Many2One `xmlrpc:"session_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// PosMakePayments represents array of pos.make.payment model.
type PosMakePayments []PosMakePayment

// PosMakePaymentModel is the odoo model name.
const PosMakePaymentModel = "pos.make.payment"

// Many2One convert PosMakePayment to *Many2One.
func (pmp *PosMakePayment) Many2One() *Many2One {
	return NewMany2One(pmp.Id.Get(), "")
}

// CreatePosMakePayment creates a new pos.make.payment model and returns its id.
func (c *Client) CreatePosMakePayment(pmp *PosMakePayment) (int64, error) {
	return c.Create(PosMakePaymentModel, pmp)
}

// UpdatePosMakePayment updates an existing pos.make.payment record.
func (c *Client) UpdatePosMakePayment(pmp *PosMakePayment) error {
	return c.UpdatePosMakePayments([]int64{pmp.Id.Get()}, pmp)
}

// UpdatePosMakePayments updates existing pos.make.payment records.
// All records (represented by ids) will be updated by pmp values.
func (c *Client) UpdatePosMakePayments(ids []int64, pmp *PosMakePayment) error {
	return c.Update(PosMakePaymentModel, ids, pmp)
}

// DeletePosMakePayment deletes an existing pos.make.payment record.
func (c *Client) DeletePosMakePayment(id int64) error {
	return c.DeletePosMakePayments([]int64{id})
}

// DeletePosMakePayments deletes existing pos.make.payment records.
func (c *Client) DeletePosMakePayments(ids []int64) error {
	return c.Delete(PosMakePaymentModel, ids)
}

// GetPosMakePayment gets pos.make.payment existing record.
func (c *Client) GetPosMakePayment(id int64) (*PosMakePayment, error) {
	pmps, err := c.GetPosMakePayments([]int64{id})
	if err != nil {
		return nil, err
	}
	if pmps != nil && len(*pmps) > 0 {
		return &((*pmps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of pos.make.payment not found", id)
}

// GetPosMakePayments gets pos.make.payment existing records.
func (c *Client) GetPosMakePayments(ids []int64) (*PosMakePayments, error) {
	pmps := &PosMakePayments{}
	if err := c.Read(PosMakePaymentModel, ids, nil, pmps); err != nil {
		return nil, err
	}
	return pmps, nil
}

// FindPosMakePayment finds pos.make.payment record by querying it with criteria.
func (c *Client) FindPosMakePayment(criteria *Criteria) (*PosMakePayment, error) {
	pmps := &PosMakePayments{}
	if err := c.SearchRead(PosMakePaymentModel, criteria, NewOptions().Limit(1), pmps); err != nil {
		return nil, err
	}
	if pmps != nil && len(*pmps) > 0 {
		return &((*pmps)[0]), nil
	}
	return nil, fmt.Errorf("pos.make.payment was not found")
}

// FindPosMakePayments finds pos.make.payment records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosMakePayments(criteria *Criteria, options *Options) (*PosMakePayments, error) {
	pmps := &PosMakePayments{}
	if err := c.SearchRead(PosMakePaymentModel, criteria, options, pmps); err != nil {
		return nil, err
	}
	return pmps, nil
}

// FindPosMakePaymentIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosMakePaymentIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PosMakePaymentModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPosMakePaymentId finds record id by querying it with criteria.
func (c *Client) FindPosMakePaymentId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosMakePaymentModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("pos.make.payment was not found")
}

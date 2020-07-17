package odoo

import (
	"fmt"
)

// PosDiscount represents pos.discount model.
type PosDiscount struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	Discount    *Float    `xmlrpc:"discount,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// PosDiscounts represents array of pos.discount model.
type PosDiscounts []PosDiscount

// PosDiscountModel is the odoo model name.
const PosDiscountModel = "pos.discount"

// Many2One convert PosDiscount to *Many2One.
func (pd *PosDiscount) Many2One() *Many2One {
	return NewMany2One(pd.Id.Get(), "")
}

// CreatePosDiscount creates a new pos.discount model and returns its id.
func (c *Client) CreatePosDiscount(pd *PosDiscount) (int64, error) {
	return c.Create(PosDiscountModel, pd)
}

// UpdatePosDiscount updates an existing pos.discount record.
func (c *Client) UpdatePosDiscount(pd *PosDiscount) error {
	return c.UpdatePosDiscounts([]int64{pd.Id.Get()}, pd)
}

// UpdatePosDiscounts updates existing pos.discount records.
// All records (represented by ids) will be updated by pd values.
func (c *Client) UpdatePosDiscounts(ids []int64, pd *PosDiscount) error {
	return c.Update(PosDiscountModel, ids, pd)
}

// DeletePosDiscount deletes an existing pos.discount record.
func (c *Client) DeletePosDiscount(id int64) error {
	return c.DeletePosDiscounts([]int64{id})
}

// DeletePosDiscounts deletes existing pos.discount records.
func (c *Client) DeletePosDiscounts(ids []int64) error {
	return c.Delete(PosDiscountModel, ids)
}

// GetPosDiscount gets pos.discount existing record.
func (c *Client) GetPosDiscount(id int64) (*PosDiscount, error) {
	pds, err := c.GetPosDiscounts([]int64{id})
	if err != nil {
		return nil, err
	}
	if pds != nil && len(*pds) > 0 {
		return &((*pds)[0]), nil
	}
	return nil, fmt.Errorf("id %v of pos.discount not found", id)
}

// GetPosDiscounts gets pos.discount existing records.
func (c *Client) GetPosDiscounts(ids []int64) (*PosDiscounts, error) {
	pds := &PosDiscounts{}
	if err := c.Read(PosDiscountModel, ids, nil, pds); err != nil {
		return nil, err
	}
	return pds, nil
}

// FindPosDiscount finds pos.discount record by querying it with criteria.
func (c *Client) FindPosDiscount(criteria *Criteria) (*PosDiscount, error) {
	pds := &PosDiscounts{}
	if err := c.SearchRead(PosDiscountModel, criteria, NewOptions().Limit(1), pds); err != nil {
		return nil, err
	}
	if pds != nil && len(*pds) > 0 {
		return &((*pds)[0]), nil
	}
	return nil, fmt.Errorf("pos.discount was not found")
}

// FindPosDiscounts finds pos.discount records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosDiscounts(criteria *Criteria, options *Options) (*PosDiscounts, error) {
	pds := &PosDiscounts{}
	if err := c.SearchRead(PosDiscountModel, criteria, options, pds); err != nil {
		return nil, err
	}
	return pds, nil
}

// FindPosDiscountIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosDiscountIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PosDiscountModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPosDiscountId finds record id by querying it with criteria.
func (c *Client) FindPosDiscountId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosDiscountModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("pos.discount was not found")
}

package odoo

import (
	"fmt"
)

// PosCategory represents pos.category model.
type PosCategory struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	ChildId     *Relation `xmlrpc:"child_id,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Image       *String   `xmlrpc:"image,omptempty"`
	ImageMedium *String   `xmlrpc:"image_medium,omptempty"`
	ImageSmall  *String   `xmlrpc:"image_small,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	ParentId    *Many2One `xmlrpc:"parent_id,omptempty"`
	Sequence    *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// PosCategorys represents array of pos.category model.
type PosCategorys []PosCategory

// PosCategoryModel is the odoo model name.
const PosCategoryModel = "pos.category"

// Many2One convert PosCategory to *Many2One.
func (pc *PosCategory) Many2One() *Many2One {
	return NewMany2One(pc.Id.Get(), "")
}

// CreatePosCategory creates a new pos.category model and returns its id.
func (c *Client) CreatePosCategory(pc *PosCategory) (int64, error) {
	return c.Create(PosCategoryModel, pc)
}

// UpdatePosCategory updates an existing pos.category record.
func (c *Client) UpdatePosCategory(pc *PosCategory) error {
	return c.UpdatePosCategorys([]int64{pc.Id.Get()}, pc)
}

// UpdatePosCategorys updates existing pos.category records.
// All records (represented by ids) will be updated by pc values.
func (c *Client) UpdatePosCategorys(ids []int64, pc *PosCategory) error {
	return c.Update(PosCategoryModel, ids, pc)
}

// DeletePosCategory deletes an existing pos.category record.
func (c *Client) DeletePosCategory(id int64) error {
	return c.DeletePosCategorys([]int64{id})
}

// DeletePosCategorys deletes existing pos.category records.
func (c *Client) DeletePosCategorys(ids []int64) error {
	return c.Delete(PosCategoryModel, ids)
}

// GetPosCategory gets pos.category existing record.
func (c *Client) GetPosCategory(id int64) (*PosCategory, error) {
	pcs, err := c.GetPosCategorys([]int64{id})
	if err != nil {
		return nil, err
	}
	if pcs != nil && len(*pcs) > 0 {
		return &((*pcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of pos.category not found", id)
}

// GetPosCategorys gets pos.category existing records.
func (c *Client) GetPosCategorys(ids []int64) (*PosCategorys, error) {
	pcs := &PosCategorys{}
	if err := c.Read(PosCategoryModel, ids, nil, pcs); err != nil {
		return nil, err
	}
	return pcs, nil
}

// FindPosCategory finds pos.category record by querying it with criteria.
func (c *Client) FindPosCategory(criteria *Criteria) (*PosCategory, error) {
	pcs := &PosCategorys{}
	if err := c.SearchRead(PosCategoryModel, criteria, NewOptions().Limit(1), pcs); err != nil {
		return nil, err
	}
	if pcs != nil && len(*pcs) > 0 {
		return &((*pcs)[0]), nil
	}
	return nil, fmt.Errorf("pos.category was not found")
}

// FindPosCategorys finds pos.category records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosCategorys(criteria *Criteria, options *Options) (*PosCategorys, error) {
	pcs := &PosCategorys{}
	if err := c.SearchRead(PosCategoryModel, criteria, options, pcs); err != nil {
		return nil, err
	}
	return pcs, nil
}

// FindPosCategoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosCategoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PosCategoryModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPosCategoryId finds record id by querying it with criteria.
func (c *Client) FindPosCategoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosCategoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("pos.category was not found")
}

package odoo

import (
	"fmt"
)

// XMarge represents x_marge model.
type XMarge struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
	XName       *String   `xmlrpc:"x_name,omptempty"`
}

// XMarges represents array of x_marge model.
type XMarges []XMarge

// XMargeModel is the odoo model name.
const XMargeModel = "x_marge"

// Many2One convert XMarge to *Many2One.
func (x *XMarge) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXMarge creates a new x_marge model and returns its id.
func (c *Client) CreateXMarge(x *XMarge) (int64, error) {
	return c.Create(XMargeModel, x)
}

// UpdateXMarge updates an existing x_marge record.
func (c *Client) UpdateXMarge(x *XMarge) error {
	return c.UpdateXMarges([]int64{x.Id.Get()}, x)
}

// UpdateXMarges updates existing x_marge records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXMarges(ids []int64, x *XMarge) error {
	return c.Update(XMargeModel, ids, x)
}

// DeleteXMarge deletes an existing x_marge record.
func (c *Client) DeleteXMarge(id int64) error {
	return c.DeleteXMarges([]int64{id})
}

// DeleteXMarges deletes existing x_marge records.
func (c *Client) DeleteXMarges(ids []int64) error {
	return c.Delete(XMargeModel, ids)
}

// GetXMarge gets x_marge existing record.
func (c *Client) GetXMarge(id int64) (*XMarge, error) {
	xs, err := c.GetXMarges([]int64{id})
	if err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of x_marge not found", id)
}

// GetXMarges gets x_marge existing records.
func (c *Client) GetXMarges(ids []int64) (*XMarges, error) {
	xs := &XMarges{}
	if err := c.Read(XMargeModel, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXMarge finds x_marge record by querying it with criteria.
func (c *Client) FindXMarge(criteria *Criteria) (*XMarge, error) {
	xs := &XMarges{}
	if err := c.SearchRead(XMargeModel, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("x_marge was not found")
}

// FindXMarges finds x_marge records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXMarges(criteria *Criteria, options *Options) (*XMarges, error) {
	xs := &XMarges{}
	if err := c.SearchRead(XMargeModel, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXMargeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXMargeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(XMargeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindXMargeId finds record id by querying it with criteria.
func (c *Client) FindXMargeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XMargeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("x_marge was not found")
}

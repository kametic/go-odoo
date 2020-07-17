package odoo

import (
	"fmt"
)

// XMesBenefices represents x_mes_benefices model.
type XMesBenefices struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
	XName       *String   `xmlrpc:"x_name,omptempty"`
}

// XMesBeneficess represents array of x_mes_benefices model.
type XMesBeneficess []XMesBenefices

// XMesBeneficesModel is the odoo model name.
const XMesBeneficesModel = "x_mes_benefices"

// Many2One convert XMesBenefices to *Many2One.
func (x *XMesBenefices) Many2One() *Many2One {
	return NewMany2One(x.Id.Get(), "")
}

// CreateXMesBenefices creates a new x_mes_benefices model and returns its id.
func (c *Client) CreateXMesBenefices(x *XMesBenefices) (int64, error) {
	return c.Create(XMesBeneficesModel, x)
}

// UpdateXMesBenefices updates an existing x_mes_benefices record.
func (c *Client) UpdateXMesBenefices(x *XMesBenefices) error {
	return c.UpdateXMesBeneficess([]int64{x.Id.Get()}, x)
}

// UpdateXMesBeneficess updates existing x_mes_benefices records.
// All records (represented by ids) will be updated by x values.
func (c *Client) UpdateXMesBeneficess(ids []int64, x *XMesBenefices) error {
	return c.Update(XMesBeneficesModel, ids, x)
}

// DeleteXMesBenefices deletes an existing x_mes_benefices record.
func (c *Client) DeleteXMesBenefices(id int64) error {
	return c.DeleteXMesBeneficess([]int64{id})
}

// DeleteXMesBeneficess deletes existing x_mes_benefices records.
func (c *Client) DeleteXMesBeneficess(ids []int64) error {
	return c.Delete(XMesBeneficesModel, ids)
}

// GetXMesBenefices gets x_mes_benefices existing record.
func (c *Client) GetXMesBenefices(id int64) (*XMesBenefices, error) {
	xs, err := c.GetXMesBeneficess([]int64{id})
	if err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of x_mes_benefices not found", id)
}

// GetXMesBeneficess gets x_mes_benefices existing records.
func (c *Client) GetXMesBeneficess(ids []int64) (*XMesBeneficess, error) {
	xs := &XMesBeneficess{}
	if err := c.Read(XMesBeneficesModel, ids, nil, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXMesBenefices finds x_mes_benefices record by querying it with criteria.
func (c *Client) FindXMesBenefices(criteria *Criteria) (*XMesBenefices, error) {
	xs := &XMesBeneficess{}
	if err := c.SearchRead(XMesBeneficesModel, criteria, NewOptions().Limit(1), xs); err != nil {
		return nil, err
	}
	if xs != nil && len(*xs) > 0 {
		return &((*xs)[0]), nil
	}
	return nil, fmt.Errorf("x_mes_benefices was not found")
}

// FindXMesBeneficess finds x_mes_benefices records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXMesBeneficess(criteria *Criteria, options *Options) (*XMesBeneficess, error) {
	xs := &XMesBeneficess{}
	if err := c.SearchRead(XMesBeneficesModel, criteria, options, xs); err != nil {
		return nil, err
	}
	return xs, nil
}

// FindXMesBeneficesIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXMesBeneficesIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(XMesBeneficesModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindXMesBeneficesId finds record id by querying it with criteria.
func (c *Client) FindXMesBeneficesId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XMesBeneficesModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("x_mes_benefices was not found")
}

package odoo

import (
	"fmt"
)

// XPosConfigStage represents x_pos.config_stage model.
type XPosConfigStage struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
	XName       *String   `xmlrpc:"x_name,omptempty"`
}

// XPosConfigStages represents array of x_pos.config_stage model.
type XPosConfigStages []XPosConfigStage

// XPosConfigStageModel is the odoo model name.
const XPosConfigStageModel = "x_pos.config_stage"

// Many2One convert XPosConfigStage to *Many2One.
func (xc *XPosConfigStage) Many2One() *Many2One {
	return NewMany2One(xc.Id.Get(), "")
}

// CreateXPosConfigStage creates a new x_pos.config_stage model and returns its id.
func (c *Client) CreateXPosConfigStage(xc *XPosConfigStage) (int64, error) {
	return c.Create(XPosConfigStageModel, xc)
}

// UpdateXPosConfigStage updates an existing x_pos.config_stage record.
func (c *Client) UpdateXPosConfigStage(xc *XPosConfigStage) error {
	return c.UpdateXPosConfigStages([]int64{xc.Id.Get()}, xc)
}

// UpdateXPosConfigStages updates existing x_pos.config_stage records.
// All records (represented by ids) will be updated by xc values.
func (c *Client) UpdateXPosConfigStages(ids []int64, xc *XPosConfigStage) error {
	return c.Update(XPosConfigStageModel, ids, xc)
}

// DeleteXPosConfigStage deletes an existing x_pos.config_stage record.
func (c *Client) DeleteXPosConfigStage(id int64) error {
	return c.DeleteXPosConfigStages([]int64{id})
}

// DeleteXPosConfigStages deletes existing x_pos.config_stage records.
func (c *Client) DeleteXPosConfigStages(ids []int64) error {
	return c.Delete(XPosConfigStageModel, ids)
}

// GetXPosConfigStage gets x_pos.config_stage existing record.
func (c *Client) GetXPosConfigStage(id int64) (*XPosConfigStage, error) {
	xcs, err := c.GetXPosConfigStages([]int64{id})
	if err != nil {
		return nil, err
	}
	if xcs != nil && len(*xcs) > 0 {
		return &((*xcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of x_pos.config_stage not found", id)
}

// GetXPosConfigStages gets x_pos.config_stage existing records.
func (c *Client) GetXPosConfigStages(ids []int64) (*XPosConfigStages, error) {
	xcs := &XPosConfigStages{}
	if err := c.Read(XPosConfigStageModel, ids, nil, xcs); err != nil {
		return nil, err
	}
	return xcs, nil
}

// FindXPosConfigStage finds x_pos.config_stage record by querying it with criteria.
func (c *Client) FindXPosConfigStage(criteria *Criteria) (*XPosConfigStage, error) {
	xcs := &XPosConfigStages{}
	if err := c.SearchRead(XPosConfigStageModel, criteria, NewOptions().Limit(1), xcs); err != nil {
		return nil, err
	}
	if xcs != nil && len(*xcs) > 0 {
		return &((*xcs)[0]), nil
	}
	return nil, fmt.Errorf("x_pos.config_stage was not found")
}

// FindXPosConfigStages finds x_pos.config_stage records by querying it
// and filtering it with criteria and options.
func (c *Client) FindXPosConfigStages(criteria *Criteria, options *Options) (*XPosConfigStages, error) {
	xcs := &XPosConfigStages{}
	if err := c.SearchRead(XPosConfigStageModel, criteria, options, xcs); err != nil {
		return nil, err
	}
	return xcs, nil
}

// FindXPosConfigStageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindXPosConfigStageIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(XPosConfigStageModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindXPosConfigStageId finds record id by querying it with criteria.
func (c *Client) FindXPosConfigStageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(XPosConfigStageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("x_pos.config_stage was not found")
}

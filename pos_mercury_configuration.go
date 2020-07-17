package odoo

import (
	"fmt"
)

// PosMercuryConfiguration represents pos_mercury.configuration model.
type PosMercuryConfiguration struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	MerchantId  *String   `xmlrpc:"merchant_id,omptempty"`
	MerchantPwd *String   `xmlrpc:"merchant_pwd,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// PosMercuryConfigurations represents array of pos_mercury.configuration model.
type PosMercuryConfigurations []PosMercuryConfiguration

// PosMercuryConfigurationModel is the odoo model name.
const PosMercuryConfigurationModel = "pos_mercury.configuration"

// Many2One convert PosMercuryConfiguration to *Many2One.
func (pc *PosMercuryConfiguration) Many2One() *Many2One {
	return NewMany2One(pc.Id.Get(), "")
}

// CreatePosMercuryConfiguration creates a new pos_mercury.configuration model and returns its id.
func (c *Client) CreatePosMercuryConfiguration(pc *PosMercuryConfiguration) (int64, error) {
	return c.Create(PosMercuryConfigurationModel, pc)
}

// UpdatePosMercuryConfiguration updates an existing pos_mercury.configuration record.
func (c *Client) UpdatePosMercuryConfiguration(pc *PosMercuryConfiguration) error {
	return c.UpdatePosMercuryConfigurations([]int64{pc.Id.Get()}, pc)
}

// UpdatePosMercuryConfigurations updates existing pos_mercury.configuration records.
// All records (represented by ids) will be updated by pc values.
func (c *Client) UpdatePosMercuryConfigurations(ids []int64, pc *PosMercuryConfiguration) error {
	return c.Update(PosMercuryConfigurationModel, ids, pc)
}

// DeletePosMercuryConfiguration deletes an existing pos_mercury.configuration record.
func (c *Client) DeletePosMercuryConfiguration(id int64) error {
	return c.DeletePosMercuryConfigurations([]int64{id})
}

// DeletePosMercuryConfigurations deletes existing pos_mercury.configuration records.
func (c *Client) DeletePosMercuryConfigurations(ids []int64) error {
	return c.Delete(PosMercuryConfigurationModel, ids)
}

// GetPosMercuryConfiguration gets pos_mercury.configuration existing record.
func (c *Client) GetPosMercuryConfiguration(id int64) (*PosMercuryConfiguration, error) {
	pcs, err := c.GetPosMercuryConfigurations([]int64{id})
	if err != nil {
		return nil, err
	}
	if pcs != nil && len(*pcs) > 0 {
		return &((*pcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of pos_mercury.configuration not found", id)
}

// GetPosMercuryConfigurations gets pos_mercury.configuration existing records.
func (c *Client) GetPosMercuryConfigurations(ids []int64) (*PosMercuryConfigurations, error) {
	pcs := &PosMercuryConfigurations{}
	if err := c.Read(PosMercuryConfigurationModel, ids, nil, pcs); err != nil {
		return nil, err
	}
	return pcs, nil
}

// FindPosMercuryConfiguration finds pos_mercury.configuration record by querying it with criteria.
func (c *Client) FindPosMercuryConfiguration(criteria *Criteria) (*PosMercuryConfiguration, error) {
	pcs := &PosMercuryConfigurations{}
	if err := c.SearchRead(PosMercuryConfigurationModel, criteria, NewOptions().Limit(1), pcs); err != nil {
		return nil, err
	}
	if pcs != nil && len(*pcs) > 0 {
		return &((*pcs)[0]), nil
	}
	return nil, fmt.Errorf("pos_mercury.configuration was not found")
}

// FindPosMercuryConfigurations finds pos_mercury.configuration records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosMercuryConfigurations(criteria *Criteria, options *Options) (*PosMercuryConfigurations, error) {
	pcs := &PosMercuryConfigurations{}
	if err := c.SearchRead(PosMercuryConfigurationModel, criteria, options, pcs); err != nil {
		return nil, err
	}
	return pcs, nil
}

// FindPosMercuryConfigurationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosMercuryConfigurationIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PosMercuryConfigurationModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPosMercuryConfigurationId finds record id by querying it with criteria.
func (c *Client) FindPosMercuryConfigurationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosMercuryConfigurationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("pos_mercury.configuration was not found")
}

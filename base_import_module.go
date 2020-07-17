package odoo

import (
	"fmt"
)

// BaseImportModule represents base.import.module model.
type BaseImportModule struct {
	LastUpdate    *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName   *String    `xmlrpc:"display_name,omptempty"`
	Force         *Bool      `xmlrpc:"force,omptempty"`
	Id            *Int       `xmlrpc:"id,omptempty"`
	ImportMessage *String    `xmlrpc:"import_message,omptempty"`
	ModuleFile    *String    `xmlrpc:"module_file,omptempty"`
	State         *Selection `xmlrpc:"state,omptempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// BaseImportModules represents array of base.import.module model.
type BaseImportModules []BaseImportModule

// BaseImportModuleModel is the odoo model name.
const BaseImportModuleModel = "base.import.module"

// Many2One convert BaseImportModule to *Many2One.
func (bim *BaseImportModule) Many2One() *Many2One {
	return NewMany2One(bim.Id.Get(), "")
}

// CreateBaseImportModule creates a new base.import.module model and returns its id.
func (c *Client) CreateBaseImportModule(bim *BaseImportModule) (int64, error) {
	return c.Create(BaseImportModuleModel, bim)
}

// UpdateBaseImportModule updates an existing base.import.module record.
func (c *Client) UpdateBaseImportModule(bim *BaseImportModule) error {
	return c.UpdateBaseImportModules([]int64{bim.Id.Get()}, bim)
}

// UpdateBaseImportModules updates existing base.import.module records.
// All records (represented by ids) will be updated by bim values.
func (c *Client) UpdateBaseImportModules(ids []int64, bim *BaseImportModule) error {
	return c.Update(BaseImportModuleModel, ids, bim)
}

// DeleteBaseImportModule deletes an existing base.import.module record.
func (c *Client) DeleteBaseImportModule(id int64) error {
	return c.DeleteBaseImportModules([]int64{id})
}

// DeleteBaseImportModules deletes existing base.import.module records.
func (c *Client) DeleteBaseImportModules(ids []int64) error {
	return c.Delete(BaseImportModuleModel, ids)
}

// GetBaseImportModule gets base.import.module existing record.
func (c *Client) GetBaseImportModule(id int64) (*BaseImportModule, error) {
	bims, err := c.GetBaseImportModules([]int64{id})
	if err != nil {
		return nil, err
	}
	if bims != nil && len(*bims) > 0 {
		return &((*bims)[0]), nil
	}
	return nil, fmt.Errorf("id %v of base.import.module not found", id)
}

// GetBaseImportModules gets base.import.module existing records.
func (c *Client) GetBaseImportModules(ids []int64) (*BaseImportModules, error) {
	bims := &BaseImportModules{}
	if err := c.Read(BaseImportModuleModel, ids, nil, bims); err != nil {
		return nil, err
	}
	return bims, nil
}

// FindBaseImportModule finds base.import.module record by querying it with criteria.
func (c *Client) FindBaseImportModule(criteria *Criteria) (*BaseImportModule, error) {
	bims := &BaseImportModules{}
	if err := c.SearchRead(BaseImportModuleModel, criteria, NewOptions().Limit(1), bims); err != nil {
		return nil, err
	}
	if bims != nil && len(*bims) > 0 {
		return &((*bims)[0]), nil
	}
	return nil, fmt.Errorf("base.import.module was not found")
}

// FindBaseImportModules finds base.import.module records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportModules(criteria *Criteria, options *Options) (*BaseImportModules, error) {
	bims := &BaseImportModules{}
	if err := c.SearchRead(BaseImportModuleModel, criteria, options, bims); err != nil {
		return nil, err
	}
	return bims, nil
}

// FindBaseImportModuleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseImportModuleIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BaseImportModuleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBaseImportModuleId finds record id by querying it with criteria.
func (c *Client) FindBaseImportModuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseImportModuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("base.import.module was not found")
}

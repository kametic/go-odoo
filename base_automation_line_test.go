package odoo

import (
	"fmt"
)

// BaseAutomationLineTest represents base.automation.line.test model.
type BaseAutomationLineTest struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	LeadId      *Many2One `xmlrpc:"lead_id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	UserId      *Many2One `xmlrpc:"user_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// BaseAutomationLineTests represents array of base.automation.line.test model.
type BaseAutomationLineTests []BaseAutomationLineTest

// BaseAutomationLineTestModel is the odoo model name.
const BaseAutomationLineTestModel = "base.automation.line.test"

// Many2One convert BaseAutomationLineTest to *Many2One.
func (balt *BaseAutomationLineTest) Many2One() *Many2One {
	return NewMany2One(balt.Id.Get(), "")
}

// CreateBaseAutomationLineTest creates a new base.automation.line.test model and returns its id.
func (c *Client) CreateBaseAutomationLineTest(balt *BaseAutomationLineTest) (int64, error) {
	return c.Create(BaseAutomationLineTestModel, balt)
}

// UpdateBaseAutomationLineTest updates an existing base.automation.line.test record.
func (c *Client) UpdateBaseAutomationLineTest(balt *BaseAutomationLineTest) error {
	return c.UpdateBaseAutomationLineTests([]int64{balt.Id.Get()}, balt)
}

// UpdateBaseAutomationLineTests updates existing base.automation.line.test records.
// All records (represented by ids) will be updated by balt values.
func (c *Client) UpdateBaseAutomationLineTests(ids []int64, balt *BaseAutomationLineTest) error {
	return c.Update(BaseAutomationLineTestModel, ids, balt)
}

// DeleteBaseAutomationLineTest deletes an existing base.automation.line.test record.
func (c *Client) DeleteBaseAutomationLineTest(id int64) error {
	return c.DeleteBaseAutomationLineTests([]int64{id})
}

// DeleteBaseAutomationLineTests deletes existing base.automation.line.test records.
func (c *Client) DeleteBaseAutomationLineTests(ids []int64) error {
	return c.Delete(BaseAutomationLineTestModel, ids)
}

// GetBaseAutomationLineTest gets base.automation.line.test existing record.
func (c *Client) GetBaseAutomationLineTest(id int64) (*BaseAutomationLineTest, error) {
	balts, err := c.GetBaseAutomationLineTests([]int64{id})
	if err != nil {
		return nil, err
	}
	if balts != nil && len(*balts) > 0 {
		return &((*balts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of base.automation.line.test not found", id)
}

// GetBaseAutomationLineTests gets base.automation.line.test existing records.
func (c *Client) GetBaseAutomationLineTests(ids []int64) (*BaseAutomationLineTests, error) {
	balts := &BaseAutomationLineTests{}
	if err := c.Read(BaseAutomationLineTestModel, ids, nil, balts); err != nil {
		return nil, err
	}
	return balts, nil
}

// FindBaseAutomationLineTest finds base.automation.line.test record by querying it with criteria.
func (c *Client) FindBaseAutomationLineTest(criteria *Criteria) (*BaseAutomationLineTest, error) {
	balts := &BaseAutomationLineTests{}
	if err := c.SearchRead(BaseAutomationLineTestModel, criteria, NewOptions().Limit(1), balts); err != nil {
		return nil, err
	}
	if balts != nil && len(*balts) > 0 {
		return &((*balts)[0]), nil
	}
	return nil, fmt.Errorf("base.automation.line.test was not found")
}

// FindBaseAutomationLineTests finds base.automation.line.test records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseAutomationLineTests(criteria *Criteria, options *Options) (*BaseAutomationLineTests, error) {
	balts := &BaseAutomationLineTests{}
	if err := c.SearchRead(BaseAutomationLineTestModel, criteria, options, balts); err != nil {
		return nil, err
	}
	return balts, nil
}

// FindBaseAutomationLineTestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseAutomationLineTestIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BaseAutomationLineTestModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBaseAutomationLineTestId finds record id by querying it with criteria.
func (c *Client) FindBaseAutomationLineTestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseAutomationLineTestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("base.automation.line.test was not found")
}

package odoo

import (
	"fmt"
)

// BaseAutomationLeadTest represents base.automation.lead.test model.
type BaseAutomationLeadTest struct {
	LastUpdate        *Time      `xmlrpc:"__last_update,omptempty"`
	Active            *Bool      `xmlrpc:"active,omptempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omptempty"`
	Customer          *Bool      `xmlrpc:"customer,omptempty"`
	DateActionLast    *Time      `xmlrpc:"date_action_last,omptempty"`
	Deadline          *Bool      `xmlrpc:"deadline,omptempty"`
	DisplayName       *String    `xmlrpc:"display_name,omptempty"`
	Id                *Int       `xmlrpc:"id,omptempty"`
	IsAssignedToAdmin *Bool      `xmlrpc:"is_assigned_to_admin,omptempty"`
	LineIds           *Relation  `xmlrpc:"line_ids,omptempty"`
	Name              *String    `xmlrpc:"name,omptempty"`
	PartnerId         *Many2One  `xmlrpc:"partner_id,omptempty"`
	Priority          *Bool      `xmlrpc:"priority,omptempty"`
	State             *Selection `xmlrpc:"state,omptempty"`
	UserId            *Many2One  `xmlrpc:"user_id,omptempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// BaseAutomationLeadTests represents array of base.automation.lead.test model.
type BaseAutomationLeadTests []BaseAutomationLeadTest

// BaseAutomationLeadTestModel is the odoo model name.
const BaseAutomationLeadTestModel = "base.automation.lead.test"

// Many2One convert BaseAutomationLeadTest to *Many2One.
func (balt *BaseAutomationLeadTest) Many2One() *Many2One {
	return NewMany2One(balt.Id.Get(), "")
}

// CreateBaseAutomationLeadTest creates a new base.automation.lead.test model and returns its id.
func (c *Client) CreateBaseAutomationLeadTest(balt *BaseAutomationLeadTest) (int64, error) {
	return c.Create(BaseAutomationLeadTestModel, balt)
}

// UpdateBaseAutomationLeadTest updates an existing base.automation.lead.test record.
func (c *Client) UpdateBaseAutomationLeadTest(balt *BaseAutomationLeadTest) error {
	return c.UpdateBaseAutomationLeadTests([]int64{balt.Id.Get()}, balt)
}

// UpdateBaseAutomationLeadTests updates existing base.automation.lead.test records.
// All records (represented by ids) will be updated by balt values.
func (c *Client) UpdateBaseAutomationLeadTests(ids []int64, balt *BaseAutomationLeadTest) error {
	return c.Update(BaseAutomationLeadTestModel, ids, balt)
}

// DeleteBaseAutomationLeadTest deletes an existing base.automation.lead.test record.
func (c *Client) DeleteBaseAutomationLeadTest(id int64) error {
	return c.DeleteBaseAutomationLeadTests([]int64{id})
}

// DeleteBaseAutomationLeadTests deletes existing base.automation.lead.test records.
func (c *Client) DeleteBaseAutomationLeadTests(ids []int64) error {
	return c.Delete(BaseAutomationLeadTestModel, ids)
}

// GetBaseAutomationLeadTest gets base.automation.lead.test existing record.
func (c *Client) GetBaseAutomationLeadTest(id int64) (*BaseAutomationLeadTest, error) {
	balts, err := c.GetBaseAutomationLeadTests([]int64{id})
	if err != nil {
		return nil, err
	}
	if balts != nil && len(*balts) > 0 {
		return &((*balts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of base.automation.lead.test not found", id)
}

// GetBaseAutomationLeadTests gets base.automation.lead.test existing records.
func (c *Client) GetBaseAutomationLeadTests(ids []int64) (*BaseAutomationLeadTests, error) {
	balts := &BaseAutomationLeadTests{}
	if err := c.Read(BaseAutomationLeadTestModel, ids, nil, balts); err != nil {
		return nil, err
	}
	return balts, nil
}

// FindBaseAutomationLeadTest finds base.automation.lead.test record by querying it with criteria.
func (c *Client) FindBaseAutomationLeadTest(criteria *Criteria) (*BaseAutomationLeadTest, error) {
	balts := &BaseAutomationLeadTests{}
	if err := c.SearchRead(BaseAutomationLeadTestModel, criteria, NewOptions().Limit(1), balts); err != nil {
		return nil, err
	}
	if balts != nil && len(*balts) > 0 {
		return &((*balts)[0]), nil
	}
	return nil, fmt.Errorf("base.automation.lead.test was not found")
}

// FindBaseAutomationLeadTests finds base.automation.lead.test records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseAutomationLeadTests(criteria *Criteria, options *Options) (*BaseAutomationLeadTests, error) {
	balts := &BaseAutomationLeadTests{}
	if err := c.SearchRead(BaseAutomationLeadTestModel, criteria, options, balts); err != nil {
		return nil, err
	}
	return balts, nil
}

// FindBaseAutomationLeadTestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseAutomationLeadTestIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(BaseAutomationLeadTestModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindBaseAutomationLeadTestId finds record id by querying it with criteria.
func (c *Client) FindBaseAutomationLeadTestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseAutomationLeadTestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("base.automation.lead.test was not found")
}

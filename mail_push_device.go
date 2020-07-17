package odoo

import (
	"fmt"
)

// MailPushDevice represents mail_push.device model.
type MailPushDevice struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	Name           *String    `xmlrpc:"name,omptempty"`
	PartnerId      *Many2One  `xmlrpc:"partner_id,omptempty"`
	ServiceType    *Selection `xmlrpc:"service_type,omptempty"`
	SubscriptionId *String    `xmlrpc:"subscription_id,omptempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MailPushDevices represents array of mail_push.device model.
type MailPushDevices []MailPushDevice

// MailPushDeviceModel is the odoo model name.
const MailPushDeviceModel = "mail_push.device"

// Many2One convert MailPushDevice to *Many2One.
func (md *MailPushDevice) Many2One() *Many2One {
	return NewMany2One(md.Id.Get(), "")
}

// CreateMailPushDevice creates a new mail_push.device model and returns its id.
func (c *Client) CreateMailPushDevice(md *MailPushDevice) (int64, error) {
	return c.Create(MailPushDeviceModel, md)
}

// UpdateMailPushDevice updates an existing mail_push.device record.
func (c *Client) UpdateMailPushDevice(md *MailPushDevice) error {
	return c.UpdateMailPushDevices([]int64{md.Id.Get()}, md)
}

// UpdateMailPushDevices updates existing mail_push.device records.
// All records (represented by ids) will be updated by md values.
func (c *Client) UpdateMailPushDevices(ids []int64, md *MailPushDevice) error {
	return c.Update(MailPushDeviceModel, ids, md)
}

// DeleteMailPushDevice deletes an existing mail_push.device record.
func (c *Client) DeleteMailPushDevice(id int64) error {
	return c.DeleteMailPushDevices([]int64{id})
}

// DeleteMailPushDevices deletes existing mail_push.device records.
func (c *Client) DeleteMailPushDevices(ids []int64) error {
	return c.Delete(MailPushDeviceModel, ids)
}

// GetMailPushDevice gets mail_push.device existing record.
func (c *Client) GetMailPushDevice(id int64) (*MailPushDevice, error) {
	mds, err := c.GetMailPushDevices([]int64{id})
	if err != nil {
		return nil, err
	}
	if mds != nil && len(*mds) > 0 {
		return &((*mds)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail_push.device not found", id)
}

// GetMailPushDevices gets mail_push.device existing records.
func (c *Client) GetMailPushDevices(ids []int64) (*MailPushDevices, error) {
	mds := &MailPushDevices{}
	if err := c.Read(MailPushDeviceModel, ids, nil, mds); err != nil {
		return nil, err
	}
	return mds, nil
}

// FindMailPushDevice finds mail_push.device record by querying it with criteria.
func (c *Client) FindMailPushDevice(criteria *Criteria) (*MailPushDevice, error) {
	mds := &MailPushDevices{}
	if err := c.SearchRead(MailPushDeviceModel, criteria, NewOptions().Limit(1), mds); err != nil {
		return nil, err
	}
	if mds != nil && len(*mds) > 0 {
		return &((*mds)[0]), nil
	}
	return nil, fmt.Errorf("mail_push.device was not found")
}

// FindMailPushDevices finds mail_push.device records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailPushDevices(criteria *Criteria, options *Options) (*MailPushDevices, error) {
	mds := &MailPushDevices{}
	if err := c.SearchRead(MailPushDeviceModel, criteria, options, mds); err != nil {
		return nil, err
	}
	return mds, nil
}

// FindMailPushDeviceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailPushDeviceIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailPushDeviceModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailPushDeviceId finds record id by querying it with criteria.
func (c *Client) FindMailPushDeviceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailPushDeviceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail_push.device was not found")
}

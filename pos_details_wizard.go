package odoo

import (
	"fmt"
)

// PosDetailsWizard represents pos.details.wizard model.
type PosDetailsWizard struct {
	LastUpdate   *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName  *String   `xmlrpc:"display_name,omptempty"`
	EndDate      *Time     `xmlrpc:"end_date,omptempty"`
	Id           *Int      `xmlrpc:"id,omptempty"`
	PosConfigIds *Relation `xmlrpc:"pos_config_ids,omptempty"`
	StartDate    *Time     `xmlrpc:"start_date,omptempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omptempty"`
}

// PosDetailsWizards represents array of pos.details.wizard model.
type PosDetailsWizards []PosDetailsWizard

// PosDetailsWizardModel is the odoo model name.
const PosDetailsWizardModel = "pos.details.wizard"

// Many2One convert PosDetailsWizard to *Many2One.
func (pdw *PosDetailsWizard) Many2One() *Many2One {
	return NewMany2One(pdw.Id.Get(), "")
}

// CreatePosDetailsWizard creates a new pos.details.wizard model and returns its id.
func (c *Client) CreatePosDetailsWizard(pdw *PosDetailsWizard) (int64, error) {
	return c.Create(PosDetailsWizardModel, pdw)
}

// UpdatePosDetailsWizard updates an existing pos.details.wizard record.
func (c *Client) UpdatePosDetailsWizard(pdw *PosDetailsWizard) error {
	return c.UpdatePosDetailsWizards([]int64{pdw.Id.Get()}, pdw)
}

// UpdatePosDetailsWizards updates existing pos.details.wizard records.
// All records (represented by ids) will be updated by pdw values.
func (c *Client) UpdatePosDetailsWizards(ids []int64, pdw *PosDetailsWizard) error {
	return c.Update(PosDetailsWizardModel, ids, pdw)
}

// DeletePosDetailsWizard deletes an existing pos.details.wizard record.
func (c *Client) DeletePosDetailsWizard(id int64) error {
	return c.DeletePosDetailsWizards([]int64{id})
}

// DeletePosDetailsWizards deletes existing pos.details.wizard records.
func (c *Client) DeletePosDetailsWizards(ids []int64) error {
	return c.Delete(PosDetailsWizardModel, ids)
}

// GetPosDetailsWizard gets pos.details.wizard existing record.
func (c *Client) GetPosDetailsWizard(id int64) (*PosDetailsWizard, error) {
	pdws, err := c.GetPosDetailsWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if pdws != nil && len(*pdws) > 0 {
		return &((*pdws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of pos.details.wizard not found", id)
}

// GetPosDetailsWizards gets pos.details.wizard existing records.
func (c *Client) GetPosDetailsWizards(ids []int64) (*PosDetailsWizards, error) {
	pdws := &PosDetailsWizards{}
	if err := c.Read(PosDetailsWizardModel, ids, nil, pdws); err != nil {
		return nil, err
	}
	return pdws, nil
}

// FindPosDetailsWizard finds pos.details.wizard record by querying it with criteria.
func (c *Client) FindPosDetailsWizard(criteria *Criteria) (*PosDetailsWizard, error) {
	pdws := &PosDetailsWizards{}
	if err := c.SearchRead(PosDetailsWizardModel, criteria, NewOptions().Limit(1), pdws); err != nil {
		return nil, err
	}
	if pdws != nil && len(*pdws) > 0 {
		return &((*pdws)[0]), nil
	}
	return nil, fmt.Errorf("pos.details.wizard was not found")
}

// FindPosDetailsWizards finds pos.details.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosDetailsWizards(criteria *Criteria, options *Options) (*PosDetailsWizards, error) {
	pdws := &PosDetailsWizards{}
	if err := c.SearchRead(PosDetailsWizardModel, criteria, options, pdws); err != nil {
		return nil, err
	}
	return pdws, nil
}

// FindPosDetailsWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosDetailsWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PosDetailsWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPosDetailsWizardId finds record id by querying it with criteria.
func (c *Client) FindPosDetailsWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosDetailsWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("pos.details.wizard was not found")
}

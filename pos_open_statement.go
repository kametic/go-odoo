package odoo

import (
	"fmt"
)

// PosOpenStatement represents pos.open.statement model.
type PosOpenStatement struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// PosOpenStatements represents array of pos.open.statement model.
type PosOpenStatements []PosOpenStatement

// PosOpenStatementModel is the odoo model name.
const PosOpenStatementModel = "pos.open.statement"

// Many2One convert PosOpenStatement to *Many2One.
func (pos *PosOpenStatement) Many2One() *Many2One {
	return NewMany2One(pos.Id.Get(), "")
}

// CreatePosOpenStatement creates a new pos.open.statement model and returns its id.
func (c *Client) CreatePosOpenStatement(pos *PosOpenStatement) (int64, error) {
	return c.Create(PosOpenStatementModel, pos)
}

// UpdatePosOpenStatement updates an existing pos.open.statement record.
func (c *Client) UpdatePosOpenStatement(pos *PosOpenStatement) error {
	return c.UpdatePosOpenStatements([]int64{pos.Id.Get()}, pos)
}

// UpdatePosOpenStatements updates existing pos.open.statement records.
// All records (represented by ids) will be updated by pos values.
func (c *Client) UpdatePosOpenStatements(ids []int64, pos *PosOpenStatement) error {
	return c.Update(PosOpenStatementModel, ids, pos)
}

// DeletePosOpenStatement deletes an existing pos.open.statement record.
func (c *Client) DeletePosOpenStatement(id int64) error {
	return c.DeletePosOpenStatements([]int64{id})
}

// DeletePosOpenStatements deletes existing pos.open.statement records.
func (c *Client) DeletePosOpenStatements(ids []int64) error {
	return c.Delete(PosOpenStatementModel, ids)
}

// GetPosOpenStatement gets pos.open.statement existing record.
func (c *Client) GetPosOpenStatement(id int64) (*PosOpenStatement, error) {
	poss, err := c.GetPosOpenStatements([]int64{id})
	if err != nil {
		return nil, err
	}
	if poss != nil && len(*poss) > 0 {
		return &((*poss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of pos.open.statement not found", id)
}

// GetPosOpenStatements gets pos.open.statement existing records.
func (c *Client) GetPosOpenStatements(ids []int64) (*PosOpenStatements, error) {
	poss := &PosOpenStatements{}
	if err := c.Read(PosOpenStatementModel, ids, nil, poss); err != nil {
		return nil, err
	}
	return poss, nil
}

// FindPosOpenStatement finds pos.open.statement record by querying it with criteria.
func (c *Client) FindPosOpenStatement(criteria *Criteria) (*PosOpenStatement, error) {
	poss := &PosOpenStatements{}
	if err := c.SearchRead(PosOpenStatementModel, criteria, NewOptions().Limit(1), poss); err != nil {
		return nil, err
	}
	if poss != nil && len(*poss) > 0 {
		return &((*poss)[0]), nil
	}
	return nil, fmt.Errorf("pos.open.statement was not found")
}

// FindPosOpenStatements finds pos.open.statement records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosOpenStatements(criteria *Criteria, options *Options) (*PosOpenStatements, error) {
	poss := &PosOpenStatements{}
	if err := c.SearchRead(PosOpenStatementModel, criteria, options, poss); err != nil {
		return nil, err
	}
	return poss, nil
}

// FindPosOpenStatementIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosOpenStatementIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PosOpenStatementModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPosOpenStatementId finds record id by querying it with criteria.
func (c *Client) FindPosOpenStatementId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosOpenStatementModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("pos.open.statement was not found")
}

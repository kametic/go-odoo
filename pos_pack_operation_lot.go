package odoo

import (
	"fmt"
)

// PosPackOperationLot represents pos.pack.operation.lot model.
type PosPackOperationLot struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String   `xmlrpc:"display_name,omptempty"`
	Id             *Int      `xmlrpc:"id,omptempty"`
	LotName        *String   `xmlrpc:"lot_name,omptempty"`
	OrderId        *Many2One `xmlrpc:"order_id,omptempty"`
	PosOrderLineId *Many2One `xmlrpc:"pos_order_line_id,omptempty"`
	ProductId      *Many2One `xmlrpc:"product_id,omptempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omptempty"`
}

// PosPackOperationLots represents array of pos.pack.operation.lot model.
type PosPackOperationLots []PosPackOperationLot

// PosPackOperationLotModel is the odoo model name.
const PosPackOperationLotModel = "pos.pack.operation.lot"

// Many2One convert PosPackOperationLot to *Many2One.
func (ppol *PosPackOperationLot) Many2One() *Many2One {
	return NewMany2One(ppol.Id.Get(), "")
}

// CreatePosPackOperationLot creates a new pos.pack.operation.lot model and returns its id.
func (c *Client) CreatePosPackOperationLot(ppol *PosPackOperationLot) (int64, error) {
	return c.Create(PosPackOperationLotModel, ppol)
}

// UpdatePosPackOperationLot updates an existing pos.pack.operation.lot record.
func (c *Client) UpdatePosPackOperationLot(ppol *PosPackOperationLot) error {
	return c.UpdatePosPackOperationLots([]int64{ppol.Id.Get()}, ppol)
}

// UpdatePosPackOperationLots updates existing pos.pack.operation.lot records.
// All records (represented by ids) will be updated by ppol values.
func (c *Client) UpdatePosPackOperationLots(ids []int64, ppol *PosPackOperationLot) error {
	return c.Update(PosPackOperationLotModel, ids, ppol)
}

// DeletePosPackOperationLot deletes an existing pos.pack.operation.lot record.
func (c *Client) DeletePosPackOperationLot(id int64) error {
	return c.DeletePosPackOperationLots([]int64{id})
}

// DeletePosPackOperationLots deletes existing pos.pack.operation.lot records.
func (c *Client) DeletePosPackOperationLots(ids []int64) error {
	return c.Delete(PosPackOperationLotModel, ids)
}

// GetPosPackOperationLot gets pos.pack.operation.lot existing record.
func (c *Client) GetPosPackOperationLot(id int64) (*PosPackOperationLot, error) {
	ppols, err := c.GetPosPackOperationLots([]int64{id})
	if err != nil {
		return nil, err
	}
	if ppols != nil && len(*ppols) > 0 {
		return &((*ppols)[0]), nil
	}
	return nil, fmt.Errorf("id %v of pos.pack.operation.lot not found", id)
}

// GetPosPackOperationLots gets pos.pack.operation.lot existing records.
func (c *Client) GetPosPackOperationLots(ids []int64) (*PosPackOperationLots, error) {
	ppols := &PosPackOperationLots{}
	if err := c.Read(PosPackOperationLotModel, ids, nil, ppols); err != nil {
		return nil, err
	}
	return ppols, nil
}

// FindPosPackOperationLot finds pos.pack.operation.lot record by querying it with criteria.
func (c *Client) FindPosPackOperationLot(criteria *Criteria) (*PosPackOperationLot, error) {
	ppols := &PosPackOperationLots{}
	if err := c.SearchRead(PosPackOperationLotModel, criteria, NewOptions().Limit(1), ppols); err != nil {
		return nil, err
	}
	if ppols != nil && len(*ppols) > 0 {
		return &((*ppols)[0]), nil
	}
	return nil, fmt.Errorf("pos.pack.operation.lot was not found")
}

// FindPosPackOperationLots finds pos.pack.operation.lot records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPackOperationLots(criteria *Criteria, options *Options) (*PosPackOperationLots, error) {
	ppols := &PosPackOperationLots{}
	if err := c.SearchRead(PosPackOperationLotModel, criteria, options, ppols); err != nil {
		return nil, err
	}
	return ppols, nil
}

// FindPosPackOperationLotIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosPackOperationLotIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PosPackOperationLotModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPosPackOperationLotId finds record id by querying it with criteria.
func (c *Client) FindPosPackOperationLotId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosPackOperationLotModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("pos.pack.operation.lot was not found")
}

package odoo

import (
	"fmt"
)

// PosOrder represents pos.order model.
type PosOrder struct {
	LastUpdate       *Time      `xmlrpc:"__last_update,omptempty"`
	AccountMove      *Many2One  `xmlrpc:"account_move,omptempty"`
	AmountPaid       *Float     `xmlrpc:"amount_paid,omptempty"`
	AmountReturn     *Float     `xmlrpc:"amount_return,omptempty"`
	AmountTax        *Float     `xmlrpc:"amount_tax,omptempty"`
	AmountTotal      *Float     `xmlrpc:"amount_total,omptempty"`
	CompanyId        *Many2One  `xmlrpc:"company_id,omptempty"`
	ConfigId         *Many2One  `xmlrpc:"config_id,omptempty"`
	CreateDate       *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omptempty"`
	DateOrder        *Time      `xmlrpc:"date_order,omptempty"`
	DisplayName      *String    `xmlrpc:"display_name,omptempty"`
	FiscalPositionId *Many2One  `xmlrpc:"fiscal_position_id,omptempty"`
	Id               *Int       `xmlrpc:"id,omptempty"`
	InvoiceId        *Many2One  `xmlrpc:"invoice_id,omptempty"`
	Lines            *Relation  `xmlrpc:"lines,omptempty"`
	LocationId       *Many2One  `xmlrpc:"location_id,omptempty"`
	LoyaltyPoints    *Float     `xmlrpc:"loyalty_points,omptempty"`
	Name             *String    `xmlrpc:"name,omptempty"`
	NbPrint          *Int       `xmlrpc:"nb_print,omptempty"`
	Note             *String    `xmlrpc:"note,omptempty"`
	PartnerId        *Many2One  `xmlrpc:"partner_id,omptempty"`
	PickingId        *Many2One  `xmlrpc:"picking_id,omptempty"`
	PickingTypeId    *Many2One  `xmlrpc:"picking_type_id,omptempty"`
	PosReference     *String    `xmlrpc:"pos_reference,omptempty"`
	PricelistId      *Many2One  `xmlrpc:"pricelist_id,omptempty"`
	SaleJournal      *Many2One  `xmlrpc:"sale_journal,omptempty"`
	SequenceNumber   *Int       `xmlrpc:"sequence_number,omptempty"`
	SessionId        *Many2One  `xmlrpc:"session_id,omptempty"`
	State            *Selection `xmlrpc:"state,omptempty"`
	StatementIds     *Relation  `xmlrpc:"statement_ids,omptempty"`
	UserId           *Many2One  `xmlrpc:"user_id,omptempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// PosOrders represents array of pos.order model.
type PosOrders []PosOrder

// PosOrderModel is the odoo model name.
const PosOrderModel = "pos.order"

// Many2One convert PosOrder to *Many2One.
func (po *PosOrder) Many2One() *Many2One {
	return NewMany2One(po.Id.Get(), "")
}

// CreatePosOrder creates a new pos.order model and returns its id.
func (c *Client) CreatePosOrder(po *PosOrder) (int64, error) {
	return c.Create(PosOrderModel, po)
}

// UpdatePosOrder updates an existing pos.order record.
func (c *Client) UpdatePosOrder(po *PosOrder) error {
	return c.UpdatePosOrders([]int64{po.Id.Get()}, po)
}

// UpdatePosOrders updates existing pos.order records.
// All records (represented by ids) will be updated by po values.
func (c *Client) UpdatePosOrders(ids []int64, po *PosOrder) error {
	return c.Update(PosOrderModel, ids, po)
}

// DeletePosOrder deletes an existing pos.order record.
func (c *Client) DeletePosOrder(id int64) error {
	return c.DeletePosOrders([]int64{id})
}

// DeletePosOrders deletes existing pos.order records.
func (c *Client) DeletePosOrders(ids []int64) error {
	return c.Delete(PosOrderModel, ids)
}

// GetPosOrder gets pos.order existing record.
func (c *Client) GetPosOrder(id int64) (*PosOrder, error) {
	pos, err := c.GetPosOrders([]int64{id})
	if err != nil {
		return nil, err
	}
	if pos != nil && len(*pos) > 0 {
		return &((*pos)[0]), nil
	}
	return nil, fmt.Errorf("id %v of pos.order not found", id)
}

// GetPosOrders gets pos.order existing records.
func (c *Client) GetPosOrders(ids []int64) (*PosOrders, error) {
	pos := &PosOrders{}
	if err := c.Read(PosOrderModel, ids, nil, pos); err != nil {
		return nil, err
	}
	return pos, nil
}

// FindPosOrder finds pos.order record by querying it with criteria.
func (c *Client) FindPosOrder(criteria *Criteria) (*PosOrder, error) {
	pos := &PosOrders{}
	if err := c.SearchRead(PosOrderModel, criteria, NewOptions().Limit(1), pos); err != nil {
		return nil, err
	}
	if pos != nil && len(*pos) > 0 {
		return &((*pos)[0]), nil
	}
	return nil, fmt.Errorf("pos.order was not found")
}

// FindPosOrders finds pos.order records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosOrders(criteria *Criteria, options *Options) (*PosOrders, error) {
	pos := &PosOrders{}
	if err := c.SearchRead(PosOrderModel, criteria, options, pos); err != nil {
		return nil, err
	}
	return pos, nil
}

// FindPosOrderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosOrderIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PosOrderModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPosOrderId finds record id by querying it with criteria.
func (c *Client) FindPosOrderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosOrderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("pos.order was not found")
}

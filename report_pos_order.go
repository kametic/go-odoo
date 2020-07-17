package odoo

import (
	"fmt"
)

// ReportPosOrder represents report.pos.order model.
type ReportPosOrder struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omptempty"`
	AveragePrice    *Float     `xmlrpc:"average_price,omptempty"`
	CompanyId       *Many2One  `xmlrpc:"company_id,omptempty"`
	ConfigId        *Many2One  `xmlrpc:"config_id,omptempty"`
	Date            *Time      `xmlrpc:"date,omptempty"`
	DelayValidation *Int       `xmlrpc:"delay_validation,omptempty"`
	DisplayName     *String    `xmlrpc:"display_name,omptempty"`
	Id              *Int       `xmlrpc:"id,omptempty"`
	Invoiced        *Bool      `xmlrpc:"invoiced,omptempty"`
	JournalId       *Many2One  `xmlrpc:"journal_id,omptempty"`
	LocationId      *Many2One  `xmlrpc:"location_id,omptempty"`
	NbrLines        *Int       `xmlrpc:"nbr_lines,omptempty"`
	OrderId         *Many2One  `xmlrpc:"order_id,omptempty"`
	PartnerId       *Many2One  `xmlrpc:"partner_id,omptempty"`
	PosCategId      *Many2One  `xmlrpc:"pos_categ_id,omptempty"`
	PriceSubTotal   *Float     `xmlrpc:"price_sub_total,omptempty"`
	PriceTotal      *Float     `xmlrpc:"price_total,omptempty"`
	PricelistId     *Many2One  `xmlrpc:"pricelist_id,omptempty"`
	ProductCategId  *Many2One  `xmlrpc:"product_categ_id,omptempty"`
	ProductId       *Many2One  `xmlrpc:"product_id,omptempty"`
	ProductQty      *Int       `xmlrpc:"product_qty,omptempty"`
	ProductTmplId   *Many2One  `xmlrpc:"product_tmpl_id,omptempty"`
	SessionId       *Many2One  `xmlrpc:"session_id,omptempty"`
	State           *Selection `xmlrpc:"state,omptempty"`
	StockLocationId *Many2One  `xmlrpc:"stock_location_id,omptempty"`
	TotalDiscount   *Float     `xmlrpc:"total_discount,omptempty"`
	UserId          *Many2One  `xmlrpc:"user_id,omptempty"`
}

// ReportPosOrders represents array of report.pos.order model.
type ReportPosOrders []ReportPosOrder

// ReportPosOrderModel is the odoo model name.
const ReportPosOrderModel = "report.pos.order"

// Many2One convert ReportPosOrder to *Many2One.
func (rpo *ReportPosOrder) Many2One() *Many2One {
	return NewMany2One(rpo.Id.Get(), "")
}

// CreateReportPosOrder creates a new report.pos.order model and returns its id.
func (c *Client) CreateReportPosOrder(rpo *ReportPosOrder) (int64, error) {
	return c.Create(ReportPosOrderModel, rpo)
}

// UpdateReportPosOrder updates an existing report.pos.order record.
func (c *Client) UpdateReportPosOrder(rpo *ReportPosOrder) error {
	return c.UpdateReportPosOrders([]int64{rpo.Id.Get()}, rpo)
}

// UpdateReportPosOrders updates existing report.pos.order records.
// All records (represented by ids) will be updated by rpo values.
func (c *Client) UpdateReportPosOrders(ids []int64, rpo *ReportPosOrder) error {
	return c.Update(ReportPosOrderModel, ids, rpo)
}

// DeleteReportPosOrder deletes an existing report.pos.order record.
func (c *Client) DeleteReportPosOrder(id int64) error {
	return c.DeleteReportPosOrders([]int64{id})
}

// DeleteReportPosOrders deletes existing report.pos.order records.
func (c *Client) DeleteReportPosOrders(ids []int64) error {
	return c.Delete(ReportPosOrderModel, ids)
}

// GetReportPosOrder gets report.pos.order existing record.
func (c *Client) GetReportPosOrder(id int64) (*ReportPosOrder, error) {
	rpos, err := c.GetReportPosOrders([]int64{id})
	if err != nil {
		return nil, err
	}
	if rpos != nil && len(*rpos) > 0 {
		return &((*rpos)[0]), nil
	}
	return nil, fmt.Errorf("id %v of report.pos.order not found", id)
}

// GetReportPosOrders gets report.pos.order existing records.
func (c *Client) GetReportPosOrders(ids []int64) (*ReportPosOrders, error) {
	rpos := &ReportPosOrders{}
	if err := c.Read(ReportPosOrderModel, ids, nil, rpos); err != nil {
		return nil, err
	}
	return rpos, nil
}

// FindReportPosOrder finds report.pos.order record by querying it with criteria.
func (c *Client) FindReportPosOrder(criteria *Criteria) (*ReportPosOrder, error) {
	rpos := &ReportPosOrders{}
	if err := c.SearchRead(ReportPosOrderModel, criteria, NewOptions().Limit(1), rpos); err != nil {
		return nil, err
	}
	if rpos != nil && len(*rpos) > 0 {
		return &((*rpos)[0]), nil
	}
	return nil, fmt.Errorf("report.pos.order was not found")
}

// FindReportPosOrders finds report.pos.order records by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportPosOrders(criteria *Criteria, options *Options) (*ReportPosOrders, error) {
	rpos := &ReportPosOrders{}
	if err := c.SearchRead(ReportPosOrderModel, criteria, options, rpos); err != nil {
		return nil, err
	}
	return rpos, nil
}

// FindReportPosOrderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindReportPosOrderIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ReportPosOrderModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindReportPosOrderId finds record id by querying it with criteria.
func (c *Client) FindReportPosOrderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ReportPosOrderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("report.pos.order was not found")
}

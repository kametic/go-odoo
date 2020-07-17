package odoo

import (
	"fmt"
)

// PosConfig represents pos.config model.
type PosConfig struct {
	LastUpdate                   *Time      `xmlrpc:"__last_update,omptempty"`
	Active                       *Bool      `xmlrpc:"active,omptempty"`
	AvailablePricelistIds        *Relation  `xmlrpc:"available_pricelist_ids,omptempty"`
	BarcodeNomenclatureId        *Many2One  `xmlrpc:"barcode_nomenclature_id,omptempty"`
	BarcodeScanner               *Bool      `xmlrpc:"barcode_scanner,omptempty"`
	CashControl                  *Bool      `xmlrpc:"cash_control,omptempty"`
	CompanyId                    *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate                   *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                    *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                   *Many2One  `xmlrpc:"currency_id,omptempty"`
	CurrentSessionId             *Many2One  `xmlrpc:"current_session_id,omptempty"`
	CurrentSessionState          *String    `xmlrpc:"current_session_state,omptempty"`
	CustomerFacingDisplayHtml    *String    `xmlrpc:"customer_facing_display_html,omptempty"`
	DefaultCashboxLinesIds       *Relation  `xmlrpc:"default_cashbox_lines_ids,omptempty"`
	DefaultFiscalPositionId      *Many2One  `xmlrpc:"default_fiscal_position_id,omptempty"`
	DisplayName                  *String    `xmlrpc:"display_name,omptempty"`
	FiscalPositionIds            *Relation  `xmlrpc:"fiscal_position_ids,omptempty"`
	GroupBy                      *Bool      `xmlrpc:"group_by,omptempty"`
	GroupPosManagerId            *Many2One  `xmlrpc:"group_pos_manager_id,omptempty"`
	GroupPosUserId               *Many2One  `xmlrpc:"group_pos_user_id,omptempty"`
	GroupPricelistItem           *Bool      `xmlrpc:"group_pricelist_item,omptempty"`
	GroupSalePricelist           *Bool      `xmlrpc:"group_sale_pricelist,omptempty"`
	Id                           *Int       `xmlrpc:"id,omptempty"`
	IfaceBigScrollbars           *Bool      `xmlrpc:"iface_big_scrollbars,omptempty"`
	IfaceCashdrawer              *Bool      `xmlrpc:"iface_cashdrawer,omptempty"`
	IfaceCustomerFacingDisplay   *Bool      `xmlrpc:"iface_customer_facing_display,omptempty"`
	IfaceDisplayCategImages      *Bool      `xmlrpc:"iface_display_categ_images,omptempty"`
	IfaceElectronicScale         *Bool      `xmlrpc:"iface_electronic_scale,omptempty"`
	IfaceInvoicing               *Bool      `xmlrpc:"iface_invoicing,omptempty"`
	IfacePaymentTerminal         *Bool      `xmlrpc:"iface_payment_terminal,omptempty"`
	IfacePrecomputeCash          *Bool      `xmlrpc:"iface_precompute_cash,omptempty"`
	IfacePrintAuto               *Bool      `xmlrpc:"iface_print_auto,omptempty"`
	IfacePrintSkipScreen         *Bool      `xmlrpc:"iface_print_skip_screen,omptempty"`
	IfacePrintViaProxy           *Bool      `xmlrpc:"iface_print_via_proxy,omptempty"`
	IfaceScanViaProxy            *Bool      `xmlrpc:"iface_scan_via_proxy,omptempty"`
	IfaceStartCategId            *Many2One  `xmlrpc:"iface_start_categ_id,omptempty"`
	IfaceTaxIncluded             *Selection `xmlrpc:"iface_tax_included,omptempty"`
	IfaceTipproduct              *Bool      `xmlrpc:"iface_tipproduct,omptempty"`
	IfaceVkeyboard               *Bool      `xmlrpc:"iface_vkeyboard,omptempty"`
	InvoiceJournalId             *Many2One  `xmlrpc:"invoice_journal_id,omptempty"`
	IsHeaderOrFooter             *Bool      `xmlrpc:"is_header_or_footer,omptempty"`
	IsInstalledAccountAccountant *Bool      `xmlrpc:"is_installed_account_accountant,omptempty"`
	IsPosbox                     *Bool      `xmlrpc:"is_posbox,omptempty"`
	JournalId                    *Many2One  `xmlrpc:"journal_id,omptempty"`
	JournalIds                   *Relation  `xmlrpc:"journal_ids,omptempty"`
	LastSessionClosingCash       *Float     `xmlrpc:"last_session_closing_cash,omptempty"`
	LastSessionClosingDate       *Time      `xmlrpc:"last_session_closing_date,omptempty"`
	LoyaltyId                    *Many2One  `xmlrpc:"loyalty_id,omptempty"`
	ModulePosDiscount            *Bool      `xmlrpc:"module_pos_discount,omptempty"`
	ModulePosLoyalty             *Bool      `xmlrpc:"module_pos_loyalty,omptempty"`
	ModulePosMercury             *Bool      `xmlrpc:"module_pos_mercury,omptempty"`
	ModulePosReprint             *Bool      `xmlrpc:"module_pos_reprint,omptempty"`
	ModulePosRestaurant          *Bool      `xmlrpc:"module_pos_restaurant,omptempty"`
	Name                         *String    `xmlrpc:"name,omptempty"`
	PickingTypeId                *Many2One  `xmlrpc:"picking_type_id,omptempty"`
	PosSessionState              *String    `xmlrpc:"pos_session_state,omptempty"`
	PosSessionUsername           *String    `xmlrpc:"pos_session_username,omptempty"`
	PricelistId                  *Many2One  `xmlrpc:"pricelist_id,omptempty"`
	ProxyIp                      *String    `xmlrpc:"proxy_ip,omptempty"`
	ReceiptFooter                *String    `xmlrpc:"receipt_footer,omptempty"`
	ReceiptHeader                *String    `xmlrpc:"receipt_header,omptempty"`
	RestrictPriceControl         *Bool      `xmlrpc:"restrict_price_control,omptempty"`
	SequenceId                   *Many2One  `xmlrpc:"sequence_id,omptempty"`
	SequenceLineId               *Many2One  `xmlrpc:"sequence_line_id,omptempty"`
	SessionIds                   *Relation  `xmlrpc:"session_ids,omptempty"`
	StartCategory                *Bool      `xmlrpc:"start_category,omptempty"`
	StockLocationId              *Many2One  `xmlrpc:"stock_location_id,omptempty"`
	TaxRegime                    *Bool      `xmlrpc:"tax_regime,omptempty"`
	TaxRegimeSelection           *Bool      `xmlrpc:"tax_regime_selection,omptempty"`
	TipProductId                 *Many2One  `xmlrpc:"tip_product_id,omptempty"`
	UseExistingLots              *Bool      `xmlrpc:"use_existing_lots,omptempty"`
	UsePricelist                 *Bool      `xmlrpc:"use_pricelist,omptempty"`
	Uuid                         *String    `xmlrpc:"uuid,omptempty"`
	WriteDate                    *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                     *Many2One  `xmlrpc:"write_uid,omptempty"`
	XStageId                     *Many2One  `xmlrpc:"x_stage_id,omptempty"`
}

// PosConfigs represents array of pos.config model.
type PosConfigs []PosConfig

// PosConfigModel is the odoo model name.
const PosConfigModel = "pos.config"

// Many2One convert PosConfig to *Many2One.
func (pc *PosConfig) Many2One() *Many2One {
	return NewMany2One(pc.Id.Get(), "")
}

// CreatePosConfig creates a new pos.config model and returns its id.
func (c *Client) CreatePosConfig(pc *PosConfig) (int64, error) {
	return c.Create(PosConfigModel, pc)
}

// UpdatePosConfig updates an existing pos.config record.
func (c *Client) UpdatePosConfig(pc *PosConfig) error {
	return c.UpdatePosConfigs([]int64{pc.Id.Get()}, pc)
}

// UpdatePosConfigs updates existing pos.config records.
// All records (represented by ids) will be updated by pc values.
func (c *Client) UpdatePosConfigs(ids []int64, pc *PosConfig) error {
	return c.Update(PosConfigModel, ids, pc)
}

// DeletePosConfig deletes an existing pos.config record.
func (c *Client) DeletePosConfig(id int64) error {
	return c.DeletePosConfigs([]int64{id})
}

// DeletePosConfigs deletes existing pos.config records.
func (c *Client) DeletePosConfigs(ids []int64) error {
	return c.Delete(PosConfigModel, ids)
}

// GetPosConfig gets pos.config existing record.
func (c *Client) GetPosConfig(id int64) (*PosConfig, error) {
	pcs, err := c.GetPosConfigs([]int64{id})
	if err != nil {
		return nil, err
	}
	if pcs != nil && len(*pcs) > 0 {
		return &((*pcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of pos.config not found", id)
}

// GetPosConfigs gets pos.config existing records.
func (c *Client) GetPosConfigs(ids []int64) (*PosConfigs, error) {
	pcs := &PosConfigs{}
	if err := c.Read(PosConfigModel, ids, nil, pcs); err != nil {
		return nil, err
	}
	return pcs, nil
}

// FindPosConfig finds pos.config record by querying it with criteria.
func (c *Client) FindPosConfig(criteria *Criteria) (*PosConfig, error) {
	pcs := &PosConfigs{}
	if err := c.SearchRead(PosConfigModel, criteria, NewOptions().Limit(1), pcs); err != nil {
		return nil, err
	}
	if pcs != nil && len(*pcs) > 0 {
		return &((*pcs)[0]), nil
	}
	return nil, fmt.Errorf("pos.config was not found")
}

// FindPosConfigs finds pos.config records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosConfigs(criteria *Criteria, options *Options) (*PosConfigs, error) {
	pcs := &PosConfigs{}
	if err := c.SearchRead(PosConfigModel, criteria, options, pcs); err != nil {
		return nil, err
	}
	return pcs, nil
}

// FindPosConfigIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosConfigIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PosConfigModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPosConfigId finds record id by querying it with criteria.
func (c *Client) FindPosConfigId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosConfigModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("pos.config was not found")
}

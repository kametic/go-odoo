package odoo

import (
	"fmt"
)

// PosSession represents pos.session model.
type PosSession struct {
	LastUpdate                     *Time      `xmlrpc:"__last_update,omptempty"`
	CashControl                    *Bool      `xmlrpc:"cash_control,omptempty"`
	CashJournalId                  *Many2One  `xmlrpc:"cash_journal_id,omptempty"`
	CashRegisterBalanceEnd         *Float     `xmlrpc:"cash_register_balance_end,omptempty"`
	CashRegisterBalanceEndReal     *Float     `xmlrpc:"cash_register_balance_end_real,omptempty"`
	CashRegisterBalanceStart       *Float     `xmlrpc:"cash_register_balance_start,omptempty"`
	CashRegisterDifference         *Float     `xmlrpc:"cash_register_difference,omptempty"`
	CashRegisterId                 *Many2One  `xmlrpc:"cash_register_id,omptempty"`
	CashRegisterTotalEntryEncoding *Float     `xmlrpc:"cash_register_total_entry_encoding,omptempty"`
	ConfigId                       *Many2One  `xmlrpc:"config_id,omptempty"`
	CreateDate                     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                      *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId                     *Many2One  `xmlrpc:"currency_id,omptempty"`
	DisplayName                    *String    `xmlrpc:"display_name,omptempty"`
	Id                             *Int       `xmlrpc:"id,omptempty"`
	JournalIds                     *Relation  `xmlrpc:"journal_ids,omptempty"`
	LoginNumber                    *Int       `xmlrpc:"login_number,omptempty"`
	Name                           *String    `xmlrpc:"name,omptempty"`
	OrderIds                       *Relation  `xmlrpc:"order_ids,omptempty"`
	PickingCount                   *Int       `xmlrpc:"picking_count,omptempty"`
	Rescue                         *Bool      `xmlrpc:"rescue,omptempty"`
	SequenceNumber                 *Int       `xmlrpc:"sequence_number,omptempty"`
	StartAt                        *Time      `xmlrpc:"start_at,omptempty"`
	State                          *Selection `xmlrpc:"state,omptempty"`
	StatementIds                   *Relation  `xmlrpc:"statement_ids,omptempty"`
	StopAt                         *Time      `xmlrpc:"stop_at,omptempty"`
	UserId                         *Many2One  `xmlrpc:"user_id,omptempty"`
	WriteDate                      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                       *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// PosSessions represents array of pos.session model.
type PosSessions []PosSession

// PosSessionModel is the odoo model name.
const PosSessionModel = "pos.session"

// Many2One convert PosSession to *Many2One.
func (ps *PosSession) Many2One() *Many2One {
	return NewMany2One(ps.Id.Get(), "")
}

// CreatePosSession creates a new pos.session model and returns its id.
func (c *Client) CreatePosSession(ps *PosSession) (int64, error) {
	return c.Create(PosSessionModel, ps)
}

// UpdatePosSession updates an existing pos.session record.
func (c *Client) UpdatePosSession(ps *PosSession) error {
	return c.UpdatePosSessions([]int64{ps.Id.Get()}, ps)
}

// UpdatePosSessions updates existing pos.session records.
// All records (represented by ids) will be updated by ps values.
func (c *Client) UpdatePosSessions(ids []int64, ps *PosSession) error {
	return c.Update(PosSessionModel, ids, ps)
}

// DeletePosSession deletes an existing pos.session record.
func (c *Client) DeletePosSession(id int64) error {
	return c.DeletePosSessions([]int64{id})
}

// DeletePosSessions deletes existing pos.session records.
func (c *Client) DeletePosSessions(ids []int64) error {
	return c.Delete(PosSessionModel, ids)
}

// GetPosSession gets pos.session existing record.
func (c *Client) GetPosSession(id int64) (*PosSession, error) {
	pss, err := c.GetPosSessions([]int64{id})
	if err != nil {
		return nil, err
	}
	if pss != nil && len(*pss) > 0 {
		return &((*pss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of pos.session not found", id)
}

// GetPosSessions gets pos.session existing records.
func (c *Client) GetPosSessions(ids []int64) (*PosSessions, error) {
	pss := &PosSessions{}
	if err := c.Read(PosSessionModel, ids, nil, pss); err != nil {
		return nil, err
	}
	return pss, nil
}

// FindPosSession finds pos.session record by querying it with criteria.
func (c *Client) FindPosSession(criteria *Criteria) (*PosSession, error) {
	pss := &PosSessions{}
	if err := c.SearchRead(PosSessionModel, criteria, NewOptions().Limit(1), pss); err != nil {
		return nil, err
	}
	if pss != nil && len(*pss) > 0 {
		return &((*pss)[0]), nil
	}
	return nil, fmt.Errorf("pos.session was not found")
}

// FindPosSessions finds pos.session records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosSessions(criteria *Criteria, options *Options) (*PosSessions, error) {
	pss := &PosSessions{}
	if err := c.SearchRead(PosSessionModel, criteria, options, pss); err != nil {
		return nil, err
	}
	return pss, nil
}

// FindPosSessionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPosSessionIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(PosSessionModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindPosSessionId finds record id by querying it with criteria.
func (c *Client) FindPosSessionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PosSessionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("pos.session was not found")
}

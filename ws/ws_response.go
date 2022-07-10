package ws

import (
	sjson "encoding/json"
	"strconv"
	"time"
)

type OrderBookL2 struct {
	ID     int64   `json:"id"`
	Price  float64 `json:"price,string"`
	Side   string  `json:"side"`
	Size   int64   `json:"size"`
	Symbol string  `json:"symbol"`
}

type OrderBookL2Delta struct {
	Delete []*OrderBookL2 `json:"delete"`
	Update []*OrderBookL2 `json:"update"`
	Insert []*OrderBookL2 `json:"insert"`
}

func (o *OrderBookL2) Key() string {
	return strconv.FormatInt(o.ID, 10)
}

type Trade struct {
	Timestamp     time.Time `json:"timestamp"`
	Symbol        string    `json:"symbol"`
	Side          string    `json:"side"`
	Size          int       `json:"size"`
	Price         float64   `json:"price"`
	TickDirection string    `json:"tick_direction"`
	TradeID       string    `json:"trade_id"`
	CrossSeq      int       `json:"cross_seq"`
}

type KLine struct {
	Start     int64   `json:"start"`
	End       int64   `json:"end"`
	Period    string  `json:"period"` // 1
	Open      float64 `json:"open"`
	Close     float64 `json:"close"`
	High      float64 `json:"high"`
	Low       float64 `json:"low"`
	Volume    string  `json:"volume"`
	Turnover  string  `json:"turnover"` // 0.0013844
	Confirm   bool    `json:"confirm"`
	CrossSeq  int64   `json:"cross_seq"`
	Timestamp int64   `json:"timestamp"` // 1657326780540640
}

type Insurance struct {
	Currency      string    `json:"currency"`
	Timestamp     time.Time `json:"timestamp"`
	WalletBalance int64     `json:"wallet_balance"`
}

type Instrument struct {
	Symbol     string  `json:"symbol"`
	MarkPrice  float64 `json:"mark_price"`
	IndexPrice float64 `json:"index_price"`
}

type Order struct {
	OrderID       string       `json:"order_id"`
	OrderLinkID   string       `json:"order_link_id"`
	Symbol        string       `json:"symbol"`
	Side          string       `json:"side"`
	OrderType     string       `json:"order_type"`
	Price         sjson.Number `json:"price"`
	Qty           float64      `json:"qty"`
	TimeInForce   string       `json:"time_in_force"` // GoodTillCancel/ImmediateOrCancel/FillOrKill/PostOnly
	CreateType    string       `json:"create_type"`
	CancelType    string       `json:"cancel_type"`
	OrderStatus   string       `json:"order_status"`
	LeavesQty     float64      `json:"leaves_qty"`
	CumExecQty    float64      `json:"cum_exec_qty"`
	CumExecValue  sjson.Number `json:"cum_exec_value"`
	CumExecFee    sjson.Number `json:"cum_exec_fee"`
	Timestamp     time.Time    `json:"timestamp"`
	TakeProfit    sjson.Number `json:"take_profit"`
	StopLoss      sjson.Number `json:"stop_loss"`
	TrailingStop  sjson.Number `json:"trailing_stop"`
	LastExecPrice sjson.Number `json:"last_exec_price"`
}

type Execution struct {
	Symbol      string    `json:"symbol"`
	Side        string    `json:"side"`
	OrderID     string    `json:"order_id"`
	ExecID      string    `json:"exec_id"`
	OrderLinkID string    `json:"order_link_id"`
	Price       float64   `json:"price,string"`
	OrderQty    float64   `json:"order_qty"`
	ExecType    string    `json:"exec_type"`
	ExecQty     float64   `json:"exec_qty"`
	ExecFee     float64   `json:"exec_fee,string"`
	LeavesQty   float64   `json:"leaves_qty"`
	IsMaker     bool      `json:"is_maker"`
	TradeTime   time.Time `json:"trade_time"`
}

type Position struct {
	UserID         string  `json:"user_id"`
	Symbol         string  `json:"symbol"`
	Size           float64 `json:"size"`
	Side           string  `json:"side"`
	PositionValue  float64 `json:"position_value,string"`
	EntryPrice     float64 `json:"entry_price,string"`
	LiqPrice       float64 `json:"liq_price,string"`
	BustPrice      float64 `json:"bust_price,string"`
	Leverage       float64 `json:"leverage,string"`
	OrderMargin    float64 `json:"order_margin,string"`
	PositionMargin float64 `json:"position_margin,string"`
	OccClosingFee  float64 `json:"occ_closing_fee,string"`
	TakeProfit     float64 `json:"take_profit,string"`
	TPTriggerBy    string  `json:"tp_trigger_by,string"`
	StopLoss       float64 `json:"stop_loss,string"`
	SLTriggerBy    string  `json:"sl_trigger_by,string"`
	TrailingStop   float64 `json:"trailing_stop,string"`
	RealisedPnl    float64 `json:"realised_pnl,string"`
	AutoAddMargin  string  `json:"auto_add_margin"`
	CumRealisedPnl float64 `json:"cum_realised_pnl,string"`
	PositionStatus string  `json:"position_status"`
	PositionId     string  `json:"position_id"`
	PositionSeq    string  `json:"position_seq"`
	FreeQty        float64 `json:"free_qty"`
	TPSLMode       string  `json:"tp_sl_mode"`
	RiskID         string  `json:"risk_id"`
	Isolated       bool    `json:"isolated"`
	Mode           string  `json:"mode"`
	PositionIdx    string  `json:"position_idx"`
}

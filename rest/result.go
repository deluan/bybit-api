package rest

import (
	sjson "encoding/json"
	"time"
)

type BaseResult struct {
	RetCode          int         `json:"ret_code"`
	RetMsg           string      `json:"ret_msg"`
	ExtCode          string      `json:"ext_code"`
	ExtInfo          string      `json:"ext_info"`
	Result           interface{} `json:"result"`
	TimeNow          string      `json:"time_now"`
	RateLimitStatus  int         `json:"rate_limit_status"`
	RateLimitResetMs int64       `json:"rate_limit_reset_ms"`
	RateLimit        int         `json:"rate_limit"`
}

type ResultStringArrayResponse struct {
	BaseResult
	Result []string `json:"result"`
}

type Item struct {
	Price float64 `json:"price,string"`
	Size  float64 `json:"size"`
}

type OrderBook struct {
	Asks []Item    `json:"asks"`
	Bids []Item    `json:"bids"`
	Time time.Time `json:"time"`
}

type RawItem struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price,string"`
	Size   float64 `json:"size"`
	Side   string  `json:"side"` // Buy/Sell
}

type GetOrderBookResult struct {
	BaseResult
	Result []RawItem `json:"result"`
}

type OHLC struct {
	Symbol   string  `json:"symbol"`
	Interval string  `json:"interval"`
	OpenTime int64   `json:"open_time"`
	Open     float64 `json:"open,string"`
	High     float64 `json:"high,string"`
	Low      float64 `json:"low,string"`
	Close    float64 `json:"close,string"`
	Volume   float64 `json:"volume,string"`
	Turnover float64 `json:"turnover,string"`
}

type GetKlineResult struct {
	BaseResult
	Result []OHLC `json:"result"`
}

type OHLCLinear struct {
	Symbol   string  `json:"symbol"`
	Period   string  `json:"period"`
	OpenTime int64   `json:"open_time"`
	Open     float64 `json:"open"`
	High     float64 `json:"high"`
	Low      float64 `json:"low"`
	Close    float64 `json:"close"`
	Volume   float64 `json:"volume"`
	Turnover float64 `json:"turnover"`
}

type GetLinearKlineResult struct {
	BaseResult
	Result []OHLCLinear `json:"result"`
}

type OpenInterest struct {
	Symbol       string       `json:"symbol"`
	OpenInterest sjson.Number `json:"open_interest"`
	Timestamp    sjson.Number `json:"timestamp"`
}

type GetOpenInterestResult struct {
	BaseResult
	Result []OpenInterest `json:"result"`
}

type AccountRatio struct {
	Symbol    string       `json:"symbol"`
	BuyRatio  sjson.Number `json:"buy_ratio"`
	SellRatio sjson.Number `json:"sell_ratio"`
	Timestamp sjson.Number `json:"timestamp"`
}

type GetAccountRatioResult struct {
	BaseResult
	Result []AccountRatio `json:"result"`
}

type Ticker struct {
	Symbol               string       `json:"symbol"`
	BidPrice             float64      `json:"bid_price,string"`
	AskPrice             float64      `json:"ask_price,string"`
	LastPrice            float64      `json:"last_price,string"`
	LastTickDirection    string       `json:"last_tick_direction"`
	PrevPrice24H         sjson.Number `json:"prev_price_24h"`
	Price24HPcnt         sjson.Number `json:"price_24h_pcnt"`
	HighPrice24H         sjson.Number `json:"high_price_24h"`
	LowPrice24H          sjson.Number `json:"low_price_24h"`
	PrevPrice1H          sjson.Number `json:"prev_price_1h"`
	Price1HPcnt          sjson.Number `json:"price_1h_pcnt"`
	MarkPrice            sjson.Number `json:"mark_price"`
	IndexPrice           sjson.Number `json:"index_price"`
	OpenInterest         float64      `json:"open_interest"`
	OpenValue            sjson.Number `json:"open_value"`
	TotalTurnover        sjson.Number `json:"total_turnover"`
	Turnover24H          sjson.Number `json:"turnover_24h"`
	TotalVolume          float64      `json:"total_volume"`
	Volume24H            float64      `json:"volume_24h"`
	FundingRate          sjson.Number `json:"funding_rate"`
	PredictedFundingRate sjson.Number `json:"predicted_funding_rate"`
	NextFundingTime      string       `json:"next_funding_time"` // string because can be empty, parse it with "2006-01-02T15:04:05Z07:00"
	CountdownHour        int          `json:"countdown_hour"`
}

type GetTickersResult struct {
	RetCode int      `json:"ret_code"`
	RetMsg  string   `json:"ret_msg"`
	ExtCode string   `json:"ext_code"`
	ExtInfo string   `json:"ext_info"`
	Result  []Ticker `json:"result"`
	TimeNow string   `json:"time_now"`
}

type TradingRecord struct {
	ID     int       `json:"id"`
	Symbol string    `json:"symbol"`
	Price  float64   `json:"price"`
	Qty    int       `json:"qty"`
	Side   string    `json:"side"`
	Time   time.Time `json:"time"`
}

type GetTradingRecordsResult struct {
	RetCode int             `json:"ret_code"`
	RetMsg  string          `json:"ret_msg"`
	ExtCode string          `json:"ext_code"`
	ExtInfo string          `json:"ext_info"`
	Result  []TradingRecord `json:"result"`
	TimeNow string          `json:"time_now"`
}

type LeverageFilter struct {
	MinLeverage  int     `json:"min_leverage"`
	MaxLeverage  int     `json:"max_leverage"`
	LeverageStep float64 `json:"leverage_step,string"`
}

type PriceFilter struct {
	MinPrice float64 `json:"min_price,string"`
	MaxPrice float64 `json:"max_price,string"`
	TickSize float64 `json:"tick_size,string"`
}

type LotSizeFilter struct {
	MaxTradingQty float64 `json:"max_trading_qty"`
	MinTradingQty float64 `json:"min_trading_qty"`
	QtyStep       float64 `json:"qty_step"`
}

type SymbolInfo struct {
	Name           string         `json:"name"`
	BaseCurrency   string         `json:"base_currency"`
	QuoteCurrency  string         `json:"quote_currency"`
	PriceScale     int            `json:"price_scale"`
	TakerFee       float64        `json:"taker_fee,string"`
	MakerFee       float64        `json:"maker_fee,string"`
	LeverageFilter LeverageFilter `json:"leverage_filter"`
	PriceFilter    PriceFilter    `json:"price_filter"`
	LotSizeFilter  LotSizeFilter  `json:"lot_size_filter"`
}

type GetSymbolsResult struct {
	RetCode int          `json:"ret_code"`
	RetMsg  string       `json:"ret_msg"`
	ExtCode string       `json:"ext_code"`
	ExtInfo string       `json:"ext_info"`
	Result  []SymbolInfo `json:"result"`
	TimeNow string       `json:"time_now"`
}

type Balance struct {
	Equity           float64 `json:"equity"`
	AvailableBalance float64 `json:"available_balance"`
	UsedMargin       float64 `json:"used_margin"`
	OrderMargin      float64 `json:"order_margin"`
	PositionMargin   float64 `json:"position_margin"`
	OccClosingFee    float64 `json:"occ_closing_fee"`
	OccFundingFee    float64 `json:"occ_funding_fee"`
	WalletBalance    float64 `json:"wallet_balance"`
	RealisedPnl      float64 `json:"realised_pnl"`
	UnrealisedPnl    float64 `json:"unrealised_pnl"`
	CumRealisedPnl   float64 `json:"cum_realised_pnl"`
	GivenCash        float64 `json:"given_cash"`
	ServiceCash      float64 `json:"service_cash"`
}

type GetBalanceResult struct {
	RetCode          int                  `json:"ret_code"`
	RetMsg           string               `json:"ret_msg"`
	ExtCode          string               `json:"ext_code"`
	ExtInfo          string               `json:"ext_info"`
	Result           GetBalanceResultData `json:"result"`
	TimeNow          string               `json:"time_now"`
	RateLimitStatus  int                  `json:"rate_limit_status"`
	RateLimitResetMs int64                `json:"rate_limit_reset_ms"`
	RateLimit        int                  `json:"rate_limit"`
}

type GetBalanceResultData struct {
	BTC  Balance `json:"BTC"`
	ETH  Balance `json:"ETH"`
	EOS  Balance `json:"EOS"`
	XRP  Balance `json:"XRP"`
	USDT Balance `json:"USDT"`
}

type Position struct {
	Id                  int       `json:"id"`
	UserId              int       `json:"user_id"`
	RiskId              int       `json:"risk_id"`
	Symbol              string    `json:"symbol"`
	Size                float64   `json:"size"`
	Side                string    `json:"side"`
	EntryPrice          float64   `json:"entry_price,string"`
	LiqPrice            float64   `json:"liq_price,string"`
	BustPrice           float64   `json:"bust_price,string"`
	TakeProfit          float64   `json:"take_profit,string"`
	StopLoss            float64   `json:"stop_loss,string"`
	TrailingStop        float64   `json:"trailing_stop,string"`
	PositionValue       float64   `json:"position_value,string"`
	Leverage            float64   `json:"leverage,string"`
	PositionStatus      string    `json:"position_status"`
	AutoAddMargin       float64   `json:"auto_add_margin"`
	OrderMargin         float64   `json:"order_margin,string"`
	PositionMargin      float64   `json:"position_margin,string"`
	OccClosingFee       float64   `json:"occ_closing_fee,string"`
	OccFundingFee       float64   `json:"occ_funding_fee,string"`
	WalletBalance       float64   `json:"wallet_balance,string"`
	CumRealisedPnl      float64   `json:"cum_realised_pnl,string"`
	CumCommission       float64   `json:"cum_commission,string"`
	RealisedPnl         float64   `json:"realised_pnl,string"`
	DeleverageIndicator float64   `json:"deleverage_indicator"`
	OcCalcData          string    `json:"oc_calc_data"`
	CrossSeq            float64   `json:"cross_seq"`
	PositionSeq         float64   `json:"position_seq"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	UnrealisedPnl       float64   `json:"unrealised_pnl"`
}

type LinearPosition struct {
	UserID              int     `json:"user_id"`
	Symbol              string  `json:"symbol"`
	Side                string  `json:"side"`
	Size                float64 `json:"size"`
	PositionValue       float64 `json:"position_value"`
	EntryPrice          float64 `json:"entry_price"`
	LiqPrice            float64 `json:"liq_price"`
	BustPrice           float64 `json:"bust_price"`
	Leverage            float64 `json:"leverage"`
	AutoAddMargin       float64 `json:"auto_add_margin"`
	IsIsolated          bool    `json:"is_isolated"`
	PositionMargin      float64 `json:"position_margin"`
	OccClosingFee       float64 `json:"occ_closing_fee"`
	RealisedPnl         float64 `json:"realised_pnl"`
	CumRealisedPnl      float64 `json:"cum_realised_pnl"`
	FreeQty             float64 `json:"free_qty"`
	TpSlMode            string  `json:"tp_sl_mode"`
	UnrealisedPnl       float64 `json:"unrealised_pnl"`
	DeleverageIndicator float64 `json:"deleverage_indicator"`
	RiskID              int     `json:"risk_id"`
	StopLoss            float64 `json:"stop_loss"`
	TakeProfit          float64 `json:"take_profit"`
	TrailingStop        float64 `json:"trailing_stop"`
	PositionIdx         int     `json:"position_idx"`
	Mode                string  `json:"mode"`
}

type PositionResponse struct {
	BaseResult
	Result Position `json:"result"`
}

type PositionArrayResponse struct {
	BaseResult
	Result []PositionData `json:"result"`
}

type PositionData struct {
	IsValid bool     `json:"is_valid"`
	Data    Position `json:"data"`
}

type LinearClosedPnL struct {
	ID            int64   `json:"id"`
	UserID        int64   `json:"user_id"`
	Symbol        string  `json:"symbol"`
	OrderID       string  `json:"order_id"`
	Side          string  `json:"side"`
	Qty           float64 `json:"qty"`
	OrderPrice    float64 `json:"order_price"`
	OrderType     string  `json:"order_type"`
	ExecType      string  `json:"exec_type"`
	ClosedSize    float64 `json:"closed_size"`
	CumEntryValue float64 `json:"cum_entry_value"`
	AvgEntryPrice float64 `json:"avg_entry_price"`
	CumExitValue  float64 `json:"cum_exit_value"`
	AvgExitPrice  float64 `json:"avg_exit_price"`
	ClosedPnL     float64 `json:"closed_pnl"`
	FillCount     int     `json:"fill_count"`
	Leverage      float64 `json:"leverage"`
	CreatedAt     int64   `json:"created_at"`
}

type LinearClosedPnLData struct {
	CurrentPage sjson.Number      `json:"current_page"`
	Data        []LinearClosedPnL `json:"data"`
}

type LinearClosedPnLArrayResponse struct {
	BaseResult
	Result LinearClosedPnLData `json:"result"`
}

type LinearPositionArrayResponse struct {
	BaseResult
	Result []LinearPosition `json:"result"`
}

type LinearPositionDataArrayResponse struct {
	BaseResult
	Result []LinearPositionData `json:"result"`
}

type LinearPositionData struct {
	IsValid bool           `json:"is_valid"`
	Data    LinearPosition `json:"data"`
}

type Order struct {
	UserId         int          `json:"user_id"`
	OrderId        string       `json:"order_id"`
	Symbol         string       `json:"symbol"`
	Side           string       `json:"side"`
	OrderType      string       `json:"order_type"`
	Price          sjson.Number `json:"price"`
	Qty            sjson.Number `json:"qty"`
	TimeInForce    string       `json:"time_in_force"`
	OrderStatus    string       `json:"order_status"`
	LastExecTime   sjson.Number `json:"last_exec_time"`
	LastExecPrice  sjson.Number `json:"last_exec_price"`
	LeavesQty      sjson.Number `json:"leaves_qty"`
	CumExecQty     sjson.Number `json:"cum_exec_qty"`
	CumExecValue   sjson.Number `json:"cum_exec_value"`
	CumExecFee     sjson.Number `json:"cum_exec_fee"`
	RejectReason   string       `json:"reject_reason"`
	OrderLinkID    string       `json:"order_link_id"`
	TakeProfit     float64      `json:"take_profit"`
	StopLoss       float64      `json:"stop_loss"`
	ReduceOnly     bool         `json:"reduce_only"`
	CloseOnTrigger bool         `json:"close_on_trigger"`
	CreatedAt      time.Time    `json:"created_time"`
	UpdatedAt      time.Time    `json:"updated_time"`
}

type OrderListResponse struct {
	BaseResult
	Result OrderListResponseResult `json:"result"`
}

type OrderListResponseResult struct {
	Data   []Order `json:"data"`
	Cursor string  `json:"cursor"`
}

type OrderListResponsePaginated struct {
	BaseResult
	Result OrderListResponseResultPaginated `json:"result"`
}

type OrderListResponseResultPaginated struct {
	CurrentPage sjson.Number `json:"current_page"`
	LastPage    sjson.Number `json:"last_page"`
	Data        []Order      `json:"data"`
}

type OrderResponse struct {
	BaseResult
	Result Order `json:"result"`
}

type OrderArrayResponse struct {
	BaseResult
	Result []Order `json:"result"`
}

type StopOrder struct {
	StopOrderId    string       `json:"stop_order_id"`
	UserId         int64        `json:"user_id"`
	Symbol         string       `json:"symbol"`
	Side           string       `json:"side"`
	OrderType      string       `json:"order_type"`
	Price          sjson.Number `json:"price"`
	Qty            sjson.Number `json:"qty"`
	TimeInForce    string       `json:"time_in_force"`
	OrderStatus    string       `json:"order_status"`
	TriggerPrice   sjson.Number `json:"trigger_price"`
	OrderLinkId    string       `json:"order_link_id"`
	CreatedTime    time.Time    `json:"created_time"`
	UpdatedTime    time.Time    `json:"updated_time"`
	TakeProfit     float64      `json:"take_profit"`
	StopLoss       float64      `json:"stop_loss"`
	TriggerBy      string       `json:"trigger_by"`
	BasePrice      sjson.Number `json:"base_price"`
	TPTriggerBy    string       `json:"tp_trigger_by"`
	SLTriggerBy    string       `json:"sl_trigger_by"`
	ReduceOnly     bool         `json:"reduce_only"`
	CloseOnTrigger bool         `json:"close_on_trigger"`
}

type StopOrderListResponse struct {
	BaseResult
	Result StopOrderListResponseResult `json:"result"`
}

type StopOrderListResponseResult struct {
	Data   []StopOrder `json:"data"`
	Cursor string      `json:"cursor"`
}

type StopOrderResponse struct {
	BaseResult
	Result StopOrder `json:"result"`
}

type StopOrderArrayResponse struct {
	BaseResult
	Result []StopOrder `json:"result"`
}

type StopOrderListResponsePaginated struct {
	BaseResult
	Result StopOrderListResponseResultPaginated `json:"result"`
}

type StopOrderListResponseResultPaginated struct {
	CurrentPage string      `json:"current_page"`
	LastPage    string      `json:"last_page"`
	Data        []StopOrder `json:"data"`
}

type WalletFundRecord struct {
	Id            int          `json:"id"`
	UserId        int          `json:"user_id"`
	Coin          string       `json:"coin"`
	WalletId      int          `json:"wallet_id"`
	Type          string       `json:"type"`
	Amount        sjson.Number `json:"amount"`
	TxId          string       `json:"tx_id"`
	Address       string       `json:"address"`
	WalletBalance sjson.Number `json:"wallet_balance"`
	ExecTime      sjson.Number `json:"exec_time"`
	CrossSeq      sjson.Number `json:"cross_seq"`
}

type WalletFundRecordResponse struct {
	BaseResult
	Result OrderListResponseArray `json:"result"`
}

type OrderListResponseArray struct {
	Data []WalletFundRecord `json:"data"`
}

type Funding struct {
	CurrentPage  int           `json:"current_page"`
	Data         []FundingData `json:"data"`
	FirstPageUrl string        `json:"first_page_url"`
	From         int           `json:"from"`
	LastPage     int           `json:"last_page"`
	LastPageUrl  string        `json:"last_page_url"`
	NextPageUrl  string        `json:"next_page_url"`
	Path         string        `json:"path"`
	PerPage      sjson.Number  `json:"per_page"`
	PrevPageUrl  string        `json:"prev_page_url"`
	To           int           `json:"to"`
	Total        int           `json:"total"`
}

type FundingData struct {
	Id     int          `json:"id"`
	Symbol string       `json:"symbol"`
	Value  sjson.Number `json:"value"`
	Time   string       `json:"time"`
}

type FundingResult struct {
	BaseResult
	Result Funding `json:"result"`
}

type IndexOHLC struct {
	Id      int          `json:"id"`
	Symbol  string       `json:"symbol"`
	Open    sjson.Number `json:"open"`
	High    sjson.Number `json:"high"`
	Low     sjson.Number `json:"low"`
	Close   sjson.Number `json:"close"`
	StartAt int          `json:"start_at"`
	Period  string       `json:"period"`
}

type IndexOHLCResult struct {
	BaseResult
	Result []IndexOHLC `json:"result"`
}

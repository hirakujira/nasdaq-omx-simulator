// Nordic Equity TotalView ITCH protocol
//
// This implementation is based on the following specification provided by
// NASDAQ OMX:
//
//   Nordic Equity TotalView-ITCH
//   Version 1.90.2
//   April 7, 2014

package main

type Seconds struct {
	MsgType byte
	Second  [5]byte
}

type Milliseconds struct {
	MsgType     byte
	Millisecond [3]byte
}

type SystemEvent struct {
	MsgType   byte
	StockLocate [2]byte
	TrackingNumber [2]byte
	Timestamp [6]byte
	EventCode byte
}

type StockDirectory struct {
	MsgType                     byte
	StockLocate                 [2]byte
	TrackingNumber              [2]byte
	Timestamp                   [6]byte
	Stock                       [8]byte
	MarketCategory              byte
	FinancialStatusIndicator    byte
	RoundLotSize                [4]byte
	RoundLotsOnly               byte
	IssueClassification         byte
	IssueSubtype                [2]byte
	Authenticity                byte
	ShortSaleThresholdIndicator byte
	IPOFlag                     byte
	LULDReferencePriceTier      byte
	ETPFlag                     byte
	ETPLeverageFactor           [4]byte
	InverseIndicator            byte
}

type StockTradingAction struct {
	MsgType            byte
	StockLocate        [2]byte
	TrackingNumber     [2]byte
	Timestamp          [6]byte
	Stock              [8]byte
	TradingState       byte
	Reserved           byte
	Reason             [4]byte
}

type RegSHORestriction struct {
	MsgType            byte
	StockLocate        [2]byte
	TrackingNumber     [2]byte
	Timestamp          [6]byte
	Stock              [8]byte
	RegSHOAction       byte
}

type MarketPartcipantPOS struct {
	MsgType            byte
	StockLocate        [2]byte
	TrackingNumber     [2]byte
	Timestamp          [6]byte
	MPID               [4]byte
	Stock              [8]byte
	PrimaryMarketMaker byte
	MarketMakerMode    byte
	MarketPartcipantS  byte
}

type MWCBDeclineLevel struct {
	MsgType            byte
	StockLocate        [2]byte
	TrackingNumber     [2]byte
	Timestamp          [6]byte
	Level1 	           [8]byte
	Level2 	           [8]byte
	Level3 	           [8]byte
}

type MWCBStatus struct {
	MsgType            byte
	StockLocate        [2]byte
	TrackingNumber     [2]byte
	Timestamp          [6]byte
	BreachedLevel      byte
}

type IPOQuotingPeriodUpdate struct {
	MsgType            byte
	StockLocate        [2]byte
	TrackingNumber     [2]byte
	Timestamp          [6]byte
	Stock              [8]byte
	IPOQuotationRelT   [4]byte
	IPOQuotationRelQ   byte
	IPOPrice           [4]byte
}

type AddOrder struct {
	MsgType            byte
	StockLocate        [2]byte
	TrackingNumber     [2]byte
	Timestamp          [6]byte
	OrderReferenceNum  [8]byte
	BuySellIndicator   byte
	Shares             [4]byte
	Stock              [8]byte
	Price              [4]byte
}

type AddOrderMPID struct {
	MsgType            byte
	StockLocate        [2]byte
	TrackingNumber     [2]byte
	Timestamp          [6]byte
	OrderReferenceNum  [8]byte
	BuySellIndicator   byte
	Shares             [4]byte
	Stock              [8]byte
	Price              [4]byte
	Attribution        [4]byte
}

type OrderExecuted struct {
	MsgType            byte
	StockLocate        [2]byte
	TrackingNumber     [2]byte
	Timestamp          [6]byte
	OrderReferenceNum  [8]byte
	ExecutedShares     [4]byte
	MatchNumber        [8]byte
}

type OrderExecutedWithPrice struct {
	MsgType            byte
	StockLocate        [2]byte
	TrackingNumber     [2]byte
	Timestamp          [6]byte
	OrderReferenceNum  [8]byte
	ExecutedShares     [4]byte
	MatchNumber        [8]byte
	Printable          byte
	ExecutionPrice     [4]byte
}

type OrderCancel struct {
	MsgType            byte
	StockLocate        [2]byte
	TrackingNumber     [2]byte
	Timestamp          [6]byte
	OrderReferenceNum  [8]byte
	CanceledShares     [4]byte
}

type OrderDelete struct {
	MsgType            byte
	StockLocate        [2]byte
	TrackingNumber     [2]byte
	Timestamp          [6]byte
	OrderReferenceNum  [8]byte
}


type OrderReplace struct {
	MsgType            byte
	StockLocate        [2]byte
	TrackingNumber     [2]byte
	Timestamp          [6]byte
	OriginalOrderRefN  [8]byte
	NewOrderRefN       [8]byte
	Shares             [4]byte
	Price              [4]byte
}

type Trade struct {
	MsgType            byte
	StockLocate        [2]byte
	TrackingNumber     [2]byte
	Timestamp          [6]byte
	OrderReferenceNum  [8]byte
	BuySellIndicator   byte
	Shares             [4]byte
	Stock              [8]byte
	Price              [4]byte
	MatchNumber        [8]byte
}

type CrossTrade struct {
	MsgType            byte
	StockLocate        [2]byte
	TrackingNumber     [2]byte
	Timestamp          [6]byte
	Shares             [8]byte
	Stock              [8]byte
	CrossPrice         [4]byte
	MatchNumber        [8]byte
	CrossType          byte
}

type BrokenTrade struct {
	MsgType            byte
	StockLocate        [2]byte
	TrackingNumber     [2]byte
	Timestamp          [6]byte
	MatchNumber        [8]byte
}

type NOII struct {
	MsgType            byte
	StockLocate        [2]byte
	TrackingNumber     [2]byte
	Timestamp          [6]byte
	PairedShares       [8]byte
	ImbalanceShares    [8]byte
	ImbalanceDirection byte
	Stock              [8]byte
	FarPrice           [4]byte
	NearPrice          [4]byte
	CurrentReferenceP  [4]byte
	CrossType          byte
	PriceVariationIndi byte
}

type RPII struct {
	MsgType            byte
	StockLocate        [2]byte
	TrackingNumber     [2]byte
	Timestamp          [6]byte
	Stock              [8]byte
	InterestFlag       byte
}

func ItchUatoi(buf []byte, len int) uint64 {
	var ret uint64 = 0
	var idx int = 0
	for idx < len {
		if buf[idx] != ' ' {
			break
		}
		idx++
	}
	for idx < len {
		ch := buf[idx]
		ret = (ret * 10) + uint64(ch-byte('0'))
		idx += 1
	}
	return ret
}

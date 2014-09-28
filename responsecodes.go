package ib

var RESPONSE ResponseHeader

type ResponseHeader struct {
	CODE ResponseCodes
}

type ResponseCodes struct {
	ERR_MSG                  string
	CONTRACT_DATA            string
	TICK_PRICE               string
	TICK_SIZE                string
	TICK_OPTION_COMPUTATION  string
	TICK_GENERIC             string
	TICK_STRING              string
	TICK_EFP                 string
	TICK_SNAPSHOT_END        string
	MARKET_DATA_TYPE         string
	MARKET_DEPTH string
	MARKET_DEPTH_LEVEL_TWO string
	HISTORICAL_DATA          string
	ACCOUNT_VALUE            string
	PORTFOLIO_VALUE          string
	ACCOUNT_UPDATE_TIME      string
	ORDER_STATUS             string
	OPEN_ORDER               string
	OPEN_ORDER_END           string
	NEXT_VALID_ID            string
	DELTA_NEUTRAL_VALIDATION string
}

func init() {
	RESPONSE.CODE.ERR_MSG = "4"
	RESPONSE.CODE.CONTRACT_DATA = "10"
	RESPONSE.CODE.TICK_PRICE = "1"
	RESPONSE.CODE.TICK_SIZE = "2"
	RESPONSE.CODE.TICK_OPTION_COMPUTATION = "21"
	RESPONSE.CODE.TICK_GENERIC = "45"
	RESPONSE.CODE.TICK_STRING = "46"
	RESPONSE.CODE.TICK_EFP = "47"
	RESPONSE.CODE.TICK_SNAPSHOT_END = "57"
	RESPONSE.CODE.MARKET_DATA_TYPE = "58"
	RESPONSE.CODE.MARKET_DEPTH = "12"
	RESPONSE.CODE.MARKET_DEPTH_LEVEL_TWO = "13"
	RESPONSE.CODE.HISTORICAL_DATA = "17"
	RESPONSE.CODE.ACCOUNT_VALUE = "6"
	RESPONSE.CODE.PORTFOLIO_VALUE = "7"
	RESPONSE.CODE.ACCOUNT_UPDATE_TIME = "8"
	RESPONSE.CODE.ORDER_STATUS = "3"
	RESPONSE.CODE.OPEN_ORDER = "5"
	RESPONSE.CODE.OPEN_ORDER_END = "52"
	RESPONSE.CODE.NEXT_VALID_ID = "9"
	RESPONSE.CODE.DELTA_NEUTRAL_VALIDATION = "56"
}

package tpcc

import (
	"github.com/slocke716/go-tpcc/helpers"
)

//
// @TODO@
//
// This is taken directly from the py-tpcc benchmark
// and should be refactored later.
//

const (
	MONEY_DECIMALS = 2
	// Item constants
	NUM_ITEMS  = 100000
	MIN_IM     = 1
	MAX_IM     = 10000
	MIN_PRICE  = 1.00
	MAX_PRICE  = 100.00
	MIN_I_NAME = 14
	MAX_I_NAME = 24
	MIN_I_DATA = 26
	MAX_I_DATA = 50

	//  Warehouse constants
	MIN_TAX       = 0
	MAX_TAX       = 0.2000
	TAX_DECIMALS  = 4
	INITIAL_W_YTD = 300000.00
	MIN_NAME      = 6
	MAX_NAME      = 10
	MIN_STREET    = 10
	MAX_STREET    = 20
	MIN_CITY      = 10
	MAX_CITY      = 20
	STATE         = 2
	ZIP_LENGTH    = 9
	ZIP_SUFFIX    = "11111"

	// Stock constants
	MIN_QUANTITY        = 10
	MAX_QUANTITY        = 100
	DIST                = 24
	STOCK_PER_WAREHOUSE = 100000

	//  District constants
	DISTRICTS_PER_WAREHOUSE = 10
	INITIAL_D_YTD           = 30000.00
	INITIAL_NEXT_O_ID       = 3001

	// Customer constants
	CUSTOMERS_PER_DISTRICT = 3000
	INITIAL_CREDIT_LIM     = 50000.00
	MIN_DISCOUNT           = 0.0000
	MAX_DISCOUNT           = 0.5000
	DISCOUNT_DECIMALS      = 4
	INITIAL_BALANCE        = -10.00
	INITIAL_YTD_PAYMENT    = 10.00
	INITIAL_PAYMENT_CNT    = 1
	INITIAL_DELIVERY_CNT   = 0
	MIN_FIRST              = 6
	MAX_FIRST              = 10
	MIDDLE                 = "OE"
	PHONE                  = 16
	MIN_C_DATA             = 300
	MAX_C_DATA             = 500
	GOOD_CREDIT            = "GC"
	BAD_CREDIT             = "BC"
	//  Order constants
	MIN_CARRIER_ID = 1
	MAX_CARRIER_ID = 10
	//  HACK: This is not strictly correct, but it works
	NULL_CARRIER_ID = 0
	//  O_ID < than this value, carrier != null, >= -> carrier == null
	NULL_CARRIER_LOWER_BOUND    = 2101
	MIN_OL_CNT                  = 5
	MAX_OL_CNT                  = 15
	INITIAL_ALL_LOCAL           = 1
	INITIAL_ORDERS_PER_DISTRICT = 3000

	//  Used to generate new order transactions
	MAX_OL_QUANTITY = 10

	//  Order line constants
	INITIAL_QUANTITY = 5
	MIN_AMOUNT       = 0.01

	//  History constants
	MIN_DATA       = 12
	MAX_DATA       = 24
	INITIAL_AMOUNT = 10.00

	//  New order constants
	INITIAL_NEW_ORDERS_PER_DISTRICT = 900

	//  TPC-C 2.4.3.4 (page 31) says this must be displayed when new order rolls back.
	INVALID_ITEM_MESSAGE = "Item number is not valid"

	//  Used to generate stock level transactions
	MIN_STOCK_LEVEL_THRESHOLD = 10
	MAX_STOCK_LEVEL_THRESHOLD = 20

	//  Used to generate payment transactions
	MIN_PAYMENT = 1.0
	MAX_PAYMENT = 5000.0

	//  Indicates "brand" items and stock in I_DATA and s_data.
	ORIGINAL_STRING = "ORIGINAL"

	// Table Names
	TABLENAME_ITEM       = "ITEM"
	TABLENAME_WAREHOUSE  = "WAREHOUSE"
	TABLENAME_DISTRICT   = "DISTRICT"
	TABLENAME_CUSTOMER   = "CUSTOMER"
	TABLENAME_STOCK      = "STOCK"
	TABLENAME_ORDERS     = "ORDERS"
	TABLENAME_NEW_ORDER  = "NEW_ORDER"
	TABLENAME_ORDER_LINE = "ORDER_LINE"
	TABLENAME_HISTORY    = "HISTORY"
)

var SYLLABLES = [...]string{"BAR", "OUGHT", "ABLE", "PRI", "PRES", "ESE", "ANTI", "CALLY", "ATION", "EING"}

type Address struct {
	street_1 string
	street_2 string
	city     string
	state    string
	zip      string
}

func (w *Worker) generateRandomAddress() Address {
	zipLength := ZIP_LENGTH - len(ZIP_SUFFIX)

	return Address{
		street_1: helpers.RandString(helpers.RandInt(MIN_STREET, MAX_STREET)),
		street_2: helpers.RandString(helpers.RandInt(MIN_STREET, MAX_STREET)),
		city:     helpers.RandString(helpers.RandInt(MIN_CITY, MAX_CITY)),
		state:    helpers.RandString(STATE),
		zip:      helpers.RandNumericString(zipLength) + ZIP_SUFFIX,
	}
}

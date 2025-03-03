package postgresql

import "context"

func (db *PostgreSQL) CreateSchema() error {

	tables := []string{`
CREATE TABLE IF NOT EXISTS WAREHOUSE (
	W_ID smallint not null,
	W_NAME varchar(10), 
	W_STREET_1 varchar(20), 
	W_STREET_2 varchar(20), 
	W_CITY varchar(20), 
	W_STATE char(2), 
	W_ZIP char(9), 
	W_TAX decimal(4,2), 
	W_YTD decimal(12,2),
	PRIMARY KEY (W_ID))`, `
CREATE TABLE STOCK (
  S_I_ID int NOT NULL,
  S_W_ID smallint NOT NULL,
  S_QUANTITY smallint DEFAULT NULL,
  S_DIST_01 char(24) DEFAULT NULL,
  S_DIST_02 char(24) DEFAULT NULL,
  S_DIST_03 char(24) DEFAULT NULL,
  S_DIST_04 char(24) DEFAULT NULL,
  S_DIST_05 char(24) DEFAULT NULL,
  S_DIST_06 char(24) DEFAULT NULL,
  S_DIST_07 char(24) DEFAULT NULL,
  S_DIST_08 char(24) DEFAULT NULL,
  S_DIST_09 char(24) DEFAULT NULL,
  S_DIST_10 char(24) DEFAULT NULL,
  S_YTD decimal(8,0) DEFAULT NULL,
  S_ORDER_CNT smallint DEFAULT NULL,
  S_REMOTE_CNT smallint DEFAULT NULL,
  S_DATA varchar(50) DEFAULT NULL,
  PRIMARY KEY (S_W_ID,S_I_ID))
`, `
CREATE TABLE ORDERS (
  O_ID int NOT NULL,
  O_D_ID smallint NOT NULL,
  O_W_ID smallint NOT NULL,
  O_C_ID int DEFAULT NULL,
  O_ENTRY_D timestamp DEFAULT NULL,
  O_CARRIER_ID smallint DEFAULT NULL,
  O_OL_CNT smallint DEFAULT NULL,
  O_ALL_LOCAL smallint DEFAULT NULL,
  PRIMARY KEY (O_W_ID,O_D_ID,O_ID)
 )
`, `
CREATE TABLE ORDER_LINE (
  OL_O_ID int NOT NULL,
  OL_D_ID smallint NOT NULL,
  OL_W_ID smallint NOT NULL,
  OL_NUMBER smallint NOT NULL,
  OL_I_ID int DEFAULT NULL,
  OL_SUPPLY_W_ID smallint DEFAULT NULL,
  OL_DELIVERY_D timestamp DEFAULT NULL,
  OL_QUANTITY smallint DEFAULT NULL,
  OL_AMOUNT decimal(6,2) DEFAULT NULL,
  OL_DIST_INFO char(24) DEFAULT NULL,
  PRIMARY KEY (OL_W_ID,OL_D_ID,OL_O_ID,OL_NUMBER))
`, `
 CREATE TABLE NEW_ORDER (
  NO_O_ID int NOT NULL,
  NO_D_ID smallint NOT NULL,
  NO_W_ID smallint NOT NULL,
  PRIMARY KEY (NO_W_ID,NO_D_ID,NO_O_ID))
`, `
CREATE TABLE ITEM (
  I_ID int NOT NULL,
  I_IM_ID int DEFAULT NULL,
  I_NAME varchar(24) DEFAULT NULL,
  I_PRICE decimal(5,2) DEFAULT NULL,
  I_DATA varchar(50) DEFAULT NULL,
  PRIMARY KEY (I_ID))
`,
		`
CREATE TABLE HISTORY (
  H_C_ID int DEFAULT NULL,
  H_C_D_ID smallint DEFAULT NULL,
  H_C_W_ID smallint DEFAULT NULL,
  H_D_ID smallint DEFAULT NULL,
  H_W_ID smallint DEFAULT NULL,
  H_DATE timestamp DEFAULT NULL,
  H_AMOUNT decimal(6,2) DEFAULT NULL,
  H_DATA varchar(24) DEFAULT NULL)
`, `
CREATE TABLE DISTRICT (
  D_ID smallint NOT NULL,
  D_W_ID smallint NOT NULL,
  D_NAME varchar(10) DEFAULT NULL,
  D_STREET_1 varchar(20) DEFAULT NULL,
  D_STREET_2 varchar(20) DEFAULT NULL,
  D_CITY varchar(20) DEFAULT NULL,
  D_STATE char(2) DEFAULT NULL,
  D_ZIP char(9) DEFAULT NULL,
  D_TAX decimal(4,2) DEFAULT NULL,
  D_YTD decimal(12,2) DEFAULT NULL,
  D_NEXT_O_ID int DEFAULT NULL,
  PRIMARY KEY (D_W_ID,D_ID))
`, `
 CREATE TABLE CUSTOMER (
  C_ID int NOT NULL,
  C_D_ID smallint NOT NULL,
  C_W_ID smallint NOT NULL,
  C_FIRST varchar(16) DEFAULT NULL,
  C_MIDDLE char(2) DEFAULT NULL,
  C_LAST varchar(16) DEFAULT NULL,
  C_STREET_1 varchar(20) DEFAULT NULL,
  C_STREET_2 varchar(20) DEFAULT NULL,
  C_CITY varchar(20) DEFAULT NULL,
  C_STATE char(2) DEFAULT NULL,
  C_ZIP char(9) DEFAULT NULL,
  C_PHONE char(16) DEFAULT NULL,
  C_SINCE timestamp DEFAULT NULL,
  C_CREDIT char(2) DEFAULT NULL,
  C_CREDIT_LIM bigint DEFAULT NULL,
  C_DISCOUNT decimal(4,2) DEFAULT NULL,
  C_BALANCE decimal(12,2) DEFAULT NULL,
  C_YTD_PAYMENT decimal(12,2) DEFAULT NULL,
  C_PAYMENT_CNT smallint DEFAULT NULL,
  C_DELIVERY_CNT smallint DEFAULT NULL,
  C_DATA text,
  PRIMARY KEY (C_W_ID,C_D_ID,C_ID))
`}

	for _, table := range tables {
		_, err := db.Client.Exec(context.Background(), table)
		if err != nil {
			return err
		}
	}

	return nil
}

func (db *PostgreSQL) CreateIndexes() error {

	queries := []string{
		"CREATE INDEX idx_customer on CUSTOMER (C_W_ID,C_D_ID,C_LAST,C_FIRST)",
		"CREATE INDEX idx_orders  ON ORDERS  (O_W_ID,O_D_ID,O_C_ID,O_ID)",
		"CREATE INDEX fkey_stock_2 ON STOCK (S_I_ID)",
		"CREATE INDEX fkey_order_line_2 ON ORDER_LINE (OL_SUPPLY_W_ID,OL_I_ID)",
		"CREATE INDEX fkey_history_1 ON HISTORY (H_C_W_ID,H_C_D_ID,H_C_ID)",
		"CREATE INDEX fkey_history_2 ON HISTORY (H_W_ID,H_D_ID)",
	}

	if db.fk {
		fkq := []string{
			"ALTER TABLE NEW_ORDER ADD CONSTRAINT fkey_new_orders_1 FOREIGN KEY(NO_W_ID,NO_D_ID,NO_O_ID) REFERENCES ORDERS(O_W_ID,O_D_ID,O_ID)",
			"ALTER TABLE ORDERS ADD CONSTRAINT fkey_orders_1 FOREIGN KEY(O_W_ID,O_D_ID,O_C_ID) REFERENCES CUSTOMER(C_W_ID,C_D_ID,C_ID)",
			"ALTER TABLE CUSTOMER ADD CONSTRAINT fkey_customer_1 FOREIGN KEY(C_W_ID,C_D_ID) REFERENCES DISTRICT(D_W_ID,D_ID)",
			"ALTER TABLE HISTORY ADD CONSTRAINT fkey_history_1 FOREIGN KEY(H_C_W_ID,H_C_D_ID,H_C_ID) REFERENCES CUSTOMER(C_W_ID,C_D_ID,C_ID)",
			"ALTER TABLE HISTORY ADD CONSTRAINT fkey_history_2 FOREIGN KEY(H_W_ID,H_D_ID) REFERENCES DISTRICT(D_W_ID,D_ID)",
			"ALTER TABLE DISTRICT ADD CONSTRAINT fkey_district_1 FOREIGN KEY(D_W_ID) REFERENCES WAREHOUSE(W_ID)",
			"ALTER TABLE ORDER_LINE ADD CONSTRAINT fkey_order_line_1 FOREIGN KEY(OL_W_ID,OL_D_ID,OL_O_ID) REFERENCES ORDERS(O_W_ID,O_D_ID,O_ID)",
			"ALTER TABLE ORDER_LINE ADD CONSTRAINT fkey_order_line_2 FOREIGN KEY(OL_SUPPLY_W_ID,OL_I_ID) REFERENCES STOCK(S_W_ID,S_I_ID)",
			"ALTER TABLE STOCK ADD CONSTRAINT fkey_stock_1 FOREIGN KEY(S_W_ID) REFERENCES WAREHOUSE(W_ID)",
			"ALTER TABLE STOCK ADD CONSTRAINT fkey_stock_2 FOREIGN KEY(S_I_ID) REFERENCES ITEM(I_ID)",
		}

		queries = append(queries, fkq...)
	}
	for _, query := range queries {
		_, err := db.Client.Exec(context.Background(), query)
		if err != nil {
			return err
		}
	}

	return nil
}

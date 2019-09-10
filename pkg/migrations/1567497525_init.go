package migrations

import migrate "github.com/rubenv/sql-migrate"

func init() {
	instance.add(&migrate.Migration{
		Id: "1567497525",

		Up:[]string{`
			CREATE TABLE products(
			ID bigint(20) NOT NULL AUTO_INCREMENT,
  			Created bigint(20) DEFAULT NULL,
  			Updated bigint(20) DEFAULT NULL,
  			Deleted tinyint(1) DEFAULT 0,
			URL varchar(512) NOT NULL,
			Title varchar(255) DEFAULT NULL,
			Price float(15,2) DEFAULT NULL,
			CompanyName varchar(100) DEFAULT NULL,
			Status int NOT NULL,
			PRIMARY KEY (ID),
			UNIQUE(URL)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8;`,
		},
		Down: []string{`
			DROP TABLE products;
		`},
	})
}
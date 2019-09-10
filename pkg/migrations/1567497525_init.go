package migrations

import migrate "github.com/rubenv/sql-migrate"

func init() {
	instance.add(&migrate.Migration{
		Id: "1567426004",

		Up:[]string{`
			CREATE TABLE links(
			ID bigint(20) NOT NULL AUTO_INCREMENT,
  			Created bigint(20) DEFAULT NULL,
  			Updated bigint(20) DEFAULT NULL,
  			Deleted tinyint(1) DEFAULT 0,
			URL varchar(65535) NOT NULL,
			PRIMARY KEY (ID),
			UNIQUE(URL)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8;
			`,`
			CREATE TABLE products(
			ID bigint(20) NOT NULL AUTO_INCREMENT,
  			Created bigint(20) DEFAULT NULL,
  			Updated bigint(20) DEFAULT NULL,
  			Deleted tinyint(1) DEFAULT 0,
			LinkID bigint(20) NOT NULL,
			ProductTitle varchar(255),
			Price float(6,2),
			CompanyName varchar(100),
			Status int,
			PRIMARY KEY (ID),
			CONSTRAINT FK_LinkID FOREIGN KEY (LinkID)
    		REFERENCES links(ID),
			CONSTRAINT CHK_Status CHECK (Status>=0 AND Status<=2)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8;`,
		},
		Down: []string{`
			DROP TABLE links;
			DROP TABLE products;
		`},
	})
}
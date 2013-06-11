package main

import (
  _ "github.com/lib/pq"
  "github.com/coopernurse/gorp"
)

import (
  "database/sql"
  "log"
)

type DBObject struct {
  AOID string `db:"AOID"`
  // AOGUID string `db:"AOGUID"`
  // PARENTGUID string `db:"PARENTGUID"`
  // FORMALNAME string `db:"FORMALNAME"`
  // OFFNAME string `db:"OFFNAME"`
  // SHORTNAME string `db:"SHORTNAME"`
  // AOLEVEL string `db:"AOLEVEL"`
  // REGIONCODE string `db:"REGIONCODE"`
  // AREACODE string `db:"AREACODE"`
  // AUTOCODE string `db:"AUTOCODE"`
  // CITYCODE string `db:"CITYCODE"`
  // CTARCODE string `db:"CTARCODE"`
  // PLACECODE string `db:"PLACECODE"`
  // STREETCODE string `db:"STREETCODE"`
  // EXTRCODE string `db:"EXTRCODE"`
  // SEXTCODE string `db:"SEXTCODE"`
  // PLAINCODE string `db:"PLAINCODE"`
  // CODE string `db:"CODE"`
  // CURRSTATUS string `db:"CURRSTATUS"`
  // ACTSTATUS string `db:"ACTSTATUS"`
  // LIVESTATUS string `db:"LIVESTATUS"`
  // CENTSTATUS string `db:"CENTSTATUS"`
  // OPERSTATUS string `db:"OPERSTATUS"`
  // IFNSFL string `db:"IFNSFL"`
  // IFNSUL string `db:"IFNSUL"`
  // TERRIFNSFL string `db:"TERRIFNSFL"`
  // TERRIFNSUL string `db:"TERRIFNSUL"`
  // OKATO string `db:"OKATO"`
  // OKTMO string `db:"OKTMO"`
  // POSTALCODE string `db:"POSTALCODE"`
  // STARTDATE string `db:"STARTDATE"`
  // ENDDATE string `db:"ENDDATE"`
  // UPDATEDATE string `db:"UPDATEDATE"`
}

func main() {
  db, err := sql.Open("postgres", "user=dev dbname=test_db password=dev")

  if err != nil {
    log.Fatal(err)
  }

  dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
  dbmap.AddTableWithName(DBObject{}, "objects")

  // create all registered tables
  dbmap.DropTables()
  dbmap.CreateTables()

  obj1 := &DBObject{"test"}
  dbmap.Insert(obj1)
}
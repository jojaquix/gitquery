package engine

import (
	gosql "database/sql"
	"gopkg.in/sqle/sqle.v0"
	"gopkg.in/sqle/sqle.v0/sql"
	fixtures "gopkg.in/src-d/go-git-fixtures.v3"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	//"gopkg.in/src-d/go-git-fixtures.v3"
)

func init() {
	fixtures.RootFolder = "vendor/gopkg.in/src-d/go-git-fixtures.v3/"
}

const (
	testDBName = "sysInfo"
)

func TestDatabase_Tables(t *testing.T) {
	assert := assert.New(t)

	f := fixtures.Basic().One()
	db := getDB(assert, f, testDBName)

	tables := db.Tables()
	var tableNames []string
	for key := range tables {
		tableNames = append(tableNames, key)
	}

	sort.Strings(tableNames)
	expected := []string{
		"os_version",
		//commitsTableName,
		//referencesTableName,
		//treeEntriesTableName,
		//tagsTableName,
		//blobsTableName,
		//objectsTableName,
	}
	sort.Strings(expected)

	assert.Equal(expected, tableNames)
}

func TestOsVersionQuery(t *testing.T) {
	assert := assert.New(t)
	f := fixtures.Basic().One()
	db := getDB(assert, f, testDBName)

	sqle.DefaultEngine.AddDatabase(db)
	cx, err := gosql.Open(sqle.DriverName, "")
	if err != nil {
		t.Fatal(err)
	}

	res, err := cx.Query("select * from os_version where major = int64(10)")

	if err != nil {
		t.Fatal(err)
	}

	for res.Next() {

	}

}

func TestDatabase_Name(t *testing.T) {
	assert := assert.New(t)

	f := fixtures.Basic().One()
	db := getDB(assert, f, testDBName)
	assert.Equal(testDBName, db.Name())
}

func getDB(assert *assert.Assertions, fixture *fixtures.Fixture,
	name string) sql.Database {

	db := NewDatabase(name)
	assert.NotNil(db)

	return db
}

func getTable(assert *assert.Assertions, fixture *fixtures.Fixture,
	name string) sql.Table {

	db := getDB(assert, fixture, "foo")
	assert.NotNil(db)
	assert.Equal(db.Name(), "foo")

	tables := db.Tables()
	table, ok := tables[name]
	assert.True(ok, "table %s does not exist", table)
	assert.NotNil(table)

	return table
}
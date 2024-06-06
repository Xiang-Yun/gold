package repository

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"time"

	_ "github.com/glebarez/go-sqlite"
)

var testRepo *SQLiteRepository

func TestMain(m *testing.M) {
	_ = os.Remove("./testdata/sql.db")
	path := "./testdata/sql.db"
	db, err := sql.Open("sqlite", path)
	if err != nil {
		log.Println(err)
	}

	testRepo = NewSQLiteRepository(db)
	testRepo.Migrate()
	os.Exit(m.Run())
}

func TestSQLiteRepository_InsertHolding(t *testing.T) {
	h := Holdings{
		Amount:        1,
		PurchaseDate:  time.Now().Local(),
		PurchasePrice: 1000,
	}

	result, err := testRepo.InsertHolding(h)
	if err != nil {
		t.Error("insert failed:", err)
	}

	if result.ID <= 0 {
		t.Error("invalid id send back:", result.ID)
	}
}

func TestSQLiteRepository_AllHoldings(t *testing.T) {
	h, err := testRepo.AllHoldings()
	if err != nil {
		t.Error("allHolding failed", err)
	}
	if len(h) != 1 {
		t.Error("wrong number of rows returned; expected 1, but get", len(h))
	}
}

func TestSQLiteRepository_GetHoldingByID(t *testing.T) {
	h, err := testRepo.GetHoldingByID(1)
	if err != nil {
		t.Error("wrong purchase price returned; expected 1000 but got", h.PurchasePrice)
	}

	_, err = testRepo.GetHoldingByID(2)
	if err == nil {
		t.Error("get one return")
	}
}

func TestSQLiteRepository_UpdateHolding(t *testing.T) {
	h, err := testRepo.GetHoldingByID(1)
	if err != nil {
		t.Error(err)
	}

	h.PurchasePrice = 1001

	err = testRepo.UpdateHolding(1, *h)
	if err != nil {
		t.Error("update failed:", err)
	}
}

func TestSQLiteRepository_DeleteHolding(t *testing.T) {
	err := testRepo.DeleteHolding(1)
	if err != nil {
		t.Error("failed to delete holding", err)
		if err != errDeleteFailed {
			t.Error("wrong error returned")
		}
	}

	err = testRepo.DeleteHolding(2)
	if err == nil {
		t.Error("no error when trying to delete non-existant record")
	}
}

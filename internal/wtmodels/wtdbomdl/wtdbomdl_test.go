package wtdbomdl

import (
	"database/sql"
	"testing"

	"github.com/syukrisuw/wanthatgo/internal/wtdatabase"
)

func TestInitDbWrongDbType(test *testing.T) {
	_, err := wtdatabase.InitDb("postgresql", "user=alaqshapbdbadmin password=5ubh4nALLAH#33x dbname=alaqshapb sslmode=disable port=5432")
	if err == nil {
		test.Fatalf("test should be error because using wrong, and error is %v", err)
	}
	test.Logf("negative test success with err:%v", err)
}

func TestInitDbWrongDbname(test *testing.T) {
	_, err := wtdatabase.InitDb("postgres", "user=alaqshapbdbadmin password=test dbname=alaqshadb sslmode=disable port=5432")
	if err == nil {
		test.Fatalf("test should be error because using wrong dbname, and error is %v", err)
	}
	test.Logf("negative test success with err:%v", err)
}

func TestInitDbWrongUserName(test *testing.T) {
	_, err := wtdatabase.InitDb("postgres", "user=test password=5ubh4nALLAH#33x dbname=alaqshapb sslmode=disable port=5432")
	if err == nil {
		test.Fatalf("test should be error because using wrong username, and error is %v", err)
	}
	test.Logf("negative test success with err:%v", err)
}

func TestInitDbWrongPassword(test *testing.T) {
	_, err := wtdatabase.InitDb("postgres", "user=alaqshapbdbadmin password=test dbname=alaqshapb sslmode=disable port=5432")
	if err == nil {
		test.Fatalf("test should be error because using wrong password, and error is %v", err)
	}
	test.Logf("negative test success with err:%v", err)
}

func TestInitDbWrongPort(test *testing.T) {
	_, err := wtdatabase.InitDb("postgres", "user=alaqshapbdbadmin password=test dbname=alaqshapb sslmode=disable port=5433")
	if err == nil {
		test.Fatalf("test should be error because using wrong port, and error is %v", err)
	}
	test.Logf("negative test success with err:%v", err)
}

func TestInitDb(test *testing.T) {
	_, err := wtdatabase.InitDb("postgres", "user=alaqshapbdbadmin password=5ubh4nALLAH#33x dbname=alaqshapb sslmode=disable port=5432")
	if err != nil {
		test.Fatalf("test initialize db failed with error: %v", err)
	}
	test.Logf("positive test success with err:%v", err)
}

func TestGetDbUninitialized(test *testing.T) {
	// wtdb = nil
	wtdatabase.ClearDb()

	_, err := wtdatabase.GetDb()
	if err == nil {
		test.Fatalf("test should be error because db is not initialized yet, and error is %v", err)
	}
	test.Logf("negative test success with err:%v", err)
}

func TestGetDb(test *testing.T) {
	var wtdb *sql.DB
	var err error
	_, err = wtdatabase.InitDb("postgres", "user=alaqshapbdbadmin password=5ubh4nALLAH#33x dbname=alaqshapb sslmode=disable port=5432")
	if err != nil {
		test.Fatalf("test initialize db failed with error: %v", err)
	}

	wtdb, err = wtdatabase.GetDb()
	if err != nil {
		test.Fatalf("test get db failed with error:  %v", err)
	}
	err = wtdb.Ping()
	if err != nil {
		test.Fatalf("test get db failed with error:  %v", err)
	}

	test.Logf("positive test success with err:%v", err)
}

func TestTotalRowUnknownTable(test *testing.T) {
	var wtdb *sql.DB
	var err error
	_, err = wtdatabase.InitDb("postgres", "user=alaqshapbdbadmin password=5ubh4nALLAH#33x dbname=alaqshapb sslmode=disable port=5432")
	if err != nil {
		test.Fatalf("test initialize db failed with error: %v", err)
	}

	wtdb, err = wtdatabase.GetDb()
	if err != nil {
		test.Fatalf("test get db failed with error:  %v", err)
	}
	defer wtdb.Close()

	namaPenerima := ""
	tbPenerimaan := WtPenerimaanDbo{
		wtDboBase: wtDboBase{TblName: "tb_penerimaanzs"},
		WtUserDbo: WtUserDbo{NomorRef: "", NamaRef: &namaPenerima},
	}
	// fmt.Println("tabel name : ", tbPenerimaan.TblName)
	totalRow, err := tbPenerimaan.TotalRow()
	if err != nil {
		test.Fatalf("test overridden total row failed on unknown table with error:  %v", err)
	}

	test.Logf("positive overridden test success with total row:%v", totalRow)
}
func TestTotalRowPenerimaan(test *testing.T) {
	var wtdb *sql.DB
	var err error
	_, err = wtdatabase.InitDb("postgres", "user=alaqshapbdbadmin password=5ubh4nALLAH#33x dbname=alaqshapb sslmode=disable port=5432")
	if err != nil {
		test.Fatalf("test initialize db failed with error: %v", err)
	}

	wtdb, err = wtdatabase.GetDb()
	if err != nil {
		test.Fatalf("test get db failed with error:  %v", err)
	}
	defer wtdb.Close()

	namaPenerima := ""
	tbPenerimaan := WtPenerimaanDbo{
		wtDboBase: wtDboBase{TblName: "tb_penerimaans"},
		WtUserDbo: WtUserDbo{NomorRef: "", NamaRef: &namaPenerima},
	}
	// fmt.Println("tabel name : ", tbPenerimaan.TblName)
	totalRow, err := tbPenerimaan.TotalRow()
	if err != nil {
		test.Fatalf("test total row failed with error:  %v", err)
	}

	test.Logf("positive test success with totalRow:%v", totalRow)
}

func TestTotalRowPengeluaran(test *testing.T) {
	var wtdb *sql.DB
	var err error
	_, err = wtdatabase.InitDb("postgres", "user=alaqshapbdbadmin password=5ubh4nALLAH#33x dbname=alaqshapb sslmode=disable port=5432")
	if err != nil {
		test.Fatalf("test initialize db failed with error: %v", err)
	}

	wtdb, err = wtdatabase.GetDb()
	if err != nil {
		test.Fatalf("test get db failed with error:  %v", err)
	}
	defer wtdb.Close()

	namaPenerima := ""
	tbPengeluaran := WtPengeluaranDbo{
		wtDboBase: wtDboBase{TblName: "tb_pengeluarans"},
		WtUserDbo: WtUserDbo{NomorRef: "", NamaRef: &namaPenerima},
	}
	// fmt.Println("tabel name : ", tbPenerimaan.TblName)
	totalRow, err := tbPengeluaran.TotalRow()
	if err != nil {
		test.Fatalf("test total row failed with error:  %v", err)
	}

	test.Logf("positive test success with totalRow:%v", totalRow)
}

func TestTableNameUnset(test *testing.T) {

	namaPenerima := ""
	namaTabel := "tb_penerimaans"
	tbPenerimaan := WtPenerimaanDbo{
		wtDboBase: wtDboBase{TblDesc: &namaTabel},
		WtUserDbo: WtUserDbo{NomorRef: "", NamaRef: &namaPenerima},
	}
	tbName, err := tbPenerimaan.TableName()
	if err == nil {
		test.Fatalf("test empty TblName failed:  %v", tbName)
	}
	test.Logf("negative test success with err:%v", err)
}

func TestTableName(test *testing.T) {

	namaPenerima := ""
	tbPenerimaan := WtPenerimaanDbo{
		wtDboBase: wtDboBase{TblName: "tb_penerimaans"},
		WtUserDbo: WtUserDbo{NomorRef: "", NamaRef: &namaPenerima},
	}
	tbName, err := tbPenerimaan.TableName()
	if err != nil {
		test.Fatalf("test TblName failed with error:  %v", err)
	}

	test.Logf("positive test success with TblName:%v", tbName)
}

func TestNewTableUsingFactory(test *testing.T) {

	tbPenerimaan, err := FactoryOfTable(TbPenerimaan)
	if err != nil {
		test.Fatalf("test new table with factory failed with error:  %v", err)
	}

	tbPengeluaran, err := FactoryOfTable(TbPenerimaan)
	if err != nil {
		test.Fatalf("test new table with factory failed with error:  %v", err)
	}
	test.Logf("positive test success with TblName:%v", tbPenerimaan.getTabelName())
	test.Logf("positive test success with TblName:%v", tbPengeluaran.getTabelName())
}

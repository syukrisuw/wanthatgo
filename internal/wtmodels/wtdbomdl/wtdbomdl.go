package wtdbomdl

import (
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/syukrisuw/wanthatgo/internal/wtdatabase"
)

type WtInfo interface {
	TotalRow() (uint32, error)
}

type WtFind interface {
	FindById(uint64) (interface{}, error)
	FindByClause(string) (interface{}, error)
}

func (wtDbBase *wtDboBase) TableName() (string, error) {
	if wtDbBase.TblName == "" {
		return "", fmt.Errorf("tabel name not set yet")
	}

	return wtDbBase.TblName, nil
}

type wtDboBase struct {
	ID        uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	TblName   string
	TblDesc   *string
	WtInfo
	WtFind
}

func (wtDboBase *wtDboBase) TotalRow() (totalRow int64, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("count total row of the table failed, error :%v", r)
			totalRow = -1
		}
	}()
	var count int
	wtdb, err := wtdatabase.GetDb()
	if err != nil {
		return -1, err
	}
	// stmt, err := wtdb.Prepare("Select count(1) from 1$")
	// row := stmt.QueryRow(wtDboBase.TblName)

	queryString := fmt.Sprintf("Select count(1) from %s", wtDboBase.TblName)
	rows, err := wtdb.Query(queryString)

	if err != nil {
		return -1, err
	}
	rows.Next() //only care for the first result
	err = rows.Scan(&count)
	totalRow = int64(count)
	err = nil
	return
}

func (wtDboBase *WtPenerimaanDbo) TotalRow() (totalRow string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("count total row of the table failed, error :%v", r)
			totalRow = "failed to get total row"
		}
	}()
	var count int
	wtdb, err := wtdatabase.GetDb()
	if err != nil {
		return "failed to get total row", err
	}
	// stmt, err := wtdb.Prepare("Select count(1) from 1$")
	// row := stmt.QueryRow(wtDboBase.TblName)

	queryString := "Select count(1) from tb_penerimaans"
	rows, err := wtdb.Query(queryString)

	if err != nil {
		return "failed to get total row", err
	}
	rows.Next() //only care for the first result
	err = rows.Scan(&count)
	if err != nil {
		return "failed to get total row", err
	}
	totalRow = fmt.Sprintf("total row = %d", uint64(count))
	err = nil
	return
}

type WtUserDbo struct {
	wtDboBase
	NomorRef string
	NamaRef  *string
}

type WtPenerimaanDbo struct {
	wtDboBase
	WtUserDbo
}

type WtPengeluaranDbo struct {
	wtDboBase
	WtUserDbo
}

type iwtDboBase interface {
	setTableName(name string)
	setTabelDesc(desc string)
	getTabelName() string
	getTabelDesc() string
}

func (tbl *wtDboBase) setTableName(name string) {
	tbl.TblName = name
}

func (tbl *wtDboBase) getTabelName() string {
	return tbl.TblName
}

func (tbl *wtDboBase) setTabelDesc(desc string) {
	tbl.TblDesc = &desc
}

func (tbl *wtDboBase) getTabelDesc() string {
	return *tbl.TblDesc
}

func newWtPenerimaanDbo() iwtDboBase {
	return &WtPenerimaanDbo{
		wtDboBase: wtDboBase{
			TblName: "tb_penerimaans",
		},
	}
}

func newWtPengeluaranDbo() iwtDboBase {
	return &WtPengeluaranDbo{
		wtDboBase: wtDboBase{
			TblName: "tb_penerimaans",
		},
	}
}

type base int

const (
	TbPenerimaan base = iota
	TbPengeluaran
	TbNeracaKas
	TbKasPenerimaan
	TbKasPengeluaran
)

func FactoryOfTable(tblNameConst base) (iwtDboBase, error) {
	switch tblNameConst {
	case TbPenerimaan:
		{
			return newWtPenerimaanDbo(), nil
		}
	case TbPengeluaran:
		{
			return newWtPengeluaranDbo(), nil
		}
	default:
		return nil, fmt.Errorf("unknown table name")
	}
}

package sqlite

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"speggo/domain"
)

type SfgService struct {
	dbPath    string
	db        *sql.DB
	baseQuery string
}

func NewSfgService(path string) SfgService {
	temp := SfgService{}
	temp.dbPath = path
	temp.db, _ = sql.Open("sqlite3", path)
	temp.baseQuery = "SELECT id, name, measured_time, wavenumbers, sfg, ir, vis FROM sfg"
	return temp
}

func ScanSfg(rows *sql.Rows) []RawSFG {
	out := make([]RawSFG, 0)
	for rows.Next() {
		temp := RawSFG{}
		err := rows.Scan(&temp.Id, &temp.Name, &temp.MeasuredTime, &temp.Wavenumbers, &temp.Sfg, &temp.Vis, &temp.Ir)
		if err != nil {
			log.Println(err)
		} else {
			out = append(out, temp)
		}

	}
	return out
}

func (svc *SfgService) GetSfgById(id string) (domain.SFG, error) {
	query := fmt.Sprintf("%s WHERE id='%s'", svc.baseQuery, id)
	log.Print(query)
	rows, err := svc.db.Query(query)
	if err != nil {
		return domain.SFG{}, err
	}
	test, err := ScanSfg(rows)[0].ToSFG()

	if err != nil {
		return domain.SFG{}, err
	}
	return test, nil
}

func (svc *SfgService) ListSfgSpectra() ([]domain.SFG, error) {
	rows, err := svc.db.Query(svc.baseQuery + " LIMIT 500")
	if err != nil {
		log.Println(err)
		return []domain.SFG{}, err
	}
	test := ScanSfg(rows)

	out := make([]domain.SFG, 0)
	for _, t := range test {
		sfg, err := t.ToSFG()
		if err != nil {
			log.Println(err)
		} else {
			out = append(out, sfg)
		}

	}
	return out, nil
}

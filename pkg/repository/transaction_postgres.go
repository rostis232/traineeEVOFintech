package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rostis232/traineeEVOFintech"
	"strings"
	"time"
)

type TransactionPostgres struct {
	db *sqlx.DB
}

func NewTransactionPostgres(db *sqlx.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}

// InsertToDB inserts data from []traineeEVOFintech.Transaction to DB
func (i *TransactionPostgres) InsertToDB(transactions []traineeEVOFintech.Transaction) error {
	query := "INSERT INTO transaction (transaction_id, request_id, terminal_id, partner_object_id, amount_total, " +
		"amount_original, commission_ps, commission_client, commission_provider, date_input, date_post, status, " +
		"payment_type, payment_number, service_id, service, payee_id, payee_name, payee_bank_mfo, payee_bank_account, " +
		"payment_narrative) VALUES "

	for i, t := range transactions {
		query += fmt.Sprintf("(%d, %d, %d, %d, %.2f, %.2f, %.2f, %.2f, %.2f, '%s', '%s', '%s', '%s', '%s', %d, '%s', %d, '%s', %d, '%s', '%s')",
			t.TransactionId, t.RequestId, t.TerminalId, t.PartnerObjectId, t.AmountTotal,
			t.AmountOriginal, t.CommissionPS, t.CommissionClient, t.CommissionProvider,
			t.DateInput, t.DatePost,
			t.Status, t.PaymentType, t.PaymentNumber, t.ServiceId, t.Service, t.PayeeId, t.PayeeName, t.PayeeBankMfo,
			t.PayeeBankAccount, t.PaymentNarrative)
		if i == len(transactions)-1 {
			query += ";"
		} else {
			query += ", "
		}
	}

	row := i.db.QueryRow(query)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// GetJSON collect data from DB uses keys stored in map[string]string and return []traineeEVOFintech.Transaction
func (i *TransactionPostgres) GetJSON(m map[string]string) ([]traineeEVOFintech.Transaction, error) {
	var transactions []traineeEVOFintech.Transaction
	query := createQuery(m)

	if err := i.db.Select(&transactions, query); err != nil {
		return nil, err
	}

	for i, _ := range transactions {
		transactions[i].DBTimeToJSON()
	}

	return transactions, nil
}

// GetJSON collect data from DB uses keys stored in map[string]string and return []traineeEVOFintech.Transaction
func (i *TransactionPostgres) GetCSV(m map[string]string) ([]traineeEVOFintech.Transaction, error) {
	var transactions []traineeEVOFintech.Transaction
	query := createQuery(m)

	if err := i.db.Select(&transactions, query); err != nil {
		return nil, err
	}

	for i, _ := range transactions {
		transactions[i].DBTimeToJSON()
	}

	return transactions, nil
}

// GetJSON collect data from DB uses keys stored in map[string]string and return []traineeEVOFintech.Transaction
func (i *TransactionPostgres) GetCSVFile(m map[string]string) ([]traineeEVOFintech.Transaction, error) {
	var transactions []traineeEVOFintech.Transaction
	query := createQuery(m)

	if err := i.db.Select(&transactions, query); err != nil {
		return nil, err
	}

	for i, _ := range transactions {
		transactions[i].DBTimeToJSON()
	}

	return transactions, nil

}

// createQuery generate SQL request
func createQuery(m map[string]string) string {
	var query string
	if len(m) == 0 {
		query = fmt.Sprintf("SELECT * FROM transaction;")
	} else {
		query = fmt.Sprintf("SELECT * FROM transaction WHERE ")

		v, ok := m["transactionId"]
		if ok == true && v != "" {
			query += fmt.Sprintf("transaction_id = %s", v)
		}

		v, ok = m["terminalId"]
		if ok == true && v != "" {
			if !strings.HasSuffix(query, "WHERE ") {
				query += " AND "
			}
			arguments := strings.Split(v, ",")
			if len(arguments) == 1 {
				query += fmt.Sprintf("terminal_id = %s", v)
			} else {
				query += "terminal_id IN ("
				for i, a := range arguments {
					query += fmt.Sprintf("'%s'", a)
					if i != len(arguments)-1 {
						query += ","
					} else {
						query += ")"
					}
				}
			}

		}

		v, ok = m["status"]
		if ok == true && v != "" {
			if !strings.HasSuffix(query, "WHERE ") {
				query += " AND "
			}
			query += fmt.Sprintf("status = '%s'", v)
		}

		v, ok = m["paymentType"]
		if ok == true && v != "" {
			if !strings.HasSuffix(query, "WHERE ") {
				query += " AND "
			}
			query += fmt.Sprintf("payment_type = '%s'", v)
		}

		v, ok = m["datePostFrom"]
		if ok == true && v != "" {
			if !strings.HasSuffix(query, "WHERE ") {
				query += " AND "
			}
			date, _ := time.Parse("2006-01-02", v)
			query += fmt.Sprintf("date_post >= '%s'", date.Format("2006-01-02 15:04:05 -07:00"))
		}

		v, ok = m["datePostTo"]
		if ok == true && v != "" {
			if !strings.HasSuffix(query, "WHERE ") {
				query += " AND "
			}
			date, _ := time.Parse("2006-01-02", v)
			date = date.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
			query += fmt.Sprintf("date_post <= '%s'", date.Format("2006-01-02 15:04:05 -07:00"))
		}

		v, ok = m["paymentNarrative"]
		if ok == true && v != "" {
			if !strings.HasSuffix(query, "WHERE ") {
				query += " AND "
			}
			v = strings.TrimPrefix(v, "'")
			v = strings.TrimSuffix(v, "'")
			query += fmt.Sprintf("payment_narrative LIKE '%%%s%%'", v)
		}

		query += ";"
	}
	return query
}

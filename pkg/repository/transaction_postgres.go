package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/rostis232/traineeEVOFintech"
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
			dateTimeToTimestampTZ(t.DateInput), dateTimeToTimestampTZ(t.DatePost),
			t.Status, t.PaymentType, t.PaymentNumber, t.ServiceId, t.Service, t.PayeeId, t.PayeeName, t.PayeeBankMfo,
			t.PayeeBankAccount, t.PaymentNarrative)
		if i == len(transactions)-1 {
			query += ";"
		} else {
			query += ", "
		}
	}
	fmt.Println(query)

	row := i.db.QueryRow(query)
	if err := row.Err(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (i *TransactionPostgres) GetJSON(m map[string]string) ([]traineeEVOFintech.Transaction, error) {
	var transactions []traineeEVOFintech.Transaction
	query := fmt.Sprintf("SELECT * FROM transaction WHERE ")
	v, ok := m["transactionId"]
	if ok == true {
		query += fmt.Sprintf("transaction_id = %s;", v)
	}
	fmt.Println(query)
	if err := i.db.Select(&transactions, query); err != nil {
		return nil, err
	}

	return transactions, nil
}

// dateTimeToTimestampTZ converts date and time data from custom DateTime type to string
func dateTimeToTimestampTZ(dt traineeEVOFintech.DateTime) string {
	return dt.Format("2006-01-02 15:04:05 -07:00")
}

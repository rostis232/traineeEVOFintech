package traineeEVOFintech

import "time"

type DateTime struct {
	time.Time
}

func (date *DateTime) UnmarshalCSV(csv string) (err error) {
	date.Time, err = time.Parse("2006-01-02 15:04:05", csv)
	return err
}

// TODO: What type of DateInput & DatePost?
// TODO: What type of moneytery fields? float32 or flot64?
// TODO: What type of Service & PaymentNarrative? string or []rune, []byte?
// TODO: What type of PayeeBankMfo? uint or uint8?
type Transaction struct {
	TransactionId      uint     `db:"transaction_id" csv:"TransactionId"`
	RequestId          uint     `db:"request_id" csv:"RequestId"`
	TerminalId         uint     `db:"terminal_id" csv:"TerminalId"`
	PartnerObjectId    uint     `db:"partner_object_id" csv:"PartnerObjectId"`
	AmountTotal        float32  `db:"amount_total" csv:"AmountTotal"`
	AmountOriginal     float32  `db:"amount_original" csv:"AmountOriginal"`
	CommissionPS       float32  `db:"commission_ps" csv:"CommissionPS"`
	CommissionClient   float32  `db:"commission_client" csv:"CommissionClient"`
	CommissionProvider float32  `db:"commission_provider" csv:"CommissionProvider"`
	DateInput          DateTime `db:"date_input" csv:"DateInput"`
	DatePost           DateTime `db:"date_post" csv:"DatePost"`
	Status             string   `db:"status" csv:"Status"`
	PaymentType        string   `db:"payment_type" csv:"PaymentType"`
	PaymentNumber      string   `db:"payment_number" csv:"PaymentNumber"`
	ServiceId          uint     `db:"service_id" csv:"ServiceId"`
	Service            string   `db:"service" csv:"Service"`
	PayeeId            uint     `db:"payee_id" csv:"PayeeId"`
	PayeeName          string   `db:"payee_name" csv:"PayeeName"`
	PayeeBankMfo       uint     `db:"payee_bank_mfo" csv:"PayeeBankMfo"`
	PayeeBankAccount   string   `db:"payee_bank_account" csv:"PayeeBankAccount"`
	PaymentNarrative   string   `db:"payment_narrative" csv:"PaymentNarrative"`
}

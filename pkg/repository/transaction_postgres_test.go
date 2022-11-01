package repository

import (
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestCreateQuery(t *testing.T) {
	var tests = []map[string]string{}
	tests = append(tests, map[string]string{
		"transactionId":    "1",
		"terminalId":       "324",
		"status":           "accepted",
		"paymentType":      "cash",
		"datePostFrom":     "2022-08-17",
		"datePostTo":       "2022-08-17",
		"paymentNarrative": "А11/27122",
	})
	tests = append(tests, map[string]string{
		"transactionId": "1",
		"terminalId":    "324",
		"status":        "accepted",
		"paymentType":   "cash",
	})
	tests = append(tests, map[string]string{
		"datePostFrom":     "2022-08-17",
		"datePostTo":       "2022-08-17",
		"paymentNarrative": "А11/27122",
	})
	tests = append(tests, map[string]string{
		"paymentNarrative": "А11/27122",
	})
	tests = append(tests, map[string]string{
		"datePostFrom": "2022-08-17",
	})
	tests = append(tests, map[string]string{
		"datePostTo": "2022-08-17",
	})
	tests = append(tests, map[string]string{
		"terminalId": "324,325",
	})
	tests = append(tests, map[string]string{
		"transactionId":    "1",
		"terminalId":       "324,325,326",
		"status":           "accepted",
		"paymentType":      "cash",
		"datePostFrom":     "2022-08-17",
		"datePostTo":       "2022-08-17",
		"paymentNarrative": "від 28.12.2022 №123",
	})

	for _, m := range tests {
		query := createQuery(m)
		fmt.Println(query)
		if len(m) != 0 && !strings.HasPrefix(query, "SELECT * FROM transaction WHERE") {
			t.Error("Query prefix does not contain 'SELECT * FROM transaction WHERE'")
		}
		if len(m) == 0 && query != "SELECT * FROM transaction;" {
			t.Errorf("Expected 'SELECT * FROM transaction;'. Got '%s'", query)
		}
		if !strings.HasSuffix(query, ";") {
			t.Error("Query suffix does not contain ';'")
		}

		v, ok := m["transactionId"]
		if ok == true && v != "" {
			if !strings.Contains(query, fmt.Sprintf("transaction_id = %s", v)) {
				t.Errorf("Query does not contain '%s'", fmt.Sprintf("transaction_id = %s", v))
			}
		}

		//v, ok = m["terminalId"]
		//if ok == true && v != "" {
		//	if !strings.HasSuffix(query, "WHERE ") {
		//		query += " AND "
		//	}
		//	arguments := strings.Split(v, ",")
		//	if len(arguments) == 1 {
		//		query += fmt.Sprintf("terminal_id = %s", v)
		//	} else {
		//		query += "terminal_id IN ("
		//		for i, a := range arguments {
		//			query += fmt.Sprintf("'%s'", a)
		//			if i != len(arguments)-1 {
		//				query += ","
		//			} else {
		//				query += ")"
		//			}
		//		}
		//	}
		//
		//}

		v, ok = m["status"]
		if ok == true && v != "" {
			if !strings.Contains(query, fmt.Sprintf("status = '%s'", v)) {
				t.Errorf("Query does not contain '%s'", fmt.Sprintf("status = '%s'", v))
			}
		}

		v, ok = m["paymentType"]
		if ok == true && v != "" {
			if !strings.Contains(query, fmt.Sprintf("payment_type = '%s'", v)) {
				t.Errorf("Query does not contain '%s'", fmt.Sprintf("payment_type = '%s'", v))
			}
		}

		v, ok = m["datePostFrom"]
		if ok == true && v != "" {
			date, _ := time.Parse("2006-01-02", v)
			nv := date.Format("2006-01-02 15:04:05 -07:00")
			if !strings.Contains(query, fmt.Sprintf("date_post >= '%s'", nv)) {
				t.Errorf("Query does not contain '%s'", fmt.Sprintf("date_post >= '%s'", nv))
			}
		}

		v, ok = m["datePostTo"]
		if ok == true && v != "" {
			date, _ := time.Parse("2006-01-02", v)
			date = date.Add(23 * time.Hour).Add(59 * time.Minute).Add(59 * time.Second)
			nv := date.Format("2006-01-02 15:04:05 -07:00")
			if !strings.Contains(query, fmt.Sprintf("date_post <= '%s'", nv)) {
				t.Errorf("Query does not contain '%s'", fmt.Sprintf("date_post <= '%s'", nv))
			}
		}

		v, ok = m["paymentNarrative"]
		if ok == true && v != "" {

			v = strings.TrimPrefix(v, "'")
			v = strings.TrimSuffix(v, "'")
			if !strings.Contains(query, fmt.Sprintf("payment_narrative LIKE '%%%s%%'", v)) {
				t.Errorf("Query does not contain '%%%s%%'", fmt.Sprintf("payment_narrative LIKE '%%%s%%'", v))
			}
		}
	}

}

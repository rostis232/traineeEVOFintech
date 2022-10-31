package main

import "testing"

func TestRun(t *testing.T) {
	err := run()
	if err != nil {
		t.Error("Faild run(). Check connection to DB and configs")
	}
}

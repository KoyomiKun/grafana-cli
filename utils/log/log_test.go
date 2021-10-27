package log

import "testing"

func TestLog(t *testing.T) {
	Infof("Info %s", "Hello")
	Warnf("Warn %s", "Hello")
	Errorf("Error %s", "Hello")
	//Fatalf("Fatalf %s", "Hello")
}

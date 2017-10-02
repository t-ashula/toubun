package core

import (
	"os"
	"testing"
)

func TestEnvMap(t *testing.T) {
	em := EnvMap("")
	if len(em) == 0 {
		t.Fatal("something wrong")
	}

	os.Clearenv()
	em = EnvMap("")
	if len(em) != 0 {
		t.Fatalf("not cleard; %v", em)
	}

	os.Setenv("foo", "val")
	os.Setenv("hoge", "huga")
	os.Setenv("PREF_BAR", "BAZ")
	os.Setenv("PREF_PIYO", "PUNI")
	em = EnvMap("")
	if len(em) != 4 ||
		em["foo"] != "val" ||
		em["hoge"] != "huga" ||
		em["PREF_BAR"] != "BAZ" ||
		em["PREF_PIYO"] != "PUNI" {
		t.Fatalf("get all failed; %v", em)
	}

	em = EnvMap("PREF_")
	if len(em) != 2 ||
		em["PREF_BAR"] != "BAZ" ||
		em["PREF_PIYO"] != "PUNI" {
		t.Fatalf("with prefix failed.%v", em)
	}
}

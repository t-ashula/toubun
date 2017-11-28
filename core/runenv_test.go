package core

import (
	"os"
	"testing"
)

func TestChangeWorkDir(t *testing.T) {
	re := NewRunEnv()

	nd := os.TempDir() // just, test
	err := re.ChangeWorkDir(nd, false)
	if err != nil {
		t.Fatalf("Chagne dir to %s failed; %v", nd, err)
	}

	cd, err := os.Getwd()
	if cd != nd || err != nil {
		t.Fatalf("something wrong, dir:%s, err:%s", cd, err)
	}
}

func TestCurrentWorkDir(t *testing.T) {
	re := NewRunEnv()

	cd := re.CurrentWorkDir()
	if cd != "" {
		t.Fatalf("initial current dir should empty, but '%s'", cd)
	}

	nd := os.TempDir() // just, test
	err := re.ChangeWorkDir(nd, false)
	if err != nil {
		t.Fatalf("Chagne dir to %s failed; %v", nd, err)
	}

	cd = re.CurrentWorkDir()
	if cd != nd {
		t.Fatalf("something wrong cd is not %s, %s", nd, cd)
	}
}

func TestCleanup(t *testing.T) {
	re := NewRunEnv()

	nd := os.TempDir() // just, test
	err := re.ChangeWorkDir(nd, false)
	if err != nil {
		t.Fatalf("Chagne dir to %s failed; %v", nd, err)
	}

	cd := re.CurrentWorkDir()
	if cd != nd {
		t.Fatalf("something wrong cd is not %s, %s", nd, cd)
	}

	re.Cleanup()

	cd = re.CurrentWorkDir()
	if cd != "" {
		t.Fatalf("cleanup failed, %s", cd)
	}
}

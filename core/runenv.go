package core

import (
	"os"
)

// RunEnv represents running environment
type RunEnv interface {
	CurrentWorkDir() string
	ChangeWorkDir(string, bool) error
	EnvVars() map[string]string
	Cleanup()
}

// NewRunEnv create new running environment
func NewRunEnv() RunEnv {
	re := &runEnv{
		workdirs: make([]workdirInfo, 5),
	}
	return re
}

type runEnv struct {
	workdirs []workdirInfo
}

type workdirInfo struct {
	dir    string
	remove bool
}

func (re *runEnv) CurrentWorkDir() string {
	l := len(re.workdirs)
	if l == 0 {
		return ""
	}

	return re.workdirs[l-1].dir
}

func (re *runEnv) ChangeWorkDir(dir string, remove bool) error {
	err := os.Chdir(dir)
	if err != nil {
		return err
	}
	re.workdirs = append(re.workdirs, workdirInfo{dir, remove})
	return nil
}

func (re *runEnv) EnvVars() map[string]string {
	return EnvMap("")
}

func (re *runEnv) Cleanup() {
	l := len(re.workdirs)
	for i := range re.workdirs {
		w := re.workdirs[l-1-i]
		if w.remove {
			os.RemoveAll(w.dir)
		}
	}
	re.workdirs = make([]workdirInfo, 5)
}

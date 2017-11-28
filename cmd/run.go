package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	k "github.com/t-ashula/toubun/core"
	"github.com/t-ashula/toubun/runner"
)

var runCommand = &cobra.Command{
	Use:   "run",
	Short: "run toubun",
	Run:   runCore,
}

var configPath string

func init() {
	runCommand.Flags().StringVarP(&configPath, "config", "c", ".toubun.yml", "path to config file")
	RootCommand.AddCommand(runCommand)
}

func runCore(cmd *cobra.Command, args []string) {
	cfg, err := loadConfig()
	if err != nil {
		log.Printf("load config failed; %s\n", err)
		return
	}
	log.Printf("config:%+v\n", cfg)

	re, err := prepareRunEnv()
	if err != nil {
		log.Printf("prepare env failed; %s\n", err)
		return
	}
	log.Printf("RunEnv Prepared\n")

	rs, err := makeRunner(cfg)
	if err != nil {
		log.Printf("make runner faild; %s\n", err)
		return
	}
	log.Printf("Runner Created %v\n", rs)

	err = validateConfig(rs)
	if err != nil {
		log.Printf("validate config faild; %s\n", err)
		return
	}
	log.Printf("Config validated\n")

	err = run(re, rs)
	if err != nil {
		log.Printf("run faild; %s\n", err)
		return
	}

	cleanRunEnv(re)
}

func loadConfig() (*k.AppConfig, error) {
	cfg, err := k.LoadConfig(configPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Printf("guess mode is not support yet\n")
		}
		return nil, err
	}
	return cfg, nil
}

func prepareRunEnv() (k.RunEnv, error) {
	re := k.NewRunEnv()

	td := os.TempDir()
	err := re.ChangeWorkDir(td, false)
	if err != nil {
		return nil, err
	}

	return re, nil
}

func makeRunner(cfg *k.AppConfig) ([]runner.Runner, error) {
	f := runner.NewFetcher(&cfg.Fetch)
	if f == nil {
		log.Printf("no fetcher <%s> matched.\n", cfg.Fetch.Module)
		return nil, fmt.Errorf("no fetcher <%s> matched", cfg.Fetch.Module)
	}

	u := runner.NewUpdater(&cfg.Update)
	if u == nil {
		log.Printf("no updater <%s> matched.\n", cfg.Update.Module)
		return nil, fmt.Errorf("no fetcher <%s> matched", cfg.Update.Module)
	}

	p := runner.NewPublisher(&cfg.Publish)
	if p == nil {
		log.Printf("no publisher <%s> matched.\n", cfg.Publish.Module)
		return nil, fmt.Errorf("no fetcher <%s> matched", cfg.Publish.Module)
	}
	return []runner.Runner{f, u, p}, nil
}

func validateConfig(runs []runner.Runner) error {
	for _, r := range runs {
		if err := r.ValidateConfig(); err != nil {
			log.Printf("%s validation failed %v", r.Name(), err)
			return err
		}
	}
	return nil
}

func run(re k.RunEnv, rs []runner.Runner) error {
	for _, r := range rs {
		log.Printf("%s:run\n", r.Name())
		if err := r.Run(re); err != nil {
			return err
		}
	}
	return nil
}

func cleanRunEnv(re k.RunEnv) {
	re.Cleanup()
}

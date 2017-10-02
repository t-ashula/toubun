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
	Run: func(cmd *cobra.Command, args []string) {
		cfg, err := loadConfig()
		if err != nil {
			log.Printf("load config failed; '%s'\n", err)
			return
		}
		log.Printf("config:%+v\n", cfg)

		re, err := prepareRunEnv()
		if err != nil {
			fmt.Printf("prepare env failed; '%s'\n", err)
			return
		}

		rs, err := makeRunner(cfg)
		if err != nil {
			fmt.Printf("make runner faild; '%s'\n", err)
			return
		}

		err = validateConfig(rs)
		if err != nil {
			fmt.Printf("validate config faild; '%s'\n", err)
			return
		}

		err = run(re, rs)
		if err != nil {
			fmt.Printf("run faild; '%s'\n", err)
			return
		}

		cleanRunEnv(re)
	},
}

var configPath string

func init() {
	runCommand.Flags().StringVarP(&configPath, "config", "c", ".toubun.yml", "path to config file")
	RootCommand.AddCommand(runCommand)
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
	err := re.ChangeWorkDir(td, true)
	if err != nil {
		return nil, err
	}

	return re, nil
}

func makeRunner(cfg *k.AppConfig) ([]runner.Runner, error) {
	f := runner.NewFetcher(&cfg.Fetch)
	if f != nil {
		return nil, fmt.Errorf("no fetcher matched.%s\n", cfg.Fetch.Module)
	}

	u := runner.NewUpdater(&cfg.Update)
	if u != nil {
		fmt.Printf("no fetcher matched.%s\n", cfg.Fetch.Module)
		return nil, fmt.Errorf("no fetcher matched.%s\n", cfg.Fetch.Module)
	}

	p := runner.NewPublisher(&cfg.Update)
	if f != nil {
		fmt.Printf("no fetcher matched.%s\n", cfg.Fetch.Module)
		return nil, fmt.Errorf("no fetcher matched.%s\n", cfg.Fetch.Module)
	}
	return []runner.Runner{f, u, p}, nil
}

func validateConfig(rs []runner.Runner) error {
	for _, r := range rs {
		if err := r.ValidateConfig(); err != nil {
			return err
		}
	}
	return nil
}

func run(re k.RunEnv, rs []runner.Runner) error {
	for _, r := range rs {
		if err := r.Run(re); err != nil {
			return err
		}
	}
	return nil
}

func cleanRunEnv(re k.RunEnv) {
	re.Cleanup()
}

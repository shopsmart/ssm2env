package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/shopsmart/ssm2env"
	"github.com/shopsmart/ssm2env/pkg/service"
)

// New creates a new cobra command for the given version
func New(version string, svc service.Service) *cobra.Command {
	config := viper.New()
	config.SetEnvPrefix("ssm2env")

	rootCmd := &cobra.Command{
		Use:   "ssm2env",
		Short: "Pulls SSM parameters into env format",
		Long: `SSM2Env pulls parameters from AWS SSM Param Store
		and puts them in env format

		ssm2env /my/prefix
		`,
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var v bool
			v = config.GetBool("verbose")
			if v {
				log.SetLevel(log.DebugLevel)
			}

			v = config.GetBool("version")
			if v {
				fmt.Println(version)
				return
			}

			multilineSupport := config.GetBool("multiline")
			recursive := config.GetBool("recursive")
			export := config.GetBool("export")

			path := config.GetString("path")
			if path == "" {
				validArgs := []string{}
				for _, arg := range args {
					if arg != "" {
						validArgs = append(validArgs, arg)
					}
				}

				if len(validArgs) < 1 {
					if err := cmd.Help(); err != nil {
						log.Fatal(err)
					}
					return
				}

				path = validArgs[0]
			}

			cfg := ssm2env.Config{
				SearchPath:       path,
				Recursive:        recursive,
				MultilineSupport: multilineSupport,
				Export:           export,
			}

			err := ssm2env.Collect(svc, os.Stdout, &cfg)
			if err != nil {
				log.Fatal(err)
				return
			}
		},
	}

	rootCmd.PersistentFlags().Bool("multiline", true, "enables multiline support; to enable, set --multiline=false")
	rootCmd.PersistentFlags().Bool("recursive", false, "searches the path recursively")
	rootCmd.PersistentFlags().Bool("export", false, "adds export before each variable")
	rootCmd.PersistentFlags().Bool("verbose", false, "enables verbose output")
	rootCmd.PersistentFlags().Bool("version", false, "prints the version and exits")

	_ = config.BindPFlag("multiline", rootCmd.Flags().Lookup("multiline"))
	_ = config.BindPFlag("recursive", rootCmd.Flags().Lookup("recursive"))
	_ = config.BindPFlag("export", rootCmd.Flags().Lookup("export"))
	_ = config.BindPFlag("verbose", rootCmd.Flags().Lookup("verbose"))

	_ = config.BindEnv("path")
	_ = config.BindEnv("multiline")
	_ = config.BindEnv("recursive")
	_ = config.BindEnv("export")
	_ = config.BindEnv("verbose")

	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string, svc service.Service) {
	cmd := New(version, svc)

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

package cmd

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/shopsmart/ssm2env"
	"github.com/shopsmart/ssm2env/pkg/service"
)

// New creates a new cobra command for the given version
func New(version string, svc *service.Service) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "ssm2env",
		Short: "Pulls SSM paramters into env format",
		Long: `SSM2Env pulls paramters from AWS SSM Param Store
		and puts them in env format

		ssm2env /my/prefix
		`,
		Args: cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var v bool
			v, _ = cmd.Flags().GetBool("verbose")
			if v {
				log.SetLevel(log.DebugLevel)
			}

			v, _ = cmd.Flags().GetBool("version")
			if v {
				fmt.Println(version)
				return
			}

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

			err := ssm2env.Collect(svc, os.Stdout, validArgs[0])
			if err != nil {
				log.Fatal(err)
				return
			}
		},
	}

	rootCmd.PersistentFlags().Bool("verbose", false, "enables verbose output")
	rootCmd.PersistentFlags().Bool("version", false, "prints the version and exits")

	return rootCmd
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute(version string, svc *service.Service) {
	cmd := New(version, svc)

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

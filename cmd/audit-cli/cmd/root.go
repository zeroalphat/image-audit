/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"flag"
	"fmt"
	"os"

	_ "crypto/sha256"
	_ "crypto/sha512"

	digest "github.com/opencontainers/go-digest"
	"github.com/spf13/cobra"
)

type ImageSpec struct {
	Name      string
	Digest    digest.Digest
	MediaType string
}

var rootCmd = &cobra.Command{
	Use:                "audit-cli",
	FParseErrWhitelist: cobra.FParseErrWhitelist{UnknownFlags: true},
	Short:              "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		cmd.SilenceUsage = true

		image := ImageSpec{}
		flag.StringVar(&image.Name, "name", "", "container image name")
		d := flag.String("digest", "", "container iamge digest")
		flag.StringVar(&image.MediaType, "stdin-media-type", "", "container image media type")
		flag.Parse()
		image.Digest = digest.FromString(*d)
		fmt.Printf("%+v", image)

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

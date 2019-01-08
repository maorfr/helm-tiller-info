package main

import (
	"fmt"
	"os"

	utils "github.com/maorfr/helm-plugin-utils/pkg"

	"github.com/spf13/cobra"
)

var (
	tillerNamespace string
)

func main() {
	cmd := &cobra.Command{
		Use:   "tiller-info [flags]",
		Short: "print information about Tiller",
		RunE:  run,
	}

	f := cmd.Flags()
	f.StringVar(&tillerNamespace, "tiller-namespace", "kube-system", "namespace of Tiller")

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) error {
	storage := utils.GetTillerStorage(tillerNamespace)
	fmt.Println("Tiller storage: " + storage)
	fmt.Println("Tiller default deployment namespace: " + "default") // possible future development as part of helm/helm#5111

	return nil
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Mostra a versão da aplicação",
	Long: `Mostra a versão atual da aplicação bast.

		Exemplos:
		  bast version        # Mostra a versão
		  bast version --help # Mostra ajuda deste comando`,
	Run: func(cmd *cobra.Command, args []string) {
		verbosePrint(cmd, "Obtendo informações de versão...\n")
		fmt.Println("bast v1.0.0")
		fmt.Println("Construído com Go e Cobra")
		verbosePrint(cmd, "Informações de versão exibidas.\n")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	name     string
	greeting string
)

var greetCmd = &cobra.Command{
	Use:   "greet",
	Short: "Cumprimenta uma pessoa",
	Long: `Cumprimenta uma pessoa pelo nome.
Você pode personalizar a saudação usando a flag --greeting.

Exemplos:
  bast greet --name "Maria"              # Cumprimenta Maria
  bast greet -n "Pedro" -g "Bem-vindo"  # Cumprimenta Pedro com saudação personalizada
  bast greet                             # Cumprimenta "Mundo" (padrão)
  bast greet --help                      # Mostra ajuda deste comando`,
	Run: func(cmd *cobra.Command, args []string) {
		verbosePrint(cmd, "Processando comando greet...")
		verbosePrint(cmd, "Nome fornecido: '%s'", name)
		verbosePrint(cmd, "Saudação fornecida: '%s'", greeting)

		if name == "" {
			name = "Mundo"
			verbosePrint(cmd, "Usando nome padrão: 'Mundo'")
		}
		if greeting == "" {
			greeting = "Olá"
			verbosePrint(cmd, "Usando saudação padrão: 'Olá'")
		}

		verbosePrint(cmd, "Gerando mensagem de cumprimento...")
		fmt.Printf("%s, %s!\n", greeting, name)
		verbosePrint(cmd, "Comando greet executado com sucesso.")
	},
}

func init() {
	rootCmd.AddCommand(greetCmd)

	greetCmd.Flags().StringVarP(&name, "name", "n", "", "Nome da pessoa a ser cumprimentada")
	greetCmd.Flags().StringVarP(&greeting, "greeting", "g", "", "Saudação personalizada")
}

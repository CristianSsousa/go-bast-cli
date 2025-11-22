package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// isVerbose verifica se o modo verbose está ativado
func isVerbose(cmd *cobra.Command) bool {
	verbose, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		return false
	}
	return verbose
}

// verbosePrint imprime mensagens apenas se o modo verbose estiver ativo
func verbosePrint(cmd *cobra.Command, format string, args ...interface{}) {
	if isVerbose(cmd) {
		fmt.Printf("[VERBOSE] "+format, args...)
	}
}

var rootCmd = &cobra.Command{
	Use:   "bast",
	Short: "Uma CLI moderna construída com Go e Cobra",
	Long: `bast é uma aplicação CLI moderna construída com Go e Cobra.
Ela fornece uma interface de linha de comando poderosa e extensível.

Exemplos:
  bast version                    # Mostra a versão
  bast greet --name "João"        # Cumprimenta alguém
  bast serve --port 3000          # Inicia servidor na porta 3000
  bast install git                # Instala o Git
  bast info                       # Mostra informações do sistema
  bast port 8080                  # Verifica se porta está em uso
  bast config list                # Lista configurações
  bast --help                     # Mostra esta mensagem de ajuda`,
	Run: func(cmd *cobra.Command, args []string) {
		verbosePrint(cmd, "Executando comando raiz...\n")
		fmt.Println("Bem-vindo ao bast!")
		fmt.Println("Use 'bast --help' para ver os comandos disponíveis.")
		verbosePrint(cmd, "Modo verbose está ativo.\n")
	},
}

// Execute adiciona todos os comandos filhos ao comando raiz e define flags apropriadas.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Flags globais podem ser adicionadas aqui
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Modo verboso")
}

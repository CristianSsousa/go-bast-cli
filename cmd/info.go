package cmd

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Mostra informações do sistema",
	Long: `Mostra informações detalhadas sobre o sistema operacional,
arquitetura, versão do Go e outras informações úteis.

Exemplos:
  bast info              # Mostra todas as informações
  bast info --os         # Mostra apenas informações do OS
  bast info --go         # Mostra apenas informações do Go
  bast info --help       # Mostra ajuda deste comando`,
	Run: func(cmd *cobra.Command, args []string) {
		showOS, _ := cmd.Flags().GetBool("os")
		showGo, _ := cmd.Flags().GetBool("go")
		showEnv, _ := cmd.Flags().GetBool("env")

		// Se nenhuma flag específica foi passada, mostra tudo
		if !showOS && !showGo && !showEnv {
			showOS = true
			showGo = true
			showEnv = true
		}

		verbosePrint(cmd, "Coletando informações do sistema...\n")

		if showOS {
			showOSInfo(cmd)
		}

		if showGo {
			showGoInfo(cmd)
		}

		if showEnv {
			showEnvInfo(cmd)
		}
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)

	infoCmd.Flags().Bool("os", false, "Mostra apenas informações do sistema operacional")
	infoCmd.Flags().Bool("go", false, "Mostra apenas informações do Go")
	infoCmd.Flags().Bool("env", false, "Mostra apenas variáveis de ambiente importantes")
}

func showOSInfo(cmd *cobra.Command) {
	fmt.Println("\n📱 Sistema Operacional:")
	fmt.Printf("  OS: %s\n", runtime.GOOS)
	fmt.Printf("  Arquitetura: %s\n", runtime.GOARCH)
	fmt.Printf("  CPUs: %d\n", runtime.NumCPU())

	hostname, err := os.Hostname()
	if err == nil {
		fmt.Printf("  Hostname: %s\n", hostname)
		verbosePrint(cmd, "Hostname obtido com sucesso.\n")
	} else {
		verbosePrint(cmd, "Erro ao obter hostname: %v\n", err)
	}

	wd, err := os.Getwd()
	if err == nil {
		fmt.Printf("  Diretório atual: %s\n", wd)
		verbosePrint(cmd, "Diretório de trabalho: %s\n", wd)
	} else {
		verbosePrint(cmd, "Erro ao obter diretório atual: %v\n", err)
	}
}

func showGoInfo(cmd *cobra.Command) {
	fmt.Println("\n🐹 Go:")
	fmt.Printf("  Versão: %s\n", runtime.Version())
	fmt.Printf("  Compilador: %s\n", runtime.Compiler)
	fmt.Printf("  Goroutines: %d\n", runtime.NumGoroutine())

	memStats := runtime.MemStats{}
	runtime.ReadMemStats(&memStats)
	fmt.Printf("  Memória alocada: %d KB\n", memStats.Alloc/1024)
	fmt.Printf("  Total de alocações: %d\n", memStats.Mallocs)
	verbosePrint(cmd, "Estatísticas de memória coletadas.\n")
}

func showEnvInfo(cmd *cobra.Command) {
	fmt.Println("\n🌍 Variáveis de Ambiente Importantes:")

	envVars := []string{
		"HOME", "USER", "USERNAME",
		"PATH", "SHELL",
		"GOPATH", "GOROOT", "GOBIN",
		"EDITOR", "LANG", "TZ",
	}

	for _, envVar := range envVars {
		value := os.Getenv(envVar)
		if value != "" {
			// Trunca PATH se for muito longo
			if envVar == "PATH" && len(value) > 100 {
				fmt.Printf("  %s: %s...\n", envVar, value[:100])
				verbosePrint(cmd, "PATH truncado para exibição.\n")
			} else {
				fmt.Printf("  %s: %s\n", envVar, value)
			}
		} else {
			verbosePrint(cmd, "Variável %s não definida.\n", envVar)
		}
	}
}

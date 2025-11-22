package cmd

import (
	"fmt"
	"os"

	"github.com/CristianSsousa/go-bast-cli/internal/config"
	"github.com/CristianSsousa/go-bast-cli/internal/constants"
	"github.com/CristianSsousa/go-bast-cli/internal/logger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	appLog  = logger.GetLogger()
)

// isVerbose verifica se o modo verbose está ativado
func isVerbose(cmd *cobra.Command) bool {
	verbose, err := cmd.Flags().GetBool("verbose")
	if err != nil {
		return false
	}
	return verbose
}

// verbosePrint imprime mensagens apenas se o modo verbose estiver ativo usando logger estruturado
func verbosePrint(cmd *cobra.Command, format string, args ...interface{}) {
	if isVerbose(cmd) {
		appLog.Debugf(format, args...)
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
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Inicializar configuração
		if err := config.Init(cfgFile); err != nil {
			appLog.Warnf("Erro ao carregar configuração: %v", err)
		}

		// Inicializar logger com configurações
		cfg := config.Get()
		logger.Init(cfg.Logging.Level, cfg.Logging.Format)

		// Atualizar referência do logger
		appLog = logger.GetLogger()

		// Aplicar verbose se necessário
		if isVerbose(cmd) {
			appLog.SetLevel(logger.GetLogger().Level)
			appLog.Debug("Modo verbose ativado")
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		verbosePrint(cmd, "Executando comando raiz...")
		cfg := config.Get()
		fmt.Printf(constants.WelcomeMessage+"\n", cfg.App.Name)
		fmt.Printf(constants.HelpMessage+"\n", cfg.App.Name)
		verbosePrint(cmd, "Modo verbose está ativo.")
	},
}

// Execute adiciona todos os comandos filhos ao comando raiz e define flags apropriadas.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		appLog.Errorf("Erro ao executar comando: %v", err)
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	// Flags globais
	rootCmd.PersistentFlags().BoolP("verbose", "v", false, "Modo verboso")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Arquivo de configuração (padrão: ~/.bast/config.yaml)")

	// Bind flags ao Viper
	if err := viper.BindPFlag("features.verbose", rootCmd.PersistentFlags().Lookup("verbose")); err != nil {
		appLog.Warnf("Erro ao vincular flag verbose: %v", err)
	}
}

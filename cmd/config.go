package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/CristianSsousa/go-bast-cli/internal/config"
	"github.com/CristianSsousa/go-bast-cli/internal/constants"
	"github.com/CristianSsousa/go-bast-cli/pkg/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Gerencia configurações do bast CLI",
	Long: `Gerencia configurações persistentes do bast CLI.
As configurações são salvas em ~/.bast/config.yaml (Linux/macOS) ou
%USERPROFILE%\.bast\config.yaml (Windows).

Subcomandos:
  list    - Lista todas as configurações
  get     - Obtém valor de uma configuração específica
  set     - Define uma configuração
  reset   - Reseta todas as configurações para valores padrão
  init    - Cria arquivo de configuração inicial`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		// Inicializar configuração sem output do rootCmd
		if err := config.Init(""); err != nil {
			appLog.Warnf("Erro ao carregar configuração: %v", err)
		}
	},
}

var configListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista todas as configurações",
	Long:  `Lista todas as configurações atuais do bast CLI.`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.Get()
		fmt.Println("📋 Configurações do bast CLI:")
		fmt.Println()
		fmt.Printf("  App:\n")
		fmt.Printf("    Nome:        %s\n", cfg.App.Name)
		fmt.Printf("    Versão:      %s\n", cfg.App.Version)
		fmt.Printf("    Descrição:   %s\n", cfg.App.Description)
		fmt.Printf("    Autor:       %s\n", cfg.App.Author)
		fmt.Println()
		fmt.Printf("  Logging:\n")
		fmt.Printf("    Nível:       %s\n", cfg.Logging.Level)
		fmt.Printf("    Formato:     %s\n", cfg.Logging.Format)
		fmt.Println()
		fmt.Printf("  Server:\n")
		fmt.Printf("    Porta padrão: %d\n", cfg.Server.DefaultPort)
		fmt.Printf("    Host padrão:  %s\n", cfg.Server.DefaultHost)
		fmt.Printf("    Timeout:      %d\n", cfg.Server.Timeout)
		fmt.Println()
		fmt.Printf("  Features:\n")
		fmt.Printf("    Auto Update: %v\n", cfg.Features.AutoUpdate)
		fmt.Printf("    Verbose:     %v\n", cfg.Features.Verbose)

		configPath, err := utils.GetConfigPath()
		if err == nil {
			fmt.Println()
			fmt.Printf("📁 Arquivo de configuração: %s\n", configPath)
		}
	},
}

var configGetCmd = &cobra.Command{
	Use:   "get [chave]",
	Short: "Obtém valor de uma configuração específica",
	Long: `Obtém o valor de uma configuração específica.
Use notação com pontos para acessar valores aninhados.

Exemplos:
  bast config get app.name
  bast config get server.default_port
  bast config get logging.level`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		value := viper.Get(key)
		if value == nil {
			fmt.Printf("❌ Chave '%s' não encontrada\n", key)
			os.Exit(1)
		}
		fmt.Printf("%s = %v\n", key, value)
	},
}

var configSetCmd = &cobra.Command{
	Use:   "set [chave] [valor]",
	Short: "Define uma configuração",
	Long: `Define o valor de uma configuração específica.
A configuração será salva no arquivo de configuração.

Exemplos:
  bast config set server.default_port 3000
  bast config set logging.level debug
  bast config set features.auto_update true`,
	Args: cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		key := args[0]
		valueStr := args[1]

		// Tentar converter para tipos apropriados
		var value interface{} = valueStr

		// Tentar converter para int
		if intVal, err := strconv.Atoi(valueStr); err == nil {
			value = intVal
		} else if boolVal, err := strconv.ParseBool(valueStr); err == nil {
			value = boolVal
		}

		config.Set(key, value)

		// Salvar configuração
		if err := config.Save(); err != nil {
			appLog.Errorf(constants.ErrConfigSave+": %v", err)
			fmt.Printf("❌ "+constants.ErrConfigSave+": %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✅ "+constants.SuccessConfigSet+"\n", key, value)
		fmt.Println("💾 " + constants.ConfigSavedMessage)
	},
}

var configResetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reseta todas as configurações para valores padrão",
	Long:  `Reseta todas as configurações para os valores padrão e salva no arquivo de configuração.`,
	Run: func(cmd *cobra.Command, args []string) {
		// Resetar para defaults
		viper.Reset()
		config.Init("")

		// Salvar
		if err := config.Save(); err != nil {
			appLog.Errorf(constants.ErrConfigSave+": %v", err)
			fmt.Printf("❌ "+constants.ErrConfigSave+": %v\n", err)
			os.Exit(1)
		}

		fmt.Println("✅ " + constants.ConfigResetMessage)
		fmt.Println("💾 " + constants.ConfigSavedMessage)
	},
}

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Cria arquivo de configuração inicial",
	Long:  `Cria o arquivo de configuração inicial com valores padrão se não existir.`,
	Run: func(cmd *cobra.Command, args []string) {
		configPath, err := utils.GetConfigPath()
		if err != nil {
			appLog.Errorf("Erro ao obter caminho de configuração: %v", err)
			fmt.Printf("❌ Erro: %v\n", err)
			os.Exit(1)
		}

		// Verificar se já existe
		if utils.FileExists(configPath) {
			fmt.Printf("ℹ️  "+constants.InfoConfigExists+"\n", configPath)
			cfg := config.Get()
			fmt.Printf("   "+constants.InfoConfigResetHint+"\n", cfg.App.Name)
			return
		}

		// Criar diretório se não existir
		if err := utils.EnsureConfigDir(); err != nil {
			appLog.Errorf("Erro ao criar diretório de configuração: %v", err)
			fmt.Printf("❌ Erro ao criar diretório: %v\n", err)
			os.Exit(1)
		}

		// Salvar configuração padrão
		if err := config.Save(); err != nil {
			appLog.Errorf(constants.ErrConfigSave+": %v", err)
			fmt.Printf("❌ "+constants.ErrConfigSave+": %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("✅ "+constants.SuccessConfigCreated+"\n", configPath)
		cfg := config.Get()
		fmt.Printf("💡 "+constants.InfoConfigEditHint+"\n", cfg.App.Name)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.AddCommand(configListCmd)
	configCmd.AddCommand(configGetCmd)
	configCmd.AddCommand(configSetCmd)
	configCmd.AddCommand(configResetCmd)
	configCmd.AddCommand(configInitCmd)
}

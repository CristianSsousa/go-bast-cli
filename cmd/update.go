package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/CristianSsousa/go-bast-cli/internal/config"
	"github.com/spf13/cobra"
)

const (
	githubAPIURL = "https://api.github.com/repos/CristianSsousa/go-bast-cli/releases/latest"
	timeout      = 10 * time.Second
)

type githubRelease struct {
	TagName string `json:"tag_name"`
	Name    string `json:"name"`
	Body    string `json:"body"`
}

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Atualiza o bast CLI para a versão mais recente",
	Long: `Verifica e atualiza o bast CLI para a versão mais recente disponível no GitHub.

O comando verifica a versão atual instalada e compara com a versão mais recente
disponível no GitHub. Se houver uma atualização disponível, o comando oferece
para atualizar usando 'go install'.

Exemplos:
  bast update              # Verifica e atualiza para a versão mais recente
  bast update --check      # Apenas verifica se há atualização disponível
  bast update --help       # Mostra ajuda deste comando`,
	Run: func(cmd *cobra.Command, args []string) {
		checkOnly, _ := cmd.Flags().GetBool("check")

		verbosePrint(cmd, "Iniciando verificação de atualização...")
		cfg := config.Get()

		fmt.Printf("Versão atual: %s\n", cfg.App.Version)
		fmt.Println("Verificando atualizações disponíveis...")

		latestRelease, err := getLatestRelease()
		if err != nil {
			verbosePrint(cmd, "Erro ao obter release: %v", err)
			fmt.Printf("Erro ao verificar atualizações: %v\n", err)
			fmt.Println("\nDica: Verifique sua conexão com a internet.")
			os.Exit(1)
		}

		latestVersion := strings.TrimPrefix(latestRelease.TagName, "v")
		currentVersion := cfg.App.Version

		verbosePrint(cmd, "Versão mais recente encontrada: %s", latestVersion)
		verbosePrint(cmd, "Versão atual: %s", currentVersion)

		if compareVersions(currentVersion, latestVersion) >= 0 {
			fmt.Printf("Você já está usando a versão mais recente (%s)!\n", currentVersion)
			return
		}

		fmt.Printf("Nova versão disponível: %s\n", latestVersion)
		if latestRelease.Name != "" {
			fmt.Printf("Release: %s\n", latestRelease.Name)
		}

		if checkOnly {
			fmt.Println("\nPara atualizar, execute: bast update")
			return
		}

		fmt.Print("\nDeseja atualizar agora? [y/n]: ")
		var response string
		_, err = fmt.Scanln(&response)
		if err != nil {
			verbosePrint(cmd, "Erro ao ler entrada do usuário: %v", err)
			return
		}

		if strings.ToLower(response) != "s" && strings.ToLower(response) != "sim" && strings.ToLower(response) != "y" && strings.ToLower(response) != "yes" {
			fmt.Println("Atualização cancelada.")
			return
		}

		fmt.Println("\nAtualizando bast CLI...")
		verbosePrint(cmd, "Executando: go install github.com/CristianSsousa/go-bast-cli@%s", latestRelease.TagName)

		updateCmd := exec.Command("go", "install", fmt.Sprintf("github.com/CristianSsousa/go-bast-cli@%s", latestRelease.TagName))
		updateCmd.Stdout = os.Stdout
		updateCmd.Stderr = os.Stderr

		if err := updateCmd.Run(); err != nil {
			verbosePrint(cmd, "Erro ao executar go install: %v", err)
			fmt.Printf("\nErro ao atualizar: %v\n", err)
			fmt.Println("\nTente atualizar manualmente:")
			fmt.Printf("  go install github.com/CristianSsousa/go-bast-cli@%s\n", latestRelease.TagName)
			os.Exit(1)
		}

		fmt.Println("\nAtualização concluída com sucesso!")
		fmt.Printf("Versão instalada: %s\n", latestVersion)
		fmt.Println("\nNota: Se o comando 'bast' não refletir a nova versão, certifique-se de que")
		fmt.Println("o diretório $GOPATH/bin ou $GOBIN está no seu PATH.")
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().BoolP("check", "c", false, "Apenas verifica se há atualização disponível, sem atualizar")
}

// getLatestRelease obtém a release mais recente do GitHub
func getLatestRelease() (*githubRelease, error) {
	client := &http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("GET", githubAPIURL, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criar requisição: %w", err)
	}

	req.Header.Set("Accept", "application/vnd.github.v3+json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("erro na resposta da API: status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler resposta: %w", err)
	}

	var release githubRelease
	if err := json.Unmarshal(body, &release); err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}

	return &release, nil
}

// compareVersions compara duas versões no formato semver (X.Y.Z)
// Retorna: -1 se v1 < v2, 0 se v1 == v2, 1 se v1 > v2
func compareVersions(v1, v2 string) int {
	parts1 := strings.Split(v1, ".")
	parts2 := strings.Split(v2, ".")

	maxLen := len(parts1)
	if len(parts2) > maxLen {
		maxLen = len(parts2)
	}

	for i := 0; i < maxLen; i++ {
		var num1, num2 int

		if i < len(parts1) {
			fmt.Sscanf(parts1[i], "%d", &num1)
		}
		if i < len(parts2) {
			fmt.Sscanf(parts2[i], "%d", &num2)
		}

		if num1 < num2 {
			return -1
		}
		if num1 > num2 {
			return 1
		}
	}

	return 0
}

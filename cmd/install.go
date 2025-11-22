package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

var installCmd = &cobra.Command{
	Use:   "install",
	Short: "Instala ferramentas e dependências",
	Long: `Instala ferramentas e dependências necessárias.
Atualmente suporta instalação do Git em diferentes sistemas operacionais.

Exemplos:
  bast install git              # Instala o Git
  bast install --help           # Mostra ajuda deste comando`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Println("Erro: especifique o que deseja instalar.")
			fmt.Println("Uso: bast install <ferramenta>")
			fmt.Println("\nFerramentas disponíveis:")
			fmt.Println("  git - Instala o Git")
			os.Exit(1)
		}

		tool := args[0]
		switch tool {
		case "git":
			installGit(cmd)
		default:
			fmt.Printf("Erro: ferramenta '%s' não é suportada.\n", tool)
			fmt.Println("\nFerramentas disponíveis:")
			fmt.Println("  git - Instala o Git")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(installCmd)
}

func installGit(cmd *cobra.Command) {
	verbosePrint(cmd, "Iniciando processo de instalação do Git...\n")
	fmt.Println("Verificando se o Git já está instalado...")

	// Verifica se o git já está instalado
	if isGitInstalled() {
		verbosePrint(cmd, "Git encontrado no sistema.\n")
		fmt.Println("✓ Git já está instalado!")
		version, err := getGitVersion()
		if version != "" {
			fmt.Printf("  Versão: %s\n", version)
		}
		if err != nil {
			verbosePrint(cmd, "Erro ao obter versão: %v\n", err)
		}
		return
	}

	fmt.Println("Git não encontrado. Iniciando instalação...")
	fmt.Printf("Sistema operacional detectado: %s\n", runtime.GOOS)
	verbosePrint(cmd, "Arquitetura: %s\n", runtime.GOARCH)

	var installCmd *exec.Cmd
	var installMethod string

	switch runtime.GOOS {
	case "windows":
		installCmd, installMethod = getWindowsInstallCommand(cmd)
	case "linux":
		installCmd, installMethod = getLinuxInstallCommand(cmd)
	case "darwin":
		installCmd, installMethod = getDarwinInstallCommand(cmd)
	default:
		fmt.Printf("Erro: sistema operacional '%s' não é suportado para instalação automática.\n", runtime.GOOS)
		fmt.Println("\nPor favor, instale o Git manualmente:")
		fmt.Println("  Windows: https://git-scm.com/download/win")
		fmt.Println("  Linux: Use o gerenciador de pacotes da sua distribuição")
		fmt.Println("  macOS: brew install git")
		os.Exit(1)
	}

	if installCmd == nil {
		fmt.Printf("Erro: não foi possível determinar o método de instalação para %s.\n", runtime.GOOS)
		fmt.Println("\nPor favor, instale o Git manualmente:")
		fmt.Println("  Windows: https://git-scm.com/download/win")
		fmt.Println("  Linux: Use o gerenciador de pacotes da sua distribuição")
		fmt.Println("  macOS: brew install git")
		os.Exit(1)
	}

	fmt.Printf("Método de instalação: %s\n", installMethod)
	verbosePrint(cmd, "Comando completo: %s\n", installCmd.String())
	fmt.Println("Executando comando de instalação...")
	fmt.Println("⚠️  Nota: Você pode precisar inserir sua senha de administrador.")

	// Para apt-get, executamos update primeiro
	if installMethod == "apt-get" && runtime.GOOS == "linux" {
		fmt.Println("Atualizando lista de pacotes...")
		updateCmd := exec.Command("sudo", "apt-get", "update")
		updateCmd.Stdout = os.Stdout
		updateCmd.Stderr = os.Stderr
		verbosePrint(cmd, "Executando: sudo apt-get update\n")
		if err := updateCmd.Run(); err != nil {
			fmt.Printf("Aviso: falha ao atualizar lista de pacotes: %v\n", err)
			verbosePrint(cmd, "Erro detalhado: %v\n", err)
			fmt.Println("Continuando com a instalação...")
		} else {
			verbosePrint(cmd, "Lista de pacotes atualizada com sucesso.\n")
		}
		// Agora criamos o comando de install sem o update
		installCmd = exec.Command("sudo", "apt-get", "install", "-y", "git")
		verbosePrint(cmd, "Comando de instalação: sudo apt-get install -y git\n")
	}

	installCmd.Stdout = os.Stdout
	installCmd.Stderr = os.Stderr
	installCmd.Stdin = os.Stdin

	verbosePrint(cmd, "Iniciando execução do comando de instalação...\n")
	if err := installCmd.Run(); err != nil {
		verbosePrint(cmd, "Erro durante execução: %v\n", err)
		fmt.Printf("\nErro ao executar instalação: %v\n", err)
		fmt.Println("\nTente instalar manualmente:")
		fmt.Println("  Windows: https://git-scm.com/download/win")
		fmt.Println("  Linux: Use o gerenciador de pacotes da sua distribuição")
		fmt.Println("  macOS: brew install git")
		os.Exit(1)
	}

	fmt.Println("\n✓ Instalação concluída!")
	fmt.Println("Verificando instalação...")
	verbosePrint(cmd, "Verificando se Git está acessível no PATH...\n")

	if isGitInstalled() {
		version, err := getGitVersion()
		if version != "" {
			fmt.Printf("✓ Git instalado com sucesso! Versão: %s\n", version)
		} else {
			fmt.Println("✓ Git instalado com sucesso!")
		}
		if err != nil {
			verbosePrint(cmd, "Aviso ao obter versão: %v\n", err)
		}
		verbosePrint(cmd, "Instalação verificada e funcionando corretamente.\n")
	} else {
		fmt.Println("⚠️  Git pode ter sido instalado, mas não foi encontrado no PATH.")
		fmt.Println("   Tente fechar e reabrir o terminal.")
		verbosePrint(cmd, "PATH atual: %s\n", os.Getenv("PATH"))
	}
}

func isGitInstalled() bool {
	cmd := exec.Command("git", "--version")
	err := cmd.Run()
	return err == nil
}

func getGitVersion() (string, error) {
	cmd := exec.Command("git", "--version")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}
	// Remove quebra de linha do final
	version := string(output)
	if len(version) > 0 && version[len(version)-1] == '\n' {
		version = version[:len(version)-1]
	}
	return version, nil
}

func getWindowsInstallCommand(cmd *cobra.Command) (*exec.Cmd, string) {
	// Tenta winget primeiro (Windows 10/11 moderno)
	if isCommandAvailable("winget") {
		verbosePrint(cmd, "Usando winget como gerenciador de pacotes.\n")
		return exec.Command("winget", "install", "--id", "Git.Git", "-e", "--source", "winget"), "winget"
	}

	// Tenta chocolatey
	if isCommandAvailable("choco") {
		verbosePrint(cmd, "Usando Chocolatey como gerenciador de pacotes.\n")
		return exec.Command("choco", "install", "git", "-y"), "chocolatey"
	}

	verbosePrint(cmd, "Nenhum gerenciador de pacotes encontrado (winget ou chocolatey).\n")
	// Se nenhum gerenciador de pacotes estiver disponível, retorna nil
	return nil, ""
}

func getLinuxInstallCommand(cmd *cobra.Command) (*exec.Cmd, string) {
	// Detecta o gerenciador de pacotes disponível
	if isCommandAvailable("apt-get") {
		verbosePrint(cmd, "Usando apt-get como gerenciador de pacotes.\n")
		// Para apt-get, retornamos um comando placeholder
		// O update será executado separadamente em installGit()
		return exec.Command("sudo", "apt-get", "install", "-y", "git"), "apt-get"
	}

	if isCommandAvailable("yum") {
		verbosePrint(cmd, "Usando yum como gerenciador de pacotes.\n")
		return exec.Command("sudo", "yum", "install", "-y", "git"), "yum"
	}

	if isCommandAvailable("dnf") {
		verbosePrint(cmd, "Usando dnf como gerenciador de pacotes.\n")
		return exec.Command("sudo", "dnf", "install", "-y", "git"), "dnf"
	}

	if isCommandAvailable("pacman") {
		verbosePrint(cmd, "Usando pacman como gerenciador de pacotes.\n")
		return exec.Command("sudo", "pacman", "-S", "--noconfirm", "git"), "pacman"
	}

	if isCommandAvailable("zypper") {
		verbosePrint(cmd, "Usando zypper como gerenciador de pacotes.\n")
		return exec.Command("sudo", "zypper", "install", "-y", "git"), "zypper"
	}

	verbosePrint(cmd, "Nenhum gerenciador de pacotes Linux encontrado.\n")
	return nil, ""
}

func getDarwinInstallCommand(cmd *cobra.Command) (*exec.Cmd, string) {
	// macOS - usa Homebrew
	if isCommandAvailable("brew") {
		verbosePrint(cmd, "Usando Homebrew como gerenciador de pacotes.\n")
		return exec.Command("brew", "install", "git"), "homebrew"
	}

	verbosePrint(cmd, "Homebrew não encontrado. Instale em: https://brew.sh\n")
	return nil, ""
}

func isCommandAvailable(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}

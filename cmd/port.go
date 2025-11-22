package cmd

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

var (
	portHost    string
	portTimeout int
)

var portCmd = &cobra.Command{
	Use:   "port",
	Short: "Verifica se uma porta está em uso",
	Long: `Verifica se uma porta específica está em uso ou disponível.
Pode verificar portas locais ou remotas.

Exemplos:
  bast port 8080              # Verifica porta 8080 em localhost
  bast port 3000 --host google.com  # Verifica porta 3000 em google.com
  bast port 22 --timeout 5     # Verifica com timeout de 5 segundos
  bast port --help             # Mostra ajuda deste comando`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		portStr := args[0]
		port, err := strconv.Atoi(portStr)
		if err != nil {
			fmt.Printf("Erro: '%s' não é uma porta válida.\n", portStr)
			verbosePrint(cmd, "Erro ao converter porta: %v\n", err)
			return
		}

		if port < 1 || port > 65535 {
			fmt.Printf("Erro: porta deve estar entre 1 e 65535.\n")
			return
		}

		verbosePrint(cmd, "Verificando porta %d em %s...\n", port, portHost)
		checkPort(cmd, port, portHost, portTimeout)
	},
}

func init() {
	rootCmd.AddCommand(portCmd)

	portCmd.Flags().StringVarP(&portHost, "host", "H", "localhost", "Host para verificar a porta")
	portCmd.Flags().IntVarP(&portTimeout, "timeout", "t", 3, "Timeout em segundos")
}

func checkPort(cmd *cobra.Command, port int, host string, timeout int) {
	address := fmt.Sprintf("%s:%d", host, port)
	verbosePrint(cmd, "Tentando conectar em %s...\n", address)

	timeoutDuration := time.Duration(timeout) * time.Second
	conn, err := net.DialTimeout("tcp", address, timeoutDuration)

	if err != nil {
		// Se não conseguiu conectar, a porta provavelmente está livre
		if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
			fmt.Printf("⏱️  Timeout ao conectar em %s:%d\n", host, port)
			fmt.Printf("   A porta pode estar fechada ou o host não está acessível.\n")
			verbosePrint(cmd, "Timeout após %d segundos.\n", timeout)
		} else {
			fmt.Printf("✅ Porta %d em %s está DISPONÍVEL\n", port, host)
			verbosePrint(cmd, "Erro de conexão (esperado para porta livre): %v\n", err)
		}
		return
	}

	defer conn.Close()

	// Se conseguiu conectar, a porta está em uso
	fmt.Printf("❌ Porta %d em %s está EM USO\n", port, host)
	fmt.Printf("   Endereço: %s\n", address)
	verbosePrint(cmd, "Conexão estabelecida com sucesso, porta está em uso.\n")

	// Tenta obter informações adicionais
	localAddr := conn.LocalAddr()
	remoteAddr := conn.RemoteAddr()
	verbosePrint(cmd, "Endereço local: %s\n", localAddr)
	verbosePrint(cmd, "Endereço remoto: %s\n", remoteAddr)
}

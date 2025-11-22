package cmd

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/CristianSsousa/go-bast-cli/internal/constants"
	"github.com/spf13/cobra"
)

var (
	port     string
	host     string
	endpoint string
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Inicia um servidor HTTP",
	Long: `Inicia um servidor HTTP na porta especificada.
Por padrão, o servidor roda na porta 8080.

Exemplos:
  bast serve                      # Inicia na porta 8080 (padrão)
  bast serve --port 3000         # Inicia na porta 3000
  bast serve -p 3000 -H localhost # Inicia na porta 3000 em localhost
  bast serve --help              # Mostra ajuda deste comando`,
	Run: func(cmd *cobra.Command, args []string) {
		startServer(cmd)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().StringVarP(&port, "port", "p", strconv.Itoa(constants.DefaultPort), "Porta do servidor")
	serveCmd.Flags().StringVarP(&host, "host", "H", constants.DefaultHost, "Host do servidor")
	serveCmd.Flags().StringVarP(&endpoint, "endpoint", "e", "/", "Endpoint principal")
}

func startServer(cmd *cobra.Command) {
	addr := fmt.Sprintf("%s:%s", host, port)

	verbosePrint(cmd, "Configurando servidor HTTP...\n")
	verbosePrint(cmd, "Host: %s\n", host)
	verbosePrint(cmd, "Porta: %s\n", port)
	verbosePrint(cmd, "Endpoint principal: %s\n", endpoint)

	http.HandleFunc("/", handler)
	http.HandleFunc("/health", healthHandler)

	srv := &http.Server{
		Addr:         addr,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	verbosePrint(cmd, "ReadTimeout: %v\n", srv.ReadTimeout)
	verbosePrint(cmd, "WriteTimeout: %v\n", srv.WriteTimeout)
	verbosePrint(cmd, "IdleTimeout: %v\n", srv.IdleTimeout)

	log.Printf("Servidor iniciando em http://%s", addr)
	log.Printf("Endpoint principal: %s", endpoint)
	verbosePrint(cmd, "Servidor pronto para receber conexões.\n")
	if err := srv.ListenAndServe(); err != nil {
		verbosePrint(cmd, "Erro ao iniciar servidor: %v\n", err)
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World! Este é um servidor CLI construído com Go e Cobra.")
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "OK")
}

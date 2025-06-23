
package main

import (
	"io/ioutil"	
	"log"
	"os"
	"fmt"
	"strconv"
	"github.com/skip2/go-qrcode"	
)

type QRConfig struct {
	InputFile  string
 	OutputFile string
	Text               string
	Size               int
	RecoveryLevel	  qrcode.RecoveryLevel
}

func main() {

	config := QRConfig{
		InputFile:       "teste.txt",
		OutputFile:      "custom-qr.png",
		Size:            512,
		RecoveryLevel:   qrcode.Medium,
	}

	// Leitura de argumentos ou pode ser modificado para ler de arquivo de configuração
	if len(os.Args) > 1 {
		config = parseConfig(os.Args)
	}

	// Carrega o texto do arquivo se não foi fornecido diretamente
	if config.Text == "" {
		content, err := ioutil.ReadFile(config.InputFile)
		if err != nil {
			log.Fatal("Erro ao ler arquivo de entrada: ", err)
		}
		config.Text = string(content)
	}

	// Gera o QR code com as configurações
	err := generateCustomQR(config)
	if err != nil {
		log.Fatal("Erro ao gerar QR code: ", err)
	}

	fmt.Printf("QR code gerado com sucesso em %s\n", config.OutputFile)
}

func parseConfig(args []string) QRConfig {
	config := QRConfig{}

	// Implementação simplificada - em produção, usar flag package ou viper
	for i := 1; i < len(args); i++ {
		switch args[i] {
		case "-i":
			config.InputFile = args[i+1]	// Nome do arquivo de entrada
			i++
		case "-o":
			config.OutputFile = args[i+1]	// Nome do arquivo de saída
			i++
		case "-t":
			config.Text = args[i+1]			// Texto para o QR code
			i++
		case "-s":
			size, _ := strconv.Atoi(args[i+1])		// Tamanho do QR code
			config.Size = size
			i++
		case "-lvl":
			config.RecoveryLevel = parseRecoveryLevel(args[i+1])	// Nível de recuperação
			i++
		}
	}

	return config
}

func parseRecoveryLevel(level string) qrcode.RecoveryLevel {
	switch level {
	case "L":
		return qrcode.Low
	case "M":
		return qrcode.Medium
	case "Q":
		return qrcode.High
	case "H":
		return qrcode.Highest
	default:
		return qrcode.Medium
	}
}

func generateCustomQR(config QRConfig) error {
	var q *qrcode.QRCode
	var err error

	// Cria o QR code básico
	q, err = qrcode.New(config.Text, config.RecoveryLevel)
	if err != nil {
		return err
	}

	// Gera a imagem final
	return q.WriteFile(config.Size, config.OutputFile)
}

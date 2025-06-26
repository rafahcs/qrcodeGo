
package main

import (
	"io/ioutil"	
	"image/color"
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
	ForegroundColor    color.RGBA
	BackgroundColor    color.RGBA
	
}

func main() {

	config := QRConfig{
		InputFile:       "teste.txt",
		OutputFile:      "custom-qr.png",
		Size:            512,
		RecoveryLevel:   qrcode.Medium,
		ForegroundColor: color.RGBA{0, 0, 0, 255},       // Preto
		BackgroundColor: color.RGBA{255, 255, 255, 255}, // Branco
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
		case "-fg":
			config.ForegroundColor = parseColor(args[i+1])	// Cor do primeiro plano
			i++
		case "-bg":
			config.BackgroundColor = parseColor(args[i+1])	// Cor do fundo
			i++
		}
	}

	return config
}

func parseColor(hex string) color.RGBA {
	if len(hex) != 6 && len(hex) != 8 {
		return color.RGBA{0, 0, 0, 255}
	}

	r, _ := strconv.ParseUint(hex[0:2], 16, 8)
	g, _ := strconv.ParseUint(hex[2:4], 16, 8)
	b, _ := strconv.ParseUint(hex[4:6], 16, 8)
	a := uint8(255)
	if len(hex) == 8 {
		a64, _ := strconv.ParseUint(hex[6:8], 16, 8)
		a = uint8(a64)
	}

	return color.RGBA{R: uint8(r), G: uint8(g), B: uint8(b), A: uint8(a)}
}

func parseRecoveryLevel(level string) qrcode.RecoveryLevel {
	switch level {
	case "L":
		return qrcode.Low
	case "M":
		return qrcode.Medium
	case "H":
		return qrcode.High
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

	// Configurações de cor
	q.ForegroundColor = config.ForegroundColor
	q.BackgroundColor = config.BackgroundColor


	// Gera a imagem final
	return q.WriteFile(config.Size, config.OutputFile)
}

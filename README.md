# QR Code em Go

Gerador de QR code com a linguagem Golang

Baseado no pacote [go-qrcode](https://github.com/skip2/go-qrcode)

# Como rodar
``
go run . -i <InputFile> -o <OutputFile> -s <Size> -fg <ForegroundColor> -bg <BackgroundColor> -lvl <RecoveryLevel>
``

* InputFile: nome do arquivo .txt com o link do QR code
* OutputFile: nome no arquivo .png com a imagem do QR code
* Size: tamanho da imagem 
* ForegroundColor: cor do desenho do código em hexadecimal 
* BackgroundColor: cor do fundo da imagem em hexadecimal 
* RecoveryLevel: nível de qualidade de leitura do qr-code (L: low, M: medium, H: high)

Exemplo:
* QR code preto-branco  
```bash
 go run . -i teste.txt -o preto-branco.png -s 256 -fg 000000 -bg FFFFFF -lvl M
```

![preto-branco](https://github.com/user-attachments/assets/1b9042cf-22ea-4b03-8c17-f5cbf68cdb45)


* QR code verde-amarelo  
```bash
go run . -i teste.txt -o verde-amarelo.png -s 128 -fg 008000 -bg FFFF00  -lvl L
```

![verde-amarelo](https://github.com/user-attachments/assets/811210e8-efdd-4d38-90d6-c97d099c1b3e)




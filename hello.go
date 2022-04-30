package main // informa que este arquivo será o ponto inicial do sistema

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const vezesAMonitorar = 3
const delay = 3

func main() {

	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	nome := "Fabiana"              // formas diferentes de declarar variável, por inferência e ser `var` (declaração curta)
	var versao float32 = 1.1       // de forma explícita
	fmt.Println("Olá, sra.", nome) // toda função externa ao nosso software (que provém de algum pacote) deve iniciar com letra maiúscula
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair do programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido) // salvar o retorno do usuário na variável `comando` (`&` refere-se ao endereço da variável passada -> ponteiro de memória)

	fmt.Println("O comando escolhido foi", comandoLido)
	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := []string{"https://random-status-code.herokuapp.com", "https://www.alura.com", "https://www.google.com"}

	for i := 0; i < vezesAMonitorar; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testSite(site)
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testSite(site string) {
	resposta, _ := http.Get(site)

	if resposta.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resposta.StatusCode)
	}
}

// para executar, podemos dar `go build hello.go` que builda o projeto e depois executamos com `./hello`.
// esse comando pode ser simplificado, usando o comando `go run hello.go` que builda e já executa na sequência

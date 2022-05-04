package main // informa que este arquivo será o ponto inicial do sistema

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
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
			imprimeLogs()
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

	sites := leSitesDoArquivo()

	for i := 0; i < vezesAMonitorar; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSite(site)
		}

		time.Sleep(delay * time.Second)
		fmt.Println("")
	}
	fmt.Println("")
}

func testaSite(site string) {
	resposta, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	var status bool
	if resposta.StatusCode == 200 {
		fmt.Println("Site:", site, "foi carregado com sucesso!")
		status = true
	} else {
		fmt.Println("Site:", site, "está com problemas. Status Code:", resposta.StatusCode)
		status = false
	}

	registraLog(site, status)
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)
	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)

		sites = append(sites, linha)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	return sites
}

func registraLog(site string, status bool) {
	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	now := time.Now().Format("02/01/2006 15:04:05") // os números representam constantes do pacote "time"
	arquivo.WriteString(now + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	arquivo, err := ioutil.ReadFile("log.txt") // retorna um byte array

	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	fmt.Println(string(arquivo))
}

// para executar, podemos dar `go build hello.go` que builda o projeto e depois executamos com `./hello`.
// esse comando pode ser simplificado, usando o comando `go run hello.go` que builda e já executa na sequência

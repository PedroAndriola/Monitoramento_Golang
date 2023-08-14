package main

import ( 
	"io"
	"strings"
	"fmt"	
    "os"
    "net/http"
    "time"
	"bufio"
	"strconv"
	"io/ioutil"
)

const monitora = 5
const delay = 3

func main()  {

	 exibeComeco()

	for { 
		exibeMenu()

		comando := leComando()
		

		switch comando {
		case 1:
			initMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
			imprimeLogs()
		case 0:
			fmt.Println("Finalizando Programa!")
			os.Exit(0)
		default:
			fmt.Println("Desconheço essa opção")		
			os.Exit(-1)	
	}
}	
}

func exibeComeco() {
	nome := "Chico"
	versao := 1.0
	fmt.Println("Olá, sr:", nome)
	fmt.Println("Versão", versao)
}	


func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("o comando escolhido foi: ", comandoLido)

	return comandoLido
}

func initMonitoramento() {
	fmt.Println("Monitorando...")

	sites := leSites()
	

	for i := 0; i < monitora; i++ { 	
		for i, site := range sites {
			fmt.Println("Estou na posição", i ," : ", site)
			testaSite(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println(":")
	}	
	fmt.Println(":")
}

func testaSite(site string) {
    resp, err := http.Get(site)

    if err != nil {
        fmt.Println("Ocorreu um erro:", err)
    }

    if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
        registraLog(site, true)
    } else {
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
        registraLog(site, false)
    }
}

func leSites() []string {

	var sites []string

	arquivo, err := os.Open("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)
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

func imprimeLogs() {

	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println("Houve um erro", err)
	}

	fmt.Println(string(arquivo))
}

func registraLog(site string, status bool) {

	arquivo, err := os.OpenFile("log.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	
	if err != nil  {
		fmt.Println(err)
	}

	fmt.Println(arquivo)

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")
	arquivo.Close()
}


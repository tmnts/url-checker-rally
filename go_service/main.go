package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

// Result описывает структуру данных ответа
// Result descreve a estrutura de dados da resposta
type Result struct {
	URL       string
	Status    int
	LatencyMS int64
	Error     string
}

func checkURL(url string, ch chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()

	client := http.Client{
		Timeout: 5 * time.Second,
	}

	startTime := time.Now()
	resp, err := client.Get(url)
	duration := time.Since(startTime).Milliseconds()

	if err != nil {
		ch <- Result{URL: url, Status: 0, LatencyMS: 0, Error: err.Error()}
		return
	}
	defer resp.Body.Close()

	ch <- Result{URL: url, Status: resp.StatusCode, LatencyMS: duration, Error: ""}
}

func main() {
	fmt.Println("Go: Начало обработки... / Processamento iniciado...")
	start := time.Now()

	baseUrls := []string{
		"https://google.com", "https://github.com", "https://microsoft.com",
		"https://kitdepontos.com.br", "https://lovefunk.com", "https://python.org",
	}

	// Размножаем список аналогично Python-скрипту
	// Replicar a lista, de forma semelhante ao script em Python
	var urls []string
	for i := 0; i < 10; i++ {
		urls = append(urls, baseUrls...)
	}

	resultsChan := make(chan Result, len(urls))
	var wg sync.WaitGroup

	// Запуск конкурентных горутин для каждого URL
	// Iniciando goroutines concorrentes para cada URL
	for _, url := range urls {
		wg.Add(1)
		go checkURL(url, resultsChan, &wg)
	}

	// Закрываем канал, когда все горутины завершат работу
	// Feche o canal quando todas as goroutines tiverem terminado a execução.
	go func() {
		wg.Wait()
		close(resultsChan)
	}()

	// Создаем файл отчета
	// Criar um arquivo de relatório
	file, err := os.Create("report_go.csv")
	if err != nil {
		fmt.Println("Ошибка создания файла / Erro ao criar arquivo", err)
		return
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Заголовки таблицы
	// Cabeçalhos da tabela
	_ = writer.Write([]string{"url", "status", "latency_ms", "error"})

	// Читаем из канала и пишем в CSV «на лету»
	// Ler do canal e gravar em CSV em tempo real
	for res := range resultsChan {
		_ = writer.Write([]string{
			res.URL,
			strconv.Itoa(res.Status),
			strconv.FormatInt(res.LatencyMS, 10),
			res.Error,
		})
	}

	fmt.Printf("Go: Готово за %v. Результаты в report_go.csv / Concluído em %v. Resultados em report_go.csv.\n", time.Since(start), time.Since(start))
}

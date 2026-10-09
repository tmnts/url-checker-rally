package main

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

// Тестируем функцию checkURL с использованием мок-сервера
// Testa a função checkURL usando um servidor mockado
func TestCheckURL_Success(t *testing.T) {
	// 1. Создаем тестовый сервер, который всегда отвечает кодом 200 OK
	// 1. Cria um servidor de teste que sempre responde com 200 OK
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	ch := make(chan Result, 1)
	var wg sync.WaitGroup

	wg.Add(1)
	// 2. Запускаем нашу реальную функцию, передавая ей адрес тестового сервера
	// 2. Executa a nossa função real, passando o endereço do servidor de teste
	go checkURL(server.URL, ch, &wg)
	wg.Wait()
	close(ch)

	// 3. Проверяем результат из канала
	// 3. Verifica o resultado vindo do canal
	res := <-ch
	if res.Status != 200 {
		t.Errorf("Ожидался статус 200, получили /// %d /// Esperado status 200, obtido", res.Status)
	}
	if res.Error != "" {
		t.Errorf("Ожидалась пустая ошибка, получили: /// %s /// Esperada mensagem de erro vazia, obtida: ", res.Error)
	}
}

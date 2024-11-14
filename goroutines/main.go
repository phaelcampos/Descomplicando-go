package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

type Pokemon struct {
	Name string
}

func getPokemonInfo(url string, c chan Pokemon) {
	fmt.Println(url)
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("O erro foi", err)
		return
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("O erro do corpo foi", err)
		return
	}

	var pokemon Pokemon
	err = json.Unmarshal(body, &pokemon)
	if err != nil {
		fmt.Println("O erro ao decodificar o json", err)
		return
	}
	c <- pokemon
}

func main() {
	pokemonurls := []string{
		"https://pokeapi.co/api/v2/pokemon/1",
		"https://pokeapi.co/api/v2/pokemon/2",
		"https://pokeapi.co/api/v2/pokemon/3",
		"https://pokeapi.co/api/v2/pokemon/4",
		"https://pokeapi.co/api/v2/pokemon/5",
		"https://pokeapi.co/api/v2/pokemon/6",
	}

	c := make(chan Pokemon)

	for _, url := range pokemonurls {
		go getPokemonInfo(url, c)
	}

	for i := 0; i < len(pokemonurls); i++ {
		pokemon := <-c
		fmt.Println(pokemon.Name)
	}
	main1()
	main2()
	main3()
}

// ------------------------------------------------------------
// Um outro exemplo de channel que podemos usar para compartilhar dados
// entre goroutines é esse, o de uma arquitetura simples (beeem simples)
// de producer/consumer dentro da sua app:
func producer(ch chan<- int) {
	for i := 1; i <= 5; i++ {
		fmt.Println("Producer: Enviando valor", i)
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch)
}

func consumer(ch <-chan int) {
	for num := range ch {
		fmt.Println("Consumer: Recebendo valor", num)
		time.Sleep(time.Millisecond * 500)
	}
}

func main1() {
	ch := make(chan int)

	go producer(ch)
	consumer(ch)

	fmt.Println("Programa finalizado")
}

// --------------------------------------------------------------
// Supondo que meu interesse nesse código aqui seja apenas saber
//
//	quando cada todos os workers finalizaram, ao invés de um channel
//	posso usar uma estrutura de dados chamada WaitGroup (o nome já diz tudo)
//	para esperar até que essa execução termine. O código:
func main2() {
	var wg sync.WaitGroup
	workers := 5

	wg.Add(workers)

	for i := 0; i < workers; i++ {
		go func(id int) {
			time.Sleep(time.Second * time.Duration(id+1))
			wg.Done()
		}(i)
	}

	wg.Wait()

}

//------------------------------------------

var counter int
var m sync.Mutex

func increment() {
	m.Lock()
	counter++
	m.Unlock()
}

func main3() {
	for i := 0; i < 1000; i++ {
		go increment()
	}

	// Wait for goroutines to complete
	time.Sleep(time.Second)

	fmt.Println("Counter:", counter)
}

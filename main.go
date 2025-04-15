package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	const cep = "80530908"

	viaCepUrl := fmt.Sprintf("https://viacep.com.br/ws/%v/json/", cep)
	brasilApiUrl := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%v", cep)

	viaCepChannel := make(chan string)
	brasilApiChannel := make(chan string)

	//ViaCEP
	go func() {
		res, err := http.Get(viaCepUrl)
		if err != nil {
			viaCepChannel <- ""
			return
		}
		defer res.Body.Close()
		r, err := io.ReadAll(res.Body)
		if err != nil {
			viaCepChannel <- ""
			return
		}
		viaCepChannel <- string(r)
	}()

	//BrasilAPI
	go func() {
		res, err := http.Get(brasilApiUrl)
		if err != nil {
			brasilApiChannel <- ""
			return
		}
		defer res.Body.Close()
		r, err := io.ReadAll(res.Body)
		if err != nil {
			brasilApiChannel <- ""
			return
		}
		brasilApiChannel <- string(r)
	}()

	select {
	case viaCep := <-viaCepChannel:
		fmt.Println("ViaCep:", viaCep)
	case brasilApi := <-brasilApiChannel:
		fmt.Println("BrasilApi:", brasilApi)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")
	}
}

package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	urls := []string{"https://comee.page.link/youtube", "https://comee.page.link/instagram", "https://comee.page.link/tikc", "https://comee.page.link/fbc"}

	for {
		rand.Seed(time.Now().UnixNano())

		for _, url := range urls {
			times := rand.Intn(1101) + 900 // Genera un número aleatorio entre 900 y 2000 para cada URL

			for i := 0; i < times; i++ {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Printf("Error al hacer la solicitud a %s: %v\n", url, err)
				} else {
					fmt.Printf("Solicitud exitosa a %s\n", url)
					resp.Body.Close()
				}
			}
		}

		time.Sleep(24 * time.Hour) // Espera 24 horas antes de repetir el ciclo
	}
}

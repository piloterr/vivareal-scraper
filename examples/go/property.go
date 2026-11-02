// Viva Real Property API – Go example
// Usage: PILOTERR_API_KEY=your-key go run property.go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"time"
)

func main() {
	params := url.Values{}
	params.Set("query", "https://www.vivareal.com.br/imovel/apartamento-3-quartos-brooklin-zona-sul-sao-paulo-com-garagem-186m2-venda-RS3710000-id-2776423622/")

	req, _ := http.NewRequest("GET", "https://api.piloterr.com/v2/vivareal/property?"+params.Encode(), nil)
	req.Header.Set("x-api-key", os.Getenv("PILOTERR_API_KEY"))

	client := &http.Client{Timeout: 120 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		fmt.Fprintf(os.Stderr, "Piloterr error %d: %s\n", resp.StatusCode, body)
		os.Exit(1)
	}

	var pretty bytes.Buffer
	json.Indent(&pretty, body, "", "  ")
	fmt.Println(pretty.String())
}

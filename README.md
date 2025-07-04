# Hexcore

![Go Version](https://img.shields.io/badge/go-1.24-blue.svg)
![License](https://img.shields.io/github/license/rico-vz/hexcore)
![Go Report Card](https://goreportcard.com/badge/github.com/rico-vz/hexcore)

Hexcore is an easy-to-use library for interacting with the League of Legends Client Update (LCU) API in Go.

---

## Features

-   Automatic finding of LCU port and authorization token.
-   Typed structs for the LCU API endpoints.
-   WebSocket support for LCU events.

---

## Installation

To install the Hexcore package, use the `go get` command:

```bash
go get https://github.com/rico-vz/hexcore/lcu
```

---

## Usage

Basic example of how to init a new client and get the current summoner's information. The client will automatically find the LoL/LCU process to connect to.

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/rico-vz/hexcore/lcu"
)

func main() {
	log.Println("Connecting to LCU API...")

	client, err := lcu.NewHexcoreClient()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	log.Printf("Connected to LCU on port %s\n", client.Params.Port)

	summonerResponse, err := client.GetLolSummonerV1CurrentSummonerWithResponse(context.Background())
	if err != nil {
		log.Fatalf("Failed to get current summoner: %v", err)
	}

	if summonerResponse.StatusCode() == http.StatusOK {
		var summonerData lcu.LolSummonerSummoner
		err := json.Unmarshal(summonerResponse.Body, &summonerData)
		if err != nil {
			log.Fatalf("Failed to response body: %v", err)
		}

		fmt.Printf("Current Player: %s\n", summonerData.GameName)
		fmt.Printf("Summoner Level: %d\n", summonerData.SummonerLevel)
	} else {
		fmt.Println("Could not retrieve summoner data.")
		fmt.Printf("Response status: %s\n", summonerResponse.Status())
		fmt.Printf("Response body: %s\n", string(summonerResponse.Body))
	}
}
```

Output example:

```text
2025/07/04 02:05:17 Connecting to LCU API...                                                                                                                                                                                                           
2025/07/04 02:05:17 Connected to LCU on port 4218                                                                                                                                                                                                      
Current Player: Z
Summoner Level: 205
```

---

## Contributing

Contributions are welcome. Please feel free to submit a pull request or open an issue for any bugs or feature requests.

Hexcore follows [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/) for all commit messages.
Allowed types: `fix`, `feat`, `build`, `chore`, `ci`, `docs`, `style`, `refactor`, `perf` & `test`

1. Fork the repository
2. Create your feature branch (`git checkout -b feat/new-thing`)
3. Commit your changes (`git commit -m 'feat: brand new thing'`)
4. Push to your branch (`git push origin feat/new-thing`)
5. Open a Pull Request

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.


***Disclaimer:***  
*Hexcore isn't endorsed by Riot Games and doesn't reflect the views or opinions of Riot Games or anyone officially involved in producing or managing League of Legends.*
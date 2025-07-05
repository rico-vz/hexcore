# Hexcore

![Go Version](https://img.shields.io/badge/go-1.24-blue.svg)
![License](https://img.shields.io/github/license/rico-vz/hexcore)
![Go Report Card](https://goreportcard.com/badge/github.com/rico-vz/hexcore)

Hexcore is an easy-to-use library for interacting with the League of Legends Client Update (LCU) API in Go.

---

## Features

-   Automatic finding of LCU port and authorization token.
-   Typed structs for the LCU API endpoints.
-   WebSocket support for LCU events with type-safe event topics.

---

## Installation

To install the Hexcore package, use the `go get` command:

```bash
go get https://github.com/rico-vz/hexcore/lcu
```

---

## Usage

Basic example of how to init a new client, get the current player's information and subscribe to an event:

```go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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

	// Getting the current player their information
	summonerResponse, err := client.GetLolSummonerV1CurrentSummonerWithResponse(context.Background())
	if err != nil {
		log.Fatalf("Failed to get current summoner: %v", err)
	}

	if summonerResponse.StatusCode() == http.StatusOK {
		var summonerData lcu.LolSummonerSummoner
		err := json.Unmarshal(summonerResponse.Body, &summonerData)
		if err != nil {
			log.Fatalf("Failed to parse response body: %v", err)
		}

		fmt.Printf("Current Player: %s\n", summonerData.GameName)
		fmt.Printf("Summoner Level: %d\n", summonerData.SummonerLevel)
	} else {
		fmt.Println("Could not retrieve summoner data.")
		fmt.Printf("Response status: %s\n", summonerResponse.Status())
		fmt.Printf("Response body: %s\n", string(summonerResponse.Body))
	}

	// Subscribing to an event (End of Game event in this example)
	log.Println("Subscribing to End of Game events...")
	err = client.Subscribe(lcu.Events.EndOfGameV1EogStatsBlock(), OnEndOfGameEvent)
	if err != nil {
		log.Fatalf("Failed to subscribe to End of Game event: %v", err)
	}

	log.Println("Subscribed to the End of Game event")
	log.Println("Listening for events...")

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
}

// This gets called when the event we subscribed to gets triggered
func OnEndOfGameEvent(payload interface{}) {
	log.Println("[OnEndOfGameEvent] Received the event!")
}
```

Output example:

```text
2025/07/05 10:33:44 Connecting to LCU API...
2025/07/05 10:33:44 Connected to LCU on port 3588
Current Player: Z
Summoner Level: 205
2025/07/05 10:33:44 Subscribing to End of Game events...
2025/07/05 10:33:44 Subscribed to the End of Game event
2025/07/05 10:33:44 Listening for events...
2025/07/05 10:35:11 [OnEndOfGameEvent] Received the event!
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

## Acknowledgments

- [dysolix/hasagi-types](https://github.com/dysolix/hasagi-types/) - Used to get the latest `swagger.json` file for the LCU API.
- [oapi-codegen](https://github.com/oapi-codegen/oapi-codegen) - Used to generate the API client code from the OpenAPI specification.
- [Riot Third Party Developer Discord](https://discord.com/invite/riotgamesdevrel) - Nice community. Useful for information about the LCU API and libraries.
- [HextechDocs](https://hextechdocs.dev/) - Community maintained developer documentation for games made by Riot Games.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.


***Disclaimer:***  
*Hexcore isn't endorsed by Riot Games and doesn't reflect the views or opinions of Riot Games or anyone officially involved in producing or managing League of Legends.*
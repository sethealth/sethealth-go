# Sethealth Go client

Sethealth Go client allows to access the backend sethealth API from a server.
The unique use case of this library today is to provide a authentication schema to delegate the "frontend" javascript library to communicate safely with the sethealth backend.

This is accomplish by the generation of a `service account` in sethealth. A service account is a long-living account for non-human users, like servers. Once a service account is created, a `api key` and a `api secret` are generated, this credentials **MUST be kept private, never exposed in a client side application.**

This "long-living" credentials can be used instead to create short-living credentials in the shape of `access tokens` in order to call the upload/download medical data from the client.

## Install

```
go get -u github.com/sethealth/go-client
```

## Usage

Get your service account credentials from the [Sethealth Dashboard](https://dashboard.set.health).

**.bashrc/.zshrc:**

```bash
export SETHEALTH_KEY="0000000000000"
export SETHEALTH_SECRET="r_xxxxxxxxxxxxxxxxxxxxxxxxxxxx="
```

**main.go:**

```go
package main

import (
    "github.com/sethealth/go-client"
    "fmt"
)

func main() {
    // Create a new sethealth client with the service account credentials
    client := sethealth.New()

    // Ask for a short-living access token
    token, err := client.GetToken()
    fmt.Println("ACCESS TOKEN", token.Token)
}
```

Alternatively, the credentials can be provided programatically using the `sethealth.NewWithCredentials()`:

```go
package main

import (
    "github.com/sethealth/go-client"
    "fmt"
)

func main() {
    // Create a new sethealth client with the service account credentials
    apiKey := "HERE THE API KEY"
    apiSecret := "HERE THE API SECRET"
    client := sethealth.NewWithCredentials(apiKey, apiSecret)

    // Ask for a short-living access token
    token, err := client.GetToken()
    fmt.Println("ACCESS TOKEN", token.Token)
}
```

>Note: Credentials should be kept secret, it's not a good practice to hard code them in the source code.

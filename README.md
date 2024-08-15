# The CRYPTO package

## The Akr78 method

This akr78 method in the crypto package is a method that is based on the sha3 and blake2b methods from the golang/crypto package and serves as a generalized method for their hashing techniques.
[![alt text](68747470733a2f2f706b672e676f2e6465762f62616467652f676f6c616e672e6f72672f782f63727970746f2e737667.svg)](https://pkg.go.dev/golang.org/x/crypto)

## Installation

```
go get github.com/akramov1ch/crypto@latest
```

## Usage Akr78

```
package main

import (
    "fmt"
    "panic"
    
    "github.com/akramov1ch/crypto/akr78"
)

func main(){
    data := "Hello, World"

    hash, err := akr78.Akr78(data)
    if err != nil {
        panic(err)
    }
    fmt.Println(hash)
}
```

## Usage VerifyAkr78

The VerifyAkr78 method takes both a non-hashed and a hashed string as inputs and determines whether they are the same string. This can be useful for logging in with a password.

```
package main

import (
    "fmt"
    "panic"
    
    "github.com/akramov1ch/crypto/akr78"
)

func main(){
    data := "Hello, World"

    hash, err := akr78.Akr78(data)
    if err != nil {
        panic(err)
    }

    ok, err := akr78.VerifyAkr78(data, hash)
    if err != nil {
        panic(err)
    }
    fmt.Println(ok)
}
```
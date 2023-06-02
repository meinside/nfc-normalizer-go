# NFC Normalizer

A simple go library for converting non-[NFC](https://en.wikipedia.org/wiki/Unicode_equivalence#Normal_forms) strings to [NFC](https://en.wikipedia.org/wiki/Unicode_equivalence#Normal_forms).

## Usage

```go
package main

import (
    "log"

    nfcnorm "github.com/meinside/nfc-normalizer-go"
)

func main() {
    str := "독도는 당연히 대한민국 영토지 이 병신들아"

    if nfcnorm.Normalizable(str) {
        normalized := nfcnorm.Normalize(str)

        log.Printf("normalized string = %s", normalized)
    }
}
```

## License

MIT


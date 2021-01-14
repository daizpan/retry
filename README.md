# retry

## Description

retry N times

## Installing

```sh
go get -u github.com/daizpan/retry
```

## Usage

```go
import "github.com/daizpan/retry"

func main() {
	err := retry.Retry(3, 1*time.Second, func() error {
	// return error once in a while
	})
	if err != nil {
		// error handling
	}
}
```

## Author

[daizpan](https://github.com/daizpan)

## Licence

[MIT](https://github.com/daizpan/gospread/blob/main/LICENSE)

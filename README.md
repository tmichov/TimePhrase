# TimePhrase - String to Time parser for Go

TimePhrase is a Go package that allows you to parse natural phrases into `time.Time` objects. You can easily parse dates like `"today"`, `"tomorrow"`, `"2 days ago"`, and exact timestamps like `"2024-12-18T10:00:00"`. This package is ideal for any application that requires handling natural language date inputs or timestamp parsing.

## Installation

To install the package, use the following command:

```sh
go get github.com/tmichov/TimePhrase
```

After installation, you can import the package into your Go project:

```go
import "github.com/tmichov/TimePhrase"
```

## Functions

### `Parse(dateStr string, base time.Time) (time.Time, error)`

This is the main function to parse a date string. It can handle both absolute and relative time phrases. If the input string represents a valid date or time phrase, it returns a `time.Time` object. Otherwise, it returns an error.

#### Arguments:
- `dateStr`: A string containing the date or time phrase to parse (e.g., `"today"`, `"2024-12-18"`, `"2 days ago"`).
- `base`: The base time, used for relative dates (e.g., `"yesterday"`, `"2 days ago"`). Typically, `time.Now()` is passed as the base time.

#### Returns:
- A `time.Time` object if the date string is parsed successfully.
- An error if the string could not be parsed.

#### Example Usage:

```go
package main

import (
    "fmt"
    "time"
    "github.com/tmichov/TimePhrase"
)

func main() {
    baseTime := time.Now()
    parsedTime, err := TimePhrase.Parse("tomorrow", baseTime)
    if err != nil {
        fmt.Println("Error:", err)
    }
    fmt.Println("Parsed time:", parsedTime)
}
```


## Supported Date Formats

### Absolute Time Formats:
- `"2006-01-02 15:04:05"`
- `"2006-01-02"`
- `"2006-01-02T15:04:05"`
- `"02 Jan 2006"`
- `"02 Jan 2006 15:04:05"`
- More formats for timezones, including `Z`, `+00:00`, and `-0700`.

### Relative Time Phrases:
- `"now"`
- `"today"`
- `"tomorrow"`
- `"yesterday"`
- `"2 days ago"`
- `"in 5 hours"`
- `"3 weeks from now"`
- `"1 year ago"`


## License

MIT License

Feel free to contribute to this project by opening issues or creating pull requests.

For any questions or issues, feel free to reach out!

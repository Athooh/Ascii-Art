# Ascii-art Project

This project is a Go program that generates ASCII art representations of text strings.

## Features

- Converts text strings into large-scale ASCII art.
- Handles numbers, letters, spaces, special characters, and newlines.
- Utilizes various banner files (shadow, standard, thinkertoy) for different styles.
- Provides unit testing for ensuring functionality.

## Getting Started

### Prerequisites

- Go compiler installed ([Download here](https://go.dev/dl/))

### Installation

1. Clone the repository

```bash
 git clone https://learn.zone01kisumu.ke/git/seodhiambo/ascii-art.git
 ```
2. Navigate to the project directory: 
```bash
cd ascii-art
```

### Usage

- Run the program with a string argument:
```bash
go run . "your_text_here" | cat -e
```

  - Example:
  ```bash
  user$ go run . "Hello\n" | cat -e
  ```
  ```bash
   _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                  $
                                  $
$
  ```

- Use `\n` within the string for manual line breaks:

```
user$ go run . "Hello\nWorld!" | cat -e
```
```console
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
__          __                 _       _   _  $
\ \        / /                | |     | | | | $
 \ \  /\  / /    ___    _ __  | |   __| | | | $
  \ \/  \/ /    / _ \  | '__| | |  / _` | | | $
   \  /\  /    | (_) | | |    | | | (_| | |_| $
    \/  \/      \___/  |_|    |_|  \__,_| (_) $
                                              $
                                              $
```

- Special characters can be escaped using `\`: `go run . "Special chars: \\n \\t"`

### Banner Files

- `standard.txt`: The default banner file with a classic ASCII art style.
- `shadow.txt`: A banner file that creates an outlined or shadowed effect.
- `thinkertoy.txt`: A more playful banner file reminiscent of construction toys.

### Tests

- Unit tests are located in `main_test.go` and can be run using: `go test`

## Roadmap

- Implement options for choosing specific banner files.
- Allow customization of the output size.
- Add support for colorized ASCII art.

## Contributing

Contributions are welcome! Please feel free to submit pull requests for bug fixes or new features.

## License

This project is licensed under the MIT License.

## Acknowledgments

- The Go programming language community.
- The creators of the various banner files used in this project.
- Group Contributors:
    - vinomondi
    - coketch
    - seodhiambo

# Scoop

Scoop is a CLI tool for effortlessly managing all your environment variables. Save environment variables locally and retrieve them with just one command. Built using only the standard libraries of Go. Scoop is simple, useful, and blazingly fast.

# Table of Contents

- [Installation](#installation)
- [Usage Example](#usage-example)
  - [Save Environment Variables](#save-environment-variables)
  - [Retrieve Environment Variables](#retrieve-environment-variables)
  - [Pro-Tip](#pro-tip)
- [Contributing](#contributing)
- [License](#license)

## Installation

To install Scoop, ensure you have Go installed on your machine.

```
go install github.com/SameerJadav/scoop
```

This will install a Go binary that automatically binds to your `$GOPATH`.

## Usage Example

Scoop is extremly simple to use. You can save environment variables with `scoop set <project> <key=value>` command and retrieve them with `scoop get <project>` command.

### Save Environment Variables

```
scoop set my-project DB_HOST=localhost DB_PORT=5432
```

### Retrieve Environment Variables

```
scoop get my-project
```

This command will display the saved environment variables.

```
DB_HOST=localhost
DB_PORT=5432
```

### Pro-Tip

You can easily export environment variables to a `.env` file by just redirecting the result of the `scoop get` command to an `.env` file.

```
scoop get my-project > .env
```

## Contributing

If you'd like to contribute to Scoop, please feel free to open an issue or submit a pull request. We welcome your suggestions and improvements!

## License

Scoop is open-source and available under the [MIT License](./LICENSE).

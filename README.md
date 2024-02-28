# Keyper

Keyper is a CLI tool for effortlessly managing your environment variables. Save environment variables locally and retrieve them with just one command. Built using only the standard libraries of Go. Keyper is simple, useful, and blazingly fast.

# Table of Contents

- [Installation](#installation)
- [Usage Example](#usage-example)
  - [Save Environment Variables](#save-environment-variables)
  - [Retrieve Environment Variables](#retrieve-environment-variables)
  - [Pro-Tip](#pro-tip)
- [Contributing](#contributing)
- [License](#license)

## Installation

To install Keyper, ensure you have Go installed on your machine. If you haven't installed Go yet, you can follow the [official Go installation instructions](https://go.dev/doc/install).

```
go install github.com/SameerJadav/keyper@latest
```

This will install a Go binary that automatically binds to your `$GOPATH`.

## Usage Example

Keyper is extremly simple to use. You can save environment variables with `keyper set <project> <key=value>` command and retrieve them with `keyper get <project>` command.

### Save Environment Variables

```
keyper set my-project DB_HOST=localhost DB_PORT=5432
```

### Retrieve Environment Variables

```
keyper get my-project
```

This command will display the saved environment variables.

```
DB_HOST=localhost
DB_PORT=5432
```

### Pro-Tip

You can easily export environment variables to a `.env` file by just redirecting the result of the `keyper get` command to an `.env` file.

```
keyper get my-project > .env
```

## Contributing

If you'd like to contribute to Keyper, please feel free to open an issue or submit a pull request. We welcome your suggestions and improvements!

## License

Keyper is open-source and available under the [MIT License](./LICENSE).

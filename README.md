# Keyper

Keyper is a CLI tool for effortlessly managing your environment variables. Save environment variables locally and retrieve them with just one command. Built using only the standard libraries of Go. Keyper is simple, useful, and blazingly fast.

# Table of Contents

- [Installation](#installation)
- [Usage Example](#usage-example)
  - [Save Environment Variables](#save-environment-variables)
  - [Retrieve Environment Variables](#retrieve-environment-variables)
  - [Remove Environment Variables](#remove-environment-variables)
  - [Purge Project and its Variables](#purge-project-and-its-variables)
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

Keyper is designed for simplicity. It's a [grug brain tool](https://grugbrain.dev/#grug-on-tools).

### Save Environment Variables

**Command:**

```
keyper set <project> <key=value> ...
```

**Example:**

```
keyper set my-project DATABASE_URL=postgresql://localhost/mydb DATABASE_AUTH_TOKEN=42069
```

Saves the specified key-value pairs as environment variables for the given project, overwriting existing variables with the same keys.

### Retrieve Environment Variables

**Command:**

```
keyper get <project>
```

**Example:**

```
keyper get my-project
```

Displays the saved environment variables for the specified project.

```
DATABASE_URL=postgresql://localhost/mydb
DATABASE_AUTH_TOKEN=42069
```

### Remove Environment Variables

**Command:**

```
keyper remove <project> <key> ...
```

**Example:**

```
keyper remove my-project DATABASE_URL
```

Permanently removes the specified environment variables from the given project.

### Purge Project and its Variables

**Command:**

```
keyper purge <project> ...
```

**Example:**

```
keyper purge my-project
```

Permanently removes the entire project and all its associated environment variables. Use this command with caution.

### Pro-Tip

You can easily export environment variables to a `.env` file by just redirecting the result of the `keyper get` command to an `.env` file.

```
keyper get my-project > .env
```

## Contributing

If you'd like to contribute to Keyper, please feel free to open an issue or submit a pull request. I welcome your suggestions and improvements!

## License

Keyper is open-source and available under the [MIT License](./LICENSE).

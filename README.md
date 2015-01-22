# valiant

A brave mail pigeon

## Installation

Golang users can compile it using the following command:

```
go get github.com/elwinar/valiant
```

Others can just go to the [releases](https://github.com/elwinar/valiant/releases) page, select the last version, and download the cross-compiled version for their OS and architecture.

## Usage

```
valiant --configuration ./path/to/the/configuration/file.json --body ./path/to/the/body/file.html --subject "Your subject here"
```

Simple, isn't it ?
The configuration option in the command-line defaults to `valiant.json`.

The configuration file goes as follow:

```json
{
	"server": {
		"tls": false,
		"host": "smtp.example.com",
		"port": 587,
		"user": "user@example.com",
		"password": "password"
	},
	"from": {
		"name": "User",
		"address": "user@example.com"
	},
	"to": [
		{
			"name": "Foo",
			"address": "foo@example.com"
		}
	]
}
```

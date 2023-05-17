# VectorGo

VectorGo is a powerful and flexible SQLite powered Vector Database in pure Go with optional embeddings generation via the OpenAI API. It allows users to store, search, and manage vector embeddings easily.

## Features

* SQLite powered vector database
* Pure Go implementation
* Option to generate embeddings using OpenAI API
* Support for various vector operations
* Easy to integrate into your projects

## Installation

To get started with VectorGo, simply add the package to your Go project:

```
go get -u github.com/chand1012/vectorgo
```

## Usage

First, create a new VectorDB instance by providing a path to the SQLite database file and an OpenAI API key:

```go
import (
        "github.com/chand1012/vectorgo"
)

config := &vectorgo.VectorDBConfig{
        Path:      "path/to/your/db/file",
        OpenAIKey: "your_openai_api_key",
}

vdb, err := vectorgo.NewVectorDB(config)
```

Now, you can start adding embeddings by providing identifiers and content:

```go
err := vdb.Add("identifier", "content")
```

To search for an embedding in plain text, use the `Search` method:

```go
results, err := vdb.Search("query text", 10)
```

Additionally, you can check if an embedding exists or matches content, delete embeddings, and perform other operations. For more information, check out the [documentation](https://pkg.go.dev/github.com/chand1012/vectorgo).

## Contributing

Feel free to contribute to VectorGo by submitting issues, pull requests, or feature requests.

## License

VectorGo is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

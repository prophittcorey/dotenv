# Dotenv

[![Go Reference](https://pkg.go.dev/badge/github.com/prophittcorey/dotenv.svg)](https://pkg.go.dev/github.com/prophittcorey/dotenv)

A golang package for the loading of localized environment variables.

## Package Usage

The package exposes a single function to call at the start of your program.

```golang
if err := dotenv.Load(); err != nil {
  log.Fatal("failed to load environment configuration")
}
```

## License

The source code for this repository is licensed under the MIT license, which you can
find in the [LICENSE](LICENSE.md) file.

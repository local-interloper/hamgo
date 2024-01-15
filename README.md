<h1 align="center">hamgo</h1>
<p align="center">
A set of CLI utilities for HAM operators.
Currently boasting an ADIF compatible logger with a bunch of neat comfort features.
</p>

---

# Why?

I wanted to learn Go and thought this could be a fun practical idea.

# Roadmap

- Add a CLI .adif viewer (`hamgo view <PATH>`)
- Add a callsign lookup (maybe hook into [qrz.com](https://qrz.com))
- Add a prefix lookup system

# Building

1. Clone the repo and step into it

```shell
git clone https://github.com/local-interloper/hamgo
cd hamgo
```

2. From here you can either:

```shell
# Run the project by running
go run .

# Build the binary by running
go build
```

# License

Check the `LICENSE` file
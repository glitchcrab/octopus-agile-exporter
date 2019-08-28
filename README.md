# Octopus Agile Exporter

Octopus Agile Exporter is a Prometheus exporter which gets the current unit price for electricity on Octopus Energy's Agile tariff.

## Build from source

```bash
go build
```

## Usage

With flags:

```bash
./octopus-agile-exporter --product-code="" --tariff-code=""
```

Or with environment variables:

```bash
export OCTOPUS_PRODUCT_CODE=""
export OCTOPUS_TARIFF_CODE=""
./octopus-agile-exporter
```

### Notes

- The exporter exposes the metrics on port `8080` under the path `/metrics`.
- It updates the value every 60 seconds.

## Docker

An image is already available on the [Docker hub](https://hub.docker.com/r/glitchcrab/octopus-agile-exporter). (currently it only works with env vars (see #3))

```bash
docker run -it --rm -e OCTOPUS_PRODUCT_CODE="" -e OCTOPUS_TARIFF_CODE="" glitchcrab/octopus-agile-exporter:latest
```

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

# Bigtable

## Prerequisite

### Install Go Packages

```bash
go mod tidy
```

### Setup local bigtable emulator

Follow the official references to set up the [Emulator](https://cloud.google.com/bigtable/docs/emulator) and optionally install the [cbt-tool](https://cloud.google.com/bigtable/docs/cbt-overview).

## Local Development

Make sure the `BIGTABLE_EMULATOR_HOST` environment variable is set. If not, run `$(gcloud beta emulators bigtable env-init)` after setup the emulator.

Create a file named `.env` and write the following content into it

```
PROJECT_ID=my-project
INSTANCE_ID=my-instance
TABLE_NAME=vaccine-reservation-system
```

### Database seeding

Run the following command to setup the database.

```bash
go run main.go
```

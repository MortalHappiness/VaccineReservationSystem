# VaccineReservationSystem


## Deploy development environment

1. Deploy backend and bigtable

```bash
docker-compose -f dev.docker-compose.yml up -d
```

2. Initialize database

```bash
export BIGTABLE_EMULATOR_HOST=localhost:8086
```

Follow the rest instructions in [bigtable - Setup local bigtable emulator](/bigtable/README.md)
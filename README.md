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

3. See spec

- See **User** and **Session** api spec at `http://localhost:7712/.spec/`
- See **Reservation** api spec at `http://localhost:7713/.spec/`
- See **Hospital** api spec at `http://localhost:7714/.spec/`


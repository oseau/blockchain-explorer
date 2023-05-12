# blockchain-explorer

run `make` to see available commands.

## backend

### create/update database schema

```bash
make login
go run -mod=mod entgo.io/ent/cmd/ent new Balance # if adding new table 'balance'
# edit backend/ent/schema/balance.go
go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/upsert ./ent/schema

go run -mod=mod entgo.io/ent/cmd/ent describe ./ent/schema # show table
```

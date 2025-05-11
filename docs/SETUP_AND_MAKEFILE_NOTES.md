# Project Setup & Makefile Usage Notes

## 1. Environment Setup

- Copy `.envrc` to your local machine and update values as needed.
- Ensure `DB_ADDR` is set to your database connection string (e.g., `postgres://user:password@localhost:5432/dbname?sslmode=disable`).
- Install the [migrate](https://github.com/golang-migrate/migrate) CLI tool if not already installed.

## 2. Makefile Commands

### Create a New Migration
```
make migration NAME=your_migration_name
```
- This creates new `.up.sql` and `.down.sql` files in `cmd/migrate/migrations/`.
- Example: `make migration NAME=add_comments`

### Apply All Up Migrations
```
make migrate-up
```
- Runs all pending migrations against the database specified in `DB_ADDR`.

### Apply All Down Migrations
```
make migrate-down
```
- Rolls back all migrations. **Use with caution!**

## 3. Handling Migration Errors

### Dirty Database State
If you see an error like `Dirty database version X. Fix and force version.`:
1. Force the migration state to the correct version:
   ```sh
   migrate -path=./cmd/migrate/migrations -database='$DB_ADDR' force X
   ```
   - Replace `X` with the version number from the error.
   - If using zsh, wrap the database URL in single quotes to avoid shell issues.
2. Re-run migrations as needed.

## 4. Additional Notes
- All migration files are located in `cmd/migrate/migrations/`.
- The Makefile expects the `migrate` CLI to be available in your PATH.
- For more details, see the [migrate documentation](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate). 
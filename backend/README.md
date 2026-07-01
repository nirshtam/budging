# Budging Backend

Budging backend is a Go backend for a personal budgeting and expense tracking application.
The backend is designed around Open Banking data, with Enable Banking used as the external provider for bank account and transaction access.

The current goal of the backend is to provide a clean, reliable API that can:

- connect to an Open Banking provider
- fetch banking/account data
- expose REST endpoints
- store transaction data in PostgreSQL
- use Redis for background sync / queueing
- keep the core application independent from infrastructure details

## Tech Stack

- Go
- net/http
- PostgreSQL
- Redis
- Docker
- Enable Banking API

## Architecture

The backend follows a lightweight hexagonal architecture to avoid complexity but respect best practices of Dependency Inversion and find a practical solution.

## File Structure

```text
--|backend
----| internal - all layers
------| adapters
--------| enablebanking - Open banking data aggregator
--------| http - handle http server
--------| postgres - DB
------| core - core layer
--------| application - services / use cases
--------| domain - models / schemas
--------| ports - abstractions

# Hanzo Payments

Hanzo Payments is the payments connectivity engine of **Hanzo Finance**. It is a framework for ingesting pay-ins and pay-outs from many different Payment Service Providers (PSPs) and normalizing them into a single, generalized format.

The framework is built around **connectors**. Each connector is a translator for a PSP: its main role is to translate a specific PSP's pay-in/pay-out formats into the generalized model used across Hanzo. Because it is a framework, it is extensible — follow the contributor guide below to add your own connector.

## Getting started

Payments runs as a standalone binary. Installing it locally with Docker is easiest for testing, as the compose stack comes prepackaged with all dependencies.

Set the `STACK_PUBLIC_URL` env variable to a publicly available URL so webhooks and redirects from PSPs can reach the application. You can use a tunnel such as [ngrok](https://ngrok.com/) for this. If you do not plan to use connectors with webhooks, you may set it to localhost.

```shell
git clone git@github.com:hanzo-fi/payments.git
cd payments
just compile-plugins
STACK_PUBLIC_URL=https://subdomain.ngrok-free.app docker compose up
```

### Debugging

Use `docker-compose.dev.yml` to run the application with Delve and Air for debugging and live reload.

## Storage

Hanzo Payments defaults to **Hanzo Base (embedded SQLite), per-tenant** — each org/project gets its own isolated store with no external database to run. **PostgreSQL remains a first-class, opt-in option** for shared/multi-instance scale (`STORAGE_DRIVER=postgres`).

## Modules

* **Community Edition** connectors live under `ce/plugins/` (MIT-licensed).
* **Enterprise Edition** connectors live under `ee/plugins/` (see [ee/LICENSE](./ee/LICENSE)); build with `-tags ee`.

## Contribute

- Connector development tutorial: [CONTRIBUTING.md](./CONTRIBUTING.md)
- General development guidelines: [CONTRIBUTING_GUIDE.md](./CONTRIBUTING_GUIDE.md)

## Attribution

Hanzo Payments is a fork of [Formance Payments](https://github.com/formancehq/payments). See [LICENSE](./LICENSE) — Community Edition is MIT, Enterprise Edition (`ee/`) is under [ee/LICENSE](./ee/LICENSE). Upstream copyright remains with Formance Solutions; Hanzo modifications © Hanzo AI, Inc.

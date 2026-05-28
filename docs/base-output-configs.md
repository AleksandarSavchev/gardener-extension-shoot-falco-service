# Base Falco Output Configs

## Overview

Landscape operators can define base Falcosidekick output configurations in the extension controller deployment. These are applied to all shoots on the seed unless overridden by the shoot's own destination configuration.

## Extension Controller Configuration

The `baseFalcoOutputConfigs` field in the extension config defines Falcosidekick output blocks. Each entry has a `key` (the Falcosidekick output name) and a `value` (the raw config map for that output).

```yaml
apiVersion: falco.extensions.config.gardener.cloud/v1alpha1
kind: Configuration
falco:
  baseFalcoOutputConfigs:
  - key: "splunk"
    value:
      hostport: "https://falco-splunk-ingestor.SEED_INGRESS"
      customheaders:
        Authorization: "Bearer SA_TOKEN"
      checkcert: true
  - key: "webhook"
    value:
      address: "https://falco-central-ingestor.SEED_INGRESS/ingest"
      customheaders:
        Authorization: "Bearer SA_TOKEN"
      checkcert: true
```

### Placeholders

| Placeholder | Resolved | Description |
|-------------|----------|-------------|
| `SEED_INGRESS` | At reconcile time | Replaced with the seed's ingress domain |
| `SA_TOKEN` | At pod startup (init container) | Replaced with the `gardener-falcosidekick` SA token |

## Go Types

```go
type FalcoOutputConfig struct {
    Key   string
    Value map[string]interface{}
}

type Falco struct {
    ...
    BaseFalcoOutputConfigs []FalcoOutputConfig
}
```

```go
type Destination struct {
    Name               string
    // Enabled defaults to true. Set to false to opt out of a base output config.
    Enabled            *bool // new field
    ResourceSecretName *string
}
```

This is the exported version of the existing internal `falcoOutputConfig` struct in `pkg/values/falcovalues.go`.

## Override Behavior

| Scenario | Result |
|----------|--------|
| Destination not listed by shoot | Base config is applied (default) |
| Destination listed with `enabled: false` | Opt-out — base config is NOT applied |
| Destination listed with secret | Shoot's config overrides the base |
| Destination listed with no secret (enabled defaults to true) | Base config is applied |

## Example: Shoot gets operator-provided splunk (default, not listed)

```yaml
# Shoot spec — splunk not mentioned, base config applies automatically
extensions:
- type: shoot-falco-service
  providerConfig:
    destinations:
    - name: logging
```

## Example: Shoot overrides with own splunk

```yaml
# Shoot spec — provides own secret, overrides base config
extensions:
- type: shoot-falco-service
  providerConfig:
    destinations:
    - name: splunk
      resourceSecretName: my-splunk-secret
    - name: logging
```

## Example: Shoot opts out of operator-provided splunk

```yaml
# Shoot spec — explicitly disables splunk
extensions:
- type: shoot-falco-service
  providerConfig:
    destinations:
    - name: splunk
      enabled: false
    - name: logging
```

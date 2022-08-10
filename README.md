# Add Missing Headers
Reads headers from app service/client, adds missing headers, but not override existing headers, and forwards them to client/app service.

Traefik Plugin: [https://plugins.traefik.io/plugins/62f3496be2bf06d4675b9445/add-missing-headers](https://plugins.traefik.io/plugins/62f3496be2bf06d4675b9445/add-missing-headers)

GitHub: [https://github.com/jerrywoo96/AddMissingHeaders](https://github.com/jerrywoo96/AddMissingHeaders)

## Configuration

### Static (traefik.yml)
```yaml
experimental:
  plugins:
    AddMissingHeaders:
      moduleName: github.com/jerrywoo96/AddMissingHeaders
      version: v1.0.0
```

### Dynamic
```yaml
http:
  middlewares:
    AddMissingHeaders:
      plugin:
        AddMissingHeaders:
          requestHeaders:
            X-Custom-RequestHeader: CustomRequestHeader
          responseHeaders:
            X-Custom-ResponseHeader: CustomResponseHeader
```

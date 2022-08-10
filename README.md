# Add Missing Headers
Reads headers from app service/client, adds missing headers, but not override existing headers, and forwards them to client/app service.

<!---
Traefik Plugin: [https://plugins.traefik.io/plugins/62cfd4129279ff6d9dd027a9/add-forwarded-header](https://plugins.traefik.io/plugins/62cfd4129279ff6d9dd027a9/add-forwarded-header)
--->

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

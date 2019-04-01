# Default backend for nginx-ingress on Kubernetes

This is a custom default backend for [nginx-ingress](https://github.com/kubernetes/ingress-nginx). The only reason for creating this project was to customize the 404 page and provide a better (and branded) user experience.

## Configuration

By default, the backend will listen on port `9090` and serve an `HTML` page with the error file at `/assets/404.html`. Override the default html page by mounting a config map with the html page at a different path.

**Important Note**

The final image is based on `scratch` image, so normal file system operations won't work. See the `Dockerfile` for more details

## Public Builds

https://registry.hub.docker.com/u/nuvo/default-backend


Build from Source

```bash
docker build -t nuvo/default-backend .
```
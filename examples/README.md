# Imagekit-go SDK examples

This module contains code examples to demonstrate different features of the SDK. Each directory(except assets) is an application which groups related features. 
assets is a helper package used by different example applications.

## Step 1
In order to run any example, just set the following imagekit.io keys and endpoint in environment.

- IMAGEKIT_PRIVATE_KEY
- IMAGEKIT_PUBLIC_KEY
- IMAGEKIT_ENDPOINT_URL

```
# Set in .bashrc or .zshrc startup file depending your shell.

export IMAGEKIT_PRIVATE_KEY=<private_key>
export IMAGEKIT_PUBLIC_KEY=<public_key>
export IMAGEKIT_ENDPOINT_URL=<endpoing_url>
```

## Step 2
Run any example as illustrated below.
```
cd examples/upload
go run main.go
```

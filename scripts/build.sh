#!/bin/sh

cd ../cmd/sitemap-generator

# Linux build
GOOS=linux GOARCH=amd64 go build -o app
mv app ../../dist/sitemap-generator-linux-x86
echo "✅ Successfully created sitemap-generator-linux-x86!"

# macOS build
GOOS=darwin GOARCH=amd64 go build -o app
mv app ../../dist/sitemap-generator-darwin-x86
echo "✅ Successfully created sitemap-generator-darwin-x86!"

# FreeBSD build
GOOS=freebsd GOARCH=amd64 go build -o app
mv app ../../dist/sitemap-generator-freebsd-x86
echo "✅ Successfully created sitemap-generator-freebsd-x86!"

echo "✅ Successfully moved binaries to dist/ folder!"

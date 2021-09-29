# `sdk-service-meta`: a goa template for Cosmos SDK services

## How to generate goa examples from this repository:

 1. add this repository as a Go dependency
 2. `goa example github.com/allinbits/sdk-service-meta`\
 3. execute `find . -type f -name '*.go' -exec sed -i "s|github.com/allinbits/sdk-service-v{your version}/gen|github.com/allinbits/sdk-service-meta/gen|g" {} +`

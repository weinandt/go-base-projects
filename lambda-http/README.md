# Http Lambda
A lambda endpoint which can handle http requests. Written in golang.

Only three steps to have a full REST api:

### 1. Build
`GOOS=linux GOARCH=amd64 go build -o api main.go`

### 2. Run Terraform
1. `cd terraform`
1. `terraform init`
1. Fill out the `ohio.tfvars` file
1. `terraform apply -var-file="ohio.tfvars"`

## 3. Call the api
`curl -d '{"key1":"value1", "key2":"value2"}' -H "Content-Type: application/json" -X POST https://<endpointFromTerraformOutputHere>/echo`

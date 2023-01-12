# Http Lambda
A lambda endpoint which can handle http requests. Written in golang.


## Setting Up
### 1. Build
`GOOS=linux GOARCH=amd64 go build -o api main.go`

### 2. Run Terraform
1. `cd terraform`
1. `terraform init`
1. Fill out the `ohio.tfvars` file
1. `terraform apply -var-file="ohio.tfvars"`
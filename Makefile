build:
	@go build -o terraform-provider-url
	@terraform init

install:
	@go install

clean:
	@rm -rf ./terraform.tfstate
	@rm -rf ./terraform.tfstate.backup

apply: build
	@terraform apply -auto-approve

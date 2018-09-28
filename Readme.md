# Terraform URL Provider

Stupid simple URL parsing data source.

Meant for those times when you're given a full URL and just want to pass on a piece of it.

## Example

```hcl
provider url {}

data "url" "main" {
  url = "postgres://user:pass@test.com:9543/hi?cool=lol"
}

output "path" {
  value = "${data.url.main.path}"
}

output "query" {
  value = "${data.url.main.query.cool}"
}

output "port" {
  value = "${data.url.main.port}"
}

output "host" {
  value = "${data.url.main.host}"
}

output "hostname" {
  value = "${data.url.main.hostname}"
}

output "scheme" {
  value = "${data.url.main.scheme}"
}

output "username" {
  value = "${data.url.main.username}"
}

output "password" {
  value = "${data.url.main.password}"
}
```

## Installation

```sh
go get github.com/matthewmueller/terraform-provider-url
go install github.com/matthewmueller/terraform-provider-url
```

## License

MIT

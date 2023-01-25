project = "shrls"

config {
  env = {
    "MONGO_USERNAME" = dynamic("vault", {
      path = "database/creds/prod_rw"
      key = "username"
    })
    "MONGO_PASSWORD" = dynamic("vault", {
      path = "database/creds/prod_rw"
      key = "password"
    })
  }
}

app "shrls" {
  build {
    use "docker" {}
    registry {
      use "docker" {
        image = "registry.brittg.com/demophoon/shrls"
        tag = gitrefhash()
      }
    }
  }

  deploy {
    use "nomad" {
      service_port = 8000
    }
  }
}

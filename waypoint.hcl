project = "shrls"

config {
  #internal = {
  #  username = dynamic("vault", {
  #    path = "database/creds/prod_rw"
  #    key = "username"
  #  })
  #  password = dynamic("vault", {
  #    path = "database/creds/prod_rw"
  #    key = "password"
  #  })
  #}

  env = {
    #"MONGO_URI" = "mongodb://${config.internal.username}:${config.internal.password}@10.211.55.6:27017/shrls"
    "MONGO_URI" = "mongodb://root:changeMe@10.211.55.6:27017/shrls"
  }
}

app "shrls" {
  build {
    use "docker" {}
    registry {
      use "docker" {
        image = "registry.services.demophoon.com/demophoon/shrls-test"
        tag = join("-", [workspace.name, gitrefhash()])
      }
    }
  }

  deploy {
    #use "docker" {
    #  service_port = 8000
    #}
    use "kubernetes" {
      service_port = 8000
    }
    #use "nomad" {
    #  service_port = 8000
    #}
  }
}

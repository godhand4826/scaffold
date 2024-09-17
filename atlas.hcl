env "local" {
  migration {
    dir = "file://migrations"
  }
  url = "postgres://postgres_user:postgres_password@localhost:5432/scaffold?search_path=public&sslmode=disable"
  dev = "docker://postgres/15/dev?search_path=public"
}

env "docker" {
  migration {
    dir = "file://migrations"
  }
  url = "postgres://postgres_user:postgres_password@postgres:5432/scaffold?search_path=public&sslmode=disable"
}
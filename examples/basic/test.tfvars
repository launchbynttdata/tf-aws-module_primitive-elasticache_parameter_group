name   = "cache-params"
family = "redis2.8"
parameters = [
  {
    name  = "activerehashing",
    value = "yes"
    }, {
    name  = "min-slaves-to-write",
    value = "2"
  }
]

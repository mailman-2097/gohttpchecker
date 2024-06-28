# HTTP Header Check Go

Simple app to check http client headers

# Deploy

Build this as a openshift s2i project based on [examples](https://examples.openshift.pub/build/)

# Simple Instructions

```

$ oc new-app golang~https://github.com/mailman-2097/gohttpchecker.git
$ curl 10.1.0.2:8888
$ oc expose $(oc project -q)/gohttpchecker --port=8888
$ oc expose service/gohttpchecker

```


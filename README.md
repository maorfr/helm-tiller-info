# Helm Tiller information Plugin

This is a Helm plugin which prints information about Tiller.

## Usage

Print information about Tiller

```
$ helm tiller-info [flags]
```

### Flags:

```
  -l, --label string              label to select tiller resources by (default "NAME=TILLER")
      --tiller-namespace string   namespace of Tiller (default "kube-system")
```


## Install

```
$ helm plugin install https://github.com/maorfr/helm-tiller-info
```

The above will fetch the latest binary release of `helm tiller info` and install it.

### Developer (From Source) Install

If you would like to handle the build yourself, instead of fetching a binary,
this is how recommend doing it.

First, set up your environment:

- You need to have [Go](http://golang.org) installed. Make sure to set `$GOPATH`
- If you don't have [Dep](https://github.com/golang/dep) installed, this will install it into
  `$GOPATH/bin` for you.

Clone this repo into your `$GOPATH`. You can use `go get -d github.com/maorfr/helm-tiller-info`
for that.

```
$ cd $GOPATH/src/github.com/maorfr/helm-tiller-info
$ make bootstrap build
$ SKIP_BIN_INSTALL=1 helm plugin install $GOPATH/src/github.com/maorfr/helm-tiller-info
```

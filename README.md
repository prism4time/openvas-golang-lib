# openvas-golang-lib

A golang library to connect and manage the OpenVAS servers using the OMP protocol.Currently I develop it with OpenVAS 7.

## install 
to install and use the lib,use

```
go get -u github.com/qzyse2017/openvas-golang-lib
```

to get the latest code,you may find the commit ID on master branch and copy it to command line, use

```
go get -u github.com/qzyse2017/openvas-golang-lib@[latest-commit-id-here]

go mod tidy
```



## supported API
currently it only supports task and target related API. And I will update it occasionally according to my needs.

## examples
[a simple example](https://github.com/qzyse2017/openvas-golang-lib-example)

## Contributions
PRs welcome!!!!!!❤️❤️❤️❤️❤️❤️

## How to contribute
basically I generate most of the commands related structs from xml example in [openvas's omp docs](https://docs.greenbone.net/API/OMP/omp.html). The tool is [zek](https://github.com/miku/zek), and it has a handy [online version](https://www.onlinetool.io/xmltogo/) you could use.

and if you would like to add some new commands or fix some, I would appreciate it a lot.

following the conventions, you may need to extract some structs which show the meaningful things in the command and use its pointer as an argument to call the commands related function.e.g. ``Target`` and ``CreateTarget``

## License
MIT 

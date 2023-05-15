#!/bin/bash

set -x

# Master device name
ext_ifc=wlp4s0

setup () {
  docker network create -d macvlan --subnet=192.168.2.0/24 --gateway=192.168.2.1 -o parent=${ext_ifc} my_macvlan
}

teardown() {
  docker network rm my_macvlan
}

run() {
  docker run --network=my_macvlan -it anfernee/network-toolbox bash
  # Then run curl google.com
}

case "$1" in
  setup) setup
  ;;
  teardown) teardown
  ;;
esac
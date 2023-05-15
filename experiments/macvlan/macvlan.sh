#!/bin/bash

set -x

# Master device name
# You cannot use the main external interface. After enabling macvlan
# on the master interface, the host will lose connection (for wifi interface).
# Try this on an ethernet interface.
ext_ifc=enp3s0


ifc=mac

setup() {
  # Mode `bridge` means that mac0 can communicate with all other macvlan interfaces
  # on the same physical interface.
  sudo ip link add ${ifc}-1 link ${ext_ifc} type macvlan mode bridge
  sudo ip addr add 192.168.1.2/24 dev ${ifc}-1
  sudo ip link set ${ifc}-1 up

  sudo ip link add ${ifc}-2 link ${ext_ifc} type macvlan mode bridge
  sudo ip addr add 192.168.1.3/24 dev ${ifc}-2
  sudo ip link set ${ifc}-2 up
}

teardown() {
  sudo ip link del ${ifc}-1
  sudo ip link del ${ifc}-2
}

run() {
  ping -I 192.168.1.2 192.168.1.3
}

case "$1" in
  setup) setup
  ;;
  teardown) teardown
  ;;
esac

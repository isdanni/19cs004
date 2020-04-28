#!/bin/bash
# Steps to hard cleaup for topologies

ip netns | xargs -L1 sudo ip netns delete
find ~ -maxdepth 1 -name ".*.json"  -print0 | xargs -0 rm -rf
find ~ -maxdepth 1 -name ".*.json.lock"  -print0 | xargs -0 rm -rf

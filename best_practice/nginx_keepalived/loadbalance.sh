#!/usr/bin/env bash


for n in $(seq 1 1 10)
do
  nohup curl -H 'Host:lb.genox.wang' http://47.103.147.164:80/ping &>/dev/null &
done


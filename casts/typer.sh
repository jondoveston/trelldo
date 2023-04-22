#!/bin/bash

function typer
{
  local text="$1"
  local delay="${2:-0.1}"

  for i in $(seq 0 $(expr length "${text}")) ; do
    printf "${text:$i:1}"
    sleep ${delay}
  done
}

function enter
{
  printf "\x0a"
}

function escape
{
  printf "\x1b"
}

function controlc
{
  printf "\x03"
}

function return
{
  printf "\x0d"
}

function formfeed
{
  printf "\x0c"
}

sleep 1
typer "trelldo version 2>/dev/null"
enter
sleep 1

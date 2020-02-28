#!/bin/bash

PASS=true

go test

if [[ $? -eq 1 ]] ; then
  PASS=false
fi
if [[ $PASS == false ]] ; then
  printf "COMMIT FAILED\n"
  exit 1
else
  printf "COMMIT SUCCEEDED\n"
fi
exit 0
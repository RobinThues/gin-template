#!/bin/bash

go test

PASS=true

if [[ $? != 0 ]] ; then
  PASS=false
fi
if [[ "$PASS$" == true ]] ; then
  printf "COMMIT FAILED\n"
  exit 1
else
  printf "COMMIT SUCCEEDED\n"
fi
exit 0
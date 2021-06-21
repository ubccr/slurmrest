#!/bin/bash
  
OUT=.

java -jar $OPENAPI_CLI_JAR generate \
  -g go \
  -i scripts/slurm-openapi.json \
  -t scripts/openapi-templates/Go \
  -o $OUT \
  --package-name=slurmrest 

rm -f $OUT/.travis.yml
rm -f $OUT/git_push.sh

#!/bin/bash

rm -rf tmp
mkdir tmp

for f in `find /Users/Rick/data/dicom/compsamples_j2k/IMAGES -type f`
do
  echo $f
  g="tmp/"$(basename -- $f)".json"
  echo $g
  ./dcm2json -path $f | python -m json.tool > $g
done

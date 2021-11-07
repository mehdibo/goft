#!/usr/bin/env zsh
for i in `seq 1 100`
do
  goft requests get /projects/$i/projects > samplejson/project$i.project.json
done

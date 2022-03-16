#!/bin/bash

FILE= # input file
LINE=1
while read -r CURRENT_LINE
  do
    ./main send "$CURRENT_LINE" "$(cat new.txt)"
    ((LINE++))
done < "$FILE"

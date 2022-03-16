#!/bin/bash

FILE="" # list of emails input file
TEMPLATE="" # template file
LINE=1
while read -r CURRENT_LINE
  do
    email-cli send "$CURRENT_LINE" "$(cat $TEMPLATE)"
    ((LINE++))
done < "$FILE"

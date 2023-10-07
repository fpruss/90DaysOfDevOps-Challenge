#!/bin/bash

# Loop to create files with consecutive numbers
for ((i = 7; i <= 90; i++))
do
    filename=day$(printf %02d $i).md
    
    sed -i '0,/\(day[0-9]\{2\}\)/s//corresponding day/' $filename
done

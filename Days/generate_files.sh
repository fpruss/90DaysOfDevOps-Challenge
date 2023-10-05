#!/bin/bash
template_file=template.md
output_directory=./

script_dir=$(cd $(dirname $BASH_SOURCE) && pwd)
cd $script_dir

# Loop to create files with consecutive numbers
for ((i = 6; i <= 90; i++))  # Change '10' to the desired number of files
do
    # Generate the filename with leading zeros
    filename=day$(printf %02d $i).md
    
    # Copy the template file and replace the placeholder
    cp $template_file $output_directory/$filename
    sed -i s/dayX/day$(printf %02d $i)/g $filename
done

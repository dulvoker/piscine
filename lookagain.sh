#! /bin/bash
find . -name "*.sh" | sort -r | sed 's/.sh//g' | sed 's/^.*\///g'

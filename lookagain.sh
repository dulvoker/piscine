#! /bin/bash
find . -type f -name  '*sh' -print | tr -d '/' | tr -d '.' | sed 's/.$//' | sed 's/.$//'

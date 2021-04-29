#! /bin/bash
find . -type f -name  '*sh' | tr -d '/' | tr -d '.' | sed 's/.$//' | sed 's/.$//'

#! /bin/bash
find . -name '*sh' -print | tr -d '/' | tr -d '.' | sed 's/.$//' | sed 's/.$//'

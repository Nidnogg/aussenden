#!/bin/sh
# This shell script kills ports currently
kill -9 $(lsof -ti :8000) 
#!/bin/sh
# This shell script kills any process currently using port 8000
kill -9 $(lsof -ti :8000) || echo "Error: port is currently not in use"


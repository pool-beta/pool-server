#!/bin/bash

# Installation Script for POOL Server

# Check if GOPATH is set
echo "Checking GOPATH..."

if test -z "$GOPATH"
then
	# NOT FOUND
	echo "GOPATH not found -- setting GOPATH..."
	export GOPATH=~/go
fi

echo "GOPATH: " $GOPATH


# Install pool-server and pool-cli
echo "Installing..."

go install ./...

# Check error
if [ $? -ne 0 ]; then
	return $?
fi

echo "Done."


# Check if GOPATH in PATH
if [[ "$PATH" != *"$GOPATH"* ]] 
then
	echo "Adding GOPATH to PATH"
	export PATH=$PATH:$GOPATH
fi


# Output Instructions for use
echo ""
echo ""
echo "POOL has been installed successfully. Please follow the instructions below to get it running!

1) Start running the POOL Server: pool-server

2) Interact using the POOL Client: pool-cli -help
"

echo ""


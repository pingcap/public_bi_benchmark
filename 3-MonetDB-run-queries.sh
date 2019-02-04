#!/bin/bash

SCRIPT_DIR=$(dirname "$(realpath "$0")")
WORKING_DIR=$(pwd)

usage() {
cat <<EOM
Usage: $(basename $0) <dataset-dir>
  dataset-dir    path to dataset root directory
EOM
}

if [ "$#" -ne 1 ]; then
    usage
	exit 1
fi
DATASET_DIR="$1"

mclient --version
ret=$?
if [ "$ret" -ne 0 ]; then
	echo "error: unable to connect to MonetDB"
	exit $ret
fi

cd $DATASET_DIR
ret=$?
if [ $ret -ne 0 ]; then
	exit $ret
fi

$SCRIPT_DIR/lib/11-MonetDB-run-queries.sh

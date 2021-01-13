#!/bin/sh

UMASK_SET=${UMASK_SET:-022}
umask "$UMASK_SET"

./filer

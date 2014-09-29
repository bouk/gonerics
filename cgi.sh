#!/bin/sh
set -e

GITDIR="/home/bouke/git$GIT_REPO_PATH"
DIR=`pwd`

if [ ! -d $GITDIR ]; then
    TDIR=`mktemp -d`
    cd $TDIR >/dev/null 2>/dev/null
    git init --bare >/dev/null 2>/dev/null

    TDIR2=`mktemp -d`
    cd $TDIR2 >/dev/null 2>/dev/null
    git clone $TDIR . >/dev/null 2>/dev/null
    /home/bouke/gonerics/gonerics --template-type="$TEMPLATE_TYPE" --name="$NAME" --parameters="$PARAMETERS" > file.go 2>>/home/bouke/git/gonerics.log
    git add file.go >/dev/null 2>/dev/null
    git commit -m "This repo is $NAME with params $PARAMETERS" >/dev/null 2>/dev/null
    git push origin master >/dev/null 2>/dev/null
    rm -rf $TDIR2 >/dev/null 2>/dev/null

    mkdir -p `dirname $GITDIR` >/dev/null 2>/dev/null
    mv -T $TDIR $GITDIR >/dev/null 2>/dev/null
fi

cd $DIR
/usr/lib/git-core/git-http-backend

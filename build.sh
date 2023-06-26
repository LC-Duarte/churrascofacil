#!/bin/bash
# create swagger docs $swag init  --output internal/docs --generalInfo cmd/mygoapp/main.go

VERSION="v0.1"
TAG="leocd123/churrascofacil:$VERSION"
LABEL="churrascoFacil"
APP="cfserver"
PORT=8080

#ENSURE BIN exists
[ ! -d "bin" ] && mkdir bin || echo "bin already exists"
[ ! -d "bin/$VERSION" ] && mkdir "bin/$VERSION" ||  echo "bin/$VERSION already exists"
#No options provided
if [ $# -eq 0 ]
then
    echo "Running docker build"
    docker build . --label $LABEL -t $TAG
    exit 0
fi
LOCAL="true" #Set To true for convinece. #replace to false before pushing
while getopts "hlr" OPTION; do
        case $OPTION in
                #local build
                l)
                    echo "Runing local build"
                    LOCAL="true"
                    mkdir -p bin #Ensure bin dir exists
                    go build -o bin/$VERSION/$APP cmd/$APP/main.go
                    ;;

                h)
                    echo " `basename $0`     #Build docker image"
                    echo " `basename $0` -r  #Build and  docker image "
                    echo " `basename $0` -l  #copile with go without docker image"
                    echo " `basename $0` -lr #copile with go without docker image and run "
                    exit 0
                    ;;
                #run server    
                r)
                    if [ $LOCAL = "true" ];then
                        echo "Starting server locally at port $PORT"
                        cd bin/$VERSION/
                        ./$APP
                    elif [ $LOCAL = "false" ];then
                        echo "Running docker build"
                        docker build . --label LABEL -t $TAG
                        echo "Starting server on docker container using local port $PORT"
                        docker run -p $PORT:$PORT $TAG
                    fi
                    ;;
        esac
done
  


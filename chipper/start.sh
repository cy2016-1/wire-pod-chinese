#!/bin/bash

if [[ $EUID -ne 0 ]]; then
  echo "This script must be run as root. sudo ./start.sh"
  exit 1
fi

if [[ -d ./chipper ]]; then
   cd chipper
fi

#if [[ ! -f ./chipper ]]; then
#   if [[ -f ./go.mod ]]; then
#     echo "You need to build chipper first. This can be done with the setup.sh script."
#   else
#     echo "You must be in the chipper directory."
#   fi
#   exit 0
#fi

if [[ ! -f ./source.sh ]]; then
  echo "You need to make a source.sh file. This can be done with the setup.sh script."
  exit 0
fi

######luowei add

#杀死占用8084端口的程序 luowei add
if sudo lsof -t -i:8084 | grep . >/dev/null; then
  sudo kill -9 $(sudo lsof -t -i:8084)
fi

export BAIDU_APIKEY="9EK7fILTS7TLilYsCGLDi7mQ"
export BADIU_APISECRET="rG9sj6Cwwx1LxZe3SbWDfyncWfIQr1b2"

export XUNFEI_APPID="9b14f37d"     #填写控制台中获取的 APPID 信息
export XUNFEI_APISECRET="ZGQ3MDlhNWEwNGUzNjliYzViZTBhOTU0"   #填写控制台中获取的 APISecret 信息
export XUNFEI_APIKEY="a7e66bd53e2de0d8922056c32d1cd280"    #填写控制台中获取的 APIKey 信息


export PYTHON_CMD="/home/luowei/anaconda3/bin/python"
export PYTHON_CHDIR="/home/luowei/space/vector/wire-pod/python"
export WIREPOD_HOME="/home/luowei/space/vector/wire-pod" #用于初始化sdk_wrapper
######


source source.sh

#./chipper
if [[ ${STT_SERVICE} == "leopard" ]]; then
	if [[ -f ./chipper ]]; then
	  ./chipper
	else
	  /usr/local/go/bin/go run cmd/leopard/main.go
	fi
elif [[ ${STT_SERVICE} == "rhino" ]]; then
	if [[ -f ./chipper ]]; then
          ./chipper
        else
          /usr/local/go/bin/go run cmd/experimental/rhino/main.go
        fi
elif [[ ${STT_SERVICE} == "houndify" ]]; then
	if [[ -f ./chipper ]]; then
          ./chipper
        else
          /usr/local/go/bin/go run cmd/experimental/houndify/main.go
        fi
elif [[ ${STT_SERVICE} == "whisper" ]]; then
        if [[ -f ./chipper ]]; then
          ./chipper
        else
          /usr/local/go/bin/go run cmd/experimental/whisper/main.go
        fi
elif [[ ${STT_SERVICE} == "baidu" ]]; then #加上baidu语音识别的选项 luowei add
        if [[ -f ./chipper ]]; then
          ./chipper
        else
          echo "/usr/local/go/bin/go run cmd/experimental/baidu/main.go"
          /usr/local/go/bin/go run cmd/experimental/baidu/main.go
        fi
elif [[ ${STT_SERVICE} == "vosk" ]]; then
	if [[ -f ./chipper ]]; then
		export CGO_ENABLED=1 
		export CGO_CFLAGS="-I/root/.vosk/libvosk"
		export CGO_LDFLAGS="-L /root/.vosk/libvosk -lvosk -ldl -lpthread"
		export LD_LIBRARY_PATH="/root/.vosk/libvosk:$LD_LIBRARY_PATH"
		./chipper
	else
		export CGO_ENABLED=1 
		export CGO_CFLAGS="-I$HOME/.vosk/libvosk"
		export CGO_LDFLAGS="-L $HOME/.vosk/libvosk -lvosk -ldl -lpthread"
		export LD_LIBRARY_PATH="$HOME/.vosk/libvosk:$LD_LIBRARY_PATH"
		/usr/local/go/bin/go run cmd/vosk/main.go
	fi
else
if [[ -f ./chipper ]]; then
      export CGO_LDFLAGS="-L/root/.coqui/"
      export CGO_CXXFLAGS="-I/root/.coqui/"
      export LD_LIBRARY_PATH="/root/.coqui/:$LD_LIBRARY_PATH"
  ./chipper
else
      export CGO_LDFLAGS="-L$HOME/.coqui/"
      export CGO_CXXFLAGS="-I$HOME/.coqui/"
      export LD_LIBRARY_PATH="$HOME/.coqui/:$LD_LIBRARY_PATH"
      /usr/local/go/bin/go run cmd/coqui/main.go
fi
fi

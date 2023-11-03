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
#if sudo lsof -t -i:8084 | grep . >/dev/null; then
#  sudo kill -9 $(sudo lsof -t -i:8084)
#fi

export BAIDU_APIKEY="xxx"
export BADIU_APISECRET="xxx"

export XUNFEI_APPID="xxx"
export XUNFEI_APISECRET="xxx"
export XUNFEI_APIKEY="xxx"

export PYTHON_CMD="python3"
export PYTHON_CHDIR="xxx/wire-pod-chinese/python"

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

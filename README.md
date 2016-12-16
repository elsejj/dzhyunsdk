# 背景

对于大智慧金融云服务，使用 HTTP+JSON 的访问组合最方便使用，而 Websocket+Protobuf的访问组合则提供最好的性能（大大降低的流量以及长链接，但 Websocket+Protobuf 具有一定的编程难度，所以提供本 SDK。

本 SDK 会创建一个到大智慧金融云服务的 Websocket 长链接，通过 Protobuf 协议进行交互，但在本地会创建一个 Websocket 代理，原生应用可以连接此代理，以 JSON 的方式进行数据读取。

`APP <--- Websocket+JSON ---> SDK <--- Websocket+Protobuf ---> RemoteServer`

在 SDK 中，同时对 YFloat 数据类型进行了处理，YFloat 是一种针对金融应用的浮点数的整数重编码，它具有足够的精度，并对数据传输进行了优化。


# API

SDK 已被简化为一个函数，传入远端服务地址和本地监听地址，对本地监听的请求将代理到远端服务

`StartSDK("remote url", "local listen address")`

# 实现

SDK 使用 `Go` 编写，通过 [gomobile](https://github.com/golang/mobile), [Go 交叉编译](http://golangcookbook.com/chapters/running/cross-compiling/) 等机制提供不同平台的 SDK

## Android

产生 Android 可用的 aar 包

`make android`

output/dzhyunsdk.aar 即为产生的 aar 包，可以用于导入到各 Android 工程

```java
import com.dzhyun.dzhyunsdk.Dzhyunsdk;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        Dzhyunsdk.startSDK("ws://gw.yundzh.com/ws?token=00000003:1481705008:48ad77788e2e588c9eab9c52618173d51a600abe", "127.0.0.1:9999");
    }

```

对 `ws://127.0.0.1:9999/ws` 进行连接后，即可正常的发送请求和接受回复。 

## iOS

产生 xcode framework 

`make ios`

output/Dzhyunsdk.framework 即为产生的 framwork ，将其加入到已有的工程中

```objc
#include "Dzhyunsdk/Dzhyunsdk.h"

- (BOOL)application:(UIApplication *)application didFinishLaunchingWithOptions:(NSDictionary *)launchOptions {
    
    DzhyunDzhyunsdkStartSDK(@"ws://gw.yundzh.com/ws?token=00000003:1481705008:48ad77788e2e588c9eab9c52618173d51a600abe", @"127.0.0.1:9999");
    
    return YES;
}
```

对 `ws://127.0.0.1:9999/ws` 进行连接后，即可正常的发送请求和接受回复。 

# PC平台

产生不同平台下的可执行文件

`make exe`

output/dzhyunsdk.exe 64bit windows 下的可执行文件

output/dzhyunsdk.linux 64bit linux 下的可执行文

output/dzhyunsdk.mac 64bit macos 下的可执行文件件

可执行文件使用说明

```
 $ output/dzhyunsdk.mac -h
Usage of output/dzhyunsdk:
  -local string
    	the local address (default "127.0.0.1:9999")
  -server string
    	the server url (default "ws://gw.yundzh.com/ws")
  -token string
    	the server access token
```

启动后，即可通过 `ws://127.0.0.1:9999/ws` 进行访问

# HTTP代理

`SDK` 也同时对 HTTP 请求进行代理，所以在某些需要使用HTTP的情况下，可以避免HTTP到远程的性能损耗，无论与`SDK`之间有多少条连接，实际上到远程的链接只有一条，并且在与远程的通讯中，启用了压缩和`Protobuf`，例如：

`curl "http://127.0.0.1:9999/stkdata?obj=SH600000&field=ZhongWenJianCheng,ZuiXinJia"`

将与直接访问远程的相应请求一样返回

```json
{
  "Qid": "",
  "Counter": 1,
  "Err": 0,
  "Data": {
  "Id": 27,
  "ObjCount": 1,
  "RepDataStkData": [
    {
      "Obj": "SH600000",
      "ZhongWenJianCheng": "浦发银行",
      "ZuiXinJia": 16.69,
      "XuHao": 0
    }
  ]
  }
}
```

同时，HTTP也被用于对SDK的控制

## /rl

`curl "http://127.0.0.1:9999/rl"`

指示 `SDK`重新链接远程服务，这特别在移动端网络切换时很有必要调用

## /rs

`curl "http://127.0.0.1:9999/rs"`

返回当前远端链接的状态，如

```json
{
  "Online": true,
  "SendBytes": 96,
  "RecvBytes": 54
}
```


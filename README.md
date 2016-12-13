# 背景

对于大智慧金融云服务，使用 HTTP+JSON 的访问组合最方便使用，而 Websocket+Protobuf的访问组合则提供最好的性能（大大降低的流量以及长链接，但 Websocket+Protobuf 具有一定的编程难度，所以提供本 SDK。

本 SDK 会创建一个到大智慧金融云服务的 Websocket 长链接，通过 Protobuf 协议进行交互，但在本地会创建一个 Websocket 代理，原生应用可以连接此代理，以 JSON 的方式进行数据读取。

`APP <--- Websocket+JSON ---> SDK <--- Websocket+Protobuf ---> RemoteServer`

在 SDK 中，同时对 YFloat 数据类型进行了处理，YFloat 是一种针对金融应用的浮点数的整数重编码，它具有足够的精度，并对数据传输进行了优化。


# API

SDK 已被简化为一个函数，传入远端服务地址和本地监听地址，对本地监听的请求将代理到远端服务

`StartSDK("remote url", "local listen address")`

# 实现

SDK 使用 `Go` 编写，通过 `gomobile`, `go shared library` 等机制提供不同平台的 SDK

## Android

产生 Android 可用的 aar 包

`make android`

`output/dzhyunsdk.aar` 即为产生的 aar 包，可以用于导入到各 Android 工程

```java
import com.dzhyun.dzhyunsdk.Dzhyunsdk;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        Dzhyunsdk.startSDK("ws://gw.yundzh.com/ws?token=00000003:1481705008:48ad77788e2e588c9eab9c52618173d51a600abe", "127.0.0.1:9999");
    }

```

对 `ws://127.0.0.1:9999/ws` 进行连接后，即可正常的发送请求和接受回复。 
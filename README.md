# vei-driver-sdk-go
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/volcengine/vei-driver-sdk-go/go.yml?branch=main&logo=github)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/volcengine/vei-driver-sdk-go?logo=go)
![GitHub License](https://img.shields.io/github/license/volcengine/vei-driver-sdk-go)
![GitHub Release](https://img.shields.io/github/v/release/volcengine/vei-driver-sdk-go)

欢迎使用边缘智能驱动开发SDK for Golang，本文档为您介绍如何开发一个自定义驱动

[边缘智能产品主页](https://www.volcengine.com/product/vei/mainpage)
## 驱动开发步骤

### 1. 程序入口
```go
package main

func main() {
    d := &driver.MinimalDriver{}
    vei.Bootstrap("device-minimal", "0.0.1", d)
}
```

### 1. 必备接口
```go
type MinimalDriver struct {
    logger  logger.Logger
    asyncCh chan<- *contracts.AsyncValues
}

/**
 *  初始化接口: 程序启动时被调用一次
 *  @param lc:              logger客户端
 *  @param asyncCh:         原生的异步数据上报通道
 */
func (m *MinimalDriver) Initialize(logger logger.Logger, asyncCh chan<- *contracts.AsyncValues) error {
    m.logger = logger
    m.asyncCh = asyncCh
    // 其他必要的初始化过程
    ...
    return nil
}

/**
 *  数据读取接口: 根据设置的采样周期定时调用，或通过在线调试主动调用
 *  @param device:          设备，包含了设备名称和通信协议内容
 *  @param reqs:            请求列表，每个请求中包含属性标识符和点表定义
 */
func (m *MinimalDriver) ReadProperty(device *contracts.Device, reqs []contracts.ReadRequest) error {
	// 通过 device.Name 或 device.Protocols 和设备建立连接，如果连接失败，可手动更新设备状态。
    if err := createConnectionWith(device);err != nil{
        device.SetOperatingState(state, reason)
    }
	
    for _, req := range reqs {
        // 根据请求中携带的模块、属性名、类型等参数访问设备获取数据
        value, err := readProperty(device, req)
        ...
		
        if err != nil {
            // 返回读取的结果
            req.SetResult(contracts.NewSimpleResult(value))
        } else {
            // 或者返回失败内容
            req.Failed(err)			
        }
    }
	return nil
}

/**
 *  数据写入接口: 通过在线调试主动调用
 *  @param device:          设备，包含了设备名称和通信协议内容
 *  @param reqs:            请求列表，每个请求中包含属性标识符和点表定义，以及写入的参数
 */
func (m *MinimalDriver) WriteProperty(device *contracts.Device, reqs []contracts.WriteRequest) error {
    for idx, req := range reqs {
        param := req.Param()
        // 根据 req 和 param 向设备写入数据
        ...
    }
    return nil
}

/**
 *  服务调用接口: 通过在线调试主动调用
 *  @param device:          设备，包含了设备名称和通信协议内容
 *  @param reqs:            请求列表，每个请求中包含了服务的标识符和相关定义，以及调用的参数
 */
func (m *MinimalDriver)  CallService(device *contracts.Device, reqs []contracts.CallRequest) error {
    type Request struct {
        X float32 `json:"x"`
        Y float32 `json:"y"`
    }
    
    type Response struct {
        Result float32 `json:"result"`
    }
    
    request, response := &Request{}, &Response{}
    if err := json.Unmarshal(req.Payload(), request); err != nil {
        return nil, err
    }
    
    // 根据请求内容调用对应的Service
    ...
    
    // 返回调用结果
    req.SetResult(contracts.NewSimpleResult(response))

    return nil
}

/**
 *  停止运行接口: 程序关闭前被调用一次
 *  @param force:           是否强制停止
 */
func (m *MinimalDriver) Stop(force bool) error {
    return nil
}
```

### 2. 系统事件接口
```go
/**
 *  添加设备回调函数: 成功添加设备时调用一次
 *  @param device:      设备实例
 */
func (m *MinimalDriver) AddDevice(device *contracts.Device) error error {}
/**
 *  更新设备回调函数: 成功更新设备时调用一次
 *  @param device:      设备实例
 */
func (m *MinimalDriver) UpdateDevice(device *contracts.Device) error {}
/**
 *  删除设备回调函数: 成功删除设备时调用一次
 *  @param device:      设备实例
 */
func (m *MinimalDriver) RemoveDevice(device *contracts.Device) error {}
```

## SDK功能规划
- [x] 事件上报
- [x] 设备发现
- [ ] 点表在线调试功能


## 许可证
[Apache-2.0 License](LICENSE).
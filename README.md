# vei-driver-sdk-go
欢迎使用边缘智能驱动开发SDK for Golang，本文档为您介绍如何开发一个自定义驱动

[边缘智能产品控制台](https://console.volcengine.com/vei)
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
    lc       logger.LoggingClient
    asyncCh  chan<- *sdkmodels.AsyncValues
    reporter interfaces.EventReporter
}

/**
 *  初始化接口: 程序启动时被调用一次
 *  @param lc:              logger客户端
 *  @param asyncCh:         原生的异步数据上报通道
 *  @param eventReporter:   事件上报函数(驱动主动调用)，
 */
func (m *MinimalDriver) Initialize(lc logger.LoggingClient, asyncCh chan<- *sdkmodels.AsyncValues, eventReporter interfaces.EventReporter) error {
    m.lc = lc
    m.asyncCh = asyncCh
    m.reporter = eventReporter
    // 其他必要的初始化过程
    ...
    return nil
}

/**
 *  数据读取接口: 根据设置的采样周期定时调用，或通过在线调试主动调用
 *  @param deviceName:      设备名称
 *  @param protocols:       设备的协议参数
 *  @param reqs:            请求列表，每个请求中包含属性标识符和点表定义
 */
func (m *MinimalDriver) HandleReadCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []sdkmodels.CommandRequest) ([]*sdkmodels.CommandValue, error) {
    result := make([]*sdkmodels.CommandValue, len(reqs))
    for i, req := range reqs {
        var cv *sdkmodels.CommandValue
        // 根据 req 中参数信息访问设备获取数据
        ...
        result[i] = cv

    }
    return result, nil
}

/**
 *  数据写入接口: 通过在线调试主动调用
 *  @param deviceName:      设备名称
 *  @param protocols:       设备的协议参数
 *  @param reqs:            请求列表，每个请求中包含属性标识符和点表定义
 *  @param params:          参数列表，与reqs长度相等且一一对应，每个参数包含了参数类型和参数值
 */
func (m *MinimalDriver) HandleWriteCommands(deviceName string, protocols map[string]models.ProtocolProperties, reqs []sdkmodels.CommandRequest, params []*sdkmodels.CommandValue) error {
    for idx, req := range reqs {
        param := params[idx]
        // 根据 req 和 param 向设备写入数据
        ...
    }
    return nil
}

/**
 *  服务调用接口: 通过在线调试主动调用
 *  @param deviceName:      设备名称
 *  @param protocols:       设备的协议参数
 *  @param req:             服务请求，包含服务标识符
 *  @param data:            服务参数，已完成解码，JSON格式
 */
func (m *MinimalDriver) HandleServiceCall(deviceName string, protocols map[string]models.ProtocolProperties, req sdkmodels.CommandRequest, data []byte) (*sdkmodels.CommandValue, error) {
    type Request struct {
        X float32 `json:"x"`
        Y float32 `json:"y"`
    }
    
    type Response struct {
        Result float32 `json:"result"`
    }
    
    request, response := &Request{}, &Response{}
    if err := json.Unmarshal(data, request); err != nil {
        return nil, err
    }
    
    // 根据请求内容调用对应的Service
    ...
    
    return sdkmodels.NewCommandValue(req.DeviceResourceName, common.ValueTypeObject, response)
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
 *  @param deviceName:      设备名称
 *  @param protocols:       设备的协议参数
 */
func (m *MinimalDriver) AddDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {}
/**
 *  更新设备回调函数: 成功更新设备时调用一次
 *  @param deviceName:      设备名称
 *  @param protocols:       设备的协议参数
 */
func (m *MinimalDriver) UpdateDevice(deviceName string, protocols map[string]models.ProtocolProperties, adminState models.AdminState) error {}
/**
 *  删除设备回调函数: 成功删除设备时调用一次
 *  @param deviceName:      设备名称
 *  @param protocols:       设备的协议参数
 */
func (m *MinimalDriver) RemoveDevice(deviceName string, protocols map[string]models.ProtocolProperties) error {}
```

## SDK功能规划
- [ ] 事件上报
- [ ] 设备自发现
- [ ] 点表在线调试功能


## 许可证
[Apache-2.0 License](LICENSE).
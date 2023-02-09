# BkAudit SDK

以下为 BkAudit SDK 使用说明文档，所有代码均有完整注释，可以在源码中查看。

## 类型说明

### 审计事件 AuditEvent

AuditEvent 定义了审计事件的标准字段，我们期望的审计事件由这些标准字段组成，详见 [bkaudit.AuditEvent](../bkaudit/models.go)

### 操作 Action

Action 实际上为 iam 模型中的 Action 对象，详见 [iam.model.models.Action](https://github.com/TencentBlueKing/iam-python-sdk/blob/master/iam/model/models.py)

### 资源类型 ResourceType

ResourceType 实际上为 iam 模型中的 ResourceType 对象，详见 [iam.model.models.ResourceType](https://github.com/TencentBlueKing/iam-python-sdk/blob/master/iam/model/models.py)

### 审计对象实例 AuditInstance

AuditInstance 定义了审计中传递实例需要的必须参数，会映射对应的字段到审计日志中，详见 [bkaudit.AuditInstance](../bkaudit/models.go)

### 审计上下文 AuditContext

上下文中定义了一个请求中常用且固定的参数，例如用户名，访问来源IP等信息，详见 [bkaudit.AuditContext](../bkaudit/models.go)

### 审计事件处理 Formatter

Formatter 用于将非标准化的参数或字段，转换为 AuditEvent 对象，详见 [bkaudit.Formatter](../bkaudit/formatters.go)

### 审计事件输出 Exporter

Exporter 用于输出审计事件，如输出到日志文件，OT Collector 等，详见 [bkaudit.Exporter](../bkaudit/exporters.go)

## 基本使用

以下以代码示例的方式进行说明，逐步完成整个SDK的配置与自定义，也可以参考 [example](../example/main.go) 的使用

### 初始化 Client

```
var Client *bkaudit.EventClient

func initClient() {
    var formatter = &bkaudit.Formatter{}
    var exporters = []bkaudit.BaseExporter{&bkaudit.Exporter{}}
    var err error
    Client, err = bkaudit.InitEventClient("BkAppCode", "BkAppSecret", formatter, exporters, 0, nil)
    if err != nil {
        fmt.Println("init client failed")
        return
    }
}
```

### 构造 Action 与 Resource

需要提前定义 Action 与 Resource 用于审计事件的生成

```
var (
	ViewHost = bkaudit.AuditAction{ActionID: "view-host"}
	Host     = bkaudit.AuditResource{ResourceTypeID: "host"}
)
```

### 构造 AuditContext

```
// "admin" 需要为实际的操作人
context := bkaudit.AuditContext{Username: "admin"}
```

### 构造实例 AuditInstance

```
instance := bkaudit.AuditInstance{
    InstanceID:          "host_01",
    InstanceName:        "主机一号",
    InstanceSensitivity: 0,
    InstanceData:        map[string]any{"ip": "127.0.0.1"},
    InstanceOriginData:  map[string]any{"ip": "127.0.0.2"},
}
```

### 输出 AuditEvent

```
# 调用 client 的 AddEvent 方法添加审计事件
# AddEvent 函数的大部分参数都具有默认值，如果无法满足审计需求，请传入必要内容
Client.AddEvent(&ViewHost, &Host, &instance, &context, "", "", 0, 0, 0, "", map[string]any{})
```

## 使用进阶

### 自定义 Formatter

在参数传入后，会调用 Formatter 进行格式化与参数提取，得到一个 AuditEvent   
为了简便处理，对于一些在请求中比较固定的内容，可以直接传递给 AuditContext，并在 Formatter 中处理   
例如，在 Gin 项目中，可以注入 gin.Context 对象用于获取 IP/UA 等内容

```
type GinFormatter struct{}

func (f *GinFormatter) Format(
	action *bkaudit.AuditAction,
	resourceType *bkaudit.AuditResource,
	instance *bkaudit.AuditInstance,
	auditContext *bkaudit.AuditContext,
	eventID string,
	eventContent string,
	startTime int64,
	endTime int64,
	resultCode int64,
	resultContent string,
	extendData map[string]any,
) (auditEvent *bkaudit.AuditEvent, err error) {
	// get context
	context := auditContext.ExtraData["context"].(*gin.Context)
	auditContext.AccessSourceIp = context.ClientIP()
	auditContext.AccessUserAgent = context.Request.UserAgent()
	// format
	return (&bkaudit.Formatter{}).Format(
		action,
		resourceType,
		instance,
		auditContext,
		eventID,
		eventContent,
		startTime,
		endTime,
		resultCode,
		resultContent,
		extendData,
	)
}
```

在定义好了 Formatter 后，需要在初始化 Client 时指定

```
Client, err = bkaudit.InitEventClient("BkAppCode", "BkAppSecret", &GinFormatter{}, exporters, nil)
```

此时我们在传递上下文时，就不需要传递 audit_context 中的所有参数，而是传递 Request 对象，在 Formatter 中进行提取补全

```
client.AddEvent(
	&ViewHome,
	&bkaudit.AuditResource{},
	&bkaudit.AuditInstance{},
	&bkaudit.AuditContext{Username: "admin", ExtraData: map[string]any{"context": context}},
	"",
	"",
	0,
	0,
	0,
	"",
	map[string]any{},
)
```

### 自定义 Exporter

审计事件默认会使用 log 输出，也可以定义自己的输出 Exporter

```
type Exporter struct{}

func (e *Exporter) Export(queue BaseQueue) {
	for {
		// read from channel
		event, ok := <-queue
		if ok {
			// get string data
			data, err := event.String()
			if err != nil {
				log.Error("export event failed: ", err)
				return
			}
			// Directly Export to Log
			log.Info(data)
		}
	}
}
```

在定义好了 Exporter 后，需要在初始化 Client 时指定

```python
Client, err = bkaudit.InitEventClient("BkAppCode", "BkAppSecret", formatter, []bkaudit.BaseExporter{&FileExport{}}, nil)
```

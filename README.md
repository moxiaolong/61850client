根据项目 https://github.com/beanit/iec61850bean 改写。

实现了Client部分。

建立连接
```go
clientSap := src.NewClientSap()
association := clientSap.Associate(hostName, port, src.NewEventListener())
```
接受SCL模型
```go
serverModel := association.RetrieveModel()
```
请求数据
```go
fcModelNode := serverModel.AskForFcModelNode("ied1lDevice1/MMXU1.TotW.mag.f", "MX")
association.GetDataValues(fcModelNode)
fcNodeBasic := fcModelNode.(src.BasicDataAttributeI)
println(fcNodeBasic.GetValueString())
```
写入数据
```go
fcModelNode := serverModel.AskForFcModelNode("ied1lDevice1/LLN0.NamPlt.vendor", "DC")
fcModelNode.(*src.BdaVisibleString).SetValue("abc")
association.SetDataValues(fcModelNode)
```

	
根据项目 https://github.com/beanit/iec61850bean 改写。

实现了Client部分。

建立连接
```golang
clientSap := src.NewClientSap()
association := clientSap.Associate(hostName, port, src.NewEventListener())
```
接受SCL模型
```golang
serverModel := association.RetrieveModel()
```
请求数据
```golang
fcModelNode := serverModel.AskForFcModelNode("ied1lDevice1/MMXU1.TotW.mag.f", "MX")
association.GetDataValues(fcModelNode)
fcNodeBasic := fcModelNode.(src.BasicDataAttributeI)
println(fcNodeBasic.GetValueString())
```
写入数据
```
TODO
```
	
##应用api
请求url：api_v1.xxxx.xxx/app
###添加应用 
http请求类型：POST

是否需要token：是（token存入http head Authorization key）

请求body：proto类型 -> App

```
string name =2; 应用名字
string region =3; 引用所属区域代码 默认“shanghai”
repeated  Pod Pods = 4 ; 应用中的实例
repeated  Service service = 5 ; 应用中的网络服务
repeated Volume volumes = 6; 应用所挂载的服务器空间
Runtime runtime = 7; 应用的运行状态
string timelyState =8; 应用动态信息 
map<string,string> conf =9;应用配置文件
repeated string editable =10; 可修改的值
```

返回数据：proto类型 -> Event

```
string id =1; 事件流水id
int32 code =2; 事件返回状态码
string msg =3; 事件返回消息
string region =4; 云端区域
string content = 5; 返回完成应用创建的id
```

###替换应用

http请求类型：PUT

是否需要token：是（token存入http head Authorization key）

请求body：proto类型 -> App

```
string id =1; 应用id
string name =2; 应用名字
string region =3; 引用所属区域代码 默认“shanghai”
repeated  Pod Pods = 4 ; 应用中的实例
repeated  Service service = 5 ; 应用中的网络服务
repeated Volume volumes = 6; 应用所挂载的服务器空间
Runtime runtime = 7; 应用的运行状态
string timelyState =8; 应用动态信息 
map<string,string> conf =9;应用配置文件
repeated string editable =10; 可修改的值
```

返回数据：proto类型 -> Event

```
string id =1; 事件流水id
int32 code =2; 事件返回状态码
string msg =3; 事件返回消息
string region =4; 云端区域
sting content = 5; 返回替换完成应用的id
```

###更新应用状态

http请求类型：Patch

是否需要token：是（token存入http head Authorization key）

请求body：proto类型 -> App

```
string id =1; 应用id
string name =2; 应用名字
string region =3; 引用所属区域代码 默认“shanghai”
Runtime runtime = 7; 应用的运行状态
string timelyState =8; 应用动态信息 
```

返回数据：proto类型 -> Event

```
string id =1; 事件流水id
int32 code =2; 事件返回状态码
string msg =3; 事件返回消息
string region =4; 云端区域
string content = 5; 返回更新完成应用的id
```

###删除应用

http请求类型：DELETE

是否需要token：是（token存入http head Authorization key）

请求body：proto类型 -> DeleteParams

```
string id =1; 应用id
string type =2; 值->app
string region =3; 云端区域
```

返回数据：proto类型 -> Event

```
string id =1; 事件流水id
int32 code =2; 事件返回状态码
string msg =3; 事件返回消息
string region =4; 云端区域
string content = 5; 返回删除完成应用的id
```

###获取单个应用

http请求类型：GET

是否需要token：是（token存入http head Authorization key）

请求body：

```
string id =1; 应用id
string type =2; 值->value
string region =3; 云端区域
```

返回数据：proto类型 -> App

```
string id =1; 应用id
string name =2; 应用名字
string region =3; 引用所属区域代码 默认“shanghai”
repeated  Pod Pods = 4 ; 应用中的实例
repeated  Service service = 5 ; 应用中的网络服务
repeated Volume volumes = 6; 应用所挂载的服务器空间
Runtime runtime = 7; 应用的运行状态
string timelyState =8; 应用动态信息 
map<string,string> conf =9;应用配置文件
repeated string editable =10; 可修改的值
```

###获取应用列表

http请求类型：GET

是否需要token：是（token存入http head Authorization key）

请求body：

```
string type =2; 值->list
string region =3; 云端区域
int64 offset =4; 查询偏移值
int64 length =5; 查询长度
string orderby = 6; 排序字段
```

返回数据：proto类型 -> AppList

```
int64 offset =1; 
int64 length =2; 列表长度
string region =3; 云端区域
repeated  App content =4; 应用列表
```


##应用模板api
请求url：api_v1.xxxx.xxx/app/temp
###获取应用模板

http请求类型：GET

是否需要token：是（token存入http head Authorization key）

请求body：

```
string id =1; 应用id
string type =2; 值->app_tmp
string region =3; 云端区域
```

返回数据：proto类型 -> App

```
string name =2; 应用名字
string region =3; 引用所属区域代码 默认“shanghai”
repeated  Pod Pods = 4 ; 应用中的实例
repeated  Service service = 5 ; 应用中的网络服务
repeated Volume volumes = 6; 应用所挂载的服务器空间
Runtime runtime = 7; 应用的运行状态
string timelyState =8; 应用动态信息 
map<string,string> conf =9;应用配置文件
repeated string editable =10; 可修改的值
```

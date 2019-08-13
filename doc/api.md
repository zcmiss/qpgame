## 接口设计要求
<ul class="list-group" style="width:60%;font-size:14px;">
  <li class="list-group-item list-group-item-info">1. 模块之间是松耦合的,本项目的设计理念为通过良好的组织设计来解决包冲突问题</li>
  <li class="list-group-item list-group-item-info">2. 接口响应字段有以下字段组成(code状态码,interalMsg给开发人员看的消息说明,clientMsg给用户看的消息说明,成功响应的data)</li>
</ul>


## 规范说明
<ul class="list-group" style="width:60%;font-size:14px;">
    <li class="list-group-item list-group-item-info">1. 所有异常，如果需要修复的在日志文件中记录,否则正常终止返回接口</li>
    <li class="list-group-item list-group-item-info">2. 所有接口返回状态码必须定义在config下面的apiStatusCode.go文件中</li>
    <li class="list-group-item list-group-item-info">3. 项目共用无依赖处理类定义在qpp/util目下的文件中 </li>
    <li class="list-group-item list-group-item-info">4. 日志名称为logrus.log存放在项目当前目录,日志统一有util.Log对象打印到文件</li>
    <li class="list-group-item list-group-item-info">5. 不要过度封装mysql,redis,频繁操作的在models下面创建表模型即可，其他情况直接从连接池中拿数据连接对象原生操作，mysql,redis一些共用处理的业务封装到models下面的redis,mysql文件夹里的commonUtil.go中</li>
    <li class="list-group-item list-group-item-info">6. 响应字段如果无内容字符串一律用""表示,不要用null</li>
    <li class="list-group-item list-group-item-info">7. 取消redis的has判断，直接获取结果即可，没有结果会有对应的返回值,可以减少一次redis操作</li>
    <li class="list-group-item list-group-item-info">8. 内部方法访问为小写首字母，方法一律采用驼峰命名</li>
    <li class="list-group-item list-group-item-info">9. 时区为中国上海时区</li>
    <li class="list-group-item list-group-item-info">10. mysql连接池的连接数量不要改，上线之后会调整出一个最佳数量，数据库的连接数要控制在300个以下(32核)</li>
    <li class="list-group-item list-group-item-info">11. </li>
 </ul>


## 架构设计
<ul class="list-group" style="width:60%;font-size:14px;">
  <li class="list-group-item list-group-item-success">1. 本项目微服务架构设计</li>
  <li class="list-group-item list-group-item-success">2. 可伸缩集群</li>
  <li class="list-group-item list-group-item-success">3. 永不宕机服务</li>
  <li class="list-group-item list-group-item-success">4. 设计语言为go</li>
  <li class="list-group-item list-group-item-success">5. https安全通道</li>
  <li class="list-group-item list-group-item-success">6. 负载均衡</li>
  <li class="list-group-item list-group-item-success">7. 关系型数据库mysql</li>
  <li class="list-group-item list-group-item-success">8. nosql数据库mongodb,redis</li>
  <li class="list-group-item list-group-item-success">9. restful接口规范</li>
</ul>









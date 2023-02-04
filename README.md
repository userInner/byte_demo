# byte_demo V1 
## Running
## 请在已加入项目组环境下，执行git
````
mkdir byte_demo
cd byte_demo
git init
git checkout -b byte_demo.V1
git remote add origin git@github.com:userInner/byte_demo.git
git pull origin byte_demo.V1
````
## 命令规范
1. 变量、函数驼峰命名法
2. 务必与变量内容相关
3. 文件名小写+下划线方式
4. `git commit -m` '注释含义明确'

## 目前进度
*[x] 数据库MySQL表设计
*[x] 用户注册、登录
* [ ] 待补充

<h4>common: 数据库工具</h4>
<h4>dao: 数据库持久层</h4>
<h4>dto: 返回前台数据</h4>
<h4>middleware: 中间件</h4>
<h4>models: 数据库模型</h4>
<h4>routers: 路由管理</h4>
<h4>utils: 工具包</h4>


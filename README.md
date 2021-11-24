### 帮某人做的毕设

1. 用户系统
   	1.  居民，回收机构
2.  回收信息发布与接收
3.  垃圾查询





页面 ：

1. 登录/注册
2. 主页面
   1. 导航栏， 用户，登出
   2. 搜索
   3. 回收信息



API : 

1. 查询垃圾分类 /rubbish?name=？
2. 登录 /login (get)
3. 注册 /register (get)
4. 发布回收信息 /recycle (post)
5. 获取回收信息列表 /recycle (get)
6. 接受回收信息 /recycle (put)



数据表

user

1. user 主键
2. pwd
3. session
4. type 用户类型

recycle

1. id
2. c_user  发布者
3. place 发布时间
4. r_user 接收者
5. time 
6. info 备注

garbages

1. id
2. name
3. fname
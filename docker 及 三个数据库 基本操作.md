# docker

docker 基本操作

```docker
docker ps               // 显示正在运行的容器
docker ps -a            // 显示所有容器
docker run              // 运行容器
docker start            // 启动容器
docker stop             // 停止容器
docker restart          // 重启容器
docker rm               // 删除容器
docker pull             // 拉取镜像
docker images           // 列出镜像
docker rmi              // 删除镜像
```

docker 运行 mysql/mongo/cockroach/mssql

```docker
docker run -d -p 127.0.0.1:3306:3306 –-name shopapi -e MYSQL_ROOT_PASSWORD=123456 mysql

docker run -d -p 3307:27017 --name shopmgo  mongo

// 运行一个节点
docker run -d --name=roach1 --hostname=roach1 --net=roachnet -p 26257:26257 -p 8080:8080  -v "${PWD}/cockroach-data/roach1:/cockroach/cockroach-data"  cockroachdb/cockroach start --insecure

docker run -d --name name_your_container -e 'ACCEPT_EULA=Y' -e 'SA_PASSWORD=P@55w0rd' -p 1433:1433 microsoft/mssql-server-linux


// -d 后台运行
// -p 本地 3306 端口对应 docker 3306
// --name 容器名字
// -e MYSQL_ROOT_PASSWORD=123456 MySQL 密码
// mysql使用的镜像
// -v 创建数据卷 
```

## MySQL 基本操作

``注意 MySQL 保留``

``MySQL 命令``

```sql
// hostname 本机使用 127.0.0.1 或 localhost
mysql -h hostname -u username -p

// 修改 root 密码
/usr/bin/mysqladmin -u root password '123456'

// 查看当前所有存在的数据库
show databases;

// 创建数据库
create database database_name;

// 查看创建好的数据库
show create database database_name\G

// 删除数据库
drop database database_name;

// 查看数据库的存储引擎
show engines\G

// 选择当前数据库
use database_name

// 创建数据表
create table <表名>
(
字段名1, 数据类型 [列级别约束条件][默认值],
字段名2, 数据类型 [列级别约束条件][默认值],
...
[表级别约束条件]
);

// 查看数据表
show tables;

// 查看具体的表
describe table_name;
```

更多 MySQL操作 请查看

``
http://www.cnblogs.com/tuhooo/p/5441897.html
``

``
http://blog.csdn.net/cleanness/article/details/42967661
``

## Mongo 基本操作
简单命令

db

```mongo
查看当前使用的数据库
    > db
 
切换数据库
    > use admin

查看MongoDB实例拥有哪些数据库
    > show dbs;

不需要显式创建数据库，当向数据库的某个collection插入文档时，数据库就被创建

    yangs-Air(mongod-3.4.0) test> show dbs
    admin  → 0.000GB
    analy  → 0.000GB
    test   → 0.000GB
    local  → 0.000GB
    yangs-Air(mongod-3.4.0) test> use user
    switched to db user
    yangs-Air(mongod-3.4.0) github> show dbs
    admin → 0.000GB
    analy → 0.000GB
    local → 0.000GB
    test  → 0.000GB
    yangs-Air(mongod-3.4.0) user> db.users.insert({"name":"yang"})
    Inserted 1 record(s) in 87ms
    WriteResult({
      "nInserted": 1
    })
    yangs-Air(mongod-3.4.0) user> show dbs
    admin  → 0.000GB
    analy  → 0.000GB
    local  → 0.000GB
    test   → 0.000GB
    user   → 0.000GB

有一些数据库名是保留的，可以直接访问这些具有特殊语义的数据库，同时自己命名数据库时注意不要使用这些名字。

insert

insert即向 collection添加新的 documents.如果插入时集合不存在,插入操作会创建该集合。

MongoDB中提供了以下方法来插入文档到一个集合:

- db.collection.insert()
- db.collection.insertOne()
- db.collection.insertMany()

    db.COLLECTION_NAME.insert(document)

实例

db.collection.insert()向集合插入一个或多个文档.要想插入一个文档,传递一个文档给该方法;要想插入多个文档,传递文档数组给该方法.

    >db.col.insert({"name":"test"})

该操作返回了含有操作状态的 WriteResult对象.插入文档成功返回如下WriteResult 对象:

    WriteResult({ "nInserted" : 1 })

nInserted字段指明了插入文档的总数.如果该操作遇到了错误, WriteResult 对象将包含该错误信息.

下面我们看下插入多个文档时的情况：

    > db.users.insert(
    ...    [
    ...      { name: "bob", age: 42, status: "A", },
    ...      { name: "ahn", age: 22, status: "A", },
    ...      { name: "xi", age: 34, status: "D", }
    ...    ]
    ... )

该方法返回了包含操作状态的 BulkWriteResult对象.成功插入文档返回如下 BulkWriteResult对象:

    BulkWriteResult({
    	"writeErrors" : [ ],
    	"writeConcernErrors" : [ ],
    	"nInserted" : 3,
    	"nUpserted" : 0,
    	"nMatched" : 0,
    	"nModified" : 0,
    	"nRemoved" : 0,
    	"upserted" : [ ]
    })

现在我们可以调用find()查看一下:

    > db.col.find()
    { "_id" : ObjectId("5870c0b3d599544ddbd1d575"), "name" : "test" }
    { "_id" : ObjectId("5870cf70d599544ddbd1d579"), "name" : "bob", "age" : 42, "status" : "A" }
    { "_id" : ObjectId("5870cf70d599544ddbd1d57a"), "name" : "ahn", "age" : 22, "status" : "A" }
    { "_id" : ObjectId("5870cf70d599544ddbd1d57b"), "name" : "xi", "age" : 34, "status" : "D" }

可以看到已经插入成功，但每个文档都多了一个字段："_id",它是那里来的呢？

MongoDB中储存的文档必须有一个"id"键。这个键可以是任意类型的，默认为ObjectId对象。在一个集合中，每个文档都有唯一的"id"值，来确保集合里面每个文档都能被唯一标识。如果插入文档的时候没有"id"键,系统会自动帮你创建一个。这就是为什会多出来一个"id"字段。

接下来看下另外两个函数。

db.collection.insertOne() 向集合插入单个document。

db.collection.insertMany()向集合插入多个 documents。

    > db.users.insertMany(
       [
         { name: "bob", age: 42, status: "A", },
         { name: "ahn", age: 22, status: "A", },
         { name: "xi", age: 34, status: "D", }
       ]
    )

例子同前。

query

MongoDB 查询数据的语法格式如下：

    >db.COLLECTION_NAME.find(<query filter>, <projection>)

find() 方法以非结构化的方式来显示所有文档。

如果你需要以易读的方式来读取数据，可以使用 pretty() 方法，语法格式如下：

    >db.COLLECTION_NAME.find().pretty()

除了 find() 方法之外，还有一个 findOne() 方法，它只返回一个文档。

条件查询

    >db.users.find( { status: "A" } )
    
    >db.users.find( { status: "A", age: { $lt: 30 } } )

上面四个例子中，第一个查询字段status的值为"A"的文档，第二和第三个查询status的值为"P"或"D"的文档,第四个查询status的值为"A",且age的值小于30的文档

  操作   	格式格式                  	范例                              	RDBMS类似语句         
  等于   	{<key>:<value>}       	db.col.find({"name":"mon"})     	where name = "mon"
  小于   	{<key>:{$lt:<value>}} 	db.col.find({"count":{$lt:50}}) 	where count < 50  
  小于或等于	{<key>:{$lte:<value>}}	db.col.find({"count":{$lte:50}})	where count <= 50 
  大于   	{<key>:{$gt:<value>}} 	db.col.find({"count":{$gt:50}}) 	where count> 50   
  大于或等于	{<key>:{$gte:<value>}}	db.col.find({"count":{$gte:50}})	where count >= 50 
  不等于  	{<key>:{$ne:<value>}} 	db.col.find({"count":{$ne:50}}) 	where count != 50 

返回指定键

    db.user.find({}, {"name" : 1, "_id" : 0})

多条件

    db.users.find(
      {
        status: "A",
        $or: [ { age: { $lt: 30 } }, { type: 1 } ]
      }
    )

查询一个条件的多个值

    db.user.find({"num" : {"$in" : [12, 13, 23]}})

取反

    db.user.find({"num" : {"$not" : [12, 13]})



update

简单实例

    db.people.update({"_id": 1},{"$set" : {"name": "yang"})

    db.users.insert(
     {
        _id: 1,
        name: "sue",
        age: 19,
        type: 1,
        status: "P",
        favorites: { artist: "Picasso", food: "pizza" },
        finished: [ 17, 3 ],
        badges: [ "blue", "black" ],
        points: [
            { points: 85, bonus: 20 },
            { points: 85, bonus: 10 }
         ]
     }
    )

修改内嵌文档

    db.users.update(
       { "favorites.artist": "Pisanello" },
       {
         $set: { "favorites.food": "pizza"}
       }
    )

 错误：

    db.users.update(
       { "favorites.artist": "Pisanello" },
       {
         { "favorites.food": "pizza"}
       }
    )

数字增加或减少

    db.user.update({"_id" : 1}, {"$inc" : {"article" : 1}})

数组

    db.comment.update({"_id" : 1}, {"$push" : {"comments" : "first"})

添加多个

    db.blog.post.update({"title" : "Post"}, {"$push" : {"tag" : {"$each" : ["go", "linux", "database"]}}})

避免重复

    db.user.update({"_id" : 1}, {"$addToSet" : {"emails" : "23132132132@qq.com"})

删除

pop

    db.comment.update({"_id" : 1}, {"$pop" : {"comments" : 1})

pull

    db.comment.update({}, {"$pull" : {"comments" : "xxxx"})

位置

    db.user.update({"article" : "post"}, {"$inc" : {"comments.0.votes" : 1}})

定位

    db.user.update({"article" : "post"}, {"$inc" : {"comments.$.author" : "Jim"}})

upsert

如果 db.collection.update()，db.collection.updateOne()， db.collection.updateMany() 或者 db.collection.replaceOne()包含 upsert : true  并且没有文档匹配指定的过滤器，那么此操作会创建一个新文档并插入它。如果有匹配的文档，那么此操作修改或替换匹配的单个或多个文档。

    db.user.update({"rep" : 25}, {"$inc" : {"rep" : 3}}, true)

remove

删除

    db.user.remove({})

    db.users.remove( { status : false }, 1)

    db.collection.deleteOne({ status: "D" })

```

## cockroachdb 基本操作

#### **1. 创建网桥**

由于您将在单个主机上运行多个 Docker 容器，因此每个容器有一个 CockroachDB 节点，您需要创建Docker所指的桥接网络。桥接网络将使容器能够作为单个群集进行通信，同时保持与外部网络的隔离。

```docker
    docker network create -d bridge roachnet
```

我们在 roachnet 这里和随后的步骤中使用了网络名称，但是请随时给您的网络任何您喜欢的名字。

#### **2. 启动第一个节点**

```docker
    docker run -d \
    --name=roach1 \
    --hostname=roach1 \
    --net=roachnet \
    -p 26257:26257 -p 8080:8080  \
    -v "${PWD}/cockroach-data/roach1:/cockroach/cockroach-data"  \
    cockroachdb/cockroach:v1.0.4 start --insecure
```

此命令创建一个容器并启动其中的第一个 CockroachDB 节点。我们来看看每个部分：

docker run：Docker 命令启动一个新的容器。

-d：这个标志在后台运行容器，所以你可以在同一个shell中继续下一步。

--name：容器的名称。这是可选的，但是自定义名称使得在其他命令中引用容器更容易，例如在容器中打开Bash会话或停止容器时。

--hostname：容器的主机名。您将使用它将其他容器/节点连接到集群。

--net：用于容器加入的网桥。有关详细信息，请参阅步骤1。

-p 26257:26257 -p 8080:8080：这些标志将用于节点间和客户端节点通信（26257）的默认端口和 8080 从容器到主机的管理UI（）的 HTTP 请求的默认端口映射。这允许集装箱间通信，并可以从浏览器调用管理 UI。

-v "${PWD}/cockroach-data/roach1:/cockroach/cockroach-data"：该标志挂载作为数据卷的主机目录。这意味着该节点的数据和日志将存储在 ${PWD}/cockroach-data/roach1 主机上，并在容器停止或删除后持续。有关更多详细信息，请参阅 Docker 将主机目录作为数据卷主题。 

cockroachdb/cockroach:v1.0.4 start --insecure：CockroachDB 命令以不安全的方式启动容器中的一个节点。

#### **3. 将节点添加到集群**

在这一点上，您的群集是实时和可操作的。只需一个节点，您就可以连接一个 SQL 客户端并开始构建数据库。然而，在实际部署中，您总是希望3个或更多节点可以利用 CockroachDB 的自动复制，重新平衡和容错能力。

要模拟真正的部署，通过添加另外两个节点来扩展集群：

```docker
    # Start the second container/node:
    docker run -d \
    --name=roach2 \
    --hostname=roach2 \
    --net=roachnet \
    -v "${PWD}/cockroach-data/roach2:/cockroach/cockroach-data" \
    cockroachdb/cockroach:v1.0.4 start --insecure --join=roach1
    
    # Start the third container/node:
    docker run -d \
    --name=roach3 \
    --hostname=roach3 \
    --net=roachnet \
    -v "${PWD}/cockroach-data/roach3:/cockroach/cockroach-data" \
    cockroachdb/cockroach:v1.0.4 start --insecure --join=roach1
```

这些命令添加了两个容器，并在其中启动 CockroachDB 节点，并将它们连接到第一个节点。从步骤2中只需要注意几点：

-v：该标志挂载作为数据卷的主机目录。数据和日志这些节点将被存储在 ${PWD}/cockroach-data/roach2与${PWD}/cockroach-data/roach3 主机上和容器停止或删除之后将继续存在。

--join：该标志使用第一个容器将新节点连接到集群 hostname。否则，所有 cockroach start 默认值都被接受。请注意，由于每个节点都在唯一的容器中，所以使用相同的默认端口不会引起冲突。

#### **4. 测试集群**

现在您已经扩展到3个节点，可以使用任何节点作为集群的 SQL 网关。为了演示这一点，使用 docker exec 命令在第一个容器中启动内置的 SQL shell：

```docker
    docker exec -it roach1 ./cockroach sql --insecure
    # Welcome to the cockroach SQL interface.
    # All statements must be terminated by a semicolon.
    # To exit: CTRL + D.
```

运行一些基本的 CockroachDB SQL 语句：

```Cockroach
    > CREATE DATABASE bank;
    
    > CREATE TABLE bank.accounts (id INT PRIMARY KEY, balance DECIMAL);
    
    > INSERT INTO bank.accounts VALUES (1, 1000.50);
    
    > SELECT * FROM bank.accounts;

    +----+---------+
    | id | balance |
    +----+---------+
    |  1 |  1000.5 |
    +----+---------+
    (1 row)
```

退出节点1上的 SQL shell：

```Cockroach
    > \q
```

然后在第二个容器中启动 SQL shell：

```docker 
    docker exec -it roach2 ./cockroach sql --insecure
    # Welcome to the cockroach SQL interface.
    # All statements must be terminated by a semicolon.
    # To exit: CTRL + D.
```

现在运行相同的 SELECT 查询：

```Cockroach
    > SELECT * FROM bank.accounts;

    +----+---------+
    | id | balance |
    +----+---------+
    |  1 |  1000.5 |
    +----+---------+

    > (1 row)
```

如您所见，节点1和节点2的行为与 SQL 网关相同。

完成后，退出节点2上的 SQL shell：

```Cockroach
    > \q
```

#### **5. 监控集群**

当您启动第一个容器/节点时，将节点的默认HTTP端口映射8080到8080主机上的端口。要查看集群的管理UI，请将浏览器指向该端口localhost，即 http://localhost:8080。



如前所述，CockroachDB会自动将您的数据复制到幕后。要验证上一步中写入的数据是否已成功复制，请向下滚动到每个商店的副本图表并将其悬停在该行上：

每个节点上的副本计数相同，表示集群中的所有数据都被复制3次（默认值）。

#### **6. 停止集群**

使用 docker stop 和 docker rm 命令停止和删除容器（因此集群）：
```docker 
    # Stop the containers:
    docker stop roach1 roach2 roach3
    
    # Remove the containers:
    docker rm roach1 roach2 roach3
```

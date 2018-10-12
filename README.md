# share 介绍
基于golang原生database/sql，写一个适合业务场景的MySQL ORM，需要实现的需求如下：
1. 支持但不限于增删改查功能；
2. 支持SQL事务；
3. 服务运行的情况下，不重启服务能修改MySQL的连接信息；
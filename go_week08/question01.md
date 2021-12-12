# Question 1
题目1：使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。


➜ ~ echo "redis test 10 byte value:" redis test 10 byte value: ➜ ~ redis-benchmark -t set,get -n 2000000 -q -d 10 SET: 168137.86 requests per second, p50=0.175 msec GET: 170299.72 requests per second, p50=0.167 msec

➜ ~ echo "redis test 20 byte value:" redis test 20 byte value: ➜ ~ redis-benchmark -t set,get -n 2000000 -q -d 20 SET: 169534.62 requests per second, p50=0.175 msec GET: 170314.23 requests per second, p50=0.167 msec

➜ ~ echo "redis test 50 byte value:" redis test 50 byte value: ➜ ~ redis-benchmark -t set,get -n 2000000 -q -d 50 SET: 168961.73 requests per second, p50=0.175 msec GET: 169750.47 requests per second, p50=0.167 msec

➜ ~ echo "redis test 100 byte value:" redis test 100 byte value: ➜ ~ redis-benchmark -t set,get -n 2000000 -q -d 100 SET: 167098.34 requests per second, p50=0.175 msec GET: 169018.84 requests per second, p50=0.167 msec

➜ ~ echo "redis test 200 byte value:" redis test 200 byte value: ➜ ~ redis-benchmark -t set,get -n 2000000 -q -d 200 SET: 166889.19 requests per second, p50=0.175 msec GET: 160346.36 requests per second, p50=0.167 msec

➜ ~ echo "redis test 1k byte value:" redis test 1k byte value: ➜ ~ redis-benchmark -t set,get -n 2000000 -q -d 1000 SET: 168293.50 requests per second, p50=0.175 msec GET: 165384.92 requests per second, p50=0.175 msec

➜ ~ echo "redis test 5k byte value:" redis test 5k byte value: ➜ ~ redis-benchmark -t set,get -n 2000000 -q -d 5000 SET: 161799.20 requests per second, p50=0.183 msec GET: 157616.83 requests per second, p50=0.183 msec

2000000

|  | SET | GET | 单位 |
| --- | --- | --- | --- |
| 10  | 168137.86 | 170299.72 | QPS |
| 20 | 169534.62 |170314.23  | QPS |
| 50 | 168961.73 | 169750.47 | QPS |
| 100 |167098.34  | 169018.84 | QPS |
| 200 | 166889.19 | 160346.36 | QPS |
| 1000 |168293.50  | 165384.92 | QPS |
| 5000 |  161799.20| 157616.83 |QPS  |

**对比结论：10 20 50 100 200 1k 5k 字节 value 大小，性能差别不大； GET的性能比SET的性能稍微差点；当超过10K，随着字节数越来越大， 性能的表现越来越差。**

### 参数的作用

| 作用分类            | 参数及作用                                                                                                             |
|-----------------|-------------------------------------------------------------------------------------------------------------------|
| 连接 Redis 服务相关参数 | -h ：Redis 服务主机地址，默认为 127.0.0.1 。                                                                                  |
| _               | -p ：Redis 服务端口，默认为 6379 。                                                                                         |
| _	              | -s ：指定连接的 Redis 服务地址，用于覆盖 -h 和 -p 参数。一般情况下，我们并不会使用。                                                               |
| _	              | -a ：Redis 认证密码。                                                                                                   |
| 	_              | –dbnum ：选择 Redis 数据库编号。                                                                                           |
| 	_              | k ：是否保持连接。默认会持续保持连接。                                                                                              |
| 请求相关参数          | 	-c ：并发的客户端数                                                                                                      |
| 	_              | -n ：总共发起的操作（请求）数                                                                                                  |
| 	_              | -d ：指定 SET/GET 操作的数据大小，单位：字节。                                                                                     |
| 	_              | -r ：SET/GET/INCR 使用随机 KEY ，SADD 使用随机值。通过设置-r参数，可以设置KEY的随机范围，比如-r 10生成的KEY范围为[0，9）                                 |
| 	_              | -P ：默认情况下，Redis 客户端一次请求只发起一个命令。通过 -P 参数，可以设置使用 _                                                                  |pipeline功能，一次发起指定个请求，从而提升 QPS 。|
| _               | 	-l ：循环，一直执行基准测试。                                                                                                 |
| _               | 	-t ：指定需要测试的 Redis 命令，多个命令通过逗号分隔。默认情况下，测试 PING_INLINE/PING_BULK/SET/GET 等等命令。若只想测试 SET/GET 命令，则可以 -t SET,GET 来指定。 |
| 	_              | -I ：Idle 模式。仅仅打开 N 个 Redis Idle 个连接，然后等待，啥也不做。不是很理解这个参数的目的，目前猜测，仅仅用于占用 Redis 连接。                                  |
| 结果输出相关参数        |                                                                                                                   |
| _	              | -e ：如果 Redis Server 返回错误，是否将错误打印出来。默认情况下不打印，通过该参数开启。                                                              |
| _               | 	-q ：精简输出结果。即只展示每个命令的 QPS 测试结果                                                                                    |
| _               | 	-csv ：按照 CSV 的格式，输出结果                                                                                            |

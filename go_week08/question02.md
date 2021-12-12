# Question 2
题目：写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息, 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

```shell
默认端口：6379

redis-benchmark -h 127.0.0.1 -q -t set -r 100000 -n 500000 -d 10
SET: 108178.27 requests per second, p50=0.407 msec

➜  ~ redis-cli info memory
used_memory:12264448
used_memory_human:11.70M
used_memory_rss:14241792
used_memory_rss_human:13.58M
used_memory_peak:16355088
used_memory_peak_human:15.60M
used_memory_peak_perc:74.99%
used_memory_overhead:6077136
used_memory_startup:1028320
used_memory_dataset:6187312
used_memory_dataset_perc:55.07%
allocator_allocated:12180944
allocator_active:14203904
allocator_resident:14203904
total_system_memory:17179869184
total_system_memory_human:16.00G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.17
allocator_frag_bytes:2022960
allocator_rss_ratio:1.00
allocator_rss_bytes:0
rss_overhead_ratio:1.00
rss_overhead_bytes:37888
mem_fragmentation_ratio:1.17
mem_fragmentation_bytes:2060848
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:0
mem_aof_buffer:0
mem_allocator:libc
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0

redis-cli info memory 看日志前后的 写入前used_memory_dataset   写入后used_memory_dataset    
 
 redis-benchmark -h 127.0.0.1 -q -t set -r 100000 -n 500000 -d 10
 -r 后面是 随机额key
 -n 后面 跟的是 总共发起的操作（请求）数
 （写入后used_memory_dataset -  写入前used_memory_dataset/B）/ 总共发起的操作（请求）数  等于 平均每个 key 的占用内存空间。
计算：（290346024-276779232)/100000  =  133.67

以此方式 类推 调用 10 20 50 100 200 1000 5000就能测出来 不同value 的key的占用内存空间大小

也可以利用 [Redis容量预估-极数云舟](http://www.redis.cn/redis_memory/) 这个网址取计算一下一个 key 大概占用的内存空间。
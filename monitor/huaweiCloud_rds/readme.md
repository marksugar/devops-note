这是一个在华为云没有提供钉钉 webhook 时候通过消息通知的 HTTP，接收 rds 的告警，匹配华为云 RDS 的触发器的 metric_name 来做警报处理。

对于慢查询的处理：
1，收到 rds074_slow_queries 后
2，开始拉取最近 2 分钟的最近的一条 slow 日志

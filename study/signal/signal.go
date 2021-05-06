/**
信号 Signature

kill {pid}

// list all kill commands
kill -l

// 默认情况下 kill 等同于 kill -15 `SIGTERM 15`
kill
kill -15

// interrupt 中断，同 `CTRL + C`, `SIGINT 2`
kill -2

// `SIGKILL 9` 立即结束程序，不能被阻塞，处理和忽略
kill -9

// 用户自定义信号
kill -USR1
kill -USR2
kill -SIGUSR1
kill -SIGUSR2

其他信号量参考 syscall.SIGINT
*/
package signal

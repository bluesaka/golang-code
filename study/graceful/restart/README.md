### 平滑重启
- grace
- endless
- overseer

grace与endless：旧的api都不会断掉，会执行原来的逻辑，但pid会变化，不支持supervisor管理
overseer：旧api不会断掉，会执行原来的逻辑，主进程pid也不会变化，支持supervisor、systemd等管理


supervisorctl signal sigusr2 {process_name}
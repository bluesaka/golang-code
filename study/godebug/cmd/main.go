/**
GODEBUG

GODEBUG变量可以控制运行时内的调试变量，主要涉及 gctrace 参数，
我们通过设置 gctrace=1 后就可以使得垃圾收集器向标准错误流发出 GC 运行信息。
*/

/**
output:

gc 1 @0.006s 2%: 0.014+0.75+0.062 ms clock, 0.11+0.48/0.54/0.32+0.49 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
gc 2 @0.008s 5%: 0.056+0.45+0.035 ms clock, 0.44+0.87/0.49/0+0.28 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
gc 3 @0.009s 6%: 0.064+0.39+0.044 ms clock, 0.51+0.76/0.44/0+0.35 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
gc 4 @0.010s 8%: 0.071+0.54+0.052 ms clock, 0.57+0.72/0.50/0+0.41 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
gc 5 @0.012s 9%: 0.045+0.62+0.054 ms clock, 0.36+0.99/0.61/0+0.43 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
gc 6 @0.013s 10%: 0.063+0.84+0.063 ms clock, 0.50+0.78/0.95/0+0.50 ms cpu, 4->5->1 MB, 5 MB goal, 8 P
gc 7 @0.015s 11%: 0.076+0.86+0.064 ms clock, 0.61+0.57/1.1/0+0.51 ms cpu, 4->5->2 MB, 5 MB goal, 8 P
gc 8 @0.017s 13%: 0.086+0.81+0.056 ms clock, 0.69+1.0/1.0/0+0.45 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
gc 9 @0.018s 13%: 0.048+0.91+0.060 ms clock, 0.38+0.90/1.2/0.030+0.48 ms cpu, 4->5->2 MB, 5 MB goal, 8 P
gc 10 @0.020s 14%: 0.041+0.52+0.020 ms clock, 0.33+0.98/0.81/0.27+0.16 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
gc 11 @0.029s 10%: 0.024+0.45+0.002 ms clock, 0.19+0.14/0.65/0.78+0.023 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
gc 12 @0.031s 11%: 0.042+0.75+0.14 ms clock, 0.34+0.95/0.84/0.008+1.1 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
gc 13 @0.032s 11%: 0.062+0.51+0.048 ms clock, 0.49+0.69/0.55/0+0.38 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
gc 14 @0.033s 11%: 0.051+0.56+0.088 ms clock, 0.41+0.56/0.78/0+0.70 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
gc 15 @0.035s 12%: 0.077+0.63+0.077 ms clock, 0.62+0.52/0.78/0+0.61 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
gc 16 @0.036s 12%: 0.066+0.61+0.064 ms clock, 0.52+0.80/0.70/0+0.51 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
gc 17 @0.037s 13%: 0.050+0.64+0.035 ms clock, 0.40+0.64/0.82/0+0.28 ms cpu, 4->5->1 MB, 5 MB goal, 8 P
gc 18 @0.038s 13%: 0.062+0.68+0.061 ms clock, 0.49+0.69/0.88/0+0.49 ms cpu, 4->5->2 MB, 5 MB goal, 8 P
gc 19 @0.040s 13%: 0.044+0.78+0.073 ms clock, 0.35+0.65/1.1/0+0.58 ms cpu, 4->5->2 MB, 5 MB goal, 8 P
gc 20 @0.041s 14%: 0.065+0.88+0.069 ms clock, 0.52+0.69/1.2/0+0.55 ms cpu, 4->5->2 MB, 5 MB goal, 8 P
gc 21 @0.043s 14%: 0.026+0.70+0.004 ms clock, 0.21+0.35/0.97/0.56+0.036 ms cpu, 4->4->1 MB, 5 MB goal, 8 P
gc 22 @0.051s 12%: 0.015+0.35+0.005 ms clock, 0.12+0.11/0.43/0.48+0.040 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
gc 23 @0.066s 10%: 0.35+0.54+0.004 ms clock, 2.8+0.47/0.82/0.034+0.039 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
gc 24 @0.091s 7%: 0.18+0.46+0.003 ms clock, 1.4+0.17/0.58/0.89+0.025 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
# command-line-arguments
gc 1 @0.003s 6%: 0.025+2.5+0.016 ms clock, 0.20+0.12/2.3/2.7+0.13 ms cpu, 4->5->3 MB, 5 MB goal, 8 P
gc 25 @0.099s 7%: 0.073+0.42+0.022 ms clock, 0.59+0.21/0.42/0.65+0.18 ms cpu, 4->4->0 MB, 5 MB goal, 8 P
# command-line-arguments
gc 1 @0.001s 9%: 0.008+2.0+0.018 ms clock, 0.067+0.15/2.4/3.2+0.14 ms cpu, 4->6->5 MB, 5 MB goal, 8 P
gc 2 @0.008s 5%: 0.007+1.2+0.017 ms clock, 0.062+0/1.7/1.1+0.13 ms cpu, 8->9->6 MB, 10 MB goal, 8 P
gc 3 @0.031s 3%: 0.29+2.0+0.019 ms clock, 2.3+0.13/1.2/2.1+0.15 ms cpu, 12->12->8 MB, 13 MB goal, 8 P
*/

/**
举例：gc 7 @0.015s 11%: 0.076+0.86+0.064 ms clock, 0.61+0.57/1.1/0+0.51 ms cpu, 4->5->2 MB, 5 MB goal, 8 P

- gc 7：第 7 次 GC。
- @0.015s：当前是程序启动后的 0.015s。
- 11%：程序启动后到现在共花费 11% 的时间在 GC 上。
- 0.076+0.86+0.064 ms clock：
	- 0.076：表示单个 P 在 mark 阶段的 STW 时间。
	- 0.86：表示所有 P 的 mark concurrent（并发标记）所使用的时间。
	- 0.064：表示单个 P 的 markTermination 阶段的 STW 时间。
- 0.61+0.57/1.1/0+0.51 ms cpu：
	- 0.61：表示整个进程在 mark 阶段 STW 停顿的时间。
	- 0.57/1.1/0：0.57 表示 mutator assist 占用的时间，1.1 表示 dedicated + fractional 占用的时间，0 表示 idle 占用的时间。
	- 0.51ms：表示整个进程在 markTermination 阶段 STW 时间。
- 4->5->2 MB：
	- 4：表示开始 mark 阶段前的 heap_live 大小。
	- 5：表示开始 markTermination 阶段前的 heap_live 大小。
	- 2：表示被标记对象的大小。
- 5 MB goal：表示下一次触发 GC 回收的阈值是 5 MB。
- 8 P：本次 GC 一共涉及多少个 P。
*/
package main

import "sync"

func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			var counter int
			for i := 0; i < 1e10; i++ {
				counter++
			}
		}(&wg)
	}
	wg.Wait()
}

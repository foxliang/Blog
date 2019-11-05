## top命令
```
[root@localhost ~]# top
top - 11:12:18 up 4 days, 11:02,  3 users,  load average: 0.00, 0.01, 0.05
Tasks: 108 total,   2 running, 106 sleeping,   0 stopped,   0 zombie
%Cpu(s):  0.0 us,  0.3 sy,  0.0 ni, 99.7 id,  0.0 wa,  0.0 hi,  0.0 si,  0.0 st
KiB Mem :   999936 total,   194136 free,   346648 used,   459152 buff/cache
KiB Swap:  2097148 total,  2095580 free,     1568 used.   420808 avail Mem 

   PID USER      PR  NI    VIRT    RES    SHR S %CPU %MEM     TIME+ COMMAND                                          
   903 root      20   0  553164  14764   4128 S  0.3  1.5   1:39.11 tuned                                            
     1 root      20   0  193628   6448   3684 S  0.0  0.6   0:33.31 systemd                                          
     2 root      20   0       0      0      0 S  0.0  0.0   0:00.09 kthreadd 
```
参数说明

第一行
11:12:18 当前时间 
4 days, 11:02 系统运行时间
3 users 当前登录用户数
load average: 0.00, 0.01, 0.05 系统负载(1分钟 10分钟 15分钟),即任务队列的平均长度(单核CPU中不超过1是正常的,超过1说明负载压力大)

第二行
108 total 进程总数 2 running 运行进程数 106 sleeping 休眠进程数 0 stopped 终止进程数 0 zombie 僵尸进程数

第三行
0.0 us 用户空间占用cpu百分比
0.3 sy 内核空间占用cpu百分比
0.0 ni 用户进程空间内改变过优先级的进程占用cpu百分比
99.7 id 空闲cpu百分比
0.0 wa 等待输入输出（I/O）的cpu百分比
0.0 hi cpu处理硬件中断的时间
0.0 si cpu处理软件中断的时间 
0.0 st 用于有虚拟cpu的情况,用来指示被虚拟机偷掉的cpu时间

第四行
999936 total 内存总数
194136 free 空闲的内存数
346648 used 已经使用的内存数   
459152 buff/cache 缓存的内存数

第五行
2097148 total 总的交换空间
2095580 free 空闲的交换空间     
1568 used 已经使用的交换空间   
420808 avail Mem 可用交换空间

进程信息区统计信息区域的下方显示了各个进程的详细信息
PID 进程ID
USER 进程所有者用户名
PR 优先级
NI nice值,负值表示高优先级,正值表示低优先级   
VIRT 进程使用的虚拟内存总量   
RES 物理内存用量  
SHR 共享内存用量
S 该进程的状态;其中S代表休眠状态,D代表不可中断的休眠状态,R代表运行状态,Z代表僵尸状态,T代表停止或跟踪状态 
%CPU 进程占用的CPU时间和总时间的百分比 
%MEM 进程占用的物理内存和总内存的百分比    
TIME+ 累计CPU占用时间 
COMMAND 命令名/命令行

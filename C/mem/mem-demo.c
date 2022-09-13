/* ~~~~~~~~~
申请1GB内存，分别查看VIRT && RES内存 大小
 Author: Leo
 Usage: gcc mem-demo.c -o mem-demo && ./mem-demo

    PID USER      PR  NI    VIRT    RES    SHR S  %CPU  %MEM     TIME+ COMMAND                                                                  
  19012 leo       20   0 1051352    980    892 S   0.0   0.0   0:00.00 mem-demo   
*/
#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>

int main(int argc, char const *argv[])
{
    void *p = malloc(1024 * 1024 * 1024);
    if (p == NULL)
    {
        printf("malloc failed\n");
    }
    // int *p2 = (int *)(p + 1024 * 1024);
    // *p2 = 1;
    printf("运行top -p %d查看内存参数\n", (int)(getpid()));
    sleep(1000);
    return 0;
}

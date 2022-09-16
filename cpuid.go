package main

/*
#include<stdio.h>
#include<string.h>
char* getcpuid() {
    static char all_buf[32*4];
    unsigned int s1,s2;
    char buf[32] = {0};
    asm volatile(
    "push %%rax;"
    "push %%rdx;"
    "push %%rbx;"
    "push %%rcx;"
    "movl $0x01, %%eax;"
    "xorl %%edx, %%edx;"
    "cpuid;"
    "movl %%edx, %0;"
    "movl %%eax, %1;"
    "pop %%rcx;"
    "pop %%rbx;"
    "pop %%rdx;"
    "pop %%rax;"
    :"=m"(s1), "=m"(s2)
    );

    if (s1) {
        memset(buf, 0, 32);
        snprintf(buf, 32, "%08X", s1);
        strncpy(all_buf,buf,sizeof(buf));
    }

    if (s2) {
        memset(buf, 0, 32);
        snprintf(buf, 32, "%08X", s2);
        strcat(all_buf,buf);
    }
    asm volatile(
        "push %%rax;"
        "push %%rdx;"
        "push %%rbx;"
        "push %%rcx;"
        "movl $0x03, %%eax;"
        "xorl %%ecx, %%ecx;"
        "xorl %%edx, %%edx;"
        "cpuid;"
        "movl %%edx, %0;"
        "movl %%ecx, %1;"
        "pop %%rcx;"
        "pop %%rbx;"
        "pop %%rdx;"
        "pop %%rax;"
        :"=m"(s1), "=m"(s2)
        );
        if (s1) {
            memset(buf, 0, 32);
            snprintf(buf, 32, "%08X", s1);
            strcat(all_buf,buf);
        }
        if (s2) {
            memset(buf, 0, 32);
            snprintf(buf, 32, "%08X", s2);
            strcat(all_buf,buf);
        }
            return all_buf;
    }
*/
import "C"
import "fmt"

func main() {
	str := C.GoStringN(C.getcpuid(), 16)
	fmt.Println(str)
}

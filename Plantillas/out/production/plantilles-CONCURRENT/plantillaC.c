#define THREADS 4
#include <pthread.h>
#include <stdio.h>

struct tdata{
    int tid;
};

int id;

void *run(void *ptr){
    int tid = ((struct tdata*)ptr)->tid;
    printf("Hi, i'm thread %d\n", tid);
    pthread_exit(NULL);

}
int main(int argc, char*argv[]){

    pthread_t threads [THREADS]:
    int rc;
    for (int i = 0; i<THREADS; i++){
        id[i].tid = i;

        rc= pthread_create(&threads[i], NULL, run, (void*) &id[i]);

    } 
    for (int i =0; i<THREADS; i++){
        pthread_join(threads[i], NULL);
    }
}

/*
*
int int pthread_create(pthread_t *restrict thread,
                          const pthread_attr_t *restrict attr,
                          void *(*start_routine)(void *),
                          void *restrict arg);
*/
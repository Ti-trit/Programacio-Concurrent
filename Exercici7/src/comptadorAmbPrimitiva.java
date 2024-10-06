/**
 * Implementacio d'un comptador usant la primitiva de sincronització get&add
 */

import java.util.concurrent.atomic.AtomicInteger;

public class comptadorAmbPrimitiva implements Runnable{
    //DECLARACIO DE VARIABLES GLOBALS
    static final int THREADS = 2;
    static final int MAX_COUNT = 1000000;
    static AtomicInteger number, torn;
    static volatile int n = 0;
    int id;
    //Constructor
    public comptadorAmbPrimitiva(int id) {
        this.id = id;
    }

    @Override
    public void run() {
        System.out.println("Thread " + id);

        for (int i = 0; i < MAX_COUNT / THREADS; i++) {

          int current = number.getAndAdd(1);
            // Get&Add
            while ( current!= torn.get()) {
                //espera activa
            }
            n = n+1; //SC
            System.out.println("Thread " + id + ": " + n);
            //unlock()
            torn.getAndAdd(1);
        }
    }

    public static void main(String[] args) throws InterruptedException {
        Thread[] threads = new Thread[THREADS];
        number = new AtomicInteger(0);
        torn = new AtomicInteger(0);
        int i;
        for (i = 0; i < THREADS; i++) {
            threads[i] = new Thread(new comptadorAmbPrimitiva(i));
            threads[i].start();
        }
        for (i = 0; i < THREADS; i++) {
            threads[i].join();
        }
        float error = (MAX_COUNT - n) / (float) MAX_COUNT * 100;
        System.out.printf("Counter value: %d Expected: %d Error: %3.6f%%\n", n, MAX_COUNT, error);
    }

}
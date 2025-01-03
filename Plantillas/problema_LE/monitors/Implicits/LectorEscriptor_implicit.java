package problema_LE.monitors.Implicits;

import java.util.logging.Level;
import java.util.logging.Logger;

/**
 * Versió del professor
 * A la regió critica hi pot haver multiples lectors a la vegada
 * Un escriptor només pot accedir a la SC no hi ha cap procés a dins
 */
public class LectorEscriptor_implicit implements Runnable {
    public final static int THREADS = 6;
    static final int MAX_COUNT = 10000;
    static volatile int counter = 0;
    int id;
    monitorLE monitor;

    public LectorEscriptor_implicit(int id, monitorLE monitor) {
        this.id = id;
        this.monitor = monitor;
    }

    @Override
    public void run() {
        int max = MAX_COUNT / THREADS;
        int c;
        System.out.printf("Thread %d\n", id);
        for (int i = 0; i < max; i++) {
            if (i % 10 == 0) {
                monitor.writerLock();
                counter += 1;
                monitor.writerUnlock();
                System.out.println(id + " Incrementa " + counter);
            } else {
                monitor.readerLock();
                c = counter;
                monitor.readerUnlock();
                System.out.println("    " + id + " Llegeix " + counter);
            }
        }
    }

    public static void main(String args[]) throws InterruptedException {
        Thread[] threads = new Thread[THREADS];
        monitorLE monitor = new monitorLE();
        int i;

        for (i = 0; i < THREADS; i++) {
            threads[i] = new Thread(new LectorEscriptor_implicit(i, monitor));
            threads[i].start();
        }
        try {
            for (i = 0; i < THREADS; i++) {
                threads[i].join();
            }
            System.out.printf("Counter value: %d Expected: %d\n", counter, MAX_COUNT / 10);


        } catch (InterruptedException ex) {
            Logger.getLogger(LectorEscriptor_implicit.class.getName()).log(Level.SEVERE, null, ex);
        }
    }
}






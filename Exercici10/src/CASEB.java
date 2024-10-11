import java.util.Random;
import java.util.concurrent.atomic.AtomicInteger;

/**
 * Compare-And-Set Back off exponential
 * CAS is the simplest possible spinlock implementation. It uses a
 * single shared memory location  for synchronization, which indicates if the lock is taken or not.
 * Problem: Very bad scalability. more threads competing to the lock --> more
 * cache line invalidations
 * Solution : back-off.
 * When a thread fails to acquire the lock, it will have
 * to wait for a short time to let
 * other threads finish before trying to enter CS again. This reduces
 * the number of threads simultaneously competing for the lock.
 * A good strategy of back-off is eh the exponential one.
 * Unfortunately, the threads that have been attempting to acquire
 * the lock longest are also backing-off longest.
 * Consequently, newly arriving threads have a higher chance to acquire the lock
 * than older threads. On top of that threads might back-off for too long, causing
 * the CS to be underutilized.
 * link <a href="https://geidav.wordpress.com/tag/exponential-back-off/">...</a>
 */
public class CASEB implements Runnable{
    static final int THREADS = 4;
    static final int MAX_COUNT = 10000;
    static final int MAX_ITERATIONS = 256;
    static final int MIN_BACKOFF_ITERS= 3;
    static AtomicInteger mutex;
    static volatile int n = 0;
    static int maxIters;
    int id;
    Random rn = new Random();
    //Constructor
    public CASEB(int id) {
        this.id = id;
    }

    @Override
    public void run() {
        maxIters = MIN_BACKOFF_ITERS;
        System.out.println("Thread " + id);
        for (int i = 0; i < MAX_COUNT / THREADS; i++) {
            int local = 0;
            try {
                Thread.sleep(0, rn.nextInt(0,3));
            } catch (InterruptedException e) {
                throw new RuntimeException(e);
            }
            //afegir la exponencial
            // CompareAndSet
            while ( !mutex.compareAndSet(local, 1)) {
                //espera activa
                exponentialBack_off();
            }
            n = n+1; //SC
           // System.out.println("Thread " + id + ": " + n);
            //unlock()
            mutex.set(0);
        }
    }

    private void exponentialBack_off(){
        assert (maxIters)>0;
        Random rand = new Random();
        int spines = rand.nextInt(maxIters);
        maxIters = Math.min(2*maxIters, MAX_ITERATIONS);
        System.out.println("Thread " + id + ": " + "spines : "+ spines);
        for (int i = 0; i <spines; i++) {
            //espera activa
        }
    }

    public static void main(String[] args) throws InterruptedException {
        Thread[] threads = new Thread[THREADS];
        mutex =new AtomicInteger(0);
        int i;
        for (i = 0; i < THREADS; i++) {
            threads[i] = new Thread(new CASEB(i));
            threads[i].start();
        }
        for (i = 0; i < THREADS; i++) {
            threads[i].join();
        }
        float error = (MAX_COUNT - n) / (float) MAX_COUNT * 100;
        System.out.printf("Counter value: %d Expected: %d Error: %3.6f%%\n", n, MAX_COUNT, error);
    }

}
import java.util.ArrayDeque;
import java.util.Deque;
import java.util.Random;
import java.util.logging.Level;
import java.util.logging.Logger;
import java.util.concurrent.Semaphore;

/**
 * Enunciat de l'exercici:  13. Implementau el problema de productors - consumidors usant els semàfors de Java
 * Versió de buffer limitat
 */
public class productorConsumidor {
    //constants del programa
    private final static int NUM_CONSUMIDORS = 3;
    private final static int NUM_PRODUCTORS = 3;
    private final static int BUFFER_SIZE =7;
    private final static int repeticions = 10;

    //semafors
    protected static Semaphore notEmpty = new Semaphore(0); // per produir
    protected static Semaphore notFull = new Semaphore(BUFFER_SIZE); // per consumir
    protected static Semaphore mutex = new Semaphore(1); //protegir el buffer
    protected static final Deque<Integer> buffer = new ArrayDeque<>(BUFFER_SIZE);

    public static void main(String[] args) {
        Thread [] consumers = new Thread[NUM_CONSUMIDORS];
        Thread [] productors = new Thread[NUM_PRODUCTORS];
        for (int i = 0; i < NUM_CONSUMIDORS; i++) {
            consumers[i] = new Thread(new consumidor(i+1));
            consumers[i].start();

        }
        for (int i = 0; i < NUM_PRODUCTORS; i++) {
            productors[i]= new Thread(new productor(i+1));
            productors[i].start();
        }
        try{
        for (int i = 0; i < NUM_CONSUMIDORS; i++) {
            consumers[i].join();
        }
        for (int i = 0; i < NUM_PRODUCTORS; i++) {
            productors[i].join();
        }

        }catch (InterruptedException ex){
            Logger.getLogger(productorConsumidor.class.getName()).log(Level.SEVERE, null, ex);
        }
    }




    public static class productor implements  Runnable{
        private static  int id;
        public productor(int id){
            this.id = id;
        }
        @Override
        public void run() {

            Random rn = new Random();
            try{
                for (int i = 0; i < repeticions; i++) {
                    // el buffer encara no esta ple?
                    notFull.acquire();
                    mutex.acquire();
                    int dada = rn.nextInt(0, 100);
                    buffer.addLast(dada);
                    mutex.release();
                    System.out.printf("El productor %d ha produit la dada %d\n", this.id, dada);
                    notEmpty.release();

                    Thread.sleep(rn.nextInt(3, 17));
                }

            }catch (InterruptedException ex){
                Logger.getLogger(productor.class.getName()).log(Level.SEVERE, null, ex);
            }
            //el buffer esta

        }
    }

    public static class consumidor implements Runnable {
        private int id;
        public consumidor(int id){
            this.id = id;
        }
        @Override
        public void run() {
            Random rn = new Random();
            try{
                for (int i = 0; i < repeticions; i++) {

                    //hi ha elements per consumir?
                    notEmpty.acquire();
                    mutex.acquire();
                    //consumir i esborrar del buffer
                    int dada = buffer.removeFirst();
                    System.out.printf("El consumidor %d ha consumit la dada %d\n", this.id, dada);
                    mutex.release();
                    Thread.sleep(rn.nextInt(5, 15));
                    notFull.release();
                }
            }catch (InterruptedException ex){
                Logger.getLogger(productor.class.getName()).log(Level.SEVERE, null, ex);
            }
            //el buffer esta

        }
        }

}

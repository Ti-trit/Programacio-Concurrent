package probelmaAbelles.Monitors.Implicits;

import java.util.logging.Level;
import java.util.logging.Logger;

public class osAbellaImplicit {

    private final static int THREADS= 10; //abelles
    private final static int BUFFER_SIZE = 10;

    public static void main(String[] args) {
        Thread[] threads = new Thread[THREADS];
        int i;
        monitorOsAbella monitor = new monitorOsAbella(BUFFER_SIZE);
        Thread os = new Thread(new os(monitor));
        os.start();
        for (i = 0; i < THREADS; i++) {
            threads[i] = new Thread(new Abella(i, monitor));
            threads[i].start();
        }
        try {
            os.join();
            for (i = 0; i < THREADS; i++) {
                threads[i].join();
            }
            System.out.println("Simulació acabada");
        }catch (InterruptedException ex){
            Logger.getLogger(osAbellaImplicit.class.getName()).log(Level.SEVERE, null, ex);
        }
    }
}


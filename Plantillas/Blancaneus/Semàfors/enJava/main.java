package Blancaneus.enJava;


import java.util.concurrent.Semaphore;
import java.util.logging.Level;
import java.util.logging.Logger;

public class main {

    protected  final static int NUM_NANS = 7;
    protected final static int NUM_CADIRES= 4;
    protected final static int NUM_MENJADES= 2;
    protected static volatile int nans_cabats = 0;
    protected static volatile int volMenjar = 0;
    //declaració de semàfors
    static Semaphore cadires = new Semaphore(NUM_CADIRES);
    static Semaphore EsperarMenjar = new Semaphore(0);
    static Semaphore mutex = new Semaphore(1);
    static Semaphore mutex_nans = new Semaphore(1);
    public static void main(String[] args) {
        System.out.println("Simulació de BlancaNeus i els 7 nans");

        Thread Blancaneus = new Thread(new Blancaneus());
        Blancaneus.start();
        Thread[] nans = new Thread[NUM_NANS];

        for (int i = 0; i < nans.length; i++) {
            nans[i] = new Thread(new nan(i + 1));
            nans[i].start();
        }

        //join
        try {
            Blancaneus.join();

            for (Thread nan : nans) {
                nan.join();
            }


            System.out.println("simulació acabada");

        } catch (InterruptedException ex) {
            Logger.getLogger(main.class.getName()).log(Level.SEVERE, null, ex);

        }

    }


    }



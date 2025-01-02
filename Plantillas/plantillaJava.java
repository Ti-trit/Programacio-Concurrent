
/**
 * Plantilla generica de programació concurrent en Java
 * @author Khaoula Ikkene
 * TOTS ELS TIPUS PRIMITIUS SÓN ATOMICS.
 * LONG I DOUBLE HO PODEN SER SI ES DECLAREN VOLATILS.
 */
import java.util.logging.Level;
import java.util.logging.Logger;
public class plantillaJava implements Runnable{
    static final int THREADS= 3;
    int id;

    @Override
    public void run() {

        // TODO Auto-generated method stub
        System.out.println("som el thread"+ id); 
   }

    public plantillaJava(int id){
        this.id = id;
    }






    // La invocació de start() i join() requereix la gestió de l'excepció InterruptedException
    public static void main(String args[]) throws InterruptedException{
        Thread[] threads = new Thread[THREADS];
        int i;

        for (i = 0; i < THREADS; i++) {
            threads[i] = new Thread(new plantillaJava(i));
            threads[i].start();
        }
        try {
            for (i = 0; i < THREADS; i++) {
                threads[i].join();
            }

        }catch (InterruptedException ex){
            Logger.getLogger(plantillaJava.class.getName()).log(Level.SEVERE, null, ex);
        }
    }




}
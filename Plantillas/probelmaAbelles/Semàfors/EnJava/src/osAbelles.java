package probelmaAbelles.Semàfors.EnJava.src;
import java.util.logging.Level;
import java.util.logging.Logger;
import java.util.concurrent.Semaphore;

public class osAbelles {

    private final static int numAbelles = 10;
    protected static int pot;
    protected static int maxim = 10;
    static Semaphore mutex = new Semaphore(1);
    static Semaphore notPle = new Semaphore(maxim);
    static Semaphore avisarOs = new Semaphore(0);
    public final static int  repeticions = 20;

    public static void main(String [] args){
        System.out.println("Simulació d'os i Abelles");

        Thread os = new Thread(new os());
        os.start();
        Thread[] abelles = new Thread[numAbelles];

        for (int i = 0; i < abelles.length; i++) {
            abelles[i]= new Thread(new abella(i+1));
            abelles[i].start();
        }

        //join
        try {
            os.join();
            for (int i = 0; i < abelles.length; i++) {
            abelles[i].join();
        }
        System.out.println("simulació acabada");

        } catch (InterruptedException ex) {
            Logger.getLogger(osAbelles.class.getName()).log(Level.SEVERE, null, ex);

        }
        



    }



    


    
}

package probelmaAbelles.Semàfors.EnJava.src;
import java.util.Random;
import java.util.logging.Level;
import java.util.logging.Logger;

public class abella implements Runnable{
    private int id;

    public abella(int id){
        this.id=id;

    }

    @Override
    public void run() {
        try {
            Random rn = new Random();
            System.out.println("Hola, som l'abella "+ this.id);
           for (int i = 0; i < osAbelles.repeticions; i++) {
            //omplir el pot
               int espera = rn.nextInt(15);
               //permis per produir
               osAbelles.notPle.acquire();
               //consultat el tamany del pot
               osAbelles.mutex.acquire();
               osAbelles.pot= osAbelles.pot+1;
               System.out.println("L'abella "+ this.id+ " ha posat una porció de mel. Ara n'hi ha " + osAbelles.pot);
               if (osAbelles.pot == osAbelles.maxim ){
                   //despertar a l'os
                   System.out.println("Abella "+ this.id+ " El pot está ple. Despertaré a l'ós");
                   osAbelles.avisarOs.release();

               }
               osAbelles.mutex.release();
               Thread.sleep(espera);



           }

            // TODO Auto-generated method stub
        } catch (InterruptedException ex) {
            Logger.getLogger(abella.class.getName()).log(Level.SEVERE, null, ex);
        }
    }
    

    
    
}

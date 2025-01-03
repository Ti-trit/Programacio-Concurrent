package probelmaAbelles.Monitors.Implicits;

import java.util.Random;

public class monitorOsAbella {
    Random rn = new Random();
    int tamany;
    int pot = 0;
    int repeticions = 20;
    public monitorOsAbella(int tamany){
        this.tamany = tamany;
    }
    synchronized void posar_mel(int id){
            while (pot == tamany) {
                try{
                    wait();
                }catch (InterruptedException e){}
            }

            pot+=1;
            System.out.println("L'abella "+ id+ " ha posat una porció de mel al pot. El pot té "+ pot+ "porcions");
            //si el pot esta ple, notificar
            if (pot == tamany){
                System.out.println("Abella "+ id+ " : el pot está ple, Os!!");
                notify();
            }
        }

    synchronized  void consumir_mel() throws InterruptedException {
            System.out.println(" Os: yummie, a menjar!!");
            Thread.sleep(rn.nextInt(10));
            pot = 0;
            notifyAll();
    }
    synchronized void esperarPotPle (){
        while (pot <tamany){
            try{
                wait();

            }catch(InterruptedException ex){

            }
        }
    }

    synchronized void esperarPotBuid(){
        while(pot==tamany){
            try{
                wait();

            }catch(InterruptedException ex){

            }
        }
    }
}


package productor.consumidor;

import javax.swing.DefaultListModel;

public class productor extends Thread{
    
    public static int i = 1;
    
    public void run(){
        try{
            while(true){
                sleep(2000); // Definimos la rapidez
                formulario.buffer.acquire();
                producirElemento();
                formulario.sem.release();
            }
        }catch(Exception e){}
    }
    
    private void producirElemento(){
        elemento e = new elemento(i);
        formulario.modeloProductor.addElement(e);
        i++;
    }
}

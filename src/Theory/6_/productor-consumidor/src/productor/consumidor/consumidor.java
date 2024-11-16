package productor.consumidor;
import javax.swing.DefaultListModel;

public class consumidor extends Thread{
    
    private DefaultListModel modelo;
    
    public consumidor(DefaultListModel modelo){
        this.modelo = modelo;
    }
    
    public void run(){
        try{
            while(true){
                sleep(5000);
                formulario.sem.acquire();
                formulario.mutex.acquire();
                    consumirElemento();
                formulario.mutex.release();
                formulario.buffer.release();
            }
        }catch(Exception e){}
    }
    
    private void consumirElemento(){
        if (formulario.modeloProductor.isEmpty())
            System.out.println("Accediendo a Buffer Vacío");
        elemento e = (elemento) formulario.modeloProductor.remove(0);
        modelo.addElement(e);
    }
}

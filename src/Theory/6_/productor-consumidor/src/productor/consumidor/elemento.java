package productor.consumidor;
public class elemento {
    
    private int id;

    public elemento(int id) {
        this.id = id;
    }

    @Override
    public String toString() {
        return "Elemento " + id;
    }
    
}

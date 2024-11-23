import java.util.HashMap; 
import java.util.Objects;

// Prototype
abstract class Prototype{
    protected String prototypeKey;
    protected int prototypeValue;

    public Prototype(){}

    public Prototype(String prototypeKey, int prototypeValue){
        this.prototypeKey = prototypeKey;
        this.prototypeValue = prototypeValue;
    }

    // Java supports method overloading, therefore we can create a separate prototype constructor.
    // For other languages that do not support method overloading, copying the object's data can be done
    // in the clone method.
    public Prototype(Prototype prototype){
        if(prototype != null){
            this.prototypeKey = prototype.prototypeKey;
            this.prototypeValue = prototype.prototypeValue;
        }
    }

    public abstract Prototype clone();

    public void print(){
        System.out.println("Called from key: "+ this.prototypeKey+", value: "+this.prototypeValue);
    }

}

// Concrete Prototype
class ConcretePrototype1 extends Prototype{
    public int concretePrototypeValue1;

    public ConcretePrototype1(){}
    public ConcretePrototype1(String prototypeKey, int prototypeValue){
        super(prototypeKey, prototypeValue);
    }
    public ConcretePrototype1(ConcretePrototype1 concretePrototype1){
        super(concretePrototype1);
        if(concretePrototype1 != null){
            this.concretePrototypeValue1 = concretePrototype1.concretePrototypeValue1;
        }
    }

    @Override
    public Prototype clone(){
        return new ConcretePrototype1(this);
    }

    @Override
    public void print(){
        System.out.println("Called from key: "+ this.prototypeKey+", value: "+this.prototypeValue+", concretePrototypeValue1: "+this.concretePrototypeValue1);
    }
}
class ConcretePrototype2 extends Prototype{
    public int concretePrototypeValue2;

    public ConcretePrototype2(){}
    public ConcretePrototype2(String prototypeKey, int prototypeValue){
        super(prototypeKey, prototypeValue);
    }
    public ConcretePrototype2(ConcretePrototype2 concretePrototype2){
        super(concretePrototype2);
        if(concretePrototype2 != null){
            this.concretePrototypeValue2 = concretePrototype2.concretePrototypeValue2;
        }
    }

    @Override
    public Prototype clone(){
        return new ConcretePrototype2(this);
    }

    @Override
    public void print(){
        System.out.println("Called from key: "+ this.prototypeKey+", value: "+this.prototypeValue+", concretePrototypeValue2: "+this.concretePrototypeValue2);
    }
}

// Prototype Registry Keys
enum PrototypeKey {
    PROTOTYPE_1,
    PROTOTYPE_2
}

// Prototype Manager
class PrototypeManager{
    private HashMap<String,Prototype> prototypes;

    public PrototypeManager(){
        prototypes = new HashMap<String,Prototype>();
        ConcretePrototype1 concretePrototype1 = new ConcretePrototype1(PrototypeKey.PROTOTYPE_1.name(), 1);
        concretePrototype1.concretePrototypeValue1 = 11;
        prototypes.put(PrototypeKey.PROTOTYPE_1.name(), concretePrototype1);
    }

    public Prototype createPrototype(String prototypeKey){
        System.out.println("Creating prototype: "+prototypeKey);
        return prototypes.get(prototypeKey).clone();
    }

    // Complimentary operations
    public void listPrototypes(){
        System.out.println("Listing available prototypes:");
        for (String i : prototypes.keySet()) {
            System.out.println(i);
        }
    }
    public void putPrototype(String prototypeKey, Prototype prototype){
        System.out.println("Putting prototype <" + prototypeKey+ ">");
        prototypes.put(prototypeKey, prototype);
    }

    public void compareWithPrototype(String prototypeKey, Prototype prototype){

        System.out.println("Comparing with Prototype:");
        if(prototypes.get(prototypeKey) == prototype){
            System.out.println("- Same objects");
        }else{
            System.out.println("- Different objects");
        }
        if(Objects.equals(prototypes.get(prototypeKey), prototype)){
            System.out.println("- Identical");
        }else{
            System.out.println("- Not Identical");
        }
    }

    public void compareWithClone(Prototype prototype, Prototype prototype2){
        System.out.println("Comparing with Clone:");
        if(prototype == prototype2){
            System.out.println("- Same objects");
        }else{
            System.out.println("- Different objects");
        }
        if(Objects.equals(prototype, prototype2)){
            System.out.println("- Identical");
        }else{
            System.out.println("- Not Identical");
        }
    }
}

// Client code
class Client{
    public void clientCode(){
        PrototypeManager prototypeManager = new PrototypeManager();
        prototypeManager.listPrototypes();
        ConcretePrototype2 concretePrototype2 = new ConcretePrototype2(PrototypeKey.PROTOTYPE_2.name(), 1);
        concretePrototype2.concretePrototypeValue2 = 22;
        prototypeManager.putPrototype(PrototypeKey.PROTOTYPE_2.name(), concretePrototype2);
        prototypeManager.listPrototypes();

        System.out.println("======PROTOTYPE_1========");
        Prototype prototype1 = prototypeManager.createPrototype(PrototypeKey.PROTOTYPE_1.name());
        prototype1.print();
        prototypeManager.compareWithPrototype(PrototypeKey.PROTOTYPE_1.name(), prototype1);

        System.out.println("");
        Prototype prototype2 = prototypeManager.createPrototype(PrototypeKey.PROTOTYPE_1.name());
        prototype2.print();
        prototypeManager.compareWithPrototype(PrototypeKey.PROTOTYPE_1.name(), prototype2);

        System.out.println("");
        prototypeManager.compareWithClone(prototype1, prototype2);

        System.out.println("");
        System.out.println("=======PROTOTYPE_2=======");
        Prototype prototype3 = prototypeManager.createPrototype(PrototypeKey.PROTOTYPE_2.name());
        prototype1.print();
        prototypeManager.compareWithPrototype(PrototypeKey.PROTOTYPE_2.name(), prototype3);

        System.out.println("");
        Prototype prototype4 = prototypeManager.createPrototype(PrototypeKey.PROTOTYPE_2.name());
        prototype2.print();
        prototypeManager.compareWithPrototype(PrototypeKey.PROTOTYPE_2.name(), prototype4);

        System.out.println("");
        prototypeManager.compareWithClone(prototype3, prototype4);


        System.out.println("=======DIFF=======");
        prototypeManager.compareWithClone(prototype1, prototype3);
        prototypeManager.compareWithClone(prototype2, prototype4);

    }
}
public class PrototypeDemo {
    public static void main(String[] args) {
        Client client = new Client();
        client.clientCode();
    }
}
interface AbstractFactory {
    AbstractProductA createProductA();
    AbstractProductB createProductB();
}
class ConcreteFactory1 implements AbstractFactory {
    @Override
    public AbstractProductA createProductA(){
        return new ConcreteProductA1();
    }
    @Override
    public AbstractProductB createProductB(){
        return new ConcreteProductB1();
    }
}
class ConcreteFactory2 implements AbstractFactory {
    @Override
    public AbstractProductA createProductA(){
        return new ConcreteProductA2();
    }
    @Override
    public AbstractProductB createProductB(){
        return new ConcreteProductB2();
    }
}

interface AbstractProductA {
    public String usefulFunctionA();
}
class ConcreteProductA1 implements AbstractProductA {
    @Override
    public String usefulFunctionA(){
        return "The result of the product A1.";
    }
}
class ConcreteProductA2 implements AbstractProductA {
    @Override
    public String usefulFunctionA(){
        return "The result of the product A2.";
    }
}

interface AbstractProductB {
    public String usefulFunctionB();
    public String anotherUsefulFunctionB(AbstractProductA collaborator);
}
class ConcreteProductB1 implements AbstractProductB {
    @Override
    public String usefulFunctionB(){
        return "The result of the product B1.";
    }
    @Override
    public String anotherUsefulFunctionB(AbstractProductA collaborator){
        String result = collaborator.usefulFunctionA();
        return "The result of the B1 collaborating with the ("+result+")";
    }
}
class ConcreteProductB2 implements AbstractProductB {
    @Override
    public String usefulFunctionB(){
        return "The result of the product B2.";
    }
    @Override
    public String anotherUsefulFunctionB(AbstractProductA collaborator){
        String result = collaborator.usefulFunctionA();
        return "The result of the B2 collaborating with the ("+result+")";
    }
}

class Client {
    private AbstractProductA productA;
    private AbstractProductB productB;

    public void clientCode(AbstractFactory factory){
        productA = factory.createProductA();
        productB = factory.createProductB();
        System.out.println(productB.usefulFunctionB());
        System.out.println(productB.anotherUsefulFunctionB(productA));
    }
}

public class AbstractFactoryDemo {
    public static void main(String[] args) {
        Client client = new Client();
        System.out.println("Client: Testing client code with the first factory type:");
        client.clientCode(new ConcreteFactory1());
        System.out.println("Client: Testing the same client code with the second factory type:");
        client.clientCode(new ConcreteFactory2());
    }
}
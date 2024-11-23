import java.util.ArrayList;

/**
 * It makes sense to use the Builder pattern only when your products are quite
 * complex and require extensive configuration.
 *
 * Unlike in other creational patterns, different concrete builders can produce
 * unrelated products. In other words, results of various builders may not
 * always follow the same interface.
 */
class Product1 {
    ArrayList<String> parts = new ArrayList<String>();
    public void listParts(){
        System.out.println("Product parts: " + String.join(", ", parts));
        System.out.println("");
    }
}

/**
 * The Builder interface specifies methods for creating the different parts of
 * the Product objects.
 */
interface Builder {
    public void producePartA();
    public void producePartB();
    public void producePartC();
    public Product1 getProduct();
}

class ConcreteBuilder1 implements Builder {
    private Product1 product;
    
    /**
     * A fresh builder instance should contain a blank product object, which is
     * used in further assembly.
     */
    public ConcreteBuilder1(){
        this.reset();
    }

    public void reset(){
        this.product = new Product1();
    }

    public void producePartA(){
        this.product.parts.add("PartA1");
    }
    public void producePartB(){
        this.product.parts.add("PartB1");
    }
    public void producePartC(){
        this.product.parts.add("PartC1");
    }

    /**
     * Concrete Builders are supposed to provide their own methods for
     * retrieving results. That's because various types of builders may create
     * entirely different products that don't follow the same interface.
     * Therefore, such methods cannot be declared in the base Builder interface
     * (at least in a statically typed programming language). Note that PHP is a
     * dynamically typed language and this method CAN be in the base interface.
     * However, we won't declare it there for the sake of clarity.
     *
     * Usually, after returning the end result to the client, a builder instance
     * is expected to be ready to start producing another product. That's why
     * it's a usual practice to call the reset method at the end of the
     * `getProduct` method body. However, this behavior is not mandatory, and
     * you can make your builders wait for an explicit reset call from the
     * client code before disposing of the previous result.
     */

    public Product1 getProduct(){
        Product1 p = this.product;
        this.reset();
        return p;
    }
}

/**
 * The Director is only responsible for executing the building steps in a
 * particular sequence. It is helpful when producing products according to a
 * specific order or configuration. Strictly speaking, the Director class is
 * optional, since the client can control builders directly.
 */
class Director{
    private Builder builder;
    
    /**
     * The Director works with any builder instance that the client code passes
     * to it. This way, the client code may alter the final type of the newly
     * assembled product.
     */
    public void setBuilder(Builder builder){
        this.builder = builder;
    }
    /**
     * The Director can construct several product variations using the same
     * building steps.
     */
    public void buildMinimalViableProduct(){
        this.builder.producePartA();
    }

    public void buildFullFeaturedProduct(){
        this.builder.producePartA();
        this.builder.producePartB();
        this.builder.producePartC();
    }
}

class Client{
    public void clientCode(Director director){
        Builder builder = new ConcreteBuilder1();
        director.setBuilder(builder);

        System.out.println("Standard basic product:");
        director.buildMinimalViableProduct();
        builder.getProduct().listParts();

        System.out.println("Standard full featured product:");
        director.buildFullFeaturedProduct();
        builder.getProduct().listParts();
        
        // Remember, the Builder pattern can be used without a Director class.
        System.out.println("Custom product:");
        builder.producePartA();
        builder.producePartC();
        builder.getProduct().listParts();
    }
}

public class BuilderDemo {
    public static void main(String[] args) {
        Client client = new Client();
        client.clientCode(new Director());
    }
}
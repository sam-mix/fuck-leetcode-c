package onjava;

class Cleanser {
    private String s = "Cleanser";

    public void append(String a) {
        s += a;
    }

    public void d() {
        this.append(" d()");
    }

    public void a() {
        this.append(" a()");
    }

    public void s() {
        this.append(" s()");
    }

    @Override
    public String toString() {
        return s;
    }

    public static void main(String[] args) {
        var c = new Cleanser();
        c.a();
        c.d();
        c.s();
        System.out.println(c);
    }

}

public class ReuseDetergent extends Cleanser {
    @Override
    public void s() {
        append(" ReuseDetergent.s()");
        super.s();
    }

    public void b() {
        append(" b()");
    }

    public static void main(String[] args) {
        var x = new ReuseDetergent();
        x.a();
        x.b();
        x.d();
        x.s();
        System.out.println(x);
        System.out.println("Test base class:");
        Cleanser.main(args);

    }

}

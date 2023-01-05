package onjava;

public class SimpleEnumUse {
    public static void main(String[] args) {
        for (var v : Spiciness.values()) {
            System.out.println(v + "ordinal: " + v.ordinal());
        }
    }
}

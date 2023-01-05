package onjava;

class Plumbus {
}

public class TypeInference {
    void method() {
        // 显式类型
        String hello1 = "Hello";
        var hello = "Hello";
        Plumbus p1 = new Plumbus();
        var p2 = new Plumbus();
        System.out.println(hello1 + " " + hello + " " + p1 + " " + p2);
    }

    static void staticMethod() {
        var hello = "hello";
        var p1 = new Plumbus();
        System.out.println(hello + " " + p1);
    }
}

class NoInference {
    String f1 = "hekki";

    // var f2 = "xxxx";
    void method() {
        // var noInit;
        // var aNull = null;

    }

    // var inferReturnType() {
    // return "";
    // }
}

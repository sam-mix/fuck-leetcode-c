package onjava;

import java.util.Arrays;

class Processor {
    public String name() {
        return getClass().getSimpleName();
    }

    public Object process(Object input) {
        return input;
    }

}

class UpCase extends Processor {
    @Override
    public Object process(Object input) {
        return ((String) input).toUpperCase();
    }
}

class DownCase extends Processor {
    @Override
    public Object process(Object input) {
        return ((String) input).toLowerCase();
    }
}

class SplitCase extends Processor {
    @Override
    public Object process(Object input) {
        return Arrays.toString(((String) input).split(" "));
    }
}

public class IApp {
    public static void apply(Processor p, Object s) {
        System.out.println("Using Processor: " + p.name());
        System.out.println(p.process(s));
    }

    public static void main(String[] args) {
        String s = "We are family!";
        apply(new UpCase(), s);
        apply(new DownCase(), s);
        apply(new SplitCase(), s);
    }
}

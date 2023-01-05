package onjava;

class Game {
    Game(int i) {
        System.out.println("Game " + i);
    }
}

class BoardGame extends Game {

    BoardGame(int i) {
        super(i);
        System.out.println("BoardGame " + i);
    }
}

public class ReuseChess extends BoardGame {

    ReuseChess(int i) {
        super(i);
        System.out.println("ReuseChess " + i);
    }

    public static void main(String[] args) {
        new ReuseChess(0);
    }

}

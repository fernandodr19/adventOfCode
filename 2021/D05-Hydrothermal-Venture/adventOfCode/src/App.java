import java.util.ArrayList;
import java.util.Collections;
import java.util.Vector;

public class App {
    public static void main(String[] args) throws Exception {
        CoordinatesParser parser = new CoordinatesParser("../input.txt");

        Vector<Line> lines = parser.Parse();

        int biggestNumber = 0;
        for (Line line : lines) {
            ArrayList<Integer> tmp = new ArrayList<>();
            tmp.add(biggestNumber);
            tmp.add(line.getFirstCoordinate().x());
            tmp.add(line.getFirstCoordinate().y());
            tmp.add(line.getSecondCoordinate().x());
            tmp.add(line.getSecondCoordinate().y());
            biggestNumber = Collections.max(tmp);
        }

        Matrix matrix = new Matrix(biggestNumber+1);

        for (Line line : lines) {
            matrix.draw(line);
        }

        // System.out.println(matrix);
        
        System.out.println("Overlap count: " + matrix.getOverlapCount());
    }
}

import java.io.BufferedReader;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;
import java.util.Vector;

public class CoordinatesParser {
    private final String file;

    public CoordinatesParser(String file) {
        this.file = file;
    }

    public Vector<Line> Parse() throws FileNotFoundException, IOException {
        FileReader fr = new FileReader(this.file);
        BufferedReader br = new BufferedReader(fr);
        
        Vector<Line> lines = new Vector<>();

        String lineStr;
        while ((lineStr = br.readLine()) != null) {
           lineStr = lineStr.replaceAll(" ", "");
           lines.add(ParseLine(lineStr));
        }

        br.close();

        return lines;
    }

    private Line ParseLine(String lineStr) {
        lineStr = lineStr.replaceAll(" ", "");
        String[] coordinates = lineStr.split("->");
        if (coordinates.length != 2) {
            throw new RuntimeException("Invalid line string: " + lineStr);  
        }

        Coordinate c0 = ParseCoordinate(coordinates[0]);
        Coordinate c1 = ParseCoordinate(coordinates[1]);

        return new Line(c0, c1);
    }

    private Coordinate ParseCoordinate(String coordinateStr) {
        String[] points = coordinateStr.split(",");
        if (points.length != 2) {
            throw new RuntimeException("Invalid coordinate string: " + coordinateStr);  
        }

        int p0 = Integer.parseInt(points[0]);
        int p1 = Integer.parseInt(points[1]);
        return new Coordinate(p0, p1);
        
    }
}

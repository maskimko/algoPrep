package serializeBinTree;

import org.junit.jupiter.api.Test;

import java.util.Arrays;
import java.util.LinkedList;
import java.util.List;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void justSerialize(){
        assertEquals("6;4;3;;;5;;;8;;;",Solution.serialize(Solution.Node.buildTree(Arrays.asList("6 4 3 x x 5 x x 8 x x".split(" ")).iterator())));
    }


    @Test
    void deserialize() {
    String source = "6 4 3 x x 5 x x 8 x x";

    Solution.Node root = Solution.Node.buildTree(Arrays.asList("6 4 3 x x 5 x x 8 x x".split(" ")).iterator());
    String dump = Solution.serialize(root);
        Solution.Node unmarshalled = Solution.deserialize(dump);
        List<String> out = new LinkedList<>();
        Solution.Node.printTree(unmarshalled,out);
        assertEquals(source, String.join(" ", out));
    }
}
package sorting;

import java.util.Arrays;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @org.junit.jupiter.api.Test
    void sortList() {
        assertEquals(Arrays.asList(1,2,3),
                Solution.sortList(Arrays.asList(3,2,1)));
        assertEquals(Arrays.asList(1,2,3,4,5,6,7,8,9),
                Solution.sortList(Arrays.asList(9,2,5,1,3,4,7,6,8)));
    }
}
package uglyNumber3;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void nthUglyNumber() {
        assertEquals(14,Solution.nthUglyNumber(10,2,3,5));
    }

    @Test
    void nthUglyNumber2() {
        assertEquals(10,Solution.nthUglyNumber(8,2,3,4));
    }

    @Test
    void nthUglyNumber3() {
        assertEquals(88244,Solution.nthUglyNumber(1,88244,55236,44288));
    }

}
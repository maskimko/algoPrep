package uglyNumber3;

import org.junit.jupiter.api.Test;

import java.math.BigInteger;

import static org.junit.jupiter.api.Assertions.*;

class SolutionTest {

    @Test
    void nthUglyNumber() {
        assertEquals(14,Solution.nthUglyNumber(10,2,3,5));
    }

    @Test
    void nthUglyNumber2() {
        assertEquals(12,Solution.nthUglyNumber(8,2,3,4));
    }

    @Test
    void nthUglyNumber3() {
        assertEquals(88244,Solution.nthUglyNumber(88244,55236,44288,1 ));
    }


    @Test
    void nthUglyNumber5() {
        assertEquals(830665264,Solution.nthUglyNumber(20231774,4234,2117,6351));
    }


    @Test
    void howManyUglyOf14() {
        assertEquals(10, Solution.howManyUglyOf(
                BigInteger.valueOf(14),
                BigInteger.valueOf(2),
                BigInteger.valueOf(3),
                BigInteger.valueOf(5)
        ).intValue());
    }
    @Test
    void howManyUglyOf4() {
        assertEquals(2, Solution.howManyUglyOf(
                BigInteger.valueOf(4),
                BigInteger.valueOf(3),
                BigInteger.valueOf(4),
                BigInteger.valueOf(5)
                ).intValue());
    }
}
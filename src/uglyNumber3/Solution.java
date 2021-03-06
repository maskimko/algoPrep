package uglyNumber3;

import java.math.BigInteger;
import java.util.Scanner;

class Solution {

    private final static int modulo = 1000000007;

    static BigInteger howManyUglyOf(BigInteger k, BigInteger a, BigInteger b, BigInteger c) {
        BigInteger acInt = a.multiply(c).divide(a.gcd(c));
        BigInteger abInt = a.multiply(b).divide(a.gcd(b));
        BigInteger bcInt = b.multiply(c).divide(b.gcd(c));
        BigInteger abcInt = abInt.multiply(c).divide(abInt.gcd(c));
        return (k.divide(a)).add(k.divide(b)).add(k.divide(c))
                .subtract(k.divide(acInt)).subtract(k.divide(abInt)).subtract(k.divide(bcInt))
                .add(k.divide(abcInt));
    }

    public static int nthUglyNumber(int n, int a, int b, int c) {
        if (n == 0 || a == 0 || b == 0 || c == 0) return 0;
        BigInteger aInt = BigInteger.valueOf(a);
        BigInteger bInt = BigInteger.valueOf(b);
        BigInteger cInt = BigInteger.valueOf(c);
        BigInteger nInt = BigInteger.valueOf(n);
        BigInteger low = BigInteger.valueOf(n);
        BigInteger mod = BigInteger.valueOf(modulo);
        BigInteger high = BigInteger.valueOf(n).multiply(BigInteger.valueOf(a).min(BigInteger.valueOf(b).min(BigInteger.valueOf(c))));
        BigInteger mid;
        BigInteger k = null;
        while (k == null || !k.equals(BigInteger.valueOf(n))) {
            mid = low.add((high.subtract(low).divide(BigInteger.valueOf(2))));
            if (high.compareTo(low) < 0) {
                break;
            }
            k = howManyUglyOf(mid, aInt, bInt, cInt);
            if (k.compareTo(nInt) > 0) {
                high = mid.subtract(BigInteger.valueOf(1));
            } else if (k.compareTo(BigInteger.valueOf(n)) < 0) {
                low = mid.add(BigInteger.valueOf(1));
            } else {
                k = howManyUglyOf(mid.subtract(BigInteger.ONE), aInt, bInt, cInt);
                if (k.compareTo(nInt) < 0) {
                    return mid.mod(mod).intValue();
                } else {
                    high = high.subtract(BigInteger.TWO);
                }
            }
        }
        return -1;
    }

    public static void main(String[] args) {
        Scanner scanner = new Scanner(System.in);
        int n = Integer.parseInt(scanner.nextLine());
        int a = Integer.parseInt(scanner.nextLine());
        int b = Integer.parseInt(scanner.nextLine());
        int c = Integer.parseInt(scanner.nextLine());
        scanner.close();
        int res = nthUglyNumber(n, a, b, c);
        System.out.println(res);
    }
}
class Solution {
public:
    /*int sub(int amount, int divisor) {
        if(amount > divisor) return 0;
        int ret = 0;
        if(amount > divisor + divisor){
            ret = sub(amount - divisor, divisor);
            if(ret < INT_MAX) ret = ret + 1;
        }
        else if(amount > divisor + divisor + divisor){
            ret = sub(amount - divisor - divisor, divisor);
            if(ret < INT_MAX - 1) ret = ret + 2;
        }
        else{
            ret = sub(amount - divisor -divisor - divisor, divisor);
            if(ret < INT_MAX - 2) ret = ret + 3;
        }
        return ret;
    }*/

    int divide(int dividend, int divisor) {
        int pos = 1;
        if(dividend <= 0){
            if(divisor > 0){
                pos = 0;
                divisor = 0 - divisor;
            }
        }
        else{
            dividend = 0 - dividend;
            if(divisor < 0) pos = 0;
            else divisor = 0 - divisor;
        }

        if(divisor == -1) {
            if(dividend <= INT_MIN){
                return (pos == 1 ? INT_MAX : INT_MIN);
            }
            else {
                return (pos == 1 ? 0 - dividend : dividend);
            }
        }

        int ret = 0;
        map<int, int> m;
        m[0] = 0;
        for(int i = 1; i <= INT_MAX; ++i){
            m[i] = m[i-1] + divisor;
            if(dividend > m[i]){
                for(int j = i - 1 ; j >=1; --j){
                    if(dividend <= m[j]){
                        ret += j;
                        dividend -= m[j];
                    }
                }
                break;
            }
            else if(dividend - m[i] > m[i]){
                ret += i;
                dividend -= m[i];

                for(int j = i - 1 ; j >=1; --j){
                    if(dividend <= m[j]){
                        ret += j;
                        dividend -= m[j];
                    }
                }
                break;
            }
            else{
                ret += i + i;
                dividend = dividend - m[i] -m[i];
            }
        }
        if(pos == 1) return ret;
        else return 0 - ret;
        //if(pos == 1) return sub(dividend, divisor);
        //else return 0 - sub(dividend, divisor);
    }
};

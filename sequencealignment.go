package main

import (
    "math"
)
// CPP program to implement sequence alignment
// problem.
//#include <bits/stdc++.h>


// function to find out the minimum penalty
//void getMinimumPenalty(string x, string y, int pxy, int pgap)
func getMinimumPenalty(x string , y string , pxy int , pgap int ){
    //var i, j; // intialising variables

    m := x.length() // length of gene1
    n := y.length() // length of gene2

    // table for storing optimal substructure answers
    //int dp[n+m+1][n+m+1] = {0}
    var dp [n+m+1][n+m+1]int
    // intialising the table
    // for (i = 0; i <= (n+m); i++)
    // {
    for i := 0; i < (n+m); i++ {
        dp[i][0] = i * pgap
        dp[0][i] = i * pgap
    }

    // calcuting the minimum penalty
    // for (i = 1; i <= m; i++)
    // {
    for i := 0; i < m; i++ {
        // for (j = 1; j <= n; j++)
        // {
        for j := 0; j <= n; j++ {
            if (x[i - 1] == y[j - 1]){
                dp[i][j] = dp[i - 1][j - 1]
            } else {
                auxmin := int(math.Min(float64(dp[i - 1][j - 1] + pxy) , float64(dp[i - 1][j] + pgap)))
                dp[i][j] = int(math.Min(float64(auxmin) ,float64(dp[i][j - 1] + pgap)))
            }
        }
    }

    // Reconstructing the solution
    l := n + m // maximum possible length

    i := m
    j := n

    xpos := l
    ypos := l

    // Final answers for the respective strings
    var xans[l+1]int
    var yans[l+1]int

    // while ( !(i == 0 || j == 0))
    // {
    for ind :=0 ;  !(i == 0 || j == 0) ; ind++{
        if x[i - 1] == y[j - 1]{
            aux_xpos := xpos - 1
            aux_ypos := ypos - 1
            xans[aux_xpos] = int(x[i - 1])
            yans[aux_ypos] = int(y[j - 1])
            i--
            j--
        }else{
            if (dp[i - 1][j - 1] + pxy == dp[i][j]){
                aux_xpos := xpos - 1
                aux_ypos := ypos - 1
                xans[aux_xpos] = int(x[i - 1])
                yans[aux_ypos] = int(y[j - 1])
                i--
                j--
            }else{
                if (dp[i - 1][j] + pgap == dp[i][j]){
                    aux_xpos := xpos - 1
                    aux_ypos := ypos - 1
                    xans[aux_xpos] = int(x[i - 1])
                    yans[aux_ypos] = int('_')
                    i--
                                                                                                                               }else
                                                                 if (dp[i][j - 1] + pgap == dp[i][j])
                                                                     aux_xpos := xpos - 1
                                                                     aux_ypos := ypos - 1
                                                                     xans[aux_xpos] = int('_')
                                                                     yans[aux_ypos] = int(y[j - 1])
                                                                     j--}
                        }
    }
}

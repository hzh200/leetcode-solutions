func isMatch(s string, p string) bool {
  dp := make([][]bool, len(s) + 1)
  for i := 0; i < len(s) + 1; i++ {
      dp[i] = make([]bool, len(p) + 1)
  }
  dp[0][0] = true
  for i := 0; i < len(p); i++ {
      if p[i] != '*' {
          break
      }
      dp[0][i + 1] = true;
  }
  for i := 0; i < len(s); i++ {
      for j := 0; j < len(p); j++ {
          if p[j] == '*' {
              dp[i + 1][j + 1] = dp[i + 1][j] || dp[i][j + 1]
          } else {
              if p[j] == s[i] || p[j] == '?' {
                  dp[i + 1][j + 1] = dp[i][j]
              } else {
                  dp[i + 1][j + 1] = false
              }
          }
      }
  }
  return dp[len(s)][len(p)]
}

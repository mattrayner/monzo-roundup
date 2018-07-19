package monzoapi

type Methods interface {
  RefreshToken   (*RefreshTokenInput)       (*RefreshTokenOutput,   error)
  GetTransaction (*GetTransactionInput)     (*GetTransactionOutput, error)
  GetCoinJar     (*GetCoinJarInput)         (*GetCoinJarOutput,     error)
  Deposit        (*DepositInput)            (*DepositOutput,        error)
}

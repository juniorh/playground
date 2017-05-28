package main

import (
  "os"
  "fmt"
  "bufio"
  "encoding/pem"
  "crypto/x509"
  "crypto/rand"
  "crypto/rsa"
)

func main() {
  privKeyFile, err := os.Open("ssl/key.pem")
  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  privKeyFileStat,_ := privKeyFile.Stat()
  privKeyFileBytes := make([]byte, privKeyFileStat.Size())
  buffer := bufio.NewReader(privKeyFile)

  _, err = buffer.Read(privKeyFileBytes)
  if err != nil {
    fmt.Println("Error load file to memory: ", err)
  } 

  privKeyPem,rest := pem.Decode(privKeyFileBytes)
  if privKeyPem == nil {
    fmt.Println("Error decode PEM to rsa.privKey: ", rest)
  } 
  privKey,_ := x509.ParsePKCS1PrivateKey(privKeyPem.Bytes)
  pubKey := privKey.PublicKey

  random := rand.Reader
  msg := []byte("sapi")
  outEncPKCS,err := rsa.EncryptPKCS1v15(random, &pubKey, msg)
  if err != nil {
    fmt.Println("Error encrypt message: ", err)
  }
  outDecPKCS,err := rsa.DecryptPKCS1v15(random, privKey, outEncPKCS)
  if err != nil {
    fmt.Println("Error decrypt message: ", err)
  }

  Priv := privKey
  PCValues := Priv.Precomputed

  fmt.Println("Private Key: ",privKey)
  fmt.Println("Private Exponent :", Priv.D.String())
  fmt.Printf("Primes : %s %s \n", Priv.Primes[0].String(), Priv.Primes[1].String())
  fmt.Printf("Precomputed Values : Dp[%s] Dq[%s] Qinv[%s]\n", PCValues.Dp.String(), PCValues.Dq.String(), PCValues.Qinv.String())

  fmt.Println()
  fmt.Println("Public Key: ",pubKey)
  fmt.Printf("Hasil Enc(msg): 0x%X \n", outEncPKCS)
  fmt.Println("Hasil Dec(Enc(msg)): ", string(outDecPKCS))
}

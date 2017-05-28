package main

import(
    "fmt"
    "encoding/pem"
    "crypto/x509"
    "crypto/rsa"
    "crypto/rand"
    "io/ioutil"
)

func main(){
  fmt.Println()

  Priv,_ := rsa.GenerateKey(rand.Reader, 128)
  Pub := Priv.PublicKey
  //privKey := Priv
  pubKey := Pub
  PCValues := Priv.Precomputed
  PrivASN1 := x509.MarshalPKCS1PrivateKey(Priv)
  PubASN1, err := x509.MarshalPKIXPublicKey(&Priv.PublicKey)

  if err != nil {
    fmt.Println("Error")
  }

  PrivBytes := pem.EncodeToMemory(&pem.Block{
    Type:  "RSA PRIVATE KEY",
    Bytes: PrivASN1,
  })

  PubBytes := pem.EncodeToMemory(&pem.Block{
    Type:  "PUBLIC KEY",
    Bytes: PubASN1,
  })

  fmt.Println("Private Key:", Priv)
  fmt.Printf("Private Exponent : %v | 0x%X\n", Priv.D.String(), Priv.D)
  fmt.Printf("Primes : %s %s \n", Priv.Primes[0].String(), Priv.Primes[1].String())
  fmt.Printf("Precomputed Values : Dp[%s] Dq[%s] Qinv[%s]\n", PCValues.Dp.String(), PCValues.Dq.String(), PCValues.Qinv.String())

  fmt.Println()
  fmt.Println("Public key ", pubKey)
  fmt.Println("Public Exponent : ", pubKey.E)
  fmt.Printf("Modulus : %v | 0x%b | 0x%X \n", pubKey.N.String(), pubKey.N, pubKey.N)
  mod := pubKey.N
  fmt.Printf("Length Modulus: %d\n", mod.BitLen)
  fmt.Println("Save file to ssl/key.pem & ssl/key.pub")
  ioutil.WriteFile("ssl/key.pem", PrivBytes, 0644)
  ioutil.WriteFile("ssl/key.pub", PubBytes, 0644)
  fmt.Println()
}


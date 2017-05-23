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
  Priv,_ := rsa.GenerateKey(rand.Reader, 128)

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

  ioutil.WriteFile("/ssl/key.pem", PrivBytes, 0644)
  ioutil.WriteFile("/ssl/key.pub", PubBytes, 0644)
}


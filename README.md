#### Letöltés

$GOPATH/src mappában kell git clone-ozni

##### OP beállítás

A kód elején levő konstansba a saját useredhez generált access tokent kell megadni:  
const access_token string = "YOUR_OP_USER_ACCESS_TOKEN"

#### Build

1. lib letöltése

**go get github.com/urfave/cli**

vagy ha van glide

**glide get github.com/urfave/cli**

2. go build
3. have fun!

#### Használat

./op-cli time [--date=string] [--ma] [--tegnap] <work_package> <hours>  

pl.:  

./op-cli --tegnap 2408 1.5  

másfél órát logol tegnapra a 2408 package-re  

![alt text](https://user-images.githubusercontent.com/19253721/100717691-ad569b80-33ba-11eb-9dc1-29dc9c2fd9d7.png)



